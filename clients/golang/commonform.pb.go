// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commonform.proto

package open_commonform_org

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Document_DocumentMeta_InputFormat int32

const (
	Document_DocumentMeta_commonmark Document_DocumentMeta_InputFormat = 0
	Document_DocumentMeta_json       Document_DocumentMeta_InputFormat = 1
	Document_DocumentMeta_markup     Document_DocumentMeta_InputFormat = 2
)

var Document_DocumentMeta_InputFormat_name = map[int32]string{
	0: "commonmark",
	1: "json",
	2: "markup",
}

var Document_DocumentMeta_InputFormat_value = map[string]int32{
	"commonmark": 0,
	"json":       1,
	"markup":     2,
}

func (x Document_DocumentMeta_InputFormat) String() string {
	return proto.EnumName(Document_DocumentMeta_InputFormat_name, int32(x))
}

func (Document_DocumentMeta_InputFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{0, 0, 0}
}

// we use an enum both for type safety as well as to
// signal what has been implemented in the server.
type ToAssemble_OutputFormat int32

const (
	ToAssemble_docx     ToAssemble_OutputFormat = 0
	ToAssemble_html     ToAssemble_OutputFormat = 1
	ToAssemble_html5    ToAssemble_OutputFormat = 2
	ToAssemble_markdown ToAssemble_OutputFormat = 3
)

var ToAssemble_OutputFormat_name = map[int32]string{
	0: "docx",
	1: "html",
	2: "html5",
	3: "markdown",
}

var ToAssemble_OutputFormat_value = map[string]int32{
	"docx":     0,
	"html":     1,
	"html5":    2,
	"markdown": 3,
}

func (x ToAssemble_OutputFormat) String() string {
	return proto.EnumName(ToAssemble_OutputFormat_name, int32(x))
}

func (ToAssemble_OutputFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{2, 0}
}

// the below are options but we should be largely
// opinionated for most uses.
type ToAssemble_NumberFormat int32

const (
	ToAssemble_decimal ToAssemble_NumberFormat = 0
	ToAssemble_outline ToAssemble_NumberFormat = 1
)

var ToAssemble_NumberFormat_name = map[int32]string{
	0: "decimal",
	1: "outline",
}

var ToAssemble_NumberFormat_value = map[string]int32{
	"decimal": 0,
	"outline": 1,
}

func (x ToAssemble_NumberFormat) String() string {
	return proto.EnumName(ToAssemble_NumberFormat_name, int32(x))
}

func (ToAssemble_NumberFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{2, 1}
}

// The Document message is a fairly straight-forward and
// object which includes a modicum of meta data which is needed
// both by the cf-server as well as clients of this server
// along with a filename and the contents of the file as a
// byte slice (golang) or a bytes buffer (node).
type Document struct {
	Meta                 *Document_DocumentMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{0}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetMeta() *Document_DocumentMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Document) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Document_DocumentMeta struct {
	Name                 string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mime                 string                            `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
	Format               Document_DocumentMeta_InputFormat `protobuf:"varint,3,opt,name=format,proto3,enum=commonform.Document_DocumentMeta_InputFormat" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Document_DocumentMeta) Reset()         { *m = Document_DocumentMeta{} }
func (m *Document_DocumentMeta) String() string { return proto.CompactTextString(m) }
func (*Document_DocumentMeta) ProtoMessage()    {}
func (*Document_DocumentMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{0, 0}
}

func (m *Document_DocumentMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document_DocumentMeta.Unmarshal(m, b)
}
func (m *Document_DocumentMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document_DocumentMeta.Marshal(b, m, deterministic)
}
func (m *Document_DocumentMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document_DocumentMeta.Merge(m, src)
}
func (m *Document_DocumentMeta) XXX_Size() int {
	return xxx_messageInfo_Document_DocumentMeta.Size(m)
}
func (m *Document_DocumentMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_Document_DocumentMeta.DiscardUnknown(m)
}

