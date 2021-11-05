package consul

import (
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/model"
)

//
// 放置consul的service和instance的实体对象信息struct
//

//
// consule这个地址获取的数据: http://localhost:8500/v1/catalog/services
//{
//    "consul": [],
//    "fgmp-5g-service": [
//        "fgmp-5g-service",
//        "logs"
//    ],
//    "fgmp-base-service": [
//        "fgmp-base-service",
//        "logs"
//    ],
//    "node_exporter": [
//        "prometheus",
//        "node_exporter"
//    ]
//}
//
//

const (
	node       = "5G"
	localhost  = "127.0.0.1"
	datacenter = "dc1"
	id         = "03f98bf7-f479-83a1-e0f9-1d1ea676f954"
)

type Services map[string][]string

func ConvertServices(nacosServices model.ServiceList) Services {
	services := Services{}
	services["consul"] = []string{}
	for _, value := range nacosServices.Doms {
		services[value] = []string{value}
	}
	return services
}

//
// consul这个地址获取的数据 curl http://localhost:8500/v1/catalog/service/node_exporter
//[
//    {
//        "ID": "03f98bf7-f479-83a1-e0f9-1d1ea676f954",
//        "Node": "liuxu",
//        "Address": "127.0.0.1",
//        "Datacenter": "dc1",
//        "TaggedAddresses": {
//            "lan": "127.0.0.1",
//            "lan_ipv4": "127.0.0.1",
//            "wan": "127.0.0.1",
//            "wan_ipv4": "127.0.0.1"
//        },
//        "NodeMeta": {
//            "consul-network-segment": ""
//        },
//        "ServiceKind": "",
//        "ServiceID": "node_exporter_172.16.16.37",
//        "ServiceName": "node_exporter",
//        "ServiceTags": [
//            "node_exporter",
//            "prometheus"
//        ],
//        "ServiceAddress": "172.16.16.37",
//        "ServiceTaggedAddresses": {
//            "lan_ipv4": {
//                "Address": "172.16.16.37",
//                "Port": 19049
//            },
//            "wan_ipv4": {
//                "Address": "172.16.16.37",
//                "Port": 19049
//            }
//        },
//        "ServiceWeights": {
//            "Passing": 1,
//            "Warning": 1
//        },
//        "ServiceMeta": {},
//        "ServicePort": 19049,
//        "ServiceSocketPath": "",
//        "ServiceEnableTagOverride": false,
//        "ServiceProxy": {
//            "Mode": "",
//            "MeshGateway": {},
//            "Expose": {}
//        },
//        "ServiceConnect": {},
//        "CreateIndex": 10804,
//        "ModifyIndex": 10804
//    },
//    {
//        "ID": "03f98bf7-f479-83a1-e0f9-1d1ea676f954",
//        "Node": "liuxu",
//        "Address": "127.0.0.1",
//        "Datacenter": "dc1",
//        "TaggedAddresses": {
//            "lan": "127.0.0.1",
//            "lan_ipv4": "127.0.0.1",
//            "wan": "127.0.0.1",
//            "wan_ipv4": "127.0.0.1"
//        },
//        "NodeMeta": {
//            "consul-network-segment": ""
//        },
//        "ServiceKind": "",
//        "ServiceID": "node_exporter_172.16.16.46",
//        "ServiceName": "node_exporter",
//        "ServiceTags": [
//            "node_exporter",
//            "prometheus"
//        ],
//        "ServiceAddress": "172.16.16.46",
//        "ServiceTaggedAddresses": {
//            "lan_ipv4": {
//                "Address": "172.16.16.46",
//                "Port": 19049
//            },
//            "wan_ipv4": {
//                "Address": "172.16.16.46",
//                "Port": 19049
//            }
//        },
//        "ServiceWeights": {
//            "Passing": 1,
//            "Warning": 1
//        },
//        "ServiceMeta": {},
//        "ServicePort": 19049,
//        "ServiceSocketPath": "",
//        "ServiceEnableTagOverride": false,
//        "ServiceProxy": {
//            "Mode": "",
//            "MeshGateway": {},
//            "Expose": {}
//        },
//        "ServiceConnect": {},
//        "CreateIndex": 10796,
//        "ModifyIndex": 10796
//    }
//]

