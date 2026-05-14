package asha

import "math"

const (
	TargetOmegaDMH2              = 0.120
	EntropyDensityTodayPerCM3    = 2891.2
	CriticalDensityH2GeVPerCM3   = 1.05375e-5
	RelativisticGStarS           = 106.75
	Zeta3                        = 1.202056903159594
	DarkEnergyTargetPlanckUnits  = 1.0e-120
	Gate344DarkEnergyTargetUnits = 1.0e-122
	AshaThresholdDeltaLambda     = -0.097846792207
	StandardVEVGeV               = 246.22
	ObservedHiggsMassGeV         = 125.25
	TopMassGeV                   = 172.69
	StrongCouplingMZ             = 0.1179
	ZBosonMassGeV                = 91.1876
	G1AtMZ                       = 0.4614
	G2AtMZ                       = 0.6517
)

type DarkMatterConditional struct {
	MassGeV                   float64 `json:"mass_gev"`
	TargetOmegaH2             float64 `json:"target_omega_h2"`
	RequiredYield             float64 `json:"required_yield"`
	ThermalRelativisticYield  float64 `json:"thermal_relativistic_yield"`
	ThermalStableOmegaH2      float64 `json:"thermal_stable_omega_h2"`
	OverclosureFactor         float64 `json:"overclosure_factor"`
	RequiredFractionOfThermal float64 `json:"required_fraction_of_thermal"`
}

func ComputeDarkMatterConditional(massGeV float64) DarkMatterConditional {
	required := TargetOmegaDMH2 * CriticalDensityH2GeVPerCM3 / (massGeV * EntropyDensityTodayPerCM3)
	thermalYield := 135 * Zeta3 / (8 * math.Pow(math.Pi, 4)) * 2 / RelativisticGStarS
	thermalOmega := massGeV * thermalYield * EntropyDensityTodayPerCM3 / CriticalDensityH2GeVPerCM3
	return DarkMatterConditional{
		MassGeV:                   massGeV,
		TargetOmegaH2:             TargetOmegaDMH2,
		RequiredYield:             required,
		ThermalRelativisticYield:  thermalYield,
		ThermalStableOmegaH2:      thermalOmega,
		OverclosureFactor:         thermalOmega / TargetOmegaDMH2,
		RequiredFractionOfThermal: required / thermalYield,
	}
}

type CosmologyConditional struct {
	BareVacuumPlanckUnits          float64 `json:"bare_vacuum_planck_units"`
	TargetDarkEnergyPlanckUnits    float64 `json:"target_dark_energy_planck_units"`
	RequiredCounterterm            float64 `json:"required_counterterm"`
	CancellationRatio              float64 `json:"cancellation_ratio"`
	DigitsOfCancellation           float64 `json:"digits_of_cancellation"`
	HolographicLMpForTarget        float64 `json:"holographic_l_mp_for_target"`
	Gate344HolographicLMpForTarget float64 `json:"gate344_holographic_l_mp_for_target"`
	PfaffianHierarchy              float64 `json:"pfaffian_hierarchy"`
	ElectroweakVacuumPlanckFourth  float64 `json:"electroweak_vacuum_planck_fourth"`
	EWVacuumOverTarget             float64 `json:"ew_vacuum_over_target"`
	EWVacuumOverGate344Target      float64 `json:"ew_vacuum_over_gate344_target"`
}

func ComputeCosmologyConditional(planckGeV, vPfGeV float64) CosmologyConditional {
	bare := 48 / (math.Pi * math.Pi)
	target := DarkEnergyTargetPlanckUnits
	rho := vPfGeV / planckGeV
	ew4 := math.Pow(rho, 4)
	return CosmologyConditional{
		BareVacuumPlanckUnits:          bare,
		TargetDarkEnergyPlanckUnits:    target,
		RequiredCounterterm:            bare - target,
		CancellationRatio:              bare / target,
		DigitsOfCancellation:           math.Log10(bare / target),
		HolographicLMpForTarget:        1 / math.Sqrt(target),
		Gate344HolographicLMpForTarget: 1 / math.Sqrt(Gate344DarkEnergyTargetUnits),
		PfaffianHierarchy:              rho,
		ElectroweakVacuumPlanckFourth:  ew4,
		EWVacuumOverTarget:             ew4 / target,
		EWVacuumOverGate344Target:      ew4 / Gate344DarkEnergyTargetUnits,
	}
}

type VacuumFateConditional struct {
	SeedMode              string  `json:"seed_mode"`
	TopRunningFactor      float64 `json:"top_running_factor"`
	InitialLambda         float64 `json:"initial_lambda"`
	InitialYTop           float64 `json:"initial_y_top"`
	InitialG3             float64 `json:"initial_g3"`
	ThresholdScaleGeV     float64 `json:"threshold_scale_gev"`
	ThresholdDeltaLambda  float64 `json:"threshold_delta_lambda"`
	LambdaBeforeThreshold float64 `json:"lambda_before_threshold"`
	LambdaAfterThreshold  float64 `json:"lambda_after_threshold"`
	InstabilityScaleGeV   float64 `json:"instability_scale_gev"`
	LambdaMin             float64 `json:"lambda_min"`
	LambdaMinScaleGeV     float64 `json:"lambda_min_scale_gev"`
	BounceAction          float64 `json:"bounce_action"`
	Log10LifetimeYears    float64 `json:"log10_lifetime_years"`
	AgeUniverseLog10Years float64 `json:"age_universe_log10_years"`
	Metastable            bool    `json:"metastable"`
	StableToPlanck        bool    `json:"stable_to_planck"`
}

type rgState struct{ lambda, yt, g3 float64 }

