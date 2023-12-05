package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Test data to the file here....")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	b := bytes.NewBufferString("Hello there ")
	fmt.Println(b.String())
	b.WriteString("Again, Hi!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Reset Done.. Forget About it!!")
	fmt.Println(b.String())

	f, err = os.Create("output2.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	p001 := person{
		first: "AlexanderBellsMan",
	}

	var b0 bytes.Buffer
	p001.writeOut(f)
	p001.writeOut(&b0)
	// b.Write([]byte("New strings in a slice of bytes"))
	fmt.Println(b0.String())

}
