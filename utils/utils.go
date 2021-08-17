package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// list of color
//goland:noinspection GoSnakeCaseUsage
var (
	COLOR_OFF     = []byte("\033[0m")
	COLOR_BLACK   = []byte("\033[0;30m")
	COLOR_RED     = []byte("\033[0;31m")
	COLOR_LRED    = []byte("\033[0;91m")
	COLOR_GREEN   = []byte("\033[0;32m")
	COLOR_LGREEN  = []byte("\033[0;92m")
	COLOR_ORANGE  = []byte("\033[0;93m")
	COLOR_BLUE    = []byte("\033[0;34m")
	COLOR_PURPLE  = []byte("\033[0;35m")
	COLOR_CYAN    = []byte("\033[0;36m")
	COLOR_GRAY    = []byte("\033[0;37m")
)

// Log a message to the console with a beautiful format
func Log(msg string) {
	date := time.Now()

	// get the class that used this method
	funcName, fileNameAndDir, lineNumber, _ := GetTheClass(2)
	fileName := strings.Split(fileNameAndDir, "\\")[1]

	// format the stuff
	fmt.Printf("%s[%s%s%s] %s[%sINFO%s]  %s--- %s[%s%s/%s:%d%s] %s%s\n",
		// time 	- "[HH:mm:ss]"
		COLOR_BLACK, COLOR_GRAY, date.Format("15:04:05"), COLOR_BLACK,
		// log type - "[INFO]"
		COLOR_GREEN, COLOR_LGREEN, COLOR_GREEN,
		// line		- "---"
		COLOR_PURPLE,
		// file 	- "[file.go/func:line]"
		COLOR_GRAY, COLOR_BLUE, fileName, funcName, lineNumber, COLOR_GRAY,
		// message
		COLOR_OFF, msg)
}

// Log an error message, also with beautiful format
// Although an error is not beautiful
func Error(msg string) {
	date := time.Now()

	funcName, fileNameAndDir, lineNumber, _ := GetTheClass(2)
	fileName := strings.Split(fileNameAndDir, "\\")[1]

	// basically the same thing but different colors
	fmt.Printf("%s[%s%s%s] %s[%sERROR%s] %s--- %s[%s%s/%s:%d%s] %s%s\n",
		COLOR_BLACK, COLOR_GRAY, date.Format("15:04:05"), COLOR_BLACK,
		COLOR_RED, COLOR_LRED, COLOR_RED,
		COLOR_PURPLE,
		COLOR_GRAY, COLOR_RED, fileName, funcName, lineNumber, COLOR_GRAY,
		COLOR_CYAN, msg)
}

func Debug(msg string) {
	date := time.Now()

	funcName, fileNameAndDir, lineNumber, _ := GetTheClass(2)
	fileName := strings.Split(fileNameAndDir, "\\")[1]

	fmt.Printf("%s[%s%s%s] %s[%sDEBUG%s] %s--- %s[%s%s/%s:%d%s] %s%s\n",
		COLOR_BLACK, COLOR_GRAY, date.Format("15:04:05"), COLOR_BLACK,
		COLOR_LRED, COLOR_ORANGE, COLOR_LRED,
		COLOR_PURPLE,
		COLOR_GRAY, COLOR_BLUE, fileName, funcName, lineNumber, COLOR_GRAY,
		COLOR_PURPLE, msg)
}

// Get the class that the function is called from
// See here: https://stackoverflow.com/questions/47218715/is-it-possible-to-get-filename-where-code-is-called-in-golang
func GetTheClass(skip ...int) (string, string, int, error) {
	sk := 1
	if len(skip) > 0 && skip[0] > 1 {
		sk = skip[0]
	}
	var pc uintptr
	var ok bool
	pc, fileName, fileLine, ok := runtime.Caller(sk)
	if !ok {
		return "", "", 0, fmt.Errorf("N/A")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	ix := strings.LastIndex(name, ".")
	if ix > 0 && (ix+1) < len(name) {
		name = name[ix+1:]
	}
	funcName := name
	nd, nf := filepath.Split(fileName)
	fileName = filepath.Join(filepath.Base(nd), nf)
	return funcName, fileName, fileLine, nil
}