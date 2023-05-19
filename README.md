# SpeedLine

SpeedLine is a versatile command-line tool that enhances your productivity by providing efficient mathematical expression evaluation, prompt-based searches, URL opening, unit conversion, and Twitter handle lookup capabilities. With SpeedLine, you can effortlessly compute mathematical expressions, perform quick web searches, open URLs directly from the command line, convert between different units, and look up Twitter profiles, all without leaving your terminal. Experience streamlined and convenient command-line operations with SpeedLine to boost your efficiency and simplify your workflow.

![SpeedLine gif](.github/SpeedLine.gif)

## Features

- **Math Expression Evaluation**: Evaluate mathematical expressions with speed and accuracy using SpeedLine's built-in expression evaluator.
- **Prompt-Based Searches**: Conduct prompt-based searches directly from the command line, fetching relevant results from the web.
- **Seamless Browser Integration**: Seamlessly open URLs within your default browser to quickly access web content related to your searches.
- **Unit Conversion**: Perform unit conversions effortlessly with SpeedLine. Convert between different units, such as time, speed, and data size, with ease and precision.
- **Twitter Handle Lookup**: Look up Twitter profiles directly from the command line. Search for usernames and retrieve relevant information in an instant. All you need is to type the @ of who you are looking for.

The default browser is Firefox and the default search engine is DuckDuckGo.

## Usage

To use SpeedLine, simply enter the desired input as a command-line argument. The program will determine if it is a mathematical expression, a prompt for web search, URL, etc. The corresponding action will be executed, providing you with the results you need quickly.

## Installation

To install SpeedLine, follow these steps:

1. Clone the repository:

```git clone https://github.com/drull1000/SpeedLine.git```

2. Navigate to the project directory:

```cd SpeedLine```

3. Install the dependencies:

```go mod tidy```

4. Build the binary:

```go build -o SpeedLine```

5. Add the binary to your system's PATH for convenience. You can also bind a specific key to run the program (if you see the gif, you'll see that mine is the letter "s").

## Contributing

Contributions are welcome! If you have any ideas, enhancements, or bug fixes, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [GPLv3 License](LICENSE).