func ComputeVacuumFateConditionals(planckGeV, thresholdGeV float64) []VacuumFateConditional {
	return []VacuumFateConditional{
		computeVacuumFateConditional("tree-pole-top-seed", 1, planckGeV, thresholdGeV),
		computeVacuumFateConditional("one-loop-QCD-MSbar-top-seed", 1-4*StrongCouplingMZ/(3*math.Pi), planckGeV, thresholdGeV),
	}
}

func betaRG(s rgState) rgState {
	const loop = 1.0 / (16.0 * math.Pi * math.Pi)
	g1sq, g2sq, g3sq := G1AtMZ*G1AtMZ, G2AtMZ*G2AtMZ, s.g3*s.g3
	yt2, lambda := s.yt*s.yt, s.lambda
	betaLambda := 24*lambda*lambda + 12*lambda*yt2 - 6*yt2*yt2 - 3*lambda*(3*g2sq+g1sq) + (3.0/8.0)*(2*g2sq*g2sq+math.Pow(g2sq+g1sq, 2))
	betaYt := s.yt * ((9.0/2.0)*yt2 - 8*g3sq - (9.0/4.0)*g2sq - (17.0/12.0)*g1sq)
	betaG3 := -7.0 * s.g3 * g3sq
	return rgState{lambda: loop * betaLambda, yt: loop * betaYt, g3: loop * betaG3}
}

func rk4Step(s rgState, h float64) rgState {
	k1 := betaRG(s)
	k2 := betaRG(rgState{s.lambda + 0.5*h*k1.lambda, s.yt + 0.5*h*k1.yt, s.g3 + 0.5*h*k1.g3})
	k3 := betaRG(rgState{s.lambda + 0.5*h*k2.lambda, s.yt + 0.5*h*k2.yt, s.g3 + 0.5*h*k2.g3})
	k4 := betaRG(rgState{s.lambda + h*k3.lambda, s.yt + h*k3.yt, s.g3 + h*k3.g3})
	return rgState{
		lambda: s.lambda + h*(k1.lambda+2*k2.lambda+2*k3.lambda+k4.lambda)/6,
		yt:     s.yt + h*(k1.yt+2*k2.yt+2*k3.yt+k4.yt)/6,
		g3:     s.g3 + h*(k1.g3+2*k2.g3+2*k3.g3+k4.g3)/6,
	}
}

func computeVacuumFateConditional(seedMode string, topRunningFactor, planckGeV, thresholdGeV float64) VacuumFateConditional {
	lambda0 := ObservedHiggsMassGeV * ObservedHiggsMassGeV / (2 * StandardVEVGeV * StandardVEVGeV)
	yt0 := math.Sqrt2 * TopMassGeV * topRunningFactor / StandardVEVGeV
	g30 := math.Sqrt(4 * math.Pi * StrongCouplingMZ)
	startT, endT, thresholdT := math.Log(ZBosonMassGeV), math.Log(planckGeV), math.Log(thresholdGeV)
	steps := 20000
	h := (endT - startT) / float64(steps)
	s := rgState{lambda: lambda0, yt: yt0, g3: g30}
	lambdaMin, lambdaMinT := s.lambda, startT
	instabilityScale := 0.0
	thresholdApplied := false
	lambdaBefore, lambdaAfter := math.NaN(), math.NaN()
	prevLambda, prevT := s.lambda, startT
	for i := 0; i < steps; i++ {
		t := startT + float64(i)*h
		nextT := t + h
		if !thresholdApplied && t < thresholdT && nextT >= thresholdT {
			partial := thresholdT - t
			if partial > 0 {
				s = rk4Step(s, partial)
			}
			lambdaBefore = s.lambda
			s.lambda += AshaThresholdDeltaLambda
			lambdaAfter = s.lambda
			thresholdApplied = true
			if prevLambda > 0 && s.lambda <= 0 && instabilityScale == 0 {
				instabilityScale = math.Exp(thresholdT)
			}
			remaining := nextT - thresholdT
			if remaining > 0 {
				s = rk4Step(s, remaining)
			}
		} else {
			s = rk4Step(s, h)
		}
		if s.lambda < lambdaMin {
			lambdaMin, lambdaMinT = s.lambda, nextT
		}
		if prevLambda > 0 && s.lambda <= 0 && instabilityScale == 0 {
			frac := prevLambda / (prevLambda - s.lambda)
			instabilityScale = math.Exp(prevT + frac*h)
		}
		prevLambda, prevT = s.lambda, nextT
	}
	stableToPlanck := instabilityScale == 0
	metastable := !stableToPlanck
	bounce := math.Inf(1)
	logLifetime := math.Inf(1)
	if lambdaMin < 0 {
		bounce = 8 * math.Pi * math.Pi / (3 * math.Abs(lambdaMin))
		logLifetime = math.Log10(6.582119569e-25/31557600.0/instabilityScale) + bounce/math.Ln10
	}
	return VacuumFateConditional{
		SeedMode: seedMode, TopRunningFactor: topRunningFactor, InitialLambda: lambda0, InitialYTop: yt0, InitialG3: g30,
		ThresholdScaleGeV: thresholdGeV, ThresholdDeltaLambda: AshaThresholdDeltaLambda, LambdaBeforeThreshold: lambdaBefore,
		LambdaAfterThreshold: lambdaAfter, InstabilityScaleGeV: instabilityScale, LambdaMin: lambdaMin, LambdaMinScaleGeV: math.Exp(lambdaMinT),
		BounceAction: bounce, Log10LifetimeYears: logLifetime, AgeUniverseLog10Years: math.Log10(13.8e9), Metastable: metastable, StableToPlanck: stableToPlanck,
	}
}
