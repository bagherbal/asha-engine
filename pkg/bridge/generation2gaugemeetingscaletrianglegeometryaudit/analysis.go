// Package generation2gaugemeetingscaletrianglegeometryaudit implements
// Gate 608: Gauge Meeting-Scale Triangle Geometry Audit.
//
// Gate 607 quantified the strong-coupling residual at Lambda_12 and computed
// the one-loop pairwise meeting scales Lambda_12, Lambda_13, and Lambda_23.
// Gate 608 treats those three scales as a transport-ledger triangle: it
// computes the log geometry, classifies boundary-scale choices, defines beta
// deformation diagnostics, records threshold-origin slots, and preserves the
// no-unification / no-threshold-existence firewalls.
package generation2gaugemeetingscaletrianglegeometryaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2strongcouplingthresholdresidualledgeraudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE608-GAUGE-MEETING-SCALE-TRIANGLE-GEOMETRY-AUDIT"

	StatusGate607Inherited             = "PASS_GATE607_MEETING_SCALE_TRIANGLE_INHERITED"
	StatusLogTriangleComputed          = "PASS_LOG_TRIANGLE_GEOMETRY_COMPUTED"
	StatusTriangleAsymmetryAudited     = "PASS_TRIANGLE_ASYMMETRY_AUDITED"
	StatusBoundaryChoiceClassified     = "PASS_BOUNDARY_CHOICE_RESIDUALS_CLASSIFIED"
	StatusBetaDeformationVectorAudited = "PASS_BETA_DEFORMATION_VECTOR_AUDIT_DEFINED"
	StatusThresholdOriginSlotsDefined  = "PASS_THRESHOLD_ORIGIN_SLOTS_CLASSIFIED"
	StatusScalarRelationClassified     = "PASS_SCALAR_ZERO_CROSSING_RELATIVE_TO_GAUGE_TRIANGLE_CLASSIFIED"
	StatusStructuredTransportLedger    = "CONDITIONAL_SUPPORT_MEETING_TRIANGLE_IS_STRUCTURED_TRANSPORT_LEDGER"
	StatusBalancedLambdaCandidate      = "CONDITIONAL_SUPPORT_GEOMETRIC_MEAN_IS_BALANCED_LOG_TRIANGLE_DIAGNOSTIC_ONLY"
	StatusMinimalNormDiagnostic        = "CONDITIONAL_SUPPORT_MINIMAL_NORM_BETA_DEFORMATION_DIAGNOSTIC_COMPUTED"
	StatusNoSingleOneLoopUnification   = "FAILED_ROUTE_NO_SINGLE_ONE_LOOP_UNIFICATION_POINT"
	StatusNoNativeThresholdTheorem     = "FAILED_ROUTE_NO_NATIVE_THRESHOLD_OR_EXTRA_FIELD_THEOREM"
	StatusNoNativeLambdaUSelection     = "FAILED_ROUTE_NO_NATIVE_LAMBDA_U_SELECTION_THEOREM"
	StatusNoFullGaugeUnificationClaim  = "FAILED_ROUTE_NO_FULL_GAUGE_UNIFICATION_CLAIM"
	StatusNoThresholdFit               = "FAILED_ROUTE_THRESHOLD_ORIGIN_SLOTS_NOT_FITTED"
	StatusNoNewFields                  = "FAILED_ROUTE_NO_NEW_FIELD_CONTENT_INTRODUCED"
	StatusNoScalarGaugeClosure         = "FAILED_ROUTE_SCALAR_ZERO_CROSSING_NOT_USED_TO_CLOSE_GAUGE_TRIANGLE"
	StatusGate608Boundary              = "FIREWALL_PRESERVED_GATE608_MEETING_TRIANGLE_BOUNDARY"
	StatusNoEndpointDerivation         = "FIREWALL_PRESERVED_NO_OBSERVED_ENDPOINT_DERIVATION"
	StatusNoUnificationFirewall        = "FIREWALL_PRESERVED_NO_GAUGE_UNIFICATION_DERIVATION"
)

const (
	b1SM = 41.0 / 10.0
	b2SM = -19.0 / 6.0
	b3SM = -7.0
)

type InheritedGate607 struct {
	Lambda12GeV             float64
	Lambda13GeV             float64
	Lambda23GeV             float64
	GStar                   float64
	G3Lambda12              float64
	Delta3ThresholdRequired float64
	DeltaAlpha3InvRequired  float64
	DeltaB3Required         float64
	Verdict                 string
}

