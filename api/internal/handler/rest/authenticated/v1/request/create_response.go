package request

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
)

type responseReq struct {
	RequestID string `json:"requestID"`
	Message   string `json:"message"`
}

func (r responseReq) ValidateAndMap() (requests.CreateResponseInput, error) {
	reqID, _ := strconv.ParseInt(r.RequestID, 10, 64)
	if reqID <= 0 {
		return requests.CreateResponseInput{}, webErrInvalidRequestID
	}

	message := strings.TrimSpace(r.Message)
	if message == "" {
		return requests.CreateResponseInput{}, webErrMessageIsRequired
	}

	return requests.CreateResponseInput{
		RequestID: reqID,
		Message:   message,
	}, nil
}

const (
	maxMemory = 100 << 20 // 100MB
	maxMedia  = 20
)

// CreateResponse create a new response for the request
func (h Handler) CreateResponse() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[CreateResponse] START processing requests")
		ctx := r.Context()

		if err := r.ParseMultipartForm(maxMemory); err != nil {
			return webErrExceedMaxMemoryAllowed
		}

		dataField := r.FormValue("data")
		var req responseReq
		if err := json.Unmarshal([]byte(dataField), &req); err != nil {
			log.Printf("[CreateResponse] Unmarshal error %v", err.Error())
			return err
		}

		inp, err := req.ValidateAndMap()
		if err != nil {
			log.Printf("[CreateResponse] ValidateAndMap error %v", err.Error())
			return err
		}

		// Retrieve all uploaded files
		files := r.MultipartForm.File["file"]
		if len(files) > maxMedia {
			return webErrExceedMaxMediaAllowed
		}

		inp.Files = files
		if err := h.requestCtrl.CreateResponse(ctx, inp); err != nil {
			return convertCtrlErr(err)
		}

		log.Printf("Send resolved signal \n")
		h.requestCtrl.Push(ctx, inp.RequestID)

		httpserver.RespondJSON(w, successResponse{Success: true})
		return nil
	})
}
