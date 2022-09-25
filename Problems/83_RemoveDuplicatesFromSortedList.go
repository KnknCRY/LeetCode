package problems

func RemoveDuplicatesFromSortedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var resultHead, resultCurrent *ListNode
	existMap := map[int]interface{}{}

	currentNode := head
	for {
		if currentNode == nil {
			break
		} else {
			if _, exist := existMap[currentNode.Val]; !exist {
				if resultHead == nil {
					resultHead = &ListNode{Val: currentNode.Val}
					resultCurrent = resultHead
				} else {
					resultCurrent.Next = &ListNode{Val: currentNode.Val}
					resultCurrent = resultCurrent.Next
				}
				existMap[currentNode.Val] = nil
			}
		}
		currentNode = currentNode.Next
	}

	return resultHead
}
