// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/rentalmgmt/rentalmgmt.proto

package rentalmgmt

import (
	fmt "fmt"
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

//USER'S SIDE
//Create New User
type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone_Number         string   `protobuf:"bytes,4,opt,name=Phone_Number,json=PhoneNumber,proto3" json:"Phone_Number,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
	District             string   `protobuf:"bytes,6,opt,name=District,proto3" json:"District,omitempty"`
	Postal_Code          string   `protobuf:"bytes,7,opt,name=Postal_Code,json=PostalCode,proto3" json:"Postal_Code,omitempty"`
	Country              string   `protobuf:"bytes,8,opt,name=Country,proto3" json:"Country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone_Number() string {
	if m != nil {
		return m.Phone_Number
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *User) GetPostal_Code() string {
	if m != nil {
		return m.Postal_Code
	}
	return ""
}

func (m *User) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

//Get the User Id
type UserId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{1}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//Get the Item Id
type ItemId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemId) Reset()         { *m = ItemId{} }
func (m *ItemId) String() string { return proto.CompactTextString(m) }
func (*ItemId) ProtoMessage()    {}
func (*ItemId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{2}
}

func (m *ItemId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemId.Unmarshal(m, b)
}
func (m *ItemId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemId.Marshal(b, m, deterministic)
}
func (m *ItemId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemId.Merge(m, src)
}
func (m *ItemId) XXX_Size() int {
	return xxx_messageInfo_ItemId.Size(m)
}
func (m *ItemId) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemId.DiscardUnknown(m)
}

var xxx_messageInfo_ItemId proto.InternalMessageInfo

func (m *ItemId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//Get a Particular Item with all Details
type Item struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Category             string   `protobuf:"bytes,4,opt,name=Category,proto3" json:"Category,omitempty"`
	Available_From       string   `protobuf:"bytes,5,opt,name=Available_From,json=AvailableFrom,proto3" json:"Available_From,omitempty"`
	Available_To         string   `protobuf:"bytes,6,opt,name=Available_To,json=AvailableTo,proto3" json:"Available_To,omitempty"`
	AmountPerDay         int32    `protobuf:"varint,7,opt,name=Amount_per_day,json=AmountPerDay,proto3" json:"Amount_per_day,omitempty"`
	UserID               int32    `protobuf:"varint,8,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{3}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Item) GetAvailable_From() string {
	if m != nil {
		return m.Available_From
	}
	return ""
}

func (m *Item) GetAvailable_To() string {
	if m != nil {
		return m.Available_To
	}
	return ""
}

func (m *Item) GetAmountPerDay() int32 {
	if m != nil {
		return m.AmountPerDay
	}
	return 0
}

func (m *Item) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

//To Show the Available Dates
type AvailableDates struct {
	Available_From       string   `protobuf:"bytes,1,opt,name=Available_From,json=AvailableFrom,proto3" json:"Available_From,omitempty"`
	Available_To         string   `protobuf:"bytes,2,opt,name=Available_To,json=AvailableTo,proto3" json:"Available_To,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailableDates) Reset()         { *m = AvailableDates{} }
func (m *AvailableDates) String() string { return proto.CompactTextString(m) }
func (*AvailableDates) ProtoMessage()    {}
func (*AvailableDates) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{4}
}

func (m *AvailableDates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableDates.Unmarshal(m, b)
}
func (m *AvailableDates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableDates.Marshal(b, m, deterministic)
}
func (m *AvailableDates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableDates.Merge(m, src)
}
func (m *AvailableDates) XXX_Size() int {
	return xxx_messageInfo_AvailableDates.Size(m)
}
func (m *AvailableDates) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableDates.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableDates proto.InternalMessageInfo

func (m *AvailableDates) GetAvailable_From() string {
	if m != nil {
		return m.Available_From
	}
	return ""
}

func (m *AvailableDates) GetAvailable_To() string {
	if m != nil {
		return m.Available_To
	}
	return ""
}

//For Listing all the Items
type Items struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Items) Reset()         { *m = Items{} }
func (m *Items) String() string { return proto.CompactTextString(m) }
func (*Items) ProtoMessage()    {}
func (*Items) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{5}
}

