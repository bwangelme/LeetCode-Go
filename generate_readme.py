#!/usr/bin/env python3
"""
LeetCode题目信息生成器
遍历目录并获取LeetCode题目信息，生成README表格
"""

import os
import re
import json
import logging
import requests
import time
from pathlib import Path
from typing import List, Dict, Optional, Tuple
from dataclasses import dataclass
from enum import Enum

# 常量定义
class Constants:
    JSON_FILE = 'leetcode_cn_all_problems.json'
    README_FILE = 'README.md'
    LEETCODE_CN_API_URL = 'https://leetcode.cn/api/problems/all/'
    LEETCODE_CN_LOGIN_URL = 'https://leetcode.cn'
    LEETCODE_CN_GRAPHQL_URL = 'https://leetcode.cn/graphql/'
    LEETCODE_COM_GRAPHQL_URL = 'https://leetcode.com/graphql/'
    
    # 特殊目录名称映射
    SPECIAL_DIRECTORY_MAPPING = {
        '3_v2': [3],
        '84_l85': [84, 85],  # 修正目录名映射
        '141_l142': [141, 142],
    }
    
    # 难度映射
    DIFFICULTY_MAPPING = {
        1: 'Easy',
        2: 'Medium', 
        3: 'Hard'
    }
    
    CHINESE_DIFFICULTY_MAPPING = {
        'Easy': '简单',
        'Medium': '中等',
        'Hard': '困难'
    }

@dataclass
class ProblemInfo:
    """题目信息数据类"""
    id: int
    title: str
    difficulty: str
    directory: str
    tags: List[str] = None
    title_cn: str = None
    
    def __post_init__(self):
        if self.tags is None:
            self.tags = []
        if self.title_cn is None:
            self.title_cn = self.title

class LeetCodeDataError(Exception):
    """LeetCode数据相关异常"""
    pass

class LeetCodeAPIClient:
    """LeetCode API客户端"""
    
    def __init__(self, use_cn=True):
        self.use_cn = use_cn
        self.graphql_url = Constants.LEETCODE_CN_GRAPHQL_URL if use_cn else Constants.LEETCODE_COM_GRAPHQL_URL
        self.session = requests.Session()
        # TODO: 加入 cookie
        self.session.headers.update({
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        })
        self.logger = logging.getLogger(__name__)
        
    def get_problem_list(self, limit: int = 50, skip: int = 0) -> Dict:
        """获取题目列表"""
        query = """
        query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
          problemsetQuestionList: questionList(
            categorySlug: $categorySlug
            limit: $limit
            skip: $skip
            filters: $filters
          ) {
            total: totalNum
            questions: data {
              acRate
              difficulty
              freqBar
              frontendQuestionId: questionFrontendId
              isFavor
              paidOnly: isPaidOnly
              status
              title
              titleSlug
              titleCn
              topicTags {
                name
                nameTranslated
                id
                slug
              }
              hasSolution
              hasVideoSolution
            }
          }
        }
        """
        
        variables = {
            "categorySlug": "",
            "skip": skip,
            "limit": limit,
            "filters": {}
        }
        
        return self._make_graphql_request(query, variables)
    
    def get_problem_detail(self, title_slug: str) -> Dict:
        """获取题目详情"""
        query = """
        query questionData($titleSlug: String!) {
          question(titleSlug: $titleSlug) {
            questionId
            questionFrontendId
            title
            titleSlug
            translatedTitle
            difficulty
            topicTags {
              name
              nameTranslated
              slug
            }
            content
            translatedContent
            stats
            hints
            similarQuestions
          }
        }
        """
        
        variables = {
            "titleSlug": title_slug
        }
        
        return self._make_graphql_request(query, variables)
    
    def get_all_problems(self) -> List[Dict]:
        """获取所有题目信息"""
        self.logger.info("开始从LeetCode API获取题目信息...")
        all_problems = []
        skip = 0
        limit = 100
        
        while True:
            try:
                self.logger.info(f"正在获取第 {skip + 1} 到 {skip + limit} 个题目...")
                
                response = self.get_problem_list(limit=limit, skip=skip)
                
                if not response or 'data' not in response:
                    self.logger.warning("API响应异常，停止获取")
                    break
                
                problem_list = response['data']['problemsetQuestionList']
                questions = problem_list['questions']
                
                if not questions:
                    self.logger.info("没有更多题目，获取完成")
                    break
                
                all_problems.extend(questions)
                skip += limit
                
                # 添加延时避免请求过快
                time.sleep(0.5)
                
                # 如果获取的题目数量小于限制，说明已经获取完毕
                if len(questions) < limit:
                    break
                    
            except Exception as e:
                self.logger.error(f"获取题目信息时出错: {e}")
                break
        
        self.logger.info(f"共获取到 {len(all_problems)} 个题目信息")
        return all_problems
    
    def _make_graphql_request(self, query: str, variables: Dict) -> Dict:
        """发送GraphQL请求"""
        payload = {
            'query': query,
            'variables': variables
        }
        
        try:
            response = self.session.post(
                self.graphql_url,
                json=payload,
                timeout=10
            )
            response.raise_for_status()
            
            result = response.json()
            
            if 'errors' in result:
                raise LeetCodeDataError(f"GraphQL错误: {result['errors']}")
            
            return result
            
        except requests.exceptions.RequestException as e:
            raise LeetCodeDataError(f"网络请求错误: {e}")
        except json.JSONDecodeError as e:
            raise LeetCodeDataError(f"JSON解析错误: {e}")
    
    def batch_get_problem_details(self, problem_ids: List[int]) -> Dict[int, Dict]:
        """批量获取题目详情"""
        self.logger.info(f"开始批量获取 {len(problem_ids)} 个题目的详情...")
        
        # 首先获取所有题目的基本信息
        all_problems = self.get_all_problems()
        
        # 建立ID到题目信息的映射
        problem_map = {}
        for problem in all_problems:
            try:
                problem_id = int(problem['frontendQuestionId'])
                problem_map[problem_id] = {
                    'id': problem_id,
                    'title': problem['title'],
                    'title_cn': problem.get('titleCn', problem['title']),
                    'difficulty': problem['difficulty'],
                    'tags': [tag['nameTranslated'] if tag.get('nameTranslated') else tag['name'] 
                            for tag in problem.get('topicTags', [])],
                    'slug': problem['titleSlug']
                }
            except (ValueError, KeyError) as e:
                self.logger.warning(f"解析题目信息时出错: {e}")
                continue
        
        # 返回请求的题目信息
        result = {}
        for problem_id in problem_ids:
            if problem_id in problem_map:
                result[problem_id] = problem_map[problem_id]
            else:
                self.logger.warning(f"未找到题目 {problem_id} 的信息")
        
        return result

