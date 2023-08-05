package ratelimiter

import (
	"encoding/json"
	"net/http"
)

// pkg/ratelimiter/ratelimiter.go

type RateLimitRule struct {
	RequestsPerInterval int `json:"requestsPerInterval"`
	Interval            int `json:"interval"`
}

type RateLimitRules map[string]RateLimitRule

type CheckRequest struct {
	ClientID string `json:"clientID"`
}

type CheckResponse struct {
	Allowed bool `json:"allowed"`
}

// pkg/ratelimiter/ratelimiter.go

func checkHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req CheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if a token is available
	allowed, err := isTokenAvailable(req.ClientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	res := CheckResponse{Allowed: allowed}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
