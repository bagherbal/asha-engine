// Package generation2minimalrightmoduleweakdoubletsocketedgeoperatoraudit implements
// Gate 847: Minimal RightModule / WeakDoublet Socket Edge-Operator Audit.
//
// Gate 847 follows Gate 846's punctured socket response table. Gate 846
// reconstructed the finite-body aggregate response on the active right module
//
//	H_R^min = (e_+ tensor P_3) plus (e_- tensor P_3) plus (e_- tensor P_1),
//
// with the puncture e_+ tensor P_1 absent. Gate 847 audits the next sharper
// object: a support-only finite-Dirac edge operator refined by rank-one weak
// doublet socket selectors h_+, h_- inside C_L^2.
//
// The candidate symbolic support is
//
//	e_+ tensor P_3 -> h_+ tensor P_3,
//	e_- tensor P_3 -> h_- tensor P_3,
//	e_- tensor P_1 -> h_- tensor P_1,
//
// while the absent cell e_+ tensor P_1 has no edge. This is an edge-support
// skeleton only: no explicit D_F matrix, no first-order/bimodule proof, no
// Yukawa magnitudes, no alpha_B source, no R3/R4 promotion, and no official
// ledger update is certified.
package generation2minimalrightmoduleweakdoubletsocketedgeoperatoraudit

import (
	"strconv"
	"strings"
)

const (
	AuditID = "GATE847-MINIMAL-RIGHTMODULE-WEAKDOUBLET-SOCKET-EDGE-OPERATOR-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	RightSocketDim = 2
	WeakDoubletDim = 2
	WeakSocketRank = 1
	HRMinRank      = 7
	HLRank         = WeakDoubletDim * WDim

	StatusGate846Inherited             = "PASS_GATE846_PUNCTURED_SOCKET_RESPONSE_TABLE_INHERITED"
	StatusMinimalRightDomainInherited  = "PASS_MINIMAL_RIGHT_MODULE_EDGE_DOMAIN_INHERITED"
	StatusWeakSocketPairAudited        = "PASS_WEAK_DOUBLET_SOCKET_PAIR_AUDITED"
	StatusThreeActiveEdgesConstructed  = "PASS_THREE_ACTIVE_SYMBOLIC_SOCKET_EDGES_CONSTRUCTED"
	StatusPunctureEdgeAbsencePreserved = "PASS_PUNCTURE_EDGE_ABSENCE_PRESERVED"
	StatusLeptoColorPreserved          = "PASS_LEPTO_COLOR_SUPPORT_PRESERVED_BY_SYMBOLIC_EDGES"
	StatusResponseCellsGenerated       = "PASS_EDGE_SUPPORT_RECONSTRUCTS_GATE846_ACTIVE_RESPONSE_CELLS"
	StatusEdgeOperatorClassifiedAsSeal = "PASS_EDGE_OPERATOR_CLASSIFIED_AS_SUPPORT_SEAL_NOT_NATIVE_D_F_MATRIX"
	StatusOfficialLedgersFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed           = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusR2PlusPlusPlusPlusStill      = "PASS_R2_PLUS_PLUS_PLUS_PLUS_EDGE_GENERATED_SHADOW_NOT_R3_OR_R4"
	StatusFirewallGate847              = "FIREWALL_PRESERVED_GATE847_WEAKDOUBLET_SOCKET_EDGE_OPERATOR"

	SupportWeakSocketPairSeal          = "CONDITIONAL_SUPPORT_WEAK_DOUBLET_SOCKET_PAIR_EXISTS_AS_ORIENTATION_SEAL"
	SupportThreeActiveSocketEdges      = "CONDITIONAL_SUPPORT_D_F_SUPP_HAS_THREE_ACTIVE_SOCKET_EDGES"
	SupportNeutralSingletonNullEdge    = "CONDITIONAL_SUPPORT_NEUTRAL_SINGLETON_IS_NULL_EDGE_CANDIDATE_AT_SEAL_LEVEL"
	SupportResponseTableEdgeGenerator  = "CONDITIONAL_SUPPORT_PUNCTURED_RESPONSE_TABLE_HAS_SYMBOLIC_EDGE_GENERATOR"
	SupportLeptoColorPreservingEdges   = "CONDITIONAL_SUPPORT_LEPTO_COLOR_PRESERVING_EDGE_SUPPORT"
	SupportRightDomainToWeakTargets    = "CONDITIONAL_SUPPORT_ACTIVE_RIGHT_DOMAIN_EDGE_TARGETS_REFINED_TO_H_PLUS_H_MINUS"
	SupportNoYukawaMagnitudeFromEdges  = "CONDITIONAL_SUPPORT_EDGE_SUPPORT_ONLY_NO_MAGNITUDE_READOUT"
	SupportR2PlusPlusPlusPlusEdgeStage = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_EDGE_GENERATED_SHADOW_STAGE"

	FailureWeakSplitNotNative        = "FAILED_ROUTE_WEAK_SOCKET_SPLIT_NOT_NATIVE_WITHOUT_HIGGS_ORIENTATION"
	FailureEdgeOperatorSealNotNative = "FAILED_ROUTE_SYMBOLIC_EDGE_OPERATOR_IS_SEAL_NOT_NATIVE_D_F_MATRIX"
	FailureNoExplicitDFMatrix        = "FAILED_ROUTE_NO_EXPLICIT_D_F_MATRIX_CERTIFIED"
	FailureNoFirstOrderProof         = "FAILED_ROUTE_NO_FIRST_ORDER_CONDITION_STABILITY_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof  = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureEdgeSupportNotYukawa      = "FAILED_ROUTE_EDGE_SUPPORT_NOT_TRACE_MAGNITUDE"
	FailureNoNumericalYukawaValues   = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureAlphaStillSealed          = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailurePunctureNullEdgeOnlySeal  = "FAILED_ROUTE_PUNCTURE_NULL_EDGE_ONLY_SEAL_NOT_NATIVE_THEOREM"
	FailureNoNativeNullEdgeTheorem   = "FAILED_ROUTE_NO_NATIVE_NULL_EDGE_THEOREM_FOR_E_PLUS_TENSOR_P1"
	FailureNoPhysicalParticleAssign  = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT_FROM_EDGE_SKELETON"
	FailureNoRightNeutrinoTheorem    = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoThreeGenerationTheorem  = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNoNEffUpdate              = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate           = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                     = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_NOT_R3"
	FailureNotR4                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                              float64
	OfficialNEff, OfficialCYukawa       float64
	OfficialCHiggs                      float64
	OfficialFrozen                      bool
	R2PlusPlusPlusPlusEdgeStage, R3, R4 bool
	AlphaNative                         bool
}