class DirectoryScanner:
    """目录扫描器"""
    
    def __init__(self, base_path: Path = None):
        self.base_path = base_path or Path.cwd()
        self.number_pattern = re.compile(r'\d+')
        self.version_pattern = re.compile(r'_.*$')
        
    def get_problem_directories(self) -> List[str]:
        """获取所有问题目录"""
        directories = []
        
        for item in self.base_path.iterdir():
            if (item.is_dir() and 
                item.name.startswith('l') and 
                not item.name.startswith('lcr') and
                item.name != 'lt'):  # 排除lt目录
                directories.append(item.name)
        
        return sorted(directories, key=self._get_sort_key)
    
    def _get_sort_key(self, directory_name: str) -> int:
        """获取排序键"""
        numbers = self.number_pattern.findall(directory_name)
        return int(numbers[0]) if numbers else 0
    
    def extract_problem_numbers(self, directory_name: str) -> List[int]:
        """从目录名中提取题目编号"""
        # 忽略特定目录
        if directory_name.startswith('lcr') or directory_name == 'algo':
            return []
        
        # 只处理l开头的目录
        if not directory_name.startswith('l'):
            return []
        
        name_without_l = directory_name[1:]
        
        # 处理特殊情况
        if name_without_l in Constants.SPECIAL_DIRECTORY_MAPPING:
            return Constants.SPECIAL_DIRECTORY_MAPPING[name_without_l]
        
        # 处理包含下划线的情况
        if '_l' in name_without_l:
            return self._extract_from_underscore_format(name_without_l)
        
        # 处理包含字母的情况
        if 'l' in name_without_l and name_without_l != name_without_l.replace('l', ''):
            return self._extract_all_numbers(name_without_l)
        
        # 普通情况：直接提取数字
        return self._extract_simple_number(name_without_l)
    
    def _extract_from_underscore_format(self, name: str) -> List[int]:
        """从下划线格式提取数字"""
        parts = name.split('_l')
        try:
            return [int(parts[0]), int(parts[1])]
        except (ValueError, IndexError):
            return []
    
    def _extract_all_numbers(self, name: str) -> List[int]:
        """提取所有数字"""
        numbers = self.number_pattern.findall(name)
        try:
            return [int(num) for num in numbers]
        except ValueError:
            return []
    
    def _extract_simple_number(self, name: str) -> List[int]:
        """提取简单数字"""
        try:
            clean_name = self.version_pattern.sub('', name)
            return [int(clean_name)]
        except ValueError:
            return []

