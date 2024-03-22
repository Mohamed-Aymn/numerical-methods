package common

import (
  "strings"
  "math"
  "strconv"
  "fmt"
)

func CalculateN(a float64, b float64, e float64) (float64, error) {
  return math.Round((math.Log2((b-a) / (e)))), nil
}

func OneVariableFunction(function string, variable float64) (float64, error) {
  // replace variables
  expression:= strings.Replace(function, "x", fmt.Sprintf("%.2f", variable), -1)

  tokens := strings.Split(expression, " ")

  var result float64
  var currentOperator string

  for _, token := range tokens {
    if isOperator(token) {
      currentOperator = token
    } else {
      num, err := strconv.ParseFloat(token, 64)
      if err != nil {
        return 0, err
      }
      switch currentOperator {
      case "":
        result = num
      case "+":
        result += num
      case "-":
        result -= num
      case "*":
        result *= num
      case "/":
        result /= num
      case "^":
        result = math.Pow(result, num)
    }
    }
  }
  return result, nil
}

func isOperator(token string) bool {
  return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

