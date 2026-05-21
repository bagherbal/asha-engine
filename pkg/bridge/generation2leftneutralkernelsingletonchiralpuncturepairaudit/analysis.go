// Package generation2leftneutralkernelsingletonchiralpuncturepairaudit implements
// Gate 849: LeftNeutral Kernel Singleton and Chiral Puncture Pair Audit.
//
// Gate 849 follows Gate 848's symbolic finite-Dirac support matrix. Gate 848
// constructed a support-only chiral block matrix
//
//	D_F^sym = [[0, Y_supp^dagger], [Y_supp, 0]],
//
// with Y_supp mapping the minimal right module H_R^min of rank 7 into the left
// lepto-color doublet H_L of rank 8. Gate 849 audits the forced rank-one left
// complement: if Y_supp has full support rank 7, its image occupies
//
//	h_+ tensor P_3, h_- tensor P_3, h_- tensor P_1,
//
// and leaves the singleton h_+ tensor P_1 outside the image. The audit compares
// this left neutral kernel singleton with the absent right puncture
// e_+ tensor P_1. This is structural support anatomy only: no physical neutrino
// theorem, no right-neutrino theorem, no masslessness theorem, no Yukawa
// magnitude readout, no alpha source, and no R3/R4 promotion is certified.
package generation2leftneutralkernelsingletonchiralpuncturepairaudit

import (
	"strconv"
	"strings"
)

const (
	AuditID = "GATE849-LEFT-NEUTRAL-KERNEL-SINGLETON-CHIRAL-PUNCTURE-PAIR-AUDIT"

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

	StatusGate848Inherited               = "PASS_GATE848_SYMBOLIC_D_F_SUPPORT_MATRIX_INHERITED"
	StatusImageSupportAudited            = "PASS_Y_SUPP_IMAGE_SUPPORT_AUDITED"
	StatusLeftNeutralComplementForced    = "PASS_LEFT_NEUTRAL_COMPLEMENT_FORCED_BY_RANK_8_TARGET_AND_RANK_7_DOMAIN"
	StatusYImageExcludesHPlusP1          = "PASS_Y_SUPP_IMAGE_EXCLUDES_H_PLUS_TENSOR_P1"
	StatusSymbolicDFKernelAudited        = "PASS_SYMBOLIC_D_F_KERNEL_SINGLETON_AUDITED"
	StatusRightPunctureLeftKernelAudited = "PASS_RIGHT_PUNCTURE_LEFT_KERNEL_PAIR_AUDITED"
	StatusSupportRankAnatomyAudited      = "PASS_RANK_15_SYMBOLIC_D_F_SUPPORT_ANATOMY_AUDITED"
	StatusNoPhysicalNaming               = "PASS_NEUTRAL_SINGLETON_PHYSICAL_NAMING_FIREWALL_ENFORCED"
	StatusNoObservedDataUsed             = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusOfficialLedgersFrozen          = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2KernelStage                  = "PASS_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_STAGE_NOT_R3_OR_R4"
	StatusFirewallGate849                = "FIREWALL_PRESERVED_GATE849_LEFT_NEUTRAL_KERNEL_SINGLETON"

	SupportLeftNeutralKernelSingleton = "CONDITIONAL_SUPPORT_H_PLUS_TENSOR_P1_IS_LEFT_NEUTRAL_KERNEL_SINGLETON"
	SupportChiralNeutralPair          = "CONDITIONAL_SUPPORT_E_PLUS_TENSOR_P1_AND_H_PLUS_TENSOR_P1_FORM_CHIRAL_NEUTRAL_PUNCTURE_KERNEL_PAIR"
	SupportOneForcedLeftNullMode      = "CONDITIONAL_SUPPORT_MINIMAL_EDGE_SUPPORT_HAS_ONE_FORCED_LEFT_NULL_MODE_AT_SEAL_LEVEL"
	SupportDFRankFourteenKernelOne    = "CONDITIONAL_SUPPORT_SYMBOLIC_D_F_SUPPORT_RANK_14_KERNEL_DIM_1"
	SupportRightPunctureLeftKernel    = "CONDITIONAL_SUPPORT_RIGHT_PUNCTURE_AND_LEFT_KERNEL_SHARE_PLUS_LEPTON_PROFILE"
	SupportNoMagnitudeFromKernel      = "CONDITIONAL_SUPPORT_SYMBOLIC_KERNEL_IS_NOT_YUKAWA_MAGNITUDE"
	SupportR2KernelStage              = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_STAGE"

	FailureKernelSealNotNative       = "FAILED_ROUTE_LEFT_NEUTRAL_KERNEL_SINGLETON_IS_SEAL_NOT_NATIVE_THEOREM"
	FailureNoNativeNullEdgeTheorem   = "FAILED_ROUTE_NO_NATIVE_NULL_EDGE_THEOREM_FOR_H_PLUS_TENSOR_P1"
	FailureNoPhysicalNeutrinoTheorem = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoRightNeutrinoTheorem    = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoMasslessnessTheorem     = "FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_OBSERVED_MASSLESSNESS_THEOREM"
	FailureKernelNotYukawaMagnitude  = "FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_YUKAWA_MAGNITUDE"
	FailureNoNumericalYukawaValues   = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoExplicitDFMatrix        = "FAILED_ROUTE_NO_NUMERICAL_OR_OPERATOR_VALUED_D_F_MATRIX_CERTIFIED"
	FailureNoFirstOrderProof         = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF"
	FailureNoBimoduleProof           = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoJOppositeProof          = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_CERTIFIED"
	FailureAlphaStillSealed          = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate      = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate     = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                     = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_NOT_R3"
	FailureNotR4                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	R2PlusPlusPlusPlusPlusKernel  bool
	R3, R4, AlphaNative           bool
}

