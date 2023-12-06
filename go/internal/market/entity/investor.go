package entity

import "fmt"

type Investor struct {
	Id            string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		Id: id,
	}
}

/** DEPRECATED - use UpdateAssetPosition insted */
func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	position := i.GetAssetPosition(assetPosition.AssetId)
	if position == nil {
		i.AssetPosition = append(i.AssetPosition, assetPosition)
	}
}

func (i *Investor) UpdateAssetPosition(assetId string, shares int) {
	position := i.GetAssetPosition(assetId)

	if position == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetId, shares))
		fmt.Printf("Investor %s has %d of %s", i.Id, shares, assetId)
	} else {
		position.Shares += shares
		fmt.Printf("Investor %s has %d of %s\n", i.Id, position.Shares, position.AssetId)
	}
}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, position := range i.AssetPosition {
		if position.AssetId == assetId {
			return position
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
