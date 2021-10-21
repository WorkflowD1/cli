package model

type Document struct {
	StatusID           int64                  `json:"status_id"`
	ProductID          int64                  `json:"product_id"`
	ModalityIdentifier string                 `json:"modality_identifier"`
	CPF                string                 `json:"cpf"`
	Name               string                 `json:"name"`
	Source             string                 `json:"source" default:"cli"`
	Phone              string                 `json:"phone"`
	Email              string                 `json:"email"`
	FilledColumns      map[string]interface{} `json:"filled_columns"`
}

type DocumentResponse struct {
	UUID         string `json:"uuid"`
	SlaStart     string `json:"sla_start"`
	SlaEnd       string `json:"sla_end"`
	SlaEndStatus string `json:"sla_end_status"`
	DateTime     string `json:"date_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ID           int    `json:"id"`
	Protocol     string `json:"protocol"`
}
