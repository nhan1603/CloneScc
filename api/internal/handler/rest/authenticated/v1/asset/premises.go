package asset

import (
	"log"
	"net/http"
	"strings"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	assetCtrl "github.com/nhan1603/CloneScc/api/internal/controller/asset"
)

// Premises is the response for each item of premises
type Premise struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	PremisesCode string `json:"premises_code"`
	Description  string `json:"description"`
	CCTVCount    int    `json:"cctv_count"`
}

// PremisesResponse is the response for GetPremises
type PremisesResponse struct {
	Items []Premise `json:"items"`
}

// GetPremises retrieves a list of premises
func (h Handler) GetPremises() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetPremises] START processing requests")
		ctx := r.Context()
		nameQuery := strings.TrimSpace(r.URL.Query().Get("name"))

		premises, err := h.assetCtrl.GetPremises(ctx, assetCtrl.GetPremisesInput{
			Name: nameQuery,
		})
		if err != nil {
			return err
		}

		items := make([]Premise, len(premises))
		for idx, item := range premises {
			items[idx] = Premise{
				ID:           item.ID,
				Name:         item.Name,
				Location:     item.Location,
				PremisesCode: item.PremisesCode,
				Description:  item.Description,
				CCTVCount:    item.CCTVCount,
			}
		}

		httpserver.RespondJSON(w, PremisesResponse{
			Items: items,
		})
		return nil
	})
}
