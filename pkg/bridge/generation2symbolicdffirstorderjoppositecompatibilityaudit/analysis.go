// Package generation2symbolicdffirstorderjoppositecompatibilityaudit implements
// Gate 850: Symbolic D_F First-Order and J-Opposite Compatibility Audit.
//
// Gate 850 follows Gate 849's chiral neutral puncture/kernel pair.  Gate 848
// constructed the symbolic finite-Dirac support matrix
//
//	D_F^sym = [[0, Y_supp^dagger], [Y_supp, 0]],
//
// and Gate 849 showed that its full support-rank anatomy forces the left neutral
// kernel singleton h_+ tensor P_1 while the right neutral singleton e_+ tensor
// P_1 is absent from the minimal right module.  Gate 850 audits the next
// finite-triple compatibility layer: whether the symbolic D_F support matrix is
// compatible with a represented finite algebra rho_F(A_F), the J_F-opposite
// action, bimodule stability, and the first-order condition.  The result is a
// firewall audit: the support matrix has the correct chiral block form and the
// kernel singleton is a stability candidate, but no full rho_F/J_F/gamma_F/D_F
// package or first-order proof is certified.
package generation2symbolicdffirstorderjoppositecompatibilityaudit

import (
	"strconv"
	"strings"
)

const (
	AuditID = "GATE850-SYMBOLIC-D-F-FIRST-ORDER-J-OPPOSITE-COMPATIBILITY-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	WeakDoubletDim = 2
	HRMinRank      = 7
	HLRank         = WeakDoubletDim * WDim
	ChiralTotalDim = HRMinRank + HLRank

	StatusGate849Inherited           = "PASS_GATE849_CHIRAL_NEUTRAL_PUNCTURE_KERNEL_PAIR_INHERITED"
	StatusSymbolicDFInherited        = "PASS_SYMBOLIC_D_F_SUPPORT_MATRIX_INHERITED"
	StatusRepresentationRequirements = "PASS_REPRESENTATION_ACTION_REQUIREMENTS_AUDITED"
	StatusJOppositeRequirements      = "PASS_J_OPPOSITE_ACTION_REQUIREMENTS_AUDITED"
	StatusFirstOrderRequirements     = "PASS_FIRST_ORDER_CONDITION_REQUIREMENTS_AUDITED"
	StatusBimoduleRequirements       = "PASS_BIMODULE_STABILITY_REQUIREMENTS_AUDITED"
	StatusChiralityOddnessInherited  = "PASS_CHIRALITY_ODDNESS_REMAINS_SUPPORT_LEVEL_ONLY"
	StatusKernelStabilityAudited     = "PASS_LEFT_NEUTRAL_KERNEL_STABILITY_REQUIREMENTS_AUDITED"
	StatusNoObservedDataUsed         = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusOfficialLedgersFrozen      = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusCompatibilityFirewall      = "FIREWALL_PRESERVED_GATE850_FIRST_ORDER_AND_J_OPPOSITE_COMPATIBILITY"

	SupportCorrectChiralSupportForm    = "CONDITIONAL_SUPPORT_SYMBOLIC_D_F_HAS_CORRECT_CHIRAL_SUPPORT_FORM"
	SupportKernelStableCandidate       = "CONDITIONAL_SUPPORT_LEFT_NEUTRAL_KERNEL_SINGLETON_IS_REPRESENTATION_STABILITY_CANDIDATE"
	SupportPunctureKernelPairInherited = "CONDITIONAL_SUPPORT_CHIRAL_NEUTRAL_PUNCTURE_KERNEL_PAIR_INHERITED"
	SupportFirstOrderDataRequirements  = "CONDITIONAL_SUPPORT_FIRST_ORDER_PROOF_REQUIRES_COMPLETE_RHO_F_J_F_D_F_DATA"
	SupportR2CompatibilityStage        = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_COMPATIBILITY_STAGE"

	FailureNoCompleteRhoFActionLedger      = "FAILED_ROUTE_NO_COMPLETE_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoCompleteFiniteTriplePackage   = "FAILED_ROUTE_NO_COMPLETE_RHO_F_J_F_GAMMA_F_D_F_PACKAGE_CERTIFIED"
	FailureNoJOppositeCompatibilityProof   = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_CERTIFIED"
	FailureNoFullFirstOrderConditionProof  = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF"
	FailureNoBimoduleCommutantProof        = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoExplicitDFOperator            = "FAILED_ROUTE_NO_NUMERICAL_OR_OPERATOR_VALUED_D_F_MATRIX_CERTIFIED"
	FailureSymbolicDFNotNativeFiniteTriple = "FAILED_ROUTE_SYMBOLIC_D_F_SUPPORT_MATRIX_NOT_NATIVE_FINITE_TRIPLE"
	FailureKernelStabilityNotCertified     = "FAILED_ROUTE_KERNEL_SINGLETON_STABILITY_NOT_CERTIFIED_WITHOUT_RHO_F_AND_J_F"
	FailureChiralityOnlySupportLevel       = "FAILED_ROUTE_CHIRALITY_ODDNESS_REMAINS_SUPPORT_LEVEL_ONLY"
	FailureNoPhysicalNeutrinoTheorem       = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoRightNeutrinoTheorem          = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoMasslessnessTheorem           = "FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_OBSERVED_MASSLESSNESS_THEOREM"
	FailureSymbolicYNotYukawaMagnitude     = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureKernelNotYukawaMagnitude        = "FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_YUKAWA_MAGNITUDE"
	FailureAlphaStillSealed                = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout         = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate            = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate           = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                           = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_COMPATIBILITY_NOT_R3"
	FailureNotR4                           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	R2KernelCompatibilityStage    bool
	R3, R4, AlphaNative           bool
}

