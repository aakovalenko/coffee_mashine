package main

import (
	"fmt"
	"coffee-module/hello"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

const coffeeMashineSay = "**** coffee mashine says *****"

var water uint = 400
var milk uint = 540
var coffeeBeans uint = 120
var disposableCups uint = 9
var money uint = 550
var checkWater, checkMilk, checkBeans, checkCups bool

func main() {
	var action string
	var chooseCoffee string

	color.Yellow(hello.HelloWorld(coffeeMashineSay))

	for {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)

		switch action {
		case "buy":
			fmt.Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
			fmt.Scan(&chooseCoffee)
			buyCoffee(chooseCoffee)
			fmt.Println()
		case "fill":
			water, milk, coffeeBeans, disposableCups = fill(water, milk, coffeeBeans, disposableCups)
			fmt.Printf("\n")
		case "take":
			money = takeMoney(money)
			fmt.Printf("\n")
		case "remaining":
			fmt.Printf("\n")
			greeting(water, milk, coffeeBeans, disposableCups, money)
		case "exit":
			return
		}
	}
}

func greeting(water, milk, coffeeBeans, disposableCups, money uint) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d of water\n", water)
	fmt.Printf("%d of milk\n", milk)
	fmt.Printf("%d of coffee beans\n", coffeeBeans)
	fmt.Printf("%d of disposable cups\n", disposableCups)
	fmt.Printf("$%d of money\n", money)
}

func espresso(water, coffeeBeans, disposableCups, money uint) (uint, uint, uint, uint) {
	water, checkWater = checkElement(water, 250, "Sorry, not enough water!", checkWater)
	if checkWater {
		coffeeBeans, checkBeans = checkElement(coffeeBeans, 16, "Not enough resources!", checkBeans)
		disposableCups, checkCups = checkElement(disposableCups, 1, "Not enough resources!", checkCups)
		color.Green("I have enough resources, making you a coffee!")
	}

	money = money + 4
	return water, coffeeBeans, disposableCups, money
}

func latte(water, milk, coffeeBeans, disposableCups, money uint) (uint, uint, uint, uint, uint) {
	water, checkWater = checkElement(water, 350, "Sorry, not enough water!", checkWater)
	if checkWater {
		milk, checkMilk = checkElement(milk, 75, "Not enough resources!", checkMilk)
		coffeeBeans, checkBeans = checkElement(coffeeBeans, 20, "Not enough resources!", checkBeans)
		disposableCups, checkCups = checkElement(disposableCups, 1, "Not enough resources!", checkCups)
		fmt.Println("I have enough resources, making you a coffee!")
	}
	money = money + 7
	return water, milk, coffeeBeans, disposableCups, money
}

func cappuccino(water, milk, coffeeBeans, disposableCups, money uint) (uint, uint, uint, uint, uint) {
	water, checkWater = checkElement(water, 200, "Sorry, not enough water!", checkWater)
	if checkWater {
		milk, checkMilk = checkElement(milk, 100, "Not enough resources!", checkMilk)
		coffeeBeans, checkBeans = checkElement(coffeeBeans, 12, "Not enough resources!", checkBeans)
		disposableCups, checkCups = checkElement(disposableCups, 1, "Not enough resources!", checkCups)
		fmt.Println("I have enough resources, making you a coffee!")
	}

	money = money + 6
	return water, milk, coffeeBeans, disposableCups, money
}

func buyCoffee(typeOfCoffee string) {
	if typeOfCoffee == "1" {
		water, coffeeBeans, disposableCups, money = espresso(water, coffeeBeans, disposableCups, money)
	} else if typeOfCoffee == "2" {
		water, milk, coffeeBeans, disposableCups, money = latte(water, milk, coffeeBeans, disposableCups, money)
	} else if typeOfCoffee == "3" {
		water, milk, coffeeBeans, disposableCups, money = cappuccino(water, milk, coffeeBeans, disposableCups, money)
	} else if typeOfCoffee == "back" {
		return
	}
}

func checkElement(element uint, cal uint, message string, check bool) (uint, bool) {

	var resources uint = element - cal

	if resources > element {
		color.Red(message)
		check = false
	} else {
		element = element - cal
		check = true
	}
	return element, check
}

func fill(water, milk, coffeeBeans, disposableCups uint) (uint, uint, uint, uint) {
	var fillWater, fillMilk, fillCoffeeBeans, fillDisposableCups uint
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&fillWater)
	water += fillWater
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&fillMilk)
	milk += fillMilk
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&fillCoffeeBeans)
	coffeeBeans += fillCoffeeBeans
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&fillDisposableCups)
	disposableCups += fillDisposableCups
	return water, milk, coffeeBeans, disposableCups
}

func takeMoney(money uint) uint {
	giveYou := money
	money -= money
	emoji.Printf("\n:moneybag:I gave you :dollar:%d :moneybag:", giveYou)
	return money
}
