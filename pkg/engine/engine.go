package engine

import (
	"log"
	"runtime"

	"github.com/philcantcode/goengine/pkg/graphics"
	"github.com/philcantcode/goengine/pkg/input"
)

func Run() {
	// Lock the main thread for OpenGL
	runtime.LockOSThread()

	// Initialize graphics
	err := graphics.Init()
	if err != nil {
		log.Fatalf("Failed to initialize graphics: %v", err)
	}

	// Initialize input
	input.Init(graphics.GetWindow())

	// Start the game loop
	startGameLoop()

	// Cleanup resources
	graphics.Cleanup()
}
