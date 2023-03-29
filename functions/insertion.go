package functions

/*func main() {
	var n = NumGenerator(10000, 100000)
	timeStart := time.Now()
	insertion(n)
	finished := time.Now().Sub(timeStart).Seconds()
	fmt.Println(n)
	fmt.Printf("Finished After %f\n", finished)
}*/

func Insertion(n []int) {
	i := 1

	for i < len(n) {
		j := i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]
			j--
		}
		i++
	}
}
