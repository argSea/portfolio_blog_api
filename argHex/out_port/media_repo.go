package out_port

type MediaRepo interface {
	UploadMedia(mime_type string, bytes []byte) (string, error)
}
