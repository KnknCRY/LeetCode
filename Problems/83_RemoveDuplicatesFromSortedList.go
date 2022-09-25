package problems

func RemoveDuplicatesFromSortedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var resultHead, resultCurrent *ListNode
	existMap := map[int]interface{}{}

	resultHead = &ListNode{Val: head.Val}
	resultCurrent = resultHead
	currentNode := head.Next
	existMap[head.Val] = nil
	for {
		if currentNode == nil {
			break
		} else {
			// fmt.Println("現在跑到的點", currentNode)
			if _, exist := existMap[currentNode.Val]; !exist {
				// fmt.Println("新點", currentNode)
				// fmt.Println("結果指向", resultCurrent)
				existMap[currentNode.Val] = nil
				resultCurrent.Next = &ListNode{Val: currentNode.Val}
				resultCurrent = resultCurrent.Next
				// fmt.Println("結果指向", resultCurrent.Val)
			}
		}
		currentNode = currentNode.Next
	}

	return resultHead
}
