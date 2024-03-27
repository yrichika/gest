package gt

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"

// long name to prevent env variable conflict
const EnvName = "GESTON20240125"

func GreenMsg(msg string) string {
	return Green + msg + Reset
}

func YellowMsg(msg string) string {
	return Yellow + msg + Reset
}

func RedMsg(msg string) string {
	return Red + msg + Reset
}

func itFuncFailMsg(description string, elapsed time.Duration) string {
	timeInSeconds := fmt.Sprintf("%.4f", elapsed.Seconds())
	return RedMsg("    - fail: it \"" + description + "\"  (" + timeInSeconds + "s)")
}

func itFuncPassMsg(description string, elapsed time.Duration) string {
	timeInSeconds := fmt.Sprintf("%.4f", elapsed.Seconds())
	// TODO: 'it' should be switchable to 'test'
	return GreenMsg("    - pass: it \"" + description + "\"  (" + timeInSeconds + "s)")
}

func itFuncSkipMsg(description string) string {
	return YellowMsg("    - skip: it \"" + description + "\"")
}

func IsInSlice[T comparable](val T, slice []T) bool {
	eq := func(a, b T) bool {
		return a == b
	}
	return ContainsElement(val, slice, eq)
}

func ContainsElement[T any](val T, slice []T, predicate func(T, T) bool) bool {
	for _, item := range slice {
		if predicate(val, item) {
			return true
		}
	}
	return false
}

func GetAllTestFileDirectories(isRunInAllDirs bool) []string {
	var directories []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip unnecessary directories
		if info.IsDir() && ignoreUsually(path, info, isRunInAllDirs) {
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "_test.go") {
			dirPath := filepath.Dir(path)
			// exclude duplicate directory names
			if !IsInSlice(dirPath, directories) {
				directories = append(directories, dirPath)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(RedMsg("Error walking the directory:"), err)
	}

	return directories
}

// Add directory names to ignore only when `-all` flag is NOT specified
func ignoreUsually(path string, info os.FileInfo, runInAllDirs bool) bool {
	// do not ignore the project root directory
	if path == "." {
		return false
	}

	if runInAllDirs {
		return alwaysIgnore(info)
	}

	return alwaysIgnore(info) ||
		// Add directory names to ignore if required
		path == "examples"
}

// Add directory names to always ignore
// hidden directories are always ignored
func alwaysIgnore(info os.FileInfo) bool {
	return strings.HasPrefix(info.Name(), ".")
}
