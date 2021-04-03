package config

type windowConfig struct {
	X int32
	Y int32
}

var WINDOW = windowConfig{800, 600}

type antConfig struct {
	MIN_V      int32
	MAX_V      int32
	MARKER_TTL float32
}

var ANT_CONFIG = antConfig{2, 4, 300}

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
