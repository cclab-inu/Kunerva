package types

import "encoding/json"

// NetworkLog Structure
type NetworkLog struct {
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

// NetworkLogRaw
type NetworkLogRaw struct {
	ID                    int             `json:"id,omitempty"`          //
	Time                  int64           `json:"time,omitempty"`        //
	Verdict               int32           `json:"verdict,omitempty"`     //
	DropReason            int32           `json:"drop_reason,omitempty"` //
	Ethernet              json.RawMessage `json:"ethernet,omitempty"`
	IP                    json.RawMessage `json:"IP,omitempty"`          //
	L4                    json.RawMessage `json:"l4,omitempty"`          //
	Source                json.RawMessage `json:"source,omitempty"`      //
	Destination           json.RawMessage `json:"destination,omitempty"` //
	Type                  int32           `json:"Type,omitempty"`
	NodeName              string          `json:"node_name,omitempty"`
	SourceNames           []string        `json:"source_names,omitempty"`
	DestinationNames      []string        `json:"destination_names,omitempty"`
	L7                    json.RawMessage `json:"l7,omitempty"` //
	Reply                 bool            `json:"reply,omitempty"`
	EventType             json.RawMessage `json:"event_type,omitempty"` //
	SourceService         json.RawMessage `json:"source_service,omitempty"`
	DestinationService    json.RawMessage `json:"destination_service,omitempty"`
	TrafficDirection      uint32          `json:"traffic_direction,omitempty"` //
	PolicyMatchType       uint32          `json:"policy_match_type,omitempty"` //
	TraceObservationPoint string          `json:"trace_observation_point,omitempty"`
	DropReasonDesc        int             `json:"drop_reason_desc,omitempty"`
	IsReply               bool            `json:"is_reply,omitempty"`
	DebugCapturePoint     int             `json:"debug_capture_point,omitempty"`
	Interface             json.RawMessage `json:"interface,omitempty"`
	ProxyPort             int             `json:"proxy_port,omitempty"`
	TraceContext          json.RawMessage `json:"trace_context,omitempty"`
	SockXLatePoint        int             `json:"sock_xlate_point,omitempty"`
	SocketCookie          int             `json:"socket_cookie,omitempty"`
	CroupID               int             `json:"cgroup_id,omitempty"`
	Summary               string          `json:"Summary,omitempty"`
}
