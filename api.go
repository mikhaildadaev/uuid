// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
)

// Публичные функции
func V1() UUID {
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
func V2(postype int) UUID {
	var uuid UUID
	// Проверка POSType (UID/GID и др.)
	postype = max(0, min(postype, maxV2POSType))
	// Извлечение POS-идентификатора
	pi := getPOSID(postype)
	// Получение MAC-адреса
	mac := initMAC.Load().([6]byte)
	// Установка рандомных данных
	genRandCrypto(uuid[4:9])
	// Установка POSID
	binary.BigEndian.PutUint32(uuid[0:4], pi)
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV2
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	// Установка POSTypeID
	uuid[9] = uint8(postype & 0x03)
	// Установка MAC-адреса
	copy(uuid[10:16], mac[:])
	return uuid
}
func V3(namespace UUID, nameline string) UUID {
	var uuid UUID
	var hashBufPool = hashMD5Pool.Get().(*struct {
		buf  [md5.Size + 36]byte
		hash [md5.Size]byte
	})
	defer hashMD5Pool.Put(hashBufPool)
	// Формирование данных и хеша
	if len(nameline) == 0 || len(nameline) > 36 {
		nameline = "default_nameline"
	}
	buf := hashBufPool.buf[:0]
	buf = append(buf, namespace[:]...)
	buf = append(buf, nameline...)
	hash := md5.Sum(buf)
	// Установка хеша
	copy(uuid[:], hash[:])
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV3
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	return uuid
}
func V4() UUID {
	var uuid UUID
	// Установка рандомных данных
	genRandCrypto(uuid[:])
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV4
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	return uuid
}
func V5(namespace UUID, nameline string) UUID {
	var uuid UUID
	var hashBufPool = hashSHA1Pool.Get().(*struct {
		buf  [sha1.Size + 36]byte
		hash [sha1.Size]byte
	})
	defer hashSHA1Pool.Put(hashBufPool)
	// Формирование данных и хеша
	if len(nameline) == 0 || len(nameline) > 36 {
		nameline = "default_nameline"
	}
	buf := hashBufPool.buf[:0]
	buf = append(buf, namespace[:]...)
	buf = append(buf, nameline...)
	hash := sha1.Sum(buf)
	// Установка хеша
	copy(uuid[:], hash[:])
	// Установка версии UUID
	uuid[6] = (uuid[6] & 0x0F) | bitV5
	// Установка варианта RFC 4122
	uuid[8] = (uuid[8] & 0x3F) | bitVRFC4122
	return uuid
}
func V6() UUID {
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
func V7() UUID {
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
func V8(nodeid int) UUID {
	var uuid UUID
	// Получение временной метки и последовательности
	ts, sq := getTimeMilliAndSequence("v8")
	// Получение идентификатора ноды
	nodeid = max(0, min(nodeid, maxV8NodeID))
	// Установка рандомных данных
	genRandCrypto(uuid[10:])
	// Установка временной метки
	binary.BigEndian.PutUint64(uuid[0:], ts<<16)
	// Установка последовательности и версии UUID
	binary.BigEndian.PutUint16(uuid[6:8], uint16(sq)&maxV8Sequence|bitV8<<8)
	// Установка идентификатора ноды и варианта RFC 4122
	binary.BigEndian.PutUint16(uuid[8:10], uint16(nodeid)&maxV8NodeID|bitVRFC4122<<8)
	return uuid
}
