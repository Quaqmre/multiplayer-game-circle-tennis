package main

import (
	"akif/multiplayer-game-circle-tennis/backend"
	"fmt"
	"os"
)

func main() {
	logger := backend.NewServerLogger(os.Stderr)

	logger.LogStatus("2.filem")
	logger.LogWarning("bu da bir warningdir")
	logger.LogFatalAlert("selam dünya")
	fmt.Println("akif")
}