var xxx_messageInfo_Document_DocumentMeta proto.InternalMessageInfo

func (m *Document_DocumentMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Document_DocumentMeta) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *Document_DocumentMeta) GetFormat() Document_DocumentMeta_InputFormat {
	if m != nil {
		return m.Format
	}
	return Document_DocumentMeta_commonmark
}

// The Blanks message is a simple wrapper around an array of
// strings since we cannot return arrays.
type Blanks struct {
	Blanks               []string `protobuf:"bytes,1,rep,name=blanks,proto3" json:"blanks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blanks) Reset()         { *m = Blanks{} }
func (m *Blanks) String() string { return proto.CompactTextString(m) }
func (*Blanks) ProtoMessage()    {}
func (*Blanks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{1}
}

func (m *Blanks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blanks.Unmarshal(m, b)
}
func (m *Blanks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blanks.Marshal(b, m, deterministic)
}
func (m *Blanks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blanks.Merge(m, src)
}
func (m *Blanks) XXX_Size() int {
	return xxx_messageInfo_Blanks.Size(m)
}
func (m *Blanks) XXX_DiscardUnknown() {
	xxx_messageInfo_Blanks.DiscardUnknown(m)
}

var xxx_messageInfo_Blanks proto.InternalMessageInfo

func (m *Blanks) GetBlanks() []string {
	if m != nil {
		return m.Blanks
	}
	return nil
}

// The ToAssemble message is fairly complex due to the structure
// of the information which commonform expects us to send it.
//
// Many of the nobs that can be turned within the commonform server
// have reasonable defaults given BOTH within the server/engine
// itself and within the ToAssemble message. As such, many of the
// below feilds are optional (or repeated, which assumes optional
// in protobuf parlance).
type ToAssemble struct {
	Document   *Document               `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Title      string                  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Blanks     []*ToAssemble_Blank     `protobuf:"bytes,3,rep,name=blanks,proto3" json:"blanks,omitempty"`
	Signatures []*ToAssemble_Signature `protobuf:"bytes,4,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// really the below should be a formalized object,
	// but protobuf cannot cope with the mixed output
	// it would require since it is really a
	// map<string, (string|bool)> . as such, until
	// we systematize the sytles object or protobuf
	// changes it's type guarantees we just send the
	// json bytes in and allow the server to Marshal
	// them into a json object the server can read.
	Styles               []byte                  `protobuf:"bytes,5,opt,name=styles,proto3" json:"styles,omitempty"`
	Format               ToAssemble_OutputFormat `protobuf:"varint,6,opt,name=format,proto3,enum=commonform.ToAssemble_OutputFormat" json:"format,omitempty"`
	Numbering            ToAssemble_NumberFormat `protobuf:"varint,7,opt,name=numbering,proto3,enum=commonform.ToAssemble_NumberFormat" json:"numbering,omitempty"`
	Hash                 bool                    `protobuf:"varint,8,opt,name=hash,proto3" json:"hash,omitempty"`
	IndentMargins        bool                    `protobuf:"varint,9,opt,name=indentMargins,proto3" json:"indentMargins,omitempty"`
	LeftAlignTitle       bool                    `protobuf:"varint,10,opt,name=leftAlignTitle,proto3" json:"leftAlignTitle,omitempty"`
	MarkFilled           bool                    `protobuf:"varint,11,opt,name=markFilled,proto3" json:"markFilled,omitempty"`
	UnfilledBlanks       string                  `protobuf:"bytes,12,opt,name=unfilledBlanks,proto3" json:"unfilledBlanks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ToAssemble) Reset()         { *m = ToAssemble{} }
func (m *ToAssemble) String() string { return proto.CompactTextString(m) }
func (*ToAssemble) ProtoMessage()    {}
func (*ToAssemble) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{2}
}

