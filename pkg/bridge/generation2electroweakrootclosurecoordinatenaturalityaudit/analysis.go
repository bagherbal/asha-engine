// Package generation2electroweakrootclosurecoordinatenaturalityaudit implements
// Gate 665: ElectroweakRoot Closure Coordinate-Naturality Audit.
//
// Gate 664 established a v1 dual-root alignment between the electroweak
// meeting root g1=g2 and the E72 scalar/flavor/boundary closure root. Gate 665
// audits whether that alignment is coordinate-natural across typed gauge
// residual coordinates or whether it is sealed to the coupling-amplitude ratio
// coordinate. It preserves the distinction between endpoint bridge coordinates
// and RG-native inverse-coupling variables.
package generation2electroweakrootclosurecoordinatenaturalityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate664 "github.com/bagherbal/asha-engine/pkg/bridge/generation2electroweakmeetingdeficitclosuredualrootaudit"
)

const (
	AuditID = "GATE665-ELECTROWEAK-ROOT-CLOSURE-COORDINATE-NATURALITY-AUDIT"

	StatusGate664DualRootInherited        = "PASS_GATE664_DUAL_ROOT_ALIGNMENT_INHERITED"
	StatusCommonRootAudited               = "PASS_COMMON_ROOT_STATEMENT_AUDITED"
	StatusLocalFactorizationAudited       = "PASS_LOCAL_FACTORIZATION_AUDITED"
	StatusGaugeCoordinateFamilyAudited    = "PASS_GAUGE_COORDINATE_FAMILY_AUDITED"
	StatusAmplitudeCoordinateSupported    = "CONDITIONAL_SUPPORT_DUAL_ROOT_ALIGNMENT_IN_AMPLITUDE_RATIO_COORDINATE"
	StatusCoordinateNaturalityUncertified = "CONDITIONAL_SUPPORT_COORDINATE_NATURALITY_REMAINS_UNCERTIFIED"
	StatusBridgeCoordinateSeal            = "CONDITIONAL_SUPPORT_BOUNDARY_WEIGHTED_DEFICIT_CLOSURE_IS_BRIDGE_COORDINATE_SEAL"
	StatusInverseCoordinateFails          = "FAILED_ROUTE_INVERSE_COUPLING_COORDINATE_DOES_NOT_YET_CERTIFY_SAME_ALIGNMENT"
	StatusNoNativeDualRootTheorem         = "FAILED_ROUTE_NO_NATIVE_DUAL_ROOT_ALIGNMENT_THEOREM"
	StatusNoNativeSevenOver72Theorem      = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoFullUncertaintyPropagation    = "FAILED_ROUTE_NO_FULL_UNCERTAINTY_PROPAGATION"
	StatusNoBoundaryStressDerivation      = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoNativeTransportTheorem        = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoHiggsStabilityGaugeFlavor     = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate665Boundary                 = "FIREWALL_PRESERVED_GATE665_COORDINATE_NATURALITY_BOUNDARY"
)

const (
	sevenOver72    = 7.0 / 72.0
	transportSteps = 2000
)

type Gate664Inheritance struct {
	DualRootInherited   bool
	Lambda12GeV         float64
	T12                 float64
	KSum                float64
	E72AtLambda12       float64
	ClosureRootRatio    float64
	ClosureRootDeltaLog float64
	DE72Dt              float64
	TransverseCrossing  bool
	NoNativeDualRoot    bool
	NoNativeSevenOver72 bool
	NoFullUncertainty   bool
	NoBoundaryStress    bool
	Verdict             string
}

type TransportSeed struct {
	Mu0GeV        float64
	Lambda12GeV   float64
	T12           float64
	G1MZ          float64
	GYMZ          float64
	G2MZ          float64
	G3MZ          float64
	InitialVector []float64
	Verdict       string
}

type CommonRootAudit struct {
	DefinitionF12       string
	DefinitionE72       string
	T12Analytic         float64
	Mu12GeV             float64
	F12AtRoot           float64
	E72AmplitudeAtRoot  float64
	WBestAtRoot         float64
	WBestMinus7Over72   float64
	ConditionalRootPass bool
	Verdict             string
}

