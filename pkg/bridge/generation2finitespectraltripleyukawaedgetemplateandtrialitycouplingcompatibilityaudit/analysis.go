// Package generation2finitespectraltripleyukawaedgetemplateandtrialitycouplingcompatibilityaudit implements
// Gate 804: Finite Spectral Triple Yukawa Edge Template and Triality Coupling Compatibility Audit.
//
// Gate 804 inherits the Gate 803 no-go that the complex D4 trilinear T_D4 is
// not a Yukawa ledger.  It audits the narrower compatibility question: whether
// the already-certified finite spectral triple Yukawa edge skeleton can host the
// airlocked trilinear as a universal local edge-kernel shape, without promoting
// it to sector operators, trace atoms, N_eff, or C_Higgs.
package generation2finitespectraltripleyukawaedgetemplateandtrialitycouplingcompatibilityaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE804-FINITE-SPECTRAL-TRIPLE-YUKAWA-EDGE-TEMPLATE-TRIALITY-COUPLING-COMPATIBILITY-AUDIT"

	StatusGate803Inherited     = "PASS_GATE803_TRIALITY_YUKAWA_MINIMALITY_INHERITED"
	StatusFSTSelected          = "PASS_FINITE_SPECTRAL_TRIPLE_SELECTED_AS_NEXT_COMPATIBILITY_HOST"
	StatusEdgeTemplateRecorded = "PASS_FINITE_SPECTRAL_TRIPLE_YUKAWA_EDGE_TEMPLATE_RECORDED"
	StatusCompatibilityTarget  = "PASS_EDGE_TRIALITY_KERNEL_COMPATIBILITY_TARGET_DEFINED"
	StatusSlotMatchingAudited  = "PASS_TRIALITY_SLOT_MATCHING_AUDITED"
	StatusFourSectorFirewall   = "PASS_FOUR_SECTOR_THREE_SLOT_FIREWALL_AUDITED"
	StatusGaugeCompatAudited   = "PASS_GAUGE_REPRESENTATION_COMPATIBILITY_AUDITED"
	StatusHiggsCompatAudited   = "PASS_HIGGS_ONE_FORM_COMPATIBILITY_AUDITED"
	StatusHermitianReaudited   = "PASS_HERMITIAN_OPERATOR_OBSTRUCTION_REAUDITED"
	StatusGenerationRecorded   = "PASS_GENERATION_CARRIER_OBSTRUCTION_RECORDED"
	StatusTraceCompatAudited   = "PASS_TRACE_FORM_COMPATIBILITY_AUDITED"
	StatusTopColorFirewall     = "PASS_TOP_COLOR_DOMINANCE_FIREWALL_PRESERVED"
	StatusOutcomeRecorded      = "PASS_COMPATIBILITY_OUTCOME_RECORDED"
	StatusPackageUpdated       = "PASS_TRIALITY_YUKAWA_READOUT_PACKAGE_STATUS_UPDATED"
	StatusCHiggsFirewall       = "PASS_C_HIGGS_FIREWALL_PRESERVED"
	StatusBranchDecision       = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusFSTEdgeSkeleton            = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_STANDARD_MODEL_YUKAWA_EDGE_SKELETON"
	StatusTD4KernelShape             = "CONDITIONAL_SUPPORT_T_D4_MAY_BE_TESTED_AS_AIRLOCKED_EDGE_KERNEL_SHAPE"
	StatusTD4CorrectArity            = "CONDITIONAL_SUPPORT_T_D4_HAS_CORRECT_TRILINEAR_ARITY_FOR_YUKAWA_EDGE_KERNEL"
	StatusTD4UniversalKernelOnly     = "CONDITIONAL_SUPPORT_T_D4_COULD_ONLY_BE_UNIVERSAL_EDGE_KERNEL_NOT_SECTOR_LIST"
	StatusFSTPartialGaugeAssignment  = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_PARTIALLY_SUPPLIES_GAUGE_REPRESENTATION_ASSIGNMENT_SEAL"
	StatusHiggsOneFormCandidate      = "CONDITIONAL_SUPPORT_HIGGS_ONE_FORM_IS_THE_ONLY_PLAUSIBLE_T_D4_BOSONIC_SLOT_CANDIDATE"
	StatusTD4PreTraceKernel          = "CONDITIONAL_SUPPORT_T_D4_COULD_BE_A_PRE_TRACE_KERNEL_ONLY_AFTER_EDGE_EMBEDDING"
	StatusFSTTraceColorThree         = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_TRACE_FORM_CONTAINS_COLOR_FACTOR_THREE"
	StatusArityOnlyCompatibility     = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_AND_T_D4_ARE_STRUCTURALLY_COMPATIBLE_ONLY_AT_EDGE_KERNEL_ARITY_LEVEL"
	StatusFSTSuppliesPartialSkeleton = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_PART_OF_GAUGE_AND_SECTOR_EDGE_SKELETON"
	StatusNextEdgeTrialityEmbedding  = "CONDITIONAL_SUPPORT_NEXT_NATIVE_GATE_SHOULD_TEST_EDGE_TRIALITY_EMBEDDING"

	StatusTD4AloneNotLedger      = "FAILED_ROUTE_T_D4_ALONE_REMAINS_NOT_YUKAWA_LEDGER"
	StatusEdgeNoEigenvalues      = "FAILED_ROUTE_EDGE_TEMPLATE_DOES_NOT_DETERMINE_YUKAWA_EIGENVALUES"
	StatusEdgeNoMixing           = "FAILED_ROUTE_EDGE_TEMPLATE_DOES_NOT_SUPPLY_GENERATION_MIXING"
	StatusKernelNotReadout       = "FAILED_ROUTE_EDGE_KERNEL_COMPATIBILITY_NOT_YUKAWA_READOUT_THEOREM"
	StatusArityNoEmbedding       = "FAILED_ROUTE_ARITY_MATCH_DOES_NOT_PROVE_CARRIER_EMBEDDING"
	StatusNoHLREmbedding         = "FAILED_ROUTE_NO_CERTIFIED_HIGGS_LEFT_RIGHT_SLOT_EMBEDDING_IN_D4_CARRIERS"
	StatusThreeSlotsNotFour      = "FAILED_ROUTE_THREE_TRIALITY_SLOTS_DO_NOT_MATCH_FOUR_YUKAWA_SECTORS"
	StatusTD4NoSectorReplacement = "FAILED_ROUTE_T_D4_DOES_NOT_REPLACE_SECTOR_ASSIGNMENT"
	StatusD4NoGaugeAssignment    = "FAILED_ROUTE_D4_TRIALITY_DOES_NOT_SUPPLY_GAUGE_REPRESENTATION_ASSIGNMENT"
	StatusNoEdgeToD4Theorem      = "FAILED_ROUTE_NO_EDGE_TO_D4_SLOT_EMBEDDING_THEOREM"
	StatusNoHiggsSlotEmbedding   = "FAILED_ROUTE_NO_NATIVE_HIGGS_SLOT_EMBEDDING_IN_D4_CARRIER"
	StatusK7PlusNotD4Vector      = "FAILED_ROUTE_K7_PLUS_HIGGS_SOCKET_NOT_D4_VECTOR_SLOT_THEOREM"
	StatusKernelNoYFMatrix       = "FAILED_ROUTE_EDGE_KERNEL_DOES_NOT_SUPPLY_Y_F_MATRIX"
	StatusNoGenerationOperator   = "FAILED_ROUTE_NO_GENERATION_OPERATOR_MULTIPLYING_T_D4_KERNEL"
	StatusNoYdaggerYTrace        = "FAILED_ROUTE_NO_Y_DAGGER_Y_TRACE_FROM_TRIALITY_EDGE_KERNEL"
	StatusFSTNoNativeGenerations = "FAILED_ROUTE_FINITE_TRIPLE_EDGE_TEMPLATE_DOES_NOT_NATIVE_DERIVE_THREE_GENERATIONS"
	StatusTD4NoGenerations       = "FAILED_ROUTE_TRIALITY_EDGE_KERNEL_DOES_NOT_NATIVE_DERIVE_THREE_GENERATIONS"
	StatusNoPMNSCKM              = "FAILED_ROUTE_NO_PMNS_CKM_WITHOUT_GENERATION_CARRIER_AND_SECTOR_OPERATORS"
	StatusTD4NoTraceInputs       = "FAILED_ROUTE_T_D4_DOES_NOT_SUPPLY_TRACE_FORM_INPUTS_Y_F"
	StatusTD4NoABNEffUpdate      = "FAILED_ROUTE_T_D4_DOES_NOT_UPDATE_A_B_N_EFF"
	StatusEdgeNoTopDominance     = "FAILED_ROUTE_EDGE_TEMPLATE_DOES_NOT_DERIVE_TOP_DOMINANCE"
	StatusKernelNoNEffMinusThree = "FAILED_ROUTE_T_D4_EDGE_KERNEL_DOES_NOT_EXPLAIN_N_EFF_MINUS_THREE"
	StatusNoEdgeEmbeddingReadout = "FAILED_ROUTE_NO_EDGE_EMBEDDING_OR_TRACE_READOUT_THEOREM"
	StatusReadoutStillMissing    = "FAILED_ROUTE_REMAINING_READOUT_PACKAGE_STILL_MISSING"
	StatusNoCYukawaUpdate        = "FAILED_ROUTE_EDGE_COMPATIBILITY_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate804        = "FIREWALL_PRESERVED_GATE804_FINITE_TRIPLE_TRIALITY_EDGE_COMPATIBILITY_BOUNDARY"
)

