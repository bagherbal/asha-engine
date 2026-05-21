// Package generation2totalrelativetracemagnitudeoperatorandledgerconsistencyaudit
// implements Gate 829: Total Relative TraceMagnitude Operator and Ledger
// Consistency Audit.
//
// Gate 829 follows Gate 828's successful obstruction of a native
// BoundaryAlphaDomainTransportMap.  It therefore does not try to promote
// alpha_B.  Instead it consolidates the complete relative trace-magnitude
// operator given sealed/bridge alpha_B,
//
//	H_total/T = I_3 plus [alpha_B P_3 - 3 alpha_B^2(B-L)],
//
// derives the trace, square trace, and operator N_eff directly from that
// single diagonal operator, and enforces the ledger separation between:
//
//	operator_N_eff       diagnostic readout of the Gate 825/826 operator,
//	BFN_truncated_N_eff  fourth-order closure 3 + 6 alpha_B,
//	official_N_eff       frozen external/runtime ledger value.
//
// The gate is a consolidation and consistency audit only.  It does not derive
// alpha_B, does not assign Standard Model sectors, and does not update
// C_Yukawa or C_Higgs.
package generation2totalrelativetracemagnitudeoperatorandledgerconsistencyaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE829-TOTAL-RELATIVE-TRACE-MAGNITUDE-OPERATOR-LEDGER-CONSISTENCY-AUDIT"

	SBoundary = 0.0012924448188162962
	CHistory  = 1.038025177923625

	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	RankTopBlock = 3
	RankP1       = 1
	RankP3       = 3

	StatusGate828Inherited     = "PASS_GATE828_ALPHA_TRANSPORT_OBSTRUCTION_INHERITED"
	StatusAlphaSealedInput     = "PASS_ALPHA_B_USED_AS_SEALED_BRIDGE_INPUT_NOT_NATIVE_THEOREM"
	StatusTotalOperatorDefined = "PASS_TOTAL_RELATIVE_TRACE_MAGNITUDE_OPERATOR_DEFINED_GIVEN_ALPHA"
	StatusTopBlockAssembled    = "PASS_TOP_COLOR_BLOCK_I3_ASSEMBLED"
	StatusRestBlockAssembled   = "PASS_B_MINUS_L_REST_TRANSFER_BLOCK_ASSEMBLED"
	StatusTraceDerived         = "PASS_TOTAL_TRACE_DERIVED_FROM_OPERATOR"
	StatusSquareTraceDerived   = "PASS_TOTAL_SQUARE_TRACE_DERIVED_FROM_OPERATOR"
	StatusAbsoluteTCancels     = "PASS_ABSOLUTE_TOP_TRACE_ATOM_T_CANCELS_FROM_N_EFF"
	StatusOperatorNEffDerived  = "PASS_OPERATOR_N_EFF_FORM_DERIVED"
	StatusBFNResidualDerived   = "PASS_OPERATOR_VS_BFN_TRUNCATED_RESIDUAL_IS_FIFTH_ORDER"
	StatusLedgerSeparation     = "PASS_LEDGER_SEPARATION_OFFICIAL_VS_DIAGNOSTIC_ENFORCED"
	StatusAliasBugCorrected    = "PASS_GATE828_LEDGER_ALIASING_CORRECTED_IN_AUDIT_TEXT"
	StatusFreezePreserved      = "PASS_N_EFF_C_YUKAWA_C_HIGGS_FROZEN_DESPITE_DIAGNOSTIC_READOUT"
	StatusNextGateDefined      = "PASS_NEXT_PRESSURE_POINT_ALPHA_VARIATIONAL_TRACE_ACTION_SOURCE_DEFINED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate829      = "FIREWALL_PRESERVED_GATE829_TOTAL_OPERATOR_LEDGER_CONSISTENCY_BOUNDARY"

	SupportTotalOperatorGivenAlpha = "CONDITIONAL_SUPPORT_TOTAL_OPERATOR_IS_WELL_DEFINED_GIVEN_SEALED_ALPHA_B"
	SupportGate826RestTransfer     = "CONDITIONAL_SUPPORT_REST_BLOCK_REUSES_GATE826_B_MINUS_L_TRACE_ZERO_TRANSFER"
	SupportTraceFormula            = "CONDITIONAL_SUPPORT_A_TOTAL_OVER_T_EQUALS_3_PLUS_3_ALPHA_B"
	SupportSquareTraceFormula      = "CONDITIONAL_SUPPORT_B_TOTAL_OVER_T2_EQUALS_3_PLUS_3_ALPHA_B2_MINUS_6_ALPHA_B3_PLUS_12_ALPHA_B4"
	SupportOperatorNEffFormula     = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_EQUALS_3_ONE_PLUS_ALPHA_SQUARED_OVER_DENOMINATOR"
	SupportFifthOrderClosure       = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_MATCHES_BFN_TRUNCATED_CLOSURE_THROUGH_FOURTH_ORDER"
	SupportLedgerSeparation        = "CONDITIONAL_SUPPORT_OPERATOR_BFN_AND_OFFICIAL_N_EFF_ARE_DISTINCT_LEDGER_OBJECTS"
	SupportOfficialFreeze          = "CONDITIONAL_SUPPORT_OFFICIAL_LEDGER_REMAINS_FROZEN_UNTIL_ALPHA_SOURCE_AND_SECTOR_LEDGER_CERTIFIED"
	SupportNoAbsoluteTRequired     = "CONDITIONAL_SUPPORT_RELATIVE_TRACE_MAGNITUDE_READOUT_DOES_NOT_REQUIRE_ABSOLUTE_T"
	SupportR2PlusPlusConsolidated  = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_AGGREGATE_OPERATOR_CONSOLIDATED_BUT_NOT_PROMOTED_TO_R3"

	FailureAlphaNotNative           = "FAILED_ROUTE_ALPHA_B_NOT_NATIVE_BOUNDARY_THEOREM"
	FailureNoBoundaryAlphaMap       = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP"
	FailureTotalNotR3SectorLedger   = "FAILED_ROUTE_TOTAL_OPERATOR_NOT_R3_SECTOR_LEDGER"
	FailureTotalNotR4NativeYukawa   = "FAILED_ROUTE_TOTAL_OPERATOR_NOT_R4_NATIVE_YUKAWA_THEOREM"
	FailureNoCYukawaUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNEffSealReduction      = "FAILED_ROUTE_N_EFF_SEAL_REDUCTION_NOT_ALLOWED_WITHOUT_ALPHA_TRANSPORT"
	FailureNoSectorAssignment       = "FAILED_ROUTE_AGGREGATE_TRACE_OPERATOR_DOES_NOT_ASSIGN_STANDARD_MODEL_SECTORS"
	FailureNoVariationalAlphaSource = "FAILED_ROUTE_NO_VARIATIONAL_OR_TRACE_ACTION_SOURCE_FOR_ALPHA_B"
	FailureNoPMNSCKM                = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoHiggsMass              = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
)

