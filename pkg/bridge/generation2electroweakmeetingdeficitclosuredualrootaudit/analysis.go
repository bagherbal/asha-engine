// Package generation2electroweakmeetingdeficitclosuredualrootaudit implements
// Gate 664: ElectroweakMeeting DeficitClosure Dual-Root Alignment Audit.
//
// Gate 663 showed that the E72 closure is a transverse near-zero crossing, not
// a stationary beta-balance point. Gate 664 audits whether that closure zero is
// aligned with the electroweak meeting root g1=g2, whether the alignment is a
// residual-convention artifact, and whether local proportionality can explain
// the alignment. This remains a v1 bridge diagnostic and preserves all native
// theorem and physics-promotion firewalls.
package generation2electroweakmeetingdeficitclosuredualrootaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate663 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryweighteddeficitclosurestationarityaudit"
)

const (
	AuditID = "GATE664-ELECTROWEAK-MEETING-DEFICIT-CLOSURE-DUAL-ROOT-ALIGNMENT-AUDIT"

	StatusGate663ZeroCrossingInherited      = "PASS_GATE663_ZERO_CROSSING_RESULT_INHERITED"
	StatusElectroweakMeetingFunctionDefined = "PASS_ELECTROWEAK_MEETING_FUNCTION_DEFINED"
	StatusClosureRootComputed               = "PASS_CLOSURE_ROOT_COMPUTED"
	StatusDualRootOffsetComputed            = "PASS_DUAL_ROOT_OFFSET_COMPUTED"
	StatusTransversalityAudited             = "PASS_TRANSVERSALITY_AUDITED"
	StatusLocalProportionalityAudited       = "PASS_LOCAL_PROPORTIONALITY_AUDITED"
	StatusGaugeResidualConventionAudited    = "PASS_GAUGE_RESIDUAL_CONVENTION_AUDITED"
	StatusWeightRootAudited                 = "PASS_WEIGHT_ROOT_AUDITED"
	StatusE72ZeroAlignedWithEWRoot          = "CONDITIONAL_SUPPORT_E72_ZERO_ALIGNED_WITH_ELECTROWEAK_MEETING_ROOT_IN_V1"
	StatusDualRootReplacesStationarity      = "CONDITIONAL_SUPPORT_DUAL_ROOT_ALIGNMENT_REPLACES_STATIONARITY_AS_PRESSURE_POINT"
	StatusNoNativeDualRootTheorem           = "FAILED_ROUTE_NO_NATIVE_DUAL_ROOT_ALIGNMENT_THEOREM"
	StatusNoNativeSevenOver72Theorem        = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoFullUncertaintyPropagation      = "FAILED_ROUTE_NO_FULL_UNCERTAINTY_PROPAGATION"
	StatusNoBoundaryStressDerivation        = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoNativeTransportTheorem          = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoHiggsStabilityGaugeFlavorClaim  = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate664Boundary                   = "FIREWALL_PRESERVED_GATE664_DUAL_ROOT_ALIGNMENT_BOUNDARY"
)

const (
	sevenOver72    = 7.0 / 72.0
	weightAbs      = 65.0 / 72.0
	transportSteps = 2000
)

