/*
Copyright 2021 Alexaps
*/

package licensecontroller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errs "github.com/pkg/errors"
)

const (
	RHSMBaseURL = "subscription.rhn.redhat.com"
)

// RHSMClient -
type RHSMClient struct {
	User       string
	Password   string
	httpClient *http.Client
}

// NewClient implements the RHSM API Client
func NewRHSMClient() *RHSMClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
	}
	httpClient := &http.Client{Transport: tr}
	return &RHSMClient{httpClient: httpClient}
}

// SetCredentials -
func (c *RHSMClient) SetCredentials(user, password string) {
	c.User = user
	c.Password = password
}

// GetEntitlement return related to the system UUID entitlement. Get the first one
func (c *RHSMClient) GetEntitlement(systemUUID string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/subscription/consumers/%s/entitlements", RHSMBaseURL, systemUUID), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.User, c.Password)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		fmt.Println("RHSMClient responce: ", string(bodyBytes))
		entitlements, err := getEntitlements(bodyBytes)
		if err != nil {
			return "", err
		}
		if len(entitlements) > 0 {
			return entitlements[0].ID, nil
		}
		return "", errs.Wrap(err, "error: the entitlements are empty")

	} else {
		return "", errs.Wrap(err, "error getting the entitlements from RHSM")
	}
}

// Delete removes an Entitlement from a Consumer By the Entitlement ID
func (c *RHSMClient) Delete(systemUUID, entitlementID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://%s/subscription/consumers/%s/entitlements/%s", RHSMBaseURL, systemUUID, entitlementID), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.User, c.Password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errs.Wrap(err, "error remove the entitlement from RHSM")
	}
	return nil
}

func getEntitlements(raw []byte) ([]*EntitlementDto, error) {
	var entitlements []*EntitlementDto
	err := json.Unmarshal(raw, &entitlements)
	return entitlements, err
}
