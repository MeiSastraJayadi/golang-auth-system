package usecase

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
  godotenv.Load() 
  privK := os.Getenv("PRIVATE_KEY")
  privK = strings.ReplaceAll(privK, "\\n", "\n")
  token, err := GenerateJWT("meisastra", time.Minute*10, privK)
  if err != nil {
    t.Fail()
  }
  assert.NotNil(t, token)
}
