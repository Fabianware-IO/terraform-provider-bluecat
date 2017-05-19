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

// ProteusAPI :: Does Something
type ProteusAPI struct {
	client *SOAPClient
}

// Added!

// Long :: Does Something
type Long struct{}

// APIEntity :: Does Something
type APIEntity struct{}

// APIEntityArray :: Does Something
type APIEntityArray struct{}

// Int :: does something
type Int struct{}

// APIDeploymentRole :: does something
type APIDeploymentRole struct{}

// APIDeploymentRoleArray :: does something
type APIDeploymentRoleArray struct{}

// APIDeploymentOption :: does something
type APIDeploymentOption struct{}

// APIDeploymentOptionArray :: does something
type APIDeploymentOptionArray struct{}

// LongArray :: Does Something
type LongArray struct{}

// Boolean :: does something
type Boolean struct{}

// String :: does something
type String struct{}

// APIAccessRight :: does something
type APIAccessRight struct{}

// APIData :: does something
type APIData struct{}

// StringArray :: does something
type StringArray struct{}

// Base64Binary :: does something
type Base64Binary struct{}

//APIAccessRightArray :: does something
type APIAccessRightArray struct{}

// APIUserDefinedFieldArray :: does something
type APIUserDefinedFieldArray struct{}

// ResponsePolicySearchResultArray :: does something
type ResponsePolicySearchResultArray struct{}

// NewProteusAPI :: returns a new proteus api client
func NewProteusAPI(url string, tls bool, auth *BasicAuth) *ProteusAPI {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &ProteusAPI{
		client: client,
	}
}

