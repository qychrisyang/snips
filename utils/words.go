// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package utils

var capitalizedToCapitalizedWordsMap = map[string]string{
	"Dns":       "DNS",
	"Dyn":       "DYN",
	"Eip":       "EIP",
	"Keypair":   "KeyPair",
	"Vxnet":     "VxNet",

	"ITEM":      "Item",
	"CREATE":    "Create",

	"UTC-1":     "UTCMinus1",
	"UTC-2":     "UTCMinus2",
	"UTC-3":     "UTCMinus3",
	"UTC-4":     "UTCMinus4",
	"UTC-5":     "UTCMinus5",
	"UTC-6":     "UTCMinus6",
	"UTC-7":     "UTCMinus7",
	"UTC-8":     "UTCMinus8",
	"UTC-9":     "UTCMinus9",
	"UTC-10":    "UTCMinus10",
	"UTC-11":    "UTCMinus11",
	"UTC-12":    "UTCMinus12",
	"UTC+1":     "UTCPlus1",
	"UTC+2":     "UTCPlus2",
	"UTC+3":     "UTCPlus3",
	"UTC+4":     "UTCPlus4",
	"UTC+5":     "UTCPlus5",
	"UTC+6":     "UTCPlus6",
	"UTC+7":     "UTCPlus7",
	"UTC+8":     "UTCPlus8",
	"UTC+9":     "UTCPlus9",
	"UTC+10":    "UTCPlus10",
	"UTC+11":    "UTCPlus11",
	"UTC+12":    "UTCPlus12",
}

var lowerCaseToLowercaseWordsMap = map[string]string{
	"lastest": "latest", // Fix typo
}

var lowerCaseToCapitalizedWordsMap = map[string]string{
	"acl":           "ACL",
	"cors":          "CORS",
	"cpu":           "CPU",
	"datadir":       "DataDir",
	"dhcp":          "DHCP",
	"dns":           "DNS",
	"dyn":           "DYN",
	"eip":           "EIP",
	"eips":          "EIPs",
	"http":          "HTTP",
	"icp":           "ICP",
	"id":            "ID",
	"ids":           "IDs",
	"innodb":        "InnoDB",
	"io":            "IO",
	"ip":            "IP",
	"ips":           "IPs",
	"ipset":         "IPSet",
	"ipsets":        "IPSets",
	"keypair":       "KeyPair",
	"keypairs":      "KeyPairs",
	"lastest":       "Latest", // Fix typo
	"loadbalancer":  "LoadBalancer",
	"loadbalancers": "LoadBalancers",
	"md5":           "MD5",
	"newsid":        "NewSID",
	"nic":           "NIC",
	"os":            "OS",
	"opt":           "OPT",
	"qingstor":      "QingStor",
	"qingcloud":     "QingCloud",
	"qs":            "QS",
	"rdb":           "RDB",
	"rdbs":          "RDBs",
	"sha1":          "SHA1",
	"sha256":        "SHA256",
	"sql":           "SQL",
	"tmp":           "TMP",
	"tmpdir":        "TMPDir",
	"topslave":      "TopSlave",
	"trx":           "TRX",
	"ui":            "UI",
	"uri":           "URI",
	"url":           "URL",
	"usb":           "USB",
	"uuid":          "UUID",
	"vcpus":         "VCPUs",
	"vxnet":         "VxNet",
	"vxnets":        "VxNets",

	"sso":           "SSO",
	"ldap":          "LDAP",
	"ad":            "AD",
	"utc":           "UTC",
}

var abbreviateWordsMap = []string{
	"ACL",
	"CORS",
	"CPU",
	"DHCP",
	"DNS",
	"DYN",
	"EIP",
	"ETag",
	"IaaS",
	"ICP",
	"ID",
	"IO",
	"IP",
	"IPSets",
	"MD5",
	"NIC",
	"OAuth",
	"OS",
	"OPT",
	"QingStor",
	"QingCloud",
	"QS",
	"RDB",
	"SQL",
	"SSO",
	"TMP",
	"TMPDir",
	"TRX",
	"UI",
	"URL",
	"UUID",
	"VCPUs",
	"VxNet",

	"SSO",
	"LDAP",
	"AD",
	"UTC",
}
