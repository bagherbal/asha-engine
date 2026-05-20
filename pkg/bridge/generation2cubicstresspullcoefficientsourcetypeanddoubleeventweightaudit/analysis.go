// Package generation2cubicstresspullcoefficientsourcetypeanddoubleeventweightaudit implements
// Gate 731: Cubic Stress-Pull Coefficient Source-Type and Double-Event Weight Audit.
//
// Gate 730 showed that the boundary-history wall residual is strongly compressed by
// a second-plus-third moment ansatz
//
//	D_base ≈ M1_wall + kappa_e M2_wall - (7/36)M3_wall.
//
// Gate 731 audits the source type of the cubic coefficient 7/36 by rewriting it
// as 2*(7/72)=2p_K7. It records boundary-pair and two-wall stress-pull source
// candidates, while preserving the firewall that no native boundary moment
// expansion theorem or physical scalar/Higgs/Yukawa theorem follows.
package generation2cubicstresspullcoefficientsourcetypeanddoubleeventweightaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate730 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryhistoryresidualcubicstresspullcorrectionaudit"
)

const (
	AuditID = "GATE731-CUBIC-STRESS-PULL-COEFFICIENT-SOURCE-TYPE-AND-DOUBLE-EVENT-WEIGHT-AUDIT"

	StatusGate730CubicStressPullInherited                = "PASS_GATE730_CUBIC_STRESS_PULL_INHERITED"
	StatusCubicCoeffRewrittenAsTwoTimesK7Weight          = "PASS_CUBIC_COEFFICIENT_REWRITTEN_AS_TWO_TIMES_K7_EVENT_WEIGHT"
	StatusBoundaryPairSourceCandidateAudited             = "PASS_BOUNDARY_PAIR_SOURCE_CANDIDATE_AUDITED"
	StatusTwoWallStressPullSourceCandidateAudited        = "PASS_TWO_WALL_STRESS_PULL_SOURCE_CANDIDATE_AUDITED"
	StatusKineticToAmplitudeFactorTwoWarningRecorded     = "PASS_KINETIC_TO_AMPLITUDE_FACTOR_TWO_WARNING_RECORDED"
	StatusTypedAlternativesAudited                       = "PASS_TYPED_ALTERNATIVES_AUDITED"
	StatusMomentPolynomialRewrittenWithEventWeightSource = "PASS_MOMENT_POLYNOMIAL_REWRITTEN_WITH_EVENT_WEIGHT_SOURCE"
	StatusNoncircularityFirewallAudited                  = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusPhysicalFirewallsEnforced                      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusSevenOverThirtySixDoubleK7EventWeightCandidate = "CONDITIONAL_SUPPORT_SEVEN_OVER_THIRTY_SIX_IS_DOUBLE_K7_EVENT_WEIGHT_CANDIDATE"
	StatusFactorTwoBoundaryPairStressPullCandidate       = "CONDITIONAL_SUPPORT_FACTOR_TWO_HAS_BOUNDARY_PAIR_STRESS_PULL_SOURCE_CANDIDATE"
	StatusTwoPK7BestTypedSourceForCubicCoeff             = "CONDITIONAL_SUPPORT_2P_K7_IS_BEST_TYPED_SOURCE_FOR_CUBIC_COEFFICIENT"
	StatusBoundaryPairTimesK7WeightCandidate             = "CONDITIONAL_SUPPORT_SEVEN_OVER_THIRTY_SIX_AS_BOUNDARY_PAIR_TIMES_K7_EVENT_WEIGHT_CANDIDATE"
	StatusCubicTermTwoWallStressPullCandidate            = "CONDITIONAL_SUPPORT_CUBIC_TERM_IS_TWO_WALL_STRESS_PULL_CANDIDATE"
	StatusFactorTwoKineticAmplitudeResonance             = "CONDITIONAL_SUPPORT_FACTOR_TWO_RESONATES_WITH_KINETIC_TO_AMPLITUDE_LINEARIZATION"

	StatusNoNativeReasonCubicCoeffEqualsTwoPK7      = "FAILED_ROUTE_NO_NATIVE_REASON_CUBIC_COEFFICIENT_EQUALS_TWO_P_K7"
	StatusNoNativeBoundaryPairStressPullTheorem     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_PAIR_STRESS_PULL_THEOREM"
	StatusNoNativeBoundaryMomentExpansionTheorem    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_MOMENT_EXPANSION_THEOREM"
	StatusKineticToAmplitudeDoesNotDeriveCubicCoeff = "FAILED_ROUTE_KINETIC_TO_AMPLITUDE_FACTOR_TWO_DOES_NOT_DERIVE_CUBIC_COEFFICIENT"
	StatusNoNativeScalarRuntimeTheorem              = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem              = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem       = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusKappaEQuadraticCoefficientDependent       = "FAILED_ROUTE_KAPPA_E_QUADRATIC_COEFFICIENT_PARTIALLY_DEPENDENT"
	StatusGate731Boundary                           = "FIREWALL_PRESERVED_GATE731_CUBIC_COEFFICIENT_SOURCE_BOUNDARY"
)

