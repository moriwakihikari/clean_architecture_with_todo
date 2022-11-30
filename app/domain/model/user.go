package model

// Entitiesレイヤ
// User, Todoはビジネスルールのためのデータ構造なのでここで定義します。
// また、今回は追加していないユーザーの名前の文字数制限などはここで定義する必要があります。
// UserとTodoは次のようになります。
// 次はinfrastructureへ趣きdbとの接続へ

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Users []User