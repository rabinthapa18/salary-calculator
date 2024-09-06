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
	var requestBody []struct {
		Input string `json:"input"`
	}

	// decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// checking if the input is empty
	if len(requestBody) == 0 {
		http.Error(w, "Empty input", http.StatusBadRequest)
		return
	}

	// results of all of the inputs
	var results []int

	// processing all of the inputs and storing the results
	for _, input := range requestBody {
		result, err := pkg.ProcessInput(input.Input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		results = append(results, result)
	}

	response := map[string]interface{}{
		"results": results,
	}

	json.NewEncoder(w).Encode(response)
}
