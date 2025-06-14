package main

import "fmt"

func main() {
	
	//var expenses float64
	//var taxRate float64

	revenue :=  getUserInput("Revenue: ")
	expenses:= getUserInput("Expenses: ")
	taxRate:= getUserInput("taxRate: ")


	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	
	fmt.Printf("%.1f\n",ebt)
	fmt.Printf("%.1f\n",profit)
	fmt.Printf("%.3f\n", ratio)
}


func calculateFinancials (revenue , expenses , taxRate float64)(float64 , float64, float64 ){
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt , profit, ratio
}

func getUserInput (infoText string )float64{
	var userInput float64
	fmt.Println(infoText)
	fmt.Scan(&userInput)
	return userInput
}