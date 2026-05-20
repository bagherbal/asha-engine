// Package generation2edgetrialityembeddingsealandhiggsfermionslotassignmentnogoaudit implements
// Gate 805: EdgeTrialityEmbeddingSeal and Higgs/Fermion Slot Assignment No-Go Audit.
//
// Gate 805 inherits Gate 804's arity-only compatibility between the finite
// spectral triple Yukawa edge skeleton and the complex-airlocked D4 trilinear.
// It audits whether the Higgs/fermion edge data can actually be embedded into
// the three D4 triality slots without violating real-form, gauge, chirality,
// Higgs-socket, boson/fermion-role, or generation firewalls.
package generation2edgetrialityembeddingsealandhiggsfermionslotassignmentnogoaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE805-EDGE-TRIALITY-EMBEDDING-SEAL-HIGGS-FERMION-SLOT-ASSIGNMENT-NO-GO-AUDIT"

	StatusGate804Inherited     = "PASS_GATE804_FINITE_TRIPLE_TRIALITY_EDGE_COMPATIBILITY_INHERITED"
	StatusAritySelected        = "PASS_ARITY_COMPATIBILITY_SELECTED_FOR_SHARP_EMBEDDING_AUDIT"
	StatusSealDefined          = "PASS_EDGE_TRIALITY_EMBEDDING_SEAL_DEFINED"
	StatusCanonicalSlotAudited = "PASS_CANONICAL_VECTOR_SPINOR_SPINOR_SLOT_CANDIDATE_AUDITED"
	StatusPermutedSlotsAudited = "PASS_TRIALITY_PERMUTED_SLOT_CANDIDATES_AUDITED"
	StatusHiggsSlotAudited     = "PASS_HIGGS_SLOT_EMBEDDING_AUDITED"
	StatusFermionSlotAudited   = "PASS_FERMION_SLOT_EMBEDDING_AUDITED"
	StatusChiralityFirewall    = "PASS_CHIRALITY_FIREWALL_AUDITED"
	StatusGaugeLabelAudited    = "PASS_GAUGE_LABEL_PRESERVATION_AUDITED"
	StatusSectorUniversality   = "PASS_SECTOR_UNIVERSALITY_AUDITED"
	StatusHermitianMatrix      = "PASS_HERMITIAN_MATRIX_OBSTRUCTION_AUDITED"
	StatusRealFormDescent      = "PASS_REAL_FORM_DESCENT_OBSTRUCTION_REAUDITED"
	StatusCandidateTable       = "PASS_SLOT_ASSIGNMENT_CANDIDATE_TABLE_RECORDED"
	StatusPackageUpdated       = "PASS_TRIALITY_YUKAWA_READOUT_PACKAGE_STATUS_UPDATED"
	StatusCHiggsFirewall       = "PASS_C_HIGGS_FIREWALL_PRESERVED"
	StatusOutcomeRecorded      = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchDecision       = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusEdgeEmbeddingRequired      = "CONDITIONAL_SUPPORT_EDGE_EMBEDDING_IS_REQUIRED_BEFORE_T_D4_CAN_HOST_FINITE_TRIPLE_EDGES"
	StatusHiggsVectorCandidate       = "CONDITIONAL_SUPPORT_HIGGS_AS_VECTOR_SLOT_IS_STRONGEST_FORMAL_T_D4_YUKAWA_KERNEL_CANDIDATE"
	StatusPermutationsAirlocked      = "CONDITIONAL_SUPPORT_TRIALITY_PERMUTATIONS_EXIST_ONLY_INSIDE_COMPLEX_D4_AIRLOCK"
	StatusK7PlusHiggsCandidate       = "CONDITIONAL_SUPPORT_K7_PLUS_HIGGS_SOCKET_IS_ONLY_CURRENT_HIGGS_CARRIER_CANDIDATE"
	StatusUniversalKernelIfEmbedding = "CONDITIONAL_SUPPORT_T_D4_COULD_ONLY_BE_COMMON_EDGE_KERNEL_IF_EMBEDDINGS_EXIST"
	StatusCandidateAStrongest        = "CONDITIONAL_SUPPORT_CANDIDATE_A_IS_STRONGEST_FORMAL_EDGE_KERNEL_ASSIGNMENT"
	StatusBranchInterestingBlocked   = "CONDITIONAL_SUPPORT_EDGE_TRIALITY_BRANCH_IS_STRUCTURALLY_INTERESTING_BUT_EMBEDDING_BLOCKED"
	StatusNextGenerationOperator     = "CONDITIONAL_SUPPORT_NEXT_NATIVE_GATE_SHOULD_AUDIT_GENERATION_OPERATOR_SEAL"

	StatusNoEmbeddingSeal        = "FAILED_ROUTE_NO_EDGE_TRIALITY_EMBEDDING_SEAL_CURRENTLY_CERTIFIED"
	StatusNoHiggsToVector        = "FAILED_ROUTE_NO_CERTIFIED_HIGGS_TO_VECTOR_SLOT_EMBEDDING"
	StatusNoFermionsToSpinors    = "FAILED_ROUTE_NO_CERTIFIED_LEFT_RIGHT_FERMION_TO_HALF_SPINOR_SLOT_EMBEDDING"
	StatusPermutationRoleFail    = "FAILED_ROUTE_TRIALITY_PERMUTATION_DOES_NOT_PRESERVE_PHYSICAL_HIGGS_FERMION_ROLE"
	StatusParityFail             = "FAILED_ROUTE_BOSON_FERMION_PARITY_NOT_PRESERVED_BY_UNTYPED_TRIALITY_SLOT_SWAP"
	StatusK7PlusC2NotD4C8        = "FAILED_ROUTE_K7_PLUS_C2_NOT_D4_VECTOR_C8"
	StatusNoHiggsC2ToD4C8        = "FAILED_ROUTE_NO_CANONICAL_HIGGS_C2_TO_D4_C8_EMBEDDING"
	StatusHiggsOneFormNotD4      = "FAILED_ROUTE_FINITE_HIGGS_ONE_FORM_NOT_IDENTIFIED_WITH_D4_VECTOR_SLOT"
	StatusNoSMFermionToD4Spinor  = "FAILED_ROUTE_NO_STANDARD_MODEL_FERMION_CARRIER_TO_D4_HALF_SPINOR_EMBEDDING"
	StatusChiralityNotD4         = "FAILED_ROUTE_LEFT_RIGHT_FINITE_TRIPLE_CHIRALITY_NOT_CERTIFIED_AS_D4_HALF_SPINOR_CHIRALITY"
	StatusNoSectorEmbeddings     = "FAILED_ROUTE_SECTOR_DEPENDENT_FERMION_EMBEDDINGS_NOT_SUPPLIED"
	StatusSMChiralityNotD4       = "FAILED_ROUTE_STANDARD_MODEL_LEFT_RIGHT_CHIRALITY_NOT_AUTOMATICALLY_D4_HALF_SPINOR_CHIRALITY"
	StatusNoChiralitySeal        = "FAILED_ROUTE_NO_CHIRALITY_COMPATIBILITY_SEAL"
	StatusD4NoGaugeLabels        = "FAILED_ROUTE_D4_SLOTS_DO_NOT_CARRY_STANDARD_MODEL_GAUGE_LABELS_BY_DEFAULT"
	StatusNoGaugePreservingMap   = "FAILED_ROUTE_NO_GAUGE_LABEL_PRESERVING_EDGE_EMBEDDING_MAP"
	StatusNoHyperchargeFromSlot  = "FAILED_ROUTE_HYPERCHARGE_ASSIGNMENT_NOT_DERIVED_FROM_T_D4_SLOT_ASSIGNMENT"
	StatusUniversalNoSectors     = "FAILED_ROUTE_UNIVERSAL_KERNEL_DOES_NOT_EXPLAIN_SECTOR_DIFFERENTIATION"
	StatusUniversalNoHierarchy   = "FAILED_ROUTE_UNIVERSAL_KERNEL_DOES_NOT_EXPLAIN_YUKAWA_HIERARCHY"
	StatusEmbeddingNoGenOperator = "FAILED_ROUTE_EDGE_EMBEDDING_DOES_NOT_SUPPLY_GENERATION_OPERATOR"
	StatusKernelNoYF             = "FAILED_ROUTE_EDGE_KERNEL_DOES_NOT_SUPPLY_Y_F_MATRIX"
	StatusNoYdaggerY             = "FAILED_ROUTE_NO_Y_DAGGER_Y_TRACE_FROM_SLOT_ASSIGNMENT"
	StatusComplexNotNative       = "FAILED_ROUTE_COMPLEX_SLOT_ASSIGNMENT_NOT_NATIVE_REAL_CL17_EDGE_EMBEDDING"
	StatusNoRealDescent          = "FAILED_ROUTE_NO_REAL_DESCENT_MAP_FOR_EDGE_TRIALITY_EMBEDDING"
	StatusNoCandidateCertified   = "FAILED_ROUTE_NO_CANDIDATE_CURRENTLY_CERTIFIES_EDGE_TRIALITY_EMBEDDING"
	StatusEmbeddingSealMissing   = "FAILED_ROUTE_EDGE_TRIALITY_EMBEDDING_SEAL_NOT_SUPPLIED"
	StatusReadoutStillMissing    = "FAILED_ROUTE_REMAINING_READOUT_PACKAGE_STILL_MISSING"
	StatusNoCYukawaUpdate        = "FAILED_ROUTE_EDGE_TRIALITY_EMBEDDING_NO_GO_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate805        = "FIREWALL_PRESERVED_GATE805_EDGE_TRIALITY_EMBEDDING_BOUNDARY"
)

