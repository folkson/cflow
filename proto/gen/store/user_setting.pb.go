// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/user_setting.proto

package store

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

type UserSettingKey int32

const (
	UserSettingKey_USER_SETTING_KEY_UNSPECIFIED     UserSettingKey = 0
	UserSettingKey_USER_SETTING_ACCESS_TOKENS       UserSettingKey = 1
	UserSettingKey_USER_SETTING_LOCALE              UserSettingKey = 2
	UserSettingKey_USER_SETTING_APPEARANCE          UserSettingKey = 3
	UserSettingKey_USER_SETTING_MEMO_VISIBILITY     UserSettingKey = 4
	UserSettingKey_USER_SETTING_MARK_WITH_TAG       UserSettingKey = 5
	UserSettingKey_USER_SETTING_SHOW_WORD_CNT       UserSettingKey = 6
	UserSettingKey_USER_SETTING_CUSTOM_SHORTCUT     UserSettingKey = 7
	UserSettingKey_USER_SETTING_FAV_TAG             UserSettingKey = 8
	UserSettingKey_USER_SETTING_REF_PREVIEW         UserSettingKey = 9
	UserSettingKey_USER_SETTING_SHOW_TODO_PAGE      UserSettingKey = 10
	UserSettingKey_USER_SETTING_SHOW_ARCHIVE_PAGE   UserSettingKey = 11
	UserSettingKey_USER_SETTING_PASTE_RENAME        UserSettingKey = 12
	UserSettingKey_USER_SETTING_SHOW_TAG_SELECTOR   UserSettingKey = 13
	UserSettingKey_USER_SETTING_SHOW_MEMO_PUBLIC    UserSettingKey = 14
	UserSettingKey_USER_SETTING_CUSTOM_CARD_STYLE   UserSettingKey = 15
	UserSettingKey_USER_SETTING_DOUBLE_CLICK_EDIT   UserSettingKey = 16
	UserSettingKey_USER_SETTING_USE_EXCALIDRAW      UserSettingKey = 17
	UserSettingKey_USER_SETTING_HIDE_MARK_BLOCK     UserSettingKey = 18
	UserSettingKey_USER_SETTING_HIDE_FULL_SCREEN    UserSettingKey = 19
	UserSettingKey_USER_SETTING_SYS_SHORTCUT_CONFIG UserSettingKey = 20
)

// Enum value maps for UserSettingKey.
var (
	UserSettingKey_name = map[int32]string{
		0:  "USER_SETTING_KEY_UNSPECIFIED",
		1:  "USER_SETTING_ACCESS_TOKENS",
		2:  "USER_SETTING_LOCALE",
		3:  "USER_SETTING_APPEARANCE",
		4:  "USER_SETTING_MEMO_VISIBILITY",
		5:  "USER_SETTING_MARK_WITH_TAG",
		6:  "USER_SETTING_SHOW_WORD_CNT",
		7:  "USER_SETTING_CUSTOM_SHORTCUT",
		8:  "USER_SETTING_FAV_TAG",
		9:  "USER_SETTING_REF_PREVIEW",
		10: "USER_SETTING_SHOW_TODO_PAGE",
		11: "USER_SETTING_SHOW_ARCHIVE_PAGE",
		12: "USER_SETTING_PASTE_RENAME",
		13: "USER_SETTING_SHOW_TAG_SELECTOR",
		14: "USER_SETTING_SHOW_MEMO_PUBLIC",
		15: "USER_SETTING_CUSTOM_CARD_STYLE",
		16: "USER_SETTING_DOUBLE_CLICK_EDIT",
		17: "USER_SETTING_USE_EXCALIDRAW",
		18: "USER_SETTING_HIDE_MARK_BLOCK",
		19: "USER_SETTING_HIDE_FULL_SCREEN",
		20: "USER_SETTING_SYS_SHORTCUT_CONFIG",
	}
	UserSettingKey_value = map[string]int32{
		"USER_SETTING_KEY_UNSPECIFIED":     0,
		"USER_SETTING_ACCESS_TOKENS":       1,
		"USER_SETTING_LOCALE":              2,
		"USER_SETTING_APPEARANCE":          3,
		"USER_SETTING_MEMO_VISIBILITY":     4,
		"USER_SETTING_MARK_WITH_TAG":       5,
		"USER_SETTING_SHOW_WORD_CNT":       6,
		"USER_SETTING_CUSTOM_SHORTCUT":     7,
		"USER_SETTING_FAV_TAG":             8,
		"USER_SETTING_REF_PREVIEW":         9,
		"USER_SETTING_SHOW_TODO_PAGE":      10,
		"USER_SETTING_SHOW_ARCHIVE_PAGE":   11,
		"USER_SETTING_PASTE_RENAME":        12,
		"USER_SETTING_SHOW_TAG_SELECTOR":   13,
		"USER_SETTING_SHOW_MEMO_PUBLIC":    14,
		"USER_SETTING_CUSTOM_CARD_STYLE":   15,
		"USER_SETTING_DOUBLE_CLICK_EDIT":   16,
		"USER_SETTING_USE_EXCALIDRAW":      17,
		"USER_SETTING_HIDE_MARK_BLOCK":     18,
		"USER_SETTING_HIDE_FULL_SCREEN":    19,
		"USER_SETTING_SYS_SHORTCUT_CONFIG": 20,
	}
)