type LocalFactorizationAudit struct {
	Window                 float64
	Samples                int
	CAmplitudeForF12       float64
	InterceptAmplitudeF12  float64
	RelativeResidualF12    float64
	CAmplitudeForU12       float64
	InterceptAmplitudeU12  float64
	RelativeResidualU12    float64
	AmplitudeFactorLikeF12 bool
	InverseFactorLikeU12   bool
	Verdict                string
}

type CoordinateRow struct {
	Name                  string
	Definition            string
	GaugeResidualAtT12    float64
	WBestAtT12            float64
	WBestMinus7Over72     float64
	E72AtSevenOver72      float64
	TZero                 float64
	DeltaLogRootFromT12   float64
	MuZeroOverLambda12    float64
	RootFoundNearLambda12 bool
	NearSevenOver72       bool
	CoordinateClass       string
	Verdict               string
}

type CoordinateFamilyAudit struct {
	Rows                    []CoordinateRow
	AmplitudeRowsNearWeight int
	InverseRowsNearWeight   int
	CoordinateRobust        bool
	AmplitudeNatural        bool
	RGNativeInverseNatural  bool
	Verdict                 string
}

type CoordinateNaturalityVerdict struct {
	Classification string
	Outcomes       []string
	Verdict        string
}

type SourceTypeInterpretation struct {
	Interpretations []string
	Verdict         string
}

type VerdictDiscipline struct {
	ClaimsNativeDualRootTheorem      bool
	ClaimsNativeSevenOver72Theorem   bool
	ClaimsFullUncertaintyPropagation bool
	ClaimsBoundaryStressDerivation   bool
	ClaimsNativeTransportTheorem     bool
	ClaimsHiggsPrediction            bool
	ClaimsScalarStability            bool
	ClaimsFlavorDerivation           bool
	ClaimsGaugeUnification           bool
	ClaimsCKMPMNSDerivation          bool
	Verdict                          string
}

type Analysis struct {
	Inherited      Gate664Inheritance
	Seed           TransportSeed
	CommonRoot     CommonRootAudit
	Factorization  LocalFactorizationAudit
	Coordinates    CoordinateFamilyAudit
	CoordinateSeal CoordinateNaturalityVerdict
	Source         SourceTypeInterpretation
	Discipline     VerdictDiscipline
	Truth          string
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
	g664, err := gate664.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate664 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g664)
	seed := buildSeed(g664.Seed)
	common := buildCommonRoot(seed, inherited.KSum)
	factor := buildLocalFactorization(seed, inherited.KSum, seed.T12)
	coords := buildCoordinateFamily(seed, inherited.KSum, seed.T12)
	seal := buildCoordinateVerdict(coords)
	source := buildSourceType(coords)
	discipline := VerdictDiscipline{Verdict: StatusGate665Boundary}
	truth := "Gate 665 classifies the Gate664 dual-root alignment as strongest in the coupling-amplitude EW-mean ratio coordinate. Squared-coupling, alpha, inverse-coupling, and log-coupling residual coordinates do not keep the 7/72 weight or the same near-root closure. Therefore the active E72 bridge is an amplitude-coordinate bridge seal in v1, not yet a coordinate-natural or RG-native inverse-coupling theorem."
	return Analysis{Inherited: inherited, Seed: seed, CommonRoot: common, Factorization: factor, Coordinates: coords, CoordinateSeal: seal, Source: source, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate664.Analysis) Gate664Inheritance {
	return Gate664Inheritance{
		DualRootInherited:   g.DualRoot.AlignedInV1 && g.Transversality.E72Transverse,
		Lambda12GeV:         g.Seed.Lambda12GeV,
		T12:                 g.Seed.T12,
		KSum:                g.Inherited.KSum,
		E72AtLambda12:       g.Inherited.E72AtLambda12,
		ClosureRootRatio:    g.DualRoot.MuEOverMu12,
		ClosureRootDeltaLog: g.DualRoot.DeltaLogMuEOverMu12,
		DE72Dt:              g.Transversality.DE72DtAtLambda12,
		TransverseCrossing:  g.Transversality.E72Transverse && !g.Transversality.SlopeTied,
		NoNativeDualRoot:    !g.Discipline.ClaimsNativeDualRootTheorem,
		NoNativeSevenOver72: !g.Discipline.ClaimsNativeSevenOver72Theorem,
		NoFullUncertainty:   !g.Discipline.ClaimsFullUncertaintyPropagation,
		NoBoundaryStress:    !g.Discipline.ClaimsBoundaryStressDerivation,
		Verdict:             StatusGate664DualRootInherited,
	}
}

