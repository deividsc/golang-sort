package functions

func CountSort(arr []int) []int {
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	indices := make([]int, max+1)

	i := 0
	for i < len(arr) {
		indices[arr[i]]++
		i++
	}

	j := 1

	for j < len(indices) {
		indices[j] += indices[j-1]
		j++
	}

	k := 0

	result := make([]int, len(arr))
	for k < len(arr) {
		result[indices[arr[k]]-1] = arr[k]
		indices[arr[k]]--

		k++
	}

	return result
}

/*func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}

	fmt.Println(countSort(n))
}*/