type Instance struct {
	ID                       string             `json:"ID"`
	Node                     string             `json:"Node"`
	Address                  string             `json:"Address"`
	Datacenter               string             `json:"Datacenter"`
	TaggedAddresses          TaggedAddresses    `json:"TaggedAddresses"`
	NodeMeta                 map[string]string  `json:"NodeMeta"`
	ServiceKind              string             `json:"ServiceKind"`
	ServiceID                string             `json:"ServiceID"`
	ServiceName              string             `json:"ServiceName"`
	ServiceTags              []string           `json:"ServiceTags"`
	ServiceAddress           string             `json:"ServiceAddress"`
	ServiceTaggedAddresses   map[string]Address `json:"ServiceTaggedAddresses"`
	ServiceWeights           ServiceWeights     `json:"ServiceWeights"`
	ServiceMeta              map[string]string  `json:"ServiceMeta"`
	ServicePort              int32              `json:"ServiceProt"`
	ServiceSocketPath        string             `json:"ServiceSocketPath"`
	ServiceEnableTagOverride bool               `json:"ServiceEnableTagOverride"`
	ServiceProxy             ServiceProxy       `json:"ServiceProxy"`
	ServiceConnect           ServiceConnect     `json:"ServiceConnect"`
	CreateIndex              int32              `json:"CreateIndex"`
	ModifyIndex              int32              `json:"ModifyIndex"`
}

type TaggedAddresses struct {
	Lan     string `json:"lan"`
	LanIpv4 string `json:"lan_ipv4"`
	Wan     string `json:"wan"`
	WanIpv4 string `json:"wan_ipv4"`
}

type Address struct {
	Address string `json:"Address"`
	Port    int32  `json:"Port"`
}

type ServiceWeights struct {
	Passing int32 `json:"Passing"`
	Warning int32 `json:"Warning"`
}

type ServiceProxy struct {
	Mode        string      `json:"Mode"`
	MeshGateway MeshGateway `json:"MeshGateway"`
	Expose      Expose      `json:"Expose"`
}

type MeshGateway struct {
}

type Expose struct {
}

type ServiceConnect struct {
}

func ConvertInstances(sources []model.Instance) []Instance {
	targets := []Instance{}
	for _, source := range sources {
		targets = append(targets, convertInstance(source))
	}
	return targets
}

func convertInstance(source model.Instance) Instance {
	target := Instance{}
	target.ID = id
	target.Node = node
	target.Address = localhost
	target.Datacenter = datacenter
	target.TaggedAddresses = TaggedAddresses{
		Lan:     localhost,
		LanIpv4: localhost,
		Wan:     localhost,
		WanIpv4: localhost,
	}
	target.NodeMeta = map[string]string{
		"consul-network-segment": "",
	}
	target.ServiceKind = ""
	target.ServiceID = source.InstanceId
	target.ServiceName = source.ServiceName
	target.ServiceTags = []string{source.ServiceName}
	target.ServiceAddress = source.Ip
	target.ServiceTaggedAddresses = map[string]Address{
		"lan_ipv4": {
			Address: source.Ip,
			Port:    int32(source.Port),
		},
		"wan_ipv4": {
			Address: source.Ip,
			Port:    int32(source.Port),
		},
	}
	target.ServiceWeights = ServiceWeights{
		Passing: int32(source.Weight),
		Warning: int32(source.Weight),
	}
	target.ServiceMeta = convertNacosServiceMeta(source.Metadata)
	target.ServicePort = int32(source.Port)
	target.ServiceSocketPath = ""
	target.ServiceEnableTagOverride = false
	target.ServiceProxy = ServiceProxy{
		Mode:        "",
		MeshGateway: MeshGateway{},
		Expose:      Expose{},
	}
	target.ServiceConnect = ServiceConnect{}
	target.CreateIndex = 1996
	target.ModifyIndex = 1996

	return target
}

