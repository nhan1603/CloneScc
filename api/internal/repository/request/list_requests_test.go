package request

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func Test_impl_GetRequests(t *testing.T) {
	tcs := map[string]struct {
		input  model.GetRequestsInput
		expRes []model.RequestSummary
		expErr error
	}{
		"success|get request with premise": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				Limit:      50,
				Page:       1,
				AssigneeID: 0,
			},
			expRes: []model.RequestSummary{
				{
					ID:              400,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              401,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              402,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "William",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "RESOLVED",
				},
				{
					ID:              403,
					AlertID:         301,
					Alert:           "CCTV 1-301",
					AlertType:       "Unauthorized Access",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "William",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              404,
					AlertID:         301,
					Alert:           "CCTV 1-301",
					AlertType:       "Unauthorized Access",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              405,
					AlertID:         302,
					Alert:           "CCTV 1-302",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "RESOLVED",
				},
				{
					ID:              406,
					AlertID:         303,
					Alert:           "CCTV 2-303",
					AlertType:       "Unauthorized Access",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "William",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
			},
		},
		"success|get request with premise for assignee": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				Limit:      50,
				Page:       1,
				AssigneeID: 101,
			},
			expRes: []model.RequestSummary{
				{
					ID:              400,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              401,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              405,
					AlertID:         302,
					Alert:           "CCTV 1-302",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "RESOLVED",
				},
			},
		},
		"success|limit 2, page 2": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				Limit:      2,
				Page:       2,
				AssigneeID: 0,
			},
			expRes: []model.RequestSummary{
				{
					ID:              402,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "William",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "RESOLVED",
				},
				{
					ID:              403,
					AlertID:         301,
					Alert:           "CCTV 1-301",
					AlertType:       "Unauthorized Access",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "William",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
			},
		},
		"not found": {
			input: model.GetRequestsInput{
				PremiseID:  200,
				Limit:      50,
				Page:       1,
				AssigneeID: 0,
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Given:
				ctx := context.Background()
				test.LoadSqlTestFile(t, tx, "testdata/requests.sql")

				// When:
				repo := New(tx)
				_, _, err := repo.GetRequests(ctx, tc.input)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
				}
			})
		})
	}
}
