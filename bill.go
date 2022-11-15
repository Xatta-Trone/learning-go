package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill

func newBill(name string) bill {

	b := bill{
		name: name,
		items: map[string]float64{
			
		},
		tip: 0,
	}

	return b

}

// format the bill

func (b bill) format() string {

	fs := "Bill breakdown: \n"
	var total float64 = 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	// tip
	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "total:", total+b.tip)

	return fs

}

// update tip 
func (b *bill) updateTip(tip float64){
	// (*b).tip = tip
	b.tip = tip
}

// add item to bill

func (b *bill) addItem(name string, price float64)  {
	b.items[name] = price
	
}