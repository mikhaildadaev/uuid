// Copyright (c) 2026 Mikhail Dadaev
// All rights reserved.
//
// This source code is licensed under the MIT License found in the
// LICENSE file in the root directory of this source tree.
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
