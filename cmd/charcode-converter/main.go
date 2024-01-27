package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func utf8ToShiftJIS(data []byte) ([]byte, error) {
	t := japanese.ShiftJIS.NewEncoder()
	b, _, err := transform.Bytes(t, data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func shiftJIStoUtf8(data []byte) ([]byte, error) {
	t := japanese.ShiftJIS.NewDecoder()
	b, _, err := transform.Bytes(t, data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func showExitPrompt() {
	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func showErrorExit(message string) {
	fmt.Println(message)
	showExitPrompt()
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		showErrorExit("Args must have 1")
	}
	filename := os.Args[1]
	if _, err := os.Stat(filename); err != nil {
		showErrorExit("File not exists")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		showErrorExit("File read error")
	}
	outDir := filepath.Dir(filename)
	if utf8.Valid(data) {
		fmt.Println("Try convert from UTF-8 to Shift-JIS")
		b, err := utf8ToShiftJIS(data)
		if err != nil {
			showErrorExit("Convert error")
		}
		outFile := filepath.Join(outDir, getFileNameWithoutExt(filename)+"_sjis"+filepath.Ext(filename))
		os.WriteFile(outFile, b, 0664)
		fmt.Println("Convert complete: " + outFile)
	} else {
		fmt.Println("Try convert from Shift-JIS to UTF-8")
		b, err := shiftJIStoUtf8(data)
		if err != nil {
			showErrorExit("Convert error")
		}
		outFile := filepath.Join(outDir, getFileNameWithoutExt(filename)+"_utf8"+filepath.Ext(filename))
		os.WriteFile(outFile, b, 0664)
		fmt.Println("Convert complete: " + outFile)
	}
	showExitPrompt()
}