func (x UserSettingKey) Enum() *UserSettingKey {
	p := new(UserSettingKey)
	*p = x
	return p
}

func (x UserSettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_store_user_setting_proto_enumTypes[0].Descriptor()
}

func (UserSettingKey) Type() protoreflect.EnumType {
	return &file_store_user_setting_proto_enumTypes[0]
}

func (x UserSettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSettingKey.Descriptor instead.
func (UserSettingKey) EnumDescriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{0}
}

type UserSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32          `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Key    UserSettingKey `protobuf:"varint,2,opt,name=key,proto3,enum=memos.store.UserSettingKey" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*UserSetting_AccessTokens
	//	*UserSetting_Locale
	//	*UserSetting_Appearance
	//	*UserSetting_MemoVisibility
	//	*UserSetting_MarkWithTag
	//	*UserSetting_ShowWordCnt
	//	*UserSetting_CustomShortcut
	//	*UserSetting_FavTag
	//	*UserSetting_RefPreview
	//	*UserSetting_ShowTodoPage
	//	*UserSetting_ShowArchivePage
	//	*UserSetting_PasteRename
	//	*UserSetting_ShowTagSelector
	//	*UserSetting_ShowMemoPublic
	//	*UserSetting_CustomCardStyle
	//	*UserSetting_DoubleClickEdit
	//	*UserSetting_UseExcalidraw
	//	*UserSetting_HideMarkBlock
	//	*UserSetting_HideFullScreen
	//	*UserSetting_SysShortcutConfig
	Value isUserSetting_Value `protobuf_oneof:"value"`
}

func (x *UserSetting) Reset() {
	*x = UserSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_user_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSetting) ProtoMessage() {}

func (x *UserSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSetting.ProtoReflect.Descriptor instead.
func (*UserSetting) Descriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{0}
}

func (x *UserSetting) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserSetting) GetKey() UserSettingKey {
	if x != nil {
		return x.Key
	}
	return UserSettingKey_USER_SETTING_KEY_UNSPECIFIED
}

func (m *UserSetting) GetValue() isUserSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *UserSetting) GetAccessTokens() *AccessTokensUserSetting {
	if x, ok := x.GetValue().(*UserSetting_AccessTokens); ok {
		return x.AccessTokens
	}
	return nil
}

func (x *UserSetting) GetLocale() string {
	if x, ok := x.GetValue().(*UserSetting_Locale); ok {
		return x.Locale
	}
	return ""
}

func (x *UserSetting) GetAppearance() string {
	if x, ok := x.GetValue().(*UserSetting_Appearance); ok {
		return x.Appearance
	}
	return ""
}

