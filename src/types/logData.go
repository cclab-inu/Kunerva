package types

import "encoding/json"

// K8sNetworkLog Structure
type K8sNetworkLog struct {
	SrcNamespace string `json:"src_namespace,omitempty" bson:"src_namespace"`
	SrcPodName   string `json:"src_pod_name,omitempty" bson:"src_pod_name"`

	DstNamespace string `json:"dst_namespace,omitempty" bson:"dst_namespace"`
	DstPodName   string `json:"dst_pod_name,omitempty" bson:"dst_pod_name"`

	EtherType int `json:"ether_type,omitempty" bson:"ether_type"` // not used, we assume all the ipv4

	Protocol int    `json:"protocol,omitempty" bson:"protocol"`
	SrcIP    string `json:"src_ip,omitempty" bson:"src_ip"`
	DstIP    string `json:"dst_ip,omitempty" bson:"dst_ip"`
	SrcPort  int    `json:"src_port,omitempty" bson:"src_port"`
	DstPort  int    `json:"dst_port,omitempty" bson:"dst_port"`

	SynFlag bool `json:"syn_flag,omitempty" bson:"syn_flag"` // for tcp

	DNSQuery string `json:"dns_query,omitempty" bson:"dns_query"` // for L7 dns

	HTTPMethod string `json:"http_method,omitempty" bson:"http_method"` // for L7 http
	HTTPPath   string `json:"http_path,omitempty" bson:"http_path"`     // for L7 http

	Direction string `json:"direction,omitempty" bson:"direction"` // ingress or egress

	Action string `json:"action,omitempty" bson:"action"`
}

// NetworkLogEvent
type NetworkLogEvent struct {
	Time                  string          `json:"time,omitempty"`
	Verdict               string          `json:"verdict,omitempty"`
	DropReason            int             `json:"drop_reason,omitempty"`
	Ethernet              json.RawMessage `json:"ethernet,omitempty"`
	IP                    json.RawMessage `json:"IP,omitempty"`
	L4                    json.RawMessage `json:"l4,omitempty"`
	Source                json.RawMessage `json:"source,omitempty"`
	Destination           json.RawMessage `json:"destination,omitempty"`
	Type                  string          `json:"Type,omitempty"`
	NodeName              string          `json:"node_name,omitempty"`
	SourceNames           string          `json:"source_names,omitempty"`
	DestinationNames      string          `json:"destination_names,omitempty"`
	L7                    json.RawMessage `json:"l7,omitempty"`
	Reply                 bool            `json:"reply,omitempty"`
	EventType             json.RawMessage `json:"event_type,omitempty"`
	SourceService         json.RawMessage `json:"source_service,omitempty"`
	DestinationService    json.RawMessage `json:"destination_service,omitempty"`
	TrafficDirection      string          `json:"traffic_direction,omitempty"`
	PolicyMatchType       int             `json:"policy_match_type,omitempty"`
	TraceObservationPoint string          `json:"trace_observation_point,omitempty"`
	DropReasonDesc        int             `json:"drop_reason_desc,omitempty"`
	IsReply               bool            `json:"is_reply,omitempty"`
	DebugCapturePoint     int             `json:"debug_capture_point,omitempty"`
	Interface             int             `json:"interface,omitempty"`
	ProxyPort             int             `json:"proxy_port,omitempty"`
	TraceContext          int             `json:"trace_context,omitempty"`
	SockXLatePoint        int             `json:"sock_xlate_point,omitempty"`
	SocketCookie          int             `json:"socket_cookie,omitempty"`
	CroupID               int             `json:"cgroup_id,omitempty"`
	Summary               string          `json:"Summary,omitempty"`
}

// Cilium - Structure for Hubble Log Flow
type CiliumLog struct {
	Verdict                     string `json:"verdict,omitempty"`
	IpSource                    string `json:"ip_source,omitempty"`
	IpDestination               string `json:"ip_destination,omitempty"`
	IpVersion                   string `json:"ip_version,omitempty"`
	IpEncrypted                 bool   `json:"ip_encrypted,omitempty"`
	L4TCPSourcePort             uint32 `json:"l4_tcp_source_port,omitempty"`
	L4TCPDestinationPort        uint32 `json:"l4_tcp_destination_port,omitempty"`
	L4UDPSourcePort             uint32 `json:"l4_udp_source_port,omitempty"`
	L4UDPDestinationPort        uint32 `json:"l4_udp_destination_port,omitempty"`
	L4ICMPv4Type                uint32 `json:"l4_icmpv4_type,omitempty"`
	L4ICMPv4Code                uint32 `json:"l4_icmpv4_code,omitempty"`
	L4ICMPv6Type                uint32 `json:"l4_icmpv6_type,omitempty"`
	L4ICMPv6Code                uint32 `json:"l4_icmpv6_code,omitempty"`
	SourceNamespace             string `json:"source_namespace,omitempty"`
	SourceLabels                string `json:"source_labels,omitempty"`
	SourcePodName               string `json:"source_pod_name,omitempty"`
	DestinationNamespace        string `json:"destination_namespace,omitempty"`
	DestinationLabels           string `json:"destination_labels,omitempty"`
	DestinationPodName          string `json:"destination_pod_name,omitempty"`
	Type                        string `json:"type,omitempty"`
	NodeName                    string `json:"node_name,omitempty"`
	L7Type                      string `json:"l7_type,omitempty"`
	L7DnsCnames                 string `json:"l7_dns_cnames,omitempty"`
	L7DnsObservationsource      string `json:"l7_dns_observation_source,omitempty"`
	L7HttpCode                  uint32 `json:"l7_http_code,omitempty"`
	L7HttpMethod                string `json:"l7_http_method,omitempty"`
	L7HttpUrl                   string `json:"l7_http_url,omitempty"`
	L7HttpProtocol              string `json:"l7_http_protocol,omitempty"`
	L7HttpHeaders               string `json:"l7_http_headers,omitempty"`
	EventTypeType               int32  `json:"event_type_type,omitempty"`
	EventTypeSubType            int32  `json:"event_type_sub_type,omitempty"`
	SourceServiceName           string `json:"source_service_name,omitempty"`
	SourceServiceNamespace      string `json:"source_service_namespace,omitempty"`
	DestinationServiceName      string `json:"destination_service_name,omitempty"`
	DestinationServiceNamespace string `json:"destination_service_namespace,omitempty"`
	TrafficDirection            string `json:"traffic_direction,omitempty"`
	TraceObservationPoint       string `json:"trace_observation_point,omitempty"`
	DropReasonDesc              string `json:"drop_reason_desc,omitempty"`
	IsReply                     bool   `json:"is_reply,omitempty"`
	StartTime                   int64  `json:"start_time,omitempty"`
	UpdatedTime                 int64  `json:"updated_time,omitempty"`
	Total                       int64  `json:"total,omitempty"`
}
