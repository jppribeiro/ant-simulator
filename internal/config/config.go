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
	MARKER_TTL float32
	VIEW_D     float32
	VIEW_ANGLE float64 // For each section. Multiply by 2 for total FOV
}

var ANT_CONFIG = antConfig{2, 4, 500, 20, math.Pi / 3}

type worldConfig struct {
	HOME_POS_X float32
	HOME_POS_Y float32
	HOME_SIZE  float32

	FOOD_POS_X float32
	FOOD_POS_Y float32
	FOOD_SIZE  float32

	ANT_NUM int32
}

var WORLD_CONFIG = worldConfig{
	HOME_POS_X: 200,
	HOME_POS_Y: 200,
	HOME_SIZE:  40,
	FOOD_POS_X: 600,
	FOOD_POS_Y: 400,
	FOOD_SIZE:  60,
	ANT_NUM:    60,
}
