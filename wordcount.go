package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

// countWords returns the number of words in the supplied string.
func countWords(s string) int {
    inWord := false
    count := 0
    for _, r := range s {
        if unicode.IsSpace(r) {
            if inWord {
                count++
                inWord = false
            }
        } else {
            inWord = true
        }
    }
    if inWord {
        count++
    }
    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        total += countWords(line)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading stdin:", err)
        os.Exit(1)
    }
    fmt.Println(total)
}
