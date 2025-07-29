package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
)

var conversationalResponses = map[string]string{
	"hi there!":           "hi!",
	"hi":                  "hi!",
	"how do you do?":      "i'm cool.",
	"fine, you?":          "always cool",
	"i'm ok":              "glad to hear that.",
	"i'm fine":            "glad to hear that.",
	"i feel awesome":      "excellent, glad to hear that.",
	"not so good":         "sorry to hear that.",
	"what's your name?":   "i'm MathBot. ask me a math question, please.",
	"pythagorean theorem": "a squared plus b squared equals c squared.",
}

func getBotResponse(userText string) string {
	// Check for conversational responses first
	lowerUserText := strings.ToLower(userText)
	if response, ok := conversationalResponses[lowerUserText]; ok {
		return response
	}

	// Try to extract a mathematical expression from the user's input
	re := regexp.MustCompile(`(?i)(what is|whats|what's|calculate|compute|evaluate|tell me|solve)\s+(.*)`)
	matches := re.FindStringSubmatch(userText)

	var expressionStr string
	if len(matches) > 2 {
		expressionStr = matches[2]
	} else {
		expressionStr = userText
	}

	// Try to evaluate the extracted expression
	expression, err := govaluate.NewEvaluableExpression(expressionStr)
	if err != nil {
		return fmt.Sprintf("Error parsing expression: %v", err)
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return fmt.Sprintf("Error evaluating expression: %v", err)
	}

	return fmt.Sprintf("%v", result)
}
