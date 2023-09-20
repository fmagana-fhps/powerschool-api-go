package ps

import (
	"fmt"

	m "github.com/fmagana-fhps/powerschool-api-go/models"
)

func (c *Client) CreateTransportRecord(record m.TransportionRecord) ([]m.Result, error) {
	sr := m.SchemaRecord{Tables: m.Tables{TransportationRecord: record}}
	model, err := c.post("/ws/schema/table/transportation", sr, &m.SchemaResponse{})
	return model.Body.(*m.SchemaResponse).Result, err
}

func (c *Client) UpdateTransportRecord(id int, record m.TransportionRecord) ([]m.Result, error) {
	table := m.Tables{TransportationRecord: record}
	url := fmt.Sprintf("/ws/schema/table/transportation/%d", id)
	model, err := c.put(url, table, &m.SchemaResponse{})
	return model.Body.(*m.SchemaResponse).Result, err
}

func (c *Client) DeleteTransportationRecord(id int) (int, error) {
	url := fmt.Sprintf("/ws/schema/table/transportation/%d", id)
	model, err := c.delete(url, nil)
	return model.StatusCode, err
}

func (c *Client) TransportationRecordById(id int) (m.TransportionRecord, error) {
	url := fmt.Sprintf("/ws/schema/table/transportation/%d?projection=*", id)
	model, err := c.get(url, &m.SchemaRecord{})
	return model.Body.(*m.SchemaRecord).Tables.TransportationRecord, err
}
