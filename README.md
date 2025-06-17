# Go AI Text Translator

This is a desktop utility written in Go that allows you to quickly translate selected text using a local Ollama language model.

## Features

*   **Hotkey Activated**: Press `CMD+SHIFT+SPACE` to trigger the translation.
*   **Clipboard Integration**: Automatically copies the selected text and pastes the translation.
*   **Local LLM**: Uses an Ollama model for translation, keeping your data local. The default model is `avanderbergh/gemma3-translator:4b`.
*   **Configurable Prompt**: The translation prompt can be modified (currently set to translate from English to German).

## How it Works

1.  The application runs in the background, listening for the `CMD+SHIFT+SPACE` hotkey.
2.  When the hotkey is detected:
    *   It simulates a "copy" command (`CMD+C`) to get the currently selected text.
    *   The copied text is sent to the specified Ollama model along with a predefined prompt.
    *   The model returns the translated text.
    *   The application then simulates a "paste" command (`CMD+V`) to replace the original selected text with the translation.

## Prerequisites

*   Go installed on your system.
*   Ollama installed and running.
*   The Ollama model `avanderbergh/gemma3-translator:4b` (or your preferred translation model) pulled and accessible by Ollama. You can pull a model using `ollama pull <model_name>`.

## Building and Running

1.  **Clone the repository (if applicable) or ensure you are in the project directory.**
2.  **Install dependencies:**
    ```shell
    go mod tidy
    ```
3.  **Run the application:**
    ```shell
    go run main.go
    ```
    You should see a log message indicating that the hotkey listener has started.

4.  **To build a binary:**
    ```shell
    go build -o go-ai-translator main.go
    ```
    Then you can run the executable:
    ```shell
    ./go-ai-translator
    ```

## Usage

1.  Ensure the application is running.
2.  Select any text in any application.
3.  Press `CMD+SHIFT+SPACE`.
4.  The selected text will be replaced by its German translation.

## Configuration

*   **Translation Model**: The Ollama model used for translation is defined in `internal/llm.go` in the `model` constant.
    ```go
    const model = "avanderbergh/gemma3-translator:4b"
    ```
*   **Translation Prompt**: The prompt instructing the model on how to translate is defined in `internal/prompt.go` in the `Prompt` variable.
    ```go
    var Prompt string = "Translate from English to German: "
    ```
    You can change this to translate to other languages or modify the translation instructions.

## Customizing the Ollama Model

The project includes a `modelfile` located at `model/gemma3-translator-4b.modelfile`. This file defines how the Ollama model behaves, including its base model, parameters, and system prompt.

You can customize this `modelfile` to change the translation behavior, such as:
*   Using a different base model (e.g., another version of Gemma or a different LLM altogether).
*   Adjusting parameters like `temperature` or `top_p` to control the creativity and randomness of the translation.
*   Modifying the `SYSTEM` prompt to change the translation guidelines, target languages, or tone.

**To create or update your local Ollama model using the modelfile:**

1.  Navigate to the directory containing the `modelfile` (i.e., the `model/` directory within this project).
2.  Run the following Ollama command in your terminal:

    ```shell
    ollama create [username]/gemma3-translator:4b -f ./gemma3-translator-4b.modelfile
    ```
    *   Replace `[username]/gemma3-translator:4b` with your desired model name if you want to create a new model or a different tag.
    *   The `-f` flag specifies the path to the modelfile.

This command will build the model according to the specifications in the `modelfile` and make it available for Ollama to use. If the model name already exists, it will be updated.

Make sure that the model name used in the `ollama create` command matches the `model` constant in `internal/llm.go` if you want the application to use your customized version.

## Project Structure

```
.
├── go.mod
├── go.sum
├── main.go         # Main application entry point
├── cmd/
│   └── run.go      # Handles hotkey registration and core application logic
├── internal/
│   ├── llm.go      # Contains the logic for interacting with the Ollama LLM
│   └── prompt.go   # Defines the translation prompt
└── model/
    └── gemma3-translator-4b.modelfile # Modelfile for the default Ollama translator model
```

## Dependencies

*   [github.com/go-vgo/robotgo](https://github.com/go-vgo/robotgo) - For simulating keyboard and mouse events (copy/paste).
*   [github.com/robotn/gohook](https://github.com/robotn/gohook) - For global hotkey listening.
*   [github.com/ollama/ollama/api](https://github.com/ollama/ollama) - For interacting with the Ollama API.

## Notes

*   The application requires accessibility permissions on macOS to simulate keyboard events. If it doesn't work, you might need to grant these permissions in System Settings.
*   The `time.Sleep` calls in `cmd/run.go` are there to give the system time to process clipboard operations. These might need adjustment based on system performance.
