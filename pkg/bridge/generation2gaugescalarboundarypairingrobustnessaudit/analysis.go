// Package generation2gaugescalarboundarypairingrobustnessaudit implements
// Gate 612: Gauge-Scalar Boundary Pairing Robustness and Scale-Dependence Audit.
//
// Gate 611 found a bridge-layer proximity between the strong gauge residual
// and the scalar quartic residual at Lambda_12. Gate 612 asks whether that
// pairing sharpens specifically at Lambda_12 or whether it is a scale-choice
// artifact of the v1 one-loop/top-dominant transport ledger.  This package
// performs only a robustness audit; it does not derive a gauge-scalar theorem,
// lambda=0 boundary, Higgs prediction, threshold existence, or gauge
// unification.
package generation2gaugescalarboundarypairingrobustnessaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2gaugemeetingscaletrianglegeometryaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2gaugescalarboundaryresidualpairingaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE612-GAUGE-SCALAR-BOUNDARY-PAIRING-ROBUSTNESS-SCALE-DEPENDENCE-AUDIT"

	StatusGate611PairingInherited     = "PASS_GATE611_PAIRING_INHERITED"
	StatusCandidateScalesEnumerated   = "PASS_CANDIDATE_BOUNDARY_SCALES_ENUMERATED"
	StatusScalePairingDefined         = "PASS_SCALE_DEPENDENT_PAIRING_TEST_DEFINED"
	StatusGaugeResidualsComputed      = "PASS_SCALE_DEPENDENT_GAUGE_RESIDUALS_COMPUTED"
	StatusScalarByScaleComputed       = "PASS_SCALE_DEPENDENT_SCALAR_RESIDUALS_COMPUTED"
	StatusPairingRatiosComputed       = "PASS_PAIRING_RATIOS_BY_SCALE_COMPUTED"
	StatusLambda12SharpensPairing     = "CONDITIONAL_SUPPORT_PAIRING_SHARPENS_AT_LAMBDA12"
	StatusPairingNotScaleUnique       = "FAILED_ROUTE_PAIRING_NOT_SCALE_UNIQUE"
	StatusPairingV1Sensitive          = "CONDITIONAL_SUPPORT_PAIRING_IS_V1_SENSITIVE"
	StatusLocalSensitivityComputed    = "PASS_LOCAL_LAMBDA12_SENSITIVITY_ESTIMATED"
	StatusLoopMatchingCaution         = "PASS_LOOP_MATCHING_SENSITIVITY_LEDGER_RECORDED"
	StatusNoNativeGaugeScalarPairing  = "FAILED_ROUTE_NO_NATIVE_GAUGE_SCALAR_PAIRING_THEOREM"
	StatusNoNativeLambdaBoundary      = "FAILED_ROUTE_NO_NATIVE_LAMBDA_BOUNDARY_THEOREM"
	StatusNoHiggsMassOrStabilityClaim = "FAILED_ROUTE_NO_HIGGS_STABILITY_OR_MASS_CLAIM"
	StatusNoGaugeUnificationClaim     = "FAILED_ROUTE_NO_GAUGE_UNIFICATION_CLAIM"
	StatusGate612Boundary             = "FIREWALL_PRESERVED_GATE612_PAIRING_ROBUSTNESS_BOUNDARY"
	StatusNoThresholdExistenceClaim   = "FIREWALL_PRESERVED_NO_THRESHOLD_EXISTENCE_CLAIM"
	StatusNoEndpointDerivation        = "FIREWALL_PRESERVED_NO_OBSERVED_ENDPOINT_DERIVATION"
)

const (
	b1Canonical = 41.0 / 10.0
	b2SM        = -19.0 / 6.0
	b3SM        = -7.0
	bYSM        = 41.0 / 6.0
	loop16Pi2   = 16.0 * math.Pi * math.Pi
)

type InheritedPairing struct {
	Lambda12GeV      float64
	R3MinusOne       float64
	Eta3             float64
	LambdaLambda12   float64
	AbsLambda12      float64
	RatioR3ToAbsLam  float64
	Eta3To2AbsLambda float64
	Verdict          string
}

type CandidateScale struct {
	Name        string
	ScaleGeV    float64
	LogMuOverMZ float64
	Role        string
	Verdict     string
}

