package output

import (
	"github.com/fatih/color"
)

type BarometerWeather uint

const (
	BARO_WEATHER_UNKNOWN = iota
	BARO_WEATHER_SUNNY
	BARO_WEATHER_CLOUDY
	BARO_WEATHER_OVERCAST
	BARO_WEATHER_RAINY
	BARO_WEATHER_THUNDER_STORM
)

var (
	BARO_WEATHER_SUNNY_STYLE         = color.New(color.Bold, color.FgGreen)
	BARO_WEATHER_CLOUDY_STYLE        = color.New(color.FgBlue)
	BARO_WEATHER_OVERCAST_STYLE      = color.New(color.FgHiYellow)
	BARO_WEATHER_RAINY_STYLE         = color.New(color.FgYellow)
	BARO_WEATHER_THUNDER_STORM_STYLE = color.New(color.FgRed, color.Bold)

	BAROMETER_TITLE = color.New(color.Bold).Sprint(("Build Barometer:"))
)

func (b BarometerWeather) String() string {
	switch b {
	case BARO_WEATHER_SUNNY:
		return "🌞"
	case BARO_WEATHER_CLOUDY:
		return "⛅"
	case BARO_WEATHER_THUNDER_STORM:
		return "🌩"
	case BARO_WEATHER_OVERCAST:
		return "☁"
	case BARO_WEATHER_RAINY:
		return "🌧"
	default:
		return "[?]"
	}
}

func (b BarometerWeather) GetTextStyle() *color.Color {
	var style *color.Color
	switch b {
	case BARO_WEATHER_SUNNY:
		style = BARO_WEATHER_SUNNY_STYLE
	case BARO_WEATHER_CLOUDY:
		style = BARO_WEATHER_CLOUDY_STYLE
	case BARO_WEATHER_THUNDER_STORM:
		style = BARO_WEATHER_THUNDER_STORM_STYLE
	case BARO_WEATHER_OVERCAST:
		style = BARO_WEATHER_OVERCAST_STYLE
	case BARO_WEATHER_RAINY:
		style = BARO_WEATHER_RAINY_STYLE
	default:
		style = color.New()
	}
	return style
}

func (b BarometerWeather) Print(name string, successTime int, failureTime int) string {
	style := b.GetTextStyle()

	text := style.Sprintf("%s    (%d/%d)", name, successTime, successTime+failureTime)
	return b.String() + "  " + text
}

func BaroByTimes(successTime int, failureTime int) BarometerWeather {
	if failureTime == 0 {
		return BARO_WEATHER_SUNNY
	} else if successTime == 0 {
		return BARO_WEATHER_THUNDER_STORM
	} else if successTime == failureTime {
		return BARO_WEATHER_OVERCAST
	} else if successTime > failureTime {
		return BARO_WEATHER_CLOUDY
	} else if successTime < failureTime {
		return BARO_WEATHER_RAINY
	} else {
		return BARO_WEATHER_UNKNOWN
	}
}

func BaroPrintByTimes(name string, successTime int, failureTime int) string {
	return BaroByTimes(successTime, failureTime).Print(name, successTime, failureTime)
}