type PairwiseMeetingScaleRow struct {
	Pair              string
	ScaleGeV          float64
	Log10Scale        float64
	LogMuOverMZ       float64
	CouplingAtMeeting float64
	ExactPair         string
	Verdict           string
}

type LogTriangleGeometry struct {
	Ratio13Over12      float64
	Ratio23Over13      float64
	Ratio23Over12      float64
	Log10Ratio13Over12 float64
	Log10Ratio23Over13 float64
	Log10Ratio23Over12 float64
	SpreadDecades      float64
	GeometricMeanGeV   float64
	Log10GeometricMean float64
	DistancesFromMean  map[string]float64
	SkewStatement      string
	Verdict            string
}

type BoundaryChoiceResidualRow struct {
	ChoiceScale     string
	ScaleGeV        float64
	ExactPair       string
	G1              float64
	G2              float64
	G3              float64
	U1              float64
	U2              float64
	U3              float64
	ResidualSummary string
	MaxDeltaU       float64
	Verdict         string
}

type BetaDeformationVectorRow struct {
	Strategy       string
	LambdaUGeV     float64
	TargetU        float64
	DeltaB1        float64
	DeltaB2        float64
	DeltaB3        float64
	Norm           float64
	Formula        string
	Interpretation string
	Verdict        string
}

type ThresholdOriginSlotRow struct {
	Candidate     string
	Kind          string
	WouldMove     string
	CurrentStatus string
	NativeStatus  string
	Verdict       string
}

type ASHANativeStatus struct {
	ProvidesNativeThresholdSpectrum bool
	ProvidesBColoredDeformation     bool
	ProvidesFiniteAlgebraExtension  bool
	ProvidesBoundaryColorCorrection bool
	ProvidesNativeLambdaUSelection  bool
	ClaimsUnification               bool
	Statement                       string
	Verdict                         string
}

type ScalarTriangleRelation struct {
	ZeroCrossingGeV float64
	Lambda12GeV     float64
	Lambda13GeV     float64
	Lambda23GeV     float64
	Statement       string
	Verdict         string
}

type Firewalls struct {
	ClaimsUnification   bool
	IntroducesNewFields bool
	FitsThresholds      bool
	PromotesLambdaU     bool
	DerivesEndpoint     bool
	UsesScalarToClose   bool
	Verdict             string
}

type Analysis struct {
	Inherited            InheritedGate607
	PairwiseScales       []PairwiseMeetingScaleRow
	LogGeometry          LogTriangleGeometry
	BoundaryChoices      []BoundaryChoiceResidualRow
	BetaDeformations     []BetaDeformationVectorRow
	ThresholdOriginSlots []ThresholdOriginSlotRow
	NativeStatus         ASHANativeStatus
	ScalarRelation       ScalarTriangleRelation
	Firewalls            Firewalls
	Truth                string
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
	g607, err := generation2strongcouplingthresholdresidualledgeraudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate607 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport bundle: %w", err)
	}
	pairwise := buildPairwiseScales(g607)
	logGeom := buildLogGeometry(pairwise)
	choices := buildBoundaryChoices(bundle, pairwise, logGeom)
	beta := buildBetaDeformations(bundle, pairwise, logGeom)
	slots := buildThresholdOriginSlots()
	native := buildNativeStatus()
	scalar := buildScalarRelation(bundle, pairwise)
	firewalls := auditFirewalls()
	truth := "Gate 608 treats the one-loop pairwise gauge meeting scales as a log-scale transport triangle. The triangle is a structured ledger of non-closure, not a unification theorem: Lambda_12, Lambda_13, and Lambda_23 are separated by about 2.93 decades, boundary choices leave different residual couplings, and beta/threshold deformations are diagnostic slots only."
	return Analysis{inherit(g607), pairwise, logGeom, choices, beta, slots, native, scalar, firewalls, truth}, nil
}

