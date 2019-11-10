package main

import "fmt"

// Payer - интерфейс плательщика
type Payer interface {
	Pay(int) error
}

// Wallet - кошелек счета
type Wallet struct {
	Cash int
}

// Pay - метод кошелька
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

//Bay - функция реализующая интерфейс Payer
func Bay(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо %T\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Bay(myWallet)
}
