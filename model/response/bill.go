/*
 * air_condition
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package response

type Bill struct {

	// follow the unix timestamps
	StartTime string `json:"start_time"`

	// follow the unix timestamps
	EndTime string `json:"end_time"`

	Speed string `json:"speed"`

	FromTem string `json:"from_tem"`

	ToTem string `json:"to_tem"`
}
