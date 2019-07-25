// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package marketplacecommerceanalytics provides a client for AWS Marketplace Commerce Analytics.
package marketplacecommerceanalytics

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGenerateDataSet = "GenerateDataSet"

// GenerateDataSetRequest generates a request for the GenerateDataSet operation.
func (c *MarketplaceCommerceAnalytics) GenerateDataSetRequest(input *GenerateDataSetInput) (req *request.Request, output *GenerateDataSetOutput) {
	op := &request.Operation{
		Name:       opGenerateDataSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataSetInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GenerateDataSetOutput{}
	req.Data = output
	return
}

// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already
// exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached permissions
// policy providing Allow permissions for the following actions: s3:PutObject,
// s3:getBucketLocation, sns:SetRegion, sns:ListTopics, sns:Publish, iam:GetRolePolicy.
func (c *MarketplaceCommerceAnalytics) GenerateDataSet(input *GenerateDataSetInput) (*GenerateDataSetOutput, error) {
	req, out := c.GenerateDataSetRequest(input)
	err := req.Send()
	return out, err
}

// Container for the parameters to the GenerateDataSet operation.
type GenerateDataSetInput struct {
	_ struct{} `type:"structure"`

	// The date a data set was published. For daily data sets, provide a date with
	// day-level granularity for the desired day. For weekly data sets, provide
	// a date with day-level granularity within the desired week (the day value
	// will be ignored). For monthly data sets, provide a date with month-level
	// granularity for the desired month (the day value will be ignored).
	DataSetPublicationDate *time.Time `locationName:"dataSetPublicationDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The type of the data set to publish.
	DataSetType *string `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"DataSetType"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GenerateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateDataSetInput) GoString() string {
	return s.String()
}

// Container for the result of the GenerateDataSet operation.
type GenerateDataSetOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the GenerateDataSet
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s GenerateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateDataSetOutput) GoString() string {
	return s.String()
}

// The type of the data set to publish.
const (
	// @enum DataSetType
	DataSetTypeCustomerSubscriberHourlyMonthlySubscriptions = "customer_subscriber_hourly_monthly_subscriptions"
	// @enum DataSetType
	DataSetTypeCustomerSubscriberAnnualSubscriptions = "customer_subscriber_annual_subscriptions"
	// @enum DataSetType
	DataSetTypeDailyBusinessUsageByInstanceType = "daily_business_usage_by_instance_type"
	// @enum DataSetType
	DataSetTypeDailyBusinessFees = "daily_business_fees"
	// @enum DataSetType
	DataSetTypeDailyBusinessFreeTrialConversions = "daily_business_free_trial_conversions"
	// @enum DataSetType
	DataSetTypeDailyBusinessNewInstances = "daily_business_new_instances"
	// @enum DataSetType
	DataSetTypeDailyBusinessNewProductSubscribers = "daily_business_new_product_subscribers"
	// @enum DataSetType
	DataSetTypeDailyBusinessCanceledProductSubscribers = "daily_business_canceled_product_subscribers"
	// @enum DataSetType
	DataSetTypeMonthlyRevenueBillingAndRevenueData = "monthly_revenue_billing_and_revenue_data"
	// @enum DataSetType
	DataSetTypeMonthlyRevenueAnnualSubscriptions = "monthly_revenue_annual_subscriptions"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByProduct = "disbursed_amount_by_product"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByCustomerGeo = "disbursed_amount_by_customer_geo"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByAgeOfUncollectedFunds = "disbursed_amount_by_age_of_uncollected_funds"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByAgeOfDisbursedFunds = "disbursed_amount_by_age_of_disbursed_funds"
	// @enum DataSetType
	DataSetTypeCustomerProfileByIndustry = "customer_profile_by_industry"
	// @enum DataSetType
	DataSetTypeCustomerProfileByRevenue = "customer_profile_by_revenue"
)
