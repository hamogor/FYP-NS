package main

import (
	"bufio"
	"github.com/faiface/pixel"
	"github.com/tidwall/gjson"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

type Assets struct {
	Sheets Sheets
	Env    map[string][]pixel.Rect
	Anims  map[string]*Animation
}

type Sheets struct {
	Environment pixel.Picture
	Sprites     pixel.Picture
}

type Unparsed struct {
	img  image.Image
	rect pixel.Rect
	json string
}

func (g *Game) buildAssets() {
	envsheet, env := BuildEnvironmentSheet()
	spritesheet, anims := BuildSpriteSheet()
	g.Assets = &Assets{
		Sheets: Sheets{
			Environment: envsheet,
			Sprites:     spritesheet,
		},
		Env:   env,
		Anims: anims,
	}
}

func BuildEnvironmentSheet() (pixel.Picture, map[string][]pixel.Rect) {
	outputLen, unparsedTiles := GetUnparsedAssets(tilePaths)
	outputImg := image.NewRGBA(image.Rectangle{Max: image.Point{X: outputLen, Y: int(TileH)}})
	BuildImage(unparsedTiles, outputImg, tileSheetOutputPath)
	sheet := GetPixelPicture(tileSheetOutputPath)
	environmentMap := make(map[string][]pixel.Rect, len(tilePaths))
	for i := range unparsedTiles {
		for j := 0; j < getJsonLength(unparsedTiles[i].json); j++ {
			key := getJsonName(unparsedTiles[i].json)
			environmentMap[key] = append(environmentMap[key], pixel.R(nextPos, 0, nextPos+TileW, TileH))
			nextPos += float64(TileW)
		}
	}
	nextPos = 0
	return sheet, environmentMap
}

func BuildSpriteSheet() (pixel.Picture, map[string]*Animation) {
	outputLen, unparsedAnims := GetUnparsedAssets(spritePaths)
	outputImg := image.NewRGBA(image.Rectangle{Max: image.Point{X: outputLen, Y: int(TileH)}})
	BuildImage(unparsedAnims, outputImg, spriteSheetOutputPath)
	sheet := GetPixelPicture(spriteSheetOutputPath)
	animMap := make(map[string]*Animation, len(spritePaths))

	frames := make([]*Frame, 0)
	for i := range unparsedAnims {
		for j := 0; j < getJsonLength(unparsedAnims[i].json); j++ {
			frame := &Frame{
				Rect:     pixel.R(nextPos, 0, nextPos+SpriteW, 48),
				Duration: getJsonDuration(unparsedAnims[i].json, j),
			}
			frames = append(frames, frame)
			nextPos += SpriteW
		}
		key := getJsonName(unparsedAnims[i].json)
		anim := &Animation{
			Frames: frames,
			Loop:   getJsonLoop(unparsedAnims[i].json),
			Active: false,
		}
		animMap[key] = anim
	}
	nextPos = 0
	return sheet, animMap
}

func GetUnparsedAssets(paths []string) (int, []*Unparsed) {
	slice := make([]*Unparsed, 0)
	x := 0
	for i := range paths {
		unparsed := &Unparsed{}
		unparsed.img = GetUnparsedImg(paths[i])
		unparsed.json = GetUnparsedJson(paths[i])
		x += unparsed.img.Bounds().Max.X
		slice = append(slice, unparsed)
	}
	return x, slice
}

func BuildImage(unparsed []*Unparsed, outputImg *image.RGBA, path string) {
	for i := range unparsed {
		if i == 0 {
			rect = unparsed[i].img.Bounds()
			draw.Draw(outputImg, rect, unparsed[i].img, image.Point{X: 0, Y: 0}, draw.Src)
			nextPos = float64(rect.Max.X)
			unparsed[i].rect = GetUnparsedRect(rect)
		} else {
			nextPoint := image.Point{X: int(nextPos), Y: 0}
			rect = image.Rectangle{
				Min: nextPoint,
				Max: nextPoint.Add(unparsed[i].img.Bounds().Size()),
			}
			draw.Draw(outputImg, rect, unparsed[i].img, image.Point{X: 0, Y: 0}, draw.Src)
			nextPos = float64(rect.Max.X)
			unparsed[i].rect = GetUnparsedRect(rect)
		}
	}
	nextPos = 0
	out, _ := os.Create(path)
	_ = png.Encode(out, outputImg)

}

func GetUnparsedRect(rect image.Rectangle) pixel.Rect {
	return pixel.Rect{
		Min: pixel.Vec{X: float64(rect.Bounds().Min.X), Y: float64(rect.Bounds().Min.Y)},
		Max: pixel.Vec{X: float64(rect.Bounds().Max.X), Y: float64(rect.Bounds().Max.Y)},
	}
}

func GetUnparsedImg(asset string) image.Image {
	file, err := os.Open(pngPath + asset + ".png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
	}
	return img
}

func GetPixelPicture(path string) pixel.Picture {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
	}
	return pixel.PictureDataFromImage(img)
}

func GetUnparsedJson(asset string) string {
	jsonFile, err := os.Open(jsonPath + asset + ".json")
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(jsonFile)
	json := ""
	for scanner.Scan() {
		json += scanner.Text()
	}
	return json
}

func getJsonLength(json string) int {
	return int(gjson.Get(json, "length").Int())
}

func getJsonName(json string) string {
	return gjson.Get(json, "name").String()
}

func getJsonDuration(json string, i int) float64 {
	durations := gjson.Get(json, "duration").Array()
	return durations[i].Float()
}

func getJsonLoop(json string) bool {
	return gjson.Get(json, "loop").Bool()
}

func rectToRectangle(r image.Rectangle) pixel.Rect {
	return pixel.R(float64(r.Min.X), float64(r.Min.Y), float64(r.Max.X), float64(r.Max.Y))
}
