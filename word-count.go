package main

/* will only work in the golang tour runner: tour.golang.org */
import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    var strArray []string = strings.Fields(s)
    wordCounter := make(map[string]int)

    /* map values will automatically be assigned to the default value */
    for _, word := range strArray {
        wordCounter[word]++
    }

    return wordCounter
}

func main() {
    wc.Test(WordCount)
}
