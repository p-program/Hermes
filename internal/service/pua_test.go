package service

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestPUA(t *testing.T) {
	godotenv.Load("../../.env")
	fmt.Println(os.Getenv("DEEPSEEK_API_KEY"))
	title := "高级产品经理"
	result := PUA(time.Now(), title)
	fmt.Println(result)
}