type Inheritance struct {
	Shape                 string
	FiniteTripleEdges     []string
	ArityOnly             bool
	CarrierEmbeddingFound bool
	Verdicts              []string
}

type SealDefinition struct {
	Defined    bool
	Name       string
	Components []string
	SectorRule string
	Verdict    string
	Supports   []string
	Failures   []string
}

type SlotCandidate struct {
	Audited    bool
	Name       string
	Assignment map[string]string
	Reason     string
	Required   []string
	Verdict    string
	Supports   []string
	Failures   []string
}

type SlotPermutationAudit struct {
	Audited        bool
	Candidates     []string
	LawfulDomain   string
	PhysicalDanger string
	Verdict        string
	Supports       []string
	Failures       []string
}

type EmbeddingAudit struct {
	Audited  bool
	Subject  string
	Current  []string
	Target   []string
	Required []string
	Verdict  string
	Supports []string
	Failures []string
}

type FirewallAudit struct {
	Audited  bool
	Subject  string
	Reason   string
	Verdict  string
	Supports []string
	Failures []string
}

type CandidateRow struct {
	Name       string
	Assignment string
	Status     string
}

type CandidateTable struct {
	Recorded bool
	Rows     []CandidateRow
	Verdict  string
	Supports []string
	Failures []string
}

