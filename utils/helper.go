package utils

import (
	"fmt"
	"os"
	"strings"
)

func IndexOfRune(r []rune, char rune) int {
	for i, v := range r {
		if v == char {
			return i
		}
	}
	return -1
}

func GetMostRepeated(array []int) int {
	counts := make(map[int]int)
	for _, char := range array {
		counts[char]++
	}
	maxVal := 0
	maxKey := 0
	for i, v := range counts {
		if v > maxVal {
			maxVal = v
			maxKey = i
		}
	}
	return maxKey
}

func ReadTxt(fileName string) string {
	mainPath := getMainPath()
	filePath := fmt.Sprintf("%s/data/files/%s", mainPath, fileName)

	f, err := os.Open(filePath)
	if err != nil {
		return ""
	}

	defer f.Close()

	var text string
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return ""
		}
		text += string(buf[:n])
	}

	return text
}

// getMainPath returns the path up to the CrypGo directory.
func getMainPath() string {
	callerPath, _ := os.Getwd()

	targetDir := "CrypGo"

	index := strings.Index(callerPath, targetDir)
	pathUpToTarget := callerPath[:index+len(targetDir)]

	return pathUpToTarget
}
