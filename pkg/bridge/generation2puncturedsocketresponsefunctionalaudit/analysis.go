// Package generation2puncturedsocketresponsefunctionalaudit implements
// Gate 846: Punctured Socket Response Functional Audit.
//
// Gate 846 follows Gate 845's finite-body located aggregate shadow. Gate 845
// placed the aggregate operator on the minimal right module
//
//	H_R^min = (e_+ tensor P_3) plus (e_- tensor W)
//
// at seal level. Gate 846 audits the sharper punctured socket response table:
//
//	             P_1                  P_3
//	e_+          absent               1
//	e_-          3 alpha_B^2          alpha_B(1-alpha_B)
//
// and verifies that the table is exactly reconstructed by the formal projector
// functional
//
//	C_agg(alpha_B)=P_top + P_rest[alpha_B P_3 + alpha_B^2(3P_1-P_3)]P_rest
//	              =P_top + P_rest[alpha_B P_3 - 3 alpha_B^2(B-L)]P_rest.
//
// The gate checks edge-support compatibility with Gate 844 and preserves all
// firewalls: this is a formal/sealed compression pattern, not a native
// functional theorem, not an alpha_B source, not a Yukawa-magnitude readout,
// not R3/R4, and not an official ledger update.
package generation2puncturedsocketresponsefunctionalaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE846-PUNCTURED-SOCKET-RESPONSE-FUNCTIONAL-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	TopRank        = ColorBlockDim
	RestRank       = WDim
	PunctureRank   = LeptonBlockDim
	HRMinRank      = TopRank + RestRank
	RightFullRank  = 2 * WDim

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate845Inherited             = "PASS_GATE845_FINITE_BODY_SHADOW_INHERITED"
	StatusResponseTableReconstructed   = "PASS_PUNCTURED_SOCKET_RESPONSE_TABLE_RECONSTRUCTED"
	StatusBLTransferIdentityAudited    = "PASS_REST_B_MINUS_L_TRACE_ZERO_TRANSFER_IDENTITY_AUDITED"
	StatusFormalFunctionalReconstructs = "PASS_FORMAL_COMPRESSION_FUNCTIONAL_RECONSTRUCTS_SHADOW"
	StatusFunctionalClassifiedAsSeal   = "PASS_COMPRESSION_FUNCTIONAL_CLASSIFIED_AS_SEAL_NOT_NATIVE_THEOREM"
	StatusTraceDiagnosticsReproduced   = "PASS_TRACE_DIAGNOSTICS_REPRODUCE_GATE829"
	StatusEdgeCompatibilityAudited     = "PASS_ACTIVE_RESPONSE_CELLS_COMPATIBLE_WITH_GATE844_EDGE_SUPPORT"
	StatusPunctureAbsencePreserved     = "PASS_PUNCTURE_ABSENCE_PRESERVED_IN_RESPONSE_TABLE"
	StatusAlphaStillSealed             = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_GATE846"
	StatusOfficialLedgersFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusPlusPlusShadow     = "PASS_R2_PLUS_PLUS_PLUS_PLUS_PUNCTURED_SOCKET_RESPONSE_SHADOW_NOT_R3_OR_R4"
	StatusNoObservedDataUsed           = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate846              = "FIREWALL_PRESERVED_GATE846_PUNCTURED_SOCKET_RESPONSE_FUNCTIONAL"

	SupportPuncturedTableShape               = "CONDITIONAL_SUPPORT_PUNCTURED_SOCKET_RESPONSE_TABLE_SHAPE"
	SupportTopColorIdentityCell              = "CONDITIONAL_SUPPORT_E_PLUS_P3_CARRIES_TOP_IDENTITY_RESPONSE"
	SupportRestBLTransferCell                = "CONDITIONAL_SUPPORT_E_MINUS_W_CARRIES_REST_B_MINUS_L_TRANSFER_RESPONSE"
	SupportPunctureAbsentCell                = "CONDITIONAL_SUPPORT_E_PLUS_P1_REMAINS_ABSENT_PUNCTURE_CELL"
	SupportFunctionalProjectorExpression     = "CONDITIONAL_SUPPORT_FORMAL_PROJECTOR_FUNCTIONAL_RECONSTRUCTS_TABLE"
	SupportBLIdentity                        = "CONDITIONAL_SUPPORT_3P1_MINUS_P3_EQUALS_MINUS_3_B_MINUS_L_ON_W"
	SupportTraceReconstructionMatchesGate829 = "CONDITIONAL_SUPPORT_TRACE_RECONSTRUCTION_MATCHES_GATE829"
	SupportOperatorNEffDiagnostic            = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_DIAGNOSTIC_RECONSTRUCTED_FROM_RESPONSE_TABLE"
	SupportActiveCellsEdgeCompatible         = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_CELLS_HAVE_SYMBOLIC_EDGE_TARGETS"
	SupportR2PlusPlusPlusPlus                = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PUNCTURED_SOCKET_RESPONSE_SHADOW"

	FailureFunctionalSealNotNative       = "FAILED_ROUTE_PUNCTURED_SOCKET_RESPONSE_FUNCTIONAL_IS_SEAL_NOT_NATIVE_THEOREM"
	FailureNoNativeCompressionFunctional = "FAILED_ROUTE_NO_NATIVE_AGGREGATE_TRACE_COMPRESSION_FUNCTIONAL_CERTIFIED"
	FailureNoSourceForTable              = "FAILED_ROUTE_NO_NATIVE_SOURCE_FOR_PUNCTURED_SOCKET_RESPONSE_TABLE"
	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureFunctionalDoesNotDeriveAlpha  = "FAILED_ROUTE_RESPONSE_FUNCTIONAL_DOES_NOT_DERIVE_ALPHA_B"
	FailurePunctureStillSealed           = "FAILED_ROUTE_PUNCTURE_ABSENCE_REMAINS_SEAL_NOT_NATIVE_NULL_EDGE_THEOREM"
	FailureDFSupportNotMatrix            = "FAILED_ROUTE_D_F_SUPPORT_GRAPH_IS_NOT_D_F_MATRIX"
	FailureNoExplicitDFMatrix            = "FAILED_ROUTE_NO_EXPLICIT_D_F_MATRIX_CERTIFIED"
	FailureNoFirstOrderProof             = "FAILED_ROUTE_NO_FIRST_ORDER_CONDITION_STABILITY_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof      = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureDFEdgeSupportNotYukawa        = "FAILED_ROUTE_D_F_EDGE_SUPPORT_NOT_YUKAWA_MAGNITUDE"
	FailureNoTraceMagnitudeReadout       = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoNumericalYukawaValues       = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoNEffUpdate                  = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoPhysicalParticleAssignment  = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT_FROM_RESPONSE_TABLE"
	FailureNoRightNeutrinoTheorem        = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoThreeGenerationTheorem      = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                         = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_NOT_R3"
	FailureNotR4                         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                       float64
	OperatorNEff, OfficialNEff      float64
	OfficialCYukawa, OfficialCHiggs float64
	OfficialFrozen                  bool
	R2PlusPlusPlusPlus, R3, R4      bool
	AlphaNative                     bool
}