func convertNacosServiceMeta(source map[string]string) map[string]string {
	target := make(map[string]string, len(source))
	for key, value := range source {
		target[strings.ReplaceAll(key, ".", "_")] = value
	}
	return target
}

const Agent string = `
{
    "Config": {
        "Datacenter": "dc1",
        "PrimaryDatacenter": "dc1",
        "NodeName": "5G",
        "NodeID": "03f98bf7-f479-83a1-e0f9-1d1ea676f954",
        "Revision": "c976ffd2d",
        "Server": true,
        "Version": "1.10.3"
    },
    "DebugConfig": {
        "ACLDatacenter": "dc1",
        "ACLDefaultPolicy": "allow",
        "ACLDisabledTTL": "2m0s",
        "ACLDownPolicy": "extend-cache",
        "ACLEnableKeyListPolicy": false,
        "ACLMasterToken": "hidden",
        "ACLPolicyTTL": "30s",
        "ACLRoleTTL": "0s",
        "ACLTokenReplication": false,
        "ACLTokenTTL": "30s",
        "ACLTokens": {
            "ACLAgentMasterToken": "hidden",
            "ACLAgentToken": "hidden",
            "ACLDefaultToken": "hidden",
            "ACLReplicationToken": "hidden",
            "DataDir": "",
            "EnablePersistence": false,
            "EnterpriseConfig": {}
        },
        "ACLsEnabled": false,
        "AEInterval": "1m0s",
        "AdvertiseAddrLAN": "127.0.0.1",
        "AdvertiseAddrWAN": "127.0.0.1",
        "AdvertiseReconnectTimeout": "0s",
        "AllowWriteHTTPFrom": [],
        "AutoConfig": {
            "Authorizer": {
                "AllowReuse": false,
                "AuthMethod": {
                    "ACLAuthMethodEnterpriseFields": {},
                    "Config": {
                        "BoundAudiences": null,
                        "BoundIssuer": "",
                        "ClaimMappings": null,
                        "ClockSkewLeeway": 0,
                        "ExpirationLeeway": 0,
                        "JWKSCACert": "",
                        "JWKSURL": "",
                        "JWTSupportedAlgs": null,
                        "JWTValidationPubKeys": null,
                        "ListClaimMappings": null,
                        "NotBeforeLeeway": 0,
                        "OIDCDiscoveryCACert": "",
                        "OIDCDiscoveryURL": ""
                    },
                    "Description": "",
                    "DisplayName": "",
                    "EnterpriseMeta": {},
                    "MaxTokenTTL": "0s",
                    "Name": "Auto Config Authorizer",
                    "RaftIndex": {
                        "CreateIndex": 0,
                        "ModifyIndex": 0
                    },
                    "TokenLocality": "",
                    "Type": "jwt"
                },
                "ClaimAssertions": [],
                "Enabled": false
            },
            "DNSSANs": [],
            "Enabled": false,
            "IPSANs": [],
            "IntroToken": "hidden",
            "IntroTokenFile": "",
            "ServerAddresses": []
        },
        "AutoEncryptAllowTLS": false,
        "AutoEncryptDNSSAN": [],
        "AutoEncryptIPSAN": [],
        "AutoEncryptTLS": false,
        "AutopilotCleanupDeadServers": true,
        "AutopilotDisableUpgradeMigration": false,
        "AutopilotLastContactThreshold": "200ms",
        "AutopilotMaxTrailingLogs": 250,
        "AutopilotMinQuorum": 0,
        "AutopilotRedundancyZoneTag": "",
        "AutopilotServerStabilizationTime": "10s",
        "AutopilotUpgradeVersionTag": "",
        "BindAddr": "127.0.0.1",
        "Bootstrap": false,
        "BootstrapExpect": 0,
        "CAFile": "",
        "CAPath": "",
        "Cache": {
            "EntryFetchMaxBurst": 2,
            "EntryFetchRate": 1.7976931348623157e+308,
            "Logger": null
        },
        "CertFile": "",
        "CheckDeregisterIntervalMin": "1m0s",
        "CheckOutputMaxSize": 4096,
        "CheckReapInterval": "30s",
        "CheckUpdateInterval": "5m0s",
        "Checks": [],
        "ClientAddrs": [
            "0.0.0.0"
        ],
        "ConfigEntryBootstrap": [],
        "ConnectCAConfig": {},
        "ConnectCAProvider": "",
        "ConnectEnabled": true,
        "ConnectMeshGatewayWANFederationEnabled": false,
        "ConnectSidecarMaxPort": 21255,
        "ConnectSidecarMinPort": 21000,
        "ConnectTestCALeafRootChangeSpread": "0s",
        "ConsulCoordinateUpdateBatchSize": 128,
        "ConsulCoordinateUpdateMaxBatches": 5,
        "ConsulCoordinateUpdatePeriod": "100ms",
        "ConsulRaftElectionTimeout": "52ms",
        "ConsulRaftHeartbeatTimeout": "35ms",
        "ConsulRaftLeaderLeaseTimeout": "20ms",
        "ConsulServerHealthInterval": "10ms",
        "DNSARecordLimit": 0,
        "DNSAddrs": [
            "tcp://0.0.0.0:8600",
            "udp://0.0.0.0:8600"
        ],
        "DNSAllowStale": true,
        "DNSAltDomain": "",
        "DNSCacheMaxAge": "0s",
        "DNSDisableCompression": false,
        "DNSDomain": "consul.",
        "DNSEnableTruncate": false,
        "DNSMaxStale": "87600h0m0s",
        "DNSNodeMetaTXT": true,
        "DNSNodeTTL": "0s",
        "DNSOnlyPassing": false,
        "DNSPort": 8600,
        "DNSRecursorTimeout": "2s",
        "DNSRecursors": [],
        "DNSSOA": {
            "Expire": 86400,
            "Minttl": 0,
            "Refresh": 3600,
            "Retry": 600
        },
        "DNSServiceTTL": {},
        "DNSUDPAnswerLimit": 3,
        "DNSUseCache": false,
        "DataDir": "",
        "Datacenter": "dc1",
        "DefaultQueryTime": "5m0s",
        "DevMode": true,
        "DisableAnonymousSignature": true,
        "DisableCoordinates": false,
        "DisableHTTPUnprintableCharFilter": false,
        "DisableHostNodeID": true,
        "DisableKeyringFile": true,
        "DisableRemoteExec": true,
        "DisableUpdateCheck": false,
        "DiscardCheckOutput": false,
        "DiscoveryMaxStale": "0s",
        "EnableAgentTLSForChecks": false,
        "EnableCentralServiceConfig": true,
        "EnableDebug": true,
        "EnableLocalScriptChecks": false,
        "EnableRemoteScriptChecks": false,
        "EncryptKey": "hidden",
        "EncryptVerifyIncoming": true,
        "EncryptVerifyOutgoing": true,
        "EnterpriseRuntimeConfig": {},
        "ExposeMaxPort": 21755,
        "ExposeMinPort": 21500,
        "GRPCAddrs": [
            "tcp://0.0.0.0:8502"
        ],
        "GRPCPort": 8502,
        "GossipLANGossipInterval": "100ms",
        "GossipLANGossipNodes": 3,
        "GossipLANProbeInterval": "100ms",
        "GossipLANProbeTimeout": "100ms",
        "GossipLANRetransmitMult": 4,
        "GossipLANSuspicionMult": 3,
        "GossipWANGossipInterval": "100ms",
        "GossipWANGossipNodes": 3,
        "GossipWANProbeInterval": "100ms",
        "GossipWANProbeTimeout": "100ms",
        "GossipWANRetransmitMult": 4,
        "GossipWANSuspicionMult": 3,
        "HTTPAddrs": [
            "tcp://0.0.0.0:8500"
        ],
        "HTTPBlockEndpoints": [],
        "HTTPMaxConnsPerClient": 200,
        "HTTPMaxHeaderBytes": 0,
        "HTTPPort": 8500,
        "HTTPResponseHeaders": {},
        "HTTPSAddrs": [],
        "HTTPSHandshakeTimeout": "5s",
        "HTTPSPort": -1,
        "HTTPUseCache": true,
        "KVMaxValueSize": 524288,
        "KeyFile": "hidden",
        "LeaveDrainTime": "5s",
        "LeaveOnTerm": false,
        "Logging": {
            "EnableSyslog": false,
            "LogFilePath": "",
            "LogJSON": false,
            "LogLevel": "DEBUG",
            "LogRotateBytes": 0,
            "LogRotateDuration": "0s",
            "LogRotateMaxFiles": 0,
            "Name": "",
            "SyslogFacility": "LOCAL0"
        },
        "MaxQueryTime": "10m0s",
        "NodeID": "03f98bf7-f479-83a1-e0f9-1d1ea676f954",
        "NodeMeta": {},
        "NodeName": "5G",
        "PidFile": "",
        "PrimaryDatacenter": "dc1",
        "PrimaryGateways": [],
        "PrimaryGatewaysInterval": "30s",
        "RPCAdvertiseAddr": "tcp://127.0.0.1:8300",
        "RPCBindAddr": "tcp://127.0.0.1:8300",
        "RPCConfig": {
            "EnableStreaming": true
        },
        "RPCHandshakeTimeout": "5s",
        "RPCHoldTimeout": "7s",
        "RPCMaxBurst": 1000,
        "RPCMaxConnsPerClient": 100,
        "RPCProtocol": 2,
        "RPCRateLimit": -1,
        "RaftProtocol": 3,
        "RaftSnapshotInterval": "30s",
        "RaftSnapshotThreshold": 16384,
        "RaftTrailingLogs": 10240,
        "ReadReplica": false,
        "ReconnectTimeoutLAN": "0s",
        "ReconnectTimeoutWAN": "0s",
        "RejoinAfterLeave": false,
        "RetryJoinIntervalLAN": "30s",
        "RetryJoinIntervalWAN": "30s",
        "RetryJoinLAN": [],
        "RetryJoinMaxAttemptsLAN": 0,
        "RetryJoinMaxAttemptsWAN": 0,
        "RetryJoinWAN": [],
        "Revision": "c976ffd2d",
        "SegmentLimit": 64,
        "SegmentName": "",
        "SegmentNameLimit": 64,
        "Segments": [],
        "SerfAdvertiseAddrLAN": "tcp://127.0.0.1:8301",
        "SerfAdvertiseAddrWAN": "tcp://127.0.0.1:8302",
        "SerfAllowedCIDRsLAN": [],
        "SerfAllowedCIDRsWAN": [],
        "SerfBindAddrLAN": "tcp://127.0.0.1:8301",
        "SerfBindAddrWAN": "tcp://127.0.0.1:8302",
        "SerfPortLAN": 8301,
        "SerfPortWAN": 8302,
        "ServerMode": true,
        "ServerName": "",
        "ServerPort": 8300,
        "Services": [],
        "SessionTTLMin": "0s",
        "SkipLeaveOnInt": true,
        "StartJoinAddrsLAN": [],
        "StartJoinAddrsWAN": [],
        "SyncCoordinateIntervalMin": "15s",
        "SyncCoordinateRateTarget": 64,
        "TLSCipherSuites": [],
        "TLSMinVersion": "tls12",
        "TLSPreferServerCipherSuites": false,
        "TaggedAddresses": {
            "lan": "127.0.0.1",
            "lan_ipv4": "127.0.0.1",
            "wan": "127.0.0.1",
            "wan_ipv4": "127.0.0.1"
        },
        "Telemetry": {
            "AllowedPrefixes": [],
            "BlockedPrefixes": [],
            "CirconusAPIApp": "",
            "CirconusAPIToken": "hidden",
            "CirconusAPIURL": "",
            "CirconusBrokerID": "",
            "CirconusBrokerSelectTag": "",
            "CirconusCheckDisplayName": "",
            "CirconusCheckForceMetricActivation": "",
            "CirconusCheckID": "",
            "CirconusCheckInstanceID": "",
            "CirconusCheckSearchTag": "",
            "CirconusCheckTags": "",
            "CirconusSubmissionInterval": "",
            "CirconusSubmissionURL": "",
            "Disable": false,
            "DisableCompatOneNine": false,
            "DisableHostname": false,
            "DogstatsdAddr": "",
            "DogstatsdTags": [],
            "FilterDefault": true,
            "MetricsPrefix": "consul",
            "PrometheusOpts": {
                "CounterDefinitions": [],
                "Expiration": "0s",
                "GaugeDefinitions": [],
                "Registerer": null,
                "SummaryDefinitions": []
            },
            "StatsdAddr": "",
            "StatsiteAddr": ""
        },
        "TranslateWANAddrs": false,
        "TxnMaxReqLen": 524288,
        "UIConfig": {
            "ContentPath": "/ui/",
            "DashboardURLTemplates": {},
            "Dir": "",
            "Enabled": true,
            "MetricsProvider": "",
            "MetricsProviderFiles": [],
            "MetricsProviderOptionsJSON": "",
            "MetricsProxy": {
                "AddHeaders": [],
                "BaseURL": "",
                "PathAllowlist": []
            }
        },
        "UnixSocketGroup": "",
        "UnixSocketMode": "",
        "UnixSocketUser": "",
        "UseStreamingBackend": true,
        "VerifyIncoming": false,
        "VerifyIncomingHTTPS": false,
        "VerifyIncomingRPC": false,
        "VerifyOutgoing": false,
        "VerifyServerHostname": false,
        "Version": "1.10.3",
        "VersionPrerelease": "",
        "Watches": []
    },
    "Coord": {
        "Vec": [
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0
        ],
        "Error": 1.5,
        "Adjustment": 0,
        "Height": 0.00001
    },
    "Member": {
        "Name": "5G",
        "Addr": "127.0.0.1",
        "Port": 8301,
        "Tags": {
            "acls": "0",
            "build": "1.10.3:c976ffd2",
            "dc": "dc1",
            "ft_fs": "1",
            "ft_si": "1",
            "id": "03f98bf7-f479-83a1-e0f9-1d1ea676f954",
            "port": "8300",
            "raft_vsn": "3",
            "role": "consul",
            "segment": "",
            "vsn": "2",
            "vsn_max": "3",
            "vsn_min": "2",
            "wan_join_port": "8302"
        },
        "Status": 1,
        "ProtocolMin": 1,
        "ProtocolMax": 5,
        "ProtocolCur": 2,
        "DelegateMin": 2,
        "DelegateMax": 5,
        "DelegateCur": 4
    },
    "Stats": {
        "agent": {
            "check_monitors": "0",
            "check_ttls": "0",
            "checks": "0",
            "services": "5"
        },
        "build": {
            "prerelease": "",
            "revision": "c976ffd2",
            "version": "1.10.3"
        },
        "consul": {
            "acl": "disabled",
            "bootstrap": "false",
            "known_datacenters": "1",
            "leader": "true",
            "leader_addr": "127.0.0.1:8300",
            "server": "true"
        },
        "raft": {
            "applied_index": "20888",
            "commit_index": "20888",
            "fsm_pending": "0",
            "last_contact": "0",
            "last_log_index": "20888",
            "last_log_term": "2",
            "last_snapshot_index": "16384",
            "last_snapshot_term": "2",
            "latest_configuration": "[{Suffrage:Voter ID:03f98bf7-f479-83a1-e0f9-1d1ea676f954 Address:127.0.0.1:8300}]",
            "latest_configuration_index": "0",
            "num_peers": "0",
            "protocol_version": "3",
            "protocol_version_max": "3",
            "protocol_version_min": "0",
            "snapshot_version_max": "1",
            "snapshot_version_min": "0",
            "state": "Leader",
            "term": "2"
        },
        "runtime": {
            "arch": "amd64",
            "cpu_count": "4",
            "goroutines": "137",
            "max_procs": "4",
            "os": "linux",
            "version": "go1.16.7"
        },
        "serf_lan": {
            "coordinate_resets": "0",
            "encrypted": "false",
            "event_queue": "1",
            "event_time": "2",
            "failed": "0",
            "health_score": "0",
            "intent_queue": "0",
            "left": "0",
            "member_time": "1",
            "members": "1",
            "query_queue": "0",
            "query_time": "1"
        },
        "serf_wan": {
            "coordinate_resets": "0",
            "encrypted": "false",
            "event_queue": "0",
            "event_time": "1",
            "failed": "0",
            "health_score": "0",
            "intent_queue": "0",
            "left": "0",
            "member_time": "1",
            "members": "1",
            "query_queue": "0",
            "query_time": "1"
        }
    },
    "Meta": {
        "consul-network-segment": ""
    },
    "xDS": {
        "SupportedProxies": {
            "envoy": [
                "1.18.4",
                "1.17.4",
                "1.16.5",
                "1.15.5"
            ]
        }
    }
}
`

