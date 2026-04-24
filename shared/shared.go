// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
)

type MessageAbortedError struct {
	Data MessageAbortedErrorData `json:"data,required"`
	Name MessageAbortedErrorName `json:"name,required"`
	JSON messageAbortedErrorJSON `json:"-"`
}

// messageAbortedErrorJSON contains the JSON metadata for the struct
// [MessageAbortedError]
type messageAbortedErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAbortedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAbortedErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageAbortedError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r MessageAbortedError) ImplementsAssistantMessageError() {}

type MessageAbortedErrorData struct {
	Message string                      `json:"message,required"`
	JSON    messageAbortedErrorDataJSON `json:"-"`
}

// messageAbortedErrorDataJSON contains the JSON metadata for the struct
// [MessageAbortedErrorData]
type messageAbortedErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAbortedErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAbortedErrorDataJSON) RawJSON() string {
	return r.raw
}

type MessageAbortedErrorName string

const (
	MessageAbortedErrorNameMessageAbortedError MessageAbortedErrorName = "MessageAbortedError"
)

func (r MessageAbortedErrorName) IsKnown() bool {
	switch r {
	case MessageAbortedErrorNameMessageAbortedError:
		return true
	}
	return false
}

type ProviderAuthError struct {
	Data ProviderAuthErrorData `json:"data,required"`
	Name ProviderAuthErrorName `json:"name,required"`
	JSON providerAuthErrorJSON `json:"-"`
}

// providerAuthErrorJSON contains the JSON metadata for the struct
// [ProviderAuthError]
type providerAuthErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthErrorJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r ProviderAuthError) ImplementsAssistantMessageError() {}

type ProviderAuthErrorData struct {
	Message    string                    `json:"message,required"`
	ProviderID string                    `json:"providerID,required"`
	JSON       providerAuthErrorDataJSON `json:"-"`
}

// providerAuthErrorDataJSON contains the JSON metadata for the struct
// [ProviderAuthErrorData]
type providerAuthErrorDataJSON struct {
	Message     apijson.Field
	ProviderID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthErrorDataJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthErrorName string

const (
	ProviderAuthErrorNameProviderAuthError ProviderAuthErrorName = "ProviderAuthError"
)

func (r ProviderAuthErrorName) IsKnown() bool {
	switch r {
	case ProviderAuthErrorNameProviderAuthError:
		return true
	}
	return false
}

type UnknownError struct {
	Data UnknownErrorData `json:"data,required"`
	Name UnknownErrorName `json:"name,required"`
	JSON unknownErrorJSON `json:"-"`
}

// unknownErrorJSON contains the JSON metadata for the struct [UnknownError]
type unknownErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownErrorJSON) RawJSON() string {
	return r.raw
}

func (r UnknownError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r UnknownError) ImplementsAssistantMessageError() {}

type UnknownErrorData struct {
	Message string               `json:"message,required"`
	JSON    unknownErrorDataJSON `json:"-"`
}

// unknownErrorDataJSON contains the JSON metadata for the struct
// [UnknownErrorData]
type unknownErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownErrorDataJSON) RawJSON() string {
	return r.raw
}

type UnknownErrorName string

const (
	UnknownErrorNameUnknownError UnknownErrorName = "UnknownError"
)

func (r UnknownErrorName) IsKnown() bool {
	switch r {
	case UnknownErrorNameUnknownError:
		return true
	}
	return false
}

type BadRequestError struct {
	Data    interface{}          `json:"data,required"`
	Errors  []map[string]interface{} `json:"errors,required"`
	Success bool                 `json:"success,required"`
	JSON    badRequestErrorJSON  `json:"-"`
}

// badRequestErrorJSON contains the JSON metadata for the struct
// [BadRequestError]
type badRequestErrorJSON struct {
	Data        apijson.Field
	Errors      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BadRequestError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r badRequestErrorJSON) RawJSON() string {
	return r.raw
}

type NotFoundError struct {
	Data NotFoundErrorData `json:"data,required"`
	Name NotFoundErrorName `json:"name,required"`
	JSON notFoundErrorJSON `json:"-"`
}

// notFoundErrorJSON contains the JSON metadata for the struct [NotFoundError]
type notFoundErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notFoundErrorJSON) RawJSON() string {
	return r.raw
}

type NotFoundErrorData struct {
	Message string                `json:"message,required"`
	JSON    notFoundErrorDataJSON `json:"-"`
}

// notFoundErrorDataJSON contains the JSON metadata for the struct
// [NotFoundErrorData]
type notFoundErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotFoundErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notFoundErrorDataJSON) RawJSON() string {
	return r.raw
}

type NotFoundErrorName string

const (
	NotFoundErrorNameNotFoundError NotFoundErrorName = "NotFoundError"
)

func (r NotFoundErrorName) IsKnown() bool {
	switch r {
	case NotFoundErrorNameNotFoundError:
		return true
	}
	return false
}