func (m *Items) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Items.Unmarshal(m, b)
}
func (m *Items) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Items.Marshal(b, m, deterministic)
}
func (m *Items) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Items.Merge(m, src)
}
func (m *Items) XXX_Size() int {
	return xxx_messageInfo_Items.Size(m)
}
func (m *Items) XXX_DiscardUnknown() {
	xxx_messageInfo_Items.DiscardUnknown(m)
}

var xxx_messageInfo_Items proto.InternalMessageInfo

func (m *Items) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

//REQUEST AND RESPONSE
//For Requesting
type Request struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{6}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

//To Conform Item is booked
type BookingId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingId) Reset()         { *m = BookingId{} }
func (m *BookingId) String() string { return proto.CompactTextString(m) }
func (*BookingId) ProtoMessage()    {}
func (*BookingId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{7}
}

func (m *BookingId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingId.Unmarshal(m, b)
}
func (m *BookingId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingId.Marshal(b, m, deterministic)
}
func (m *BookingId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingId.Merge(m, src)
}
func (m *BookingId) XXX_Size() int {
	return xxx_messageInfo_BookingId.Size(m)
}
func (m *BookingId) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingId.DiscardUnknown(m)
}

var xxx_messageInfo_BookingId proto.InternalMessageInfo

func (m *BookingId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Booking struct {
	StartDate            string   `protobuf:"bytes,1,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string   `protobuf:"bytes,2,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	UserId               int32    `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ItemId               int32    `protobuf:"varint,4,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Booking) Reset()         { *m = Booking{} }
func (m *Booking) String() string { return proto.CompactTextString(m) }
func (*Booking) ProtoMessage()    {}
func (*Booking) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{8}
}

func (m *Booking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Booking.Unmarshal(m, b)
}
func (m *Booking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Booking.Marshal(b, m, deterministic)
}
func (m *Booking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Booking.Merge(m, src)
}
func (m *Booking) XXX_Size() int {
	return xxx_messageInfo_Booking.Size(m)
}
func (m *Booking) XXX_DiscardUnknown() {
	xxx_messageInfo_Booking.DiscardUnknown(m)
}

var xxx_messageInfo_Booking proto.InternalMessageInfo

func (m *Booking) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Booking) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *Booking) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Booking) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

//To Show the Booked Items
type BookedItems struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ItemName             string   `protobuf:"bytes,2,opt,name=ItemName,proto3" json:"ItemName,omitempty"`
	StartDate            string   `protobuf:"bytes,3,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              string   `protobuf:"bytes,4,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	TotalAmount          int32    `protobuf:"varint,5,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookedItems) Reset()         { *m = BookedItems{} }
func (m *BookedItems) String() string { return proto.CompactTextString(m) }
func (*BookedItems) ProtoMessage()    {}
func (*BookedItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{9}
}

func (m *BookedItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookedItems.Unmarshal(m, b)
}
func (m *BookedItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookedItems.Marshal(b, m, deterministic)
}
func (m *BookedItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookedItems.Merge(m, src)
}
func (m *BookedItems) XXX_Size() int {
	return xxx_messageInfo_BookedItems.Size(m)
}
func (m *BookedItems) XXX_DiscardUnknown() {
	xxx_messageInfo_BookedItems.DiscardUnknown(m)
}

var xxx_messageInfo_BookedItems proto.InternalMessageInfo

func (m *BookedItems) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BookedItems) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *BookedItems) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *BookedItems) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *BookedItems) GetTotalAmount() int32 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

type ReviewId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReviewId) Reset()         { *m = ReviewId{} }
func (m *ReviewId) String() string { return proto.CompactTextString(m) }
func (*ReviewId) ProtoMessage()    {}
func (*ReviewId) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{10}
}

