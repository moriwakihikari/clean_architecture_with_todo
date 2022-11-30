package infrastructure

// 最外殻
// リクエストは外部から来るので最外殻のinfrastructureレイヤにルーティングを実装します。

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moriwakihikari/clean_architecture_with_todo.git/interfaces/controllers"
)

// SetUpRouting()でSqlhandlerを定義しControllerに渡してあげることで
// 各レイヤの依存関係を保つことができます。
func SetUpRouting() *http.ServeMux {
    mux := http.NewServeMux()

    sqlhandler := NewSqlhandler()
    userController := controllers.NewUserController(sqlhandler)
    todoController := controllers.NewTodoController(sqlhandler)

    mux.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            userController.Create(w, r)
        default:
            ResponseError(w, http.StatusNotFound, "")
        }
    })

    mux.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            userController.Show(w, r)
        default:
            ResponseError(w, http.StatusNotFound, "")
        }
    })
	mux.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Update(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/user/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Delete(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/todo/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Create(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/get", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Show(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Update(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Delete(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/search", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Search(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	return mux
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}