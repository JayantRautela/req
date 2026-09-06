package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JayantRautela/req/internal/input"
	"github.com/JayantRautela/req/internal/normalizer"
	"github.com/JayantRautela/req/internal/validator"
)

func main() {
	fmt.Println("~ req");

	reader := bufio.NewReader(os.Stdin);

	requestConfig, err := input.CollectRequest(reader)

	if err != nil {
		fmt.Println("Error collecting request:", err)
		return
	}

	requestConfig = normalizer.NormalizeRequest(requestConfig)

	err = validator.ValidateRequest(requestConfig)

	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("All validation passed")
}