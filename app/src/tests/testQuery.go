package tests

import "boilerplate/app/database"

func GetTestQuery() []TestModel {
	result := []TestModel{}

	database.DBConn.Unscoped().
		Table("customer").
		Find(&result)

	return result
}
