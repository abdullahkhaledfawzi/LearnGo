package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
===============================================
Stage 7: Core Standard Library Packages
Part 2: IO, OS, and JSON
===============================================

- `io`: Core interfaces for reading and writing data streams.
- `os`: Interacting with the underlying operating system (files, env vars).
- `encoding/json`: Essential for web APIs to serialize/deserialize data.
*/

/*
===============================================
os Package (Operating System)
===============================================
*/

func osDemo() {
	fmt.Println("======== os Package ========")

	// 1. Command Line Arguments
	fmt.Println("\n--- Command Line Arguments ---")
	fmt.Printf("Program executed path: %s\n", os.Args[0])
	// To pass arguments: `go run main.go arg1 arg2` (Available via os.Args[1:])

	// 2. Environment Variables
	fmt.Println("\n--- Environment Variables ---")
	path := os.Getenv("PATH")
	if len(path) > 50 {
		fmt.Printf("PATH (first 50 chars): %s...\n", path[:50])
	}

	// Setting an env variable dynamically
	os.Setenv("MY_APP_MODE", "development")
	fmt.Printf("MY_APP_MODE: %s\n", os.Getenv("MY_APP_MODE"))

	// 3. Current Working Directory
	fmt.Println("\n--- Working Directory ---")
	wd, _ := os.Getwd()
	fmt.Printf("Current Directory: %s\n", wd)

	// 4. File Stat (Checking file information)
	fmt.Println("\n--- File Info ---")
	info, err := os.Stat("go.mod")
	if err == nil {
		fmt.Printf("File '%s' size: %d bytes\n", info.Name(), info.Size())
	} else {
		fmt.Println("go.mod not found here.")
	}

	// 5. Exiting the program
	// os.Exit(1) // Immediately terminates the program with status code 1. Defers are NOT run!
}

/*
===============================================
io Package (Reader/Writer Interfaces)
===============================================
*/

func ioDemo() {
	fmt.Println("\n======== io Package ========")

	// The two most important interfaces in Go:

	// type Reader interface {
	//     Read(p []byte) (n int, err error)
	// }

	// type Writer interface {
	//     Write(p []byte) (n int, err error)
	// }

	// Implementations include: os.File, bytes.Buffer, net.Conn, http.ResponseWriter
	fmt.Println("io.Reader and io.Writer are the backbone of all Go I/O streams.")
}

/*
===============================================
encoding/json Package
===============================================

⚠️ Crucial Concepts (Guaranteed MCQ):
1. Marshal: Converts Go structs to JSON (Bytes).
2. Unmarshal: Converts JSON (Bytes) to Go structs.
3. Struct fields MUST BE EXPORTED (Uppercase) for the JSON package to see them!
4. Struct Tags (`json:"name"`) control how fields are mapped.
*/

// User struct with JSON Tags
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age,omitempty"` // omitempty: Excludes field from JSON if it has the zero value
	Password string `json:"-"`             // - : Completely ignores this field in JSON
}

func jsonMarshalDemo() {
	fmt.Println("\n======== JSON Marshal (Struct -> JSON) ========")

	user := User{
		ID:       1,
		Name:     "Ahmed",
		Email:    "ahmed@example.com",
		Age:      0,             // Because of `omitempty`, this will NOT appear in the JSON
		Password: "supersecret", // Because of `-`, this will NOT appear in the JSON
	}

	// json.Marshal returns []byte, not string
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Printf("Raw JSON bytes: %s\n", jsonBytes)

	// MarshalIndent adds newlines and spacing for readability
	jsonPretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:")
	fmt.Println(string(jsonPretty))
}

func jsonUnmarshalDemo() {
	fmt.Println("\n======== JSON Unmarshal (JSON -> Struct) ========")

	jsonStr := `{
		"id": 2,
		"name": "Fatima",
		"email": "fatima@example.com",
		"age": 23
	}`

	var user User
	// ⚠️ MUST pass a pointer to the struct so Unmarshal can modify it!
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	fmt.Printf("Struct populated: %+v\n", user)
}

/*
===============================================
JSON with Collections
===============================================
*/

func jsonSlicesAndMaps() {
	fmt.Println("\n======== JSON with Slices and Maps ========")

	// Slice of structs becomes a JSON Array
	users := []User{
		{1, "Ahmed", "ahmed@example.com", 25, ""},
		{2, "Fatima", "fatima@example.com", 23, ""},
	}

	jsonBytes, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println("Slice JSON Array:")
	fmt.Println(string(jsonBytes))

	// Go Maps become JSON Objects
	data := map[string]interface{}{
		"role":   "admin",
		"active": true,
		"score":  95.5,
	}

	jsonMap, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println("\nMap JSON Object:")
	fmt.Println(string(jsonMap))
}

/*
===============================================
Generic JSON Parsing
===============================================

⚠️ When you receive a JSON payload but don't know its structure, parse it into `interface{}`.
Go will map it to `map[string]interface{}`.
*/

func jsonGeneric() {
	fmt.Println("\n======== Generic JSON Parsing ========")

	jsonStr := `{
		"product": "Laptop",
		"price": 1200.50,
		"in_stock": true,
		"tags": ["electronics", "computers"]
	}`

	var data interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	// Type Assert the interface to a map
	if m, ok := data.(map[string]interface{}); ok {
		for key, value := range m {
			fmt.Printf("%s: %v (Type: %T)\n", key, value, value)
		}
	}
	// Note: JSON numbers are always parsed as `float64` by default when parsing into interface{}!
}

/*
===============================================
JSON String Tag Trick
===============================================
Sometimes APIs send numbers as strings (`"price": "99.99"`). You can tell Go to convert them automatically.
*/

type Item struct {
	ID    int     `json:"id"`
	Price float64 `json:"price,string"` // Parses string numbers into float64 directly!
}

func jsonTagsDemo() {
	fmt.Println("\n======== JSON String Tags ========")

	jsonStr := `{"id": 101, "price": "149.99"}`
	var item Item
	
	json.Unmarshal([]byte(jsonStr), &item)
	fmt.Printf("Item unmarshaled: ID=%d, Price=%.2f (Type: %T)\n", item.ID, item.Price, item.Price)
}


