package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint      `json:"id" example:"1"`
	CreatedAt time.Time `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	DeletedAt time.Time `json:"deletedAt,omitempty" example:"2021-01-01T00:00:00Z"`
	Role      string    `json:"role" example:"Software Engineer"`
	Company   string    `json:"company" example:"Google"`
	Location  string    `json:"location" example:"Remote"`
	Remote    bool      `json:"remote" example:"true"`
	Link      string    `json:"link" example:"https://google.com"`
	Salary    int64     `json:"salary" example:"100000"`
}

type OpeningRequest struct {
	Role     string `json:"role" example:"Software Engineer"`
	Company  string `json:"company" example:"Google"`
	Location string `json:"location" example:"Remote"`
	Remote   *bool  `json:"remote" example:"true"`
	Link     string `json:"link" example:"https://google.com"`
	Salary   int64  `json:"salary" example:"100000"`
}

func (o *Opening) ToResponse() *OpeningResponse {
	return &OpeningResponse{
		ID:        o.ID,
		CreatedAt: o.CreatedAt,
		DeletedAt: o.DeletedAt.Time,
		UpdatedAt: o.UpdatedAt,

		Role:     o.Role,
		Company:  o.Company,
		Location: o.Location,
		Remote:   o.Remote,
		Link:     o.Link,
		Salary:   o.Salary,
	}
}

func (o *OpeningRequest) ToOpening() *Opening {
	return &Opening{
		Role:     o.Role,
		Company:  o.Company,
		Location: o.Location,
		Remote:   *o.Remote,
		Link:     o.Link,
		Salary:   o.Salary,
	}
}

// ValidateCreation validates the request body for creating a new opening
func (o *OpeningRequest) ValidateCreation() error {
	fieldRequired := func(s string) error {
		msg := fmt.Sprintf("missing required field: %s", s)
		return fmt.Errorf(msg)
	}

	switch {
	case o.Role == "" &&
		o.Company == "" &&
		o.Location == "" &&
		o.Link == "" &&
		o.Remote == nil &&
		o.Salary <= 00:
		return fmt.Errorf("request body is empty")

	case o.Role == "":
		return fieldRequired("Role (string)")
	case o.Link == "":
		return fieldRequired("Link (string)")
	case o.Remote == nil:
		return fieldRequired("Remote (bool)")
	case o.Company == "":
		return fieldRequired("Company (string)")
	case o.Location == "":
		return fmt.Errorf("location (string)")
	case o.Salary <= 00:
		return fieldRequired("Salary (int32)")

	default:
		return nil
	}
}

// ValidateUpdate validates the request body for updating an existing opening
func (o *OpeningRequest) ValidateUpdate() error {

	//truthy if any field is provided
	if o.Role != "" ||
		o.Company != "" ||
		o.Location != "" ||
		o.Link != "" ||
		o.Remote != nil ||
		o.Salary > 00 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
