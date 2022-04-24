package algorithm

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 寻找树中最左下结点的值
// https://www.lintcode.com/problem/1197/
// 给定一棵二叉树，找到这棵树最后一行中最左边的值。
// 输入:{1,2,3,4,5,6,#,#,7}
// 输出:7
// 解释：
//         1
//        /  \
//      2     3
//    /  \    /
//  4     5 6
//   \
//    7

func FindBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	ans := root.Val
	for len(nodes) > 0 {
		tmp := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		if len(tmp) > 0 {
			ans = tmp[0].Val
		}
		nodes = tmp
	}
	return ans
}

func TestFindBottomLeftValue(t *testing.T) {
	ar := assert.New(t)
	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n4 := TreeNode{Val: 4}
	n5 := TreeNode{Val: 5}
	n6 := TreeNode{Val: 6}
	n7 := TreeNode{Val: 7}
	n1.Left = &n2
	n1.Right = &n3
	ar.Equal(2, FindBottomLeftValue(&n1))
	n2.Left = &n4
	n2.Right = &n5
	ar.Equal(4, FindBottomLeftValue(&n1))
	n4.Right = &n7
	n3.Left = &n6
	ar.Equal(7, FindBottomLeftValue(&n1))
}

// 终止进程
// https://www.lintcode.com/problem/872/

// 每个进程都有一个唯一的 PID(进程id) 和 PPID(父进程id)。每个进程只有一个父进程，但可能有一个或多个子进程，这就像一个树形结构。
// 并且只有一个进程的PPID是0，这意味着这个进程没有父进程。所有的pid都是不同的正整数。
// 我们使用两个整数列表来表示进程列表，其中第一个列表包含每个进程的PID，第二个列表包含对应的PPID。
// 现在给定这两个列表，以及一个你要终止(kill)的进程的ID，返回将在最后被终止的进程的PID列表。
// （当一个进程被终止时，它的所有子进程将被终止。返回的列表没有顺序要求）
//
// 给定的kill id一定是PID列表中的某个id
// 给定的PID列表中至少含有一个进程
//
// 输入: PID = [1, 3, 10, 5], PPID = [3, 0, 5, 3], killID = 5
// 输出: [5, 10]
// 解释: 终止进程5同样会终止进程10.
//      3
//    /   \
//   1     5
//        /
//       10
func KillProcess(pid []int, ppid []int, kill int) []int {
	if len(pid) != len(ppid) || len(pid) == 0 || kill < 1 {
		return nil
	}
	childMap := map[int][]int{}
	for i := range pid {
		childMap[ppid[i]] = append(childMap[ppid[i]], pid[i])
	}
	kills := []int{kill}
	var ans []int
	for len(kills) > 0 {
		var tmp []int
		for _, k := range kills {
			ans = append(ans, k)
			tmp = append(tmp, childMap[k]...)
		}
		kills = tmp
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func TestKillProcess(t *testing.T) {
	fmt.Println(KillProcess([]int{1, 3, 10, 5}, []int{3, 0, 5, 3}, 5))
}
