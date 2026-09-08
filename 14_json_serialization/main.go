package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Example 1: Basic struct with JSON tags
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Example 2: Struct with more complex tags
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	Active   bool   `json:"active"`
	CreatedAt string `json:"created_at"`
}

// Example 3: Nested structs
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	ZIP    string `json:"zip"`
}

type Employee struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Salary  float64 `json:"salary"`
	Address Address `json:"address"`
}

// Example 4: Array/Slice JSON
type Company struct {
	Name      string     `json:"name"`
	Employees []Employee `json:"employees"`
}

func main() {
	fmt.Println("=== JSON & SERIALIZATION ===\n")

	// Example 1: Basic Marshal (Go → JSON)
	fmt.Println("1. Basic Marshaling (Struct to JSON):")
	person := Person{
		Name:    "Alice",
		Age:     30,
		City:    "New York",
		Country: "USA",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("JSON: %s\n\n", string(jsonData))

	// Example 2: Pretty JSON (indented)
	fmt.Println("2. Pretty JSON (Indented):")
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Pretty JSON:\n%s\n\n", string(prettyJSON))

	// Example 3: Unmarshal (JSON → Struct)
	fmt.Println("3. Unmarshaling (JSON to Struct):")
	jsonStr := `{"name":"Bob","age":25,"city":"San Francisco","country":"USA"}`
	var person2 Person

	err = json.Unmarshal([]byte(jsonStr), &person2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Parsed Person: %+v\n\n", person2)

	// Example 4: Struct with tag options (omitempty, -)
	fmt.Println("4. Advanced Tags (omitempty, -, etc):")
	user1 := User{
		ID:        1,
		Username:  "alice",
		Email:     "alice@example.com",
		Password:  "secret123",
		Active:    true,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	user2 := User{
		ID:       2,
		Username: "bob",
		Password: "secret456",
		Active:   false,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	user1JSON, _ := json.MarshalIndent(user1, "", "  ")
	user2JSON, _ := json.MarshalIndent(user2, "", "  ")

	fmt.Println("User 1 (with email):")
	fmt.Println(string(user1JSON))
	fmt.Println("\nUser 2 (without email - omitempty):")
	fmt.Println(string(user2JSON))
	fmt.Println()

	// Example 5: Nested structures
	fmt.Println("5. Nested Structures:")
	employee := Employee{
		ID:     1,
		Name:   "Charlie",
		Salary: 75000,
		Address: Address{
			Street: "123 Main St",
			City:   "Boston",
			ZIP:    "02101",
		},
	}

	empJSON, _ := json.MarshalIndent(employee, "", "  ")
	fmt.Println(string(empJSON))
	fmt.Println()

	// Example 6: Arrays of objects
	fmt.Println("6. Array of Objects:")
	company := Company{
		Name: "Tech Corp",
		Employees: []Employee{
			{
				ID:     1,
				Name:   "Alice",
				Salary: 80000,
				Address: Address{
					Street: "456 Oak Ave",
					City:   "New York",
					ZIP:    "10001",
				},
			},
			{
				ID:     2,
				Name:   "Bob",
				Salary: 75000,
				Address: Address{
					Street: "789 Pine Rd",
					City:   "San Francisco",
					ZIP:    "94102",
				},
			},
		},
	}

	companyJSON, _ := json.MarshalIndent(company, "", "  ")
	fmt.Println(string(companyJSON))
	fmt.Println()

	// Example 7: Unmarshal array
	fmt.Println("7. Unmarshal Array:")
	jsonArray := `[{"name":"Dave","age":28,"city":"Seattle","country":"USA"},{"name":"Eve","age":32,"city":"Portland","country":"USA"}]`
	var people []Person

	err = json.Unmarshal([]byte(jsonArray), &people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Parsed %d people:\n", len(people))
	for i, p := range people {
		fmt.Printf("%d. %+v\n", i+1, p)
	}
	fmt.Println()

	// Example 8: Map for dynamic JSON
	fmt.Println("8. Using Maps for Dynamic JSON:")
	data := map[string]interface{}{
		"name":  "Frank",
		"age":   35,
		"email": "frank@example.com",
		"tags":  []string{"developer", "golang"},
	}

	dynamicJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(dynamicJSON))
	fmt.Println()

	// Example 9: Save to file
	fmt.Println("9. Save JSON to File:")
	filename := "person.json"
	ioutil.WriteFile(filename, companyJSON, 0644)
	fmt.Printf("JSON saved to '%s'\n", filename)

	// Example 10: Read from file
	fmt.Println("\n10. Read JSON from File:")
	fileData, _ := ioutil.ReadFile(filename)
	var loadedCompany Company
	json.Unmarshal(fileData, &loadedCompany)
	fmt.Printf("Loaded company: %+v\n", loadedCompany.Name)
	fmt.Printf("Number of employees: %d\n", len(loadedCompany.Employees))
}
