package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
  "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func check_array (a string, list []string) bool {
	for i:= 0; i < len(list); i++ {
		if a == list[i] {
			return true
		}
	}
	return false
}

func read_file (fname string, selector string) ([]string, error){
	loc := fname + selector + ".txt"
	file, err := os.Open(loc)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func write_file (fname string, selector string, lines []string) error{
	loc := fname + selector + ".txt"
	file, err := os.Create(loc)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func rand_num(a bool, num int, max int) int{
	if a {
		return num
	} else {
		a1 := rand.NewSource(time.Now().UnixNano())
		a2 := rand.New(a1)
		comp_choice := a2.Intn(max)
		return comp_choice
	}
}

func main () {
	var max int
	var comp_choice int
	var user_choice int
	var play_again string
	var new_input []string
	var save string
	var user_name string
	var err1, err2 error
	o := false
	found := true
	win_count := 0
	loss_count := 0
	key := "mertcanselcuk"
	input_yes := []string{"yes", "y", "yea", "sure", "yeh"}
	input_no := []string{"no", "n", "nope", "nah", "nay"}
	fmt.Println("Have you played before?")
	fmt.Scan(&play_again)
	if play_again == key {
		o = true
		play_again = "yes"
	}
	for i := 0; i < 1; i = i{
		if check_array(play_again, input_yes) {
			fmt.Println("What is your user name?")
			fmt.Scan(&user_name)
			if found == true {
				input_yes, err1 = read_file(user_name, "yes")
				check(err1)
				input_no, err2 = read_file(user_name, "no")
				check(err2)
				i = 1
			} else {
				fmt.Println("Sorry we could not find you. Would you like to try a different name?")
				fmt.Scan(&play_again)
			}
		} else if check_array(play_again, input_no) {
			fmt.Println("Okay no problem.")
			i = 1
		} else {
			fmt.Println("Sorry I didn't quite understand that. Please input yes or no?")
			if play_again == key {
				o = !o
				play_again = "yes"
			}
			fmt.Scan(&play_again)
		}
	}
	fmt.Println("Lets hop right into it!")
	for i := 0; i < 1; i = i {
		fmt.Println("Odds 0 to what?")
		fmt.Scan(&max)
		fmt.Println("Okay enter your number.")
		fmt.Scan(&user_choice)
		comp_choice = rand_num(o, user_choice, max)
		if user_choice == comp_choice {
			fmt.Println("Congratulations you win!")
			fmt.Println("Your Number: ", user_choice)
			fmt.Println("Our Number: ", comp_choice)
			win_count++
		} else {
		fmt.Println("Sorry you lose!")
		fmt.Println("Your Number: ", user_choice)
		fmt.Println("Our Number: ", comp_choice)
		loss_count++
		}
		fmt.Println("You have", win_count, "win(s) and", loss_count, "loss(es).")
		fmt.Println("Would you like to play again?")
		for j := 0; j < 1; j = j {
			fmt.Scan(&play_again)
			if play_again == key {
				o = !o
				play_again = "yes"
			}
			if check_array(play_again, input_yes) {
				fmt.Println("Okay, good luck!")
				j = 1
			} else if check_array(play_again, input_no) {
				fmt.Println("Okay, thanks for playing!")
				j = 1
				i = 1
			} else {
				for k := 0; k < 1; k = k {
					fmt.Println("Sorry I didn't quite understand that. Please input yes or no?")
					new_input = append(new_input, play_again)
					fmt.Scan(&play_again)
					if play_again == key {
						o = !o
						play_again = "yes"
					}
					if play_again == "yes" || check_array(play_again, input_yes) {
						fmt.Println("Okay I'll make sure to take a note of that.")
						input_yes = append(input_yes, new_input...)
						new_input = nil
						fmt.Println("Good luck!")
						k = 1
						j = 1
					} else if play_again == "no" || check_array(play_again, input_no) {
						fmt.Println("Okay I'll make sure to take a note of that.")
						input_no = append(input_no, new_input...)
						new_input = nil
						fmt.Println("Thanks for playing!")
						k = 1
						j = 1
						i = 1
					}
				}
			}
		}
	}
	fmt.Println(input_yes)
	fmt.Println(input_no)
	if user_name != "" {
		fmt.Println("Saving to ", user_name, "...")
		err3 := write_file(user_name, "yes", input_yes)
		check(err3)
		err4 := write_file(user_name, "no", input_no)
		check(err4)
		fmt.Println("Goodbye!")
	} else {
		fmt.Println("Would you like to save your responses?")
		fmt.Scan(&save)
		for i := 0; i < 1; i = i {
			if check_array(save, input_yes) {
				fmt.Println("Okay please enter a user name.")
				fmt.Scan(&user_name)
				fmt.Println("Saving...")
				err3 := write_file(user_name, "yes", input_yes)
				check(err3)
				err4 := write_file(user_name, "no", input_no)
				check(err4)
				fmt.Println("Goodbye!")
				i = 1
			} else if check_array(save, input_no){
				fmt.Println("Okay your responses wont be saved.")
				fmt.Println("Goodbye!")
				i = 1
			}	else {
				fmt.Println("Sorry I didn't quite understand that. Please input yes or no?")
				fmt.Scan(&play_again)
			}
		}
	}
}
