package ps

import (
	"fmt"

	m "github.com/fmagana-fhps/powerschool-api-go/models"
)

func (c *Client) CreateTransportRecord(record m.TransportionRecord) (m.Response[m.SchemaResponse], error) {
	sr := m.SchemaRecord{Tables: m.Tables{TransportationRecord: record}}
	resp, err := c.post("/ws/schema/table/transportation", sr, &m.SchemaResponse{})
	if err != nil {
		return m.Response[m.SchemaResponse]{}, err
	}

	model := convertBody(*resp, m.SchemaResponse{})
	return model, err
}

func (c *Client) UpdateTransportRecord(id int, record m.TransportionRecord) (m.Response[m.SchemaResponse], error) {
	table := m.SchemaRecord{Tables: m.Tables{TransportationRecord: record}}
	url := fmt.Sprintf("/ws/schema/table/transportation/%d", id)
	resp, err := c.put(url, table, &m.SchemaResponse{})
	if err != nil {
		return m.Response[m.SchemaResponse]{}, err
	}

	model := convertBody(*resp, m.SchemaResponse{})
	return model, err
}

func (c *Client) DeleteTransportationRecord(id int) (int, error) {
	url := fmt.Sprintf("/ws/schema/table/transportation/%d", id)
	resp, err := c.delete(url)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, err
}

func (c *Client) TransportationRecordById(id int) (m.Response[m.SchemaRecord], error) {
	url := fmt.Sprintf("/ws/schema/table/transportation/%d?projection=*", id)
	resp, err := c.get(url, &m.SchemaRecord{})
	if err != nil {
		return m.Response[m.SchemaRecord]{}, err
	}

	model := convertBody(*resp, m.SchemaRecord{})
	return model, err
}
