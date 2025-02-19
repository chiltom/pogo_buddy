package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnv reads key=value pairs from a .env file and sets them as environment variables
func LoadEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Warning: Could not open .env file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // Ignore empty lines and comments
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Warning: Invalid line in .env file")
			continue // Ignore invalid lines
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
}
