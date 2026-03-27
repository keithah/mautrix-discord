// mautrix-discord - A Matrix-Discord puppeting bridge.
// Copyright (C) 2026 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package discordauth

import "fmt"

// TODO(skip): Some overlap with this and discordgo. Sort that out.

type APIError struct {
	Message string  `json:"message"`
	Code    ErrCode `json:"code"`

	// Detailed errors. Returned for e.g. [InvalidFormBody].
	Errors map[string]any `json:"errors"`

	// The raw HTTP response body.
	ResponseBody []byte `json:"-"`
}

var _ error = (*APIError)(nil)

func (err APIError) Error() string {
	return fmt.Sprintf("Discord API error %d: \"%s\"", err.Code, err.Message)
}

type ErrCode int

const (
	RateLimited         ErrCode = 31001
	RateLimitedResource ErrCode = 31002

	AccountScheduledForDeletion ErrCode = 20011
	AccountDisabled             ErrCode = 20013

	Unauthorized              ErrCode = 40001
	AccountVerificationNeeded ErrCode = 40002
	CloudflareBlocked         ErrCode = 40333

	InvalidFormBody ErrCode = 50035

	MFAAlreadyEnrolled                ErrCode = 60001
	MFANotEnrolled                    ErrCode = 60002
	MFARequired                       ErrCode = 60003
	MustBeVerified                    ErrCode = 60004
	MFAInvalidSecret                  ErrCode = 60005
	MFAInvalidAuthTicket              ErrCode = 60006
	MFAInvalidCode                    ErrCode = 60008
	MFAInvalidSession                 ErrCode = 60009
	SMSAuthNotEnrolled                ErrCode = 60010
	InvalidKey                        ErrCode = 60011
	SMSAuthCannotBeEnabled            ErrCode = 60012
	MFARequiredForShopListings        ErrCode = 60015
	MFAEmailIneligible                ErrCode = 60019
	CredentialUndiscoverableOrInvalid ErrCode = 60021

	SMSAuthUnableToSendMessage              ErrCode = 70003
	SMSAuthPhoneNumberRecentlyUsedElsewhere ErrCode = 70004
	SMSAuthPhoneNumberIsVoIPOrLandline      ErrCode = 70005
	SMSAuthVerificationNeeded               ErrCode = 70007
	SMSAuthPhoneNumberAlreadyUsedElsewhere  ErrCode = 70008
	PasswordResetLinkSentToEmail            ErrCode = 70009
	SMSAuthPhoneNumberCannotBeAssociated    ErrCode = 70011
)
