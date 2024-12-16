/*
    Схема БД для информационной системы
    отслеживания выполнения задач.
*/

DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;

-- пользователи системы
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- метки задач
CREATE TABLE labels (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- задачи
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    opened BIGINT NOT NULL DEFAULT extract(epoch from now()), -- время создания задачи
    closed BIGINT DEFAULT 0, -- время выполнения задачи
    author_id INTEGER REFERENCES users(id) DEFAULT 0, -- автор задачи
    assigned_id INTEGER REFERENCES users(id) DEFAULT 0, -- ответственный
    title TEXT, -- название задачи
    content TEXT -- задачи
);

-- связь многие - ко- многим между задачами и метками
CREATE TABLE tasks_labels (
    task_id INTEGER REFERENCES tasks(id),
    label_id INTEGER REFERENCES labels(id)
);

TRUNCATE TABLE tasks_labels, tasks, labels, users;

INSERT INTO users(name) VALUES
	('Ivanov'), ('Petrov'), ('Velon'),('Makin');

INSERT INTO labels(name) VALUES	('hard'), ('old'), ('easy'),('today');

INSERT INTO tasks(opened, closed, author_id,assigned_id,title, content) VALUES	
(extract(epoch from now()),extract(epoch from now()),4,1, 'task23', 'Задача 23'),
(extract(epoch from now()),extract(epoch from now()),1,2, 'task74', 'Задача 74'),
(extract(epoch from now()),extract(epoch from now()),3,1, 'task83', 'Задача 83'),
(extract(epoch from now()),extract(epoch from now()),4,2, 'task261', 'Задача 261');

INSERT INTO tasks_labels(task_id,label_id) VALUES	(2,3), (4,1), (1,4),(2,4);