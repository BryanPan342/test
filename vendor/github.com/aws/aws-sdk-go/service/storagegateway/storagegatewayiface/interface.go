// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package storagegatewayiface provides an interface for the AWS Storage Gateway.
package storagegatewayiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/storagegateway"
)

// StorageGatewayAPI is the interface type for storagegateway.StorageGateway.
type StorageGatewayAPI interface {
	ActivateGatewayRequest(*storagegateway.ActivateGatewayInput) (*request.Request, *storagegateway.ActivateGatewayOutput)

	ActivateGateway(*storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error)

	AddCacheRequest(*storagegateway.AddCacheInput) (*request.Request, *storagegateway.AddCacheOutput)

	AddCache(*storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error)

	AddTagsToResourceRequest(*storagegateway.AddTagsToResourceInput) (*request.Request, *storagegateway.AddTagsToResourceOutput)

	AddTagsToResource(*storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error)

	AddUploadBufferRequest(*storagegateway.AddUploadBufferInput) (*request.Request, *storagegateway.AddUploadBufferOutput)

	AddUploadBuffer(*storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error)

	AddWorkingStorageRequest(*storagegateway.AddWorkingStorageInput) (*request.Request, *storagegateway.AddWorkingStorageOutput)

	AddWorkingStorage(*storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error)

	CancelArchivalRequest(*storagegateway.CancelArchivalInput) (*request.Request, *storagegateway.CancelArchivalOutput)

	CancelArchival(*storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error)

	CancelRetrievalRequest(*storagegateway.CancelRetrievalInput) (*request.Request, *storagegateway.CancelRetrievalOutput)

	CancelRetrieval(*storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error)

	CreateCachediSCSIVolumeRequest(*storagegateway.CreateCachediSCSIVolumeInput) (*request.Request, *storagegateway.CreateCachediSCSIVolumeOutput)

	CreateCachediSCSIVolume(*storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error)

	CreateSnapshotRequest(*storagegateway.CreateSnapshotInput) (*request.Request, *storagegateway.CreateSnapshotOutput)

	CreateSnapshot(*storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error)

	CreateSnapshotFromVolumeRecoveryPointRequest(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*request.Request, *storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput)

	CreateSnapshotFromVolumeRecoveryPoint(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error)

	CreateStorediSCSIVolumeRequest(*storagegateway.CreateStorediSCSIVolumeInput) (*request.Request, *storagegateway.CreateStorediSCSIVolumeOutput)

	CreateStorediSCSIVolume(*storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error)

	CreateTapesRequest(*storagegateway.CreateTapesInput) (*request.Request, *storagegateway.CreateTapesOutput)

	CreateTapes(*storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error)

	DeleteBandwidthRateLimitRequest(*storagegateway.DeleteBandwidthRateLimitInput) (*request.Request, *storagegateway.DeleteBandwidthRateLimitOutput)

	DeleteBandwidthRateLimit(*storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error)

	DeleteChapCredentialsRequest(*storagegateway.DeleteChapCredentialsInput) (*request.Request, *storagegateway.DeleteChapCredentialsOutput)

	DeleteChapCredentials(*storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error)

	DeleteGatewayRequest(*storagegateway.DeleteGatewayInput) (*request.Request, *storagegateway.DeleteGatewayOutput)

	DeleteGateway(*storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error)

	DeleteSnapshotScheduleRequest(*storagegateway.DeleteSnapshotScheduleInput) (*request.Request, *storagegateway.DeleteSnapshotScheduleOutput)

	DeleteSnapshotSchedule(*storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error)

	DeleteTapeRequest(*storagegateway.DeleteTapeInput) (*request.Request, *storagegateway.DeleteTapeOutput)

	DeleteTape(*storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error)

	DeleteTapeArchiveRequest(*storagegateway.DeleteTapeArchiveInput) (*request.Request, *storagegateway.DeleteTapeArchiveOutput)

	DeleteTapeArchive(*storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error)

	DeleteVolumeRequest(*storagegateway.DeleteVolumeInput) (*request.Request, *storagegateway.DeleteVolumeOutput)

	DeleteVolume(*storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error)

	DescribeBandwidthRateLimitRequest(*storagegateway.DescribeBandwidthRateLimitInput) (*request.Request, *storagegateway.DescribeBandwidthRateLimitOutput)

	DescribeBandwidthRateLimit(*storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error)

	DescribeCacheRequest(*storagegateway.DescribeCacheInput) (*request.Request, *storagegateway.DescribeCacheOutput)

	DescribeCache(*storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error)

	DescribeCachediSCSIVolumesRequest(*storagegateway.DescribeCachediSCSIVolumesInput) (*request.Request, *storagegateway.DescribeCachediSCSIVolumesOutput)

	DescribeCachediSCSIVolumes(*storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error)

	DescribeChapCredentialsRequest(*storagegateway.DescribeChapCredentialsInput) (*request.Request, *storagegateway.DescribeChapCredentialsOutput)

	DescribeChapCredentials(*storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error)

	DescribeGatewayInformationRequest(*storagegateway.DescribeGatewayInformationInput) (*request.Request, *storagegateway.DescribeGatewayInformationOutput)

	DescribeGatewayInformation(*storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error)

	DescribeMaintenanceStartTimeRequest(*storagegateway.DescribeMaintenanceStartTimeInput) (*request.Request, *storagegateway.DescribeMaintenanceStartTimeOutput)

	DescribeMaintenanceStartTime(*storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error)

	DescribeSnapshotScheduleRequest(*storagegateway.DescribeSnapshotScheduleInput) (*request.Request, *storagegateway.DescribeSnapshotScheduleOutput)

	DescribeSnapshotSchedule(*storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error)

	DescribeStorediSCSIVolumesRequest(*storagegateway.DescribeStorediSCSIVolumesInput) (*request.Request, *storagegateway.DescribeStorediSCSIVolumesOutput)

	DescribeStorediSCSIVolumes(*storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error)

	DescribeTapeArchivesRequest(*storagegateway.DescribeTapeArchivesInput) (*request.Request, *storagegateway.DescribeTapeArchivesOutput)

	DescribeTapeArchives(*storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error)

	DescribeTapeArchivesPages(*storagegateway.DescribeTapeArchivesInput, func(*storagegateway.DescribeTapeArchivesOutput, bool) bool) error

	DescribeTapeRecoveryPointsRequest(*storagegateway.DescribeTapeRecoveryPointsInput) (*request.Request, *storagegateway.DescribeTapeRecoveryPointsOutput)

	DescribeTapeRecoveryPoints(*storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error)

	DescribeTapeRecoveryPointsPages(*storagegateway.DescribeTapeRecoveryPointsInput, func(*storagegateway.DescribeTapeRecoveryPointsOutput, bool) bool) error

	DescribeTapesRequest(*storagegateway.DescribeTapesInput) (*request.Request, *storagegateway.DescribeTapesOutput)

	DescribeTapes(*storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error)

	DescribeTapesPages(*storagegateway.DescribeTapesInput, func(*storagegateway.DescribeTapesOutput, bool) bool) error

	DescribeUploadBufferRequest(*storagegateway.DescribeUploadBufferInput) (*request.Request, *storagegateway.DescribeUploadBufferOutput)

	DescribeUploadBuffer(*storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error)

	DescribeVTLDevicesRequest(*storagegateway.DescribeVTLDevicesInput) (*request.Request, *storagegateway.DescribeVTLDevicesOutput)

	DescribeVTLDevices(*storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error)

	DescribeVTLDevicesPages(*storagegateway.DescribeVTLDevicesInput, func(*storagegateway.DescribeVTLDevicesOutput, bool) bool) error

	DescribeWorkingStorageRequest(*storagegateway.DescribeWorkingStorageInput) (*request.Request, *storagegateway.DescribeWorkingStorageOutput)

	DescribeWorkingStorage(*storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error)

	DisableGatewayRequest(*storagegateway.DisableGatewayInput) (*request.Request, *storagegateway.DisableGatewayOutput)

	DisableGateway(*storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error)

	ListGatewaysRequest(*storagegateway.ListGatewaysInput) (*request.Request, *storagegateway.ListGatewaysOutput)

	ListGateways(*storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error)

	ListGatewaysPages(*storagegateway.ListGatewaysInput, func(*storagegateway.ListGatewaysOutput, bool) bool) error

	ListLocalDisksRequest(*storagegateway.ListLocalDisksInput) (*request.Request, *storagegateway.ListLocalDisksOutput)

	ListLocalDisks(*storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error)

	ListTagsForResourceRequest(*storagegateway.ListTagsForResourceInput) (*request.Request, *storagegateway.ListTagsForResourceOutput)

	ListTagsForResource(*storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error)

	ListVolumeInitiatorsRequest(*storagegateway.ListVolumeInitiatorsInput) (*request.Request, *storagegateway.ListVolumeInitiatorsOutput)

	ListVolumeInitiators(*storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error)

	ListVolumeRecoveryPointsRequest(*storagegateway.ListVolumeRecoveryPointsInput) (*request.Request, *storagegateway.ListVolumeRecoveryPointsOutput)

	ListVolumeRecoveryPoints(*storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error)

	ListVolumesRequest(*storagegateway.ListVolumesInput) (*request.Request, *storagegateway.ListVolumesOutput)

	ListVolumes(*storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error)

	ListVolumesPages(*storagegateway.ListVolumesInput, func(*storagegateway.ListVolumesOutput, bool) bool) error

	RemoveTagsFromResourceRequest(*storagegateway.RemoveTagsFromResourceInput) (*request.Request, *storagegateway.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error)

	ResetCacheRequest(*storagegateway.ResetCacheInput) (*request.Request, *storagegateway.ResetCacheOutput)

	ResetCache(*storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error)

	RetrieveTapeArchiveRequest(*storagegateway.RetrieveTapeArchiveInput) (*request.Request, *storagegateway.RetrieveTapeArchiveOutput)

	RetrieveTapeArchive(*storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error)

	RetrieveTapeRecoveryPointRequest(*storagegateway.RetrieveTapeRecoveryPointInput) (*request.Request, *storagegateway.RetrieveTapeRecoveryPointOutput)

	RetrieveTapeRecoveryPoint(*storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error)

	ShutdownGatewayRequest(*storagegateway.ShutdownGatewayInput) (*request.Request, *storagegateway.ShutdownGatewayOutput)

	ShutdownGateway(*storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error)

	StartGatewayRequest(*storagegateway.StartGatewayInput) (*request.Request, *storagegateway.StartGatewayOutput)

	StartGateway(*storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error)

	UpdateBandwidthRateLimitRequest(*storagegateway.UpdateBandwidthRateLimitInput) (*request.Request, *storagegateway.UpdateBandwidthRateLimitOutput)

	UpdateBandwidthRateLimit(*storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error)

	UpdateChapCredentialsRequest(*storagegateway.UpdateChapCredentialsInput) (*request.Request, *storagegateway.UpdateChapCredentialsOutput)

	UpdateChapCredentials(*storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error)

	UpdateGatewayInformationRequest(*storagegateway.UpdateGatewayInformationInput) (*request.Request, *storagegateway.UpdateGatewayInformationOutput)

	UpdateGatewayInformation(*storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error)

	UpdateGatewaySoftwareNowRequest(*storagegateway.UpdateGatewaySoftwareNowInput) (*request.Request, *storagegateway.UpdateGatewaySoftwareNowOutput)

	UpdateGatewaySoftwareNow(*storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error)

	UpdateMaintenanceStartTimeRequest(*storagegateway.UpdateMaintenanceStartTimeInput) (*request.Request, *storagegateway.UpdateMaintenanceStartTimeOutput)

	UpdateMaintenanceStartTime(*storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error)

	UpdateSnapshotScheduleRequest(*storagegateway.UpdateSnapshotScheduleInput) (*request.Request, *storagegateway.UpdateSnapshotScheduleOutput)

	UpdateSnapshotSchedule(*storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error)

	UpdateVTLDeviceTypeRequest(*storagegateway.UpdateVTLDeviceTypeInput) (*request.Request, *storagegateway.UpdateVTLDeviceTypeOutput)

	UpdateVTLDeviceType(*storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error)
}

var _ StorageGatewayAPI = (*storagegateway.StorageGateway)(nil)