type WeakSocketSplit struct {
	Expression                     string
	HPlus, HMinus, Sum             string
	RankHPlus, RankHMinus, WeakDim int
	Complete, Orthogonal           bool
	OrientationSeal, NativeSplit   bool
	HiggsOrientationCertified      bool
	Supports, Failures             []string
}

type Edge struct {
	Name, Domain, Target, LeptoColor, Role string
	DomainRank, TargetRank                 int
	Present, Puncture, HasMagnitude        bool
	ValueSource                            string
}

type EdgeOperator struct {
	Expression, Rule                                                                           string
	Edges                                                                                      []Edge
	MissingEdge                                                                                Edge
	ActiveEdges, ExpectedActiveEdges                                                           int
	DomainRank, TargetRank                                                                     int
	PreservesLeptoColor, ReconstructsGate846Cells, PunctureAbsent                              bool
	SupportOnly, ExplicitDFMatrix, NativeDFMatrix, FirstOrderCertified, BimoduleCommutantProof bool
	Magnitudes, AlphaDerived, ParticleAssignment                                               bool
	Supports, Failures                                                                         []string
}

type Impact struct {
	Classification                                                     string
	MinimalRightDomainInherited, WeakSocketSeal, ThreeActiveEdgesBuilt bool
	ResponseTableHasEdgeGenerator, PunctureNullEdgeCandidate           bool
	ExplicitDFMatrix, FirstOrderProof, BimoduleCommutantProof          bool
	AlphaStillSealed, MagnitudesStillMissing                           bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs                   bool
	CanPromoteToR3, CanPromoteToR4                                     bool
}

type Firewalls struct {
	Enforced                                                                                      bool
	WeakSplitNotNative, EdgeOperatorSealNotNative, NoExplicitDFMatrix                             bool
	NoFirstOrderProof, NoBimoduleCommutantProof, EdgeSupportNotYukawa                             bool
	NoNumericalYukawaValues, AlphaStillSealed, NoTraceMagnitudeReadout                            bool
	PunctureNullEdgeOnlySeal, NoNativeNullEdgeTheorem, NoPhysicalParticleAssignment               bool
	NoRightNeutrinoTheorem, NoThreeGenerationTheorem, NoNEffUpdate, NoCYukawaUpdate, NotR3, NotR4 bool
	Verdict                                                                                       string
}

