//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

const (
	module  = "armpeering"
	version = "v0.1.0"
)

// ConnectionState - The state of the connection.
type ConnectionState string

const (
	ConnectionStateActive                ConnectionState = "Active"
	ConnectionStateApproved              ConnectionState = "Approved"
	ConnectionStateNone                  ConnectionState = "None"
	ConnectionStatePendingApproval       ConnectionState = "PendingApproval"
	ConnectionStateProvisioningCompleted ConnectionState = "ProvisioningCompleted"
	ConnectionStateProvisioningFailed    ConnectionState = "ProvisioningFailed"
	ConnectionStateProvisioningStarted   ConnectionState = "ProvisioningStarted"
	ConnectionStateValidating            ConnectionState = "Validating"
)

// PossibleConnectionStateValues returns the possible values for the ConnectionState const type.
func PossibleConnectionStateValues() []ConnectionState {
	return []ConnectionState{
		ConnectionStateActive,
		ConnectionStateApproved,
		ConnectionStateNone,
		ConnectionStatePendingApproval,
		ConnectionStateProvisioningCompleted,
		ConnectionStateProvisioningFailed,
		ConnectionStateProvisioningStarted,
		ConnectionStateValidating,
	}
}

// ToPtr returns a *ConnectionState pointing to the current value.
func (c ConnectionState) ToPtr() *ConnectionState {
	return &c
}

// DirectPeeringType - The type of direct peering.
type DirectPeeringType string

const (
	DirectPeeringTypeCdn      DirectPeeringType = "Cdn"
	DirectPeeringTypeEdge     DirectPeeringType = "Edge"
	DirectPeeringTypeInternal DirectPeeringType = "Internal"
	DirectPeeringTypeTransit  DirectPeeringType = "Transit"
)

// PossibleDirectPeeringTypeValues returns the possible values for the DirectPeeringType const type.
func PossibleDirectPeeringTypeValues() []DirectPeeringType {
	return []DirectPeeringType{
		DirectPeeringTypeCdn,
		DirectPeeringTypeEdge,
		DirectPeeringTypeInternal,
		DirectPeeringTypeTransit,
	}
}

// ToPtr returns a *DirectPeeringType pointing to the current value.
func (c DirectPeeringType) ToPtr() *DirectPeeringType {
	return &c
}

type Enum0 string

const (
	Enum0Available   Enum0 = "Available"
	Enum0UnAvailable Enum0 = "UnAvailable"
)

// PossibleEnum0Values returns the possible values for the Enum0 const type.
func PossibleEnum0Values() []Enum0 {
	return []Enum0{
		Enum0Available,
		Enum0UnAvailable,
	}
}

// ToPtr returns a *Enum0 pointing to the current value.
func (c Enum0) ToPtr() *Enum0 {
	return &c
}

type Enum1 string

const (
	Enum1Direct   Enum1 = "Direct"
	Enum1Exchange Enum1 = "Exchange"
)

// PossibleEnum1Values returns the possible values for the Enum1 const type.
func PossibleEnum1Values() []Enum1 {
	return []Enum1{
		Enum1Direct,
		Enum1Exchange,
	}
}

// ToPtr returns a *Enum1 pointing to the current value.
func (c Enum1) ToPtr() *Enum1 {
	return &c
}

type Enum14 string

const (
	Enum14Direct   Enum14 = "Direct"
	Enum14Exchange Enum14 = "Exchange"
)

// PossibleEnum14Values returns the possible values for the Enum14 const type.
func PossibleEnum14Values() []Enum14 {
	return []Enum14{
		Enum14Direct,
		Enum14Exchange,
	}
}

// ToPtr returns a *Enum14 pointing to the current value.
func (c Enum14) ToPtr() *Enum14 {
	return &c
}

type Enum15 string

const (
	Enum15Cdn      Enum15 = "Cdn"
	Enum15Edge     Enum15 = "Edge"
	Enum15Internal Enum15 = "Internal"
	Enum15Transit  Enum15 = "Transit"
)

// PossibleEnum15Values returns the possible values for the Enum15 const type.
func PossibleEnum15Values() []Enum15 {
	return []Enum15{
		Enum15Cdn,
		Enum15Edge,
		Enum15Internal,
		Enum15Transit,
	}
}

// ToPtr returns a *Enum15 pointing to the current value.
func (c Enum15) ToPtr() *Enum15 {
	return &c
}

// Family - The family of the peering SKU.
type Family string

const (
	FamilyDirect   Family = "Direct"
	FamilyExchange Family = "Exchange"
)

