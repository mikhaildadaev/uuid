package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Публичные типы
type UUID [16]byte

// Публичные структуры
type NullUUID struct {
	UUID  UUID
	Valid bool
}

// Публичные константы
const (
	Author  = "Mikhail Dadaev"
	Version = "1.26.11"
)

// Публичные переменные
var (
	ErrInvalidUUIDLength   = errors.New("invalid UUID length")
	ErrInvalidUUIDMAC      = errors.New("invalid UUID mac")
	ErrInvalidUUIDPOSIX    = errors.New("invalid UUID posix")
	ErrInvalidUUIDString   = errors.New("invalid UUID string")
	ErrInvalidUUIDVariant  = errors.New("invalid UUID variant")
	ErrInvalidUUIDVersion  = errors.New("invalid UUID version")
	ErrNullUUID            = errors.New("null UUID")
	ErrNullUUIDNotAllowed  = errors.New("null UUID not allowed")
	ErrUnsupportedUUIDType = errors.New("unsupported UUID type")
)
var (
	NameSpaceDNS   = [16]byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NameSpaceURL   = [16]byte{0x6b, 0xa7, 0xb8, 0x11, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NameSpaceOID   = [16]byte{0x6b, 0xa7, 0xb8, 0x12, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NameSpaceX500  = [16]byte{0x6b, 0xa7, 0xb8, 0x14, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NullUUIDBinary = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	NullUUIDString = "00000000-0000-0000-0000-000000000000"
)

// Публичные функции
func NewNull() NullUUID {
	return NullUUID{Valid: false}
}
func NewV1() UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	// Получение временной метки и последовательности
	ts, sq := getTimeNanoAndSequence("v1")
	// Получение MAC-адреса
	mac := initMAC.Load().([6]byte)
	// Установка временной метки и версии UUID
	binary.BigEndian.PutUint32(uuid[0:], uint32(ts)&0xFFFFFFFF)
	binary.BigEndian.PutUint16(uuid[4:], uint16(ts>>32)&0xFFFF)
	binary.BigEndian.PutUint16(uuid[6:], uint16(ts>>48)&0x0FFF|bitV1<<8)
	// Установка последовательности и RFC 4122 вариант
	uuid[8] = byte((sq>>8)&0x3F | bitVRFC4122)
	uuid[9] = byte(sq & 0xFF)
	// Установка MAC-адреса
	copy(uuid[10:16], mac[:])
	return uuid
}
func NewV2(posType int, posValue ...int) UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var pi uint32
	var uuid UUID
	// Проверка POSType (UID/GID и др.)
	posType = max(0, min(posType, maxV2POSType))
	if len(posValue) == 1 {
		// Переданный POS-идентификатора
		pi = uint32(posValue[0])
	} else {
		// Извлечение POS-идентификатора
		pi = getPOSIX(posType)
	}
	// Получение MAC-адреса
	mac := initMAC.Load().([6]byte)
	// Установка рандомных данных
	genRandCrypto(uuid[4:9])
	// Установка POSIX
	binary.BigEndian.PutUint32(uuid[0:4], pi)
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV2
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	// Установка POSTypeID
	uuid[9] = uint8(posType & 0x03)
	// Установка MAC-адреса
	copy(uuid[10:16], mac[:])
	return uuid
}
func NewV3(nameSpace UUID, nameString string) UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	var hashBufPool = hashMD5Pool.Get().(*struct {
		buf  [md5.Size + bufSize]byte
		hash [md5.Size]byte
	})
	defer hashMD5Pool.Put(hashBufPool)
	// Проверка nameString
	if len(nameString) == 0 || len(nameString) > bufSize {
		return NullUUIDBinary
	}
	// Формирование данных и хеша
	buf := hashBufPool.buf[:0]
	buf = append(buf, nameSpace[:]...)
	buf = append(buf, nameString...)
	hash := md5.Sum(buf)
	// Установка хеша
	copy(uuid[:], hash[:])
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV3
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	return uuid
}
func NewV4() UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	// Установка рандомных данных
	genRandCrypto(uuid[:])
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV4
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	return uuid
}
func NewV5(nameSpace UUID, nameString string) UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	var hashBufPool = hashSHA1Pool.Get().(*struct {
		buf  [sha1.Size + bufSize]byte
		hash [sha1.Size]byte
	})
	defer hashSHA1Pool.Put(hashBufPool)
	// Проверка nameString
	if len(nameString) == 0 || len(nameString) > bufSize {
		return NullUUIDBinary
	}
	// Формирование данных и хеша
	buf := hashBufPool.buf[:0]
	buf = append(buf, nameSpace[:]...)
	buf = append(buf, nameString...)
	hash := sha1.Sum(buf)
	// Установка хеша
	copy(uuid[:], hash[:])
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV5
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	return uuid
}
func NewV6() UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	// Получение временной метки и последовательности
	ts, sq := getTimeNanoAndSequence("v6")
	// Получение MAC-адреса
	mac := initMAC.Load().([6]byte)
	// Установка временной метки и версии UUID
	binary.BigEndian.PutUint32(uuid[0:], uint32(ts>>28)&0xFFFFFFFF)
	binary.BigEndian.PutUint16(uuid[4:], uint16(ts>>12)&0xFFFF)
	binary.BigEndian.PutUint16(uuid[6:], uint16(ts)&0x0FFF|bitV6<<8)
	// Установка последовательности и RFC 4122 вариант
	uuid[8] = byte((sq>>8)&0x3F | bitVRFC4122)
	uuid[9] = byte(sq & 0xFF)
	// Установка MAC-адреса
	copy(uuid[10:16], mac[:])
	return uuid
}
func NewV7() UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	// Получение временной метки и последовательности
	ts, sq := getTimeMilliAndSequence("v7")
	// Установка временной метки и версии UUID
	binary.BigEndian.PutUint64(uuid[0:], ts<<16)
	binary.BigEndian.PutUint16(uuid[6:8], uint16(sq)&maxV7Sequence|bitV7<<8)
	// Установка рандомных данных
	genRandCrypto(uuid[8:])
	// Установка RFC 4122 вариант
	uuid[8] = (uuid[8]&0x3F | bitVRFC4122)
	return uuid
}
func NewV8(nodeID int) UUID {
	initSync.Do(func() {
		initError = initGlobal()
	})
	var uuid UUID
	// Получение временной метки и последовательности
	ts, sq := getTimeMilliAndSequence("v8")
	// Получение идентификатора ноды
	nodeID = max(0, min(nodeID, maxV8NodeID))
	// Установка рандомных данных
	genRandCrypto(uuid[10:])
	// Установка временной метки
	binary.BigEndian.PutUint64(uuid[0:], ts<<16)
	// Установка последовательности и версии UUID
	binary.BigEndian.PutUint16(uuid[6:8], uint16(sq)&maxV8Sequence|bitV8<<8)
	// Установка идентификатора ноды и варианта RFC 4122
	binary.BigEndian.PutUint16(uuid[8:10], uint16(nodeID)&maxV8NodeID|bitVRFC4122<<8)
	return uuid
}
func GetAuthor() string {
	return Author
}
func GetCopyright() string {
	Copyright := fmt.Sprintf("Copyright © 2022-%d %s. All rights reserved.", time.Now().Year(), Author)
	return Copyright
}
func GetVersion() string {
	return Version
}
func Parse(str string) (UUID, error) {
	var ln = strings.NewReplacer("{", "", "}", "", "-", "").Replace(str)
	var ui UUID
	if len(ln) != 32 {
		return UUID{}, ErrInvalidUUIDLength
	}
	if _, err := hex.Decode(ui[:], []byte(ln)); err != nil {
		return UUID{}, ErrInvalidUUIDString
	}
	return ui, nil
}

