package handler

import (
	"encoding/json"
	"fmt"
	"log"
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
	resp := nmodel.NewResponseFormat()
	// Get the Body
	decodedBody := nmodel.Siswa{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decodedBody)
	if err != nil {
		resp.AddError(errorDecodingJSONReq, errorDecodingJSONReq)
		sendResponse(http.StatusBadRequest, resp, w, r)
		return
	}
	addedSiswa, err := h.db.AddSiswa(decodedBody)
	if err != nil {
		resp.AddError(errorAddNewSiswa, err.Error())
		sendResponse(http.StatusBadRequest, resp, w, r)
		return
	}
	resp.SetData(addedSiswa)
	sendResponse(http.StatusCreated, resp, w, r)
	return
}

// GetAllSiswa handles the GET ALL SISWA
func (h *Handler) GetAllSiswa(w http.ResponseWriter, r *http.Request) {
	resp := nmodel.NewResponseFormat()
	daftarSiswa := h.db.GetAllSiswa()
	if len(daftarSiswa) <= 0 {
		resp.AddError(emptySiswa, emptySiswa)
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	resp.SetData(daftarSiswa)
	sendResponse(http.StatusOK, resp, w, r)
	return
}

// GetDetailSiswa handles the GET SPECIFIC SISWA
func (h *Handler) GetDetailSiswa(w http.ResponseWriter, r *http.Request) {
	resp := nmodel.NewResponseFormat()
	id, err := getVarsID(r)
	if err != nil {
		resp.AddError(errorParsingID, errorParsingID)
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	siswa := h.db.GetDetailSiswa(id)
	if siswa == nil {
		resp.AddError(errorFindingID, errorFindingID)
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	resp.SetData(siswa)
	sendResponse(http.StatusOK, resp, w, r)
	return
}

// DeleteSiswa handles the DELETE SPECIFIC SISWA
func (h *Handler) DeleteSiswa(w http.ResponseWriter, r *http.Request) {
	resp := nmodel.NewResponseFormat()
	id, err := getVarsID(r)
	if err != nil {
		resp.AddError(errorParsingID, errorParsingID)
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	err = h.db.DeleteSiswa(id)
	if err != nil {
		resp.AddError(errorFindingID, errorFindingID)
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	resp.SetData(map[string]string{"message": fmt.Sprintf("siswa id %v deleted successfully", id)})
	sendResponse(http.StatusOK, resp, w, r)
	return
}

// UpdateSiswa handles the PUT SPECIFIC SISWA
func (h *Handler) UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	resp := nmodel.NewResponseFormat()
	// Get ID
	id, err := getVarsID(r)
	if err != nil {
		resp.AddError(errorParsingID, errorParsingID)
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	// Get the Body
	decodedBody := nmodel.Siswa{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&decodedBody)
	if err != nil {
		resp.AddError(errorDecodingJSONReq, errorDecodingJSONReq)
		sendResponse(http.StatusBadRequest, resp, w, r)
		return
	}
	updatedSiswa, err := h.db.UpdateSiswa(id, decodedBody)
	if err != nil {
		resp.AddError(errorUpdatingSiswa, err.Error())
		sendResponse(http.StatusBadRequest, resp, w, r)
		return
	}
	resp.SetData(updatedSiswa)
	sendResponse(http.StatusOK, resp, w, r)
	return
}

// NewHandler return handler
func NewHandler(db DB) *Handler {
	handler := Handler{db}
	return &handler
}

func sendResponse(statusCode int, resp *nmodel.ResponseFormat, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	encodedResponse, err := resp.EncodeToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		log.Printf("Source: %v| Destination: %v| ResponseCode: %v| ResponseLen: %v", r.RemoteAddr, r.RequestURI, statusCode, "error while encoding response")
		return fmt.Errorf("unable to encode JSON: %v", err)
	}
	w.Write(encodedResponse)
	log.Printf("| Source: %v | Destination: %v | Mehod: %v | ResponseCode: %v | ResponseLen: %v", r.RemoteAddr, r.RequestURI, r.Method, statusCode, len(encodedResponse))
	return nil
}

func getVarsID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			return id, err
		}
		id = convertedVal
	}
	return
}
