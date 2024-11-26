package main

import (
	"os/exec"
	"testing"
)

func TestDuplicateUser(t *testing.T) {
	expectErr, expectNoErr := getHelpers(t)
	expectNoErr(resetDatabase())
	expectNoErr(addUser("peter"))
	expectErr(addUser("peter"))
}

func resetDatabase() ([]byte, error) {
	return exec.Command("go", "run", ".", "reset").CombinedOutput()
}

func addUser(username string) ([]byte, error) {
	return exec.Command("go", "run", ".", "register", username).CombinedOutput()
}

func getHelpers(t *testing.T) (func([]byte, error), func([]byte, error)) {
	expectErr := func(_ []byte, err error) {
		if err == nil {
			t.FailNow()
		}
	}
	expectNoErr := func(_ []byte, err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
	return expectErr, expectNoErr
}
