package repository

// Task stuff //

var taskCreateTable string = `CREATE TABLE IF NOT EXISTS task (
	id INTEGER PRIMARY KEY,
	topic TEXT NOT NULL,
	notes TEXT,
	done BOOLEAN NOT NULL DEFAULT FALSE
);`

var taskReadAllRecordSchema string = `SELECT * from task;`
var taskCreateRecordSchema string = `INSERT INTO task (topic, notes) VALUES (?, ?);`
var taskUpdateTopicRecordSchema string = `UPDATE task SET (topic) = (?) WHERE id = (?);`
var taskUpdateNotesRecordSchema string = `UPDATE task SET (notes) = (?) WHERE id = (?);`
var taskSetCompleteRecordSchema string = `UPDATE task SET done = true WHERE id = (?);`
var taskSetIncompleteRecordSchema string = `UPDATE task SET done = false WHERE id = (?);`
var taskDeleteRecord string = `DELETE FROM task WHERE id = (?);`
