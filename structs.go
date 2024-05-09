package npglib

import (
	"image"
	"image/color"
	"strings"
)

// A 2D slice of color.RGBA, meant to be used as a sprite container
type Sprite struct {
	Size   [2]int
	pixels [][]color.RGBA
}

func (spr *Sprite) Generate(initColor color.RGBA) {
	spr.pixels = [][]color.RGBA{}
	for range spr.Size[1] {
		var o []color.RGBA
		for range spr.Size[0] {
			o = append(o, initColor)
		}
		spr.pixels = append(spr.pixels, o)
	}
}

// =< 0: whitespace is black, otherwise white
// => 1: 'r' is red, 'g' is green, 'b' is blue; whitespace is interpreted as black; otherwise white
func (spr *Sprite) GenerateFromString(str string, mode int) {
	lines := strings.Split(str, "\n")
	spr.pixels = [][]color.RGBA{}
	for i1, l := range lines {
		if i1 > spr.Size[1]-1 {
			continue
		}
		var o []color.RGBA
		for i2, char := range l {
			if i2 > spr.Size[0]-1 {
				continue
			}
			var col color.RGBA
			if mode <= 0 {
				if char == '\t' || char == ' ' {
					col = color.RGBA{0, 0, 0, 255}
				} else {
					col = color.RGBA{255, 255, 255, 255}
				}
			} else {
				switch char {
				case 'r':
					col = color.RGBA{255, 0, 0, 255}
				case 'g':
					col = color.RGBA{0, 255, 0, 255}
				case 'b':
					col = color.RGBA{0, 0, 255, 255}
				case '\t', ' ':
					col = color.RGBA{0, 0, 0, 255}
				default:
					col = color.RGBA{255, 255, 255, 255}
				}
			}

			o = append(o, col)
		}
		spr.pixels = append(spr.pixels, o)
	}
}

func (spr *Sprite) GenerateFromSprite(sprite [][]color.Color, overwriteOwnSize bool) {
	spr.pixels = [][]color.RGBA{}
	if overwriteOwnSize {
		spr.Size = [2]int{len(sprite[0]), len(sprite)}
	}
	for iy, y := range sprite {
		if iy > spr.Size[1]-1 {
			continue
		}
		var o []color.RGBA
		for ix, x := range y {
			if ix > spr.Size[0]-1 {
				continue
			}
			var col = color.RGBA{0, 0, 0, 255}

			r, g, b, a := x.RGBA()

			col.A = uint8(a)
			col.R = uint8(r / a)
			col.G = uint8(g / a)
			col.B = uint8(b / a)

			o = append(o, col)
		}
		spr.pixels = append(spr.pixels, o)
	}
}

func (spr *Sprite) GenerateFromImage(img image.Image, overwriteOwnSize bool) {
	spr.pixels = [][]color.RGBA{}
	imgXY := [2]int{img.Bounds().Max.X, img.Bounds().Max.Y}
	if overwriteOwnSize {
		spr.Size = imgXY
	}
	for y := range imgXY[1] {
		if y > spr.Size[1]-1 {
			continue
		}
		var o []color.RGBA
		for x := range imgXY[0] {
			if x > spr.Size[0]-1 {
				continue
			}
			var col = color.RGBA{0, 0, 0, 255}

			r, g, b, a := img.At(x, y).RGBA()

			col.A = uint8(a)
			col.R = uint8(r / a)
			col.G = uint8(g / a)
			col.B = uint8(b / a)

			o = append(o, col)
		}
		spr.pixels = append(spr.pixels, o)
	}
}

func (spr Sprite) DrawSpriteOnBoard(offsetX, offsetY int, board *[][]color.RGBA) {
	mBoard := *board

	for iy, y := range spr.pixels {
		iy += offsetY
		if iy > len(mBoard) {
			continue
		}

		for ix, x := range y {
			ix += offsetX
			if ix > len(mBoard[iy]) {
				continue
			}
			mBoard[iy][ix] = x
		}
	}

	*board = mBoard
}

func (spr Sprite) GetColors() map[color.RGBA]int {
	out := make(map[color.RGBA]int)
	for _, y := range spr.pixels {
		for _, x := range y {
			out[x]++
		}
	}
	return out
}

// Returns true if the pixel is found, else false
func (spr Sprite) GetPixel(x, y int) (color.RGBA, bool) {
	if y < len(spr.pixels) {
		if x < len(spr.pixels[y]) {
			return spr.pixels[y][x], true
		}
	}
	return color.RGBA{0, 0, 0, 0}, false
}

func (spr *Sprite) SetPixel(x, y int, color color.RGBA) bool {
	if y < len(spr.pixels) {
		if x < len(spr.pixels[y]) {
			spr.pixels[y][x] = color
			return true
		}
	}
	return false
}

// A 3D slice of color.RGBA, meant to be used as a container for a collection of colors in a 3D space, such as a voxel model
type VoxelSprite struct {
	Size   [3]int
	pixels [][][]color.RGBA
}

func (spr *VoxelSprite) Generate(initColor color.RGBA) {
	spr.pixels = [][][]color.RGBA{}
	for range spr.Size[2] {
		var o [][]color.RGBA
		for range spr.Size[1] {
			var o2 []color.RGBA
			for range spr.Size[0] {
				o2 = append(o2, initColor)
			}
			o = append(o, o2)
		}
		spr.pixels = append(spr.pixels, o)
	}
}

func (spr VoxelSprite) GetColors() map[color.RGBA]int {
	out := make(map[color.RGBA]int)
	for _, z := range spr.pixels {
		for _, y := range z {
			for _, x := range y {
				out[x]++
			}
		}
	}
	return out
}

// Returns true if the voxel is found, else false
func (spr VoxelSprite) GetVoxel(x, y, z int) (color.RGBA, bool) {
	if z < len(spr.pixels) {
		if y < len(spr.pixels[z]) {
			if x < len(spr.pixels[y]) {
				return spr.pixels[z][y][x], true
			}
		}
	}
	return color.RGBA{0, 0, 0, 0}, false
}

func (spr *VoxelSprite) SetVoxel(x, y, z int, color color.RGBA) bool {
	if z < len(spr.pixels) {
		if y < len(spr.pixels[z]) {
			if x < len(spr.pixels[y]) {
				spr.pixels[z][y][x] = color
				return true
			}
		}
	}
	return false
}
