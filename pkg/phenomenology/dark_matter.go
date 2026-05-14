package phenomenology

import "math"

type DarkMatterResult struct {
	Executed                  bool
	Candidate                 string
	MassGeV                   float64
	TargetOmegaH2             float64
	RequiredYield             float64
	ThermalRelativisticYield  float64
	ThermalStableOmegaH2      float64
	OverclosureFactor         float64
	RequiredFractionOfThermal float64
	ReheatingTemperatureFree  bool
	ExactAbundancePredicted   bool
	Interpretation            string
	Statuses                  []string
}

// RequiredYieldForOmega returns Y=n/s required for a stable relic of mass m to
// explain Omega h^2.
func RequiredYieldForOmega(omegaH2, massGeV float64) float64 {
	return omegaH2 * CriticalDensityH2GeVPerCM3 / (massGeV * EntropyDensityTodayPerCM3)
}

func RelativisticFermionYield(gInternal, gStarS float64) float64 {
	// Y = n/s = [3/4*g*zeta(3)/pi^2*T^3] / [2*pi^2/45*g*S*T^3]
	// = 135*zeta(3)/(8*pi^4) * g/g*S
	return 135 * 1.202056903159594 / (8 * math.Pow(math.Pi, 4)) * gInternal / gStarS
}

func OmegaFromYield(massGeV, yield float64) float64 {
	return massGeV * yield * EntropyDensityTodayPerCM3 / CriticalDensityH2GeVPerCM3
}

func ComputeDarkMatterConstraint() DarkMatterResult {
	mass := AshaBGapMajoranaMassGeV
	required := RequiredYieldForOmega(TargetRelicDensity, mass)
	thermalYield := RelativisticFermionYield(2, 106.75)
	thermalOmega := OmegaFromYield(mass, thermalYield)
	overclosure := thermalOmega / TargetRelicDensity
	requiredFraction := required / thermalYield
	return DarkMatterResult{
		Executed:                  true,
		Candidate:                 "ASHA B-gap heavy Majorana sector",
		MassGeV:                   mass,
		TargetOmegaH2:             TargetRelicDensity,
		RequiredYield:             required,
		ThermalRelativisticYield:  thermalYield,
		ThermalStableOmegaH2:      thermalOmega,
		OverclosureFactor:         overclosure,
		RequiredFractionOfThermal: requiredFraction,
		ReheatingTemperatureFree:  true,
		ExactAbundancePredicted:   false,
		Interpretation:            "A stable thermal relic at the ASHA B-gap mass grossly overcloses the universe; matching Planck requires an extremely suppressed/nonthermal yield or entropy dilution. Reheating temperature, production channel and decay/stability theorem remain empirical seals.",
		Statuses: []string{
			StatusEmpiricalSealLoaded,
			StatusDarkYieldConstraintComputed,
			StatusThermalRelicOverclosureComputed,
			StatusDarkMatterNotPredicted,
			StatusNotNativePrediction,
		},
	}
}
