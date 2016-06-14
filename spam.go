package main

// Spam filter

type SpamFilter struct {
	storage map[string]int
	max     int
}

func CreateSpamFilter(max int) *SpamFilter {
	return &SpamFilter{
		make(map[string]int),
		max,
	}
}

func (s *SpamFilter) OK(ip string) bool {
	if s.storage[ip] > s.max {
		return false
	}
	s.storage[ip]++
	return true
}
