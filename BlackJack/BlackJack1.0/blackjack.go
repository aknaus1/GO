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

func make_deck() []int {
  var deck []int
  for i := 1; i < 14; i++ {
    for j := 0; j < 4; j++ {
      deck = append(deck, i)
    }
  }
  return deck
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

func value(a int) string{
  var x string
  if a == 1 {
    x = "A"
  } else if a == 11 {
    x = "J"
  } else if a == 12 {
    x = "Q"
  } else if a == 13 {
    x = "K"
  } else {
    x = string(a)
  }
  return x
}

func print_hand(deck []int) {
  var a []string
  for i := 0; i < len(deck); i++ {
    a = append(a, value(deck[i]))
  }
  fmt.Println(a)
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

func main() {
  var choice string
  win := 0
  loss := 0
  for j := 0; j < 1; j = j {
    fmt.Println("Making a new deck...")
    deck := make_deck()
    fmt.Println("Shuffling deck..")
    deck = shuffle_deck(deck)
    fmt.Println("Okay time to play!")
    var dealer, player []int
    k := 0
    for i := 0; i < 2; i++ {
      player = append(player, deck[k])
      k++
      dealer = append(dealer, deck[k])
      k++
    }
    fmt.Println("Dealer is showing", dealer[0])
    for i := 0; i < 1; i = i {
      fmt.Println("You are at:", get_value(player))
      fmt.Println("Would you like to hit or stay?")
      fmt.Scan(&choice)
      if choice == "hit" {
        player = append(player, deck[k])
        k++
      } else {
        i = 1
      }
      if get_value(player) > 21 {
        i = 1
        fmt.Println("You are at:", get_value(player))
        fmt.Println("You busted! The dealer wins")
        loss++
      }
    }
    if get_value(player) <= 21 {
      for get_value(dealer) < get_value(player) && get_value(dealer) < 17 {
        dealer = append(dealer, deck[k])
        k++
      }
      fmt.Println("Your Hand:", get_value(player))
      fmt.Println("Dealers Hand:", get_value(dealer))
      if get_value(player) > get_value(dealer) || get_value(dealer) > 21{
        fmt.Println("Win: Congrats")
        win++
      } else if get_value(player) == get_value(dealer) {
        fmt.Println("Push: You scored the same of the dealer.")
      } else {
        fmt.Println("Loss: Sorry")
        loss++
      }
    }
    fmt.Println("Would you like to play again?")
    fmt.Scan(&choice)
    if choice == "no" {
      j = 1
    }
  }
  fmt.Println("Wins:",win)
  fmt.Println("Losses:",loss)
}
