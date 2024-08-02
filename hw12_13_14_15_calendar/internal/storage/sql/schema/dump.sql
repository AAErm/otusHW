CREATE TABLE USERS (
    ID   VARCHAR(32) PRIMARY KEY,
    NAME VARCHAR(255)
);

CREATE TABLE EVENTS (
    ID                  VARCHAR(32) PRIMARY KEY,
    UserID              VARCHAR(32) REFERENCES USER (ID),
	Title               VARCHAR(255),
	DateAt              TIMESTAMP
	DateTo              TIMESTAMP
	Description         VARCHAR(255),
	NotificationAdvance TIMESTAMP
);