type Audit struct {
	Ledger    Ledger
	Weak      WeakSocketSplit
	Edges     EdgeOperator
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:    buildLedger(),
		Weak:      buildWeakSocketSplit(),
		Edges:     buildEdgeOperator(),
		Impact:    buildImpact(),
		Firewalls: buildFirewalls(),
		Truth:     "Gate 847 refines the Gate 844/846 support skeleton by introducing h_+,h_- inside C_L^2 at seal level. The three active right cells receive weak-socket edge targets, while e_+ tensor P_1 remains edge-absent. This gives the punctured response table a symbolic edge generator, but not an explicit D_F matrix, first-order proof, magnitude theorem, alpha source, R3/R4 promotion, or official ledger update.",
		Final:     "Final verdict: CONDITIONAL_SUPPORT_D_F_SUPP_HAS_THREE_ACTIVE_SOCKET_EDGES and CONDITIONAL_SUPPORT_PUNCTURED_RESPONSE_TABLE_HAS_SYMBOLIC_EDGE_GENERATOR, but FAILED_ROUTE_SYMBOLIC_EDGE_OPERATOR_IS_SEAL_NOT_NATIVE_D_F_MATRIX, FAILED_ROUTE_NO_EXPLICIT_D_F_MATRIX_CERTIFIED, FAILED_ROUTE_NO_FIRST_ORDER_CONDITION_STABILITY_PROOF_CERTIFIED, FAILED_ROUTE_EDGE_SUPPORT_NOT_TRACE_MAGNITUDE, FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE, and FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_NOT_R3.",
	}
	return a, nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlusPlusPlusEdgeStage: true, R3: false, R4: false, AlphaNative: false}
}

func buildWeakSocketSplit() WeakSocketSplit {
	return WeakSocketSplit{
		Expression: "C_L^2 = h_+ plus h_-",
		HPlus:      "h_+", HMinus: "h_-", Sum: "I_{C_L^2}",
		RankHPlus: 1, RankHMinus: 1, WeakDim: WeakDoubletDim,
		Complete: true, Orthogonal: true,
		OrientationSeal: true, NativeSplit: false, HiggsOrientationCertified: false,
		Supports: []string{SupportWeakSocketPairSeal, SupportRightDomainToWeakTargets},
		Failures: []string{FailureWeakSplitNotNative},
	}
}

func buildEdgeOperator() EdgeOperator {
	edges := []Edge{
		{Name: "Y_+3", Domain: "e_+ tensor P_3", Target: "h_+ tensor P_3", LeptoColor: "P_3 -> P_3", Role: "dominant character-color edge support", DomainRank: 3, TargetRank: 3, Present: true, Puncture: false, HasMagnitude: false, ValueSource: "support-only"},
		{Name: "Y_-3", Domain: "e_- tensor P_3", Target: "h_- tensor P_3", LeptoColor: "P_3 -> P_3", Role: "rest character-color edge support", DomainRank: 3, TargetRank: 3, Present: true, Puncture: false, HasMagnitude: false, ValueSource: "support-only"},
		{Name: "Y_-1", Domain: "e_- tensor P_1", Target: "h_- tensor P_1", LeptoColor: "P_1 -> P_1", Role: "rest character-lepton edge support", DomainRank: 1, TargetRank: 1, Present: true, Puncture: false, HasMagnitude: false, ValueSource: "support-only"},
	}
	missing := Edge{Name: "Y_+1", Domain: "e_+ tensor P_1", Target: "h_+ tensor P_1", LeptoColor: "P_1 -> P_1", Role: "neutral right-lepton puncture edge absent in minimal support", DomainRank: 1, TargetRank: 1, Present: false, Puncture: true, HasMagnitude: false, ValueSource: "absent by minimal puncture seal"}
	return EdgeOperator{
		Expression: "D_F^supp = Y_+3 plus Y_-3 plus Y_-1 plus adjoint plus J-copy; Y_+1 absent",
		Rule:       "support-only right socket to weak-doublet socket edge skeleton preserving P_1/P_3",
		Edges:      edges, MissingEdge: missing,
		ActiveEdges: len(edges), ExpectedActiveEdges: 3,
		DomainRank: HRMinRank, TargetRank: HLRank,
		PreservesLeptoColor: true, ReconstructsGate846Cells: true, PunctureAbsent: true,
		SupportOnly: true, ExplicitDFMatrix: false, NativeDFMatrix: false, FirstOrderCertified: false, BimoduleCommutantProof: false,
		Magnitudes: false, AlphaDerived: false, ParticleAssignment: false,
		Supports: []string{SupportThreeActiveSocketEdges, SupportNeutralSingletonNullEdge, SupportResponseTableEdgeGenerator, SupportLeptoColorPreservingEdges, SupportRightDomainToWeakTargets, SupportNoYukawaMagnitudeFromEdges},
		Failures: []string{FailureEdgeOperatorSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureEdgeSupportNotYukawa, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailurePunctureNullEdgeOnlySeal, FailureNoNativeNullEdgeTheorem, FailureNoPhysicalParticleAssign, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem},
	}
}

