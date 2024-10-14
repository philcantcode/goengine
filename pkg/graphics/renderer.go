package graphics

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/philcantcode/goengine/pkg/graphics/geometry"
	"github.com/philcantcode/goengine/pkg/graphics/shaders"
)

var window *glfw.Window
var program uint32

func Init() error {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("failed to initialize GLFW: %v", err)
	}

	// Set GLFW window hints for OpenGL version and profile
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True) // Necessary for macOS

	var err error
	window, err = glfw.CreateWindow(800, 600, "My Game Engine", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create GLFW window: %v", err)
	}

	window.MakeContextCurrent()

	// Initialize Glow (OpenGL bindings)
	if err := gl.Init(); err != nil {
		return fmt.Errorf("failed to initialize OpenGL bindings: %v", err)
	}

	// Print OpenGL version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version:", version)

	// Set up viewport size
	gl.Viewport(0, 0, 800, 600)

	// Enable depth testing
	gl.Enable(gl.DEPTH_TEST)

	// Initialize shaders
	var shaderErr error
	program, shaderErr = shaders.InitShaders()
	if shaderErr != nil {
		return fmt.Errorf("failed to initialize shaders: %v", shaderErr)
	}

	// Initialize geometry
	geometry.InitGeometry()

	return nil
}

func Render() {
	// Clear the color and depth buffers
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Use the shader program
	gl.UseProgram(program)

	// Draw the geometry
	geometry.Draw()

	// Swap the front and back buffers
	window.SwapBuffers()
}

func PollEvents() {
	glfw.PollEvents()
}

func Cleanup() {
	geometry.Cleanup()
	shaders.Cleanup(program)
	window.Destroy()
	glfw.Terminate()
}

func WindowShouldClose() bool {
	return window.ShouldClose()
}

func GetWindow() *glfw.Window {
	return window
}
