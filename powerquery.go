package ps

import "fmt"

func (c *Client) Powerquery(query string, data interface{}) (interface{}, error) {
	url := fmt.Sprintf("/ws/schema/query/%s", query)
	i, err := c.post(url, data, nil)
	return i, err
}