type Inheritance struct {
	TD4Status            string
	FiniteTripleNextHost bool
	TD4IsLedger          bool
	HasYukawaReadout     bool
	Verdicts             []string
}

type EdgeTemplate struct {
	Recorded bool
	Edges    []string
	Knows    []string
	Missing  []string
	Verdict  string
	Supports []string
	Failures []string
}

type CompatibilityTarget struct {
	Defined      bool
	SealName     string
	Components   []string
	LawfulGoal   string
	RejectedGoal string
	Verdict      string
	Supports     []string
	Failures     []string
}

type SlotMatching struct {
	Audited     bool
	D4Slots     []string
	YukawaSlots []string
	Candidate   string
	Required    []string
	Verdict     string
	Supports    []string
	Failures    []string
}

type FourSectorFirewall struct {
	Audited      bool
	Sectors      []string
	D4Slots      []string
	LawfulForm   string
	BlockedClaim string
	Verdict      string
	Supports     []string
	Failures     []string
}

type GaugeCompatibility struct {
	Audited          bool
	FiniteTripleRole string
	TD4Role          string
	Missing          []string
	Verdict          string
	Supports         []string
	Failures         []string
}

type HiggsCompatibility struct {
	Audited      bool
	RequiredSeal string
	Components   []string
	Candidate    string
	Verdict      string
	Supports     []string
	Failures     []string
}

