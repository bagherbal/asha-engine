// Package generation2strongcouplingthresholdresidualledgeraudit implements
// Gate 607: Strong-Coupling Threshold Residual Ledger Audit.
//
// Gate 606 exposed the boundary-to-endpoint RG transport spine and the clean
// g3 mismatch at the electroweak g1=g2 meeting scale. Gate 607 converts that
// mismatch into invariant transport variables, defines the exact strong-sector
// threshold slot required to close the mismatch, audits one-loop beta-coefficient
// deformation size, computes the one-loop meeting-scale triangle, and preserves
// the no-unification / no-new-physics firewalls.
package generation2strongcouplingthresholdresidualledgeraudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryendpointthresholdtransportspineaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE607-STRONG-COUPLING-THRESHOLD-RESIDUAL-LEDGER-AUDIT"

	StatusGate606Inherited                = "PASS_GATE606_GAUGE_SPINE_INHERITED"
	StatusStrongResidualConverted         = "PASS_STRONG_RESIDUAL_CONVERTED_IN_MULTIPLE_SCHEMES"
	StatusStrongThresholdSlotDefined      = "PASS_STRONG_THRESHOLD_SLOT_DEFINED"
	StatusRequiredThresholdQuantified     = "CONDITIONAL_SUPPORT_REQUIRED_STRONG_THRESHOLD_LEDGER_QUANTIFIED"
	StatusBetaDeformationComputed         = "CONDITIONAL_SUPPORT_BETA_DEFORMATION_SIZE_COMPUTED"
	StatusMeetingScaleTriangleComputed    = "PASS_MEETING_SCALE_TRIANGLE_COMPUTED_ONE_LOOP"
	StatusMeetingScaleShiftNotClosure     = "CONDITIONAL_SUPPORT_MEETING_SCALE_SHIFT_SHOWS_TRIANGLE_NOT_SINGLE_UNIFICATION_POINT"
	StatusHigherLoopThresholdSlotsClassed = "PASS_HIGHER_LOOP_THRESHOLD_SOURCE_SLOTS_CLASSIFIED"
	StatusScalarSectorSeparated           = "PASS_SCALAR_SECTOR_SEPARATED_FROM_STRONG_THRESHOLD_LEDGER"
	StatusNoNativeStrongThresholdTheorem  = "FAILED_ROUTE_NO_NATIVE_STRONG_THRESHOLD_THEOREM"
	StatusNoFullGaugeUnificationClaim     = "FAILED_ROUTE_NO_FULL_GAUGE_UNIFICATION_CLAIM"
	StatusNoNativeExtraColoredContent     = "FAILED_ROUTE_NO_NATIVE_EXTRA_COLORED_FIELD_CONTENT"
	StatusNoNativeBoundaryG3Correction    = "FAILED_ROUTE_NO_NATIVE_G3_BOUNDARY_MATCHING_TERM"
	StatusNoEndpointDerivation            = "FAILED_ROUTE_NO_ENDPOINT_COUPLING_DERIVATION"
	StatusNoThresholdExistenceClaim       = "FAILED_ROUTE_THRESHOLD_SLOT_DEFINED_NOT_ASSERTED_REAL"
	StatusNoNewPhysicsDiscoveryClaim      = "FAILED_ROUTE_NO_NEW_PHYSICS_DISCOVERY_CLAIM"
	StatusGate607Boundary                 = "FIREWALL_PRESERVED_GATE607_STRONG_THRESHOLD_RESIDUAL_BOUNDARY"
	StatusNoWZPhotonDerivationFirewall    = "FIREWALL_PRESERVED_NO_WZ_PHOTON_OR_MASS_DERIVATION"
	StatusScalarNotMixedFirewall          = "FIREWALL_PRESERVED_SCALAR_THRESHOLD_NOT_MIXED_WITH_STRONG_LEDGER"
	StatusNoObservedEndpointPromotion     = "FIREWALL_PRESERVED_NO_OBSERVED_ENDPOINT_DERIVATION"
)

const (
	b1SM = 41.0 / 10.0
	b2SM = -19.0 / 6.0
	b3SM = -7.0
)

type InheritedGate606 struct {
	GaugeSpinePresent bool
	Lambda12GeV       float64
	GStar             float64
	G3Lambda          float64
	Delta3Runtime     float64
	R3                float64
	Verdict           string
}

type StrongResidualConversionRow struct {
	Quantity       string
	Formula        string
	Value          float64
	SignConvention string
	Interpretation string
	Verdict        string
}

type ThresholdCorrectionSlotRow struct {
	Slot           string
	Definition     string
	RequiredValue  float64
	Unit           string
	SignConvention string
	Interpretation string
	Verdict        string
}

