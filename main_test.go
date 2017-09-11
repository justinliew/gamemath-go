package main

import "testing"

func TestMagnitude(t *testing.T) {
	v := Vector3D{
		X : 1,
		Y : 2,
		Z : 2,
	}
	if v.Magnitude() != 3 {
		t.Errorf("Magnitude should be 3")
	}
}

func TestScalarMult(t *testing.T) {
	v := Vector3D{
		X : 1,
		Y : 2,
		Z : 2,
	}
	vres := Vector3D{
		X : 10,
		Y : 20,
		Z : 20,
	}

	if v.Mult(10) != vres {
		t.Errorf("Scalar multiply doesn't work")
	}
}

func TestScalarMultFactor(t *testing.T) {
	v := Vector3D{
		X : 1,
		Y : 2,
		Z : 2,
	}
	msg1 := v.Mult(10).Magnitude()
	msg2 := v.Magnitude() * 10
	if msg1 != msg2 {
		t.Errorf("Magnitude and Scalar Multiply Factoring is incorrect")
	}
}