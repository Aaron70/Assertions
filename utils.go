package assertions

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func getFilePath() string {
	_, file, line, _ := runtime.Caller(4)
	return fmt.Sprintf("%v:%v", filepath.Base(file), line)
}

func getPrefixMessage(description string) string {
	if description == "" {
		return fmt.Sprintf("%v: Condition not met: ", getFilePath())
	}
	return fmt.Sprintf("%v: Condition '%v' not met: ", getFilePath(), description)
}

func getExpectedGivenMessage(description string, expected any, given any) string {
	return fmt.Sprintf("%vexpected '%v' given '%v'.", getPrefixMessage(description), expected, given)
}
