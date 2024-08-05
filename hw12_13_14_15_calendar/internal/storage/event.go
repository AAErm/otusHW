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
	ID                  ID
	Title               string
	DateAt              time.Time
	DateTo              time.Time
	Description         *string
	UserID              ID
	NotificationAdvance *time.Time
}
