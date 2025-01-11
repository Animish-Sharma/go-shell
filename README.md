# Golang Shell: A Custom Command-Line Interface

![Golang Logo](https://miro.medium.com/v2/resize:fit:1400/1*Ifpd_HtDiK9u6h68SZgNuA.png)

A lightweight, fast, and feature-rich custom shell written in Go. This shell supports **12 essential commands**, including `cd`, `ls`, `pwd`, and more. Additionally, it features an **inbuilt text editor** for seamless editing within the shell.

---

## Features

1. **Command Execution**  
   Execute commonly used shell commands like `cd`, `ls`, `pwd`, and more.
2. **Inbuilt Text Editor**  
   Edit text files without leaving the shell environment.
3. **Simple and Lightweight**  
   Designed for performance and minimalism.
4. **Extensible**  
   Easily extendable with additional commands and features.

---

## Supported Commands

| Command   | Description                                                                 |
|-----------|-----------------------------------------------------------------------------|
| `cd`      | Change the current directory.                                              |
| `ls`      | List files and directories in the current directory.                       |
| `pwd`     | Print the current working directory.                                       |
| `touch`   | Create a new empty file.                                                   |
| `mkdir`   | Create a new directory.                                                    |
| `rm`      | Remove files or directories.                                               |
| `cat`     | Display the contents of a file.                                            |
| `echo`    | Print text to the shell or write text to a file.                           |
| `exit`    | Exit the shell.                                                            |
| `help`    | Display a list of available commands.                                      |
| `clear`   | Clear the terminal screen.                                                 |
| `edit`    | Open the inbuilt text editor for creating or editing text files.           |

---

## Installation

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/Animish-Sharma/golang-shell.git
   cd golang-shell
2. **Build the Shell**  
   Ensure that [Go](https://golang.org/dl/) is installed on your system.  
   Then, use the following command to build the executable file:  
   ```bash
   go build -o shell main.go
3. **Run the Shell**  
   After building the shell, execute the following command to start it:  
   ```bash
   ./shell
4. **Add to PATH (Optional)**  
   To make the shell accessible globally, move the executable to a directory included in your system's `PATH`. For example:  
   ```bash
   sudo mv shell /usr/local/bin/