type PackageStatus struct {
	Updated       bool
	SuppliedByFST []string
	NotSupplied   []string
	Verdict       string
	Failures      []string
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
	Seal         SealDefinition
	Canonical    SlotCandidate
	Permutations SlotPermutationAudit
	Higgs        EmbeddingAudit
	Fermion      EmbeddingAudit
	Chirality    FirewallAudit
	Gauge        FirewallAudit
	Sector       FirewallAudit
	Hermitian    FirewallAudit
	RealForm     FirewallAudit
	Table        CandidateTable
	Package      PackageStatus
	CHiggs       CHiggsFirewall
	Outcome      Outcome
	Branch       BranchDecision
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func BuildDefault() (Analysis, error) {
	inheritance := Inheritance{
		Shape:                 "T_D4(H, psi_L, psi_R) with Higgs, left fermion, and right fermion slots",
		FiniteTripleEdges:     []string{"Y_u: Q_L -> u_R", "Y_d: Q_L -> d_R", "Y_e: L_L -> e_R", "Y_nu: L_L -> nu_R or chosen neutrino convention"},
		ArityOnly:             true,
		CarrierEmbeddingFound: false,
		Verdicts:              []string{StatusGate804Inherited, StatusAritySelected, "FAILED_ROUTE_ARITY_MATCH_DOES_NOT_PROVE_SLOT_ASSIGNMENT"},
	}
	if !inheritance.ArityOnly || inheritance.CarrierEmbeddingFound {
		return Analysis{}, fmt.Errorf("Gate 805 requires arity-only inheritance and no certified carrier embedding")
	}

	seal := SealDefinition{
		Defined:    true,
		Name:       "EdgeTrialityEmbeddingSeal",
		Components: []string{"finite spectral triple edge carrier E_f", "D4 triality slot assignment", "Higgs-slot embedding", "left-fermion slot embedding", "right-fermion slot embedding", "real-form descent", "gauge-label preservation", "chirality compatibility", "boson/fermion parity firewall", "normalization convention"},
		SectorRule: "for each f in {u,d,e,nu}, require E_f -> V_C x S_plus_C x S_minus_C or a triality-equivalent permutation",
		Verdict:    StatusSealDefined,
		Supports:   []string{StatusEdgeEmbeddingRequired},
		Failures:   []string{StatusNoEmbeddingSeal},
	}

	canonical := SlotCandidate{
		Audited:    true,
		Name:       "Candidate A — vector-spinor-spinor",
		Assignment: map[string]string{"Higgs": "V_C", "psi_L": "S_plus_C", "psi_R": "S_minus_C"},
		Reason:     "strongest formal Clifford shape: gamma(H): S_plus_C -> S_minus_C and T_D4(H,psi_L,psi_R)=<gamma(H)psi_L,psi_R>",
		Required:   []string{"Higgs one-form carrier -> V_C", "left finite fermion carrier -> S_plus_C", "right finite fermion carrier -> S_minus_C"},
		Verdict:    StatusCanonicalSlotAudited,
		Supports:   []string{StatusHiggsVectorCandidate},
		Failures:   []string{StatusNoHiggsToVector, StatusNoFermionsToSpinors},
	}

	permutations := SlotPermutationAudit{
		Audited:        true,
		Candidates:     []string{"Higgs -> S_plus_C, left -> S_minus_C, right -> V_C", "Higgs -> S_minus_C, left -> V_C, right -> S_plus_C"},
		LawfulDomain:   "complex D4 airlock representation-type permutations",
		PhysicalDanger: "representation-type triality is not physical spin-statistics, Standard Model chirality, Higgs/fermion parity, or gauge representation assignment",
		Verdict:        StatusPermutedSlotsAudited,
		Supports:       []string{StatusPermutationsAirlocked},
		Failures:       []string{StatusPermutationRoleFail, StatusParityFail},
	}

	higgs := EmbeddingAudit{
		Audited:  true,
		Subject:  "Higgs slot embedding",
		Current:  []string{"K7+_J(n) ~= C^2", "K7+ ~= R^4", "finite spectral triple Higgs one-form / Higgs doublet"},
		Target:   []string{"V_C, dim_C=8", "or another D4 triality slot"},
		Required: []string{"nontrivial embedding or projection", "specified complement", "normalization", "real-form airlock compatibility"},
		Verdict:  StatusHiggsSlotAudited,
		Supports: []string{StatusK7PlusHiggsCandidate},
		Failures: []string{StatusK7PlusC2NotD4C8, StatusNoHiggsC2ToD4C8, StatusHiggsOneFormNotD4},
	}

	fermion := EmbeddingAudit{
		Audited:  true,
		Subject:  "Fermion slot embedding",
		Current:  []string{"Q_L", "L_L", "u_R", "d_R", "e_R", "nu_R", "gauge labels", "chirality labels", "color/lepton distinctions"},
		Target:   []string{"S_plus_C", "S_minus_C"},
		Required: []string{"Q_L or L_L -> S_plus_C", "u_R,d_R,e_R,nu_R -> S_minus_C", "sector-dependent variants if needed"},
		Verdict:  StatusFermionSlotAudited,
		Failures: []string{StatusNoSMFermionToD4Spinor, StatusChiralityNotD4, StatusNoSectorEmbeddings},
	}

	chirality := FirewallAudit{Audited: true, Subject: "Chirality compatibility", Reason: "finite triple left/right chirality and complex D4 half-spinor chirality are not automatically the same grading", Verdict: StatusChiralityFirewall, Failures: []string{StatusSMChiralityNotD4, StatusNoChiralitySeal}}
	gauge := FirewallAudit{Audited: true, Subject: "Gauge-label preservation", Reason: "D4 slots know representation type only; Standard Model SU(3)c, SU(2)L, U(1)Y, quark/lepton, up/down, charged/neutral labels require a separate embedding map", Verdict: StatusGaugeLabelAudited, Failures: []string{StatusD4NoGaugeLabels, StatusNoGaugePreservingMap, StatusNoHyperchargeFromSlot}}
	sector := FirewallAudit{Audited: true, Subject: "Sector universality", Reason: "a universal T_D4 edge kernel could only supply a shared structural form; it cannot explain four sector matrices, sector differentiation, hierarchy, Clebsch factors, or PMNS/CKM misalignment", Verdict: StatusSectorUniversality, Supports: []string{StatusUniversalKernelIfEmbedding}, Failures: []string{StatusUniversalNoSectors, StatusUniversalNoHierarchy}}
	hermitian := FirewallAudit{Audited: true, Subject: "Hermitian matrix trace extraction", Reason: "slot assignment gives at most H x psi_L x psi_R -> scalar; finite spectral action needs Y_f matrices, Y_f†Y_f, and traces over generation space", Verdict: StatusHermitianMatrix, Failures: []string{StatusEmbeddingNoGenOperator, StatusKernelNoYF, StatusNoYdaggerY}}
	realForm := FirewallAudit{Audited: true, Subject: "Real-form descent", Reason: "all D4 slot data remain under ComplexD4TrialityAirlock; a native ASHA edge embedding requires descent to the real Cl(1,7)/finite triple board", Verdict: StatusRealFormDescent, Failures: []string{StatusComplexNotNative, StatusNoRealDescent}}

	table := CandidateTable{
		Recorded: true,
		Rows: []CandidateRow{
			{Name: "Candidate A", Assignment: "H -> V_C, psi_L -> S_plus_C, psi_R -> S_minus_C", Status: "strongest formal arity match; blocked by Higgs C2-to-C8 embedding, fermion slot embedding, real descent, and gauge labels"},
			{Name: "Candidate B", Assignment: "H -> S_plus_C, psi_L -> V_C, psi_R -> S_minus_C", Status: "triality-permuted; blocked by boson/fermion role firewall and gauge-label mismatch"},
			{Name: "Candidate C", Assignment: "H -> S_minus_C, psi_L -> S_plus_C, psi_R -> V_C", Status: "triality-permuted; blocked by boson/fermion role firewall and gauge-label mismatch"},
			{Name: "Candidate D", Assignment: "sector-dependent embeddings E_f -> D4 slots", Status: "formally flexible but requires four independent embedding seals and is no longer explained by triality alone"},
		},
		Verdict:  StatusCandidateTable,
		Supports: []string{StatusCandidateAStrongest},
		Failures: []string{StatusNoCandidateCertified},
	}

	pkg := PackageStatus{
		Updated:       true,
		SuppliedByFST: []string{"GaugeRepresentationAssignmentSeal: still supplied by finite spectral triple", "SectorAssignmentSeal: still supplied by finite spectral triple edge labels"},
		NotSupplied:   []string{"EdgeTrialityEmbeddingSeal", "HiggsSlotEmbeddingSeal", "FermionSlotEmbeddingSeal", "RealDescentSeal", "GenerationCarrierSeal", "HermitianOperatorSeal", "TraceAtomExtractionSeal"},
		Verdict:       StatusPackageUpdated,
		Failures:      []string{StatusEmbeddingSealMissing, StatusReadoutStillMissing},
	}

	chiggs := CHiggsFirewall{Preserved: true, Formula: "C_Higgs=(3/N_eff)C_History", Unchanged: []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs", "lambda_H_bridge", "m_H_tree_proxy"}, Verdict: StatusCHiggsFirewall, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}}
	outcome := Outcome{Recorded: true, Items: []string{"T_D4 has the right formal arity for a Yukawa edge kernel", "finite spectral triple supplies the Standard Model edge skeleton", "no Higgs one-form / K7+ socket to D4 slot map", "no finite triple left/right fermion to D4 half-spinor map", "no real-form descent or gauge-label preservation theorem", "no Hermitian generation operator or trace atom readout"}, Verdict: StatusOutcomeRecorded, Supports: []string{StatusBranchInterestingBlocked}}
	branch := BranchDecision{Recorded: true, Next: "Gate 806 — GenerationOperatorSeal and Yukawa Matrix Source Minimality Audit", Seal: "GenerationOperatorSeal", Purpose: "audit the minimal object required to turn finite spectral triple edge templates into sector Yukawa matrices Y_u,Y_d,Y_e,Y_nu", Verdict: StatusBranchDecision, Supports: []string{StatusNextGenerationOperator}}
	firewalls := Firewalls{Enforced: true, NoYukawa: true, NoEigenvalues: true, NoPMNSCKM: true, NoFlavor: true, NoNEff: true, NoGJ: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoNativeTriality: true, NoHistoryLoop: true, Verdict: StatusFirewallGate805}

	return Analysis{Inheritance: inheritance, Seal: seal, Canonical: canonical, Permutations: permutations, Higgs: higgs, Fermion: fermion, Chirality: chirality, Gauge: gauge, Sector: sector, Hermitian: hermitian, RealForm: realForm, Table: table, Package: pkg, CHiggs: chiggs, Outcome: outcome, Branch: branch, Firewalls: firewalls, Truth: "Gate 805 finds that T_D4 has the right trilinear arity to resemble a Yukawa edge kernel, but no typed Higgs/fermion slot embedding is certified.", Final: "Triality remains an airlocked structural guide, not a Yukawa source. The next native obstruction is GenerationOperatorSeal."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate804Inherited, StatusAritySelected, StatusSealDefined, StatusCanonicalSlotAudited,
		StatusPermutedSlotsAudited, StatusHiggsSlotAudited, StatusFermionSlotAudited, StatusChiralityFirewall,
		StatusGaugeLabelAudited, StatusSectorUniversality, StatusHermitianMatrix, StatusRealFormDescent,
		StatusCandidateTable, StatusPackageUpdated, StatusCHiggsFirewall, StatusOutcomeRecorded, StatusBranchDecision,
		StatusPhysicalFirewalls, StatusEdgeEmbeddingRequired, StatusHiggsVectorCandidate, StatusPermutationsAirlocked,
		StatusK7PlusHiggsCandidate, StatusUniversalKernelIfEmbedding, StatusCandidateAStrongest,
		StatusBranchInterestingBlocked, StatusNextGenerationOperator, StatusNoEmbeddingSeal, StatusNoHiggsToVector,
		StatusNoFermionsToSpinors, StatusPermutationRoleFail, StatusParityFail, StatusK7PlusC2NotD4C8,
		StatusNoHiggsC2ToD4C8, StatusHiggsOneFormNotD4, StatusNoSMFermionToD4Spinor, StatusChiralityNotD4,
		StatusNoSectorEmbeddings, StatusSMChiralityNotD4, StatusNoChiralitySeal, StatusD4NoGaugeLabels,
		StatusNoGaugePreservingMap, StatusNoHyperchargeFromSlot, StatusUniversalNoSectors, StatusUniversalNoHierarchy,
		StatusEmbeddingNoGenOperator, StatusKernelNoYF, StatusNoYdaggerY, StatusComplexNotNative, StatusNoRealDescent,
		StatusNoCandidateCertified, StatusEmbeddingSealMissing, StatusReadoutStillMissing, StatusNoCYukawaUpdate,
		StatusCHiggsLevelB, StatusFirewallGate805,
	}
}

func FormatSeal(s SealDefinition) string {
	return fmt.Sprintf("%s components=[%s] sectorRule=%q supports=[%s] failures=[%s]", s.Name, strings.Join(s.Components, "; "), s.SectorRule, strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatSlotCandidate(c SlotCandidate) string {
	pairs := make([]string, 0, len(c.Assignment))
	for k, v := range c.Assignment {
		pairs = append(pairs, k+"->"+v)
	}
	return fmt.Sprintf("%s assignment=[%s] reason=%q required=[%s] supports=[%s] failures=[%s]", c.Name, strings.Join(pairs, "; "), c.Reason, strings.Join(c.Required, "; "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatPermutation(p SlotPermutationAudit) string {
	return fmt.Sprintf("candidates=[%s] domain=%q danger=%q supports=[%s] failures=[%s]", strings.Join(p.Candidates, "; "), p.LawfulDomain, p.PhysicalDanger, strings.Join(p.Supports, "; "), strings.Join(p.Failures, "; "))
}

func FormatEmbedding(e EmbeddingAudit) string {
	return fmt.Sprintf("%s current=[%s] target=[%s] required=[%s] supports=[%s] failures=[%s]", e.Subject, strings.Join(e.Current, "; "), strings.Join(e.Target, "; "), strings.Join(e.Required, "; "), strings.Join(e.Supports, "; "), strings.Join(e.Failures, "; "))
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("%s reason=%q supports=[%s] failures=[%s]", f.Subject, f.Reason, strings.Join(f.Supports, "; "), strings.Join(f.Failures, "; "))
}

func FormatTable(t CandidateTable) string {
	rows := make([]string, 0, len(t.Rows))
	for _, r := range t.Rows {
		rows = append(rows, r.Name+": "+r.Assignment+" => "+r.Status)
	}
	return fmt.Sprintf("rows=[%s] supports=[%s] failures=[%s]", strings.Join(rows, " | "), strings.Join(t.Supports, "; "), strings.Join(t.Failures, "; "))
}

func FormatPackage(p PackageStatus) string {
	return fmt.Sprintf("suppliedByFST=[%s] missing=[%s] failures=[%s]", strings.Join(p.SuppliedByFST, "; "), strings.Join(p.NotSupplied, "; "), strings.Join(p.Failures, "; "))
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
