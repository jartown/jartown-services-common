package common

import "time"

type TimeSource interface {
	Now() time.Time
}

type RealTime struct{}

func (r *RealTime) Now() time.Time {
	return time.Now()
}

type MockTime struct {
	Time time.Time
}

func (m *MockTime) Now() time.Time {
	return m.Time
}

func (m *MockTime) Set(t time.Time) {
	m.Time = t
}

func (m *MockTime) NowNextWeek() time.Time {
	return m.Time.Add(7 * 24 * time.Hour)
}