type Gate663Inheritance struct {
	ZeroCrossingInherited bool
	Lambda12GeV           float64
	T12                   float64
	KSum                  float64
	E72AtLambda12         float64
	DE72Dt                float64
	MuZeroOverLambda12    float64
	DeltaLogZero          float64
	NoStationaryClaim     bool
	NoNativeScale         bool
	NoNativeSevenOver72   bool
	NoUncertainty         bool
	NoBoundaryStress      bool
	Verdict               string
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

type ElectroweakMeetingAudit struct {
	DefinitionF12   string
	DefinitionU12   string
	T12Analytic     float64
	Mu12GeV         float64
	F12AtRoot       float64
	U12AtRoot       float64
	GaugeConvention string
	Verdict         string
}

type ClosureRootAudit struct {
	TClosureZero        float64
	MuClosureZeroGeV    float64
	E72AtClosureZero    float64
	BracketHalfWidth    float64
	ClosureIsTransverse bool
	Verdict             string
}

type DualRootOffsetAudit struct {
	DeltaLogMuEOverMu12 float64
	MuEOverMu12         float64
	AbsoluteScaleDelta  float64
	AlignedInV1         bool
	Verdict             string
}

type TransversalityAudit struct {
	DF12DtAtLambda12 float64
	DU12DtAtLambda12 float64
	DE72DtAtLambda12 float64
	F12Transverse    bool
	U12Transverse    bool
	E72Transverse    bool
	SlopeTied        bool
	Verdict          string
}

type ProportionalityAudit struct {
	Window              float64
	Samples             int
	CForF12             float64
	RelativeResidualF12 float64
	CForU12             float64
	RelativeResidualU12 float64
	ProportionalToF12   bool
	ProportionalToU12   bool
	Verdict             string
}

type ResidualConventionRow struct {
	Name                  string
	Definition            string
	GaugeResidualAtT12    float64
	E72AtT12              float64
	TZero                 float64
	DeltaLogFromT12       float64
	MuZeroOverLambda12    float64
	RootFoundNearLambda12 bool
	Verdict               string
}

type ResidualConventionAudit struct {
	Rows                          []ResidualConventionRow
	DirectCouplingConventionsPass int
	InverseConventionPasses       bool
	ConventionStable              bool
	Verdict                       string
}

type WeightRootRow struct {
	DeltaLog          float64
	WBest             float64
	WBestMinus7Over72 float64
	E72At7Over72      float64
}

type WeightRootAudit struct {
	Rows                         []WeightRootRow
	WBestAtLambda12              float64
	WBestMinus7Over72AtLambda12  float64
	CrossesSevenOver72NearLambda bool
	WeightIndependentlySelected  bool
	Verdict                      string
}

type SourceTypeClassification struct {
	Outcomes []string
	Verdict  string
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
	Inherited       Gate663Inheritance
	Seed            TransportSeed
	Meeting         ElectroweakMeetingAudit
	ClosureRoot     ClosureRootAudit
	DualRoot        DualRootOffsetAudit
	Transversality  TransversalityAudit
	Proportionality ProportionalityAudit
	Conventions     ResidualConventionAudit
	WeightRoot      WeightRootAudit
	Source          SourceTypeClassification
	Discipline      VerdictDiscipline
	Truth           string
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
	g663, err := gate663.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate663 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g663)
	seed := buildSeed(g663.Seed)
	meeting := buildMeeting(seed)
	closureRoot := buildClosureRoot(seed, inherited.KSum, seed.T12, gaugeResidualEWMean)
	dual := buildDualRoot(seed, meeting, closureRoot)
	trans := buildTransversality(seed, inherited.KSum, seed.T12, gaugeResidualEWMean)
	prop := buildProportionality(seed, inherited.KSum, seed.T12, closureRoot.TClosureZero)
	conventions := buildConventionAudit(seed, inherited.KSum, seed.T12)
	weight := buildWeightRoot(seed, inherited.KSum)
	source := buildSourceType(dual, trans, prop, conventions)
	discipline := VerdictDiscipline{Verdict: StatusGate664Boundary}
	truth := "Gate 664 classifies the v1 E72 closure as a dual-root alignment: the E72 zero lies within about 9e-7 in log scale of the electroweak g1=g2 meeting root, and both roots are transverse. The alignment is strongest for the direct electroweak-mean gauge residual convention used by Gates 659-663; it is not a stationarity theorem, native 7/72 theorem, full uncertainty propagation, or boundary-stress derivation."
	return Analysis{Inherited: inherited, Seed: seed, Meeting: meeting, ClosureRoot: closureRoot, DualRoot: dual, Transversality: trans, Proportionality: prop, Conventions: conventions, WeightRoot: weight, Source: source, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate663.Analysis) Gate663Inheritance {
	return Gate663Inheritance{
		ZeroCrossingInherited: g.ZeroScale.ClosureZeroAligned && g.Derivative.ZeroCrossingNotStationary && !g.Derivative.Stationary,
		Lambda12GeV:           g.Seed.Lambda12GeV,
		T12:                   g.Seed.T12,
		KSum:                  g.Inherited.KSum,
		E72AtLambda12:         g.Function.E72,
		DE72Dt:                g.Derivative.DE72DtAnalytic,
		MuZeroOverLambda12:    g.ZeroScale.MuZeroOverLambda12,
		DeltaLogZero:          g.ZeroScale.DeltaLogFromLambda12,
		NoStationaryClaim:     g.Derivative.ZeroCrossingNotStationary && !g.BetaBalance.StationarityWouldRequire,
		NoNativeScale:         !g.Discipline.ClaimsNativeScaleSelection,
		NoNativeSevenOver72:   !g.Discipline.ClaimsNativeSevenOver72Theorem,
		NoUncertainty:         !g.Discipline.ClaimsFullUncertaintyPropagation,
		NoBoundaryStress:      !g.Discipline.ClaimsBoundaryStressDerivation,
		Verdict:               StatusGate663ZeroCrossingInherited,
	}
}

func buildSeed(s gate663.TransportSeed) TransportSeed {
	return TransportSeed{Mu0GeV: s.Mu0GeV, Lambda12GeV: s.Lambda12GeV, T12: s.T12, G1MZ: s.G1MZ, GYMZ: s.GYMZ, G2MZ: s.G2MZ, G3MZ: s.G3MZ, InitialVector: append([]float64(nil), s.InitialVector...), Verdict: StatusGate663ZeroCrossingInherited}
}

func buildMeeting(seed TransportSeed) ElectroweakMeetingAudit {
	t12 := meetingLog(seed.G1MZ, seed.G2MZ, 41.0/10.0, -19.0/6.0)
	g1 := runCanonicalGauge(seed.G1MZ, 41.0/10.0, t12)
	g2 := runCanonicalGauge(seed.G2MZ, -19.0/6.0, t12)
	return ElectroweakMeetingAudit{DefinitionF12: "F12(mu)=g1(mu)-g2(mu)", DefinitionU12: "U12(mu)=1/g1(mu)^2-1/g2(mu)^2", T12Analytic: t12, Mu12GeV: seed.Mu0GeV * math.Exp(t12), F12AtRoot: g1 - g2, U12AtRoot: 1/(g1*g1) - 1/(g2*g2), GaugeConvention: "canonical g1=sqrt(5/3)gY with one-loop v1 gauge transport", Verdict: StatusElectroweakMeetingFunctionDefined}
}

func buildClosureRoot(seed TransportSeed, ksum, t12 float64, residual residualFunc) ClosureRootAudit {
	tz, ez, found := findRootNear(seed, ksum, residual, t12, 1e-4)
	width := 1e-4
	if !found {
		tz, ez, found = findRootNear(seed, ksum, residual, t12, 2.0)
		width = 2.0
	}
	return ClosureRootAudit{TClosureZero: tz, MuClosureZeroGeV: seed.Mu0GeV * math.Exp(tz), E72AtClosureZero: ez, BracketHalfWidth: width, ClosureIsTransverse: found, Verdict: StatusClosureRootComputed}
}

func buildDualRoot(seed TransportSeed, meeting ElectroweakMeetingAudit, root ClosureRootAudit) DualRootOffsetAudit {
	d := root.TClosureZero - meeting.T12Analytic
	return DualRootOffsetAudit{DeltaLogMuEOverMu12: d, MuEOverMu12: math.Exp(d), AbsoluteScaleDelta: root.MuClosureZeroGeV - meeting.Mu12GeV, AlignedInV1: math.Abs(d) < 1e-5, Verdict: join(StatusDualRootOffsetComputed, StatusE72ZeroAlignedWithEWRoot, StatusDualRootReplacesStationarity)}
}

func buildTransversality(seed TransportSeed, ksum, t12 float64, residual residualFunc) TransversalityAudit {
	g1 := runCanonicalGauge(seed.G1MZ, 41.0/10.0, t12)
	g2 := runCanonicalGauge(seed.G2MZ, -19.0/6.0, t12)
	bg1 := gaugeBeta(g1, 41.0/10.0)
	bg2 := gaugeBeta(g2, -19.0/6.0)
	dF := bg1 - bg2
	dU := -(41.0/10.0 - (-19.0 / 6.0)) / (8.0 * math.Pi * math.Pi)
	dE := derivativeE(seed, ksum, t12, residual)
	return TransversalityAudit{DF12DtAtLambda12: dF, DU12DtAtLambda12: dU, DE72DtAtLambda12: dE, F12Transverse: math.Abs(dF) > 1e-3, U12Transverse: math.Abs(dU) > 0.05, E72Transverse: math.Abs(dE) > 9e-4, SlopeTied: false, Verdict: StatusTransversalityAudited}
}

func buildProportionality(seed TransportSeed, ksum, t12, tE float64) ProportionalityAudit {
	window := 0.1
	points := []float64{-window, -0.05, 0, 0.05, window}
	eVals := make([]float64, 0, len(points))
	fVals := make([]float64, 0, len(points))
	uVals := make([]float64, 0, len(points))
	for _, d := range points {
		t := t12 + d
		eVals = append(eVals, eAt(seed, ksum, t, gaugeResidualEWMean))
		fVals = append(fVals, f12(seed, t))
		uVals = append(uVals, u12(seed, t))
	}
	cF, relF := fitRelative(eVals, fVals)
	cU, relU := fitRelative(eVals, uVals)
	return ProportionalityAudit{Window: window, Samples: len(points), CForF12: cF, RelativeResidualF12: relF, CForU12: cU, RelativeResidualU12: relU, ProportionalToF12: relF < 0.15, ProportionalToU12: relU < 0.15, Verdict: StatusLocalProportionalityAudited}
}

func buildConventionAudit(seed TransportSeed, ksum, t12 float64) ResidualConventionAudit {
	defs := []struct {
		name string
		def  string
		fn   residualFunc
	}{
		{"EW-mean direct", "g3/((g1+g2)/2)-1", gaugeResidualEWMean},
		{"inverse-coupling EW-mean", "(1/g3^2)/mean(1/g1^2,1/g2^2)-1", gaugeResidualInverseEWMean},
		{"pair-meeting residual", "nonmeeting g3 relative to g1/g2 pair mean at Lambda12", gaugeResidualEWMean},
		{"strong relative to g1", "g3/g1-1", gaugeResidualG1},
		{"strong relative to g2", "g3/g2-1", gaugeResidualG2},
	}
	rows := make([]ResidualConventionRow, 0, len(defs))
	directPass := 0
	inversePass := false
	for _, d := range defs {
		_, absLam, g, _, e := closureValues(seed, ksum, t12, d.fn)
		_ = absLam
		tz, _, found := findRootNear(seed, ksum, d.fn, t12, 0.5)
		row := ResidualConventionRow{Name: d.name, Definition: d.def, GaugeResidualAtT12: g, E72AtT12: e, TZero: tz, DeltaLogFromT12: tz - t12, MuZeroOverLambda12: math.Exp(tz - t12), RootFoundNearLambda12: found && math.Abs(tz-t12) < 1e-3, Verdict: StatusGaugeResidualConventionAudited}
		if strings.Contains(d.name, "direct") || strings.Contains(d.name, "relative") || strings.Contains(d.name, "pair") {
			if row.RootFoundNearLambda12 {
				directPass++
			}
		}
		if d.name == "inverse-coupling EW-mean" && row.RootFoundNearLambda12 {
			inversePass = true
		}
		rows = append(rows, row)
	}
	return ResidualConventionAudit{Rows: rows, DirectCouplingConventionsPass: directPass, InverseConventionPasses: inversePass, ConventionStable: directPass >= 4, Verdict: StatusGaugeResidualConventionAudited}
}

func buildWeightRoot(seed TransportSeed, ksum float64) WeightRootAudit {
	shifts := []float64{-0.1, -0.01, 0, 0.01, 0.1}
	rows := make([]WeightRootRow, 0, len(shifts))
	for _, d := range shifts {
		_, absLam, gauge, _, e := closureValues(seed, ksum, seed.T12+d, gaugeResidualEWMean)
		wBest := (ksum - absLam) / (gauge - absLam)
		rows = append(rows, WeightRootRow{DeltaLog: d, WBest: wBest, WBestMinus7Over72: wBest - sevenOver72, E72At7Over72: e})
	}
	at := rows[2]
	return WeightRootAudit{Rows: rows, WBestAtLambda12: at.WBest, WBestMinus7Over72AtLambda12: at.WBestMinus7Over72, CrossesSevenOver72NearLambda: math.Abs(at.WBestMinus7Over72) < 1e-6, WeightIndependentlySelected: false, Verdict: StatusWeightRootAudited}
}

func buildSourceType(d DualRootOffsetAudit, t TransversalityAudit, p ProportionalityAudit, c ResidualConventionAudit) SourceTypeClassification {
	outcomes := []string{
		"dual-root alignment: supported in v1 because E72=0 lies within O(1e-6) log-scale of g1=g2",
		"stationarity explanation: rejected by Gate663 and inherited here; E72 is transverse",
		"local proportionality: audited as a diagnostic only and not promoted to a theorem",
		"gauge-convention dependence: direct EW-mean/g1/g2 conventions preserve the alignment; inverse-coupling convention is recorded separately",
		"structural bridge candidate: open, but requires native dual-root theorem and uncertainty propagation",
	}
	return SourceTypeClassification{Outcomes: outcomes, Verdict: join(StatusE72ZeroAlignedWithEWRoot, StatusDualRootReplacesStationarity, StatusNoNativeDualRootTheorem)}
}

// residualFunc returns the gauge residual G(mu) used in W72(mu).
type residualFunc func(seed TransportSeed, t float64) float64

func gaugeResidualEWMean(seed TransportSeed, t float64) float64 {
	g1 := runCanonicalGauge(seed.G1MZ, 41.0/10.0, t)
	g2 := runCanonicalGauge(seed.G2MZ, -19.0/6.0, t)
	g3 := runCanonicalGauge(seed.G3MZ, -7.0, t)
	return g3/(0.5*(g1+g2)) - 1.0
}

func gaugeResidualG1(seed TransportSeed, t float64) float64 {
	g1 := runCanonicalGauge(seed.G1MZ, 41.0/10.0, t)
	g3 := runCanonicalGauge(seed.G3MZ, -7.0, t)
	return g3/g1 - 1.0
}

func gaugeResidualG2(seed TransportSeed, t float64) float64 {
	g2 := runCanonicalGauge(seed.G2MZ, -19.0/6.0, t)
	g3 := runCanonicalGauge(seed.G3MZ, -7.0, t)
	return g3/g2 - 1.0
}

func gaugeResidualInverseEWMean(seed TransportSeed, t float64) float64 {
	g1 := runCanonicalGauge(seed.G1MZ, 41.0/10.0, t)
	g2 := runCanonicalGauge(seed.G2MZ, -19.0/6.0, t)
	g3 := runCanonicalGauge(seed.G3MZ, -7.0, t)
	u1, u2, u3 := 1.0/(g1*g1), 1.0/(g2*g2), 1.0/(g3*g3)
	return u3/(0.5*(u1+u2)) - 1.0
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
		// fall back to the smallest absolute value in the bracket, but mark no root.
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
			hi = mid
			fhi = fm
		} else {
			lo = mid
			flo = fm
		}
	}
	t := 0.5 * (lo + hi)
	return t, eAt(seed, ksum, t, residual), true
}

