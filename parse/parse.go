package parse

import (
	"day11/operations"
	"os"
	"strconv"
	"strings"
)

type InputInfo struct {
	Items                []int
	Operation            operations.Operation
	Devider              int
	PositiveReciverIndex int
	NegativeReciverindex int
}

func ParseInput(filepath string) ([]InputInfo, error) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return []InputInfo{}, err
	}
	inputInfos := parseInput(string(fileContent))
	return inputInfos, nil
}

func parseInput(filContent string) []InputInfo {
	monkeyBlocks := strings.Split(filContent, "\r\n\r\n")
	inputs := []InputInfo{}
	for _, monkeyBlock := range monkeyBlocks {
		InputInfo := parsOneMonkey(monkeyBlock)
		inputs = append(inputs, InputInfo)
	}
	return inputs
}

func parsOneMonkey(monkeyText string) InputInfo {
	stdOpr := operations.Add{K: 0}
	input := InputInfo{Items: []int{}, Operation: stdOpr, Devider: 0, PositiveReciverIndex: 0, NegativeReciverindex: 0}
	lines := strings.Split(monkeyText, "\r\n")
	for _, line := range lines {
		if strings.Contains(line, "items") {
			worries := parseItems(line)
			input.Items = worries
		} else if strings.Contains(line, "Test") {
			divider := parseTest(line)
			input.Devider = divider
		} else if strings.Contains(line, "If true") {
			positiveIndex := parseIndex(line)
			input.PositiveReciverIndex = positiveIndex
		} else if strings.Contains(line, "If false") {
			negativeIndex := parseIndex(line)
			input.NegativeReciverindex = negativeIndex
		} else if strings.Contains(line, "Operation") {
			opr := parseOperation(line)
			input.Operation = opr
		}
	}
	return input
}

func parseItems(itemLine string) []int {
	worry := []int{}
	startIntex := strings.Index(itemLine, ":")
	substring := itemLine[startIntex+1:]
	substring = strings.Trim(substring, " ")
	numberStrings := strings.Split(substring, ", ")
	for i := range numberStrings {
		itemWorry, _ := strconv.Atoi(numberStrings[i])
		worry = append(worry, itemWorry)
	}
	return worry
}

func parseTest(testLine string) int {
	return extractLastNumber(testLine)
}

func parseIndex(line string) int {
	return extractLastNumber(line)
}

func extractLastNumber(line string) int {
	splited := strings.Split(line, " ")
	numberString := splited[len(splited)-1]
	ret, _ := strconv.Atoi(numberString)
	return ret
}

func parseOperation(line string) operations.Operation {
	startIntex := strings.Index(line, ":")
	subString := line[startIntex+1:]
	subString = strings.Trim(subString, " ")
	isAdd := strings.Contains(subString, "+")
	indexOfOperand := 0
	indexOfEqual := strings.Index(subString, "=")
	if isAdd {
		indexOfOperand = strings.Index(subString, "+")
	} else {
		indexOfOperand = strings.Index(subString, "*")
	}
	left := strings.Trim(subString[indexOfEqual:indexOfOperand], " ")
	right := strings.Trim(subString[indexOfOperand+1:], " ")
	if !isNumeric(left) && !isNumeric(right) {
		if isAdd {
			return operations.Multiply{K: 2}
		} else {
			return operations.Square{}
		}
	} else {
		number, _ := strconv.Atoi(right)
		if isAdd {
			return operations.Add{K: number}
		} else {
			return operations.Multiply{K: number}
		}
	}
}

func isNumeric(maybeInt string) bool {
	_, err := strconv.ParseInt(maybeInt, 10, 0)
	return err == nil
}
