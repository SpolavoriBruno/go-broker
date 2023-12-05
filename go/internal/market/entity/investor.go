package entity

type Investor struct {
	Id            string
	AssetPosition []InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		Id: id,
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	position := i.GetAssetPosition(assetPosition.AssetId)

	if position == nil {
		i.AssetPosition = append(i.AssetPosition, *assetPosition)
	} else {
		position.Shares += assetPosition.Shares
	}

}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, position := range i.AssetPosition {
		if position.AssetId == assetId {
			return &position
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetId string
	Shares  int
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetId: assetId,
		Shares:  shares,
	}
}
