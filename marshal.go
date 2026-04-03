// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import (
	"encoding/hex"
	"encoding/json"
)

// Публичные методы
func (uuid *UUID) MarshalBinary() ([]byte, error) {
	if uuid.IsZero() {
		return nil, nil
	}
	return uuid[:], nil
}
func (uuid *UUID) MarshalJSON() ([]byte, error) {
	if uuid.IsZero() {
		return []byte("null"), nil
	}
	text, err := uuid.MarshalText()
	if err != nil {
		return nil, err
	}
	jsonData := make([]byte, 1, 38)
	jsonData[0] = '"'
	jsonData = append(jsonData, text...)
	jsonData = append(jsonData, '"')
	return jsonData, nil
}
func (uuid *UUID) MarshalText() ([]byte, error) {
	if uuid.IsZero() {
		return []byte(NilUUIDString), nil
	}
	buf := initCachePool.Get().(*[36]byte)
	defer initCachePool.Put(buf)
	encodeHex(buf[:], *uuid)
	return buf[:], nil
}
func (uuid *UUID) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		*uuid = NilUUIDByte
		return nil
	}
	if len(data) != 16 {
		return ErrInvalidUUIDLength
	}
	copy(uuid[:], data)
	return nil
}
func (uuid *UUID) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		*uuid = NilUUIDByte
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return ErrInvalidUUIDString
	}
	return uuid.UnmarshalText(data[1 : len(data)-1])
}
func (uuid *UUID) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*uuid = NilUUIDByte
		return nil
	}
	j := 0
	var raw [32]byte
	for _, b := range data {
		switch b {
		case '{', '}', '-':
			continue
		default:
			if j >= 32 {
				return ErrInvalidUUIDLength
			}
			raw[j] = b
			j++
		}
	}
	if j != 32 {
		return ErrInvalidUUIDLength
	}
	_, err := hex.Decode(uuid[:], raw[:])
	if err != nil {
		return ErrInvalidUUIDString
	}
	return nil
}
func (nulluuid NullUUID) MarshalBinary() ([]byte, error) {
	if nulluuid.Valid {
		return nulluuid.UUID[:], nil
	}
	return []byte(nil), nil
}
func (nulluuid NullUUID) MarshalJSON() ([]byte, error) {
	if !nulluuid.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nulluuid.UUID.String())
}
func (nulluuid NullUUID) MarshalText() ([]byte, error) {
	if !nulluuid.Valid {
		return nil, nil
	}
	return nulluuid.UUID.MarshalText()
}
func (nulluuid *NullUUID) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return ErrInvalidUUIDLength
	}
	copy(nulluuid.UUID[:], data)
	nulluuid.Valid = true
	return nil
}
func (nulluuid *NullUUID) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nulluuid.UUID, nulluuid.Valid = UUID{}, false
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	u, err := Parse(s)
	if err != nil {
		return err
	}
	nulluuid.UUID, nulluuid.Valid = u, true
	return nil
}
func (nulluuid *NullUUID) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		nulluuid.UUID, nulluuid.Valid = NilUUIDByte, false
		return nil
	}
	if err := nulluuid.UUID.UnmarshalText(data); err != nil {
		return err
	}
	nulluuid.Valid = true
	return nil
}
