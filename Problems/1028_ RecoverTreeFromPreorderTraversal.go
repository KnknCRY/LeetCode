package problems

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// "1-2--3--4-5--6--7" // accept
// "1-2--3---4-5--6---7" // wrong answer
// "1-401--349---90--88" // wrong answer
// "1-401--349--90---88" // wrong answer 要知道這個node他爸是誰
// 打掉重練：
// 如果下一個點在下一階，現在跑到的點 會是下一個點的爸爸
// 如果下一個點在同一階，現在跑到的點 會是下一個點的平輩，並共享同個爸爸
// 如果下一個點在上一階，現在跑到的點 不會是下一個點的爸爸，如果是上n階 就要往前找爸爸
// 用一個map[階數]TreeNode，然後用stack存起來，如果下一個點是上n階，就要pop stack直到stack中的點可以當下一個點的爸爸
// 後來發現，stack只要是TreeNode即可，因為"1-401--349---90--88"

func RecoverFromPreorder(traversal string) *TreeNode {
	var root TreeNode
	if traversal == "" {
		return &root
	}

	if len(traversal) == 1 {
		val, _ := strconv.Atoi(string(traversal[0]))
		root.Val = val
		return &root
	}

	var parentStack []*TreeNode
	var currentDepth, nextDepth int

	// val, _ := strconv.Atoi(string(traversal[0]))
	// root.Val = val

	// parentStack = append(parentStack, map[int]*TreeNode{0: &root})

	for i := 0; i < len(traversal); i++ {
		//fmt.Println("!!!!!!!!!!!!!!!!!!!!", i)
		if string(traversal[i]) == "-" {
			currentDepth++
		} else { // 如果遇到非"-"，要先看i往後有幾位數
			j := i + 1
			for { // 無限迴圈往i的後面找，直到再次碰到"-"
				if j == len(traversal) { // 走到底了
					break
				}
				if string(traversal[j]) == "-" { // 再次碰到"-"代表該停了
					break
				}
				j++
			}

			// 切出i~j並轉成數字
			var _val int
			// if i == j { // 走到底 且 個位數
			// 	_val, _ = strconv.Atoi(string(traversal[i]))
			// } else {
			_val, _ = strconv.Atoi(string(traversal[i:j]))
			// }
			//fmt.Println("要切出的數字", _val)

			i = j - 1

			// 如果此時i尚未走完，才要判斷下一個點是幾階
			node := TreeNode{Val: _val}
			if i != len(traversal)-1 {
				for k := i + 1; ; k++ {
					// //fmt.Println("i,j,k", i, j, k)
					//fmt.Println("字串", string(traversal[k]))
					if string(traversal[k]) == "-" {
						nextDepth++
					} else {
						break
					}
				}
				//fmt.Println("這一個點的階數", currentDepth)
				//fmt.Println("下一個點的階數", nextDepth)

				if nextDepth > currentDepth { // 如果階數 下一點 大於 此點，必定只會大1階
					// 此點是下一點的爸爸
					// 第一次進入這裡，必為currentDepth=0，root
					if currentDepth == 0 {
						root = node
						parentStack = append(parentStack, &root)
						//fmt.Println(root.Val, "是樹根")
					} else {
						stackLastElement := parentStack[len(parentStack)-1]
						if stackLastElement.Left == nil {
							stackLastElement.Left = &node
							//fmt.Println(node.Val, "放到", stackLastElement.Val, "左邊")
						} else if stackLastElement.Right == nil {
							stackLastElement.Right = &node
							//fmt.Println(node.Val, "放到", stackLastElement.Val, "右邊")
						}
						parentStack = append(parentStack, &node)
					}
				} else if nextDepth < currentDepth { // 如果階數 下一點 大於 此點，可能會大n階
					// 把stack pop到只剩下nextDepth個
					stackLastElement := parentStack[len(parentStack)-1]
					if stackLastElement.Left == nil {
						stackLastElement.Left = &node
						//fmt.Println(node.Val, "放到", stackLastElement.Val, "左邊")
					} else if stackLastElement.Right == nil {
						stackLastElement.Right = &node
						//fmt.Println(node.Val, "放到", stackLastElement.Val, "右邊")
					}

					//fmt.Println("stack準備回頭 要剩下", nextDepth, "個")
					for len(parentStack) > nextDepth {
						parentStack = parentStack[:len(parentStack)-1]
					}
					//fmt.Println("stack的狀態", parentStack)
				} else { // 如果階數 下一點 等於 此點，此點不能是爸爸
					stackLastElement := parentStack[len(parentStack)-1]
					if stackLastElement.Left == nil {
						stackLastElement.Left = &node
						//fmt.Println(node.Val, "放到", stackLastElement.Val, "左邊")
					} else if stackLastElement.Right == nil {
						stackLastElement.Right = &node
						//fmt.Println(node.Val, "放到", stackLastElement.Val, "右邊")
					}
				}
			}

			// j是下一個點的數字的末位置
			if i == len(traversal)-1 {
				stackLastElement := parentStack[len(parentStack)-1]
				if stackLastElement.Left == nil {
					stackLastElement.Left = &node
					//fmt.Println(node.Val, "放到", stackLastElement.Val, "左邊")
				} else if stackLastElement.Right == nil {
					stackLastElement.Right = &node
					//fmt.Println(node.Val, "放到", stackLastElement.Val, "右邊")
				}
			}

			currentDepth = 0
			nextDepth = 0

			//fmt.Println("stack的狀態", parentStack)
		}
	}

	return &root
}

