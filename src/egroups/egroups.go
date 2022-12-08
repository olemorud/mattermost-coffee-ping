package egroups

import (
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type PrivilegeTypeCode string

const (
	PrivilegeTypeCodeSeeMembers PrivilegeTypeCode = "SeeMembers"

	PrivilegeTypeCodeAdmin PrivilegeTypeCode = "Admin"
)

type EgroupTypeCode string

const (
	EgroupTypeCodeStaticEgroup EgroupTypeCode = "StaticEgroup"

	EgroupTypeCodeDynamicEgroup EgroupTypeCode = "DynamicEgroup"
)

type StatusCode string

const (
	StatusCodeCreated StatusCode = "Created"

	StatusCodeActive StatusCode = "Active"

	StatusCodeBlocked StatusCode = "Blocked"

	StatusCodeDeleted StatusCode = "Deleted"
)

type BlockingReasonCode string

const (
	BlockingReasonCodeManual BlockingReasonCode = "Manual"

	BlockingReasonCodeExpired BlockingReasonCode = "Expired"
)

type UsageCode string

const (
	UsageCodeEgroupsOnly UsageCode = "EgroupsOnly"

	UsageCodeSecurityMailing UsageCode = "SecurityMailing"
)

type AdministratorTypeCode string

const (
	AdministratorTypeCodeStaticEgroup AdministratorTypeCode = "StaticEgroup"

	AdministratorTypeCodeDynamicEgroup AdministratorTypeCode = "DynamicEgroup"
)

type PrivacyType string

const (
	PrivacyTypeMembers PrivacyType = "Members"

	PrivacyTypeOpen PrivacyType = "Open"

	PrivacyTypeAdministrators PrivacyType = "Administrators"

	PrivacyTypeUsers PrivacyType = "Users"
)

type SelfsubscriptionType string

const (
	SelfsubscriptionTypeClosed SelfsubscriptionType = "Closed"

	SelfsubscriptionTypeOpen SelfsubscriptionType = "Open"

	SelfsubscriptionTypeMembers SelfsubscriptionType = "Members"

	SelfsubscriptionTypeUsers SelfsubscriptionType = "Users"

	SelfsubscriptionTypeOpenWithAdminApproval SelfsubscriptionType = "OpenWithAdminApproval"

	SelfsubscriptionTypeUsersWithAdminApproval SelfsubscriptionType = "UsersWithAdminApproval"
)

type MemberTypeCode string

const (
	MemberTypeCodeExternal MemberTypeCode = "External"

	MemberTypeCodePerson MemberTypeCode = "Person"

	MemberTypeCodeServiceProvider MemberTypeCode = "ServiceProvider"

	MemberTypeCodeStaticEgroup MemberTypeCode = "StaticEgroup"

	MemberTypeCodeDynamicEgroup MemberTypeCode = "DynamicEgroup"

	MemberTypeCodeAccount MemberTypeCode = "Account"
)

type PostingRestrictionType string

const (
	PostingRestrictionTypeEveryone PostingRestrictionType = "Everyone"

	PostingRestrictionTypeCernUsers PostingRestrictionType = "CernUsers"

	PostingRestrictionTypeOwnerAdminsAndOthers PostingRestrictionType = "OwnerAdminsAndOthers"
)

type WhoReceivesDeliveryErrorsType string

const (
	WhoReceivesDeliveryErrorsTypeGroupOwner WhoReceivesDeliveryErrorsType = "GroupOwner"

	WhoReceivesDeliveryErrorsTypeSender WhoReceivesDeliveryErrorsType = "Sender"

	WhoReceivesDeliveryErrorsTypeNone WhoReceivesDeliveryErrorsType = "None"
)

type ArchivePropertiesType string

const (
	ArchivePropertiesTypeDoesNotExist ArchivePropertiesType = "DoesNotExist"

	ArchivePropertiesTypeActive ArchivePropertiesType = "Active"

	ArchivePropertiesTypeNotActive ArchivePropertiesType = "NotActive"
)

type ErrorCode string

const (
	ErrorCodeNOT_VALID_USER ErrorCode = "NOT_VALID_USER"

	ErrorCodeNOT_FOUND ErrorCode = "NOT_FOUND"

	ErrorCodeEGROUP_NAME_BAD_FORMATED ErrorCode = "EGROUP_NAME_BAD_FORMATED"

	ErrorCodeINSUFFICIENT_PRIVILEGES ErrorCode = "INSUFFICIENT_PRIVILEGES"

	ErrorCodePERSON_ID_NOT_FOUND ErrorCode = "PERSON_ID_NOT_FOUND"

	ErrorCodeEXPIRE_DATE_BAD_FORMATTED ErrorCode = "EXPIRE_DATE_BAD_FORMATTED"

	ErrorCodeNAME_AND_ID_NOT_CORRESPONDS ErrorCode = "NAME_AND_ID_NOT_CORRESPONDS"

	ErrorCodeEGROUP_RECENTLY_DELETED ErrorCode = "EGROUP_RECENTLY_DELETED"

	ErrorCodeEMAIL_PROP_ARCHIVE_NOT_VALID ErrorCode = "EMAIL_PROP_ARCHIVE_NOT_VALID"

	ErrorCodeEMAIL_PROP_MAIL_SIZE_NOT_VALID ErrorCode = "EMAIL_PROP_MAIL_SIZE_NOT_VALID"

	ErrorCodeEXPIRATION_DATE_MANDATORY ErrorCode = "EXPIRATION_DATE_MANDATORY"

	ErrorCodeEXPIRATION_DATE_NOT_VALID ErrorCode = "EXPIRATION_DATE_NOT_VALID"

	ErrorCodeINTERNAL_DB_ERROR ErrorCode = "INTERNAL_DB_ERROR"

	ErrorCodeDATABASE_CONFIGURATION_NOT_FOUND ErrorCode = "DATABASE_CONFIGURATION_NOT_FOUND"

	ErrorCodeEGROUP_MEMBER_OF_ANOTHER ErrorCode = "EGROUP_MEMBER_OF_ANOTHER"

	ErrorCodeNOT_MAILING_SEGURITY_USAGE ErrorCode = "NOT_MAILING_SEGURITY_USAGE"

	ErrorCodeIS_ALREADY_MEMBER ErrorCode = "IS_ALREADY_MEMBER"

	ErrorCodeTOPIC_ALREADY_EXISTS ErrorCode = "TOPIC_ALREADY_EXISTS"

	ErrorCodeALIAS_MUST_HAVE_HYPHEN ErrorCode = "ALIAS_MUST_HAVE_HYPHEN"

	ErrorCodeALIAS_ALREADY_EXISTS ErrorCode = "ALIAS_ALREADY_EXISTS"

	ErrorCodeNAME_ALREADY_RESERVED ErrorCode = "NAME_ALREADY_RESERVED"

	ErrorCodeNAME_TOO_LONG ErrorCode = "NAME_TOO_LONG"

	ErrorCodeEGROUP_ALREADY_EXISTS ErrorCode = "EGROUP_ALREADY_EXISTS"

	ErrorCodeMEMBER_NOT_FOUND ErrorCode = "MEMBER_NOT_FOUND"

	ErrorCodeALIAS_NOT_FOUND ErrorCode = "ALIAS_NOT_FOUND"

	ErrorCodeSELF_EGROUP_ALREADY_EXISTS ErrorCode = "SELF_EGROUP_ALREADY_EXISTS"

	ErrorCodeALREADY_ACTIVE ErrorCode = "ALREADY_ACTIVE"

	ErrorCodeIS_BLOCKED ErrorCode = "IS_BLOCKED"

	ErrorCodeMUST_BE_MODERATOR ErrorCode = "MUST_BE_MODERATOR"

	ErrorCodeSTATUS_CHANGE_NOT_ALLOWED ErrorCode = "STATUS_CHANGE_NOT_ALLOWED"

	ErrorCodeEXPIRATION_DATE_CANT_BE_PROLONGUED ErrorCode = "EXPIRATION_DATE_CANT_BE_PROLONGUED"

	ErrorCodeBLOCKING_REASON_UNDEFINED ErrorCode = "BLOCKING_REASON_UNDEFINED"

	ErrorCodeUSAGE_TYPE_NOT_VALID ErrorCode = "USAGE_TYPE_NOT_VALID"

	ErrorCodeNOT_LOGGED ErrorCode = "NOT_LOGGED"

	ErrorCodeALREADY_DELETED ErrorCode = "ALREADY_DELETED"

	ErrorCodeNAME_MUST_HAVE_HYPHEN ErrorCode = "NAME_MUST_HAVE_HYPHEN"

	ErrorCodeIS_ALREADY_ALLOWED_TO_POST ErrorCode = "IS_ALREADY_ALLOWED_TO_POST"

	ErrorCodeHAS_ALREADY_PRIVILEGE ErrorCode = "HAS_ALREADY_PRIVILEGE"

	ErrorCodeOWNER_ID_NOT_FOUND ErrorCode = "OWNER_ID_NOT_FOUND"

	ErrorCodeWARNING ErrorCode = "WARNING"

	ErrorCodeUNEXPECTED_ERROR ErrorCode = "UNEXPECTED_ERROR"
)

type UpdateEmailPropertiesRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema UpdateEmailPropertiesRequest"`

	EgroupName string `xml:"egroupName,omitempty"`

	EmailProperties *EmailPropertiesType `xml:"emailProperties,omitempty"`
}

type UpdateEmailPropertiesResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema UpdateEmailPropertiesResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type AddEgroupEmailMembersRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema AddEgroupEmailMembersRequest"`

	EgroupName string `xml:"egroupName,omitempty"`

	OverwriteMembers bool `xml:"overwriteMembers,omitempty"`

	Emails []string `xml:"emails,omitempty"`
}

type AddEgroupEmailMembersResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema AddEgroupEmailMembersResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type AddEgroupMembersRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema AddEgroupMembersRequest"`

	EgroupName string `xml:"egroupName,omitempty"`

	OverwriteMembers bool `xml:"overwriteMembers,omitempty"`

	Members []*MemberType `xml:"members,omitempty"`
}

type AddEgroupMembersResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema AddEgroupMembersResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type RemoveEgroupEmailMembersRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema RemoveEgroupEmailMembersRequest"`

	EgroupName string `xml:"egroupName,omitempty"`

	Emails []string `xml:"emails,omitempty"`
}

type RemoveEgroupEmailMembersResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema RemoveEgroupEmailMembersResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type RemoveEgroupMembersRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema RemoveEgroupMembersRequest"`

	EgroupName string `xml:"egroupName,omitempty"`

	Members []*MemberType `xml:"members,omitempty"`
}

type RemoveEgroupMembersResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema RemoveEgroupMembersResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type GetEgroupsUserOwnOrManageRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema GetEgroupsUserOwnOrManageRequest"`

	UserName string `xml:"userName,omitempty"`
}

type GetEgroupsUserOwnOrManageResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema GetEgroupsUserOwnOrManageResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`

	Results []*EgroupType `xml:"results,omitempty"`
}

type ChangeExternalEmailAddressRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema ChangeExternalEmailAddressRequest"`

	OldEmail string `xml:"oldEmail,omitempty"`

	NewEmail string `xml:"newEmail,omitempty"`
}

type ChangeExternalEmailAddressResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema ChangeExternalEmailAddressResponse"`

	Result int32 `xml:"result,omitempty"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`
}

type FindEgroupByNameRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema FindEgroupByNameRequest"`

	Name string `xml:"name,omitempty"`
}

type FindEgroupByNameResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema FindEgroupByNameResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`

	Result *EgroupType `xml:"result,omitempty"`
}

type FindEgroupByIdRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema FindEgroupByIdRequest"`

	Id int64 `xml:"id,omitempty"`
}

type FindEgroupByIdResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema FindEgroupByIdResponse"`

	Result *EgroupType `xml:"result,omitempty"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type DeleteEgroupRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema DeleteEgroupRequest"`

	EgroupName string `xml:"egroupName,omitempty"`
}

type DeleteEgroupResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema DeleteEgroupResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`
}

type SynchronizeEgroupRequest struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema SynchronizeEgroupRequest"`

	Egroup *EgroupType `xml:"egroup,omitempty"`
}