func buildSeed(s gate664.TransportSeed) TransportSeed {
	return TransportSeed{Mu0GeV: s.Mu0GeV, Lambda12GeV: s.Lambda12GeV, T12: s.T12, G1MZ: s.G1MZ, GYMZ: s.GYMZ, G2MZ: s.G2MZ, G3MZ: s.G3MZ, InitialVector: append([]float64(nil), s.InitialVector...), Verdict: StatusGate664DualRootInherited}
}

func buildCommonRoot(seed TransportSeed, ksum float64) CommonRootAudit {
	t12 := meetingLog(seed.G1MZ, seed.G2MZ, 41.0/10.0, -19.0/6.0)
	e := eAt(seed, ksum, t12, gaugeResidualAmplitude)
	_, absLam, g, _, _ := closureValues(seed, ksum, t12, gaugeResidualAmplitude)
	wBest := (ksum - absLam) / (g - absLam)
	return CommonRootAudit{
		DefinitionF12:       "F12(mu)=g1(mu)-g2(mu)",
		DefinitionE72:       "E72^g(mu)=K_sum-[(65/72)|lambda(mu)|+(7/72)(g3/gEW-1)]",
		T12Analytic:         t12,
		Mu12GeV:             seed.Mu0GeV * math.Exp(t12),
		F12AtRoot:           f12(seed, t12),
		E72AmplitudeAtRoot:  e,
		WBestAtRoot:         wBest,
		WBestMinus7Over72:   wBest - sevenOver72,
		ConditionalRootPass: math.Abs(e) < 1e-8 && math.Abs(wBest-sevenOver72) < 1e-6,
		Verdict:             join(StatusCommonRootAudited, StatusAmplitudeCoordinateSupported),
	}
}

func buildLocalFactorization(seed TransportSeed, ksum, t12 float64) LocalFactorizationAudit {
	window := 0.1
	offsets := []float64{-window, -0.05, 0, 0.05, window}
	eVals := make([]float64, 0, len(offsets))
	fVals := make([]float64, 0, len(offsets))
	uVals := make([]float64, 0, len(offsets))
	for _, d := range offsets {
		t := t12 + d
		eVals = append(eVals, eAt(seed, ksum, t, gaugeResidualAmplitude))
		fVals = append(fVals, f12(seed, t))
		uVals = append(uVals, u12(seed, t))
	}
	cF, bF, relF := fitAffineRelative(eVals, fVals)
	cU, bU, relU := fitAffineRelative(eVals, uVals)
	return LocalFactorizationAudit{Window: window, Samples: len(offsets), CAmplitudeForF12: cF, InterceptAmplitudeF12: bF, RelativeResidualF12: relF, CAmplitudeForU12: cU, InterceptAmplitudeU12: bU, RelativeResidualU12: relU, AmplitudeFactorLikeF12: relF < 0.01, InverseFactorLikeU12: relU < 0.01, Verdict: StatusLocalFactorizationAudited}
}

