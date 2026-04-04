// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import (
	"strconv"
	"testing"
	"time"
)

// Benchmarks
func Benchmark_Parse(b *testing.B) {
	UUID := testUUIDVUString
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Parse(UUID)
	}
}
func Benchmark_String(b *testing.B) {
	uuid, _ := Parse(testUUIDVUString)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.String()
	}
}
func Benchmark_Time(b *testing.B) {
	for i := 0; i < b.N; i++ {
		waitTime(time.Nanosecond * 100)
	}
}
func Benchmark_UUIDs(b *testing.B) {
	b.Run("v1 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V1()
		}
	})
	b.Run("v2 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V2(testPOSType)
		}
	})
	b.Run("v3 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V3(NameSpaceDNS, testNameString)
		}
	})
	b.Run("v4 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V4()
		}
	})
	b.Run("v5 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V5(NameSpaceDNS, testNameString)
		}
	})
	b.Run("v6 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V6()
		}
	})
	b.Run("v7 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V7()
		}
	})
	b.Run("v8 UUID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = V8(testNodeID)
		}
	})
}
func Benchmark_Validate(b *testing.B) {
	uuid, _ := Parse(testUUIDVUString)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.Validate()
	}
}
func Benchmark_V1_Info(b *testing.B) {
	b.Run("With Mock Time", func(b *testing.B) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			b.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v1.lastTime.Load()
		realSequence := v1.lastSequence.Load()
		b.Cleanup(func() {
			initClock = realClock
			v1.lastTime.Store(realTime)
			v1.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v1.lastTime.Store(uint64((initClock.Now().UnixNano() / 100) + offsetTime))
		v1.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V1()
			b.Logf("UUIDv1-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V1()
			b.Logf("UUIDv1-%d: %s", i, ui)
		}
	})
}
func Benchmark_V1_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V1()
		}
	})
}
func Benchmark_V1_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V1()
	}
}
func Benchmark_V2_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V2(testPOSType)
			b.Logf("UUIDv2-%d: %s", i, ui)
		}
	})
}
func Benchmark_V2_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V2(testPOSType)
		}
	})
}
func Benchmark_V2_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V2(testPOSType)
	}
}
func Benchmark_V3_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V3(NameSpaceDNS, testNameString+strconv.Itoa(i))
			b.Logf("UUIDv3-%d: %s", i, ui)
		}
	})
}
func Benchmark_V3_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V3(NameSpaceDNS, testNameString)
		}
	})
}
func Benchmark_V3_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V3(NameSpaceDNS, testNameString+strconv.Itoa(i))
	}
}
func Benchmark_V4_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V4()
			b.Logf("UUIDv4-%d: %s", i, ui)
		}
	})
}
func Benchmark_V4_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V4()
		}
	})
}
func Benchmark_V4_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V4()
	}
}
func Benchmark_V5_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V5(NameSpaceDNS, testNameString+strconv.Itoa(i))
			b.Logf("UUIDv5-%d: %s", i, ui)
		}
	})
}
func Benchmark_V5_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V5(NameSpaceDNS, testNameString)
		}
	})
}
func Benchmark_V5_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V5(NameSpaceDNS, testNameString+strconv.Itoa(i))
	}
}
func Benchmark_V6_Info(b *testing.B) {
	b.Run("With Mock Time", func(b *testing.B) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			b.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v6.lastTime.Load()
		realSequence := v6.lastSequence.Load()
		b.Cleanup(func() {
			initClock = realClock
			v6.lastTime.Store(realTime)
			v6.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v6.lastTime.Store(uint64((initClock.Now().UnixNano() / 100) + offsetTime))
		v6.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V6()
			b.Logf("UUIDv6-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V6()
			b.Logf("UUIDv6-%d: %s", i, ui)
		}
	})
}
func Benchmark_V6_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V6()
		}
	})
}
func Benchmark_V6_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V6()
	}
}
func Benchmark_V7_Info(b *testing.B) {
	b.Run("With Mock Time", func(b *testing.B) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			b.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v7.lastTime.Load()
		realSequence := v7.lastSequence.Load()
		b.Cleanup(func() {
			initClock = realClock
			v7.lastTime.Store(realTime)
			v7.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v7.lastTime.Store(uint64(initClock.Now().UnixMicro()))
		v7.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V7()
			b.Logf("UUIDv7-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V7()
			b.Logf("UUIDv7-%d: %s", i, ui)
		}
	})
}
func Benchmark_V7_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V7()
		}
	})
}
func Benchmark_V7_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V7()
	}
}
func Benchmark_V8_Info(b *testing.B) {
	b.Run("With Mock Time", func(b *testing.B) {
		// Ленивая инициализация глобального состояния
		initSync.Do(func() {
			initError = initGlobal()
		})
		if initError != nil {
			b.Fatal(initError)
		}
		// Фиксация текущего состояние
		realClock := initClock
		realTime := v8.lastTime.Load()
		realSequence := v8.lastSequence.Load()
		b.Cleanup(func() {
			initClock = realClock
			v8.lastTime.Store(realTime)
			v8.lastSequence.Store(realSequence)
		})
		// Мокирование
		mockClock := &mockClock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = mockClock
		v8.lastTime.Store(uint64(initClock.Now().UnixMicro()))
		v8.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V8(testNodeID)
			b.Logf("UUIDv8-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := V8(testNodeID)
			b.Logf("UUIDv8-%d: %s", i, ui)
		}
	})
}
func Benchmark_V8_Multi(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = V8(testNodeID)
		}
	})
}
func Benchmark_v8_Single(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = V8(testNodeID)
	}
}
