## Gogeta
Gogeta is a tool that generates valid Go types (structs) from JSON files.

Gogeta - I was just thinking of **Go** **Ge**nerate **Ty**pes -> GoGeTy -> Gogeta. That's how I named it.

### Features
- Generates Go structs from JSON
- Supports nested JSON objects
- Supports arrays of JSON objects
- Generating Go types from JSON files or stdin (WIP)

### Usage

To use Gogeta, simply run the following command:

Use code with caution. Learn more
```
gogeta json [flags]
```

### Flags
* -i, --input: The input JSON filename or stdin. If not provided, Gogeta will read the JSON from stdin.
* -o, --output: The file to write the generated types to. If not provided, Gogeta will print the generated types to stdout.
* -f, --output-to-file: Output the generated type to a file.

### Examples
Generate a Go struct from a JSON file and write to _output.go_ or anywhere you want:
```
gogeta json -i input.json -o output.go
```

Generate Go struct and write to a file: 
```
gogeta json -i input.json -f
```

Generate a Go struct from JSON to stdout:
Simply omit the _-o_ flag
```
gogeta json -i input.json
```

### Installation

Clone the Repository
```
git clone https://github.com/rajenderK7/gogeta.git
cd gogeta
```

Build the CLI Application
```
go build -o bin/gogeta
```


Add to System Environment Variables

Linux / macOS:

```
export PATH=$PATH:/path/to/gogeta/bin
```
To make it permanent, add the above line to your shell profile file (e.g., ~/.bashrc, ~/.zshrc).

Windows:

1. Copy the gogeta.exe from the bin directory to a folder of your choice.
2. Add the folder path to the system's PATH variable:
   - Right-click on "This PC" or "Computer" on the desktop or in File Explorer.
   - Select "Properties" -> "Advanced system settings" -> "Environment Variables..."
   - Under "System variables," find and select the "Path" variable, then click "Edit..."
   - Click "New" and add the path to the folder where gogeta.exe is located.

Verify Installation
```
gogeta --help
```

### Contributing

If you would like to contribute to Gogeta, please open a pull request.
