# Lesson 15: Mini Projects

This lesson brings together everything you've learned with complete, practical projects.

## Project 1: Command-Line Todo App

A simple todo application with file persistence.

**Features:**
- Add, delete, mark complete
- Save to JSON file
- Load todos on startup

**Key Concepts:**
- File I/O
- JSON serialization
- Struct methods
- Error handling

## Project 2: Weather App (API Client)

Fetch weather data from an API and display it.

**Features:**
- HTTP requests
- JSON parsing
- Error handling
- Formatted output

**Key Concepts:**
- net/http package
- Unmarshaling JSON
- Interface usage
- Error handling

## Project 3: Concurrent Downloader

Download multiple files concurrently.

**Features:**
- Goroutines for parallel downloads
- Channels for coordination
- Progress tracking
- Error recovery

**Key Concepts:**
- Goroutines
- Channels
- WaitGroup
- File I/O

## Project 4: Simple Web Server

Build a basic HTTP server.

**Features:**
- Handle HTTP requests
- Route different endpoints
- Serve JSON responses
- Log requests

**Key Concepts:**
- net/http package
- Handler functions
- Routing
- JSON responses

## Project 5: Chat Application

Simple chat system with multiple clients.

**Features:**
- Multiple goroutines
- Channel communication
- Broadcasting messages
- Client management

**Key Concepts:**
- Goroutines
- Channels
- Select statement
- Concurrency patterns

## Getting Started

Choose any project and implement it! Each project teaches different aspects of Go.

### Example: Todo Project Structure

```
todo-app/
├── main.go
├── todo.go (Todo struct and methods)
├── storage.go (File operations)
└── todos.json (Data file)
```

## Challenge Projects

Once you've mastered the basics:

1. **Build a REST API** with multiple endpoints
2. **Create a game** using concurrent goroutines
3. **Implement a cache** with mutex synchronization
4. **Build a crawler** that scrapes websites
5. **Create a worker pool** for parallel processing

## Resources

- [Go Blog - Real-world Go](https://golang.org/blog)
- [Go by Example](https://gobyexample.com) - See full examples
- [Go Playground](https://play.golang.org) - Write and run code online

## Tips for Success

1. **Start small** - Implement basic features first
2. **Test frequently** - Run your code often
3. **Read errors carefully** - They point you to problems
4. **Use goroutines wisely** - Not everything needs concurrency
5. **Handle errors** - Always check error returns

## What's Next?

- Advanced concurrency patterns
- Dependency injection
- Testing and benchmarking
- Building CLI tools with flags
- Database integration
- Web frameworks (Gin, Echo)

## Congratulations!

You've completed the Go Learning Course! You now understand:
- ✅ Fundamentals (variables, control flow, functions)
- ✅ Data structures (arrays, slices, maps, structs)
- ✅ OOP concepts (methods, interfaces)
- ✅ Error handling
- ✅ Concurrency (goroutines, channels)
- ✅ File I/O and JSON
- ✅ Building real applications

**Keep learning and building!** 🎉
