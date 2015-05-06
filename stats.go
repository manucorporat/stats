package stats

import "sync"

type StatsType map[string]float64

type StatsCollector struct {
	lock  sync.RWMutex
	stats StatsType
}

var defaultCollector = New()

func New() *StatsCollector {
	s := new(StatsCollector)
	s.Reset()
	return s
}

func (s *StatsCollector) Reset() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stats = make(StatsType)
}

func (s *StatsCollector) Set(key string, value float64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stats[key] = value
}

func (s *StatsCollector) Add(key string, delta float64) float64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	value := s.stats[key]
	value += delta
	s.stats[key] = value
	return value
}

func (s *StatsCollector) Get(key string) float64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.stats[key]
}

func (s *StatsCollector) Data() StatsType {
	s.lock.Lock()
	defer s.lock.Unlock()

	cp := make(StatsType)
	for key, value := range s.stats {
		cp[key] = value
	}
	return cp
}

func Reset() {
	defaultCollector.Reset()
}

func Set(key string, value float64) {
	defaultCollector.Set(key, value)
}

func Add(key string, delta float64) float64 {
	return defaultCollector.Add(key, delta)
}

func Get(key string) float64 {
	return defaultCollector.Get(key)
}

func Data() StatsType {
	return defaultCollector.Data()
}
