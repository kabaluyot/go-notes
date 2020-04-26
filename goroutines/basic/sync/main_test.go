package main

import (
	"testing"
)

func TestDirectGoroutine(t *testing.T) {
	directGoRoutines()
}

func TestSimpleWaitGroup(t *testing.T) {
	simpleWaitGroup()
}

func TestWaitGroupAndRWMutex(t *testing.T) {
	waitGroupAndRWMutex()
}
