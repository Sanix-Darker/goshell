package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ProcessCode() {
	var line, finalCode string

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
				fmt.Println("Help menu... comming soon")
				goto ASK_AGAIN
			case strings.HasPrefix(line, "resetvar"):
				IMPORTS_INPUT = nil
				goto ASK_AGAIN
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

		writeFile(HIST_FILE, line)
		finalCode = extractCodeAndImport(line)

		writeFile(TEMP_FILE, finalCode)
		if CODE_INPUT != nil && len(CODE_INPUT) > 0 {

			out := ExecCode(TEMP_FILE)

			if strings.Contains(out, "exit status") {
				// we want to remove the last element if it's not a var decalration
				if isVarDeclaration(line) {
				} else {
					CODE_INPUT = CODE_INPUT[:len(CODE_INPUT)-1]
					fmt.Println(out)
				}
			} else {
				fmt.Println(out)
			}
		}
	}
}

func isVarDeclaration(ss string) bool {

	if strings.Contains(ss, " := ") ||
		strings.Contains(ss, " += ") ||
		strings.Contains(ss, " -= ") ||
		strings.Contains(ss, " == ") ||
		strings.HasPrefix(ss, "var ") {
		return true
	}
	return false
}

func ExecCode(path string) string {
	cmd := exec.Command("go", "run", path)
	output, err := cmd.Output()

	// when a variable is created this may fail. do we need to print that out ?
	if err != nil {
		return err.Error()
	}

	return string(output)
}

func saveCodeInFile(fTmpfile *os.File, codeString string) {
	// fmt.Println("code: " + codeString)
	_, err := fTmpfile.WriteString(codeString)

	if err != nil {
		fmt.Println("Error writing inside the file " + err.Error())
		return
	}
}

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
	// else if strings.Contains(stdIn, " := ") || strings.HasPrefix(stdIn, "var ") {
	// 		VAR_INPUT = append(VAR_INPUT, " "+stdIn)
	// 		CODE_INPUT = VAR_INPUT
	// 	}

	if len(IMPORTS_INPUT) > 0 {
		importCode = `import (
` + strings.Join(IMPORTS_INPUT, ",\n") + `
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
