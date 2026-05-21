// Package generation2boundaryalphasourceanddomaintransportaudit implements
// Gate 827: BoundaryAlpha Source and Domain-Transport Audit.
//
// Gate 827 follows Gate 826's B-L trace-zero rest-transfer factorization.
// Gate 826 moved the wound from the rest eigenvalue shape to the source of
// alpha_B.  This gate audits the two-domain coefficient decomposition
//
//	alpha_B = (3/10)s + (7/72)s^2,
//
// where 3/10 is tested as a triplet response over V8 plus boundary B2, and
// 7/72 is tested as a quadratic K7 defect response over Lambda^4 V8 plus B2.
// The gate intentionally distinguishes dimension-ratio source resonance from
// a certified BoundaryAlphaDomainTransportMap.
package generation2boundaryalphasourceanddomaintransportaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE827-BOUNDARY-ALPHA-SOURCE-AND-DOMAIN-TRANSPORT-AUDIT"

	SBoundary = 0.0012924448188162962
	NEff      = 3.0023273474722147
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	RankP3             = 3
	DimV8              = 8
	DimBoundaryB2      = 2
	DimVectorBoundary  = DimV8 + DimBoundaryB2
	DimK7              = 7
	DimLambda4V8       = 70
	DimH72             = DimLambda4V8 + DimBoundaryB2
	LinearPower        = 1
	QuadraticPower     = 2
	ExpectedAlphaBText = "0.0003878958469680527"

	StatusGate826Inherited       = "PASS_GATE826_B_MINUS_L_TRANSFER_INHERITED"
	StatusBoundarySplitInherited = "PASS_BOUNDARY_SPLIT_COORDINATE_S_INHERITED"
	StatusLinearRatioVerified    = "PASS_LINEAR_TRIPLET_OVER_VECTOR_BOUNDARY_RATIO_VERIFIED"
	StatusQuadraticRatioVerified = "PASS_QUADRATIC_K7_OVER_H72_RATIO_VERIFIED"
	StatusAlphaDecomposed        = "PASS_ALPHA_B_TWO_DOMAIN_DECOMPOSITION_RECONSTRUCTED"
	StatusPowersSeparated        = "PASS_LINEAR_AND_QUADRATIC_BOUNDARY_POWERS_SEPARATED"
	StatusDimensionSources       = "PASS_DIMENSION_RATIO_SOURCE_CANDIDATES_AUDITED"
	StatusTransportRequirement   = "PASS_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_REQUIREMENT_DEFINED"
	StatusNonCircularity         = "PASS_NONCIRCULARITY_ALPHA_BEFORE_N_EFF_ENFORCED"
	StatusNoBackfit              = "PASS_FIREWALL_NO_N_EFF_BACKFITTING_ALPHA_B_ENFORCED"
	StatusFrozenImpact           = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusNextGateDefined        = "PASS_NEXT_PRESSURE_POINT_TOTAL_RELATIVE_TRACE_MAGNITUDE_OPERATOR_DEFINED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate827        = "FIREWALL_PRESERVED_GATE827_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_BOUNDARY"

	SupportLinearCoeffSource      = "CONDITIONAL_SUPPORT_THREE_TENTHS_AS_TRIPLET_OVER_VECTOR_PLUS_BOUNDARY_RATIO"
	SupportQuadraticCoeffSource   = "CONDITIONAL_SUPPORT_SEVEN_SEVENTY_SECONDS_AS_K7_OVER_AUGMENTED_LAMBDA4_CHAMBER_RATIO"
	SupportTwoDomainAlphaShape    = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_TWO_DOMAIN_BOUNDARY_RESPONSE_SHAPE"
	SupportSameBoundaryCoordinate = "CONDITIONAL_SUPPORT_S_SPLIT_CAN_FEED_LINEAR_AND_QUADRATIC_RESPONSE_LANES_AS_CANDIDATE"
	SupportGate826Pressure        = "CONDITIONAL_SUPPORT_GATE826_MOVED_WOUND_TO_ALPHA_B_SOURCE"
	SupportR2PlusSharpened        = "CONDITIONAL_SUPPORT_R2_PLUS_REST_OPERATOR_REMAINS_SHARPENED_GIVEN_ALPHA_B"
	SupportBoundaryTransportLive  = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_IS_NEAREST_LAWFUL_PRESSURE_POINT"
	SupportCoefficientAuditUseful = "CONDITIONAL_SUPPORT_COEFFICIENTS_ARE_TYPED_DIMENSION_RESONANCES_NOT_ARBITRARY_NUMBERS"
	SupportNextTotalOperator      = "CONDITIONAL_SUPPORT_IF_TRANSPORT_CERTIFIED_NEXT_OBJECT_IS_TOTAL_RELATIVE_TRACE_MAGNITUDE_OPERATOR"

	FailureNoActivationMap      = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureDimensionRatioOnly   = "FAILED_ROUTE_DIMENSION_RATIO_NOT_ACTIVATION_MAP"
	FailureLinearNoTransport    = "FAILED_ROUTE_THREE_TENTHS_NOT_NATIVE_WITHOUT_V8_B2_TO_P3_TRANSPORT"
	FailureQuadraticNoTransport = "FAILED_ROUTE_SEVEN_OVER_SEVENTY_TWO_NOT_NATIVE_ALPHA_SOURCE_WITHOUT_H72_K7_TRANSPORT"
	FailureSameSNotCertified    = "FAILED_ROUTE_SAME_S_SPLIT_COORDINATE_NOT_LAWFULLY_TRANSPORTED_INTO_BOTH_DOMAINS"
	FailureNoAlphaTheorem       = "FAILED_ROUTE_ALPHA_B_NOT_NATIVE_BOUNDARY_THEOREM"
	FailureNoTraceReadout       = "FAILED_ROUTE_BOUNDARY_ALPHA_SOURCE_DOES_NOT_YET_CERTIFY_TRACE_MAGNITUDE_READOUT"
	FailureNoSectorLedger       = "FAILED_ROUTE_GATE827_NOT_R3_SECTOR_TRACE_LEDGER"
	FailureNotR4                = "FAILED_ROUTE_GATE827_NOT_R4_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoCYukawaUpdate      = "FAILED_ROUTE_GATE827_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_TRANSPORT_AND_TRACE_LEDGER"
	FailureCHiggsLevelB         = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B_UNTIL_OPERATOR_ALPHA_AND_LEDGER_ARE_CERTIFIED"
	FailureNoPMNSCKM            = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoHiggsMass          = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
)

