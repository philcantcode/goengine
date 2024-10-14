// pkg/input/input.go
package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var window *glfw.Window

func Init(win *glfw.Window) {
	window = win
	// Set up input callbacks if needed
	// For now, we're not handling any input
}

func Update() {
	// Placeholder for input processing
	// For example, you can check if the ESC key was pressed:
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
