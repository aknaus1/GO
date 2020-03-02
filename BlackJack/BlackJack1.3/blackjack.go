package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Description:
	Makes a random integer that is less then m
Parameters:
	m - upper bound limit
Return:
	comp_choice - random integer that is less then m
*/
func rand_num(m int) int {
	a1 := rand.NewSource(time.Now().UnixNano())
	a2 := rand.New(a1)
	comp_choice := a2.Intn(m)
	return comp_choice
}

func checkBalances(bank []int, numPlayers int) bool{
	for i := 0; i < numPlayers; i++ {
		if bank[i] > 0 {
			return true
		}
	}
	return false
}

/*
Description:
	Creates a new deck of cards
Parameters:
	N/A
Return:
	shuffle_deck(deck) - a shuffled deck of cards
*/
func new_deck() []int {
  var deck []int
  for i := 1; i < 14; i++ {
    for j := 0; j < 4; j++ {
      deck = append(deck, i)
    }
  }
  return shuffle_deck(deck)
}

/*
Description:
	Shuffles a deck of cards
Parameters:
	deck - deck of cards
Return:
	shuffled - a shuffled deck of cards
*/
func shuffle_deck(deck []int) []int {
  var shuffled []int
  for len(deck) > 0 {
    x := rand_num(len(deck))
    shuffled = append(shuffled, deck[x])
    for i := x; i > 0; i-- {
      deck[i] = deck[i-1]
    }
    deck = deck[1:]
  }
  return shuffled
}

/*
Description:
	Determins the value of any array of cards
Parameters:
	deck - array of cards
Return:
	value - value of deck
*/
func get_value(deck []int) int {
  value := 0
  ace := 0
  for i := 0; i < len(deck); i++ {
    if deck[i] > 1 && deck[i] < 10 {
      value += deck[i]
    } else if deck[i] > 9 && deck[i] < 14 {
      value += 10
    } else {
      ace++
      value += 11
    }
  }
  for value > 21 && ace > 0 {
    ace--
    value -= 10
  }
  return value
}

/*
Description:
	Determins the best score held by a non busted player
Parameters:
	players - cards held by all players
Return:
	best - best score held by a non busted player
*/
func best_score(players [][]int) (int) {
	best := 0
	for i := 0; i < len(players); i++ {
		temp := get_value(players[i])
		if temp > best && temp <= 21 {
			best = temp
		}
	}
	return best
}

func print_card(card int) {
	if card == 1 {
		fmt.Printf("A ")
	} else if card == 11 {
		fmt.Printf("J ")
	} else if card == 12 {
		fmt.Printf("Q ")
	} else if card == 13 {
		fmt.Printf("K ")
	} else {
		fmt.Printf("%d ",card)
	}
}

func print_hand(player []int) {
	for i := 0; i < len(player); i++ {
		print_card(player[i])
	}
	fmt.Println("Total Value:", get_value(player))
}

/*
Description:
	Deals one card
Parameters:
	dealto - array of cards held by a player or the dealer
	deck - cards that hevent been dealt
Return:
	dealto - updated array of cards
	deck - cards that hevent been dealt
*/
func deal(dealto []int, deck []int) ([]int, []int){
	dealto = append(dealto, deck[0])
	deck = deck[1:]
	return dealto, deck
}

/*
Description:
	Gives each player $1000 to start the game
Parameters:
	a - number of players in game
Return:
	bank - balances of all players
*/
func initiate(a int) []int {
	bank := make([]int, a)
	for i := 0; i < a; i++ {
		bank[i] = 1000
	}
	return bank
}

func betting(bank []int) []int{
	var bet int
	bets := make([]int, len(bank))
	for i := 0; i < len(bank); i++ {
		if bank[i] > 0 {
			fmt.Println("Player", (i+1), ": how much would you like to bet?")
			fmt.Println("Current balance:", bank[i])
			fmt.Scan(&bet)
			if bet >= bank[i] {
				fmt.Println("Player", (i+1), "is all in!")
				bet = bank[i]
				bets[i] = bet
			} else {
				fmt.Println("Player", (i+1), "is betting $", bet)
				bets[i] = bet
			}
		} else {
			fmt.Println("Player", (i+1), "is out of money")
			bets[i] = 0
		}
	}
	return bets
}