// GetEntities :: returns a list of entitities
func (service *ProteusAPI) GetEntities(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddHostRecord :: adds a host A record
func (service *ProteusAPI) AddHostRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddAliasRecord :: adds an CNAME alias record
func (service *ProteusAPI) AddAliasRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddMXRecord :: Adds a mail MX record
func (service *ProteusAPI) AddMXRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddSRVRecord :: adds an SRV record
func (service *ProteusAPI) AddSRVRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTXTRecord :: adds a TXT record
func (service *ProteusAPI) AddTXTRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddHINFORecord :: adds an HINFO record
func (service *ProteusAPI) AddHINFORecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddNAPTRRecord :: adds a NAPTR record
func (service *ProteusAPI) AddNAPTRRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddGenericRecord :: adds a generic record
func (service *ProteusAPI) AddGenericRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddStartOfAuthority :: adds an SOA
func (service *ProteusAPI) AddStartOfAuthority(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddACL :: adds an ACL
func (service *ProteusAPI) AddACL(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddUserGroup :: Adds a user group
func (service *ProteusAPI) AddUserGroup(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTagGroup :: Adds a tag group
func (service *ProteusAPI) AddTagGroup(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTFTPGroup :: adds a TFTP Group
func (service *ProteusAPI) AddTFTPGroup(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTFTPFolder :: adds a TFTP folder
func (service *ProteusAPI) AddTFTPFolder(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTFTPFile :: adds a tftp file
func (service *ProteusAPI) AddTFTPFile(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTFTPDeploymentRole :: adds a tftp deployment role
func (service *ProteusAPI) AddTFTPDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddServer :: adds a server
func (service *ProteusAPI) AddServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateXHAPair :: creates an XHA Pair
func (service *ProteusAPI) CreateXHAPair(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// EditXHAPair :: edits an XHA Pair
func (service *ProteusAPI) EditXHAPair(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeployServerServices :: deploy a server services
func (service *ProteusAPI) DeployServerServices(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeployServerConfig :: deploys a server config
func (service *ProteusAPI) DeployServerConfig(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// AddAdditionalIPAddresses :: add additional ip addresses
func (service *ProteusAPI) AddAdditionalIPAddresses(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// RemoveAdditionalIPAddresses :: removes additional IP Addresses
func (service *ProteusAPI) RemoveAdditionalIPAddresses(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAdditionalIPAddresses :: get additional ip addresses
func (service *ProteusAPI) GetAdditionalIPAddresses(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AssignIP4Address :: assign an ipv4 address
func (service *ProteusAPI) AssignIP4Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AssignNextAvailableIP4Address :: assigns the next available IPv4 address
func (service *ProteusAPI) AssignNextAvailableIP4Address(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetNextIP4Address :: gets the next IPv4 address
func (service *ProteusAPI) GetNextIP4Address(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMaxAllowedRange :: gets the max allowed range of ipv4 addresses
func (service *ProteusAPI) GetMaxAllowedRange(request *Long) (*StringArray, error) {
	response := new(StringArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetNextAvailableIPRanges :: gets the next available IP ranges
func (service *ProteusAPI) GetNextAvailableIPRanges(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// TagEntity :: tags an entity
func (service *ProteusAPI) TagEntity(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UntagEntity :: untags an entity
func (service *ProteusAPI) UntagEntity(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// QuickDeploy :: quick deploys
func (service *ProteusAPI) QuickDeploy(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDependentRecords :: get dependent records
func (service *ProteusAPI) GetDependentRecords(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetNetworkLinkedProperties :: get network linked properties
func (service *ProteusAPI) GetNetworkLinkedProperties(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// LinkEntities :: links entities
func (service *ProteusAPI) LinkEntities(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UnlinkEntities :: unlinks entities
func (service *ProteusAPI) UnlinkEntities(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetServerForRole :: gets a server for a given role
func (service *ProteusAPI) GetServerForRole(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMACAddressesInPool :: gets mac addresses in the pool
func (service *ProteusAPI) GetMACAddressesInPool(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMACAddress :: gets a mac address for an entitiy
func (service *ProteusAPI) GetMACAddress(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntityByName :: gets an entity by name
func (service *ProteusAPI) GetEntityByName(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntitiesByName :: gets multiple entities by an array of entity names
func (service *ProteusAPI) GetEntitiesByName(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntitiesByNameUsingOptions :: gets entities by name using options
func (service *ProteusAPI) GetEntitiesByNameUsingOptions(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAliasesByHint :: get aliases by a hint string
func (service *ProteusAPI) GetAliasesByHint(request *Int) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHostRecordsByHint :: get hosts records by hint string
func (service *ProteusAPI) GetHostRecordsByHint(request *Int) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetZonesByHint :: get zones by hint string
func (service *ProteusAPI) GetZonesByHint(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIP4NetworksByHint :: get ipv4 networks by hint string
func (service *ProteusAPI) GetIP4NetworksByHint(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIP6ObjectsByHint :: get IPv6 objects by Hint string
func (service *ProteusAPI) GetIP6ObjectsByHint(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateWithOptions :: updates with options
func (service *ProteusAPI) UpdateWithOptions(request *APIEntity) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccessRightsForEntity :: gets access rights for an entity
func (service *ProteusAPI) GetAccessRightsForEntity(request *Long) (*APIAccessRightArray, error) {
	response := new(APIAccessRightArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccessRightsForUser :: gets access rights for a particular user
func (service *ProteusAPI) GetAccessRightsForUser(request *Long) (*APIAccessRightArray, error) {
	response := new(APIAccessRightArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteAccessRight :: deletes access right
func (service *ProteusAPI) DeleteAccessRight(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateAccessRight :: updates access right
func (service *ProteusAPI) UpdateAccessRight(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDeploymentRoles :: get deployment roles
func (service *ProteusAPI) GetDeploymentRoles(request *Long) (*APIDeploymentRoleArray, error) {
	response := new(APIDeploymentRoleArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetServerDeploymentRoles :: get server deployment roles
func (service *ProteusAPI) GetServerDeploymentRoles(request *Long) (*APIDeploymentRoleArray, error) {
	response := new(APIDeploymentRoleArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDNSDeploymentRole :: get dns deployment role
func (service *ProteusAPI) GetDNSDeploymentRole(request *Long) (*APIDeploymentRole, error) {
	response := new(APIDeploymentRole)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDNSDeploymentRoleForView :: get dns deployment role for view
func (service *ProteusAPI) GetDNSDeploymentRoleForView(request *Long) (*APIDeploymentRole, error) {
	response := new(APIDeploymentRole)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDHCPDeploymentRole :: get dhcp deployment role
func (service *ProteusAPI) GetDHCPDeploymentRole(request *Long) (*APIDeploymentRole, error) {
	response := new(APIDeploymentRole)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// AddDNSDeploymentRole :: add dns deployment role
func (service *ProteusAPI) AddDNSDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// AddDHCPDeploymentRole :: add dhcp deployment role
func (service *ProteusAPI) AddDHCPDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDNSDeploymentRole :: delete dns deployment role
func (service *ProteusAPI) DeleteDNSDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDNSDeploymentRoleForView :: delete DNS deployment role for view
func (service *ProteusAPI) DeleteDNSDeploymentRoleForView(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDHCPDeploymentRole :: delete dhcp deployment role
func (service *ProteusAPI) DeleteDHCPDeploymentRole(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateDNSDeploymentRole :: update DNS deployment role
func (service *ProteusAPI) UpdateDNSDeploymentRole(request *APIDeploymentRole) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateDHCPDeploymentRole :: update dhcp deployment role
func (service *ProteusAPI) UpdateDHCPDeploymentRole(request *APIDeploymentRole) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// MoveDeploymentRoles :: move deployment roles
func (service *ProteusAPI) MoveDeploymentRoles(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDeploymentOptions :: get deployment options
func (service *ProteusAPI) GetDeploymentOptions(request *Long) (*APIDeploymentOptionArray, error) {
	response := new(APIDeploymentOptionArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDNSDeploymentOption :: get dns deployment option
func (service *ProteusAPI) GetDNSDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDHCPClientDeploymentOption :: get dhcp client deployment option
func (service *ProteusAPI) GetDHCPClientDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDHCP6ClientDeploymentOption :: get dhcpv6 client deployment option
func (service *ProteusAPI) GetDHCP6ClientDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDHCPServiceDeploymentOption :: get dhcp service deployment option
func (service *ProteusAPI) GetDHCPServiceDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDHCP6ServiceDeploymentOption :: get dhcp v6 service deployment option
func (service *ProteusAPI) GetDHCP6ServiceDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDNSDeploymentOption :: add dns deployment option
func (service *ProteusAPI) AddDNSDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCPClientDeploymentOption :: add dhcp client deployment option
func (service *ProteusAPI) AddDHCPClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCP6ClientDeploymentOption :: add dhcp v6 client deployment option
func (service *ProteusAPI) AddDHCP6ClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCPServiceDeploymentOption :: add dhcp service deployment option
func (service *ProteusAPI) AddDHCPServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCP6ServiceDeploymentOption :: Add DHCP v6 service deployment option
func (service *ProteusAPI) AddDHCP6ServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDNSDeploymentOption :: delete dns deployment option
func (service *ProteusAPI) DeleteDNSDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteDHCPClientDeploymentOption :: delete dhcp client deployment option
func (service *ProteusAPI) DeleteDHCPClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDHCP6ClientDeploymentOption :: deletes dhcp v6 client deployment option
func (service *ProteusAPI) DeleteDHCP6ClientDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDHCPServiceDeploymentOption :: delete dhcp service deployment option
func (service *ProteusAPI) DeleteDHCPServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDHCP6ServiceDeploymentOption :: delete dhcp v6 service deployment option
func (service *ProteusAPI) DeleteDHCP6ServiceDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateDNSDeploymentOption :: updates dns deployment option
func (service *ProteusAPI) UpdateDNSDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateDHCPClientDeploymentOption :: update dhcp client deployment option
func (service *ProteusAPI) UpdateDHCPClientDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateDHCP6ClientDeploymentOption :: update dhcp v6 client deployment option
func (service *ProteusAPI) UpdateDHCP6ClientDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateDHCPServiceDeploymentOption :: update dhcp service deployment option
func (service *ProteusAPI) UpdateDHCPServiceDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateDHCP6ServiceDeploymentOption :: update dhcp v6 service deployment option
func (service *ProteusAPI) UpdateDHCP6ServiceDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddMACAddress :: add MAC address
func (service *ProteusAPI) AddMACAddress(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// AssociateMACAddressWithPool :: associate MAC address with pool
func (service *ProteusAPI) AssociateMACAddressWithPool(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DenyMACAddress :: deny MAC address
func (service *ProteusAPI) DenyMACAddress(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// IsAddressAllocated :: checks if a specific address is allocated
func (service *ProteusAPI) IsAddressAllocated(request *Long) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// MergeSelectedBlocksOrNetworks :: merges selected blocks or networks together
func (service *ProteusAPI) MergeSelectedBlocksOrNetworks(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddParentBlock :: add parent block
func (service *ProteusAPI) AddParentBlock(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddParentBlockWithProperties :: adds a parent block with properties
func (service *ProteusAPI) AddParentBlockWithProperties(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// MergeBlocksWithParent :: merges a block with a parent
func (service *ProteusAPI) MergeBlocksWithParent(request *LongArray) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ResizeRange :: resizes a range
func (service *ProteusAPI) ResizeRange(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// MoveIP4Object :: moves an IPv4 object
func (service *ProteusAPI) MoveIP4Object(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// MoveIPObject :: moves an IP object
func (service *ProteusAPI) MoveIPObject(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// MoveResourceRecord :: moves a resource record
func (service *ProteusAPI) MoveResourceRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateBulkUdf :: updates a bulk Udf
func (service *ProteusAPI) UpdateBulkUdf(request *Base64Binary) (*Base64Binary, error) {
	response := new(Base64Binary)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddVendorProfile :: adds a vendor profile
func (service *ProteusAPI) AddVendorProfile(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddVendorOptionDefinition :: adds a vendor option definition
func (service *ProteusAPI) AddVendorOptionDefinition(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCPVendorDeploymentOption :: adds a dhcp vendor deployment option
func (service *ProteusAPI) AddDHCPVendorDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDHCPVendorDeploymentOption :: gets a DHCP vendor deployment option
func (service *ProteusAPI) GetDHCPVendorDeploymentOption(request *Long) (*APIDeploymentOption, error) {
	response := new(APIDeploymentOption)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateDHCPVendorDeploymentOption :: updates a dhcp vendor deployment option
func (service *ProteusAPI) UpdateDHCPVendorDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDHCPVendorDeploymentOption :: deletes a dhcp vendor deployment option
func (service *ProteusAPI) DeleteDHCPVendorDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP6BlockByPrefix :: adds ipv6 block by a prefix
func (service *ProteusAPI) AddIP6BlockByPrefix(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP6BlockByMACAddress :: adds an ipv6 block by mac address
func (service *ProteusAPI) AddIP6BlockByMACAddress(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntityByPrefix :: gets an entity by prefix
func (service *ProteusAPI) GetEntityByPrefix(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP6NetworkByPrefix :: adds an ipv6 network by prefix
func (service *ProteusAPI) AddIP6NetworkByPrefix(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP6Address :: add an ipv6 address
func (service *ProteusAPI) AddIP6Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AssignIP6Address :: assigns an ipv6 address
func (service *ProteusAPI) AssignIP6Address(request *Long) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReassignIP6Address :: reassigns an ipv6 address
func (service *ProteusAPI) ReassignIP6Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ClearIP6Address :: clears an ipv6 address
func (service *ProteusAPI) ClearIP6Address(request *Long) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchByCategory :: searches by category
func (service *ProteusAPI) SearchByCategory(request *String) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchByObjectTypes :: searches by object types
func (service *ProteusAPI) SearchByObjectTypes(request *String) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CustomSearch :: custom search
func (service *ProteusAPI) CustomSearch(request *StringArray) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PurgeHistoryNow :: purge history now
func (service *ProteusAPI) PurgeHistoryNow(request *String) (*Int, error) {
	response := new(Int)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddZoneTemplate :: adds a zone template
func (service *ProteusAPI) AddZoneTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// AssignOrUpdateTemplate :: assign or update a template
func (service *ProteusAPI) AssignOrUpdateTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ReapplyTemplate :: reapply template
func (service *ProteusAPI) ReapplyTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddCustomOptionDefinition :: add custom option definition
func (service *ProteusAPI) AddCustomOptionDefinition(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDeviceType :: add a device type
func (service *ProteusAPI) AddDeviceType(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDeviceSubtype :: add a device subtype
func (service *ProteusAPI) AddDeviceSubtype(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDevice :: add a device
func (service *ProteusAPI) AddDevice(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIPRangedByIP :: gets an ip ranged by an ip
func (service *ProteusAPI) GetIPRangedByIP(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddBulkHostRecord :: gets a bulk host record
func (service *ProteusAPI) AddBulkHostRecord(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// MigrateFile :: migrates a file
func (service *ProteusAPI) MigrateFile(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// IsMigrationRunning :: gets whether a migration is running
func (service *ProteusAPI) IsMigrationRunning(request *String) (*Boolean, error) {
	response := new(Boolean)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4ReconciliationPolicy :: add an ipv4 reconciliation policy
func (service *ProteusAPI) AddIP4ReconciliationPolicy(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ChangeStateIP4Address :: changes the state of a ipv4 address
func (service *ProteusAPI) ChangeStateIP4Address(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserDefinedFields :: gets the user defined fields
func (service *ProteusAPI) GetUserDefinedFields(request *String) (*APIUserDefinedFieldArray, error) {
	response := new(APIUserDefinedFieldArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4IPGroupByRange :: add an ipv4 ip group by range
func (service *ProteusAPI) AddIP4IPGroupByRange(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4IPGroupBySize :: add an ipv4 group by size
func (service *ProteusAPI) AddIP4IPGroupBySize(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SplitIP4Network :: splits an ipv4 network
func (service *ProteusAPI) SplitIP4Network(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SplitIP6Range :: splits an ipv6 range
func (service *ProteusAPI) SplitIP6Range(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDeviceInstance :: adds a device instance
func (service *ProteusAPI) AddDeviceInstance(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteDeviceInstance :: deletes a device instance
func (service *ProteusAPI) DeleteDeviceInstance(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetKSK :: gets a KSK
func (service *ProteusAPI) GetKSK(request *Long) (*StringArray, error) {
	response := new(StringArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddResponsePolicy :: adds a response policy
func (service *ProteusAPI) AddResponsePolicy(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UploadResponsePolicyItems :: upload a response policy's items
func (service *ProteusAPI) UploadResponsePolicyItems(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ConfigureReplication :: configures replication
func (service *ProteusAPI) ConfigureReplication(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDevices :: gets discovered devices
func (service *ProteusAPI) GetDiscoveredDevices(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDevice :: gets discovered device
func (service *ProteusAPI) GetDiscoveredDevice(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDeviceInterfaces :: gets discovered device interfaces
func (service *ProteusAPI) GetDiscoveredDeviceInterfaces(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDeviceNetworks :: gets discovered device networks
func (service *ProteusAPI) GetDiscoveredDeviceNetworks(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDeviceHosts :: gets discovered device hosts
func (service *ProteusAPI) GetDiscoveredDeviceHosts(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDeviceVlans :: gets discovered device vlans
func (service *ProteusAPI) GetDiscoveredDeviceVlans(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDeviceArpEntries :: gets discovered device ARP entries
func (service *ProteusAPI) GetDiscoveredDeviceArpEntries(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDiscoveredDeviceMacAddressEntries :: gets discovered device mac address entries
func (service *ProteusAPI) GetDiscoveredDeviceMacAddressEntries(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// StartProbe :: starts a probe
func (service *ProteusAPI) StartProbe(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetProbeData :: gets the probe data
func (service *ProteusAPI) GetProbeData(request *String) (*APIData, error) {
	response := new(APIData)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ShareNetwork :: share a network
func (service *ProteusAPI) ShareNetwork(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UnshareNetwork :: unshares a network
func (service *ProteusAPI) UnshareNetwork(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSharedNetworks :: gets shared networks
func (service *ProteusAPI) GetSharedNetworks(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetLocationByCode :: gets location by code
func (service *ProteusAPI) GetLocationByCode(request *String) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllUsedLocations :: gets all used locations
func (service *ProteusAPI) GetAllUsedLocations() (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchResponsePolicyItems :: searches response policy items
func (service *ProteusAPI) SearchResponsePolicyItems(request *String) (*ResponsePolicySearchResultArray, error) {
	response := new(ResponsePolicySearchResultArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetConfigurationSetting :: gets configuration settings
func (service *ProteusAPI) GetConfigurationSetting(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateConfigurationSetting :: updates the configuration settings
func (service *ProteusAPI) UpdateConfigurationSetting(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddRawDeploymentOption :: adds the raw deployment option
func (service *ProteusAPI) AddRawDeploymentOption(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// UpdateRawDeploymentOption :: updates the raw deployment option
func (service *ProteusAPI) UpdateRawDeploymentOption(request *APIDeploymentOption) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// Update :: update
func (service *ProteusAPI) Update(request *APIEntity) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetParent :: gets the parent
func (service *ProteusAPI) GetParent(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// Delete :: deletes
func (service *ProteusAPI) Delete(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// Login :: logs in
func (service *ProteusAPI) Login(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// Logout :: logs out
func (service *ProteusAPI) Logout() (*Long, error) {
	response := new(Long)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddUser :: adds a user
func (service *ProteusAPI) AddUser(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddAccessRight :: adds an access right
func (service *ProteusAPI) AddAccessRight(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIP6Address :: gets an IPv6 address
func (service *ProteusAPI) GetIP6Address(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddEntity :: adds an entity
func (service *ProteusAPI) AddEntity(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddResourceRecord :: adds a resource record
func (service *ProteusAPI) AddResourceRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ReplaceServer :: replace a server
func (service *ProteusAPI) ReplaceServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetServerDeploymentStatus :: get server deployment status
func (service *ProteusAPI) GetServerDeploymentStatus(request *Long) (*Int, error) {
	response := new(Int)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetLinkedEntities :: get linked entities
func (service *ProteusAPI) GetLinkedEntities(request *Long) (*APIEntityArray, error) {
	response := new(APIEntityArray)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddView :: adds a view
func (service *ProteusAPI) AddView(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCPMatchClass :: adds a dhcp match class
func (service *ProteusAPI) AddDHCPMatchClass(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddEnumNumber :: add an enum number
func (service *ProteusAPI) AddEnumNumber(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSystemInfo :: get system info
func (service *ProteusAPI) GetSystemInfo() (*String, error) {
	response := new(String)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeployServer :: deploy server
func (service *ProteusAPI) DeployServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// ImportServer :: imports a server
func (service *ProteusAPI) ImportServer(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetProbeStatus :: get probe status
func (service *ProteusAPI) GetProbeStatus(request *String) (*Int, error) {
	response := new(Int)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccessRight :: get an access right
func (service *ProteusAPI) GetAccessRight(request *Long) (*APIAccessRight, error) {
	response := new(APIAccessRight)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIP4Address :: get ipv4 address
func (service *ProteusAPI) GetIP4Address(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4Network :: add ipv4 network
func (service *ProteusAPI) AddIP4Network(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4NetworkTemplate :: add an ipv4 network template
func (service *ProteusAPI) AddIP4NetworkTemplate(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetNextAvailableIP4Network ;: gets the next available ipv4 network
func (service *ProteusAPI) GetNextAvailableIP4Network(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetNextAvailableIP4Address :: get net available ipv4 address
func (service *ProteusAPI) GetNextAvailableIP4Address(request *Long) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetNextAvailableIPRange :: get next available ip range
func (service *ProteusAPI) GetNextAvailableIPRange(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCPSubClass :: add dhcp sub class
func (service *ProteusAPI) AddDHCPSubClass(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddTag :: adds a tag
func (service *ProteusAPI) AddTag(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// DeleteWithOptions :: delete with options
func (service *ProteusAPI) DeleteWithOptions(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4BlockByCIDR :: adds an ipv4 block by CIDR
func (service *ProteusAPI) AddIP4BlockByCIDR(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// FailoverXHA :: failover XHA
func (service *ProteusAPI) FailoverXHA(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddZone :: adds a zone
func (service *ProteusAPI) AddZone(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddExternalHostRecord :: adds an external host record
func (service *ProteusAPI) AddExternalHostRecord(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntityByID :: gets an entity by ID
func (service *ProteusAPI) GetEntityByID(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCP4Range :: adds a DHCP v4 range
func (service *ProteusAPI) AddDHCP4Range(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCP4RangeBySize :: adds a dhcp v4 range by size
func (service *ProteusAPI) AddDHCP4RangeBySize(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddDHCP6Range :: adds dhcp v6 range
func (service *ProteusAPI) AddDHCP6Range(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntityByCIDR :: get entity by CIDR
func (service *ProteusAPI) GetEntityByCIDR(request *Long) (*APIEntity, error) {
	response := new(APIEntity)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddIP4BlockByRange :: adds an ipv4 block by range
func (service *ProteusAPI) AddIP4BlockByRange(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AddEnumZone :: adds an enum by zone
func (service *ProteusAPI) AddEnumZone(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// LoginWithOptions :: logs in with options
func (service *ProteusAPI) LoginWithOptions(request *String) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Could be broken, no type in last parens, first param, * */

// BreakXHAPair :: breaks the XHA pair
func (service *ProteusAPI) BreakXHAPair(request *Long) (*Long, error) {
	response := new(Long)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEntityByRange :: get entity by range
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

// SOAPEnvelope :: container for xml envelope
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

// SOAPHeader :: SOAP header
type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

// SOAPBody :: SOAP body
type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// SOAPFault :: SOAP Fault
type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

// BasicAuth :: basic auth credentials
type BasicAuth struct {
	Login    string
	Password string
}

// SOAPClient :: a soap client
type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

// UnmarshalXML :: unmarshals xml to internal structure
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

// NewSOAPClient :: returns a new soap client
func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

// Call :: calls to the WSDL SOAP API
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
