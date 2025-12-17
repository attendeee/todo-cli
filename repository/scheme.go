package repository

// Task stuff //

var taskCreateTable string = `CREATE TABLE IF NOT EXISTS task (
	id INTEGER PRIMARY KEY,
	topic TEXT NOT NULL,
	notes TEXT,
	done BOOLEAN NOT NULL DEFAULT FALSE,

	cathegoryID INTEGER NOT NULL,
	FOREIGN KEY (cathegoryID) REFERENCES cathegory(id) ON DELETE CASCADE
);`

var cathegoryCreateTable string = `CREATE TABLE IF NOT EXISTS cathegory (
	id INTEGER PRIMARY KEY,
	body TEXT NOT NULL UNIQUE
);`

var taskCreateRecordSchema string = `INSERT INTO task (cathegoryID, topic, notes) VALUES (?, ?, ?);`

// Task //
var taskReadAllRecordSchema string = `SELECT * from task;`

var taskReadAssignedRecordsSchema string = `SELECT * from task WHERE cathegoryID = (SELECT id FROM cathegory WHERE body = (?));`

var taskUpdateTopicRecordSchema string = `UPDATE task SET (topic) = (?) WHERE id = (?);`
var taskUpdateNotesRecordSchema string = `UPDATE task SET (notes) = (?) WHERE id = (?);`
var taskSetCompleteRecordSchema string = `UPDATE task SET done = true WHERE id = (?);`
var taskSetIncompleteRecordSchema string = `UPDATE task SET done = false WHERE id = (?);`

var taskDeleteRecord string = `DELETE FROM task WHERE id = (?);`

// Cathegory //
var cathegoryCreateRecordSchema string = `INSERT OR IGNORE INTO cathegory (body) VALUES (?);`
var cathegoryUpdateRecordSchema string = `UPDATE cathegory SET body = (?) WHERE id = (?);`
var cathegoryReadRecordsSchema string = `SELECT * from cathegory;`
var cathegoryDeleteRecord string = `DELETE FROM cathegory WHERE id = (?);`
var cathegoryAssignToTaskSchema string = `UPDATE task SET cathegoryID = (SELECT id FROM cathegory WHERE body = (?)) WHERE id = (?);`
