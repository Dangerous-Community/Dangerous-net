package art_link

import (
    "bufio"
    "fmt"
    "strings"
    "time"
    "embed"
)
//go:embed *.txt
var artFiles embed.FS


func PrintFileSlowly(fileName string) error {
    fileData, err := artFiles.ReadFile(fileName)
    if err != nil {
        return err
    }

    scanner := bufio.NewScanner(strings.NewReader(string(fileData)))
    for scanner.Scan() {
        line := scanner.Text()
        for _, char := range line {
            fmt.Print(string(char))
            time.Sleep(2 * time.Millisecond) // Adjust the delay here to mimic the baud rate
        }
        fmt.Println()
    }

    return scanner.Err()
}
