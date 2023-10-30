package ps

import (
	"fmt"

	m "github.com/fmagana-fhps/powerschool-api-go/models"
)

// I want to adjust and change the output of this method. The type of record could
// be represented in a better way so that it is easier to be transformed into a struct
//
// Also there are possible url parameters that I need to implement.
func (c *Client) Powerquery(query string, data interface{}) (m.Response[m.PowerqueryResponse], error) {
	url := fmt.Sprintf("/ws/schema/query/%s", query)
	resp, err := c.post(url, data, &m.PowerqueryResponse{})
	if err != nil {
		return m.Response[m.PowerqueryResponse]{}, err
	}

	model := convertBody(*resp, m.PowerqueryResponse{})
	return model, err
}
