// Copyright [2026] [Mikhail Dadaev]
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
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Публичные константы
const (
	Author  = "Mikhail Dadaev"
	Version = "1.26.10"
)

// Публичные переменные
var (
	ErrFailedCryptoRand    = errors.New("failed crypto/rand")
	ErrFailedGenMAC        = errors.New("failed gen/mac")
	ErrFailedGenSequences  = errors.New("failed gen/sequences")
	ErrInvalidNameLine     = errors.New("invalid nameline")
	ErrInvalidNameSpase    = errors.New("invalid namespace")
	ErrInvalidUUIDLength   = errors.New("invalid UUID length")
	ErrInvalidUUIDMAC      = errors.New("invalid UUID mac")
	ErrInvalidUUIDNode     = errors.New("invalid UUID node")
	ErrInvalidUUIDPOSIX    = errors.New("invalid UUID posix")
	ErrInvalidUUIDPosType  = errors.New("invalid UUID postype")
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
	bufSize          = 512
	hexTable         = "0123456789abcdef"
	maxV1Sequence    = 0x3FFF
	maxV2POSType     = 0xFF
	maxV6Sequence    = 0x3FFF
	maxV7Sequence    = 0xFFF
	maxV8NodeID      = 0x3FFF
	maxV8Sequence    = 0xFFF
	offsetTime       = 0x01B21DD213814000
	variantInvalid   = -1
	variantNCS       = 0
	variantRFC4122   = 1
	variantMicrosoft = 2
	variantReserved  = 3
)

// Приватные типы
type poolBuffer struct {
	buf []byte
	pos atomic.Uint32
}
type version struct {
	lastSequence *atomic.Uint32
	lastTime     *atomic.Uint64
}

// Приватные переменные
var (
	cacheGID      atomic.Uint32
	cacheUID      atomic.Uint32
	initCachePool = sync.Pool{
		New: func() any {
			return new([36]byte)
		},
	}
	initClock    clock = realClock{}
	initError    error
	initMAC      atomic.Value
	initRandPool = sync.Pool{
		New: func() any {
			buf := make([]byte, bufSize)
			_, _ = rand.Read(buf)
			return &poolBuffer{
				buf: buf,
				pos: atomic.Uint32{},
			}
		},
	}
	initSync    sync.Once
	hashMD5Pool = sync.Pool{
		New: func() any {
			return &struct {
				buf  [md5.Size + bufSize]byte
				hash [md5.Size]byte
			}{}
		},
	}
	hashSHA1Pool = sync.Pool{
		New: func() any {
			return &struct {
				buf  [sha1.Size + bufSize]byte
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
	if len(b) > bufSize {
		if _, err := rand.Read(b); err != nil {
			clear(b)
		}
		return
	}
	poolBuffer := initRandPool.Get().(*poolBuffer)
	defer initRandPool.Put(poolBuffer)
	for {
		pos := poolBuffer.pos.Load()
		if pos+uint32(len(b)) > bufSize {
			if _, err := rand.Read((poolBuffer.buf)); err != nil {
				clear(b)
			}
			poolBuffer.pos.Store(0)
			continue
		}
		if poolBuffer.pos.CompareAndSwap(pos, pos+uint32(len(b))) {
			copy(b, poolBuffer.buf[pos:pos+uint32(len(b))])
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
	ts = uint64(initClock.now().UnixMilli())
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
	ts = uint64(initClock.now().UnixNano()/100) + offsetTime
	for {
		prevTime := lastTime.Load()
		if ts < prevTime {
			ts = uint64(initClock.now().UnixNano()/100 + offsetTime)
			continue
		}
		if ts == prevTime {
			sq := lastSequence.Add(1)
			if sq > maxSequence {
				if _, isMock := initClock.(*mockClock); !isMock {
					waitTime(100 * time.Nanosecond)
					ts = uint64(initClock.now().UnixNano()/100 + offsetTime)
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
func getPOSIX(p int) (pi uint32) {
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
	initPOSIX()
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
func initPOSIX() {
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
func transformVariant(variant int) string {
	switch variant {
	case variantNCS:
		return "NCS"
	case variantRFC4122:
		return "RFC4122"
	case variantMicrosoft:
		return "Microsoft"
	case variantReserved:
		return "Reserved"
	default:
		return "Invalid"
	}
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