class ProblemDataLoader:
    """题目数据加载器"""
    
    def __init__(self, json_file: str = Constants.JSON_FILE):
        self.json_file = json_file
        self.logger = logging.getLogger(__name__)
        
    def load_problems_from_json(self) -> Dict[int, Dict]:
        """从JSON文件加载题目信息"""
        try:
            with open(self.json_file, 'r', encoding='utf-8') as f:
                data = json.load(f)
            
            problems_dict = self._parse_json_data(data)
            self.logger.info(f"从 {self.json_file} 加载了 {len(problems_dict)} 个题目信息")
            return problems_dict
            
        except FileNotFoundError:
            self.logger.warning(f"未找到文件 {self.json_file}，将使用API获取数据")
            return {}
        except json.JSONDecodeError:
            raise LeetCodeDataError(f"文件 {self.json_file} 格式错误，请检查JSON格式是否正确")
        except Exception as e:
            raise LeetCodeDataError(f"读取文件 {self.json_file} 时出错: {e}")
    
    def _parse_json_data(self, data: Dict) -> Dict[int, Dict]:
        """解析JSON数据"""
        problems_dict = {}
        
        if 'stat_status_pairs' in data:
            problems_dict = self._parse_leetcode_api_format(data)
        else:
            problems_dict = self._parse_direct_format(data)
        
        return problems_dict
    
    def _parse_leetcode_api_format(self, data: Dict) -> Dict[int, Dict]:
        """解析LeetCode API格式"""
        problems_dict = {}
        
        for item in data['stat_status_pairs']:
            stat = item.get('stat', {})
            difficulty = item.get('difficulty', {})
            
            frontend_id = stat.get('frontend_question_id', '')
            if frontend_id and not frontend_id.startswith('LCP'):
                try:
                    problem_id = int(frontend_id)
                    problems_dict[problem_id] = {
                        'title': stat.get('question__title', f'题目{problem_id}'),
                        'difficulty': Constants.DIFFICULTY_MAPPING.get(
                            difficulty.get('level', 2), 'Medium'
                        ),
                        'tags': []
                    }
                except (ValueError, TypeError):
                    continue
        
        return problems_dict
    
    def _parse_direct_format(self, data: Dict) -> Dict[int, Dict]:
        """解析直接格式"""
        problems_dict = {}
        
        for key, value in data.items():
            try:
                problem_id = int(key)
                problems_dict[problem_id] = value
            except (ValueError, TypeError):
                continue
        
        return problems_dict

class ReadmeGenerator:
    """README生成器"""
    
    @staticmethod
    def generate_table(problems_info: List[ProblemInfo]) -> str:
        """生成README表格"""
        lines = [
            "# LeetCode Solutions in Go",
            "",
            "## 题目列表",
            "",
            "| 题号 | 题目名称 | 中文名称 | 难度 | 标签 | 目录 |",
            "|------|----------|----------|------|------|------|"
        ]
        
        for info in problems_info:
            difficulty = Constants.CHINESE_DIFFICULTY_MAPPING.get(
                info.difficulty, info.difficulty
            )
            tags_str = ", ".join(info.tags[:3]) if info.tags else "无"  # 最多显示3个标签
            if len(info.tags) > 3:
                tags_str += "..."
            
            line = f"| {info.id} | {info.title} | {info.title_cn} | {difficulty} | {tags_str} | [`{info.directory}/`](./{info.directory}/) |"
            lines.append(line)
        
        return "\n".join(lines) + "\n"
    
    @staticmethod
    def write_readme(content: str, output_file: str = Constants.README_FILE):
        """写入README文件"""
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(content)

