package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    value := 1
    oldValue := 0

    return func() int {
        answer := value + oldValue
        oldValue = value
        value = answer

        return answer
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
