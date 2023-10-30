package ps

import (
	"fmt"

	m "github.com/fmagana-fhps/powerschool-api-go/models"
)

func (c *Client) StudentById(id string) (m.Response[m.StudentResponse], error) {
	url := fmt.Sprintf("/ws/v1/student/%s", id)
	resp, err := c.get(url, &m.StudentResponse{})
	if err != nil {
		return m.Response[m.StudentResponse]{}, err
	}

	model := convertBody(*resp, m.StudentResponse{})
	return model, err
}
