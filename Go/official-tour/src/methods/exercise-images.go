package methods

import
(
	"golang-tour/mypic"
	"image"
	"image/color"
)

type Image struct{
	w int
	h int
}

func (self Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.w, self.h)
}

func (self Image) At(x, y int) color.Color {
	r := (uint8)((float64)(x) / (float64)(self.w) * 255.0)
	g := (uint8)((float64)(y) / (float64)(self.h) * 255.0)
	b := (uint8)((float64)(x * y) / (float64)(self.w * self.h) * 255.0)
	return color.RGBA{r, g, b, 255}
}

func ExerciseImages() {
	m := Image{255, 255}
	mypic.ShowImage(m)
}