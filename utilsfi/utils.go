package utilsfi

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

var Utils Utilization

// func init() {
// 	Utils = new(Utilization)
// 	for {
// 		loadErr := godotenv.Load(".env")
// 		if loadErr != nil {
// 			fmt.Println("preparing to new setup")
// 			setup()
// 			continue
// 		}
// 		break
// 	}

// 	Utils.GeminiApiKey = os.Getenv("AISTUDIO_API")
// 	Utils.BaseURL = os.Getenv("BASE_URL")

// }

func InputGeminiApiKey() string {
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

// please use GEMINI_API_KEY for key of api key
func LoadConfig(envPath string) {
	_ = godotenv.Load(envPath)
	// Utils = new(Utilization)
	BaseURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-preview-image-generation:generateContent"
	// GemApiKey := inputGeminiApiKey()
	GemApiKey := os.Getenv("GEMINI_API_KEY")
	if GemApiKey == "" {
		log.Fatal("API_KEY not found, please setting config or use utils.Utils.GeminiApiKey")
	}
	Utils.GeminiApiKey = GemApiKey
	Utils.BaseURL = BaseURL
	env := fmt.Sprintf(`GEMINI_API_KEY="%v"`,
		GemApiKey,
	)
	writeErr := ioutil.WriteFile(".env", []byte(env), 0664)
	if writeErr != nil {
		log.Fatal("write env file Error :\n\t" + writeErr.Error())
	}
}
