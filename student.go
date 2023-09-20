package ps

import (
	"fmt"

	m "github.com/fmagana-fhps/powerschool-api-go/models"
)

func (c *Client) StudentById(id string) (m.Student, error) {
	url := fmt.Sprintf("/ws/v1/student/%s", id)
	model, err := c.get(url, &m.StudentResponse{})
	return model.Body.(*m.StudentResponse).Student, err
}
