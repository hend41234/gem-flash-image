package main

import (
	"flag"
	"fmt"
	genimage "gemflashimage/gen_image"
	"log"
)

var NewPrompt string
var Output string

func init() {
	flag.StringVar(&NewPrompt, "promt", "", "--promt 'generate image cat in the moon'")
	flag.StringVar(&NewPrompt, "p", "", "-p 'generate image cat in the moon'")
	flag.StringVar(&Output, "output", "", "-output output_name")
	flag.StringVar(&Output, "o", "", "-o output_name")
	flag.Parse()
}

func CLI() {
	// CreateFlag()
	//  run
	if flag.Lookup("p").Value.String() == "" {
		fmt.Println("-p / --promt is require")
		return
	}
	genimage.Promt = flag.Lookup("p").Value.String()
	genimage.OutPut = flag.Lookup("o").Value.String()
	newGenImage := genimage.GenerateImage()
	ok := genimage.ConvertDataToImage(newGenImage)
	if !ok {
		log.Fatal("error generate image")
	}

}
