package query

import (
	"fmt"
	"strings"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

var fieldsNotification = []string{"EventID", "Status"}

func BuildAddNotificationQuery(eventID storage.ID) (string, []any) {
	qArgs := ""
	for i := 1; i < len(fieldsNotification); i++ {
		qArgs += fmt.Sprintf(", $%d", i)
	}

	return fmt.Sprintf(
			"INSERT INTO NOTIFICATION (%s) VALUES (%s)",
			strings.Join(fieldsEvent[1:], ", "),
			qArgs[1:]),
		[]any{
			eventID,
			"done",
		}
}
