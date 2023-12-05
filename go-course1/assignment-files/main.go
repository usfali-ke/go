package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// fmt.Println(os.Args[1])
	filename := os.Args[1]

	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// s := strings.Split(string(bs), ",")
	fmt.Println(string(bs))

	// Option two
	f, err := os.Open(filename)
	defer f.Close() //
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

	// output00, err := getContents("test.txt")
	fmt.Println()
	fmt.Println("Testing getContents func")
	fmt.Println(getContents("test.txt"))

}

// Contents returns the file's contents as a string.
func getContents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}
