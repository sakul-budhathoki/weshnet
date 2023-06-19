// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode/errcode.proto

package errcode

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                                            ErrCode = 0
	TODO                                                 ErrCode = 666
	ErrNotImplemented                                    ErrCode = 777
	ErrInternal                                          ErrCode = 888
	ErrInvalidInput                                      ErrCode = 100
	ErrInvalidRange                                      ErrCode = 101
	ErrMissingInput                                      ErrCode = 102
	ErrSerialization                                     ErrCode = 103
	ErrDeserialization                                   ErrCode = 104
	ErrStreamRead                                        ErrCode = 105
	ErrStreamWrite                                       ErrCode = 106
	ErrStreamTransform                                   ErrCode = 110
	ErrStreamSendAndClose                                ErrCode = 111
	ErrStreamHeaderWrite                                 ErrCode = 112
	ErrStreamHeaderRead                                  ErrCode = 115
	ErrStreamSink                                        ErrCode = 113
	ErrStreamCloseAndRecv                                ErrCode = 114
	ErrMissingMapKey                                     ErrCode = 107
	ErrDBWrite                                           ErrCode = 108
	ErrDBRead                                            ErrCode = 109
	ErrDBDestroy                                         ErrCode = 120
	ErrDBMigrate                                         ErrCode = 121
	ErrDBReplay                                          ErrCode = 122
	ErrDBRestore                                         ErrCode = 123
	ErrDBOpen                                            ErrCode = 124
	ErrDBClose                                           ErrCode = 125
	ErrCryptoRandomGeneration                            ErrCode = 200
	ErrCryptoKeyGeneration                               ErrCode = 201
	ErrCryptoNonceGeneration                             ErrCode = 202
	ErrCryptoSignature                                   ErrCode = 203
	ErrCryptoSignatureVerification                       ErrCode = 204
	ErrCryptoDecrypt                                     ErrCode = 205
	ErrCryptoDecryptPayload                              ErrCode = 206
	ErrCryptoEncrypt                                     ErrCode = 207
	ErrCryptoKeyConversion                               ErrCode = 208
	ErrCryptoCipherInit                                  ErrCode = 209
	ErrCryptoKeyDerivation                               ErrCode = 210
	ErrMap                                               ErrCode = 300
	ErrForEach                                           ErrCode = 301
	ErrKeystoreGet                                       ErrCode = 400
	ErrKeystorePut                                       ErrCode = 401
	ErrNotFound                                          ErrCode = 404
	ErrOrbitDBInit                                       ErrCode = 1000
	ErrOrbitDBOpen                                       ErrCode = 1001
	ErrOrbitDBAppend                                     ErrCode = 1002
	ErrOrbitDBDeserialization                            ErrCode = 1003
	ErrOrbitDBStoreCast                                  ErrCode = 1004
	ErrHandshakeOwnEphemeralKeyGenSend                   ErrCode = 1100
	ErrHandshakePeerEphemeralKeyRecv                     ErrCode = 1101
	ErrHandshakeRequesterAuthenticateBoxKeyGen           ErrCode = 1102
	ErrHandshakeResponderAcceptBoxKeyGen                 ErrCode = 1103
	ErrHandshakeRequesterHello                           ErrCode = 1104
	ErrHandshakeResponderHello                           ErrCode = 1105
	ErrHandshakeRequesterAuthenticate                    ErrCode = 1106
	ErrHandshakeResponderAccept                          ErrCode = 1107
	ErrHandshakeRequesterAcknowledge                     ErrCode = 1108
	ErrContactRequestSameAccount                         ErrCode = 1200
	ErrContactRequestContactAlreadyAdded                 ErrCode = 1201
	ErrContactRequestContactBlocked                      ErrCode = 1202
	ErrContactRequestContactUndefined                    ErrCode = 1203
	ErrContactRequestIncomingAlreadyReceived             ErrCode = 1204
	ErrGroupMemberLogEventOpen                           ErrCode = 1300
	ErrGroupMemberLogEventSignature                      ErrCode = 1301
	ErrGroupMemberUnknownGroupID                         ErrCode = 1302
	ErrGroupSecretOtherDestMember                        ErrCode = 1303
	ErrGroupSecretAlreadySentToMember                    ErrCode = 1304
	ErrGroupInvalidType                                  ErrCode = 1305
	ErrGroupMissing                                      ErrCode = 1306
	ErrGroupActivate                                     ErrCode = 1307
	ErrGroupDeactivate                                   ErrCode = 1308
	ErrGroupInfo                                         ErrCode = 1309
	ErrGroupUnknown                                      ErrCode = 1310
	ErrGroupOpen                                         ErrCode = 1311
	ErrMessageKeyPersistencePut                          ErrCode = 1500
	ErrMessageKeyPersistenceGet                          ErrCode = 1501
	ErrServiceReplication                                ErrCode = 4100
	ErrServiceReplicationServer                          ErrCode = 4101
	ErrServiceReplicationMissingEndpoint                 ErrCode = 4102
	ErrServicesDirectory                                 ErrCode = 4200
	ErrServicesDirectoryInvalidVerifiedCredentialSubject ErrCode = 4201
	ErrServicesDirectoryExistingRecordNotFound           ErrCode = 4202
	ErrServicesDirectoryRecordLockedAndCantBeReplaced    ErrCode = 4203
	ErrServicesDirectoryExplicitReplaceFlagRequired      ErrCode = 4204
	ErrServicesDirectoryInvalidVerifiedCredential        ErrCode = 4205
	ErrServicesDirectoryExpiredVerifiedCredential        ErrCode = 4206
	ErrServicesDirectoryInvalidVerifiedCredentialID      ErrCode = 4207
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrInternal",
	100:  "ErrInvalidInput",
	101:  "ErrInvalidRange",
	102:  "ErrMissingInput",
	103:  "ErrSerialization",
	104:  "ErrDeserialization",
	105:  "ErrStreamRead",
	106:  "ErrStreamWrite",
	110:  "ErrStreamTransform",
	111:  "ErrStreamSendAndClose",
	112:  "ErrStreamHeaderWrite",
	115:  "ErrStreamHeaderRead",
	113:  "ErrStreamSink",
	114:  "ErrStreamCloseAndRecv",
	107:  "ErrMissingMapKey",
	108:  "ErrDBWrite",
	109:  "ErrDBRead",
	120:  "ErrDBDestroy",
	121:  "ErrDBMigrate",
	122:  "ErrDBReplay",
	123:  "ErrDBRestore",
	124:  "ErrDBOpen",
	125:  "ErrDBClose",
	200:  "ErrCryptoRandomGeneration",
	201:  "ErrCryptoKeyGeneration",
	202:  "ErrCryptoNonceGeneration",
	203:  "ErrCryptoSignature",
	204:  "ErrCryptoSignatureVerification",
	205:  "ErrCryptoDecrypt",
	206:  "ErrCryptoDecryptPayload",
	207:  "ErrCryptoEncrypt",
	208:  "ErrCryptoKeyConversion",
	209:  "ErrCryptoCipherInit",
	210:  "ErrCryptoKeyDerivation",
	300:  "ErrMap",
	301:  "ErrForEach",
	400:  "ErrKeystoreGet",
	401:  "ErrKeystorePut",
	404:  "ErrNotFound",
	1000: "ErrOrbitDBInit",
	1001: "ErrOrbitDBOpen",
	1002: "ErrOrbitDBAppend",
	1003: "ErrOrbitDBDeserialization",
	1004: "ErrOrbitDBStoreCast",
	1100: "ErrHandshakeOwnEphemeralKeyGenSend",
	1101: "ErrHandshakePeerEphemeralKeyRecv",
	1102: "ErrHandshakeRequesterAuthenticateBoxKeyGen",
	1103: "ErrHandshakeResponderAcceptBoxKeyGen",
	1104: "ErrHandshakeRequesterHello",
	1105: "ErrHandshakeResponderHello",
	1106: "ErrHandshakeRequesterAuthenticate",
	1107: "ErrHandshakeResponderAccept",
	1108: "ErrHandshakeRequesterAcknowledge",
	1200: "ErrContactRequestSameAccount",
	1201: "ErrContactRequestContactAlreadyAdded",
	1202: "ErrContactRequestContactBlocked",
	1203: "ErrContactRequestContactUndefined",
	1204: "ErrContactRequestIncomingAlreadyReceived",
	1300: "ErrGroupMemberLogEventOpen",
	1301: "ErrGroupMemberLogEventSignature",
	1302: "ErrGroupMemberUnknownGroupID",
	1303: "ErrGroupSecretOtherDestMember",
	1304: "ErrGroupSecretAlreadySentToMember",
	1305: "ErrGroupInvalidType",
	1306: "ErrGroupMissing",
	1307: "ErrGroupActivate",
	1308: "ErrGroupDeactivate",
	1309: "ErrGroupInfo",
	1310: "ErrGroupUnknown",
	1311: "ErrGroupOpen",
	1500: "ErrMessageKeyPersistencePut",
	1501: "ErrMessageKeyPersistenceGet",
	4100: "ErrServiceReplication",
	4101: "ErrServiceReplicationServer",
	4102: "ErrServiceReplicationMissingEndpoint",
	4200: "ErrServicesDirectory",
	4201: "ErrServicesDirectoryInvalidVerifiedCredentialSubject",
	4202: "ErrServicesDirectoryExistingRecordNotFound",
	4203: "ErrServicesDirectoryRecordLockedAndCantBeReplaced",
	4204: "ErrServicesDirectoryExplicitReplaceFlagRequired",
	4205: "ErrServicesDirectoryInvalidVerifiedCredential",
	4206: "ErrServicesDirectoryExpiredVerifiedCredential",
	4207: "ErrServicesDirectoryInvalidVerifiedCredentialID",
}

