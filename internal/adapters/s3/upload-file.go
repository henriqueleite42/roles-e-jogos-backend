package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *s3Adapter) UploadFile(i *adapters.UploadFileInput) error {
	r := bytes.NewReader(i.File)

	_, err := self.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &i.StorageId,
		Key:    &i.FileName,
		Body:   r,
	})
	if err != nil {
		return err
	}

	return nil
}