func (x *UserSetting) GetMemoVisibility() string {
	if x, ok := x.GetValue().(*UserSetting_MemoVisibility); ok {
		return x.MemoVisibility
	}
	return ""
}

func (x *UserSetting) GetMarkWithTag() bool {
	if x, ok := x.GetValue().(*UserSetting_MarkWithTag); ok {
		return x.MarkWithTag
	}
	return false
}

func (x *UserSetting) GetShowWordCnt() bool {
	if x, ok := x.GetValue().(*UserSetting_ShowWordCnt); ok {
		return x.ShowWordCnt
	}
	return false
}

func (x *UserSetting) GetCustomShortcut() string {
	if x, ok := x.GetValue().(*UserSetting_CustomShortcut); ok {
		return x.CustomShortcut
	}
	return ""
}

func (x *UserSetting) GetFavTag() string {
	if x, ok := x.GetValue().(*UserSetting_FavTag); ok {
		return x.FavTag
	}
	return ""
}

func (x *UserSetting) GetRefPreview() bool {
	if x, ok := x.GetValue().(*UserSetting_RefPreview); ok {
		return x.RefPreview
	}
	return false
}

func (x *UserSetting) GetShowTodoPage() bool {
	if x, ok := x.GetValue().(*UserSetting_ShowTodoPage); ok {
		return x.ShowTodoPage
	}
	return false
}

func (x *UserSetting) GetShowArchivePage() bool {
	if x, ok := x.GetValue().(*UserSetting_ShowArchivePage); ok {
		return x.ShowArchivePage
	}
	return false
}

func (x *UserSetting) GetPasteRename() bool {
	if x, ok := x.GetValue().(*UserSetting_PasteRename); ok {
		return x.PasteRename
	}
	return false
}

func (x *UserSetting) GetShowTagSelector() bool {
	if x, ok := x.GetValue().(*UserSetting_ShowTagSelector); ok {
		return x.ShowTagSelector
	}
	return false
}

func (x *UserSetting) GetShowMemoPublic() bool {
	if x, ok := x.GetValue().(*UserSetting_ShowMemoPublic); ok {
		return x.ShowMemoPublic
	}
	return false
}

func (x *UserSetting) GetCustomCardStyle() string {
	if x, ok := x.GetValue().(*UserSetting_CustomCardStyle); ok {
		return x.CustomCardStyle
	}
	return ""
}

func (x *UserSetting) GetDoubleClickEdit() bool {
	if x, ok := x.GetValue().(*UserSetting_DoubleClickEdit); ok {
		return x.DoubleClickEdit
	}
	return false
}

func (x *UserSetting) GetUseExcalidraw() bool {
	if x, ok := x.GetValue().(*UserSetting_UseExcalidraw); ok {
		return x.UseExcalidraw
	}
	return false
}

func (x *UserSetting) GetHideMarkBlock() bool {
	if x, ok := x.GetValue().(*UserSetting_HideMarkBlock); ok {
		return x.HideMarkBlock
	}
	return false
}

func (x *UserSetting) GetHideFullScreen() bool {
	if x, ok := x.GetValue().(*UserSetting_HideFullScreen); ok {
		return x.HideFullScreen
	}
	return false
}

func (x *UserSetting) GetSysShortcutConfig() string {
	if x, ok := x.GetValue().(*UserSetting_SysShortcutConfig); ok {
		return x.SysShortcutConfig
	}
	return ""
}

type isUserSetting_Value interface {
	isUserSetting_Value()
}

type UserSetting_AccessTokens struct {
	AccessTokens *AccessTokensUserSetting `protobuf:"bytes,3,opt,name=access_tokens,json=accessTokens,proto3,oneof"`
}

type UserSetting_Locale struct {
	Locale string `protobuf:"bytes,4,opt,name=locale,proto3,oneof"`
}

type UserSetting_Appearance struct {
	Appearance string `protobuf:"bytes,5,opt,name=appearance,proto3,oneof"`
}

