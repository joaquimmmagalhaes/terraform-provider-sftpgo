package models

type Admin struct {
	ID             int          `json:"id,omitempty"`
	Status         int          `json:"status,omitempty"`
	Username       string       `json:"username,omitempty"`
	Description    string       `json:"description,omitempty"`
	Password       string       `json:"password,omitempty"`
	Email          string       `json:"email,omitempty"`
	Permissions    []string     `json:"permissions,omitempty"`
	Filters        AdminFilters `json:"filters,omitempty"`
	AdditionalInfo string       `json:"additional_info,omitempty"`
}

type AdminFilters struct {
	AllowList []string `json:"allow_list,omitempty"`
}