// func RecoverFromPreorder(traversal string) *TreeNode {
// 	var root TreeNode
// 	if traversal == "" {
// 		return &root
// 	}

// 	val, _ := strconv.Atoi(string(traversal[0]))
// 	root.Val = val

// 	var depth, direction int // direction = 1 root往左，2 root往右
// 	for i := 1; i < len(traversal); i++ {
// 		if string(traversal[i]) == "-" {
// 			depth++
// 		} else {
// 			// 第一次遇到depth == 1，direction=1往左
// 			// 第二次遇到depth == 1，direction=2往右
// 			if depth == 1 {
// 				direction++
// 			}

// 			// 遇到非'-'的數字，繼續往後找他是不是多位數，轉型，深度計數器歸0
// 			j := i + 1
// 			for {
// 				if j == len(traversal) { // 走到底了
// 					// j--
// 					break
// 				}
// 				if string(traversal[j]) == "-" {
// 					break
// 				}
// 				j++
// 			}

// 			var _val int
// 			if i == j { // 走到底 且 個位數
// 				_val, _ = strconv.Atoi(string(traversal[i]))
// 			} else {
// 				_val, _ = strconv.Atoi(string(traversal[i:j]))
// 			}
// 			goToDepth(depth, _val, direction, &root, true)
// 			depth = 0
// 		}
// 	}

// 	return &root
// }

// func goToDepth(restDepth, val, direction int, node *TreeNode, isRoot bool) {
// 	// "-2" restDepth是1
// 	// 要從root開始往下走，先走左邊再走右邊
// 	// 遇到nil寫值，非nil return
// 	if restDepth == 0 {
// 		return
// 	}

// 	// 剩餘深度是1代表走到了，才要寫值
// 	// 如果走到了，卻兩邊都有node，就會回到前面的func然後往右走
// 	if restDepth == 1 {
// 		newNode := TreeNode{Val: val}
// 		if node.Left == nil {
// 			node.Left = &newNode
// 		} else if node.Right == nil {
// 			node.Right = &newNode
// 		}
// 	}

// 	restDepth -= 1
// 	if isRoot {
// 		// 如果是樹根，只能往direction的方向走
// 		if direction == 1 {
// 			goToDepth(restDepth, val, direction, node.Left, false)
// 		} else {
// 			goToDepth(restDepth, val, direction, node.Right, false)
// 		}
// 	} else {
// 		// 如果非樹根，要把左邊走完再走右邊
// 		if node.Left != nil {
// 			goToDepth(restDepth, val, direction, node.Left, false)
// 		}
// 		if node.Right != nil {
// 			goToDepth(restDepth, val, direction, node.Right, false)
// 		}
// 	}
// }
