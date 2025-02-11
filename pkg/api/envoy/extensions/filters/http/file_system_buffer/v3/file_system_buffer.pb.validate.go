// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/file_system_buffer/v3/file_system_buffer.proto

package file_system_bufferv3

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

// Validate checks the field values on BufferBehavior with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BufferBehavior) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BufferBehavior with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BufferBehaviorMultiError,
// or nil if none found.
func (m *BufferBehavior) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferBehavior) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Behavior.(type) {

	case *BufferBehavior_StreamWhenPossible_:

		if all {
			switch v := interface{}(m.GetStreamWhenPossible()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "StreamWhenPossible",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "StreamWhenPossible",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStreamWhenPossible()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferBehaviorValidationError{
					field:  "StreamWhenPossible",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *BufferBehavior_Bypass_:

		if all {
			switch v := interface{}(m.GetBypass()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "Bypass",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "Bypass",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBypass()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferBehaviorValidationError{
					field:  "Bypass",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *BufferBehavior_InjectContentLengthIfNecessary_:

		if all {
			switch v := interface{}(m.GetInjectContentLengthIfNecessary()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "InjectContentLengthIfNecessary",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "InjectContentLengthIfNecessary",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInjectContentLengthIfNecessary()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferBehaviorValidationError{
					field:  "InjectContentLengthIfNecessary",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *BufferBehavior_FullyBufferAndAlwaysInjectContentLength_:

		if all {
			switch v := interface{}(m.GetFullyBufferAndAlwaysInjectContentLength()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "FullyBufferAndAlwaysInjectContentLength",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "FullyBufferAndAlwaysInjectContentLength",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFullyBufferAndAlwaysInjectContentLength()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferBehaviorValidationError{
					field:  "FullyBufferAndAlwaysInjectContentLength",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *BufferBehavior_FullyBuffer_:

		if all {
			switch v := interface{}(m.GetFullyBuffer()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "FullyBuffer",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BufferBehaviorValidationError{
						field:  "FullyBuffer",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFullyBuffer()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferBehaviorValidationError{
					field:  "FullyBuffer",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := BufferBehaviorValidationError{
			field:  "Behavior",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return BufferBehaviorMultiError(errors)
	}

	return nil
}

// BufferBehaviorMultiError is an error wrapping multiple validation errors
// returned by BufferBehavior.ValidateAll() if the designated constraints
// aren't met.
type BufferBehaviorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferBehaviorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferBehaviorMultiError) AllErrors() []error { return m }

// BufferBehaviorValidationError is the validation error returned by
// BufferBehavior.Validate if the designated constraints aren't met.
type BufferBehaviorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferBehaviorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferBehaviorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferBehaviorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferBehaviorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferBehaviorValidationError) ErrorName() string { return "BufferBehaviorValidationError" }

// Error satisfies the builtin error interface
func (e BufferBehaviorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferBehavior.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferBehaviorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferBehaviorValidationError{}

// Validate checks the field values on StreamConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StreamConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StreamConfigMultiError, or
// nil if none found.
func (m *StreamConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBehavior()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamConfigValidationError{
					field:  "Behavior",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamConfigValidationError{
					field:  "Behavior",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBehavior()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamConfigValidationError{
				field:  "Behavior",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetMemoryBufferBytesLimit(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := StreamConfigValidationError{
				field:  "MemoryBufferBytesLimit",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetStorageBufferBytesLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamConfigValidationError{
					field:  "StorageBufferBytesLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamConfigValidationError{
					field:  "StorageBufferBytesLimit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStorageBufferBytesLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamConfigValidationError{
				field:  "StorageBufferBytesLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStorageBufferQueueHighWatermarkBytes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StreamConfigValidationError{
					field:  "StorageBufferQueueHighWatermarkBytes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamConfigValidationError{
					field:  "StorageBufferQueueHighWatermarkBytes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStorageBufferQueueHighWatermarkBytes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamConfigValidationError{
				field:  "StorageBufferQueueHighWatermarkBytes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StreamConfigMultiError(errors)
	}

	return nil
}

// StreamConfigMultiError is an error wrapping multiple validation errors
// returned by StreamConfig.ValidateAll() if the designated constraints aren't met.
type StreamConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamConfigMultiError) AllErrors() []error { return m }

// StreamConfigValidationError is the validation error returned by
// StreamConfig.Validate if the designated constraints aren't met.
type StreamConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamConfigValidationError) ErrorName() string { return "StreamConfigValidationError" }

// Error satisfies the builtin error interface
func (e StreamConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamConfigValidationError{}

// Validate checks the field values on FileSystemBufferFilterConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FileSystemBufferFilterConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileSystemBufferFilterConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FileSystemBufferFilterConfigMultiError, or nil if none found.
func (m *FileSystemBufferFilterConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *FileSystemBufferFilterConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetManagerConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "ManagerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "ManagerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetManagerConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FileSystemBufferFilterConfigValidationError{
				field:  "ManagerConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStorageBufferPath()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "StorageBufferPath",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "StorageBufferPath",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStorageBufferPath()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FileSystemBufferFilterConfigValidationError{
				field:  "StorageBufferPath",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FileSystemBufferFilterConfigValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FileSystemBufferFilterConfigValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FileSystemBufferFilterConfigValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FileSystemBufferFilterConfigMultiError(errors)
	}

	return nil
}

// FileSystemBufferFilterConfigMultiError is an error wrapping multiple
// validation errors returned by FileSystemBufferFilterConfig.ValidateAll() if
// the designated constraints aren't met.
type FileSystemBufferFilterConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileSystemBufferFilterConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileSystemBufferFilterConfigMultiError) AllErrors() []error { return m }

// FileSystemBufferFilterConfigValidationError is the validation error returned
// by FileSystemBufferFilterConfig.Validate if the designated constraints
// aren't met.
type FileSystemBufferFilterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileSystemBufferFilterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileSystemBufferFilterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileSystemBufferFilterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileSystemBufferFilterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileSystemBufferFilterConfigValidationError) ErrorName() string {
	return "FileSystemBufferFilterConfigValidationError"
}

// Error satisfies the builtin error interface
func (e FileSystemBufferFilterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileSystemBufferFilterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileSystemBufferFilterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileSystemBufferFilterConfigValidationError{}

// Validate checks the field values on BufferBehavior_StreamWhenPossible with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *BufferBehavior_StreamWhenPossible) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BufferBehavior_StreamWhenPossible
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// BufferBehavior_StreamWhenPossibleMultiError, or nil if none found.
func (m *BufferBehavior_StreamWhenPossible) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferBehavior_StreamWhenPossible) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BufferBehavior_StreamWhenPossibleMultiError(errors)
	}

	return nil
}

// BufferBehavior_StreamWhenPossibleMultiError is an error wrapping multiple
// validation errors returned by
// BufferBehavior_StreamWhenPossible.ValidateAll() if the designated
// constraints aren't met.
type BufferBehavior_StreamWhenPossibleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferBehavior_StreamWhenPossibleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferBehavior_StreamWhenPossibleMultiError) AllErrors() []error { return m }

// BufferBehavior_StreamWhenPossibleValidationError is the validation error
// returned by BufferBehavior_StreamWhenPossible.Validate if the designated
// constraints aren't met.
type BufferBehavior_StreamWhenPossibleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferBehavior_StreamWhenPossibleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferBehavior_StreamWhenPossibleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferBehavior_StreamWhenPossibleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferBehavior_StreamWhenPossibleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferBehavior_StreamWhenPossibleValidationError) ErrorName() string {
	return "BufferBehavior_StreamWhenPossibleValidationError"
}

// Error satisfies the builtin error interface
func (e BufferBehavior_StreamWhenPossibleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferBehavior_StreamWhenPossible.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferBehavior_StreamWhenPossibleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferBehavior_StreamWhenPossibleValidationError{}

// Validate checks the field values on BufferBehavior_Bypass with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BufferBehavior_Bypass) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BufferBehavior_Bypass with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BufferBehavior_BypassMultiError, or nil if none found.
func (m *BufferBehavior_Bypass) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferBehavior_Bypass) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BufferBehavior_BypassMultiError(errors)
	}

	return nil
}

// BufferBehavior_BypassMultiError is an error wrapping multiple validation
// errors returned by BufferBehavior_Bypass.ValidateAll() if the designated
// constraints aren't met.
type BufferBehavior_BypassMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferBehavior_BypassMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferBehavior_BypassMultiError) AllErrors() []error { return m }

// BufferBehavior_BypassValidationError is the validation error returned by
// BufferBehavior_Bypass.Validate if the designated constraints aren't met.
type BufferBehavior_BypassValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferBehavior_BypassValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferBehavior_BypassValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferBehavior_BypassValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferBehavior_BypassValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferBehavior_BypassValidationError) ErrorName() string {
	return "BufferBehavior_BypassValidationError"
}

// Error satisfies the builtin error interface
func (e BufferBehavior_BypassValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferBehavior_Bypass.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferBehavior_BypassValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferBehavior_BypassValidationError{}

// Validate checks the field values on
// BufferBehavior_InjectContentLengthIfNecessary with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BufferBehavior_InjectContentLengthIfNecessary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BufferBehavior_InjectContentLengthIfNecessary with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// BufferBehavior_InjectContentLengthIfNecessaryMultiError, or nil if none found.
func (m *BufferBehavior_InjectContentLengthIfNecessary) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferBehavior_InjectContentLengthIfNecessary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BufferBehavior_InjectContentLengthIfNecessaryMultiError(errors)
	}

	return nil
}

// BufferBehavior_InjectContentLengthIfNecessaryMultiError is an error wrapping
// multiple validation errors returned by
// BufferBehavior_InjectContentLengthIfNecessary.ValidateAll() if the
// designated constraints aren't met.
type BufferBehavior_InjectContentLengthIfNecessaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferBehavior_InjectContentLengthIfNecessaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferBehavior_InjectContentLengthIfNecessaryMultiError) AllErrors() []error { return m }

// BufferBehavior_InjectContentLengthIfNecessaryValidationError is the
// validation error returned by
// BufferBehavior_InjectContentLengthIfNecessary.Validate if the designated
// constraints aren't met.
type BufferBehavior_InjectContentLengthIfNecessaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferBehavior_InjectContentLengthIfNecessaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferBehavior_InjectContentLengthIfNecessaryValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e BufferBehavior_InjectContentLengthIfNecessaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferBehavior_InjectContentLengthIfNecessaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferBehavior_InjectContentLengthIfNecessaryValidationError) ErrorName() string {
	return "BufferBehavior_InjectContentLengthIfNecessaryValidationError"
}

// Error satisfies the builtin error interface
func (e BufferBehavior_InjectContentLengthIfNecessaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferBehavior_InjectContentLengthIfNecessary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferBehavior_InjectContentLengthIfNecessaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferBehavior_InjectContentLengthIfNecessaryValidationError{}

// Validate checks the field values on
// BufferBehavior_FullyBufferAndAlwaysInjectContentLength with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BufferBehavior_FullyBufferAndAlwaysInjectContentLength) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// BufferBehavior_FullyBufferAndAlwaysInjectContentLength with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BufferBehavior_FullyBufferAndAlwaysInjectContentLengthMultiError, or nil if
// none found.
func (m *BufferBehavior_FullyBufferAndAlwaysInjectContentLength) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferBehavior_FullyBufferAndAlwaysInjectContentLength) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BufferBehavior_FullyBufferAndAlwaysInjectContentLengthMultiError(errors)
	}

	return nil
}

// BufferBehavior_FullyBufferAndAlwaysInjectContentLengthMultiError is an error
// wrapping multiple validation errors returned by
// BufferBehavior_FullyBufferAndAlwaysInjectContentLength.ValidateAll() if the
// designated constraints aren't met.
type BufferBehavior_FullyBufferAndAlwaysInjectContentLengthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferBehavior_FullyBufferAndAlwaysInjectContentLengthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferBehavior_FullyBufferAndAlwaysInjectContentLengthMultiError) AllErrors() []error {
	return m
}

// BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError is the
// validation error returned by
// BufferBehavior_FullyBufferAndAlwaysInjectContentLength.Validate if the
// designated constraints aren't met.
type BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError) ErrorName() string {
	return "BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError"
}

// Error satisfies the builtin error interface
func (e BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferBehavior_FullyBufferAndAlwaysInjectContentLength.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferBehavior_FullyBufferAndAlwaysInjectContentLengthValidationError{}

// Validate checks the field values on BufferBehavior_FullyBuffer with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BufferBehavior_FullyBuffer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BufferBehavior_FullyBuffer with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BufferBehavior_FullyBufferMultiError, or nil if none found.
func (m *BufferBehavior_FullyBuffer) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferBehavior_FullyBuffer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BufferBehavior_FullyBufferMultiError(errors)
	}

	return nil
}

// BufferBehavior_FullyBufferMultiError is an error wrapping multiple
// validation errors returned by BufferBehavior_FullyBuffer.ValidateAll() if
// the designated constraints aren't met.
type BufferBehavior_FullyBufferMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferBehavior_FullyBufferMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferBehavior_FullyBufferMultiError) AllErrors() []error { return m }

// BufferBehavior_FullyBufferValidationError is the validation error returned
// by BufferBehavior_FullyBuffer.Validate if the designated constraints aren't met.
type BufferBehavior_FullyBufferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferBehavior_FullyBufferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferBehavior_FullyBufferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferBehavior_FullyBufferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferBehavior_FullyBufferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferBehavior_FullyBufferValidationError) ErrorName() string {
	return "BufferBehavior_FullyBufferValidationError"
}

// Error satisfies the builtin error interface
func (e BufferBehavior_FullyBufferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferBehavior_FullyBuffer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferBehavior_FullyBufferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferBehavior_FullyBufferValidationError{}
