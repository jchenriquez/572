package main

import "regexp"

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isEqual(s *TreeNode, t *TreeNode) bool {
  if s == nil && t == nil {
    return true
  }

  if s == nil {
    return false
  }

  if t == nil {
    return false
  }

  if s.Val != t.Val {
    return false
  }

  return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
  if isEqual(s, t) {
    return true
  }

  if s != nil && s.Left != nil && s.Right != nil {
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
  }

  if s != nil && s.Left == nil {
    return isSubtree(s.Right, t)
  }

  if s != nil && s.Right == nil {
    return isSubtree(s.Left, t)
  }

  return false
}

func main() {

}