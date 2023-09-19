package ps_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	ps "github.com/fmagana-fhps/powerschool-api-go"
	m "github.com/fmagana-fhps/powerschool-api-go/models"
	"github.com/joho/godotenv"
)

var options *ps.Options

func client() (*ps.Client, error) {
	godotenv.Load("../.env.local")
	expires := os.Getenv("EXPIRES")
	expire, err := strconv.Atoi(expires)

	if expires != "" {
		if err != nil {
			return nil, err
		}
	}

	options = &ps.Options{
		ClientId:     os.Getenv("ID"),
		ClientSecret: os.Getenv("SECRET"),
		Access: &m.AccessTokenResponse{
			AccessToken: os.Getenv("TOKEN"),
			ExpiresAt:   int64(expire),
		},
	}

	return ps.NewClient(os.Getenv("HOST"), options)
}

func TestClient(t *testing.T) {
	c, err := client()
	if c == nil {
		t.Fatalf("client is nil, error is %v", err)
	}

	fmt.Println(options.Access)

	s, _ := c.StudentById("79201")
	fmt.Println(s.ExtensionData)
}
