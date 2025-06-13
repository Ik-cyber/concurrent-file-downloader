# 🚀 Concurrent File Downloader in Go

A fast, minimal, terminal-based concurrent file downloader written in Go.  
Supports multi-file downloads with individual progress bars and automatic retries.

---

## 📦 Features

- Concurrent file downloads
- Progress bars for each file
- Automatic retries on failure
- Clean, modular project structure
- Easy CLI integration (future milestone)

---

## 🛠️ Installation

### Option 1: Manual Build (Recommended for Local Use)

```bash
git clone https://github.com/Ik-cyber/concurrent-file-downloader
cd concurrent-file-downloader
go build -o downloader
sudo mv downloader /usr/local/bin/
```

### Option 2: Use Go Install

```bash
go install github.com/Ik-cyber/concurrent-file-downloader
```

## Project structure

```text
.
├── cmd/            # CLI commands (for future milestone)
├── downloader/     # Core downloader logic
├── utils/          # Utility functions
├── progress/       # Progress bar management
├── main.go         # Application entry point
├── queue.txt       # List of files to download
├── TODO.md         # Project milestones
└── README.md       # Project guide
```

## 📜 License

This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for more details.
