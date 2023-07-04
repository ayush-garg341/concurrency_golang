package race

import "fmt"

func RaceConditions() {
	var balance = 100
	go deposit(&balance, 10)
	withdraw(&balance, 50)
	fmt.Println(balance)
}

func deposit(balance *int, amount int) {
	*balance += amount
}

func withdraw(balance *int, amount int) {
	*balance -= amount
}