//
// 支持 /v1/health/service/服务名字的信息
//
//[
//    {
//        "Node": {
//            "ID": "03f98bf7-f479-83a1-e0f9-1d1ea676f954",
//            "Node": "liuxu",
//            "Address": "127.0.0.1",
//            "Datacenter": "dc1",
//            "TaggedAddresses": {
//                "lan": "127.0.0.1",
//                "lan_ipv4": "127.0.0.1",
//                "wan": "127.0.0.1",
//                "wan_ipv4": "127.0.0.1"
//            },
//            "Meta": {
//                "consul-network-segment": ""
//            },
//            "CreateIndex": 7,
//            "ModifyIndex": 7
//        },
//        "Service": {
//            "ID": "node_exporter_172.16.16.37",
//            "Service": "node_exporter",
//            "Tags": [
//                "node_exporter",
//                "prometheus"
//            ],
//            "Address": "172.16.16.37",
//            "TaggedAddresses": {
//                "lan_ipv4": {
//                    "Address": "172.16.16.37",
//                    "Port": 19049
//                },
//                "wan_ipv4": {
//                    "Address": "172.16.16.37",
//                    "Port": 19049
//                }
//            },
//            "Meta": null,
//            "Port": 19049,
//            "Weights": {
//                "Passing": 1,
//                "Warning": 1
//            },
//            "EnableTagOverride": false,
//            "Proxy": {
//                "Mode": "",
//                "MeshGateway": {},
//                "Expose": {}
//            },
//            "Connect": {},
//            "CreateIndex": 10804,
//            "ModifyIndex": 10804
//        },
//        "Checks": [
//            {
//                "Node": "liuxu",
//                "CheckID": "serfHealth",
//                "Name": "Serf Health Status",
//                "Status": "passing",
//                "Notes": "",
//                "Output": "Agent alive and reachable",
//                "ServiceID": "",
//                "ServiceName": "",
//                "ServiceTags": [],
//                "Type": "",
//                "Interval": "",
//                "Timeout": "",
//                "ExposedPort": 0,
//                "Definition": {},
//                "CreateIndex": 12,
//                "ModifyIndex": 12
//            }
//        ]
//    },
//]

