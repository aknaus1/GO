package main

import (
	"fmt"
	"time"
)

/*
Description:
	Cycles through one players turn allowing them to "hit"
	until they decide to stay or go over 21
Parameters:
	player - cards held by the player
	deck - cards that hevent been dealt
Return:
	player - cards now held by the player
	deck - cards that hevent been dealt
*/
func nplayer_play(player []int, deck []int, dealer int) ([]int, []int){
	var choice string
	fmt.Printf("The dealer is showing a ")
	print_card(dealer)
	fmt.Println()
	fmt.Printf("Your cards: ")
	print_hand(player)
	for get_value(player) < 21{
		fmt.Println("Hit or stay?")
		fmt.Scan(&choice)
		if choice == "hit" {
			player, deck = deal(player, deck)
			fmt.Printf("Your cards: ")
			print_hand(player)
		} else if choice == "stay" {
			return player, deck
		} else {
			fmt.Println("Sorry, I did not understand that.")
		}
	}
	return player, deck
}

/*
Description:
	Cycles through one round of blackjack
Parameters:
	bank - balances held by all players
Return:
	bank - new balances held by all players
*/
func nplay (bank []int) ([]int){
	players, dealer, deck, bets := start(bank)
	blackjack := check_Blackjack(players, dealer)
	if blackjack {
		bank = blackjack_winners(players, bank, bets, dealer)
		return bank
	}
	for i := 0; i < len(players); i++ {
		if bank[i] > 0 {
			fmt.Println("Player", (i+1),"s turn.")
			players[i], deck = nplayer_play(players[i], deck, dealer[0])
			if get_value(players[i]) > 21 {
				fmt.Println("Player", (i+1)," busted.")
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
	x := best_score(players)
	dealer, deck = dealer_play(dealer, deck, x)
	time.Sleep(1000 * time.Millisecond)
	bank = winners(players, bank, bets, dealer)
	return bank
}

func normal(bank []int, numPlayers int, input_yes []string, input_no []string) {
	var choice string
	var response bool
  for i := 0; i < 1; i = i {
		bank = nplay(bank)
		left := checkBalances(bank, numPlayers)
		if left {
			fmt.Println("Would you like to play again?")
			fmt.Scan(&choice)
			response, input_yes, input_no = get_response(choice, input_yes, input_no)
			if response == false {
				i = 1
			}
		} else {
			fmt.Println("Game Over: Everyone is out of money.")
			i = 1
		}
	}
}