type Cell struct {
	Name, Support, Chirality, Socket, LeptoColor, Role string
	Rank                                               int
	Absent, Kernel, InMinimalSupport, StableCandidate  bool
	PhysicalAssignment                                 bool
}

type SymbolicDF struct {
	Expression, Domain, Codomain, Space                       string
	LeftRank, RightRank, TotalRank, YRank, DFRank, KernelDim  int
	SelfAdjointByBlock, ChiralityOddByBlock, SupportOnly      bool
	NativeFiniteTriple, ExplicitDFOperator, NumericalDFMatrix bool
	Supports, Failures                                        []string
}

type RepresentationData struct {
	Algebra, Representation, RequiredPackage             string
	RhoFAvailable, CompleteActionLedger, CompletePackage bool
	GammaFAvailable, JFAvailable, DFOperatorAvailable    bool
	CentralSupportLedger, KernelStabilityCertified       bool
	Kernel                                               Cell
	Supports, Failures                                   []string
}

type FirstOrderAudit struct {
	Expression                                                 string
	CanFormCommutator, CanFormOppositeAction, FirstOrderProven bool
	BimoduleStable, JOppositeCompatible, KernelStable          bool
	MissingData                                                []string
	Supports, Failures                                         []string
}

type ChiralityAudit struct {
	Expression                                      string
	SupportLevelOddness, NativeGammaCertified       bool
	KernelChiralityConsistent, PunctureChiralitySet bool
	Supports, Failures                              []string
}

type Impact struct {
	Classification                                                            string
	Gate849Inherited, SymbolicDFInherited, CompatibilityAudited               bool
	NativeFiniteTriple, FirstOrderCertified, JOppositeCertified, KernelStable bool
	PhysicalNeutrinoTheorem, MasslessTheorem                                  bool
	AlphaStillSealed, MagnitudesStillMissing                                  bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs                          bool
	CanPromoteToR3, CanPromoteToR4                                            bool
}