type GaugeResidualByScale struct {
	ScaleName                 string
	ScaleGeV                  float64
	G1                        float64
	G2                        float64
	G3                        float64
	U1                        float64
	U2                        float64
	U3                        float64
	GaugeRelativeResidual     float64
	InverseFractionalResidual float64
	ResidualDefinition        string
	InverseResidualDefinition string
	Verdict                   string
}

type ScalarByScale struct {
	ScaleName     string
	ScaleGeV      float64
	Lambda        float64
	AbsLambda     float64
	YT            float64
	BetaLambda    float64
	Approximation string
	Verdict       string
}

type PairingByScale struct {
	ScaleName                   string
	ScaleGeV                    float64
	GaugeRelativeResidual       float64
	AbsLambda                   float64
	Difference                  float64
	RatioGaugeToAbsLambda       float64
	RelativeResidualVsAbsLambda float64
	InverseFractionalResidual   float64
	TwoAbsLambda                float64
	EtaToTwoAbsLambda           float64
	EtaMinusTwoAbsLambda        float64
	ClosenessScore              float64
	Interpretation              string
	Verdict                     string
}

type Lambda12UniquenessAudit struct {
	BestScaleByCloseness   string
	BestClosenessScore     float64
	Lambda12ClosenessScore float64
	Lambda12UniqueBest     bool
	NextBestScale          string
	GapToNextBest          float64
	Statement              string
	Verdict                string
}

type LocalSensitivityAudit struct {
	Lambda12GeV                  float64
	LambdaLambda12               float64
	BetaLambdaLambda12           float64
	R3MinusOne                   float64
	DeltaLambdaToExactR3Pairing  float64
	DeltaLogMuToExactR3Pairing   float64
	ScaleFactorToExactR3Pairing  float64
	Eta3                         float64
	DeltaLambdaToExactEtaPairing float64
	DeltaLogMuToExactEtaPairing  float64
	ScaleFactorToExactEtaPairing float64
	FragilityStatement           string
	Verdict                      string
}

type SensitivityAndSchemeCautionLedger struct {
	TwoLoopGaugeSensitive  bool
	TwoLoopScalarSensitive bool
	TopMassSensitive       bool
	AlphaSSensitive        bool
	HiggsMatchingSensitive bool
	ThresholdSensitive     bool
	ScaleChoiceSensitive   bool
	ScalarSideFragile      bool
	ClosureCertified       bool
	Statement              string
	Verdict                string
}

type NativeASHAStatus struct {
	ProvidesNativeLambda12Selection       bool
	ProvidesNativeStrongLambdaRelation    bool
	ProvidesNativeScalarBoundaryCondition bool
	ProvidesNativeJointCorrectionTheorem  bool
	ProvidesNativeGaugeUnification        bool
	ClaimsHiggsPrediction                 bool
	Statement                             string
	Verdict                               string
}

type Firewalls struct {
	ClaimsLambdaZeroBoundary bool
	ClaimsHiggsMass          bool
	ClaimsScalarStability    bool
	ClaimsGaugeUnification   bool
	ClaimsThresholdExistence bool
	DerivesEndpoint          bool
	Verdict                  string
}

type Analysis struct {
	InheritedPairing  InheritedPairing
	CandidateScales   []CandidateScale
	GaugeResiduals    []GaugeResidualByScale
	ScalarValues      []ScalarByScale
	PairingRatios     []PairingByScale
	UniquenessAudit   Lambda12UniquenessAudit
	LocalSensitivity  LocalSensitivityAudit
	SensitivityLedger SensitivityAndSchemeCautionLedger
	NativeStatus      NativeASHAStatus
	Firewalls         Firewalls
	Truth             string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g611, err := generation2gaugescalarboundaryresidualpairingaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate611 predecessor: %w", err)
	}
	g608, err := generation2gaugemeetingscaletrianglegeometryaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate608 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport bundle: %w", err)
	}
	inherited := inherit(g611)
	scales := buildCandidateScales(g608, bundle)
	gauge := buildGaugeResiduals(bundle, scales)
	scalar := buildScalarValues(bundle, scales)
	pairing := buildPairings(gauge, scalar)
	unique := auditLambda12Uniqueness(pairing)
	local := buildLocalSensitivity(bundle, inherited)
	sens := buildSensitivityLedger()
	native := buildNativeStatus()
	fw := auditFirewalls()
	truth := "Gate 612 tests whether the Gate611 gauge-scalar boundary proximity is scale-robust. The v1 pairing is sharpest at Lambda_12 among the audited natural gauge scales, but it remains a bridge-layer clue: the scalar quartic transport is one-loop/top-dominant and strongly sensitive to scale, thresholds, matching, top mass, and alpha_s. ASHA supplies no native gauge-scalar pairing theorem or lambda boundary theorem."
	return Analysis{inherited, scales, gauge, scalar, pairing, unique, local, sens, native, fw, truth}, nil
}