type BetaCoefficientDeformationAudit struct {
	LogInterval              float64
	SMb3                     float64
	RequiredDeltaUCorrection float64
	DeltaB3Required          float64
	EffectiveB3              float64
	FractionOfAbsSMb3        float64
	Formula                  string
	Interpretation           string
	Verdict                  string
}

type MeetingScaleRow struct {
	Pair              string
	LogMuOverMZ       float64
	ScaleGeV          float64
	CouplingAtMeeting float64
	Formula           string
	Interpretation    string
	Verdict           string
}

type SourceClassificationRow struct {
	Candidate     string
	Type          string
	CouldAffectG3 bool
	CurrentStatus string
	NativeStatus  string
	Verdict       string
}

type NativeASHAStatus struct {
	ProvidesNativeStrongThreshold bool
	ProvidesExtraColoredContent   bool
	ProvidesBoundaryG3Correction  bool
	ClaimsFullUnification         bool
	Statement                     string
	Verdict                       string
}

type ScalarRelation struct {
	LambdaLambda12        float64
	Statement             string
	MixedIntoStrongLedger bool
	Verdict               string
}

type Firewalls struct {
	ClaimsGaugeUnification bool
	ClaimsNewPhysics       bool
	DerivesEndpoint        bool
	DerivesWZPhoton        bool
	ThresholdAssertedReal  bool
	ScalarMixedIntoStrong  bool
	Verdict                string
}

type Analysis struct {
	Inherited             InheritedGate606
	ResidualConversions   []StrongResidualConversionRow
	ThresholdSlots        []ThresholdCorrectionSlotRow
	BetaDeformation       BetaCoefficientDeformationAudit
	MeetingScales         []MeetingScaleRow
	SourceClassifications []SourceClassificationRow
	NativeStatus          NativeASHAStatus
	ScalarRelation        ScalarRelation
	Firewalls             Firewalls
	Truth                 string
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
	g606, err := generation2boundaryendpointthresholdtransportspineaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate606 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport bundle: %w", err)
	}
	inherited := inherit(g606, bundle)
	residuals := buildResidualConversions(bundle)
	slots := buildThresholdSlots(bundle)
	beta := buildBetaDeformation(bundle)
	meetings := buildMeetingScales(bundle)
	sources := buildSourceClassifications()
	native := buildNativeStatus()
	scalar := buildScalarRelation(bundle)
	firewalls := auditFirewalls()
	truth := "Gate 607 converts the one-loop strong-coupling mismatch at Lambda_12 into coupling, inverse-coupling, alpha-inverse, threshold-slot, beta-deformation, and meeting-scale-triangle variables. The required strong correction is quantified, but no threshold, unification, endpoint derivation, or new physics sector is asserted."
	return Analysis{inherited, residuals, slots, beta, meetings, sources, native, scalar, firewalls, truth}, nil
}

func inherit(a generation2boundaryendpointthresholdtransportspineaudit.Analysis, b historytransport.Bundle) InheritedGate606 {
	return InheritedGate606{
		GaugeSpinePresent: len(a.GaugeTransport) > 0 && containsGauge606(a.GaugeTransport, "Delta_3"),
		Lambda12GeV:       b.GaugeBoundary.Lambda12GeV,
		GStar:             b.GaugeBoundary.GStar,
		G3Lambda:          b.GaugeBoundary.G3Lambda,
		Delta3Runtime:     b.GaugeBoundary.Delta3,
		R3:                b.GaugeBoundary.R3,
		Verdict:           StatusGate606Inherited,
	}
}

