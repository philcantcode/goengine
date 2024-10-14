package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// GLFW event handling must run on the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(fmt.Errorf("could not initialize glfw: %w", err))
	}
	defer glfw.Terminate()

	// Set the required OpenGL version and profile
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "OpenGL Example", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create window: %w", err))
	}

	window.MakeContextCurrent()

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	for !window.ShouldClose() {
		// Poll for and process events
		glfw.PollEvents()

		// Swap front and back buffers
		window.SwapBuffers()
	}
}
