package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Generate Log File
	log.Println("Sup Ninjas!")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("Next Sup Ninjas!")

	// Set output file, APPEND = untuk menambahkan dari file yang ada sebelumnya
	// CREATE = untuk membuat file, WRONLY = untuk write only
	// 0644 merupakan permission untuk membaca dan menulis file dan user lain hanya bisa membaca (referensi "https://www.maths.cam.ac.uk/computing/linux/unixinfo/perms#:~:text=644%20means%20you%20can%20read,users%20can%20only%20read%20it.")
	file, _ := os.OpenFile("file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	log.SetOutput(file)
	log.Println("Hello World")

	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(file, "INFO :", flags)
	warnLogger := log.New(file, "WARN :", flags)
	errorLogger := log.New(file, "ERROR :", flags)
	infoLogger.Println("This is info log")
	warnLogger.Println("This is warn log")
	errorLogger.Println("This is error log")

	file.Close()

	// Cara membaca file ".env", sebagai berikut
	// 1. Panggil di terminal "go mod init golang_env" untuk membuat go.mod
	// 2. Panggil "go get github.com/joho/godotenv" di terminal
	// 3. Buat file dengan nama ".env"
	// 4. Isi file ".env" dengan settingan yang ingin digunakan
	// 5. Panggil file ".env" dalam main function
	// 6. Jalankan aplikasi

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Gagal memuat file .env:", err)
		return
	}

	// Gunakan variabel lingkungan yang dimuat
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	fmt.Println("DB Host:", dbHost)
	fmt.Println("DB Port:", dbPort)
	fmt.Println("DB User:", dbUser)
	fmt.Println("DB Password:", dbPassword)
}
