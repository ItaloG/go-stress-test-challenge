package internal

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GenerateReport(duration time.Duration, totalRequests int, statusCodes map[int]int) {
	filePath := filepath.Join("files", "report.txt")
	err := os.MkdirAll("files", os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar a pasta 'files':", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de relatório:", err)
		return
	}
	defer file.Close()

	successfulRequests := statusCodes[http.StatusOK]

	formattedDuration := fmt.Sprintf("%.2fs", duration.Seconds())

	spendTimeMsg := fmt.Sprintf("Tempo total gasto na execução: %s\n", formattedDuration)
	totalRequestMsg := fmt.Sprintf("Quantidade total de requests realizados: %d\n", totalRequests)
	successfulRequestMsg := fmt.Sprintf("Quantidade de requests com status HTTP 200: %d\n", successfulRequests)
	codesTitle := fmt.Sprintln("Outros de status:")

	fmt.Print(spendTimeMsg)
	fmt.Print(totalRequestMsg)
	fmt.Print(successfulRequestMsg)
	fmt.Print(codesTitle)

	file.WriteString(spendTimeMsg)
	file.WriteString(totalRequestMsg)
	file.WriteString(successfulRequestMsg)
	file.WriteString(codesTitle)

	for code, count := range statusCodes {
		if code == http.StatusOK {
			continue
		}
		fmt.Printf("Status code %d - quantidade de requests: %d\n", code, count)
		fmt.Fprintf(file, "Status code %d - quantidade de requests: %d\n", code, count)
	}
}