// PossibleFamilyValues returns the possible values for the Family const type.
func PossibleFamilyValues() []Family {
	return []Family{
		FamilyDirect,
		FamilyExchange,
	}
}

// ToPtr returns a *Family pointing to the current value.
func (c Family) ToPtr() *Family {
	return &c
}

// Kind - The kind of the peering.
type Kind string

const (
	KindDirect   Kind = "Direct"
	KindExchange Kind = "Exchange"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindDirect,
		KindExchange,
	}
}

// ToPtr returns a *Kind pointing to the current value.
func (c Kind) ToPtr() *Kind {
	return &c
}

// LearnedType - The prefix learned type
type LearnedType string

const (
	LearnedTypeNone       LearnedType = "None"
	LearnedTypeViaPartner LearnedType = "ViaPartner"
	LearnedTypeViaSession LearnedType = "ViaSession"
)

// PossibleLearnedTypeValues returns the possible values for the LearnedType const type.
func PossibleLearnedTypeValues() []LearnedType {
	return []LearnedType{
		LearnedTypeNone,
		LearnedTypeViaPartner,
		LearnedTypeViaSession,
	}
}

// ToPtr returns a *LearnedType pointing to the current value.
func (c LearnedType) ToPtr() *LearnedType {
	return &c
}

// Name - The name of the peering SKU.
type Name string

const (
	NameBasicDirectFree        Name = "Basic_Direct_Free"
	NameBasicExchangeFree      Name = "Basic_Exchange_Free"
	NamePremiumDirectFree      Name = "Premium_Direct_Free"
	NamePremiumDirectMetered   Name = "Premium_Direct_Metered"
	NamePremiumDirectUnlimited Name = "Premium_Direct_Unlimited"
	NamePremiumExchangeMetered Name = "Premium_Exchange_Metered"
)

// PossibleNameValues returns the possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{
		NameBasicDirectFree,
		NameBasicExchangeFree,
		NamePremiumDirectFree,
		NamePremiumDirectMetered,
		NamePremiumDirectUnlimited,
		NamePremiumExchangeMetered,
	}
}

// ToPtr returns a *Name pointing to the current value.
func (c Name) ToPtr() *Name {
	return &c
}

// PrefixValidationState - The prefix validation state
type PrefixValidationState string

const (
	PrefixValidationStateFailed   PrefixValidationState = "Failed"
	PrefixValidationStateInvalid  PrefixValidationState = "Invalid"
	PrefixValidationStateNone     PrefixValidationState = "None"
	PrefixValidationStatePending  PrefixValidationState = "Pending"
	PrefixValidationStateUnknown  PrefixValidationState = "Unknown"
	PrefixValidationStateVerified PrefixValidationState = "Verified"
)

// PossiblePrefixValidationStateValues returns the possible values for the PrefixValidationState const type.
func PossiblePrefixValidationStateValues() []PrefixValidationState {
	return []PrefixValidationState{
		PrefixValidationStateFailed,
		PrefixValidationStateInvalid,
		PrefixValidationStateNone,
		PrefixValidationStatePending,
		PrefixValidationStateUnknown,
		PrefixValidationStateVerified,
	}
}

// ToPtr returns a *PrefixValidationState pointing to the current value.
func (c PrefixValidationState) ToPtr() *PrefixValidationState {
	return &c
}

// ProvisioningState - The provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// SessionAddressProvider - The field indicating if Microsoft provides session ip addresses.
type SessionAddressProvider string

const (
	SessionAddressProviderMicrosoft SessionAddressProvider = "Microsoft"
	SessionAddressProviderPeer      SessionAddressProvider = "Peer"
)

// PossibleSessionAddressProviderValues returns the possible values for the SessionAddressProvider const type.
func PossibleSessionAddressProviderValues() []SessionAddressProvider {
	return []SessionAddressProvider{
		SessionAddressProviderMicrosoft,
		SessionAddressProviderPeer,
	}
}

// ToPtr returns a *SessionAddressProvider pointing to the current value.
func (c SessionAddressProvider) ToPtr() *SessionAddressProvider {
	return &c
}

// SessionStateV4 - The state of the IPv4 session.
type SessionStateV4 string

