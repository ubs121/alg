package main

func kadane(arr []int, n int) (max, left, right int) {
	sum := arr[0]
	max = arr[0]
	start := 0
	for i := 1; i < n; i++ {
		sum += arr[i]
		if arr[i] > sum {
			sum = arr[i]
			start = i + 1
		}
		if max < sum {
			max = sum
			left = start
			right = i
		}
	}
	return
}
