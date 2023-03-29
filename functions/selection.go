package functions

/*func main() {
	var n = NumGenerator(10000, 100000)

	timeStart := time.Now()
	selection(n)
	finished := time.Now().Sub(timeStart).Seconds()
	fmt.Println(n)
	fmt.Printf("Finished After %f\n", finished)
}*/

func Selection(n []int) {
	i := 1
	for i < len(n)-1 {
		j := i + 1
		var minIndex = i
		for j < len(n) {
			if n[j] < n[minIndex] {
				minIndex = j
			}
			j++
		}

		if n[minIndex] < n[i] {
			n[i], n[minIndex] = n[minIndex], n[i]
		}
		i++
	}
}