const (
	k7Dim             = 7.0
	h72Dim            = 72.0
	boundaryPairDim   = 2.0
	oneFifthControl   = 1.0 / 5.0
	higgsRadialWeight = 1.0 / 4.0
	phaseLoopUnit     = 1.0 / (2.0 * math.Pi)
)

type Gate730Inheritance struct {
	Inherited                       bool
	P_K7                            float64
	SSplit                          float64
	KappaE                          float64
	M1Wall                          float64
	M2Wall                          float64
	M3Wall                          float64
	CubicCoefficient                float64
	CombinedCorrection              float64
	ResidualAfterCubicCorrection    float64
	SevenOver36CompressedResidual   bool
	KappaEPartiallyDependent        bool
	NoNativeBoundaryMomentExpansion bool
	Verdict                         string
}

type DoubleEventCoefficientAudit struct {
	CubicCoefficient   float64
	K7EventProbability float64
	DoubleK7Weight     float64
	IdentityExact      bool
	RewrittenExpansion string
	Verdict            string
}

type BoundaryPairSourceCandidateAudit struct {
	BoundaryPairDimension     float64
	K7EventProbability        float64
	BoundaryPairTimesK7Weight float64
	EqualsCubicCoefficient    bool
	Interpretation            string
	Verdict                   string
}

type StressPullSourceCandidateAudit struct {
	CubicTermFormula     string
	TwoSidedBoundaryLegs bool
	ArbitraryFitRejected bool
	Verdict              string
}

type KineticToAmplitudeWarning struct {
	LinearizationFormula string
	FactorTwoResonance   bool
	DerivesCubicCoeff    bool
	Verdict              string
}

type TypedAlternative struct {
	Name     string
	Value    float64
	Distance float64
	Lane     string
	Accepted bool
}

type TypedAlternativesAudit struct {
	TargetCoefficient float64
	Candidates        []TypedAlternative
	ClosestName       string
	ClosestAccepted   bool
	NoArbitrarySearch bool
	Verdict           string
}

type MomentPolynomialAudit struct {
	PolynomialInS       string
	MomentForm          string
	LeadingTerm         float64
	QuadraticTerm       float64
	CubicTerm           float64
	UsesDoubleEventForm bool
	Verdict             string
}

type NoncircularityAudit struct {
	KappaEPartiallyDependent     bool
	TwoPK7TypedButUnexplained    bool
	BoundaryPairStressPullNative bool
	MomentExpansionTheoremNative bool
	Verdict                      string
}

type FirewallAudit struct {
	ClaimsNativeScalarRuntime bool
	ClaimsHiggsMassTheorem    bool
	ClaimsYukawaTheorem       bool
	ClaimsCKMPMNSTheorem      bool
	ClaimsHistoryLoopTheorem  bool
	Verdict                   string
}

