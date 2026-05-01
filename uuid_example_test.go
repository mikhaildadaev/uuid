package uuid_test

import (
	"fmt"

	"github.com/mikhaildadaev/uuid"
)

// Примеры использования публичных функций
func ExampleNewNull() {
	un := uuid.NewNull()
	fmt.Println(un.IsZero())
	fmt.Println(un.String())
	// Output:
	// true
	// 00000000-0000-0000-0000-000000000000
}
func ExampleNewV1() {
	u1 := uuid.NewV1()
	fmt.Println(u1.Version())
	// Output:
	// 1
}
func ExampleNewV2() {
	u2 := uuid.NewV2(uuidPosType, uuidPosValue)
	fmt.Println(u2.Version())
	// Output:
	// 2
}
func ExampleNewV3() {
	u3 := uuid.NewV3(uuid.NameSpaceDNS, uuidName)
	fmt.Println(u3.Version())
	// Output:
	// 3
}
func ExampleNewV4() {
	u4 := uuid.NewV4()
	fmt.Println(u4.Version())
	// Output:
	// 4
}
func ExampleNewV5() {
	u5 := uuid.NewV5(uuid.NameSpaceDNS, uuidName)
	fmt.Println(u5.Version())
	// Output:
	// 5
}
func ExampleNewV6() {
	u6 := uuid.NewV6()
	fmt.Println(u6.Version())
	// Output:
	// 6
}
func ExampleNewV7() {
	u7 := uuid.NewV7()
	fmt.Println(u7.Version())
	// Output:
	// 7
}
func ExampleNewV8() {
	u8 := uuid.NewV8(uuidNode)
	fmt.Println(u8.Version())
	// Output:
	// 8
}
func ExampleParse() {
	u8, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u8.IsZero())
	// Output:
	// false
}