var ErrCode_value = map[string]int32{
	"Undefined":                          0,
	"TODO":                               666,
	"ErrNotImplemented":                  777,
	"ErrInternal":                        888,
	"ErrInvalidInput":                    100,
	"ErrInvalidRange":                    101,
	"ErrMissingInput":                    102,
	"ErrSerialization":                   103,
	"ErrDeserialization":                 104,
	"ErrStreamRead":                      105,
	"ErrStreamWrite":                     106,
	"ErrStreamTransform":                 110,
	"ErrStreamSendAndClose":              111,
	"ErrStreamHeaderWrite":               112,
	"ErrStreamHeaderRead":                115,
	"ErrStreamSink":                      113,
	"ErrStreamCloseAndRecv":              114,
	"ErrMissingMapKey":                   107,
	"ErrDBWrite":                         108,
	"ErrDBRead":                          109,
	"ErrDBDestroy":                       120,
	"ErrDBMigrate":                       121,
	"ErrDBReplay":                        122,
	"ErrDBRestore":                       123,
	"ErrDBOpen":                          124,
	"ErrDBClose":                         125,
	"ErrCryptoRandomGeneration":          200,
	"ErrCryptoKeyGeneration":             201,
	"ErrCryptoNonceGeneration":           202,
	"ErrCryptoSignature":                 203,
	"ErrCryptoSignatureVerification":     204,
	"ErrCryptoDecrypt":                   205,
	"ErrCryptoDecryptPayload":            206,
	"ErrCryptoEncrypt":                   207,
	"ErrCryptoKeyConversion":             208,
	"ErrCryptoCipherInit":                209,
	"ErrCryptoKeyDerivation":             210,
	"ErrMap":                             300,
	"ErrForEach":                         301,
	"ErrKeystoreGet":                     400,
	"ErrKeystorePut":                     401,
	"ErrNotFound":                        404,
	"ErrOrbitDBInit":                     1000,
	"ErrOrbitDBOpen":                     1001,
	"ErrOrbitDBAppend":                   1002,
	"ErrOrbitDBDeserialization":          1003,
	"ErrOrbitDBStoreCast":                1004,
	"ErrHandshakeOwnEphemeralKeyGenSend": 1100,
	"ErrHandshakePeerEphemeralKeyRecv":   1101,
	"ErrHandshakeRequesterAuthenticateBoxKeyGen":           1102,
	"ErrHandshakeResponderAcceptBoxKeyGen":                 1103,
	"ErrHandshakeRequesterHello":                           1104,
	"ErrHandshakeResponderHello":                           1105,
	"ErrHandshakeRequesterAuthenticate":                    1106,
	"ErrHandshakeResponderAccept":                          1107,
	"ErrHandshakeRequesterAcknowledge":                     1108,
	"ErrContactRequestSameAccount":                         1200,
	"ErrContactRequestContactAlreadyAdded":                 1201,
	"ErrContactRequestContactBlocked":                      1202,
	"ErrContactRequestContactUndefined":                    1203,
	"ErrContactRequestIncomingAlreadyReceived":             1204,
	"ErrGroupMemberLogEventOpen":                           1300,
	"ErrGroupMemberLogEventSignature":                      1301,
	"ErrGroupMemberUnknownGroupID":                         1302,
	"ErrGroupSecretOtherDestMember":                        1303,
	"ErrGroupSecretAlreadySentToMember":                    1304,
	"ErrGroupInvalidType":                                  1305,
	"ErrGroupMissing":                                      1306,
	"ErrGroupActivate":                                     1307,
	"ErrGroupDeactivate":                                   1308,
	"ErrGroupInfo":                                         1309,
	"ErrGroupUnknown":                                      1310,
	"ErrGroupOpen":                                         1311,
	"ErrMessageKeyPersistencePut":                          1500,
	"ErrMessageKeyPersistenceGet":                          1501,
	"ErrServiceReplication":                                4100,
	"ErrServiceReplicationServer":                          4101,
	"ErrServiceReplicationMissingEndpoint":                 4102,
	"ErrServicesDirectory":                                 4200,
	"ErrServicesDirectoryInvalidVerifiedCredentialSubject": 4201,
	"ErrServicesDirectoryExistingRecordNotFound":           4202,
	"ErrServicesDirectoryRecordLockedAndCantBeReplaced":    4203,
	"ErrServicesDirectoryExplicitReplaceFlagRequired":      4204,
	"ErrServicesDirectoryInvalidVerifiedCredential":        4205,
	"ErrServicesDirectoryExpiredVerifiedCredential":        4206,
	"ErrServicesDirectoryInvalidVerifiedCredentialID":      4207,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb5abb189af31c1a, []int{0}
}