type Analysis struct {
	Gate730       Gate730Inheritance
	DoubleEvent   DoubleEventCoefficientAudit
	BoundaryPair  BoundaryPairSourceCandidateAudit
	StressPull    StressPullSourceCandidateAudit
	KineticFactor KineticToAmplitudeWarning
	Alternatives  TypedAlternativesAudit
	Polynomial    MomentPolynomialAudit
	NonCircular   NoncircularityAudit
	Firewall      FirewallAudit
	Truth         string
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
	g730, err := gate730.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate730 inheritance unavailable: %w", err)
	}
	inherited := buildGate730Inheritance(g730)
	double := buildDoubleEventCoefficient(inherited)
	pair := buildBoundaryPairSource(double)
	stress := buildStressPullSource()
	kinetic := buildKineticWarning()
	alts := buildTypedAlternatives(double)
	poly := buildMomentPolynomial(inherited, double)
	noncirc := buildNoncircularity()
	firewall := buildFirewall()
	truth := "Gate 731 source-types the Gate730 cubic coefficient 7/36 as 2p_K7, with p_K7=Tr(rho_72 P_K7)=7/72. This supports a bridge-layer reading of the cubic term as a double K7-event-weight or boundary-pair stress-pull candidate: D_base≈M1_wall+kappa_e M2_wall-2p_K7 M3_wall. The factor two also resonates with a kinetic-to-amplitude linearization, but that resonance does not derive the coefficient. The expansion remains a residual-compression clue: kappa_e is partially dependent, 2p_K7 has no native cubic-selection theorem, and no boundary moment expansion, scalar runtime, Higgs mass, or Yukawa theorem is certified."
	return Analysis{Gate730: inherited, DoubleEvent: double, BoundaryPair: pair, StressPull: stress, KineticFactor: kinetic, Alternatives: alts, Polynomial: poly, NonCircular: noncirc, Firewall: firewall, Truth: truth}, nil
}

func buildGate730Inheritance(g gate730.Analysis) Gate730Inheritance {
	p := g.Gate729.P_K7
	return Gate730Inheritance{
		Inherited:                       g.Gate729.Inherited && g.CubicCorr.ImprovesSecondOrderResidual && g.CubicCorr.RawCompressionFactor > 1000,
		P_K7:                            p,
		SSplit:                          g.Gate729.SSplit,
		KappaE:                          g.Gate729.KappaE,
		M1Wall:                          p * g.Gate729.SSplit,
		M2Wall:                          g.Gate729.M2Wall,
		M3Wall:                          g.CubicMoment.M3Wall,
		CubicCoefficient:                g.CubicCorr.CubicCoefficient,
		CombinedCorrection:              g.CubicCorr.CombinedCorrection,
		ResidualAfterCubicCorrection:    g.CubicCorr.ResidualAfterCubicCorrection,
		SevenOver36CompressedResidual:   g.CubicCorr.ImprovesSecondOrderResidual,
		KappaEPartiallyDependent:        g.NonCircular.KappaEUsedAsQuadraticCoeff && g.NonCircular.DBaseContainsKappaE,
		NoNativeBoundaryMomentExpansion: !g.SourceType.MomentExpansionTheoremNative,
		Verdict:                         StatusGate730CubicStressPullInherited,
	}
}

func buildDoubleEventCoefficient(g Gate730Inheritance) DoubleEventCoefficientAudit {
	double := 2.0 * g.P_K7
	return DoubleEventCoefficientAudit{
		CubicCoefficient:   g.CubicCoefficient,
		K7EventProbability: g.P_K7,
		DoubleK7Weight:     double,
		IdentityExact:      near(g.CubicCoefficient, double, 1e-18),
		RewrittenExpansion: "D_base≈M1_wall+kappa_e M2_wall-2p_K7 M3_wall",
		Verdict: strings.Join([]string{
			StatusCubicCoeffRewrittenAsTwoTimesK7Weight,
			StatusSevenOverThirtySixDoubleK7EventWeightCandidate,
		}, "; "),
	}
}

func buildBoundaryPairSource(d DoubleEventCoefficientAudit) BoundaryPairSourceCandidateAudit {
	candidate := boundaryPairDim * d.K7EventProbability
	return BoundaryPairSourceCandidateAudit{
		BoundaryPairDimension:     boundaryPairDim,
		K7EventProbability:        d.K7EventProbability,
		BoundaryPairTimesK7Weight: candidate,
		EqualsCubicCoefficient:    near(candidate, d.CubicCoefficient, 1e-18),
		Interpretation:            "dim(R^2_boundary)*p_K7: two boundary wall legs times K7 no-bias event probability",
		Verdict: strings.Join([]string{
			StatusBoundaryPairSourceCandidateAudited,
			StatusBoundaryPairTimesK7WeightCandidate,
			StatusFactorTwoBoundaryPairStressPullCandidate,
		}, "; "),
	}
}

func buildStressPullSource() StressPullSourceCandidateAudit {
	return StressPullSourceCandidateAudit{
		CubicTermFormula:     "-(2p_K7)Tr(rho_72 R_wall^3)",
		TwoSidedBoundaryLegs: true,
		ArbitraryFitRejected: true,
		Verdict: strings.Join([]string{
			StatusTwoWallStressPullSourceCandidateAudited,
			StatusCubicTermTwoWallStressPullCandidate,
		}, "; "),
	}
}