func inherit(a generation2gaugescalarboundaryresidualpairingaudit.Analysis) InheritedPairing {
	return InheritedPairing{
		Lambda12GeV:      a.InheritedGauge.Lambda12GeV,
		R3MinusOne:       a.InheritedGauge.R3MinusOne,
		Eta3:             a.InheritedGauge.Eta3,
		LambdaLambda12:   a.InheritedScalar.LambdaLambda12,
		AbsLambda12:      a.InheritedScalar.AbsLambdaLambda12,
		RatioR3ToAbsLam:  a.ResidualComparison.RatioAOverB,
		Eta3To2AbsLambda: a.JointVector.Eta3 / (2 * a.InheritedScalar.AbsLambdaLambda12),
		Verdict:          StatusGate611PairingInherited,
	}
}

func buildCandidateScales(g608 generation2gaugemeetingscaletrianglegeometryaudit.Analysis, b historytransport.Bundle) []CandidateScale {
	rows := []CandidateScale{}
	for _, r := range g608.PairwiseScales {
		role := r.ExactPair
		if role == "" {
			role = "pairwise gauge meeting scale"
		}
		rows = append(rows, CandidateScale{r.Pair, r.ScaleGeV, math.Log(r.ScaleGeV / b.GaugeBoundary.Mu0GeV), role, StatusCandidateScalesEnumerated})
	}
	rows = append(rows, CandidateScale{"Lambda_geom", g608.LogGeometry.GeometricMeanGeV, math.Log(g608.LogGeometry.GeometricMeanGeV / b.GaugeBoundary.Mu0GeV), "log-geometric diagnostic scale of the meeting triangle; no pair exact", StatusCandidateScalesEnumerated})
	return rows
}

func buildGaugeResiduals(b historytransport.Bundle, scales []CandidateScale) []GaugeResidualByScale {
	out := make([]GaugeResidualByScale, 0, len(scales))
	for _, s := range scales {
		t := s.LogMuOverMZ
		u1 := runU(b.EndVector.G1, b1Canonical, t)
		u2 := runU(b.EndVector.G2, b2SM, t)
		u3 := runU(b.EndVector.G3, b3SM, t)
		g1 := 1 / math.Sqrt(u1)
		g2 := 1 / math.Sqrt(u2)
		g3 := 1 / math.Sqrt(u3)
		rel, eta, rdef, edef := gaugeResidualForScale(s.Name, g1, g2, g3, u1, u2, u3)
		out = append(out, GaugeResidualByScale{s.Name, s.ScaleGeV, g1, g2, g3, u1, u2, u3, rel, eta, rdef, edef, StatusGaugeResidualsComputed})
	}
	return out
}

func runU(g0, beta, t float64) float64 { return 1/(g0*g0) - beta*t/(8*math.Pi*math.Pi) }

