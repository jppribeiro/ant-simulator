package config

import "math"

type windowConfig struct {
	X int32
	Y int32
}

var WINDOW = windowConfig{800, 600}

type antConfig struct {
	MIN_V      int32
	MAX_V      int32
	MARKER_TTL float64
	VIEW_D     float64
	VIEW_ANGLE float64 // For each section. Multiply by 2 for total FOV
}

var ANT_CONFIG = antConfig{2, 4, 500, 20, math.Pi / 2}

type worldConfig struct {
	HOME_POS_X float64
	HOME_POS_Y float64
	HOME_SIZE  float64

	FOOD_POS_X float64
	FOOD_POS_Y float64
	FOOD_SIZE  float64

	ANT_NUM          int32
	ANT_SOURCE_DECAY float64

	MARKER_TTL float64
}

var WORLD_CONFIG = worldConfig{
	HOME_POS_X: 200,
	HOME_POS_Y: 200,
	HOME_SIZE:  40,

	FOOD_POS_X: 600,
	FOOD_POS_Y: 400,
	FOOD_SIZE:  60,

	ANT_NUM:          150,
	ANT_SOURCE_DECAY: 600,

	MARKER_TTL: 300,
}
