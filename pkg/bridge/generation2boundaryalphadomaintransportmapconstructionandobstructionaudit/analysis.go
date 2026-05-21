// Package generation2boundaryalphadomaintransportmapconstructionandobstructionaudit
// implements Gate 828: BoundaryAlphaDomainTransportMap Construction and
// Obstruction Audit.
//
// Gate 828 follows Gate 827's controlled source-typing success.  Gate 827
// verified the two visible coefficient ratios in
//
//	alpha_B = (3/10)s + (7/72)s^2
//
// but refused to promote those ratios into a native activation theorem.  This
// gate attacks the named missing object directly: it audits whether the same
// boundary split coordinate S_split lawfully transports into two normalized
// domains,
//
//	S_split   -> V_8 plus B_2              -> P_3,
//	S_split^2 -> Lambda^4 V_8 plus B_2    -> K_7,
//
// or whether the result is still only a typed bridge-rule / dimension-ratio
// resonance.  The gate deliberately preserves the firewall: normalized support
// traces are not activation maps without a concrete typed transport law.
package generation2boundaryalphadomaintransportmapconstructionandobstructionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE828-BOUNDARY-ALPHA-DOMAIN-TRANSPORT-MAP-CONSTRUCTION-OBSTRUCTION-AUDIT"

	SBoundary = 0.0012924448188162962
	NEff      = 3.0023273474722147
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	RankP3        = 3
	DimV8         = 8
	DimBoundaryB2 = 2
	DimV8B2       = DimV8 + DimBoundaryB2
	DimK7         = 7
	DimLambda4V8  = 70
	DimH72        = DimLambda4V8 + DimBoundaryB2

	StatusGate827Inherited       = "PASS_GATE827_ALPHA_SOURCE_TYPING_INHERITED"
	StatusCandidateLawRebuilt    = "PASS_BOUNDARY_ALPHA_CANDIDATE_LAW_REBUILT"
	StatusSupportTraceWeights    = "PASS_NORMALIZED_SUPPORT_TRACE_WEIGHTS_VERIFIED"
	StatusLinearLaneSpecified    = "PASS_LINEAR_VECTOR_BOUNDARY_TRIPLET_LANE_SPECIFIED"
	StatusQuadraticLaneSpecified = "PASS_QUADRATIC_H72_K7_DEFECT_LANE_SPECIFIED"
	StatusCriteriaEvaluated      = "PASS_TRANSPORT_MAP_CERTIFICATION_CRITERIA_EVALUATED"
	StatusNoConcreteTransport    = "PASS_NO_CONCRETE_TYPED_TRANSPORT_MAP_FOUND"
	StatusDimensionResonance     = "PASS_CLASSIFIED_AS_DIMENSION_RATIO_RESONANCE_NOT_ACTIVATION_THEOREM"
	StatusNonCircularity         = "PASS_FIREWALL_NO_N_EFF_BACKFITTING_ALPHA_B_ENFORCED"
	StatusNoOperatorPromotion    = "PASS_TOTAL_OPERATOR_PROMOTION_DEFERRED_UNTIL_ALPHA_TRANSPORT_CERTIFIED"
	StatusFrozenImpact           = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusNextGateDefined        = "PASS_NEXT_PRESSURE_POINT_TOTAL_RELATIVE_TRACE_MAGNITUDE_OPERATOR_CONSOLIDATION_DEFINED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate828        = "FIREWALL_PRESERVED_GATE828_BOUNDARY_ALPHA_TRANSPORT_MAP_BOUNDARY"

	SupportTwoLaneSourceShape       = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_TWO_LANE_SUPPORT_TRACE_SHAPE"
	SupportLinearTraceWeight        = "CONDITIONAL_SUPPORT_LINEAR_WEIGHT_IS_TRIPLET_SUPPORT_TRACE_OVER_V8_PLUS_B2"
	SupportQuadraticTraceWeight     = "CONDITIONAL_SUPPORT_QUADRATIC_WEIGHT_IS_K7_SUPPORT_TRACE_OVER_H72"
	SupportBridgeRuleCandidate      = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_AS_BRIDGE_RULE_CANDIDATE"
	SupportSameSAsPressurePoint     = "CONDITIONAL_SUPPORT_SAME_S_SPLIT_LINEAR_AND_QUADRATIC_POWER_SPLIT_IS_THE_LIVE_PRESSURE_POINT"
	SupportGate826StillValid        = "CONDITIONAL_SUPPORT_GATE826_B_MINUS_L_REST_TRANSFER_REMAINS_VALID_GIVEN_ALPHA_B"
	SupportGate827StillValid        = "CONDITIONAL_SUPPORT_GATE827_COEFFICIENT_SOURCE_TYPING_REMAINS_VALID"
	SupportNextTotalOperatorUseful  = "CONDITIONAL_SUPPORT_TOTAL_RELATIVE_TRACE_MAGNITUDE_OPERATOR_AUDIT_IS_USEFUL_AS_CONSOLIDATION"
	SupportObstructionIsInformative = "CONDITIONAL_SUPPORT_OBSTRUCTION_SHARPENS_MISSING_OBJECT_TO_TYPED_TRANSPORT_FUNCTOR_OR_VARIATIONAL_LAW"

	FailureNoBoundaryAlphaMap      = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureDimensionRatioOnly      = "FAILED_ROUTE_DIMENSION_RATIO_NOT_ACTIVATION_MAP"
	FailureNoLinearTransport       = "FAILED_ROUTE_NO_TYPED_S_SPLIT_TO_V8_B2_TO_P3_LINEAR_TRANSPORT"
	FailureNoQuadraticTransport    = "FAILED_ROUTE_NO_TYPED_S_SPLIT_SQUARED_TO_H72_TO_K7_QUADRATIC_TRANSPORT"
	FailureNoSharedFunctor         = "FAILED_ROUTE_NO_SHARED_FUNCTOR_TRANSPORTS_S_SPLIT_INTO_BOTH_DOMAINS"
	FailurePowerLawNotDerived      = "FAILED_ROUTE_LINEAR_VS_QUADRATIC_RESPONSE_ORDER_NOT_DERIVED"
	FailureNoVariationalPrinciple  = "FAILED_ROUTE_NO_VARIATIONAL_OR_TRACE_ACTION_PRINCIPLE_DERIVES_ALPHA_POLYNOMIAL"
	FailureNoNativeAlphaTheorem    = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_RULE_NOT_NATIVE_BOUNDARY_THEOREM"
	FailureNoTraceMagnitudeReadout = "FAILED_ROUTE_NO_TOTAL_TRACE_MAGNITUDE_READOUT_CERTIFIED_FROM_ALPHA_TRANSPORT"
	FailureNoSectorLedger          = "FAILED_ROUTE_GATE828_NOT_R3_SECTOR_TRACE_LEDGER"
	FailureNotR4                   = "FAILED_ROUTE_GATE828_NOT_R4_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoCYukawaUpdate         = "FAILED_ROUTE_GATE828_DOES_NOT_UPDATE_C_YUKAWA_OR_C_HIGGS"
	FailureNoPMNSCKM               = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoHiggsMass             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
)