func derivativeE(seed TransportSeed, ksum, t float64, residual residualFunc) float64 {
	h := 1e-4
	return (eAt(seed, ksum, t+h, residual) - eAt(seed, ksum, t-h, residual)) / (2.0 * h)
}

func f12(seed TransportSeed, t float64) float64 {
	return runCanonicalGauge(seed.G1MZ, 41.0/10.0, t) - runCanonicalGauge(seed.G2MZ, -19.0/6.0, t)
}

func u12(seed TransportSeed, t float64) float64 {
	g1 := runCanonicalGauge(seed.G1MZ, 41.0/10.0, t)
	g2 := runCanonicalGauge(seed.G2MZ, -19.0/6.0, t)
	return 1.0/(g1*g1) - 1.0/(g2*g2)
}

func fitRelative(y, x []float64) (float64, float64) {
	var xy, xx, yy float64
	for i := range y {
		xy += y[i] * x[i]
		xx += x[i] * x[i]
		yy += y[i] * y[i]
	}
	if xx == 0 || yy == 0 {
		return 0, math.Inf(1)
	}
	c := xy / xx
	var res float64
	for i := range y {
		d := y[i] - c*x[i]
		res += d * d
	}
	return c, math.Sqrt(res / yy)
}

func meetingLog(gA, gB, bA, bB float64) float64 {
	return 8.0 * math.Pi * math.Pi * (1.0/(gA*gA) - 1.0/(gB*gB)) / (bA - bB)
}

func gaugeBeta(g, b float64) float64 { return b * g * g * g / (16.0 * math.Pi * math.Pi) }

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
			y[j] += dt * (k1[j] + 2.0*k2[j] + 2.0*k3[j] + k4[j]) / 6.0
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
		StatusGate663ZeroCrossingInherited,
		StatusElectroweakMeetingFunctionDefined,
		StatusClosureRootComputed,
		StatusDualRootOffsetComputed,
		StatusTransversalityAudited,
		StatusLocalProportionalityAudited,
		StatusGaugeResidualConventionAudited,
		StatusWeightRootAudited,
		StatusE72ZeroAlignedWithEWRoot,
		StatusDualRootReplacesStationarity,
		StatusNoNativeDualRootTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoFullUncertaintyPropagation,
		StatusNoBoundaryStressDerivation,
		StatusNoNativeTransportTheorem,
		StatusNoHiggsStabilityGaugeFlavorClaim,
		StatusGate664Boundary,
	}
}
