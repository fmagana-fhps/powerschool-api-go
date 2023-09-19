package models

type ExtensionData struct {
	TableExtension TableExtension `json:"_table_extension,omitempty"`
}

type TableExtension struct {
	RecordFound bool    `json:"recordFound,omitempty"`
	Field       []Field `json:"_field,omitempty"`
	Name        string  `json:"name,omitempty"`
}

type Field struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value any    `json:"value,omitempty"`
}
