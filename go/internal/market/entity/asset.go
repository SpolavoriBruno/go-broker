package entity

type Asset struct {
	Name   string
	Id     string
	Value  float32
	Shares int
}

func NewAsset(id, name string, value float32) *Asset {
	return &Asset{
		Name:  name,
		Id:    id,
		Value: value,
	}
}
