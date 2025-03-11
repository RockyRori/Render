package tests

import (
	"render/models"
	"testing"
)

func TestGameModel(t *testing.T) {
	game := models.Game{Title: "Test Game", Price: 29.99}
	if game.Title != "Test Game" {
		t.Errorf("Expected Title to be 'Test Game', got %s", game.Title)
	}
}
