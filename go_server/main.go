package main

import (
	"encoding/json"
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the /api endpoint

	

    // Register the handler for the /api endpoint
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Create a dummy response as a Go struct
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Hello, this is a dummy response from the /api endpoint!",
		}

		// Convert the response struct to JSON
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			return
		}

		// Set the response headers and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

    // Start the HTTP server on port 8080
    fmt.Println("Server is running on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
