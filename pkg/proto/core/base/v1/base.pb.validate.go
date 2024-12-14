// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/core/base/v1/base.proto

package base

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on HelloRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HelloRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetStrVal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "StrVal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFloatVal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "FloatVal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDoubleVal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "DoubleVal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBoolVal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "BoolVal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBytesVal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "BytesVal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInt32Val()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "Int32Val",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUint32Val()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "Uint32Val",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInt64Val()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "Int64Val",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUint64Val()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloRequestValidationError{
				field:  "Uint64Val",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HelloRequestValidationError is the validation error returned by
// HelloRequest.Validate if the designated constraints aren't met.
type HelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloRequestValidationError) ErrorName() string { return "HelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e HelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloRequestValidationError{}

// Validate checks the field values on HelloReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HelloReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Message

	return nil
}

// HelloReplyValidationError is the validation error returned by
// HelloReply.Validate if the designated constraints aren't met.
type HelloReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloReplyValidationError) ErrorName() string { return "HelloReplyValidationError" }

// Error satisfies the builtin error interface
func (e HelloReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloReplyValidationError{}
