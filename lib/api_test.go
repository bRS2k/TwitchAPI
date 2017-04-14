package twitchapi

import "testing"

// TestReadAPIConfig tests readAPIConfig in config/api.go
func TestReadAPIConfig(t *testing.T) {
	// clientID is assigned the result of readAPIConfig in api.go
	if clientID == "" {
		t.Error("Expected nonempty value for clientID")
	}
}
