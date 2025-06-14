package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hend41234/gem-flash-image/genimage"

	"github.com/hend41234/gem-flash-image/utilsfi"
)

var (
	NewPrompt string
	Dir       string
	Output    string
	Config    string
)

func init() {
	flag.StringVar(&NewPrompt, "promt", "", "--promt 'generate image cat in the moon'")
	flag.StringVar(&NewPrompt, "p", "", "-p 'generate image cat in the moon'")
	flag.StringVar(&Dir, "dir", "output", "-dir directory_to_save")
	flag.StringVar(&Dir, "d", "output", "-d directory_to_save")
	flag.StringVar(&Output, "output", "output", "--output output_name")
	flag.StringVar(&Output, "o", "output", "-o output_name")
	flag.StringVar(&Config, "config", ".env", "--config env_path")
	flag.StringVar(&Config, "c", ".env", "-c env_path")

	flag.Parse()
}

func CLI() {
	//  run
	if flag.Lookup("p").Value.String() == "" {
		fmt.Println("-p / --promt is require")
		return
	}

	checkConfig()
	genimage.Promt = flag.Lookup("p").Value.String()
	// genimage.OutPut = flag.Lookup("o").Value.String()

	newGenImage := genimage.GenerateImage()
	ok := genimage.ConvertDataToImage(newGenImage, Dir, Output)
	if !ok {
		log.Fatal("error generate image")
	}

}

func checkConfig() {
	utilsfi.LoadConfig(Config)
	if utilsfi.Utils.GeminiApiKey == "" {
		utilsfi.Utils.GeminiApiKey = utilsfi.InputGeminiApiKey()
		utilsfi.Utils.BaseURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-preview-image-generation:generateContent"

		var saveConfig string
		fmt.Println("save config [y/n]? ")
		fmt.Scan(&saveConfig)
		if saveConfig != "y" {
			return
		}
		env := fmt.Sprintf(`
	API_KEY="%v"`,
			utilsfi.Utils.GeminiApiKey,
		)
		writeErr := ioutil.WriteFile(".env", []byte(env), 0664)
		if writeErr != nil {
			log.Fatal("write env file Error :\n\t" + writeErr.Error())
		}
	}
}
