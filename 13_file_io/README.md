# Lesson 13: File I/O

## Reading Files

### Using ioutil.ReadFile()

```go
import "io/ioutil"

data, err := ioutil.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

### Using os.Open()

```go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

## Writing Files

### Using ioutil.WriteFile()

```go
data := []byte("Hello, World!")
err := ioutil.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

### Using os.Create()

```go
file, err := os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

file.WriteString("Hello, World!")
```

## File Operations

### Check if file exists

```go
if _, err := os.Stat("file.txt"); err == nil {
    fmt.Println("File exists")
} else if os.IsNotExist(err) {
    fmt.Println("File not found")
}
```

### Append to file

```go
file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

file.WriteString("\nNew line")
```

### Delete file

```go
err := os.Remove("file.txt")
if err != nil {
    log.Fatal(err)
}
```

## Directories

### Create directory

```go
os.Mkdir("newdir", 0755)
os.MkdirAll("path/to/dir", 0755)
```

### List directory

```go
entries, err := ioutil.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name())
}
```

## Reading CSV Files

```go
import "encoding/csv"

file, _ := os.Open("data.csv")
reader := csv.NewReader(file)
records, _ := reader.ReadAll()

for _, record := range records {
    fmt.Println(record)
}
```

## Running the Example

```bash
cd 13_file_io
go run main.go
```

## Next Steps

→ [14 - JSON & Serialization](../14_json_serialization/)

## Resources

- [os Package](https://pkg.go.dev/os)
- [ioutil Package](https://pkg.go.dev/io/ioutil)
- [bufio Package](https://pkg.go.dev/bufio)
