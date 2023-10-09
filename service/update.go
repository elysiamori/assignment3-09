package service

func UpdateStatusWater(v int) string {
	switch {
	case v < 5:
		return "aman"
	case v >= 6 && v <= 8:
		return "siaga"
	default:
		return "bahaya"
	}
}

func UpdateStatusWind(v int) string {
	switch {
	case v < 6:
		return "aman"
	case v >= 7 && v <= 15:
		return "siaga"
	default:
		return "bahaya"
	}
}