type UserSetting_MemoVisibility struct {
	MemoVisibility string `protobuf:"bytes,6,opt,name=memo_visibility,json=memoVisibility,proto3,oneof"`
}

type UserSetting_MarkWithTag struct {
	MarkWithTag bool `protobuf:"varint,7,opt,name=mark_with_tag,json=markWithTag,proto3,oneof"`
}

type UserSetting_ShowWordCnt struct {
	ShowWordCnt bool `protobuf:"varint,8,opt,name=show_word_cnt,json=showWordCnt,proto3,oneof"`
}

type UserSetting_CustomShortcut struct {
	CustomShortcut string `protobuf:"bytes,9,opt,name=custom_shortcut,json=customShortcut,proto3,oneof"`
}

type UserSetting_FavTag struct {
	FavTag string `protobuf:"bytes,10,opt,name=fav_tag,json=favTag,proto3,oneof"`
}

type UserSetting_RefPreview struct {
	RefPreview bool `protobuf:"varint,11,opt,name=ref_preview,json=refPreview,proto3,oneof"`
}

type UserSetting_ShowTodoPage struct {
	ShowTodoPage bool `protobuf:"varint,12,opt,name=show_todo_page,json=showTodoPage,proto3,oneof"`
}

type UserSetting_ShowArchivePage struct {
	ShowArchivePage bool `protobuf:"varint,13,opt,name=show_archive_page,json=showArchivePage,proto3,oneof"`
}

type UserSetting_PasteRename struct {
	PasteRename bool `protobuf:"varint,14,opt,name=paste_rename,json=pasteRename,proto3,oneof"`
}

type UserSetting_ShowTagSelector struct {
	ShowTagSelector bool `protobuf:"varint,15,opt,name=show_tag_selector,json=showTagSelector,proto3,oneof"`
}

type UserSetting_ShowMemoPublic struct {
	ShowMemoPublic bool `protobuf:"varint,16,opt,name=show_memo_public,json=showMemoPublic,proto3,oneof"`
}

type UserSetting_CustomCardStyle struct {
	CustomCardStyle string `protobuf:"bytes,17,opt,name=custom_card_style,json=customCardStyle,proto3,oneof"`
}

type UserSetting_DoubleClickEdit struct {
	DoubleClickEdit bool `protobuf:"varint,18,opt,name=double_click_edit,json=doubleClickEdit,proto3,oneof"`
}

type UserSetting_UseExcalidraw struct {
	UseExcalidraw bool `protobuf:"varint,19,opt,name=use_excalidraw,json=useExcalidraw,proto3,oneof"`
}

type UserSetting_HideMarkBlock struct {
	HideMarkBlock bool `protobuf:"varint,20,opt,name=hide_mark_block,json=hideMarkBlock,proto3,oneof"`
}

type UserSetting_HideFullScreen struct {
	HideFullScreen bool `protobuf:"varint,21,opt,name=hide_full_screen,json=hideFullScreen,proto3,oneof"`
}

type UserSetting_SysShortcutConfig struct {
	SysShortcutConfig string `protobuf:"bytes,22,opt,name=sys_shortcut_config,json=sysShortcutConfig,proto3,oneof"`
}

func (*UserSetting_AccessTokens) isUserSetting_Value() {}

func (*UserSetting_Locale) isUserSetting_Value() {}

func (*UserSetting_Appearance) isUserSetting_Value() {}

func (*UserSetting_MemoVisibility) isUserSetting_Value() {}

func (*UserSetting_MarkWithTag) isUserSetting_Value() {}

func (*UserSetting_ShowWordCnt) isUserSetting_Value() {}

func (*UserSetting_CustomShortcut) isUserSetting_Value() {}

func (*UserSetting_FavTag) isUserSetting_Value() {}

func (*UserSetting_RefPreview) isUserSetting_Value() {}

func (*UserSetting_ShowTodoPage) isUserSetting_Value() {}

