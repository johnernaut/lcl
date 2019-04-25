package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseLocaleFile(fileName string) (map[string]string, error) {
	lines := []string{}

	file, err := os.Open(fileName)
	// file, err := os.Open("/Users/john/Desktop/test.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileStats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bufCapacity := fileStats.Size()
	buf := make([]byte, bufCapacity)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, int(bufCapacity))
	scanner.Split(splitOnSemiColon)

	// remove empty lines
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	// remove comments
	reComment := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(re.ReplaceAllString(line, ""), "\r\n")
		line = reComment.ReplaceAllString(line, "")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	parsedLines := parseLines(lines)
	return parsedLines, nil
}

func splitOnSemiColon(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ';' {
			return i + 1, data[:i], nil
		}
	}

	return 0, data, bufio.ErrFinalToken
}

func parseLines(lines []string) map[string]string {
	parsedLineMap := map[string]string{}

	if len(lines) > 0 {
		for _, line := range lines {
			split := strings.Split(line, "=")
			if len(split) > 1 {
				key, skip := formatKey(split[0])
				if skip {
					continue
				}

				val := formatValue(split[1])
				parsedLineMap[key] = val
			}
		}
	}

	return parsedLineMap
}

func formatKey(key string) (string, bool) {
	removedQuotes := strings.TrimSpace(strings.Replace(key, "\"", "", -1))
	removedLeadingNewLine := strings.TrimLeft(removedQuotes, "\n")
	if len(removedLeadingNewLine) > 2 {
		// remove leading comments on strings
		if removedLeadingNewLine[0:2] == "//" {
			return "", true
		}
	}

	return removedLeadingNewLine, false
}

func formatValue(value string) string {
	return strings.TrimSpace(strings.Replace(value, "\"", "", -1))
}
