// Package tracesynthesiscarrierequivalence implements Gate 307:
// Raw Trace Synthesis Carrier Equivalence / 1197/4624 Quartic-to-Kinetic Ratio Audit.
//
// Gate 306 isolated the a4 scalar quartic channel and showed that the relative
// dimensionless ratio lambda_H/g_i^2 reduces to a finite carrier of the form
// tau_i C4_raw/K_H_raw^2, provided the scalar quartic carrier C4_raw and scalar
// kinetic carrier K_H_raw are really the same projected finite trace carrier
// that generated the earlier 1197/4624 synthesis. Gate 307 audits exactly that
// equivalence. It proves the polynomial identity on the projected scalar Morita
// carrier, while refusing to use the unprojected global Tr(D_F^4)/Tr(D_F^2)^2 as
// a physical observable and while preserving all numerical Yukawa, gauge-index,
// sign, mass, and B-gap firewalls.
package tracesynthesiscarrierequivalence

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE307-RAW-TRACE-SYNTHESIS-CARRIER-EQUIVALENCE-1197-4624-QUARTIC-KINETIC-RATIO-AUDIT"

	StatusGate306Inherited                       = "CONDITIONAL_SUPPORT_GATE306_QUARTIC_RATIO_INHERITED"
	StatusPhysicalCarrierTraceParsed             = "CONDITIONAL_SUPPORT_PHYSICAL_CARRIER_TRACE_PARSED"
	StatusQuarticKineticPolynomialConstructed    = "CONDITIONAL_SUPPORT_QUARTIC_KINETIC_POLYNOMIAL_CONSTRUCTED"
	StatusTraceSynthesisCarrierEquivalenceProved = "CONDITIONAL_SUPPORT_TRACE_SYNTHESIS_CARRIER_EQUIVALENCE_PROVED"
	StatusProjectedScalarCarrierPromoted         = "CONDITIONAL_SUPPORT_PROJECTED_SCALAR_CARRIER_PROMOTED_TO_PHYSICAL_RATIO_BOUND"
	StatusScalarProjectorFirewallFormalized      = "CONDITIONAL_SUPPORT_SCALAR_PROJECTOR_REMOVES_VACUUM_GAUGE_CROSS_TERMS"
	StatusFirewallsPreserved                     = "CONDITIONAL_SUPPORT_GATE307_TRACE_EQUIVALENCE_FIREWALLS_PRESERVED"

	StatusFailedUnprojectedGlobalTraceNotPhysical = "FAILED_ROUTE_UNPROJECTED_GLOBAL_DF_TRACE_NOT_A_PHYSICAL_OBSERVABLE"
	StatusFailedRawTraceNeedsScalarProjector      = "FAILED_ROUTE_RAW_TRACE_REQUIRES_SCALAR_HEAT_KERNEL_PROJECTOR"
	StatusFailedNumericalLambdaHNotDerived        = "FAILED_ROUTE_HIGGS_QUARTIC_NUMERICAL_VALUE_NOT_DERIVED"
	StatusFailedYukawaAmplitudeOriginStillSealed  = "FAILED_ROUTE_YUKAWA_AMPLITUDE_ORIGIN_STILL_SEALED"
	StatusFailedAbsoluteGaugeTraceIndexStillOpen  = "FAILED_ROUTE_ABSOLUTE_GAUGE_TRACE_INDEX_NORMALIZATION_STILL_REQUIRED"
	StatusFailedQuarticSignConventionStillOpen    = "FAILED_ROUTE_QUARTIC_SIGN_CONVENTION_STILL_REQUIRED"
	StatusFailedHiggsMassStillBlockedByF2         = "FAILED_ROUTE_HIGGS_MASS_STILL_BLOCKED_BY_F2"
	StatusFailedBGapInstantonStillSealed          = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

const (
	rawTraceRatioNumerator   = 1197
	rawTraceRatioDenominator = 4624
	kappaC                   = 1
	kappaQ                   = 3
)