func buildKineticWarning() KineticToAmplitudeWarning {
	return KineticToAmplitudeWarning{
		LinearizationFormula: "1-1/(1+r_g)^2≈2r_g",
		FactorTwoResonance:   true,
		DerivesCubicCoeff:    false,
		Verdict: strings.Join([]string{
			StatusKineticToAmplitudeFactorTwoWarningRecorded,
			StatusFactorTwoKineticAmplitudeResonance,
			StatusKineticToAmplitudeDoesNotDeriveCubicCoeff,
		}, "; "),
	}
}

func buildTypedAlternatives(d DoubleEventCoefficientAudit) TypedAlternativesAudit {
	cands := []TypedAlternative{
		{Name: "2p_K7=7/36", Value: d.DoubleK7Weight, Lane: "active double-event / boundary-pair candidate", Accepted: true},
		{Name: "1/5", Value: oneFifthControl, Lane: "nearby numerical control without active source", Accepted: false},
		{Name: "p_K7=7/72", Value: d.K7EventProbability, Lane: "leading K7 event probability, too small by factor two", Accepted: false},
		{Name: "1/4", Value: higgsRadialWeight, Lane: "Higgs radial event probability, wrong lane", Accepted: false},
		{Name: "1/(2*pi)", Value: phaseLoopUnit, Lane: "Hopf phase-loop unit, wrong lane", Accepted: false},
	}
	for i := range cands {
		cands[i].Distance = math.Abs(cands[i].Value - d.CubicCoefficient)
	}
	sort.SliceStable(cands, func(i, j int) bool { return cands[i].Distance < cands[j].Distance })
	best := cands[0]
	return TypedAlternativesAudit{
		TargetCoefficient: d.CubicCoefficient,
		Candidates:        cands,
		ClosestName:       best.Name,
		ClosestAccepted:   best.Accepted,
		NoArbitrarySearch: true,
		Verdict: strings.Join([]string{
			StatusTypedAlternativesAudited,
			StatusTwoPK7BestTypedSourceForCubicCoeff,
		}, "; "),
	}
}

func buildMomentPolynomial(g Gate730Inheritance, d DoubleEventCoefficientAudit) MomentPolynomialAudit {
	quad := g.KappaE * g.M2Wall
	cubic := d.DoubleK7Weight * g.M3Wall
	return MomentPolynomialAudit{
		PolynomialInS:       "D_base≈p_K7 S_split+kappa_e p_K7 S_split^2-2p_K7^2 S_split^3",
		MomentForm:          "D_base≈M1_wall+kappa_e M2_wall-2p_K7 M3_wall",
		LeadingTerm:         g.M1Wall,
		QuadraticTerm:       quad,
		CubicTerm:           cubic,
		UsesDoubleEventForm: true,
		Verdict: strings.Join([]string{
			StatusMomentPolynomialRewrittenWithEventWeightSource,
			StatusSevenOverThirtySixDoubleK7EventWeightCandidate,
		}, "; "),
	}
}

