package main

import "fmt"

type Node4 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node4
	TopRight    *Node4
	BottomLeft  *Node4
	BottomRight *Node4
}

type Node struct {
	Val       int
	Left      *Node
	Right     *Node
	Next      *Node
	Neighbors []*Node
	Children  []*Node
}

var (
	//    5
	//   4  6
	//     3  7
	Tree1     = &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}
	ListNode1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (s *TreeNode) Print() {
	var dfs func(node *TreeNode)
	var values []int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		values = append(values, node.Val)
		dfs(node.Right)
	}
	dfs(s)
	fmt.Println(values)
}

var DemoList = &ListNode{
	Val: 4,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
				},
			},
		},
	},
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) Print() {
	arr := make([]int, 0, 10)
	for cur := s; cur != nil; cur = cur.Next {
		arr = append(arr, cur.Val)
	}
	fmt.Println(arr)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func init2DIntSlice(row, col int) [][]int {
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	return dp
}

// 初始化二维数组
func init2DStringSlice(row, col int) [][]string {
	dp := make([][]string, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]string, col)
	}
	return dp
}

func init2DIntSliceWithDefault(row, col, def int) [][]int {
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = def
		}
	}
	return dp
}

func minIntSlice(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
