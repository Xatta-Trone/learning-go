package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	// struct and customs

	// mybill := newBill("xatta's bill")

	// mybill.updateTip(10)

	// mybill.addItem("xatta",10.0)
	// mybill.addItem("trone",15.0)

	// fmt.Println(mybill.format())

	mybill := createBill()

	promptOptions(mybill)

	fmt.Println(mybill)

}

func getInput(question string, r *bufio.Reader) (string, error) {
	fmt.Print(question)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option: (a- add item, s - save bill, t - add tip) ", reader)

	fmt.Println(opt)

	switch opt {
	case "a":

		name, _ := getInput("add item name ", reader)
		price, _ := getInput("add item price ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			// panic("Can not parse the value")
			fmt.Println("Can not parse the value")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You choose s")
	case "t":
		tip, _ := getInput("add tip ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			// panic("Can not parse the value")
			fmt.Println("Can not parse the value")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("tip added", tip)
		promptOptions(b)
	default:
		fmt.Println("Invalid option..")
		promptOptions(b)

	}

}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	// print(name)

	name, _ := getInput("Create a new bill name ", reader)

	b := newBill(name)
	fmt.Println("Created bill for ", b.name)

	return b

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
