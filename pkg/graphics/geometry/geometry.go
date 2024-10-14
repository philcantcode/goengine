package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var vao uint32

func InitGeometry() {
	// Define the vertices of a triangle
	vertices := []float32{
		// Positions
		0.0, 0.5, 0.0, // Top
		-0.5, -0.5, 0.0, // Left
		0.5, -0.5, 0.0, // Right
	}

	// Generate and bind a Vertex Array Object
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Generate and bind a Vertex Buffer Object
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Configure the vertex attribute pointers
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))

	// Unbind the VAO and VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func Draw() {
	// Bind the VAO
	gl.BindVertexArray(vao)

	// Draw the triangle
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	// Unbind the VAO
	gl.BindVertexArray(0)
}

func Cleanup() {
	gl.DeleteVertexArrays(1, &vao)
}
