// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// NewUserDelegationCredential creates a new UserDelegationCredential using a Storage account's name and a user delegation key from it
func NewUserDelegationCredential(accountName string, key UserDelegationKey) UserDelegationCredential {
	return UserDelegationCredential{
		accountName: accountName,
		accountKey:  key,
	}
}

type UserDelegationCredential struct {
	accountName string
	accountKey  UserDelegationKey
}

// AccountName returns the Storage account's name
func (f UserDelegationCredential) AccountName() string {
	return f.accountName
}

// ComputeHMACSHA256 generates a hash signature for an HTTP request or for a SAS.
func (f UserDelegationCredential) ComputeHMACSHA256(message string) (base64String string) {
	bytes, _ := base64.StdEncoding.DecodeString(*f.accountKey.Value)
	h := hmac.New(sha256.New, bytes)
	_, err := h.Write([]byte(message))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// GetUDKParams : Private method to return important parameters for NewSASQueryParameters
func (f UserDelegationCredential) GetUDKParams() UserDelegationKey {
	return f.accountKey
}