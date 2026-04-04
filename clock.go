// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
package uuid

import "time"

// Публичные методы
func (mockClock *mockClock) Add(d time.Duration) {
	mockClock.time = mockClock.time.Add(d)
}
func (mockClock *mockClock) Now() time.Time {
	return mockClock.time
}
func (realClock) Now() time.Time {
	return time.Now().UTC()
}

// Приватные интерфейсы
type clock interface {
	Now() time.Time
}

// Приватные структуры
type mockClock struct {
	time time.Time
}
type realClock struct{}
