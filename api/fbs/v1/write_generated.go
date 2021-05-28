// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package v1

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type FieldValue byte

const (
	FieldValueNONE   FieldValue = 0
	FieldValueString FieldValue = 1
	FieldValueUint   FieldValue = 2
)

var EnumNamesFieldValue = map[FieldValue]string{
	FieldValueNONE:   "NONE",
	FieldValueString: "String",
	FieldValueUint:   "Uint",
}

var EnumValuesFieldValue = map[string]FieldValue{
	"NONE":   FieldValueNONE,
	"String": FieldValueString,
	"Uint":   FieldValueUint,
}

func (v FieldValue) String() string {
	if s, ok := EnumNamesFieldValue[v]; ok {
		return s
	}
	return "FieldValue(" + strconv.FormatInt(int64(v), 10) + ")"
}

type String struct {
	_tab flatbuffers.Table
}

func GetRootAsString(buf []byte, offset flatbuffers.UOffsetT) *String {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &String{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsString(buf []byte, offset flatbuffers.UOffsetT) *String {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &String{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *String) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *String) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *String) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StringStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StringAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func StringEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Uint struct {
	_tab flatbuffers.Table
}

func GetRootAsUint(buf []byte, offset flatbuffers.UOffsetT) *Uint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uint{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUint(buf []byte, offset flatbuffers.UOffsetT) *Uint {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Uint{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Uint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Uint) Value() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uint) MutateValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func UintStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UintAddValue(builder *flatbuffers.Builder, value int64) {
	builder.PrependInt64Slot(0, value, 0)
}
func UintEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Field struct {
	_tab flatbuffers.Table
}

func GetRootAsField(buf []byte, offset flatbuffers.UOffsetT) *Field {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Field{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsField(buf []byte, offset flatbuffers.UOffsetT) *Field {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Field{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Field) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Field) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Field) ValueType() FieldValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return FieldValue(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Field) MutateValueType(n FieldValue) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *Field) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func FieldStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FieldAddValueType(builder *flatbuffers.Builder, valueType FieldValue) {
	builder.PrependByteSlot(0, byte(valueType), 0)
}
func FieldAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func FieldEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type entityValue struct {
	_tab flatbuffers.Table
}

func GetRootAsentityValue(buf []byte, offset flatbuffers.UOffsetT) *entityValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &entityValue{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsentityValue(buf []byte, offset flatbuffers.UOffsetT) *entityValue {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &entityValue{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *entityValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *entityValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *entityValue) EntityId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *entityValue) TimestampNanoseconds() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *entityValue) MutateTimestampNanoseconds(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *entityValue) DataBinary(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *entityValue) DataBinaryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *entityValue) DataBinaryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *entityValue) MutateDataBinary(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *entityValue) Fields(obj *Field, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *entityValue) FieldsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func entityValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func entityValueAddEntityId(builder *flatbuffers.Builder, entityId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(entityId), 0)
}
func entityValueAddTimestampNanoseconds(builder *flatbuffers.Builder, timestampNanoseconds uint64) {
	builder.PrependUint64Slot(1, timestampNanoseconds, 0)
}
func entityValueAddDataBinary(builder *flatbuffers.Builder, dataBinary flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dataBinary), 0)
}
func entityValueStartDataBinaryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func entityValueAddFields(builder *flatbuffers.Builder, fields flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(fields), 0)
}
func entityValueStartFieldsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func entityValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type WriteEntity struct {
	_tab flatbuffers.Table
}

func GetRootAsWriteEntity(buf []byte, offset flatbuffers.UOffsetT) *WriteEntity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WriteEntity{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWriteEntity(buf []byte, offset flatbuffers.UOffsetT) *WriteEntity {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WriteEntity{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WriteEntity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WriteEntity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WriteEntity) MateData(obj *Metadata) *Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *WriteEntity) Entity(obj *entityValue) *entityValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(entityValue)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func WriteEntityStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WriteEntityAddMateData(builder *flatbuffers.Builder, mateData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mateData), 0)
}
func WriteEntityAddEntity(builder *flatbuffers.Builder, entity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(entity), 0)
}
func WriteEntityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type WriteResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsWriteResponse(buf []byte, offset flatbuffers.UOffsetT) *WriteResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WriteResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWriteResponse(buf []byte, offset flatbuffers.UOffsetT) *WriteResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WriteResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WriteResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WriteResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func WriteResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func WriteResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
