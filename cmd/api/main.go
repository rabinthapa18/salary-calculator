// this is the main entry point of the application for API

package main

import (
	"log"
	"net/http"
	"os/exec"
	"salary-calculator/api"
)

func main() {
	// building frontend
	log.Println("Building frontend...")
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = "cmd/frontend"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to build frontend: %s\n%s", err, output)
	}
	log.Println("Frontend built successfully")

	// serving the frontend
	fs := http.FileServer(http.Dir("cmd/frontend/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// setting headers for CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		api.ProcessInputAPI(w, r)
	})
	log.Println("Starting server on :8080. You can visit http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
