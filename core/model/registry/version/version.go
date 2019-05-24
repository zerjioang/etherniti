package version

import "strconv"

type ContractVersion struct {
	Address string `json:"address"`
	Minor   int    `json:"minor"`
	Major   int    `json:"major"`
}

// return the stringified data of version information as
// major.minor
// Example: 1.5
func (v ContractVersion) String() string {
	return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor)
}

// check if contract information version is valid or not
func (v ContractVersion) Valid() bool {
	return v.Minor != 0 && v.Major != 0
}
