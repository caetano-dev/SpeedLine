package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/drull1000/utils"
	"github.com/Knetic/govaluate"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a mathematical expression, prompt, or URL as a command-line argument.")
		return
	}

	input := strings.Join(os.Args[1:], " ")

	switch {
	case isMathExpression(input):
		evaluateExpression(input)
	case isURL(input):
		openURL(input)
	case isUsername(input):
		searchUsername(input)
	case isUnit(input):
		convertUnits(input)
	default:
		searchPrompt(input)
	}
}

func isMathExpression(input string) bool {
	expressionRegex := regexp.MustCompile(`^[\d+\-*^/()\s]+$`)
	return expressionRegex.MatchString(input)
}

func evaluateExpression(expression string) {
	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, "^", "**")
	evalExpression, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		log.Fatalf("Error evaluating the expression: %v", err)
	}
	result, err := evalExpression.Evaluate(nil)
	if err != nil {
		log.Fatalf("Error evaluating the expression: %v", err)
	}
	fmt.Printf("%v = %v\n", expression, result)
}

func isURL(input string) bool {
	urlRegex := regexp.MustCompile(`^(https?://)?[a-zA-Z0-9-]+\.[a-zA-Z]{2,}(?:/[a-zA-Z0-9-]+)*$`)
	return urlRegex.MatchString(input)
}

func openURL(urlStr string) {
	cmd := exec.Command("firefox", urlStr)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error opening the browser: %v", err)
	}
}

func searchPrompt(prompt string) {
	query := url.QueryEscape(prompt)
	searchURL := fmt.Sprintf("https://www.duckduckgo.com/?q=%s", query)

	cmd := exec.Command("firefox", searchURL)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error opening the browser: %v", err)
	}
}

func isUsername(input string) bool {
	usernameRegex := regexp.MustCompile(`^(@)?[a-zA-Z0-9_]+$`)
	return usernameRegex.MatchString(input)
}

func searchUsername(username string) {
	username = strings.ReplaceAll(username, "@", "")
	query := url.QueryEscape(username)
	searchURL := fmt.Sprintf("https://www.twitter.com/%s", query)
	cmd := exec.Command("firefox", searchURL)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error opening the browser: %v", err)
	}

}

func isUnit(input string) bool {
	conversionRegex := regexp.MustCompile(`([\d.]+)\s*([a-zA-Z/]+)\s+into\s+([a-zA-Z/]+)`)
	return conversionRegex.MatchString(input)
}

func convertUnits(input string) error {
	conversionRegex := regexp.MustCompile(`([\d.]+)\s*([a-zA-Z/]+)\s+into\s+([a-zA-Z/]+)`)
	matches := conversionRegex.FindStringSubmatch(input)
	if len(matches) != 4 {
		return fmt.Errorf("invalid input format")
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return fmt.Errorf("invalid value: %v", err)
	}

	fromUnit := strings.ToLower(matches[2])
	toUnit := strings.ToLower(matches[3])

	if conversionMap, ok := u.ConversionsTable[fromUnit]; ok {
		if conversionFunc, ok := conversionMap[toUnit]; ok {
			result := conversionFunc(value)
			fmt.Printf("%.2f %s", result, toUnit)
			return nil
		}
	}

	return fmt.Errorf("unsupported unit conversion")
}
