package in_port

type MediaService interface {
	UploadMedia(mime_type string, bytes []byte) (string, error)
}
