// Representation of code ownership for a repository as described in a CODEOWNERS file.
// As various implementations have slightly different syntax for CODEOWNERS files,
// this algebraic representation servers as a unified funnel.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: codeowners.proto

package v1

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

// File represents the contents of a single CODEOWNERS file.
// As specified by various CODEOWNERS implementations the following apply:
//   - There is at most one CODEOWNERS file per repository.
//   - The semantic contents of the file boil down to rules.
//   - Order matters: When discerning ownership for a path
//     only the owners from the last rule that matches the path
//     is applied.
//   - Except if using sections - then every section is considered
//     separately. That is, an owner is potentially extracted
//     for every section.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule []*Rule `protobuf:"bytes,1,rep,name=rule,proto3" json:"rule,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeowners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_codeowners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_codeowners_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetRule() []*Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

// Rule associates a single pattern to match a path with an owner.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Patterns are familliar glob patterns that match file paths.
	// * `filename` matches any file with that name, for example:
	//   * `/filename` and `/src/filename` match.
	// * `directory/path/` matches any tree of subdirectories rooted
	//   at this pattern, for example:
	//   * `/src/directory/path/file` matches.
	//   * `/src/directory/path/another/directory/file` matches.
	// * `directory/*` matches only files with specified parent,
	//   but not descendants, for example:
	//   * `/src/foo/bar/directory/file` matches.
	//   * `/src/foo/bar/directory/another/file` does not match.
	// * Any of the above can be prefixed with `/`, which further
	//   filters the match, by requiring the file path match to be
	//   rooted at the directory root, for `/src/dir/*`:
	//   * `/src/dir/file` matches.
	//   * `/main/src/dir/file` does not match, as `src` is not top-level.
	//   * `/src/dir/another/file` does not match as `*` matches
	//     only files directly contained in specified directory.
	// * In the above patterns `/**/` can be used to match any sub-path
	//   between two parts of a pattern. For example: `/docs/**/internal/`
	//   will match `/docs/foo/bar/internal/file`.
	// * The file part of the pattern can use a `*` wildcard like so:
	//   `docs/*.md` will match `/src/docs/index.md` but not `/src/docs/index.js`.
	// * In BITBUCKET plugin, patterns that serve to exclude ownership
	//   start with an exclamation mark `!/src/noownershere`. These are
	//   translated to a pattern without the `!` and now owners.
	Pattern string `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// Owners list all the parties that claim ownership over files
	// matched by a given pattern.
	// This list may be empty. In such case it denotes an abandoned
	// codebase, and can be used if there is an un-owned subdirectory
	// within otherwise owned directory structure.
	Owner []*Owner `protobuf:"bytes,2,rep,name=owner,proto3" json:"owner,omitempty"`
	// Optionally a rule can be associated with a section name.
	// The name must be lowercase, as the names of sections in text
	// representation of the codeowners file are case-insensitive.
	// Each section represents a kind-of-ownership. That is,
	// when evaluating an owner for a path, only one rule can apply
	// for a path, but that is within the scope of a section.
	// For instance a CODEOWNERS file could specify a [PM] section
	// associating product managers with codebases. This rule set
	// can be completely independent of the others. In that case,
	// when evaluating owners, the result also contains a separate
	// owners for the PM section.
	SectionName string `protobuf:"bytes,3,opt,name=section_name,json=sectionName,proto3" json:"section_name,omitempty"`
	// The line number this rule originally appeared in in the input data.
	LineNumber int32 `protobuf:"varint,4,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeowners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_codeowners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_codeowners_proto_rawDescGZIP(), []int{1}
}

func (x *Rule) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *Rule) GetOwner() []*Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Rule) GetSectionName() string {
	if x != nil {
		return x.SectionName
	}
	return ""
}

func (x *Rule) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

// Owner is denoted by either a handle or an email.
// We expect exactly one of the fields to be present.
type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Handle can refer to a user or a team defined externally.
	// In the text config, a handle always starts with `@`.
	// In can contain `/` to denote a sub-group.
	// The string content of the handle stored here DOES NOT CONTAIN
	// the initial `@` sign.
	Handle string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
	// E-mail can be used instead of a handle to denote an owner account.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeowners_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_codeowners_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_codeowners_proto_rawDescGZIP(), []int{2}
}

func (x *Owner) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *Owner) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_codeowners_proto protoreflect.FileDescriptor

var file_codeowners_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x64, 0x65, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x6f, 0x77, 0x6e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x33, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x77,
	0x6e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x2e, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f,
	0x77, 0x6e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x35, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x77, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_codeowners_proto_rawDescOnce sync.Once
	file_codeowners_proto_rawDescData = file_codeowners_proto_rawDesc
)

func file_codeowners_proto_rawDescGZIP() []byte {
	file_codeowners_proto_rawDescOnce.Do(func() {
		file_codeowners_proto_rawDescData = protoimpl.X.CompressGZIP(file_codeowners_proto_rawDescData)
	})
	return file_codeowners_proto_rawDescData
}

var file_codeowners_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_codeowners_proto_goTypes = []interface{}{
	(*File)(nil),  // 0: own.codeowners.v1.File
	(*Rule)(nil),  // 1: own.codeowners.v1.Rule
	(*Owner)(nil), // 2: own.codeowners.v1.Owner
}
var file_codeowners_proto_depIdxs = []int32{
	1, // 0: own.codeowners.v1.File.rule:type_name -> own.codeowners.v1.Rule
	2, // 1: own.codeowners.v1.Rule.owner:type_name -> own.codeowners.v1.Owner
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_codeowners_proto_init() }
func file_codeowners_proto_init() {
	if File_codeowners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codeowners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_codeowners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_codeowners_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
			RawDescriptor: file_codeowners_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_codeowners_proto_goTypes,
		DependencyIndexes: file_codeowners_proto_depIdxs,
		MessageInfos:      file_codeowners_proto_msgTypes,
	}.Build()
	File_codeowners_proto = out.File
	file_codeowners_proto_rawDesc = nil
	file_codeowners_proto_goTypes = nil
	file_codeowners_proto_depIdxs = nil
}
