package api

import (
	"context"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
	"net/http"
)

// Client -
type Client interface {
	GetAdmin(ctx context.Context, username string) (*models.Admin, error)
	CreateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error)
	UpdateAdmin(ctx context.Context, username string, admin models.Admin) error
	DeleteAdmin(ctx context.Context, username string) error
	doRequest(ctx context.Context, req *http.Request) ([]byte, *http.Response, error)
}

// client -
type client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}