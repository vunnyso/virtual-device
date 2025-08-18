package utils

import (
	"fmt"
	"os"
	"time"
)

func WaitForEventFile(eventPath string, timeout time.Duration) error {
	start := time.Now()
	for {
		file, err := os.Open(eventPath)
		if err == nil {
			fmt.Println("vunny closing file")
			file.Close()
			return nil
		} else {
			fmt.Printf("vunny  %s and %v", eventPath, err)
		}

		if time.Since(start) > timeout {
			return fmt.Errorf("event file %s is not ready within the timeout1", eventPath)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