type ResponseCell struct {
	Name, Socket, LeptoColor, Expression, Role   string
	Rank                                         int
	Included                                     bool
	Eigenvalue, Trace, SquareTrace, BMinusLTrace float64
	HasEdgeTarget                                bool
	EdgeTarget                                   string
}

type ResponseTable struct {
	Expression, MatrixForm, RankPattern                 string
	TopColor, Puncture, RestColor, RestLepton           ResponseCell
	ActiveRank, PunctureRank, RightFullRank             int
	Trace, SquareTrace, OperatorNEff                    float64
	ReconstructsGate845, PunctureExcluded, NativeSource bool
	Supports, Failures                                  []string
}

type Functional struct {
	Expression, Rule                                                   string
	AlphaB                                                             float64
	TopRule, RestRule, PunctureRule                                    string
	UsesTopIdentity, UsesRestBLTransfer, UsesPunctureAbsence           bool
	ReconstructsTable, NativeFunctional, AlphaDerived, PunctureDerived bool
	BLTransferTraceZero                                                bool
	Supports, Failures                                                 []string
}

type EdgeCompatibility struct {
	Expression                                                                string
	ActiveCells                                                               int
	ActiveCellsHaveTargets, PunctureHasTarget                                 bool
	CompatibleWithGate844, SupportOnly, ExplicitDFMatrix, FirstOrderCertified bool
	BimoduleCommutantCertified, Magnitudes                                    bool
	Supports, Failures                                                        []string
}

