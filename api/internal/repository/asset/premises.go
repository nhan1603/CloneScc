package asset

import (
	"context"
	"fmt"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetPremisesInput represents for each item of an input
type GetPremisesInput struct {
	Name string
}

// GetPremises returns list of premises with provided parameters
func (i impl) GetPremises(ctx context.Context, input GetPremisesInput) ([]model.Premises, error) {
	qms := []qm.QueryMod{
		qm.OrderBy(dbmodel.PremiseColumns.Name),
	}

	if input.Name != "" {
		qms = append(qms, qm.Where(dbmodel.PremiseColumns.Name+" ILIKE ?", fmt.Sprintf("%%%s%%", input.Name)))
	}

	premises, err := dbmodel.Premises(qms...).All(ctx, i.dbConn)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	if len(premises) == 0 {
		return nil, nil
	}

	// Translate the ORM object to internal model object
	result := make([]model.Premises, len(premises))
	for idx, record := range premises {
		result[idx] = model.Premises{
			ID:           record.ID,
			Name:         record.Name,
			Location:     record.Location,
			PremisesCode: record.PremisesCode,
			Description:  record.Description,
			CCTVCount:    record.CCTVCount,
		}
	}

	return result, nil
}
