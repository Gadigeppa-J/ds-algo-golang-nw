package utils

// stack implementation
type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack{
		arr: []int{},
	}
}

func (s *Stack) Push(i int) {
	s.arr = append(s.arr, i)
}

func (s *Stack) Pop() int {

	if s.IsEmpty() {
		panic("stack is empty!")
	}

	i := s.arr[len(s.arr)-1]

	s.arr = s.arr[:len(s.arr)-1]

	return i
}

func (s *Stack) IsEmpty() bool {

	if len(s.arr) == 0 {
		return true
	}

	return false
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty!")
	}
	return s.arr[len(s.arr)-1]
}

// Queue implementation
type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	return &Queue{
		arr: []int{},
	}
}

func (q *Queue) Enqueue(i int) {
	q.arr = append(q.arr, i)
}

func (q *Queue) Dequeue() int {
	if len(q.arr) == 0 {
		panic("Queue is empty!")
	}
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}

func (q *Queue) IsEmpty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

func (q *Queue) Peek() int {

	if len(q.arr) == 0 {
		panic("Queue is empty!")
	}
	return q.arr[0]
}

func (q *Queue) Size() int {
	return len(q.arr)
}

// Queue generic
type GenericQueue struct {
	arr []interface{}
}

func NewGenericQueue() *GenericQueue {
	return &GenericQueue{
		arr: []interface{}{},
	}
}

func (q *GenericQueue) Enqueue(i interface{}) {
	q.arr = append(q.arr, i)
}

func (q *GenericQueue) Dequeue() interface{} {
	if len(q.arr) == 0 {
		panic("Queue is empty!")
	}
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v
}

func (q *GenericQueue) IsEmpty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}

func (q *GenericQueue) Peek() interface{} {

	if len(q.arr) == 0 {
		panic("Queue is empty!")
	}
	return q.arr[0]
}

func (q *GenericQueue) Size() int {
	return len(q.arr)
}

// Linked list node defination
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(i int) {

	node := &ListNode{
		Val: i,
	}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}

}

func (l *LinkedList) Head() *ListNode {
	return l.head
}

func (l *LinkedList) Tail() *ListNode {
	return l.tail
}

// Binary Tree Node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}
