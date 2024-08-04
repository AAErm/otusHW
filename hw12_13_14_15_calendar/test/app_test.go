package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type EventCreateResponse int

type EventCreate struct {
	Title            string    `json:"title,omitempty"`
	DateAt           time.Time `json:"date_at,omitempty"`
	DateTo           time.Time `json:"date_to,omitempty"`
	Description      string    `json:"description,omitempty"`
	UserID           int       `json:"user_id,omitempty"`
	NotificationTime time.Time `json:"notification_time,omitempty"`
}

type EventsResponse []EventResponse

type EventResponse struct {
	ID               int       `json:"ID,omitempty"`
	Title            string    `json:"title,omitempty"`
	DateAt           time.Time `json:"date_at,omitempty"`
	DateTo           time.Time `json:"date_to,omitempty"`
	Description      string    `json:"description,omitempty"`
	UserID           int       `json:"user_id,omitempty"`
	NotificationTime time.Time `json:"notification_time,omitempty"`
}

func TestIntegration(t *testing.T) {
	timeFrom := time.Date(2024, 8, 04, 14, 4, 5, 0, time.UTC)
	unixTime := timeFrom.Unix()

	newEvent := EventCreate{
		Title:            "New Event",
		DateAt:           timeFrom,
		DateTo:           timeFrom.Add(24 * time.Hour),
		Description:      "Event description",
		UserID:           1,
		NotificationTime: timeFrom.Add(12 * time.Hour),
	}

	var eventId int

	t.Run("creating event", func(t *testing.T) {
		payload, err := json.Marshal(newEvent)
		require.NoError(t, err)

		statusCode, result, err := sendRequest(http.MethodPost, "add", payload)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, statusCode)

		var response EventCreateResponse

		err = json.Unmarshal(result, &response)
		require.NoError(t, err)
		require.NotEmpty(t, response)

		eventId = int(response)
	})

	t.Run("getting event by day", func(t *testing.T) {
		statusCode, result, err := sendRequest(http.MethodGet, "listByDay?time="+strconv.FormatInt(unixTime, 10), nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		var response EventsResponse

		err = json.Unmarshal(result, &response)
		require.NoError(t, err)
		require.Equal(t, 1, len(response))
		require.Equal(t, eventId, response[0].ID)
	})

	t.Run("getting event by weak", func(t *testing.T) {
		statusCode, result, err := sendRequest(http.MethodGet, "listByWeak?time="+strconv.FormatInt(unixTime, 10), nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		var response EventsResponse

		err = json.Unmarshal(result, &response)
		require.NoError(t, err)
		require.Equal(t, 1, len(response))
		require.Equal(t, eventId, response[0].ID)
	})

	t.Run("getting event by month", func(t *testing.T) {
		statusCode, result, err := sendRequest(http.MethodGet, "listByMonth?time="+strconv.FormatInt(unixTime, 10), nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		var response EventsResponse

		err = json.Unmarshal(result, &response)
		require.NoError(t, err)
		require.Equal(t, 1, len(response))
		require.Equal(t, eventId, response[0].ID)
	})

	t.Run("deleting event", func(t *testing.T) {
		statusCode, _, err := sendRequest(http.MethodDelete, fmt.Sprintf("delete/%d", eventId), nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		// Checking not exists event
		statusCode, result, err := sendRequest(http.MethodGet, "listByDay?time="+strconv.FormatInt(unixTime, 10), nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		var eventsResponse EventsResponse

		err = json.Unmarshal(result, &eventsResponse)
		require.NoError(t, err)
		require.Equal(t, 0, len(eventsResponse))
	})

}

func sendRequest(method string, endpoint string, payload []byte) (int, []byte, error) {
	host := "http://calendar:8080/"
	// host := "http://localhost:4444/"

	req, err := http.NewRequest(method, host+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, result, nil
}
