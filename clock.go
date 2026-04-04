package uuid

import "time"

// Приватные интерфейсы
type clock interface {
	now() time.Time
}

// Приватные структуры
type mockClock struct {
	time time.Time
}
type realClock struct{}

// Приватные методы
func (mockClock *mockClock) add(d time.Duration) {
	mockClock.time = mockClock.time.Add(d)
}
func (mockClock *mockClock) now() time.Time {
	return mockClock.time
}
func (realClock) now() time.Time {
	return time.Now().UTC()
}