const (
	SessionStateV4Active        SessionStateV4 = "Active"
	SessionStateV4Connect       SessionStateV4 = "Connect"
	SessionStateV4Established   SessionStateV4 = "Established"
	SessionStateV4Idle          SessionStateV4 = "Idle"
	SessionStateV4None          SessionStateV4 = "None"
	SessionStateV4OpenConfirm   SessionStateV4 = "OpenConfirm"
	SessionStateV4OpenReceived  SessionStateV4 = "OpenReceived"
	SessionStateV4OpenSent      SessionStateV4 = "OpenSent"
	SessionStateV4PendingAdd    SessionStateV4 = "PendingAdd"
	SessionStateV4PendingRemove SessionStateV4 = "PendingRemove"
	SessionStateV4PendingUpdate SessionStateV4 = "PendingUpdate"
)

// PossibleSessionStateV4Values returns the possible values for the SessionStateV4 const type.
func PossibleSessionStateV4Values() []SessionStateV4 {
	return []SessionStateV4{
		SessionStateV4Active,
		SessionStateV4Connect,
		SessionStateV4Established,
		SessionStateV4Idle,
		SessionStateV4None,
		SessionStateV4OpenConfirm,
		SessionStateV4OpenReceived,
		SessionStateV4OpenSent,
		SessionStateV4PendingAdd,
		SessionStateV4PendingRemove,
		SessionStateV4PendingUpdate,
	}
}

// ToPtr returns a *SessionStateV4 pointing to the current value.
func (c SessionStateV4) ToPtr() *SessionStateV4 {
	return &c
}

// SessionStateV6 - The state of the IPv6 session.
type SessionStateV6 string

const (
	SessionStateV6Active        SessionStateV6 = "Active"
	SessionStateV6Connect       SessionStateV6 = "Connect"
	SessionStateV6Established   SessionStateV6 = "Established"
	SessionStateV6Idle          SessionStateV6 = "Idle"
	SessionStateV6None          SessionStateV6 = "None"
	SessionStateV6OpenConfirm   SessionStateV6 = "OpenConfirm"
	SessionStateV6OpenReceived  SessionStateV6 = "OpenReceived"
	SessionStateV6OpenSent      SessionStateV6 = "OpenSent"
	SessionStateV6PendingAdd    SessionStateV6 = "PendingAdd"
	SessionStateV6PendingRemove SessionStateV6 = "PendingRemove"
	SessionStateV6PendingUpdate SessionStateV6 = "PendingUpdate"
)

// PossibleSessionStateV6Values returns the possible values for the SessionStateV6 const type.
func PossibleSessionStateV6Values() []SessionStateV6 {
	return []SessionStateV6{
		SessionStateV6Active,
		SessionStateV6Connect,
		SessionStateV6Established,
		SessionStateV6Idle,
		SessionStateV6None,
		SessionStateV6OpenConfirm,
		SessionStateV6OpenReceived,
		SessionStateV6OpenSent,
		SessionStateV6PendingAdd,
		SessionStateV6PendingRemove,
		SessionStateV6PendingUpdate,
	}
}

// ToPtr returns a *SessionStateV6 pointing to the current value.
func (c SessionStateV6) ToPtr() *SessionStateV6 {
	return &c
}

// Size - The size of the peering SKU.
type Size string

const (
	SizeFree      Size = "Free"
	SizeMetered   Size = "Metered"
	SizeUnlimited Size = "Unlimited"
)

// PossibleSizeValues returns the possible values for the Size const type.
func PossibleSizeValues() []Size {
	return []Size{
		SizeFree,
		SizeMetered,
		SizeUnlimited,
	}
}

// ToPtr returns a *Size pointing to the current value.
func (c Size) ToPtr() *Size {
	return &c
}

// Tier - The tier of the peering SKU.
type Tier string

const (
	TierBasic   Tier = "Basic"
	TierPremium Tier = "Premium"
)

// PossibleTierValues returns the possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{
		TierBasic,
		TierPremium,
	}
}

// ToPtr returns a *Tier pointing to the current value.
func (c Tier) ToPtr() *Tier {
	return &c
}

// ValidationState - The validation state of the ASN associated with the peer.
type ValidationState string

const (
	ValidationStateApproved ValidationState = "Approved"
	ValidationStateFailed   ValidationState = "Failed"
	ValidationStateNone     ValidationState = "None"
	ValidationStatePending  ValidationState = "Pending"
)

// PossibleValidationStateValues returns the possible values for the ValidationState const type.
func PossibleValidationStateValues() []ValidationState {
	return []ValidationState{
		ValidationStateApproved,
		ValidationStateFailed,
		ValidationStateNone,
		ValidationStatePending,
	}
}

// ToPtr returns a *ValidationState pointing to the current value.
func (c ValidationState) ToPtr() *ValidationState {
	return &c
}
