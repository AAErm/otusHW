CREATE TABLE USERS (
    ID   SERIAL PRIMARY KEY,
    NAME VARCHAR(255)
);

CREATE TABLE EVENTS (
    ID                  SERIAL PRIMARY KEY,
    UserID              INTEGER REFERENCES USER (ID),
	Title               VARCHAR(255),
	DateAt              TIMESTAMP
	DateTo              TIMESTAMP
	Description         VARCHAR(255),
	NotificationTime TIMESTAMP
);