func inherit(a generation2strongcouplingthresholdresidualledgeraudit.Analysis) InheritedGate607 {
	return InheritedGate607{
		Lambda12GeV:             meetingScale(a.MeetingScales, "Lambda_12"),
		Lambda13GeV:             meetingScale(a.MeetingScales, "Lambda_13"),
		Lambda23GeV:             meetingScale(a.MeetingScales, "Lambda_23"),
		GStar:                   a.Inherited.GStar,
		G3Lambda12:              a.Inherited.G3Lambda,
		Delta3ThresholdRequired: a.ThresholdSlots[0].RequiredValue,
		DeltaAlpha3InvRequired:  a.ThresholdSlots[1].RequiredValue,
		DeltaB3Required:         a.BetaDeformation.DeltaB3Required,
		Verdict:                 StatusGate607Inherited,
	}
}

func buildPairwiseScales(a generation2strongcouplingthresholdresidualledgeraudit.Analysis) []PairwiseMeetingScaleRow {
	rows := make([]PairwiseMeetingScaleRow, 0, len(a.MeetingScales))
	for _, m := range a.MeetingScales {
		exact := ""
		switch m.Pair {
		case "Lambda_12":
			exact = "g1=g2"
		case "Lambda_13":
			exact = "g1=g3"
		case "Lambda_23":
			exact = "g2=g3"
		}
		rows = append(rows, PairwiseMeetingScaleRow{m.Pair, m.ScaleGeV, math.Log10(m.ScaleGeV), m.LogMuOverMZ, m.CouplingAtMeeting, exact, StatusGate607Inherited})
	}
	return rows
}

func buildLogGeometry(rows []PairwiseMeetingScaleRow) LogTriangleGeometry {
	l12 := scale(rows, "Lambda_12")
	l13 := scale(rows, "Lambda_13")
	l23 := scale(rows, "Lambda_23")
	geom := math.Pow(l12*l13*l23, 1.0/3.0)
	d := map[string]float64{
		"Lambda_12": math.Log10(l12 / geom),
		"Lambda_13": math.Log10(l13 / geom),
		"Lambda_23": math.Log10(l23 / geom),
	}
	return LogTriangleGeometry{
		Ratio13Over12:      l13 / l12,
		Ratio23Over13:      l23 / l13,
		Ratio23Over12:      l23 / l12,
		Log10Ratio13Over12: math.Log10(l13 / l12),
		Log10Ratio23Over13: math.Log10(l23 / l13),
		Log10Ratio23Over12: math.Log10(l23 / l12),
		SpreadDecades:      math.Log10(l23 / l12),
		GeometricMeanGeV:   geom,
		Log10GeometricMean: math.Log10(geom),
		DistancesFromMean:  d,
		SkewStatement:      "The triangle is strongly skewed: Lambda_13 sits much closer to Lambda_12 than to Lambda_23 in log space, so the pairwise crossings do not form a symmetric one-loop unification point.",
		Verdict:            StatusLogTriangleComputed,
	}
}

func buildBoundaryChoices(b historytransport.Bundle, rows []PairwiseMeetingScaleRow, geom LogTriangleGeometry) []BoundaryChoiceResidualRow {
	mu0 := b.GaugeBoundary.Mu0GeV
	choices := []struct {
		name  string
		scale float64
		pair  string
		stat  string
	}{
		{"Lambda_12", scale(rows, "Lambda_12"), "g1=g2", StatusBoundaryChoiceClassified},
		{"Lambda_13", scale(rows, "Lambda_13"), "g1=g3", StatusBoundaryChoiceClassified},
		{"Lambda_23", scale(rows, "Lambda_23"), "g2=g3", StatusBoundaryChoiceClassified},
		{"Lambda_geom", geom.GeometricMeanGeV, "none exact; log-balanced diagnostic", StatusBalancedLambdaCandidate},
	}
	out := make([]BoundaryChoiceResidualRow, 0, len(choices))
	for _, c := range choices {
		t := math.Log(c.scale / mu0)
		u1 := runU(b.EndVector.G1, b1SM, t)
		u2 := runU(b.EndVector.G2, b2SM, t)
		u3 := runU(b.EndVector.G3, b3SM, t)
		g1 := 1 / math.Sqrt(u1)
		g2 := 1 / math.Sqrt(u2)
		g3 := 1 / math.Sqrt(u3)
		maxdu := maxAbsPairDelta(u1, u2, u3)
		summary := residualSummary(c.name, u1, u2, u3)
		out = append(out, BoundaryChoiceResidualRow{c.name, c.scale, c.pair, g1, g2, g3, u1, u2, u3, summary, maxdu, c.stat})
	}
	return out
}

