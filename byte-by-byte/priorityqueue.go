package main

/*
Priority Queue
https://www.youtube.com/watch?v=HqPJF2L5h9U
TC: O(logn)  both insertion and deletion
SC: O(1)
*/

/*
func main() {

	pq := NewPriorityQueue(7)
	pq.enqueue(1)
	pq.enqueue(2)
	pq.enqueue(3)
	pq.enqueue(4)
	pq.enqueue(5)
	pq.enqueue(6)
	pq.enqueue(7)
	//pq.enqueue(8)

	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())
	fmt.Println(pq.dequeue())

}
*/
type PriorityQueue struct {
	cap int
	arr []int
	ptr int
}

func NewPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{
		cap: cap,
		arr: make([]int, cap),
		ptr: -1,
	}
}

func (p *PriorityQueue) enqueue(el int) {
	if p.ptr+1 >= p.cap {
		panic("Queue is full!")
	}

	p.ptr = p.ptr + 1
	p.arr[p.ptr] = el

	curr := p.ptr

	for curr > 0 {
		parent := ((curr + 1) / 2) - 1

		if p.arr[parent] < p.arr[curr] {
			swapEl(p.arr, parent, curr)
			curr = parent
		} else {
			break
		}
	}

}

func (p *PriorityQueue) dequeue() int {

	if p.ptr < 0 {
		panic("Queue is empty!")
	}

	el := p.arr[0]

	swapEl(p.arr, p.ptr, 0)
	p.ptr = p.ptr - 1

	curr := 0

	for curr <= p.ptr {

		lIndex := ((curr + 1) * 2) - 1
		rIndex := ((curr + 1) * 2)

		if lIndex <= p.ptr && rIndex <= p.ptr {

			if p.arr[lIndex] > p.arr[rIndex] && p.arr[lIndex] > p.arr[curr] {
				swapEl(p.arr, curr, lIndex)
				curr = lIndex
			} else if p.arr[rIndex] > p.arr[curr] {
				swapEl(p.arr, curr, rIndex)
				curr = rIndex
			} else {
				break
			}

		} else if lIndex <= p.ptr {

			if p.arr[lIndex] > p.arr[curr] {
				swapEl(p.arr, curr, lIndex)
				curr = lIndex
			} else {
				break
			}

		} else {
			break
		}

	}

	return el
}

func (p *PriorityQueue) peek() int {
	return -1
}

func swapEl(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
