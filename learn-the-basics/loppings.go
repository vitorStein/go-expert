package main

func main() {

	// For básico
	for i := 0; i <= 10; i++ {
		println(i)
	}

	// Range
	numbers := []int{1, 2, 3, 4, 5, 6}

	for key, value := range numbers {
		println(key, value)

	}

	// For parecido com while
	j := 0
	for j < 10 {
		println(j)
		j++
	}

	// loop infinito

	/* for {
		print("Ok")
	} */

}