type Obstruction struct {
	Audited  bool
	Subject  string
	Reason   string
	Verdict  string
	Supports []string
	Failures []string
}

type TraceCompatibility struct {
	Audited    bool
	TraceForms []string
	Question   string
	Answer     string
	Verdict    string
	Supports   []string
	Failures   []string
}

type Outcome struct {
	Recorded bool
	Items    []string
	Verdict  string
	Supports []string
	Failures []string
}

type PackageStatus struct {
	Updated           bool
	PartiallySupplied []string
	NotSupplied       []string
	Verdict           string
	Supports          []string
	Failures          []string
}

type CHiggsFirewall struct {
	Preserved bool
	Formula   string
	Unchanged []string
	Verdict   string
	Failures  []string
}

type BranchDecision struct {
	Recorded bool
	Next     string
	Seal     string
	Purpose  string
	Verdict  string
	Supports []string
}

type Firewalls struct {
	Enforced         bool
	NoYukawa         bool
	NoEigenvalues    bool
	NoPMNSCKM        bool
	NoFlavor         bool
	NoNEff           bool
	NoGJ             bool
	NoScalar         bool
	NoPoleMass       bool
	NoVEVGF          bool
	NoNativeTriality bool
	NoHistoryLoop    bool
	Verdict          string
}