type Ledger struct {
	S, S2                                            float64
	LinearWeight, QuadraticWeight                    float64
	LinearAlpha, QuadraticAlpha                      float64
	AlphaB, ExpectedAlphaB, AlphaResidual            float64
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
}

type SupportTraceWeightAudit struct {
	LinearNumerator, LinearDenominator       int
	QuadraticNumerator, QuadraticDenominator int
	LinearWeight, ExpectedLinearWeight       float64
	QuadraticWeight, ExpectedQuadraticWeight float64
	LinearResidual, QuadraticResidual        float64
	LinearFormula, QuadraticFormula          string
	WeightsVerified                          bool
	Verdicts, Supports, Failures             []string
}

type Lane struct {
	Name                         string
	SourcePower                  int
	SourceObject                 string
	Domain                       string
	Support                      string
	WeightFormula                string
	Weight                       float64
	Contribution                 float64
	CarrierTyped                 bool
	SupportProjectorTyped        bool
	ConcreteTransportMap         bool
	ResponseOrderDerived         bool
	Verdicts, Supports, Failures []string
}

type TransportCriteria struct {
	CandidateMapName               string
	SourceCoordinate               string
	HasSourceScalar                bool
	HasTypedTargetCarriers         bool
	HasSupportTraceWeights         bool
	HasConcreteLinearMap           bool
	HasConcreteQuadraticMap        bool
	HasSharedFunctor               bool
	HasPowerLawDerivation          bool
	HasVariationalPrinciple        bool
	HasNonCircularDirection        bool
	CertifiesNativeAlphaTheorem    bool
	CertifiesTraceMagnitudeReadout bool
	Classification                 string
	Verdicts, Supports, Failures   []string
}

