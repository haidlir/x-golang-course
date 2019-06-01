package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	nmodel "github.com/haidlir/x-golang-course/021-simple-rest-api/model"
)

const (
	errorDecodingJSONReq = "decoding JSON Error"
	errorAddNewSiswa     = "unable to add new siswa"
	errorUpdatingSiswa   = "unable to update siswa"
	emptySiswa           = "siswa is empty"
	errorParsingID       = "unable to parse id"
	errorFindingID       = "id not found"
)

// DB is the method signature for every DB.
type DB interface {
	GetAllSiswa() []nmodel.Siswa
	GetDetailSiswa(id int) *nmodel.Siswa
	AddSiswa(nmodel.Siswa) (*nmodel.Siswa, error)
	UpdateSiswa(id int, data nmodel.Siswa) (*nmodel.Siswa, error)
	DeleteSiswa(id int) error
}

// Handler is the handler object
type Handler struct {
	db DB
}

// AddSiswa handles the POST SISWA
func (h *Handler) AddSiswa(w http.ResponseWriter, r *http.Request) {
	// Get the Body
	decodedBody := nmodel.Siswa{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decodedBody)
	if err != nil {
		singleErrResp := nmodel.Error{
			Title:  errorDecodingJSONReq,
			Detail: errorDecodingJSONReq,
		}
		sendNOKResponse(http.StatusBadRequest, []nmodel.Error{singleErrResp}, w)
		return
	}
	addedSiswa, err := h.db.AddSiswa(decodedBody)
	if err != nil {
		singleErrResp := nmodel.Error{
			Title:  errorAddNewSiswa,
			Detail: fmt.Sprintf("%v", err),
		}
		sendNOKResponse(http.StatusBadRequest, []nmodel.Error{singleErrResp}, w)
		return
	}
	sendOKResponse(http.StatusCreated, addedSiswa, nil, w)
	return
}

// GetAllSiswa handles the GET ALL SISWA
func (h *Handler) GetAllSiswa(w http.ResponseWriter, r *http.Request) {
	daftarSiswa := h.db.GetAllSiswa()
	if len(daftarSiswa) <= 0 {
		singleErrResp := nmodel.Error{
			Title:  emptySiswa,
			Detail: emptySiswa,
		}
		sendNOKResponse(http.StatusInternalServerError, []nmodel.Error{singleErrResp}, w)
		return
	}
	sendOKResponse(http.StatusOK, daftarSiswa, nil, w)
	return
}

// GetDetailSiswa handles the GET SPECIFIC SISWA
func (h *Handler) GetDetailSiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id int
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			singleErrResp := nmodel.Error{
				Title:  errorParsingID,
				Detail: errorParsingID,
			}
			sendNOKResponse(http.StatusInternalServerError, []nmodel.Error{singleErrResp}, w)
			return
		}
		id = convertedVal
	}
	siswa := h.db.GetDetailSiswa(id)
	if siswa == nil {
		singleErrResp := nmodel.Error{
			Title:  errorFindingID,
			Detail: errorFindingID,
		}
		sendNOKResponse(http.StatusInternalServerError, []nmodel.Error{singleErrResp}, w)
		return
	}
	sendOKResponse(http.StatusOK, siswa, nil, w)
	return
}

// DeleteSiswa handles the DELETE SPECIFIC SISWA
func (h *Handler) DeleteSiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id int
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			singleErrResp := nmodel.Error{
				Title:  errorParsingID,
				Detail: errorParsingID,
			}
			sendNOKResponse(http.StatusInternalServerError, []nmodel.Error{singleErrResp}, w)
			return
		}
		id = convertedVal
	}
	err := h.db.DeleteSiswa(id)
	if err != nil {
		singleErrResp := nmodel.Error{
			Title:  errorFindingID,
			Detail: errorFindingID,
		}
		sendNOKResponse(http.StatusInternalServerError, []nmodel.Error{singleErrResp}, w)
		return
	}
	resp := map[string]string{"message": fmt.Sprintf("siswa id %v deleted successfully", id)}
	sendOKResponse(http.StatusOK, resp, nil, w)
	return
}

// UpdateSiswa handles the PUT SPECIFIC SISWA
func (h *Handler) UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	// Get ID
	vars := mux.Vars(r)
	var id int
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			singleErrResp := nmodel.Error{
				Title:  errorParsingID,
				Detail: errorParsingID,
			}
			sendNOKResponse(http.StatusInternalServerError, []nmodel.Error{singleErrResp}, w)
			return
		}
		id = convertedVal
	}
	// Get the Body
	decodedBody := nmodel.Siswa{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decodedBody)
	if err != nil {
		singleErrResp := nmodel.Error{
			Title:  errorDecodingJSONReq,
			Detail: errorDecodingJSONReq,
		}
		sendNOKResponse(http.StatusBadRequest, []nmodel.Error{singleErrResp}, w)
		return
	}
	updatedSiswa, err := h.db.UpdateSiswa(id, decodedBody)
	if err != nil {
		singleErrResp := nmodel.Error{
			Title:  errorUpdatingSiswa,
			Detail: fmt.Sprintf("%v", err),
		}
		sendNOKResponse(http.StatusBadRequest, []nmodel.Error{singleErrResp}, w)
		return
	}
	sendOKResponse(http.StatusOK, updatedSiswa, nil, w)
	return
}

// NewHandler return handler
func NewHandler(db DB) *Handler {
	handler := Handler{db}
	return &handler
}

func sendOKResponse(statusCode int, data interface{}, meta nmodel.MetaResponse, w http.ResponseWriter) error {
	response := nmodel.ResponseFormat{}
	if data != nil {
		response.Data = data
	}
	if meta != nil {
		response.Meta = meta
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return fmt.Errorf("Unable to encode response: %v", err)
	}
	w.Write(encodedResponse)
	return nil
}

func sendNOKResponse(statusCode int, errResp []nmodel.Error, w http.ResponseWriter) error {
	response := nmodel.ResponseFormat{}
	if errResp != nil {
		response.Errors = errResp
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	encodedResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return fmt.Errorf("Unable to encode response: %v", err)
	}
	w.Write(encodedResponse)
	return nil
}
