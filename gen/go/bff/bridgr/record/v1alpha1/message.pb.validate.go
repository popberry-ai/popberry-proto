// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bff/bridgr/record/v1alpha1/message.proto

package recordv1alpha1

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
)

// Validate checks the field values on SendRecordRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendRecordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendRecordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendRecordRequestMultiError, or nil if none found.
func (m *SendRecordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendRecordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DeviceToken

	if all {
		switch v := interface{}(m.GetRecord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SendRecordRequestValidationError{
					field:  "Record",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SendRecordRequestValidationError{
					field:  "Record",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRecord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendRecordRequestValidationError{
				field:  "Record",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SendRecordRequestMultiError(errors)
	}

	return nil
}

// SendRecordRequestMultiError is an error wrapping multiple validation errors
// returned by SendRecordRequest.ValidateAll() if the designated constraints
// aren't met.
type SendRecordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendRecordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendRecordRequestMultiError) AllErrors() []error { return m }

// SendRecordRequestValidationError is the validation error returned by
// SendRecordRequest.Validate if the designated constraints aren't met.
type SendRecordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendRecordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendRecordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendRecordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendRecordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendRecordRequestValidationError) ErrorName() string {
	return "SendRecordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendRecordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendRecordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendRecordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendRecordRequestValidationError{}

// Validate checks the field values on SendRecordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendRecordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendRecordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendRecordResponseMultiError, or nil if none found.
func (m *SendRecordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SendRecordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	// no validation rules for Success

	if len(errors) > 0 {
		return SendRecordResponseMultiError(errors)
	}

	return nil
}

// SendRecordResponseMultiError is an error wrapping multiple validation errors
// returned by SendRecordResponse.ValidateAll() if the designated constraints
// aren't met.
type SendRecordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendRecordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendRecordResponseMultiError) AllErrors() []error { return m }

// SendRecordResponseValidationError is the validation error returned by
// SendRecordResponse.Validate if the designated constraints aren't met.
type SendRecordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendRecordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendRecordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendRecordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendRecordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendRecordResponseValidationError) ErrorName() string {
	return "SendRecordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SendRecordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendRecordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendRecordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendRecordResponseValidationError{}

// Validate checks the field values on ParsedRecord with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ParsedRecord) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ParsedRecord with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ParsedRecordMultiError, or
// nil if none found.
func (m *ParsedRecord) ValidateAll() error {
	return m.validate(true)
}

func (m *ParsedRecord) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Recipient

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Priority

	// no validation rules for Metadata

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ParsedRecordValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ParsedRecordValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ParsedRecordValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ParsedRecordMultiError(errors)
	}

	return nil
}

// ParsedRecordMultiError is an error wrapping multiple validation errors
// returned by ParsedRecord.ValidateAll() if the designated constraints aren't met.
type ParsedRecordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParsedRecordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParsedRecordMultiError) AllErrors() []error { return m }

// ParsedRecordValidationError is the validation error returned by
// ParsedRecord.Validate if the designated constraints aren't met.
type ParsedRecordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParsedRecordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParsedRecordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParsedRecordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParsedRecordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParsedRecordValidationError) ErrorName() string { return "ParsedRecordValidationError" }

// Error satisfies the builtin error interface
func (e ParsedRecordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParsedRecord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParsedRecordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParsedRecordValidationError{}

// Validate checks the field values on RawRecord with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RawRecord) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RawRecord with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RawRecordMultiError, or nil
// if none found.
func (m *RawRecord) ValidateAll() error {
	return m.validate(true)
}

func (m *RawRecord) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Source

	// no validation rules for Sender

	// no validation rules for Header

	// no validation rules for Body

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RawRecordValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RawRecordValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RawRecordValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.EncodingType != nil {
		// no validation rules for EncodingType
	}

	if len(errors) > 0 {
		return RawRecordMultiError(errors)
	}

	return nil
}

// RawRecordMultiError is an error wrapping multiple validation errors returned
// by RawRecord.ValidateAll() if the designated constraints aren't met.
type RawRecordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RawRecordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RawRecordMultiError) AllErrors() []error { return m }

// RawRecordValidationError is the validation error returned by
// RawRecord.Validate if the designated constraints aren't met.
type RawRecordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RawRecordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RawRecordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RawRecordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RawRecordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RawRecordValidationError) ErrorName() string { return "RawRecordValidationError" }

// Error satisfies the builtin error interface
func (e RawRecordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRawRecord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RawRecordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RawRecordValidationError{}

// Validate checks the field values on Record with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Record) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Record with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RecordMultiError, or nil if none found.
func (m *Record) ValidateAll() error {
	return m.validate(true)
}

func (m *Record) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RecordValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RecordValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RecordValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.MessageData.(type) {
	case *Record_Parsed:
		if v == nil {
			err := RecordValidationError{
				field:  "MessageData",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetParsed()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RecordValidationError{
						field:  "Parsed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RecordValidationError{
						field:  "Parsed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetParsed()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RecordValidationError{
					field:  "Parsed",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Record_Raw:
		if v == nil {
			err := RecordValidationError{
				field:  "MessageData",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRaw()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RecordValidationError{
						field:  "Raw",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RecordValidationError{
						field:  "Raw",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRaw()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RecordValidationError{
					field:  "Raw",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RecordMultiError(errors)
	}

	return nil
}

// RecordMultiError is an error wrapping multiple validation errors returned by
// Record.ValidateAll() if the designated constraints aren't met.
type RecordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecordMultiError) AllErrors() []error { return m }

// RecordValidationError is the validation error returned by Record.Validate if
// the designated constraints aren't met.
type RecordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecordValidationError) ErrorName() string { return "RecordValidationError" }

// Error satisfies the builtin error interface
func (e RecordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecordValidationError{}
