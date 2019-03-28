package fileManager

import (
	"fmt"
	"log"
	"os"
)

var files []*os.File

func CreatNewFile(fileName string) {
	newFile, err := os.Create(fileName + ".v")
	if err != nil {
		fmt.Println(err)
		return
	}
	files = append(files, newFile)
}

func WriteToFile(fileName string, data string) {
	file, err := os.OpenFile(
		fileName+".v",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	byteSlice := []byte(data)
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

func fileManager() {

}
