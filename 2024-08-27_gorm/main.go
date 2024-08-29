package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// user=postgres.xmysxvupqzihcoftnbxq password=[YOUR-PASSWORD] host=aws-0-eu-central-1.pooler.supabase.com port=6543 dbname=postgres

func main() {
	dsn := getSqlConfig(".env")
	db := initDB(dsn)
	checkDBConnection(db)

}

// Функция загружает переменные окружения из файла .env и формирует строку dsn (Data Source Name) для подключения к PostgreSQL
func getSqlConfig(envFileName string) string {
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Printf("Ошибка при загрузке %s файла:", envFileName)
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("TIME_ZONE"),
	)

	return dsn
}

// Функция принимает параметры для подключения к базе данных и инициирует соединение с PostgreSQL.
// Возвращает указатель на объект gorm.DB при успешном подключении или завершает работу программы в случае ошибки.
func initDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка в подключении к базе данных")
	}
	return db
}

// Функция для проверки статуса соединения и получения информации о базе данных
func checkDBConnection(db *gorm.DB) {
	// Проверяем подключение
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Ошибка при получении DB объекта: %v", err)
	}

	// Пингуем базу данных, чтобы проверить соединение
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("Соединение с базой данных не установлено: %v", err)
	}
	fmt.Println("Соединение с базой данных успешно установлено")

	// Получаем информацию о базе данных
	var version string
	db.Raw("SELECT version()").Scan(&version)
	fmt.Printf("Версия PostgreSQL: %s\n", version)

	var currentDB string
	db.Raw("SELECT current_database()").Scan(&currentDB)
	fmt.Printf("Текущая база данных: %s\n", currentDB)
}
