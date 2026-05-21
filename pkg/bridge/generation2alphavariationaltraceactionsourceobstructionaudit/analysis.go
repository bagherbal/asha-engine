// Package generation2alphavariationaltraceactionsourceobstructionaudit
// implements Gate 830: Alpha Variational / Trace-Action Source Obstruction Audit.
//
// Gate 830 follows Gate 829's clean aggregate relative trace-magnitude
// consolidation.  The downstream chain
//
//	alpha_B -> H_rest -> H_total -> operator_N_eff
//
// is now coherent given sealed alpha_B.  The only live wound is the upstream
// source law
//
//	S_split -> alpha_B = (3/10)s + (7/72)s^2.
//
// Gate 830 tests whether this law can be promoted from a two-lane
// support-trace bridge rule into a native finite trace-action or variational
// source.  It deliberately accepts the formal trace reconstruction, but rejects
// it as a theorem when the scalar insertions X1(s)=sI and X2(s)=s^2I, the
// response-order split, and the variational action are not independently typed.
package generation2alphavariationaltraceactionsourceobstructionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE830-ALPHA-VARIATIONAL-TRACE-ACTION-SOURCE-OBSTRUCTION-AUDIT"

	SBoundary = 0.0012924448188162962
	CHistory  = 1.038025177923625

	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	RankP3  = 3
	DimV8B2 = 10
	DimK7   = 7
	DimH72  = 72

	StatusGate829Inherited       = "PASS_GATE829_TOTAL_OPERATOR_CONSOLIDATION_INHERITED"
	StatusLiveWoundSelected      = "PASS_LIVE_WOUND_SELECTED_S_SPLIT_TO_ALPHA_B"
	StatusTraceAnsatzBuilt       = "PASS_FORMAL_TRACE_EXPANSION_RECONSTRUCTS_ALPHA_RULE"
	StatusTraceAnsatzRestatement = "PASS_TRACE_EXPANSION_CLASSIFIED_AS_RESTATEMENT_NOT_SOURCE"
	StatusResponseOrderAudited   = "PASS_LINEAR_AND_QUADRATIC_RESPONSE_ORDER_AUDITED"
	StatusResponseOrderOpen      = "PASS_RESPONSE_ORDER_SOURCE_REMAINS_OPEN"
	StatusVariationalAudited     = "PASS_FORMAL_VARIATIONAL_STATIONARITY_AUDITED"
	StatusVariationalRejected    = "PASS_VARIATIONAL_ACTION_CLASSIFIED_AS_FORMAL_REPACKAGING"
	StatusNonCircularity         = "PASS_NONCIRCULAR_ALPHA_SOURCE_FIREWALL_ENFORCED"
	StatusNoSealReduction        = "PASS_N_EFF_SEAL_REDUCTION_BLOCKED_AFTER_ALPHA_SOURCE_OBSTRUCTION"
	StatusNextGateDefined        = "PASS_NEXT_PRESSURE_POINT_SECTOR_LEDGER_FIREWALL_DEFINED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate830        = "FIREWALL_PRESERVED_GATE830_ALPHA_VARIATIONAL_TRACE_ACTION_SOURCE_OBSTRUCTION"

	SupportTraceRuleReconstructsAlpha = "CONDITIONAL_SUPPORT_SUPPORT_TRACE_RULE_RECONSTRUCTS_ALPHA_B_EXACTLY"
	SupportTwoLaneSourceTyping        = "CONDITIONAL_SUPPORT_TWO_LANE_SUPPORT_TRACE_ANATOMY_REMAINS_VALID"
	SupportFirstSecondOrderStory      = "CONDITIONAL_SUPPORT_LINEAR_AS_FIRST_ORDER_AND_QUADRATIC_AS_SECOND_ORDER_RESPONSE_CANDIDATE"
	SupportFormalActionStationarity   = "CONDITIONAL_SUPPORT_FORMAL_ACTION_HAS_STATIONARY_ALPHA_RULE_IF_RULE_INSERTED"
	SupportGate826829Downstream       = "CONDITIONAL_SUPPORT_DOWNSTREAM_OPERATOR_CHAIN_REMAINS_VALID_GIVEN_ALPHA_B"
	SupportAlphaMarkedSealed          = "CONDITIONAL_SUPPORT_ALPHA_B_SHOULD_BE_MARKED_SEALED_BRIDGE_RESPONSE"
	SupportR2PlusPlusStatus           = "CONDITIONAL_SUPPORT_STATUS_REMAINS_R2_PLUS_PLUS_NOT_R3"

	FailureTraceExpansionRestatesRule     = "FAILED_ROUTE_TRACE_EXPANSION_RESTATES_ALPHA_RULE"
	FailureX1NotNative                    = "FAILED_ROUTE_X1_EQUALS_S_I_NOT_NATIVELY_PRODUCED_BY_BOUNDARY_SPLIT"
	FailureX2NotNative                    = "FAILED_ROUTE_X2_EQUALS_S_SQUARED_I_NOT_NATIVELY_PRODUCED_BY_BOUNDARY_SPLIT"
	FailureResponseOrderNotDerived        = "FAILED_ROUTE_LINEAR_AND_QUADRATIC_RESPONSE_ORDER_NOT_DERIVED"
	FailureNoBoundaryTraceAction          = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_TRACE_ACTION_FUNCTIONAL_CERTIFIED"
	FailureVariationalRepackaging         = "FAILED_ROUTE_VARIATIONAL_ACTION_IS_FORMAL_REPACKAGING"
	FailureNoEulerLagrangeAlphaTheorem    = "FAILED_ROUTE_NO_NATIVE_EULER_LAGRANGE_ALPHA_THEOREM"
	FailureNoBoundaryAlphaDomainTransport = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureAlphaNotNative                 = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_BRIDGE_RESPONSE_NOT_NATIVE_THEOREM"
	FailureNoNEffSealReduction            = "FAILED_ROUTE_N_EFF_SEAL_REDUCTION_NOT_ALLOWED_AFTER_ALPHA_SOURCE_OBSTRUCTION"
	FailureNoCYukawaUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3SectorLedger              = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3_SECTOR_TRACE_LEDGER"
	FailureNotR4Yukawa                    = "FAILED_ROUTE_NOT_R4_NATIVE_YUKAWA_THEOREM"
	FailureNoSectorAssignment             = "FAILED_ROUTE_NO_STANDARD_MODEL_SECTOR_ASSIGNMENT"
	FailureNoPMNSCKM                      = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoHiggsMass                    = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
)

