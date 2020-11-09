package score

// Variables are input parameters for the score calculation
type Variables struct {
	SlashedDealsCount uint32
	FaultsCount       uint32
	RelativePower     float32
	SectorSize        uint64
}

const (
	slashingsWeight  = 100
	faultsWeight     = 100
	powerWeight      = 100
	sectorSizeWeight = 10
)

const (
	relativePowerBaseline = 0.1
	sectorSizeBaseline    = 32 * 1024 * 1024 * 1024 // 32 GiB
)

// CalculateScore computes a miner score based on a set of variables
func CalculateScore(vars Variables) uint32 {
	slashingsScore := 1 / (1 + float32(vars.SlashedDealsCount))
	faultsScore := 1 / (1 + float32(vars.FaultsCount))
	powerScore := vars.RelativePower / relativePowerBaseline
	sectorSizeScore := vars.SectorSize / sectorSizeBaseline

	return uint32(slashingsScore*slashingsWeight) +
		uint32(faultsScore*faultsWeight) +
		uint32(powerScore*powerWeight) +
		uint32(sectorSizeScore*sectorSizeWeight)
}
