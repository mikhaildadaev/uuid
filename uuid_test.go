// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"strings"
	"testing"
	"time"
)

// Tests
func Test_Init(t *testing.T) {
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		t.Error("Invalid initial global state")
	}
	if initMAC.Load().([6]byte) == [6]byte{} {
		t.Error("MAC address was not initialized")
	}
	if sq1 := v1.lastSequence.Load(); sq1 > maxV1Sequence {
		t.Errorf("Invalid initial v1 sequence: %d", sq1)
	}
	if sq6 := v6.lastSequence.Load(); sq6 > maxV6Sequence {
		t.Errorf("Invalid initial v6 sequence: %d", sq6)
	}
	if sq7 := v7.lastSequence.Load(); sq7 > maxV7Sequence {
		t.Errorf("Invalid initial v1 sequence: %d", sq7)
	}
	if sq8 := v8.lastSequence.Load(); sq8 > maxV8Sequence {
		t.Errorf("Invalid initial v6 sequence: %d", sq8)
	}
}
func Test_Info(t *testing.T) {
	testCases := []struct {
		name string
		uuid UUID
	}{
		{"V1 UUID", func() UUID { u := V1(); return u }()},
		{"V2 UUID", func() UUID { u := V2(testPOSType); return u }()},
		{"V3 UUID", func() UUID { u := V3(NameSpaceDNS, testNameString); return u }()},
		{"V4 UUID", func() UUID { u := V4(); return u }()},
		{"V5 UUID", func() UUID { u := V5(NameSpaceDNS, testNameString); return u }()},
		{"V6 UUID", func() UUID { u := V6(); return u }()},
		{"V7 UUID", func() UUID { u := V7(); return u }()},
		{"V8 UUID", func() UUID { u := V8(testNodeID); return u }()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info := tc.uuid.Info()
			if info == "" {
				t.Error("Info() returned empty string")
			}
			if !strings.Contains(info, "UUID:") {
				t.Error("Info() does not contain UUID")
			}
			if !strings.Contains(info, "VAR.:") {
				t.Error("Info() does not contain VAR.")
			}
			if !strings.Contains(info, "VER.:") {
				t.Error("Info() does not contain VER.")
			}
			if !strings.Contains(info, "FORM:") {
				t.Error("Info() does not contain FORM")
			}
			if !strings.Contains(info, "INFO:") {
				t.Error("Info() does not contain INFO")
			}
		})
	}
}
func Test_Nil(t *testing.T) {
	uinil, _ := Parse(NilUUIDString)
	uinonil, _ := Parse(testUUIDVUString)
	if !uinil.IsZero() {
		t.Error("returned false nil UUID")
	}
	if uinonil.IsZero() {
		t.Error("returned true nil UUID")
	}
}
func Test_Parse(t *testing.T) {
	validCases := []struct {
		name string
		line string
		want UUID
	}{
		{
			name: "Nil UUID",
			line: NilUUIDString,
			want: NilUUIDByte,
		},
		{
			name: "V1 UUID",
			line: testUUIDV1String,
			want: testUUIDV1Byte,
		},
		{
			name: "V2 UUID",
			line: testUUIDV2String,
			want: testUUIDV2Byte,
		},
		{
			name: "V3 UUID",
			line: testUUIDV3String,
			want: testUUIDV3Byte,
		},
		{
			name: "V4 UUID",
			line: testUUIDV4String,
			want: testUUIDV4Byte,
		},
		{
			name: "V5 UUID",
			line: testUUIDV5String,
			want: testUUIDV5Byte,
		},
		{
			name: "V6 UUID",
			line: testUUIDV6String,
			want: testUUIDV6Byte,
		},
		{
			name: "V7 UUID",
			line: testUUIDV7String,
			want: testUUIDV7Byte,
		},
		{
			name: "V8 UUID",
			line: testUUIDV8String,
			want: testUUIDV8Byte,
		},
	}
	for _, tc := range validCases {
		t.Run(tc.name, func(t *testing.T) {
			got := must(Parse(tc.line))
			if got != tc.want {
				t.Errorf("Parse() = %v, want %v", got, tc.want)
			}
		})
	}
	invalidCases := []struct {
		name string
		line string
		want UUID
	}{
		{
			name: "Empty string UUID",
			line: testUUIDErrStringEmpty,
		},
		{
			name: "Invalid string UUID",
			line: testUUIDErrStringInvalid,
		},
		{
			name: "Incorrect string character UUID",
			line: testUUIDErrStringCharacter,
		},
		{
			name: "Long string UUID",
			line: testUUIDErrStringLengthLong,
		},
		{
			name: "Short string UUID",
			line: testUUIDErrStringLengthShort,
		},
	}
	for _, tc := range invalidCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Parse(tc.line)
			if err == nil {
				t.Errorf("Expected error for invalid UUID %s", tc.line)
			}
		})
	}
}
func Test_Scan(t *testing.T) {
	tests := []struct {
		name    string
		src     any
		want    UUID
		wantErr bool
	}{
		{
			name:    "Empty byte format UUID",
			src:     testUUIDErrByteEmpty,
			want:    NilUUIDByte,
			wantErr: true,
		},
		{
			name:    "Empty string format UUID",
			src:     testUUIDErrStringEmpty,
			want:    NilUUIDByte,
			wantErr: true,
		},
		{
			name:    "Long byte format UUID",
			src:     testUUIDErrByteLengthLong,
			want:    NilUUIDByte,
			wantErr: true,
		},
		{
			name:    "Short byte format UUID",
			src:     testUUIDErrByteLengthShort,
			want:    NilUUIDByte,
			wantErr: true,
		},
		{
			name:    "Invalid type",
			src:     testUUIDErrTypeInt,
			want:    NilUUIDByte,
			wantErr: true,
		},
		{
			name:    "Nil input",
			src:     nil,
			want:    NilUUIDByte,
			wantErr: true,
		},
		{
			name:    "Valid byte format UUID",
			src:     must(Parse(testUUIDVUString)).Bytes(),
			want:    testUUIDVUByte,
			wantErr: false,
		},
		{
			name:    "Valid string format UUID",
			src:     testUUIDVUString,
			want:    testUUIDVUByte,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u UUID
			err := u.Scan(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if u != tt.want {
				t.Errorf("Scan() = %v, want %v", u, tt.want)
			}
		})
	}
}
func Test_String(t *testing.T) {
	validCases := []struct {
		name string
		uuid UUID
		want string
	}{
		{
			name: "Nil UUID",
			uuid: NilUUIDByte,
			want: NilUUIDString,
		},
		{
			name: "V1 UUID",
			uuid: testUUIDV1Byte,
			want: testUUIDV1String,
		},
		{
			name: "V2 UUID",
			uuid: testUUIDV2Byte,
			want: testUUIDV2String,
		},
		{
			name: "V3 UUID",
			uuid: testUUIDV3Byte,
			want: testUUIDV3String,
		},
		{
			name: "V4 UUID",
			uuid: testUUIDV4Byte,
			want: testUUIDV4String,
		},
		{
			name: "V5 UUID",
			uuid: testUUIDV5Byte,
			want: testUUIDV5String,
		},
		{
			name: "V6 UUID",
			uuid: testUUIDV6Byte,
			want: testUUIDV6String,
		},
		{
			name: "V7 UUID",
			uuid: testUUIDV7Byte,
			want: testUUIDV7String,
		},
		{
			name: "V8 UUID",
			uuid: testUUIDV8Byte,
			want: testUUIDV8String,
		},
	}
	for _, tc := range validCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.uuid.String()
			if got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}
