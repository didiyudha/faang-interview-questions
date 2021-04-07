package main

func getPartitionIndex(arr []int, left, right int) int {
	i := left
	for j := left; j <= right; j++ {
		if arr[j] <= arr[right] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	return i - 1
}

func quickSelect(arr []int, left, right, idxToFind int) int {

	if right < left {
		return -1
	}

	partitionIdx := getPartitionIndex(arr, left, right)

	if partitionIdx == idxToFind {
		return arr[partitionIdx]
	}

	if partitionIdx < idxToFind {
		return quickSelect(arr, partitionIdx+1, right, idxToFind)
	}

	return quickSelect(arr, left, partitionIdx-1, idxToFind)

}

func kthLargest(arr []int, k int) int {
	left := 0
	right := len(arr) - 1
	idxToFind := len(arr) - k
	return quickSelect(arr, left, right, idxToFind)
}

func main() {

}
