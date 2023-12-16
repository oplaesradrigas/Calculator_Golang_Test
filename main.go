package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RomanNumsStruct struct {
	integer_value int
	roman_number  string
}

var romanNumsSlice = []RomanNumsStruct{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var roman_nums_to_integers_map = map[string]int{"I": 1, "II": 2, "III": 3,
	"IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

func intToRoman(value_to_convert int) string {
	var result string = ""

	for _, slice_element := range romanNumsSlice {
		for value_to_convert >= slice_element.integer_value {
			result += slice_element.roman_number
			value_to_convert -= slice_element.integer_value
		}
	}

	return result
}

func calculate(the_sign string, first_num int, second_num int) (int, error) {

	switch the_sign {
	case "+":
		return first_num + second_num, nil
	case "-":
		return first_num - second_num, nil
	case "*":
		return first_num * second_num, nil
	case "/":
		return first_num / second_num, nil
	case "":
		err := errors.New("ERROR: undefined behavior in <func calculate>")
		return 0, err
	default:
		err := errors.New("ERROR: undefined behavior in <func calculate>")
		return 0, err
	}
}

func check_all_arab_nums(operation string) bool {

	for i := 0; i < 11; i++ {
		if strings.Contains(operation, fmt.Sprint(i)) {
			return true
		}
	}

	return false
}

func check_all_roman_nums(operation string) bool {

	for key := range roman_nums_to_integers_map {
		if strings.Contains(operation, key) {
			return true
		}
	}

	return false
}

func remove_symbols(str string) string {
	// Define a regular expression pattern to match symbols other than "+", "-", "*", "/"
	pattern := "[^+\\-*/]"

	// Create a regular expression object
	reg := regexp.MustCompile(pattern)

	// Remove symbols using the regular expression
	result := reg.ReplaceAllString(str, "")

	return result
}

func douplicate_exception(operation string) bool {

	arithemtic_ops_only := remove_symbols(operation)
	operations := strings.Split(arithemtic_ops_only, "")
	var signs_arr = [4]string{"+", "-", "*", "/"}

	counter := 0

	for _, op := range operations {
		for _, sign := range signs_arr {
			if op == sign {
				counter++
			}
			if counter > 1 {
				return true
			}
		}
	}

	return false
}

func no_sign_exception(operation string) bool {

	arithemtic_ops_only := remove_symbols(operation)
	operations := strings.Split(arithemtic_ops_only, "")
	var signs_arr = [4]string{"+", "-", "*", "/"}

	counter := 0

	for _, op := range operations {
		for _, sign := range signs_arr {
			if op == sign {
				counter++
			}
		}
	}

	return counter == 0
}

func search_for_sign(operation string) (string, error) {

	var signs_arr = [4]string{"+", "-", "*", "/"}

	if no_sign_exception(operation) {
		err := errors.New("ERROR: string is not a mathematical operation")
		return "", err
	}

	if douplicate_exception(operation) {
		err := errors.New("ERROR: you have entered more than one arithmetic operation or the operation is not supported")
		return "", err
	}

	for _, sign := range signs_arr {
		if strings.Contains(operation, sign) {
			return sign, nil
		}
	}

	err := errors.New("ERROR: this operation is not supported")
	return "", err
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter the operation to calculate: ")

		text, _ := reader.ReadString('\n') // Waiting for data in String format
		// text := "1 + 1"
		clear_text := strings.TrimSpace(text) // Clearing the string (removing whitespaces/tabs)
		operation := strings.ReplaceAll(clear_text, " ", "")

		if operation == "stop" {
			fmt.Println("The program is forcibly terminated")
			break
		}

		the_sign, err := search_for_sign(operation)
		if err != nil {
			fmt.Println(err)
			continue
		}

		the_nums := strings.Split(operation, the_sign)

		if len(the_nums) != 2 {
			err := errors.New("ERROR: The calculator supports operations with only two numbers")
			fmt.Println(err)
			continue
		}

		if check_all_roman_nums(operation) {
			if check_all_arab_nums(operation) {
				err := errors.New("ERROR: You have entered Arabic and Roman numerals in one operation")
				fmt.Println(err)
				continue

			} else {

				first_num := roman_nums_to_integers_map[the_nums[0]]
				second_num := roman_nums_to_integers_map[the_nums[1]]

				if first_num > 10 || second_num > 10 || first_num == 0 || second_num == 0 {
					err := errors.New("ERROR: The calculator supports operations only with numbers from I to X")
					fmt.Println(err)
					continue
				}

				result, err := calculate(the_sign, first_num, second_num)

				if result <= 0 {
					err := errors.New("ERROR: The result of an operation with Roman numerals cannot be zero or less than zero")
					fmt.Println(err)
					continue
				}

				if err != nil {
					fmt.Println(err)
					continue
				} else {
					fmt.Println(intToRoman(result))
				}
			}

		} else if check_all_arab_nums(operation) {

			first_num, err1 := strconv.Atoi(the_nums[0])
			second_num, err2 := strconv.Atoi(the_nums[1])

			if first_num > 10 || second_num > 10 || first_num < 1 || second_num < 1 {
				err := errors.New("ERROR: The calculator supports operations only with numbers from 1 to 10")
				fmt.Println(err)
				continue
			}

			if err1 != nil && err2 != nil {
				fmt.Println("Error", err1, err2)
				continue
			}

			result, err := calculate(the_sign, first_num, second_num)

			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println(result)
			}

		}

	}

}
