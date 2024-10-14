package engine

import (
	"github.com/philcantcode/goengine/pkg/graphics"
	"github.com/philcantcode/goengine/pkg/input"
)

func startGameLoop() {
	for !graphics.WindowShouldClose() {
		// Process input
		input.Update()

		// Update game state (placeholder)
		// Game logic would go here

		// Render frame
		graphics.Render()

		// Poll for events
		graphics.PollEvents()
	}
}