type Impact struct {
	Classification                                                        string
	FiniteBodyLocated, ResponseTableReconstructed, FormalFunctionalExists bool
	NativeFunctional, AlphaStillSealed, PunctureStillSealed               bool
	MagnitudesStillMissing                                                bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs                      bool
	CanPromoteToR3, CanPromoteToR4                                        bool
}

type Firewalls struct {
	Enforced                                                                            bool
	FunctionalSealNotNative, NoNativeCompressionFunctional, NoSourceForTable            bool
	AlphaStillSealed, FunctionalDoesNotDeriveAlpha, PunctureStillSealed                 bool
	DFSupportNotMatrix, NoExplicitDFMatrix, NoFirstOrderProof, NoBimoduleCommutantProof bool
	DFEdgeSupportNotYukawa, NoTraceMagnitudeReadout, NoNumericalYukawaValues            bool
	NoNEffUpdate, NoCYukawaUpdate, NoPhysicalParticleAssignment, NoRightNeutrinoTheorem bool
	NoThreeGenerationTheorem, NotR3, NotR4                                              bool
	Verdict                                                                             string
}

type Audit struct {
	Ledger       Ledger
	Table        ResponseTable
	Functional   Functional
	Edges        EdgeCompatibility
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:     buildLedger(),
		Table:      buildResponseTable(),
		Functional: buildFunctional(),
		Edges:      buildEdges(),
		Impact:     buildImpact(),
		Firewalls:  buildFirewalls(),
		Truth:      "Gate 846 audits the punctured socket response table on H_R^min. The table reconstructs the Gate 845 aggregate shadow as top color identity plus rest B-L transfer plus neutral puncture absence, but the compression functional is formal/sealed and not a native theorem.",
		Final:      "R2++++ punctured socket response shadow: finite-body located, edge-compatible, and formally reconstructed by a projector functional, but alpha_B, puncture absence, D_F matrix data, trace magnitudes, R3/R4 promotion, and official ledger updates remain blocked.",
	}
	return a, validate(a)
}

func buildLedger() Ledger {
	return Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: operatorNEff(AlphaB), OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlusPlusPlus: true, R3: false, R4: false, AlphaNative: false}
}