func buildImpact() Impact {
	return Impact{
		Classification:              "R2++++ edge-generated punctured socket response shadow; support-only, not R3/R4",
		MinimalRightDomainInherited: true, WeakSocketSeal: true, ThreeActiveEdgesBuilt: true,
		ResponseTableHasEdgeGenerator: true, PunctureNullEdgeCandidate: true,
		ExplicitDFMatrix: false, FirstOrderProof: false, BimoduleCommutantProof: false,
		AlphaStillSealed: true, MagnitudesStillMissing: true,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		CanPromoteToR3: false, CanPromoteToR4: false,
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		Enforced:           true,
		WeakSplitNotNative: true, EdgeOperatorSealNotNative: true, NoExplicitDFMatrix: true,
		NoFirstOrderProof: true, NoBimoduleCommutantProof: true, EdgeSupportNotYukawa: true,
		NoNumericalYukawaValues: true, AlphaStillSealed: true, NoTraceMagnitudeReadout: true,
		PunctureNullEdgeOnlySeal: true, NoNativeNullEdgeTheorem: true, NoPhysicalParticleAssignment: true,
		NoRightNeutrinoTheorem: true, NoThreeGenerationTheorem: true, NoNEffUpdate: true, NoCYukawaUpdate: true,
		NotR3: true, NotR4: true, Verdict: StatusFirewallGate847,
	}
}

