package main

type ListNode struct {
	Val  int
	Next ListNode
}

func main() {

}

func mergeListNodes(p, q *ListNode) *ListNode {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	if p.Val < q.Val {
		p.Next = mergeListNodes(p.Next, q)
		return p
	}
	q.Next = mergeListNodes(p, q.Next)
	return q
}

func mergeListNodes1(p, q *ListNode) *ListNode {
	//p，q长短
	for p != nil || q != nil {
		if p.Val < q.Val {

			p = p.Next
		} else {
			tmp := p.Next
			p.Next = q
			p.Next.Next = tmp
			q = q.Next
		}
		//p空q不空
		if p.Next == nil && q.Next != nil {
			p.Next = q
		}
	}
	return p
}
