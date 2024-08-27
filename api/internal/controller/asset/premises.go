package asset

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
)

// GetPremisesInput is the search input for GetPremises
type GetPremisesInput struct {
	Name string
}

// GetPremises get list of premises with provided prameters
func (i impl) GetPremises(ctx context.Context, input GetPremisesInput) ([]model.Premises, error) {
	return i.repo.Asset().GetPremises(ctx, asset.GetPremisesInput{
		Name: input.Name,
	})
}