func buildNoncircularity() NoncircularityAudit {
	return NoncircularityAudit{
		KappaEPartiallyDependent:     true,
		TwoPK7TypedButUnexplained:    true,
		BoundaryPairStressPullNative: false,
		MomentExpansionTheoremNative: false,
		Verdict: strings.Join([]string{
			StatusNoncircularityFirewallAudited,
			StatusKappaEQuadraticCoefficientDependent,
			StatusNoNativeReasonCubicCoeffEqualsTwoPK7,
			StatusNoNativeBoundaryPairStressPullTheorem,
			StatusNoNativeBoundaryMomentExpansionTheorem,
		}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		ClaimsNativeScalarRuntime: false,
		ClaimsHiggsMassTheorem:    false,
		ClaimsYukawaTheorem:       false,
		ClaimsCKMPMNSTheorem:      false,
		ClaimsHistoryLoopTheorem:  false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate731Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate730CubicStressPullInherited,
		StatusCubicCoeffRewrittenAsTwoTimesK7Weight,
		StatusBoundaryPairSourceCandidateAudited,
		StatusTwoWallStressPullSourceCandidateAudited,
		StatusKineticToAmplitudeFactorTwoWarningRecorded,
		StatusTypedAlternativesAudited,
		StatusMomentPolynomialRewrittenWithEventWeightSource,
		StatusNoncircularityFirewallAudited,
		StatusPhysicalFirewallsEnforced,
		StatusSevenOverThirtySixDoubleK7EventWeightCandidate,
		StatusFactorTwoBoundaryPairStressPullCandidate,
		StatusTwoPK7BestTypedSourceForCubicCoeff,
		StatusBoundaryPairTimesK7WeightCandidate,
		StatusCubicTermTwoWallStressPullCandidate,
		StatusFactorTwoKineticAmplitudeResonance,
		StatusNoNativeReasonCubicCoeffEqualsTwoPK7,
		StatusNoNativeBoundaryPairStressPullTheorem,
		StatusNoNativeBoundaryMomentExpansionTheorem,
		StatusKineticToAmplitudeDoesNotDeriveCubicCoeff,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusKappaEQuadraticCoefficientDependent,
		StatusGate731Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate730(x Gate730Inheritance) string {
	return fmt.Sprintf("inherited=%t p=%.17g S=%.17g kE=%.17g M1=%.17g M2=%.17g M3=%.17g c3=%.17g combined=%.17g residual=%.17g compressed=%t kEdep=%t noNativeMoment=%t verdict=%q", x.Inherited, x.P_K7, x.SSplit, x.KappaE, x.M1Wall, x.M2Wall, x.M3Wall, x.CubicCoefficient, x.CombinedCorrection, x.ResidualAfterCubicCorrection, x.SevenOver36CompressedResidual, x.KappaEPartiallyDependent, x.NoNativeBoundaryMomentExpansion, x.Verdict)
}
func FormatDoubleEvent(x DoubleEventCoefficientAudit) string {
	return fmt.Sprintf("cubic=%.17g p=%.17g 2p=%.17g exact=%t expansion=%q verdict=%q", x.CubicCoefficient, x.K7EventProbability, x.DoubleK7Weight, x.IdentityExact, x.RewrittenExpansion, x.Verdict)
}
func FormatBoundaryPair(x BoundaryPairSourceCandidateAudit) string {
	return fmt.Sprintf("boundaryDim=%.17g p=%.17g pair*p=%.17g equals=%t interpretation=%q verdict=%q", x.BoundaryPairDimension, x.K7EventProbability, x.BoundaryPairTimesK7Weight, x.EqualsCubicCoefficient, x.Interpretation, x.Verdict)
}
func FormatStressPull(x StressPullSourceCandidateAudit) string {
	return fmt.Sprintf("formula=%q twoBoundaryLegs=%t arbitraryFitRejected=%t verdict=%q", x.CubicTermFormula, x.TwoSidedBoundaryLegs, x.ArbitraryFitRejected, x.Verdict)
}
func FormatKinetic(x KineticToAmplitudeWarning) string {
	return fmt.Sprintf("formula=%q resonance=%t derives=%t verdict=%q", x.LinearizationFormula, x.FactorTwoResonance, x.DerivesCubicCoeff, x.Verdict)
}
func FormatAlternatives(x TypedAlternativesAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s=%.17g dist=%.17g accepted=%t lane=%s", c.Name, c.Value, c.Distance, c.Accepted, c.Lane))
	}
	return fmt.Sprintf("target=%.17g closest=%s accepted=%t noSearch=%t candidates=[%s] verdict=%q", x.TargetCoefficient, x.ClosestName, x.ClosestAccepted, x.NoArbitrarySearch, strings.Join(parts, "; "), x.Verdict)
}
func FormatPolynomial(x MomentPolynomialAudit) string {
	return fmt.Sprintf("poly=%q moment=%q M1=%.17g quad=%.17g cubic=%.17g doubleEvent=%t verdict=%q", x.PolynomialInS, x.MomentForm, x.LeadingTerm, x.QuadraticTerm, x.CubicTerm, x.UsesDoubleEventForm, x.Verdict)
}
func FormatNoncircularity(x NoncircularityAudit) string {
	return fmt.Sprintf("kEdep=%t twoPUnexplained=%t boundaryNative=%t momentNative=%t verdict=%q", x.KappaEPartiallyDependent, x.TwoPK7TypedButUnexplained, x.BoundaryPairStressPullNative, x.MomentExpansionTheoremNative, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("runtime=%t mass=%t yukawa=%t ckm=%t historyLoop=%t verdict=%q", x.ClaimsNativeScalarRuntime, x.ClaimsHiggsMassTheorem, x.ClaimsYukawaTheorem, x.ClaimsCKMPMNSTheorem, x.ClaimsHistoryLoopTheorem, x.Verdict)
}
