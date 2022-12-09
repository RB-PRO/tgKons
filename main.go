package tgkons

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	fileToken, errorToken := os.Open(os.Args[1])
	if errorToken != nil {
		log.Fatal(errorToken)
	}

	defer func() {
		if errorToken = fileToken.Close(); errorToken != nil {
			log.Fatal(errorToken)
		}
	}()
	data, errFileToken := ioutil.ReadAll(fileToken)
	if errFileToken != nil {
		log.Fatal(errFileToken)
	}
	fmt.Print(data)
}
