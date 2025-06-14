package genimage

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/hend41234/gem-flash-image/models"
	"github.com/hend41234/gem-flash-image/utilsfi"
)

var Promt string = ""

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
		break
	}
}

// to generate image, you must have the API KEY
//
// if you have the env file, you can use the utils.LoadConfig(envPath)
//
// in env file, using GEMINI_API_KEY for key.
//
// or you can using parameter apiKey.
//
//	'make sure the API KEY is CORRECT'
func GenerateImage(apiKey ...string) (dataResponse models.ResGenImageModel) {
	utilsfi.Utils.BaseURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-preview-image-generation:generateContent"
	if len(apiKey) > 0 {
		utilsfi.Utils.GeminiApiKey = apiKey[0]
	}
	if utilsfi.Utils.GeminiApiKey == "" {
		log.Fatal("config is not valid, please check your config\nnote: if you not inputed the parameter apiKey in genimage.GenerateImage(), you need to use utils.LoadConfig(envPath)")
	}

	//  try request
	if Promt == "" {
		log.Fatal("error: genimage.Promt is require")
	}
	part := models.Part{Text: Promt}
	content := models.Content{Parts: []models.Part{part}}
	genConf := models.GenConfig{ResponseModalities: []string{"TEXT", "IMAGE"}}
	body := models.ReqGenImageModel{
		Contents:         []models.Content{content},
		GenerationConfig: genConf,
	}
	byteModels, _ := json.Marshal(body)
	//
	url := utilsfi.Utils.BaseURL + "?key=" + utilsfi.Utils.GeminiApiKey
	// fmt.Println(url)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteModels))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error request :\n\t" + err.Error())
	}
	defer res.Body.Close()
	{
		if res.StatusCode != 200 {
			log.Fatal("error : please check your authentication, API_KEY or something")
		}
	}
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

// data
//
//	data = models.ResGenImageModel{}
//
// savePath
// directory / path that you use to save result, by default is "output" in your work directory
//
//	savePath = "path_directory"
//
// nameFile
// name file for output name, by default is "output"
//
//	nameFile = "output_name"
func ConvertDataToImage(data models.ResGenImageModel, savePath string, nameFile string) bool {

	{
		// x, _ := json.Marshal(data)
		// fmt.Println(string(x))
		// ext := data.Candidates[0].Content.Parts[1]
		// fmt.Println(ext)

	}
	ext := getExt(data.Candidates[0].Content.Parts[1].InlineData.MimeType)
	if ext == "" {
		ext = "png"
	}
	// fileJson, _ := os.Open("res.json")
	base64data := data.Candidates[0].Content.Parts[1].InlineData.Data
	byteData, _ := base64.StdEncoding.DecodeString(base64data)

	// if utilsfi.NotExistPath(savePath) {
	// 	fmt.Println("create directory")
	// 	utilsfi.CreatePath(savePath)
	// }

	nameForImage := fmt.Sprintf("%v/%v%v.%v", savePath, nameFile, len(savePath), ext)

	writeImageErr := ioutil.WriteFile(nameForImage, []byte(byteData), 0664)
	if writeImageErr != nil {
		// log.Fatal(writeImageErr)
		return false
	}
	// fmt.Println("generate image successfully.\nsaved in : " + nameForImage)
	return true
}

func getExt(imageType string) string {
	extProb := []string{"jpg", "jpeg", "png", "webp", "bmp", "tiff", "tif", "gif", "svg", "ico", "avif", "heif", "heic", "dds", "exr", "raw", "psd"}
	for _, p := range extProb {
		if imageType == p {
			return p
		}
	}
	re := regexp.MustCompile(`/([^/]+)`)
	match := re.FindStringSubmatch(imageType)

	if len(match) > 1 {
		// fmt.Println(match[1])
		return match[1] // Output: "png", "img", etc
	} else {
		fmt.Println("No match found")
		return ""
	}
}