type Ledger struct {
	S, S2                               float64
	LinearWeight, QuadraticWeight       float64
	LinearAlpha, QuadraticAlpha, AlphaB float64
	ExpectedAlphaB, AlphaResidual       float64
	OperatorNEff, OfficialNEff          float64
	OperatorCYukawa, OperatorCHiggs     float64
	OfficialCYukawa, OfficialCHiggs     float64
	AlphaSealedBridgeResponse           bool
}

type TraceExpansionAudit struct {
	Expression                                string
	X1, X2                                    string
	LinearTrace, QuadraticTrace               float64
	LinearContribution, QuadraticContribution float64
	AlphaFromTrace, ExpectedAlpha             float64
	ReconstructsAlpha                         bool
	X1NaturallyProduced, X2NaturallyProduced  bool
	TraceActionCertified                      bool
	ClassifiedAsRestatement                   bool
	Verdicts, Supports, Failures              []string
}

type ResponseOrderAudit struct {
	LinearLane, QuadraticLane                        string
	LinearPower, QuadraticPower                      int
	LinearOrderCandidate, QuadraticOrderCandidate    string
	LinearOrderDerived, QuadraticOrderDerived        bool
	FirstSecondOrderInterpretationAllowedAsCandidate bool
	ResponseOrderTheoremCertified                    bool
	Verdicts, Supports, Failures                     []string
}

type VariationalAudit struct {
	ActionName, ActionExpression, StationaryCondition string
	StationaryAlpha, ExpectedAlpha                    float64
	StationarityWorksFormally                         bool
	WeightsAllTraceSourced                            bool
	PowersTypedByResponseOrder                        bool
	ActionNative                                      bool
	UsesInsertedAlphaRule                             bool
	IsFormalRepackaging                               bool
	CertifiesAlphaTheorem                             bool
	Verdicts, Supports, Failures                      []string
}

