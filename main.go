package main

import "log"

func main() {
	v := Vector3D{
		X : 1,
		Y : 2,
		Z : 3,
	}
	log.Printf("V: %f,%f,%f", v.X, v.Y, v.Z)
}
