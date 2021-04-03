package config

type windowConfig struct {
	X int32
	Y int32
}

var WINDOW = windowConfig{800, 600}

type antConfig struct {
	MIN_V        int32
	MAX_V        int32
	FORAGING_TTL float32
}

var ANT_CONFIG = antConfig{2, 4, 300}
