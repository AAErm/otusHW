package query

import (
	"fmt"
	"strings"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

var fieldsEvent = []string{
	"ID",
	"UserID",
	"Title",
	"DateAt",
	"DateTo",
	"Description",
	"NotificationAdvance",
}

func BuildAddEventQuery(event storage.Event) (string, []any) {
	qArgs := ""
	for i := 1; i <= len(fieldsEvent); i++ {
		qArgs += fmt.Sprintf(", $%d", i)
	}

	return fmt.Sprintf(
			"INSERT INTO EVENTS (%s) VALUES (%s) RETURNING ID",
			strings.Join(fieldsEvent, ", "),
			qArgs[1:]),
		[]any{
			event.ID,
			event.UserID,
			event.Title,
			event.DateAt.Truncate(time.Second),
			event.DateTo.Truncate(time.Second),
			event.Description,
			event.NotificationAdvance,
		}
}

func BuildEditEventQuery(event storage.Event) (string, []any) {
	partSQL := ""
	for k, v := range fieldsEvent {
		if v == "ID" {
			continue
		}
		partSQL += fmt.Sprintf(", %s = $%d", v, k)
	}
	return fmt.Sprintf("UPDATE EVENTS SET %s WHERE ID = %d",
			partSQL,
			len(fieldsEvent)),
		[]any{
			event.UserID,
			event.Title,
			event.DateAt.Truncate(time.Second),
			event.DateTo.Truncate(time.Second),
			event.Description,
			event.NotificationAdvance,
			event.ID,
		}
}

func BuildDeleteEventQuery(id storage.ID) (string, []any) {
	return "DELETE FROM EVENTS WHERE ID = $1", []any{id}
}

func BuildListEventByDay(t time.Time) (string, []any) {
	start := t.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)
	return fmt.Sprintf("SELECT %s WHERE DateAt >= $1 AND DateAt <= $2",
			strings.Join(fieldsEvent, ", ")),
		[]any{
			start,
			end,
		}
}

func BuildListEventByWeak(t time.Time) (string, []any) {
	start := t.AddDate(0, 0, -int(t.Weekday()-time.Monday))
	end := start.AddDate(0, 0, 7)
	return fmt.Sprintf("SELECT %s WHERE DateAt >= $1 AND DateAt <= $2",
			strings.Join(fieldsEvent, ", ")),
		[]any{
			start,
			end,
		}
}

func BuildListEventByMonth(t time.Time) (string, []any) {
	start := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	end := start.AddDate(0, 1, 0)
	return fmt.Sprintf("SELECT %s WHERE DateAt >= $1 AND DateAt <= $2",
			strings.Join(fieldsEvent, ", ")),
		[]any{
			start,
			end,
		}
}