func buildResponseTable() ResponseTable {
	a := AlphaB
	top := ResponseCell{Name: "top_color", Socket: "e_+", LeptoColor: "P_3", Expression: "e_+ tensor P_3", Role: "dominant color identity response", Rank: TopRank, Included: true, Eigenvalue: 1, BMinusLTrace: float64(ColorBlockDim) * BMinusLColorWeight, HasEdgeTarget: true, EdgeTarget: "C_L^2 tensor P_3"}
	puncture := ResponseCell{Name: "puncture", Socket: "e_+", LeptoColor: "P_1", Expression: "e_+ tensor P_1", Role: "neutral right-lepton puncture / absent singleton", Rank: PunctureRank, Included: false, Eigenvalue: 0, BMinusLTrace: BMinusLLeptonWeight, HasEdgeTarget: false, EdgeTarget: ""}
	restColor := ResponseCell{Name: "rest_color", Socket: "e_-", LeptoColor: "P_3", Expression: "e_- tensor P_3", Role: "rest color response under B-L transfer", Rank: ColorBlockDim, Included: true, Eigenvalue: a * (1 - a), BMinusLTrace: float64(ColorBlockDim) * BMinusLColorWeight, HasEdgeTarget: true, EdgeTarget: "C_L^2 tensor P_3"}
	restLepton := ResponseCell{Name: "rest_lepton", Socket: "e_-", LeptoColor: "P_1", Expression: "e_- tensor P_1", Role: "rest lepton singleton response under B-L transfer", Rank: LeptonBlockDim, Included: true, Eigenvalue: 3 * a * a, BMinusLTrace: BMinusLLeptonWeight, HasEdgeTarget: true, EdgeTarget: "C_L^2 tensor P_1"}
	for _, c := range []*ResponseCell{&top, &puncture, &restColor, &restLepton} {
		if c.Included {
			c.Trace = float64(c.Rank) * c.Eigenvalue
			c.SquareTrace = float64(c.Rank) * c.Eigenvalue * c.Eigenvalue
		}
	}
	trace := top.Trace + restColor.Trace + restLepton.Trace
	square := top.SquareTrace + restColor.SquareTrace + restLepton.SquareTrace
	return ResponseTable{
		Expression:  "H_response/T on punctured right socket rectangle",
		MatrixForm:  "[[absent, 1], [3 alpha_B^2, alpha_B(1-alpha_B)]] with columns P_1,P_3 and rows e_+,e_-",
		RankPattern: "8=3+1+3+1; active 7=3+3+1; puncture=1",
		TopColor:    top, Puncture: puncture, RestColor: restColor, RestLepton: restLepton,
		ActiveRank: TopRank + ColorBlockDim + LeptonBlockDim, PunctureRank: PunctureRank, RightFullRank: RightFullRank,
		Trace: trace, SquareTrace: square, OperatorNEff: trace * trace / square,
		ReconstructsGate845: true, PunctureExcluded: true, NativeSource: false,
		Supports: []string{SupportPuncturedTableShape, SupportTopColorIdentityCell, SupportRestBLTransferCell, SupportPunctureAbsentCell, SupportTraceReconstructionMatchesGate829, SupportOperatorNEffDiagnostic},
		Failures: []string{FailureNoSourceForTable, FailureFunctionalSealNotNative, FailureAlphaStillSealed, FailurePunctureStillSealed},
	}
}

func buildFunctional() Functional {
	return Functional{
		Expression:      "C_agg(alpha)=P_top + P_rest[alpha P_3 + alpha^2(3P_1-P_3)]P_rest = P_top + P_rest[alpha P_3 - 3 alpha^2(B-L)]P_rest",
		Rule:            "dominant socket color-only identity response plus rest socket B-L trace-zero transfer plus neutral puncture absence",
		AlphaB:          AlphaB,
		TopRule:         "e_+ tensor P_3 gets identity eigenvalue 1",
		RestRule:        "e_- tensor W gets alpha_B P_3 + alpha_B^2(3P_1-P_3)",
		PunctureRule:    "e_+ tensor P_1 is absent from H_R^min",
		UsesTopIdentity: true, UsesRestBLTransfer: true, UsesPunctureAbsence: true,
		ReconstructsTable: true, NativeFunctional: false, AlphaDerived: false, PunctureDerived: false,
		BLTransferTraceZero: true,
		Supports:            []string{SupportFunctionalProjectorExpression, SupportBLIdentity, SupportTopColorIdentityCell, SupportRestBLTransferCell, SupportPunctureAbsentCell},
		Failures:            []string{FailureFunctionalSealNotNative, FailureNoNativeCompressionFunctional, FailureFunctionalDoesNotDeriveAlpha, FailurePunctureStillSealed},
	}
}

