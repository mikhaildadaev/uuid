// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import (
	"database/sql/driver"
	"fmt"
)

// Публичные методы
func (uuid *UUID) Scan(src any) error {
	un, valid, err := scanUUID(src)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("cannot scan NULL into UUID")
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
		return UUID{}, false, fmt.Errorf("Scan: unsupported type %T", src)
	}
}