type Firewalls struct {
	Enforced                                                                                     bool
	NoRhoFActionLedger, NoCompletePackage, NoJOppositeProof, NoFirstOrderProof, NoBimoduleProof  bool
	NoExplicitDFOperator, SymbolicDFNotNative, KernelStabilityNotCertified, ChiralitySupportOnly bool
	NoPhysicalNeutrino, NoRightNeutrino, NoMasslessness, NoYukawaMagnitudes                      bool
	AlphaStillSealed, NoTraceMagnitudeReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate       bool
	NotR3, NotR4                                                                                 bool
	Verdict                                                                                      string
}

type Audit struct {
	Ledger         Ledger
	SymbolicDF     SymbolicDF
	Representation RepresentationData
	FirstOrder     FirstOrderAudit
	Chirality      ChiralityAudit
	Impact         Impact
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:         buildLedger(),
		SymbolicDF:     buildSymbolicDF(),
		Representation: buildRepresentationData(),
		FirstOrder:     buildFirstOrderAudit(),
		Chirality:      buildChiralityAudit(),
		Impact:         buildImpact(),
		Firewalls:      buildFirewalls(),
		Truth:          "Gate 850 follows Gate 849's chiral neutral puncture/kernel pair.  The symbolic support matrix D_F^sym has the correct left/right chiral block form and inherits the right puncture e_+ tensor P_1 and left kernel h_+ tensor P_1.  The next theorem-level question is whether this support skeleton is compatible with rho_F(A_F), the J_F-opposite action, bimodule stability, and the first-order condition.  Current project data do not certify the complete rho_F/J_F/gamma_F/D_F package, so the first-order expression cannot be promoted beyond a firewall audit.",
		Final:          "Final verdict: CONDITIONAL_SUPPORT_SYMBOLIC_D_F_HAS_CORRECT_CHIRAL_SUPPORT_FORM and CONDITIONAL_SUPPORT_FIRST_ORDER_PROOF_REQUIRES_COMPLETE_RHO_F_J_F_D_F_DATA, but FAILED_ROUTE_NO_COMPLETE_RHO_F_ACTION_LEDGER_CERTIFIED, FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_CERTIFIED, FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF, FAILED_ROUTE_SYMBOLIC_D_F_SUPPORT_MATRIX_NOT_NATIVE_FINITE_TRIPLE, and FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_COMPATIBILITY_NOT_R3.",
	}
	return a, nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2KernelCompatibilityStage: true, R3: false, R4: false, AlphaNative: false}
}

func buildSymbolicDF() SymbolicDF {
	return SymbolicDF{
		Expression:          "D_F^sym=[[0,Y_supp^dagger],[Y_supp,0]], Y_supp=y_+3Y_+3+y_-3Y_-3+y_-1Y_-1, y_+1=0",
		Domain:              "H_R^min=(e_+ tensor P_3) plus (e_- tensor P_3) plus (e_- tensor P_1)",
		Codomain:            "H_L=(h_+ plus h_-) tensor (P_1 plus P_3)",
		Space:               "H_L plus H_R^min",
		LeftRank:            HLRank,
		RightRank:           HRMinRank,
		TotalRank:           ChiralTotalDim,
		YRank:               7,
		DFRank:              14,
		KernelDim:           1,
		SelfAdjointByBlock:  true,
		ChiralityOddByBlock: true,
		SupportOnly:         true,
		NativeFiniteTriple:  false,
		ExplicitDFOperator:  false,
		NumericalDFMatrix:   false,
		Supports:            []string{SupportCorrectChiralSupportForm, SupportPunctureKernelPairInherited, SupportR2CompatibilityStage},
		Failures:            []string{FailureSymbolicDFNotNativeFiniteTriple, FailureNoExplicitDFOperator, FailureNoFullFirstOrderConditionProof, FailureSymbolicYNotYukawaMagnitude},
	}
}

