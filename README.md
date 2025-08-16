# Alexa OpenAI Go

This project integrates **Amazon Alexa** with the **OpenAI (GPT) API** using **Go (Golang)** and the **Gin** framework.  
It allows Alexa to send spoken intents to your own server, which processes the request, retrieves a GPT-generated response, and returns it to Alexa in **SSML** format.

---

## Features

-   Receives HTTPS requests from Alexa via **Custom Skill**
-   Validates Alexa's signature and certificate (Amazon's security requirement)
-   Integrates with **OpenAI's GPT API** for intelligent responses
-   Formats responses in **SSML** for natural Alexa speech
-   Can be publicly exposed via **Cloudflare Tunnel** with valid HTTPS

---

## Project Structure

```
alexa-ai/
├── clients/
│   └── openai.go          # OpenAI API client
├── tmp/
│   └── main               # Compiled binary
├── .gitignore             # Git ignore patterns
├── go.mod                 # Go module dependencies
├── go.sum                 # Go dependency checksums
├── env.exemple            # Environment variables template
├── main.go                # Server initialization and main logic
└── README.md              # This file
```

---

## How It Works

1. **User speaks to Alexa** → Amazon Alexa Cloud
2. **Alexa sends HTTPS request** → Your public server (via Cloudflare Tunnel)
3. **Server validates** request signature & certificate
4. **Intent is extracted** from Alexa's JSON payload
5. **Request is sent** to OpenAI GPT API
6. **GPT response is formatted** into SSML
7. **Response is returned** to Alexa → Alexa speaks it

---

## 🔧 Prerequisites

### Install Go

Make sure you have Go installed on your system:

#### macOS (using Homebrew)

```bash
brew install go
```

#### macOS/Linux (using official installer)

```bash
# Download and install from https://golang.org/dl/
# Or use the following commands:
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

#### Windows

Download and install from [https://golang.org/dl/](https://golang.org/dl/)

#### Verify Installation

```bash
go version
```

---

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/your-username/alexa-openai-go.git
cd alexa-openai-go
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Configure environment variables

Create a `.env` file based on the template:

```bash
cp env.exemple .env
```

Edit the `.env` file with your actual values:

```ini
PORT=3000
ALEXA_APP_ID=your-alexa-skill-id
OPENAI_API_KEY=your-openai-api-key
DISABLE_ALEXA_VERIFY=false
```

### 4. Run the server

#### Development mode (with auto-reload)

```bash
go run main.go
```

#### Build and run

```bash
# Build the binary
go build -o tmp/main main.go

# Run the compiled binary
./tmp/main
```

#### Run with specific environment file

```bash
go run main.go -env .env
```

The server will run at `http://localhost:3000`.

### 5. Expose server with Cloudflare Tunnel

You can expose your local server publicly using Cloudflare Tunnel:

1. **Install Cloudflare CLI (cloudflared)**

    ```bash
    # macOS
    brew install cloudflared
    ```

2. **Authenticate with Cloudflare**

    ```bash
    cloudflared tunnel login
    ```

3. **Run the tunnel**

    ```bash
    cloudflared tunnel --url http://localhost:3000
    ```

4. **Copy the generated HTTPS URL** and set it as your Alexa skill endpoint.

### 6. Configure your Alexa Skill

In the Alexa Developer Console, set your endpoint to the Cloudflare Tunnel URL.

You can find more about cloudflared [here](https://github.com/cloudflare/cloudflared).

---

## Testing

### Local Testing

```bash
# Health check
curl -X GET http://localhost:3000/health

# Test with sample Alexa request (if endpoint is available)
curl -X POST http://localhost:3000/alexa \
  -H "Content-Type: application/json" \
  -d '{"request": {"type": "IntentRequest", "intent": {"name": "HelloWorldIntent"}}}'
```

### Alexa Simulator

Use the Alexa Developer Console's simulator to test responses.

### Physical Device

Speak to Alexa and validate live responses.

---

## Development

### Hot Reload (Development)

For development with auto-reload, you can use:

```bash
# Install air for hot reloading (optional)
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Build for Production

```bash
# Build for current platform
go build -o alexa-ai main.go

# Build for Linux (if deploying to server)
GOOS=linux GOARCH=amd64 go build -o alexa-ai-linux main.go

# Build for multiple platforms
make build-all  # if you have a Makefile
```

---

## Security Notes

-   **Never keep `DISABLE_ALEXA_VERIFY=true` in production**
-   **Keep your `OPENAI_API_KEY` secret** (do not commit `.env` to version control)
-   **Use environment variables** for all sensitive data
-   **Limit request size and processing time** to avoid high API costs
-   **Only expose `/health` and `/alexa` endpoints** publicly
-   **Use HTTPS in production** (Alexa requires it)

---

## Learning Goals

This project is designed for learning Go, API integration, and Alexa skill development. It covers:

-   **HTTP server creation** with Gin framework
-   **JSON parsing and struct mapping** in Go
-   **Request verification & HTTPS security**
-   **Calling external APIs** in Go
-   **SSML response formatting** for Alexa
-   **Public server exposure** using Cloudflare Tunnel
-   **Environment configuration** and secrets management
-   **Error handling** and logging in Go
