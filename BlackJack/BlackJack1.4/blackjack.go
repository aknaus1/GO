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
	input_yes, input_no, user_name := sign_in()
	fmt.Println("How many players?")
	fmt.Scan(&numPlayers)
	bank, easy, input_yes, input_no := initiate(numPlayers, input_yes, input_no)
	if easy{
		hints(bank, numPlayers, input_yes, input_no)
	} else {
		basic(bank, numPlayers, input_yes, input_no)
	}
	sign_out(user_name, input_yes, input_no)
}
