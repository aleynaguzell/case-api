package handler

import (
	"case-api/model/errormessage"
	"case-api/model/inmemory"
	"case-api/pkg/logger"
	"case-api/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type MemoryHandler struct {
}

var MemoryService services.MemoryService

func NewMemoryHandler(memoryService services.MemoryService) *MemoryHandler {
	MemoryService = memoryService
	return &MemoryHandler{}
}

// Set func set data to an in-memory database.
// @Description set data to an in-memory .
// @Summary set data to an in-memory .
// @Tags Memory
// @Accept json
// @Produce json
// @Success      201  {string}  "ok"
// @Failure      400  {string}  "error"
// @Failure      404  {string}  "notfound"
// @Failure      500  {string}  "error"
// @Router /in-memory/ [post]
// @Param inmemoryRequest body inmemory.Request true "InmemoryRequest"
func (m *MemoryHandler) Set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request inmemory.Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			_, err := fmt.Fprintf(w, "%+v", err.Error())
			if err != nil {
				http.Error(w, "Request can not decoded", http.StatusBadRequest)
				logger.Error("Request can not decoded", err)
				return
			}
		}

		err = MemoryService.Set(request.Key, request.Value)
		if err != nil {
			http.Error(w, "Can not set key/value pair", http.StatusInternalServerError)
			logger.Error("Can not set key/value pair", err)
			return
		}

		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(&request)
		if err != nil {
			http.Error(w, "Request can not encoded", http.StatusBadRequest)
			logger.Error("Request can not encoded", err)
			return
		}

		return
	} else {
		http.Error(w, "Method not found", http.StatusNotFound)
	}
}

// Get func fetch data from an in-memory database.
// @Description Get func fetch data from an in-memory database.
// @Summary gets value of key
// @Tags Memory
// @Accept json
// @Produce json
// @Param        key query string true "Key"
// @Success      200  {string}  string
// @Failure      400  {string}  string  "error"
// @Failure      404  {string}  string  "notfound"
// @Failure      500  {string}  string  "error"
// @Router       /in-memory [get]
func (m *MemoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		keyQuery := r.URL.Query()
		key := keyQuery.Get("key")
		value, err := MemoryService.Get(key)
		if err != nil {
			switch err.Error() {
			case errormessage.KeyEmpty:
				http.Error(w, errormessage.KeyEmpty, http.StatusInternalServerError)
				return
			case errormessage.KeyNotFound:
				http.Error(w, errormessage.KeyNotFound, http.StatusNotFound)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logger.CustomError(err)
				return
			}
		} else {
			out := inmemory.Request{Key: key, Value: value}
			err = json.NewEncoder(w).Encode(out)
			if err != nil {
				http.Error(w, "Can not encode value", http.StatusInternalServerError)
				logger.Error("Can not encode value", err)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	} else {
		http.Error(w, "Method not found", http.StatusNotFound)
	}
}
