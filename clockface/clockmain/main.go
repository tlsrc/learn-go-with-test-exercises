package main

import (
	"learngo/clockface"
	"os"
	"time"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
