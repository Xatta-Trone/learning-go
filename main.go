package main

import (
	"fmt"
	"strings"
)

func main() {

	// fmt.Println("test")

	// // variables

	// var name string = "xatta \n";
	// var name2 = "xatta";
	// name3 := "asdf";

	// var num int8 = 127;

	// println(name);
	// println(name2);
	// println(name3);
	// println(num);

	// fmt.Print(name,name2)

	// fmt.Printf("%s %d",name2,num)

	// // arrays

	// var ages [3]int = [3]int{10, 12, 13}

	// names := []string{"asdf", "asdf"}

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // // slices

	// var scores = []int{10, 12, 13}
	// scores[0] = 100

	// // append
	// scores = append(scores, 100)

	// fmt.Println(scores)

	// // slice range

	// rangeOne := scores[1:3]

	// fmt.Println(rangeOne)

	// greet := "my name is xatta"

	// fmt.Println(strings.Contains(greet,"xatta"))
	// fmt.Println(strings.ReplaceAll(greet,"my","your"))

	// loops

	// x := 0

	// for x < 5 {
	// 	fmt.Println("index :: ",x)
	// 	x++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("index i:: ",i)
	// }

	// names := []string{"asdf", "asdf", "asd"}

	// // for i := 0; i < len(names); i++ {
	// // 	fmt.Println(names[i])
	// // }

	// for index, v := range names {
	// 	fmt.Println(v, index)
	// }

	// begum("asdf")

	// testFunc(names,begum)
	// fmt.Println(add(10,20))

	// one,two := funcWithMultipleReturn("Xatta Trone")

	// fmt.Println(one,two)

	// go run main.go greetings.go
	// sayHello("asdf")

	// // maps

	// menu := map[string]float64{
	// 	"soup":   6.44,
	// 	"pie":    7.99,
	// 	"salad":  6.99,
	// 	"toffee": 3.55,
	// }

	// fmt.Println(menu, menu["pie"])

	// // map loops

	// for k, v := range menu {
	// 	fmt.Println(k, v)

	// }

	// // int as key type

	// phonebook := map[int]string{
	// 	1 : "asdf",
	// 	2 : "SDF",
	// }
	// fmt.Println(phonebook, phonebook[1])

	// // map loops

	// for k, v := range phonebook {
	// 	fmt.Println(k, v)

	// }

	// // update item 

	// phonebook[1] = "xatta"


	// fmt.Println(phonebook, phonebook[1])



}

func begum(n string) {
	fmt.Println("begum", n)
}

func add(a float64, b float64) float64 {
	return (a + b)
}

func testFunc(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func funcWithMultipleReturn(n string) (string, string) {

	s := strings.ToUpper(n)

	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	// check the length
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""

}
