package main

import (
	"log"

	"github.com/hend41234/gem-flash-image/genimage"
	"github.com/hend41234/gem-flash-image/utilsfi"
)

func main() {
	// sample usage CLI
	// CLI()

	// sample usage in your code
	UseInYourCode()
}

func UseInYourCode() {
	// input prompt
	genimage.Promt = "create the picture the one person whos drinks"
	// using env file for get GEMINI_API_KEY
	utilsfi.LoadConfig(".env")
	// or you input API KEY in your code
	// newGen := genimage.GenerateImage("GEMINI_API_KEY")
	newGen := genimage.GenerateImage()

	if ok := genimage.ConvertDataToImage(newGen, "test_output", "person_drinking"); !ok {
		log.Fatal("error generate image")
	}
}
