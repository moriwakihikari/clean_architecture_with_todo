package database

// このレイヤにはUserRepositoryとTodoRepositoryを宣言します。
// 例としてuserRepository.goを示します。SQLを叩いていることがわかります。

// UserRepositoryにinfrastructureレイヤで定義したSqlhandlerを埋め込んでいます。
// しかし、先ほど説明したように内側の層は外側の層の物を使用してはいけないので新しくSqlhandlerをinterfaceレイヤに定義します。

import (
	"log"

	"github.com/moriwakihikari/clean_architecture_with_todo.git/domain/model"
)

type UserRepository struct {
	Sqlhandler
}

func (repo *UserRepository) Store(u model.User) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		`INSERT INTO users (first_name, last_name) VALUES (?,?)`, u.FirstName, u.LastName,
	)
	if err != nil {
		log.Fatal("Db open error:", err.Error())
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *UserRepository) Update(user model.User) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"UPDATE users SET first_name=?, last_name=? WHERE id=?", user.FirstName, user.LastName, user.ID,
	)

	if err != nil {
		log.Fatal("Db open error:", err.Error())
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *UserRepository) Delete(userID int) (err error) {
	_, err = repo.Sqlhandler.Query("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindById(userID int) (user model.User, err error) {
	row, err := repo.Sqlhandler.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", userID)
	if err != nil {
		log.Fatal("Db open error:", err.Error())
		return
	}
	row.Next()
	if err = row.Scan(&(user.ID), &(user.FirstName), &(user.LastName)); err != nil {
		log.Fatal("Db open error:", err.Error())
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	rows, err := repo.Sqlhandler.Query(`SELECT id, first_name, last_name FROM users`)
	if err != nil {
		log.Fatal("Db open error:", err.Error())
		return
	}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&(user.ID), &(user.FirstName), &(user.LastName)); err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}