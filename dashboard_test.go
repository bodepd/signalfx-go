package signalfx

import (
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDashboard(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboard", verifyRequest(t, "POST", http.StatusOK, nil, "dashboard/create_success.json"))

	result, err := client.CreateDashboard(&Dashboard{
		Name: "string",
	})
	assert.NoError(t, err, "Unexpected error creating dashboard")
	assert.Equal(t, "string", result.Name, "Name does not match")
}

func TestDeleteDashboard(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboard/string", verifyRequest(t, "DELETE", http.StatusOK, nil, ""))

	err := client.DeleteDashboard("string")
	assert.NoError(t, err, "Unexpected error getting dashboard")
}

func TestGetDashboard(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboard/string", verifyRequest(t, "GET", http.StatusOK, nil, "dashboard/get_success.json"))

	result, err := client.GetDashboard("string")
	assert.NoError(t, err, "Unexpected error getting dashboard")
	assert.Equal(t, result.Name, "string", "Name does not match")
}

func TestSearchDashboard(t *testing.T) {
	teardown := setup()
	defer teardown()

	limit := 10
	name := "foo"
	offset := 2
	tags := "bar"
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("name", name)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("tags", tags)

	mux.HandleFunc("/v2/dashboard", verifyRequest(t, "GET", http.StatusOK, params, "dashboard/get_success.json"))

	results, err := client.SearchDashboard(limit, name, offset, tags)
	assert.NoError(t, err, "Unexpected error search dashboard")
	assert.Equal(t, int64(0), results.Count, "Incorrect number of results")
}

func TestUpdateDashboard(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboard/string", verifyRequest(t, "PUT", http.StatusOK, nil, "dashboard/update_success.json"))

	result, err := client.UpdateDashboard("string", &Dashboard{
		Name: "string",
	})
	assert.NoError(t, err, "Unexpected error updating dashboard")
	assert.Equal(t, "string", result.Name, "Name does not match")
}