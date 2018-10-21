package main

import (
	"flag"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
)

func resizeJPEG(srcPath string, dstDir string) error {
	fp, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	filename := filepath.Base(srcPath)
	img, err := jpeg.Decode(fp)
	//r := resize.Resize(128, 128, img, resize.Lanczos3)
	r := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	dstPath := filepath.Join(dstDir, filename)
	out, err := os.Create(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dstDir, 0666)
			if err != nil {
				return err
			}
			out, err = os.Create(dstPath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer out.Close()
	jpeg.Encode(out, r, nil)
	return nil
}

func resizePNG(srcPath string, dstDir string) error {
	fp, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	filename := filepath.Base(srcPath)
	img, err := png.Decode(fp)
	//r := resize.Resize(128, 128, img, resize.Lanczos3)
	r := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	dstPath := filepath.Join(dstDir, filename)
	out, err := os.Create(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dstDir, 0666)
			if err != nil {
				return err
			}
			out, err = os.Create(dstPath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer out.Close()
	png.Encode(out, r)
	return nil
}

func resizeBMP(srcPath string, dstDir string) error {
	fp, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	filename := filepath.Base(srcPath)
	img, err := bmp.Decode(fp)
	//r := resize.Resize(128, 128, img, resize.Lanczos3)
	r := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	dstPath := filepath.Join(dstDir, filename)
	out, err := os.Create(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dstDir, 0666)
			if err != nil {
				return err
			}
			out, err = os.Create(dstPath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer out.Close()
	bmp.Encode(out, r)
	return nil
}

func resizeGIF(srcPath string, dstDir string) error {
	fp, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	filename := filepath.Base(srcPath)
	img, err := gif.Decode(fp)
	//r := resize.Resize(128, 128, img, resize.Lanczos3)
	r := resize.Thumbnail(128, 128, img, resize.Lanczos3)
	dstPath := filepath.Join(dstDir, filename)
	out, err := os.Create(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dstDir, 0666)
			if err != nil {
				return err
			}
			out, err = os.Create(dstPath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer out.Close()
	gif.Encode(out, r, nil)
	return nil
}

func resizeSVG(srcPath string, dstDir string) error {
	//w := 500
	//h := 500
	//c := svg.New(os.Stdout)
	//c.Start(w, h)
	//c.Circle(w/2, h/2, 100)
	//c.End()
	return nil
}

// ResizePic 生成缩略图
func ResizePic(srcPath string, dstDir string) error {
	ext := path.Ext(srcPath)
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		return resizeJPEG(srcPath, dstDir)
	case ".png":
		return resizePNG(srcPath, dstDir)
	case ".bmp":
		return resizeBMP(srcPath, dstDir)
	case ".gif":
		return resizeGIF(srcPath, dstDir)
	case ".svg":
		return resizeSVG(srcPath, dstDir)
	default:
		return fmt.Errorf("unsupport ext: %s", ext)
	}
}

var (
	srcFile = flag.String("src", "test.jpg", "源文件")
	dstDir  = flag.String("o", "dist", "输出文件目录")
)

func main() {
	flag.Parse()
	err := ResizePic(*srcFile, *dstDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("resize success")
}
