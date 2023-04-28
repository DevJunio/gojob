package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func FieldRequired(s string) error {
	msg := fmt.Sprintf("missing required field: %s", s)
	return fmt.Errorf(msg)
}

func (req CreateOpeningRequest) Validate() error {
	switch {
	case req.Role    == ""  &&
		req.Company  == ""  &&
		req.Location == ""  &&
		req.Link     == ""  &&
		req.Remote   == nil &&
		req.Salary   <= 00  :
		return fmt.Errorf("request body is empty")

	case req.Role     == ""  : return FieldRequired("Role (string)")
	case req.Link     == ""  : return FieldRequired("Link (string)")
	case req.Remote   == nil : return FieldRequired("Remote (bool)")
	case req.Company  == ""  : return FieldRequired("Company (string)")
	case req.Location == ""  : return fmt.Errorf("location (string)")
	case req.Salary   <= 00  : return FieldRequired("Salary (int32)")

	default:
		return nil
	}
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	//truthy if any field is provided
	if r.Role  !=  ""  ||
	r.Company  !=  ""  ||
	r.Location !=  ""  ||
	r.Link     !=  ""  ||
	r.Remote   !=  nil ||
	r.Salary   >   00   {
		return nil
	}

	// Otherwise return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
