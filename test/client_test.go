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

	resp, err := c.Powerquery("net.fhps.transportation.transportation.get_records?pagesize=5", nil)
	trs := ps.CreatePowerqueryModels[m.TransportionRecord](resp.Body.Records)

	fmt.Println(trs, err)

	// fmt.Println(trans.Record[0], err)
}

func TestTransportation(t *testing.T) {
	c, err := client()
	if c == nil || err != nil {
		t.Fatalf("client is nil, error is %v", err)
	}

	r := m.TransportionRecord{
		Studentid:     "6510",
		Description:   "FH Data Test",
		Schoolid:      "1265",
		Startdate:     "2023-09-19",
		Enddate:       "2023-12-31",
		Departuretime: "34980",
		Arrivaltime:   "36900",
		FromTo:        "To",
		Type:          "R",
		Monday:        "0",
		Tuesday:       "0",
		Wednesday:     "1",
		Thursday:      "0",
		Friday:        "0",
		Busnumber:     "7",
		Drivername:    "Danis, Ben",
		Address:       "1725 Hall St SW",
	}

	re, err := c.CreateTransportRecord(r)
	if err != nil {
		fmt.Println(re)
		t.Fatalf("err is not nil, error is %v", err)
	}

	fmt.Println(re.Body)

	fmt.Println(c.TransportationRecordById(re.Body.Result[0].SuccessMessage.ID))

	fmt.Println(c.DeleteTransportationRecord(re.Body.Result[0].SuccessMessage.ID))
}
