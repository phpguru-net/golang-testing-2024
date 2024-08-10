package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_readUserInput(t *testing.T) {
	// to test this function we need a channel, and an instance of io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	stdin.Write([]byte("2\n"))
	stdin.Write([]byte("q\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "quit", input: "q", expected: ""},
		{name: "quit", input: "Q", expected: ""},
	}
	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s expected %s but got %s", e.name, e.expected, res)
		}
	}
}

func Test_alpha_intro(t *testing.T) {
	// save a copy of os stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our write
	_ = w.Close()

	// reset os stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt function from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro not correct but got %s", string(out))
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our write
	_ = w.Close()

	// reset os stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt function from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_alpha_isPrime(t *testing.T) {

	// Table Test
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -1, false, "Negative numbers are not prime!"},
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
