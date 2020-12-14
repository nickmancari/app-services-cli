/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package managedservices

// ServiceAccountListItem struct for ServiceAccountListItem
type ServiceAccountListItem struct {
	// server generated unique id of the service account
	Id          string      `json:"id,omitempty"`
	Kind        string      `json:"kind,omitempty"`
	Href        string      `json:"href,omitempty"`
	ClientID    string      `json:"clientID,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description interface{} `json:"description,omitempty"`
}