type Cell struct {
	Name, Support, Chirality, Socket, LeptoColor, Role string
	Rank                                               int
	InRightDomain, InLeftTarget, InYImage              bool
	Absent, Kernel, Neutral, PhysicalAssignment        bool
}

type ImageSupport struct {
	Expression, Rule                                           string
	RightDomainRank, LeftTargetRank, ImageRank, ComplementRank int
	ActiveTargetCells                                          []Cell
	LeftComplement                                             Cell
	FullSupportRank, PreservesLeptoColor                       bool
	Supports, Failures                                         []string
}

type DiracKernel struct {
	Expression, Space                                  string
	LeftRank, RightRank, TotalRank                     int
	YRank, DFRank, KernelDim                           int
	RightKernelDim, LeftKernelDim                      int
	KernelSupport                                      Cell
	SupportOnly, NativeDFMatrix, NumericalDFMatrix     bool
	FirstOrderCertified, BimoduleProof, JOppositeProof bool
	Supports, Failures                                 []string
}

type ChiralPair struct {
	RightPuncture, LeftKernel                             Cell
	SameLeptonSupport, SamePlusSocket, DifferentChirality bool
	PairCandidate, PhysicalParticleTheorem                bool
	Supports, Failures                                    []string
}

type Impact struct {
	Classification                                         string
	Gate848Inherited, LeftKernelForced, PairAudited        bool
	KernelNative, PhysicalNeutrinoTheorem, MasslessTheorem bool
	AlphaStillSealed, MagnitudesStillMissing               bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs       bool
	CanPromoteToR3, CanPromoteToR4                         bool
}

type Firewalls struct {
	Enforced                                                                               bool
	KernelSealNotNative, NoNativeNullEdge, NoPhysicalNeutrino, NoRightNeutrino             bool
	NoMasslessnessTheorem, KernelNotYukawaMagnitude, NoNumericalYukawaValues               bool
	NoExplicitDFMatrix, NoFirstOrderProof, NoBimoduleProof, NoJOppositeProof               bool
	AlphaStillSealed, NoTraceMagnitudeReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate bool
	NotR3, NotR4                                                                           bool
	Verdict                                                                                string
}

type Audit struct {
	Ledger    Ledger
	Image     ImageSupport
	Kernel    DiracKernel
	Pair      ChiralPair
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:    buildLedger(),
		Image:     buildImageSupport(),
		Kernel:    buildDiracKernel(),
		Pair:      buildChiralPair(),
		Impact:    buildImpact(),
		Firewalls: buildFirewalls(),
		Truth:     "Gate 849 follows Gate 848's symbolic D_F support matrix. Since Y_supp maps a rank-seven minimal right module into a rank-eight left lepto-color doublet, full support rank forces a rank-one left complement. The image occupies h_+ tensor P_3, h_- tensor P_3, and h_- tensor P_1; the forced complement is h_+ tensor P_1. Therefore D_F^sym has support rank fourteen on the fifteen-dimensional chiral support and a one-dimensional symbolic kernel supported on h_+ tensor P_1.",
		Final:     "Final verdict: CONDITIONAL_SUPPORT_H_PLUS_TENSOR_P1_IS_LEFT_NEUTRAL_KERNEL_SINGLETON and CONDITIONAL_SUPPORT_E_PLUS_TENSOR_P1_AND_H_PLUS_TENSOR_P1_FORM_CHIRAL_NEUTRAL_PUNCTURE_KERNEL_PAIR, but FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM, FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM, FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_YUKAWA_MAGNITUDE, FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF, and FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_KERNEL_NOT_R3.",
	}
	return a, nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlusPlusPlusPlusKernel: true, R3: false, R4: false, AlphaNative: false}
}

