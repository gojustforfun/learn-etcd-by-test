// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: trip.proto

package trippb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_NOT_SPECIFIED Status = 0
	Status_NOT_STARTED   Status = 1
	Status_IN_PROGRESS   Status = 2
	Status_FINISHED      Status = 3
	Status_PAID          Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "NOT_SPECIFIED",
		1: "NOT_STARTED",
		2: "IN_PROGRESS",
		3: "FINISHED",
		4: "PAID",
	}
	Status_value = map[string]int32{
		"NOT_SPECIFIED": 0,
		"NOT_STARTED":   1,
		"IN_PROGRESS":   2,
		"FINISHED":      3,
		"PAID":          4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_trip_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_trip_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// 不是按照声明位置,而是按照“=”右侧的数字排序
	StartPosition *Location   `protobuf:"bytes,5,opt,name=start_position,json=startPosition,proto3" json:"start_position,omitempty"`
	PathPositions []*Location `protobuf:"bytes,7,rep,name=path_positions,json=pathPositions,proto3" json:"path_positions,omitempty"`
	End           string      `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	EndPosition   *Location   `protobuf:"bytes,6,opt,name=end_position,json=endPosition,proto3" json:"end_position,omitempty"`
	DurationInSec int64       `protobuf:"varint,3,opt,name=duration_in_sec,json=durationInSec,proto3" json:"duration_in_sec,omitempty"`
	FeeInCent     int64       `protobuf:"varint,4,opt,name=fee_in_cent,json=feeInCent,proto3" json:"fee_in_cent,omitempty"`
	Status        Status      `protobuf:"varint,8,opt,name=status,proto3,enum=trip.Status" json:"status,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

func (x *Trip) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Trip) GetStartPosition() *Location {
	if x != nil {
		return x.StartPosition
	}
	return nil
}

func (x *Trip) GetPathPositions() []*Location {
	if x != nil {
		return x.PathPositions
	}
	return nil
}

func (x *Trip) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *Trip) GetEndPosition() *Location {
	if x != nil {
		return x.EndPosition
	}
	return nil
}

func (x *Trip) GetDurationInSec() int64 {
	if x != nil {
		return x.DurationInSec
	}
	return 0
}

func (x *Trip) GetFeeInCent() int64 {
	if x != nil {
		return x.FeeInCent
	}
	return 0
}

func (x *Trip) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_NOT_SPECIFIED
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GetTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTripRequest) Reset() {
	*x = GetTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRequest) ProtoMessage() {}

func (x *GetTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRequest.ProtoReflect.Descriptor instead.
func (*GetTripRequest) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{2}
}

func (x *GetTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Trip *Trip  `protobuf:"bytes,2,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *GetTripResponse) Reset() {
	*x = GetTripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripResponse) ProtoMessage() {}

func (x *GetTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripResponse.ProtoReflect.Descriptor instead.
func (*GetTripResponse) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{3}
}

func (x *GetTripResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTripResponse) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

var File_trip_proto protoreflect.FileDescriptor

var file_trip_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x72,
	0x69, 0x70, 0x22, 0xbd, 0x02, 0x0a, 0x04, 0x54, 0x72, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x35, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0e, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x70, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x31, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x1e, 0x0a, 0x0b,
	0x66, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x66, 0x65, 0x65, 0x49, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74,
	0x72, 0x69, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x72,
	0x69, 0x70, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70, 0x2a, 0x55, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41,
	0x49, 0x44, 0x10, 0x04, 0x32, 0x45, 0x0a, 0x0b, 0x54, 0x72, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x14,
	0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trip_proto_rawDescOnce sync.Once
	file_trip_proto_rawDescData = file_trip_proto_rawDesc
)

func file_trip_proto_rawDescGZIP() []byte {
	file_trip_proto_rawDescOnce.Do(func() {
		file_trip_proto_rawDescData = protoimpl.X.CompressGZIP(file_trip_proto_rawDescData)
	})
	return file_trip_proto_rawDescData
}

var file_trip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trip_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_trip_proto_goTypes = []interface{}{
	(Status)(0),             // 0: trip.Status
	(*Trip)(nil),            // 1: trip.Trip
	(*Location)(nil),        // 2: trip.Location
	(*GetTripRequest)(nil),  // 3: trip.GetTripRequest
	(*GetTripResponse)(nil), // 4: trip.GetTripResponse
}
var file_trip_proto_depIdxs = []int32{
	2, // 0: trip.Trip.start_position:type_name -> trip.Location
	2, // 1: trip.Trip.path_positions:type_name -> trip.Location
	2, // 2: trip.Trip.end_position:type_name -> trip.Location
	0, // 3: trip.Trip.status:type_name -> trip.Status
	1, // 4: trip.GetTripResponse.trip:type_name -> trip.Trip
	3, // 5: trip.TripService.GetTrip:input_type -> trip.GetTripRequest
	4, // 6: trip.TripService.GetTrip:output_type -> trip.GetTripResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_trip_proto_init() }
func file_trip_proto_init() {
	if File_trip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trip); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trip_proto_goTypes,
		DependencyIndexes: file_trip_proto_depIdxs,
		EnumInfos:         file_trip_proto_enumTypes,
		MessageInfos:      file_trip_proto_msgTypes,
	}.Build()
	File_trip_proto = out.File
	file_trip_proto_rawDesc = nil
	file_trip_proto_goTypes = nil
	file_trip_proto_depIdxs = nil
}