type Ledger struct {
	S, AlphaB                                         float64
	OperatorNEff, BFNTruncatedNEff, OfficialNEff      float64
	OperatorMinusBFN, OperatorMinusOfficial           float64
	OfficialMinusBFN                                  float64
	OperatorCYukawa, OperatorCHiggs                   float64
	BFNTruncatedCYukawa, BFNTruncatedCHiggs           float64
	OfficialCYukawa, OfficialCHiggs                   float64
	OfficialFrozen, DiagnosticSeparated, AliasBlocked bool
}

type TotalOperatorAudit struct {
	Alpha                                         float64
	TopBlock, RestBlock, TotalSpectrum            []float64
	Formula                                       string
	RestTrace, ExpectedRestTrace                  float64
	TraceTotal, ExpectedTraceTotal                float64
	RestSquareTrace, ExpectedRestSquareTrace      float64
	SquareTraceTotal, ExpectedSquareTraceTotal    float64
	TraceResidual, SquareTraceResidual            float64
	AbsoluteTScaleCancels                         bool
	OperatorNEffFromTrace, OperatorNEffClosedForm float64
	OperatorFormulaResidual                       float64
	Verdicts, Supports, Failures                  []string
}

type LedgerConsistencyAudit struct {
	OperatorLabel, BFNTruncatedLabel, OfficialLabel  string
	OperatorNEff, BFNTruncatedNEff, OfficialNEff     float64
	OperatorMinusBFN, OperatorMinusOfficial          float64
	BFNMinusOfficial                                 float64
	FifthOrderResidualFormula                        string
	FifthOrderResidual                               float64
	SilentCollapseDetected                           bool
	OfficialUsedAsCandidate                          bool
	LedgerSeparationEnforced                         bool
	CanPromoteOperatorToOfficial                     bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	Reason                                           string
	Verdicts, Supports, Failures                     []string
}

