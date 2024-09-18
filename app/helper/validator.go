package helper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gookit/validate"
)

// ParseAndValidateForm adalah fungsi helper untuk melakukan parsing form dan validasi.
func ParseAndValidateForm(r *http.Request, form interface{}) (err error) {
	// Parsing body request ke dalam struct yang diberikan (form).
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// Mengembalikan error jika body request kosong atau parsing gagal.
		return errors.New("Empty payload")
	}

	// Membuat validator berdasarkan struct yang diparsing.
	v := validate.Struct(form)
	// Melakukan validasi berdasarkan aturan yang sudah didefinisikan di struct.
	if !v.Validate() {
		// Mengembalikan error jika validasi gagal, berisi pesan kesalahan validasi.
		return v.Errors
	}
	// Mengembalikan nil jika tidak ada error (validasi berhasil).
	return
}
