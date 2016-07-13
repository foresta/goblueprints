package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
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
	t.templ.Execute(w, r) // 本来はt.templ.Executeの戻り値をチェックすべき
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します

	gomniauth.SetSecurityKey("lkjlkjasdfasdf")
	gomniauth.WithProviders(
		google.New("626496761979-drq45a44rv4j2v5b2624vf61c2va5kni.apps.googleusercontent.com", "SfKHtHj1Qswt1-0r3ab_hLGK", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	//	r.tracer = trace.New(os.Stdout)

	//	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// チャットルームの開始
	go r.run()

	// webサーバーを開始する
	log.Println("Webサーバーを起動します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
