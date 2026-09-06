package input

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/JayantRautela/req/internal/model"
)

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func readURL(reader *bufio.Reader) (string, error) {
	return readLine(reader, "Enter the URL :- ")
}

func readMethod(reader *bufio.Reader) (string, error) {
	return readLine(reader, "Enter the HTTP method: ")
}

func readBody(reader *bufio.Reader) (string, error) {
	fmt.Println("Enter the JSON body.")
	fmt.Println("Press Enter on an empty line when finished:")

	var lines []string

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			return "", err
		}

		line = strings.TrimRight(line, "\r\n")

		if strings.TrimSpace(line) == "" {
			break
		}

		lines = append(lines, line)

		if err == io.EOF {
			break
		}
	}

	return strings.Join(lines, "\n"), nil
}

func readHeaders(reader *bufio.Reader) (map[string]string, error) {
	fmt.Println("Enter headers in the format 'Key: Value'.")
	fmt.Println("Press Enter on an empty line when finished:")

	headers := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			return nil, err
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid header format. Use 'Key: Value'.")
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		headers[key] = value

		if err == io.EOF {
			break
		}
	}

	return headers, nil
}

func CollectRequest(reader *bufio.Reader) (model.RequestConfig, error) {
	url, err := readURL(reader)
	if err != nil {
		return model.RequestConfig{}, err
	}

	method, err := readMethod(reader)
	if err != nil {
		return model.RequestConfig{}, err
	}

	body, err := readBody(reader)
	if err != nil {
		return model.RequestConfig{}, err
	}

	headers, err := readHeaders(reader)
	if err != nil {
		return model.RequestConfig{}, err
	}

	requestConfig := model.RequestConfig{
		URL:     url,
		Method:  method,
		Body:    body,
		Headers: headers,
	}

	return requestConfig, nil
}
