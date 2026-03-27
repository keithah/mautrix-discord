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

type AuthMachineLogFilters struct {
	EveryHTTPRequest  bool
	EveryHTTPResponse bool

	SuccessfulLogin bool
	LoggedInUserID  bool
	Fingerprint     bool

	// The following fields are likely to log credentials and other sensitive
	// stuff when enabled. ONLY FOR USE DURING DEVELOPMENT.

	DangerouslyLeakyHTTPHeaders bool
}

var DefaultAuthMachineLogFilters = AuthMachineLogFilters{
	EveryHTTPRequest:  true,
	EveryHTTPResponse: true,

	SuccessfulLogin: true,
	LoggedInUserID:  false,
	Fingerprint:     false,

	DangerouslyLeakyHTTPHeaders: false,
}

var LeakyDevelopmentAuthMachineLogFilters = AuthMachineLogFilters{
	EveryHTTPRequest:  true,
	EveryHTTPResponse: true,

	SuccessfulLogin: true,
	LoggedInUserID:  true,
	Fingerprint:     true,

	DangerouslyLeakyHTTPHeaders: true,
}
