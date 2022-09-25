package problems

import "fmt"

func LinkedListCycle(head *ListNode) bool {
	nodeMap := map[*ListNode]interface{}{}
	var currentNode *ListNode = head
	for {
		if currentNode == nil {
			return false
		} else {
			if _, exist := nodeMap[currentNode]; !exist {
				nodeMap[currentNode] = nil
				fmt.Println("找到新的點", currentNode, "值", currentNode.Val)
			} else {
				return true
			}
		}
		currentNode = currentNode.Next
	}
}
