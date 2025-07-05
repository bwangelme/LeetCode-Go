#!/usr/bin/env python3
"""
LeetCode题目信息生成器
遍历目录并获取LeetCode题目信息，生成README表格
"""

import os
import re
import requests
import json
import time
from typing import List, Dict, Tuple

def extract_problem_numbers(directory_name: str) -> List[int]:
    """
    从目录名中提取题目编号
    """
    # 忽略lcr和algo目录
    if directory_name.startswith('lcr') or directory_name == 'algo':
        return []
    
    # 只处理l开头的目录
    if not directory_name.startswith('l'):
        return []
    
    # 移除l前缀
    name_without_l = directory_name[1:]
    
    # 特殊情况处理
    special_cases = {
        '3_v2': [3],
        '84l85': [84, 85],
        '141_l142': [141, 142],
    }
    
    if name_without_l in special_cases:
        return special_cases[name_without_l]
    
    # 处理包含下划线的情况（如l141_l142）
    if '_l' in name_without_l:
        parts = name_without_l.split('_l')
        try:
            return [int(parts[0]), int(parts[1])]
        except ValueError:
            return []
    
    # 处理包含字母的情况（如l84l85）
    if 'l' in name_without_l and name_without_l != name_without_l.replace('l', ''):
        # 尝试提取所有数字
        numbers = re.findall(r'\d+', name_without_l)
        try:
            return [int(num) for num in numbers]
        except ValueError:
            return []
    
    # 普通情况：直接提取数字
    try:
        # 移除可能的后缀（如_v2）
        clean_name = re.sub(r'_.*$', '', name_without_l)
        return [int(clean_name)]
    except ValueError:
        return []

def load_problems_from_json() -> Dict[int, Dict]:
    """
    从本地JSON文件加载题目信息
    """
    json_file = 'leetcode_cn_all_problems.json'
    problems_dict = {}
    
    try:
        with open(json_file, 'r', encoding='utf-8') as f:
            data = json.load(f)
            
        # 解析JSON数据结构
        if 'stat_status_pairs' in data:
            # 如果是LeetCode API格式
            for item in data['stat_status_pairs']:
                stat = item.get('stat', {})
                difficulty = item.get('difficulty', {})
                
                frontend_id = stat.get('frontend_question_id', '')
                if frontend_id and not frontend_id.startswith('LCP'):  # 跳过LCP题目
                    try:
                        problem_id = int(frontend_id)
                        difficulty_map = {1: 'Easy', 2: 'Medium', 3: 'Hard'}
                        problems_dict[problem_id] = {
                            'title': stat.get('question__title', f'题目{problem_id}'),
                            'difficulty': difficulty_map.get(difficulty.get('level', 2), 'Medium'),
                            'tags': []  # 这个API格式通常不包含标签
                        }
                    except (ValueError, TypeError):
                        continue
        else:
            # 如果是其他格式，尝试直接解析
            for key, value in data.items():
                try:
                    problem_id = int(key)
                    problems_dict[problem_id] = value
                except (ValueError, TypeError):
                    continue
                    
        print(f"从 {json_file} 加载了 {len(problems_dict)} 个题目信息")
        
    except FileNotFoundError:
        print(f"未找到文件 {json_file}")
        print("请按以下步骤获取题目数据：")
        print("1. 在浏览器中登录 https://leetcode.cn")
        print("2. 登录后访问 https://leetcode.cn/api/problems/all/")
        print("3. 将返回的JSON数据保存为 leetcode_cn_all_problems.json 文件")
        print("4. 将文件放在当前目录下，然后重新运行脚本")
        print("\n程序终止。")
        exit(1)
    except json.JSONDecodeError:
        print(f"文件 {json_file} 格式错误，请检查JSON格式是否正确")
        print("程序终止。")
        exit(1)
    except Exception as e:
        print(f"读取文件 {json_file} 时出错: {e}")
        print("程序终止。")
        exit(1)
    
    return problems_dict

def get_leetcode_problem_info(problem_id: int, json_problems: Dict[int, Dict]) -> Dict:
    """
    获取LeetCode题目信息，优先使用JSON文件中的数据
    """
    # 检查JSON文件中的数据
    if problem_id in json_problems:
        json_data = json_problems[problem_id]
        return {
            'id': problem_id,
            'title': json_data.get('title', f'题目{problem_id}'),
            'difficulty': json_data.get('difficulty', '未知'),
            'tags': json_data.get('tags', [])
        }
    
    # 如果没有找到，返回无信息
    return {
        'id': problem_id,
        'title': '无信息',
        'difficulty': '未知',
        'tags': []
    }


def get_all_problem_directories() -> List[str]:
    """
    获取所有l开头的目录
    """
    current_dir = os.getcwd()
    directories = []
    
    for item in os.listdir(current_dir):
        if os.path.isdir(item) and item.startswith('l') and not item.startswith('lcr'):
            directories.append(item)
    
    return sorted(directories, key=lambda x: int(re.findall(r'\d+', x)[0]) if re.findall(r'\d+', x) else 0)

def difficulty_to_chinese(difficulty: str) -> str:
    """
    将英文难度转换为中文
    """
    mapping = {
        'Easy': '简单',
        'Medium': '中等',
        'Hard': '困难'
    }
    return mapping.get(difficulty, difficulty)

def generate_readme_table(problems_info: List[Dict]) -> str:
    """
    生成README表格
    """
    table = """# LeetCode Solutions in Go

## 题目列表

| 题号 | 题目名称 | 难度 | 目录 |
|------|----------|------|------|
"""
    
    for info in problems_info:
        problem_id = info['id']
        title = info['title']
        difficulty = difficulty_to_chinese(info['difficulty'])
        directory = info['directory']
        
        table += f"| {problem_id} | {title} | {difficulty} | [`{directory}/`](./{directory}/) |\n"
    
    return table

def main():
    """
    主函数
    """
    print("开始扫描目录...")
    directories = get_all_problem_directories()
    print(f"找到 {len(directories)} 个目录")
    
    # 加载JSON文件中的题目信息
    json_problems = load_problems_from_json()
    
    all_problems = []
    
    for directory in directories:
        print(f"处理目录: {directory}")
        problem_numbers = extract_problem_numbers(directory)
        
        if not problem_numbers:
            print(f"  跳过目录 {directory}（无法提取题号）")
            continue
        
        for problem_id in problem_numbers:
            print(f"  获取题目 {problem_id} 的信息...")
            problem_info = get_leetcode_problem_info(problem_id, json_problems)
            problem_info['directory'] = directory
            all_problems.append(problem_info)
    
    # 按题号排序
    all_problems.sort(key=lambda x: x['id'])
    
    print("\n生成README表格...")
    readme_content = generate_readme_table(all_problems)
    
    # 写入README.md
    with open('README.md', 'w', encoding='utf-8') as f:
        f.write(readme_content)
    
    print(f"完成！共处理了 {len(all_problems)} 个题目")
    print("README.md 已更新")

if __name__ == "__main__":
    main()
