package main

import (
	"fmt"
	"bufio"
  "os"
)

func check(e error) {
    if e != nil {
        panic(e)
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

func sign_in() ([]string, []string, string){
	var user_name, input string
	var response bool
	input_yes, err1 := read_file("default", "yes")
	check(err1)
	input_no, err2 := read_file("default", "no")
	check(err2)
	fmt.Println("Have you played with us before?")
	fmt.Scan(&input)
	response, input_yes, input_no= get_response(input, input_yes, input_no)
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
	return input_yes, input_no, user_name
}

func sign_out(user_name string, input_yes []string, input_no []string) {
	var err error
	var save bool
	var input string
	fmt.Println("Would you like to save your responses?")
	fmt.Scan(&input)
	save, input_yes, input_no = get_response(input, input_yes, input_no)
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
