// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/squash/v3/squash.proto

package envoy_extensions_filters_http_squash_v3

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _squash_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Squash with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Squash) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCluster()) < 1 {
		return SquashValidationError{
			field:  "Cluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	{
		tmp := m.GetAttachmentTemplate()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return SquashValidationError{
					field:  "AttachmentTemplate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRequestTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return SquashValidationError{
					field:  "RequestTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetAttachmentTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return SquashValidationError{
					field:  "AttachmentTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetAttachmentPollPeriod()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return SquashValidationError{
					field:  "AttachmentPollPeriod",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// SquashValidationError is the validation error returned by Squash.Validate if
// the designated constraints aren't met.
type SquashValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SquashValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SquashValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SquashValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SquashValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SquashValidationError) ErrorName() string { return "SquashValidationError" }

// Error satisfies the builtin error interface
func (e SquashValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSquash.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SquashValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SquashValidationError{}