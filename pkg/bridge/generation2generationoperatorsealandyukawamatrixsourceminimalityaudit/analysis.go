// Package generation2generationoperatorsealandyukawamatrixsourceminimalityaudit implements
// Gate 806: GenerationOperatorSeal and Yukawa Matrix Source Minimality Audit.
//
// Gate 806 inherits Gate 805's edge-triality no-go and identifies the true
// Yukawa source obstruction: the finite spectral triple tells ASHA where Yukawa
// edges may exist, but a separate operator on generation/family space is needed
// to determine what numerical matrices live on those edges.
package generation2generationoperatorsealandyukawamatrixsourceminimalityaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE806-GENERATION-OPERATOR-SEAL-YUKAWA-MATRIX-SOURCE-MINIMALITY-AUDIT"

	StatusGate805Inherited       = "PASS_GATE805_EDGE_TRIALITY_NO_GO_INHERITED"
	StatusFactorizationRecorded  = "PASS_FINITE_TRIPLE_EDGE_AND_GENERATION_OPERATOR_FACTORIZATION_RECORDED"
	StatusSealDefined            = "PASS_GENERATION_OPERATOR_SEAL_DEFINED"
	StatusMinimalityAudited      = "PASS_GENERATION_OPERATOR_MINIMALITY_AUDITED"
	StatusLayersSeparated        = "PASS_MAGNITUDE_AND_ORIENTATION_READOUT_LAYERS_SEPARATED"
	StatusFSTSourceAudited       = "PASS_FINITE_SPECTRAL_TRIPLE_SOURCE_AUDITED"
	StatusTD4SourceAudited       = "PASS_T_D4_SOURCE_AUDITED"
	StatusAggregateSourceAudited = "PASS_AGGREGATE_TRACE_LEDGER_SOURCE_AUDITED"
	StatusExternalSourceAudited  = "PASS_EXTERNAL_LEDGER_SOURCE_AUDITED"
	StatusK7ProjectiveAudited    = "PASS_K7_AND_PROJECTIVE_SOURCE_AUDITED"
	StatusNormalFormRecorded     = "PASS_YUKAWA_EDGE_TIMES_GENERATION_OPERATOR_NORMAL_FORM_RECORDED"
	StatusTraceObstruction       = "PASS_TRACE_READOUT_OBSTRUCTION_AUDITED"
	StatusMixingObstruction      = "PASS_MIXING_READOUT_OBSTRUCTION_AUDITED"
	StatusHierarchyObstruction   = "PASS_HIERARCHY_BREAKING_OBSTRUCTION_AUDITED"
	StatusCHiggsFirewall         = "PASS_C_HIGGS_FIREWALL_PRESERVED"
	StatusOutcomeRecorded        = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchDecision         = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusYukawaRequiresOperators      = "CONDITIONAL_SUPPORT_YUKAWA_TRACE_LEDGER_REQUIRES_SECTOR_OPERATORS_ON_GENERATION_SPACE"
	StatusSubobjectsNoncosmetic        = "CONDITIONAL_SUPPORT_ALL_GENERATION_OPERATOR_SUBOBJECTS_ARE_NONCOSMETIC"
	StatusNEffNeedsSpectra             = "CONDITIONAL_SUPPORT_N_EFF_NEEDS_HERMITIAN_SPECTRA_BUT_NOT_MIXING_FRAMES"
	StatusKappaOrientNeedsFrames       = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_NEEDS_MIXING_FRAMES_NOT_ONLY_SPECTRA"
	StatusFSTSuppliesEdgeDomain        = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_EDGE_DOMAIN_FOR_Y_F"
	StatusExternalCanPopulateSeal      = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_POPULATE_GENERATION_OPERATOR_DATA_AS_SEAL"
	StatusK7ProjectiveCandidates       = "CONDITIONAL_SUPPORT_K7_MINUS_OR_PROJECTIVE_THREE_ARE_FUTURE_GENERATION_CARRIER_CANDIDATES"
	StatusFSTWhereNotWhat              = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_WHERE_NOT_WHAT"
	StatusGenerationOperatorBottleneck = "CONDITIONAL_SUPPORT_GENERATION_OPERATOR_IS_NOW_THE_TRUE_YUKAWA_SOURCE_BOTTLENECK"
	StatusNextTraceMagnitude           = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_SPLIT_TRACE_MAGNITUDE_FROM_MIXING_ORIENTATION"

	StatusEdgeNoYF                   = "FAILED_ROUTE_EDGE_TEMPLATE_ALONE_DOES_NOT_SUPPLY_Y_F"
	StatusNoNativeGenerationOperator = "FAILED_ROUTE_NO_CURRENT_NATIVE_GENERATION_OPERATOR_SEAL"
	StatusCannotCompress             = "FAILED_ROUTE_GENERATION_OPERATOR_SEAL_CANNOT_BE_COMPRESSED_TO_EDGE_TEMPLATE_OR_T_D4"
	StatusSingularNoMixing           = "FAILED_ROUTE_SINGULAR_VALUES_ALONE_DO_NOT_DERIVE_PMNS_CKM"
	StatusFSTNoYF                    = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_NATIVE_DERIVE_Y_F_OPERATORS"
	StatusTD4NoGenOperator           = "FAILED_ROUTE_T_D4_EDGE_KERNEL_DOES_NOT_SUPPLY_GENERATION_OPERATOR"
	StatusTD4NoTraceAtoms            = "FAILED_ROUTE_T_D4_DOES_NOT_SUPPLY_Y_DAGGER_Y_TRACE_ATOMS"
	StatusAggregateNoOperator        = "FAILED_ROUTE_A_B_N_EFF_DO_NOT_IDENTIFY_GENERATION_OPERATOR"
	StatusExternalNotNative          = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_GENERATION_OPERATOR_THEOREM"
	StatusK7NotOperator              = "FAILED_ROUTE_K7_3_DIMENSIONAL_RESONANCE_NOT_GENERATION_OPERATOR_THEOREM"
	StatusProjectiveNotSource        = "FAILED_ROUTE_PROJECTIVE_1_PLUS_3_NOT_YUKAWA_MATRIX_SOURCE"
	StatusNoNativeWhatOperator       = "FAILED_ROUTE_NO_NATIVE_WHAT_OPERATOR_FOR_YUKAWA_VALUES"
	StatusCarrierAloneNoSpectra      = "FAILED_ROUTE_GENERATION_CARRIER_ALONE_DOES_NOT_SUPPLY_YUKAWA_SPECTRA"
	StatusGenerationsAloneNoNEff     = "FAILED_ROUTE_THREE_GENERATIONS_ALONE_DO_NOT_EXPLAIN_N_EFF_NEAR_THREE"
	StatusNoTraceAtomsWithoutH       = "FAILED_ROUTE_NO_TRACE_ATOM_EXTRACTION_WITHOUT_HERMITIAN_OPERATORS"
	StatusTraceLedgerNoKappaOrient   = "FAILED_ROUTE_TRACE_ATOM_LEDGER_ALONE_DOES_NOT_SOURCE_KAPPA_ORIENT"
	StatusNoPMNSCKMFrames            = "FAILED_ROUTE_NO_PMNS_CKM_WITHOUT_SECTOR_FRAME_MISALIGNMENT"
	StatusNoTopDominanceOperator     = "FAILED_ROUTE_NO_NATIVE_TOP_DOMINANCE_OPERATOR"
	StatusNoLightSuppressionOperator = "FAILED_ROUTE_NO_NATIVE_LIGHT_FAMILY_SUPPRESSION_OPERATOR"
	StatusNoNEffMinusThreeSource     = "FAILED_ROUTE_NO_NATIVE_N_EFF_MINUS_THREE_SOURCE"
	StatusNoCYukawaUpdate            = "FAILED_ROUTE_GENERATION_OPERATOR_AUDIT_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB               = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate806            = "FIREWALL_PRESERVED_GATE806_GENERATION_OPERATOR_SEAL_BOUNDARY"
)