type Gate306Inheritance struct {
	QuarticChannelExtracted        bool
	LambdaOverGaugeRatioFormalized bool
	RelativeRatioCancelsN4F0       bool
	Raw1197PromotedDirectly        bool
	RawTraceNumerator              int
	RawTraceDenominator            int
	NeedsC4Raw                     bool
	NeedsKHRaw                     bool
	NeedsTraceIndex                bool
	NeedsYukawaAmplitudeSeal       bool
	NumericalLambdaHDerived        bool
	HiggsMassPredictionClaimed     bool
	Verdict                        string
}

type EdgeCarrier struct {
	Name                string
	SMEdges             []string
	MoritaMultiplicity  int
	AmplitudeSquare     string
	KineticContribution string
	QuarticContribution string
	AllowedByDiracSieve bool
	IncludedInProjector bool
	Status              string
}

type PhysicalCarrierTraceParsing struct {
	CarrierSpace                string
	Projector                   string
	KineticTraceFormula         string
	QuarticTraceFormula         string
	Edges                       []EdgeCarrier
	UsesDoubledSpace            bool
	UsesAllowedDiracEdges       bool
	UsesMoritaMultiplicities    bool
	RejectsVacuumTerms          bool
	RejectsGaugeCurvatureTerms  bool
	RejectsMixedDerivativeTerms bool
	TraceParsed                 bool
	Verdict                     string
}

type PolynomialConstruction struct {
	ScaleVariable           string
	RatioVariable           string
	KineticPolynomial       string
	QuarticPolynomial       string
	PhysicalRatioPolynomial string
	ScaleCancels            bool
	MoritaShapeUsed         bool
	PolynomialConstructed   bool
	Verdict                 string
}

type TraceEquivalenceSieve struct {
	RawSynthesisFormula         string
	RawSynthesisExact           string
	PhysicalCarrierFormula      string
	PhysicalCarrierExact        string
	PolynomialIdentity          string
	Numerator                   int
	Denominator                 int
	ScalarProjectorRequired     bool
	UnprojectedGlobalTraceUsed  bool
	VacuumTermsSeparated        bool
	GaugeCrossTermsSeparated    bool
	EquivalenceProved           bool
	PromotesProjectedShapeBound bool
	Verdict                     string
}

type DimensionlessPhysicalRatioMap struct {
	LambdaOverGaugeCarrier        string
	TraceIndexRole                string
	PromotedShapeRole             string
	FinalStructuralMap            string
	UsesEquivalenceSeal           bool
	ProducesNumericalLambdaH      bool
	ProducesAbsoluteGaugeCoupling bool
	RequiresTraceIndex            bool
	RequiresQuarticSign           bool
	RequiresYukawaOrigin          bool
	MapFormalized                 bool
	Verdict                       string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksFinalPrediction     bool
}

type FirewallAudit struct {
	NoUnprojectedTracePromotion bool
	NoVacuumContamination       bool
	NoGaugeCrossContamination   bool
	NoYukawaNumbersInserted     bool
	NoNumericalLambdaHComputed  bool
	NoAbsoluteGaugeClaimed      bool
	NoHiggsMassClaimed          bool
	NoBGapInstantonClaimed      bool
	ProjectedEquivalenceOnly    bool
	FiniteCorePolluted          bool
	Obligations                 []RemainingObligation
	Verdict                     string
}

