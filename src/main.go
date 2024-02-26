package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("fieldName is required. Usage: ./tippecanone-concat-array <fieldName>")
	}

	fieldName := args[1]

	readWriter := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	replaceInStream(fieldName, readWriter)
}

func replaceInStream(fieldName string, readWriter *bufio.ReadWriter) {
	scanner := bufio.NewScanner(readWriter.Reader)
	for scanner.Scan() {
		line := scanner.Text()
		output := replace(fieldName, line)

		fmt.Fprint(readWriter.Writer, output)
	}

	readWriter.Writer.Flush()
}

func replace(fieldName string, input string) string {
	pattern := fmt.Sprintf("\"%s\": \"(\\d+(?:,\\d+)*)\"", fieldName)
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(input, -1)

	output := input

	for _, groups := range matches {
		match := groups[0]
		values := groups[1]

		split := strings.Split(values, ",")
		distinct := distinct(split)
		new := fmt.Sprintf("\"%s\": \"%s\"", fieldName, strings.Join(distinct, ","))

		output = strings.Replace(output, match, new, 1)
	}

	return output
}

func distinct(values []string) []string {
	var distinctValues []string
	seen := make(map[string]bool)

	for _, value := range values {
		if seen[value] {
			continue
		}

		distinctValues = append(distinctValues, value)
		seen[value] = true
	}

	return distinctValues
}
