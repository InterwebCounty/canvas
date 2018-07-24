package glimplgogl

import (
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/tfriedel6/canvas"
)

type GLImpl struct{}

var _ canvas.GL = GLImpl{}

func (_ GLImpl) Ptr(data interface{}) unsafe.Pointer {
	return gl.Ptr(data)
}
func (_ GLImpl) ActiveTexture(texture uint32) {
	gl.ActiveTexture(texture)
}
func (_ GLImpl) AttachShader(program uint32, shader uint32) {
	gl.AttachShader(program, shader)
}
func (_ GLImpl) BindBuffer(target uint32, buffer uint32) {
	gl.BindBuffer(target, buffer)
}
func (_ GLImpl) BindFramebuffer(target uint32, framebuffer uint32) {
	gl.BindFramebuffer(target, framebuffer)
}
func (_ GLImpl) BindRenderbuffer(target uint32, renderbuffer uint32) {
	gl.BindRenderbuffer(target, renderbuffer)
}
func (_ GLImpl) BindTexture(target uint32, texture uint32) {
	gl.BindTexture(target, texture)
}
func (_ GLImpl) BlendFunc(sfactor uint32, dfactor uint32) {
	gl.BlendFunc(sfactor, dfactor)
}
func (_ GLImpl) BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	gl.BufferData(target, size, data, usage)
}
func (_ GLImpl) CheckFramebufferStatus(target uint32) uint32 {
	return gl.CheckFramebufferStatus(target)
}
func (_ GLImpl) Clear(mask uint32) {
	gl.Clear(mask)
}
func (_ GLImpl) ClearColor(red float32, green float32, blue float32, alpha float32) {
	gl.ClearColor(red, green, blue, alpha)
}
func (_ GLImpl) ColorMask(red bool, green bool, blue bool, alpha bool) {
	gl.ColorMask(red, green, blue, alpha)
}
func (_ GLImpl) CompileShader(shader uint32) {
	gl.CompileShader(shader)
}
func (_ GLImpl) CreateProgram() uint32 {
	return gl.CreateProgram()
}
func (_ GLImpl) CreateShader(xtype uint32) uint32 {
	return gl.CreateShader(xtype)
}
func (_ GLImpl) DeleteShader(shader uint32) {
	gl.DeleteShader(shader)
}
func (_ GLImpl) DeleteFramebuffers(n int32, framebuffers *uint32) {
	gl.DeleteFramebuffers(n, framebuffers)
}
func (_ GLImpl) DeleteRenderbuffers(n int32, renderbuffers *uint32) {
	gl.DeleteRenderbuffers(n, renderbuffers)
}
func (_ GLImpl) DeleteTextures(n int32, textures *uint32) {
	gl.DeleteTextures(n, textures)
}
func (_ GLImpl) Disable(cap uint32) {
	gl.Disable(cap)
}
func (_ GLImpl) DisableVertexAttribArray(index uint32) {
	gl.DisableVertexAttribArray(index)
}
func (_ GLImpl) DrawArrays(mode uint32, first int32, count int32) {
	gl.DrawArrays(mode, first, count)
}
func (_ GLImpl) Enable(cap uint32) {
	gl.Enable(cap)
}
func (_ GLImpl) EnableVertexAttribArray(index uint32) {
	gl.EnableVertexAttribArray(index)
}
func (_ GLImpl) FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	gl.FramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer)
}
func (_ GLImpl) FramebufferTexture(target uint32, attachment uint32, texture uint32, level int32) {
	gl.FramebufferTexture(target, attachment, texture, level)
}
func (_ GLImpl) GenBuffers(n int32, buffers *uint32) {
	gl.GenBuffers(n, buffers)
}
func (_ GLImpl) GenFramebuffers(n int32, framebuffers *uint32) {
	gl.GenFramebuffers(n, framebuffers)
}
func (_ GLImpl) GenRenderbuffers(n int32, renderbuffers *uint32) {
	gl.GenRenderbuffers(n, renderbuffers)
}
func (_ GLImpl) GenTextures(n int32, textures *uint32) {
	gl.GenTextures(n, textures)
}
func (_ GLImpl) GenerateMipmap(target uint32) {
	gl.GenerateMipmap(target)
}
func (_ GLImpl) GetAttribLocation(program uint32, name string) int32 {
	return gl.GetAttribLocation(program, gl.Str(name+"\x00"))
}
func (_ GLImpl) GetError() uint32 {
	return gl.GetError()
}
func (_ GLImpl) GetProgramInfoLog(program uint32) string {
	var length int32
	gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &length)
	if length == 0 {
		return ""
	}
	log := strings.Repeat("\x00", int(length+1))
	gl.GetProgramInfoLog(program, length, nil, gl.Str(log))
	return log
}
func (_ GLImpl) GetProgramiv(program uint32, pname uint32, params *int32) {
	gl.GetProgramiv(program, pname, params)
}
func (_ GLImpl) GetShaderInfoLog(program uint32) string {
	var length int32
	gl.GetShaderiv(program, gl.INFO_LOG_LENGTH, &length)
	if length == 0 {
		return ""
	}
	log := strings.Repeat("\x00", int(length+1))
	gl.GetShaderInfoLog(program, length, nil, gl.Str(log))
	return log
}
func (_ GLImpl) GetShaderiv(shader uint32, pname uint32, params *int32) {
	gl.GetShaderiv(shader, pname, params)
}
func (_ GLImpl) GetUniformLocation(program uint32, name string) int32 {
	return gl.GetUniformLocation(program, gl.Str(name+"\x00"))
}
func (_ GLImpl) LinkProgram(program uint32) {
	gl.LinkProgram(program)
}
func (_ GLImpl) RenderbufferStorage(target uint32, internalformat uint32, width int32, height int32) {
	gl.RenderbufferStorage(target, internalformat, width, height)
}
func (_ GLImpl) ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.ReadPixels(x, y, width, height, format, xtype, pixels)
}
func (_ GLImpl) Scissor(x int32, y int32, width int32, height int32) {
	gl.Scissor(x, y, width, height)
}
func (_ GLImpl) ShaderSource(shader uint32, source string) {
	csource, freeFunc := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csource, nil)
	freeFunc()
}
func (_ GLImpl) StencilFunc(xfunc uint32, ref int32, mask uint32) {
	gl.StencilFunc(xfunc, ref, mask)
}
func (_ GLImpl) StencilMask(mask uint32) {
	gl.StencilMask(mask)
}
func (_ GLImpl) StencilOp(fail uint32, zfail uint32, zpass uint32) {
	gl.StencilOp(fail, zfail, zpass)
}
func (_ GLImpl) TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(target, level, internalformat, width, height, border, format, xtype, pixels)
}
func (_ GLImpl) TexParameteri(target uint32, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}
func (_ GLImpl) TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexSubImage2D(target, level, xoffset, yoffset, width, height, format, xtype, pixels)
}
func (_ GLImpl) Uniform1f(location int32, v0 float32) {
	gl.Uniform1f(location, v0)
}
func (_ GLImpl) Uniform1i(location int32, v0 int32) {
	gl.Uniform1i(location, v0)
}
func (_ GLImpl) Uniform2f(location int32, v0 float32, v1 float32) {
	gl.Uniform2f(location, v0, v1)
}
func (_ GLImpl) Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	gl.Uniform4f(location, v0, v1, v2, v3)
}
func (_ GLImpl) UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	gl.UniformMatrix3fv(location, count, transpose, value)
}
func (_ GLImpl) UseProgram(program uint32) {
	gl.UseProgram(program)
}
func (_ GLImpl) VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, offset uint32) {
	gl.VertexAttribPointer(index, size, xtype, normalized, stride, gl.PtrOffset(int(offset)))
}
func (_ GLImpl) Viewport(x int32, y int32, width int32, height int32) {
	gl.Viewport(x, y, width, height)
}
