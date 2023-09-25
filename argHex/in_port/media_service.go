package in_port

type MediaService interface {
	UploadMedia(file_type string, bytes []byte) (string, error)
}
