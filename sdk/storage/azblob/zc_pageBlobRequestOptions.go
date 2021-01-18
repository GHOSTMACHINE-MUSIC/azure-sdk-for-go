package azblob

import (
	"fmt"
	"strconv"
)

func rangeToString(offset, count int64) string {
	return "bytes=" + strconv.FormatInt(offset, 10) + "-" + strconv.FormatInt(offset+count-1, 10)
}

func rangeToStringPtr(offset, count int64) *string {
	out := rangeToString(offset, count)
	return &out
}

func (pr PageRange) pointers() *string {
	startOffset := strconv.FormatInt(*pr.Start, 10)
	endOffset := strconv.FormatInt(*pr.End, 10)
	asString := fmt.Sprintf("bytes=%v-%s", startOffset, endOffset)
	return &asString
}

type CreatePageBlobOptions struct {
	// Set for page blobs only. The sequence number is a user-controlled value that you can use to track requests. The value of
	// the sequence number must be between 0 and 2^63 - 1.
	BlobSequenceNumber *int64
	// Optional. Used to set blob tags in various blob operations.
	BlobTagsMap *map[string]string
	// Optional. Specifies a user-defined name-value pair associated with the blob. If no name-value pairs are specified, the
	// operation will copy the metadata from the source blob or file to the destination blob. If one or more name-value pairs
	// are specified, the destination blob is created with the specified metadata, and metadata is not copied from the source
	// blob or file. Note that beginning with version 2009-09-19, metadata names must adhere to the naming rules for C# identifiers.
	// See Naming and Referencing Containers, Blobs, and Metadata for more information.
	Metadata *map[string]string
	// Optional. Indicates the tier to be set on the page blob.
	Tier *PremiumPageBlobAccessTier

	BlobHttpHeaders *BlobHttpHeaders
	CpkInfo         *CpkInfo
	CpkScopeInfo    *CpkScopeInfo
	BlobAccessConditions
}

func (o *CreatePageBlobOptions) pointers() (*PageBlobCreateOptions, *BlobHttpHeaders, *CpkInfo, *CpkScopeInfo, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil, nil
	}

	options := &PageBlobCreateOptions{
		BlobSequenceNumber: o.BlobSequenceNumber,
		BlobTagsString:     SerializeBlobTagsToStrPtr(o.BlobTagsMap),
		Metadata:           o.Metadata,
		Tier:               o.Tier,
	}

	return options, o.BlobHttpHeaders, o.CpkInfo, o.CpkScopeInfo, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type UploadPagesOptions struct {
	// Specify the transactional crc64 for the body, to be validated by the service.
	Offset                    *int64
	Count                     *int64
	TransactionalContentCrc64 *[]byte
	// Specify the transactional md5 for the body, to be validated by the service.
	TransactionalContentMd5 *[]byte

	CpkInfo                        *CpkInfo
	CpkScopeInfo                   *CpkScopeInfo
	SequenceNumberAccessConditions *SequenceNumberAccessConditions
	BlobAccessConditions
}

func (o *UploadPagesOptions) pointers() (*PageBlobUploadPagesOptions, *CpkInfo, *CpkScopeInfo, *SequenceNumberAccessConditions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil, nil
	}

	options := &PageBlobUploadPagesOptions{
		RangeParameter:            getSourceRange(o.Offset, o.Count),
		TransactionalContentCrc64: o.TransactionalContentCrc64,
		TransactionalContentMd5:   o.TransactionalContentMd5,
	}

	return options, o.CpkInfo, o.CpkScopeInfo, o.SequenceNumberAccessConditions, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type UploadPagesFromURLOptions struct {
	// Specify the md5 calculated for the range of bytes that must be read from the copy source.
	SourceContentMd5 *[]byte
	// Specify the crc64 calculated for the range of bytes that must be read from the copy source.
	SourceContentcrc64 *[]byte

	CpkInfo                        *CpkInfo
	CpkScopeInfo                   *CpkScopeInfo
	SequenceNumberAccessConditions *SequenceNumberAccessConditions
	SourceModifiedAccessConditions *SourceModifiedAccessConditions
	BlobAccessConditions
}

func (o *UploadPagesFromURLOptions) pointers() (*PageBlobUploadPagesFromURLOptions, *CpkInfo, *CpkScopeInfo, *SequenceNumberAccessConditions, *SourceModifiedAccessConditions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil, nil, nil
	}

	options := &PageBlobUploadPagesFromURLOptions{
		SourceContentMd5:   o.SourceContentMd5,
		SourceContentcrc64: o.SourceContentcrc64,
	}

	return options, o.CpkInfo, o.CpkScopeInfo, o.SequenceNumberAccessConditions, o.SourceModifiedAccessConditions, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type ClearPagesOptions struct {
	CpkInfo                        *CpkInfo
	CpkScopeInfo                   *CpkScopeInfo
	SequenceNumberAccessConditions *SequenceNumberAccessConditions
	BlobAccessConditions
}

func (o *ClearPagesOptions) pointers() (*CpkInfo, *CpkScopeInfo, *SequenceNumberAccessConditions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil, nil
	}

	return o.CpkInfo, o.CpkScopeInfo, o.SequenceNumberAccessConditions, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type GetPageRangesOptions struct {
	Snapshot *string

	BlobAccessConditions
}

func (o *GetPageRangesOptions) pointers() (*string, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	return o.Snapshot, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type ResizePageBlobOptions struct {
	CpkInfo      *CpkInfo
	CpkScopeInfo *CpkScopeInfo
	BlobAccessConditions
}

func (o *ResizePageBlobOptions) pointers() (*CpkInfo, *CpkScopeInfo, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil
	}

	return o.CpkInfo, o.CpkScopeInfo, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type UpdateSequenceNumberPageBlob struct {
	ActionType         *SequenceNumberActionType
	BlobSequenceNumber *int64

	BlobAccessConditions
}

func (o *UpdateSequenceNumberPageBlob) pointers() (*PageBlobUpdateSequenceNumberOptions, *SequenceNumberActionType, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil, nil
	}

	options := &PageBlobUpdateSequenceNumberOptions{
		BlobSequenceNumber: o.BlobSequenceNumber,
	}

	if *o.ActionType == SequenceNumberActionTypeIncrement {
		options.BlobSequenceNumber = nil
	}

	return options, o.ActionType, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type CopyIncrementalPageBlobOptions struct {
	ModifiedAccessConditions *ModifiedAccessConditions

	RequestId *string

	Timeout *int32
}

func (o *CopyIncrementalPageBlobOptions) pointers() (*PageBlobCopyIncrementalOptions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil
	}

	options := PageBlobCopyIncrementalOptions{
		RequestId: o.RequestId,
		Timeout:   o.Timeout,
	}

	return &options, o.ModifiedAccessConditions
}