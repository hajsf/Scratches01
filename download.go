package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	image := "https://drive.google.com/open?id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"
	//	url := "https://googledrive.com/host/1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"
	//	url := "https://drive.google.com/download?id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"
	// url := "https://docs.google.com/uc?export=download&id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"
	//	id := "1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"

	id := image[33:]
	fmt.Println(id)
	url := "https://docs.google.com/uc?export=download&id=" + id
	fileName := "file2.jpg"
	fmt.Println("Downloading file...")

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	fmt.Println(n, "bytes downloaded")
}
