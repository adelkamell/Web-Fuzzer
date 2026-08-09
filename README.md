# Mini Web Fuzzer - V1

A lightweight directory enumeration tool for discovering hidden web paths and resources.

## 📌 Overview

This is a minimal web fuzzer that checks for common directory and file paths on a target web server. It's designed to be simple, fast, and easy to understand - perfect for learning the fundamentals of web fuzzing and HTTP probing.

## 🚀 Features

- **Lightweight & Fast**: Uses HTTP HEAD requests to minimize bandwidth usage
- **Simple Implementation**: Clean code structure ideal for beginners
- **Extensible Design**: Built with future enhancements in mind
- **Basic Directory Discovery**: Tests common paths like `/admin`, `/login`, `/config`, etc.

## 🛠️ Installation

```bash
git clone https://github.com/adelkamell/web-fuzzer.git
cd web-fuzzer
go build
```

### 💻 Usage
```bash
./web-fuzzer
```

### Current Behavior:

Targets: http://example.com

Tests these paths: admin, login, config, backup

Outputs: Lists all discovered paths with status code 200

### 📝 Configuration
Currently, the fuzzer uses a hardcoded target and wordlist. To customize:

Change the base variable to your target URL

Modify the words slice with your custom wordlist

### 🔮 Future Features
□ Custom wordlist loading from files
□ Multi-threading support
□ Support for different HTTP methods (GET, POST)
□ Custom headers and cookies
□ Response filtering and validation
□ Export results to file formats (JSON, CSV)
□ Recursive directory scanning
□ Status code filtering

### 🎯 Why Go?
I chose Go for this project because:

Fast Execution: Go's compiled nature makes it excellent for networking tools

Standard Library Power: The net/http package provides everything needed for HTTP operations without external dependencies

Concurrency Support: Built-in goroutines make it easy to add parallel scanning later

Simplicity: Clean syntax makes the code readable and maintainable

Cross-Platform: Single binary deployment works on any system

Performance: Efficient memory management and fast I/O operations

### ⚠️ Legal & Ethical Use
Important: This tool should only be used on systems you own or have explicit permission to test. Unauthorized access to computer systems is illegal in most jurisdictions.

### 🤝 Contributing
Contributions are welcome! Please feel free to submit issues and pull requests.

### 👤 Author
- GitHub: @adelkamell

### Made with **❤️** for the security community