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

	switch {
	case fromUnit == "hours" && toUnit == "days":
		result := value / 24
		fmt.Printf("%.2f days", result)
		return nil
	case fromUnit == "hours" && toUnit == "seconds":
		result := value * 3600
		fmt.Printf("%.2f seconds", result)
		return nil
	case fromUnit == "hours" && toUnit == "minutes":
		result := value * 60
		fmt.Printf("%.2f minutes", result)
		return nil
	case fromUnit == "days" && toUnit == "hours":
		result := value * 24
		fmt.Printf("%.2f hours", result)
		return nil
	case fromUnit == "seconds" && toUnit == "hours":
		result := value / 3600
		fmt.Printf("%.2f hours", result)
		return nil
	case fromUnit == "minutes" && toUnit == "hours":
		result := value / 60
		fmt.Printf("%.2f hours", result)
		return nil
	case fromUnit == "km/h" && toUnit == "m/s":
		result := value * 0.277778
		fmt.Printf("%.2f m/s", result)
		return nil
	case fromUnit == "m/s" && toUnit == "km/h":
		result := value * 3.6
		fmt.Printf("%.2f km/h", result)
		return nil
	case fromUnit == "bytes" && toUnit == "kb":
		result := value / 1024
		fmt.Printf("%.2f KB", result)
		return nil
	case fromUnit == "bytes" && toUnit == "mb":
		result := value / 1024 / 1024
		fmt.Printf("%.2f MB", result)
		return nil
	case fromUnit == "bytes" && toUnit == "gb":
		result := value / 1024 / 1024 / 1024
		fmt.Printf("%.2f GB", result)
		return nil
	case fromUnit == "bytes" && toUnit == "tb":
		result := value / 1024 / 1024 / 1024 / 1024
		fmt.Printf("%.2f TB", result)
		return nil
	case fromUnit == "kb" && toUnit == "bytes":
		result := value * 1024
		fmt.Printf("%.2f bytes", result)
		return nil
	case fromUnit == "mb" && toUnit == "bytes":
		result := value * 1024 * 1024
		fmt.Printf("%.2f bytes", result)
		return nil
	case fromUnit == "gb" && toUnit == "bytes":
		result := value * 1024 * 1024 * 1024
		fmt.Printf("%.2f bytes", result)
		return nil
	case fromUnit == "tb" && toUnit == "bytes":
		result := value * 1024 * 1024 * 1024 * 1024
		fmt.Printf("%.2f bytes", result)
		return nil
	default:
		return fmt.Errorf("unsupported unit conversion")
	}
}
