# [Testing in Golang] Hello World

> Target:

1. Build a simple command line program that decides if a number is prime or not
2. Write unit tests to check the program for accuracy
3. Test coverage
4. How to test for user input from the command line

```sh
go mod init helloworld
```

```sh
go test .
go test -v . # -v will detail your test
go test -cover . # -cover will return your test coverage
go test -coverprofile=coverage.out # will return the comprehensive report of which statements had been covered
go tool cover --html=coverage.out # will open a web browser and show me what code is covered
# quick command to check test coverage
go test -coverprofile=coverage.out && go tool cover --html=coverage.out
```

## Terminologies

1. Table test

```go
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"not prime", 1, false, "1 is not prime, by definition!"},
		{"not prime", -1, false, "Negative numbers are not prime!"},
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
	}
```

2. Test coverage

> Percentage of my code has been covered by tests.

3. Test the output to console

```go
func Test_promt(t *testing.T) {
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
```

4. Single test

```sh
go test -v -run Test_checkNumbers
```

5. Groups of tests (test suites)

```sh
go test -v -run Test_alpha
```

```go
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
```

```go
func Test_alpha_intro(t *testing.T) {
	// save a copy of os stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w
```
