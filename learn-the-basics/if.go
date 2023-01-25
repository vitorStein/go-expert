package main

func main() {

	a := 1
	b := 2

	if a > b {
		println("A maior que B")
	} else {
		println("B maior que A")
	}

	if b > a && a == 1 {
		println("B maior que A e a = 1")
	}

	if a > b || a == 1 {
		println("A maior que B ou a = 1")
	}

	switch a {
	case 0:
		println("a = 0")
	case 2:
		println("a = 1")
	default:
		println("a não é 0 nem 2")
	}
}
