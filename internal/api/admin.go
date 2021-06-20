package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
	"net/http"
	"strings"
)

func (c *client) GetAdmin(ctx context.Context, username string) (*models.Admin, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/admins/%s", c.HostURL, username), nil)
	if err != nil {
		return nil, err
	}

	body, res, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	order := models.Admin{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *client) CreateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	rb, err := json.Marshal(admin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/admins", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, res, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	order := models.Admin{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *client) UpdateAdmin(ctx context.Context, username string, admin models.Admin) error {
	rb, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/admins/%s", c.HostURL, username), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(ctx, req)

	return err
}

func (c *client) DeleteAdmin(ctx context.Context, username string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/admins/%s", c.HostURL, username), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(ctx, req)

	return err
}
