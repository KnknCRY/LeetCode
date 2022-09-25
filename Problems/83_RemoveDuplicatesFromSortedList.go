package problems

func RemoveDuplicatesFromSortedList(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return head

	// ------------------------以下是第一次寫的大便------------------------
	// if head == nil {
	// 	return nil
	// }
	// var resultHead, resultCurrent *ListNode
	// existMap := map[int]interface{}{}

	// currentNode := head
	// for {
	// 	if currentNode == nil {
	// 		break
	// 	} else {
	// 		if _, exist := existMap[currentNode.Val]; !exist {
	// 			if resultHead == nil {
	// 				resultHead = &ListNode{Val: currentNode.Val}
	// 				resultCurrent = resultHead
	// 			} else {
	// 				resultCurrent.Next = &ListNode{Val: currentNode.Val}
	// 				resultCurrent = resultCurrent.Next
	// 			}
	// 			existMap[currentNode.Val] = nil
	// 		}
	// 	}
	// 	currentNode = currentNode.Next
	// }

	// return resultHead
}