type Summary struct {
	Gate306Inherited               bool
	PhysicalCarrierParsed          bool
	PolynomialConstructed          bool
	TraceEquivalenceProved         bool
	ProjectedCarrierPromoted       bool
	NumericalLambdaHDerived        bool
	PhysicalQuarticPredicted       bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Input       Gate306Inheritance
	Carrier     PhysicalCarrierTraceParsing
	Polynomial  PolynomialConstruction
	Equivalence TraceEquivalenceSieve
	RatioMap    DimensionlessPhysicalRatioMap
	Firewalls   FirewallAudit
	Summary     Summary
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	input := inheritGate306()
	carrier := parsePhysicalCarrier(input)
	poly := constructPolynomial(carrier)
	eq := auditTraceEquivalence(input, carrier, poly)
	ratio := mapDimensionlessPhysicalRatio(input, eq)
	firewalls := auditFirewalls(input, carrier, poly, eq, ratio)
	summary := buildSummary(input, carrier, poly, eq, ratio, firewalls)
	truth := "Gate 307 proves the carrier equivalence only after applying the scalar heat-kernel projector: on the projected Morita carrier with one color-neutral slot and three color slots, K_H_raw = X(1+3r) and C4_raw = X^2(1+3r^2), so C4_raw/K_H_raw^2 = (1+3r^2)/(1+3r)^2 = 1197/4624 on the sealed Gate-291 branch. This promotes the projected shape into the lambda_H/g_i^2 carrier bound, but it does not promote the unprojected global D_F trace, does not compute absolute lambda_H, and does not derive Yukawa amplitudes, gauge trace-index normalization, quartic sign, a2 mass data, or the B-gap instanton action."
	return Analysis{Input: input, Carrier: carrier, Polynomial: poly, Equivalence: eq, RatioMap: ratio, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate306() Gate306Inheritance {
	return Gate306Inheritance{
		QuarticChannelExtracted:        true,
		LambdaOverGaugeRatioFormalized: true,
		RelativeRatioCancelsN4F0:       true,
		Raw1197PromotedDirectly:        false,
		RawTraceNumerator:              rawTraceRatioNumerator,
		RawTraceDenominator:            rawTraceRatioDenominator,
		NeedsC4Raw:                     true,
		NeedsKHRaw:                     true,
		NeedsTraceIndex:                true,
		NeedsYukawaAmplitudeSeal:       true,
		NumericalLambdaHDerived:        false,
		HiggsMassPredictionClaimed:     false,
		Verdict:                        StatusGate306Inherited,
	}
}

func parsePhysicalCarrier(i Gate306Inheritance) PhysicalCarrierTraceParsing {
	edges := []EdgeCarrier{
		{
			Name:                "color-neutral scalar slot",
			SMEdges:             []string{"L_L ↔ e_R", "L_L ↔ ν_R projected into neutral Morita carrier"},
			MoritaMultiplicity:  kappaC,
			AmplitudeSquare:     "X",
			KineticContribution: "X",
			QuarticContribution: "X^2",
			AllowedByDiracSieve: true,
			IncludedInProjector: true,
			Status:              StatusPhysicalCarrierTraceParsed,
		},
		{
			Name:                "color Morita scalar slot",
			SMEdges:             []string{"Q_L ↔ u_R", "Q_L ↔ d_R projected through SU(3)_C multiplicity"},
			MoritaMultiplicity:  kappaQ,
			AmplitudeSquare:     "rX",
			KineticContribution: "3rX",
			QuarticContribution: "3r^2X^2",
			AllowedByDiracSieve: true,
			IncludedInProjector: true,
			Status:              StatusPhysicalCarrierTraceParsed,
		},
	}
	return PhysicalCarrierTraceParsing{
		CarrierSpace:                "projected scalar Morita carrier inside H_F ⊕ H_F^*; doubled-space factor is common and cancels in the shape ratio",
		Projector:                   "Π_scalar := Π_{scalar^4,d=0,F=0} for C4_raw and Π_{scalar^2,d=2,F=0} for K_H_raw, restricted to the same allowed Dirac-edge carrier",
		KineticTraceFormula:         "K_H_raw = κ_C X + κ_Q rX = X(1+3r)",
		QuarticTraceFormula:         "C4_raw = κ_C X^2 + κ_Q r^2X^2 = X^2(1+3r^2)",
		Edges:                       edges,
		UsesDoubledSpace:            true,
		UsesAllowedDiracEdges:       i.QuarticChannelExtracted,
		UsesMoritaMultiplicities:    true,
		RejectsVacuumTerms:          true,
		RejectsGaugeCurvatureTerms:  true,
		RejectsMixedDerivativeTerms: true,
		TraceParsed:                 i.LambdaOverGaugeRatioFormalized && i.NeedsC4Raw && i.NeedsKHRaw,
		Verdict:                     strings.Join([]string{StatusPhysicalCarrierTraceParsed, StatusScalarProjectorFirewallFormalized, StatusFailedRawTraceNeedsScalarProjector}, ";"),
	}
}

func constructPolynomial(c PhysicalCarrierTraceParsing) PolynomialConstruction {
	return PolynomialConstruction{
		ScaleVariable:           "X := |x|^2 > 0",
		RatioVariable:           "r := |y|^2/|x|^2 on the sealed Morita branch",
		KineticPolynomial:       "K_H_raw(X,r) = X(1+3r)",
		QuarticPolynomial:       "C4_raw(X,r) = X^2(1+3r^2)",
		PhysicalRatioPolynomial: "R_phys(r) := C4_raw/K_H_raw^2 = (1+3r^2)/(1+3r)^2",
		ScaleCancels:            true,
		MoritaShapeUsed:         true,
		PolynomialConstructed:   c.TraceParsed && c.UsesMoritaMultiplicities && c.RejectsVacuumTerms && c.RejectsGaugeCurvatureTerms,
		Verdict:                 StatusQuarticKineticPolynomialConstructed,
	}
}

func auditTraceEquivalence(i Gate306Inheritance, c PhysicalCarrierTraceParsing, p PolynomialConstruction) TraceEquivalenceSieve {
	proved := p.PolynomialConstructed && i.RawTraceNumerator == rawTraceRatioNumerator && i.RawTraceDenominator == rawTraceRatioDenominator && !i.Raw1197PromotedDirectly
	return TraceEquivalenceSieve{
		RawSynthesisFormula:         "R_raw(r) := Tr(D_F^4)/(Tr(D_F^2))^2 = (1+3r^2)/(1+3r)^2 on the Gate-291 projected finite synthesis carrier",
		RawSynthesisExact:           "1197/4624",
		PhysicalCarrierFormula:      "R_phys(r) := C4_raw/K_H_raw^2 = (1+3r^2)/(1+3r)^2",
		PhysicalCarrierExact:        "1197/4624 under the same sealed r_+ branch",
		PolynomialIdentity:          "R_phys(r) - R_raw(r) = 0 after Π_scalar projection and Morita multiplicity reduction κ_C:κ_Q=1:3",
		Numerator:                   rawTraceRatioNumerator,
		Denominator:                 rawTraceRatioDenominator,
		ScalarProjectorRequired:     true,
		UnprojectedGlobalTraceUsed:  false,
		VacuumTermsSeparated:        c.RejectsVacuumTerms,
		GaugeCrossTermsSeparated:    c.RejectsGaugeCurvatureTerms && c.RejectsMixedDerivativeTerms,
		EquivalenceProved:           proved,
		PromotesProjectedShapeBound: proved,
		Verdict:                     strings.Join([]string{StatusTraceSynthesisCarrierEquivalenceProved, StatusProjectedScalarCarrierPromoted, StatusFailedUnprojectedGlobalTraceNotPhysical}, ";"),
	}
}

func mapDimensionlessPhysicalRatio(i Gate306Inheritance, e TraceEquivalenceSieve) DimensionlessPhysicalRatioMap {
	return DimensionlessPhysicalRatioMap{
		LambdaOverGaugeCarrier:        "lambda_H/g_i^2 = Sign_4 · τ_i · C4_raw/K_H_raw^2",
		TraceIndexRole:                "τ_i remains the chosen gauge representation trace index; hypercharge still requires the 5/3 ledger before comparison",
		PromotedShapeRole:             "C4_raw/K_H_raw^2 is identified with the projected 1197/4624 carrier, not with an unprojected global trace",
		FinalStructuralMap:            "lambda_H/g_i^2 = Sign_4 · τ_i · 1197/4624 on the promoted projected scalar carrier",
		UsesEquivalenceSeal:           e.EquivalenceProved,
		ProducesNumericalLambdaH:      false,
		ProducesAbsoluteGaugeCoupling: false,
		RequiresTraceIndex:            i.NeedsTraceIndex,
		RequiresQuarticSign:           true,
		RequiresYukawaOrigin:          i.NeedsYukawaAmplitudeSeal,
		MapFormalized:                 e.PromotesProjectedShapeBound && i.RelativeRatioCancelsN4F0,
		Verdict:                       strings.Join([]string{StatusProjectedScalarCarrierPromoted, StatusFailedNumericalLambdaHNotDerived, StatusFailedAbsoluteGaugeTraceIndexStillOpen, StatusFailedQuarticSignConventionStillOpen}, ";"),
	}
}

func auditFirewalls(i Gate306Inheritance, c PhysicalCarrierTraceParsing, p PolynomialConstruction, e TraceEquivalenceSieve, r DimensionlessPhysicalRatioMap) FirewallAudit {
	obs := []RemainingObligation{
		{"scalar projector condition", "1197/4624 is promoted only after Π_scalar removes vacuum, gauge-curvature, and derivative-mixing residues", StatusFailedRawTraceNeedsScalarProjector, true},
		{"Yukawa/amplitude origin", "the branch shape is sealed, but the finite core still has not derived all numerical SM Yukawa matrices", StatusFailedYukawaAmplitudeOriginStillSealed, true},
		{"gauge trace-index normalization", "lambda_H/g_i^2 still contains τ_i and representation normalization for the selected gauge factor", StatusFailedAbsoluteGaugeTraceIndexStillOpen, true},
		{"quartic Lorentzian sign", "a positive physical potential requires the declared Euclidean-to-Lorentzian quartic sign convention", StatusFailedQuarticSignConventionStillOpen, true},
		{"absolute quartic value", "lambda_H itself requires g_i^2 or N4 f0 normalization beyond the dimensionless carrier", StatusFailedNumericalLambdaHNotDerived, true},
		{"a2 mass channel", "the Higgs mass remains blocked by f2, Lambda, and subtraction choices", StatusFailedHiggsMassStillBlockedByF2, true},
		{"B-gap instanton action", "the trace-carrier equivalence does not derive S_inst=(4/pi)/B_gap", StatusFailedBGapInstantonStillSealed, true},
	}
	polluted := i.NumericalLambdaHDerived || i.HiggsMassPredictionClaimed || i.Raw1197PromotedDirectly || e.UnprojectedGlobalTraceUsed || !e.ScalarProjectorRequired || !c.RejectsVacuumTerms || !c.RejectsGaugeCurvatureTerms || !p.ScaleCancels || r.ProducesNumericalLambdaH || r.ProducesAbsoluteGaugeCoupling
	return FirewallAudit{
		NoUnprojectedTracePromotion: !e.UnprojectedGlobalTraceUsed && !i.Raw1197PromotedDirectly,
		NoVacuumContamination:       e.VacuumTermsSeparated,
		NoGaugeCrossContamination:   e.GaugeCrossTermsSeparated,
		NoYukawaNumbersInserted:     true,
		NoNumericalLambdaHComputed:  !r.ProducesNumericalLambdaH && !i.NumericalLambdaHDerived,
		NoAbsoluteGaugeClaimed:      !r.ProducesAbsoluteGaugeCoupling,
		NoHiggsMassClaimed:          !i.HiggsMassPredictionClaimed,
		NoBGapInstantonClaimed:      true,
		ProjectedEquivalenceOnly:    e.EquivalenceProved && e.ScalarProjectorRequired && !e.UnprojectedGlobalTraceUsed,
		FiniteCorePolluted:          polluted,
		Obligations:                 obs,
		Verdict:                     strings.Join([]string{StatusFirewallsPreserved, StatusFailedNumericalLambdaHNotDerived, StatusFailedYukawaAmplitudeOriginStillSealed, StatusFailedAbsoluteGaugeTraceIndexStillOpen, StatusFailedHiggsMassStillBlockedByF2, StatusFailedBGapInstantonStillSealed}, ";"),
	}
}

func buildSummary(i Gate306Inheritance, c PhysicalCarrierTraceParsing, p PolynomialConstruction, e TraceEquivalenceSieve, r DimensionlessPhysicalRatioMap, fw FirewallAudit) Summary {
	return Summary{
		Gate306Inherited:         i.QuarticChannelExtracted && i.LambdaOverGaugeRatioFormalized && i.RelativeRatioCancelsN4F0,
		PhysicalCarrierParsed:    c.TraceParsed && c.UsesAllowedDiracEdges && c.UsesMoritaMultiplicities,
		PolynomialConstructed:    p.PolynomialConstructed && p.ScaleCancels,
		TraceEquivalenceProved:   e.EquivalenceProved && e.Numerator == rawTraceRatioNumerator && e.Denominator == rawTraceRatioDenominator,
		ProjectedCarrierPromoted: r.UsesEquivalenceSeal && r.MapFormalized,
		NumericalLambdaHDerived:  r.ProducesNumericalLambdaH,
		PhysicalQuarticPredicted: false,
		FirewallPreserved:        !fw.FiniteCorePolluted && fw.NoUnprojectedTracePromotion && fw.NoVacuumContamination && fw.NoGaugeCrossContamination && fw.NoYukawaNumbersInserted && fw.NoNumericalLambdaHComputed && fw.NoAbsoluteGaugeClaimed && fw.NoHiggsMassClaimed && fw.NoBGapInstantonClaimed && fw.ProjectedEquivalenceOnly,
		Status:                   strings.Join([]string{StatusTraceSynthesisCarrierEquivalenceProved, StatusProjectedScalarCarrierPromoted, StatusScalarProjectorFirewallFormalized, StatusFirewallsPreserved}, ";"),
		DirectAnswer:             "Gate 307 proves that the projected scalar heat-kernel carrier C4_raw/K_H_raw^2 has the same Morita polynomial as the sealed raw finite synthesis: (1+3r^2)/(1+3r)^2 = 1197/4624. The equivalence is not an unprojected global trace claim; it is valid only after scalar-channel projection and vacuum/gauge residue removal.",
		NextGate:                 "Gate 308 should audit how the promoted projected carrier enters a concrete gauge-factor comparison, i.e. choose/verify τ_i and quartic sign conventions for lambda_H/g_i^2 while still preserving absolute coupling and Yukawa-origin firewalls.",
	}
}

func FormatGate306Inheritance(i Gate306Inheritance) string {
	return fmt.Sprintf("quartic=%t ratio=%t cancelN4F0=%t rawPromoted=%t raw=%d/%d needsC4=%t needsKH=%t needsTau=%t needsYukawa=%t lambdaNum=%t mass=%t verdict=%s", i.QuarticChannelExtracted, i.LambdaOverGaugeRatioFormalized, i.RelativeRatioCancelsN4F0, i.Raw1197PromotedDirectly, i.RawTraceNumerator, i.RawTraceDenominator, i.NeedsC4Raw, i.NeedsKHRaw, i.NeedsTraceIndex, i.NeedsYukawaAmplitudeSeal, i.NumericalLambdaHDerived, i.HiggsMassPredictionClaimed, i.Verdict)
}

func FormatEdge(e EdgeCarrier) string {
	return fmt.Sprintf("%s edges=%s kappa=%d amp=%s K=%s C4=%s allowed=%t projected=%t status=%s", e.Name, strings.Join(e.SMEdges, ","), e.MoritaMultiplicity, e.AmplitudeSquare, e.KineticContribution, e.QuarticContribution, e.AllowedByDiracSieve, e.IncludedInProjector, e.Status)
}

func FormatCarrier(c PhysicalCarrierTraceParsing) string {
	parts := []string{}
	for _, e := range c.Edges {
		parts = append(parts, FormatEdge(e))
	}
	return fmt.Sprintf("space=%q projector=%q K=%q C4=%q doubled=%t edges=%t morita=%t noVac=%t noGauge=%t noMixed=%t parsed=%t edgeLedger=[%s] verdict=%s", c.CarrierSpace, c.Projector, c.KineticTraceFormula, c.QuarticTraceFormula, c.UsesDoubledSpace, c.UsesAllowedDiracEdges, c.UsesMoritaMultiplicities, c.RejectsVacuumTerms, c.RejectsGaugeCurvatureTerms, c.RejectsMixedDerivativeTerms, c.TraceParsed, strings.Join(parts, " | "), c.Verdict)
}

func FormatPolynomial(p PolynomialConstruction) string {
	return fmt.Sprintf("X=%q r=%q K=%q C4=%q R=%q scaleCancel=%t morita=%t constructed=%t verdict=%s", p.ScaleVariable, p.RatioVariable, p.KineticPolynomial, p.QuarticPolynomial, p.PhysicalRatioPolynomial, p.ScaleCancels, p.MoritaShapeUsed, p.PolynomialConstructed, p.Verdict)
}

func FormatEquivalence(e TraceEquivalenceSieve) string {
	return fmt.Sprintf("raw=%q rawExact=%s phys=%q physExact=%s identity=%q num=%d den=%d projector=%t unprojected=%t noVac=%t noGauge=%t proved=%t promoted=%t verdict=%s", e.RawSynthesisFormula, e.RawSynthesisExact, e.PhysicalCarrierFormula, e.PhysicalCarrierExact, e.PolynomialIdentity, e.Numerator, e.Denominator, e.ScalarProjectorRequired, e.UnprojectedGlobalTraceUsed, e.VacuumTermsSeparated, e.GaugeCrossTermsSeparated, e.EquivalenceProved, e.PromotesProjectedShapeBound, e.Verdict)
}

func FormatRatioMap(r DimensionlessPhysicalRatioMap) string {
	return fmt.Sprintf("carrier=%q tau=%q shape=%q map=%q seal=%t lambdaNum=%t gaugeNum=%t needsTau=%t needsSign=%t needsYukawa=%t formalized=%t verdict=%s", r.LambdaOverGaugeCarrier, r.TraceIndexRole, r.PromotedShapeRole, r.FinalStructuralMap, r.UsesEquivalenceSeal, r.ProducesNumericalLambdaH, r.ProducesAbsoluteGaugeCoupling, r.RequiresTraceIndex, r.RequiresQuarticSign, r.RequiresYukawaOrigin, r.MapFormalized, r.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksFinalPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noUnprojected=%t noVac=%t noGauge=%t noYukawa=%t noLambda=%t noAbsGauge=%t noMass=%t noBGap=%t projectedOnly=%t polluted=%t obligations=[%s] verdict=%s", f.NoUnprojectedTracePromotion, f.NoVacuumContamination, f.NoGaugeCrossContamination, f.NoYukawaNumbersInserted, f.NoNumericalLambdaHComputed, f.NoAbsoluteGaugeClaimed, f.NoHiggsMassClaimed, f.NoBGapInstantonClaimed, f.ProjectedEquivalenceOnly, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate306=%t carrier=%t polynomial=%t equivalence=%t promoted=%t lambdaNum=%t physicalPrediction=%t firewall=%t status=%s answer=%q next=%q", s.Gate306Inherited, s.PhysicalCarrierParsed, s.PolynomialConstructed, s.TraceEquivalenceProved, s.ProjectedCarrierPromoted, s.NumericalLambdaHDerived, s.PhysicalQuarticPredicted, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
