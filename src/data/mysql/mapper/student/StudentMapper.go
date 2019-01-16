package student

import (
	"../../../../common/conn"
	"../../../../common/utils/DatetimeUtils"
	"../../../../common/utils/LoggerUtils"
	"encoding/json"
	"fmt"
)

type Student struct {
	Id        int    `json:"id"`
	Name      string `json:"user"`
	Value     int    `db:"score" json:"score"`
	ValueTime string `db:"score_time" json:"datetime"`
}

func GetById(id int) string {
	c := conn.MySQLConn()
	var s []Student
	sql := `select id, name, score, score_time from student where id = ?`
	err := c.Select(&s, sql, id)
	LoggerUtils.Error(err)
	j, _ := json.Marshal(s)
	return string(j)
}

func insertExample() string {
	c := conn.MySQLConn()
	sql := `insert into student(name, score, score_time) values(?, ?, ?)`
	r := c.MustExec(sql, "dd", 99, DatetimeUtils.GetDatetime())
	rowsAffected, _ := r.RowsAffected()
	lastInsertId, _ := r.LastInsertId()
	return fmt.Sprintf("rows affected : %d, last insert id : %d", rowsAffected, lastInsertId)
}
