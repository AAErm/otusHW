package storage

import "time"

type ID int64

// Событие - основная сущность, содержит в себе поля:
// * ID - уникальный идентификатор события (можно воспользоваться UUID);
// * Заголовок - короткий текст;
// * Дата и время события;
// * Длительность события (или дата и время окончания);
// * Описание события - длинный текст, опционально;
// * ID пользователя, владельца события;
// * За сколько времени высылать уведомление, опционально.
type Event struct {
	ID               ID         `json:"id,omitempty"`
	Title            string     `json:"title"`
	DateAt           time.Time  `json:"date_at"`
	DateTo           time.Time  `json:"date_to"`
	Description      *string    `json:"description,omitempty"`
	UserID           ID         `json:"user_id"`
	NotificationTime *time.Time `json:"notification_time"`
}
