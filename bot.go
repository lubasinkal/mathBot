package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/expr-lang/expr"
)

var conversationalResponses = map[string]string{
	"hi":             "hi!",
	"how do you do?": "i'm cool.",
	// ...
}

var variables = map[string]interface{}{}

var sqrtFn = expr.Function("sqrt", func(params ...any) (any, error) {
	return math.Sqrt(params[0].(float64)), nil
}, new(func(float64) float64))

var powFn = expr.Function("pow", func(params ...any) (any, error) {
	return math.Pow(params[0].(float64), params[1].(float64)), nil
}, new(func(float64, float64) float64))

func getBotResponse(userText string) string {
	text := strings.ToLower(strings.TrimSpace(userText))
	if resp, ok := conversationalResponses[text]; ok {
		return resp
	}

	re := regexp.MustCompile(`(?i)(?:(?:what is|calculate|solve)\s+)?(.+)`)
	match := re.FindStringSubmatch(userText)
	exprStr := strings.TrimSpace(match[1])

	if parts := strings.SplitN(exprStr, "=", 2); len(parts) == 2 {
		name := strings.TrimSpace(parts[0])
		valExpr := strings.TrimSpace(parts[1])
		prog, err := expr.Compile(valExpr, sqrtFn, powFn, expr.Env(variables), expr.AllowUndefinedVariables())
		if err != nil {
			return fmt.Sprintf("Invalid expression for %s: %v", name, err)
		}
		result, err := expr.Run(prog, variables)
		if err != nil {
			return fmt.Sprintf("Error evaluating assignment: %v", err)
		}
		variables[name] = result
		return fmt.Sprintf("Assigned: %s = %v", name, result)
	}

	prog, err := expr.Compile(exprStr, sqrtFn, powFn, expr.Env(variables), expr.AllowUndefinedVariables())
	if err != nil {
		return "Sorry, I didn't understand. Try: 'What is 2 + 2?'"
	}
	result, err := expr.Run(prog, variables)
	if err != nil {
		return "Hmm... that doesn't compute. Maybe try simpler math or add variable bindings first."
	}
	return fmt.Sprintf("Answer: %v", result)
}
