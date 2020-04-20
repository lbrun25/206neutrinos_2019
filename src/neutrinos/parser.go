package neutrinos

import (
	"fmt"
	"os"
	"regexp"
	"utils"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBePositiveInteger = "must be a positive integer.\n"
	mustBePositiveFloat = "must be a positive float.\n"
	mustBeGreatherThanZero = "must be greater than zero.\n"
	theValueMustBeGreatherThanZero = "The value must be greater than zero.\n"
	wrongNextValueFormat = "Wrong next value's format which must be a positive integer."
	maxArg = 4
	minArg = 4
)

// NumberValues - number of values
var NumberValues int = 0

// ArithmeticMean - arithmetic mean
var ArithmeticMean float64 = 0.0

// HarmonicMean - harmonic mean
var HarmonicMean float64 = 0.0

// StandardDeviation - standard deviation
var StandardDeviation float64 = 0.0

func printError(valueName string, errorMessage string) {
	fmt.Printf("Error: '%s' %s\n", valueName, errorMessage)
}

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if (arg == "-h") {
			return true
		}
	}
	return false
}

// CheckNextValueFormat - check next value that is entered in the input
func CheckNextValueFormat(input string) bool {
	var re = regexp.MustCompile("[0-9]")

	match := re.FindString(input)
	if len(match) != len(input) {
		fmt.Println(wrongNextValueFormat)
		return false
	}
	resInt := utils.ConvertStringToInt(input)
	if (resInt == 0) {
		fmt.Printf(mustBeGreatherThanZero)
		return false
	}
	return true;
}

func getIntegerPositiveValueGreaterThanZero(valueName string, arg string) (bool, int) {
	if !utils.IsPositiveInteger(arg) {
		printError(valueName, mustBePositiveInteger)
		return false, -1
	}
	integer := utils.ConvertStringToInt(arg)
	if (integer <= 0) {
		printError(valueName, mustBeGreatherThanZero)
		return false, -1
	}
	return true, integer
}

func getFloatPositiveValueGreaterThanZero(valueName string, arg string) (bool, float64) {
	if !utils.IsPositiveFloat(arg) {
		printError(valueName, mustBePositiveFloat)
		return false, -1
	}
	float := utils.ConvertStringToFloat(arg)
	if (float <= 0) {
		printError(valueName, mustBeGreatherThanZero)
		return false, -1
	}
	return true, float
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < minArg {
		fmt.Println(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > maxArg {
		fmt.Println(tooManyArgs)
		return false
	}
	valueNames := [4]string{"n", "a", "h", "sd"}
	for i, arg := range argsWithoutProg {
		valueName := valueNames[i]

		// Check and assign n
		if valueName == "n" {
			status, integer := getIntegerPositiveValueGreaterThanZero(valueName, arg)
			if (!status) {
				return false;
			}
			NumberValues = integer
		}

		// Check and assign a
		if valueName == "a" {
			status, float := getFloatPositiveValueGreaterThanZero(valueName, arg)
			if (!status) {
				return false;
			}
			ArithmeticMean = float
		}

		// Check and assign h
		if valueName == "h" {
			status, float := getFloatPositiveValueGreaterThanZero(valueName, arg)
			if (!status) {
				return false;
			}
			HarmonicMean = float
		}

		// Check and assign sd
		if valueName == "sd" {
			status, float := getFloatPositiveValueGreaterThanZero(valueName, arg)
			if (!status) {
				return false;
			}
			StandardDeviation = float
		}
	}
	return true
}