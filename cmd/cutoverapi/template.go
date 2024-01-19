package cutoverapi

import "time"

type Template struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Archived          bool      `json:"archived"`
			AutoStart         bool      `json:"auto_start"`
			CreatedAt         time.Time `json:"created_at"`
			CustomFieldValues []struct {
				Name  string   `json:"name"`
				Value []string `json:"value"`
			} `json:"custom_field_values"`
			Description     string    `json:"description"`
			EndActual       any       `json:"end_actual"`
			EndForecast     any       `json:"end_forecast"`
			EndPlanned      time.Time `json:"end_planned"`
			EndScheduled    any       `json:"end_scheduled"`
			ExternalID      any       `json:"external_id"`
			IsTemplate      bool      `json:"is_template"`
			Name            string    `json:"name"`
			RunComms        any       `json:"run_comms"`
			RunType         any       `json:"run_type"`
			Stage           string    `json:"stage"`
			StartActual     any       `json:"start_actual"`
			StartPlanned    time.Time `json:"start_planned"`
			StartScheduled  any       `json:"start_scheduled"`
			Status          string    `json:"status"`
			StatusAuthor    any       `json:"status_author"`
			StatusMessage   any       `json:"status_message"`
			StatusUpdatedAt any       `json:"status_updated_at"`
			TemplateStatus  string    `json:"template_status"`
			TemplateType    string    `json:"template_type"`
			Timezone        any       `json:"timezone"`
			TimingMode      string    `json:"timing_mode"`
			UpdatedAt       time.Time `json:"updated_at"`
		} `json:"attributes"`
		Relationships struct {
			Author struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"author"`
			CurrentVersion struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"current_version"`
			Folder struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"folder"`
			RunbookType struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"runbook_type"`
			Workspace struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"workspace"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Meta struct {
			ApprovalStatus      string `json:"approval_status"`
			TemplateVersion     int    `json:"template_version"`
			TasksCount          int    `json:"tasks_count"`
			CompletedTasksCount int    `json:"completed_tasks_count"`
		} `json:"meta"`
	} `json:"data"`
	Included []any `json:"included"`
	Meta     struct {
		Page struct {
			Number int `json:"number"`
			Total  int `json:"total"`
		} `json:"page"`
	} `json:"meta"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}
