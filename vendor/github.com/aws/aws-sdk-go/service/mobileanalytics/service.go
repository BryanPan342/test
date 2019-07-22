// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package mobileanalytics

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Amazon Mobile Analytics is a service for collecting, visualizing, and understanding
// app usage data at scale.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type MobileAnalytics struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "mobileanalytics"

// New creates a new instance of the MobileAnalytics client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MobileAnalytics client from just a session.
//     svc := mobileanalytics.New(mySession)
//
//     // Create a MobileAnalytics client with additional configuration
//     svc := mobileanalytics.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *MobileAnalytics {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *MobileAnalytics {
	svc := &MobileAnalytics{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-06-05",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(restjson.Build)
	svc.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a MobileAnalytics operation and runs any
// custom request initialization.
func (c *MobileAnalytics) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
