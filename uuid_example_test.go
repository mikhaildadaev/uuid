// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.

package uuid_test

import (
	"fmt"
	"log"

	"github.com/mikhaildadaev/uuid"
)

// Example functions
func ExampleNull() {
	u := uuid.Null()
	fmt.Println(u.IsZero())
	fmt.Println(u.String())
	// Output:
	// true
	// 00000000-0000-0000-0000-000000000000
}
func ExampleParse() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.IsZero())
	// Output:
	// false
}
func ExampleV1() {
	u1 := uuid.V1()
	fmt.Println(u1.Version())
	// Output:
	// 1
}
func ExampleV2() {
	u2 := uuid.V2(posix)
	fmt.Println(u2.Version())
	// Output:
	// 2
}
func ExampleV3() {
	u3 := uuid.V3(uuid.NameSpaceDNS, name)
	fmt.Println(u3.Version())
	// Output:
	// 3
}
func ExampleV4() {
	u4 := uuid.V4()
	fmt.Println(u4.Version())
	// Output:
	// 4
}
func ExampleV5() {
	u5 := uuid.V5(uuid.NameSpaceDNS, name)
	fmt.Println(u5.Version())
	// Output:
	// 5
}
func ExampleV6() {
	u6 := uuid.V6()
	fmt.Println(u6.Version())
	// Output:
	// 6
}
func ExampleV7() {
	u7 := uuid.V7()
	fmt.Println(u7.Version())
	// Output:
	// 7
}
func ExampleV8() {
	u8 := uuid.V8(node)
	fmt.Println(u8.Version())
	// Output:
	// 8
}

// Example methods
func ExampleUUID_Bytes() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	b := u.Bytes()
	fmt.Printf("%x", b)
	// Output:
	// 019687278c7e800087cbbdba4f634d9f
}
func ExampleUUID_Equal() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	other := u
	fmt.Println(u.Equal(other))
	// Output:
	// true
}
func ExampleUUID_Info() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	info := u.Info()
	fmt.Println(info)
	// Output:
	// UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
	// VAR.: RFC4122
	// VER.: 8
	// FORM: TTTTTTTT-TTTT-8SSS-VNNN-RRRRRRRRRRRR
	// INFO: TIME (1-milliseconds interval since 1970-01-01) + SEQUENCE (0-4095) + NODE (0-16383) + RANDOM
	// TIME: 1746024238206 [2025-04-30 14:43:58.206 +0000 UTC]
	// SEQ.: 0
	// NODE: 1995
	// RAND: bdba4f634d9f
}
func ExampleUUID_IsZero() {
	var u uuid.UUID
	fmt.Println(u.IsZero())
	// Output:
	// true
}
func ExampleUUID_Node() {
	u8 := uuid.V8(node)
	fmt.Println(u8.Node())
	// Output:
	// 1995
}
func ExampleUUID_Posix() {
	u, err := uuid.Parse(uuidStringV2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Posix())
	// Output:
	// UID 501
}
func ExampleUUID_Sequence() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Sequence())
	// Output:
	// 0
}
func ExampleUUID_String() {
	u4 := uuid.V4()
	fmt.Println(u4.Version())
	fmt.Println(len(u4.String()))
	// Output:
	// 4
	// 36
}
func ExampleUUID_Timestamp() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Timestamp())
	// Output:
	// 1746024238206
}
func ExampleUUID_Variant() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Variant())
	// Output:
	// 1
}
func ExampleUUID_Version() {
	u, err := uuid.Parse(uuidStringV8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Version())
	// Output:
	// 8
}

// Приватные переменные
var (
	name         = "github.com/mikhaildadaev/uuid"
	node         = 1995
	posix        = 0
	uuidStringV2 = "000001f5-dd95-2565-9600-acde48001122"
	uuidStringV8 = "01968727-8c7e-8000-87cb-bdba4f634d9f"
)
