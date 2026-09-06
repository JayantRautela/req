package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JayantRautela/req/internal/input"
)

func main() {
	fmt.Println("~ req");

	reader := bufio.NewReader(os.Stdin);

	requestConfig, err := input.CollectRequest(reader)

	if err != nil {
		fmt.Println("Error collecting request:", err)
		return
	}

	fmt.Println()
	fmt.Println("Request collected successfully!")
	fmt.Println()

	fmt.Println("URL:", requestConfig.URL)
	fmt.Println("Method:", requestConfig.Method)
	fmt.Println("Body:")
	fmt.Println(requestConfig.Body)

	fmt.Println("Headers:")
	for key, value := range requestConfig.Headers {
		fmt.Printf("%s: %s\n", key, value)
	}
}