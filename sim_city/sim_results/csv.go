package sim_results

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func AddLine(filename, lineToAdd string) error {
	// Open the file with the 'os' package
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a buffered writer to efficiently write to the file
	writer := bufio.NewWriter(file)

	// Write the new line to the file
	_, err = writer.WriteString(lineToAdd + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	// Flush the buffer to ensure the data is written to the file
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush buffer: %v", err)
	}

	return nil
}
