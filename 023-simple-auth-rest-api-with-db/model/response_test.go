package model_test

import (
	"bytes"
	"testing"

	nmodel "github.com/haidlir/x-golang-course/021-simple-rest-api/model"
)

func TestResponseFormat(t *testing.T) {
	t.Run("Data without Meta", func(t *testing.T) {
		resp := nmodel.NewResponseFormat()
		siswa1 := nmodel.Siswa{0, "A", 1}
		siswa2 := nmodel.Siswa{1, "B", 2}
		listOfData := []nmodel.Siswa{siswa1, siswa2}
		resp.SetData(listOfData)
		encodedData, err := resp.EncodeToJSON()
		if err != nil {
			t.Fatalf("There should no error while encoding to json but: %v", err)
		}
		encodedDataFixture := []byte(`{"data":[{"id":0,"nama":"A","kelas":1},{"id":1,"nama":"B","kelas":2}]}`)
		if !bytes.Equal(encodedData, encodedDataFixture) {
			t.Fatalf("Encoded data is not identical to the fixture.\nEncoded Data:\n%v\nFixture:\n%v", string(encodedData), string(encodedDataFixture))
		}
	})
	t.Run("Data with Meta", func(t *testing.T) {
		resp := nmodel.NewResponseFormat()
		siswa1 := nmodel.Siswa{0, "A", 1}
		siswa2 := nmodel.Siswa{1, "B", 2}
		listOfData := []nmodel.Siswa{siswa1, siswa2}
		resp.SetData(listOfData)
		resp.AddMeta("jumlah_siswa", 2)
		encodedData, err := resp.EncodeToJSON()
		if err != nil {
			t.Fatalf("There should no error while encoding to json but: %v", err)
		}
		encodedDataFixture := []byte(`{"data":[{"id":0,"nama":"A","kelas":1},{"id":1,"nama":"B","kelas":2}],"meta":{"jumlah_siswa":2}}`)
		if !bytes.Equal(encodedData, encodedDataFixture) {
			t.Fatalf("Encoded data is not identical to the fixture.\nEncoded Data:\n%v\nFixture:\n%v", string(encodedData), string(encodedDataFixture))
		}
	})
	t.Run("Error", func(t *testing.T) {
		resp := nmodel.NewResponseFormat()
		resp.AddError("error title 0", "error detail 0")
		resp.AddError("error title 1", "error detail 1")
		resp.AddError("error title 2", "error detail 2")
		encodedData, err := resp.EncodeToJSON()
		if err != nil {
			t.Fatalf("There should no error while encoding to json but: %v", err)
		}
		encodedDataFixture := []byte(`{"errors":[{"title":"error title 0","detail":"error detail 0"},{"title":"error title 1","detail":"error detail 1"},{"title":"error title 2","detail":"error detail 2"}]}`)
		if !bytes.Equal(encodedData, encodedDataFixture) {
			t.Fatalf("Encoded data is not identical to the fixture.\nEncoded Data:\n%v\nFixture:\n%v", string(encodedData), string(encodedDataFixture))
		}
	})
}