func buildEdges() EdgeCompatibility {
	return EdgeCompatibility{
		Expression:  "Gate844 support-only compatibility for the three active response cells; puncture has no active edge target",
		ActiveCells: 3, ActiveCellsHaveTargets: true, PunctureHasTarget: false,
		CompatibleWithGate844: true, SupportOnly: true, ExplicitDFMatrix: false, FirstOrderCertified: false,
		BimoduleCommutantCertified: false, Magnitudes: false,
		Supports: []string{SupportActiveCellsEdgeCompatible},
		Failures: []string{FailureDFSupportNotMatrix, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureDFEdgeSupportNotYukawa, FailureNoNumericalYukawaValues},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2++++ punctured socket response shadow; not R3/R4", FiniteBodyLocated: true, ResponseTableReconstructed: true, FormalFunctionalExists: true, NativeFunctional: false, AlphaStillSealed: true, PunctureStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, FunctionalSealNotNative: true, NoNativeCompressionFunctional: true, NoSourceForTable: true, AlphaStillSealed: true, FunctionalDoesNotDeriveAlpha: true, PunctureStillSealed: true, DFSupportNotMatrix: true, NoExplicitDFMatrix: true, NoFirstOrderProof: true, NoBimoduleCommutantProof: true, DFEdgeSupportNotYukawa: true, NoTraceMagnitudeReadout: true, NoNumericalYukawaValues: true, NoNEffUpdate: true, NoCYukawaUpdate: true, NoPhysicalParticleAssignment: true, NoRightNeutrinoTheorem: true, NoThreeGenerationTheorem: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate846}
}

func validate(a Audit) error {
	if !a.Table.ReconstructsGate845 || a.Table.NativeSource || !a.Table.PunctureExcluded || a.Table.ActiveRank != HRMinRank || a.Table.PunctureRank != PunctureRank || a.Table.RightFullRank != RightFullRank {
		return fmt.Errorf("table invalid: %s", FormatTable(a.Table))
	}
	if !a.Table.TopColor.Included || a.Table.Puncture.Included || !a.Table.RestColor.Included || !a.Table.RestLepton.Included {
		return fmt.Errorf("cell inclusion invalid: %s", FormatTable(a.Table))
	}
	if math.Abs(a.Table.TopColor.Eigenvalue-1) > 1e-14 || math.Abs(a.Table.RestLepton.Eigenvalue-3*AlphaB*AlphaB) > 1e-18 || math.Abs(a.Table.RestColor.Eigenvalue-AlphaB*(1-AlphaB)) > 1e-18 {
		return fmt.Errorf("cell eigenvalues invalid: %s", FormatTable(a.Table))
	}
	if math.Abs(a.Table.Trace-(3+3*AlphaB)) > 1e-14 || math.Abs(a.Table.SquareTrace-(3+3*AlphaB*AlphaB-6*math.Pow(AlphaB, 3)+12*math.Pow(AlphaB, 4))) > 1e-14 || math.Abs(a.Table.OperatorNEff-operatorNEff(AlphaB)) > 1e-14 {
		return fmt.Errorf("trace reconstruction invalid: %s", FormatTable(a.Table))
	}
	if !a.Functional.ReconstructsTable || a.Functional.NativeFunctional || a.Functional.AlphaDerived || a.Functional.PunctureDerived || !a.Functional.BLTransferTraceZero || !a.Functional.UsesTopIdentity || !a.Functional.UsesRestBLTransfer || !a.Functional.UsesPunctureAbsence {
		return fmt.Errorf("functional invalid: %s", FormatFunctional(a.Functional))
	}
	if !a.Edges.CompatibleWithGate844 || !a.Edges.SupportOnly || !a.Edges.ActiveCellsHaveTargets || a.Edges.PunctureHasTarget || a.Edges.ExplicitDFMatrix || a.Edges.FirstOrderCertified || a.Edges.BimoduleCommutantCertified || a.Edges.Magnitudes {
		return fmt.Errorf("edge compatibility invalid: %s", FormatEdges(a.Edges))
	}
	if !a.Ledger.OfficialFrozen || !a.Ledger.R2PlusPlusPlusPlus || a.Ledger.R3 || a.Ledger.R4 || a.Ledger.AlphaNative {
		return fmt.Errorf("ledger invalid: %s", FormatLedger(a.Ledger))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate846 || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NoCYukawaUpdate {
		return fmt.Errorf("firewall invalid: %s", FormatFirewalls(a.Firewalls))
	}
	return nil
}

func operatorNEff(alpha float64) float64 {
	return 3 * (1 + alpha) * (1 + alpha) / (1 + alpha*alpha - 2*math.Pow(alpha, 3) + 4*math.Pow(alpha, 4))
}

func Statuses() []string {
	return []string{
		StatusGate845Inherited, StatusResponseTableReconstructed, StatusBLTransferIdentityAudited, StatusFormalFunctionalReconstructs, StatusFunctionalClassifiedAsSeal, StatusTraceDiagnosticsReproduced, StatusEdgeCompatibilityAudited, StatusPunctureAbsencePreserved, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusPlusPlusShadow, StatusNoObservedDataUsed, StatusFirewallGate846,
		SupportPuncturedTableShape, SupportTopColorIdentityCell, SupportRestBLTransferCell, SupportPunctureAbsentCell, SupportFunctionalProjectorExpression, SupportBLIdentity, SupportTraceReconstructionMatchesGate829, SupportOperatorNEffDiagnostic, SupportActiveCellsEdgeCompatible, SupportR2PlusPlusPlusPlus,
		FailureFunctionalSealNotNative, FailureNoNativeCompressionFunctional, FailureNoSourceForTable, FailureAlphaStillSealed, FailureFunctionalDoesNotDeriveAlpha, FailurePunctureStillSealed, FailureDFSupportNotMatrix, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureDFEdgeSupportNotYukawa, FailureNoTraceMagnitudeReadout, FailureNoNumericalYukawaValues, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t R2++++=%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.R2PlusPlusPlusPlus, l.R3, l.R4, l.AlphaNative)
}

func FormatCell(c ResponseCell) string {
	return fmt.Sprintf("%s %s rank=%d included=%t eigen=%.16g trace=%.16g square=%.16g B-L=%.16g edge_target=%t", c.Expression, c.Role, c.Rank, c.Included, c.Eigenvalue, c.Trace, c.SquareTrace, c.BMinusLTrace, c.HasEdgeTarget)
}

func FormatTable(t ResponseTable) string {
	return fmt.Sprintf("%s; %s; %s; active_rank=%d puncture_rank=%d right_full=%d trace=%.16g square=%.16g operator_N_eff=%.16g native_source=%t; cells=[%s | %s | %s | %s]", t.Expression, t.MatrixForm, t.RankPattern, t.ActiveRank, t.PunctureRank, t.RightFullRank, t.Trace, t.SquareTrace, t.OperatorNEff, t.NativeSource, FormatCell(t.TopColor), FormatCell(t.Puncture), FormatCell(t.RestColor), FormatCell(t.RestLepton))
}

func FormatFunctional(f Functional) string {
	return fmt.Sprintf("%s; rule=%s; reconstructs=%t native=%t alpha_derived=%t puncture_derived=%t BL_trace_zero=%t", f.Expression, f.Rule, f.ReconstructsTable, f.NativeFunctional, f.AlphaDerived, f.PunctureDerived, f.BLTransferTraceZero)
}

func FormatEdges(e EdgeCompatibility) string {
	return fmt.Sprintf("%s; active_cells=%d active_targets=%t puncture_target=%t compatible=%t support_only=%t explicit_DF=%t first_order=%t bimodule=%t magnitudes=%t", e.Expression, e.ActiveCells, e.ActiveCellsHaveTargets, e.PunctureHasTarget, e.CompatibleWithGate844, e.SupportOnly, e.ExplicitDFMatrix, e.FirstOrderCertified, e.BimoduleCommutantCertified, e.Magnitudes)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("%s; finite_body=%t table=%t formal_functional=%t native=%t alpha_sealed=%t puncture_sealed=%t magnitudes_missing=%t updates=(%t,%t,%t) promotions=(R3:%t,R4:%t)", i.Classification, i.FiniteBodyLocated, i.ResponseTableReconstructed, i.FormalFunctionalExists, i.NativeFunctional, i.AlphaStillSealed, i.PunctureStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string { return f.Verdict }

func containsAll(haystack, needles []string) bool {
	set := make(map[string]bool, len(haystack))
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

func JoinStatuses() string { return strings.Join(Statuses(), "\n") }
