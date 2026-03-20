# Word-it

Word-it is a simple web-based utility tool written in Go. It offers a unique Shakespearean translator alongside a standard English dictionary, making it a handy companion for writers, students, or anyone looking to add a touch of the Bard to their modern prose.

## Features

- **Shakespearean Translator**: Convert modern English sentences into their Early Modern English equivalents using a custom-built mapping of common terms (e.g., "you" to "thou", "are" to "art").
- **English Dictionary**: Search for definitions of English words using a reliable external API.
- **Clean UI**: Built with [Pico CSS](https://picocss.com/) for a minimal, semantic, and responsive user experience.
- **Server-Side Rendering**: Powered by Go's `html/template` for fast and efficient page delivery.

## Prerequisites

- [Go](https://go.dev/doc/install) 1.16 or higher installed on your machine.

## Getting Started

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/word-it.git
    cd word-it
    ```

2.  **Run the application**:
    ```bash
    go run main.go
    ```

3.  **Access the app**:
    Open your web browser and navigate to `http://localhost:8080`.

## Docker Usage

If you have [Docker](https://www.docker.com/) installed, you can easily containerize and run the application:

1.  **Build the Docker image**:
    ```bash
    docker build -t word-it .
    ```

2.  **Run the Docker container**:
    ```bash
    docker run -p 8080:8080 word-it
    ```

3.  **Access the app**:
    Navigate to `http://localhost:8080` in your browser.

## Project Structure

- `main.go`: The core application logic, including the HTTP server, routing, Shakespearean translation logic, and dictionary API integration.
- `index.html`: The frontend template utilizing Pico CSS and Go's templating engine.
- `template/`: (If applicable) Additional template files.
- `.qodo/`: Configuration for Qodo agents and workflows.

## How it Works

### Shakespearean Translation
The translator uses a local lookup table to replace modern English words with their Shakespearean counterparts. It handles common pronouns, verbs, and descriptive adjectives to give your text a classic feel.

### Dictionary Search
The dictionary feature fetches real-time data from the [Free Dictionary API](https://dictionaryapi.dev/), providing multiple definitions for any valid English word.

