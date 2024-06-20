package utils

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}

// GetBaseURL mengembalikan base URL dari host dan port yang diberikan
// func GetBaseURL(host string, port int) string {
// 	return fmt.Sprintf("https://%s:%d", host, port)
// }

// GetFileURL mengembalikan URL lengkap dari base URL dan path file
// func GetFileURL(baseURL, filePath string) string {
// 	return fmt.Sprintf("%s/%s", baseURL, filePath)
// }

// GetFileName mengembalikan nama file dari path lengkap
// func GetFileName(filePath string) string {
// 	// Menggunakan filepath.Base untuk mendapatkan nama file dari path lengkap
// 	fileName := filepath.Base(filePath)
// 	// Mencegah karakter backslash dan mengganti dengan slash untuk URL yang konsisten
// 	fileName = strings.Replace(fileName, "\\", "/", -1)
// 	return fileName
// }

// GetLocalFileURL mengembalikan URL lokal untuk file
// func GetLocalFileURL(basePath, fileName string) string {
// 	return fmt.Sprintf("file://%s/%s", basePath, fileName)
// }
