package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	nhandler "github.com/haidlir/x-golang-course/021-simple-rest-api/handler"
	nmodel "github.com/haidlir/x-golang-course/021-simple-rest-api/model"

	"github.com/gorilla/mux"
)

// MockedDB was made to mock DB
type MockedDB struct {
	ErrMap      map[string]error
	MockedSiswa []nmodel.Siswa
}

func newMockedDB() *MockedDB {
	siswaBaru := nmodel.Siswa{
		Nama:  "Ardi",
		Kelas: 6,
	}
	db := new(MockedDB)
	db.MockedSiswa = append(db.MockedSiswa, siswaBaru)
	db.ErrMap = map[string]error{}
	return db
}

func (db *MockedDB) GetAllSiswa() []nmodel.Siswa {
	funcName := "GetAllSiswa"
	if db.ErrMap[funcName] != nil {
		return nil
	}
	return db.MockedSiswa
}
func (db *MockedDB) GetDetailSiswa(id int) *nmodel.Siswa {
	funcName := "GetDetailSiswa"
	if db.ErrMap[funcName] != nil {
		return nil
	}
	return &db.MockedSiswa[0]
}
func (db *MockedDB) AddSiswa(nmodel.Siswa) (*nmodel.Siswa, error) {
	funcName := "AddSiswa"
	if db.ErrMap[funcName] != nil {
		return nil, db.ErrMap[funcName]
	}
	return &db.MockedSiswa[0], nil
}
func (db *MockedDB) UpdateSiswa(id int, data nmodel.Siswa) (*nmodel.Siswa, error) {
	funcName := "UpdateSiswa"
	if db.ErrMap[funcName] != nil {
		return nil, db.ErrMap[funcName]
	}
	db.MockedSiswa[0].Nama = data.Nama
	db.MockedSiswa[0].Kelas = data.Kelas
	return &db.MockedSiswa[0], nil
}
func (db *MockedDB) DeleteSiswa(id int) error {
	funcName := "DeleteSiswa"
	if db.ErrMap[funcName] != nil {
		return db.ErrMap[funcName]
	}
	return nil
}

func setupHandlerWithMockedDB(db *MockedDB, t *testing.T) *nhandler.Handler {
	siswa := db.GetAllSiswa()
	if len(siswa) != 1 {
		t.Fatalf("The number of siswa should 1 instead of %v", len(siswa))
	}
	handler := nhandler.NewHandler(db)
	if handler == nil {
		t.Fatal("handler should not nil")
	}
	return handler
}

func TestNewHandler(t *testing.T) {
	db := newMockedDB()
	if db == nil {
		t.Fatal("db should not nil")
	}
	setupHandlerWithMockedDB(db, t)
}

