package models

type Admin struct {
	ID             int          `json:"id"`
	Status         int          `json:"status"`
	Username       string       `json:"username"`
	Description    string       `json:"description"`
	Password       string       `json:"password,omitempty"`
	Email          string       `json:"email"`
	Permissions    []string     `json:"permissions"`
	Filters        AdminFilters `json:"filters"`
	AdditionalInfo string       `json:"additional_info"`
}

type AdminFilters struct {
	AllowList []string `json:"allow_list"`
}