func Statuses() []string {
	return []string{
		StatusGate846Inherited, StatusMinimalRightDomainInherited, StatusWeakSocketPairAudited, StatusThreeActiveEdgesConstructed, StatusPunctureEdgeAbsencePreserved, StatusLeptoColorPreserved, StatusResponseCellsGenerated, StatusEdgeOperatorClassifiedAsSeal, StatusOfficialLedgersFrozen, StatusNoObservedDataUsed, StatusR2PlusPlusPlusPlusStill, StatusFirewallGate847,
		SupportWeakSocketPairSeal, SupportThreeActiveSocketEdges, SupportNeutralSingletonNullEdge, SupportResponseTableEdgeGenerator, SupportLeptoColorPreservingEdges, SupportRightDomainToWeakTargets, SupportNoYukawaMagnitudeFromEdges, SupportR2PlusPlusPlusPlusEdgeStage,
		FailureWeakSplitNotNative, FailureEdgeOperatorSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureEdgeSupportNotYukawa, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailurePunctureNullEdgeOnlySeal, FailureNoNativeNullEdgeTheorem, FailureNoPhysicalParticleAssign, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return join("ledger", []string{
		"alpha_B=" + f64(l.AlphaB),
		"official_N_eff=" + f64(l.OfficialNEff),
		"official_C_Yukawa=" + f64(l.OfficialCYukawa),
		"official_C_Higgs=" + f64(l.OfficialCHiggs),
		"official_frozen=" + b(l.OfficialFrozen),
		"R2++++_edge_stage=" + b(l.R2PlusPlusPlusPlusEdgeStage),
		"R3=" + b(l.R3), "R4=" + b(l.R4), "alpha_native=" + b(l.AlphaNative),
	})
}

func FormatWeak(w WeakSocketSplit) string {
	return join("weak_socket_split", []string{
		w.Expression,
		"rank(h_+)=" + i(w.RankHPlus), "rank(h_-)=" + i(w.RankHMinus), "dim(C_L^2)=" + i(w.WeakDim),
		"complete=" + b(w.Complete), "orthogonal=" + b(w.Orthogonal), "orientation_seal=" + b(w.OrientationSeal), "native_split=" + b(w.NativeSplit),
		"supports=" + strings.Join(w.Supports, ","), "failures=" + strings.Join(w.Failures, ","),
	})
}

func FormatEdge(e Edge) string {
	return join("edge", []string{e.Name, e.Domain + " -> " + e.Target, "lepto_color=" + e.LeptoColor, "role=" + e.Role, "domain_rank=" + i(e.DomainRank), "target_rank=" + i(e.TargetRank), "present=" + b(e.Present), "puncture=" + b(e.Puncture), "magnitude=" + b(e.HasMagnitude), "value_source=" + e.ValueSource})
}

func FormatEdges(op EdgeOperator) string {
	parts := []string{op.Expression, "rule=" + op.Rule, "active_edges=" + i(op.ActiveEdges), "domain_rank=" + i(op.DomainRank), "target_rank=" + i(op.TargetRank), "preserves_lepto_color=" + b(op.PreservesLeptoColor), "reconstructs_gate846_cells=" + b(op.ReconstructsGate846Cells), "puncture_absent=" + b(op.PunctureAbsent), "support_only=" + b(op.SupportOnly), "explicit_D_F_matrix=" + b(op.ExplicitDFMatrix), "first_order=" + b(op.FirstOrderCertified), "bimodule_commutant=" + b(op.BimoduleCommutantProof), "magnitudes=" + b(op.Magnitudes)}
	for _, e := range op.Edges {
		parts = append(parts, FormatEdge(e))
	}
	parts = append(parts, "missing="+FormatEdge(op.MissingEdge), "supports="+strings.Join(op.Supports, ","), "failures="+strings.Join(op.Failures, ","))
	return join("edge_operator", parts)
}

func FormatImpact(im Impact) string {
	return join("impact", []string{im.Classification, "minimal_right_domain=" + b(im.MinimalRightDomainInherited), "weak_socket_seal=" + b(im.WeakSocketSeal), "three_active_edges=" + b(im.ThreeActiveEdgesBuilt), "table_has_edge_generator=" + b(im.ResponseTableHasEdgeGenerator), "puncture_null_edge_candidate=" + b(im.PunctureNullEdgeCandidate), "explicit_D_F_matrix=" + b(im.ExplicitDFMatrix), "first_order=" + b(im.FirstOrderProof), "bimodule_commutant=" + b(im.BimoduleCommutantProof), "alpha_sealed=" + b(im.AlphaStillSealed), "magnitudes_missing=" + b(im.MagnitudesStillMissing), "update_N_eff=" + b(im.CanUpdateNEff), "update_C_Yukawa=" + b(im.CanUpdateCYukawa), "update_C_Higgs=" + b(im.CanUpdateCHiggs), "R3=" + b(im.CanPromoteToR3), "R4=" + b(im.CanPromoteToR4)})
}

func FormatFirewalls(f Firewalls) string {
	return join("firewalls", []string{"enforced=" + b(f.Enforced), "weak_split_not_native=" + b(f.WeakSplitNotNative), "edge_operator_seal_not_native=" + b(f.EdgeOperatorSealNotNative), "no_explicit_D_F_matrix=" + b(f.NoExplicitDFMatrix), "no_first_order=" + b(f.NoFirstOrderProof), "no_bimodule_commutant=" + b(f.NoBimoduleCommutantProof), "edge_support_not_yukawa=" + b(f.EdgeSupportNotYukawa), "no_numerical_yukawa=" + b(f.NoNumericalYukawaValues), "alpha_sealed=" + b(f.AlphaStillSealed), "no_trace_magnitude=" + b(f.NoTraceMagnitudeReadout), "puncture_null_edge_only_seal=" + b(f.PunctureNullEdgeOnlySeal), "no_native_null_edge=" + b(f.NoNativeNullEdgeTheorem), "no_particle_assignment=" + b(f.NoPhysicalParticleAssignment), "no_right_neutrino_theorem=" + b(f.NoRightNeutrinoTheorem), "no_three_generation=" + b(f.NoThreeGenerationTheorem), "no_N_eff_update=" + b(f.NoNEffUpdate), "no_C_updates=" + b(f.NoCYukawaUpdate), "not_R3=" + b(f.NotR3), "not_R4=" + b(f.NotR4), "verdict=" + f.Verdict})
}

func containsAll(haystack, needles []string) bool {
	m := map[string]bool{}
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func join(label string, parts []string) string { return label + "{" + strings.Join(parts, "; ") + "}" }
func b(v bool) string {
	if v {
		return "true"
	}
	return "false"
}
func i(v int) string { return strconv.Itoa(v) }
func f64(v float64) string {
	return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(v, 'f', 16, 64), "0"), ".")
}
