package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templは一つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil) // 本来はt.templ.Executeの戻り値をチェックすべき
}

func main() {
	// ルート
	// http.Handle(string, http.Handler)
	// http.Hander
	// type Handler interface {
	//      ServeHTTP(ResponseWriter, *Request)
	// }
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// webサーバーを開始する
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