type NonCircularityAudit struct {
	AllowedInputs, ForbiddenInputs []string
	Direction                      string
	UsesNEffToDefineAlpha          bool
	UsesOfficialLedger             bool
	UsesObservedYukawas            bool
	UsesHiggsMass                  bool
	ComputesAlphaBeforeReadout     bool
	Verdicts, Supports, Failures   []string
}

type Impact struct {
	OperatorNEff, OfficialNEff                       float64
	CanPromoteAlpha, CanPromoteOperatorNEff          bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	RecommendedAlphaStatus, NextGate                 string
	Reason                                           string
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced, TraceExpansionRestatement  bool
	X1X2NotNative, ResponseOrderOpen     bool
	NoTraceAction, VariationalRepackage  bool
	AlphaSealed, NoBoundaryAlphaMap      bool
	NoNEffSealReduction, NoCYukawaUpdate bool
	NotR3, NotR4, NoSectorAssignment     bool
	NoPMNSCKM, NoHiggs                   bool
	Verdict                              string
}

type Analysis struct {
	Ledger        Ledger
	Trace         TraceExpansionAudit
	ResponseOrder ResponseOrderAudit
	Variational   VariationalAudit
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

func FormalTraceResponse(s float64) (linear, quadratic, alpha float64) {
	linear = LinearWeight() * s
	quadratic = QuadraticWeight() * s * s
	alpha = linear + quadratic
	return linear, quadratic, alpha
}

func FormalVariationalStationaryAlpha(s float64) float64 {
	// The formal action tested by the gate is
	//     A(alpha;s)=1/2[alpha - ((3/10)s+(7/72)s^2)]^2.
	// Its stationary point is exactly the inserted bridge rule.  This function
	// therefore verifies formal stationarity while the audit rejects native status.
	return AlphaB(s)
}

func BuildDefault() (Analysis, error) {
	s := SBoundary
	s2 := s * s
	lw := LinearWeight()
	qw := QuadraticWeight()
	linearAlpha, quadraticAlpha, alphaTrace := FormalTraceResponse(s)
	alpha := AlphaB(s)
	expectedAlpha := 0.0003878958469680527
	operatorNEff := NEffOperator(alpha)

	ledger := Ledger{
		S: s, S2: s2, LinearWeight: lw, QuadraticWeight: qw,
		LinearAlpha: linearAlpha, QuadraticAlpha: quadraticAlpha, AlphaB: alpha,
		ExpectedAlphaB: expectedAlpha, AlphaResidual: alpha - expectedAlpha,
		OperatorNEff: operatorNEff, OfficialNEff: OfficialNEff,
		OperatorCYukawa: CYukawaFromNEff(operatorNEff), OperatorCHiggs: CHiggsFromNEff(operatorNEff),
		OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs,
		AlphaSealedBridgeResponse: true,
	}

	trace := TraceExpansionAudit{
		Expression:  "R_trace(s)=Tr_V8+B2(P_3 sI)/dim(V_8 plus B_2)+Tr_H72(P_K7 s^2 I)/dim(H_72)",
		X1:          "X_1(s)=sI on V_8 plus B_2",
		X2:          "X_2(s)=s^2I on H_72",
		LinearTrace: float64(RankP3), QuadraticTrace: float64(DimK7),
		LinearContribution: linearAlpha, QuadraticContribution: quadraticAlpha,
		AlphaFromTrace: alphaTrace, ExpectedAlpha: alpha,
		ReconstructsAlpha:   math.Abs(alphaTrace-alpha) < 1e-18,
		X1NaturallyProduced: false, X2NaturallyProduced: false,
		TraceActionCertified: false, ClassifiedAsRestatement: true,
		Verdicts: []string{StatusGate829Inherited, StatusLiveWoundSelected, StatusTraceAnsatzBuilt, StatusTraceAnsatzRestatement},
		Supports: []string{SupportTraceRuleReconstructsAlpha, SupportTwoLaneSourceTyping},
		Failures: []string{FailureTraceExpansionRestatesRule, FailureX1NotNative, FailureX2NotNative, FailureNoBoundaryTraceAction, FailureNoBoundaryAlphaDomainTransport},
	}

	responseOrder := ResponseOrderAudit{
		LinearLane: "V_8 plus B_2 -> P_3", QuadraticLane: "H_72 -> K_7",
		LinearPower: 1, QuadraticPower: 2,
		LinearOrderCandidate:    "first-order vector-boundary triplet activation",
		QuadraticOrderCandidate: "second-order K_7/H_72 defect or self-intersection response",
		LinearOrderDerived:      false, QuadraticOrderDerived: false,
		FirstSecondOrderInterpretationAllowedAsCandidate: true,
		ResponseOrderTheoremCertified:                    false,
		Verdicts:                                         []string{StatusResponseOrderAudited, StatusResponseOrderOpen},
		Supports:                                         []string{SupportFirstSecondOrderStory, SupportTwoLaneSourceTyping},
		Failures:                                         []string{FailureResponseOrderNotDerived, FailureNoBoundaryTraceAction, FailureNoBoundaryAlphaDomainTransport},
	}

	stationaryAlpha := FormalVariationalStationaryAlpha(s)
	variational := VariationalAudit{
		ActionName:          "formal quadratic penalty action",
		ActionExpression:    "A(alpha;s)=1/2[alpha - ((Tr(P_3)/10)s+(Tr(P_K7)/72)s^2)]^2",
		StationaryCondition: "dA/dalpha=0 -> alpha=(3/10)s+(7/72)s^2",
		StationaryAlpha:     stationaryAlpha, ExpectedAlpha: alpha,
		StationarityWorksFormally:  math.Abs(stationaryAlpha-alpha) < 1e-18,
		WeightsAllTraceSourced:     true,
		PowersTypedByResponseOrder: false,
		ActionNative:               false,
		UsesInsertedAlphaRule:      true,
		IsFormalRepackaging:        true,
		CertifiesAlphaTheorem:      false,
		Verdicts:                   []string{StatusVariationalAudited, StatusVariationalRejected},
		Supports:                   []string{SupportFormalActionStationarity},
		Failures:                   []string{FailureVariationalRepackaging, FailureNoEulerLagrangeAlphaTheorem, FailureNoBoundaryTraceAction, FailureResponseOrderNotDerived},
	}

	nonCircular := NonCircularityAudit{
		AllowedInputs:         []string{"S_split", "rank(P_3)=3", "dim(V_8 plus B_2)=10", "dim(K_7)=7", "dim(H_72)=72", "Gate 826 B-L transfer", "Gate 829 aggregate operator only downstream"},
		ForbiddenInputs:       []string{"official N_eff", "operator_N_eff", "C_Yukawa", "C_Higgs", "observed Yukawa ratios", "Higgs mass", "PMNS/CKM", "sector assignment"},
		Direction:             "s -> alpha_B bridge response -> H_total -> operator_N_eff; never N_eff -> alpha_B",
		UsesNEffToDefineAlpha: false, UsesOfficialLedger: false, UsesObservedYukawas: false, UsesHiggsMass: false, ComputesAlphaBeforeReadout: true,
		Verdicts: []string{StatusNonCircularity},
		Supports: []string{SupportGate826829Downstream},
		Failures: []string{FailureAlphaNotNative},
	}

	impact := Impact{
		OperatorNEff: operatorNEff, OfficialNEff: OfficialNEff,
		CanPromoteAlpha: false, CanPromoteOperatorNEff: false,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		RecommendedAlphaStatus: "alpha_B = sealed bridge response / source-typed support-trace rule, not native theorem",
		NextGate:               "Gate 831 — R2++ / R3 Firewall and Sector Trace Ledger Obstruction Audit",
		Reason:                 "Gate 830 reconstructs the alpha polynomial by a formal trace expression and a formal stationary action, but both insert the response powers rather than deriving them; alpha_B cannot be promoted and N_eff seal reduction remains blocked.",
		Verdicts:               []string{StatusNoSealReduction, StatusNextGateDefined},
		Supports:               []string{SupportAlphaMarkedSealed, SupportR2PlusPlusStatus, SupportGate826829Downstream},
		Failures:               []string{FailureNoNEffSealReduction, FailureNoCYukawaUpdate, FailureNotR3SectorLedger, FailureNotR4Yukawa},
	}

	firewalls := Firewalls{
		Enforced: true, TraceExpansionRestatement: true, X1X2NotNative: true, ResponseOrderOpen: true,
		NoTraceAction: true, VariationalRepackage: true, AlphaSealed: true, NoBoundaryAlphaMap: true,
		NoNEffSealReduction: true, NoCYukawaUpdate: true, NotR3: true, NotR4: true, NoSectorAssignment: true,
		NoPMNSCKM: true, NoHiggs: true, Verdict: StatusFirewallGate830,
	}

	analysis := Analysis{
		Ledger: ledger, Trace: trace, ResponseOrder: responseOrder, Variational: variational,
		NonCircular: nonCircular, Impact: impact, Firewalls: firewalls,
		Truth: "Gate 830 audits the only broken arrow S_split -> alpha_B and finds that formal trace and variational rewrites reproduce the bridge rule but do not source it natively.",
		Final: "alpha_B remains a sealed bridge response: its support-trace weights are typed, but the powers s and s^2 and the action producing them are not derived by a native ASHA response law.",
	}
	if err := validate(analysis); err != nil {
		return Analysis{}, err
	}
	return analysis, nil
}

func validate(a Analysis) error {
	if math.Abs(a.Ledger.LinearWeight-0.3) > 1e-15 || math.Abs(a.Ledger.QuadraticWeight-7.0/72.0) > 1e-15 {
		return fmt.Errorf("support-trace weights failed: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.AlphaB-a.Ledger.ExpectedAlphaB) > 1e-18 || math.Abs(a.Trace.AlphaFromTrace-a.Ledger.AlphaB) > 1e-18 {
		return fmt.Errorf("alpha reconstruction failed: %s %s", FormatLedger(a.Ledger), FormatTrace(a.Trace))
	}
	if !a.Trace.ReconstructsAlpha || !a.Trace.ClassifiedAsRestatement || a.Trace.TraceActionCertified || a.Trace.X1NaturallyProduced || a.Trace.X2NaturallyProduced {
		return fmt.Errorf("trace-source firewall failed: %s", FormatTrace(a.Trace))
	}
	if a.ResponseOrder.ResponseOrderTheoremCertified || a.ResponseOrder.LinearOrderDerived || a.ResponseOrder.QuadraticOrderDerived || !a.ResponseOrder.FirstSecondOrderInterpretationAllowedAsCandidate {
		return fmt.Errorf("response-order firewall failed: %s", FormatResponseOrder(a.ResponseOrder))
	}
	if !a.Variational.StationarityWorksFormally || !a.Variational.UsesInsertedAlphaRule || !a.Variational.IsFormalRepackaging || a.Variational.ActionNative || a.Variational.CertifiesAlphaTheorem || a.Variational.PowersTypedByResponseOrder {
		return fmt.Errorf("variational firewall failed: %s", FormatVariational(a.Variational))
	}
	if a.NonCircular.UsesNEffToDefineAlpha || a.NonCircular.UsesOfficialLedger || a.NonCircular.UsesObservedYukawas || a.NonCircular.UsesHiggsMass || !a.NonCircular.ComputesAlphaBeforeReadout {
		return fmt.Errorf("noncircularity failed: %s", FormatNonCircularity(a.NonCircular))
	}
	if a.Impact.CanPromoteAlpha || a.Impact.CanPromoteOperatorNEff || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		return fmt.Errorf("impact freeze failed: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.TraceExpansionRestatement || !a.Firewalls.X1X2NotNative || !a.Firewalls.ResponseOrderOpen || !a.Firewalls.NoTraceAction || !a.Firewalls.VariationalRepackage || !a.Firewalls.AlphaSealed || !a.Firewalls.NoBoundaryAlphaMap || !a.Firewalls.NoNEffSealReduction || !a.Firewalls.NoCYukawaUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoSectorAssignment {
		return fmt.Errorf("firewall failed: %s", a.Firewalls.Verdict)
	}
	return nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.17g s2=%.17g w_linear=%.16g w_quadratic=%.16g alpha_linear=%.16g alpha_quadratic=%.16g alpha_B=%.17g expected=%.17g residual=%.3e operator_N_eff=%.16g official_N_eff=%.16g operator_C_Yukawa=%.16g operator_C_Higgs=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g alpha_sealed=%t", l.S, l.S2, l.LinearWeight, l.QuadraticWeight, l.LinearAlpha, l.QuadraticAlpha, l.AlphaB, l.ExpectedAlphaB, l.AlphaResidual, l.OperatorNEff, l.OfficialNEff, l.OperatorCYukawa, l.OperatorCHiggs, l.OfficialCYukawa, l.OfficialCHiggs, l.AlphaSealedBridgeResponse)
}

func FormatTrace(t TraceExpansionAudit) string {
	return fmt.Sprintf("%s X1=%s X2=%s Tr_linear=%.16g Tr_quadratic=%.16g alpha_linear=%.16g alpha_quadratic=%.16g alpha_trace=%.17g expected=%.17g reconstructs=%t X1_native=%t X2_native=%t trace_action=%t restatement=%t", t.Expression, t.X1, t.X2, t.LinearTrace, t.QuadraticTrace, t.LinearContribution, t.QuadraticContribution, t.AlphaFromTrace, t.ExpectedAlpha, t.ReconstructsAlpha, t.X1NaturallyProduced, t.X2NaturallyProduced, t.TraceActionCertified, t.ClassifiedAsRestatement)
}

func FormatResponseOrder(r ResponseOrderAudit) string {
	return fmt.Sprintf("linear=%s power=%d candidate=%s derived=%t quadratic=%s power=%d candidate=%s derived=%t first_second_candidate=%t theorem=%t", r.LinearLane, r.LinearPower, r.LinearOrderCandidate, r.LinearOrderDerived, r.QuadraticLane, r.QuadraticPower, r.QuadraticOrderCandidate, r.QuadraticOrderDerived, r.FirstSecondOrderInterpretationAllowedAsCandidate, r.ResponseOrderTheoremCertified)
}

func FormatVariational(v VariationalAudit) string {
	return fmt.Sprintf("%s %s stationarity=%s stationary_alpha=%.17g expected=%.17g formal_stationarity=%t trace_weights=%t powers_typed=%t action_native=%t inserted_rule=%t repackaging=%t certifies=%t", v.ActionName, v.ActionExpression, v.StationaryCondition, v.StationaryAlpha, v.ExpectedAlpha, v.StationarityWorksFormally, v.WeightsAllTraceSourced, v.PowersTypedByResponseOrder, v.ActionNative, v.UsesInsertedAlphaRule, v.IsFormalRepackaging, v.CertifiesAlphaTheorem)
}

func FormatNonCircularity(n NonCircularityAudit) string {
	return fmt.Sprintf("direction=%s allowed=%s forbidden=%s uses_N_eff=%t uses_official=%t uses_yukawas=%t uses_higgs=%t computes_alpha_first=%t", n.Direction, strings.Join(n.AllowedInputs, ", "), strings.Join(n.ForbiddenInputs, ", "), n.UsesNEffToDefineAlpha, n.UsesOfficialLedger, n.UsesObservedYukawas, n.UsesHiggsMass, n.ComputesAlphaBeforeReadout)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("operator_N_eff=%.16g official_N_eff=%.16g promote_alpha=%t promote_operator=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t alpha_status=%s next=%s reason=%s", i.OperatorNEff, i.OfficialNEff, i.CanPromoteAlpha, i.CanPromoteOperatorNEff, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.RecommendedAlphaStatus, i.NextGate, i.Reason)
}

func Statuses() []string {
	return []string{
		StatusGate829Inherited, StatusLiveWoundSelected, StatusTraceAnsatzBuilt, StatusTraceAnsatzRestatement,
		StatusResponseOrderAudited, StatusResponseOrderOpen, StatusVariationalAudited, StatusVariationalRejected,
		StatusNonCircularity, StatusNoSealReduction, StatusNextGateDefined, StatusPhysicalFirewalls,
		SupportTraceRuleReconstructsAlpha, SupportTwoLaneSourceTyping, SupportFirstSecondOrderStory,
		SupportFormalActionStationarity, SupportGate826829Downstream, SupportAlphaMarkedSealed, SupportR2PlusPlusStatus,
		FailureTraceExpansionRestatesRule, FailureX1NotNative, FailureX2NotNative, FailureResponseOrderNotDerived,
		FailureNoBoundaryTraceAction, FailureVariationalRepackaging, FailureNoEulerLagrangeAlphaTheorem,
		FailureNoBoundaryAlphaDomainTransport, FailureAlphaNotNative, FailureNoNEffSealReduction,
		FailureNoCYukawaUpdate, FailureNotR3SectorLedger, FailureNotR4Yukawa, FailureNoSectorAssignment,
		FailureNoPMNSCKM, FailureNoHiggsMass, StatusFirewallGate830,
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