func buildCoordinateFamily(seed TransportSeed, ksum, t12 float64) CoordinateFamilyAudit {
	defs := []struct {
		name  string
		def   string
		class string
		fn    residualFunc
	}{
		{"amplitude ratio", "G_g=g3/gEW-1", "amplitude", gaugeResidualAmplitude},
		{"squared-coupling ratio", "G_g2=g3^2/gEW^2-1", "strength", gaugeResidualSquared},
		{"alpha ratio", "G_alpha=alpha3/alphaEW-1", "strength", gaugeResidualAlpha},
		{"inverse-coupling ratio", "G_u=uEW/u3-1", "inverse", gaugeResidualInverse},
		{"log-coupling residual", "G_log=ln(g3/gEW)", "log", gaugeResidualLog},
	}
	rows := make([]CoordinateRow, 0, len(defs))
	ampNear, invNear := 0, 0
	for _, d := range defs {
		_, absLam, gauge, _, e := closureValues(seed, ksum, t12, d.fn)
		wBest := (ksum - absLam) / (gauge - absLam)
		tz, _, found := findRootNear(seed, ksum, d.fn, t12, 0.5)
		nearWeight := math.Abs(wBest-sevenOver72) < 1e-4
		nearRoot := found && math.Abs(tz-t12) < 1e-3
		if d.class == "amplitude" && nearWeight && nearRoot {
			ampNear++
		}
		if d.class == "inverse" && nearWeight && nearRoot {
			invNear++
		}
		rows = append(rows, CoordinateRow{Name: d.name, Definition: d.def, GaugeResidualAtT12: gauge, WBestAtT12: wBest, WBestMinus7Over72: wBest - sevenOver72, E72AtSevenOver72: e, TZero: tz, DeltaLogRootFromT12: tz - t12, MuZeroOverLambda12: math.Exp(tz - t12), RootFoundNearLambda12: nearRoot, NearSevenOver72: nearWeight, CoordinateClass: d.class, Verdict: StatusGaugeCoordinateFamilyAudited})
	}
	coordRobust := ampNear == len(rows)
	return CoordinateFamilyAudit{Rows: rows, AmplitudeRowsNearWeight: ampNear, InverseRowsNearWeight: invNear, CoordinateRobust: coordRobust, AmplitudeNatural: ampNear == 1, RGNativeInverseNatural: invNear > 0, Verdict: join(StatusGaugeCoordinateFamilyAudited, StatusAmplitudeCoordinateSupported, StatusCoordinateNaturalityUncertified, StatusInverseCoordinateFails)}
}

func buildCoordinateVerdict(c CoordinateFamilyAudit) CoordinateNaturalityVerdict {
	outcomes := []string{
		"amplitude-natural: the EW-mean coupling-amplitude residual gives w_best within O(1e-6) of 7/72 and an aligned E72 root",
		"not RG-native yet: the inverse-coupling coordinate does not keep the 7/72 closure or near-root test",
		"not coordinate-robust: squared-coupling, alpha, and log coordinates shift the best weight away from 7/72",
		"bridge-coordinate seal: current evidence selects an endpoint/amplitude coordinate layer rather than a native one-loop inverse-coupling layer",
	}
	return CoordinateNaturalityVerdict{Classification: "amplitude-coordinate bridge seal", Outcomes: outcomes, Verdict: join(StatusAmplitudeCoordinateSupported, StatusCoordinateNaturalityUncertified, StatusBridgeCoordinateSeal, StatusInverseCoordinateFails)}
}

func buildSourceType(c CoordinateFamilyAudit) SourceTypeInterpretation {
	interpretations := []string{
		"if amplitude coordinates win, the closure belongs to the endpoint/canonical coupling-amplitude bridge layer",
		"if inverse-coupling variables won, it would be closer to the RG-native gauge transport lane; Gate665 does not certify this",
		"the 7/72 weight remains active in the amplitude boundary interpolation formula only",
		"coordinate naturality remains an open theorem target and must not be replaced by a source claim",
	}
	return SourceTypeInterpretation{Interpretations: interpretations, Verdict: join(StatusBridgeCoordinateSeal, StatusNoNativeDualRootTheorem, StatusNoNativeSevenOver72Theorem)}
}

// residualFunc returns the gauge residual G(mu) used in W72(mu).
type residualFunc func(seed TransportSeed, t float64) float64

func gaugeResidualAmplitude(seed TransportSeed, t float64) float64 {
	g1, g2, g3 := gaugeValues(seed, t)
	return g3/(0.5*(g1+g2)) - 1.0
}

func gaugeResidualSquared(seed TransportSeed, t float64) float64 {
	g1, g2, g3 := gaugeValues(seed, t)
	gEW := 0.5 * (g1 + g2)
	return (g3*g3)/(gEW*gEW) - 1.0
}

