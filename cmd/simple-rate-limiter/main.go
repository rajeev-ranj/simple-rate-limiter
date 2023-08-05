package simple_rate_limiter

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func loadRateLimitRules(filename string) (RateLimitRules, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rules RateLimitRules
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&rules)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

// cmd/rate-limiter/main.go

func main() {
	rules, err := loadRateLimitRules("ratelimitrules.json")
	if err != nil {
		log.Fatalf("Failed to load rate limit rules: %v", err)
	}

	// Initialize the rate limiter with the rules
	// ...
	// Register the /check endpoint
	http.HandleFunc("/check", checkHandler)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
