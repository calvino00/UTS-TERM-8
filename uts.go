package main

import (
 "bufio"
 "os"
 "strings"
 "fmt"
)

func main() {

 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan() {
  katas := strings.Split(scanner.Text(), " ")
  hasil := make(map[string]int) 
     for _, kata := range katas { 
        hasil[kata] += 1 
    } 
    fmt.Println(hasil)
 }
}