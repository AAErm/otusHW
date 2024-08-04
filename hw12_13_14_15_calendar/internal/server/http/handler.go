package internalhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/gorilla/mux"
)

func (s *server) hello(w http.ResponseWriter, r *http.Request) {
	_ = r
	fmt.Fprintln(w, "Hello, World!")
}

func (s *server) add(w http.ResponseWriter, r *http.Request) {
	event, err := getEventFromReq(r)
	s.logger.Errorf("event %+v", event)
	if err != nil {
		s.logger.Errorf("failed to get event from req %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	id, err := s.storage.Add(r.Context(), event)
	if err != nil {
		s.logger.Errorf("failed to add event %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprint(id)))
}

func (s *server) edit(w http.ResponseWriter, r *http.Request) {
	event, err := getEventFromReq(r)
	if err != nil {
		s.logger.Errorf("failed to get event from req %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	err = s.storage.Edit(r.Context(), event)
	if err != nil {
		s.logger.Errorf("failed to edit event %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func getEventFromReq(r *http.Request) (storage.Event, error) {
	defer r.Body.Close()
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		return storage.Event{}, fmt.Errorf("failed to read body %v", err)
	}
	var event storage.Event
	if err := json.Unmarshal(bb, &event); err != nil {
		return storage.Event{}, fmt.Errorf("failed to unmarshal event body %s with err %v", string(bb), err)
	}
	return event, nil
}

func (s *server) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		s.logger.Error("failed to get parameter 'id' from URL")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	iId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		s.logger.Error("id is not int")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = s.app.DeleteEvent(r.Context(), storage.ID(iId))
	if err != nil {
		s.logger.Errorf("failed to delete event with err %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (s *server) listByDay(w http.ResponseWriter, r *http.Request) {
	timeStart, err := getTimeToList(r)
	if err != nil {
		s.logger.Errorf("failed to get time with err %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	events, err := s.app.ListEventsByDay(r.Context(), timeStart)
	s.writeListResponce(w, events, err)
}

func getTimeToList(r *http.Request) (time.Time, error) {
	unixTime := r.URL.Query().Get("time")
	if unixTime == "" {
		return time.Time{}, fmt.Errorf("failed to get parameter 'time' from URL")
	}

	unixInt, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("unixTime is not int")
	}

	return time.Unix(unixInt, 0), nil
}

func (s *server) writeListResponce(w http.ResponseWriter, events []storage.Event, err error) {
	if err != nil {
		s.logger.Errorf("failed to get list events %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	bb, err := json.Marshal(events)
	if err != nil {
		s.logger.Errorf("failed to marshal events %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bb)
}

func (s *server) listByWeak(w http.ResponseWriter, r *http.Request) {
	timeStart, err := getTimeToList(r)
	if err != nil {
		s.logger.Errorf("failed to get time with err %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	events, err := s.app.ListEventsByWeek(r.Context(), timeStart)
	s.writeListResponce(w, events, err)
}

func (s *server) listByMonth(w http.ResponseWriter, r *http.Request) {
	timeStart, err := getTimeToList(r)
	if err != nil {
		s.logger.Errorf("failed to get time with err %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	events, err := s.app.ListEventsByMonth(r.Context(), timeStart)
	s.writeListResponce(w, events, err)
}
