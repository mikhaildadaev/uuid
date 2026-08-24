// UUID (Universal Unique IDentifier)
// Copyright (C) 2026 Mikhail Dadaev
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uuid

import (
	"database/sql/driver"
)

// Публичные методы
func (uuid *UUID) Scan(src any) error {
	un, valid, err := scanUUID(src)
	if err != nil {
		return err
	}
	if !valid {
		return ErrNullUUIDNotAllowed
	}
	*uuid = un
	return nil
}
func (uuid UUID) Value() (driver.Value, error) {
	if uuid.IsZero() {
		return nil, nil
	}
	return uuid.String(), nil
}
func (nulluuid *NullUUID) Scan(src any) error {
	u, valid, err := scanUUID(src)
	if err != nil {
		return err
	}
	nulluuid.UUID, nulluuid.Valid = u, valid
	return nil
}
func (nulluuid NullUUID) Value() (driver.Value, error) {
	if !nulluuid.Valid {
		return nil, nil
	}
	return nulluuid.UUID.String(), nil
}

// Приватные функции
func scanUUID(src any) (UUID, bool, error) {
	if src == nil {
		return UUID{}, false, nil
	}
	switch src := src.(type) {
	case string:
		if src == "" {
			return UUID{}, false, nil
		}
		u, err := Parse(src)
		return u, err == nil, err
	case []byte:
		if len(src) == 0 {
			return UUID{}, false, nil
		}
		if len(src) == 16 {
			var u UUID
			copy(u[:], src)
			return u, true, nil
		}
		return scanUUID(string(src))
	default:
		return UUID{}, false, ErrUnsupportedUUIDType
	}
}
