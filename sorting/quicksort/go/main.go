package main

func Partition(arr []int, left, right int) int {
	i := left
	for j := left; j <= right; j++ {
		if arr[j] <= arr[right] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	return i - 1
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		partitionIndex := Partition(arr, left, right)
		QuickSort(arr, left, partitionIndex-1)
		QuickSort(arr, partitionIndex+1, right)
	}
}

func main() {

}
