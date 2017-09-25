/* 
 * JumpCloud APIs
 *
 * V1 and V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package v1

type InlineResponse200Results struct {

	Id string `json:"_id,omitempty"`

	Active bool `json:"active,omitempty"`

	Name string `json:"name,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	DisplayLabel string `json:"displayLabel,omitempty"`

	Organization string `json:"organization,omitempty"`

	SsoUrl string `json:"ssoUrl,omitempty"`

	LearnMore string `json:"learnMore,omitempty"`

	Config InlineResponse200Config `json:"config,omitempty"`
}