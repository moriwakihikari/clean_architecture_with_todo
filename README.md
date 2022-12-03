## 処理の流れ

### ユーザー取得の場合
- main.go
    - infrastructure/router.goのSetUpRoutingを呼び出す
- infrastructure/router.go
    - infrastructure/sqlhandler.goのNewSqlhandlerを呼び出す
        - db接続を行う
- infrastructure/router.go
    - dbの接続情報を引数に設定しcontrollerに渡す
    - この時点で色々状態を渡している
        - Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				Sqlhandler: sqlHandler,
			},
		},
        - usecaseを呼び出すとき
            - usecase.UserInteractor
            - UserRepositoryと流れる
            - domain/modelから構造体を取得
        - そしてdatabase.UserRepositoryを呼び出す
            - sql文が入っている

- rooterからcontrollerを呼び出す時に色々呼び出していると推測
    - このディレクトリだとNewUserControllerメソッドが鍵な気がする      
