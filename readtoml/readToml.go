package readtoml

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// Method for to read a toml file details
func ReadToml(pFilename string) (lFileDetails any) {
	_, lErr := toml.DecodeFile(pFilename, &lFileDetails)
	if lErr != nil {
		log.Println("Error (RRT01) :", lErr.Error())
	}
	return
}

func GetConfigValue(pConfig any, pKey string) string {
	return fmt.Sprintf("%v", pConfig.(map[string]any)[pKey])
}
