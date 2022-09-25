package problems

import "fmt"

func LinkedListCycle2(head *ListNode) *ListNode {
	nodeMap := map[*ListNode]int{}
	var currentNode *ListNode = head
	var idx int = 0
	for {
		if currentNode == nil {
			fmt.Println("沒有cycle")
			return nil
		} else {
			if _, exist := nodeMap[currentNode]; !exist {
				nodeMap[currentNode] = idx
				fmt.Println("找到新的點", currentNode, "值", currentNode.Val, "位置", idx)
			} else {
				fmt.Println("找到cycle, 當前位置", idx, "指向", currentNode.Next, "位置", nodeMap[currentNode])
				return nil
			}
		}
		idx++
		currentNode = currentNode.Next
	}
}
