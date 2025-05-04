package adapters

const MAX_IMAGE_SIZE = 5 * 1024 * 1024                // 5 MB
const MAX_AUDIO_SIZE = 1 * 1024 * 1024 * 1024 * 1024  // 1 GB
const MAX_VIDEO_SIZE = 10 * 1024 * 1024 * 1024 * 1024 // 10 GB

type GetFileInput struct {
	// The bucket name
	StorageId string
	// File path + name, ex: foo/bar/xyz.png
	FileName string
}

type UploadFileInput struct {
	// The bucket name
	StorageId string
	// File path + name, ex: foo/bar/xyz.png
	FileName string
	File     []byte
}

type Storage interface {
	GetFile(i *GetFileInput) ([]byte, error)
	UploadFile(i *UploadFileInput) error
}
