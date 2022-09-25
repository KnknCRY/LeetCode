package problems

import (
	"fmt"
	"sort"
)

func RemoveDuplicatesFromSortedList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var resultHead, resultCurrent *ListNode
	existMap := map[int]int{}

	currentNode := head
	for {
		if currentNode == nil {
			break
		} else {
			existMap[currentNode.Val]++
		}
		currentNode = currentNode.Next
	}

	fmt.Println("map", existMap)

	var appearOnceKeys []int
	for k, v := range existMap {
		if v == 1 {
			appearOnceKeys = append(appearOnceKeys, k)
		}
	}

	fmt.Println("array", appearOnceKeys)

	sort.Ints(appearOnceKeys)

	for _, item := range appearOnceKeys {
		if resultHead == nil {
			resultHead = &ListNode{Val: item}
			resultCurrent = resultHead
		} else {
			resultCurrent.Next = &ListNode{Val: item}
			resultCurrent = resultCurrent.Next
		}
	}

	return resultHead
}
