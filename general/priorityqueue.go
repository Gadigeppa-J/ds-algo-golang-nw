package main

/*
func main() {

	pq := NewPriorityQueue()

	pq.Insert(2)
	pq.Insert(5)
	pq.Insert(1)
	pq.Insert(4)
	pq.Insert(3)
	pq.Insert(6)

	fmt.Println(pq.Remove())
	fmt.Println(pq.Remove())
	fmt.Println(pq.Remove())
	fmt.Println(pq.Remove())
	fmt.Println(pq.Remove())
	fmt.Println(pq.Remove())

	fmt.Println(pq.Remove())

}
*/

type PriorityQueue struct {
	arr []*PriorityQueueItem
}

type PriorityQueueItem struct {
	Val int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		arr: []*PriorityQueueItem{},
	}
}

func (q *PriorityQueue) Insert(i int) {
	if len(q.arr) == 0 {
		q.arr = append(q.arr, &PriorityQueueItem{Val: i})
	} else {

		// insert at last
		currIndex := len(q.arr)
		q.arr = append(q.arr, &PriorityQueueItem{Val: i})

		for true {

			parentIndex := ((currIndex + 1) / 2) - 1

			if q.arr[parentIndex].Val < q.arr[currIndex].Val {
				temp := q.arr[currIndex]
				q.arr[currIndex] = q.arr[parentIndex]
				q.arr[parentIndex] = temp

				currIndex = ((currIndex + 1) / 2) - 1

				if currIndex == 0 {
					break
				}

			} else {
				break
			}

		}
	}

}

func (q *PriorityQueue) Remove() int {

	if len(q.arr) == 0 {
		panic("Queue is empty!")
	}

	val := q.arr[0].Val

	// Move last element to root
	q.arr[0] = q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]

	if len(q.arr) == 0 {
		return val
	}

	currIndex := 0

	for true {

		leftChild := ((currIndex + 1) * 2) - 1
		rightChild := ((currIndex + 1) * 2)

		if leftChild >= len(q.arr) {
			// Exit if no left child
			break
		}

		if rightChild >= len(q.arr) {

			// if no Right child, use only left

			if q.arr[leftChild].Val > q.arr[currIndex].Val {
				// swap
				temp := q.arr[leftChild]
				q.arr[leftChild] = q.arr[currIndex]
				q.arr[currIndex] = temp

				currIndex = leftChild
				continue
			} else {
				break
			}

		} else {

			// if boht left and right child exists

			if q.arr[leftChild].Val > q.arr[rightChild].Val {

				if q.arr[leftChild].Val > q.arr[currIndex].Val {

					// if left is greater than right

					// swap
					temp := q.arr[leftChild]
					q.arr[leftChild] = q.arr[currIndex]
					q.arr[currIndex] = temp

					currIndex = leftChild
					continue
				}

			} else {

				// if right is greater than left

				if q.arr[rightChild].Val > q.arr[currIndex].Val {
					// swap
					temp := q.arr[rightChild]
					q.arr[rightChild] = q.arr[currIndex]
					q.arr[currIndex] = temp

					currIndex = rightChild
					continue
				}
			}

		}
	}

	return val

}