func gaugeResidualAlpha(seed TransportSeed, t float64) float64 { return gaugeResidualSquared(seed, t) }

func gaugeResidualInverse(seed TransportSeed, t float64) float64 {
	g1, g2, g3 := gaugeValues(seed, t)
	uEW := 0.5 * (1.0/(g1*g1) + 1.0/(g2*g2))
	u3 := 1.0 / (g3 * g3)
	return uEW/u3 - 1.0
}

func gaugeResidualLog(seed TransportSeed, t float64) float64 {
	g1, g2, g3 := gaugeValues(seed, t)
	return math.Log(g3 / (0.5 * (g1 + g2)))
}

func gaugeValues(seed TransportSeed, t float64) (g1, g2, g3 float64) {
	return runCanonicalGauge(seed.G1MZ, 41.0/10.0, t), runCanonicalGauge(seed.G2MZ, -19.0/6.0, t), runCanonicalGauge(seed.G3MZ, -7.0, t)
}

func closureValues(seed TransportSeed, ksum, t float64, residual residualFunc) (lambda, absLambda, gaugeResidual, w72, e72 float64) {
	state := integrate(seed.InitialVector, t, transportSteps)
	lambda = state[12]
	absLambda = math.Abs(lambda)
	gaugeResidual = residual(seed, t)
	w72 = absLambda + sevenOver72*(gaugeResidual-absLambda)
	e72 = ksum - w72
	return
}

func eAt(seed TransportSeed, ksum, t float64, residual residualFunc) float64 {
	_, _, _, _, e := closureValues(seed, ksum, t, residual)
	return e
}

func findRootNear(seed TransportSeed, ksum float64, residual residualFunc, center, halfWidth float64) (float64, float64, bool) {
	lo, hi := center-halfWidth, center+halfWidth
	flo, fhi := eAt(seed, ksum, lo, residual), eAt(seed, ksum, hi, residual)
	if flo == 0 {
		return lo, flo, true
	}
	if fhi == 0 {
		return hi, fhi, true
	}
	if flo*fhi > 0 {
		bestT, bestE := lo, flo
		for i := 0; i <= 200; i++ {
			t := lo + (hi-lo)*float64(i)/200.0
			e := eAt(seed, ksum, t, residual)
			if math.Abs(e) < math.Abs(bestE) {
				bestT, bestE = t, e
			}
		}
		return bestT, bestE, false
	}
	for i := 0; i < 80; i++ {
		mid := 0.5 * (lo + hi)
		fm := eAt(seed, ksum, mid, residual)
		if flo*fm <= 0 {
			hi, fhi = mid, fm
		} else {
			lo, flo = mid, fm
		}
		_ = fhi
	}
	t := 0.5 * (lo + hi)
	return t, eAt(seed, ksum, t, residual), true
}

func f12(seed TransportSeed, t float64) float64 {
	g1, g2, _ := gaugeValues(seed, t)
	return g1 - g2
}

func u12(seed TransportSeed, t float64) float64 {
	g1, g2, _ := gaugeValues(seed, t)
	return 1.0/(g1*g1) - 1.0/(g2*g2)
}

func fitAffineRelative(y, x []float64) (slope, intercept, rel float64) {
	n := float64(len(y))
	var sx, sy, sxx, sxy, yy float64
	for i := range y {
		sx += x[i]
		sy += y[i]
		sxx += x[i] * x[i]
		sxy += x[i] * y[i]
		yy += y[i] * y[i]
	}
	denom := n*sxx - sx*sx
	if denom == 0 || yy == 0 {
		return 0, 0, math.Inf(1)
	}
	slope = (n*sxy - sx*sy) / denom
	intercept = (sy - slope*sx) / n
	var rss float64
	for i := range y {
		d := y[i] - (slope*x[i] + intercept)
		rss += d * d
	}
	return slope, intercept, math.Sqrt(rss / yy)
}

func meetingLog(gA, gB, bA, bB float64) float64 {
	return 8.0 * math.Pi * math.Pi * (1.0/(gA*gA) - 1.0/(gB*gB)) / (bA - bB)
}