func TestSiswaHandler(t *testing.T) {
	db := newMockedDB()
	if db == nil {
		t.Fatal("db should not nil")
	}
	handler := setupHandlerWithMockedDB(db, t)
	testURL := "/api/test"
	testSpecificURL := fmt.Sprintf("%v/{id:[0-9]+}", testURL)
	// Setup Routing
	r := mux.NewRouter()
	r.HandleFunc(testURL, handler.AddSiswa).Methods(http.MethodPost)
	r.HandleFunc(testURL, handler.GetAllSiswa).Methods(http.MethodGet)
	r.HandleFunc(testSpecificURL, handler.GetDetailSiswa).Methods(http.MethodGet)
	r.HandleFunc(testSpecificURL, handler.UpdateSiswa).Methods(http.MethodPut)
	r.HandleFunc(testSpecificURL, handler.DeleteSiswa).Methods(http.MethodDelete)
	// Create httptest Server
	httpServer := httptest.NewServer(r)
	defer httpServer.Close()
	serverURL, _ := url.Parse(httpServer.URL)
	// Hit API Endpoint
	targetPath := fmt.Sprintf("%v%v", serverURL, testURL)
	specificTargetPath := fmt.Sprintf("%v%v/1", serverURL, testURL)
	// Insert Scenario, OK
	t.Run("Insert Siswa OK", func(t *testing.T) {
		// Hit API Endpoint
		var jsonRequest = []byte(`{"nama":"Ardi", "kelas":7}`)
		req, _ := http.NewRequest(http.MethodPost, targetPath, bytes.NewBuffer(jsonRequest))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("status code %v instead of 201", resp.StatusCode)
		}
		// Get the Body
		encodedBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response body: %v", err)
		}
		var siswa nmodel.Siswa
		_, _, err = extractContent(encodedBody, &siswa)
		if siswa.Nama != "Ardi" {
			t.Fatalf("nama in siswa should be ardi instead of  %v", siswa.Nama)
		}
		resp.Body.Close()
	})
	// Insert Scenario, NOK
	t.Run("Insert Siswa NOK", func(t *testing.T) {
		db.ErrMap["AddSiswa"] = fmt.Errorf("Intentionally Error")
		// Hit API Endpoint
		var jsonRequest = []byte(`{"nama":"Ardi", "kelas":7}`)
		req, _ := http.NewRequest(http.MethodPost, targetPath, bytes.NewBuffer(jsonRequest))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("status code %v instead of 400", resp.StatusCode)
		}
		resp.Body.Close()
	})
	// Get All Scenario, OK
	t.Run("Get All Siswa OK", func(t *testing.T) {
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodGet, targetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("status code %v instead of 200", resp.StatusCode)
		}
		// Get the Body
		encodedBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response body: %v", err)
		}
		var siswa []nmodel.Siswa
		_, _, err = extractContent(encodedBody, &siswa)
		if len(siswa) != 1 {
			t.Fatalf("The amount of Siswa should be 1 instead of %v", len(siswa))
		}
		if siswa[0].Nama != "Ardi" {
			t.Fatalf("nama in siswa[0] should be ardi instead of  %v", siswa[0].Nama)
		}
		resp.Body.Close()
	})
	// Get All Scenario, NOK
	t.Run("Get All Siswa NOK", func(t *testing.T) {
		db.ErrMap["GetAllSiswa"] = fmt.Errorf("Intentionally Error")
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodGet, targetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("status code %v instead of 500", resp.StatusCode)
		}
		resp.Body.Close()
	})
	// Get Specific Siswa Scenario, OK
	t.Run("Get Specific Siswa Siswa OK", func(t *testing.T) {
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodGet, specificTargetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("status code %v instead of 200", resp.StatusCode)
		}
		// Get the Body
		encodedBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response body: %v", err)
		}
		var siswa nmodel.Siswa
		_, _, err = extractContent(encodedBody, &siswa)
		if err != nil {
			t.Fatalf("error in decoding data: %v", err)
		}
		if siswa.Nama != "Ardi" {
			t.Fatalf("nama in siswa should be ardi instead of  %v", siswa.Nama)
		}
		resp.Body.Close()
	})
	// Get Specific Siswa Scenario, NOK
	t.Run("Get Specific Siswa Siswa NOK", func(t *testing.T) {
		db.ErrMap["GetDetailSiswa"] = fmt.Errorf("Intentionally Error")
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodGet, specificTargetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("status code %v instead of 500", resp.StatusCode)
		}
		resp.Body.Close()
	})
	// Delete Specific Siswa Scenario, OK
	t.Run("Delete Specific Siswa Siswa OK", func(t *testing.T) {
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodDelete, specificTargetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("status code %v instead of 200", resp.StatusCode)
		}
		// Get the Body
		encodedBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response body: %v", err)
		}
		var message map[string]string
		_, _, err = extractContent(encodedBody, &message)
		if err != nil {
			t.Fatalf("error in decoding data: %v", err)
		}
		if message["message"] != "siswa id 1 deleted successfully" {
			t.Fatalf("failed to delete")
		}
		resp.Body.Close()
	})
	// Delete Specific Siswa Scenario, NOK
	t.Run("Delete Specific Siswa Siswa NOK", func(t *testing.T) {
		db.ErrMap["DeleteSiswa"] = fmt.Errorf("Intentionally Error")
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodDelete, specificTargetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("status code %v instead of 500", resp.StatusCode)
		}
		resp.Body.Close()
	})
	// Update Specific Siswa Scenario, OK
	t.Run("Update Specific Siswa Siswa OK", func(t *testing.T) {
		var jsonRequest = []byte(`{"nama":"Azzam", "kelas":10}`)
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodPut, specificTargetPath, bytes.NewBuffer(jsonRequest))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("status code %v instead of 200", resp.StatusCode)
		}
		// Get the Body
		encodedBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response body: %v", err)
		}
		var siswa nmodel.Siswa
		_, _, err = extractContent(encodedBody, &siswa)
		if err != nil {
			t.Fatalf("error in decoding data: %v", err)
		}
		if siswa.Nama != "Azzam" {
			t.Fatalf("nama in siswa should be Azzam instead of  %v", siswa.Nama)
		}
		if siswa.Kelas != 10 {
			t.Fatalf("kelas in siswa should be 10 instead of  %v", siswa.Kelas)
		}
		resp.Body.Close()
	})
	// Update Specific Siswa Scenario, NOK
	t.Run("Update Specific Siswa Siswa NOK", func(t *testing.T) {
		db.ErrMap["UpdateSiswa"] = fmt.Errorf("Intentionally Error")
		var jsonRequest = []byte(`{"nama":"Azzam", "kelas":10}`)
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodPut, specificTargetPath, bytes.NewBuffer(jsonRequest))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("status code %v instead of 500", resp.StatusCode)
		}
		resp.Body.Close()
	})
	// Update Specific Siswa Scenario, NOK 2
	t.Run("Update Specific Siswa Siswa NOK 2", func(t *testing.T) {
		db.ErrMap["UpdateSiswa"] = fmt.Errorf("Intentionally Error")
		var jsonRequest = []byte(`{"nama":"Azzam", kelas":10}`)
		// Hit API Endpoint
		req, _ := http.NewRequest(http.MethodPut, specificTargetPath, bytes.NewBuffer(jsonRequest))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to hit API: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("status code %v instead of 400", resp.StatusCode)
		}
		resp.Body.Close()
	})
}

func extractContent(encodedBody []byte, content interface{}) (metaResp *nmodel.MetaResponse, errResp *nmodel.Error, err error) {
	// Unmarshal JSON Resp
	var body nmodel.ResponseFormat
	err = json.Unmarshal(encodedBody, &body)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to unmarshal json response: %v", err)
	}
	if body.Errors != nil {
		encodedContent, _ := json.Marshal(body.Errors)
		err = json.Unmarshal(encodedContent, &errResp)
		if err != nil {
			return nil, nil, fmt.Errorf("Unable to unmarshal erros: %v", err)
		}
		return nil, errResp, fmt.Errorf("error response")
	}
	encodedContent, _ := json.Marshal(body.Data)
	err = json.Unmarshal(encodedContent, &content)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to unmarshal data: %v", err)
	}
	if body.Meta != nil {
		encodedContent, _ := json.Marshal(body.Meta)
		err = json.Unmarshal(encodedContent, &metaResp)
		if err != nil {
			return nil, nil, fmt.Errorf("Unable to unmarshal erros: %v", err)
		}
	} else {
		metaResp = nil
	}
	return metaResp, nil, nil
}