type ErrDetails struct {
	Codes                []ErrCode `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=weshnet.errcode.ErrCode" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrDetails) Reset()         { *m = ErrDetails{} }
func (m *ErrDetails) String() string { return proto.CompactTextString(m) }
func (*ErrDetails) ProtoMessage()    {}
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb5abb189af31c1a, []int{0}
}
func (m *ErrDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrDetails.Merge(m, src)
}
func (m *ErrDetails) XXX_Size() int {
	return m.Size()
}
func (m *ErrDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrDetails proto.InternalMessageInfo

func (m *ErrDetails) GetCodes() []ErrCode {
	if m != nil {
		return m.Codes
	}
	return nil
}

func init() {
	proto.RegisterEnum("weshnet.errcode.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*ErrDetails)(nil), "weshnet.errcode.ErrDetails")
}

func init() { proto.RegisterFile("errcode/errcode.proto", fileDescriptor_fb5abb189af31c1a) }

var fileDescriptor_fb5abb189af31c1a = []byte{
	// 1290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5b, 0x6f, 0x5b, 0xc5,
	0x16, 0xae, 0xdb, 0xd3, 0x5c, 0x26, 0xa7, 0xcd, 0xea, 0x24, 0x6d, 0xd3, 0x9b, 0xeb, 0xf6, 0xf4,
	0x9c, 0x53, 0x2a, 0x35, 0x81, 0x52, 0x21, 0x21, 0xf1, 0xe2, 0xc4, 0x6e, 0x6b, 0xa5, 0x69, 0x22,
	0x3b, 0x05, 0x89, 0xb7, 0xc9, 0x9e, 0x95, 0xed, 0x69, 0xb6, 0x67, 0x76, 0x67, 0x8f, 0xdd, 0xb8,
	0xc0, 0x1b, 0x20, 0xf1, 0x06, 0x52, 0xb9, 0x95, 0xdb, 0x1f, 0x00, 0x89, 0xdb, 0x8f, 0x28, 0xd0,
	0x4b, 0x5a, 0x78, 0x03, 0x24, 0xe8, 0x53, 0x6f, 0xc0, 0x2b, 0x8f, 0x68, 0xcf, 0x4c, 0x1c, 0x3b,
	0x76, 0x8a, 0x78, 0xb2, 0xf7, 0xb7, 0xbe, 0x75, 0x9f, 0x59, 0x6b, 0xc8, 0x4e, 0xd4, 0x3a, 0x50,
	0x1c, 0x27, 0xfc, 0xef, 0x78, 0xac, 0x95, 0x51, 0x74, 0xf8, 0x12, 0x26, 0x55, 0x89, 0x66, 0xdc,
	0xc3, 0x7b, 0x47, 0x43, 0x15, 0x2a, 0x2b, 0x9b, 0x48, 0xff, 0x39, 0xda, 0xe1, 0xe7, 0x08, 0x29,
	0x6a, 0x5d, 0x40, 0xc3, 0x44, 0x94, 0xd0, 0x71, 0xb2, 0x35, 0xe5, 0x26, 0x63, 0x99, 0xdc, 0x96,
	0xa3, 0xdb, 0x4f, 0x8c, 0x8d, 0xaf, 0x33, 0x32, 0x5e, 0xd4, 0x7a, 0x4a, 0x71, 0x2c, 0x3b, 0xda,
	0xb1, 0x1f, 0x47, 0x48, 0xbf, 0x87, 0xe8, 0x36, 0x32, 0x78, 0x5e, 0x72, 0x5c, 0x14, 0x12, 0x39,
	0x6c, 0xa2, 0x83, 0xe4, 0x5f, 0xf3, 0xb3, 0x85, 0x59, 0xb8, 0xba, 0x95, 0xee, 0x22, 0x3b, 0x8a,
	0x5a, 0x9f, 0x53, 0xa6, 0x54, 0x8b, 0x23, 0xac, 0xa1, 0x34, 0xc8, 0xe1, 0x8d, 0x3e, 0x0a, 0x64,
	0xa8, 0xa8, 0x75, 0x49, 0x1a, 0xd4, 0x92, 0x45, 0xf0, 0x67, 0x1f, 0x1d, 0x21, 0xc3, 0x16, 0x69,
	0xb0, 0x48, 0xf0, 0x92, 0x8c, 0xeb, 0x06, 0x78, 0x27, 0x58, 0x66, 0x32, 0x44, 0x40, 0x0f, 0xce,
	0x88, 0x24, 0x11, 0x32, 0x74, 0xcc, 0x45, 0x3a, 0x4a, 0xa0, 0xa8, 0x75, 0x05, 0xb5, 0x60, 0x91,
	0xb8, 0xcc, 0x8c, 0x50, 0x12, 0x42, 0xba, 0x8b, 0x50, 0x9b, 0x62, 0xd2, 0x81, 0x57, 0xe9, 0x0e,
	0xb2, 0x2d, 0x65, 0x1b, 0x8d, 0xac, 0x56, 0x46, 0xc6, 0x41, 0x50, 0x4a, 0xb6, 0xb7, 0xa0, 0x17,
	0xb4, 0x30, 0x08, 0x17, 0xbc, 0xba, 0xc3, 0xe6, 0x35, 0x93, 0xc9, 0xa2, 0xd2, 0x35, 0x90, 0x74,
	0x0f, 0xd9, 0xd9, 0xc2, 0x2b, 0x28, 0x79, 0x5e, 0xf2, 0xa9, 0x48, 0x25, 0x08, 0x8a, 0x8e, 0x91,
	0xd1, 0x96, 0xe8, 0x0c, 0x32, 0x8e, 0xda, 0x19, 0x8b, 0xe9, 0x6e, 0x32, 0xb2, 0x4e, 0x62, 0x3d,
	0x27, 0x1d, 0xc1, 0x54, 0x84, 0x5c, 0x82, 0x8b, 0x1d, 0x0e, 0xac, 0xe5, 0xbc, 0xe4, 0x65, 0x0c,
	0x1a, 0xa0, 0x7d, 0xa2, 0x3e, 0xfb, 0x19, 0x16, 0x4f, 0x63, 0x13, 0x96, 0xe8, 0x76, 0xd7, 0xcb,
	0x49, 0xe7, 0x2c, 0x4a, 0x3b, 0x62, 0xbf, 0xad, 0x8b, 0x1a, 0x05, 0xf2, 0x6f, 0xfb, 0x59, 0xc0,
	0xc4, 0x68, 0xd5, 0x84, 0xe5, 0x16, 0x32, 0x23, 0x42, 0xcd, 0x0c, 0x42, 0x93, 0x0e, 0xdb, 0x96,
	0xa4, 0x2a, 0x71, 0xc4, 0x9a, 0x70, 0xb9, 0x45, 0x29, 0x63, 0x62, 0x94, 0x46, 0x78, 0xa9, 0x65,
	0x75, 0x36, 0x46, 0x09, 0x2f, 0xb7, 0x9c, 0xba, 0xdc, 0x5f, 0xa1, 0x59, 0xb2, 0x27, 0x3d, 0x11,
	0xba, 0x19, 0x1b, 0x55, 0x66, 0x92, 0xab, 0xda, 0x69, 0x94, 0xa8, 0x5d, 0xd1, 0xaf, 0x65, 0xe8,
	0x3e, 0xb2, 0xab, 0x25, 0x9f, 0xc6, 0x66, 0x9b, 0xf0, 0x9b, 0x0c, 0x3d, 0x40, 0xc6, 0x5a, 0xc2,
	0x73, 0x4a, 0x06, 0xd8, 0x26, 0xfe, 0x36, 0x43, 0x77, 0xdb, 0x56, 0x38, 0x71, 0x45, 0x84, 0x92,
	0x99, 0xba, 0x46, 0xf8, 0x2e, 0x43, 0xff, 0x43, 0xb2, 0xdd, 0x82, 0xe7, 0x51, 0x8b, 0x45, 0x11,
	0x38, 0xed, 0xeb, 0x19, 0xba, 0xd3, 0x16, 0xcd, 0x91, 0x0a, 0x18, 0xa4, 0xbf, 0x70, 0x23, 0x43,
	0xf7, 0x93, 0xdd, 0xeb, 0xe1, 0x39, 0xd6, 0x8c, 0x14, 0xe3, 0x70, 0xb3, 0x53, 0xa9, 0x28, 0x9d,
	0xd2, 0xad, 0xae, 0x2c, 0xa6, 0x94, 0x6c, 0xa0, 0x4e, 0x52, 0x47, 0x2b, 0x19, 0x3a, 0x66, 0x9b,
	0xec, 0x84, 0x53, 0x22, 0xae, 0xa2, 0x2e, 0x49, 0x61, 0xe0, 0x76, 0x97, 0x5a, 0x01, 0xb5, 0x68,
	0xb8, 0xf8, 0xee, 0x64, 0xe8, 0x10, 0xe9, 0x4b, 0x9b, 0xca, 0x62, 0xf8, 0x74, 0x33, 0x1d, 0xb6,
	0x65, 0x3d, 0xa5, 0x74, 0x91, 0x05, 0x55, 0xf8, 0x6c, 0x33, 0x1d, 0xb1, 0x47, 0x73, 0x1a, 0x9b,
	0xb6, 0x0f, 0xa7, 0xd1, 0xc0, 0x9b, 0x5b, 0xd6, 0x81, 0x73, 0x75, 0x03, 0x6f, 0x6d, 0xf1, 0xd7,
	0xea, 0x9c, 0x32, 0xa7, 0x54, 0x5d, 0x72, 0xb8, 0xb2, 0x4a, 0x9b, 0xd5, 0x0b, 0xc2, 0x14, 0x26,
	0x6d, 0x2c, 0xf7, 0xfa, 0x3b, 0x41, 0xdb, 0xcc, 0xfb, 0xfd, 0x3e, 0x5d, 0x0f, 0xe6, 0xe3, 0x18,
	0x25, 0x87, 0x07, 0xfd, 0xbe, 0xa9, 0x1e, 0x5e, 0x7f, 0x93, 0x1e, 0xf6, 0xfb, 0x8c, 0xbd, 0xbc,
	0x92, 0xc6, 0x32, 0xc5, 0x12, 0x03, 0x8f, 0xfa, 0xe9, 0xff, 0xc9, 0xe1, 0xa2, 0xd6, 0x67, 0x98,
	0xe4, 0x49, 0x95, 0x2d, 0xe1, 0xec, 0x25, 0x59, 0x8c, 0xab, 0x58, 0x43, 0xcd, 0x22, 0xd7, 0xfd,
	0xf4, 0xea, 0xc0, 0xf5, 0x01, 0xfa, 0x5f, 0x92, 0x6b, 0x27, 0xce, 0x21, 0xea, 0x76, 0xa6, 0x3d,
	0xf8, 0x37, 0x06, 0xe8, 0x04, 0x39, 0xd6, 0x4e, 0x2b, 0xe3, 0xc5, 0x3a, 0x26, 0x06, 0x75, 0xbe,
	0x6e, 0xaa, 0x28, 0x4d, 0xda, 0x6e, 0x9c, 0x54, 0xcb, 0xce, 0x36, 0xdc, 0x1c, 0xa0, 0x4f, 0x90,
	0x23, 0x9d, 0x0a, 0x49, 0xac, 0x24, 0x47, 0x9d, 0x0f, 0x02, 0x8c, 0xcd, 0x1a, 0xf5, 0xd6, 0x00,
	0x3d, 0x48, 0xf6, 0xf6, 0xb4, 0x7d, 0x06, 0xa3, 0x48, 0xc1, 0x4a, 0x0f, 0x82, 0xb7, 0xe5, 0x08,
	0xb7, 0x07, 0xe8, 0xff, 0xc8, 0xa1, 0xbf, 0x8d, 0x0e, 0xee, 0x0c, 0xd0, 0x1c, 0xd9, 0xf7, 0x98,
	0xa0, 0xe0, 0xfb, 0xae, 0x72, 0xac, 0x59, 0x0a, 0x96, 0xa4, 0xba, 0x14, 0x21, 0x0f, 0x11, 0x7e,
	0x18, 0xa0, 0x87, 0xc8, 0x7e, 0x3b, 0x7f, 0xa5, 0x61, 0x81, 0xf1, 0xa4, 0x0a, 0xab, 0x61, 0x3e,
	0x08, 0x54, 0x5d, 0x1a, 0xf8, 0x7c, 0xd0, 0x17, 0xa0, 0x93, 0xe2, 0xbf, 0xf2, 0x91, 0x46, 0xc6,
	0x9b, 0x79, 0xce, 0x91, 0xc3, 0x17, 0x83, 0xf4, 0x08, 0x39, 0xb8, 0x11, 0x75, 0x32, 0x52, 0xc1,
	0x12, 0x72, 0xf8, 0x72, 0xd0, 0x27, 0xd9, 0x93, 0xb5, 0xb6, 0x00, 0xbe, 0x1a, 0xa4, 0xc7, 0xc9,
	0xd1, 0x2e, 0x5e, 0x49, 0x06, 0xaa, 0x26, 0x64, 0xe8, 0x3d, 0x97, 0x31, 0x40, 0xd1, 0x40, 0x0e,
	0x5f, 0x0f, 0xfa, 0xe2, 0x9e, 0xd6, 0xaa, 0x1e, 0xcf, 0x60, 0x6d, 0x01, 0xf5, 0x59, 0x15, 0x16,
	0x1b, 0x28, 0x8d, 0x3d, 0x9b, 0x57, 0x88, 0x8f, 0xae, 0x07, 0x61, 0x6d, 0x14, 0xbc, 0x4d, 0x7c,
	0x45, 0xda, 0x58, 0xe7, 0x65, 0x5a, 0x31, 0x69, 0x91, 0x52, 0x01, 0xde, 0x21, 0xf4, 0x30, 0x39,
	0xb0, 0x4a, 0xa9, 0x60, 0xa0, 0xd1, 0xcc, 0x9a, 0x2a, 0xa6, 0x0b, 0xc2, 0x38, 0x0d, 0x78, 0x97,
	0xf8, 0x24, 0xdb, 0x38, 0x3e, 0xe2, 0x0a, 0x4a, 0x33, 0xaf, 0x3c, 0xef, 0x3d, 0xe2, 0x4f, 0xbe,
	0x33, 0xee, 0x36, 0xd4, 0x7c, 0x33, 0x46, 0x78, 0x9f, 0xd0, 0x51, 0xbb, 0xa1, 0x5c, 0x20, 0x6e,
	0x50, 0xc3, 0x55, 0xe2, 0x2f, 0x98, 0x45, 0xf3, 0x81, 0x49, 0x6f, 0x3f, 0xc2, 0x07, 0xc4, 0x4f,
	0x36, 0x0b, 0x17, 0x90, 0xad, 0x0a, 0x3e, 0x24, 0x74, 0x87, 0x9d, 0xbf, 0xde, 0xfe, 0xa2, 0x82,
	0x8f, 0x3a, 0x0c, 0xfb, 0xdc, 0xe0, 0xe3, 0x0e, 0xa2, 0x2d, 0xd8, 0x27, 0xc4, 0x9f, 0xb2, 0x19,
	0x4c, 0x12, 0x16, 0xe2, 0x34, 0x36, 0xe7, 0xd2, 0x11, 0x95, 0x18, 0x94, 0x81, 0x1d, 0x15, 0x3f,
	0x0d, 0x3d, 0x8e, 0x91, 0x4e, 0x98, 0x9f, 0x87, 0xe8, 0x5e, 0xb7, 0x84, 0x50, 0x37, 0x44, 0x80,
	0xe9, 0x56, 0x58, 0x1d, 0xa8, 0xaf, 0xe6, 0xbc, 0x76, 0xb7, 0x2c, 0x45, 0x50, 0xc3, 0x6b, 0x39,
	0x7f, 0xf6, 0xba, 0x19, 0xbe, 0x20, 0x45, 0xc9, 0x63, 0x25, 0xa4, 0x81, 0xd7, 0x73, 0x74, 0x8f,
	0xdb, 0x99, 0x8e, 0x9a, 0x14, 0x84, 0xc6, 0xc0, 0x28, 0xdd, 0x84, 0x7b, 0x39, 0xfa, 0x2c, 0x39,
	0xd9, 0x4b, 0xe4, 0xeb, 0xed, 0xc6, 0x3c, 0xf2, 0x29, 0x8d, 0x3c, 0xbd, 0x63, 0x2c, 0xaa, 0xd4,
	0x17, 0x2e, 0x60, 0x60, 0xe0, 0x7e, 0xce, 0x8f, 0x8b, 0x2e, 0xd5, 0xe2, 0xb2, 0x48, 0x8c, 0x90,
	0x61, 0x19, 0x03, 0xa5, 0x79, 0x6b, 0x54, 0x3e, 0xc8, 0xd1, 0x67, 0xc8, 0x53, 0xbd, 0x14, 0x1c,
	0xf1, 0xac, 0xbd, 0x03, 0xe9, 0xa2, 0x67, 0xd2, 0x4c, 0xda, 0x84, 0x58, 0x80, 0x1c, 0x1e, 0xe6,
	0xe8, 0x49, 0x32, 0xd1, 0xdb, 0x51, 0x9a, 0xb3, 0x30, 0x9e, 0x7a, 0x2a, 0x62, 0x61, 0x7a, 0x13,
	0x84, 0x46, 0x0e, 0x8f, 0x72, 0xf4, 0x04, 0x39, 0xfe, 0x8f, 0x32, 0x83, 0xdf, 0x36, 0xd4, 0x29,
	0x2e, 0xc7, 0xa9, 0xd5, 0x1e, 0x3a, 0xbf, 0x6f, 0x18, 0xdd, 0x86, 0x7e, 0x4a, 0x05, 0xf8, 0x23,
	0x37, 0xf9, 0xe4, 0xca, 0xaf, 0xd9, 0x4d, 0xd7, 0xee, 0x66, 0x33, 0x2b, 0x77, 0xb3, 0x99, 0x5f,
	0xee, 0x66, 0x33, 0x2f, 0x66, 0x17, 0x50, 0x9b, 0xe6, 0xb8, 0xc1, 0xa0, 0x3a, 0xe1, 0x9f, 0x86,
	0x13, 0xf1, 0x52, 0xb8, 0xfa, 0xf4, 0x5c, 0xe8, 0xb3, 0x8f, 0xca, 0xa7, 0xff, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x89, 0xad, 0x04, 0x21, 0x94, 0x0a, 0x00, 0x00,
}

func (m *ErrDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Codes) > 0 {
		dAtA2 := make([]byte, len(m.Codes)*10)
		var j1 int
		for _, num := range m.Codes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintErrcode(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrcode(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrcode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ErrDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Codes) > 0 {
		l = 0
		for _, e := range m.Codes {
			l += sovErrcode(uint64(e))
		}
		n += 1 + sovErrcode(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrcode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrcode(x uint64) (n int) {
	return sovErrcode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ErrDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrcode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v ErrCode
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowErrcode
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ErrCode(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Codes = append(m.Codes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowErrcode
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthErrcode
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthErrcode
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Codes) == 0 {
					m.Codes = make([]ErrCode, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ErrCode
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowErrcode
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ErrCode(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Codes = append(m.Codes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Codes", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErrcode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrcode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrcode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrcode
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrcode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrcode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthErrcode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrcode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrcode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrcode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrcode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrcode = fmt.Errorf("proto: unexpected end of group")
)
