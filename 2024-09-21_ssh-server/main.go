package main

import (
	"log"

	"github.com/gliderlabs/ssh"
)

// Структура для хранения учетных записей
type User struct {
	Username string
	Password string
}

// Массив данных с учетными записями
var users = []User{
	{"alice", "password"},
	{"bob", "password"},
	{"smith", "password"},
}


// 192.168.100.155 (svc-admin/qaz123ZX)
// Bash (create shell) 
// terminal  

// Функция для проверки логина и пароля
func checkCred(username, password string) bool {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		s.Write([]byte("Добро пожаловать, " + s.User() + "!\n"))
	})

	log.Println("Запуск SSH-сервера на порту 2222...")
	log.Println(ssh.ListenAndServe(":2222", nil, ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		// Проверка логина и пароля (Fatal)
		return checkCred(ctx.User(), password)
	})))

	// authentication := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
	// 	return checkCred(ctx.User(), password)
	// })

	// ssh.ListenAndServe(
	//     addr: ":222",
	//     handler: nil,

	// )

	// func(addr string, handler ssh.Handler, options ...ssh.Option) error
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler
	// to handle sessions on incoming connections. Handler is typically nil,
	// in which case the DefaultHandler is used.

}
