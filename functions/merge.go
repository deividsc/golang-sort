package functions

func merge(fp, sp []int) []int {
	n := make([]int, len(fp)+len(sp))

	fpIndex := 0
	spIndex := 0
	nIndex := 0

	for fpIndex < len(fp) && spIndex < len(sp) {
		if fp[fpIndex] < sp[spIndex] {
			n[nIndex] = fp[fpIndex]
			fpIndex++
		} else if sp[spIndex] < fp[fpIndex] {
			n[nIndex] = sp[spIndex]
			spIndex++
		}
		nIndex++
	}

	for fpIndex < len(fp) {
		n[nIndex] = fp[fpIndex]
		fpIndex++
		nIndex++
	}

	for spIndex < len(sp) {
		n[nIndex] = sp[spIndex]
		spIndex++
		nIndex++
	}

	return n
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	fp := MergeSort(arr[0 : len(arr)/2])
	sp := MergeSort(arr[len(arr)/2:])

	return merge(fp, sp)
}

/*func main() {
	var n = NumGenerator(10000, 100000)
	timeStart := time.Now()
	res := mergeSort(n)
	finished := time.Now().Sub(timeStart).Seconds()
	fmt.Println(res)
	fmt.Printf("Finished After %f\n", finished)
}*/
