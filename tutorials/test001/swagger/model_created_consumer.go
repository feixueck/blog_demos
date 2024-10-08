/*
 * Strimzi Kafka Bridge API Reference
 *
 * The Strimzi Kafka Bridge provides a REST API for integrating HTTP based client applications with a Kafka cluster. You can use the API to create and manage consumers and send and receive records over HTTP rather than the native Kafka protocol. 
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CreatedConsumer struct {
	// Unique ID for the consumer instance in the group.
	InstanceId string `json:"instance_id,omitempty"`
	// Base URI used to construct URIs for subsequent requests against this consumer instance.
	BaseUri string `json:"base_uri,omitempty"`
}
