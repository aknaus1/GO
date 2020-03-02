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

func check_o (key string, input string, o bool) (bool, string){
	if key == input {
		return !o, "yes"
	} else {
		return o, input
	}
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

func get_response(input string, input_yes []string, input_no []string, key string, o bool) (bool, bool, []string, []string) {
	var new_input []string
	for true {
		if input == key {
			return true, !o, input_yes, input_no
		} else if check_array(input, input_yes) {
			input_yes = append(input_yes, new_input...)
			return true, o, input_yes, input_no
		} else if check_array(input, input_no) {
			input_no = append(input_no, new_input...)
			return false, o, input_yes, input_no
		} else {
			fmt.Println("Sorry I didn't quite understand that. Please input yes or no?")
			new_input = append(new_input, input)
			fmt.Scan(&input)
		}
	}
	return false, o, input_yes, input_no
}

func sign_in(key string) ([]string, []string, string, bool){
	var user_name, input string
	var response bool
	o := false
	input_yes, err1 := read_file("default", "yes")
	check(err1)
	input_no, err2 := read_file("default", "no")
	check(err2)
	fmt.Println("Have you played before?")
	fmt.Scan(&input)
	response, o, input_yes, input_no= get_response(input, input_yes, input_no, key, o)
	if response {
		fmt.Println("What is your user name?")
		fmt.Scan(&user_name)
		input_yes, err1 = read_file(user_name, "yes")
		check(err1)
		input_no, err2 = read_file(user_name, "no")
		check(err2)
	} else {
		fmt.Println("Okay no problem.")
	}
	return input_yes, input_no, user_name, o
}

func sign_out(user_name string, input_yes []string, input_no []string, key string, o bool) {
	var err error
	var save bool
	var input string
	fmt.Println("Would you like to save your responses?")
	fmt.Scan(&input)
	save, o, input_yes, input_no = get_response(input, input_yes, input_no, key, o)
	if save {
		if user_name != "" {
			fmt.Println("Saving to", user_name, "...")
			err = write_file(user_name, "yes", input_yes)
			check(err)
			err = write_file(user_name, "no", input_no)
			check(err)
		} else {
			fmt.Println("Okay please enter a user name:")
			fmt.Scan(&user_name)
			fmt.Println("Saving...")
			err = write_file(user_name, "yes", input_yes)
			check(err)
			err = write_file(user_name, "no", input_no)
			check(err)
		}
	} else {
		fmt.Println("Okay your responses wont be saved.")
	}
	fmt.Println("Goodbye!")
}

func play(input_yes []string, input_no []string, o bool) int{
	var max, comp_choice, user_choice int
	fmt.Println("Odds 0 to what?")
	fmt.Scan(&max)
	fmt.Println("Okay enter your number:")
	fmt.Scan(&user_choice)
	comp_choice = rand_num(o, user_choice, max)
	if user_choice == comp_choice {
		fmt.Println("Congratulations you win!")
		fmt.Println("Your Number:", user_choice)
		fmt.Println("Our Number:", comp_choice)
		return 1
	} else {
	fmt.Println("Sorry you lose!")
	fmt.Println("Your Number:", user_choice)
	fmt.Println("Our Number:", comp_choice)
	return 0
	}
}

func main () {
	var play_again bool
	var input string
	win_count := 0
	loss_count := 0
	key := "mertcanselcuk"
	input_yes, input_no, user_name, o := sign_in(key)
	fmt.Println("Lets hop right into it!")
	for i := 0; i < 1; i = i {
		if play(input_yes, input_no, o) == 1 {
			win_count++
		} else {
			loss_count++
		}
		fmt.Println("You have", win_count, "win(s) and", loss_count, "loss(es).")
		fmt.Println("Would you like to play again?")
		fmt.Scan(&input)
		play_again, o, input_yes, input_no = get_response(input, input_yes, input_no, key, o)
		if !play_again {
			fmt.Println("Okay, thanks for playing!")
			i = 1
		} else {
			fmt.Println("Okay, good luck!")
		}
	}
	sign_out(user_name, input_yes, input_no, key, o)
}
