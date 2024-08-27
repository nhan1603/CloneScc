package asset

import (
	"context"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_GetPremises(t *testing.T) {
	type mockAssetRepo struct {
		expCall bool
		input   asset.GetPremisesInput
		output  []model.Premises
		err     error
	}

	tcs := map[string]struct {
		input         GetPremisesInput
		mockAssetRepo mockAssetRepo
		expRes        []model.Premises
		expErr        error
	}{
		"success": {
			input: GetPremisesInput{
				Name: "MyTower",
			},
			mockAssetRepo: mockAssetRepo{
				expCall: true,
				input: asset.GetPremisesInput{
					Name: "MyTower",
				},
				output: []model.Premises{
					{
						ID:           50,
						Name:         "Sunrise Tower",
						Location:     "307/12 Nguyen Van Troi St, W1, Tan Binh",
						PremisesCode: "P001",
						Description:  "Sunrise Tower",
						CCTVCount:    4,
					},
					{
						ID:           51,
						Name:         "Bitexco Financial Tower",
						Location:     "2 Hai Ba Trung St, Ben Nghe, District 1",
						PremisesCode: "P002",
						Description:  "Bitexco Financial Tower",
						CCTVCount:    4,
					},
				},
			},
			expRes: []model.Premises{
				{
					ID:           50,
					Name:         "Sunrise Tower",
					Location:     "307/12 Nguyen Van Troi St, W1, Tan Binh",
					PremisesCode: "P001",
					Description:  "Sunrise Tower",
					CCTVCount:    4,
				},
				{
					ID:           51,
					Name:         "Bitexco Financial Tower",
					Location:     "2 Hai Ba Trung St, Ben Nghe, District 1",
					PremisesCode: "P002",
					Description:  "Bitexco Financial Tower",
					CCTVCount:    4,
				},
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			mockAssetRepo := &asset.MockRepository{}
			if tc.mockAssetRepo.expCall {
				mockAssetRepo.ExpectedCalls = []*mock.Call{
					mockAssetRepo.On("GetPremises",
						ctx,
						tc.mockAssetRepo.input,
					).Return(
						tc.mockAssetRepo.output,
						tc.mockAssetRepo.err,
					),
				}
			}
			repo := repository.MockRegistry{}
			repo.ExpectedCalls = []*mock.Call{
				repo.On("Asset").Return(mockAssetRepo),
			}
			s := New(&repo)
			// When:
			actRes, err := s.GetPremises(ctx, tc.input)
			// Then:
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, len(tc.expRes), len(actRes))
				require.Equal(t, tc.expRes, actRes)
			}
		})
	}
}
