package config

import "sync"

// package config

var (
	MLock sync.Mutex
	Conf  Config
)
