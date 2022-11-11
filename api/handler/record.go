package handler

import (
	"case-api/model/record"
	"case-api/services"
	"encoding/json"
	"net/http"
)

type RecordHandler struct {
}

var RecordService services.RecordService

func NewRecordHandler(recordService services.RecordService) *RecordHandler {
	RecordService = recordService
	return &RecordHandler{}
}

// Get func fetch data from records collection
// @Description func fetch data from records collection.
// @Summary func fetch data from records collection.
// @Tags Record
// @Accept json
// @Produce json
// @Success      200  {array}  record.Response
// @Failure      400  {string}  string  "error"
// @Failure      404  {string}  string  "notfound"
// @Failure      500  {string}  string  "error"
// @Router /records [post]
// @Param recordRequest body record.Request true "RecordRequest"
func (r *RecordHandler) Get(w http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodPost {
		var result record.Response
		var recordQuery record.Request

		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&recordQuery)

		records, err := RecordService.GetRecords(recordQuery)

		if err != nil {
			result.Code = http.StatusBadRequest
			result.Msg = err.Error()
			jData, _ := json.Marshal(result)

			w.WriteHeader(http.StatusBadRequest)
			w.Write(jData)
			return
		}
		if len(records) == 0 {
			result.Msg = "Not Found"
		} else {
			result.Msg = "Success"
		}

		result.Code = 0
		result.Records = records
		jData, _ := json.Marshal(result)

		w.WriteHeader(http.StatusOK)
		w.Write(jData)
		return
	} else {
		http.Error(w, "Method not found", http.StatusNotFound)
	}
}
