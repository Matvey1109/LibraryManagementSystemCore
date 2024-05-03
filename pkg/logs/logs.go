package logs

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Matvey1109/LibraryManagementSystemCore/pkg/loadenv"
)

var logFile *os.File

func OpenLogFile() error {
	_, _, pathToLogFile := loadenv.LoadGlobalEnv()
	file, err := os.OpenFile(pathToLogFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	logFile = file
	return nil
}

func CloseLogFile() {
	if logFile != nil {
		if err := logFile.Close(); err != nil {
			log.Fatal("failed to close log file: %w", err)
		}
	}
	logFile = nil
}

func LogWriter(method, path string, statusCode int) error {
	var message string
	if statusCode == 0 {
		message = "Starting server on port 8080\n"
	} else {
		message = fmt.Sprintf("[%s] %s %s - %d\n", time.Now().Format("2006-01-02 15:04:05"), method, path, statusCode)
	}

	fmt.Print(message)

	if logFile == nil {
		if err := OpenLogFile(); err != nil {
			return err
		}
	}

	if _, err := logFile.WriteString(message); err != nil {
		return errors.New("failed to write to log file")
	}

	return nil
}
