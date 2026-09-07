package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/JayantRautela/req/internal/header"
	"github.com/JayantRautela/req/internal/validator"
)

func ReadLine(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func CollectHeaders(
	reader *bufio.Reader,
) (map[string]string, error) {
	headers := make(map[string]string)

	commonHeaders := header.CommonHeaders()

	for {
		printHeaderMenu(commonHeaders)

		fmt.Print("\nChoose an option: ")

		choice, err := readChoice(reader)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		customHeaderChoice := len(commonHeaders) + 1
		doneChoice := len(commonHeaders) + 2

		switch {
		case choice >= 1 && choice <= len(commonHeaders):
			selectedHeader := commonHeaders[choice-1]

			headerName, canAdd, err := canAddHeader(
				reader,
				headers,
				selectedHeader.Name,
			)
			if err != nil {
				return nil, err
			}

			if !canAdd {
				continue
			}

			value, err := selectHeaderValue(
				reader,
				selectedHeader,
			)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			headers[headerName] = value

			fmt.Printf(
				"Added header: %s: %s\n",
				headerName,
				value,
			)

		case choice == customHeaderChoice:
			name, err := readCustomHeaderName(reader)
			if err != nil {
				return nil, err
			}

			headerName, canAdd, err := canAddHeader(
				reader,
				headers,
				name,
			)
			if err != nil {
				return nil, err
			}

			if !canAdd {
				continue
			}

			value, err := readCustomHeaderValue(reader)
			if err != nil {
				return nil, err
			}

			headers[headerName] = value

			fmt.Printf(
				"Added header: %s: %s\n",
				headerName,
				value,
			)

		case choice == doneChoice:
			return headers, nil

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func printHeaderMenu(headers []header.Definition) {
	fmt.Println("\nSelect a header:")

	for index, definition := range headers {
		fmt.Printf("%d. %s\n", index+1, definition.Name)
	}

	fmt.Printf("%d. Custom Header\n", len(headers)+1)
	fmt.Printf("%d. Done\n", len(headers)+2)
}

func selectHeaderValue(reader *bufio.Reader, definition header.Definition) (string, error) {
	if len(definition.Values) == 0 {
		fmt.Printf("Enter value for %s: ", definition.Name)

		value, err := ReadLine(reader)
		if err != nil {
			return "", err
		}

		return value, nil
	}

	fmt.Printf("\nSelect a value for %s:\n", definition.Name)

	for index, value := range definition.Values {
		fmt.Printf("%d. %s\n", index+1, value)
	}

	customValueChoice := len(definition.Values) + 1

	fmt.Printf("%d. Custom value\n", customValueChoice)

	fmt.Print("\nChoose an option: ")

	choice, err := readChoice(reader)
	if err != nil {
		return "", err
	}

	switch {
	case choice >= 1 && choice <= len(definition.Values):
		return definition.Values[choice-1], nil

	case choice == customValueChoice:
		fmt.Printf("Enter value for %s: ", definition.Name)

		value, err := ReadLine(reader)
		if err != nil {
			return "", err
		}

		return value, nil

	default:
		return "", fmt.Errorf("invalid option")
	}
}

func readChoice(reader *bufio.Reader) (int, error) {
	input, err := ReadLine(reader)
	if err != nil {
		return 0, err
	}

	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}

	return choice, nil
}

func confirmReplacement(reader *bufio.Reader, headerName string, currentValue string) (bool, error) {
	fmt.Printf(
		"\n%s is already set to: %s\n",
		headerName,
		currentValue,
	)

	for {
		fmt.Print("Do you want to replace it? (y/n): ")

		answer, err := ReadLine(reader)
		if err != nil {
			return false, err
		}

		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "y", "yes":
			return true, nil

		case "n", "no":
			return false, nil

		default:
			fmt.Println("Invalid input. Please enter y or n.")
		}
	}
}

func canAddHeader(reader *bufio.Reader, headers map[string]string, name string) (string, bool, error) {
	existingName, exists := findHeader(headers, name)

	if !exists {
		return name, true, nil
	}

	currentValue := headers[existingName]

	replace, err := confirmReplacement(
		reader,
		existingName,
		currentValue,
	)
	if err != nil {
		return "", false, err
	}

	if !replace {
		return "", false, nil
	}

	return existingName, true, nil
}

func readCustomHeaderValue(reader *bufio.Reader) (string, error) {
	for {
		fmt.Print("Enter header value: ")

		value, err := ReadLine(reader)
		if err != nil {
			return "", err
		}

		if err := validator.ValidateHeaderValue(value); err != nil {
			fmt.Printf("Invalid header value: %v\n", err)
			continue
		}

		return value, nil
	}
}

func readCustomHeaderName(reader *bufio.Reader) (string, error) {
	for {
		fmt.Print("Enter header name: ")

		name, err := ReadLine(reader)
		if err != nil {
			return "", err
		}

		err = validator.ValidateHeaderName(name)
		if err != nil {
			fmt.Printf(
				"Invalid header name: %v\n",
				err,
			)
			continue
		}

		return name, nil
	}
}

func findHeader(headers map[string]string, name string) (string, bool) {
	for existingName := range headers {
		if strings.EqualFold(existingName, name) {
			return existingName, true
		}
	}

	return "", false
}
