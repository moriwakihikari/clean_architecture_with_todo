package infrastructure

// このレイヤでは外部のツールであるdbとの接続を行います。
// 次に実際にSQlとのデータのやり取りを行う処理をinterfaces/databaseに書きます。

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/moriwakihikari/clean_architecture_with_todo.git/interfaces/database"
)

type Sqlhandler struct {
	DB *sql.DB
}

func NewSqlhandler() *Sqlhandler {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	path := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	db, err := sql.Open("mysql", path); 
	if err != nil {
		return nil
		log.Fatal("Db open error:", err.Error())
	}
	err = db.Ping()
    if err != nil {
        fmt.Println("here")
        return nil
    }
	fmt.Println(db)

	return &Sqlhandler{db}
}

// func CreateTable(handler *Sqlhandler) (err error) {
// 	_, err = handler.Execute("CREATE TABLE IF NOT EXISTS users (id SERIAL NOT NULL, FirstName varchar(30) NOT NULL, LastName varchar(30) NOT NULL);")
// 	if err != nil {
// 		return
// 	}
// 	_, err = handler.Execute("CREATE TABLE IF NOT EXISTS todos (id SERIAL NOT NULL, title varchar(30) NOT NULL, note varchar(30), duedate DATE, userid INTEGER);")
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (handler *Sqlhandler) Execute(statement string, args ...interface{}) (database.Result, error) {
    res := SqlResult{}
    result, err := handler.DB.Exec(statement, args...)
    if err != nil {
        return res, err
    }
    res.Result = result
    return res, nil
}

func (handler *Sqlhandler) Query(statement string, args ...interface{}) (database.Row, error) {
    rows, err := handler.DB.Query(statement, args...)
    if err != nil {
        return new(SqlRow), err
    }
    row := new(SqlRow)
    row.Rows = rows
    return row, nil
}

type SqlResult struct {
    Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
    return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
    return r.Result.RowsAffected()
}

type SqlRow struct {
    Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
    return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
    return r.Rows.Next()
}

func (r SqlRow) Close() error {
    return r.Rows.Close()
}