type NonCircularityAudit struct {
	AllowedInputs, ForbiddenInputs []string
	Direction                      string
	UsesNEffToDefineAlpha          bool
	UsesObservedYukawas            bool
	UsesHiggsMass                  bool
	ComputesAlphaBeforeReadout     bool
	Verdicts, Supports, Failures   []string
}

type Impact struct {
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	CanPromoteTotalOperator, CanUpdateNEff           bool
	CanUpdateCYukawa, CanUpdateCHiggs                bool
	Reason                                           string
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced, NoBoundaryAlphaMap, DimensionRatioOnly bool
	NoLinearTransport, NoQuadraticTransport          bool
	NoSharedFunctor, PowerLawNotDerived              bool
	NoTraceMagnitudeReadout, NoSectorLedger, NotR4   bool
	NoCYukawaUpdate, NoPMNSCKM, NoHiggs              bool
	Verdict                                          string
}

type Analysis struct {
	Ledger        Ledger
	Weights       SupportTraceWeightAudit
	LinearLane    Lane
	QuadraticLane Lane
	Criteria      TransportCriteria
	NonCircular   NonCircularityAudit
	Impact        Impact
	Firewalls     Firewalls
	Truth         string
	Final         string
}

func LinearWeight() float64    { return float64(RankP3) / float64(DimV8B2) }
func QuadraticWeight() float64 { return float64(DimK7) / float64(DimH72) }
func AlphaB(s float64) float64 { return LinearWeight()*s + QuadraticWeight()*s*s }
func NEffOperator(alpha float64) float64 {
	return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
}
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	s := SBoundary
	linearWeight := LinearWeight()
	quadraticWeight := QuadraticWeight()
	expectedLinear := 3.0 / 10.0
	expectedQuadratic := 7.0 / 72.0
	if math.Abs(linearWeight-expectedLinear) > 1e-15 {
		return Analysis{}, fmt.Errorf("linear support trace weight mismatch: %.17g", linearWeight)
	}
	if math.Abs(quadraticWeight-expectedQuadratic) > 1e-15 {
		return Analysis{}, fmt.Errorf("quadratic support trace weight mismatch: %.17g", quadraticWeight)
	}
	linearAlpha := linearWeight * s
	quadraticAlpha := quadraticWeight * s * s
	alpha := linearAlpha + quadraticAlpha
	expectedAlpha := 0.0003878958469680527
	candidateNEff := NEffOperator(alpha)

	ledger := Ledger{
		S: s, S2: s * s, LinearWeight: linearWeight, QuadraticWeight: quadraticWeight,
		LinearAlpha: linearAlpha, QuadraticAlpha: quadraticAlpha, AlphaB: alpha, ExpectedAlphaB: expectedAlpha,
		AlphaResidual: alpha - expectedAlpha, CandidateNEff: candidateNEff,
		CandidateCYukawa: CYukawaFromNEff(candidateNEff), CandidateCHiggs: CHiggsFromNEff(candidateNEff),
		OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
	}

	weights := SupportTraceWeightAudit{
		LinearNumerator: RankP3, LinearDenominator: DimV8B2, QuadraticNumerator: DimK7, QuadraticDenominator: DimH72,
		LinearWeight: linearWeight, ExpectedLinearWeight: expectedLinear, LinearResidual: linearWeight - expectedLinear,
		QuadraticWeight: quadraticWeight, ExpectedQuadraticWeight: expectedQuadratic, QuadraticResidual: quadraticWeight - expectedQuadratic,
		LinearFormula:    "w_3|10 = Tr(P_3)/dim(V_8 plus B_2) = 3/10",
		QuadraticFormula: "w_7|72 = Tr(P_K7)/dim(Lambda^4 V_8 plus B_2) = 7/72",
		WeightsVerified:  true,
		Verdicts:         []string{StatusGate827Inherited, StatusCandidateLawRebuilt, StatusSupportTraceWeights},
		Supports:         []string{SupportLinearTraceWeight, SupportQuadraticTraceWeight, SupportGate827StillValid, SupportTwoLaneSourceShape},
		Failures:         []string{FailureDimensionRatioOnly, FailureNoBoundaryAlphaMap},
	}

	linearLane := Lane{
		Name: "linear vector-boundary triplet lane", SourcePower: 1, SourceObject: "S_split", Domain: "V_8 plus B_2",
		Support: "P_3 triplet support", WeightFormula: "Tr(P_3)/dim(V_8 plus B_2)", Weight: linearWeight,
		Contribution: linearAlpha, CarrierTyped: true, SupportProjectorTyped: true, ConcreteTransportMap: false, ResponseOrderDerived: false,
		Verdicts: []string{StatusLinearLaneSpecified},
		Supports: []string{SupportLinearTraceWeight, SupportBridgeRuleCandidate},
		Failures: []string{FailureNoLinearTransport, FailurePowerLawNotDerived, FailureNoVariationalPrinciple},
	}

	quadraticLane := Lane{
		Name: "quadratic H72 K7 defect lane", SourcePower: 2, SourceObject: "S_split^2", Domain: "Lambda^4 V_8 plus B_2 = H_72",
		Support: "K_7 defect support", WeightFormula: "Tr(P_K7)/dim(H_72)", Weight: quadraticWeight,
		Contribution: quadraticAlpha, CarrierTyped: true, SupportProjectorTyped: true, ConcreteTransportMap: false, ResponseOrderDerived: false,
		Verdicts: []string{StatusQuadraticLaneSpecified},
		Supports: []string{SupportQuadraticTraceWeight, SupportBridgeRuleCandidate},
		Failures: []string{FailureNoQuadraticTransport, FailurePowerLawNotDerived, FailureNoVariationalPrinciple},
	}

	criteria := TransportCriteria{
		CandidateMapName: "BoundaryAlphaDomainTransportMap", SourceCoordinate: "S_split = s",
		HasSourceScalar: true, HasTypedTargetCarriers: true, HasSupportTraceWeights: true,
		HasConcreteLinearMap: false, HasConcreteQuadraticMap: false, HasSharedFunctor: false,
		HasPowerLawDerivation: false, HasVariationalPrinciple: false, HasNonCircularDirection: true,
		CertifiesNativeAlphaTheorem: false, CertifiesTraceMagnitudeReadout: false,
		Classification: "BRIDGE_RULE_CANDIDATE_AND_DIMENSION_RATIO_RESONANCE_NOT_CERTIFIED_TRANSPORT_MAP",
		Verdicts:       []string{StatusCriteriaEvaluated, StatusNoConcreteTransport, StatusDimensionResonance},
		Supports:       []string{SupportSameSAsPressurePoint, SupportBridgeRuleCandidate, SupportObstructionIsInformative, SupportGate826StillValid},
		Failures:       []string{FailureNoBoundaryAlphaMap, FailureNoLinearTransport, FailureNoQuadraticTransport, FailureNoSharedFunctor, FailurePowerLawNotDerived, FailureNoVariationalPrinciple, FailureNoNativeAlphaTheorem, FailureNoTraceMagnitudeReadout},
	}

	nonCircular := NonCircularityAudit{
		AllowedInputs:         []string{"S_split", "rank(P_3)=3", "dim(V_8 plus B_2)=10", "dim(K_7)=7", "dim(H_72)=72", "Gate 826 B-L rest transfer", "Gate 827 coefficient source typing"},
		ForbiddenInputs:       []string{"N_eff", "C_Yukawa", "C_Higgs", "observed Yukawa ratios", "Higgs mass", "PMNS/CKM", "sector assignment"},
		Direction:             "s -> candidate alpha_B -> Gate 826 H_rest; N_eff remains downstream diagnostic only",
		UsesNEffToDefineAlpha: false, UsesObservedYukawas: false, UsesHiggsMass: false, ComputesAlphaBeforeReadout: true,
		Verdicts: []string{StatusNonCircularity},
		Supports: []string{SupportBridgeRuleCandidate},
		Failures: []string{FailureNoBoundaryAlphaMap},
	}

	impact := Impact{
		CandidateNEff: candidateNEff, CandidateCYukawa: CYukawaFromNEff(candidateNEff), CandidateCHiggs: CHiggsFromNEff(candidateNEff),
		OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		CanPromoteTotalOperator: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		Reason:   "Gate 828 finds typed support-trace weights and a noncircular bridge-rule candidate, but no concrete BoundaryAlphaDomainTransportMap, no response-order theorem, and no total trace-magnitude readout certification.",
		Verdicts: []string{StatusNoOperatorPromotion, StatusFrozenImpact, StatusNextGateDefined},
		Supports: []string{SupportNextTotalOperatorUseful},
		Failures: []string{FailureNoTraceMagnitudeReadout, FailureNoSectorLedger, FailureNoCYukawaUpdate},
	}

	firewalls := Firewalls{
		Enforced: true, NoBoundaryAlphaMap: true, DimensionRatioOnly: true, NoLinearTransport: true,
		NoQuadraticTransport: true, NoSharedFunctor: true, PowerLawNotDerived: true,
		NoTraceMagnitudeReadout: true, NoSectorLedger: true, NotR4: true, NoCYukawaUpdate: true,
		NoPMNSCKM: true, NoHiggs: true, Verdict: StatusFirewallGate828,
	}

	analysis := Analysis{
		Ledger: ledger, Weights: weights, LinearLane: linearLane, QuadraticLane: quadraticLane,
		Criteria: criteria, NonCircular: nonCircular, Impact: impact, Firewalls: firewalls,
		Truth: "Gate 828 tests the named missing object directly.  The two normalized support-trace weights survive as a coherent bridge-rule candidate, but no typed transport functor or variational law is present to make S_split activate linearly in V_8 plus B_2 and quadratically in H_72.",
		Final: "The alpha_B formula is now classified more sharply: not arbitrary, not native, but a noncircular two-lane support-trace resonance awaiting a BoundaryAlphaDomainTransportMap.  The next useful gate can consolidate the total relative trace-magnitude operator while preserving the alpha-source firewall.",
	}
	if err := validate(analysis); err != nil {
		return Analysis{}, err
	}
	return analysis, nil
}