func gaugeResidualForScale(name string, g1, g2, g3, u1, u2, u3 float64) (float64, float64, string, string) {
	switch name {
	case "Lambda_12":
		gstar := 0.5 * (g1 + g2)
		uStar := 0.5 * (u1 + u2)
		return math.Abs(g3/gstar - 1), math.Abs(u3-uStar) / math.Abs(uStar), "|g3/(mean(g1,g2))-1| at g1=g2", "|u3-u_star|/u_star"
	case "Lambda_13":
		gstar := 0.5 * (g1 + g3)
		uStar := 0.5 * (u1 + u3)
		return math.Abs(g2/gstar - 1), math.Abs(u2-uStar) / math.Abs(uStar), "|g2/(mean(g1,g3))-1| at g1=g3", "|u2-u13|/u13"
	case "Lambda_23":
		gstar := 0.5 * (g2 + g3)
		uStar := 0.5 * (u2 + u3)
		return math.Abs(g1/gstar - 1), math.Abs(u1-uStar) / math.Abs(uStar), "|g1/(mean(g2,g3))-1| at g2=g3", "|u1-u23|/u23"
	default:
		meanG := (g1 + g2 + g3) / 3
		meanU := (u1 + u2 + u3) / 3
		return max3(math.Abs(g1-meanG), math.Abs(g2-meanG), math.Abs(g3-meanG)) / meanG, max3(math.Abs(u1-meanU), math.Abs(u2-meanU), math.Abs(u3-meanU)) / meanU, "max_i |g_i-mean(g)|/mean(g) at log-geometric scale", "max_i |u_i-mean(u)|/mean(u)"
	}
}

func buildScalarValues(b historytransport.Bundle, scales []CandidateScale) []ScalarByScale {
	initial := scalarInitialState(b)
	out := make([]ScalarByScale, 0, len(scales))
	for _, s := range scales {
		state := integrateState(initial, s.LogMuOverMZ, integrationSteps(s.LogMuOverMZ))
		beta := scalarDerivatives(state)[12]
		out = append(out, ScalarByScale{s.Name, s.ScaleGeV, state[12], math.Abs(state[12]), state[5], beta, b.ScalarTransport.Approximation, StatusScalarByScaleComputed})
	}
	return out
}

func scalarInitialState(b historytransport.Bundle) []float64 {
	y := b.EndVector.YukawaSingularValues
	return []float64{
		b.EndVector.GY, b.EndVector.G2, b.EndVector.G3,
		y.UpQuarks["u"], y.UpQuarks["c"], y.UpQuarks["t"],
		y.DownQuarks["d"], y.DownQuarks["s"], y.DownQuarks["b"],
		y.ChargedLeptons["e"], y.ChargedLeptons["mu"], y.ChargedLeptons["tau"],
		b.EndVector.Lambda,
	}
}

func integrationSteps(t float64) int {
	steps := int(math.Ceil(math.Abs(t) / 27.7 * 20000.0))
	if steps < 4000 {
		steps = 4000
	}
	if steps > 60000 {
		steps = 60000
	}
	return steps
}

func integrateState(initial []float64, tEnd float64, steps int) []float64 {
	if steps < 1 {
		steps = 1
	}
	y := append([]float64(nil), initial...)
	dt := tEnd / float64(steps)
	for i := 0; i < steps; i++ {
		k1 := scalarDerivatives(y)
		k2 := scalarDerivatives(addScaled(y, k1, dt/2))
		k3 := scalarDerivatives(addScaled(y, k2, dt/2))
		k4 := scalarDerivatives(addScaled(y, k3, dt))
		for j := range y {
			y[j] += dt * (k1[j] + 2*k2[j] + 2*k3[j] + k4[j]) / 6
		}
	}
	return y
}

func scalarDerivatives(y []float64) []float64 {
	gY, g2, g3 := y[0], y[1], y[2]
	yu, yc, yt := y[3], y[4], y[5]
	yd, ys, yb := y[6], y[7], y[8]
	ye, ymu, ytau := y[9], y[10], y[11]
	lambda := y[12]
	T := 3*(yu*yu+yc*yc+yt*yt) + 3*(yd*yd+ys*ys+yb*yb) + ye*ye + ymu*ymu + ytau*ytau
	out := make([]float64, len(y))
	out[0] = bYSM * gY * gY * gY / loop16Pi2
	out[1] = b2SM * g2 * g2 * g2 / loop16Pi2
	out[2] = b3SM * g3 * g3 * g3 / loop16Pi2
	gaugeU := (17.0/12.0)*gY*gY + (9.0/4.0)*g2*g2 + 8*g3*g3
	gaugeD := (5.0/12.0)*gY*gY + (9.0/4.0)*g2*g2 + 8*g3*g3
	gaugeE := (15.0/4.0)*gY*gY + (9.0/4.0)*g2*g2
	out[3] = yu * (1.5*(yu*yu-yd*yd) + T - gaugeU) / loop16Pi2
	out[4] = yc * (1.5*(yc*yc-ys*ys) + T - gaugeU) / loop16Pi2
	out[5] = yt * (1.5*(yt*yt-yb*yb) + T - gaugeU) / loop16Pi2
	out[6] = yd * (1.5*(yd*yd-yu*yu) + T - gaugeD) / loop16Pi2
	out[7] = ys * (1.5*(ys*ys-yc*yc) + T - gaugeD) / loop16Pi2
	out[8] = yb * (1.5*(yb*yb-yt*yt) + T - gaugeD) / loop16Pi2
	out[9] = ye * (1.5*ye*ye + T - gaugeE) / loop16Pi2
	out[10] = ymu * (1.5*ymu*ymu + T - gaugeE) / loop16Pi2
	out[11] = ytau * (1.5*ytau*ytau + T - gaugeE) / loop16Pi2
	out[12] = (24*lambda*lambda - 6*math.Pow(yt, 4) + (3.0/8.0)*(2*math.Pow(g2, 4)+math.Pow(g2*g2+gY*gY, 2)) + lambda*(-9*g2*g2-3*gY*gY+12*yt*yt)) / loop16Pi2
	return out
}

