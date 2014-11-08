package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt struct {
    NegativeValue float64
}

func (e *ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e.NegativeValue)
}

func Sqrt(x float64, deltaMax float64) (float64, error) {
    var z, oldZ, delta float64 = 10.0, 0.0, 10.0
    count := 0

    if x < 0 {
        return -1, &ErrNegativeSqrt {x}
    }

    for delta > deltaMax {
        oldZ = z
        z = z - (z * z - x) / (2 * z)
        delta = oldZ - z
        count++
    }

    fmt.Println(count)
    return z, nil
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("sqrt for", i)
        fmt.Println(Sqrt(float64(i), .01))
        fmt.Println(math.Sqrt(float64(i)))
    }

    fmt.Println("there should be an error")
    fmt.Println(Sqrt(-2, .01))
}
