package main

import (
	"fmt"

	"github.com/paolochang/learngo/wordbook"
)

func main() {
	
	// fmt.Println("Hello")
	// message := greetings.Hello("Paolo")
	// fmt.Println(message)

	// account := accounts.Account{Owner: "Paolo", Balance: 1000}
	// fmt.Println(account)

	// account := accounts.NewAccount("Paolo")
	// fmt.Println(account)
	// account.Deposit(100)
	// fmt.Println(account.Balance())
	// account.Withdraw(20)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(30)
	// if (err != nil) {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// account.ChangeOwner("Samuel")
	// fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account)

	wb := wordbook.Wordbook{"hello": "greeting"}
	// dictionary["hello"] = "hello"
	fmt.Println(wb)
	err := wb.Add("goodbye", "greeting when you leave")
	if err != nil {
		fmt.Println(err)
	}
	err3 := wb.Update("goodbye", "헤어질때 인사")
	if err3 != nil {
		fmt.Println(err3)
	}
	def, err := wb.Search("goodbye")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	err2 := wb.Add("hello", "greeting")
	if err2 != nil {
		fmt.Println(err2)
	}

}
