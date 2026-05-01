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
