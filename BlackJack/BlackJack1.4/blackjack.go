package main

import (
	"fmt"
)

/*
Description: Lets play some blackjack!
Parameters: N/A
Return: N/A
*/
func main() {
	var numPlayers int
	//RULES AND DIRECTIONS
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println("Welcome to our Blackjack game!")
	fmt.Println("The goal of Blackjack is to get as close to 21 as possible without going over.")
	fmt.Println("You will only be competing against the dealer and not your fellow players.")
	fmt.Println("All betting is at the beginning of each round.")
	fmt.Println("")
	fmt.Println("WINS --- LOSSS --- PUSH")
	fmt.Println("Win: You score better then the dealer while staying below 21.")
	fmt.Println("Loss: You go over 21 or score less then the dealer.")
	fmt.Println("Push: You have the same score as the dealer. You neither lose nor win.")
	fmt.Println("")
	fmt.Println("Blackjack")
	fmt.Println("If anyone is initially dealt a hand with a value of 21 they automatically win and the round is over.")
	fmt.Println("All players with Blackjack recieve 1.5 x the value of their bets.")
	fmt.Println("Those without Blackjack lose the value of their bets as they would normally.")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	//sign_in LOCATED IN file_management.go
	input_yes, input_no, user_name := sign_in()
	fmt.Println("How many players?")
	fmt.Scan(&numPlayers)
	//initiate LOCATED IN base_func.go
	bank, easy, input_yes, input_no := initiate(numPlayers, input_yes, input_no)
	if easy == 0{
		//goto hints.go
		hints(bank, numPlayers, input_yes, input_no)
	} else if easy == 1{
		//goto normal.go
		normal(bank, numPlayers, input_yes, input_no)
	} else if easy == 2{
		sym(bank, numPlayers, input_yes, input_no)
	} else {
		cheat(bank, numPlayers, input_yes, input_no)
	}
	//sign_out LOCATED IN file_management.go
	sign_out(user_name, input_yes, input_no)
}