type Analysis struct {
	Inheritance  Inheritance
	EdgeTemplate EdgeTemplate
	Target       CompatibilityTarget
	SlotMatching SlotMatching
	FourSector   FourSectorFirewall
	Gauge        GaugeCompatibility
	Higgs        HiggsCompatibility
	Hermitian    Obstruction
	Generation   Obstruction
	Trace        TraceCompatibility
	TopColor     Obstruction
	Outcome      Outcome
	Package      PackageStatus
	CHiggs       CHiggsFirewall
	Branch       BranchDecision
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func BuildDefault() (Analysis, error) {
	inheritance := Inheritance{
		TD4Status:            "lawful complex airlocked pre-Yukawa shape; not sector-labeled, Hermitian, generation-resolved, scale-local, or N_eff readout",
		FiniteTripleNextHost: true,
		TD4IsLedger:          false,
		HasYukawaReadout:     false,
		Verdicts:             []string{StatusGate803Inherited, StatusFSTSelected, StatusTD4AloneNotLedger},
	}
	if inheritance.TD4IsLedger || inheritance.HasYukawaReadout || !inheritance.FiniteTripleNextHost {
		return Analysis{}, fmt.Errorf("Gate 804 requires T_D4 to remain non-ledger while finite spectral triple is selected as host")
	}

	edge := EdgeTemplate{
		Recorded: true,
		Edges: []string{
			"Y_u edge: Q_L -> u_R",
			"Y_d edge: Q_L -> d_R",
			"Y_e edge: L_L -> e_R",
			"Y_nu edge: L_L -> nu_R or chosen neutrino convention",
		},
		Knows:    []string{"chirality", "gauge representation compatibility", "Higgs one-form edge location", "sector labels"},
		Missing:  []string{"numerical Yukawa entries", "generation mixing", "Yukawa eigenvalues"},
		Verdict:  StatusEdgeTemplateRecorded,
		Supports: []string{StatusFSTEdgeSkeleton},
		Failures: []string{StatusEdgeNoEigenvalues, StatusEdgeNoMixing},
	}

	target := CompatibilityTarget{
		Defined:      true,
		SealName:     "EdgeTrialityKernelCompatibilitySeal",
		Components:   []string{"edge label f in {u,d,e,nu}", "finite spectral triple edge E_f", "airlocked triality trilinear T_D4", "embedding of edge carrier into D4 slots", "Higgs-slot compatibility", "chirality compatibility", "gauge-label preservation"},
		LawfulGoal:   "T_D4 may act as a universal pre-Yukawa local coupling kernel after the finite spectral triple has supplied the sector edge.",
		RejectedGoal: "T_D4 -> all Yukawa data",
		Verdict:      StatusCompatibilityTarget,
		Supports:     []string{StatusTD4KernelShape},
		Failures:     []string{StatusKernelNotReadout},
	}

	slot := SlotMatching{
		Audited:     true,
		D4Slots:     []string{"V_C", "S_plus_C", "S_minus_C"},
		YukawaSlots: []string{"Higgs", "left fermion", "right fermion"},
		Candidate:   "T_D4(H, psi_L, psi_R)",
		Required:    []string{"Higgs carrier embeds into chosen D4 slot", "left/right finite triple fermion carriers embed into spinor slots", "sector labels preserved", "real form descends correctly", "hypercharge and gauge charges match"},
		Verdict:     StatusSlotMatchingAudited,
		Supports:    []string{StatusTD4CorrectArity},
		Failures:    []string{StatusArityNoEmbedding, StatusNoHLREmbedding},
	}

	four := FourSectorFirewall{
		Audited:      true,
		Sectors:      []string{"u", "d", "e", "nu"},
		D4Slots:      []string{"V_C", "S_plus_C", "S_minus_C"},
		LawfulForm:   "finite triple supplies four sector edges; triality supplies a possible common trilinear kernel shape per edge",
		BlockedClaim: "three triality slots = four Standard Model Yukawa sectors",
		Verdict:      StatusFourSectorFirewall,
		Supports:     []string{StatusTD4UniversalKernelOnly},
		Failures:     []string{StatusThreeSlotsNotFour, StatusTD4NoSectorReplacement},
	}

	gauge := GaugeCompatibility{
		Audited:          true,
		FiniteTripleRole: "supplies gauge-compatible Standard Model Yukawa edge templates",
		TD4Role:          "may supply only an airlocked coupling tensor shape",
		Missing:          []string{"embedding of each finite-triple edge carrier into D4 triality slots"},
		Verdict:          StatusGaugeCompatAudited,
		Supports:         []string{StatusFSTPartialGaugeAssignment},
		Failures:         []string{StatusD4NoGaugeAssignment, StatusNoEdgeToD4Theorem},
	}

	higgs := HiggsCompatibility{
		Audited:      true,
		RequiredSeal: "HiggsSlotEmbeddingSeal",
		Components:   []string{"finite Higgs one-form carrier", "D4 slot assignment", "compatibility with K7+ Higgs socket", "complex structure / real-form airlock", "normalization convention"},
		Candidate:    "finite Higgs one-form as the only plausible T_D4 bosonic slot candidate",
		Verdict:      StatusHiggsCompatAudited,
		Supports:     []string{StatusHiggsOneFormCandidate},
		Failures:     []string{StatusNoHiggsSlotEmbedding, StatusK7PlusNotD4Vector},
	}

	hermitian := Obstruction{Audited: true, Subject: "Hermitian operator and singular-value trace extraction", Reason: "An edge-kernel shape is an amplitude template; the trace ledger requires matrices Y_f and positive Y_f†Y_f traces on generation space.", Verdict: StatusHermitianReaudited, Failures: []string{StatusKernelNoYFMatrix, StatusNoGenerationOperator, StatusNoYdaggerYTrace}}
	generation := Obstruction{Audited: true, Subject: "Generation carrier", Reason: "Finite triple edges give sector/gauge/chirality templates; T_D4 gives representation-type coupling shape; neither natively derives a three-family carrier or PMNS/CKM frames.", Verdict: StatusGenerationRecorded, Failures: []string{StatusFSTNoNativeGenerations, StatusTD4NoGenerations, StatusNoPMNSCKM}}

	trace := TraceCompatibility{
		Audited: true,
		TraceForms: []string{
			"a = Tr(Y_e†Y_e + Y_nu†Y_nu + 3Y_u†Y_u + 3Y_d†Y_d)",
			"b = Tr((Y_e†Y_e)^2 + (Y_nu†Y_nu)^2 + 3(Y_u†Y_u)^2 + 3(Y_d†Y_d)^2)",
		},
		Question: "Does T_D4 supply the Y_f whose traces enter a,b?",
		Answer:   "No. At most, after an embedding theorem, T_D4 could provide a common pre-trace kernel shape; numerical Y_f matrices and singular values remain unsourced.",
		Verdict:  StatusTraceCompatAudited,
		Supports: []string{StatusTD4PreTraceKernel},
		Failures: []string{StatusTD4NoTraceInputs, StatusTD4NoABNEffUpdate},
	}

	top := Obstruction{Audited: true, Subject: "Top-color dominance", Reason: "The finite spectral-action trace forms contain color factor three, but edge templates and triality kernels do not derive why top dominates or why N_eff-3 is the observed small positive pressure.", Verdict: StatusTopColorFirewall, Supports: []string{StatusFSTTraceColorThree}, Failures: []string{StatusEdgeNoTopDominance, StatusKernelNoNEffMinusThree}}

	outcome := Outcome{
		Recorded: true,
		Items: []string{
			"finite spectral triple supplies the Standard Model Yukawa edge skeleton",
			"T_D4 has the right trilinear arity to be tested as an airlocked universal edge-kernel shape",
			"no embedding theorem maps finite-triple edge carriers into D4 slots",
			"no Hermitian generation operator or singular-value trace readout follows",
			"no C_Higgs update follows",
		},
		Verdict:  StatusOutcomeRecorded,
		Supports: []string{StatusArityOnlyCompatibility},
		Failures: []string{StatusNoEdgeEmbeddingReadout},
	}

	pkg := PackageStatus{
		Updated: true,
		PartiallySupplied: []string{
			"GaugeRepresentationAssignmentSeal: partially supplied by finite spectral triple edge templates",
			"SectorAssignmentSeal: partially supplied by finite spectral triple sector labels",
		},
		NotSupplied: []string{"HiggsSlotEmbeddingSeal", "GenerationCarrierSeal", "HermitianOperatorSeal", "SymmetryBreakingHierarchySeal", "TraceAtomExtractionSeal", "RealDescentSeal"},
		Verdict:     StatusPackageUpdated,
		Supports:    []string{StatusFSTSuppliesPartialSkeleton},
		Failures:    []string{StatusReadoutStillMissing},
	}

	chiggs := CHiggsFirewall{Preserved: true, Formula: "C_Higgs=(3/N_eff)C_History", Unchanged: []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs", "lambda_H_bridge", "m_H_tree_proxy"}, Verdict: StatusCHiggsFirewall, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}}
	branch := BranchDecision{Recorded: true, Next: "Gate 805 — EdgeTrialityEmbeddingSeal and Higgs/Fermion Slot Assignment No-Go Audit", Seal: "EdgeTrialityEmbeddingSeal", Purpose: "test whether any current ASHA carrier can embed finite spectral triple Higgs/fermion edge data into D4 trilinear slots without violating real-form, gauge, chirality, or generation firewalls", Verdict: StatusBranchDecision, Supports: []string{StatusNextEdgeTrialityEmbedding}}
	firewalls := Firewalls{Enforced: true, NoYukawa: true, NoEigenvalues: true, NoPMNSCKM: true, NoFlavor: true, NoNEff: true, NoGJ: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoNativeTriality: true, NoHistoryLoop: true, Verdict: StatusFirewallGate804}

	return Analysis{Inheritance: inheritance, EdgeTemplate: edge, Target: target, SlotMatching: slot, FourSector: four, Gauge: gauge, Higgs: higgs, Hermitian: hermitian, Generation: generation, Trace: trace, TopColor: top, Outcome: outcome, Package: pkg, CHiggs: chiggs, Branch: branch, Firewalls: firewalls, Truth: "Gate 804 finds the first real compatibility point: finite spectral triple gives the Standard Model Yukawa edge skeleton, while T_D4 has only the right trilinear arity to be tested as an airlocked edge-kernel shape.", Final: "Gate 804 proves edge-kernel compatibility only. It does not produce D4 slot embeddings, real-form descent, Hermitian Yukawa matrices, generation operators, singular values, N_eff, PMNS/CKM, or C_Higgs updates. The next precise obstruction is EdgeTrialityEmbeddingSeal."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate803Inherited, StatusFSTSelected, StatusEdgeTemplateRecorded, StatusCompatibilityTarget,
		StatusSlotMatchingAudited, StatusFourSectorFirewall, StatusGaugeCompatAudited, StatusHiggsCompatAudited,
		StatusHermitianReaudited, StatusGenerationRecorded, StatusTraceCompatAudited, StatusTopColorFirewall,
		StatusOutcomeRecorded, StatusPackageUpdated, StatusCHiggsFirewall, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusFSTEdgeSkeleton, StatusTD4KernelShape, StatusTD4CorrectArity, StatusTD4UniversalKernelOnly,
		StatusFSTPartialGaugeAssignment, StatusHiggsOneFormCandidate, StatusTD4PreTraceKernel,
		StatusFSTTraceColorThree, StatusArityOnlyCompatibility, StatusFSTSuppliesPartialSkeleton,
		StatusNextEdgeTrialityEmbedding, StatusTD4AloneNotLedger, StatusEdgeNoEigenvalues, StatusEdgeNoMixing,
		StatusKernelNotReadout, StatusArityNoEmbedding, StatusNoHLREmbedding, StatusThreeSlotsNotFour,
		StatusTD4NoSectorReplacement, StatusD4NoGaugeAssignment, StatusNoEdgeToD4Theorem,
		StatusNoHiggsSlotEmbedding, StatusK7PlusNotD4Vector, StatusKernelNoYFMatrix, StatusNoGenerationOperator,
		StatusNoYdaggerYTrace, StatusFSTNoNativeGenerations, StatusTD4NoGenerations, StatusNoPMNSCKM,
		StatusTD4NoTraceInputs, StatusTD4NoABNEffUpdate, StatusEdgeNoTopDominance, StatusKernelNoNEffMinusThree,
		StatusNoEdgeEmbeddingReadout, StatusReadoutStillMissing, StatusNoCYukawaUpdate, StatusCHiggsLevelB,
		StatusFirewallGate804,
	}
}

func FormatEdgeTemplate(e EdgeTemplate) string {
	return fmt.Sprintf("edges=[%s] knows=[%s] missing=[%s] supports=[%s] failures=[%s]", strings.Join(e.Edges, "; "), strings.Join(e.Knows, "; "), strings.Join(e.Missing, "; "), strings.Join(e.Supports, "; "), strings.Join(e.Failures, "; "))
}

func FormatTarget(t CompatibilityTarget) string {
	return fmt.Sprintf("%s components=[%s] lawful=%q rejected=%q supports=[%s] failures=[%s]", t.SealName, strings.Join(t.Components, "; "), t.LawfulGoal, t.RejectedGoal, strings.Join(t.Supports, "; "), strings.Join(t.Failures, "; "))
}

func FormatSlot(s SlotMatching) string {
	return fmt.Sprintf("D4=[%s] Yukawa=[%s] candidate=%s required=[%s] supports=[%s] failures=[%s]", strings.Join(s.D4Slots, "; "), strings.Join(s.YukawaSlots, "; "), s.Candidate, strings.Join(s.Required, "; "), strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatFour(f FourSectorFirewall) string {
	return fmt.Sprintf("sectors=[%s] D4slots=[%s] lawful=%q blocked=%q supports=[%s] failures=[%s]", strings.Join(f.Sectors, "; "), strings.Join(f.D4Slots, "; "), f.LawfulForm, f.BlockedClaim, strings.Join(f.Supports, "; "), strings.Join(f.Failures, "; "))
}

func FormatGauge(g GaugeCompatibility) string {
	return fmt.Sprintf("finiteTriple=%q T_D4=%q missing=[%s] supports=[%s] failures=[%s]", g.FiniteTripleRole, g.TD4Role, strings.Join(g.Missing, "; "), strings.Join(g.Supports, "; "), strings.Join(g.Failures, "; "))
}

func FormatHiggs(h HiggsCompatibility) string {
	return fmt.Sprintf("%s candidate=%q components=[%s] supports=[%s] failures=[%s]", h.RequiredSeal, h.Candidate, strings.Join(h.Components, "; "), strings.Join(h.Supports, "; "), strings.Join(h.Failures, "; "))
}

func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("%s reason=%q supports=[%s] failures=[%s]", o.Subject, o.Reason, strings.Join(o.Supports, "; "), strings.Join(o.Failures, "; "))
}

func FormatTrace(t TraceCompatibility) string {
	return fmt.Sprintf("forms=[%s] question=%q answer=%q supports=[%s] failures=[%s]", strings.Join(t.TraceForms, "; "), t.Question, t.Answer, strings.Join(t.Supports, "; "), strings.Join(t.Failures, "; "))
}

func FormatOutcome(o Outcome) string {
	return fmt.Sprintf("items=[%s] supports=[%s] failures=[%s]", strings.Join(o.Items, "; "), strings.Join(o.Supports, "; "), strings.Join(o.Failures, "; "))
}

func FormatPackage(p PackageStatus) string {
	return fmt.Sprintf("partial=[%s] missing=[%s] supports=[%s] failures=[%s]", strings.Join(p.PartiallySupplied, "; "), strings.Join(p.NotSupplied, "; "), strings.Join(p.Supports, "; "), strings.Join(p.Failures, "; "))
}

func FormatCHiggs(c CHiggsFirewall) string {
	return fmt.Sprintf("%s unchanged=[%s] failures=[%s]", c.Formula, strings.Join(c.Unchanged, "; "), strings.Join(c.Failures, "; "))
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
