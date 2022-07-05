// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: banyandb/model/v1/common.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	structpb "google.golang.org/protobuf/types/known/structpb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = structpb.NullValue(0)
)

// Validate checks the field values on ID with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *ID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ID with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IDMultiError, or nil if none found.
func (m *ID) ValidateAll() error {
	return m.validate(true)
}

func (m *ID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return IDMultiError(errors)
	}

	return nil
}

// IDMultiError is an error wrapping multiple validation errors returned by
// ID.ValidateAll() if the designated constraints aren't met.
type IDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IDMultiError) AllErrors() []error { return m }

// IDValidationError is the validation error returned by ID.Validate if the
// designated constraints aren't met.
type IDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IDValidationError) ErrorName() string { return "IDValidationError" }

// Error satisfies the builtin error interface
func (e IDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IDValidationError{}

// Validate checks the field values on Str with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Str) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Str with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StrMultiError, or nil if none found.
func (m *Str) ValidateAll() error {
	return m.validate(true)
}

func (m *Str) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return StrMultiError(errors)
	}

	return nil
}

// StrMultiError is an error wrapping multiple validation errors returned by
// Str.ValidateAll() if the designated constraints aren't met.
type StrMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrMultiError) AllErrors() []error { return m }

// StrValidationError is the validation error returned by Str.Validate if the
// designated constraints aren't met.
type StrValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrValidationError) ErrorName() string { return "StrValidationError" }

// Error satisfies the builtin error interface
func (e StrValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStr.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrValidationError{}

// Validate checks the field values on Int with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Int) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Int with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IntMultiError, or nil if none found.
func (m *Int) ValidateAll() error {
	return m.validate(true)
}

func (m *Int) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return IntMultiError(errors)
	}

	return nil
}

// IntMultiError is an error wrapping multiple validation errors returned by
// Int.ValidateAll() if the designated constraints aren't met.
type IntMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntMultiError) AllErrors() []error { return m }

// IntValidationError is the validation error returned by Int.Validate if the
// designated constraints aren't met.
type IntValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntValidationError) ErrorName() string { return "IntValidationError" }

// Error satisfies the builtin error interface
func (e IntValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntValidationError{}

// Validate checks the field values on StrArray with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StrArray) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrArray with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StrArrayMultiError, or nil
// if none found.
func (m *StrArray) ValidateAll() error {
	return m.validate(true)
}

func (m *StrArray) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StrArrayMultiError(errors)
	}

	return nil
}

// StrArrayMultiError is an error wrapping multiple validation errors returned
// by StrArray.ValidateAll() if the designated constraints aren't met.
type StrArrayMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrArrayMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrArrayMultiError) AllErrors() []error { return m }

// StrArrayValidationError is the validation error returned by
// StrArray.Validate if the designated constraints aren't met.
type StrArrayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrArrayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrArrayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrArrayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrArrayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrArrayValidationError) ErrorName() string { return "StrArrayValidationError" }

// Error satisfies the builtin error interface
func (e StrArrayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrArray.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrArrayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrArrayValidationError{}

// Validate checks the field values on IntArray with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IntArray) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntArray with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IntArrayMultiError, or nil
// if none found.
func (m *IntArray) ValidateAll() error {
	return m.validate(true)
}

func (m *IntArray) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IntArrayMultiError(errors)
	}

	return nil
}

// IntArrayMultiError is an error wrapping multiple validation errors returned
// by IntArray.ValidateAll() if the designated constraints aren't met.
type IntArrayMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntArrayMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntArrayMultiError) AllErrors() []error { return m }

// IntArrayValidationError is the validation error returned by
// IntArray.Validate if the designated constraints aren't met.
type IntArrayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntArrayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntArrayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntArrayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntArrayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntArrayValidationError) ErrorName() string { return "IntArrayValidationError" }

