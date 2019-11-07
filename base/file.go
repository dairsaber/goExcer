package base

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// CreateFile is a function
func CreateFile(fileName string) bool {
	f, error := os.Create(fileName)
	defer f.Close()
	if error != nil {
		fmt.Println("CreateFile==>", error)
		return false
	}
	return true
}

//WriteFile is a function
func WriteFile(mymap map[string]string, fileName string) bool {
	f, err := os.OpenFile(fileName, os.O_RDWR, 6)
	defer f.Close()
	if err != nil {
		fmt.Println("WriteFile==>", err)
		return false
	}
	var writeStr string = createStr(mymap)
	_, writeErr := f.WriteString(writeStr)
	if writeErr != nil {
		fmt.Println("error==>", writeErr)
		return false
	}
	return true
}

// createStr is a function
func createStr(mymap map[string]string) string {
	strs := make([]string, 0)
	for key, value := range mymap {
		var buffer bytes.Buffer
		buffer.WriteString(key)
		buffer.WriteString(":")
		buffer.WriteString(value)

		strs = append(strs, buffer.String())
	}
	return strings.Join(strs, "\n")
}

//Person ren
type Person struct {
	name  string
	hobby string
}

//CreateStructFromFile is a function
func CreateStructFromFile(fileName string) *Person {
	p := new(Person)
	f, err := os.OpenFile(fileName, os.O_RDONLY, 6)
	defer f.Close()
	if err != nil {
		fmt.Println("CreateStructFromFile==>", err)
		return p
	}

	// fileData := make([]byte, 0)
	// _, readErr := f.Read(fileData)
	// if readErr != nil {
	// 	return p
	// }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text() // or
		//line := scanner.Bytes()

		//do_your_function(line)
		// fmt.Printf("%s\n", line)
		strs := strings.Split(line, ":")
		// p[strs[0]] = strs[1]
		if strs[0] == "name" {
			p.name = strs[1]
		} else if strs[0] == "hobby" {
			p.hobby = strs[1]
		}
	}
	return p
	// io.ReaderFrom()
}
