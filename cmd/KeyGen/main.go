package main

import (
	"fmt"

	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/NetChannel"
)

func main() {
	key, err := NetChannel.GenRsaKeyPair()
	if err != nil {
		return
	}
	fmt.Print(string(key))
}
