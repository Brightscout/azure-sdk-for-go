//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package blockblob

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/shared"
)

// Type Declarations ---------------------------------------------------------------------

// Block - Represents a single block in a block blob. It describes the block's ID and size.
type Block = generated.Block

// BlockList - type of blocklist (committed/uncommitted)
type BlockList = generated.BlockList

// Request Model Declaration -------------------------------------------------------------------------------------------

// UploadOptions contains the optional parameters for the Client.Upload method.
type UploadOptions struct {
	// Optional. Used to set blob tags in various blob operations.
	Tags map[string]string

	// Optional. Specifies a user-defined name-value pair associated with the blob.
	Metadata map[string]*string

	// Optional. Indicates the tier to be set on the blob.
	Tier *blob.AccessTier

	// Specify the transactional md5 for the body, to be validated by the service.
	TransactionalContentMD5 []byte

	HTTPHeaders                  *blob.HTTPHeaders
	CpkInfo                      *blob.CpkInfo
	CpkScopeInfo                 *blob.CpkScopeInfo
	AccessConditions             *blob.AccessConditions
	LegalHold                    *bool
	ImmutabilityPolicyMode       *blob.ImmutabilityPolicySetting
	ImmutabilityPolicyExpiryTime *time.Time
}

func (o *UploadOptions) format() (*generated.BlockBlobClientUploadOptions, *generated.BlobHTTPHeaders, *generated.LeaseAccessConditions,
	*generated.CpkInfo, *generated.CpkScopeInfo, *generated.ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil, nil
	}

	basics := generated.BlockBlobClientUploadOptions{
		BlobTagsString:           shared.SerializeBlobTagsToStrPtr(o.Tags),
		Metadata:                 o.Metadata,
		Tier:                     o.Tier,
		TransactionalContentMD5:  o.TransactionalContentMD5,
		LegalHold:                o.LegalHold,
		ImmutabilityPolicyMode:   o.ImmutabilityPolicyMode,
		ImmutabilityPolicyExpiry: o.ImmutabilityPolicyExpiryTime,
	}

	leaseAccessConditions, modifiedAccessConditions := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &basics, o.HTTPHeaders, leaseAccessConditions, o.CpkInfo, o.CpkScopeInfo, modifiedAccessConditions
}

// ---------------------------------------------------------------------------------------------------------------------

// StageBlockOptions contains the optional parameters for the Client.StageBlock method.
type StageBlockOptions struct {
	CpkInfo *blob.CpkInfo

	CpkScopeInfo *blob.CpkScopeInfo

	LeaseAccessConditions *blob.LeaseAccessConditions

	// TransactionalValidation specifies the transfer validation type to use.
	// The default is nil (no transfer validation).
	TransactionalValidation blob.TransferValidationType
}

// StageBlockOptions contains the optional parameters for the Client.StageBlock method.
func (o *StageBlockOptions) format() (*generated.BlockBlobClientStageBlockOptions, *generated.LeaseAccessConditions, *generated.CpkInfo, *generated.CpkScopeInfo) {
	if o == nil {
		return nil, nil, nil, nil
	}

	return &generated.BlockBlobClientStageBlockOptions{}, o.LeaseAccessConditions, o.CpkInfo, o.CpkScopeInfo
}

// ---------------------------------------------------------------------------------------------------------------------

// StageBlockFromURLOptions contains the optional parameters for the Client.StageBlockFromURL method.
type StageBlockFromURLOptions struct {
	// Only Bearer type is supported. Credentials should be a valid OAuth access token to copy source.
	CopySourceAuthorization *string

	LeaseAccessConditions *blob.LeaseAccessConditions

	SourceModifiedAccessConditions *blob.SourceModifiedAccessConditions

	// SourceContentValidation contains the validation mechanism used on the range of bytes read from the source.
	SourceContentValidation blob.SourceContentValidationType

	// Range specifies a range of bytes.  The default value is all bytes.
	Range blob.HTTPRange

	CpkInfo *blob.CpkInfo

	CpkScopeInfo *blob.CpkScopeInfo
}