func validate(a Analysis) error {
	if !a.Weights.WeightsVerified || math.Abs(a.Weights.LinearWeight-0.3) > 1e-15 || math.Abs(a.Weights.QuadraticWeight-7.0/72.0) > 1e-15 {
		return fmt.Errorf("support trace weights failed: %s", FormatWeights(a.Weights))
	}
	if math.Abs(a.Ledger.AlphaB-a.Ledger.ExpectedAlphaB) > 1e-18 || math.Abs(a.Ledger.LinearAlpha+a.Ledger.QuadraticAlpha-a.Ledger.AlphaB) > 1e-21 {
		return fmt.Errorf("candidate alpha reconstruction failed: %s", FormatLedger(a.Ledger))
	}
	if a.LinearLane.ConcreteTransportMap || a.QuadraticLane.ConcreteTransportMap || a.Criteria.HasConcreteLinearMap || a.Criteria.HasConcreteQuadraticMap || a.Criteria.HasSharedFunctor || a.Criteria.CertifiesNativeAlphaTheorem {
		return fmt.Errorf("transport over-promoted: %s", FormatCriteria(a.Criteria))
	}
	if !a.Criteria.HasSourceScalar || !a.Criteria.HasTypedTargetCarriers || !a.Criteria.HasSupportTraceWeights || !a.Criteria.HasNonCircularDirection {
		return fmt.Errorf("candidate criteria under-specified: %s", FormatCriteria(a.Criteria))
	}
	if a.NonCircular.UsesNEffToDefineAlpha || a.NonCircular.UsesObservedYukawas || a.NonCircular.UsesHiggsMass || !a.NonCircular.ComputesAlphaBeforeReadout {
		return fmt.Errorf("noncircularity failed: %s", FormatNonCircularity(a.NonCircular))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoBoundaryAlphaMap || !a.Firewalls.DimensionRatioOnly || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoSectorLedger || !a.Firewalls.NotR4 || !a.Firewalls.NoCYukawaUpdate {
		return fmt.Errorf("firewall failed: %s", a.Firewalls.Verdict)
	}
	return nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.17g s^2=%.17g alpha_B=%.17g linear=%.17g quadratic=%.17g residual_to_expected=%.3e candidate_N_eff=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g", l.S, l.S2, l.AlphaB, l.LinearAlpha, l.QuadraticAlpha, l.AlphaResidual, l.CandidateNEff, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs)
}

func FormatWeights(w SupportTraceWeightAudit) string {
	return fmt.Sprintf("%s -> %.17g residual=%.3e; %s -> %.17g residual=%.3e; verified=%t", w.LinearFormula, w.LinearWeight, w.LinearResidual, w.QuadraticFormula, w.QuadraticWeight, w.QuadraticResidual, w.WeightsVerified)
}

func FormatLane(l Lane) string {
	return fmt.Sprintf("%s: source=%s power=%d domain=%s support=%s weight=%s=%.17g contribution=%.17g carrier_typed=%t support_typed=%t concrete_transport=%t response_order_derived=%t", l.Name, l.SourceObject, l.SourcePower, l.Domain, l.Support, l.WeightFormula, l.Weight, l.Contribution, l.CarrierTyped, l.SupportProjectorTyped, l.ConcreteTransportMap, l.ResponseOrderDerived)
}

func FormatCriteria(c TransportCriteria) string {
	return fmt.Sprintf("map=%s source=%s source_scalar=%t target_carriers=%t support_weights=%t concrete_linear=%t concrete_quadratic=%t shared_functor=%t power_law=%t variational=%t native_alpha=%t trace_readout=%t classification=%s", c.CandidateMapName, c.SourceCoordinate, c.HasSourceScalar, c.HasTypedTargetCarriers, c.HasSupportTraceWeights, c.HasConcreteLinearMap, c.HasConcreteQuadraticMap, c.HasSharedFunctor, c.HasPowerLawDerivation, c.HasVariationalPrinciple, c.CertifiesNativeAlphaTheorem, c.CertifiesTraceMagnitudeReadout, c.Classification)
}

func FormatNonCircularity(n NonCircularityAudit) string {
	return fmt.Sprintf("direction=%s computes_alpha_before_readout=%t uses_N_eff_to_define_alpha=%t uses_observed_yukawas=%t uses_higgs_mass=%t forbidden=%s", n.Direction, n.ComputesAlphaBeforeReadout, n.UsesNEffToDefineAlpha, n.UsesObservedYukawas, n.UsesHiggsMass, strings.Join(n.ForbiddenInputs, ", "))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("candidate_N_eff=%.16g candidate_C_Yukawa=%.16g candidate_C_Higgs=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g promote_total_operator=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t reason=%s", i.CandidateNEff, i.CandidateCYukawa, i.CandidateCHiggs, i.OfficialNEff, i.OfficialCYukawa, i.OfficialCHiggs, i.CanPromoteTotalOperator, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.Reason)
}

func Statuses() []string {
	return []string{
		StatusGate827Inherited, StatusCandidateLawRebuilt, StatusSupportTraceWeights, StatusLinearLaneSpecified,
		StatusQuadraticLaneSpecified, StatusCriteriaEvaluated, StatusNoConcreteTransport, StatusDimensionResonance,
		StatusNonCircularity, StatusNoOperatorPromotion, StatusFrozenImpact, StatusNextGateDefined, StatusPhysicalFirewalls,
		SupportTwoLaneSourceShape, SupportLinearTraceWeight, SupportQuadraticTraceWeight, SupportBridgeRuleCandidate,
		SupportSameSAsPressurePoint, SupportGate826StillValid, SupportGate827StillValid, SupportNextTotalOperatorUseful,
		SupportObstructionIsInformative, FailureNoBoundaryAlphaMap, FailureDimensionRatioOnly, FailureNoLinearTransport,
		FailureNoQuadraticTransport, FailureNoSharedFunctor, FailurePowerLawNotDerived, FailureNoVariationalPrinciple,
		FailureNoNativeAlphaTheorem, FailureNoTraceMagnitudeReadout, FailureNoSectorLedger, FailureNotR4,
		FailureNoCYukawaUpdate, FailureNoPMNSCKM, FailureNoHiggsMass, StatusFirewallGate828,
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
