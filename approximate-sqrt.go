package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64, deltaMax float64) float64 {
    var z, oldZ, delta float64 = 10.0, 0.0, 10.0
    count := 0

    for delta > deltaMax {
        oldZ = z
        z = z - (z * z - x) / (2 * z)
        delta = oldZ - z
        count++
    }

    fmt.Println(count)
    return z
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println("sqrt for", i)
        fmt.Println(Sqrt(float64(i), .01))
        fmt.Println(math.Sqrt(float64(i)))
    }
}
