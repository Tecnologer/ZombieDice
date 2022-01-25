package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader     = bufio.NewReader(os.Stdin)
	boolValues = map[string]bool{
		"1":     true,
		"t":     true,
		"T":     true,
		"TRUE":  true,
		"true":  true,
		"True":  true,
		"0":     false,
		"f":     false,
		"F":     false,
		"FALSE": false,
		"false": false,
		"False": false,
		"yes":   true,
		"y":     true,
		"si":    true,
		"s":     true,
		"no":    false,
		"n":     false,
	}
)

func AskString(msg, defaultValue string) (input string) {
	fmt.Print(msg)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		input = defaultValue
	}

	return
}

func AskRequiredString(msg string) (input string) {
	for input == "" {
		input = AskString(msg, "")

		if input == "" {
			fmt.Println("Invalid input, try again.")
		}
	}

	return
}

func AskBoolf(format string, defaultValue bool, values ...interface{}) (input bool) {
	return AskBool(fmt.Sprintf(format, values...), defaultValue)
}

func AskBool(msg string, defaultValue bool) (input bool) {
	defaultStr := "no"
	if defaultValue {
		defaultStr = "yes"
	}

	inputStr := strings.ToLower(AskString(msg, defaultStr))
	input = parseInputToBool(inputStr)

	return
}

func parseInputToBool(input string) bool {
	if input == "" {
		return false
	}
	for key, v := range boolValues {
		if strings.EqualFold(input, key) {
			return v
		}
	}

	return false
}
