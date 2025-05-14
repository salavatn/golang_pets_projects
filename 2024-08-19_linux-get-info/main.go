package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Настройки SSH-подключения
	config := &ssh.ClientConfig{
		User: "svc-admin", // Ваш логин
		Auth: []ssh.AuthMethod{
			ssh.Password("qaz123ZX"), // Ваш пароль
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Подключение к серверу
	client, _ := ssh.Dial("tcp", "192.168.100.155:22", config)
	defer client.Close()

	// Создание новой сессии
	session, _ := client.NewSession()
	defer session.Close()

	// Подключение stdin, stdout и stderr к текущему процессу
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	// Настройка терминала для интерактивного режима
	// PTY (pseudo-terminal) запрашивается для того, чтобы получить полноценный терминал
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // Включить/отключить эхо
		ssh.TTY_OP_ISPEED: 14400, // Скорость ввода (baudrate)
		ssh.TTY_OP_OSPEED: 14400, // Скорость вывода (baudrate)
	}
	session.RequestPty("xterm", 80, 40, modes)

	// Запуск оболочки (bash, sh и т.д.)
	session.Shell()

	// Обработка сигнала прерывания (Ctrl+C) для корректного завершения сессии
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		session.Signal(ssh.SIGINT)
	}()

	// Ожидание завершения сессии
	session.Wait()
}
