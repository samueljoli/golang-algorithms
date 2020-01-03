package main

func main(arr []int, key int) int {

	// block size to jump
	sz := len(arr)
	step := int(math.Sqrt(float(sz)))
	prev := 0

	// finding the block
	for arr[int(math.Min(float64(step), float64(sz))) - 1] < key {
		prev = step
		step += int(math.Sqrt(float64(sz)))
		if prev >= sz {
			return -1
		}
	}

	// linear search the block
	for arr[prev] < key {
		prev++
		if prev == int(math.Min(float64(step), float64(sz))) {
			return -1
		}
	}

	if arr[prev] == key {
		return prev
	}
	
	return -1
}
