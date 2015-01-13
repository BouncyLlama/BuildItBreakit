package main

import (
	"fmt"
	"testing"
)

func TestDerpin(t *testing.T) {
	if !testing.Short() {
		fmt.Println("success")
		return
	}
	t.FailNow()
}
