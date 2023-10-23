package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReadJSON reads JSON from the request body into the data object.
// It returns an error if the body is empty or if the body has more than one JSON value.
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBodyBytes := int64(1048576) // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("ERROR: read json:", err)
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		log.Println("ERROR: read json:", err)
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

// WriteJSON writes the data object as JSON to the response body.
// It returns an error if the data object cannot be encoded as JSON.
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println("ERROR: to marshal json:", err)
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header().Set(k, v[0])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Println("ERROR: write json:", err)
		return err
	}

	return nil
}

// ErrorJSON writes the error message as JSON to the response body.
// It returns an error if the error message cannot be encoded as JSON.
func ErrorJSON(w http.ResponseWriter, status int, message string) error {
	return WriteJSON(w, status, map[string]string{"error": message})
}

// StringToPrimitiveObjectID converts a string to primitive.ObjectID
func StringToPrimitiveObjectID(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID
	}

	return oid
}
