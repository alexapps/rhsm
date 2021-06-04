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

	"github.com/go-logr/logr"

	errs "github.com/pkg/errors"
)

const (
	RHSMBaseURL      = "subscription.rhn.redhat.com"
	RHELManagementV1 = "api.access.redhat.com/management/v1"
)

// RHSMClient -
type RHSMClient struct {
	User       string
	Password   string
	httpClient *http.Client
	log        logr.Logger
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
func (c *RHSMClient) DeleteSubscription(systemUUID, entitlementID string) error {
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

func (c *RHSMClient) DeleteConsumer(consumerUUID string) error {
	path := fmt.Sprintf("https://%s/subscription/consumers/%s", RHSMBaseURL, consumerUUID)
	fmt.Println("GetConsumer path: ", path)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.User, c.Password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errs.Wrap(err, "error in DeleteConsumer: read responce")
	}
	fmt.Println("RHSMClient responce: ", string(bodyBytes))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errs.Wrap(err, "error DeleteConsumer from RHSM")
	}
	return nil
}

func (c *RHSMClient) GetConsumer(consumerUUID string) error {
	path := fmt.Sprintf("https://%s/subscription/consumers/%s", RHSMBaseURL, consumerUUID)
	fmt.Println("GetConsumer path: ", path)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.User, c.Password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errs.Wrap(err, "error in GetSystem: read responce")
	}
	fmt.Println("RHSMClient responce: ", string(bodyBytes))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errs.Wrap(err, "error get the consumer from RHSM")
	}
	return nil
}

func getEntitlements(raw []byte) ([]*EntitlementDto, error) {
	var entitlements []*EntitlementDto
	err := json.Unmarshal(raw, &entitlements)
	return entitlements, err
}

func (c *RHSMClient) SetLogger(l logr.Logger) {
	c.log = l
}
