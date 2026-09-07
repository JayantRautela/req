# req

`req` is a simple command-line HTTP client built with Go.

It allows you to interactively create and send HTTP requests directly from the terminal. The tool collects request details such as the URL, HTTP method, request body, and headers, validates the input, sends the request, and displays the response along with the total time taken.

This project was built primarily as a learning exercise to understand Go, HTTP communication, project structure, and separation of responsibilities.

## Features

Currently, `req` supports:

- Interactive URL input
- Multiple HTTP methods
- Multi-line JSON request bodies
- Custom request headers
- Input normalization
- Request validation
- HTTP request execution
- Configurable HTTP client timeout
- Response status display
- Response header display
- Pretty-printed JSON responses
- Plain text response handling
- Total request duration measurement

## Architecture

The application follows a simple pipeline:

```text
User Input
    ↓
RequestConfig
    ↓
Normalization
    ↓
Validation
    ↓
HTTP Request
    ↓
Response
    ↓
Output
```

## Project Structure

```text
req/
├── cmd/
│   └── req/
│       └── main.go
│
├── internal/
│   ├── input/
│   │   └── reader.go
│   │
│   ├── model/
│   │   ├── request.go
│   │   └── response.go
│   │
│   ├── normalizer/
│   │   └── request.go
│   │
│   ├── validator/
│   │   └── request.go
│   │
│   ├── request/
│   │   └── request.go
│   │
│   └── output/
│       └── response.go
│
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

You need Go installed on your machine.

Check your Go installation:

```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/JayantRautela/req.git
cd req
```

### Run the Application

From the project root:

```bash
go run ./cmd/req
```

The CLI will then guide you through creating an HTTP request.

You will be prompted to enter:

1. The request URL
2. The HTTP method
3. The request body
4. Request headers

After the input is collected, the application will:

```text
Normalize Input
      ↓
Validate Request
      ↓
Send HTTP Request
      ↓
Receive Response
      ↓
Display Result
```

## Building the Binary

You can build the application using:

```bash
go build -o req ./cmd/req
```

Then run:

```bash
./req
```