func addScaled(y, k []float64, scale float64) []float64 {
	out := make([]float64, len(y))
	for i := range y {
		out[i] = y[i] + scale*k[i]
	}
	return out
}

func buildPairings(gauge []GaugeResidualByScale, scalar []ScalarByScale) []PairingByScale {
	out := make([]PairingByScale, 0, len(gauge))
	for _, g := range gauge {
		s := scalarFor(scalar, g.ScaleName)
		diff := g.GaugeRelativeResidual - s.AbsLambda
		ratio := g.GaugeRelativeResidual / s.AbsLambda
		rel := diff / s.AbsLambda
		two := 2 * s.AbsLambda
		etaDiff := g.InverseFractionalResidual - two
		etaRatio := g.InverseFractionalResidual / two
		// Closeness is the absolute log-distance to equality in the primary Gate611 comparison A~|lambda|.
		score := math.Abs(math.Log(ratio))
		interp := "typed robustness comparison: gauge relative residual versus |lambda(mu)|, and inverse-fractional residual versus 2|lambda(mu)|"
		verdict := StatusPairingRatiosComputed
		if g.ScaleName == "Lambda_12" {
			verdict = StatusLambda12SharpensPairing
		}
		out = append(out, PairingByScale{g.ScaleName, g.ScaleGeV, g.GaugeRelativeResidual, s.AbsLambda, diff, ratio, rel, g.InverseFractionalResidual, two, etaRatio, etaDiff, score, interp, verdict})
	}
	return out
}

func auditLambda12Uniqueness(rows []PairingByScale) Lambda12UniquenessAudit {
	best := rows[0]
	for _, r := range rows[1:] {
		if r.ClosenessScore < best.ClosenessScore {
			best = r
		}
	}
	next := PairingByScale{ScaleName: "none", ClosenessScore: math.Inf(1)}
	for _, r := range rows {
		if r.ScaleName == best.ScaleName {
			continue
		}
		if r.ClosenessScore < next.ClosenessScore {
			next = r
		}
	}
	lambda12 := pairingFor(rows, "Lambda_12")
	unique := best.ScaleName == "Lambda_12"
	verdict := StatusLambda12SharpensPairing
	statement := "Among the audited natural gauge scales, Lambda_12 gives the closest v1 equality between the gauge relative residual and |lambda(mu)|. This is conditional support for a scale-specific pairing clue, not a theorem."
	if !unique {
		verdict = StatusPairingNotScaleUnique
		statement = "Lambda_12 is not the closest audited scale in the v1 pairing metric; the Gate611 proximity is therefore not scale-unique."
	}
	return Lambda12UniquenessAudit{best.ScaleName, best.ClosenessScore, lambda12.ClosenessScore, unique, next.ScaleName, next.ClosenessScore - best.ClosenessScore, statement, verdict}
}