type SynchronizeEgroupResponse struct {
	XMLName xml.Name `xml:"https://foundservices.cern.ch/ws/egroups/v1/schema/EgroupsServicesSchema SynchronizeEgroupResponse"`

	TransactionId string `xml:"transactionId,omitempty"`

	Error *ErrorType `xml:"error,omitempty"`

	Warnings []*ErrorType `xml:"warnings,omitempty"`
}

type EgroupType struct {
	Name string `xml:"Name,omitempty"`

	Aliases []string `xml:"Aliases,omitempty"`

	ID int64 `xml:"ID,omitempty"`

	Type *EgroupTypeCode `xml:"Type,omitempty"`

	Status *StatusCode `xml:"Status,omitempty"`

	BlockingReason *BlockingReasonCode `xml:"BlockingReason,omitempty"`

	Usage *UsageCode `xml:"Usage,omitempty"`

	Topic string `xml:"Topic,omitempty"`

	Description string `xml:"Description,omitempty"`

	Comments string `xml:"Comments,omitempty"`

	ExpiryDate string `xml:"ExpiryDate,omitempty"`

	Owner *UserType `xml:"Owner,omitempty"`

	AdministratorEgroup string `xml:"AdministratorEgroup,omitempty"`

	Privacy *PrivacyType `xml:"Privacy,omitempty"`

	Selfsubscription *SelfsubscriptionType `xml:"Selfsubscription,omitempty"`

	SelfsubscriptionEgroups []*SelfsubscriptionEgroupType `xml:"SelfsubscriptionEgroups,omitempty"`

	EgroupsWithPrivileges []*EgroupWithPrivilegeType `xml:"EgroupsWithPrivileges,omitempty"`

	Members []*MemberType `xml:"Members,omitempty"`

	EmailProperties *EmailPropertiesType `xml:"EmailProperties,omitempty"`
}

type ErrorType struct {
	Code *ErrorCode `xml:"Code,omitempty"`

	Field string `xml:"Field,omitempty"`

	Message string `xml:"Message,omitempty"`
}

type EgroupWithPrivilegeType struct {
	Name string `xml:"Name,omitempty"`

	Privilege *PrivilegeTypeCode `xml:"Privilege,omitempty"`
}

