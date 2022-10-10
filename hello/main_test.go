package main

import (
	"fmt"

	"testing"
	"golang.org/x/example/stringutil"
)

func TestUpper(t *testing.T) {
    fmt.Println(stringutil.ToUpper("Hello"))
}
