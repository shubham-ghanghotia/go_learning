# Lesson 14: JSON & Serialization

## What is JSON?

JSON (JavaScript Object Notation) is a human-readable data format widely used in APIs and data storage.

```json
{
  "name": "Alice",
  "age": 30,
  "city": "New York"
}
```

## Struct Tags

Tags tell the JSON encoder how to handle struct fields:

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}
```

## Marshaling (Go → JSON)

Convert Go struct to JSON string:

```go
person := Person{Name: "Alice", Age: 30, City: "New York"}
jsonData, err := json.Marshal(person)
fmt.Println(string(jsonData))
// Output: {"name":"Alice","age":30,"city":"New York"}
```

## Unmarshaling (JSON → Go)

Convert JSON string to Go struct:

```go
jsonStr := `{"name":"Bob","age":25,"city":"SF"}`
var person Person
err := json.Unmarshal([]byte(jsonStr), &person)
fmt.Printf("%+v\n", person)
```

## Pretty JSON

Format JSON with indentation:

```go
jsonData, err := json.MarshalIndent(person, "", "  ")
fmt.Println(string(jsonData))
```

## Tag Options

```go
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email,omitempty"`  // Omit if empty
    Password string `json:"-"`                // Ignore completely
    Admin    bool   `json:"admin"`
}
```

### Common Tags

| Tag | Meaning |
|-----|---------|
| `json:"field"` | Map to JSON field |
| `json:"field,omitempty"` | Omit if zero value |
| `json:"-"` | Ignore field |
| `json:"field,string"` | Convert to/from string |

## Reading JSON from File

```go
import "encoding/json"

data, _ := ioutil.ReadFile("data.json")
var people []Person
json.Unmarshal(data, &people)
```

## Writing JSON to File

```go
jsonData, _ := json.MarshalIndent(person, "", "  ")
ioutil.WriteFile("output.json", jsonData, 0644)
```

## JSON Array Handling

```go
type People struct {
    Users []Person `json:"users"`
}

// Or directly with a slice
var users []Person
json.Unmarshal(jsonData, &users)
```

## Custom Marshal/Unmarshal

Implement custom JSON handling:

```go
func (p Person) MarshalJSON() ([]byte, error) {
    // Custom marshal logic
}

func (p *Person) UnmarshalJSON(data []byte) error {
    // Custom unmarshal logic
}
```

## Running the Example

```bash
cd 14_json_serialization
go run main.go
```

## 🎯 Exercises

### Exercise 1: Person to JSON
Create Person struct and convert to/from JSON.

### Exercise 2: API Response
Parse JSON API response into Go structs.

### Exercise 3: Configuration file
Read JSON config file and use values in program.

## Next Steps

→ [15 - Mini Projects](../15_mini_projects/)

## Resources

- [encoding/json Package](https://pkg.go.dev/encoding/json)
- [JSON and Go](https://golang.org/blog/json)