/*
Description:
	Takes bets from players and then deals 2 cards to every player
	and the dealer from a new randomly generated deck of cards
Parameters:
	bank - balances held by all players
Return:
	players - array of players each with 2 cards
	dealer - array with 2 cards held by dealer
	deck - cards that hevent been dealt
	bets - empty array of player bets
*/
func start(bank []int) ([][]int, []int, []int, []int){
	players := make([][]int, len(bank))
	var dealer []int
	deck := new_deck()
	bets  := betting(bank)
	for i := 0; i < 2; i++ {
		for j := 0; j < len(bank); j++ {
			players[j], deck = deal(players[j], deck)
		}
		dealer, deck = deal(dealer, deck)
	}
	return players, dealer, deck, bets
}

/*
Description:
	Determins the winners of each round, and updates everyone's balance
	according to how much they bet at the beginning of their turn
Parameters:
	players - cards held by all players
	bank - balances held by all players
	bets - bets made by all players
	dealer - cards held by the dealer
Return:
	bank - new balances held by all players
*/
func winners(players [][]int, bank []int, bets []int, dealer []int) []int {
	for i := 0; i < len(players); i++ {
		if get_value(dealer) > 21 && get_value(players[i]) <= 21 {
			fmt.Println("Player", (i+1), "won")
			bank[i] += bets[i]
			fmt.Println("New balance:", bank[i])
		} else if get_value(players[i]) > 21 {
			fmt.Println("Player", (i+1), "lost")
			bank[i] -= bets[i]
			fmt.Println("New balance:", bank[i])
		} else if get_value(dealer) < get_value(players[i]) {
			fmt.Println("Player", (i+1), "won")
			bank[i] += bets[i]
			fmt.Println("New balance:", bank[i])
		} else if get_value(dealer) == get_value(players[i]) {
			fmt.Println("Player", (i+1), "pushed")
			fmt.Println("Same Balance:", bank[i])
		} else {
			fmt.Println("Player", (i+1), "lost")
			bank[i] -= bets[i]
			fmt.Println("New balance:", bank[i])
		}
	}
	return bank
}

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
func player_play(player []int, deck []int, dealer int) ([]int, []int){
	var choice string
	fmt.Printf("The dealer is showing a ")
	print_card(dealer)
	fmt.Println()
	fmt.Printf("Your cards: ")
	print_hand(player)
	for get_value(player) <= 21{
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
	Cycles through dealers turn hitting until it is higher then all other
	non busted players or it reaches at least 17 whichever is lower
Parameters:
	dealer - cards held by the dealer
	deck - cards that hevent been dealt
	max - highest score held by the players
Return:
	dealer - cards now held by the dealer
	deck - cards that hevent been dealt
*/
func dealer_play(dealer []int, deck []int, max int) ([]int, []int) {
	fmt.Println("Dealer is starting with:", get_value(dealer))
	for get_value(dealer) < 17 && get_value(dealer) < max {
		fmt.Println("Dealer hits...")
		dealer, deck = deal(dealer, deck)
		fmt.Println("Dealer now has:", get_value(dealer))
		time.Sleep(1000 * time.Millisecond)
	}
	if get_value(dealer) > 21 {
		fmt.Println("Dealer busted")
	} else {
		fmt.Println("Dealer is ending with:", get_value(dealer))
	}
	return dealer, deck
}

/*
Description:
	Cycles through one round of blackjack
Parameters:
	bank - balances held by all players
Return:
	bank - new balances held by all players
*/
func play (bank []int) ([]int){
	players, dealer, deck, bets := start(bank)
	for i := 0; i < len(players); i++ {
		if bank[i] > 0 {
			fmt.Println("Player", (i+1),"s turn.")
			players[i], deck = player_play(players[i], deck, dealer[0])
			time.Sleep(1000 * time.Millisecond)
		}
	}
	x := best_score(players)
	dealer, deck = dealer_play(dealer, deck, x)
	time.Sleep(1000 * time.Millisecond)
	bank = winners(players, bank, bets, dealer)
	return bank
}

/*
Description: Lets play some blackjack!
Parameters: N/A
Return: N/A
*/
func main() {
	var numPlayers int
	var choice string
	fmt.Println("How many players?")
	fmt.Scan(&numPlayers)
	bank := initiate(numPlayers)
	for i := 0; i < 1; i = i {
		bank = play(bank)
		left := checkBalances(bank, numPlayers)
		if left {
			fmt.Println("Would you like to play again?")
			fmt.Scan(&choice)
			if choice == "no" {
				i = 1
			}
		} else {
			fmt.Println("Game Over: Everyone is out of money.")
			fmt.Println("Good Bye")
			i = 1
		}
	}
}
