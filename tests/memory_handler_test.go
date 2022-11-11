package tests

import (
	"bytes"
	"case-api/api/handler"
	"case-api/services"
	"case-api/storage/cache"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//TODO: change with test mongo url
const testMongoUrl = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"

var memoryService = services.NewMemoryService(cache.New())
var memoryHandler = handler.NewMemoryHandler(*memoryService)

func TestInMemorySetController(t *testing.T) {

	expected := []byte(`{"key":"active-tabs","value":"getir"}`)
	postBody := []byte(`{"key":  "active-tabs","value": "getir"}`)

	req := httptest.NewRequest(http.MethodPost, "/in-memory/", bytes.NewBuffer(postBody))

	ctx := context.Background()
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(testMongoUrl))
	if err != nil {
		t.Fail()
	}

	err = mClient.Ping(ctx, nil)
	if err != nil {
		t.Fail()
	}

	defer mClient.Disconnect(context.TODO())

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(memoryHandler.Set)

	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	if !strings.Contains(string(body), string(expected)) {
		t.Error("expected "+string(expected)+" got", string(body))
	}
}

func TestInMemoryGetController(t *testing.T) {

	postBody := []byte(`{"key":  "active-tabs","value": "getir"}`)
	reqPost := httptest.NewRequest(http.MethodPost, "/in-memory/", bytes.NewBuffer(postBody))

	ctx := context.Background()
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(testMongoUrl))
	if err != nil {
		t.Fail()
	}

	err = mClient.Ping(ctx, nil)
	if err != nil {
		t.Fail()
	}

	defer mClient.Disconnect(context.TODO())

	wp := httptest.NewRecorder()
	setHandler := http.HandlerFunc(memoryHandler.Set)

	setHandler.ServeHTTP(wp, reqPost)
	expected := []byte(`{"key":"active-tabs","value":"getir"}`)
	req := httptest.NewRequest(http.MethodGet, "/in-memory?key=active-tabs", nil)

	w := httptest.NewRecorder()
	getHandler := http.HandlerFunc(memoryHandler.Get)

	getHandler.ServeHTTP(w, req)

	res := w.Result()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if http.StatusOK != res.StatusCode {
		t.Error("expected", http.StatusOK, "got status ", res.StatusCode)
	}
	if !strings.Contains(string(body), string(expected)) {
		t.Error("expected"+string(expected)+"got", string(body))
	}
}
