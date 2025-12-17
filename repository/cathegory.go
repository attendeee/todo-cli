package repository

type Cathegory struct {
	Id   int    `db:"id"`
	Body string `db:"body"`
}

func AddCathegory(cathegory string) error {
	_, err := Db().Exec(cathegoryCreateRecordSchema, cathegory)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCathegory(id int) error {
	_, err := Db().Exec(cathegoryDeleteRecord, id)
	if err != nil {
		return err
	}
	return nil

}

func ReadCathegory() (*[]Cathegory, error) {
	c := []Cathegory{}
	err := Db().Select(&c, cathegoryReadRecordsSchema)
	if err != nil {
		return nil, err
	}
	return &c, nil

}
func UpdateCathegory(id int, body string) error {
	_, err := Db().Exec(cathegoryUpdateRecordSchema, body, id)
	if err != nil {
		return err
	}

	return nil
}

func AssignCathegory(cathegory string, taskID int) error {
	_, err := Db().Exec(cathegoryAssignToTaskSchema, cathegory, taskID)
	if err != nil {
		return err
	}
	return nil

}
