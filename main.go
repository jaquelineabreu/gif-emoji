package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

func main() {
	imageFile := "images2.jpeg"
	outputFile := "output.jpg"

	im, err := gg.LoadImage(imageFile)
	if err != nil {
		log.Fatalf("Error reading image file %s: %v", imageFile, err)
	}

	extraSpace := 25.0
	// Criar um novo contexto com dimensões maiores
	dc := gg.NewContext(im.Bounds().Dx(), im.Bounds().Dy()+int(extraSpace))
	dc.SetRGB(1, 1, 1) // Cor branca
	dc.Clear()

	// Desenhar a imagem original no contexto branco
	dc.DrawImage(im, 0, 0)

	text := "Teste1 😀"
	parts := strings.Split(text, " ")
	textPart := parts[0]
	emojiPart := parts[1]

	emojiMap := map[string]string{
		"😀": "sorriso.png",
	}
	iconPath, ok := emojiMap[emojiPart]
	if !ok {
		log.Fatalf("Icon not found for emoji: %s", emojiPart)
	}

	icon, err := gg.LoadImage(iconPath)
	if err != nil {
		log.Fatalf("Error loading icon for emoji %s: %v", emojiPart, err)
	}

	desiredSize := 20.0
	scale := desiredSize / float64(icon.Bounds().Dx())
	iconWidth := float64(icon.Bounds().Dx()) * scale
	iconHeight := float64(icon.Bounds().Dy()) * scale

	textWidth, _ := dc.MeasureString(textPart)

	textX := (im.Bounds().Dx() - int(textWidth)) / 2
	textY := float64(im.Bounds().Dy()) + 20

	resizedIcon := resize.Resize(uint(iconWidth), uint(iconHeight), icon, resize.Lanczos3)

	iconX := float64(textX) + textWidth + 20

	iconY := im.Bounds().Dy() + int(extraSpace) - int(iconHeight) - 1

	dc.SetRGB(0, 0, 0) // Cor do texto preto
	dc.DrawString(textPart, float64(textX), textY)

	dc.DrawImage(resizedIcon, int(iconX), iconY)

	newImage := dc.Image()
	err = saveImage(newImage, outputFile)
	if err != nil {
		log.Fatalf("Error saving image: %v", err)
	}

	log.Println("Image with icon and text saved to", outputFile)
}

func saveImage(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return err
	}

	return nil
}
