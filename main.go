package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
)

func main() {
	//Path to file
	filePath := "randomfile.txt"

	err := Shred(filePath)
	if err != nil {
		log.Fatalf("File failed to shred: %v", err)
	} else {
		fmt.Println("File shredded successfully")
	}

	// err := createRandFile(filePath)
	// if err != nil {
	// 	log.Fatalf("Failed to create file: %v", err)
	// } else {
	// 	fmt.Println("File created successfully")
	// }

	// err = randomDataGen(filePath)
	// //Error handler for randomDataGen function
	// if err != nil {
	// 	log.Fatalf("Failed to randomize/overwrite file: %v", err)
	// } else {
	// 	fmt.Println("File overwritten successfully")
	// }

	// err = shredOnly(filePath)
	// //Error handler for shred function
	// if err != nil {
	// 	log.Fatalf("Failed to shred file: %v", err)
	// } else {
	// 	fmt.Println("File shredded successfully")
	// }
}

func Shred(path string) error {
	// Checking to see if file exists
	file, err := os.OpenFile(path, os.O_WRONLY, 0) // For write access.
	if err != nil {
		return err
	}
	defer file.Close() // Ensures file deletion even in case of error

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// Randomize data 3x
	for i := 0; i < 3; i++ {
		randomData := make([]byte, fileSize)
		if _, err := rand.Read(randomData); err != nil {
			return err
		}
		if _, err := file.WriteAt(randomData, 0); err != nil {
			return err
		}
	}

	file.Close()
	return os.Remove(path)
}

func shredOnly(path string) error {
	// Checking to see if file exists
	file, err := os.OpenFile(path, os.O_WRONLY, 0) // For write access.
	if err != nil {
		return err
	}
	defer file.Close() // Ensures file deletion even in case of error

	file.Close()
	return os.Remove(path)
}

func randomDataGen(path string) error {
	// Checking to see if file exists
	file, err := os.OpenFile(path, os.O_WRONLY, 0) // For write access.
	if err != nil {
		return err
	}
	defer file.Close() // Ensures file deleltion even in case of error

	// Randomize data 3x
	for i := 0; i < 3; i++ {
		randomData := make([]byte, 1024)
		if _, err := rand.Read(randomData); err != nil {
			return err
		}
		if _, err := file.WriteAt(randomData, 0); err != nil {
			return err
		}
	}

	return file.Close()
}

func createRandFile(path string) error {
	// Checking to see if file exists
	file, err := os.OpenFile(path, os.O_WRONLY, 0) // For write access.
	if err != nil {
		return err
	}
	defer file.Close() // Ensures file deleltion even in case of error

	content := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	file.Close()
	return err
}
