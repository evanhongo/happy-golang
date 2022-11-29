package service

type IJobService interface {
	CompressImage(data []byte) (string, error)
}