type UserType struct {
	PersonId int64 `xml:"PersonId,omitempty"`

	Name string `xml:"Name,omitempty"`

	ComputingRulesSigned bool `xml:"ComputingRulesSigned,omitempty"`

	Pem string `xml:"Pem,omitempty"`

	PrimaryGem *GemType `xml:"PrimaryGem,omitempty"`

	CernUnit string `xml:"CernUnit,omitempty"`

	CernDepartment string `xml:"CernDepartment,omitempty"`

	CernGroup string `xml:"CernGroup,omitempty"`

	Telephone1 string `xml:"Telephone1,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Building string `xml:"Building,omitempty"`

	Floor string `xml:"Floor,omitempty"`

	Room string `xml:"Room,omitempty"`

	Mailbox string `xml:"Mailbox,omitempty"`
}

type GemType struct {
	EmailAddress string `xml:"EmailAddress,omitempty"`

	DisplayInPhonebook bool `xml:"DisplayInPhonebook,omitempty"`
}

type SelfsubscriptionEgroupType struct {
	ID int64 `xml:"ID,omitempty"`

	ApprovalNeeded bool `xml:"ApprovalNeeded,omitempty"`
}

type MemberType struct {
	ID int64 `xml:"ID,omitempty"`

	Type *MemberTypeCode `xml:"Type,omitempty"`

	Name string `xml:"Name,omitempty"`

	Email string `xml:"Email,omitempty"`

	PrimaryAccount string `xml:"PrimaryAccount,omitempty"`

	Comments string `xml:"Comments,omitempty"`
}

type EmailPropertiesType struct {
	MailPostingRestrictions *MailPostingRestrictionType `xml:"MailPostingRestrictions,omitempty"`

	SenderAuthenticationEnabled bool `xml:"SenderAuthenticationEnabled,omitempty"`

	WhoReceivesDeliveryErrors *WhoReceivesDeliveryErrorsType `xml:"WhoReceivesDeliveryErrors,omitempty"`

	//MaxMailSize *PositiveInteger `xml:"MaxMailSize,omitempty"`

	ArchiveProperties *ArchivePropertiesType `xml:"ArchiveProperties,omitempty"`
}

type MailPostingRestrictionType struct {
	PostingRestrictions *PostingRestrictionType `xml:"PostingRestrictions,omitempty"`

	OtherRecipientsAllowedToPost []*MemberType `xml:"OtherRecipientsAllowedToPost,omitempty"`
}

type EgroupsService interface {
	UpdateEmailProperties(request *UpdateEmailPropertiesRequest) (*UpdateEmailPropertiesResponse, error)

	AddEgroupEmailMembers(request *AddEgroupEmailMembersRequest) (*AddEgroupEmailMembersResponse, error)

	RemoveEgroupEmailMembers(request *RemoveEgroupEmailMembersRequest) (*RemoveEgroupEmailMembersResponse, error)

	AddEgroupMembers(request *AddEgroupMembersRequest) (*AddEgroupMembersResponse, error)

	RemoveEgroupMembers(request *RemoveEgroupMembersRequest) (*RemoveEgroupMembersResponse, error)

	GetEgroupsUserOwnOrManage(request *GetEgroupsUserOwnOrManageRequest) (*GetEgroupsUserOwnOrManageResponse, error)

	FindEgroupByName(request *FindEgroupByNameRequest) (*FindEgroupByNameResponse, error)

	FindEgroupById(request *FindEgroupByIdRequest) (*FindEgroupByIdResponse, error)

	ChangeExternalEmailAddress(request *ChangeExternalEmailAddressRequest) (*ChangeExternalEmailAddressResponse, error)

	SynchronizeEgroup(request *SynchronizeEgroupRequest) (*SynchronizeEgroupResponse, error)

	DeleteEgroup(request *DeleteEgroupRequest) (*DeleteEgroupResponse, error)
}

type egroupsService struct {
	client *soap.Client
}

func NewEgroupsService(client *soap.Client) EgroupsService {
	return &egroupsService{
		client: client,
	}
}

func (service *egroupsService) UpdateEmailProperties(request *UpdateEmailPropertiesRequest) (*UpdateEmailPropertiesResponse, error) {
	response := new(UpdateEmailPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) AddEgroupEmailMembers(request *AddEgroupEmailMembersRequest) (*AddEgroupEmailMembersResponse, error) {
	response := new(AddEgroupEmailMembersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) RemoveEgroupEmailMembers(request *RemoveEgroupEmailMembersRequest) (*RemoveEgroupEmailMembersResponse, error) {
	response := new(RemoveEgroupEmailMembersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) AddEgroupMembers(request *AddEgroupMembersRequest) (*AddEgroupMembersResponse, error) {
	response := new(AddEgroupMembersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) RemoveEgroupMembers(request *RemoveEgroupMembersRequest) (*RemoveEgroupMembersResponse, error) {
	response := new(RemoveEgroupMembersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) GetEgroupsUserOwnOrManage(request *GetEgroupsUserOwnOrManageRequest) (*GetEgroupsUserOwnOrManageResponse, error) {
	response := new(GetEgroupsUserOwnOrManageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) FindEgroupByName(request *FindEgroupByNameRequest) (*FindEgroupByNameResponse, error) {
	response := new(FindEgroupByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) FindEgroupById(request *FindEgroupByIdRequest) (*FindEgroupByIdResponse, error) {
	response := new(FindEgroupByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) ChangeExternalEmailAddress(request *ChangeExternalEmailAddressRequest) (*ChangeExternalEmailAddressResponse, error) {
	response := new(ChangeExternalEmailAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) SynchronizeEgroup(request *SynchronizeEgroupRequest) (*SynchronizeEgroupResponse, error) {
	response := new(SynchronizeEgroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *egroupsService) DeleteEgroup(request *DeleteEgroupRequest) (*DeleteEgroupResponse, error) {
	response := new(DeleteEgroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
