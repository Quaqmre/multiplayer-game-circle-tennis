package log_test

import (
	"akif/multiplayer-game-circle-tennis/log"
	"io"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	file, _ := os.OpenFile("default.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	mw := io.MultiWriter(file, os.Stderr)
	log.SetOutput(mw)
	log.Info.Println("Testtir")
}
