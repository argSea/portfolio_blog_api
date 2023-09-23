package out_port

type MediaRepo interface {
	UploadMedia() (string, error)
}
