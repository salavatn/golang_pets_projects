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

	// // Выполнение запроса для добавления данных
	// insertQuery := `
	// INSERT INTO car_statistic (odometer, fuel_level, average_speed)
	// VALUES (12345, 80, 65.5);
	// `
	// if err := db.Exec(insertQuery).Error; err != nil {
	// log.Fatalf("Ошибка при вставке данных: %v", err)
	// }

	// // Чтение данных из таблицы
	// var carStats []struct {
	// 	Odometer     int
	// 	FuelLevel    int
	// 	AverageSpeed float64
	// }
	// selectQuery := "SELECT odometer, fuel_level, average_speed FROM car_statistic;"
	// if err := db.Raw(selectQuery).Scan(&carStats).Error; err != nil {
	// 	log.Fatalf("Ошибка при чтении данных: %v", err)
	// }

	// // Вывод данных
	// for _, stat := range carStats {
	// 	fmt.Printf("Odometer: %d, Fuel Level: %d, Average Speed: %.2f\n", stat.Odometer, stat.FuelLevel, stat.AverageSpeed)
	// }

	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
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

func createTable(db *gorm.DB) {
	var tableName string = "TABLE"

	// Возможна SQL инъекция!!!
	createTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
		odometer INTEGER,
		fuel_level INTEGER,
		average_speed DECIMAL
	);
	`, tableName)

	if err := db.Exec(createTableQuery).Error; err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}
}
