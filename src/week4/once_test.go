package week4

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var a string
var c int
var once sync.Once

func setup() {
	c++
	t := time.Now()
	s := string(t.Format(time.RFC3339Nano))
	fmt.Printf("c=%d\n", c)
	a = "Hello World " + s
}

func doprint() {
	once.Do(setup)
	fmt.Printf("%s: %s\n", string(time.Now().Format(time.RFC3339Nano)), a)
}

func twoprint() {
	go doprint()
	go doprint()
	time.Sleep(time.Second * 1)
}

func TestTwoprint(t *testing.T) {
	twoprint()
}
