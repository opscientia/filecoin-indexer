package score

import "math"

// Variables are input parameters for the score calculation
type Variables struct {
	SlashedDealsCount uint32
	FaultsCount       uint32
	RelativePower     float32
	SectorSize        uint64
}

const (
	_slashingsWeight  = 100
	_faultsWeight     = 100
	_powerWeight      = 100
	_sectorSizeWeight = 10
)

const (
	_relativePowerBaseline = 0.1
	_sectorSizeBaseline    = 32 * 1024 * 1024 * 1024 // 32 GiB
)

// CalculateScore computes a miner score based on a set of variables
func CalculateScore(vars Variables) uint32 {
	slashingsScore := 1 / (1 + math.Pow(float64(vars.SlashedDealsCount), 2))
	faultsScore := 1 / (1 + float64(vars.FaultsCount))
	powerScore := vars.RelativePower / _relativePowerBaseline
	sectorSizeScore := vars.SectorSize / _sectorSizeBaseline

	return uint32(slashingsScore*_slashingsWeight) +
		uint32(faultsScore*_faultsWeight) +
		uint32(powerScore*_powerWeight) +
		uint32(sectorSizeScore*_sectorSizeWeight)
}
