package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init(){
    rand.Seed(time.Now().Unix()) 
}

func main() {

    instructorSlice := []string{"Jane", "Zach", "Rey", "Mike", "Alice"}

    if x := instructorSlice[rand.Intn(len(instructorSlice))]; x == "Zach" {
        fmt.Println("Zach works out of the NorthEast")
    } else if x == "Jane" {
        fmt.Println("Jane lives in the Pacific NorthWest")
    } else {
        fmt.Println("Hmm, I don't know much about the instructor,", x)
    }
}

