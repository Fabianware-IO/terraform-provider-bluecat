package bluecat

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type ProteusAPI struct {
	client *SOAPClient
}

// Added!
type Long struct{}
type APIEntity struct{}
type APIEntityArray struct{}
type Int struct{}
type APIDeploymentRole struct{}
type APIDeploymentRoleArray struct{}
type APIDeploymentOption struct{}
type APIDeploymentOptionArray struct{}
type LongArray struct{}
type Boolean struct{}
type String struct{}
type APIAccessRight struct{}
type APIData struct{}
type StringArray struct{}
type Base64Binary struct{}
type APIAccessRightArray struct{}
type APIUserDefinedFieldArray struct{}
type ResponsePolicySearchResultArray struct{}

func NewProteusAPI(url string, tls bool, auth *BasicAuth) *ProteusAPI {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &ProteusAPI{
		client: client,
	}
}

func (service *ProteusAPI) GetEntities(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddHostRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddAliasRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddMXRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddSRVRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTXTRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddHINFORecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddNAPTRRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddGenericRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddStartOfAuthority(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddACL(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddUserGroup(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTagGroup(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTFTPGroup(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTFTPFolder(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTFTPFile(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTFTPDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) CreateXHAPair(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) EditXHAPair(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeployServerServices(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeployServerConfig(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) AddAdditionalIPAddresses(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) RemoveAdditionalIPAddresses(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetAdditionalIPAddresses(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AssignIP4Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AssignNextAvailableIP4Address(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetNextIP4Address(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetMaxAllowedRange(request *Long) (*StringArray, error) {
	response := new(StringArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetNextAvailableIPRanges(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) TagEntity(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UntagEntity(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) QuickDeploy(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDependentRecords(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetNetworkLinkedProperties(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) LinkEntities(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UnlinkEntities(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetServerForRole(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetMACAddressesInPool(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetMACAddress(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntityByName(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntitiesByName(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntitiesByNameUsingOptions(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetAliasesByHint(request *Int) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetHostRecordsByHint(request *Int) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetZonesByHint(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetIP4NetworksByHint(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetIP6ObjectsByHint(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateWithOptions(request *APIEntity) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetAccessRightsForEntity(request *Long) (*APIAccessRightArray, error) {
	response := new(APIAccessRightArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetAccessRightsForUser(request *Long) (*APIAccessRightArray, error) {
	response := new(APIAccessRightArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) DeleteAccessRight(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) UpdateAccessRight(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDeploymentRoles(request *Long) (*APIDeploymentRoleArray, error) {
	response := new(APIDeploymentRoleArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetServerDeploymentRoles(request *Long) (*APIDeploymentRoleArray, error) {
	response := new(APIDeploymentRoleArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDNSDeploymentRole(request *Long) (*APIDeploymentRole, error) {
	response := new(APIDeploymentRole)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDNSDeploymentRoleForView(request *Long) (*APIDeploymentRole, error) {
	response := new(APIDeploymentRole)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDHCPDeploymentRole(request *Long) (*APIDeploymentRole, error) {
	response := new(APIDeploymentRole)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) AddDNSDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) AddDHCPDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDNSDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDNSDeploymentRoleForView(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDHCPDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateDNSDeploymentRole(request *APIDeploymentRole) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateDHCPDeploymentRole(request *APIDeploymentRole) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) MoveDeploymentRoles(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDeploymentOptions(request *Long) (*APIDeploymentOptionArray, error) {
	response := new(APIDeploymentOptionArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDNSDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDHCPClientDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDHCP6ClientDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDHCPServiceDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDHCP6ServiceDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDNSDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCPClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCP6ClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCPServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCP6ServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDNSDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) DeleteDHCPClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDHCP6ClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDHCPServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDHCP6ServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) UpdateDNSDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) UpdateDHCPClientDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateDHCP6ClientDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateDHCPServiceDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateDHCP6ServiceDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddMACAddress(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) AssociateMACAddressWithPool(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DenyMACAddress(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) IsAddressAllocated(request *Long) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) MergeSelectedBlocksOrNetworks(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddParentBlock(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddParentBlockWithProperties(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) MergeBlocksWithParent(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ResizeRange(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) MoveIP4Object(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) MoveIPObject(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) MoveResourceRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) UpdateBulkUdf(request *Base64Binary) (*Base64Binary, error) {
	response := new(Base64Binary)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddVendorProfile(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddVendorOptionDefinition(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCPVendorDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDHCPVendorDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateDHCPVendorDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDHCPVendorDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP6BlockByPrefix(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP6BlockByMACAddress(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntityByPrefix(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP6NetworkByPrefix(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP6Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AssignIP6Address(request *Long) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) ReassignIP6Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) ClearIP6Address(request *Long) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) SearchByCategory(request *String) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) SearchByObjectTypes(request *String) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) CustomSearch(request *StringArray) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) PurgeHistoryNow(request *String) (*Int, error) {
	response := new(Int)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddZoneTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) AssignOrUpdateTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ReapplyTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddCustomOptionDefinition(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDeviceType(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDeviceSubtype(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDevice(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetIPRangedByIP(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddBulkHostRecord(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) MigrateFile(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) IsMigrationRunning(request *String) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4ReconciliationPolicy(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ChangeStateIP4Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetUserDefinedFields(request *String) (*APIUserDefinedFieldArray, error) {
	response := new(APIUserDefinedFieldArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4IPGroupByRange(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4IPGroupBySize(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) SplitIP4Network(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) SplitIP6Range(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDeviceInstance(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteDeviceInstance(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetKSK(request *Long) (*StringArray, error) {
	response := new(StringArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddResponsePolicy(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UploadResponsePolicyItems(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ConfigureReplication(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDevices(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDevice(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDeviceInterfaces(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDeviceNetworks(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDeviceHosts(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDeviceVlans(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDeviceArpEntries(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetDiscoveredDeviceMacAddressEntries(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) StartProbe(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetProbeData(request *String) (*APIData, error) {
	response := new(APIData)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ShareNetwork(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UnshareNetwork(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetSharedNetworks(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetLocationByCode(request *String) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetAllUsedLocations() (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) SearchResponsePolicyItems(request *String) (*ResponsePolicySearchResultArray, error) {
	response := new(ResponsePolicySearchResultArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetConfigurationSetting(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateConfigurationSetting(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddRawDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) UpdateRawDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) Update(request *APIEntity) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetParent(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) Delete(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) Login(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) Logout() (*Long, error) {
	response := new(Long)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddUser(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddAccessRight(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetIP6Address(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddEntity(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddResourceRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ReplaceServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetServerDeploymentStatus(request *Long) (*Int, error) {
	response := new(Int)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetLinkedEntities(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddView(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCPMatchClass(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddEnumNumber(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetSystemInfo() (*String, error) {
	response := new(String)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeployServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) ImportServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetProbeStatus(request *String) (*Int, error) {
	response := new(Int)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetAccessRight(request *Long) (*APIAccessRight, error) {
	response := new(APIAccessRight)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetIP4Address(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4Network(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4NetworkTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetNextAvailableIP4Network(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetNextAvailableIP4Address(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetNextAvailableIPRange(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCPSubClass(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddTag(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) DeleteWithOptions(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4BlockByCIDR(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) FailoverXHA(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddZone(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddExternalHostRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntityById(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCP4Range(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCP4RangeBySize(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddDHCP6Range(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntityByCIDR(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddIP4BlockByRange(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) AddEnumZone(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) LoginWithOptions(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */
func (service *ProteusAPI) BreakXHAPair(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ProteusAPI) GetEntityByRange(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
