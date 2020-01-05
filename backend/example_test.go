package backend_test

import (
	"akif/multiplayer-game-circle-tennis/backend"
	"fmt"
	"os"
	"testing"
)

func TestLoga(t *testing.T) {

	logger := backend.NewServerLogger(os.Stderr)

	logger.LogStatus("2.filem")
	logger.LogWarning("bu da bir warningdir")
	logger.LogFatalAlert("selam dünya")
	fmt.Println("akif")
}
