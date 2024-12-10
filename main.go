package main

import (
    "fmt"
    "log"
    "os"
    "crypto/rand"
)

func main() {
   //Path to file
   filePath := "randomfile.txt"
   
   err := shred(filePath)
   if err != nil {
       log.Fatalf("Failed to shred file: %v", err)
   } else {
   	fmt.Println("File shredded successfully")
   }
}

func shred(path string) error {
    file, err := os.OpenFile(path, os.O_WRONLY, 0) // For write access.
    if err != nil {
        return err
    }
    defer file.Close()
    
    
    file.Close()
    return os.Remove(path)
}
