#!/bin/bash

# 创建答题目录
# 用法: ./newq.sh 514

function NewSolutionDir() {
  num=$1

  local _dirname="l$num"
  mkdir $_dirname
  echo "package $_dirname" > $_dirname/solution.go
  echo "package $_dirname" > $_dirname/solution_test.go
}

NewSolutionDir $1