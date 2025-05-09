package s3

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/h2non/bimg"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *s3Adapter) DownloadFromUrl(i *adapters.DownloadFromUrlInput) (string, error) {
	res, err := http.Get(i.Url)
	if err != nil {
		return "", err
	}

	if res.ContentLength <= 0 {
		return "", fmt.Errorf("size not specified: %d", res.ContentLength)
	}
	if res.ContentLength > adapters.MAX_IMAGE_SIZE {
		return "", fmt.Errorf("max size exceeded: %d", res.ContentLength)
	}

	bodyReader := io.LimitReader(res.Body, res.ContentLength)
	defer res.Body.Close()

	file, err := io.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}

	ext := bimg.DetermineImageTypeName(file)

	fileName := strings.Replace(i.FileName, "{{ext}}", ext, -1)

	err = self.UploadFile(&adapters.UploadFileInput{
		StorageId: i.StorageId,
		FileName:  fileName,
		File:      file,
	})
	if err != nil {
		return "", err
	}

	return fileName, nil
}