func runCanonicalGauge(g0, b, t float64) float64 {
	inv := 1.0/(g0*g0) - b/(8.0*math.Pi*math.Pi)*t
	return 1.0 / math.Sqrt(inv)
}

func integrate(initial []float64, tEnd float64, n int) []float64 {
	if n < 1 {
		n = 1
	}
	y := append([]float64(nil), initial...)
	dt := tEnd / float64(n)
	for i := 0; i < n; i++ {
		k1 := historyDerivatives(y)
		k2 := historyDerivatives(addScaled(y, k1, dt/2.0))
		k3 := historyDerivatives(addScaled(y, k2, dt/2.0))
		k4 := historyDerivatives(addScaled(y, k3, dt))
		for j := range y {
			y[j] += dt * (k1[j] + 2*k2[j] + 2*k3[j] + k4[j]) / 6.0
		}
	}
	return y
}

func historyDerivatives(y []float64) []float64 {
	gY, g2, g3 := y[0], y[1], y[2]
	yu, yc, yt := y[3], y[4], y[5]
	yd, ys, yb := y[6], y[7], y[8]
	ye, ymu, ytau := y[9], y[10], y[11]
	lambda := y[12]
	loop := 16.0 * math.Pi * math.Pi
	T := 3.0*(yu*yu+yc*yc+yt*yt) + 3.0*(yd*yd+ys*ys+yb*yb) + ye*ye + ymu*ymu + ytau*ytau
	out := make([]float64, len(y))
	out[0] = (41.0 / 6.0) * gY * gY * gY / loop
	out[1] = (-19.0 / 6.0) * g2 * g2 * g2 / loop
	out[2] = (-7.0) * g3 * g3 * g3 / loop
	gaugeU := (17.0/12.0)*gY*gY + (9.0/4.0)*g2*g2 + 8.0*g3*g3
	gaugeD := (5.0/12.0)*gY*gY + (9.0/4.0)*g2*g2 + 8.0*g3*g3
	gaugeE := (15.0/4.0)*gY*gY + (9.0/4.0)*g2*g2
	out[3] = yu * (1.5*(yu*yu-yd*yd) + T - gaugeU) / loop
	out[4] = yc * (1.5*(yc*yc-ys*ys) + T - gaugeU) / loop
	out[5] = yt * (1.5*(yt*yt-yb*yb) + T - gaugeU) / loop
	out[6] = yd * (1.5*(yd*yd-yu*yu) + T - gaugeD) / loop
	out[7] = ys * (1.5*(ys*ys-yc*yc) + T - gaugeD) / loop
	out[8] = yb * (1.5*(yb*yb-yt*yt) + T - gaugeD) / loop
	out[9] = ye * (1.5*ye*ye + T - gaugeE) / loop
	out[10] = ymu * (1.5*ymu*ymu + T - gaugeE) / loop
	out[11] = ytau * (1.5*ytau*ytau + T - gaugeE) / loop
	out[12] = (24.0*lambda*lambda - 6.0*yt*yt*yt*yt + (3.0/8.0)*(2.0*math.Pow(g2, 4)+math.Pow(g2*g2+gY*gY, 2)) + lambda*(-9.0*g2*g2-3.0*gY*gY+12.0*yt*yt)) / loop
	return out
}

func addScaled(y, k []float64, scale float64) []float64 {
	out := make([]float64, len(y))
	for i := range y {
		out[i] = y[i] + scale*k[i]
	}
	return out
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate664DualRootInherited,
		StatusCommonRootAudited,
		StatusLocalFactorizationAudited,
		StatusGaugeCoordinateFamilyAudited,
		StatusAmplitudeCoordinateSupported,
		StatusCoordinateNaturalityUncertified,
		StatusBridgeCoordinateSeal,
		StatusInverseCoordinateFails,
		StatusNoNativeDualRootTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoFullUncertaintyPropagation,
		StatusNoBoundaryStressDerivation,
		StatusNoNativeTransportTheorem,
		StatusNoHiggsStabilityGaugeFlavor,
		StatusGate665Boundary,
	}
}
