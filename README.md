# NYAN Directory Structure Viewer 📂

This Go program displays a clean directory structure in a tree view format right in your terminal. It lists each file and subdirectory, indented to show the hierarchy. Emojis distinguish between directories and files for easy visualization.
Example Output

```plaintext
📁./
┣📁 subfolder1
┃ ┣📄 file1.txt
┃ ┗📄 file2.txt
┗📁 subfolder2
  ┗📄 file3.txt
```

## Build

To build this program, you’ll need Go.

`go build`

## Usage

Run the program with the path of the directory you want to display:

`./nyan -path /your/directory`

If no path is provided, it will use the current directory (./) by default.

## Code Overview

### How dirstruct Works

It iterates over the directory’s entries.  
Uses ┣ and ┗ as connectors to build the directory structure visually.  
For subdirectories, it creates a new prefix to indent properly.  

# License

This project is licensed under the CC-BY-SA License.
