package entity

type Asset struct {
	Id           string
	MarketVolume int
	Name         string
	Shares       int
	Value        float32
}

func NewAsset(id, name string, marketVolume int) *Asset {
	return &Asset{
		Id:           id,
		MarketVolume: marketVolume,
		Name:         name,
	}
}
