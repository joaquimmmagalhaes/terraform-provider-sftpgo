package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
	"net/http"
	"strings"
)

func (c *client) GetUser(ctx context.Context, username string) (*models.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/users/%s", c.HostURL, username), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	order := models.User{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *client) CreateUser(ctx context.Context, admin models.User) (*models.User, error) {
	rb, err := json.Marshal(admin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/admins", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	order := models.User{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *client) UpdateUser(ctx context.Context, username string, admin models.User) error {
	rb, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/users/%s", c.HostURL, username), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(ctx, req)

	return err
}

func (c *client) DeleteUser(ctx context.Context, username string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/users/%s", c.HostURL, username), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(ctx, req)

	return err
}
