// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: accounts/region.proto

package acctspb

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

// Validate checks the field values on RegionReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegionReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegionReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegionReplyMultiError, or
// nil if none found.
func (m *RegionReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RegionReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRegions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegionReplyValidationError{
						field:  fmt.Sprintf("Regions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegionReplyValidationError{
						field:  fmt.Sprintf("Regions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegionReplyValidationError{
					field:  fmt.Sprintf("Regions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RegionReplyMultiError(errors)
	}

	return nil
}

// RegionReplyMultiError is an error wrapping multiple validation errors
// returned by RegionReply.ValidateAll() if the designated constraints aren't met.
type RegionReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegionReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegionReplyMultiError) AllErrors() []error { return m }

// RegionReplyValidationError is the validation error returned by
// RegionReply.Validate if the designated constraints aren't met.
type RegionReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegionReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegionReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegionReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegionReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegionReplyValidationError) ErrorName() string { return "RegionReplyValidationError" }

// Error satisfies the builtin error interface
func (e RegionReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegionReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegionReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegionReplyValidationError{}

// Validate checks the field values on RegionReply_Region with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegionReply_Region) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegionReply_Region with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegionReply_RegionMultiError, or nil if none found.
func (m *RegionReply_Region) ValidateAll() error {
	return m.validate(true)
}

func (m *RegionReply_Region) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Area

	// no validation rules for Img

	// no validation rules for Name

	if len(errors) > 0 {
		return RegionReply_RegionMultiError(errors)
	}

	return nil
}

// RegionReply_RegionMultiError is an error wrapping multiple validation errors
// returned by RegionReply_Region.ValidateAll() if the designated constraints
// aren't met.
type RegionReply_RegionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegionReply_RegionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegionReply_RegionMultiError) AllErrors() []error { return m }

// RegionReply_RegionValidationError is the validation error returned by
// RegionReply_Region.Validate if the designated constraints aren't met.
type RegionReply_RegionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegionReply_RegionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegionReply_RegionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegionReply_RegionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegionReply_RegionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegionReply_RegionValidationError) ErrorName() string {
	return "RegionReply_RegionValidationError"
}

// Error satisfies the builtin error interface
func (e RegionReply_RegionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegionReply_Region.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegionReply_RegionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegionReply_RegionValidationError{}