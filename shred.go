package shred

import (
    "log"
    "fmt"
    "os"
)

func Shred(path string) error {
    file, err := os.OpenFile(path, os.O_WRONLY, 0) // For write access.
    if err != nil {
        return err
    }
    
    return os.Remove(path)
}
