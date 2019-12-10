package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func getSubTree(currTreeNode *TreeNode, search int) *TreeNode {
  if currTreeNode == nil {
    return nil
  }

  if currTreeNode.Val == search {
    return currTreeNode
  }

  if search < currTreeNode.Val {
    return getSubTree(currTreeNode.Left, search)
  }

  if search > currTreeNode.Val {
    return getSubTree(currTreeNode.Right, search)
  }
}

func isEquals(root1 *TreeNode, root2 *TreeNode) bool {
  if root1 == nil && root2 == nil {
    return true
  }

  if root1 == nil {
    return false
  }

  if root2 == nil {
    return false
  }

  return root1.Val == root2.Val && isEquals(root1.Left, root2.Left) && isEquals(root1.Right, root2.Right)

}

func isSubTree(s *TreeNode, t *TreeNode) bool {

  if s == nil && t == nil {
    return false
  }

  if s == nil {
    return false
  }

  if t == nil {
    return false
  }

  subTree := getSubTree(s, t.Val)

  if subTree == nil {
    return false
  }

  return isEquals(subTree, t)

}

func main() {

}