func (o *StageBlockFromURLOptions) format() (*generated.BlockBlobClientStageBlockFromURLOptions, *generated.CpkInfo, *generated.CpkScopeInfo, *generated.LeaseAccessConditions, *generated.SourceModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil
	}

	options := &generated.BlockBlobClientStageBlockFromURLOptions{
		CopySourceAuthorization: o.CopySourceAuthorization,
		SourceRange:             exported.FormatHTTPRange(o.Range),
	}

	if o.SourceContentValidation != nil {
		o.SourceContentValidation.Apply(options)
	}

	return options, o.CpkInfo, o.CpkScopeInfo, o.LeaseAccessConditions, o.SourceModifiedAccessConditions
}

// ---------------------------------------------------------------------------------------------------------------------

// CommitBlockListOptions contains the optional parameters for Client.CommitBlockList method.
type CommitBlockListOptions struct {
	Tags                         map[string]string
	Metadata                     map[string]*string
	RequestID                    *string
	Tier                         *blob.AccessTier
	Timeout                      *int32
	TransactionalContentCRC64    []byte
	TransactionalContentMD5      []byte
	HTTPHeaders                  *blob.HTTPHeaders
	CpkInfo                      *blob.CpkInfo
	CpkScopeInfo                 *blob.CpkScopeInfo
	AccessConditions             *blob.AccessConditions
	LegalHold                    *bool
	ImmutabilityPolicyMode       *blob.ImmutabilityPolicySetting
	ImmutabilityPolicyExpiryTime *time.Time
}

// ---------------------------------------------------------------------------------------------------------------------

// GetBlockListOptions contains the optional parameters for the Client.GetBlockList method.
type GetBlockListOptions struct {
	Snapshot         *string
	AccessConditions *blob.AccessConditions
}

func (o *GetBlockListOptions) format() (*generated.BlockBlobClientGetBlockListOptions, *generated.LeaseAccessConditions, *generated.ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	leaseAccessConditions, modifiedAccessConditions := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &generated.BlockBlobClientGetBlockListOptions{Snapshot: o.Snapshot}, leaseAccessConditions, modifiedAccessConditions
}

// ------------------------------------------------------------

// uploadFromReaderOptions identifies options used by the UploadBuffer and UploadFile functions.
type uploadFromReaderOptions struct {
	// BlockSize specifies the block size to use; the default (and maximum size) is MaxStageBlockBytes.
	BlockSize int64

	// Progress is a function that is invoked periodically as bytes are sent to the BlockBlobClient.
	// Note that the progress reporting is not always increasing; it can go down when retrying a request.
	Progress func(bytesTransferred int64)

	// HTTPHeaders indicates the HTTP headers to be associated with the blob.
	HTTPHeaders *blob.HTTPHeaders

	// Metadata indicates the metadata to be associated with the blob when PutBlockList is called.
	Metadata map[string]*string

	// AccessConditions indicates the access conditions for the block blob.
	AccessConditions *blob.AccessConditions

	// AccessTier indicates the tier of blob
	AccessTier *blob.AccessTier

	// BlobTags
	Tags map[string]string

	// ClientProvidedKeyOptions indicates the client provided key by name and/or by value to encrypt/decrypt data.
	CpkInfo      *blob.CpkInfo
	CpkScopeInfo *blob.CpkScopeInfo

	// Concurrency indicates the maximum number of blocks to upload in parallel (0=default)
	Concurrency uint16

	TransactionalValidation blob.TransferValidationType

	// Optional header, Specifies the transactional crc64 for the body, to be validated by the service.
	TransactionalContentCRC64 uint64
	// Specify the transactional md5 for the body, to be validated by the service.
	TransactionalContentMD5 []byte
}

// UploadBufferOptions provides set of configurations for UploadBuffer operation
type UploadBufferOptions = uploadFromReaderOptions

// UploadFileOptions provides set of configurations for UploadFile operation
type UploadFileOptions = uploadFromReaderOptions

func (o *uploadFromReaderOptions) getStageBlockOptions() *StageBlockOptions {
	leaseAccessConditions, _ := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &StageBlockOptions{
		CpkInfo:               o.CpkInfo,
		CpkScopeInfo:          o.CpkScopeInfo,
		LeaseAccessConditions: leaseAccessConditions,

		TransactionalValidation: o.TransactionalValidation,
	}
}

