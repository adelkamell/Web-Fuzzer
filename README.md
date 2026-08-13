# 🔥 Web-Fuzzer - Advanced Directory Bruteforcer

[![Go Version](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go)](https://go.dev)
[![CI Status](https://img.shields.io/github/actions/workflow/status/adelkamell/Web-Fuzzer/go-test.yml?branch=main&style=for-the-badge&logo=githubactions)](https://github.com/adelkamell/Web-Fuzzer/actions)
[![License](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/adelkamell/web-fuzzer?style=for-the-badge&logo=docker)](https://hub.docker.com/r/adelkamell/web-fuzzer)

## 🚀 What is Web-Fuzzer?
A high-performance **Go-based web path fuzzer** designed for security researchers to discover hidden directories and files with multithreading support.

## 📦 Installation
### Using Go:
```bash
go install github.com/adelkamell/Web-Fuzzer/cmd/web-fuzzer@latest
```

### Using Docker (Recommended):
```bash
docker run -it adelkamell/web-fuzzer http://example.com wordlist.txt
```

### 🛠 Usage Example
```bash
web-fuzzer http://target.com /path/to/wordlist.txt
```
Sample Output:

```text
[FOUND] http://target.com/admin -> Status: 200
[FOUND] http://target.com/api -> Status: 403
```

### 🧪 Running Tests
```bash
go test -v -cover ./...
```

### 🏗 Architecture
/cmd: Entrypoint

/pkg: Reusable core logic (Fuzzer engine)

/internal: Private utilities

### 🤝 Contributing
Contributions are welcome! Here's how you can help:

Fork the repository

Create your feature branch (git checkout -b feature/AmazingFeature)

Commit your changes (git commit -m 'Add some AmazingFeature')

Push to the branch (git push origin feature/AmazingFeature)

Open a Pull Request

### 👤 Author
- Adel Kamell

- GitHub: @adelkamell

### 🙏 Acknowledgments
Thanks to the Go community for excellent tools and libraries

Inspired by tools like dirb, gobuster, ffuf, and dirsearch

Special thanks to security researchers and the open-source community

### **Made with ❤️ for the security community**
