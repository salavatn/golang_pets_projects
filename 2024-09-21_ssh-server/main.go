package main

import (
	"log"

	"github.com/gliderlabs/ssh"
)

var users = map[string]string{
	"alice": "pa$$word",
	"bob":   "p@ssword",
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		s.Write([]byte("Добро пожаловать, " + s.User() + "!\n"))
	})

	log.Println("Запуск SSH-сервера на порту 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		expectedPassword, ok := users[ctx.User()]
		if !ok {
			return false // Пользователь не найден
		}
		return password == expectedPassword
	})))
}