func runU(g0, b, t float64) float64 { return 1/(g0*g0) - b*t/(8*math.Pi*math.Pi) }

func residualSummary(name string, u1, u2, u3 float64) string {
	switch name {
	case "Lambda_12":
		return fmt.Sprintf("g1=g2 exact; u3-u*=%.15g", u3-u1)
	case "Lambda_13":
		return fmt.Sprintf("g1=g3 exact; u2-u13=%.15g", u2-u1)
	case "Lambda_23":
		return fmt.Sprintf("g2=g3 exact; u1-u23=%.15g", u1-u2)
	default:
		mean := (u1 + u2 + u3) / 3
		return fmt.Sprintf("no pair exact; u deviations from mean: [%.15g, %.15g, %.15g]", u1-mean, u2-mean, u3-mean)
	}
}

func buildBetaDeformations(b historytransport.Bundle, rows []PairwiseMeetingScaleRow, geom LogTriangleGeometry) []BetaDeformationVectorRow {
	lambda12 := scale(rows, "Lambda_12")
	t12 := math.Log(lambda12 / b.GaugeBoundary.Mu0GeV)
	uStar := runU(b.EndVector.G1, b1SM, t12)
	u3 := runU(b.EndVector.G3, b3SM, t12)
	delta := uStar - u3
	db3 := -8 * math.Pi * math.Pi * delta / t12
	db12 := 8 * math.Pi * math.Pi * delta / t12
	geomRow := minimalNormDeformationAt(b, geom.GeometricMeanGeV)
	return []BetaDeformationVectorRow{
		{"hold b1,b2 fixed at Lambda_12; deform b3 only", lambda12, uStar, 0, 0, db3, math.Abs(db3), "Delta b3=-8*pi^2*(u_star-u3)/t12", "diagnostic strong-sector deformation; reproduces Gate607 13.3% |b3| size", StatusBetaDeformationVectorAudited},
		{"hold b3 fixed at Lambda_12; deform b1,b2 to meet g3", lambda12, u3, db12, db12, 0, math.Sqrt(2) * math.Abs(db12), "Delta b1=Delta b2=+8*pi^2*(u_star-u3)/t12", "diagnostic electroweak-sector deformation; not a proposed boundary law", StatusBetaDeformationVectorAudited},
		geomRow,
	}
}

func minimalNormDeformationAt(b historytransport.Bundle, lambdaU float64) BetaDeformationVectorRow {
	t := math.Log(lambdaU / b.GaugeBoundary.Mu0GeV)
	u1 := runU(b.EndVector.G1, b1SM, t)
	u2 := runU(b.EndVector.G2, b2SM, t)
	u3 := runU(b.EndVector.G3, b3SM, t)
	target := (u1 + u2 + u3) / 3
	coef := 8 * math.Pi * math.Pi / t
	db1 := coef * (u1 - target)
	db2 := coef * (u2 - target)
	db3 := coef * (u3 - target)
	norm := math.Sqrt(db1*db1 + db2*db2 + db3*db3)
	return BetaDeformationVectorRow{"minimal ||Delta b|| at Lambda_geom", lambdaU, target, db1, db2, db3, norm, "Delta b_i=(8*pi^2/t)*(u_i(lambda_U)-mean_j u_j(lambda_U))", "balanced deformation diagnostic at log-geometric meeting scale; no native Lambda_U selection theorem", StatusMinimalNormDiagnostic}
}

func buildThresholdOriginSlots() []ThresholdOriginSlotRow {
	return []ThresholdOriginSlotRow{
		{"two-loop SM RG", "higher-loop transport", "curvature of gauge running", "not included in v1", "bridge calculation slot", StatusThresholdOriginSlotsDefined},
		{"low-energy matching", "scheme / pole-MSbar threshold", "endpoint initial couplings", "not included in v1", "observed matching ledger", StatusThresholdOriginSlotsDefined},
		{"heavy threshold near Lambda_U", "high-scale threshold", "one or more inverse-coupling jumps", "symbolic slot only", "no native spectrum", StatusNoNativeThresholdTheorem},
		{"finite spectral-action boundary correction", "boundary matching", "initial inverse-coupling offsets", "not constructed", "no native correction theorem", StatusNoNativeThresholdTheorem},
		{"extra colored states", "field-content deformation", "mostly b3", "not introduced", "no new fields supplied", StatusNoNewFields},
		{"extra colorless electroweak states", "field-content deformation", "mostly b1/b2", "not introduced", "no new fields supplied", StatusNoNewFields},
		{"renormalization scheme dependence", "transport convention", "all couplings depending on scheme", "explicitly labeled", "bridge ledger only", StatusThresholdOriginSlotsDefined},
	}
}