type Health struct {
	Node    HealthNode    `json:"Node"`
	Service HealthService `json:"Service"`
	Checks  []HealthCheck `json:"Checks"`
}

type HealthNode struct {
	ID              string            `json:"ID"`
	Node            string            `json:"Node"`
	Address         string            `json:"Address"`
	Datacenter      string            `json:"Datecenter"`
	TaggedAddresses TaggedAddresses   `json:"TaggedAddresses"`
	Meta            map[string]string `json:"Meta"`
	CreateIndex     int32             `json:"CreateIndex"`
	ModifyIndex     int32             `json:"ModifyIndex"`
}

type HealthService struct {
	ID                string             `json:"ID"`
	Service           string             `json:"Service"`
	Tags              []string           `json:"Tags"`
	Address           string             `json:"Address"`
	TaggedAddresses   map[string]Address `json:"TaggedAddresses"`
	Meta              map[string]string  `json:"Meta"`
	Port              int32              `json:"Port"`
	Weights           ServiceWeights     `json:"Weights"`
	EnableTagOverride bool               `json:"EnableTagOverride"`
	Proxy             ServiceProxy       `json:"Proxy"`
	Connect           ServiceConnect     `json:"Connect"`
	CreateIndex       int32              `json:"CreateIndex"`
	ModifyIndex       int32              `json:"ModifyIndex"`
}

