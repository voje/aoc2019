package vector

import "math"

// Vector represents a vector in 3D space.
type Vector struct {
	X float64
	Y float64
	Z float64
}

// AddVector returns v1 + v2.
func AddVector(v1, v2 Vector) Vector {
	return Vector{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
	}
}

// SubVector returns v1 - v2.
func SubVector(v1, v2 Vector) Vector {
	return Vector{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
	}
}

// MulScalar returns v1 multiplied by scalar s.
func MulScalar(v1 Vector, s float64) Vector {
	return Vector{
		v1.X * s,
		v1.Y * s,
		v1.Z * s,
	}
}

// CrossVector returns v1 x v2.
func CrossVector(v1, v2 Vector) Vector {
	return Vector{
		v1.Y*v2.Z - v2.Y*v1.Z,
		v2.X*v1.Z - v1.X*v2.Z,
		v1.X*v2.Y - v2.X*v1.Y,
	}
}

// Intersect returns (intersection point, true/false)
// line1 = a b
// line2 = c d
func Intersect(a, b, c, d Vector) (p *Vector) {
	var r, u, ac Vector
	r = SubVector(b, a)
	u = SubVector(d, c)
	ac = SubVector(a, c)

	// We only want the Z coordinate of the cross product (direction).
	var xur, t, s float64
	xur = CrossVector(u, r).Z

	t = CrossVector(ac, r).Z / xur
	s = CrossVector(ac, u).Z / xur

	if 0 <= t && t <= 1 && 0 <= s && s <= 1 {
		// p = a + s * r
		p := AddVector(a, MulScalar(r, s))
		return &p
	}
	return nil
}

// Manhattan calculates the Manhattan destance between two pointes (represented by vectors).
func Manhattan(a, b Vector) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y) + math.Abs(a.Z-b.Z)
}
