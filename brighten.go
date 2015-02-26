package monet

import (
  "math"
  "errors"
  "image"
)

/*
Brighten takes an image and brightens it up by `factor`, which should be a
float between -1.0 and 1.0. Providing 1.0 will make the image completely white
while a factor of -1.0 will make the image completely black.

Note that this method mutates the image.
*/
func Brighten(img *image.Image, factor float64) (*image.Image, error) {
  bnd := img.Bounds()

  var color image.Color
  var r, g, b, a uint32
  for x := bnd.Min.X; x < bnd.Max.X; x++ {
    for y := bnd.Min.Y; y < bnd.Max.Y; y++ {
      color = img.At(x, y)
      r, g, b, a = computeUpdatedRGBA(color, factor)

      switch input := img.(type) {
      case *image.RGBA:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetRGBA(x, y, updatedColor)
      case *image.RGBA64:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetRGBA64(x, y, updatedColor)
      case *image.Gray:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetGray(x, y, updatedColor)
      case *image.Gray16:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetGray16(x, y, updatedColor)
      case *image.NRGBA:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetNRGBA(x, y, updatedColor)
      case *image.NRGBA64:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetNRGBA64(x, y, updatedColor)
      case *image.Alpha:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetAlpha(x, y, updatedColor)
      case *image.Alpha16:
        updatedColor := input.ColorModel().RGBA(r, g, b, a)
        input.SetAlpha16(x, y, updatedColor)
      default:
        return nil, errors.New("Unable to brighten image type.")
      }
    }
  }

  return img, nil
}

func computeUpdatedRGBA(color image.Color, factor float64) (r, g, b, a uint32) {
  r, g, b, a := color.RGBA()

  var endpoint uint32
  if factor > 0.0 {
    endpoint = 255
  } else {
    endpoint = 0
  }

  r = r + (endpoint - r) * math.Abs(factor)
  g = g + (endpoint - g) * math.Abs(factor)
  b = b + (endpoint - b) * math.Abs(factor)
  a = a + (endpoint - a) * math.Abs(factor)
  return
}
