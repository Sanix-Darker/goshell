package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Fix the space issue (when we are writing an input)
// concatenate multiple line when we're executing them
// keep an history of what we are executing
// capture the arrows keys to put to the stdin the appropriage value
const checkTrueThis = true
const VERSION = "0.0.1-alpha"
const WATERMARK = `GoShell ` + VERSION + `
Type "help" for more information.
Hit "reset" to clean all instructions
Hit "resetimport" to clean all imports`
const TEMP_FILE = "/tmp/golang-shell-tmpcode.go"
const HIST_FILE = "/tmp/golang-shell-tmphistory.txt"

func main() {
	deleteFile(TEMP_FILE)

	// we set up files
	// (clearning them or create them if they don't exist)
	fTmpfile := createFile(TEMP_FILE)
	defer fTmpfile.Close()

	fHistory := createFile(TEMP_FILE)
	defer fHistory.Close()

	// then we write the code and execute it
	ProcessCode(fTmpfile, fHistory)
}

func ProcessCode(fTmpfile *os.File, fHistory *os.File) {
	var line, finalCode string

	fmt.Println(WATERMARK)

	for {
	ASK_AGAIN:
		fmt.Print(">> ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line = scanner.Text()

			// we need to handle early lines
			// quit if the use want to exit
			switch checkTrueThis {
			case strings.HasPrefix(line, "help"):
				fmt.Println("Help menu...")
			case strings.HasPrefix(line, "resetimport"):
				IMPORTS_INPUT = nil
				goto ASK_AGAIN
			case strings.HasPrefix(line, "reset"):
				CODE_INPUT = nil
				goto ASK_AGAIN
			case strings.HasPrefix(line, "q"),
				strings.HasPrefix(line, "exit"),
				strings.HasPrefix(line, "quit"):
				return
			case len(line) == 0 || line == "\n":
				goto ASK_AGAIN
			}
		}

		fHistory.WriteString(line)
		finalCode = extractCodeAndImport(line)

		// fmt.Println(finalCode)
		saveCodeInFile(fTmpfile, finalCode)
		if len(CODE_INPUT) > 0 {
			out, _ := ExecCode(TEMP_FILE)
			fmt.Println(out)
		}
	}
}

func ExecCode(path string) (string, error) {
	cmd := exec.Command("go", "run", path)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error running the temp go file " + err.Error())
		fmt.Println(err)
		return "", err
	}

	return string(output), nil
}

func saveCodeInFile(fTmpfile *os.File, codeString string) {
	// fmt.Println("code: " + codeString)
	_, err := fTmpfile.WriteString(codeString)

	if err != nil {
		fmt.Println("Error writing inside the file" + err.Error())
		return
	}
}

// we are going to store all instructions here
var CODE_INPUT []string
var IMPORTS_INPUT []string

func extractCodeAndImport(stdIn string) string {
	var importCode, finalCode string

	if strings.HasPrefix(stdIn, "import ") {
		IMPORTS_INPUT = append(
			IMPORTS_INPUT,
			" 	\""+strings.ReplaceAll(stdIn, "import ", "")+"\"",
		)
	} else {
		CODE_INPUT = append(CODE_INPUT, " "+stdIn)
	}

	if len(IMPORTS_INPUT) > 0 {
		importCode = `import (
` + strings.Join(IMPORTS_INPUT, "\n") + `
)`
	}

	// We need to find all lines starting with import
	// and write imports at the begining of the file
	// just after the package main
	finalCode = `package main
` + importCode + `
func main(){
	` + strings.Join(CODE_INPUT, "\n") + `
}`

	return finalCode
}