func buildRepresentationData() RepresentationData {
	kernel := Cell{Name: "h_+ tensor P_1", Support: "left neutral kernel singleton", Chirality: "left", Socket: "h_+", LeptoColor: "P_1", Role: "forced kernel candidate whose algebraic stability requires rho_F and J_F data", Rank: 1, Kernel: true, InMinimalSupport: true, StableCandidate: true, PhysicalAssignment: false}
	return RepresentationData{
		Algebra:                  "A_F=C plus H plus M_3(C)",
		Representation:           "rho_F:A_F -> End(H_F)",
		RequiredPackage:          "(H_F,rho_F,J_F,gamma_F,D_F)",
		RhoFAvailable:            false,
		CompleteActionLedger:     false,
		CompletePackage:          false,
		GammaFAvailable:          false,
		JFAvailable:              false,
		DFOperatorAvailable:      false,
		CentralSupportLedger:     false,
		KernelStabilityCertified: false,
		Kernel:                   kernel,
		Supports:                 []string{SupportFirstOrderDataRequirements, SupportKernelStableCandidate},
		Failures:                 []string{FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTriplePackage, FailureNoJOppositeCompatibilityProof, FailureKernelStabilityNotCertified},
	}
}

func buildFirstOrderAudit() FirstOrderAudit {
	missing := []string{"complete rho_F action ledger", "J_F opposite/right action", "operator-valued D_F matrix", "bimodule/commutant decomposition", "first-order commutator calculation"}
	return FirstOrderAudit{
		Expression:            "[[D_F,rho_F(a)],J_F rho_F(b) J_F^{-1}]=0",
		CanFormCommutator:     false,
		CanFormOppositeAction: false,
		FirstOrderProven:      false,
		BimoduleStable:        false,
		JOppositeCompatible:   false,
		KernelStable:          false,
		MissingData:           missing,
		Supports:              []string{SupportFirstOrderDataRequirements},
		Failures:              []string{FailureNoCompleteRhoFActionLedger, FailureNoJOppositeCompatibilityProof, FailureNoFullFirstOrderConditionProof, FailureNoBimoduleCommutantProof, FailureKernelStabilityNotCertified},
	}
}

