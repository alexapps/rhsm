/*
Copyright 2020 The Mirantis Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
		entitlements, err := getEntitlements(bodyBytes)
		if err != nil {
			return "", err
		}
		if len(entitlements) > 1 {
			return entitlements[0].ID, nil
		}
	} else {
		return "", errs.Wrap(err, "error getting the entitlements from RHSM")
	}
	return "", nil
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

func getEntitlements(raw []byte) ([]EntitlementDto, error) {
	var entitlements []EntitlementDto
	err := json.Unmarshal(raw, &entitlements)
	return entitlements, err
}
