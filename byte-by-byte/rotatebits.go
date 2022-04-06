package main

func rotate(num uint32, n int) uint32 {
	return (num >> n) | (num << (32 - n))
}
