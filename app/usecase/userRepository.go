package usecase

//このようにRepositoryをもつInteractorを定義する。
//Repositoryはinterfaces/databaseレイヤに存在するので、
//この層でも先ほどと同じようにインターフェースを定義することで依存関係を解決する。
//そのためUsecaseレイヤでもRepositoryを定義する。

import "github.com/moriwakihikari/clean_architecture_with_todo.git/domain/model"

type UserRepository interface {
	Store(model.User) (int, error)
	Update(user model.User) (id int, err error)
	Delete(userID int) (err error)
	FindById(int) (model.User, error)
	FindAll() (model.Users, error)
}