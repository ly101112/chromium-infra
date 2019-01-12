// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/qscheduler/qslib/scheduler/worker.proto

package scheduler

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Worker represents a resource that can run 1 task at a time. This corresponds
// to the swarming concept of a Bot. This representation considers only the
// subset of Labels that are Provisionable (can be changed by running a task),
// because the quota scheduler algorithm is expected to run against a pool of
// otherwise homogenous workers.
type Worker struct {
	// Labels represents the set of labels that this worker possesses.
	Labels []string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// RunningTask is, if non-nil, the task that is currently running on the
	// worker.
	RunningTask *TaskRun `protobuf:"bytes,2,opt,name=running_task,json=runningTask,proto3" json:"running_task,omitempty"`
	// ConfirmedTime is the most recent time at which the Worker state was
	// directly confirmed as idle by external authority (via a call to MarkIdle or
	// NotifyRequest).
	ConfirmedTime        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=confirmed_time,json=confirmedTime,proto3" json:"confirmed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_9853fc992bf091d9, []int{0}
}

func (m *Worker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker.Unmarshal(m, b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
}
func (m *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(m, src)
}
func (m *Worker) XXX_Size() int {
	return xxx_messageInfo_Worker.Size(m)
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Worker) GetRunningTask() *TaskRun {
	if m != nil {
		return m.RunningTask
	}
	return nil
}

func (m *Worker) GetConfirmedTime() *timestamp.Timestamp {
	if m != nil {
		return m.ConfirmedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Worker)(nil), "scheduler.Worker")
}

func init() {
	proto.RegisterFile("infra/qscheduler/qslib/scheduler/worker.proto", fileDescriptor_9853fc992bf091d9)
}

var fileDescriptor_9853fc992bf091d9 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xdd, 0x4a, 0x80, 0x30,
	0x1c, 0xc5, 0x31, 0x41, 0x70, 0x56, 0x17, 0xbb, 0x08, 0xf1, 0x26, 0xe9, 0x4a, 0x88, 0x36, 0x28,
	0x7a, 0x80, 0x5e, 0x61, 0x08, 0x5d, 0xca, 0xa6, 0xd3, 0x86, 0x73, 0xd3, 0x7d, 0xd0, 0xbb, 0xf4,
	0xb4, 0xb1, 0xf9, 0xd1, 0x65, 0x97, 0xe7, 0xec, 0x77, 0xce, 0xfe, 0x07, 0xbc, 0x08, 0x35, 0x1a,
	0x8a, 0x37, 0xdb, 0x7f, 0xf1, 0xc1, 0x4b, 0x6e, 0xf0, 0x66, 0xa5, 0x60, 0xf8, 0x4f, 0x7f, 0x6b,
	0x33, 0x73, 0x83, 0x56, 0xa3, 0x9d, 0x86, 0xf9, 0xe5, 0x57, 0x8f, 0x93, 0xd6, 0x93, 0xe4, 0x38,
	0x3e, 0x30, 0x3f, 0x62, 0x27, 0x16, 0x6e, 0x1d, 0x5d, 0xd6, 0x9d, 0xad, 0x9e, 0xff, 0xad, 0x76,
	0xd4, 0xce, 0x3b, 0xfc, 0xf4, 0x93, 0x80, 0xec, 0x33, 0xfe, 0x04, 0x1f, 0x40, 0x26, 0x29, 0xe3,
	0xd2, 0x96, 0x49, 0x9d, 0x36, 0x39, 0x39, 0x14, 0x7c, 0x07, 0xb7, 0xc6, 0x2b, 0x25, 0xd4, 0xd4,
	0x85, 0x60, 0x79, 0x53, 0x27, 0x4d, 0xf1, 0x0a, 0xd1, 0xd5, 0x87, 0x5a, 0x6a, 0x67, 0xe2, 0x15,
	0x29, 0x0e, 0x2e, 0x68, 0xf8, 0x01, 0xee, 0x7b, 0xad, 0x46, 0x61, 0x16, 0x3e, 0x74, 0xe1, 0xc6,
	0x32, 0x8d, 0xc1, 0x0a, 0xed, 0x03, 0xd0, 0x39, 0x00, 0xb5, 0xe7, 0x00, 0x72, 0x77, 0x25, 0x82,
	0xc7, 0xb2, 0x88, 0xbc, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x21, 0xdc, 0xa2, 0x8a, 0x2d, 0x01,
	0x00, 0x00,
}