type Inheritance struct {
	Gate805NoGo      bool
	FiniteTripleRole string
	TD4Role          string
	Missing          []string
	FactorizedNormal string
	Verdicts         []string
}

type GenerationOperatorSeal struct {
	Defined     bool
	Name        string
	Components  []string
	TraceChain  []string
	MixingChain []string
	Verdict     string
	Supports    []string
	Failures    []string
}

type MinimalityItem struct {
	Removed string
	Breaks  string
}

type MinimalityAudit struct {
	Audited  bool
	Items    []MinimalityItem
	Verdict  string
	Supports []string
	Failures []string
}

type ReadoutLayer struct {
	Name      string
	Needs     []string
	Produces  []string
	NotNeeded []string
	Blocked   []string
}

type ReadoutLayers struct {
	Separated bool
	Layers    []ReadoutLayer
	Verdict   string
	Supports  []string
	Failures  []string
}

type SourceAudit struct {
	Audited  bool
	Source   string
	Supplies []string
	Missing  []string
	Verdict  string
	Supports []string
	Failures []string
}

type NormalForm struct {
	Recorded bool
	Forms    []string
	Meaning  string
	Verdict  string
	Supports []string
	Failures []string
}

type Obstruction struct {
	Audited  bool
	Subject  string
	Reason   string
	Required []string
	Verdict  string
	Failures []string
}