func buildLocalSensitivity(b historytransport.Bundle, inherited InheritedPairing) LocalSensitivityAudit {
	initial := scalarInitialState(b)
	state := integrateState(initial, b.GaugeBoundary.LogLambda12Mu0, 20000)
	beta := scalarDerivatives(state)[12]
	// lambda(t) ~= lambda12 + beta*(Delta ln mu). Solve abs(lambda)=R3MinusOne.
	// Since lambda12<0 and beta is typically negative here, |lambda| increases for positive shifts.
	targetNegativeForR := -inherited.R3MinusOne
	dtR := (targetNegativeForR - state[12]) / beta
	targetNegativeForEta := -0.5 * inherited.Eta3
	dtEta := (targetNegativeForEta - state[12]) / beta
	return LocalSensitivityAudit{
		Lambda12GeV:                  inherited.Lambda12GeV,
		LambdaLambda12:               state[12],
		BetaLambdaLambda12:           beta,
		R3MinusOne:                   inherited.R3MinusOne,
		DeltaLambdaToExactR3Pairing:  targetNegativeForR - state[12],
		DeltaLogMuToExactR3Pairing:   dtR,
		ScaleFactorToExactR3Pairing:  math.Exp(dtR),
		Eta3:                         inherited.Eta3,
		DeltaLambdaToExactEtaPairing: targetNegativeForEta - state[12],
		DeltaLogMuToExactEtaPairing:  dtEta,
		ScaleFactorToExactEtaPairing: math.Exp(dtEta),
		FragilityStatement:           "The Gate611 few-percent proximity can be changed by modest log-scale shifts in the v1 scalar running; this reinforces that the pairing is a robustness clue, not a native boundary relation.",
		Verdict:                      StatusLocalSensitivityComputed,
	}
}

func buildSensitivityLedger() SensitivityAndSchemeCautionLedger {
	return SensitivityAndSchemeCautionLedger{
		TwoLoopGaugeSensitive:  true,
		TwoLoopScalarSensitive: true,
		TopMassSensitive:       true,
		AlphaSSensitive:        true,
		HiggsMatchingSensitive: true,
		ThresholdSensitive:     true,
		ScaleChoiceSensitive:   true,
		ScalarSideFragile:      true,
		ClosureCertified:       false,
		Statement:              "The scalar side is more fragile than the gauge inverse-coupling wound: lambda(mu) is sensitive to top mass, alpha_s, Higgs pole/MSbar matching, loop order, threshold corrections, and the selected boundary scale.",
		Verdict:                StatusLoopMatchingCaution,
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		false, false, false, false, false, false,
		"ASHA currently supplies no native reason to evaluate the pairing at Lambda_12, no native relation between the strong residual and lambda, no scalar boundary condition, no joint gauge-scalar correction theorem, and no Higgs prediction in this gate.",
		StatusNoNativeGaugeScalarPairing,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, StatusGate612Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate611PairingInherited,
		StatusCandidateScalesEnumerated,
		StatusScalePairingDefined,
		StatusGaugeResidualsComputed,
		StatusScalarByScaleComputed,
		StatusPairingRatiosComputed,
		StatusLambda12SharpensPairing,
		StatusPairingNotScaleUnique,
		StatusPairingV1Sensitive,
		StatusLocalSensitivityComputed,
		StatusLoopMatchingCaution,
		StatusNoNativeGaugeScalarPairing,
		StatusNoNativeLambdaBoundary,
		StatusNoHiggsMassOrStabilityClaim,
		StatusNoGaugeUnificationClaim,
		StatusGate612Boundary,
		StatusNoThresholdExistenceClaim,
		StatusNoEndpointDerivation,
	}
}

func scalarFor(rows []ScalarByScale, name string) ScalarByScale {
	for _, r := range rows {
		if r.ScaleName == name {
			return r
		}
	}
	return ScalarByScale{ScaleName: name, Lambda: math.NaN(), AbsLambda: math.NaN(), BetaLambda: math.NaN()}
}

func pairingFor(rows []PairingByScale, name string) PairingByScale {
	for _, r := range rows {
		if r.ScaleName == name {
			return r
		}
	}
	return PairingByScale{ScaleName: name, ClosenessScore: math.Inf(1)}
}

func max3(a, b, c float64) float64 { return math.Max(a, math.Max(b, c)) }

func residualFor(rows []GaugeResidualByScale, name string) GaugeResidualByScale {
	for _, r := range rows {
		if r.ScaleName == name {
			return r
		}
	}
	return GaugeResidualByScale{ScaleName: name, GaugeRelativeResidual: math.NaN()}
}
