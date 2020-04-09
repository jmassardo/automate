// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/compliance/ingest/events/compliance/compliance.proto

package compliance

import (
	fmt "fmt"
	common "github.com/chef/automate/api/interservice/compliance/common"
	inspec "github.com/chef/automate/api/interservice/compliance/ingest/events/inspec"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Report struct {
	// inspec full json report fields
	Version     string             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	Platform    *inspec.Platform   `protobuf:"bytes,16,opt,name=platform,proto3" json:"platform,omitempty" toml:"platform,omitempty" mapstructure:"platform,omitempty"`
	Statistics  *inspec.Statistics `protobuf:"bytes,17,opt,name=statistics,proto3" json:"statistics,omitempty" toml:"statistics,omitempty" mapstructure:"statistics,omitempty"`
	Profiles    []*inspec.Profile  `protobuf:"bytes,18,rep,name=profiles,proto3" json:"profiles,omitempty" toml:"profiles,omitempty" mapstructure:"profiles,omitempty"`
	OtherChecks []string           `protobuf:"bytes,19,rep,name=other_checks,json=otherChecks,proto3" json:"other_checks,omitempty" toml:"other_checks,omitempty" mapstructure:"other_checks,omitempty"`
	// extra report fields added by the audit cookbook
	ReportUuid           string       `protobuf:"bytes,20,opt,name=report_uuid,json=reportUuid,proto3" json:"report_uuid,omitempty" toml:"report_uuid,omitempty" mapstructure:"report_uuid,omitempty"`
	NodeUuid             string       `protobuf:"bytes,21,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty" toml:"node_uuid,omitempty" mapstructure:"node_uuid,omitempty"`
	JobUuid              string       `protobuf:"bytes,22,opt,name=job_uuid,json=jobUuid,proto3" json:"job_uuid,omitempty" toml:"job_uuid,omitempty" mapstructure:"job_uuid,omitempty"`
	NodeName             string       `protobuf:"bytes,23,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty" toml:"node_name,omitempty" mapstructure:"node_name,omitempty"`
	Environment          string       `protobuf:"bytes,24,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	Roles                []string     `protobuf:"bytes,25,rep,name=roles,proto3" json:"roles,omitempty" toml:"roles,omitempty" mapstructure:"roles,omitempty"`
	Recipes              []string     `protobuf:"bytes,26,rep,name=recipes,proto3" json:"recipes,omitempty" toml:"recipes,omitempty" mapstructure:"recipes,omitempty"`
	EndTime              string       `protobuf:"bytes,27,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" toml:"end_time,omitempty" mapstructure:"end_time,omitempty"`
	Type                 string       `protobuf:"bytes,28,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty" mapstructure:"type,omitempty"`
	SourceId             string       `protobuf:"bytes,29,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty" toml:"source_id,omitempty" mapstructure:"source_id,omitempty"`
	SourceRegion         string       `protobuf:"bytes,30,opt,name=source_region,json=sourceRegion,proto3" json:"source_region,omitempty" toml:"source_region,omitempty" mapstructure:"source_region,omitempty"`
	SourceAccountId      string       `protobuf:"bytes,31,opt,name=source_account_id,json=sourceAccountId,proto3" json:"source_account_id,omitempty" toml:"source_account_id,omitempty" mapstructure:"source_account_id,omitempty"`
	PolicyName           string       `protobuf:"bytes,32,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	PolicyGroup          string       `protobuf:"bytes,33,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	OrganizationName     string       `protobuf:"bytes,34,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty" toml:"organization_name,omitempty" mapstructure:"organization_name,omitempty"`
	SourceFqdn           string       `protobuf:"bytes,35,opt,name=source_fqdn,json=sourceFqdn,proto3" json:"source_fqdn,omitempty" toml:"source_fqdn,omitempty" mapstructure:"source_fqdn,omitempty"`
	ChefTags             []string     `protobuf:"bytes,36,rep,name=chef_tags,json=chefTags,proto3" json:"chef_tags,omitempty" toml:"chef_tags,omitempty" mapstructure:"chef_tags,omitempty"`
	Ipaddress            string       `protobuf:"bytes,37,opt,name=ipaddress,proto3" json:"ipaddress,omitempty" toml:"ipaddress,omitempty" mapstructure:"ipaddress,omitempty"`
	Fqdn                 string       `protobuf:"bytes,38,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	Tags                 []*common.Kv `protobuf:"bytes,39,rep,name=tags,proto3" json:"tags,omitempty" toml:"tags,omitempty" mapstructure:"tags,omitempty"`
	AutomateManagerId    string       `protobuf:"bytes,40,opt,name=automate_manager_id,json=automateManagerId,proto3" json:"automate_manager_id,omitempty" toml:"automate_manager_id,omitempty" mapstructure:"automate_manager_id,omitempty"`
	RunTimeLimit         float32      `protobuf:"fixed32,41,opt,name=run_time_limit,json=runTimeLimit,proto3" json:"run_time_limit,omitempty" toml:"run_time_limit,omitempty" mapstructure:"run_time_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32        `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d8e0277b12342b3, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Report) GetPlatform() *inspec.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func (m *Report) GetStatistics() *inspec.Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *Report) GetProfiles() []*inspec.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Report) GetOtherChecks() []string {
	if m != nil {
		return m.OtherChecks
	}
	return nil
}

func (m *Report) GetReportUuid() string {
	if m != nil {
		return m.ReportUuid
	}
	return ""
}

func (m *Report) GetNodeUuid() string {
	if m != nil {
		return m.NodeUuid
	}
	return ""
}

func (m *Report) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

func (m *Report) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Report) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *Report) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Report) GetRecipes() []string {
	if m != nil {
		return m.Recipes
	}
	return nil
}