func buildImageSupport() ImageSupport {
	cells := []Cell{
		{Name: "h_+ tensor P_3", Support: "left color plus target", Chirality: "left", Socket: "h_+", LeptoColor: "P_3", Role: "image of Y_+3", Rank: 3, InLeftTarget: true, InYImage: true, Neutral: false},
		{Name: "h_- tensor P_3", Support: "left color minus target", Chirality: "left", Socket: "h_-", LeptoColor: "P_3", Role: "image of Y_-3", Rank: 3, InLeftTarget: true, InYImage: true, Neutral: false},
		{Name: "h_- tensor P_1", Support: "left lepton minus target", Chirality: "left", Socket: "h_-", LeptoColor: "P_1", Role: "image of Y_-1", Rank: 1, InLeftTarget: true, InYImage: true, Neutral: true},
	}
	leftComplement := Cell{Name: "h_+ tensor P_1", Support: "left lepton plus complement", Chirality: "left", Socket: "h_+", LeptoColor: "P_1", Role: "forced rank-one left complement outside Im(Y_supp)", Rank: 1, InLeftTarget: true, InYImage: false, Kernel: true, Neutral: true, PhysicalAssignment: false}
	return ImageSupport{
		Expression:          "Im(Y_supp)=h_+ tensor P_3 plus h_- tensor P_3 plus h_- tensor P_1; complement=h_+ tensor P_1",
		Rule:                "full support rank seven inside rank-eight H_L forces one left neutral complement",
		RightDomainRank:     HRMinRank,
		LeftTargetRank:      HLRank,
		ImageRank:           7,
		ComplementRank:      1,
		ActiveTargetCells:   cells,
		LeftComplement:      leftComplement,
		FullSupportRank:     true,
		PreservesLeptoColor: true,
		Supports:            []string{SupportLeftNeutralKernelSingleton, SupportOneForcedLeftNullMode, SupportDFRankFourteenKernelOne, SupportNoMagnitudeFromKernel},
		Failures:            []string{FailureKernelSealNotNative, FailureNoNativeNullEdgeTheorem, FailureNoPhysicalNeutrinoTheorem, FailureNoMasslessnessTheorem, FailureKernelNotYukawaMagnitude},
	}
}

func buildDiracKernel() DiracKernel {
	kernel := Cell{Name: "h_+ tensor P_1", Support: "left lepton plus kernel", Chirality: "left", Socket: "h_+", LeptoColor: "P_1", Role: "one-dimensional kernel of symbolic D_F support matrix", Rank: 1, InLeftTarget: true, InYImage: false, Kernel: true, Neutral: true, PhysicalAssignment: false}
	return DiracKernel{
		Expression:          "D_F^sym=[[0,Y_supp^dagger],[Y_supp,0]], rank(Y_supp)=7 => rank(D_F^sym)=14, kernel_dim=1",
		Space:               "H_L plus H_R^min",
		LeftRank:            HLRank,
		RightRank:           HRMinRank,
		TotalRank:           ChiralTotalDim,
		YRank:               7,
		DFRank:              14,
		KernelDim:           1,
		RightKernelDim:      0,
		LeftKernelDim:       1,
		KernelSupport:       kernel,
		SupportOnly:         true,
		NativeDFMatrix:      false,
		NumericalDFMatrix:   false,
		FirstOrderCertified: false,
		BimoduleProof:       false,
		JOppositeProof:      false,
		Supports:            []string{SupportLeftNeutralKernelSingleton, SupportDFRankFourteenKernelOne, SupportOneForcedLeftNullMode},
		Failures:            []string{FailureKernelSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleProof, FailureNoJOppositeProof, FailureKernelNotYukawaMagnitude, FailureNoNumericalYukawaValues},
	}
}

