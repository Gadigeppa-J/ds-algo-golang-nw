package main

/*
Implement N > 0 stacks using a single array to store all stack data (you may
use auxiliary arrays in your stack object, but all of the objects in all of the stacks must
be in the same array). No stack should be full unless the entire array is full.

N = 3;
capacity = 10;
Stacks stacks = new Stacks(N, capacity);
stacks.put(0, 10);
stacks.put(2, 11);
stacks.pop(0) = 10;
stacks.pop(2) = 11;

TC: O(1)
SC: O(n)

*/

/* func main() {

	s := NewNStack(3, 10)

	s.push(0, 1)
	s.push(0, 2)
	s.push(1, 13)
	s.push(0, 4)
	s.push(1, 15)
	s.push(2, 26)
	s.push(0, 7)
	s.push(2, 28)
	s.push(0, 9)
	s.push(1, 110)

	fmt.Println(s.pop(0))
	fmt.Println(s.pop(0))
	fmt.Println(s.pop(0))
	fmt.Println(s.pop(0))
	fmt.Println(s.pop(0))

	fmt.Println(s.pop(1))
	fmt.Println(s.pop(1))
	fmt.Println(s.pop(1))

	fmt.Println(s.pop(2))

	fmt.Println(s.peek(2))

	s.push(2, 29)
	s.push(2, 30)
	fmt.Println(s.peek(2))

} */

type NStack struct {
	top     []int
	arr     []int
	indexes []int
	next    int
}

func NewNStack(numOfStack int, size int) *NStack {

	s := &NStack{
		top:     make([]int, numOfStack),
		arr:     make([]int, size),
		indexes: make([]int, size),
	}

	for i := 0; i < size; i++ {

		if i+1 == size {
			s.indexes[i] = -1
		} else {
			s.indexes[i] = i + 1
		}

	}

	for i := 0; i < numOfStack; i++ {
		s.top[i] = -1
	}

	return s
}

func (s *NStack) push(stack int, item int) {

	if stack < 0 || stack > len(s.top)-1 {
		panic("invalid stack number")
	}

	if s.next == -1 {
		panic("Stack is full!")
	}

	next := s.next
	prev := s.top[stack]

	s.arr[next] = item
	s.top[stack] = next

	s.next = s.indexes[next]
	s.indexes[next] = prev

}

func (s *NStack) pop(stack int) int {

	if stack < 0 || stack > len(s.top)-1 {
		panic("invalid stack number")
	}

	t := s.top[stack]

	if t <= -1 {
		panic("Stack is empty!")
	}

	item := s.arr[t]
	s.top[stack] = s.indexes[t]
	s.indexes[t] = s.next
	s.next = t

	return item
}

func (s *NStack) peek(stack int) int {

	if stack < 0 || stack > len(s.top)-1 {
		panic("invalid stack number")
	}

	t := s.top[stack]

	if t <= -1 {
		panic("stack is empty!")
	}

	return s.arr[s.top[stack]]
}