// Примеры использования публичных методов
func ExampleUUID_Bytes() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x\n", uu.Bytes())
	// Output:
	// 019687278c7e800087cbbdba4f634d9f
}
func ExampleUUID_Equal() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	other := uu
	fmt.Println(uu.Equal(other))
	// Output:
	// true
}
func ExampleUUID_Info() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.Info())
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
	uu := uuid.NewNull()
	fmt.Println(uu.IsZero())
	// Output:
	// true
}
func ExampleUUID_Node() {
	u8 := uuid.NewV8(uuidNode)
	fmt.Println(u8.Node())
	// Output:
	// 1995
}
func ExampleUUID_Posix() {
	u2 := uuid.NewV2(uuidPosType, uuidPosValue)
	fmt.Println(u2.Posix())
	// Output:
	// UID 501
}
func ExampleUUID_Sequence() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.Sequence())
	// Output:
	// 0
}
func ExampleUUID_String() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.String())
	// Output:
	// 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleUUID_Timestamp() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.Timestamp())
	// Output:
	// 1746024238206
}
func ExampleUUID_Validate() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.Validate())
	// Output:
	// <nil>
}
func ExampleUUID_Variant() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.Variant())
	// Output:
	// 1
}
func ExampleUUID_Version() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.Version())
	// Output:
	// 8
}
func ExampleUUID_MarshalBinary() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	data, err := uu.MarshalBinary()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x", data)
	// Output: 019687278c7e800087cbbdba4f634d9f
}
func ExampleUUID_MarshalJson() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	json, err := uu.MarshalJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
	// Output: "01968727-8c7e-8000-87cb-bdba4f634d9f"
}
func ExampleUUID_MarshalText() {
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	text, err := uu.MarshalText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(text))
	// Output: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleUUID_UnmarshalBinary() {
	var uu uuid.UUID
	if err := uu.UnmarshalBinary(uuidV8Binary); err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.String())
	// Output: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleUUID_UnmarshalJson() {
	var uu uuid.UUID
	if err := uu.UnmarshalJson(uuidV8Json); err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.String())
	// Output: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleUUID_UnmarshalText() {
	var uu uuid.UUID
	if err := uu.UnmarshalText(uuidV8Text); err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.String())
	// Output: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleNullUUID_MarshalBinary() {
	var nu uuid.NullUUID
	data, err := nu.MarshalBinary()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x\n", data)
	nu.Scan(uuidV8String)
	data, err = nu.MarshalBinary()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x\n", data)
	// Output:
	//
	// 019687278c7e800087cbbdba4f634d9f
}
func ExampleNullUUID_MarshalJson() {
	var nu uuid.NullUUID
	json, err := nu.MarshalJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
	nu.Scan(uuidV8String)
	json, err = nu.MarshalJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
	// Output:
	// null
	// "01968727-8c7e-8000-87cb-bdba4f634d9f"
}
func ExampleNullUUID_MarshalText() {
	var nu uuid.NullUUID
	text, err := nu.MarshalText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(text))
	nu.Scan(uuidV8String)
	text, err = nu.MarshalText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(text))
	// Output:
	//
	// 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleNullUUID_UnmarshalBinary() {
	var nu uuid.NullUUID
	data := []byte{}
	if err := nu.UnmarshalBinary(data); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	data = uuidV8Binary
	if err := nu.UnmarshalBinary(data); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	fmt.Println("UUID:", nu.UUID)
	// Output:
	// Valid: false
	// Valid: true
	// UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleNullUUID_UnmarshalJson() {
	var nu uuid.NullUUID
	json := []byte(`null`)
	if err := nu.UnmarshalJson(json); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	json = []byte(`"01968727-8c7e-8000-87cb-bdba4f634d9f"`)
	if err := nu.UnmarshalJson(json); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	fmt.Println("UUID:", nu.UUID)
	// Output:
	// Valid: false
	// Valid: true
	// UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleNullUUID_UnmarshalText() {
	var nu uuid.NullUUID
	text := []byte("")
	if err := nu.UnmarshalText(text); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	text = uuidV8Text
	if err := nu.UnmarshalText(text); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	fmt.Println("UUID:", nu.UUID)
	// Output:
	// Valid: false
	// Valid: true
	// UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleUUID_Scan() {
	var uu uuid.UUID
	err := uu.Scan(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uu.String())
	// Output: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleUUID_Value() {
	var nilUUID uuid.UUID
	uu, err := uuid.Parse(uuidV8String)
	if err != nil {
		fmt.Println(err)
	}
	value, err := uu.Value()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
	value, err = nilUUID.Value()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
	// Output:
	// 01968727-8c7e-8000-87cb-bdba4f634d9f
	// <nil>
}
func ExampleNullUUID_Scan() {
	var nu uuid.NullUUID
	if err := nu.Scan(nil); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	if err := nu.Scan(uuidV8String); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Valid:", nu.Valid)
	fmt.Println("UUID:", nu.UUID)
	// Output:
	// Valid: false
	// Valid: true
	// UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
}
func ExampleNullUUID_Value() {
	var nu uuid.NullUUID
	value, _ := nu.Value()
	fmt.Println("NULL value:", value)
	nu.Scan(uuidV8String)
	value, _ = nu.Value()
	fmt.Println("UUID value:", value)
	// Output:
	// NULL value: <nil>
	// UUID value: 01968727-8c7e-8000-87cb-bdba4f634d9f
}

// Приватные переменные
var (
	uuidName     = "github.com/mikhaildadaev/uuid"
	uuidNode     = 1995
	uuidPosType  = 0
	uuidPosValue = 501
	uuidV8Binary = []byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
	uuidV8Json   = []byte(`"01968727-8c7e-8000-87cb-bdba4f634d9f"`)
	uuidV8String = "01968727-8c7e-8000-87cb-bdba4f634d9f"
	uuidV8Text   = []byte("01968727-8c7e-8000-87cb-bdba4f634d9f")
)
