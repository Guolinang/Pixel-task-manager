package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func ParseJSON(r *http.Request, paylaod any) error {
	if r.Body == nil {
		return fmt.Errorf("empty request body")
	}

	return json.NewDecoder(r.Body).Decode(paylaod)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	return WriteJSON(w, status, map[string]string{"error": err.Error()})
}

var Validate = validator.New()

type JsonDate time.Time

func (j JsonDate) MarshalJSON() ([]byte, error) {

	return []byte(fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))), nil
}

func (j *JsonDate) UnmarshalJSON(b []byte) error {

	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) String() string {
	return time.Time(j).Format("2006-01-02")
}