type CHiggsFirewall struct {
	Preserved bool
	Formula   string
	Unchanged []string
	Verdict   string
	Failures  []string
}

type Outcome struct {
	Recorded bool
	Items    []string
	Verdict  string
	Supports []string
}

type BranchDecision struct {
	Recorded bool
	Next     string
	Branch   string
	Reason   string
	Verdict  string
	Supports []string
}

type Firewalls struct {
	Enforced      bool
	NoYukawa      bool
	NoEigenvalues bool
	NoPMNSCKM     bool
	NoFlavor      bool
	NoNEff        bool
	NoGJ          bool
	NoScalar      bool
	NoPoleMass    bool
	NoVEVGF       bool
	NoTriality    bool
	NoHistoryLoop bool
	Verdict       string
}

type Analysis struct {
	Inheritance  Inheritance
	Seal         GenerationOperatorSeal
	Minimality   MinimalityAudit
	Layers       ReadoutLayers
	FST          SourceAudit
	TD4          SourceAudit
	Aggregate    SourceAudit
	External     SourceAudit
	K7Projective SourceAudit
	NormalForm   NormalForm
	Trace        Obstruction
	Mixing       Obstruction
	Hierarchy    Obstruction
	CHiggs       CHiggsFirewall
	Outcome      Outcome
	Branch       BranchDecision
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func BuildDefault() (Analysis, error) {
	inheritance := Inheritance{
		Gate805NoGo:      true,
		FiniteTripleRole: "sector/gauge/chirality edge templates",
		TD4Role:          "possible airlocked trilinear kernel shape only",
		Missing:          []string{"actual operators on generation/family space", "Hermitian Yukawa matrices", "trace atoms", "sector-frame misalignment"},
		FactorizedNormal: "D_F^Y = sum_f Edge_f ⊗ Y_f + adjoint, with f in {u,d,e,nu}",
		Verdicts:         []string{StatusGate805Inherited, StatusFactorizationRecorded, StatusEdgeNoYF},
	}
	if !inheritance.Gate805NoGo || !strings.Contains(inheritance.FactorizedNormal, "Edge_f ⊗ Y_f") {
		return Analysis{}, fmt.Errorf("Gate 806 requires Gate 805 no-go inheritance and edge-times-generation-operator factorization")
	}

	seal := GenerationOperatorSeal{
		Defined: true,
		Name:    "GenerationOperatorSeal",
		Components: []string{
			"G_gen", "sector generation spaces", "sector Yukawa operators Y_u,Y_d,Y_e,Y_nu",
			"Hermitian trace operators H_f=Y_f†Y_f", "singular-value spectrum", "diagonalization frames",
			"PMNS/CKM misalignment readouts", "hierarchy/breaking operator", "scale and scheme convention",
			"color multiplicity rule", "neutrino convention", "noncircularity proof",
		},
		TraceChain:  []string{"Y_u,Y_d,Y_e,Y_nu", "H_f=Y_f†Y_f", "x_i=eigenvalues(H_f)", "a,b,N_eff"},
		MixingChain: []string{"Y_e,Y_nu -> U_PMNS -> theta13", "Y_u,Y_d -> V_CKM -> J_CKM"},
		Verdict:     StatusSealDefined,
		Supports:    []string{StatusYukawaRequiresOperators},
		Failures:    []string{StatusNoNativeGenerationOperator},
	}

	minimality := MinimalityAudit{
		Audited: true,
		Items: []MinimalityItem{
			{Removed: "remove G_gen", Breaks: "no family/generation domain"},
			{Removed: "remove sector spaces", Breaks: "no distinction between u,d,e,nu operators"},
			{Removed: "remove Y_f", Breaks: "no Yukawa matrices"},
			{Removed: "remove H_f=Y_f†Y_f", Breaks: "no positive trace atoms"},
			{Removed: "remove singular values", Breaks: "no a,b,N_eff"},
			{Removed: "remove diagonalization frames", Breaks: "no PMNS/CKM"},
			{Removed: "remove hierarchy/breaking operator", Breaks: "no top dominance, light-family suppression, or N_eff-3 source"},
			{Removed: "remove scale/scheme", Breaks: "no M_Z ledger or high-scale diagnostic"},
			{Removed: "remove color rule", Breaks: "no lawful factor 3 in a,b"},
			{Removed: "remove neutrino convention", Breaks: "no well-typed Y_nu sector"},
			{Removed: "remove noncircularity", Breaks: "no prediction status"},
		},
		Verdict:  StatusMinimalityAudited,
		Supports: []string{StatusSubobjectsNoncosmetic},
		Failures: []string{StatusCannotCompress},
	}

	layers := ReadoutLayers{
		Separated: true,
		Layers: []ReadoutLayer{
			{Name: "Magnitude / trace layer", Needs: []string{"H_f=Y_f†Y_f", "x_i=eigenvalues(H_f)"}, Produces: []string{"a", "b", "N_eff"}, NotNeeded: []string{"PMNS phases", "CKM phases"}},
			{Name: "Orientation / mixing layer", Needs: []string{"U_e", "U_nu", "U_u", "U_d", "sector-frame misalignment"}, Produces: []string{"U_PMNS", "V_CKM", "sin²(theta13)", "J_CKM"}, Blocked: []string{StatusSingularNoMixing}},
		},
		Verdict:  StatusLayersSeparated,
		Supports: []string{StatusNEffNeedsSpectra, StatusKappaOrientNeedsFrames},
		Failures: []string{StatusSingularNoMixing},
	}

	fst := SourceAudit{Audited: true, Source: "Finite spectral triple", Supplies: []string{"sector edge skeleton", "gauge-compatible chirality edges", "trace-form templates", "edge domain for Y_f"}, Missing: []string{"Y_f entries", "generation carrier", "eigenvalues", "mixing frames"}, Verdict: StatusFSTSourceAudited, Supports: []string{StatusFSTSuppliesEdgeDomain}, Failures: []string{StatusFSTNoYF}}
	td4 := SourceAudit{Audited: true, Source: "Complex D4 trilinear", Supplies: []string{"airlocked trilinear kernel shape"}, Missing: []string{"sector matrices", "Hermitian operators", "positive trace atoms", "generation hierarchy"}, Verdict: StatusTD4SourceAudited, Failures: []string{StatusTD4NoGenOperator, StatusTD4NoTraceAtoms}}
	agg := SourceAudit{Audited: true, Source: "Aggregate a,b,N_eff", Supplies: []string{"sealed aggregate trace values"}, Missing: []string{"operators", "sectors", "atoms", "eigenvectors"}, Verdict: StatusAggregateSourceAudited, Failures: []string{StatusAggregateNoOperator}}
	ext := SourceAudit{Audited: true, Source: "External Yukawa ledger", Supplies: []string{"sector values", "trace atoms", "N_eff audit", "possibly PMNS/CKM inputs"}, Missing: []string{"native derivation"}, Verdict: StatusExternalSourceAudited, Supports: []string{StatusExternalCanPopulateSeal}, Failures: []string{StatusExternalNotNative}}
	k7 := SourceAudit{Audited: true, Source: "K7 / Fock / projective structures", Supplies: []string{"K7=K7+⊕K7- with dim K7-=3", "Fock/projective split 4=1+3", "future carrier-search resonance"}, Missing: []string{"G_gen", "Y_f", "H_f", "PMNS/CKM"}, Verdict: StatusK7ProjectiveAudited, Supports: []string{StatusK7ProjectiveCandidates}, Failures: []string{StatusK7NotOperator, StatusProjectiveNotSource}}

	normal := NormalForm{
		Recorded: true,
		Forms: []string{
			"D_u = Edge_u ⊗ Y_u", "D_d = Edge_d ⊗ Y_d", "D_e = Edge_e ⊗ Y_e", "D_nu = Edge_nu ⊗ Y_nu",
		},
		Meaning:  "the finite edge template tells ASHA where a Yukawa coupling is allowed; the generation operator tells ASHA what the coupling is",
		Verdict:  StatusNormalFormRecorded,
		Supports: []string{StatusFSTWhereNotWhat},
		Failures: []string{StatusNoNativeWhatOperator},
	}

	trace := Obstruction{Audited: true, Subject: "Trace readout", Reason: "even G_gen ≅ C^3 would not determine eigenvalues, hierarchy, top dominance, N_eff-3, Koide, Froggatt-Nielsen powers, or Georgi-Jarlskog factors", Required: []string{"H_u,H_d,H_e,H_nu", "positive spectra", "trace atom extraction"}, Verdict: StatusTraceObstruction, Failures: []string{StatusCarrierAloneNoSpectra, StatusGenerationsAloneNoNEff, StatusNoTraceAtomsWithoutH}}
	mix := Obstruction{Audited: true, Subject: "Mixing readout", Reason: "kappa_orient requires relative sector frames and phases, not just trace atoms", Required: []string{"Y_e,Y_nu -> U_e,U_nu -> U_PMNS", "Y_u,Y_d -> U_u,U_d -> V_CKM"}, Verdict: StatusMixingObstruction, Failures: []string{StatusTraceLedgerNoKappaOrient, StatusNoPMNSCKMFrames}}
	hier := Obstruction{Audited: true, Subject: "Hierarchy and breaking", Reason: "top-color dominance requires one dominant colored singular value, small non-top participation, and controlled N_eff-3", Required: []string{"HierarchyBreakingOperator", "top-dominance mechanism", "light-family suppression", "rest-pressure mechanism"}, Verdict: StatusHierarchyObstruction, Failures: []string{StatusNoTopDominanceOperator, StatusNoLightSuppressionOperator, StatusNoNEffMinusThreeSource}}

	chiggs := CHiggsFirewall{Preserved: true, Formula: "C_Higgs = (3/N_eff) C_History", Unchanged: []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs", "lambda_H_bridge", "m_H_tree_proxy"}, Verdict: StatusCHiggsFirewall, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}}
	outcome := Outcome{Recorded: true, Items: []string{"finite spectral triple supplies lawful Yukawa edge locations", "D4 trilinear may still serve as an airlocked edge-kernel shape if embeddings are later found", "actual Yukawa matrices require independent generation-sector operators", "trace magnitude and mixing orientation are separate readout layers", "current ASHA does not supply the native GenerationOperatorSeal", "C_Higgs remains Level B"}, Verdict: StatusOutcomeRecorded, Supports: []string{StatusGenerationOperatorBottleneck}}
	branch := BranchDecision{Recorded: true, Next: "Gate 807 — TraceMagnitudeOperatorSeal and N_eff Source Audit", Branch: "magnitude / testability path", Reason: "N_eff has the highest numerical leverage in C_Higgs and requires only Hermitian trace spectra, not full PMNS/CKM orientation", Verdict: StatusBranchDecision, Supports: []string{StatusNextTraceMagnitude}}
	firewalls := Firewalls{Enforced: true, NoYukawa: true, NoEigenvalues: true, NoPMNSCKM: true, NoFlavor: true, NoNEff: true, NoGJ: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoTriality: true, NoHistoryLoop: true, Verdict: StatusFirewallGate806}

	return Analysis{Inheritance: inheritance, Seal: seal, Minimality: minimality, Layers: layers, FST: fst, TD4: td4, Aggregate: agg, External: ext, K7Projective: k7, NormalForm: normal, Trace: trace, Mixing: mix, Hierarchy: hier, CHiggs: chiggs, Outcome: outcome, Branch: branch, Firewalls: firewalls, Truth: "Gate 806 identifies the true Yukawa source obstruction: finite spectral triple edges supply where Yukawa couplings may live, but GenerationOperatorSeal is required to supply what numerical matrices live there.", Final: "N_eff needs Hermitian trace spectra H_f=Y_f†Y_f; kappa_orient needs sector-frame misalignment and phases. The next best gate is TraceMagnitudeOperatorSeal and N_eff Source Audit."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate805Inherited, StatusFactorizationRecorded, StatusSealDefined, StatusMinimalityAudited,
		StatusLayersSeparated, StatusFSTSourceAudited, StatusTD4SourceAudited, StatusAggregateSourceAudited,
		StatusExternalSourceAudited, StatusK7ProjectiveAudited, StatusNormalFormRecorded, StatusTraceObstruction,
		StatusMixingObstruction, StatusHierarchyObstruction, StatusCHiggsFirewall, StatusOutcomeRecorded,
		StatusBranchDecision, StatusPhysicalFirewalls, StatusYukawaRequiresOperators, StatusSubobjectsNoncosmetic,
		StatusNEffNeedsSpectra, StatusKappaOrientNeedsFrames, StatusFSTSuppliesEdgeDomain, StatusExternalCanPopulateSeal,
		StatusK7ProjectiveCandidates, StatusFSTWhereNotWhat, StatusGenerationOperatorBottleneck, StatusNextTraceMagnitude,
		StatusEdgeNoYF, StatusNoNativeGenerationOperator, StatusCannotCompress, StatusSingularNoMixing, StatusFSTNoYF,
		StatusTD4NoGenOperator, StatusTD4NoTraceAtoms, StatusAggregateNoOperator, StatusExternalNotNative,
		StatusK7NotOperator, StatusProjectiveNotSource, StatusNoNativeWhatOperator, StatusCarrierAloneNoSpectra,
		StatusGenerationsAloneNoNEff, StatusNoTraceAtomsWithoutH, StatusTraceLedgerNoKappaOrient, StatusNoPMNSCKMFrames,
		StatusNoTopDominanceOperator, StatusNoLightSuppressionOperator, StatusNoNEffMinusThreeSource, StatusNoCYukawaUpdate,
		StatusCHiggsLevelB, StatusFirewallGate806,
	}
}

func FormatSeal(s GenerationOperatorSeal) string {
	return fmt.Sprintf("%s components=[%s] trace=[%s] mixing=[%s] supports=[%s] failures=[%s]", s.Name, strings.Join(s.Components, "; "), strings.Join(s.TraceChain, " -> "), strings.Join(s.MixingChain, " | "), strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatMinimality(m MinimalityAudit) string {
	parts := make([]string, 0, len(m.Items))
	for _, it := range m.Items {
		parts = append(parts, it.Removed+" => "+it.Breaks)
	}
	return fmt.Sprintf("items=[%s] supports=[%s] failures=[%s]", strings.Join(parts, " | "), strings.Join(m.Supports, "; "), strings.Join(m.Failures, "; "))
}

func FormatLayers(l ReadoutLayers) string {
	parts := make([]string, 0, len(l.Layers))
	for _, layer := range l.Layers {
		parts = append(parts, fmt.Sprintf("%s needs=[%s] produces=[%s] notNeeded=[%s] blocked=[%s]", layer.Name, strings.Join(layer.Needs, "; "), strings.Join(layer.Produces, "; "), strings.Join(layer.NotNeeded, "; "), strings.Join(layer.Blocked, "; ")))
	}
	return fmt.Sprintf("layers=[%s] supports=[%s] failures=[%s]", strings.Join(parts, " | "), strings.Join(l.Supports, "; "), strings.Join(l.Failures, "; "))
}

func FormatSource(s SourceAudit) string {
	return fmt.Sprintf("%s supplies=[%s] missing=[%s] supports=[%s] failures=[%s]", s.Source, strings.Join(s.Supplies, "; "), strings.Join(s.Missing, "; "), strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatNormal(n NormalForm) string {
	return fmt.Sprintf("forms=[%s] meaning=%q supports=[%s] failures=[%s]", strings.Join(n.Forms, "; "), n.Meaning, strings.Join(n.Supports, "; "), strings.Join(n.Failures, "; "))
}

func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("%s reason=%q required=[%s] failures=[%s]", o.Subject, o.Reason, strings.Join(o.Required, "; "), strings.Join(o.Failures, "; "))
}

func FormatCHiggs(c CHiggsFirewall) string {
	return fmt.Sprintf("%s unchanged=[%s] failures=[%s]", c.Formula, strings.Join(c.Unchanged, "; "), strings.Join(c.Failures, "; "))
}

func FormatOutcome(o Outcome) string {
	return fmt.Sprintf("items=[%s] supports=[%s]", strings.Join(o.Items, "; "), strings.Join(o.Supports, "; "))
}

func containsAll(hay []string, needles []string) bool {
	joined := strings.Join(hay, "\n")
	for _, n := range needles {
		if !strings.Contains(joined, n) {
			return false
		}
	}
	return true
}
