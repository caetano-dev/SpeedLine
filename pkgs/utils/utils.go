package utils

type conversionfunc func(float64) float64

var ConversionsTable = map[string]map[string]conversionfunc{
	"hours": {
		"days":    func(value float64) float64 { return value / 24 },
		"minutes": func(value float64) float64 { return value * 60 },
		"seconds": func(value float64) float64 { return value * 3600 },
		"years":   func(value float64) float64 { return value / (24 * 365) },
	},
	"days": {
		"hours":   func(value float64) float64 { return value * 24 },
		"minutes": func(value float64) float64 { return value * 24 * 60 },
		"seconds": func(value float64) float64 { return value * 24 * 60 * 60 },
		"years":   func(value float64) float64 { return value / 365 },
	},
	"minutes": {
		"hours":   func(value float64) float64 { return value / 60 },
		"days":    func(value float64) float64 { return value / (24 * 60) },
		"seconds": func(value float64) float64 { return value * 60 },
		"years":   func(value float64) float64 { return value / (365 * 24 * 60) },
	},
	"seconds": {
		"hours":   func(value float64) float64 { return value / 3600 },
		"days":    func(value float64) float64 { return value / (24 * 60 * 60) },
		"minutes": func(value float64) float64 { return value / 60 },
		"years":   func(value float64) float64 { return value / (365 * 24 * 60 * 60) },
	},
	"years": {
		"hours":   func(value float64) float64 { return value * 24 * 365 },
		"days":    func(value float64) float64 { return value * 365 },
		"minutes": func(value float64) float64 { return value * 24 * 60 * 365 },
		"seconds": func(value float64) float64 { return value * 24 * 60 * 60 * 365 },
	},
	"km/h": {
		"m/s":  func(value float64) float64 { return value * 0.277778 },
		"mph":  func(value float64) float64 { return value * 0.621371 },
		"ft/s": func(value float64) float64 { return value * 0.911344 },
	},
	"m/s": {
		"km/h": func(value float64) float64 { return value * 3.6 },
		"mph":  func(value float64) float64 { return value * 2.23694 },
		"ft/s": func(value float64) float64 { return value * 3.28084 },
	},
	"mph": {
		"km/h": func(value float64) float64 { return value * 1.60934 },
		"m/s":  func(value float64) float64 { return value * 0.44704 },
		"ft/s": func(value float64) float64 { return value * 1.46667 },
	},
	"ft/s": {
		"km/h": func(value float64) float64 { return value * 1.09728 },
		"m/s":  func(value float64) float64 { return value * 0.3048 },
		"mph":  func(value float64) float64 { return value * 0.681818 },
	},
	"bytes": {
		"kb": func(value float64) float64 { return value / 1024 },
		"mb": func(value float64) float64 { return value / (1024 * 1024) },
		"gb": func(value float64) float64 { return value / (1024 * 1024 * 1024) },
		"tb": func(value float64) float64 { return value / (1024 * 1024 * 1024 * 1024) },
	},
	"kb": {
		"bytes": func(value float64) float64 { return value * 1024 },
		"mb":    func(value float64) float64 { return value / 1024 },
		"gb":    func(value float64) float64 { return value / (1024 * 1024) },
		"tb":    func(value float64) float64 { return value / (1024 * 1024 * 1024) },
	},
	"mb": {
		"bytes": func(value float64) float64 { return value * 1024 * 1024 },
		"kb":    func(value float64) float64 { return value * 1024 },
		"gb":    func(value float64) float64 { return value / 1024 },
		"tb":    func(value float64) float64 { return value / (1024 * 1024) },
	},
	"gb": {
		"bytes": func(value float64) float64 { return value * 1024 * 1024 * 1024 },
		"kb":    func(value float64) float64 { return value * 1024 * 1024 },
		"mb":    func(value float64) float64 { return value * 1024 },
		"tb":    func(value float64) float64 { return value / 1024 },
	},
	"tb": {
		"bytes": func(value float64) float64 { return value * 1024 * 1024 * 1024 * 1024 },
		"kb":    func(value float64) float64 { return value * 1024 * 1024 * 1024 },
		"mb":    func(value float64) float64 { return value * 1024 * 1024 },
		"gb":    func(value float64) float64 { return value * 1024 },
	},
}
