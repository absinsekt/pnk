package mixins

import (
	"crypto/sha1"
	"fmt"
	"image/png"
	"os"
	"path"
	"strings"

	"github.com/disintegration/imaging"

	"github.com/absinsekt/pnk/lib/core"
)

// Thumbnailable mixin to be embedded in a model for an entity with an image to be «thumbnailed»
type Thumbnailable struct {
	Image string `pg:"type:varchar(256)"`
}

// GetThumbnail checks if a thumbnail with given params exists or generates it on disk
func (f *Thumbnailable) GetThumbnail(width, height int) string {
	thumbPath := getThumbnailPath(f.Image, width, height)
	thumbDiskPath := path.Join(core.Config.MediaPath, thumbPath)
	srcDiskPath := path.Join(core.Config.MediaPath, f.Image)

	if !checkIfExists(thumbDiskPath) {
		generateThumb(srcDiskPath, thumbDiskPath, width, height)
	}

	return path.Join(core.Config.MediaURL, thumbPath)
}

func getThumbnailPath(imagePath string, width, height int) string {
	dir, file := path.Split(imagePath)
	fileExtension := path.Ext(file)
	fileName := strings.TrimSuffix(file, fileExtension)

	thumbsPath := path.Join("thumbs", dir)
	thumbsRawFileName := fmt.Sprintf("%s_%dx%d%s", fileName, width, height, fileExtension)

	h := sha1.New()
	h.Write([]byte(thumbsRawFileName))

	return path.Join(thumbsPath, fmt.Sprintf("%x%s", h.Sum(nil), fileExtension))
}

func checkIfExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func generateThumb(srcPath string, thumbPath string, width, height int) {
	var (
		imgOptions imaging.EncodeOption
		err        error
	)

	ext := strings.TrimPrefix(path.Ext(srcPath), ".")

	switch ext {
	case imaging.JPEG.String():
		imgOptions = imaging.JPEGQuality(80)

	case imaging.PNG.String():
		imgOptions = imaging.PNGCompressionLevel(png.BestCompression)

	default:
		imgOptions = nil
	}

	err = os.MkdirAll(path.Dir(thumbPath), 0o755)
	core.Check(err, false)

	src, err := imaging.Open(srcPath)
	core.Check(err, false)

	src = imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)

	err = imaging.Save(src, thumbPath, imgOptions)
	core.Check(err, false)
}