func (*UserSetting_ShowArchivePage) isUserSetting_Value() {}

func (*UserSetting_PasteRename) isUserSetting_Value() {}

func (*UserSetting_ShowTagSelector) isUserSetting_Value() {}

func (*UserSetting_ShowMemoPublic) isUserSetting_Value() {}

func (*UserSetting_CustomCardStyle) isUserSetting_Value() {}

func (*UserSetting_DoubleClickEdit) isUserSetting_Value() {}

func (*UserSetting_UseExcalidraw) isUserSetting_Value() {}

func (*UserSetting_HideMarkBlock) isUserSetting_Value() {}

func (*UserSetting_HideFullScreen) isUserSetting_Value() {}

func (*UserSetting_SysShortcutConfig) isUserSetting_Value() {}

type AccessTokensUserSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessTokens []*AccessTokensUserSetting_AccessToken `protobuf:"bytes,1,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
}

func (x *AccessTokensUserSetting) Reset() {
	*x = AccessTokensUserSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_user_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokensUserSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokensUserSetting) ProtoMessage() {}

func (x *AccessTokensUserSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokensUserSetting.ProtoReflect.Descriptor instead.
func (*AccessTokensUserSetting) Descriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{1}
}

func (x *AccessTokensUserSetting) GetAccessTokens() []*AccessTokensUserSetting_AccessToken {
	if x != nil {
		return x.AccessTokens
	}
	return nil
}

type AccessTokensUserSetting_AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The access token is a JWT token.
	// Including expiration time, issuer, etc.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// A description for the access token.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AccessTokensUserSetting_AccessToken) Reset() {
	*x = AccessTokensUserSetting_AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_user_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokensUserSetting_AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokensUserSetting_AccessToken) ProtoMessage() {}

func (x *AccessTokensUserSetting_AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokensUserSetting_AccessToken.ProtoReflect.Descriptor instead.
func (*AccessTokensUserSetting_AccessToken) Descriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AccessTokensUserSetting_AccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AccessTokensUserSetting_AccessToken) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_store_user_setting_proto protoreflect.FileDescriptor

var file_store_user_setting_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xa9, 0x07, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x4b, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x6f,
	0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x61,
	0x72, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x6f,
	0x77, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x57, 0x6f, 0x72, 0x64, 0x43, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63,
	0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x07, 0x66, 0x61,
	0x76, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x66,
	0x61, 0x76, 0x54, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x5f, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x77,
	0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x73,
	0x68, 0x6f, 0x77, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x0a, 0x0c, 0x70, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x74, 0x61, 0x67, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x0f, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x61, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0e, 0x73,
	0x68, 0x6f, 0x77, 0x4d, 0x65, 0x6d, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x2c, 0x0a,
	0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x79,
	0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x65, 0x64, 0x69, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x45, 0x64, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x75, 0x73, 0x65,
	0x5f, 0x65, 0x78, 0x63, 0x61, 0x6c, 0x69, 0x64, 0x72, 0x61, 0x77, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x45, 0x78, 0x63, 0x61, 0x6c, 0x69, 0x64, 0x72,
	0x61, 0x77, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x68,
	0x69, 0x64, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x10,
	0x68, 0x69, 0x64, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0e, 0x68, 0x69, 0x64, 0x65, 0x46, 0x75,
	0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x79, 0x73, 0x5f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x73, 0x79, 0x73, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x63, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x55, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x52, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xc3, 0x05, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a,
	0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x53, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x50, 0x50, 0x45, 0x41, 0x52, 0x41,
	0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x05, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x43, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x53,
	0x48, 0x4f, 0x52, 0x54, 0x43, 0x55, 0x54, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x41, 0x56, 0x5f, 0x54, 0x41,
	0x47, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54,
	0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x46, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10,
	0x09, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x50, 0x41, 0x47, 0x45,
	0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x5f,
	0x50, 0x41, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53,
	0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x53, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x0c, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x53,
	0x45, 0x4c, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x10, 0x0d, 0x12, 0x21, 0x0a, 0x1d, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x4d,
	0x45, 0x4d, 0x4f, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x0e, 0x12, 0x22, 0x0a, 0x1e,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x10, 0x0f,
	0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x45, 0x44,
	0x49, 0x54, 0x10, 0x10, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x54,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x41, 0x4c, 0x49, 0x44,
	0x52, 0x41, 0x57, 0x10, 0x11, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x49, 0x44, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x12, 0x12, 0x21, 0x0a, 0x1d, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x49, 0x44, 0x45, 0x5f, 0x46, 0x55, 0x4c,
	0x4c, 0x5f, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x13, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x53,
	0x48, 0x4f, 0x52, 0x54, 0x43, 0x55, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x14,
	0x42, 0x9b, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_user_setting_proto_rawDescOnce sync.Once
	file_store_user_setting_proto_rawDescData = file_store_user_setting_proto_rawDesc
)

func file_store_user_setting_proto_rawDescGZIP() []byte {
	file_store_user_setting_proto_rawDescOnce.Do(func() {
		file_store_user_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_user_setting_proto_rawDescData)
	})
	return file_store_user_setting_proto_rawDescData
}

var file_store_user_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_user_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_user_setting_proto_goTypes = []interface{}{
	(UserSettingKey)(0),                         // 0: memos.store.UserSettingKey
	(*UserSetting)(nil),                         // 1: memos.store.UserSetting
	(*AccessTokensUserSetting)(nil),             // 2: memos.store.AccessTokensUserSetting
	(*AccessTokensUserSetting_AccessToken)(nil), // 3: memos.store.AccessTokensUserSetting.AccessToken
}
var file_store_user_setting_proto_depIdxs = []int32{
	0, // 0: memos.store.UserSetting.key:type_name -> memos.store.UserSettingKey
	2, // 1: memos.store.UserSetting.access_tokens:type_name -> memos.store.AccessTokensUserSetting
	3, // 2: memos.store.AccessTokensUserSetting.access_tokens:type_name -> memos.store.AccessTokensUserSetting.AccessToken
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_store_user_setting_proto_init() }
func file_store_user_setting_proto_init() {
	if File_store_user_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_user_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSetting); i {
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
		file_store_user_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokensUserSetting); i {
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
		file_store_user_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokensUserSetting_AccessToken); i {
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
	file_store_user_setting_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UserSetting_AccessTokens)(nil),
		(*UserSetting_Locale)(nil),
		(*UserSetting_Appearance)(nil),
		(*UserSetting_MemoVisibility)(nil),
		(*UserSetting_MarkWithTag)(nil),
		(*UserSetting_ShowWordCnt)(nil),
		(*UserSetting_CustomShortcut)(nil),
		(*UserSetting_FavTag)(nil),
		(*UserSetting_RefPreview)(nil),
		(*UserSetting_ShowTodoPage)(nil),
		(*UserSetting_ShowArchivePage)(nil),
		(*UserSetting_PasteRename)(nil),
		(*UserSetting_ShowTagSelector)(nil),
		(*UserSetting_ShowMemoPublic)(nil),
		(*UserSetting_CustomCardStyle)(nil),
		(*UserSetting_DoubleClickEdit)(nil),
		(*UserSetting_UseExcalidraw)(nil),
		(*UserSetting_HideMarkBlock)(nil),
		(*UserSetting_HideFullScreen)(nil),
		(*UserSetting_SysShortcutConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_user_setting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_user_setting_proto_goTypes,
		DependencyIndexes: file_store_user_setting_proto_depIdxs,
		EnumInfos:         file_store_user_setting_proto_enumTypes,
		MessageInfos:      file_store_user_setting_proto_msgTypes,
	}.Build()
	File_store_user_setting_proto = out.File
	file_store_user_setting_proto_rawDesc = nil
	file_store_user_setting_proto_goTypes = nil
	file_store_user_setting_proto_depIdxs = nil
}
