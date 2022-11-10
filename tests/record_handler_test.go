package tests

import (
	"bytes"
	"case-api/api/handler"
	"case-api/services"
	"case-api/storage/repository"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestRecordController(t *testing.T) {

	expected := []byte(`{"code":0,"msg":"Success","records":[{"key":"KrZIErky","createdAt":"2016-08-15T01:12:05.989Z","totalCount":2993},{"key":"KrZIErky","createdAt":"2016-08-15T01:12:05.989Z","totalCount":2992},{"key":"bxoQiSKL","createdAt":"2016-01-29T01:59:53.494Z","totalCount":2991}]}`)

	postBody := []byte(`{
		"startDate": "2016-01-26",
		"endDate": "2018-02-02",
		"minCount": 2990,
		"maxCount": 3000
	}`)

	req := httptest.NewRequest(http.MethodPost, "/records", bytes.NewBuffer(postBody))

	mClient, err := GetMongoClient(t)
	defer mClient.Disconnect(context.TODO())


	recordRepository := repository.NewRecordsRepository(mClient)
	recordService := services.NewRecordService(*recordRepository)
	recordHandler := handler.NewRecordHandler(*recordService)

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(recordHandler.Get)

	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err oluştu ", err)
		t.Fatal(err)
	}

	if http.StatusOK != res.StatusCode {
		t.Error("expected", http.StatusOK, "got status", res.StatusCode)
	}
	if !strings.Contains(string(body), string(expected)) {
		t.Error("expected"+string(expected)+"got", string(body))
	}
}
