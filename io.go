package bit

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func ReadImage(filename string) Image {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
	}

	return Image{image: img}
}
