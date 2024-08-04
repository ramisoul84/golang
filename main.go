package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var romanNumeralMap = []struct {
	Value int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
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

// remove spaces and new line char and convert to Upper case..
func removeSpaces(input string)string{
	res := strings.TrimSpace(input)
	res = strings.ReplaceAll(res," ","")
	res = strings.ToUpper(res)
	return res
}

// Number of Math operators and the last operator
func MathOperators(input string) (int,rune) {
	mathOperators := "+-*/"
 	operator := ' '
 	sum := 0
	for _, char := range input {
		if strings.ContainsRune(mathOperators, char) {
		sum+=1
      	operator = char
		}
	}
	return sum,operator
}

// a function that split input to operands
func splitToOperands(input string,operator rune)[]string{
	return strings.Split(input,string(operator))
	
}

// a Function returns the type of operands
func checkType(input string) string {
	isNumeric := regexp.MustCompile(`^[0-9]+$`).MatchString(input)
	if isNumeric {
		return "DECIMAL"
	}
	isValidRoman := func(s string) bool {
		for i := len(s) - 1; i >= 0; i-- {
			_, found := romanNumerals[rune(s[i])]
			if !found {
				return false
			}
	
		}
		return true
	}
	if isValidRoman(input) {
		return "ROMAN"
	}

	return "STRING"
}

func romanToDecimal(s string) int {
	result := 0
	previousValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		currentValue :=  romanNumerals[rune(s[i])]
		if currentValue < previousValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		previousValue = currentValue
	
	}
	return result
}

func intToRoman(num int) string {
	romanNumeral := ""

	for _, mapping := range romanNumeralMap {
		for num >= mapping.Value {
			romanNumeral += mapping.Symbol
			num -= mapping.Value
		}
	}

	return romanNumeral
}

func moreTen(n int)bool{
	if n > 10 || n < 1{
		return true
	} else{
		return false
	}
}

func result(n1,n2 int, op rune) int{
	var result int
	switch op {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		if n2 ==0 {
			panic("integer divide by zero")
		}
		result = n1 / n2
	}	
	return result
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a mathematical expression (e.g., 1 + 3): ")
	input, _ := reader.ReadString('\n')
	input  = removeSpaces((input))

	sum,op := MathOperators((input))
	if sum == 0 {
		panic("строка не является математической операцией.")	
	}
	if sum > 1 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")	
	}
	operands := splitToOperands(input,op)
	if len(operands) != 2  || operands[0] =="" || operands[1] =="" {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	for i := 0; i < len(operands); i++ {
		if checkType(operands[i]) == "STRING"{
			panic("операнд не является числом")
		}
	}

	num1 := checkType(operands[0])
	num2 := checkType(operands[1])
	var res string
	if num1 != num2 {
		panic("используются одновременно разные системы счисления")
	}else{
		if num1 == "ROMAN"{
			number1D := romanToDecimal(operands[0])
			number2D := romanToDecimal(operands[1])	
			if moreTen(number1D) || moreTen(number2D){
				panic("числа только от 1 до 10 включительно")
			}
			x := result(number1D,number2D,op)
			if x<1{
				panic("в римской системе нет отрицательных чисел.")
			}
			roman :=intToRoman(x)
			res = roman


		}else{
			number1D, _ := strconv.Atoi(operands[0])
			number2D, _ := strconv.Atoi(operands[1])
			if moreTen(number1D) || moreTen(number2D){
				panic("числа только от 1 до 10 включительно")
			}
			
			res =strconv.Itoa(result(number1D,number2D,op) )
		}
	}








		
	fmt.Println()
	fmt.Println("INPUT")
	fmt.Println(operands[0],string(op),operands[1])
	fmt.Println("OUTPUT")
	fmt.Println(res)
}