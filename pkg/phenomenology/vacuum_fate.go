package phenomenology

import "math"

type RGPoint struct {
	MuGeV  float64
	T      float64
	Lambda float64
	YTop   float64
	G3     float64
}

type VacuumFateResult struct {
	Executed              bool
	Method                string
	SeedMode              string
	TopRunningFactor      float64
	InitialLambda         float64
	InitialYTop           float64
	InitialG3             float64
	ThresholdScaleGeV     float64
	ThresholdDeltaLambda  float64
	ThresholdApplied      bool
	LambdaBeforeThreshold float64
	LambdaAfterThreshold  float64
	InstabilityScaleGeV   float64
	LambdaMin             float64
	LambdaMinScaleGeV     float64
	BounceAction          float64
	Log10LifetimeYears    float64
	AgeUniverseLog10Years float64
	Metastable            bool
	StableToPlanck        bool
	PrecisionWarning      string
	Statuses              []string
}

type rgState struct {
	lambda float64
	yt     float64
	g3     float64
}

func betaRG(s rgState) rgState {
	const loop = 1.0 / (16.0 * math.Pi * math.Pi)
	g1 := G1AtMZ
	g2 := G2AtMZ

	g1sq := g1 * g1
	g2sq := g2 * g2
	g3sq := s.g3 * s.g3
	yt2 := s.yt * s.yt
	lambda := s.lambda

	betaLambda := 24*lambda*lambda + 12*lambda*yt2 - 6*yt2*yt2 - 3*lambda*(3*g2sq+g1sq) + (3.0/8.0)*(2*g2sq*g2sq+math.Pow(g2sq+g1sq, 2))
	betaYt := s.yt * ((9.0/2.0)*yt2 - 8*g3sq - (9.0/4.0)*g2sq - (17.0/12.0)*g1sq)
	betaG3 := -7.0 * s.g3 * g3sq

	return rgState{lambda: loop * betaLambda, yt: loop * betaYt, g3: loop * betaG3}
}

func rk4Step(s rgState, h float64) rgState {
	k1 := betaRG(s)
	k2 := betaRG(rgState{lambda: s.lambda + 0.5*h*k1.lambda, yt: s.yt + 0.5*h*k1.yt, g3: s.g3 + 0.5*h*k1.g3})
	k3 := betaRG(rgState{lambda: s.lambda + 0.5*h*k2.lambda, yt: s.yt + 0.5*h*k2.yt, g3: s.g3 + 0.5*h*k2.g3})
	k4 := betaRG(rgState{lambda: s.lambda + h*k3.lambda, yt: s.yt + h*k3.yt, g3: s.g3 + h*k3.g3})
	return rgState{
		lambda: s.lambda + h*(k1.lambda+2*k2.lambda+2*k3.lambda+k4.lambda)/6,
		yt:     s.yt + h*(k1.yt+2*k2.yt+2*k3.yt+k4.yt)/6,
		g3:     s.g3 + h*(k1.g3+2*k2.g3+2*k3.g3+k4.g3)/6,
	}
}

// ComputeVacuumFate runs the more accurate of the two built-in audits: a
// one-loop QCD pole-to-running correction for the top Yukawa before the
// continuum RG flow.  Use ComputeVacuumFatePoleSeed to expose the raw pole-mass
// tree seed used as a stress-test.
func ComputeVacuumFate() VacuumFateResult {
	return computeVacuumFate("one-loop-QCD-MSbar-top-seed", 1-4*StrongCoupling/(3*math.Pi))
}

func ComputeVacuumFatePoleSeed() VacuumFateResult {
	return computeVacuumFate("tree-pole-top-seed", 1)
}

// ComputeVacuumFateEnsemble returns both the raw pole-seed stress test and the
// minimally matched MSbar-like seed. This is not a replacement for precision
// three-loop metastability analyses; it is an ASHA phenomenology audit.
func ComputeVacuumFateEnsemble() []VacuumFateResult {
	return []VacuumFateResult{ComputeVacuumFatePoleSeed(), ComputeVacuumFate()}
}

func computeVacuumFate(seedMode string, topRunningFactor float64) VacuumFateResult {
	lambda0 := HiggsMassGeV * HiggsMassGeV / (2 * StandardVEVGeV * StandardVEVGeV)
	yt0 := math.Sqrt2 * TopMassGeV * topRunningFactor / StandardVEVGeV
	g30 := math.Sqrt(4 * math.Pi * StrongCoupling)

	startT := math.Log(ZBosonMassGeV)
	endT := math.Log(PlanckMassGeV)
	thresholdT := math.Log(AshaBGapMajoranaMassGeV)
	steps := 20000
	h := (endT - startT) / float64(steps)

	s := rgState{lambda: lambda0, yt: yt0, g3: g30}
	lambdaMin := s.lambda
	lambdaMinT := startT
	instabilityScale := 0.0
	thresholdApplied := false
	lambdaBefore := math.NaN()
	lambdaAfter := math.NaN()
	prevLambda := s.lambda
	prevT := startT

	for i := 0; i < steps; i++ {
		t := startT + float64(i)*h
		nextT := t + h

		if !thresholdApplied && t < thresholdT && nextT >= thresholdT {
			// Step to threshold first.
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
			lambdaMin = s.lambda
			lambdaMinT = nextT
		}
		if prevLambda > 0 && s.lambda <= 0 && instabilityScale == 0 {
			frac := prevLambda / (prevLambda - s.lambda)
			instabilityScale = math.Exp(prevT + frac*h)
		}
		prevLambda = s.lambda
		prevT = nextT
	}

	stableToPlanck := instabilityScale == 0
	metastable := !stableToPlanck
	bounce := math.Inf(1)
	logLifetime := math.Inf(1)
	if metastable && lambdaMin < 0 {
		bounce = 8 * math.Pi * math.Pi / (3 * math.Abs(lambdaMin))
		logLifetime = math.Log10(HbarSecondsPerGeV/SecondsPerYear/instabilityScale) + bounce/math.Ln10
	}

	statuses := []string{StatusEmpiricalSealLoaded, StatusVacuumRGExecuted, StatusAshaThresholdApplied, StatusNotNativePrediction}
	if metastable {
		statuses = append(statuses, StatusVacuumMetastableConditional, StatusUniverseLifetimeNotNative)
	}

	return VacuumFateResult{
		Executed:              true,
		Method:                "one-loop SM RG for lambda, y_t and g3; g1/g2 held fixed at mZ; ASHA threshold applied at B-gap",
		SeedMode:              seedMode,
		TopRunningFactor:      topRunningFactor,
		InitialLambda:         lambda0,
		InitialYTop:           yt0,
		InitialG3:             g30,
		ThresholdScaleGeV:     AshaBGapMajoranaMassGeV,
		ThresholdDeltaLambda:  AshaThresholdDeltaLambda,
		ThresholdApplied:      thresholdApplied,
		LambdaBeforeThreshold: lambdaBefore,
		LambdaAfterThreshold:  lambdaAfter,
		InstabilityScaleGeV:   instabilityScale,
		LambdaMin:             lambdaMin,
		LambdaMinScaleGeV:     math.Exp(lambdaMinT),
		BounceAction:          bounce,
		Log10LifetimeYears:    logLifetime,
		AgeUniverseLog10Years: math.Log10(AgeOfUniverseYears),
		Metastable:            metastable,
		StableToPlanck:        stableToPlanck,
		PrecisionWarning:      "conditional one-loop audit; precision vacuum fate requires full MSbar matching, 2/3-loop beta functions, pole-to-running conversion and threshold convention",
		Statuses:              statuses,
	}
}
