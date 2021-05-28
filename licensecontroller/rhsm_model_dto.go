/*
Copyright 2021 Alexaps
*/

package licensecontroller

//EntitlementDto DTO representing an entitlement
type EntitlementDto struct {
	Created      string            `json:"created,omitempty"`
	Updated      string            `json:"updated,omitempty"`
	ID           string            `json:"id,omitempty"`
	Consumer     NestedConsumerDto `json:"consumer,omitempty"`
	Pool         Pool              `json:"pool,omitempty"`
	Quantity     int32             `json:"quantity,omitempty"`
	Certificates []CertificateDto  `json:"certificates,omitempty"`
	StartDate    string            `json:"startDate,omitempty"`
	EndDate      string            `json:"endDate,omitempty"`
	Href         string            `json:"href,omitempty"`
}

// NestedConsumerDto DTO representing an upstream consumer
type NestedConsumerDto struct {
	ID   string `json:"id,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

// CertificateDto DTO representing a certificate
type CertificateDto struct {
	Created string               `json:"created,omitempty"`
	Updated string               `json:"updated,omitempty"`
	ID      string               `json:"id,omitempty"`
	Key     string               `json:"key,omitempty"`
	Cert    string               `json:"cert,omitempty"`
	Serial  CertificateSerialDto `json:"serial,omitempty"`
}

// CertificateSerialDto DTO representing a certificate serial
type CertificateSerialDto struct {
	Created    string `json:"created,omitempty"`
	Updated    string `json:"updated,omitempty"`
	ID         int64  `json:"id,omitempty"`
	Serial     int64  `json:"serial,omitempty"`
	Expiration string `json:"expiration,omitempty"`
	Collected  bool   `json:"collected,omitempty"`
	Revoked    bool   `json:"revoked,omitempty"`
}

// Pool DTO representing a Pool
type Pool struct {
	Created                  string               `json:"created,omitempty"`
	Updated                  string               `json:"updated,omitempty"`
	ID                       string               `json:"id,omitempty"`
	Type                     string               `json:"type,omitempty"`
	Owner                    NestedOwnerDto       `json:"owner,omitempty"`
	ActiveSubscription       bool                 `json:"activeSubscription,omitempty"`
	CreatedByShare           bool                 `json:"createdByShare,omitempty"`
	HasSharedAncestor        bool                 `json:"hasSharedAncestor,omitempty"`
	SourceEntitlement        NestedEntitlementDto `json:"sourceEntitlement,omitempty"`
	Quantity                 int64                `json:"quantity,omitempty"`
	StartDate                string               `json:"startDate,omitempty"`
	EndDate                  string               `json:"endDate,omitempty"`
	Attributes               interface{}          `json:"attributes,omitempty"`
	RestrictedToUsername     string               `json:"restrictedToUsername,omitempty"`
	ContractNumber           string               `json:"contractNumber,omitempty"`
	AccountNumber            string               `json:"accountNumber,omitempty"`
	OrderNumber              string               `json:"orderNumber,omitempty"`
	Consumed                 int64                `json:"consumed,omitempty"`
	Exported                 int64                `json:"exported,omitempty"`
	Shared                   int64                `json:"shared,omitempty"`
	Branding                 []BrandingDto        `json:"branding,omitempty"`
	CalculatedAttributes     interface{}          `json:"calculatedAttributes,omitempty"`
	UpstreamPoolID           string               `json:"upstreamPoolId,omitempty"`
	UpstreamEntitlementID    string               `json:"upstreamEntitlementId,omitempty"`
	UpstreamConsumerID       string               `json:"upstreamConsumerId,omitempty"`
	ProductName              string               `json:"productName,omitempty"`
	ProductID                string               `json:"productId,omitempty"`
	ProductAttributes        interface{}          `json:"productAttributes,omitempty"`
	StackID                  string               `json:"stackId,omitempty"`
	Stacked                  bool                 `json:"stacked,omitempty"`
	SourceStackID            string               `json:"sourceStackId,omitempty"`
	DevelopmentPool          bool                 `json:"developmentPool,omitempty"`
	DerivedProductAttributes interface{}          `json:"derivedProductAttributes,omitempty"`
	DerivedProductID         string               `json:"derivedProductId,omitempty"`
	DerivedProductName       string               `json:"derivedProductName,omitempty"`
	ProvidedProducts         []ProvidedProductDto `json:"providedProducts,omitempty"`
	DerivedProvidedProducts  []ProvidedProductDto `json:"derivedProvidedProducts,omitempty"`
	SubscriptionSubKey       string               `json:"subscriptionSubKey,omitempty"`
	SubscriptionID           string               `json:"subscriptionId,omitempty"`
	Href                     string               `json:"href,omitempty"`
}

// NestedOwnerDto DTO representing an owner/organization
type NestedOwnerDto struct {
	ID          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Href        string `json:"href,omitempty"`
}

// NestedEntitlementDto DTO representing an entitlement
type NestedEntitlementDto struct {
	ID   string `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

// BrandingDto DTO representing an entitlement
type BrandingDto struct {
	Created   string `json:"created,omitempty"`
	Updated   string `json:"updated,omitempty"`
	ProductID string `json:"productId,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
}

// ProvidedProductDto DTO representing an entitlement
type ProvidedProductDto struct {
	ProductID   string `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
}
