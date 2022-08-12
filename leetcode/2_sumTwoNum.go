/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// var tail *ListNode

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	if sum > 9 {
		return &ListNode{
			Val:  sum - 10,
			Next: addTwoNumbers(&ListNode{Val: 1, Next: nil}, addTwoNumbers(l1.Next, l2.Next)),
		}
	} else {
		return &ListNode{
			Val:  sum,
			Next: addTwoNumbers(l1.Next, l2.Next),
		}
	}
}