func Test_TextFormatVariations(t *testing.T) {
	formats := []struct {
		name   string
		input  string
		expect UUID
	}{
		{"Long_format_UUID", testUUIDErrStringLong, testUUIDVUByte},
		{"Short_format_UUID", testUUIDErrStringShort, testUUIDVUByte},
		{"Standard_format_UUID", testUUIDVUString, testUUIDVUByte},
	}
	for _, tt := range formats {
		t.Run(tt.name, func(t *testing.T) {
			u := must(Parse(tt.input))
			if u != tt.expect {
				t.Errorf("UnmarshalText() = %v, want %v", u, tt.expect)
			}
		})
	}
}
func Test_Timestamp(t *testing.T) {
	t.Run("V1 UUID", func(t *testing.T) {
		var prevUUID UUID
		for i := 0; i < 10; i++ {
			uuid := V1()
			if i > 0 {
				currentTime := uuid.Timestamp()
				prevTime := prevUUID.Timestamp()
				currentSequence := binary.BigEndian.Uint16(uuid[8:10])
				prevSequence := binary.BigEndian.Uint16(prevUUID[8:10])
				if currentTime < prevTime {
					t.Errorf("Timestamps decreased: %d < %d", currentTime, prevTime)
				} else if currentTime == prevTime {
					if currentSequence <= prevSequence {
						t.Errorf("Same timestamp but sequence not increased: %d <= %d", currentSequence, prevSequence)
					}
				}
			}
			prevUUID = uuid
		}
	})
	t.Run("V6 UUID", func(t *testing.T) {
		var prevUUID UUID
		for i := 0; i < 10; i++ {
			uuid := V6()
			if i > 0 {
				currentTime := uuid.Timestamp()
				prevTime := prevUUID.Timestamp()
				currentSequence := binary.BigEndian.Uint16(uuid[8:10])
				prevSequence := binary.BigEndian.Uint16(prevUUID[8:10])
				if currentTime < prevTime {
					t.Errorf("Timestamps decreased: %d < %d", currentTime, prevTime)
				} else if currentTime == prevTime {
					if currentSequence <= prevSequence {
						t.Errorf("Same timestamp but sequence not increased: %d <= %d", currentSequence, prevSequence)
					}
				}
			}
			prevUUID = uuid
		}
	})
	t.Run("V7 UUID", func(t *testing.T) {
		var prevUUID UUID
		for i := 0; i < 10; i++ {
			uuid := V7()
			if i > 0 {
				currentTime := uuid.Timestamp()
				prevTime := prevUUID.Timestamp()
				if currentTime < prevTime {
					t.Errorf("Timestamps decreased: %d < %d", currentTime, prevTime)
				}
			}
			prevUUID = uuid
		}
	})
	t.Run("V8 UUID", func(t *testing.T) {
		var prevUUID UUID
		for i := 0; i < 10; i++ {
			uuid := V8(testNodeID)
			if i > 0 {
				currentTime := uuid.Timestamp()
				prevTime := prevUUID.Timestamp()
				if currentTime < prevTime {
					t.Errorf("Timestamps decreased: %d < %d", currentTime, prevTime)
				}
			}
			prevUUID = uuid
		}
	})
}
func Test_Validate(t *testing.T) {
	t.Run("Valid UUID", func(t *testing.T) {
		tests := []struct {
			name string
			uuid UUID
		}{
			{"V1 UUID", V1()},
			{"V2 UUID", V2(testPOSType)},
			{"V3 UUID", V3(NameSpaceDNS, testNameString)},
			{"V4 UUID", V4()},
			{"V5 UUID", V5(NameSpaceDNS, testNameString)},
			{"V6 UUID", V6()},
			{"V7 UUID", V7()},
			{"V8 UUID", V8(testNodeID)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := tt.uuid.Validate(); err != nil {
					t.Errorf("Validate() error = %v", err)
				}
				switch tt.uuid.Version() {
				case 1, 6:
					if tt.uuid[8]&0xC0 != 0x80 {
						t.Error("Invalid variant bits")
					}
				case 7, 8:
					if tt.uuid.Timestamp() == 0 {
						t.Error("Zero timestamp")
					}
				}
			})
		}
	})
	t.Run("Invalid UUID", func(t *testing.T) {
		validV1 := V1()
		tests := []struct {
			name    string
			uuid    UUID
			wantErr error
		}{
			{
				name:    "Invalid variant",
				uuid:    func() UUID { u := validV1; u[8] = 0x00; return u }(),
				wantErr: ErrInvalidUUIDVariant,
			},
			{
				name:    "Invalid version",
				uuid:    func() UUID { u := validV1; u[6] = 0x00; return u }(),
				wantErr: ErrInvalidUUIDVersion,
			},
			{
				name:    "Nil UUID",
				uuid:    NilUUIDByte,
				wantErr: ErrNullUUID,
			},
			{
				name:    "Nil MAC UUIDV1",
				uuid:    func() UUID { u := validV1; copy(u[10:16], make([]byte, 6)); return u }(),
				wantErr: ErrInvalidUUIDMAC,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.uuid.Validate()
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Validate() error = %v, want %v", err, tt.wantErr)
				}
			})
		}
	})
}
func Test_Value(t *testing.T) {
	tests := []struct {
		name    string
		u       UUID
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "Null UUID",
			u:       NilUUIDByte,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Valid UUID",
			u:       testUUIDVUByte,
			want:    testUUIDVUString,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_MarshalUnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		wantErr bool
	}{
		{"Nil UUID", NilUUIDByte, false},
		{"Valid UUID", testUUIDVUByte, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal
			data, err := tt.uuid.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.uuid.IsZero() && data != nil {
				t.Error("nil UUID should marshal to nil slice")
			}
			if !tt.uuid.IsZero() && !bytes.Equal(data, tt.uuid[:]) {
				t.Error("marshaled data doesn't match UUID bytes")
			}
			// Unmarshal
			var u UUID
			err = u.UnmarshalBinary(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if u != tt.uuid {
				t.Errorf("Unmarshaled UUID = %v, want %v", u, tt.uuid)
			}
		})
	}
	t.Run("Invalid UUID", func(t *testing.T) {
		var u UUID
		err := u.UnmarshalBinary([]byte{1, 2, 3})
		if err == nil {
			t.Error("expected error for invalid length")
		}
	})
}
func Test_MarshalUnmarshalJson(t *testing.T) {
	tests := []struct {
		name    string
		u       UUID
		want    string
		wantErr bool
	}{
		{
			name:    "Null UUID",
			u:       NilUUIDByte,
			want:    "null",
			wantErr: false,
		},
		{
			name:    "Valid UUID",
			u:       testUUIDVUByte,
			want:    testUUIDVUJson,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// MarshalJson
			json, err := tt.u.MarshalJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(json) != tt.want {
				t.Errorf("MarshalJson() = %s, want %s", json, tt.want)
			}
			if tt.u.IsZero() {
				return
			}
			// UnmarshalJson
			var u UUID
			err = u.UnmarshalJson(json)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if u != tt.u {
				t.Errorf("UnmarshalJson() = %v, want %v", u, tt.u)
			}
		})
	}
	t.Run("Invalid UUID", func(t *testing.T) {
		var u UUID
		err := u.UnmarshalJson([]byte(`"invalid"`))
		if err == nil {
			t.Error("UnmarshalJson() expected error, got nil")
		}
	})
}
func Test_MarshalUnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		u       UUID
		want    string
		wantErr bool
	}{
		{
			name:    "Null UUID",
			u:       NilUUIDByte,
			want:    testUUIDVUNull,
			wantErr: false,
		},
		{
			name:    "Valid UUID",
			u:       testUUIDVUByte,
			want:    testUUIDVUText,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// MarshalText
			textData, err := tt.u.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(textData) != tt.want {
				t.Errorf("MarshalText() = %s, want %s", textData, tt.want)
			}
			// UnmarshalText
			var u UUID
			err = u.UnmarshalText(textData)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if u != tt.u {
				t.Errorf("UnmarshalText() = %v, want %v", u, tt.u)
			}
		})
	}
	invalidTests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"Invalid chars UUID", []byte(testUUIDErrStringCharacter), true},
		{"Long format UUID", []byte(testUUIDErrStringLong), false},
		{"Short format UUID", []byte(testUUIDErrStringLengthShort), true},
	}
	for _, tt := range invalidTests {
		t.Run(tt.name, func(t *testing.T) {
			var u UUID
			err := u.UnmarshalText(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_UUIDV1_Generate(t *testing.T) {
	// Ленивая инициализация глобального состояния
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		t.Fatal(initError)
	}
	// Фиксация текущего состояние
	savedClock := initClock
	savedTime := v1.lastTime.Load()
	savedSequence := v1.lastSequence.Load()
	t.Cleanup(func() {
		initClock = savedClock
		v1.lastTime.Store(savedTime)
		v1.lastSequence.Store(savedSequence)
	})
	// Мокирование
	initClock = realClock{}
	v1.lastTime.Store(0)
	v1.lastSequence.Store(0)
	// Генерация идентификаторов
	ui1 := V1()
	ui2 := V1()
	// Проверка вариантов
	vt1 := ui1.Variant()
	vt2 := ui2.Variant()
	if vt1 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv1-1: got %x, expected %x", vt1, bitVRFC4122)
	}
	if vt2 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv1-2: got %x, expected %x", vt2, bitVRFC4122)
	}
	// Проверка версий
	vn1 := ui1.Version()
	vn2 := ui2.Version()
	if vn1 != bitV1>>4 {
		t.Errorf("Invalid version for UUIDv1-1: got %x, expected %x", vn1, bitV1)
	}
	if vn2 != bitV1>>4 {
		t.Errorf("Invalid version for UUIDv1-2: got %x, expected %x", vn2, bitV1)
	}
	// Проверка временных меток
	ts1 := ui1.Timestamp()
	ts2 := ui2.Timestamp()
	if ts1 <= 0 || ts2 <= 0 {
		t.Error("Invalid timestamp UUIDv1")
	}
	if ts1 > ts2 {
		t.Error("Failed monotonically increasing timestamp UUIDv1")
	}
	// Проверка детерминированности
	if ui1 == ui2 {
		t.Error("Failed to generate different UUIDs")
	}
}
func Test_UUIDV1_Sequence(t *testing.T) {
	t.Run("Sequence_Increment", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v1.lastTime.Load()
		realSequence := v1.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v1.lastTime.Store(realTime)
			v1.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v1.lastTime.Store(uint64((initClock.now().UnixNano() / 100) + offsetTime))
		v1.lastSequence.Store(0)
		// Генерация идентификаторов
		ui1 := V1()
		ui2 := V1()
		// Получение временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if ts1 == ts2 && sq1 >= sq2 {
			t.Errorf("Failed to increment sequence: ts1=%d ts2=%d sq1=%d sq2=%d", ts1, ts2, sq1, sq2)
		}
	})
	t.Run("Sequence_Overflow", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v1.lastTime.Load()
		realSequence := v1.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v1.lastTime.Store(realTime)
			v1.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v1.lastTime.Store(uint64((initClock.now().UnixNano() / 100) + offsetTime))
		v1.lastSequence.Store(maxV1Sequence - 1)
		// Генерация идентификаторов
		ui1 := V1()
		mockClock.add(time.Nanosecond * 100)
		ui2 := V1()
		// Получение временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if sq2 != 0 {
			t.Errorf("Failed to overflow sequence: sq1=%d sq2=%d", sq1, sq2)
		}
	})
}
func Test_UUIDV2_Generate(t *testing.T) {
	for i := int(0); i < 2; i++ {
		// Генерация идентификаторов
		ui := V2(i)
		// Проверка вариантов
		vt := ui.Variant()
		if vt != variantRFC4122 {
			t.Errorf("Invalid variant for UUIDv2-%d: got %x, expected %x", i+1, vt, bitVRFC4122)
		}
		// Проверка версий
		vn := ui.Version()
		if vn != bitV2>>4 {
			t.Errorf("Invalid version for UUIDv2-%d: got %x, expected %x", i+1, vn, bitV2)
		}
		// Проверка идентификаторов систем
		if ui[9] != uint8(i&0x03) {
			t.Errorf("Invalid posixtype in UUIDv2-%d: got %d, expected %d", i+1, ui[9], i)
		}
	}
}
func Test_UUIDV3_Generate(t *testing.T) {
	testCases := []struct {
		namespace UUID
		name      string
	}{
		{NameSpaceDNS, testNameString},
		{NameSpaceURL, testNameString},
		{NameSpaceOID, testNameString},
		{NameSpaceX500, testNameString},
	}
	for i, tc := range testCases {
		// Генерация идентификаторов
		ui := V3(tc.namespace, tc.name)
		// Проверка вариантов
		vt := ui.Variant()
		if vt != variantRFC4122 {
			t.Errorf("Invalid variant for UUIDv3-%d: got %x, expected %x", i+1, vt, bitVRFC4122)
		}
		// Проверка версий
		vn := ui.Version()
		if vn != bitV3>>4 {
			t.Errorf("Invalid version for UUIDv3-%d: got %x, expected %x", i+1, vn, bitV3)
		}
	}
}
func Test_UUIDV3_Hash(t *testing.T) {
	// Генерация идентификаторов
	ui1a := V3(NameSpaceDNS, testNameString+"/TestA")
	ui1b := V3(NameSpaceDNS, testNameString+"/TestB")
	ui2a := V3(NameSpaceURL, testNameString+"/TestA")
	ui2b := V3(NameSpaceURL, testNameString+"/TestB")
	// Проверка детерминированности при разных NameSpace
	if (ui1a == ui2a) || (ui1b == ui2b) {
		t.Error("Invalid UUIDv3 with the same value for different Namespace")
	}
	// Проверка детерминированности при разных TestNameString
	if (ui1a == ui1b) || (ui2a == ui2b) {
		t.Error("Invalid UUIDv3 with the same value for different TestNameString")
	}
}
func Test_UUIDV4_Generate(t *testing.T) {
	// Генерация идентификаторов
	ui1 := V4()
	ui2 := V4()
	// Проверка вариантов
	vt1 := ui1.Variant()
	vt2 := ui2.Variant()
	if vt1 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv4-1: got %x, expected %x", vt1, bitVRFC4122)
	}
	if vt2 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv4-2: got %x, expected %x", vt2, bitVRFC4122)
	}
	// Проверка версий
	vn1 := ui1.Version()
	vn2 := ui2.Version()
	if vn1 != bitV4>>4 {
		t.Errorf("Invalid version for UUIDv4-1: got %x, expected %x", vn1, bitV4)
	}
	if vn2 != bitV4>>4 {
		t.Errorf("Invalid version for UUIDv4-2: got %x, expected %x", vn2, bitV4)
	}
	// Проверка детерминированности
	if ui1 == ui2 {
		t.Error("Generated same UUIDv4 for different data")
	}
}
func Test_UUIDV5_Generate(t *testing.T) {
	testCases := []struct {
		namespace UUID
		name      string
	}{
		{NameSpaceDNS, testNameString},
		{NameSpaceURL, testNameString},
		{NameSpaceOID, testNameString},
		{NameSpaceX500, testNameString},
	}
	for i, tc := range testCases {
		// Генерация идентификаторов
		ui := V5(tc.namespace, tc.name)
		// Проверка вариантов
		vt := ui.Variant()
		if vt != variantRFC4122 {
			t.Errorf("Invalid variant for UUIDv5-%d: got %x, expected %x", i+1, vt, bitVRFC4122)
		}
		// Проверка версий
		vn := ui.Version()
		if vn != bitV5>>4 {
			t.Errorf("Invalid version for UUIDv5-%d: got %x, expected %x", i+1, vn, bitV5)
		}
	}
}
func Test_UUIDV5_Hash(t *testing.T) {
	// Генерация идентификаторов
	ui1a := V5(NameSpaceDNS, testNameString+"/TestA")
	ui1b := V5(NameSpaceDNS, testNameString+"/TestB")
	ui2a := V5(NameSpaceURL, testNameString+"/TestA")
	ui2b := V5(NameSpaceURL, testNameString+"/TestB")
	// Проверка детерминированности при разных NameSpace
	if (ui1a == ui2a) || (ui1b == ui2b) {
		t.Error("Invalid UUIDv5 with the same value for different Namespace")
	}
	// Проверка детерминированности при разных TestNameString
	if (ui1a == ui1b) || (ui2a == ui2b) {
		t.Error("Invalid UUIDv5 with the same value for different TestNameString")
	}
}
func Test_UUIDV6_Generate(t *testing.T) {
	// Ленивая инициализация глобального состояния
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		t.Fatal(initError)
	}
	// Фиксация текущего состояние
	savedClock := initClock
	savedTime := v6.lastTime.Load()
	savedSequence := v6.lastSequence.Load()
	t.Cleanup(func() {
		initClock = savedClock
		v6.lastTime.Store(savedTime)
		v6.lastSequence.Store(savedSequence)
	})
	// Мокирование
	initClock = realClock{}
	v6.lastTime.Store(0)
	v6.lastSequence.Store(0)
	// Генерация идентификаторов
	ui1 := V6()
	ui2 := V6()
	// Проверка вариантов
	vt1 := ui1.Variant()
	vt2 := ui2.Variant()
	if vt1 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv6-1: got %x, expected %x", vt1, bitVRFC4122)
	}
	if vt2 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv6-2: got %x, expected %x", vt2, bitVRFC4122)
	}
	// Проверка версий
	vn1 := ui1.Version()
	vn2 := ui2.Version()
	if vn1 != bitV6>>4 {
		t.Errorf("Invalid version for UUIDv6-1: got %x, expected %x", vn1, bitV6)
	}
	if vn2 != bitV6>>4 {
		t.Errorf("Invalid version for UUIDv6-2: got %x, expected %x", vn2, bitV6)
	}
	// Проверка временных меток
	ts1 := ui1.Timestamp()
	ts2 := ui2.Timestamp()
	if ts1 <= 0 || ts2 <= 0 {
		t.Error("Invalid timestamp UUIDv6")
	}
	if ts1 > ts2 {
		t.Error("Failed monotonically increasing timestamp UUIDv6")
	}
	// Проверка детерминированности
	if ui1 == ui2 {
		t.Error("Failed to generate different UUIDs")
	}
}
func Test_UUIDV6_Sequence(t *testing.T) {
	t.Run("Sequence_Increment", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v6.lastTime.Load()
		realSequence := v6.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v6.lastTime.Store(realTime)
			v6.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v6.lastTime.Store(uint64((initClock.now().UnixNano() / 100) + offsetTime))
		v6.lastSequence.Store(0)
		// Генерация идентификаторов
		ui1 := V6()
		ui2 := V6()
		// Проверка временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if ts1 == ts2 && sq1 >= sq2 {
			t.Errorf("Failed to increment sequence: ts1=%d ts2=%d sq1=%d sq2=%d", ts1, ts2, sq1, sq2)
		}
	})
	t.Run("Sequence_Overflow", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v6.lastTime.Load()
		realSequence := v6.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v6.lastTime.Store(realTime)
			v6.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v6.lastTime.Store(uint64((initClock.now().UnixNano() / 100) + offsetTime))
		v6.lastSequence.Store(maxV6Sequence - 1)
		// Генерация идентификаторов
		ui1 := V6()
		mockClock.add(time.Nanosecond * 100)
		ui2 := V6()
		// Проверка временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if sq2 != 0 {
			t.Errorf("Failed to overflow sequence: sq1=%d sq2=%d", sq1, sq2)
		}
	})
}
func Test_UUIDV7_Generate(t *testing.T) {
	// Ленивая инициализация глобального состояния
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		t.Fatal(initError)
	}
	// Фиксация текущего состояние
	savedClock := initClock
	savedTime := v7.lastTime.Load()
	savedSequence := v7.lastSequence.Load()
	t.Cleanup(func() {
		initClock = savedClock
		v7.lastTime.Store(savedTime)
		v7.lastSequence.Store(savedSequence)
	})
	// Мокирование
	initClock = realClock{}
	v7.lastTime.Store(0)
	v7.lastSequence.Store(0)
	// Генерация идентификаторов
	ui1 := V7()
	ui2 := V7()
	// Проверка вариантов
	vt1 := ui1.Variant()
	vt2 := ui2.Variant()
	if vt1 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv7-1: got %x, expected %x", vt1, bitVRFC4122)
	}
	if vt2 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv7-2: got %x, expected %x", vt2, bitVRFC4122)
	}
	// Проверка версий
	vn1 := ui1.Version()
	vn2 := ui2.Version()
	if vn1 != bitV7>>4 {
		t.Errorf("Invalid version for UUIDv7-1: got %x, expected %x", vn1, bitV7)
	}
	if vn2 != bitV7>>4 {
		t.Errorf("Invalid version for UUIDv7-2: got %x, expected %x", vn2, bitV7)
	}
	// Проверка временных меток
	ts1 := ui1.Timestamp()
	ts2 := ui2.Timestamp()
	if ts1 <= 0 || ts2 <= 0 {
		t.Error("Invalid timestamp UUIDv7")
	}
	if ts1 > ts2 {
		t.Error("Failed monotonically increasing timestamp UUIDv7")
	}
	// Проверка детерминированности
	if ui1 == ui2 {
		t.Error("Failed to generate different UUIDs")
	}
}
func Test_UUIDV7_Sequence(t *testing.T) {
	// Подтест 1: sequence увеличение
	t.Run("Sequence_Increment", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v7.lastTime.Load()
		realSequence := v7.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v7.lastTime.Store(realTime)
			v7.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v7.lastTime.Store(uint64(initClock.now().UnixMilli()))
		v7.lastSequence.Store(0)
		// Генерация идентификаторов
		ui1 := V7()
		ui2 := V7()
		// Проверка временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if ts1 == ts2 && sq1 >= sq2 {
			t.Errorf("Failed to increment sequence: ts1=%d ts2=%d sq1=%d sq2=%d", ts1, ts2, sq1, sq2)
		}
	})
	// Подтест 2: sequence переполнение
	t.Run("Sequence_Overflow", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v7.lastTime.Load()
		realSequence := v7.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v7.lastTime.Store(realTime)
			v7.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v7.lastTime.Store(uint64(initClock.now().UnixMilli()))
		v7.lastSequence.Store(maxV7Sequence - 1)
		// Генерация идентификаторов
		ui1 := V7()
		mockClock.add(time.Millisecond)
		ui2 := V7()
		// Проверка временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if sq2 != 0 {
			t.Errorf("Failed to overflow sequence: sq1=%d sq2=%d", sq1, sq2)
		}
	})
}
func Test_UUIDV8_Generate(t *testing.T) {
	// Ленивая инициализация глобального состояния
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		t.Fatal(initError)
	}
	// Фиксация текущего состояние
	realClock := initClock
	realTime := v8.lastTime.Load()
	realSequence := v8.lastSequence.Load()
	t.Cleanup(func() {
		initClock = realClock
		v8.lastTime.Store(realTime)
		v8.lastSequence.Store(realSequence)
	})
	// Мокирование
	mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
	initClock = mockClock
	v8.lastTime.Store(uint64(initClock.now().UnixMilli()))
	v8.lastSequence.Store(0)
	// Генерация идентификаторов
	ui1 := V8(testNodeID)
	ui2 := V8(testNodeID)
	// Проверка вариантов
	vt1 := ui1.Variant()
	vt2 := ui2.Variant()
	if vt1 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv8-1: got %x, expected %x", vt1, bitVRFC4122)
	}
	if vt2 != variantRFC4122 {
		t.Errorf("Invalid variant for UUIDv8-2: got %x, expected %x", vt2, bitVRFC4122)
	}
	// Проверка версий
	vn1 := ui1.Version()
	vn2 := ui2.Version()
	if vn1 != bitV8>>4 {
		t.Errorf("Invalid version for UUIDv8-1: got %x, expected %x", vn1, bitV8)
	}
	if vn2 != bitV8>>4 {
		t.Errorf("Invalid version for UUIDv8-2: got %x, expected %x", vn2, bitV8)
	}
	// Проверка временных меток
	ts1 := ui1.Timestamp()
	ts2 := ui2.Timestamp()
	if ts1 <= 0 || ts2 <= 0 {
		t.Error("Invalid timestamp UUIDv8")
	}
	if ts1 > ts2 {
		t.Error("Failed monotonically increasing timestamp UUIDv8")
	}
	// Проверка идентификаторов нод
	nd1 := ui1.Node()
	nd2 := ui2.Node()
	if nd1 != testNodeID {
		t.Errorf("Invalid NodeID in UUIDv8-1: got %d, expected %d", nd1, testNodeID)
	}
	if nd2 != testNodeID {
		t.Errorf("Invalid NodeID in UUIDv8-2: got %d, expected %d", nd2, testNodeID)
	}
	// Проверка детерминированности
	if ui1 == ui2 {
		t.Error("Failed to generate different UUIDs")
	}
}
func Test_UUIDV8_Sequence(t *testing.T) {
	t.Run("Sequence_Increment", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v8.lastTime.Load()
		realSequence := v8.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v8.lastTime.Store(realTime)
			v8.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v8.lastTime.Store(uint64(initClock.now().UnixMilli()))
		v8.lastSequence.Store(0)
		// Генерация идентификаторов
		ui1 := V8(testNodeID)
		ui2 := V8(testNodeID)
		// Проверка временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if ts1 == ts2 && sq1 >= sq2 {
			t.Errorf("Failed to increment sequence: ts1=%d ts2=%d sq1=%d sq2=%d", ts1, ts2, sq1, sq2)
		}
	})
	t.Run("Sequence_Overflow", func(t *testing.T) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			t.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v8.lastTime.Load()
		realSequence := v8.lastSequence.Load()
		t.Cleanup(func() {
			initClock = realClock
			v8.lastTime.Store(realTime)
			v8.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v8.lastTime.Store(uint64(initClock.now().UnixMilli()))
		v8.lastSequence.Store(maxV8Sequence - 1)
		// Генерация идентификаторов
		ui1 := V8(testNodeID)
		mockClock.add(time.Millisecond)
		ui2 := V8(testNodeID)
		// Проверка временных меток и последовательностей
		ts1 := ui1.Timestamp()
		sq1 := ui1.Sequence()
		ts2 := ui2.Timestamp()
		sq2 := ui2.Sequence()
		if ts2 < ts1 {
			t.Errorf("Failed timestamp, time went backwards: ts1=%d ts2=%d", ts1, ts2)
		} else if sq2 != 0 {
			t.Errorf("Failed to overflow sequence: sq1=%d sq2=%d", sq1, sq2)
		}
	})
}

