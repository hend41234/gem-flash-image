package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Utilization struct {
	GeminiApiKey string
	BaseURL      string
}

var Utils *Utilization

func init() {
	Utils = new(Utilization)
	for {
		loadErr := godotenv.Load(".env")
		if loadErr != nil {
			fmt.Println("preparing to new setup")
			setup()
			continue
		}
		break
	}

	Utils.GeminiApiKey = os.Getenv("AISTUDIO_API")
	Utils.BaseURL = os.Getenv("BASE_URL")

}

func inputGeminiApiKey() string {
	var input string
	for {
		fmt.Println("input GEMINI API KEY : ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("input invalid..!!")
			var garbage string
			fmt.Scanln(&garbage)
			continue
		}
		return input
	}
}

func setup() {
	Utils = new(Utilization)
	BaseURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-preview-image-generation:generateContent"
	GemApiKey := inputGeminiApiKey()
	env := fmt.Sprintf(`BASE_URL="%v"
	GEMINI_API_KEY="%v"`,
		BaseURL,
		GemApiKey,
	)
	writeErr := ioutil.WriteFile(".env", []byte(env), 0664)
	if writeErr != nil {
		log.Fatal("write env file Error :\n\t" + writeErr.Error())
	}

}
