package out_port

type MediaRepo interface {
	UploadMedia(file_type string, bytes []byte) (string, error)
}
