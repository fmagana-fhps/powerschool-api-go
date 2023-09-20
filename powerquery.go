package ps

import (
	"fmt"

	"github.com/fmagana-fhps/powerschool-api-go/models"
)

func (c *Client) Powerquery(query string, data interface{}) (models.PowerqueryResponse, error) {
	url := fmt.Sprintf("/ws/schema/query/%s", query)
	resp, err := c.post(url, data, &models.PowerqueryResponse{})
	return *resp.Body.(*models.PowerqueryResponse), err
}