class LeetCodeReadmeGenerator:
    """LeetCode README生成器主类"""
    
    def __init__(self, use_api: bool = True):
        self.scanner = DirectoryScanner()
        self.loader = ProblemDataLoader()
        self.generator = ReadmeGenerator()
        self.use_api = use_api
        self.api_client = LeetCodeAPIClient(use_cn=True) if use_api else None
        self.logger = self._setup_logger()
    
    def _setup_logger(self) -> logging.Logger:
        """设置日志记录器"""
        logger = logging.getLogger(__name__)
        logger.setLevel(logging.INFO)
        
        if not logger.handlers:
            handler = logging.StreamHandler()
            formatter = logging.Formatter(
                '%(asctime)s - %(levelname)s - %(message)s'
            )
            handler.setFormatter(formatter)
            logger.addHandler(handler)
        
        return logger
    
    def generate(self):
        """生成README文件"""
        try:
            self.logger.info("开始扫描目录...")
            directories = self.scanner.get_problem_directories()
            # for test
            # directories = directories[:2]
            self.logger.info(f"找到 {len(directories)} 个目录")
            
            # 处理每个目录
            all_problems = self._process_directories(directories)
            
            # 生成并写入README
            readme_content = self.generator.generate_table(all_problems)
            self.generator.write_readme(readme_content)
            
            self.logger.info(f"完成！共处理了 {len(all_problems)} 个题目")
            self.logger.info("README.md 已更新")
            
        except LeetCodeDataError as e:
            self.logger.error(f"数据错误: {e}")
            raise
        except Exception as e:
            self.logger.error(f"生成过程中发生错误: {e}")
            raise
    
    def _process_directories(self, directories: List[str]) -> List[ProblemInfo]:
        """处理目录列表"""
        # 首先收集所有需要查询的题目ID
        all_problem_ids = set()
        directory_to_ids = {}
        
        for directory in directories:
            problem_numbers = self.scanner.extract_problem_numbers(directory)
            if problem_numbers:
                directory_to_ids[directory] = problem_numbers
                all_problem_ids.update(problem_numbers)
        
        # 获取题目信息
        if self.use_api and self.api_client:
            self.logger.info("使用API获取题目信息...")
            api_problems = self.api_client.batch_get_problem_details(list(all_problem_ids))
        else:
            api_problems = {}
        
        # 处理每个目录
        all_problems = []
        self.logger.info("使用本地JSON文件获取题目信息...")
        json_problems = self.loader.load_problems_from_json()
        for directory in directories:
            if directory not in directory_to_ids:
                continue
                
            problem_numbers = directory_to_ids[directory]
            
            for problem_id in problem_numbers:
                self.logger.info(f"  获取题目 {problem_id} 的信息...")
                
                if self.use_api and problem_id in api_problems:
                    # 使用API获取的数据
                    data = api_problems[problem_id]
                    problem_info = ProblemInfo(
                        id=problem_id,
                        title=data.get('title', f'题目{problem_id}'),
                        title_cn=data.get('title_cn', data.get('title', f'题目{problem_id}')),
                        difficulty=data.get('difficulty', '未知'),
                        directory=directory,
                        tags=data.get('tags', [])
                    )
                else:
                    # 使用JSON文件或默认数据
                    problem_info = self._get_problem_info_from_json(problem_id, directory, json_problems)
                
                all_problems.append(problem_info)
        
        # 按题号排序
        all_problems.sort(key=lambda x: x.id)
        return all_problems
    
    def _get_problem_info_from_json(self, problem_id: int, directory: str, json_problems: Dict[int, Dict]) -> ProblemInfo:
        """从JSON数据获取题目信息"""
        if problem_id in json_problems:
            data = json_problems[problem_id]
            return ProblemInfo(
                id=problem_id,
                title=data.get('title', f'题目{problem_id}'),
                title_cn=data.get('title_cn', data.get('title', f'题目{problem_id}')),
                difficulty=data.get('difficulty', '未知'),
                directory=directory,
                tags=data.get('tags', [])
            )
        
        return ProblemInfo(
            id=problem_id,
            title='无信息',
            title_cn='无信息',
            difficulty='未知',
            directory=directory
        )

def main():
    """主函数"""
    try:
        # 默认使用API模式
        generator = LeetCodeReadmeGenerator(use_api=True)
        generator.generate()
    except Exception as e:
        print(f"程序执行失败: {e}")
        return 1
    
    return 0

if __name__ == "__main__":
    exit(main())
