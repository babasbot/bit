package bit

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
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

func (i *Image) WriteImage(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	err = png.Encode(file, i.image)
	if err != nil {
		fmt.Println(err)
	}
}