// Приватные переменные
var (
	testNameString               = "md.local"
	testNodeID                   = 1995
	testPOSType                  = 0
	testUUIDErrByteEmpty         = []byte{}
	testUUIDErrByteLengthLong    = [17]byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
	testUUIDErrByteLengthShort   = [15]byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
	testUUIDErrStringCharacter   = "01968727-8c7e-800x-87cb-bdba4f634d9f"
	testUUIDErrStringEmpty       = ""
	testUUIDErrStringInvalid     = "invalid-format"
	testUUIDErrStringLengthLong  = "01968727-8c7e-800000-87cb-bdba4f634d9f"
	testUUIDErrStringLengthShort = "01968727-8c7e-80-87cb-bdba4f634d9f"
	testUUIDErrStringLong        = "{01968727-8c7e-8000-87cb-bdba4f634d9f}"
	testUUIDErrStringShort       = "019687278c7e800087cbbdba4f634d9f"
	testUUIDErrTypeInt           = 19687278
	testUUIDV1Byte               = [16]byte{0x2b, 0xa0, 0x17, 0x4a, 0x20, 0x9d, 0x11, 0xf0, 0x80, 0x00, 0xac, 0xde, 0x48, 0x00, 0x11, 0x22}
	testUUIDV1String             = "2ba0174a-209d-11f0-8000-acde48001122"
	testUUIDV2Byte               = [16]byte{0x00, 0x00, 0x01, 0xf5, 0x2b, 0xa0, 0x20, 0x9d, 0x80, 0x00, 0xac, 0xde, 0x48, 0x00, 0x11, 0x22}
	testUUIDV2String             = "000001f5-2ba0-209d-8000-acde48001122"
	testUUIDV3Byte               = [16]byte{0x90, 0x73, 0x92, 0x6b, 0x92, 0x9f, 0x31, 0xc2, 0xab, 0xc9, 0xfa, 0xd7, 0x7a, 0xe3, 0xe8, 0xeb}
	testUUIDV3String             = "9073926b-929f-31c2-abc9-fad77ae3e8eb"
	testUUIDV4Byte               = [16]byte{0xae, 0x68, 0x2b, 0x8f, 0x49, 0xff, 0x46, 0x9c, 0x85, 0x28, 0xa3, 0xed, 0xe0, 0x52, 0xc6, 0x90}
	testUUIDV4String             = "ae682b8f-49ff-469c-8528-a3ede052c690"
	testUUIDV5Byte               = [16]byte{0x4f, 0xd3, 0x5a, 0x71, 0x71, 0xef, 0x5a, 0x55, 0xa9, 0xd9, 0xaa, 0x75, 0xc8, 0x89, 0xa6, 0xd0}
	testUUIDV5String             = "4fd35a71-71ef-5a55-a9d9-aa75c889a6d0"
	testUUIDV6Byte               = [16]byte{0x1f, 0x02, 0x09, 0xd2, 0xba, 0x01, 0x67, 0x9a, 0x80, 0x00, 0xac, 0xde, 0x48, 0x00, 0x11, 0x22}
	testUUIDV6String             = "1f0209d2-ba01-679a-8000-acde48001122"
	testUUIDV7Byte               = [16]byte{0x01, 0x96, 0x65, 0x0b, 0xad, 0x3b, 0x70, 0x00, 0x82, 0xb1, 0xce, 0x73, 0x41, 0x49, 0x23, 0x30}
	testUUIDV7String             = "0196650b-ad3b-7000-82b1-ce7341492330"
	testUUIDV8Byte               = [16]byte{0xaa, 0xbb, 0xcc, 0xdd, 0x11, 0x22, 0x83, 0x44, 0x95, 0x66, 0x4c, 0x84, 0xeb, 0x01, 0x58, 0x16}
	testUUIDV8String             = "aabbccdd-1122-8344-9566-4c84eb015816"
	testUUIDVUByte               = [16]byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
	testUUIDVUJson               = `"01968727-8c7e-8000-87cb-bdba4f634d9f"`
	testUUIDVUNull               = "00000000-0000-0000-0000-000000000000"
	testUUIDVUString             = "01968727-8c7e-8000-87cb-bdba4f634d9f"
	testUUIDVUText               = "01968727-8c7e-8000-87cb-bdba4f634d9f"
)

// Приватные функции
func must(uuid UUID, err error) UUID {
	if err != nil {
		panic(err)
	}
	return uuid
}