func buildResidualConversions(b historytransport.Bundle) []StrongResidualConversionRow {
	gStar := b.GaugeBoundary.GStar
	g3 := b.GaugeBoundary.G3Lambda
	uStar := 1 / (gStar * gStar)
	u3 := 1 / (g3 * g3)
	deltaRuntime := u3 - uStar
	deltaRequired := uStar - u3
	return []StrongResidualConversionRow{
		{"Delta g3", "g3(Lambda12)-g_star", g3 - gStar, "positive means runtime strong coupling is too large", "g3 is above the electroweak meeting coupling", StatusStrongResidualConverted},
		{"R_3 - 1", "g3/g_star - 1", b.GaugeBoundary.R3 - 1, "positive means g3 exceeds g_star", "relative strong coupling excess", StatusStrongResidualConverted},
		{"u_star", "1/g_star^2", uStar, "inverse coupling target", "electroweak inverse coupling at g1=g2 meeting", StatusStrongResidualConverted},
		{"u_3", "1/g3(Lambda12)^2", u3, "runtime inverse strong coupling", "strong inverse coupling is too small", StatusStrongResidualConverted},
		{"Delta u3 runtime", "1/g3^2 - 1/g_star^2", deltaRuntime, "negative means g3 is too large", "same as Gate606 Delta_3", StatusStrongResidualConverted},
		{"required positive inverse correction", "1/g_star^2 - 1/g3^2", deltaRequired, "positive correction added to runtime u3 closes the mismatch", "threshold slot magnitude in u=1/g^2 convention", StatusRequiredThresholdQuantified},
		{"Delta alpha3^-1 runtime", "4*pi*(1/g3^2 - 1/g_star^2)", 4 * math.Pi * deltaRuntime, "negative alpha-inverse residual", "runtime alpha_3^{-1} is below target", StatusStrongResidualConverted},
		{"required Delta alpha3^-1", "4*pi*(1/g_star^2 - 1/g3^2)", 4 * math.Pi * deltaRequired, "positive alpha-inverse correction closes mismatch", "alpha-inverse threshold ledger magnitude", StatusRequiredThresholdQuantified},
	}
}

func buildThresholdSlots(b historytransport.Bundle) []ThresholdCorrectionSlotRow {
	gStar := b.GaugeBoundary.GStar
	g3 := b.GaugeBoundary.G3Lambda
	delta := 1/(gStar*gStar) - 1/(g3*g3)
	return []ThresholdCorrectionSlotRow{
		{"delta_3^threshold", "1/g3_eff^2 = 1/g3_runtime^2 + delta_3^threshold", delta, "dimensionless inverse-coupling", "positive value raises u3 to u_star", "exact slot required to close g3 to g_star", StatusStrongThresholdSlotDefined},
		{"Delta_3^alpha", "4*pi*delta_3^threshold", 4 * math.Pi * delta, "alpha inverse", "positive alpha^{-1} correction", "same correction in alpha_3^{-1} convention", StatusStrongThresholdSlotDefined},
		{"delta_3^matching", "low/high-scale matching contribution to delta_3^threshold", math.NaN(), "symbolic", "unfitted", "matching slot only; no existence claim", StatusNoThresholdExistenceClaim},
		{"delta_3^higher_loop", "two-loop and higher-loop RG contribution to strong inverse coupling", math.NaN(), "symbolic", "unfitted", "higher-loop slot only; not computed in v1", StatusNoThresholdExistenceClaim},
		{"delta_3^heavy_threshold", "heavy colored or boundary threshold contribution near Lambda_12", math.NaN(), "symbolic", "unfitted", "no native extra content supplied", StatusNoNativeExtraColoredContent},
	}
}

func buildBetaDeformation(b historytransport.Bundle) BetaCoefficientDeformationAudit {
	gStar := b.GaugeBoundary.GStar
	g3 := b.GaugeBoundary.G3Lambda
	delta := 1/(gStar*gStar) - 1/(g3*g3)
	t := b.GaugeBoundary.LogLambda12Mu0
	db := -8 * math.Pi * math.Pi * delta / t
	return BetaCoefficientDeformationAudit{
		LogInterval:              t,
		SMb3:                     b3SM,
		RequiredDeltaUCorrection: delta,
		DeltaB3Required:          db,
		EffectiveB3:              b3SM + db,
		FractionOfAbsSMb3:        math.Abs(db / b3SM),
		Formula:                  "Delta b3_required = -8*pi^2*delta_u_correction / ln(Lambda12/M_Z)",
		Interpretation:           "If the entire mismatch were represented as a constant one-loop b3 deformation over the full interval, b3 would need to be more negative by about 13.3% of |b3_SM|. This is a diagnostic size, not a proposed model.",
		Verdict:                  StatusBetaDeformationComputed,
	}
}

func buildMeetingScales(b historytransport.Bundle) []MeetingScaleRow {
	g1 := b.EndVector.G1
	g2 := b.EndVector.G2
	g3 := b.EndVector.G3
	mu0 := b.GaugeBoundary.Mu0GeV
	return []MeetingScaleRow{
		meetingRow("Lambda_12", g1, g2, b1SM, b2SM, mu0, "g1=g2 electroweak canonical meeting", StatusMeetingScaleTriangleComputed),
		meetingRow("Lambda_13", g1, g3, b1SM, b3SM, mu0, "g1=g3 meeting lies above Lambda_12 in v1", StatusMeetingScaleShiftNotClosure),
		meetingRow("Lambda_23", g2, g3, b2SM, b3SM, mu0, "g2=g3 meeting lies far above Lambda_12 in v1", StatusMeetingScaleShiftNotClosure),
	}
}

