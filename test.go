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

	greet := "my name is xatta"

	fmt.Println(strings.Contains(greet,"xatta"))
	fmt.Println(strings.ReplaceAll(greet,"my","your"))

}
