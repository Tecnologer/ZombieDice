package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	lang "github.com/tecnologer/dicegame/language"
)

var (
	lFmt       = lang.GetCurrent()
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

func AskRequiredStringf(format string, values ...interface{}) string {
	return AskRequiredString(lFmt.Sprintf(format, values...))
}

func AskString(msg, defaultValue string) (input string) {
	lFmt.Printf(msg)
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
			lFmt.Printlnf("Invalid input, try again.")
		}
	}

	return
}

func AskBoolf(format string, defaultValue bool, values ...interface{}) (input bool) {
	return AskBool(lFmt.Sprintf(format, values...), defaultValue)
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

func AskInt(msg string, defaultValue int) int {
	inputStr := strings.ToLower(AskString(msg, fmt.Sprint(defaultValue)))
	input, e := strconv.Atoi(inputStr)
	if e != nil {
		return defaultValue
	}

	return input
}

func AskEnter(msg string) {
	_ = AskString(msg, "")
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
