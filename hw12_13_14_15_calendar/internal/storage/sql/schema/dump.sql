CREATE DATABASE CALENDAR;

\c calendar;

CREATE TABLE EVENTS (  
    ID SERIAL PRIMARY KEY,  
    UserID INTEGER,  
    Title VARCHAR(255),  
    DateAt TIMESTAMP,  
    DateTo TIMESTAMP,  
    Description VARCHAR(255),  
    NotificationTime TIMESTAMP  
);  

CREATE TABLE NOTIFICATION (
	EventID INTEGER,
	Status VARCHAR(32)
);

-- Вставка событий с годом  
INSERT INTO EVENTS (UserID, Title, DateAt, DateTo, Description, NotificationTime) VALUES  
(1, 'Annual Review', '2023-09-01 10:00:00', '2023-09-01 11:00:00', 'Annual performance review meeting', '2023-08-25 09:00:00'),  
(1, 'Team Building Retreat', '2023-10-10 09:00:00', '2023-10-12 17:00:00', 'Weekend retreat for team bonding', '2023-10-05 08:00:00');  

-- Вставка событий за текущую неделю  
INSERT INTO EVENTS (UserID, Title, DateAt, DateTo, Description, NotificationTime) VALUES  
(1, 'Weekly Standup', '2024-08-05 09:00:00', '2024-08-05 09:30:00', 'Weekly team standup meeting', '2024-08-04 08:00:00'),  
(2, 'Project Kickoff', '2024-08-06 14:00:00', '2024-08-06 15:00:00', 'Kickoff meeting for new project', '2024-08-05 12:00:00'),  
(1, 'Client Call', '2024-08-07 11:00:00', '2024-08-07 11:30:00', 'Call with important client', '2024-08-06 10:00:00'),  
(2, 'Team Sync', '2024-08-08 16:00:00', '2024-08-08 17:00:00', 'Sync meeting with team members', '2024-08-07 15:00:00'),  
(2, 'Weekly Review', '2024-08-09 10:00:00', '2024-08-09 11:00:00', 'Review of the weeks progress', '2024-08-08 09:00:00');  

-- Вставка событий за текущий месяц  
INSERT INTO EVENTS (UserID, Title, DateAt, DateTo, Description, NotificationTime) VALUES  
(1, 'Monthly Check-in', '2024-08-15 10:00:00', '2024-08-15 11:00:00', 'Monthly check-in with the manager', '2024-08-10 09:00:00'),  
(1, 'Budget Planning', '2024-08-20 14:00:00', '2024-08-20 15:00:00', 'Planning budget for next quarter', '2024-08-15 13:00:00'),  
(2, 'End of Month Report', '2024-08-25 09:00:00', '2024-08-25 10:00:00', 'Preparation of end of month report', '2024-08-20 08:00:00'); 