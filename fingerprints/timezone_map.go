package fingerprint

import (
	"strconv"
	"strings"
	"time"
)

func MapTimezoneToLanguage(timezone string) string {
	timezoneLanguageMap := map[string]string{
		"Pacific/Midway":                 "en-US",
		"Pacific/Honolulu":               "en-US",
		"America/Anchorage":              "en-US",
		"America/Los_Angeles":            "en-US",
		"America/Vancouver":              "en-US",
		"America/Tijuana":                "en-US",
		"America/Phoenix":                "en-US",
		"America/Chihuahua":              "es-MX",
		"America/Denver":                 "en-US",
		"America/Edmonton":               "en-US",
		"America/Mazatlan":               "es-MX",
		"America/Regina":                 "en-US",
		"America/Costa_Rica":             "es",
		"America/Chicago":                "en-US",
		"America/Mexico_City":            "es-MX",
		"America/Tegucigalpa":            "es",
		"America/Winnipeg":               "en-US",
		"Pacific/Easter":                 "es-CL",
		"America/Bogota":                 "es",
		"America/Lima":                   "es",
		"America/New_York":               "en-US",
		"America/Toronto":                "en-US",
		"America/Caracas":                "es",
		"America/Barbados":               "en",
		"America/Halifax":                "en-US",
		"America/Manaus":                 "pt-BR",
		"America/Santiago":               "es-CL",
		"America/St_Johns":               "en-US",
		"America/Araguaina":              "pt-BR",
		"America/Argentina/Buenos_Aires": "es-AR",
		"America/Argentina/San_Luis":     "es-AR",
		"America/Montevideo":             "es-UY",
		"America/Sao_Paulo":              "pt-BR",
		"America/Godthab":                "kl",
		"Atlantic/South_Georgia":         "en-GB",
		"Atlantic/Cape_Verde":            "pt",
		"Atlantic/Azores":                "pt-PT",
		"Atlantic/Reykjavik":             "is",
		"Atlantic/St_Helena":             "en",
		"Africa/Casablanca":              "ar",
		"Atlantic/Faroe":                 "fo",
		"Europe/Dublin":                  "en-IE",
		"Europe/Lisbon":                  "pt-PT",
		"Europe/London":                  "en-GB",
		"Europe/Amsterdam":               "nl",
		"Europe/Belgrade":                "sr",
		"Europe/Berlin":                  "de",
		"Europe/Bratislava":              "sk",
		"Europe/Brussels":                "nl",
		"Europe/Budapest":                "hu",
		"Europe/Copenhagen":              "da",
		"Europe/Ljubljana":               "sl",
		"Europe/Madrid":                  "es-ES",
		"Europe/Malta":                   "mt",
		"Europe/Oslo":                    "nb",
		"Europe/Paris":                   "fr",
		"Europe/Prague":                  "cs",
		"Europe/Rome":                    "it",
		"Europe/Stockholm":               "sv",
		"Europe/Sarajevo":                "hr",
		"Europe/Tirane":                  "sq",
		"Europe/Vaduz":                   "de",
		"Europe/Vienna":                  "de",
		"Europe/Warsaw":                  "pl",
		"Europe/Zagreb":                  "hr",
		"Europe/Zurich":                  "de",
		"Africa/Windhoek":                "en",
		"Africa/Lagos":                   "en",
		"Africa/Brazzaville":             "fr",
		"Africa/Cairo":                   "ar",
		"Africa/Harare":                  "en",
		"Africa/Maputo":                  "pt",
		"Africa/Johannesburg":            "af",
		"Europe/Kaliningrad":             "ru",
		"Europe/Athens":                  "el",
		"Europe/Bucharest":               "ro",
		"Europe/Chisinau":                "ro",
		"Europe/Helsinki":                "fi",
		"Europe/Istanbul":                "tr",
		"Europe/Kiev":                    "uk",
		"Europe/Riga":                    "lv",
		"Europe/Sofia":                   "bg",
		"Europe/Tallinn":                 "et",
		"Europe/Vilnius":                 "lt",
		"Asia/Amman":                     "ar",
		"Asia/Beirut":                    "ar",
		"Asia/Jerusalem":                 "he",
		"Africa/Nairobi":                 "sw",
		"Asia/Baghdad":                   "ar",
		"Asia/Riyadh":                    "ar",
		"Asia/Kuwait":                    "ar",
		"Europe/Minsk":                   "be",
		"Europe/Moscow":                  "ru",
		"Asia/Tehran":                    "fa",
		"Europe/Samara":                  "ru",
		"Asia/Dubai":                     "ar",
		"Asia/Tbilisi":                   "ka",
		"Indian/Mauritius":               "fr",
		"Asia/Baku":                      "az",
		"Asia/Yerevan":                   "hy",
		"Asia/Kabul":                     "fa",
		"Asia/Karachi":                   "ur",
		"Asia/Aqtobe":                    "kk",
		"Asia/Ashgabat":                  "tk",
		"Asia/Oral":                      "kk",
		"Asia/Yekaterinburg":             "ru",
		"Asia/Calcutta":                  "en-IN",
		"Asia/Colombo":                   "si",
		"Asia/Katmandu":                  "ne",
		"Asia/Omsk":                      "ru",
		"Asia/Almaty":                    "kk",
		"Asia/Dhaka":                     "bn",
		"Asia/Novosibirsk":               "ru",
		"Asia/Rangoon":                   "my",
		"Asia/Bangkok":                   "th",
		"Asia/Jakarta":                   "id",
		"Asia/Krasnoyarsk":               "ru",
		"Asia/Novokuznetsk":              "ru",
		"Asia/Ho_Chi_Minh":               "vi",
		"Asia/Phnom_Penh":                "km",
		"Asia/Vientiane":                 "lo",
		"Asia/Shanghai":                  "zh-CN",
		"Asia/Hong_Kong":                 "zh-HK",
		"Asia/Kuala_Lumpur":              "ms",
		"Asia/Singapore":                 "zh-SG",
		"Asia/Manila":                    "fil",
		"Asia/Taipei":                    "zh-TW",
		"Asia/Ulaanbaatar":               "mn",
		"Asia/Makassar":                  "id",
		"Asia/Irkutsk":                   "ru",
		"Asia/Yakutsk":                   "ru",
		"Australia/Perth":                "en",
		"Australia/Eucla":                "en",
		"Asia/Seoul":                     "ko",
		"Asia/Tokyo":                     "ja",
		"Asia/Jayapura":                  "id",
		"Asia/Sakhalin":                  "ru",
		"Asia/Vladivostok":               "ru",
		"Asia/Magadan":                   "ru",
		"Australia/Darwin":               "en",
		"Australia/Adelaide":             "en",
		"Pacific/Guam":                   "en",
		"Australia/Brisbane":             "en",
		"Australia/Hobart":               "en",
		"Australia/Sydney":               "en",
		"Asia/Anadyr":                    "ru",
		"Pacific/Port_Moresby":           "en",
		"Asia/Kamchatka":                 "ru",
		"Pacific/Fiji":                   "fj",
		"Pacific/Majuro":                 "mh",
		"Pacific/Auckland":               "en",
		"Pacific/Tongatapu":              "to",
		"Pacific/Apia":                   "sm",
		"Pacific/Kiritimati":             "en",
	}

	language, found := timezoneLanguageMap[timezone]
	if !found {
		return "未知语言"
	}
	return language
}

func GetTimeZoneOffset(timezone string) (int, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	_, offset := time.Now().In(loc).Zone()
	return offset / 3600, nil
}

func getTimeZoneOffsetInHours(timezone string) float64 {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0
	}

	offset := time.Now().In(loc).Format("-07:00")
	parts := strings.Split(offset, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	return float64(hour) + float64(minute)/60
}
