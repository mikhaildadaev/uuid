package uuid

import (
	"strconv"
	"testing"
	"time"
)

// Бенчмарки компонентов
func Benchmark_NewV1(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV1()
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV1()
		}
	})
}
func Benchmark_NewV1_Info(b *testing.B) {
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
		clockMock := &clockMock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = clockMock
		v1.lastTime.Store(uint64((initClock.now().UnixNano() / 100) + offsetTime))
		v1.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV1()
			b.Logf("UUIDv1-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV1()
			b.Logf("UUIDv1-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV2(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV2(testPOSType, testPOSValue)
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV2(testPOSType, testPOSValue)
		}
	})
}
func Benchmark_NewV2_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV2(testPOSType, testPOSValue)
			b.Logf("UUIDv2-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV3(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV3(NameSpaceDNS, testNameString)
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV3(NameSpaceDNS, testNameString+strconv.Itoa(i))
		}
	})
}
func Benchmark_NewV3_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV3(NameSpaceDNS, testNameString+strconv.Itoa(i))
			b.Logf("UUIDv3-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV4(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV4()
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV4()
		}
	})
}
func Benchmark_NewV4_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV4()
			b.Logf("UUIDv4-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV5(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV5(NameSpaceDNS, testNameString)
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV5(NameSpaceDNS, testNameString+strconv.Itoa(i))
		}
	})
}
func Benchmark_NewV5_Info(b *testing.B) {
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV5(NameSpaceDNS, testNameString+strconv.Itoa(i))
			b.Logf("UUIDv5-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV6(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV6()
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV6()
		}
	})
}
func Benchmark_NewV6_Info(b *testing.B) {
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
		clockMock := &clockMock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = clockMock
		v6.lastTime.Store(uint64((initClock.now().UnixNano() / 100) + offsetTime))
		v6.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV6()
			b.Logf("UUIDv6-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV6()
			b.Logf("UUIDv6-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV7(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV7()
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV7()
		}
	})
}
func Benchmark_NewV7_Info(b *testing.B) {
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
		clockMock := &clockMock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = clockMock
		v7.lastTime.Store(uint64(initClock.now().UnixMicro()))
		v7.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV7()
			b.Logf("UUIDv7-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV7()
			b.Logf("UUIDv7-%d: %s", i, ui)
		}
	})
}
func Benchmark_NewV8(b *testing.B) {
	b.Run("Multi", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = NewV8(testNodeID)
			}
		})
	})
	b.Run("Single", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewV8(testNodeID)
		}
	})
}
func Benchmark_NewV8_Info(b *testing.B) {
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
		clockMock := &clockMock{time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
		initClock = clockMock
		v8.lastTime.Store(uint64(initClock.now().UnixMicro()))
		v8.lastSequence.Store(0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV8(testNodeID)
			b.Logf("UUIDv8-%d: %s", i, ui)
		}
	})
	b.Run("Without Mock Time", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ui := NewV8(testNodeID)
			b.Logf("UUIDv8-%d: %s", i, ui)
		}
	})
}
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
func Benchmark_Validate(b *testing.B) {
	uuid, _ := Parse(testUUIDVUString)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.Validate()
	}
}