func buildNativeStatus() ASHANativeStatus {
	return ASHANativeStatus{
		ProvidesNativeThresholdSpectrum: false,
		ProvidesBColoredDeformation:     false,
		ProvidesFiniteAlgebraExtension:  false,
		ProvidesBoundaryColorCorrection: false,
		ProvidesNativeLambdaUSelection:  false,
		ClaimsUnification:               false,
		Statement:                       "Current ASHA data give the native electroweak normalization and expose the gauge meeting triangle, but provide no threshold spectrum, no extra field theorem, no boundary color correction, and no native reason to select Lambda_12, Lambda_13, Lambda_23, or Lambda_geom as a physical unification scale.",
		Verdict:                         StatusNoNativeLambdaUSelection,
	}
}

func buildScalarRelation(b historytransport.Bundle, rows []PairwiseMeetingScaleRow) ScalarTriangleRelation {
	zero := math.NaN()
	if b.ScalarTransport.ZeroCrossingScaleGeV != nil {
		zero = *b.ScalarTransport.ZeroCrossingScaleGeV
	}
	return ScalarTriangleRelation{
		ZeroCrossingGeV: zero,
		Lambda12GeV:     scale(rows, "Lambda_12"),
		Lambda13GeV:     scale(rows, "Lambda_13"),
		Lambda23GeV:     scale(rows, "Lambda_23"),
		Statement:       "The v1 scalar zero crossing lies far below all three one-loop gauge meeting scales. Gate 608 records this relative position but does not combine scalar stability with gauge-triangle closure.",
		Verdict:         StatusScalarRelationClassified,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, StatusGate608Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate607Inherited,
		StatusLogTriangleComputed,
		StatusTriangleAsymmetryAudited,
		StatusBoundaryChoiceClassified,
		StatusBetaDeformationVectorAudited,
		StatusThresholdOriginSlotsDefined,
		StatusScalarRelationClassified,
		StatusStructuredTransportLedger,
		StatusBalancedLambdaCandidate,
		StatusMinimalNormDiagnostic,
		StatusNoSingleOneLoopUnification,
		StatusNoNativeThresholdTheorem,
		StatusNoNativeLambdaUSelection,
		StatusNoFullGaugeUnificationClaim,
		StatusNoThresholdFit,
		StatusNoNewFields,
		StatusNoScalarGaugeClosure,
		StatusGate608Boundary,
		StatusNoEndpointDerivation,
		StatusNoUnificationFirewall,
	}
}

func meetingScale(rows []generation2strongcouplingthresholdresidualledgeraudit.MeetingScaleRow, pair string) float64 {
	for _, r := range rows {
		if r.Pair == pair {
			return r.ScaleGeV
		}
	}
	return math.NaN()
}

func scale(rows []PairwiseMeetingScaleRow, pair string) float64 {
	for _, r := range rows {
		if r.Pair == pair {
			return r.ScaleGeV
		}
	}
	return math.NaN()
}

func maxAbsPairDelta(u1, u2, u3 float64) float64 {
	return math.Max(math.Abs(u1-u2), math.Max(math.Abs(u1-u3), math.Abs(u2-u3)))
}

func containsPair(rows []PairwiseMeetingScaleRow, pair string) bool {
	for _, r := range rows {
		if r.Pair == pair {
			return true
		}
	}
	return false
}

func containsBoundaryChoice(rows []BoundaryChoiceResidualRow, name string) bool {
	for _, r := range rows {
		if r.ChoiceScale == name {
			return true
		}
	}
	return false
}

func containsBetaStrategy(rows []BetaDeformationVectorRow, key string) bool {
	for _, r := range rows {
		if strings.Contains(r.Strategy, key) {
			return true
		}
	}
	return false
}

func containsThresholdSlot(rows []ThresholdOriginSlotRow, key string) bool {
	for _, r := range rows {
		if strings.Contains(r.Candidate, key) || strings.Contains(r.Kind, key) {
			return true
		}
	}
	return false
}
