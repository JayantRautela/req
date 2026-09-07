package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/JayantRautela/req/internal/model"
)

func PrintResponse(response model.Response) {
	fmt.Println()
	fmt.Println("Status:", response.Status)
	fmt.Println("Time Taken:", response.Duration)

	fmt.Println()
	fmt.Println("Headers:")
	printHeaders(response.Headers)

	for key, values := range response.Headers {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	fmt.Println()
	fmt.Println("Body:")
	fmt.Println(formatBody(response.Body))
}

func formatBody(body []byte) string {
	if len(body) == 0 {
		return "(empty)"
	}

	var formattedJSON bytes.Buffer

	err := json.Indent(
		&formattedJSON,
		body,
		"",
		"  ",
	)

	if err == nil {
		return formattedJSON.String()
	}

	return string(body)
}

func printHeaders(headers map[string][]string) {
	keys := make([]string, 0, len(headers))

	for key := range headers {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		for _, value := range headers[key] {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}