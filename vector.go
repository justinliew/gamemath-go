package main // TODO move this to its' own package

import "math"

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z + v.Z)
}

func (v Vector3D) Mult(s float64) Vector3D {
	return Vector3D{
		v.X * s,
		v.Y * s,
		v.Z * s,
	}
}