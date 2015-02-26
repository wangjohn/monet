package main

import (
  "fmt"
  "os"
  "bufio"
  "image"
  "image/png"

  "github.com/wangjohn/monet/brighten"
)

const (
  filename = "/home/wangjohn/Pictures/fax_test.png"
  outputFilename = "/home/wangjohn/Pictures/fax_test_output.png"
)

func main() {
  img, _, err := decode(filename)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(brighten.StandardLuminance(img))
  img, err = brighten.Brighten(img, 0.45)
  if err != nil {
    fmt.Println(err)
  }

  w, _ := os.Create(outputFilename)
  defer w.Close()
  png.Encode(w, img)
  fmt.Println(brighten.StandardLuminance(img))
}

func decode(filename string) (image.Image, string, error) {
  f, err := os.Open(filename)
  if err != nil {
    return nil, "", err
  }
  defer f.Close()
  return image.Decode(bufio.NewReader(f))
}
