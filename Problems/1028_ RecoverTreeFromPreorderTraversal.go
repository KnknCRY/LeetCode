package problems

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RecoverFromPreorder(traversal string) *TreeNode {
	// "1-2--3--4-5--6--7" // accept
	// "1-2--3---4-5--6---7" // wrong answer
	// "1-401--349---90--88" // wrong answer
	// "1-401--349--90---88" // wrong answer
	var root TreeNode
	if traversal == "" {
		return &root
	}

	val, _ := strconv.Atoi(string(traversal[0]))
	root.Val = val

	var depth, direction int // direction = 1 root往左，2 root往右
	for i := 1; i < len(traversal); i++ {
		if string(traversal[i]) == "-" {
			depth++
		} else {
			// 第一次遇到depth == 1，direction=1往左
			// 第二次遇到depth == 1，direction=2往右
			if depth == 1 {
				direction++
			}

			// 遇到非'-'的數字，繼續往後找他是不是多位數，轉型，深度計數器歸0
			j := i + 1
			for {
				if j == len(traversal) { // 走到底了
					// j--
					break
				}
				if string(traversal[j]) == "-" {
					break
				}
				j++
			}

			var _val int
			if i == j { // 走到底 且 個位數
				_val, _ = strconv.Atoi(string(traversal[i]))
			} else {
				_val, _ = strconv.Atoi(string(traversal[i:j]))
			}
			goToDepth(depth, _val, direction, &root, true)
			depth = 0
		}
	}

	return &root
}

func goToDepth(restDepth, val, direction int, node *TreeNode, isRoot bool) {
	// "-2" restDepth是1
	// 要從root開始往下走，先走左邊再走右邊
	// 遇到nil寫值，非nil return
	if restDepth == 0 {
		return
	}

	// 剩餘深度是1代表走到了，才要寫值
	// 如果走到了，卻兩邊都有node，就會回到前面的func然後往右走
	if restDepth == 1 {
		newNode := TreeNode{Val: val}
		if node.Left == nil {
			node.Left = &newNode
		} else if node.Right == nil {
			node.Right = &newNode
		}
	}

	restDepth -= 1
	if isRoot {
		// 如果是樹根，只能往direction的方向走
		if direction == 1 {
			goToDepth(restDepth, val, direction, node.Left, false)
		} else {
			goToDepth(restDepth, val, direction, node.Right, false)
		}
	} else {
		// 如果非樹根，要把左邊走完再走右邊
		if node.Left != nil {
			goToDepth(restDepth, val, direction, node.Left, false)
		}
		if node.Right != nil {
			goToDepth(restDepth, val, direction, node.Right, false)
		}
	}
}
