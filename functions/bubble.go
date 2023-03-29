package functions

func Bubble(n []int) {
	var Done bool
	for !Done {
		Done = true
		for i, _ := range n {
			if i != len(n)-1 && n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				Done = false
			}
		}
	}
}

/*func main() {
	var n = NumGenerator(10000, 100000)

	timeStart := time.Now()
	bubble(n)
	finished := time.Now().Sub(timeStart).Seconds()

	fmt.Println(n)
	fmt.Printf("Finished After %f\n", finished)
}*/