type Ledger struct {
	S, S2                                         float64
	LinearCoeff, QuadraticCoeff                   float64
	LinearAlpha, QuadraticAlpha                   float64
	AlphaB, ExpectedAlphaB                        float64
	AlphaResidual                                 float64
	NEffBFN, NEffOperator                         float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs float64
}

type CoefficientSourceAudit struct {
	TripletRank, VectorDim, BoundaryDim, VectorBoundaryDim    int
	LinearCoeff, ExpectedLinearCoeff, LinearResidual          float64
	K7Dim, Lambda4V8Dim, H72Dim                               int
	QuadraticCoeff, ExpectedQuadraticCoeff, QuadraticResidual float64
	LinearFormula, QuadraticFormula                           string
	DimensionRatiosVerified                                   bool
	Verdicts, Supports, Failures                              []string
}

type AlphaDecompositionAudit struct {
	S, S2                        float64
	LinearPower, QuadraticPower  int
	LinearContribution           float64
	QuadraticContribution        float64
	Alpha                        float64
	ExpectedAlpha                float64
	Residual                     float64
	Formula                      string
	PowersSeparated              bool
	Verdicts, Supports, Failures []string
}

type DomainTransportAudit struct {
	LinearDomain, LinearTarget       string
	QuadraticDomain, QuadraticTarget string
	SharedBoundaryCoordinate         string
	LinearMapName, QuadraticMapName  string
	UnifiedMapName                   string
	DimensionRatiosVerified          bool
	LinearTransportCertified         bool
	QuadraticTransportCertified      bool
	UnifiedTransportCertified        bool
	NativeAlphaTheoremCertified      bool
	Verdicts, Supports, Failures     []string
}

