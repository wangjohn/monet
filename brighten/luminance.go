package brighten

import (
  "fmt"
  "image"
  "math"
)

func StandardLuminance(img image.Image) (float64) {
  bnd := img.Bounds()

  count := 0
  totalLum := 0.0
  for x := bnd.Min.X; x < bnd.Max.X; x++ {
    for y := bnd.Min.Y; y < bnd.Max.Y; y++ {
      r, g, b, _ := img.At(x, y).RGBA()

      totalLum += Luminance(r, g, b)
      count++
    }
  }
  return totalLum / float64(count)
}

func Luminance(r, g, b uint32) (float64) {
  return 0.2126 * adjustValue(r) + 0.7152 * adjustValue(g) + 0.0722 * adjustValue(b)
}

func adjustValue(val uint32) (float64) {
  v := float64(val >> 8) / 255.0
  if v < 0.03928 {
    return v / 12.92
  } else {
    return math.Pow((v + 0.055) / 1.055, 2.4)
  }
}
