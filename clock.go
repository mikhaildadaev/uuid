// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import "time"

// Публичные интерфейсы
type Clock interface {
	Now() time.Time
}

// Публичные структуры
type MockClock struct {
	time time.Time
}
type RealClock struct{}

// Публичные методы
func (mockClock *MockClock) Add(d time.Duration) {
	mockClock.time = mockClock.time.Add(d)
}
func (mockClock *MockClock) Now() time.Time {
	return mockClock.time
}
func (realClock *RealClock) Now() time.Time {
	return time.Now().UTC()
}
