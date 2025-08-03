package calculator

type PriceConfig struct {
	BasePrice    float64
	FixedCost    float64
	PricePerCm   float64
	PricePerHour float64
}

func EstimateByDetails(cfg PriceConfig, styleMult, bodyMult, colorMult, sizeCM float64) float64 {
	base := cfg.BasePrice + cfg.FixedCost + cfg.PricePerCm*sizeCM
	return base * styleMult * bodyMult * colorMult
}

func EstimateByHour(cfg PriceConfig, hours float64) float64 {
	return cfg.PricePerHour * hours
}
