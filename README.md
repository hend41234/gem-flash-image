# Google Cloud ( Gemini ) flash-image


## USING CLI

### setup
* ```bash
    chmod +x build
    ```
* ```bash
    ./build
    ```
* ```bash
    ./gemCLI
    ```
### note 
if you don't have an API KEY, you can create an API [here](https://aistudio.google.com/apikey).

## RUN
* example
    ```bash
    ./gemCLI -p "generate image cat fly to the moon"
    ```
    ```bash
    ./gemCLI -p "generate image cat fly to the moon" -o output_name
    ```
    ```bash
    ./gemCLI -p "generate image cat fly to the moon" -o output_name -d directory_to_save
    ```
    ```bash
    ./gemCLI -p "generate image cat fly to the moon" -o output_name -d directory_to_save -c env_file
    ```
* for more details
    ```bash
        ./gemCLI -h
    ```

## USING in YOUR CODE

### install
* ```bash
    go get "github.com/hend41234/gem-flash-image"
    ```
### sample use
* ```golang
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
    	// newGen := genimage.GenerateImage("your_api_key")
    	newGen := genimage.GenerateImage()

    	if ok := genimage.ConvertDataToImage(newGen, "test_output", "person_drinking"); !ok {
    		log.Fatal("error generate image")
    	}
    }
    
    ```
## note
    in the env file, the key that using is GEMINI_API_KEY