func (m *ToAssemble) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToAssemble.Unmarshal(m, b)
}
func (m *ToAssemble) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToAssemble.Marshal(b, m, deterministic)
}
func (m *ToAssemble) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToAssemble.Merge(m, src)
}
func (m *ToAssemble) XXX_Size() int {
	return xxx_messageInfo_ToAssemble.Size(m)
}
func (m *ToAssemble) XXX_DiscardUnknown() {
	xxx_messageInfo_ToAssemble.DiscardUnknown(m)
}

var xxx_messageInfo_ToAssemble proto.InternalMessageInfo

func (m *ToAssemble) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *ToAssemble) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToAssemble) GetBlanks() []*ToAssemble_Blank {
	if m != nil {
		return m.Blanks
	}
	return nil
}

func (m *ToAssemble) GetSignatures() []*ToAssemble_Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *ToAssemble) GetStyles() []byte {
	if m != nil {
		return m.Styles
	}
	return nil
}

func (m *ToAssemble) GetFormat() ToAssemble_OutputFormat {
	if m != nil {
		return m.Format
	}
	return ToAssemble_docx
}

func (m *ToAssemble) GetNumbering() ToAssemble_NumberFormat {
	if m != nil {
		return m.Numbering
	}
	return ToAssemble_decimal
}

func (m *ToAssemble) GetHash() bool {
	if m != nil {
		return m.Hash
	}
	return false
}

func (m *ToAssemble) GetIndentMargins() bool {
	if m != nil {
		return m.IndentMargins
	}
	return false
}

func (m *ToAssemble) GetLeftAlignTitle() bool {
	if m != nil {
		return m.LeftAlignTitle
	}
	return false
}

func (m *ToAssemble) GetMarkFilled() bool {
	if m != nil {
		return m.MarkFilled
	}
	return false
}

func (m *ToAssemble) GetUnfilledBlanks() string {
	if m != nil {
		return m.UnfilledBlanks
	}
	return ""
}

// we use a formalized object to enhance type safety
type ToAssemble_Blank struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToAssemble_Blank) Reset()         { *m = ToAssemble_Blank{} }
func (m *ToAssemble_Blank) String() string { return proto.CompactTextString(m) }
func (*ToAssemble_Blank) ProtoMessage()    {}
func (*ToAssemble_Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{2, 0}
}