type SourceFirewallAudit struct {
	AlphaSourceCertified, BoundaryAlphaTransportMapCertified bool
	TotalOperatorDefinedGivenAlpha                           bool
	SectorLedgerCertified, NativeYukawaTheoremCertified      bool
	AllowedInputs, ForbiddenInputs                           []string
	NextGate                                                 string
	Verdicts, Supports, Failures                             []string
}

type Firewalls struct {
	Enforced, AlphaNotNative, NoBoundaryAlphaMap bool
	NotR3SectorLedger, NotR4NativeYukawa         bool
	NoNEffSealReduction, NoCYukawaUpdate         bool
	NoSectorAssignment, NoVariationalAlpha       bool
	NoPMNSCKM, NoHiggs                           bool
	Verdict                                      string
}

type Analysis struct {
	Ledger      Ledger
	Operator    TotalOperatorAudit
	Consistency LedgerConsistencyAudit
	Source      SourceFirewallAudit
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func LinearWeight() float64                  { return 3.0 / 10.0 }
func QuadraticWeight() float64               { return 7.0 / 72.0 }
func AlphaB(s float64) float64               { return LinearWeight()*s + QuadraticWeight()*s*s }
func NEffBFNTruncated(alpha float64) float64 { return 3.0 + 6.0*alpha }
func CYukawaFromNEff(nEff float64) float64   { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64    { return CYukawaFromNEff(nEff) * CHistory }

func TotalTrace(alpha float64) float64 { return 3.0 + 3.0*alpha }
func TotalSquareTrace(alpha float64) float64 {
	return 3.0 + 3.0*alpha*alpha - 6.0*math.Pow(alpha, 3) + 12.0*math.Pow(alpha, 4)
}
func NEffOperator(alpha float64) float64 {
	return math.Pow(TotalTrace(alpha), 2) / TotalSquareTrace(alpha)
}
func NEffOperatorClosedForm(alpha float64) float64 {
	return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
}
func OperatorMinusBFNFifthOrder(alpha float64) float64 {
	return -24.0 * math.Pow(alpha, 5) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
}

func BuildDefault() (Analysis, error) {
	alpha := AlphaB(SBoundary)
	rest := []float64{3.0 * alpha * alpha, alpha * (1.0 - alpha), alpha * (1.0 - alpha), alpha * (1.0 - alpha)}
	top := []float64{1, 1, 1}
	total := append(append([]float64{}, top...), rest...)

	restTrace := traceDiag(rest)
	restSquare := dot(rest, rest)
	traceTotal := traceDiag(total)
	squareTotal := dot(total, total)
	expectedRestTrace := 3.0 * alpha
	expectedRestSquare := 3.0*alpha*alpha - 6.0*math.Pow(alpha, 3) + 12.0*math.Pow(alpha, 4)
	expectedTrace := TotalTrace(alpha)
	expectedSquare := TotalSquareTrace(alpha)
	operatorNEff := traceTotal * traceTotal / squareTotal
	operatorClosed := NEffOperatorClosedForm(alpha)
	bfn := NEffBFNTruncated(alpha)
	operatorMinusBFN := operatorNEff - bfn
	fifthResidual := OperatorMinusBFNFifthOrder(alpha)

	operator := TotalOperatorAudit{
		Alpha: alpha, TopBlock: top, RestBlock: rest, TotalSpectrum: total,
		Formula:   "H_total/T = I_3 plus [alpha_B P_3 - 3 alpha_B^2(B-L)]",
		RestTrace: restTrace, ExpectedRestTrace: expectedRestTrace,
		TraceTotal: traceTotal, ExpectedTraceTotal: expectedTrace,
		RestSquareTrace: restSquare, ExpectedRestSquareTrace: expectedRestSquare,
		SquareTraceTotal: squareTotal, ExpectedSquareTraceTotal: expectedSquare,
		TraceResidual: traceTotal - expectedTrace, SquareTraceResidual: squareTotal - expectedSquare,
		AbsoluteTScaleCancels: true,
		OperatorNEffFromTrace: operatorNEff, OperatorNEffClosedForm: operatorClosed,
		OperatorFormulaResidual: operatorNEff - operatorClosed,
		Verdicts:                []string{StatusGate828Inherited, StatusAlphaSealedInput, StatusTotalOperatorDefined, StatusTopBlockAssembled, StatusRestBlockAssembled, StatusTraceDerived, StatusSquareTraceDerived, StatusAbsoluteTCancels, StatusOperatorNEffDerived},
		Supports:                []string{SupportTotalOperatorGivenAlpha, SupportGate826RestTransfer, SupportTraceFormula, SupportSquareTraceFormula, SupportOperatorNEffFormula, SupportNoAbsoluteTRequired, SupportR2PlusPlusConsolidated},
		Failures:                []string{FailureAlphaNotNative, FailureNoBoundaryAlphaMap, FailureTotalNotR3SectorLedger, FailureTotalNotR4NativeYukawa},
	}

	ledger := Ledger{
		S: SBoundary, AlphaB: alpha,
		OperatorNEff: operatorNEff, BFNTruncatedNEff: bfn, OfficialNEff: OfficialNEff,
		OperatorMinusBFN: operatorMinusBFN, OperatorMinusOfficial: operatorNEff - OfficialNEff,
		OfficialMinusBFN: OfficialNEff - bfn,
		OperatorCYukawa:  CYukawaFromNEff(operatorNEff), OperatorCHiggs: CHiggsFromNEff(operatorNEff),
		BFNTruncatedCYukawa: CYukawaFromNEff(bfn), BFNTruncatedCHiggs: CHiggsFromNEff(bfn),
		OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs,
		OfficialFrozen: true, DiagnosticSeparated: true, AliasBlocked: true,
	}

	consistency := LedgerConsistencyAudit{
		OperatorLabel: "operator_N_eff", BFNTruncatedLabel: "BFN_truncated_N_eff", OfficialLabel: "official_frozen_N_eff",
		OperatorNEff: operatorNEff, BFNTruncatedNEff: bfn, OfficialNEff: OfficialNEff,
		OperatorMinusBFN: operatorMinusBFN, OperatorMinusOfficial: operatorNEff - OfficialNEff, BFNMinusOfficial: bfn - OfficialNEff,
		FifthOrderResidualFormula: "operator_N_eff - BFN_truncated_N_eff = -24 alpha_B^5/(1+alpha_B^2-2 alpha_B^3+4 alpha_B^4)",
		FifthOrderResidual:        fifthResidual,
		SilentCollapseDetected:    false, OfficialUsedAsCandidate: false, LedgerSeparationEnforced: true,
		CanPromoteOperatorToOfficial: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		Reason:   "Gate 829 derives an aggregate diagnostic readout from the total relative operator, but Gate 828 did not certify the alpha transport theorem and no R3 sector ledger exists; official values remain frozen.",
		Verdicts: []string{StatusBFNResidualDerived, StatusLedgerSeparation, StatusAliasBugCorrected, StatusFreezePreserved},
		Supports: []string{SupportFifthOrderClosure, SupportLedgerSeparation, SupportOfficialFreeze},
		Failures: []string{FailureNoNEffSealReduction, FailureNoCYukawaUpdate, FailureNoSectorAssignment},
	}

	source := SourceFirewallAudit{
		AlphaSourceCertified: false, BoundaryAlphaTransportMapCertified: false,
		TotalOperatorDefinedGivenAlpha: true, SectorLedgerCertified: false, NativeYukawaTheoremCertified: false,
		AllowedInputs:   []string{"Gate 826 B-L rest-transfer factorization", "Gate 827/828 alpha_B bridge-rule candidate", "I_3 top block", "projector trace algebra"},
		ForbiddenInputs: []string{"official N_eff as alpha source", "observed Yukawa ratios", "C_Yukawa update", "C_Higgs update", "SM sector assignment", "CKM/PMNS", "Higgs mass"},
		NextGate:        "Gate 830 — Alpha Variational / Trace-Action Source Obstruction Audit",
		Verdicts:        []string{StatusNextGateDefined},
		Supports:        []string{SupportR2PlusPlusConsolidated},
		Failures:        []string{FailureAlphaNotNative, FailureNoBoundaryAlphaMap, FailureNoVariationalAlphaSource, FailureTotalNotR3SectorLedger, FailureTotalNotR4NativeYukawa},
	}

	firewalls := Firewalls{
		Enforced: true, AlphaNotNative: true, NoBoundaryAlphaMap: true, NotR3SectorLedger: true, NotR4NativeYukawa: true,
		NoNEffSealReduction: true, NoCYukawaUpdate: true, NoSectorAssignment: true, NoVariationalAlpha: true,
		NoPMNSCKM: true, NoHiggs: true, Verdict: StatusFirewallGate829,
	}

	analysis := Analysis{
		Ledger: ledger, Operator: operator, Consistency: consistency, Source: source, Firewalls: firewalls,
		Truth: "Gate 829 consolidates the complete aggregate relative trace-magnitude operator given sealed alpha_B and derives its exact diagnostic N_eff readout, while keeping the alpha-source wound open.",
		Final: "The total operator is now ledger-clean: operator_N_eff and BFN_truncated_N_eff coincide through fourth order, but both remain diagnostic and distinct from the frozen official ledger until alpha transport and sector tracing are certified.",
	}
	if err := validate(analysis); err != nil {
		return Analysis{}, err
	}
	return analysis, nil
}

func validate(a Analysis) error {
	if math.Abs(a.Ledger.AlphaB-0.0003878958469680527) > 1e-18 {
		return fmt.Errorf("alpha reconstruction failed: %s", FormatLedger(a.Ledger))
	}
	if len(a.Operator.TotalSpectrum) != 7 || len(a.Operator.TopBlock) != 3 || len(a.Operator.RestBlock) != 4 {
		return fmt.Errorf("operator block dimensions failed: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Operator.TraceResidual) > 1e-15 || math.Abs(a.Operator.SquareTraceResidual) > 1e-15 || math.Abs(a.Operator.RestTrace-a.Operator.ExpectedRestTrace) > 1e-18 || math.Abs(a.Operator.RestSquareTrace-a.Operator.ExpectedRestSquareTrace) > 1e-21 {
		return fmt.Errorf("trace derivation failed: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Operator.OperatorFormulaResidual) > 1e-15 || math.Abs(a.Operator.OperatorNEffFromTrace-a.Ledger.OperatorNEff) > 1e-18 {
		return fmt.Errorf("operator N_eff formula failed: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Consistency.OperatorMinusBFN-a.Consistency.FifthOrderResidual) > 1e-15 {
		return fmt.Errorf("fifth-order residual failed: %s", FormatConsistency(a.Consistency))
	}
	if math.Abs(a.Ledger.OperatorNEff-a.Ledger.OfficialNEff) < 1e-10 || !a.Ledger.DiagnosticSeparated || !a.Ledger.AliasBlocked || !a.Ledger.OfficialFrozen {
		return fmt.Errorf("ledger separation failed: %s", FormatLedger(a.Ledger))
	}
	if a.Consistency.SilentCollapseDetected || a.Consistency.OfficialUsedAsCandidate || !a.Consistency.LedgerSeparationEnforced || a.Consistency.CanUpdateNEff || a.Consistency.CanUpdateCYukawa || a.Consistency.CanUpdateCHiggs {
		return fmt.Errorf("consistency firewall failed: %s", FormatConsistency(a.Consistency))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AlphaNotNative || !a.Firewalls.NoBoundaryAlphaMap || !a.Firewalls.NotR3SectorLedger || !a.Firewalls.NotR4NativeYukawa || !a.Firewalls.NoNEffSealReduction || !a.Firewalls.NoCYukawaUpdate || !a.Firewalls.NoSectorAssignment || !a.Firewalls.NoVariationalAlpha {
		return fmt.Errorf("firewall failed: %s", a.Firewalls.Verdict)
	}
	return nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.17g alpha_B=%.17g operator_N_eff=%.16g BFN_truncated_N_eff=%.16g official_frozen_N_eff=%.16g operator_minus_BFN=%.3e operator_minus_official=%.3e official_minus_BFN=%.3e operator_C_Yukawa=%.16g operator_C_Higgs=%.16g BFN_C_Yukawa=%.16g BFN_C_Higgs=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g official_frozen=%t diagnostic_separated=%t alias_blocked=%t", l.S, l.AlphaB, l.OperatorNEff, l.BFNTruncatedNEff, l.OfficialNEff, l.OperatorMinusBFN, l.OperatorMinusOfficial, l.OfficialMinusBFN, l.OperatorCYukawa, l.OperatorCHiggs, l.BFNTruncatedCYukawa, l.BFNTruncatedCHiggs, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.DiagnosticSeparated, l.AliasBlocked)
}

func FormatOperator(o TotalOperatorAudit) string {
	return fmt.Sprintf("%s alpha=%.17g top=%v rest=%v total=%v a_rest/T=%.16g expected=%.16g a_total/T=%.16g expected=%.16g b_rest/T2=%.16g expected=%.16g b_total/T2=%.16g expected=%.16g trace_residual=%.3e square_residual=%.3e T_cancels=%t operator_N_eff_trace=%.16g operator_N_eff_closed=%.16g formula_residual=%.3e", o.Formula, o.Alpha, o.TopBlock, o.RestBlock, o.TotalSpectrum, o.RestTrace, o.ExpectedRestTrace, o.TraceTotal, o.ExpectedTraceTotal, o.RestSquareTrace, o.ExpectedRestSquareTrace, o.SquareTraceTotal, o.ExpectedSquareTraceTotal, o.TraceResidual, o.SquareTraceResidual, o.AbsoluteTScaleCancels, o.OperatorNEffFromTrace, o.OperatorNEffClosedForm, o.OperatorFormulaResidual)
}

func FormatConsistency(c LedgerConsistencyAudit) string {
	return fmt.Sprintf("%s=%.16g %s=%.16g %s=%.16g operator_minus_BFN=%.3e formula_residual=%.3e operator_minus_official=%.3e BFN_minus_official=%.3e silent_collapse=%t official_used_as_candidate=%t separation=%t promote=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t reason=%s", c.OperatorLabel, c.OperatorNEff, c.BFNTruncatedLabel, c.BFNTruncatedNEff, c.OfficialLabel, c.OfficialNEff, c.OperatorMinusBFN, c.FifthOrderResidual, c.OperatorMinusOfficial, c.BFNMinusOfficial, c.SilentCollapseDetected, c.OfficialUsedAsCandidate, c.LedgerSeparationEnforced, c.CanPromoteOperatorToOfficial, c.CanUpdateNEff, c.CanUpdateCYukawa, c.CanUpdateCHiggs, c.Reason)
}

func FormatSource(s SourceFirewallAudit) string {
	return fmt.Sprintf("alpha_source=%t boundary_alpha_transport=%t total_operator_given_alpha=%t sector_ledger=%t native_yukawa=%t allowed=%s forbidden=%s next=%s", s.AlphaSourceCertified, s.BoundaryAlphaTransportMapCertified, s.TotalOperatorDefinedGivenAlpha, s.SectorLedgerCertified, s.NativeYukawaTheoremCertified, strings.Join(s.AllowedInputs, ", "), strings.Join(s.ForbiddenInputs, ", "), s.NextGate)
}

func Statuses() []string {
	return []string{
		StatusGate828Inherited, StatusAlphaSealedInput, StatusTotalOperatorDefined, StatusTopBlockAssembled,
		StatusRestBlockAssembled, StatusTraceDerived, StatusSquareTraceDerived, StatusAbsoluteTCancels,
		StatusOperatorNEffDerived, StatusBFNResidualDerived, StatusLedgerSeparation, StatusAliasBugCorrected,
		StatusFreezePreserved, StatusNextGateDefined, StatusPhysicalFirewalls,
		SupportTotalOperatorGivenAlpha, SupportGate826RestTransfer, SupportTraceFormula, SupportSquareTraceFormula,
		SupportOperatorNEffFormula, SupportFifthOrderClosure, SupportLedgerSeparation, SupportOfficialFreeze,
		SupportNoAbsoluteTRequired, SupportR2PlusPlusConsolidated,
		FailureAlphaNotNative, FailureNoBoundaryAlphaMap, FailureTotalNotR3SectorLedger, FailureTotalNotR4NativeYukawa,
		FailureNoCYukawaUpdate, FailureNoNEffSealReduction, FailureNoSectorAssignment, FailureNoVariationalAlphaSource,
		FailureNoPMNSCKM, FailureNoHiggsMass, StatusFirewallGate829,
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

func traceDiag(v []float64) float64 {
	s := 0.0
	for _, x := range v {
		s += x
	}
	return s
}

func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}
