package brighten

import (
  "math"
  "errors"
  "image"
  "image/color"
)

/*
Brighten takes an image and brightens it up by `factor`, which should be a
float between -1.0 and 1.0. Providing 1.0 will make the image completely white
while a factor of -1.0 will make the image completely black.

Note that this method mutates the image.
*/
func Brighten(img image.Image, factor float64) (image.Image, error) {
  bnd := img.Bounds()

  var col color.Color
  var r, g, b, a uint32
  for x := bnd.Min.X; x < bnd.Max.X; x++ {
    for y := bnd.Min.Y; y < bnd.Max.Y; y++ {
      col = img.At(x, y)
      r, g, b, a = computeUpdatedRGBA(col, factor)

      switch input := img.(type) {
      case *image.RGBA:
        updatedColor := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
        input.SetRGBA(x, y, updatedColor)
      case *image.RGBA64:
        updatedColor := color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
        input.SetRGBA64(x, y, updatedColor)
      case *image.NRGBA:
        updatedColor := color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
        input.SetNRGBA(x, y, updatedColor)
      case *image.NRGBA64:
        updatedColor := color.NRGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
        input.SetNRGBA64(x, y, updatedColor)
      default:
        return nil, errors.New("Unable to brighten image type.")
      }
    }
  }

  return img, nil
}

func computeUpdatedRGBA(col color.Color, factor float64) (r, g, b, a uint32) {
  r, g, b, a = col.RGBA()
  r = r >> 8
  g = g >> 8
  b = b >> 8
  a = a >> 8

  var endpoint float64
  if factor > 0.0 {
    endpoint = 255.0
  } else {
    endpoint = 0.0
  }

  r = r + uint32(finalAdjustment((endpoint - float64(r)) * math.Abs(factor)))
  g = g + uint32(finalAdjustment((endpoint - float64(g)) * math.Abs(factor)))
  b = b + uint32(finalAdjustment((endpoint - float64(b)) * math.Abs(factor)))
  a = a + uint32(finalAdjustment((endpoint - float64(a)) * math.Abs(factor)))
  return
}

func finalAdjustment(diff float64) (float64) {
  adjustment := 1.05
  if diff > 0 {
    return math.Pow(diff, adjustment)
  } else {
    return -math.Pow(math.Abs(diff), adjustment)
  }
}
