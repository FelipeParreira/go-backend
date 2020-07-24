package app

import (
	"bytes"
	"encoding/json"
	"go-backend/docs"
	"go-backend/domain/servers"
	"go-backend/utils/errors"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"
)

func TestHealthCheckEndpoint(t *testing.T) {
	req := createRequest("GET", "/health_check", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, "OK", strings.TrimSpace(response.Body.String()))
}

func TestGetServersSummaryEndpoint(t *testing.T) {
	t.Run("should return an error message when the request body is invalid", func(t *testing.T) {
		requests := []*http.Request{
			createRequest("GET", "/servers/summary", servers.ServerInfoRequest{Hostname: ""}),
			createRequest("GET", "/servers/summary", ""),
		}
		expectedBody := errors.RestErr{
			Message: "Invalid JSON body. Body should have a non-empty property called `hostname` of type string.",
			Status:  http.StatusBadRequest,
		}

		for _, req := range requests {
			response := executeRequest(req)
			actualBody := errors.RestErr{}
			json.Unmarshal(response.Body.Bytes(), &actualBody)

			checkResponseCode(t, http.StatusBadRequest, response.Code)
			checkResponseBody(t, expectedBody, actualBody)
		}
	})

	t.Run("should return an error message when the hostname does not exist", func(t *testing.T) {
		req := createRequest("GET", "/servers/summary", servers.ServerInfoRequest{Hostname: "I don't exist!"})
		response := executeRequest(req)

		expectedBody := errors.RestErr{
			Message: "No server with hostname: `I don't exist!`",
			Status:  http.StatusNotFound,
		}
		actualBody := errors.RestErr{}
		json.Unmarshal(response.Body.Bytes(), &actualBody)

		checkResponseCode(t, http.StatusNotFound, response.Code)
		checkResponseBody(t, expectedBody, actualBody)
	})

	t.Run("should return summary info for an existing server", func(t *testing.T) {
		req := createRequest("GET", "/servers/summary", servers.ServerInfoRequest{Hostname: "server7"})
		response := executeRequest(req)

		expectedBody := servers.SummaryInfo{
			Hostname: "server7",
			CPULoad: servers.BaseSummaryInfo{
				Average: 0.49709810074639305,
				Mode:    0,
				Median:  0.4942173363737737,
				Unit:    "%",
			},
			MemoryUsage: servers.BaseSummaryInfo{
				Average: 5.027998799016362,
				Mode:    0,
				Median:  5.04783595043911,
				Unit:    "GB",
			},
			DiskUsage: servers.BaseSummaryInfo{
				Average: 26.743430454932618,
				Mode:    0,
				Median:  27.37304686426988,
				Unit:    "GB",
			},
		}
		actualBody := servers.SummaryInfo{}
		json.Unmarshal(response.Body.Bytes(), &actualBody)

		checkResponseCode(t, http.StatusOK, response.Code)
		checkResponseBody(t, expectedBody, actualBody)
	})
}

func checkResponseBody(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected response body %s. Got %s.\n", expected, actual)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected response code %d. Got %d.\n", expected, actual)
	}
}

func createRequest(method, url string, body interface{}) *http.Request {
	bodyByte, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(bodyByte)
	req, _ := http.NewRequest(method, url, bodyReader)

	return req
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	router := setUpRouter()
	req.URL.Path = path.Join(docs.SwaggerInfo.BasePath, req.URL.Path)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}