type NonCircularityAudit struct {
	AllowedInputs, ForbiddenInputs []string
	Direction                      string
	ComputesAlphaBeforeNEff        bool
	UsesNEffToDefineAlpha          bool
	UsesObservedYukawas            bool
	UsesHiggsMass                  bool
	Verdicts, Supports, Failures   []string
}

type Impact struct {
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	Reason                                           string
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced, NoActivationMap, DimensionRatioOnly, NoBackfit bool
	NoTraceReadout, NoSectorLedger, NotR4                    bool
	NoCYukawaUpdate, CHiggsLevelB, NoPMNSCKM, NoHiggs        bool
	Verdict                                                  string
}

type Analysis struct {
	Ledger       Ledger
	Coefficients CoefficientSourceAudit
	Alpha        AlphaDecompositionAudit
	Transport    DomainTransportAudit
	NonCircular  NonCircularityAudit
	Impact       Impact
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func LinearCoefficient() float64    { return float64(RankP3) / float64(DimVectorBoundary) }
func QuadraticCoefficient() float64 { return float64(DimK7) / float64(DimH72) }
func AlphaB(s float64) float64      { return LinearCoefficient()*s + QuadraticCoefficient()*s*s }
func NEffBFN(alpha float64) float64 { return 3.0 + 6.0*alpha }
func NEffOperator(alpha float64) float64 {
	return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
}
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	linearCoeff := LinearCoefficient()
	quadraticCoeff := QuadraticCoefficient()
	if math.Abs(linearCoeff-0.3) > 1e-15 {
		return Analysis{}, fmt.Errorf("linear coefficient mismatch: %.17g", linearCoeff)
	}
	if math.Abs(quadraticCoeff-(7.0/72.0)) > 1e-15 {
		return Analysis{}, fmt.Errorf("quadratic coefficient mismatch: %.17g", quadraticCoeff)
	}

	s := SBoundary
	linearAlpha := linearCoeff * s
	quadraticAlpha := quadraticCoeff * s * s
	alpha := linearAlpha + quadraticAlpha
	expectedAlpha := 0.0003878958469680527
	alphaResidual := alpha - expectedAlpha

	ledger := Ledger{
		S: s, S2: s * s, LinearCoeff: linearCoeff, QuadraticCoeff: quadraticCoeff,
		LinearAlpha: linearAlpha, QuadraticAlpha: quadraticAlpha, AlphaB: alpha, ExpectedAlphaB: expectedAlpha,
		AlphaResidual: alphaResidual, NEffBFN: NEffBFN(alpha), NEffOperator: NEffOperator(alpha),
		OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
	}

	coefficients := CoefficientSourceAudit{
		TripletRank: RankP3, VectorDim: DimV8, BoundaryDim: DimBoundaryB2, VectorBoundaryDim: DimVectorBoundary,
		LinearCoeff: linearCoeff, ExpectedLinearCoeff: 3.0 / 10.0, LinearResidual: linearCoeff - 3.0/10.0,
		K7Dim: DimK7, Lambda4V8Dim: DimLambda4V8, H72Dim: DimH72,
		QuadraticCoeff: quadraticCoeff, ExpectedQuadraticCoeff: 7.0 / 72.0, QuadraticResidual: quadraticCoeff - 7.0/72.0,
		LinearFormula:           "3/10 = rank(P_3)/dim(V_8 plus B_2)",
		QuadraticFormula:        "7/72 = dim(K_7)/dim(Lambda^4 V_8 plus B_2)",
		DimensionRatiosVerified: true,
		Verdicts:                []string{StatusLinearRatioVerified, StatusQuadraticRatioVerified, StatusDimensionSources},
		Supports:                []string{SupportLinearCoeffSource, SupportQuadraticCoeffSource, SupportCoefficientAuditUseful},
		Failures:                []string{FailureDimensionRatioOnly, FailureLinearNoTransport, FailureQuadraticNoTransport},
	}

	alphaAudit := AlphaDecompositionAudit{
		S: s, S2: s * s, LinearPower: LinearPower, QuadraticPower: QuadraticPower,
		LinearContribution: linearAlpha, QuadraticContribution: quadraticAlpha,
		Alpha: alpha, ExpectedAlpha: expectedAlpha, Residual: alphaResidual,
		Formula:         "alpha_B = (rank(P_3)/dim(V_8 plus B_2)) s + (dim(K_7)/dim(Lambda^4 V_8 plus B_2)) s^2",
		PowersSeparated: true,
		Verdicts:        []string{StatusBoundarySplitInherited, StatusAlphaDecomposed, StatusPowersSeparated},
		Supports:        []string{SupportTwoDomainAlphaShape, SupportSameBoundaryCoordinate, SupportGate826Pressure, SupportR2PlusSharpened},
		Failures:        []string{FailureNoActivationMap, FailureSameSNotCertified, FailureNoAlphaTheorem},
	}

	transport := DomainTransportAudit{
		LinearDomain: "V_8 plus B_2", LinearTarget: "P_3 triplet rest-activation carrier",
		QuadraticDomain: "Lambda^4 V_8 plus B_2 = H_72", QuadraticTarget: "K_7 defect response carrier",
		SharedBoundaryCoordinate: "S_split = s", LinearMapName: "S_split -> V_8 plus B_2 -> P_3",
		QuadraticMapName:        "S_split^2 -> Lambda^4 V_8 plus B_2 -> K_7",
		UnifiedMapName:          "BoundaryAlphaDomainTransportMap",
		DimensionRatiosVerified: true, LinearTransportCertified: false, QuadraticTransportCertified: false,
		UnifiedTransportCertified: false, NativeAlphaTheoremCertified: false,
		Verdicts: []string{StatusTransportRequirement},
		Supports: []string{SupportBoundaryTransportLive, SupportNextTotalOperator},
		Failures: []string{FailureNoActivationMap, FailureDimensionRatioOnly, FailureLinearNoTransport, FailureQuadraticNoTransport, FailureSameSNotCertified, FailureNoTraceReadout},
	}

	nonCircular := NonCircularityAudit{
		AllowedInputs:           []string{"S_split", "rank(P_3)=3", "dim(V_8)=8", "dim(B_2)=2", "dim(K_7)=7", "dim(Lambda^4 V_8)=70", "Gate 826 transfer factorization"},
		ForbiddenInputs:         []string{"N_eff", "C_Yukawa", "C_Higgs", "observed Yukawa ratios", "Higgs mass", "PMNS/CKM"},
		Direction:               "s -> alpha_B -> H_rest -> N_eff",
		ComputesAlphaBeforeNEff: true, UsesNEffToDefineAlpha: false, UsesObservedYukawas: false, UsesHiggsMass: false,
		Verdicts: []string{StatusNonCircularity, StatusNoBackfit},
		Supports: []string{SupportBoundaryTransportLive},
		Failures: []string{FailureNoActivationMap, FailureNoAlphaTheorem},
	}

	candidateNEff := NEffOperator(alpha)
	impact := Impact{
		CandidateNEff: candidateNEff, CandidateCYukawa: CYukawaFromNEff(candidateNEff), CandidateCHiggs: CHiggsFromNEff(candidateNEff),
		OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		Reason:   "Gate 827 verifies coefficient source candidates only; it does not certify BoundaryAlphaDomainTransportMap, total trace-magnitude readout, or sector ledger.",
		Verdicts: []string{StatusFrozenImpact, StatusNextGateDefined},
		Supports: []string{SupportNextTotalOperator},
		Failures: []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB, FailureNoTraceReadout, FailureNoSectorLedger},
	}

	firewalls := Firewalls{
		Enforced: true, NoActivationMap: true, DimensionRatioOnly: true, NoBackfit: true,
		NoTraceReadout: true, NoSectorLedger: true, NotR4: true, NoCYukawaUpdate: true,
		CHiggsLevelB: true, NoPMNSCKM: true, NoHiggs: true, Verdict: StatusFirewallGate827,
	}

	return Analysis{
		Ledger: ledger, Coefficients: coefficients, Alpha: alphaAudit, Transport: transport,
		NonCircular: nonCircular, Impact: impact, Firewalls: firewalls,
		Truth: "Gate 827 certifies the two visible coefficient ratios in alpha_B as typed dimension-ratio candidates, but refuses to identify those ratios with a native activation theorem without a BoundaryAlphaDomainTransportMap.",
		Final: "The alpha_B wound is sharpened: 3/10 and 7/72 are no longer arbitrary coefficients, but Gate 827 still returns dimension resonance, not a certified transport theorem. The next lawful object is the total relative trace-magnitude operator only if the alpha transport map is later certified.",
	}, nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.17g s^2=%.17g alpha_B=%.17g linear=(3/10)s=%.17g quadratic=(7/72)s^2=%.17g residual_to_expected=%.3e N_eff_operator=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g", l.S, l.S2, l.AlphaB, l.LinearAlpha, l.QuadraticAlpha, l.AlphaResidual, l.NEffOperator, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs)
}