// Error satisfies the builtin error interface
func (e IntArrayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntArray.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntArrayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntArrayValidationError{}

// Validate checks the field values on TagValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TagValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TagValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TagValueMultiError, or nil
// if none found.
func (m *TagValue) ValidateAll() error {
	return m.validate(true)
}

func (m *TagValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Value.(type) {

	case *TagValue_Null:
		// no validation rules for Null

	case *TagValue_Str:

		if all {
			switch v := interface{}(m.GetStr()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "Str",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "Str",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStr()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagValueValidationError{
					field:  "Str",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TagValue_StrArray:

		if all {
			switch v := interface{}(m.GetStrArray()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "StrArray",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "StrArray",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStrArray()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagValueValidationError{
					field:  "StrArray",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TagValue_Int:

		if all {
			switch v := interface{}(m.GetInt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "Int",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "Int",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagValueValidationError{
					field:  "Int",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TagValue_IntArray:

		if all {
			switch v := interface{}(m.GetIntArray()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "IntArray",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "IntArray",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetIntArray()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagValueValidationError{
					field:  "IntArray",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TagValue_BinaryData:
		// no validation rules for BinaryData

	case *TagValue_Id:

		if all {
			switch v := interface{}(m.GetId()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "Id",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagValueValidationError{
						field:  "Id",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagValueValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TagValueMultiError(errors)
	}

	return nil
}

// TagValueMultiError is an error wrapping multiple validation errors returned
// by TagValue.ValidateAll() if the designated constraints aren't met.
type TagValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagValueMultiError) AllErrors() []error { return m }

// TagValueValidationError is the validation error returned by
// TagValue.Validate if the designated constraints aren't met.
type TagValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagValueValidationError) ErrorName() string { return "TagValueValidationError" }

// Error satisfies the builtin error interface
func (e TagValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagValueValidationError{}

// Validate checks the field values on TagFamilyForWrite with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TagFamilyForWrite) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TagFamilyForWrite with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TagFamilyForWriteMultiError, or nil if none found.
func (m *TagFamilyForWrite) ValidateAll() error {
	return m.validate(true)
}

func (m *TagFamilyForWrite) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagFamilyForWriteValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagFamilyForWriteValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagFamilyForWriteValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TagFamilyForWriteMultiError(errors)
	}

	return nil
}

// TagFamilyForWriteMultiError is an error wrapping multiple validation errors
// returned by TagFamilyForWrite.ValidateAll() if the designated constraints
// aren't met.
type TagFamilyForWriteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagFamilyForWriteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagFamilyForWriteMultiError) AllErrors() []error { return m }

// TagFamilyForWriteValidationError is the validation error returned by
// TagFamilyForWrite.Validate if the designated constraints aren't met.
type TagFamilyForWriteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagFamilyForWriteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagFamilyForWriteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagFamilyForWriteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagFamilyForWriteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagFamilyForWriteValidationError) ErrorName() string {
	return "TagFamilyForWriteValidationError"
}

// Error satisfies the builtin error interface
func (e TagFamilyForWriteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagFamilyForWrite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagFamilyForWriteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagFamilyForWriteValidationError{}

// Validate checks the field values on FieldValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FieldValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FieldValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FieldValueMultiError, or
// nil if none found.
func (m *FieldValue) ValidateAll() error {
	return m.validate(true)
}

func (m *FieldValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Value.(type) {

	case *FieldValue_Null:
		// no validation rules for Null

	case *FieldValue_Str:

		if all {
			switch v := interface{}(m.GetStr()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FieldValueValidationError{
						field:  "Str",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FieldValueValidationError{
						field:  "Str",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStr()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldValueValidationError{
					field:  "Str",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FieldValue_Int:

		if all {
			switch v := interface{}(m.GetInt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FieldValueValidationError{
						field:  "Int",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FieldValueValidationError{
						field:  "Int",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldValueValidationError{
					field:  "Int",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FieldValue_BinaryData:
		// no validation rules for BinaryData

	}

	if len(errors) > 0 {
		return FieldValueMultiError(errors)
	}

	return nil
}

// FieldValueMultiError is an error wrapping multiple validation errors
// returned by FieldValue.ValidateAll() if the designated constraints aren't met.
type FieldValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldValueMultiError) AllErrors() []error { return m }

// FieldValueValidationError is the validation error returned by
// FieldValue.Validate if the designated constraints aren't met.
type FieldValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldValueValidationError) ErrorName() string { return "FieldValueValidationError" }

// Error satisfies the builtin error interface
func (e FieldValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldValueValidationError{}
