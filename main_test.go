package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout // keep backup of the real stdout
	read, write, _ := os.Pipe()
	os.Stdout = write

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	
	go printSomething("Hello, World!", &waitgroup)
	waitgroup.Wait()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut // restore the real stdout

	if !strings.Contains(output, "Hello, World!") {
		t.Errorf("Expected output to contain 'Hello, World!', but got: %s", output)
	}
}