func (m *Report) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Report) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Report) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Report) GetSourceRegion() string {
	if m != nil {
		return m.SourceRegion
	}
	return ""
}

func (m *Report) GetSourceAccountId() string {
	if m != nil {
		return m.SourceAccountId
	}
	return ""
}

func (m *Report) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Report) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Report) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Report) GetSourceFqdn() string {
	if m != nil {
		return m.SourceFqdn
	}
	return ""
}

func (m *Report) GetChefTags() []string {
	if m != nil {
		return m.ChefTags
	}
	return nil
}

func (m *Report) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *Report) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Report) GetTags() []*common.Kv {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Report) GetAutomateManagerId() string {
	if m != nil {
		return m.AutomateManagerId
	}
	return ""
}

func (m *Report) GetRunTimeLimit() float32 {
	if m != nil {
		return m.RunTimeLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*Report)(nil), "chef.automate.domain.compliance.ingest.events.compliance.Report")
}

func init() {
	proto.RegisterFile("api/interservice/compliance/ingest/events/compliance/compliance.proto", fileDescriptor_2d8e0277b12342b3)
}

var fileDescriptor_2d8e0277b12342b3 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0x13, 0x3d,
	0x14, 0x55, 0xbe, 0xf6, 0x2b, 0x89, 0x13, 0xa0, 0x71, 0x0b, 0xb8, 0x0f, 0x68, 0xfa, 0x00, 0x42,
	0x91, 0x32, 0x52, 0xd9, 0xb0, 0x69, 0xc5, 0x43, 0x80, 0x2a, 0x28, 0xa0, 0xa1, 0x2c, 0xe8, 0x26,
	0x38, 0xf6, 0xcd, 0xc4, 0x25, 0x63, 0x4f, 0x6d, 0x4f, 0xa4, 0xf2, 0xd3, 0xf8, 0x75, 0xc8, 0xd7,
	0x93, 0x36, 0x6c, 0x78, 0x74, 0x35, 0xf6, 0x39, 0xf7, 0x1e, 0x9f, 0x39, 0xbe, 0x33, 0xe4, 0x15,
	0x2f, 0x54, 0xa2, 0xb4, 0x07, 0xeb, 0xc0, 0x4e, 0x94, 0x80, 0x44, 0x98, 0xbc, 0x18, 0x2b, 0xae,
	0x05, 0x24, 0x4a, 0x67, 0xe0, 0x7c, 0x02, 0x13, 0xd0, 0xde, 0xcd, 0x12, 0x97, 0xcb, 0x5e, 0x61,
	0x8d, 0x37, 0xf4, 0xa9, 0x18, 0xc1, 0xb0, 0xc7, 0x4b, 0x6f, 0x72, 0xee, 0xa1, 0x27, 0x4d, 0xce,
	0x95, 0xee, 0xcd, 0x94, 0x45, 0xa9, 0x5e, 0x94, 0x9a, 0x21, 0x56, 0xf7, 0xff, 0xde, 0x80, 0xd2,
	0xae, 0x00, 0x51, 0x3d, 0xe2, 0xc1, 0xab, 0xc9, 0xef, 0xda, 0x85, 0xc9, 0x73, 0xa3, 0xab, 0x47,
	0x6c, 0xd8, 0xfa, 0x51, 0x27, 0x0b, 0x29, 0x14, 0xc6, 0x7a, 0xca, 0xc8, 0xb5, 0x09, 0x58, 0xa7,
	0x8c, 0x66, 0xb5, 0x4e, 0xad, 0xdb, 0x48, 0xa7, 0x5b, 0x7a, 0x42, 0xea, 0xc5, 0x98, 0xfb, 0xa1,
	0xb1, 0x39, 0x5b, 0xec, 0xd4, 0xba, 0xcd, 0xbd, 0x83, 0xde, 0xbf, 0xbd, 0x61, 0x65, 0xf2, 0x63,
	0xa5, 0x92, 0x5e, 0xe8, 0xd1, 0xaf, 0x84, 0x38, 0xcf, 0xbd, 0x72, 0x5e, 0x09, 0xc7, 0xda, 0xa8,
	0xfe, 0xec, 0x6a, 0xea, 0x9f, 0x2e, 0x74, 0xd2, 0x19, 0x4d, 0xfa, 0x85, 0xd4, 0x0b, 0x6b, 0x86,
	0x6a, 0x0c, 0x8e, 0xd1, 0xce, 0x5c, 0xb7, 0xb9, 0xb7, 0x7f, 0x45, 0xf7, 0x51, 0x25, 0xbd, 0x90,
	0xa3, 0x9b, 0xa4, 0x65, 0xfc, 0x08, 0x6c, 0x5f, 0x8c, 0x40, 0x7c, 0x73, 0x6c, 0xa9, 0x33, 0xd7,
	0x6d, 0xa4, 0x4d, 0xc4, 0x5e, 0x22, 0x44, 0x37, 0x48, 0xd3, 0x62, 0xbe, 0xfd, 0xb2, 0x54, 0x92,
	0x2d, 0x63, 0xb2, 0x24, 0x42, 0x9f, 0x4b, 0x25, 0xe9, 0x1a, 0x69, 0x68, 0x23, 0x21, 0xd2, 0xb7,
	0x90, 0xae, 0x07, 0x00, 0xc9, 0x15, 0x52, 0x3f, 0x35, 0x83, 0xc8, 0xdd, 0x8e, 0x97, 0x72, 0x6a,
	0x06, 0xbf, 0xf4, 0x69, 0x9e, 0x03, 0xbb, 0x73, 0xd9, 0xf7, 0x9e, 0xe7, 0x40, 0x3b, 0xa4, 0x09,
	0x7a, 0xa2, 0xac, 0xd1, 0x39, 0x68, 0xcf, 0x18, 0xd2, 0xb3, 0x10, 0x5d, 0x26, 0xff, 0x5b, 0x13,
	0x22, 0x59, 0x41, 0xcf, 0x71, 0x13, 0x66, 0xc0, 0x82, 0x50, 0x05, 0x38, 0xb6, 0x8a, 0xf8, 0x74,
	0x1b, 0x9c, 0x80, 0x96, 0x7d, 0xaf, 0x72, 0x60, 0x6b, 0xd1, 0x09, 0x68, 0x79, 0xac, 0x72, 0xa0,
	0x94, 0xcc, 0xfb, 0xf3, 0x02, 0xd8, 0x3a, 0xc2, 0xb8, 0x0e, 0xee, 0x9c, 0x29, 0xad, 0x80, 0xbe,
	0x92, 0xec, 0x6e, 0x74, 0x17, 0x81, 0x43, 0x49, 0xb7, 0xc9, 0xf5, 0x8a, 0xb4, 0x90, 0x85, 0x79,
	0xbb, 0x87, 0x05, 0xad, 0x08, 0xa6, 0x88, 0xd1, 0x5d, 0xd2, 0xae, 0x8a, 0xb8, 0x10, 0xa6, 0xd4,
	0x3e, 0x28, 0x6d, 0x60, 0xe1, 0xcd, 0x48, 0x3c, 0x8f, 0xf8, 0xa1, 0x0c, 0x21, 0x17, 0x66, 0xac,
	0xc4, 0x79, 0x4c, 0xa3, 0x13, 0x43, 0x8e, 0x10, 0xe6, 0xb1, 0x49, 0x5a, 0x55, 0x41, 0x66, 0x4d,
	0x59, 0xb0, 0xcd, 0x18, 0x48, 0xc4, 0xde, 0x04, 0x88, 0x3e, 0x26, 0x6d, 0x63, 0x33, 0xae, 0xd5,
	0x77, 0xee, 0x95, 0xd1, 0x51, 0x69, 0x0b, 0xeb, 0x16, 0x67, 0x09, 0xd4, 0xdb, 0x20, 0xcd, 0xca,
	0xdc, 0xf0, 0x4c, 0x6a, 0xb6, 0x1d, 0x0f, 0x8c, 0xd0, 0xeb, 0x33, 0xa9, 0xc3, 0xfb, 0x87, 0x19,
	0xeb, 0x7b, 0x9e, 0x39, 0xb6, 0x83, 0x51, 0xd6, 0x03, 0x70, 0xcc, 0x33, 0x47, 0xd7, 0x49, 0x43,
	0x15, 0x5c, 0x4a, 0x0b, 0xce, 0xb1, 0xfb, 0xd8, 0x7b, 0x09, 0x84, 0x38, 0x51, 0xf4, 0x41, 0x8c,
	0x33, 0xac, 0xe9, 0x01, 0x99, 0x47, 0xa5, 0x87, 0x38, 0xbf, 0xbb, 0x7f, 0x9c, 0xdf, 0xea, 0x1b,
	0x7f, 0x3b, 0x49, 0xb1, 0x8f, 0xf6, 0xc8, 0xd2, 0xb4, 0xba, 0x9f, 0x73, 0xcd, 0x33, 0xb0, 0x21,
	0xce, 0x2e, 0x1e, 0xd1, 0x9e, 0x52, 0x47, 0x91, 0x39, 0x94, 0x74, 0x87, 0xdc, 0xb0, 0xa5, 0xc6,
	0xdb, 0xee, 0x8f, 0x55, 0xae, 0x3c, 0x7b, 0xd4, 0xa9, 0x75, 0xff, 0x4b, 0x5b, 0xb6, 0xd4, 0xe1,
	0xce, 0xdf, 0x05, 0xec, 0xc5, 0x87, 0x93, 0xa3, 0x4c, 0xf9, 0x51, 0x39, 0x08, 0xe7, 0x25, 0xc1,
	0x53, 0x32, 0x95, 0x4a, 0xae, 0xf2, 0x23, 0x1d, 0x2c, 0xe0, 0x4f, 0xe9, 0xc9, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfd, 0xe5, 0x1b, 0xd2, 0x87, 0x05, 0x00, 0x00,
}
