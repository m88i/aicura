package nexus

type User struct {
	UserId        string   `json:"userId"`
	FirstName     string   `json:"firstName,omitempty"`
	LastName      string   `json:"lastName,omitempty"`
	Email         string   `json:"emailAddress,omitempty"`
	Source        string   `json:"source,omitempty"`
	ReadOnly      bool     `json:"readOnly,omitempty"`
	Roles         []string `json:"roles"`
	ExternalRoles []string `json:"externalRoles"`
}

type UserService service

func (u *UserService) ListUsers() ([]User, error) {
	req, err := u.client.newRequest("GET", u.client.appendVersion("/security/users"), nil)
	if err != nil {
		return nil, err
	}
	var users []User
	_, err = u.client.do(req, &users)
	return users, err
}
