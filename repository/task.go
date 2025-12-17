package repository

import (
	"database/sql"
)

type Task struct {
	Id    int    `db:"id" yaml:"id"`
	Topic string `db:"topic" yaml:"topic"`
	Notes string `db:"notes" yaml:"notes"`
	Done  bool   `db:"done" yaml:"done"`

	CathegoryID sql.NullInt64 `db:"cathegoryID" yaml:"-"`
}

func AddTask(c int16, topic string, notes string) (int64, error) {
	r, err := Db().Exec(taskCreateRecordSchema, c, topic, notes)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ReadTasks() (*[]Task, error) {
	t := []Task{}
	err := Db().Select(&t, taskReadAllRecordSchema)
	if err != nil {
		return nil, err
	}

	return &t, err
}

func ReadAssignedTasks(cathegoryName string) (*[]Task, error) {
	t := []Task{}
	var err error

	err = Db().Select(&t, taskReadAssignedRecordsSchema, cathegoryName)
	if err != nil {
		return nil, err
	}

	return &t, err
}

func DeleteTask(id int) error {
	_, err := Db().Exec(taskDeleteRecord, id)
	if err != nil {
		return err
	}

	return nil
}

func IncompleteTask(id int) error {
	_, err := Db().Exec(taskSetIncompleteRecordSchema, id)
	if err != nil {
		return err
	}

	return nil

}

func CompleteTask(id int) error {
	_, err := Db().Exec(taskSetCompleteRecordSchema, id)
	if err != nil {
		return err
	}

	return nil

}

func UpdateTask(id int, topic string, notes string) error {
	if topic != "" {
		_, err := Db().Exec(taskUpdateTopicRecordSchema, topic, id)
		if err != nil {
			return err
		}

	}

	if notes != "" {
		_, err := Db().Exec(taskUpdateNotesRecordSchema, notes, id)
		if err != nil {
			return err
		}
	}

	return nil
}
