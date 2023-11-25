package main

import "fmt"

func main() {
	// variable
	var nama string
	nama = "lia"

	fmt.Println(nama)

	nama = "alea"

	fmt.Println(nama)

	angka := 10

	fmt.Println(angka)
	//

	// constant

	type noKtp string

	var noKtpa noKtp = "blay"

	fmt.Println(noKtpa)

	// operasi matematika

	a := 1
	b := 2

	c := a + b

	fmt.Println(c)

	// operasi perbandingan

	var benar bool = 1 == 1

	fmt.Println(benar)

	// array

	var names [3]string

	names[0] = "lia"
	names[1] = "loli"
	names[2] = "qila"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var value = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(value))

	// slice

	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)

	//
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "diubah"
	fmt.Println(slice1)

	fmt.Println(slice1[1])

	var slice2 = months[11:]

	fmt.Println(slice2)

	var slice3 = append(slice2, "abbam")
	fmt.Println(slice3)

	fmt.Println(slice2)

	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "q"
	newSlice[1] = "p"

	fmt.Println(newSlice)

}
