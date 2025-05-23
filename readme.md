# 🧩 JSON Parser in Go

A lightweight JSON parser implemented in Go, built from scratch without using `encoding/json`. This project is created as part of the [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-json-parser/) to incrementally build a JSON parser using regex and basic parsing logic.

---
## ✅ Features

| Step | Feature |
|------|---------|
| 1 | Parse `{}` and detect invalid JSON |
| 2 | Support string key-value pairs (`{"key": "value"}`) |
| 3 | Support booleans, null, numbers, and typed values |
| 4 | Support empty arrays and objects as values |

---
## 🚀 Getting Started

### ✅ Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

### 🔧 Build

```bash
git clone https://github.com/hiabhi-cpu/ccwc.git
go build -o json-parser
```

This creates an executable named json-parser.

---
## 🧪 Usage
### ▶️ Run

```bash
./json-parser <path-to-json-file>
```

### 📥 Example
```bash
./json-parser tests/step4/valid.json
```

### ✅ Sample Output
For Valid JSON
```bash
✅ Valid JSON with typed values:
{
  "key": value (string)
  "key-n": 101 (float64)
  "key-o": {
  }
  "key-l": [
  ]
}
```

For InValid json
```bash
❌ Invalid JSON
```

## 🧪 Run Tests

```bash
go test -v
```