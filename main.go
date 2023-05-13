package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a mathematical expression, prompt, or URL as a command-line argument.")
		return
	}

	input := strings.Join(os.Args[1:], " ")

	if isMathExpression(input) {
		evaluateExpression(input)
	} else if isURL(input) {
		openURL(input)
	} else {
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
