package internal

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateReport(t *testing.T) {
	duration := 5 * time.Second
	totalRequests := 100
	statusCodes := map[int]int{
		http.StatusOK:                  50,
		http.StatusNotFound:            30,
		http.StatusInternalServerError: 20,
	}

	filePath := filepath.Join("files", "report.txt")

	GenerateReport(duration, totalRequests, statusCodes)

	expectedContent := []string{
		"Tempo total gasto na execução: 5.00s",
		"Quantidade total de requests realizados: 100",
		"Quantidade de requests com status HTTP 200: 50",
		"Outros de status:",
		"Status code 404 - quantidade de requests: 30",
		"Status code 500 - quantidade de requests: 20",
		"",
	}

	actualContent, err := os.ReadFile(filePath)
	assert.NoError(t, err, "Error reading report file")

	lines := strings.Split(string(actualContent), "\n")
	for index, expectedLine := range expectedContent {
		assert.Equal(t, lines[index], expectedLine)
	}

	os.Remove(filePath)
}