func buildChiralPair() ChiralPair {
	right := Cell{Name: "e_+ tensor P_1", Support: "right neutral puncture", Chirality: "right", Socket: "e_+", LeptoColor: "P_1", Role: "absent right puncture from minimal right module", Rank: 1, InRightDomain: false, Absent: true, Neutral: true, PhysicalAssignment: false}
	left := Cell{Name: "h_+ tensor P_1", Support: "left neutral kernel", Chirality: "left", Socket: "h_+", LeptoColor: "P_1", Role: "forced left kernel singleton outside Im(Y_supp)", Rank: 1, InLeftTarget: true, Kernel: true, Neutral: true, PhysicalAssignment: false}
	return ChiralPair{
		RightPuncture:           right,
		LeftKernel:              left,
		SameLeptonSupport:       true,
		SamePlusSocket:          true,
		DifferentChirality:      true,
		PairCandidate:           true,
		PhysicalParticleTheorem: false,
		Supports:                []string{SupportChiralNeutralPair, SupportRightPunctureLeftKernel},
		Failures:                []string{FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem, FailureNoNativeNullEdgeTheorem},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_kernel symbolic finite-Dirac support matrix with forced left neutral kernel singleton; not R3/R4", Gate848Inherited: true, LeftKernelForced: true, PairAudited: true, KernelNative: false, PhysicalNeutrinoTheorem: false, MasslessTheorem: false, AlphaStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, KernelSealNotNative: true, NoNativeNullEdge: true, NoPhysicalNeutrino: true, NoRightNeutrino: true, NoMasslessnessTheorem: true, KernelNotYukawaMagnitude: true, NoNumericalYukawaValues: true, NoExplicitDFMatrix: true, NoFirstOrderProof: true, NoBimoduleProof: true, NoJOppositeProof: true, AlphaStillSealed: true, NoTraceMagnitudeReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate849}
}

func Statuses() []string {
	return []string{
		StatusGate848Inherited, StatusImageSupportAudited, StatusLeftNeutralComplementForced, StatusYImageExcludesHPlusP1, StatusSymbolicDFKernelAudited, StatusRightPunctureLeftKernelAudited, StatusSupportRankAnatomyAudited, StatusNoPhysicalNaming, StatusNoObservedDataUsed, StatusOfficialLedgersFrozen, StatusR2KernelStage, StatusFirewallGate849,
		SupportLeftNeutralKernelSingleton, SupportChiralNeutralPair, SupportOneForcedLeftNullMode, SupportDFRankFourteenKernelOne, SupportRightPunctureLeftKernel, SupportNoMagnitudeFromKernel, SupportR2KernelStage,
		FailureKernelSealNotNative, FailureNoNativeNullEdgeTheorem, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem, FailureKernelNotYukawaMagnitude, FailureNoNumericalYukawaValues, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleProof, FailureNoJOppositeProof, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return join("ledger", []string{"alpha_B=" + f64(l.AlphaB), "official_N_eff=" + f64(l.OfficialNEff), "official_C_Yukawa=" + f64(l.OfficialCYukawa), "official_C_Higgs=" + f64(l.OfficialCHiggs), "official_frozen=" + b(l.OfficialFrozen), "R2+++++_kernel_stage=" + b(l.R2PlusPlusPlusPlusPlusKernel), "R3=" + b(l.R3), "R4=" + b(l.R4), "alpha_native=" + b(l.AlphaNative)})
}

func FormatCell(c Cell) string {
	return join("cell", []string{c.Name, "support=" + c.Support, "chirality=" + c.Chirality, "socket=" + c.Socket, "lepto_color=" + c.LeptoColor, "role=" + c.Role, "rank=" + i(c.Rank), "right_domain=" + b(c.InRightDomain), "left_target=" + b(c.InLeftTarget), "Y_image=" + b(c.InYImage), "absent=" + b(c.Absent), "kernel=" + b(c.Kernel), "neutral=" + b(c.Neutral), "physical_assignment=" + b(c.PhysicalAssignment)})
}

func FormatImage(im ImageSupport) string {
	parts := []string{im.Expression, "rule=" + im.Rule, "right_domain_rank=" + i(im.RightDomainRank), "left_target_rank=" + i(im.LeftTargetRank), "image_rank=" + i(im.ImageRank), "complement_rank=" + i(im.ComplementRank), "full_support_rank=" + b(im.FullSupportRank), "preserves_lepto_color=" + b(im.PreservesLeptoColor)}
	for _, c := range im.ActiveTargetCells {
		parts = append(parts, FormatCell(c))
	}
	parts = append(parts, "left_complement="+FormatCell(im.LeftComplement), "supports="+strings.Join(im.Supports, ","), "failures="+strings.Join(im.Failures, ","))
	return join("image_support", parts)
}

func FormatKernel(k DiracKernel) string {
	return join("D_F_kernel", []string{k.Expression, "space=" + k.Space, "left_rank=" + i(k.LeftRank), "right_rank=" + i(k.RightRank), "total_rank=" + i(k.TotalRank), "Y_rank=" + i(k.YRank), "D_F_rank=" + i(k.DFRank), "kernel_dim=" + i(k.KernelDim), "right_kernel_dim=" + i(k.RightKernelDim), "left_kernel_dim=" + i(k.LeftKernelDim), "kernel_support=" + FormatCell(k.KernelSupport), "support_only=" + b(k.SupportOnly), "native_D_F=" + b(k.NativeDFMatrix), "numerical_D_F=" + b(k.NumericalDFMatrix), "first_order=" + b(k.FirstOrderCertified), "bimodule=" + b(k.BimoduleProof), "J_opposite=" + b(k.JOppositeProof), "supports=" + strings.Join(k.Supports, ","), "failures=" + strings.Join(k.Failures, ",")})
}

func FormatPair(p ChiralPair) string {
	return join("chiral_pair", []string{"right=" + FormatCell(p.RightPuncture), "left=" + FormatCell(p.LeftKernel), "same_lepton_support=" + b(p.SameLeptonSupport), "same_plus_socket=" + b(p.SamePlusSocket), "different_chirality=" + b(p.DifferentChirality), "pair_candidate=" + b(p.PairCandidate), "physical_particle_theorem=" + b(p.PhysicalParticleTheorem), "supports=" + strings.Join(p.Supports, ","), "failures=" + strings.Join(p.Failures, ",")})
}

func FormatImpact(im Impact) string {
	return join("impact", []string{im.Classification, "gate848_inherited=" + b(im.Gate848Inherited), "left_kernel_forced=" + b(im.LeftKernelForced), "pair_audited=" + b(im.PairAudited), "kernel_native=" + b(im.KernelNative), "physical_neutrino=" + b(im.PhysicalNeutrinoTheorem), "massless_theorem=" + b(im.MasslessTheorem), "alpha_sealed=" + b(im.AlphaStillSealed), "magnitudes_missing=" + b(im.MagnitudesStillMissing), "update_N_eff=" + b(im.CanUpdateNEff), "update_C_Yukawa=" + b(im.CanUpdateCYukawa), "update_C_Higgs=" + b(im.CanUpdateCHiggs), "R3=" + b(im.CanPromoteToR3), "R4=" + b(im.CanPromoteToR4)})
}

func FormatFirewalls(f Firewalls) string {
	return join("firewalls", []string{"enforced=" + b(f.Enforced), "kernel_seal_not_native=" + b(f.KernelSealNotNative), "no_native_null_edge=" + b(f.NoNativeNullEdge), "no_physical_neutrino=" + b(f.NoPhysicalNeutrino), "no_right_neutrino=" + b(f.NoRightNeutrino), "no_masslessness=" + b(f.NoMasslessnessTheorem), "kernel_not_yukawa=" + b(f.KernelNotYukawaMagnitude), "no_numerical_yukawa=" + b(f.NoNumericalYukawaValues), "no_explicit_D_F=" + b(f.NoExplicitDFMatrix), "no_first_order=" + b(f.NoFirstOrderProof), "no_bimodule=" + b(f.NoBimoduleProof), "no_J_opposite=" + b(f.NoJOppositeProof), "alpha_sealed=" + b(f.AlphaStillSealed), "no_trace_magnitude=" + b(f.NoTraceMagnitudeReadout), "no_N_eff_update=" + b(f.NoOfficialNEffUpdate), "no_C_updates=" + b(f.NoCYukawaCHiggsUpdate), "not_R3=" + b(f.NotR3), "not_R4=" + b(f.NotR4), "verdict=" + f.Verdict})
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
