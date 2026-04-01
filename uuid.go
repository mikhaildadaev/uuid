// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Публичные константы
const (
	Author  = "Mikhail Dadaev"
	Version = "1.26.3"
)

// Публичные переменные
var (
	ErrFailedCryptoRand   = errors.New("failed crypto/rand")
	ErrFailedGenMAC       = errors.New("failed gen/mac")
	ErrFailedGenSequences = errors.New("failed gen/sequences")
	ErrInvalidNameLine    = errors.New("invalid nameline")
	ErrInvalidNameSpase   = errors.New("invalid namespace")
	ErrInvalidUUIDLength  = errors.New("invalid UUID length")
	ErrInvalidUUIDMAC     = errors.New("invalid UUID mac")
	ErrInvalidUUIDNode    = errors.New("invalid UUID node")
	ErrInvalidUUIDPOSIX   = errors.New("invalid UUID posix")
	ErrInvalidUUIDPosType = errors.New("invalid UUID postype")
	ErrInvalidUUIDString  = errors.New("invalid UUID string")
	ErrInvalidUUIDVariant = errors.New("invalid UUID variant")
	ErrInvalidUUIDVersion = errors.New("invalid UUID version")
	ErrNilUUID            = errors.New("nil UUID")
)
var (
	NameSpaceDNS  = [16]byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NameSpaceURL  = [16]byte{0x6b, 0xa7, 0xb8, 0x11, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NameSpaceOID  = [16]byte{0x6b, 0xa7, 0xb8, 0x12, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NameSpaceX500 = [16]byte{0x6b, 0xa7, 0xb8, 0x14, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NilUUIDByte   = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	NilUUIDString = "00000000-0000-0000-0000-000000000000"
)

// Публичные типы
type UUID [16]byte

// Публичные структуры
type NullUUID struct {
	UUID  UUID
	Valid bool
}

// Публичные функции
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
func Null() NullUUID {
	return NullUUID{Valid: false}
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
	var vt = "Unknown"
	info.Grow(256)
	info.WriteString(fmt.Sprintf("UUID: %s\n", uuid.String()))
	switch uuid.Variant() {
	case variantNCS:
		vt = "NCS"
	case variantRFC4122:
		vt = "RFC 4122"
	case variantMicrosoft:
		vt = "Microsoft"
	case variantReserved:
		vt = "Reserved"
	default:
		vt = "Invalid"
	}
	switch uuid.Version() {
	case bitV1 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: TTTTTTTT-TTTT-1TTT-VSSS-MMMMMMMMMMMM\n")
		info.WriteString("INFO: TIME (100-nanoseconds interval since 1582-10-15) + SEQUENCE (0-16383) + MAC\n")
		info.WriteString(fmt.Sprintf("TIME: %d [%s]\n", ts, (time.Unix(0, (ts-offsetTime)*100).UTC())))
		info.WriteString(fmt.Sprintf("SEQ.: %d\n", sq))
		info.WriteString(fmt.Sprintf("MAC.: %02x:%02x:%02x:%02x:%02x:%02x\n", uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))
	case bitV2 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: PPPPPPPP-RRRR-1RRR-VRXX-MMMMMMMMMMMM\n")
		info.WriteString("INFO: POSID + RANDOM + POSTYPE + MAC\n")
		switch uuid[9] {
		case 0:
			info.WriteString(fmt.Sprintf("POS.: %d (UID)\n", binary.BigEndian.Uint32(uuid[0:4])))
		case 1:
			info.WriteString(fmt.Sprintf("POS.: %d (GID)\n", binary.BigEndian.Uint32(uuid[0:4])))
		default:
			info.WriteString(fmt.Sprintf("POS.: %d (ReservedID)\n", binary.BigEndian.Uint32(uuid[0:4])))
		}
		info.WriteString(fmt.Sprintf("RAND: %x\n", uuid[4:10]))
		info.WriteString(fmt.Sprintf("MAC.: %02x:%02x:%02x:%02x:%02x:%02x\n", uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))
	case bitV3 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: HHHHHHHH-HHHH-3HHH-VHHH-HHHHHHHHHHHH\n")
		info.WriteString("INFO: HASH-MD5 (namespase+nameline)\n")
		info.WriteString(fmt.Sprintf("HASH: %x\n", uuid[0:16]))
	case bitV4 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: RRRRRRRR-RRRR-4RRR-VRRR-RRRRRRRRRRRR\n")
		info.WriteString("INFO: RANDOM\n")
		info.WriteString(fmt.Sprintf("RAND: %x\n", uuid[0:16]))
	case bitV5 >> 4:
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: HHHHHHHH-HHHH-5HHH-VHHH-HHHHHHHHHHHH\n")
		info.WriteString("INFO: HASH-SHA1 (namespase+nameline)\n")
		info.WriteString(fmt.Sprintf("HASH: %x\n", uuid[0:16]))
	case bitV6 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
		info.WriteString(fmt.Sprintf("VER.: %d\n", vn))
		info.WriteString("FORM: TTTTTTTT-TTTT-6TTT-VSSS-MMMMMMMMMMMM\n")
		info.WriteString("INFO: TIME (Reordered 100-nanoseconds interval since 1582-10-15) + SEQUENCE (0-16383) + MAC\n")
		info.WriteString(fmt.Sprintf("TIME: %d [%s]\n", ts, (time.Unix(0, (ts-offsetTime)*100).UTC())))
		info.WriteString(fmt.Sprintf("SEQ.: %d\n", sq))
		info.WriteString(fmt.Sprintf("MAC.: %02x:%02x:%02x:%02x:%02x:%02x\n", uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))
	case bitV7 >> 4:
		ts := uuid.Timestamp()
		sq := uuid.Sequence()
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
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
		info.WriteString(fmt.Sprintf("VAR.: %s\n", vt))
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
	return uuid == NilUUIDByte
}
func (uuid UUID) Node() int {
	switch uuid.Version() {
	case bitV8 >> 4:
		return int(binary.BigEndian.Uint16(uuid[8:10]) & maxV8NodeID)
	default:
		return 0
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
		return ErrNilUUID
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
		macValid := true
		for _, b := range uuid[10:16] {
			if b != 0 {
				macValid = false
				break
			}
		}
		if macValid {
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
	if uuid == NilUUIDByte {
		return variantInvalid
	}
	switch {
	case uuid[8]&0xE0 == bitVNCS:
		return variantNCS
	case uuid[8]&0xE0 == bitVRFC4122:
		return variantRFC4122
	case uuid[8]&0xE0 == bitVMS:
		return variantMicrosoft
	case uuid[8]&0xE0 == bitVReserved:
		return variantReserved
	default:
		return variantInvalid
	}
}
func (uuid UUID) Version() int {
	if uuid == NilUUIDByte {
		return 0
	}
	return int(uuid[6] >> 4)
}
func (nulluuid NullUUID) Validate() error {
	if !nulluuid.Valid {
		return ErrNilUUID
	}
	return nulluuid.UUID.Validate()
}

// Приватные константы
const (
	bitMulticast     = 0x01
	bitLocal         = 0x02
	bitV1            = 0x10
	bitV2            = 0x20
	bitV3            = 0x30
	bitV4            = 0x40
	bitV5            = 0x50
	bitV6            = 0x60
	bitV7            = 0x70
	bitV8            = 0x80
	bitVNCS          = 0x00
	bitVRFC4122      = 0x80
	bitVMS           = 0xC0
	bitVReserved     = 0xE0
	hexTable         = "0123456789abcdef"
	maxV1Sequence    = 0x3FFF
	maxV2POSType     = 0xFF
	maxV6Sequence    = 0x3FFF
	maxV7Sequence    = 0xFFF
	maxV8NodeID      = 0x3FFF
	maxV8Sequence    = 0xFFF
	offsetTime       = 0x01B21DD213814000
	randbufSize      = 512
	variantInvalid   = -1
	variantNCS       = 0
	variantRFC4122   = 1
	variantMicrosoft = 2
	variantReserved  = 3
)

// Приватные типы
type version struct {
	lastSequence *atomic.Uint32
	lastTime     *atomic.Uint64
}

// Приватные переменные
var (
	cacheGID atomic.Uint32
	cacheUID atomic.Uint32

	initCachePool = sync.Pool{
		New: func() any {
			return new([36]byte)
		},
	}
	initClock    Clock = &RealClock{}
	initError    error
	initMAC      atomic.Value
	initRandPool = sync.Pool{
		New: func() any {
			buf := make([]byte, randbufSize)
			_, _ = rand.Read(buf)
			return &buf
		},
	}
	initRandPos atomic.Uint32
	initSync    sync.Once

	hashMD5Pool = sync.Pool{
		New: func() any {
			return &struct {
				buf  [md5.Size + 36]byte
				hash [md5.Size]byte
			}{}
		},
	}
	hashSHA1Pool = sync.Pool{
		New: func() any {
			return &struct {
				buf  [sha1.Size + 36]byte
				hash [sha1.Size]byte
			}{}
		},
	}

	v1 = &version{
		lastSequence: new(atomic.Uint32),
		lastTime:     new(atomic.Uint64),
	}
	v6 = &version{
		lastSequence: new(atomic.Uint32),
		lastTime:     new(atomic.Uint64),
	}
	v7 = &version{
		lastSequence: new(atomic.Uint32),
		lastTime:     new(atomic.Uint64),
	}
	v8 = &version{
		lastSequence: new(atomic.Uint32),
		lastTime:     new(atomic.Uint64),
	}
)

// Приватные функции
func encodeHex(buf []byte, uuid UUID) {
	buf[0] = hexTable[uuid[0]>>4]
	buf[1] = hexTable[uuid[0]&0x0f]
	buf[2] = hexTable[uuid[1]>>4]
	buf[3] = hexTable[uuid[1]&0x0f]
	buf[4] = hexTable[uuid[2]>>4]
	buf[5] = hexTable[uuid[2]&0x0f]
	buf[6] = hexTable[uuid[3]>>4]
	buf[7] = hexTable[uuid[3]&0x0f]
	buf[8] = '-'
	buf[9] = hexTable[uuid[4]>>4]
	buf[10] = hexTable[uuid[4]&0x0f]
	buf[11] = hexTable[uuid[5]>>4]
	buf[12] = hexTable[uuid[5]&0x0f]
	buf[13] = '-'
	buf[14] = hexTable[uuid[6]>>4]
	buf[15] = hexTable[uuid[6]&0x0f]
	buf[16] = hexTable[uuid[7]>>4]
	buf[17] = hexTable[uuid[7]&0x0f]
	buf[18] = '-'
	buf[19] = hexTable[uuid[8]>>4]
	buf[20] = hexTable[uuid[8]&0x0f]
	buf[21] = hexTable[uuid[9]>>4]
	buf[22] = hexTable[uuid[9]&0x0f]
	buf[23] = '-'
	buf[24] = hexTable[uuid[10]>>4]
	buf[25] = hexTable[uuid[10]&0x0f]
	buf[26] = hexTable[uuid[11]>>4]
	buf[27] = hexTable[uuid[11]&0x0f]
	buf[28] = hexTable[uuid[12]>>4]
	buf[29] = hexTable[uuid[12]&0x0f]
	buf[30] = hexTable[uuid[13]>>4]
	buf[31] = hexTable[uuid[13]&0x0f]
	buf[32] = hexTable[uuid[14]>>4]
	buf[33] = hexTable[uuid[14]&0x0f]
	buf[34] = hexTable[uuid[15]>>4]
	buf[35] = hexTable[uuid[15]&0x0f]
}
func genRandCrypto(b []byte) {
	if len(b) == 0 {
		return
	}
	if len(b) > randbufSize {
		if _, err := rand.Read(b); err != nil {
			clear(b)
		}
		return
	}
	buf := initRandPool.Get().(*[]byte)
	defer initRandPool.Put(buf)
	for {
		pos := initRandPos.Load()
		if pos+uint32(len(b)) > randbufSize {
			if _, err := rand.Read((*buf)); err != nil {
				clear(b)
			}
			initRandPos.Store(0)
			continue
		}
		if initRandPos.CompareAndSwap(pos, pos+uint32(len(b))) {
			copy(b, (*buf)[pos:pos+uint32(len(b))])
			return
		}
	}
}
func getTimeMilliAndSequence(v string) (ts uint64, sq uint32) {
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		return 0, 0
	}
	var (
		lastSequence *atomic.Uint32
		lastTime     *atomic.Uint64
		maxSequence  uint32
	)
	switch v {
	case "v7":
		lastTime, lastSequence, maxSequence = v7.lastTime, v7.lastSequence, maxV7Sequence
	case "v8":
		lastTime, lastSequence, maxSequence = v8.lastTime, v8.lastSequence, maxV8Sequence
	default:
		return 0, 0
	}
	ts = uint64(initClock.Now().UnixMilli())
	if prev := lastTime.Load(); ts != prev {
		if ts > prev && lastTime.CompareAndSwap(prev, ts) {
			lastSequence.Store(0)
		}
		return ts, 0
	} else {
		sq = lastSequence.Add(1) % (maxSequence + 1)
		return ts, sq
	}
}
func getTimeNanoAndSequence(v string) (ts uint64, sq uint32) {
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		return 0, 0
	}
	var (
		lastSequence *atomic.Uint32
		lastTime     *atomic.Uint64
		maxSequence  uint32
	)
	switch v {
	case "v1":
		lastTime, lastSequence, maxSequence = v1.lastTime, v1.lastSequence, maxV1Sequence
	case "v6":
		lastTime, lastSequence, maxSequence = v6.lastTime, v6.lastSequence, maxV6Sequence
	default:
		return 0, 0
	}
	ts = uint64(initClock.Now().UnixNano()/100) + offsetTime
	for {
		prevTime := lastTime.Load()
		if ts < prevTime {
			ts = uint64(initClock.Now().UnixNano()/100 + offsetTime)
			continue
		}
		if ts == prevTime {
			sq := lastSequence.Add(1)
			if sq > maxSequence {
				if _, isMock := initClock.(*MockClock); !isMock {
					waitTime(100 * time.Nanosecond)
					ts = uint64(initClock.Now().UnixNano()/100 + offsetTime)
					continue
				}
			}
			return ts, sq
		}
		if lastTime.CompareAndSwap(prevTime, ts) {
			lastSequence.Store(0)
			return ts, 0
		}
	}
}
func getPOSID(p int) (pi uint32) {
	initSync.Do(func() {
		initError = initGlobal()
	})
	if initError != nil {
		return 0
	}
	switch p {
	case 0:
		return cacheUID.Load()
	case 1:
		return cacheGID.Load()
	default:
		return 0
	}
}
func initGlobal() error {
	initMACAddress()
	initPOSID()
	initSequences()
	return nil
}
func initMACAddress() {
	var mac [6]byte
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if iface.Flags&net.FlagLoopback == 0 && len(iface.HardwareAddr) >= 6 {
				copy(mac[:], iface.HardwareAddr[:6])
				break
			}
		}
	}
	if mac == [6]byte{} {
		if _, err := rand.Read(mac[:]); err != nil {
			mac = [6]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x01}
		}
		mac[0] = (mac[0] & 0xFC) | bitLocal | bitMulticast
	}
	initMAC.Store(mac)
}
func initPOSID() {
	cacheGID.Store(uint32(os.Getgid()))
	cacheUID.Store(uint32(os.Getuid()))
}
func initSequences() {
	buf := make([]byte, 4*4)
	if _, err := rand.Read(buf); err != nil {
		clear(buf)
	}
	v1.lastSequence.Store(binary.BigEndian.Uint32(buf[0:4]) % (maxV1Sequence + 1))
	v6.lastSequence.Store(binary.BigEndian.Uint32(buf[4:8]) % (maxV6Sequence + 1))
	v7.lastSequence.Store(binary.BigEndian.Uint32(buf[8:12]) % (maxV7Sequence + 1))
	v8.lastSequence.Store(binary.BigEndian.Uint32(buf[12:16]) % (maxV8Sequence + 1))
}
func waitTime(d time.Duration) {
	start := time.Now()
	if d > time.Microsecond*10 {
		time.Sleep(d - time.Microsecond*10)
	}
	for time.Since(start) < d {
		runtime.Gosched()
	}
}
