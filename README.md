# Go Learning Repository

A comprehensive guide to learning Go (Golang) from scratch with practical examples and exercises.

## 📚 Curriculum Overview

This repository is structured to take you from beginner to intermediate Go developer through hands-on learning.

### Part 1: Fundamentals (Lessons 1-5)
- [01 - Getting Started](./01_getting_started/)
- [02 - Variables & Data Types](./02_variables_datatypes/)
- [03 - Control Flow](./03_control_flow/)
- [04 - Functions](./04_functions/)
- [05 - Arrays & Slices](./05_arrays_slices/)

### Part 2: Intermediate Concepts (Lessons 6-10)
- [06 - Maps](./06_maps/)
- [07 - Structs & Methods](./07_structs_methods/)
- [08 - Interfaces](./08_interfaces/)
- [09 - Error Handling](./09_error_handling/)
- [10 - Packages & Imports](./10_packages_imports/)

### Part 3: Advanced Topics (Lessons 11-15)
- [11 - Goroutines](./11_goroutines/)
- [12 - Channels](./12_channels/)
- [13 - File I/O](./13_file_io/)
- [14 - JSON & Serialization](./14_json_serialization/)
- [15 - Mini Projects](./15_mini_projects/)

## 🚀 Getting Started

### Prerequisites
- Go 1.21+ installed ([Download](https://golang.org/dl))
- A code editor (VS Code, GoLand, etc.)
- GitHub Codespaces access (optional but recommended)

### Running Examples
Each lesson contains `main.go` files that can be run:

```bash
cd 01_getting_started
go run main.go
```

### Running Tests
Most lessons include test files (`*_test.go`):

```bash
go test ./...
```

### Running Exercises
Practice problems are in `exercises/` folders with solution files marked as `*_solution.go`:

```bash
# Try solving the exercise first!
# Then check the solution
go run exercises/solution_example.go
```

## 💻 Using GitHub Codespaces

1. Click "Code" → "Codespaces" → "Create codespace on main"
2. Once loaded, open the terminal
3. Navigate to any lesson: `cd 01_getting_started`
4. Run examples: `go run main.go`
5. Edit files and experiment!

## 📖 Learning Path

**Week 1-2: Foundations**
- Learn syntax, variables, and control flow
- Write simple programs
- Understand packages and imports

**Week 3-4: Intermediate**
- Master structs and interfaces
- Error handling patterns
- Work with collections (slices, maps)

**Week 5-6: Advanced**
- Concurrency with goroutines and channels
- File I/O operations
- JSON handling for APIs

**Week 7-8: Project Work**
- Build a complete CLI application
- Create a simple web server
- Implement concurrent programs

## 🎯 Learning Tips

1. **Read the comments** - Each example file is heavily commented
2. **Modify and experiment** - Don't just run; tweak the code
3. **Complete exercises** - Try solving without looking at solutions
4. **Read Go docs** - Reference [golang.org](https://golang.org)
5. **Join communities** - Reddit: r/golang, Discord: Gophers Slack

## 📚 Recommended Resources

- [Official Go Tour](https://tour.golang.org) - Interactive introduction
- [Effective Go](https://golang.org/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com) - Example-driven guide
- [Go Standard Library](https://pkg.go.dev) - Documentation

## 🤝 Contributing

Found a mistake or have suggestions? Open an issue or submit a PR!

## 📝 License

This learning material is licensed under MIT License.

---

**Happy Learning! 🎓**

Start with `01_getting_started/` and progress through the lessons sequentially.
