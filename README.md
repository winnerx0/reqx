# reqx

reqx is a command-line HTTP client that executes requests defined in a YAML configuration file. It simplifies managing and running multiple HTTP requests by storing them as code.

## Features

- **YAML Configuration**: Define HTTP requests in valid YAML files.
- **Batch Execution**: Run multiple requests in sequence from a single file.
- **Custom Headers**: Support for setting custom request headers (Authorization, Content-Type, etc.).
- **JSON Body Support**: Easily define request bodies as structured data.

## Requirements

- Go 1.25 or higher

## Installation

Ensure you have Go installed on your system.

```bash
go install github.com/winnerx0/reqx@latest
```

Or clone the repository and build from source:

```bash
git clone https://github.com/winnerx0/reqx.git
cd reqx
go build -o reqx
```

## Usage

The basic command to send requests is `send`. `request.yaml` is the default file used if no file is specified.

```bash
reqx send [file]
```

### Example

```bash
# Run using default request.yaml
reqx send

# Run using a specific configuration file
reqx send my-requests.yaml
```

## Supported HTTP Methods

The following HTTP methods are supported:

- GET
- POST
- PUT
- DELETE
- PATCH
- HEAD
- OPTIONS

## Configuration

Requests are defined in a YAML file using the following structure.

### Structure

- **requests**: A list of request objects.
  - **name**: A descriptive name for the request (printed in output).
  - **url**: The full target URL.
  - **method**: The HTTP method (e.g., GET, POST).
  - **headers**: A map of header key-value pairs (optional).
  - **body**: The request body (automatically converted to JSON) (optional).

### Example `request.yaml`

```yaml
requests:
  - name: "Get User"
    url: "https://api.example.com/users/1"
    method: "GET"
    headers:
      Authorization: "Bearer token"
      Content-Type: "application/json"

  - name: "Create Post"
    url: "https://api.example.com/posts"
    method: "POST"
    headers:
      Content-Type: "application/json"
    body:
      title: "Hello World"
      content: "This is a test post"
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