func (m *ToAssemble_Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToAssemble_Blank.Unmarshal(m, b)
}
func (m *ToAssemble_Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToAssemble_Blank.Marshal(b, m, deterministic)
}
func (m *ToAssemble_Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToAssemble_Blank.Merge(m, src)
}
func (m *ToAssemble_Blank) XXX_Size() int {
	return xxx_messageInfo_ToAssemble_Blank.Size(m)
}
func (m *ToAssemble_Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_ToAssemble_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_ToAssemble_Blank proto.InternalMessageInfo

func (m *ToAssemble_Blank) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ToAssemble_Blank) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// we use a formalized object to enhance type safety
type ToAssemble_Signature struct {
	Term                 string                         `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Header               string                         `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	Information          []string                       `protobuf:"bytes,4,rep,name=information,proto3" json:"information,omitempty"`
	SamePage             bool                           `protobuf:"varint,5,opt,name=samePage,proto3" json:"samePage,omitempty"`
	Entities             []*ToAssemble_Signature_Entity `protobuf:"bytes,6,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ToAssemble_Signature) Reset()         { *m = ToAssemble_Signature{} }
func (m *ToAssemble_Signature) String() string { return proto.CompactTextString(m) }
func (*ToAssemble_Signature) ProtoMessage()    {}
func (*ToAssemble_Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{2, 1}
}

func (m *ToAssemble_Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToAssemble_Signature.Unmarshal(m, b)
}
func (m *ToAssemble_Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToAssemble_Signature.Marshal(b, m, deterministic)
}
func (m *ToAssemble_Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToAssemble_Signature.Merge(m, src)
}
func (m *ToAssemble_Signature) XXX_Size() int {
	return xxx_messageInfo_ToAssemble_Signature.Size(m)
}
func (m *ToAssemble_Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_ToAssemble_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_ToAssemble_Signature proto.InternalMessageInfo

func (m *ToAssemble_Signature) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *ToAssemble_Signature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ToAssemble_Signature) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *ToAssemble_Signature) GetInformation() []string {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *ToAssemble_Signature) GetSamePage() bool {
	if m != nil {
		return m.SamePage
	}
	return false
}

func (m *ToAssemble_Signature) GetEntities() []*ToAssemble_Signature_Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

type ToAssemble_Signature_Entity struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Form                 string   `protobuf:"bytes,2,opt,name=form,proto3" json:"form,omitempty"`
	Jurisdiction         string   `protobuf:"bytes,3,opt,name=jurisdiction,proto3" json:"jurisdiction,omitempty"`
	By                   string   `protobuf:"bytes,4,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToAssemble_Signature_Entity) Reset()         { *m = ToAssemble_Signature_Entity{} }
func (m *ToAssemble_Signature_Entity) String() string { return proto.CompactTextString(m) }
func (*ToAssemble_Signature_Entity) ProtoMessage()    {}
func (*ToAssemble_Signature_Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1f3a923a12f016, []int{2, 1, 0}
}

func (m *ToAssemble_Signature_Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToAssemble_Signature_Entity.Unmarshal(m, b)
}
func (m *ToAssemble_Signature_Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToAssemble_Signature_Entity.Marshal(b, m, deterministic)
}
func (m *ToAssemble_Signature_Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToAssemble_Signature_Entity.Merge(m, src)
}
func (m *ToAssemble_Signature_Entity) XXX_Size() int {
	return xxx_messageInfo_ToAssemble_Signature_Entity.Size(m)
}
func (m *ToAssemble_Signature_Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_ToAssemble_Signature_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_ToAssemble_Signature_Entity proto.InternalMessageInfo

func (m *ToAssemble_Signature_Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ToAssemble_Signature_Entity) GetForm() string {
	if m != nil {
		return m.Form
	}
	return ""
}

func (m *ToAssemble_Signature_Entity) GetJurisdiction() string {
	if m != nil {
		return m.Jurisdiction
	}
	return ""
}

func (m *ToAssemble_Signature_Entity) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func init() {
	proto.RegisterEnum("commonform.Document_DocumentMeta_InputFormat", Document_DocumentMeta_InputFormat_name, Document_DocumentMeta_InputFormat_value)
	proto.RegisterEnum("commonform.ToAssemble_OutputFormat", ToAssemble_OutputFormat_name, ToAssemble_OutputFormat_value)
	proto.RegisterEnum("commonform.ToAssemble_NumberFormat", ToAssemble_NumberFormat_name, ToAssemble_NumberFormat_value)
	proto.RegisterType((*Document)(nil), "commonform.Document")
	proto.RegisterType((*Document_DocumentMeta)(nil), "commonform.Document.DocumentMeta")
	proto.RegisterType((*Blanks)(nil), "commonform.Blanks")
	proto.RegisterType((*ToAssemble)(nil), "commonform.ToAssemble")
	proto.RegisterType((*ToAssemble_Blank)(nil), "commonform.ToAssemble.Blank")
	proto.RegisterType((*ToAssemble_Signature)(nil), "commonform.ToAssemble.Signature")
	proto.RegisterType((*ToAssemble_Signature_Entity)(nil), "commonform.ToAssemble.Signature.Entity")
}

func init() { proto.RegisterFile("commonform.proto", fileDescriptor_2f1f3a923a12f016) }

var fileDescriptor_2f1f3a923a12f016 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xc1, 0x6e, 0xdb, 0x38,
	0x10, 0xb5, 0x64, 0x5b, 0x96, 0xc6, 0x5e, 0x43, 0xe0, 0x06, 0x81, 0x60, 0x2c, 0x16, 0x5a, 0xed,
	0x62, 0xeb, 0x1e, 0xea, 0xb4, 0x4e, 0x73, 0x69, 0x2e, 0x4d, 0x52, 0x07, 0xe8, 0xa1, 0x6d, 0xa0,
	0xe6, 0x94, 0x1b, 0x6d, 0x31, 0x32, 0x13, 0x91, 0x32, 0x24, 0xaa, 0x8d, 0x3f, 0xa0, 0x97, 0x7e,
	0x46, 0x3f, 0xa1, 0xe8, 0xff, 0xf4, 0x57, 0x0a, 0x8e, 0x65, 0x59, 0x49, 0x9c, 0xe6, 0xf6, 0xe6,
	0xe9, 0xbd, 0xe1, 0x70, 0x66, 0x68, 0x83, 0x3b, 0x4b, 0x85, 0x48, 0xe5, 0x65, 0x9a, 0x89, 0xd1,
	0x22, 0x4b, 0x55, 0x4a, 0x60, 0xc3, 0x04, 0x5f, 0x4d, 0xb0, 0xdf, 0xa4, 0xb3, 0x42, 0x30, 0xa9,
	0xc8, 0x01, 0xb4, 0x04, 0x53, 0xd4, 0x33, 0x7c, 0x63, 0xd8, 0x1d, 0xff, 0x33, 0xaa, 0x39, 0xd7,
	0x9a, 0x0a, 0xbc, 0x63, 0x8a, 0x86, 0x28, 0x27, 0x04, 0x5a, 0x11, 0x55, 0xd4, 0x33, 0x7d, 0x63,
	0xd8, 0x0b, 0x11, 0x0f, 0xbe, 0x1b, 0xd0, 0xab, 0x4b, 0xb5, 0x48, 0x52, 0xc1, 0x30, 0xb7, 0x13,
	0x22, 0xd6, 0x9c, 0xe0, 0x82, 0xa1, 0xd1, 0x09, 0x11, 0x93, 0x09, 0x58, 0xfa, 0x40, 0xaa, 0xbc,
	0xa6, 0x6f, 0x0c, 0xfb, 0xe3, 0x67, 0x8f, 0x56, 0x31, 0x7a, 0x2b, 0x17, 0x85, 0x3a, 0x45, 0x53,
	0x58, 0x9a, 0x83, 0x7d, 0xe8, 0xd6, 0x68, 0xd2, 0x87, 0xf2, 0xd2, 0x82, 0x66, 0xd7, 0x6e, 0x83,
	0xd8, 0xd0, 0xba, 0xca, 0x53, 0xe9, 0x1a, 0x04, 0xc0, 0xd2, 0x5c, 0xb1, 0x70, 0xcd, 0xc0, 0x07,
	0xeb, 0x38, 0xa1, 0xf2, 0x3a, 0x27, 0xbb, 0x60, 0x4d, 0x11, 0x79, 0x86, 0xdf, 0x1c, 0x3a, 0x61,
	0x19, 0x05, 0x3f, 0x3b, 0x00, 0xe7, 0xe9, 0x51, 0x9e, 0x33, 0x31, 0x4d, 0x18, 0x79, 0x0e, 0x76,
	0x54, 0x56, 0x52, 0x36, 0x6d, 0x67, 0x5b, 0xb9, 0x61, 0xa5, 0x22, 0x3b, 0xd0, 0x56, 0x5c, 0x25,
	0xeb, 0x3b, 0xaf, 0x02, 0xf2, 0xb2, 0x3a, 0xae, 0xe9, 0x37, 0x87, 0xdd, 0xf1, 0x5f, 0xf5, 0x2c,
	0x9b, 0xf3, 0x46, 0x58, 0xdd, 0xba, 0x18, 0xf2, 0x1a, 0x20, 0xe7, 0xb1, 0xa4, 0xaa, 0xc8, 0x58,
	0xee, 0xb5, 0xd0, 0xe9, 0x3f, 0xe0, 0xfc, 0xb8, 0x16, 0x86, 0x35, 0x8f, 0xbe, 0x66, 0xae, 0x96,
	0x09, 0xcb, 0xbd, 0x36, 0xce, 0xae, 0x8c, 0xc8, 0x61, 0x35, 0x04, 0x0b, 0x87, 0xf0, 0xef, 0x03,
	0x59, 0x3f, 0x14, 0xea, 0x5e, 0xeb, 0xc9, 0x11, 0x38, 0xb2, 0x10, 0x53, 0x96, 0x71, 0x19, 0x7b,
	0x9d, 0xdf, 0xfa, 0xdf, 0xa3, 0xae, 0xf4, 0x6f, 0x5c, 0x7a, 0x31, 0xe6, 0x34, 0x9f, 0x7b, 0xb6,
	0x6f, 0x0c, 0xed, 0x10, 0x31, 0xf9, 0x0f, 0xfe, 0xe0, 0x32, 0xd2, 0x33, 0xa7, 0x59, 0xcc, 0x65,
	0xee, 0x39, 0xf8, 0xf1, 0x36, 0x49, 0xfe, 0x87, 0x7e, 0xc2, 0x2e, 0xd5, 0x51, 0xc2, 0x63, 0x79,
	0x8e, 0x8d, 0x06, 0x94, 0xdd, 0x61, 0xc9, 0xdf, 0x00, 0x7a, 0xec, 0xa7, 0x3c, 0x49, 0x58, 0xe4,
	0x75, 0x51, 0x53, 0x63, 0x74, 0x9e, 0x42, 0x5e, 0x22, 0x5e, 0xad, 0x84, 0xd7, 0xc3, 0x81, 0xdd,
	0x61, 0x07, 0x2f, 0xa0, 0x8d, 0x68, 0xeb, 0x7e, 0xef, 0x40, 0xfb, 0x13, 0x4d, 0x8a, 0x6a, 0xd8,
	0x18, 0x0c, 0x7e, 0x98, 0xe0, 0x54, 0xe3, 0xd0, 0x3e, 0xc5, 0x32, 0xb1, 0xf6, 0x69, 0x5c, 0xe5,
	0x32, 0x6b, 0xb9, 0x76, 0xc1, 0x9a, 0x33, 0x1a, 0xb1, 0x0c, 0xdf, 0x85, 0x13, 0x96, 0x11, 0xf1,
	0xa1, 0xcb, 0xe5, 0xaa, 0xf3, 0x3c, 0x95, 0xb8, 0x05, 0x4e, 0x58, 0xa7, 0xc8, 0x00, 0xec, 0x9c,
	0x0a, 0x76, 0x46, 0x63, 0x86, 0x63, 0xb6, 0xc3, 0x2a, 0x26, 0x27, 0x60, 0x33, 0xa9, 0xb8, 0xe2,
	0x2c, 0xf7, 0x2c, 0x5c, 0xa0, 0x27, 0x8f, 0x2d, 0xd0, 0x68, 0xa2, 0x0d, 0xcb, 0xb0, 0x32, 0x0e,
	0x22, 0xb0, 0x56, 0xdc, 0x43, 0x8f, 0x5c, 0xe7, 0x5a, 0x5f, 0x46, 0x63, 0x12, 0x40, 0xef, 0xaa,
	0xc8, 0x78, 0x1e, 0xf1, 0x19, 0x56, 0xbd, 0xba, 0xd2, 0x2d, 0x8e, 0xf4, 0xc1, 0x9c, 0x2e, 0xbd,
	0x16, 0x7e, 0x31, 0xa7, 0xcb, 0xe0, 0x10, 0x7a, 0xf5, 0x75, 0xd3, 0x4f, 0x38, 0x4a, 0x67, 0x37,
	0xab, 0xc7, 0x3c, 0x57, 0x22, 0x71, 0x0d, 0xe2, 0x40, 0x5b, 0xa3, 0x03, 0xd7, 0x24, 0x3d, 0xb0,
	0xf5, 0x38, 0xa3, 0xf4, 0xb3, 0x74, 0x9b, 0xc1, 0x10, 0x7a, 0xf5, 0x5d, 0x23, 0x5d, 0xe8, 0x44,
	0x6c, 0xc6, 0x05, 0x4d, 0xdc, 0x86, 0x0e, 0xd2, 0x42, 0x25, 0x5c, 0x32, 0xd7, 0x18, 0x7f, 0x31,
	0xc0, 0x3d, 0xc1, 0x0e, 0x68, 0xe9, 0x44, 0xc6, 0x5c, 0x32, 0x72, 0x00, 0x9d, 0xc9, 0x8d, 0xca,
	0xe8, 0x4c, 0x91, 0xad, 0x0f, 0x7c, 0x40, 0xea, 0xec, 0x6a, 0x35, 0x82, 0x06, 0x79, 0x05, 0x76,
	0xf5, 0x53, 0xb1, 0xbb, 0xbd, 0xaf, 0x83, 0xad, 0xf9, 0x82, 0xc6, 0xf1, 0x05, 0xfc, 0x99, 0x66,
	0x71, 0xfd, 0x63, 0xba, 0x60, 0xf2, 0xf8, 0x5e, 0x6d, 0x67, 0xc6, 0xc5, 0xd3, 0x98, 0xab, 0x79,
	0x31, 0xd5, 0xda, 0xbd, 0x8d, 0x7e, 0x4f, 0xeb, 0x6f, 0xf9, 0xb3, 0xf8, 0x9b, 0xd9, 0x3c, 0x39,
	0x9d, 0x4c, 0x2d, 0xfc, 0x1f, 0xd8, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x81, 0xb0, 0x9a, 0x4a,
	0x1b, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommonFormEngineClient is the client API for CommonFormEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonFormEngineClient interface {
	// Extracts the paramters.
	Extract(ctx context.Context, in *Document, opts ...grpc.CallOption) (*Blanks, error)
	// Assembles the document.
	Assemble(ctx context.Context, in *ToAssemble, opts ...grpc.CallOption) (*Document, error)
}

type commonFormEngineClient struct {
	cc *grpc.ClientConn
}

func NewCommonFormEngineClient(cc *grpc.ClientConn) CommonFormEngineClient {
	return &commonFormEngineClient{cc}
}

func (c *commonFormEngineClient) Extract(ctx context.Context, in *Document, opts ...grpc.CallOption) (*Blanks, error) {
	out := new(Blanks)
	err := c.cc.Invoke(ctx, "/commonform.CommonFormEngine/Extract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonFormEngineClient) Assemble(ctx context.Context, in *ToAssemble, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, "/commonform.CommonFormEngine/Assemble", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonFormEngineServer is the server API for CommonFormEngine service.
type CommonFormEngineServer interface {
	// Extracts the paramters.
	Extract(context.Context, *Document) (*Blanks, error)
	// Assembles the document.
	Assemble(context.Context, *ToAssemble) (*Document, error)
}

func RegisterCommonFormEngineServer(s *grpc.Server, srv CommonFormEngineServer) {
	s.RegisterService(&_CommonFormEngine_serviceDesc, srv)
}

func _CommonFormEngine_Extract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Document)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonFormEngineServer).Extract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonform.CommonFormEngine/Extract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonFormEngineServer).Extract(ctx, req.(*Document))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonFormEngine_Assemble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToAssemble)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonFormEngineServer).Assemble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonform.CommonFormEngine/Assemble",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonFormEngineServer).Assemble(ctx, req.(*ToAssemble))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommonFormEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commonform.CommonFormEngine",
	HandlerType: (*CommonFormEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Extract",
			Handler:    _CommonFormEngine_Extract_Handler,
		},
		{
			MethodName: "Assemble",
			Handler:    _CommonFormEngine_Assemble_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commonform.proto",
}
