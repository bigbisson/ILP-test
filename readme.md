# ILP Technical Test – Golang Backend

This repository contains my technical test submission for **Back End Engineer** position.  
The implementation is written in **Golang**, designed to be clean, modular, and production-ready, including:
- Functional decomposition for each task
- Example test runner
- Automatic result generation script
- CI automation using GitHub Actions

---

## 🚀 Project Structure
- tasks.go # All GO1–GO5 task implementations
- test_runner.go # Executes all task examples
- generate_result.sh # Script to generate technical_test.txt
- technical_test.txt # Output file (generated)
- .github/workflows/ci.yml # GitHub Actions workflow
- go.mod

---

## 🧩 Tasks Summary

### **Task 1 – GO1**
- Detect array category:
  - mixed positive & negative → return 0  
  - only positive → smallest missing positive  
  - only negative → smallest missing negative  

### **Task 2 – GO2**
- Merge integers → reverse digits → remove excluded digits  
- If result empty → return error `"not found"`

### **Task 3 – GO3**
- Find intersection between two arrays  
- Remove duplicates  
- Return sorted descending  

### **Task 4 – GO4**
- Count digit occurrences in a number  
- Return `map[int]int`

### **Task 5 – GO5**
- Concurrent multiplication of index pairs  
- Sum computed results  

---

## 📦 Installation & Usage

### **Run all tasks**

go run .

### **Generate `technical_test.txt`**

./generate_result.sh

The file will appear in the root folder.

---

## ⚙️ GitHub Actions (CI)

This repository includes:
- Build verification (Go compile test)
- Auto-run technical test runner
- Upload generated file as workflow artifact

Workflow file:  
`.github/workflows/ci.yml`

---

## 📝 Author

**Indra Nurwibisono**  
Backend Engineer (Golang, Java)  
GitHub: https://github.com/bigbisson

---

## 📄 License

This project is part of a technical assessment and is not licensed for commercial use.
