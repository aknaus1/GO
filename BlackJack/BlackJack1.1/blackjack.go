package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand_num(m int) int {
	a1 := rand.NewSource(time.Now().UnixNano())
	a2 := rand.New(a1)
	comp_choice := a2.Intn(m)
	return comp_choice
}

func new_deck() []int {
  var deck []int
  for i := 1; i < 14; i++ {
    for j := 0; j < 4; j++ {
      deck = append(deck, i)
    }
  }
  return shuffle_deck(deck)
}

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

func deal(dealto []int, deck []int) ([]int, []int){
	dealto = append(dealto, deck[0])
	deck = deck[1:]
	return dealto, deck
}

func start(a int) ([][]int, []int, []int){
	var dealer []int
	players := make([][]int, a)
	deck := new_deck()
	for i := 0; i < 2; i++ {
		for j := 0; j < a; j++ {
			players[j], deck = deal(players[j], deck)
		}
		dealer, deck = deal(dealer, deck)
	}
	return players, dealer, deck
}

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

func winners(players [][]int, dealer []int) {
	for i := 0; i < len(players); i++ {
		if get_value(dealer) <= 21 {
			if get_value(players[i]) > get_value(dealer) && get_value(players[i]) <= 21 {
				fmt.Println("Player", (i+1), "won")
			} else {
				fmt.Println("Player", (i+1), "lost")
			}
		} else {
			if get_value(players[i]) <= 21 {
				fmt.Println("Player", (i+1), "won")
			} else {
				fmt.Println("Player", (i+1), "lost")
			}
		}
	}
}

func play_round(player []int, deck []int) ([]int, []int){
	var choice string
	play := true
	fmt.Println("You are at:", get_value(player))
	for play && get_value(player) <= 21{
		fmt.Println("Hit or stay?")
		fmt.Scan(&choice)
		if choice == "hit" {
			player, deck = deal(player, deck)
			fmt.Println("You are now at:", get_value(player))
		} else if choice == "stay" {
			return player, deck
		} else {
			fmt.Println("Sorry, I'm did not understand that.")
		}
	}
	return player, deck
}

func dealer_play(dealer []int, deck []int, max int) ([]int, []int) {
	fmt.Println("Dealer is starting with:", get_value(dealer))
	for get_value(dealer) < 17 && get_value(dealer) < max {
		dealer, deck = deal(dealer, deck)
	}
	if get_value(dealer) > 21 {
		fmt.Println("Dealer is bust")
	} else {
		fmt.Println("Dealer is ending with:", get_value(dealer))
	}
	return dealer, deck
}

func play(players [][]int, dealer []int, deck []int) ([][]int, []int, []int){
	for i := 0; i < len(players); i++ {
		players[i], deck = play_round(players[i], deck)
	}
	x := best_score(players)
	dealer, deck = dealer_play(dealer, deck, x)
	return players, dealer, deck
}

func main() {
	var numPlayers int
	var choice string
	fmt.Println("How many players?")
	fmt.Scan(&numPlayers)
	for i := 0; i < 1; i = i {
		players, dealer, deck := start(numPlayers)
		players, dealer, deck = play(players, dealer, deck)
		winners(players, dealer)
		fmt.Println("would you like to play again?")
		fmt.Scan(&choice)
		if choice == "no" {
			i = 1
		}
	}
}
