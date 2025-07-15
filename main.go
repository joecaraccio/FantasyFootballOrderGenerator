package main

// Import necessary packages
//  All these packages can be found at https://pkg.go.dev/
import (
	"crypto/sha256"
	"embed"
	"fmt"
	"math/rand"
	"time"
)

//go:embed main.go
var sourceFile embed.FS

// Constant, Alphabetical list of names
var names = [12]string{
	"Andre",
	"Ben",
	"Buddy",
	"Goph",
	"Hunty",
	"Jae",
	"Jarod",
	"Joe",
	"Jon",
	"Nick",
	"Spencer",
	"Teigue",
}

func main() {
	// Create a random seed based on the current time
	// 
	rand.Seed(time.Now().UnixNano())

	// Iterate 100 times, shuffling the names, print on the final iteration
	iterations := 100
	for i := 0; i < iterations; i++ {
		shuffled := names
		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})

		if(iterations - 1 == i) {
			fmt.Printf("========== ORDER OUTPUT ==========\n")
			for j, name := range shuffled {
				fmt.Printf("%2d. %s\n", j+1, name)
			}
			fmt.Println()
		}
	}

	// Compute the SHA256 checksum of the embedded file
	data, err := sourceFile.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(data)
	fmt.Printf("SHA256:::: %x\n", sum)

	// Wait for user input before exiting
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
