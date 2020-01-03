package main

func main() {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return i
		}
	}

	return -1
}
