package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *s3Adapter) GetFile(i *adapters.GetFileInput) ([]byte, error) {
	file, err := self.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &i.StorageId,
		Key:    &i.FileName,
	})
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(file.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
