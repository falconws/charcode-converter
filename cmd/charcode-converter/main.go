package main

import (
	"fmt"
	"log"
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

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Args must have 1")
	}
	filename := os.Args[1]
	if _, err := os.Stat(filename); err != nil {
		log.Fatal("File not exists")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("File read error")
	}
	outDir := filepath.Dir(filename)
	if utf8.Valid(data) {
		fmt.Println("Try convert from UTF-8 to Shift-JIS")
		b, err := utf8ToShiftJIS(data)
		if err != nil {
			log.Fatal("Convert error")
		}
		outFile := filepath.Join(outDir, getFileNameWithoutExt(filename)+"_sjis"+filepath.Ext(filename))
		os.WriteFile(outFile, b, 0664)
		fmt.Println("Convert complete: " + outFile)
	} else {
		fmt.Println("Try convert from Shift-JIS to UTF-8")
		b, err := shiftJIStoUtf8(data)
		if err != nil {
			log.Fatal("Convert error")
		}
		outFile := filepath.Join(outDir, getFileNameWithoutExt(filename)+"_utf8"+filepath.Ext(filename))
		os.WriteFile(outFile, b, 0664)
		fmt.Println("Convert complete: " + outFile)
	}
}