func (m *ReviewId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewId.Unmarshal(m, b)
}
func (m *ReviewId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewId.Marshal(b, m, deterministic)
}
func (m *ReviewId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewId.Merge(m, src)
}
func (m *ReviewId) XXX_Size() int {
	return xxx_messageInfo_ReviewId.Size(m)
}
func (m *ReviewId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewId.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewId proto.InternalMessageInfo

func (m *ReviewId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Review struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Rating               int32    `protobuf:"varint,3,opt,name=Rating,proto3" json:"Rating,omitempty"`
	UserId               int32    `protobuf:"varint,4,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ItemId               int32    `protobuf:"varint,5,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_3231b20ba83710a0, []int{11}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Review) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Review) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Review) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Review) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "rentalmgmt.User")
	proto.RegisterType((*UserId)(nil), "rentalmgmt.UserId")
	proto.RegisterType((*ItemId)(nil), "rentalmgmt.ItemId")
	proto.RegisterType((*Item)(nil), "rentalmgmt.Item")
	proto.RegisterType((*AvailableDates)(nil), "rentalmgmt.AvailableDates")
	proto.RegisterType((*Items)(nil), "rentalmgmt.Items")
	proto.RegisterType((*Request)(nil), "rentalmgmt.Request")
	proto.RegisterType((*BookingId)(nil), "rentalmgmt.BookingId")
	proto.RegisterType((*Booking)(nil), "rentalmgmt.Booking")
	proto.RegisterType((*BookedItems)(nil), "rentalmgmt.BookedItems")
	proto.RegisterType((*ReviewId)(nil), "rentalmgmt.ReviewId")
	proto.RegisterType((*Review)(nil), "rentalmgmt.Review")
}

func init() { proto.RegisterFile("pkg/rentalmgmt/rentalmgmt.proto", fileDescriptor_3231b20ba83710a0) }

var fileDescriptor_3231b20ba83710a0 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x4e, 0xdb, 0x40,
	0x10, 0xc5, 0x21, 0xce, 0x65, 0x1c, 0x50, 0xb5, 0xa5, 0xad, 0x15, 0x5a, 0x91, 0xba, 0x17, 0xf1,
	0x04, 0x15, 0xbd, 0xa1, 0xb6, 0xaa, 0x14, 0xe2, 0x80, 0x22, 0x55, 0x08, 0xb9, 0xf0, 0xc2, 0x8b,
	0xb5, 0xb0, 0xab, 0x74, 0x85, 0x2f, 0x61, 0xbd, 0xa1, 0x4a, 0x7f, 0xa3, 0x1f, 0xd6, 0x97, 0x7e,
	0x48, 0x3f, 0xa1, 0xda, 0x4b, 0x1c, 0x87, 0x98, 0x5e, 0x78, 0xdb, 0x39, 0x67, 0x67, 0x67, 0xe6,
	0xcc, 0x8c, 0x0d, 0x1b, 0xa3, 0x8b, 0xe1, 0x36, 0xa7, 0x89, 0xc0, 0x51, 0x3c, 0x8c, 0x45, 0xe1,
	0xb8, 0x35, 0xe2, 0xa9, 0x48, 0x11, 0xcc, 0x10, 0xef, 0xa7, 0x05, 0xd5, 0x93, 0x8c, 0x72, 0xb4,
	0x0a, 0x95, 0x01, 0x71, 0xad, 0x8e, 0xb5, 0x69, 0x07, 0x95, 0x01, 0x41, 0x08, 0xaa, 0x87, 0x38,
	0xa6, 0x6e, 0xa5, 0x63, 0x6d, 0x36, 0x03, 0x75, 0x46, 0x6b, 0x60, 0xf7, 0x63, 0xcc, 0x22, 0x77,
	0x59, 0x81, 0xda, 0x40, 0x8f, 0xa1, 0x75, 0xf4, 0x25, 0x4d, 0x68, 0x78, 0x38, 0x8e, 0xcf, 0x28,
	0x77, 0xab, 0x8a, 0x74, 0x14, 0xa6, 0x21, 0xe4, 0x42, 0xbd, 0x4b, 0x08, 0xa7, 0x59, 0xe6, 0xda,
	0x8a, 0x9d, 0x9a, 0xa8, 0x0d, 0x0d, 0x9f, 0x65, 0x82, 0xb3, 0x73, 0xe1, 0xd6, 0x14, 0x95, 0xdb,
	0x68, 0x03, 0x9c, 0xa3, 0x34, 0x13, 0x38, 0x0a, 0x7b, 0x29, 0xa1, 0x6e, 0x5d, 0xd1, 0xa0, 0x21,
	0x89, 0xc8, 0x67, 0x7b, 0xe9, 0x38, 0x11, 0x7c, 0xe2, 0x36, 0xf4, 0xb3, 0xc6, 0xf4, 0x5c, 0xa8,
	0xc9, 0xaa, 0x06, 0x44, 0xd6, 0xc5, 0xf2, 0xba, 0x18, 0x91, 0xcc, 0x40, 0xd0, 0xb8, 0x84, 0xf9,
	0x65, 0x41, 0x55, 0x52, 0x4a, 0x0a, 0x3f, 0x97, 0xc2, 0x2f, 0x95, 0xa2, 0x03, 0x8e, 0x4f, 0xb3,
	0x73, 0xce, 0x46, 0x82, 0xa5, 0x89, 0x11, 0xa4, 0x08, 0xc9, 0xca, 0x7a, 0x58, 0xd0, 0x61, 0xca,
	0x27, 0x46, 0x92, 0xdc, 0x46, 0xcf, 0x60, 0xb5, 0x7b, 0x85, 0x59, 0x84, 0xcf, 0x22, 0x1a, 0xee,
	0xf3, 0x34, 0x36, 0xb2, 0xac, 0xe4, 0xa8, 0x04, 0xa5, 0xb2, 0xb3, 0x6b, 0xc7, 0xa9, 0x11, 0xc8,
	0xc9, 0xb1, 0xe3, 0x14, 0x3d, 0x85, 0xd5, 0x6e, 0x2c, 0x8b, 0x0e, 0x47, 0x94, 0x87, 0x04, 0x4f,
	0x94, 0x4c, 0x76, 0xd0, 0xd2, 0xe8, 0x11, 0xe5, 0x3e, 0x9e, 0xa0, 0xfb, 0x46, 0x0e, 0x5f, 0xe9,
	0x64, 0x07, 0xc6, 0xf2, 0x4e, 0x0b, 0x79, 0xf8, 0x58, 0xd0, 0xac, 0x24, 0x33, 0xeb, 0x5f, 0x32,
	0xab, 0x2c, 0x64, 0xe6, 0x6d, 0x83, 0x2d, 0xd5, 0xcc, 0xd0, 0x73, 0xb0, 0x99, 0x3c, 0xb8, 0x56,
	0x67, 0x79, 0xd3, 0xd9, 0xb9, 0xb3, 0x55, 0x18, 0x48, 0x79, 0x23, 0xd0, 0xb4, 0xf7, 0x04, 0xea,
	0x01, 0xbd, 0x1c, 0xd3, 0x4c, 0xc8, 0xc6, 0x72, 0x7d, 0x34, 0xe1, 0xa7, 0xa6, 0xb7, 0x0e, 0xcd,
	0xbd, 0x34, 0xbd, 0x60, 0xc9, 0x50, 0x77, 0xb0, 0x38, 0xb3, 0xde, 0x25, 0xd4, 0x0d, 0x89, 0x1e,
	0x42, 0xf3, 0xb3, 0xc0, 0x5c, 0xc8, 0xaa, 0xcc, 0x1b, 0x33, 0x40, 0xbe, 0xdf, 0x4f, 0x88, 0xe2,
	0x74, 0xe6, 0x53, 0x33, 0x57, 0x8a, 0xa8, 0x96, 0x4e, 0x95, 0x22, 0x12, 0xd7, 0x63, 0xa3, 0x7a,
	0x69, 0x07, 0xc6, 0xf2, 0xbe, 0x5b, 0xe0, 0xc8, 0x98, 0x94, 0xe8, 0x62, 0xaf, 0xaf, 0x51, 0x1b,
	0x1a, 0x92, 0x28, 0xcc, 0x4f, 0x6e, 0xcf, 0xe7, 0xb8, 0xfc, 0x87, 0x1c, 0xab, 0xf3, 0x39, 0x76,
	0xc0, 0x39, 0x4e, 0x05, 0x8e, 0x74, 0x8b, 0xd5, 0xe8, 0xd8, 0x41, 0x11, 0xf2, 0xda, 0xd0, 0x08,
	0xe8, 0x15, 0xa3, 0x5f, 0x4b, 0x44, 0xfa, 0x06, 0x35, 0xcd, 0x2d, 0xe4, 0xaa, 0xd6, 0x29, 0x8e,
	0x69, 0x22, 0xa6, 0xaa, 0x18, 0x53, 0x56, 0x1f, 0x60, 0xc1, 0x92, 0xe1, 0x54, 0x15, 0x6d, 0x15,
	0xd4, 0xaa, 0xde, 0xa0, 0x96, 0x5d, 0x54, 0x6b, 0xe7, 0x47, 0x0d, 0xd6, 0x03, 0xd5, 0xfd, 0xb0,
	0x8f, 0xb3, 0x49, 0xb8, 0x3f, 0x4e, 0xce, 0xe5, 0xb2, 0xe0, 0x88, 0x09, 0x46, 0x33, 0xf4, 0x0a,
	0xa0, 0xc7, 0x29, 0x16, 0x54, 0x7d, 0x92, 0xe6, 0x26, 0x45, 0x22, 0x6d, 0x74, 0x1d, 0x19, 0x10,
	0x6f, 0x49, 0x7a, 0x9d, 0x8c, 0xc8, 0x2d, 0xbc, 0x74, 0x2c, 0xb5, 0xf3, 0x0b, 0x53, 0x39, 0xef,
	0x65, 0xba, 0xbd, 0x84, 0x76, 0xc1, 0x39, 0xa0, 0xa2, 0x1b, 0x45, 0xba, 0xdd, 0x77, 0x8b, 0x97,
	0xcc, 0xf4, 0xb6, 0x17, 0xde, 0xf2, 0x96, 0x5e, 0x58, 0xe8, 0xb5, 0xf2, 0x94, 0xe6, 0xde, 0x44,
	0x7e, 0x5f, 0x17, 0x9f, 0x2f, 0x73, 0x9c, 0x15, 0xf7, 0x5f, 0x69, 0xbe, 0x01, 0xf0, 0x69, 0x44,
	0x8d, 0x57, 0x59, 0xac, 0x72, 0xbf, 0x8f, 0x80, 0x0e, 0xa8, 0x90, 0x1a, 0x7d, 0xa2, 0x38, 0x9b,
	0x0e, 0x75, 0x89, 0x80, 0x37, 0x14, 0xb9, 0x0b, 0x0d, 0xb9, 0x0d, 0x2a, 0xea, 0x9c, 0x36, 0x66,
	0x2f, 0xdb, 0xf7, 0x4a, 0x40, 0x15, 0xb9, 0x9f, 0x47, 0x2e, 0xae, 0x53, 0x59, 0xe4, 0x07, 0xd7,
	0x9f, 0x30, 0x97, 0x55, 0x02, 0x6f, 0xa1, 0xd9, 0x25, 0xc4, 0x0c, 0x38, 0x9a, 0xef, 0x8e, 0xc4,
	0xda, 0x6b, 0x8b, 0x98, 0x8a, 0xff, 0x0e, 0x5a, 0x5a, 0xe7, 0x5b, 0xf8, 0x7e, 0x80, 0x96, 0x56,
	0xdb, 0xf8, 0x96, 0xde, 0xbb, 0xd1, 0xfb, 0x3d, 0xac, 0xe8, 0x91, 0xd2, 0x58, 0xf6, 0xf7, 0x76,
	0xe9, 0x8b, 0xb2, 0xde, 0xbd, 0x47, 0xa7, 0xeb, 0x9a, 0x08, 0x29, 0xce, 0x26, 0x5b, 0x2c, 0xd9,
	0x8e, 0x0b, 0x3f, 0xfc, 0xb3, 0x9a, 0xfa, 0xe3, 0xbf, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x09, 0x34, 0x3d, 0x14, 0x08, 0x00, 0x00,
}