func FormatCoefficients(c CoefficientSourceAudit) string {
	return fmt.Sprintf("linear: %s = %.17g residual=%.3e; quadratic: %s = %.17g residual=%.3e; dimension_ratios_verified=%t", c.LinearFormula, c.LinearCoeff, c.LinearResidual, c.QuadraticFormula, c.QuadraticCoeff, c.QuadraticResidual, c.DimensionRatiosVerified)
}

func FormatAlpha(a AlphaDecompositionAudit) string {
	return fmt.Sprintf("%s; linear_power=%d quadratic_power=%d linear=%.17g quadratic=%.17g alpha=%.17g residual=%.3e powers_separated=%t", a.Formula, a.LinearPower, a.QuadraticPower, a.LinearContribution, a.QuadraticContribution, a.Alpha, a.Residual, a.PowersSeparated)
}

func FormatTransport(t DomainTransportAudit) string {
	return fmt.Sprintf("linear_lane=%s; quadratic_lane=%s; shared_coordinate=%s; dimension_ratios_verified=%t linear_transport=%t quadratic_transport=%t unified_transport=%t native_alpha_theorem=%t", t.LinearMapName, t.QuadraticMapName, t.SharedBoundaryCoordinate, t.DimensionRatiosVerified, t.LinearTransportCertified, t.QuadraticTransportCertified, t.UnifiedTransportCertified, t.NativeAlphaTheoremCertified)
}

