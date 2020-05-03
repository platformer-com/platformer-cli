package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

const (
	orgListURL = "https://auth-module.dev.x.platformer.com/api/v1/organization/list"
)

// Organization models a Platformer Organization
type Organization struct {
	ID          string `json:"organization_id"`
	Name        string `json:"name"`
	UserName    string `json:"user_name"`
	UID         string `json:"uid"`
	UserID      string `json:"id"`
	UserEmail   string `json:"user_email"`
	Pending     bool   `json:"pending"`
	Owner       bool   `json:"owner"`
	CreatedDate struct {
		Seconds     int64 `json:"_seconds"`
		NanoSeconds int64 `json:"nano_seconds"`
	} `json:"created_date"`
}

// OrganizationList is a map of organizations by Name
type OrganizationList map[string]Organization

// Names returns the names of the organizations
func (o OrganizationList) Names() []string {
	names := []string{}
	for n := range o {
		names = append(names, n)
	}
	return names
}

// LoadOrganizationList fetches a list of organizations for the logged in
// user and also reads the currently selected org from the local config
// (and validates if the currently selected org is still valid)
func LoadOrganizationList() (OrganizationList, error) {
	req, _ := http.NewRequest("GET", orgListURL, nil)
	req.Header.Add("Authorization", config.GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orgListResponse struct {
		// ignoring other fields
		Data []Organization `json:"data"`
	}

	if err = json.Unmarshal([]byte(body), &orgListResponse); err != nil {
		return nil, err
	}

	orgList := OrganizationList{}
	for _, org := range orgListResponse.Data {
		orgList[org.Name] = org
		if currentOrgName == org.Name {
			currentOrg = &org
		}
	}

	return orgList, nil
}
