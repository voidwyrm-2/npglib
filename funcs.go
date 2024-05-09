package npglib

import (
	"bufio"
	"os"
)

// Function to read the content of a file
func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func writeFile(filename string, data string) error {
	// Open the file with write permissions, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func clamp(number, min, max int) int {
	if number < min {
		return min
	} else if number > max {
		return max
	} else {
		return number
	}
}
