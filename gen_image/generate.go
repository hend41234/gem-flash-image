package genimage

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gemflashimage/models"
	"gemflashimage/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var Promt string = ""
var OutPut string = ""

func InputPromts() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("🟡 input prompt : ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			log.Println("🔴 prompt is empty")
			continue
		}
		Promt = input
		// tag = Tags{Tag: strings.Split(input, ",")} //strings.Split(input, ",")
		break
	}
}

func GenerateImage() (dataResponse models.ResGenImageModel) {
	//  try request
	if Promt == "" {
		// InputPromts()
		log.Fatal("error: genimage.Promt is require")
	}
	part := models.Part{Text: Promt}
	content := models.Content{Parts: []models.Part{part}}
	//
	genConf := models.GenConfig{ResponseModalities: []string{"TEXT", "IMAGE"}}
	//
	body := models.ReqGenImageModel{
		Contents:         []models.Content{content},
		GenerationConfig: genConf,
	}
	byteModels, _ := json.Marshal(body)
	//
	url := utils.Utils.BaseURL + "?key=" + utils.Utils.GeminiApiKey
	// fmt.Println(url)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteModels))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error request :\n\t" + err.Error())
	}
	defer res.Body.Close()
	// readAll, _ := io.ReadAll(res.Body)
	// fmt.Println(string(readAll))
	//

	models := models.ResGenImageModel{}
	decodeErr := json.NewDecoder(res.Body).Decode(&models)
	if decodeErr != nil {
		log.Fatal("error decode response :\n\t" + decodeErr.Error())
	}
	// fmt.Println(models.Candidates[0].Content.Parts[1])
	return models
}

func ConvertDataToImage(data models.ResGenImageModel) bool {
	ext := getExt(data.Candidates[0].Content.Parts[1].InlineData.MimeType)
	if ext == "" {
		ext = "png"
	}

	// fileJson, _ := os.Open("res.json")
	base64data := data.Candidates[0].Content.Parts[1].InlineData.Data
	byteData, _ := base64.StdEncoding.DecodeString(base64data)

	dirSave, _ := os.ReadDir("output/")
	nameForImage := fmt.Sprintf("output/%v.%v", len(dirSave), ext)
	if OutPut != "" {
		nameForImage = "output/" + OutPut + "." + ext
	}
	writeImageErr := ioutil.WriteFile(nameForImage, []byte(byteData), 0664)
	if writeImageErr != nil {
		// log.Fatal(writeImageErr)
		return false
	}
	fmt.Println("generate image successfully.\nsaved in : " + nameForImage)
	return true
}

func getExt(imageType string) string {
	re := regexp.MustCompile(`/([^/]+)`)
	match := re.FindStringSubmatch(imageType)

	if len(match) > 1 {
		fmt.Println(match[1])
		return match[1] // Output: "png", "img", etc
	} else {
		fmt.Println("No match found")
		return ""
	}
}