type HealthCheck struct {
	Node        string           `json:"Node"`
	CheckID     string           `json:"CheckID"`
	Name        string           `json:"Name"`
	Status      string           `json:"Status"`
	Notes       string           `json:"Notes"`
	Output      string           `json:"Output"`
	ServiceID   string           `json:"ServiceID"`
	ServiceName string           `json:"ServiceName"`
	ServiceTags []string         `json:"ServiceTags"`
	Type        string           `json:"Type"`
	Interval    string           `json:"Interval"`
	Timeout     string           `json:"Timeout"`
	ExposePort  int32            `json:"ExposePort"`
	Definition  HealthDefinition `json:"Definition"`
	CreateIndex int32            `json:"CreateIndex"`
	ModifyIndex int32            `json:"ModifyIndex"`
}

type HealthDefinition struct {
}

func ConvertHealths(sources []model.Instance) []Health {
	targets := []Health{}
	for _, source := range sources {
		targets = append(targets, convertHealth(source))
	}
	return targets
}

func convertHealth(source model.Instance) Health {
	port := source.Port
	if value, ok := source.Metadata["management.port"]; ok {
		temp, _ := strconv.ParseInt(value, 10, 64)
		port = uint64(temp)
	}
	healthNode := HealthNode{
		ID:         id,
		Node:       node,
		Address:    localhost,
		Datacenter: datacenter,
		TaggedAddresses: TaggedAddresses{
			Lan:     localhost,
			LanIpv4: localhost,
			Wan:     localhost,
			WanIpv4: localhost,
		},
		Meta:        map[string]string{"consul-network-segment": ""},
		CreateIndex: 12,
		ModifyIndex: 12,
	}

	healthService := HealthService{
		ID:      source.InstanceId,
		Service: source.ServiceName,
		Tags:    []string{source.ServiceName},
		Address: source.Ip,
		TaggedAddresses: map[string]Address{
			"lan_ipv4": {
				Address: source.Ip,
				Port:    int32(port),
			},
			"wan_ipv4": {
				Address: source.Ip,
				Port:    int32(port),
			},
		},
		Meta: convertNacosServiceMeta(source.Metadata),
		Port: int32(port),
		Weights: ServiceWeights{
			Passing: int32(source.Weight),
			Warning: int32(source.Weight),
		},
		EnableTagOverride: false,
		Proxy: ServiceProxy{
			Mode:        "",
			MeshGateway: MeshGateway{},
			Expose:      Expose{},
		},
		Connect:     ServiceConnect{},
		CreateIndex: 10796,
		ModifyIndex: 10796,
	}

	healthChecks := []HealthCheck{
		{
			Node:        node,
			CheckID:     "serfHealth",
			Name:        "Serf Health Status",
			Status:      "passing",
			Notes:       "",
			Output:      "Agent alive and reachable",
			ServiceID:   "",
			ServiceName: "",
			ServiceTags: []string{},
			Type:        "",
			Interval:    "",
			Timeout:     "",
			ExposePort:  0,
			Definition:  HealthDefinition{},
			CreateIndex: 12,
			ModifyIndex: 12,
		},
	}

	return Health{
		Node:    healthNode,
		Service: healthService,
		Checks:  healthChecks,
	}
}