// Публичные методы
func (nulluuid NullUUID) IsZero() bool {
	return !nulluuid.Valid || nulluuid.UUID.IsZero()
}
func (nulluuid NullUUID) String() string {
	if !nulluuid.Valid {
		return NullUUIDString
	}
	return nulluuid.UUID.String()
}
func (nulluuid NullUUID) Validate() error {
	if nulluuid.Valid {
		return nulluuid.UUID.Validate()
	}
	return nil
}
func (uuid UUID) Bytes() []byte {
	if len(uuid) != 16 {
		return nil
	}
	buf := make([]byte, 16)
	copy(buf, uuid[:])
	return buf
}
func (uuid UUID) Equal(other UUID) bool {
	return uuid == other
}
func (uuid UUID) Info() string {
	var info strings.Builder
	var vn = uuid.Version()
	var vt = uuid.Variant()
	info.Grow(256)
	info.WriteString(fmt.Sprintf("UUID: %s\n", uuid.String()))
	switch uuid.Version() {
	case bitV1 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: TTTTTTTT-TTTT-1TTT-VSSS-MMMMMMMMMMMM\n")
		info.WriteString("INFO: TIME (100-nanoseconds interval since 1582-10-15) + SEQUENCE (0-16383) + MAC\n")
		info.WriteString(fmt.Sprintf("TIME: %d [%s]\n", ts, (time.Unix(0, (ts-offsetTime)*100).UTC())))
		info.WriteString(fmt.Sprintf("SEQ.: %d\n", sq))
		info.WriteString(fmt.Sprintf("MAC.: %02x:%02x:%02x:%02x:%02x:%02x\n", uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))
	case bitV2 >> 4:
		psstr, psuint := uuid.Posix()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: PPPPPPPP-RRRR-1RRR-VRXX-MMMMMMMMMMMM\n")
		info.WriteString("INFO: POSIX + RANDOM + POSTYPE + MAC\n")
		info.WriteString(fmt.Sprintf("POS.: %d (%s)\n", psuint, psstr))
		info.WriteString(fmt.Sprintf("RAND: %x\n", uuid[4:10]))
		info.WriteString(fmt.Sprintf("MAC.: %02x:%02x:%02x:%02x:%02x:%02x\n", uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))
	case bitV3 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: HHHHHHHH-HHHH-3HHH-VHHH-HHHHHHHHHHHH\n")
		info.WriteString("INFO: HASH-MD5 (namespase+nameline)\n")
		info.WriteString(fmt.Sprintf("HASH: %x\n", uuid[0:16]))
	case bitV4 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: RRRRRRRR-RRRR-4RRR-VRRR-RRRRRRRRRRRR\n")
		info.WriteString("INFO: RANDOM\n")
		info.WriteString(fmt.Sprintf("RAND: %x\n", uuid[0:16]))
	case bitV5 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: HHHHHHHH-HHHH-5HHH-VHHH-HHHHHHHHHHHH\n")
		info.WriteString("INFO: HASH-SHA1 (namespase+nameline)\n")
		info.WriteString(fmt.Sprintf("HASH: %x\n", uuid[0:16]))
	case bitV6 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: TTTTTTTT-TTTT-6TTT-VSSS-MMMMMMMMMMMM\n")
		info.WriteString("INFO: TIME (Reordered 100-nanoseconds interval since 1582-10-15) + SEQUENCE (0-16383) + MAC\n")
		info.WriteString(fmt.Sprintf("TIME: %d [%s]\n", ts, (time.Unix(0, (ts-offsetTime)*100).UTC())))
		info.WriteString(fmt.Sprintf("SEQ.: %d\n", sq))
		info.WriteString(fmt.Sprintf("MAC.: %02x:%02x:%02x:%02x:%02x:%02x\n", uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))
	case bitV7 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: TTTTTTTT-TTTT-7SSS-VRRR-RRRRRRRRRRRR\n")
		info.WriteString("INFO: TIME (1-milliseconds interval since 1970-01-01) + SEQUENCE (0-4095) + RANDOM\n")
		info.WriteString(fmt.Sprintf("TIME: %d [%s]\n", ts, (time.UnixMilli(ts).UTC())))
		info.WriteString(fmt.Sprintf("SEQ.: %d\n", sq))
		info.WriteString(fmt.Sprintf("RAND: %x\n", uuid[8:16]))
	case bitV8 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		nd := uuid.Node()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", transformVariant(vt)))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: TTTTTTTT-TTTT-8SSS-VNNN-RRRRRRRRRRRR\n")
		info.WriteString("INFO: TIME (1-milliseconds interval since 1970-01-01) + SEQUENCE (0-4095) + NODE (0-16383) + RANDOM\n")
		info.WriteString(fmt.Sprintf("TIME: %d [%s]\n", ts, (time.UnixMilli(ts).UTC())))
		info.WriteString(fmt.Sprintf("SEQ.: %d\n", sq))
		info.WriteString(fmt.Sprintf("NODE: %d\n", nd))
		info.WriteString(fmt.Sprintf("RAND: %x\n", uuid[10:16]))
	}
	return info.String()
}
func (uuid UUID) IsZero() bool {
	return uuid == NullUUIDBinary
}
func (uuid UUID) Node() int {
	switch uuid.Version() {
	case bitV8 >> 4:
		return int(binary.BigEndian.Uint16(uuid[8:10]) & maxV8NodeID)
	default:
		return 0
	}
}
func (uuid UUID) Posix() (string, uint32) {
	switch uuid[9] {
	case 0:
		return "UID", binary.BigEndian.Uint32(uuid[0:4])
	case 1:
		return "GID", binary.BigEndian.Uint32(uuid[0:4])
	default:
		return "RID", binary.BigEndian.Uint32(uuid[0:4])
	}
}
func (uuid UUID) Sequence() int64 {
	switch uuid.Version() {
	case bitV1 >> 4:
		return int64(binary.BigEndian.Uint16(uuid[8:10]) & maxV1Sequence)
	case bitV6 >> 4:
		return int64(binary.BigEndian.Uint16(uuid[8:10]) & maxV6Sequence)
	case bitV7 >> 4:
		return int64(binary.BigEndian.Uint16(uuid[6:8]) & maxV7Sequence)
	case bitV8 >> 4:
		return int64(binary.BigEndian.Uint16(uuid[6:8]) & maxV8Sequence)
	default:
		return 0
	}
}
func (uuid UUID) String() string {
	buf := initCachePool.Get().(*[36]byte)
	defer initCachePool.Put(buf)
	encodeHex(buf[:], uuid)
	return string(buf[:])
}
func (uuid UUID) Timestamp() int64 {
	switch uuid.Version() {
	case bitV1 >> 4:
		return (int64(binary.BigEndian.Uint16(uuid[6:8])&0x0FFF)<<48 | int64(binary.BigEndian.Uint16(uuid[4:6]))<<32 | int64(binary.BigEndian.Uint32(uuid[0:4])))
	case bitV6 >> 4:
		return (int64(binary.BigEndian.Uint16(uuid[6:8])&0x0FFF) | int64(binary.BigEndian.Uint16(uuid[4:6]))<<12 | int64(binary.BigEndian.Uint32(uuid[0:4]))<<28)
	case bitV7 >> 4:
		return (int64(binary.BigEndian.Uint32(uuid[0:4]))<<16 | int64(binary.BigEndian.Uint16(uuid[4:6])))
	case bitV8 >> 4:
		return (int64(binary.BigEndian.Uint32(uuid[0:4]))<<16 | int64(binary.BigEndian.Uint16(uuid[4:6])))
	default:
		return 0
	}
}
func (uuid UUID) Validate() error {
	if uuid.IsZero() {
		return ErrNullUUID
	}
	switch uuid.Variant() {
	case variantRFC4122:
		// Корректный вариант
	case variantNCS, variantMicrosoft, variantReserved:
		return ErrInvalidUUIDVariant
	default:
		return ErrInvalidUUIDVariant
	}
	version := uuid.Version()
	if version < 1 || version > 8 {
		return ErrInvalidUUIDVersion
	}
	switch version {
	case 1, 6:
		macIsZero := true
		for _, b := range uuid[10:16] {
			if b != 0 {
				macIsZero = false
				break
			}
		}
		if macIsZero {
			return ErrInvalidUUIDMAC
		}
	case 2:
		if binary.BigEndian.Uint32(uuid[0:4]) == 0 {
			return ErrInvalidUUIDPOSIX
		}
	}
	return nil
}
func (uuid UUID) Variant() int {
	if uuid == NullUUIDBinary {
		return variantInvalid
	}
	switch {
	case uuid[8] == bitVNCS:
		return variantNCS
	case uuid[8]&bitVRFC4122 == bitVRFC4122:
		return variantRFC4122
	case uuid[8]&bitVMS == bitVMS:
		return variantMicrosoft
	case uuid[8]&bitVReserved == bitVReserved:
		return variantReserved
	default:
		return variantInvalid
	}
}
func (uuid UUID) Version() int {
	if uuid == NullUUIDBinary {
		return 0
	}
	return int(uuid[6] >> 4)
}
