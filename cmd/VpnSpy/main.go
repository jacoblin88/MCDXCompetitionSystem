package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/SecretKey"
	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/VpnSpy"
)

func main() {
	f, err := VpnSpy.New("0.0.0.0", 8443, 3, SecretKey.RsaSpyKey.PrivateKeyJsonStr, SecretKey.RsaSpyKey.PrivateKeyJsonStr)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}
