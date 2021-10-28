package consul

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

type Services map[string][]string

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
	WanIpv4 string `json:"WanIpv4"`
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
