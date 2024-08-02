package sqlstorage

import (
	"fmt"
	"strings"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

func BuildAddEventQuery(event storage.Event) (string, []any) {
	to_insert := []string{
		"ID",
		"UserID",
		"Title",
		"DateAt",
		"DateTo",
		"Description",
		"NotificationAdvance",
	}
	q_args := ""
	for i := 1; i <= len(to_insert); i++ {
		q_args += fmt.Sprintf(", $%d", i)
	}

	return fmt.Sprintf(
			"INSERT INTO EVENTS (%s) VALUES (%s) RETURNING ID",
			strings.Join(to_insert, ", "),
			q_args[1:]),
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