func meetingRow(pair string, ga, gb, ba, bb, mu0 float64, interp string, verdict string) MeetingScaleRow {
	ua := 1 / (ga * ga)
	ub := 1 / (gb * gb)
	t := 8 * math.Pi * math.Pi * (ua - ub) / (ba - bb)
	scale := mu0 * math.Exp(t)
	uMeet := ua - ba*t/(8*math.Pi*math.Pi)
	gMeet := 1 / math.Sqrt(uMeet)
	return MeetingScaleRow{pair, t, scale, gMeet, "t=8*pi^2*(g_a^-2-g_b^-2)/(b_a-b_b)", interp, verdict}
}

func buildSourceClassifications() []SourceClassificationRow {
	return []SourceClassificationRow{
		{"two-loop SM RG", "higher-loop transport", true, "not included in v1", "bridge calculation slot, not native theorem", StatusHigherLoopThresholdSlotsClassed},
		{"top/Higgs/W/Z matching", "low-scale threshold/matching", true, "not included in v1", "observed endpoint matching slot", StatusHigherLoopThresholdSlotsClassed},
		{"heavy threshold near Lambda_12", "high-scale threshold", true, "open symbolic slot", "no native heavy field content", StatusNoNativeExtraColoredContent},
		{"finite spectral-action boundary correction", "boundary matching", true, "not constructed", "no native g3 correction theorem", StatusNoNativeBoundaryG3Correction},
		{"B-sector or extra colored content", "new-sector deformation", true, "not supplied by current gates", "no new physics discovery claim", StatusNoNewPhysicsDiscoveryClaim},
		{"scheme conversion", "renormalization convention", true, "explicitly labeled as missing", "bridge scheme ledger", StatusHigherLoopThresholdSlotsClassed},
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		ProvidesNativeStrongThreshold: false,
		ProvidesExtraColoredContent:   false,
		ProvidesBoundaryG3Correction:  false,
		ClaimsFullUnification:         false,
		Statement:                     "Current ASHA data certify the g1=g2 boundary normalization and expose a g3 residual, but supply no native strong threshold, no extra colored field content, no boundary g3 correction, and no full gauge-unification theorem.",
		Verdict:                       StatusNoNativeStrongThresholdTheorem,
	}
}

func buildScalarRelation(b historytransport.Bundle) ScalarRelation {
	return ScalarRelation{
		LambdaLambda12:        b.ScalarTransport.LambdaLambda12,
		Statement:             "The scalar quartic residual lambda(Lambda_12) is a separate v1 scalar transport wound and is not used to close the strong coupling threshold slot in Gate 607.",
		MixedIntoStrongLedger: false,
		Verdict:               StatusScalarSectorSeparated,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, StatusGate607Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate606Inherited,
		StatusStrongResidualConverted,
		StatusStrongThresholdSlotDefined,
		StatusRequiredThresholdQuantified,
		StatusBetaDeformationComputed,
		StatusMeetingScaleTriangleComputed,
		StatusMeetingScaleShiftNotClosure,
		StatusHigherLoopThresholdSlotsClassed,
		StatusScalarSectorSeparated,
		StatusNoNativeStrongThresholdTheorem,
		StatusNoFullGaugeUnificationClaim,
		StatusNoNativeExtraColoredContent,
		StatusNoNativeBoundaryG3Correction,
		StatusNoEndpointDerivation,
		StatusNoThresholdExistenceClaim,
		StatusNoNewPhysicsDiscoveryClaim,
		StatusGate607Boundary,
		StatusNoWZPhotonDerivationFirewall,
		StatusScalarNotMixedFirewall,
		StatusNoObservedEndpointPromotion,
	}
}

func containsGauge606(rows []generation2boundaryendpointthresholdtransportspineaudit.GaugeTransportRow, q string) bool {
	for _, r := range rows {
		if r.Quantity == q {
			return true
		}
	}
	return false
}

func containsResidual(rows []StrongResidualConversionRow, q string) bool {
	for _, r := range rows {
		if r.Quantity == q {
			return true
		}
	}
	return false
}

func containsSlot(rows []ThresholdCorrectionSlotRow, s string) bool {
	for _, r := range rows {
		if r.Slot == s {
			return true
		}
	}
	return false
}

func containsMeeting(rows []MeetingScaleRow, pair string) bool {
	for _, r := range rows {
		if r.Pair == pair {
			return true
		}
	}
	return false
}

func containsSource(rows []SourceClassificationRow, candidate string) bool {
	for _, r := range rows {
		if r.Candidate == candidate || strings.Contains(r.Candidate, candidate) {
			return true
		}
	}
	return false
}