func buildChiralityAudit() ChiralityAudit {
	return ChiralityAudit{
		Expression:                "{D_F^sym,gamma_F}=0 by left/right block support, pending native gamma_F matrix",
		SupportLevelOddness:       true,
		NativeGammaCertified:      false,
		KernelChiralityConsistent: true,
		PunctureChiralitySet:      true,
		Supports:                  []string{SupportCorrectChiralSupportForm, SupportPunctureKernelPairInherited},
		Failures:                  []string{FailureChiralityOnlySupportLevel, FailureNoCompleteFiniteTriplePackage},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_kernel_compatibility symbolic finite-Dirac support matrix with first-order/J-opposite compatibility firewall; not R3/R4", Gate849Inherited: true, SymbolicDFInherited: true, CompatibilityAudited: true, NativeFiniteTriple: false, FirstOrderCertified: false, JOppositeCertified: false, KernelStable: false, PhysicalNeutrinoTheorem: false, MasslessTheorem: false, AlphaStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, NoRhoFActionLedger: true, NoCompletePackage: true, NoJOppositeProof: true, NoFirstOrderProof: true, NoBimoduleProof: true, NoExplicitDFOperator: true, SymbolicDFNotNative: true, KernelStabilityNotCertified: true, ChiralitySupportOnly: true, NoPhysicalNeutrino: true, NoRightNeutrino: true, NoMasslessness: true, NoYukawaMagnitudes: true, AlphaStillSealed: true, NoTraceMagnitudeReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, Verdict: StatusCompatibilityFirewall}
}

func Statuses() []string {
	return []string{
		StatusGate849Inherited, StatusSymbolicDFInherited, StatusRepresentationRequirements, StatusJOppositeRequirements, StatusFirstOrderRequirements, StatusBimoduleRequirements, StatusChiralityOddnessInherited, StatusKernelStabilityAudited, StatusNoObservedDataUsed, StatusOfficialLedgersFrozen, StatusCompatibilityFirewall,
		SupportCorrectChiralSupportForm, SupportKernelStableCandidate, SupportPunctureKernelPairInherited, SupportFirstOrderDataRequirements, SupportR2CompatibilityStage,
		FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTriplePackage, FailureNoJOppositeCompatibilityProof, FailureNoFullFirstOrderConditionProof, FailureNoBimoduleCommutantProof, FailureNoExplicitDFOperator, FailureSymbolicDFNotNativeFiniteTriple, FailureKernelStabilityNotCertified, FailureChiralityOnlySupportLevel, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem, FailureSymbolicYNotYukawaMagnitude, FailureKernelNotYukawaMagnitude, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return join("ledger", []string{"alpha_B=" + f64(l.AlphaB), "official_N_eff=" + f64(l.OfficialNEff), "official_C_Yukawa=" + f64(l.OfficialCYukawa), "official_C_Higgs=" + f64(l.OfficialCHiggs), "official_frozen=" + b(l.OfficialFrozen), "R2+++++_kernel_compatibility_stage=" + b(l.R2KernelCompatibilityStage), "R3=" + b(l.R3), "R4=" + b(l.R4), "alpha_native=" + b(l.AlphaNative)})
}

func FormatCell(c Cell) string {
	return join("cell", []string{c.Name, "support=" + c.Support, "chirality=" + c.Chirality, "socket=" + c.Socket, "lepto_color=" + c.LeptoColor, "role=" + c.Role, "rank=" + i(c.Rank), "absent=" + b(c.Absent), "kernel=" + b(c.Kernel), "minimal_support=" + b(c.InMinimalSupport), "stable_candidate=" + b(c.StableCandidate), "physical_assignment=" + b(c.PhysicalAssignment)})
}

func FormatSymbolicDF(d SymbolicDF) string {
	return join("symbolic_D_F", []string{d.Expression, "domain=" + d.Domain, "codomain=" + d.Codomain, "space=" + d.Space, "left_rank=" + i(d.LeftRank), "right_rank=" + i(d.RightRank), "total_rank=" + i(d.TotalRank), "Y_rank=" + i(d.YRank), "D_F_rank=" + i(d.DFRank), "kernel_dim=" + i(d.KernelDim), "self_adjoint_by_block=" + b(d.SelfAdjointByBlock), "chirality_odd_by_block=" + b(d.ChiralityOddByBlock), "support_only=" + b(d.SupportOnly), "native_finite_triple=" + b(d.NativeFiniteTriple), "explicit_D_F_operator=" + b(d.ExplicitDFOperator), "numerical_D_F=" + b(d.NumericalDFMatrix), "supports=" + strings.Join(d.Supports, ","), "failures=" + strings.Join(d.Failures, ",")})
}

func FormatRepresentation(r RepresentationData) string {
	return join("representation", []string{"algebra=" + r.Algebra, "representation=" + r.Representation, "required_package=" + r.RequiredPackage, "rho_F_available=" + b(r.RhoFAvailable), "complete_action_ledger=" + b(r.CompleteActionLedger), "complete_package=" + b(r.CompletePackage), "gamma_F_available=" + b(r.GammaFAvailable), "J_F_available=" + b(r.JFAvailable), "D_F_operator_available=" + b(r.DFOperatorAvailable), "central_support_ledger=" + b(r.CentralSupportLedger), "kernel_stability_certified=" + b(r.KernelStabilityCertified), "kernel=" + FormatCell(r.Kernel), "supports=" + strings.Join(r.Supports, ","), "failures=" + strings.Join(r.Failures, ",")})
}

func FormatFirstOrder(f FirstOrderAudit) string {
	return join("first_order", []string{f.Expression, "can_form_commutator=" + b(f.CanFormCommutator), "can_form_opposite=" + b(f.CanFormOppositeAction), "first_order_proven=" + b(f.FirstOrderProven), "bimodule_stable=" + b(f.BimoduleStable), "J_opposite_compatible=" + b(f.JOppositeCompatible), "kernel_stable=" + b(f.KernelStable), "missing=" + strings.Join(f.MissingData, ","), "supports=" + strings.Join(f.Supports, ","), "failures=" + strings.Join(f.Failures, ",")})
}

func FormatChirality(c ChiralityAudit) string {
	return join("chirality", []string{c.Expression, "support_level_oddness=" + b(c.SupportLevelOddness), "native_gamma_certified=" + b(c.NativeGammaCertified), "kernel_chirality_consistent=" + b(c.KernelChiralityConsistent), "puncture_chirality_set=" + b(c.PunctureChiralitySet), "supports=" + strings.Join(c.Supports, ","), "failures=" + strings.Join(c.Failures, ",")})
}

func FormatImpact(im Impact) string {
	return join("impact", []string{im.Classification, "gate849_inherited=" + b(im.Gate849Inherited), "symbolic_D_F_inherited=" + b(im.SymbolicDFInherited), "compatibility_audited=" + b(im.CompatibilityAudited), "native_finite_triple=" + b(im.NativeFiniteTriple), "first_order=" + b(im.FirstOrderCertified), "J_opposite=" + b(im.JOppositeCertified), "kernel_stable=" + b(im.KernelStable), "physical_neutrino=" + b(im.PhysicalNeutrinoTheorem), "massless_theorem=" + b(im.MasslessTheorem), "alpha_sealed=" + b(im.AlphaStillSealed), "magnitudes_missing=" + b(im.MagnitudesStillMissing), "update_N_eff=" + b(im.CanUpdateNEff), "update_C_Yukawa=" + b(im.CanUpdateCYukawa), "update_C_Higgs=" + b(im.CanUpdateCHiggs), "R3=" + b(im.CanPromoteToR3), "R4=" + b(im.CanPromoteToR4)})
}

func FormatFirewalls(f Firewalls) string {
	return join("firewalls", []string{"enforced=" + b(f.Enforced), "no_rho_F_action_ledger=" + b(f.NoRhoFActionLedger), "no_complete_package=" + b(f.NoCompletePackage), "no_J_opposite_proof=" + b(f.NoJOppositeProof), "no_first_order_proof=" + b(f.NoFirstOrderProof), "no_bimodule_proof=" + b(f.NoBimoduleProof), "no_explicit_D_F_operator=" + b(f.NoExplicitDFOperator), "symbolic_D_F_not_native=" + b(f.SymbolicDFNotNative), "kernel_stability_not_certified=" + b(f.KernelStabilityNotCertified), "chirality_support_only=" + b(f.ChiralitySupportOnly), "no_physical_neutrino=" + b(f.NoPhysicalNeutrino), "no_right_neutrino=" + b(f.NoRightNeutrino), "no_masslessness=" + b(f.NoMasslessness), "no_yukawa_magnitudes=" + b(f.NoYukawaMagnitudes), "alpha_sealed=" + b(f.AlphaStillSealed), "no_trace_magnitude_readout=" + b(f.NoTraceMagnitudeReadout), "no_official_N_eff_update=" + b(f.NoOfficialNEffUpdate), "no_C_Yukawa_C_Higgs_update=" + b(f.NoCYukawaCHiggsUpdate), "not_R3=" + b(f.NotR3), "not_R4=" + b(f.NotR4), "verdict=" + f.Verdict})
}

func containsAll(hay []string, needles []string) bool {
	m := map[string]bool{}
	for _, s := range hay {
		m[s] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func join(name string, parts []string) string { return name + "{" + strings.Join(parts, "; ") + "}" }
func f64(x float64) string                    { return strconv.FormatFloat(x, 'g', 17, 64) }
func i(x int) string                          { return strconv.Itoa(x) }
func b(x bool) string                         { return strconv.FormatBool(x) }
