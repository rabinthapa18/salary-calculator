package api

import (
	"encoding/json"
	"net/http"
	"salary-calculator/pkg"
)

func ProcessInputAPI(w http.ResponseWriter, r *http.Request) {

	// restricting to only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var requestBody struct {
		Input string `json:"input"`
	}

	// decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// processing the input
	result, err := pkg.ProcessInput(requestBody.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"result": result,
	}

	json.NewEncoder(w).Encode(response)
}