func (o *uploadFromReaderOptions) getUploadBlockBlobOptions() *UploadOptions {
	return &UploadOptions{
		Tags:             o.Tags,
		Metadata:         o.Metadata,
		Tier:             o.AccessTier,
		HTTPHeaders:      o.HTTPHeaders,
		AccessConditions: o.AccessConditions,
		CpkInfo:          o.CpkInfo,
		CpkScopeInfo:     o.CpkScopeInfo,
	}
}

func (o *uploadFromReaderOptions) getCommitBlockListOptions() *CommitBlockListOptions {
	return &CommitBlockListOptions{
		Tags:         o.Tags,
		Metadata:     o.Metadata,
		Tier:         o.AccessTier,
		HTTPHeaders:  o.HTTPHeaders,
		CpkInfo:      o.CpkInfo,
		CpkScopeInfo: o.CpkScopeInfo,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// UploadStreamOptions provides set of configurations for UploadStream operation
type UploadStreamOptions struct {
	// BlockSize defines the size of the buffer used during upload. The default and mimimum value is 1 MiB.
	BlockSize int64

	// Concurrency defines the max number of concurrent uploads to be performed to upload the file.
	// Each concurrent upload will create a buffer of size BlockSize.  The default value is one.
	Concurrency int

	TransactionalValidation blob.TransferValidationType

	HTTPHeaders      *blob.HTTPHeaders
	Metadata         map[string]*string
	AccessConditions *blob.AccessConditions
	AccessTier       *blob.AccessTier
	Tags             map[string]string
	CpkInfo          *blob.CpkInfo
	CpkScopeInfo     *blob.CpkScopeInfo
}

func (u *UploadStreamOptions) setDefaults() {
	if u.Concurrency == 0 {
		u.Concurrency = 1
	}

	if u.BlockSize < _1MiB {
		u.BlockSize = _1MiB
	}
}

func (u *UploadStreamOptions) getStageBlockOptions() *StageBlockOptions {
	if u == nil {
		return nil
	}

	leaseAccessConditions, _ := exported.FormatBlobAccessConditions(u.AccessConditions)
	return &StageBlockOptions{
		TransactionalValidation: u.TransactionalValidation,
		CpkInfo:                 u.CpkInfo,
		CpkScopeInfo:            u.CpkScopeInfo,
		LeaseAccessConditions:   leaseAccessConditions,
	}
}

func (u *UploadStreamOptions) getCommitBlockListOptions() *CommitBlockListOptions {
	if u == nil {
		return nil
	}

	return &CommitBlockListOptions{
		Tags:             u.Tags,
		Metadata:         u.Metadata,
		Tier:             u.AccessTier,
		HTTPHeaders:      u.HTTPHeaders,
		CpkInfo:          u.CpkInfo,
		CpkScopeInfo:     u.CpkScopeInfo,
		AccessConditions: u.AccessConditions,
	}
}

func (u *UploadStreamOptions) getUploadOptions() *UploadOptions {
	if u == nil {
		return nil
	}

	return &UploadOptions{
		Tags:             u.Tags,
		Metadata:         u.Metadata,
		Tier:             u.AccessTier,
		HTTPHeaders:      u.HTTPHeaders,
		CpkInfo:          u.CpkInfo,
		CpkScopeInfo:     u.CpkScopeInfo,
		AccessConditions: u.AccessConditions,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// ExpiryType defines values for ExpiryType
type ExpiryType = exported.ExpiryType

// ExpiryTypeAbsolute defines the absolute time for the blob expiry
type ExpiryTypeAbsolute = exported.ExpiryTypeAbsolute

// ExpiryTypeRelativeToNow defines the duration relative to now for the blob expiry
type ExpiryTypeRelativeToNow = exported.ExpiryTypeRelativeToNow

// ExpiryTypeRelativeToCreation defines the duration relative to creation for the blob expiry
type ExpiryTypeRelativeToCreation = exported.ExpiryTypeRelativeToCreation

// ExpiryTypeNever defines that the blob will be set to never expire
type ExpiryTypeNever = exported.ExpiryTypeNever

// SetExpiryOptions contains the optional parameters for the Client.SetExpiry method.
type SetExpiryOptions = exported.SetExpiryOptions
