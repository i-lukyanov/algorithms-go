package common

import (
	"fmt"
	"strconv"
)

func StringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func FloatToString(num float64, point int) string {
	switch point {
	case 2:
		return fmt.Sprintf("%.2f", num)
	case 6:
		return fmt.Sprintf("%.6f", num)
	case 7:
		return fmt.Sprintf("%.7f", num)
	case 10:
		return fmt.Sprintf("%.10f", num)
	}

	return fmt.Sprintf("%.2f", num)
}
