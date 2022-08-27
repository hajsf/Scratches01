package global

func NormalizeNumber(e rune) int {
	switch e {
	case 1632:
		return 0
	case 1633:
		return 1
	case 1634:
		return 2
	case 1635:
		return 3
	case 1636:
		return 4
	case 1637:
		return 5
	case 1638:
		return 6
	case 1639:
		return 7
	case 1640:
		return 8
	case 1641:
		return 9
	}
	return 0
}

/*
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	latin := "0123456789"

	for i, e := range latin {
		fmt.Println(i, e)
	}

	fmt.Println("*****")

	arabic := "٩٨٧٦٥٤٣٢١٠"
	//	fmt.Println("first: ", []byte(arabic[0]))
	var new string //[]int
	for i, e := range arabic {
		fmt.Println(i, e)
		switch e {

		case 1632:
			new = fmt.Sprintf("%s%v", new, "0")
			//new = fmt.append(new, 0)
		case 1633:
			new = fmt.Sprintf("%s%v", new, "1")
			// new = append(new, 1)
		case 1634:
			new = fmt.Sprintf("%s%v", new, "2")
			// new = append(new, 2)
		case 1635:
			new = fmt.Sprintf("%s%v", new, "3")
			// new = append(new, 3)
		case 1636:
			new = fmt.Sprintf("%s%v", new, "4")
			// new = append(new, 4)
		case 1637:
			new = fmt.Sprintf("%s%v", new, "5")
			// new = append(new, 5)
		case 1638:
			new = fmt.Sprintf("%s%v", new, "6")
			// new = append(new, 6)
		case 1639:
			new = fmt.Sprintf("%s%v", new, "7")
			// new = append(new, 7)
		case 1640:
			new = fmt.Sprintf("%s%v", new, "8")
			// new = append(new, 8)
		case 1641:
			new = fmt.Sprintf("%s%v", new, "9")
			// new = append(new, 9)
		}
	}
	fmt.Println("*****")
	fmt.Println(arabic)
	fmt.Println(new)
	fmt.Println("*****")
	for i, e := range new {
		fmt.Println(i, e)
	}
}

*/
