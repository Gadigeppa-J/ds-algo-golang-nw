package main

/*
func main() {
	aq := NewArrayQueue(5)

	aq.enqueue(1)
	aq.enqueue(2)
	aq.enqueue(3)
	aq.enqueue(4)
	aq.enqueue(5)

	fmt.Println(aq.dequeue())
	fmt.Println(aq.dequeue())
	fmt.Println(aq.dequeue())
	fmt.Println(aq.dequeue())
	fmt.Println(aq.peek())
	fmt.Println(aq.dequeue())

}
*/

type ArrayQueue struct {
	arr []int
	p   int
	q   int
	cap int
}

func NewArrayQueue(cap int) *ArrayQueue {
	aq := &ArrayQueue{
		arr: make([]int, cap),
		p:   -1,
		q:   -1,
		cap: cap,
	}
	return aq
}

func (a *ArrayQueue) isempty() bool { // O(1)
	if a.p == -1 && a.q == -1 {
		return true
	}
	return false
}

func (a *ArrayQueue) enqueue(item int) { // O(1)

	if a.isempty() {
		a.p, a.q = 0, 0
	} else {
		next := (a.q + 1) % a.cap
		if next == a.p {
			panic("queue is full")
		}
		a.q = next
	}

	a.arr[a.q] = item
}

func (a *ArrayQueue) dequeue() int { // O(1)

	var item int
	if a.isempty() {
		panic("queue is empty")
	} else {
		item = a.arr[a.p]
		next := (a.p + 1) % a.cap
		if a.p == a.q {
			a.p, a.q = -1, -1
		} else {
			a.p = next
		}
	}
	return item
}

func (a *ArrayQueue) peek() int { // O(1)

	if a.isempty() {
		panic("queue is empty")
	}

	return a.arr[a.p]
}
