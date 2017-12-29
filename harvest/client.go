package harvest

import (
	"context"
	"time"
	"fmt"
	"net/http"
)

// ClientService handles communication with the client related
// methods of the Harvest API.
//
// Harvest API docs: https://help.getharvest.com/api-v2/clients-api/clients/clients/
type ClientService service

type Client struct {
	Id *int `json:"id,omitempty"` // Unique ID for the client.
	Name *string `json:"name,omitempty"` // A textual description of the client.
	IsActive *bool `json:"is_active,omitempty"` // Whether the client is active or archived.
	Address *string `json:"address,omitempty"` // The physical address for the client.
	Currency *string `json:"currency,omitempty"` // The currency code associated with this client.
	CreatedAt *time.Time `json:"created_at,omitempty"` // Date and time the client was created.
	UpdatedAt *time.Time `json:"updated_at,omitempty"` // Date and time the client was last updated.
}

type ClientList struct {
	Clients []*Client `json:"clients"`

	Pagination
}

func (c Client) String() string {
	return Stringify(c)
}

func (c ClientList) String() string {
	return Stringify(c)
}

type ClientListOptions struct {
	// Pass true to only return active projects and false to return inactive projects.
	IsActive	bool `url:"is_active,omitempty"`
	// Only return projects that have been updated since the given date and time.
	UpdatedSince	time.Time `url:"updated_since,omitempty"`

	ListOptions
}

func (s *ClientService) List(ctx context.Context, opt *ClientListOptions) (*ClientList, *http.Response, error) {
	u := "clients"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}


	clientList := new(ClientList)
	resp, err := s.client.Do(ctx, req, &clientList)
	if err != nil {
		return nil, resp, err
	}

	return clientList, resp, nil
}

func (s *ClientService) Get(ctx context.Context, clientId int) (*Client, *http.Response, error) {
	u := fmt.Sprintf("clients/%d", clientId)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	client := new(Client)
	resp, err := s.client.Do(ctx, req, &client)
	if err != nil {
		return nil, resp, err
	}

	return client, resp, nil
}