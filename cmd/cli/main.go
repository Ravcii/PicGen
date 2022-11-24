package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	picgen "github.com/Ravcii/PicGen"
)

func main() {

	pics := []struct {
		Gradient picgen.Gradient
		File     string
	}{
		{picgen.Gradient{color.RGBA{0, 0, 0, 0}, color.RGBA{255, 255, 255, 255}}, "Test1"},
		{picgen.Gradient{color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 0}}, "Test2"},
		{picgen.Gradient{color.RGBA{255, 0, 0, 0}, color.RGBA{0, 0, 0, 255}}, "Test3"},
		{picgen.Gradient{color.RGBA{0, 255, 0, 255}, color.RGBA{255, 255, 255, 255}}, "Test4"},
	}

	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	for _, pic := range pics {
		wg.Add(1)
		go func(pic struct {
			Gradient picgen.Gradient
			File     string
		}, wg *sync.WaitGroup) {
			defer wg.Done()
			img := picgen.NewImage(rand.Intn(1000)+50, rand.Intn(1000)+50, picgen.WithGradientBackground(&pic.Gradient))
			img.Generate()
			f, err := os.Create(pic.File + ".png")
			if err != nil {
				panic(err)
			}
			err = img.AsPng(f)
			if err != nil {
				panic(err)
			}

			log.Printf("%s created \n", pic.File)
		}(pic, &wg)
	}

	wg.Wait()
}
