package datatypes

import (
	"time"
)

type COTemplate struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Archived        bool      `json:"archived"`
			AutoStart       bool      `json:"auto_start"`
			CreatedAt       time.Time `json:"created_at"`
			CustomFields    any       `json:"custom_field_values"`
			Description     string    `json:"description"`
			ExternalID      any       `json:"external_id"`
			IsTemplate      bool      `json:"is_template"`
			Name            string    `json:"name"`
			Status          string    `json:"status"`
			StatusAuthor    any       `json:"status_author"`
			StatusMessage   any       `json:"status_message"`
			StatusUpdatedAt any       `json:"status_updated_at"`
			TemplateStatus  string    `json:"template_status"`
			TemplateType    string    `json:"template_type"`
			UpdatedAt       time.Time `json:"updated_at"`
		}
	}
}