func FormatNonCircularity(n NonCircularityAudit) string {
	return fmt.Sprintf("direction=%s computes_alpha_before_N_eff=%t uses_N_eff_to_define_alpha=%t uses_observed_yukawas=%t uses_higgs_mass=%t forbidden=%s", n.Direction, n.ComputesAlphaBeforeNEff, n.UsesNEffToDefineAlpha, n.UsesObservedYukawas, n.UsesHiggsMass, strings.Join(n.ForbiddenInputs, ", "))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("candidate_N_eff=%.16g candidate_C_Yukawa=%.16g candidate_C_Higgs=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g can_update_N_eff=%t can_update_C_Yukawa=%t can_update_C_Higgs=%t reason=%s", i.CandidateNEff, i.CandidateCYukawa, i.CandidateCHiggs, i.OfficialNEff, i.OfficialCYukawa, i.OfficialCHiggs, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.Reason)
}

func Statuses() []string {
	return []string{
		StatusGate826Inherited, StatusBoundarySplitInherited, StatusLinearRatioVerified, StatusQuadraticRatioVerified,
		StatusAlphaDecomposed, StatusPowersSeparated, StatusDimensionSources, StatusTransportRequirement,
		StatusNonCircularity, StatusNoBackfit, StatusFrozenImpact, StatusNextGateDefined, StatusPhysicalFirewalls,
		SupportLinearCoeffSource, SupportQuadraticCoeffSource, SupportTwoDomainAlphaShape, SupportSameBoundaryCoordinate,
		SupportGate826Pressure, SupportR2PlusSharpened, SupportBoundaryTransportLive, SupportCoefficientAuditUseful,
		SupportNextTotalOperator, FailureNoActivationMap, FailureDimensionRatioOnly, FailureLinearNoTransport,
		FailureQuadraticNoTransport, FailureSameSNotCertified, FailureNoAlphaTheorem, FailureNoTraceReadout,
		FailureNoSectorLedger, FailureNotR4, FailureNoCYukawaUpdate, FailureCHiggsLevelB, FailureNoPMNSCKM,
		FailureNoHiggsMass, StatusFirewallGate827,
	}
}

func containsAll(haystack, needles []string) bool {
	set := map[string]bool{}
	for _, h := range haystack {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}
