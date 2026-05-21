// Package generation2minimalnulledgeorientationprincipleaudit implements
// Gate 894: MinimalNullEdge Orientation Principle Audit.
//
// Gate 894 follows Gate 893's weak-socket selector obstruction. It audits
// whether the Higgs / weak socket frame can be selected by a minimal null-edge
// principle: choose the weak line h_+ so that the absent right-neutral puncture
// e_+ tensor P_1 maps to a left kernel h_+ tensor P_1, while the remaining
// three active edges preserve lepto-color support and minimize the active edge
// domain. The gate source-types the principle but does not certify it natively.
package generation2minimalnulledgeorientationprincipleaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE894-MINIMAL-NULL-EDGE-ORIENTATION-PRINCIPLE-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	FullUnbrokenAlgebra    = "A_F=C plus H plus M_3(C)"
	PostOrientationAlgebra = "A_F^orient=C_R plus C_H plus M_3(C)"
	WeakFrame              = "C_L^2=h_+ plus h_-"

	RightPuncture = "e_+ tensor P_1"
	LeftKernel    = "h_+ tensor P_1"
	MissingEdge   = "Y_+1:e_+ tensor P_1 -> h_+ tensor P_1"

	EdgePiPlus3  = "Y_+3:e_+ tensor P_3 -> h_+ tensor P_3"
	EdgePiMinus3 = "Y_-3:e_- tensor P_3 -> h_- tensor P_3"
	EdgePiMinus1 = "Y_-1:e_- tensor P_1 -> h_- tensor P_1"

	PiPlus3  = "Pi_+3=e_+ tensor P_3"
	PiMinus3 = "Pi_-3=e_- tensor P_3"
	PiMinus1 = "Pi_-1=e_- tensor P_1"

	HRMinRank  = 7
	HLeftRank  = 8
	DFRank     = 14
	KernelRank = 1

	Classification = "R3_DUALSEAL_MINIMAL_NULL_EDGE_ORIENTATION_CANDIDATE_NOT_NATIVE"
	ShortStatus    = "R3_CANDIDATE_MINIMAL_NULL_EDGE_ORIENTATION_SOURCE_TYPED_OBSTRUCTION"
	NextFrontier   = "NONCIRCULAR_WEAK_SOCKET_SELECTOR_FUNCTIONAL_OR_VARIATIONAL_MINIMALITY_THEOREM"

	StatusGate893Inherited            = "PASS_GATE893_WEAK_SOCKET_SELECTOR_OBSTRUCTION_INHERITED"
	StatusNullEdgeMinimizationAudited = "PASS_NULL_EDGE_MINIMIZATION_AUDITED"
	StatusImageKernelReconstructed    = "PASS_IMAGE_KERNEL_RECONSTRUCTION_AUDITED"
	StatusNonCircularityAudited       = "PASS_MINIMAL_NULL_EDGE_NONCIRCULARITY_AUDITED"
	StatusOfficialFreeze              = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict             = "FIREWALL_PRESERVED_GATE894_MINIMAL_NULL_EDGE_NOT_NATIVE"

	SupportMinimalNullEdgeSelectsHPlusCandidate    = "CONDITIONAL_SUPPORT_MINIMAL_NULL_EDGE_PRINCIPLE_SELECTS_H_PLUS_AS_KERNEL_LINE_CANDIDATE"
	SupportPunctureKernelPairSourceTypesWeakSocket = "CONDITIONAL_SUPPORT_RIGHT_PUNCTURE_LEFT_KERNEL_PAIR_SOURCE_TYPES_WEAK_SOCKET_ORIENTATION"
	SupportHPlusKernelEqualsLeftQuotientCandidate  = "CONDITIONAL_SUPPORT_H_PLUS_TENSOR_P1_EQUALS_H_L_OVER_IMAGE_Y_CANDIDATE"
	SupportImageYHasThreeActiveTargets             = "CONDITIONAL_SUPPORT_IMAGE_Y_HAS_THREE_ACTIVE_LEPTO_COLOR_PRESERVING_TARGETS"
	SupportMinimalRankSevenEdgeDomain              = "CONDITIONAL_SUPPORT_MINIMAL_RANK_SEVEN_EDGE_DOMAIN_RECONSTRUCTS_LEFT_KERNEL_SINGLETON"
	SupportHiggsOrientationSealWeakenedNotRemoved  = "CONDITIONAL_SUPPORT_HIGGS_ORIENTATION_SEAL_WEAKENED_BUT_NOT_REMOVED"
	SupportMissingTheoremIsSelectorFunctional      = "CONDITIONAL_SUPPORT_MISSING_THEOREM_IS_VARIATIONAL_MINIMALITY_OR_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	SupportOperatorNEffReproduced                  = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"

	FailureNotNativeR3                          = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                     = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeHiggsOrientationSource       = "FAILED_ROUTE_NO_NATIVE_HIGGS_ORIENTATION_SOURCE_CERTIFIED"
	FailureNoNativeMinimalNullEdgeOrientation   = "FAILED_ROUTE_NO_NATIVE_MINIMAL_NULL_EDGE_ORIENTATION_PRINCIPLE"
	FailureNoNativeVariationalMinimalityTheorem = "FAILED_ROUTE_NO_NATIVE_VARIATIONAL_MINIMALITY_THEOREM"
	FailureKernelLineDependsOnEdgeSupportChoice = "FAILED_ROUTE_KERNEL_LINE_SELECTION_DEPENDS_ON_EDGE_SUPPORT_CHOICE"
	FailureNoNonCircularWeakSocketSelector      = "FAILED_ROUTE_NO_NONCIRCULAR_WEAK_SOCKET_SELECTOR_FUNCTIONAL_CERTIFIED"
	FailureDFPatternRestatesOrientation         = "FAILED_ROUTE_D_F_EDGE_PATTERN_RESTATES_ORIENTATION_WITHOUT_INDEPENDENT_SELECTOR"
	FailureFullHMixesWeakSockets                = "FAILED_ROUTE_FULL_H_ACTION_MIXES_H_PLUS_H_MINUS"
	FailureNoNativeDescentFullToOrient          = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailurePostOrientationNotFullAF             = "FAILED_ROUTE_POST_ORIENTATION_PROJECTORS_NOT_FULL_UNBROKEN_A_F_SECTORS"
	FailureNoNativeIncidenceFunctor             = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeR3SectorLedger               = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoPhysicalParticleAssignment         = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap               = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap               = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues             = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                 = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem              = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type NullEdgeMinimizationAudit struct {
	MissingEdge               string
	RightPuncture             string
	LeftKernel                string
	YPlus1Zero                bool
	AmbientRightRank          int
	ActiveRightRank           int
	RankReduction             int
	MinimalRankSevenCandidate bool
	SelectsHPlusCandidate     bool
	NativePrinciple           bool
	Supports, Failures        []string
}

type ImageKernelAudit struct {
	ActiveEdges               []string
	ImageTargets              []string
	HLeftRank                 int
	ImageRank                 int
	Kernel                    string
	KernelRank                int
	QuotientIsHPlusP1         bool
	ReconstructsKernel        bool
	SelectsFrameNonCircularly bool
	Supports, Failures        []string
}

type NonCircularityAudit struct {
	CanDefineHPlusFromKernelWithoutPriorFrame bool
	EdgeSupportAssumesFrame                   bool
	RequiresVariationalMinimality             bool
	MissingObject                             string
	NativeSelectorFunctional                  bool
	Supports, Failures                        []string
}

type FreezeAudit struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                           bool
	NotNativeR3                        bool
	AlphaStillSealed                   bool
	NoNativeHiggsOrientationSource     bool
	NoNativeMinimalNullEdgeOrientation bool
	NoNativeVariationalMinimality      bool
	NoNonCircularWeakSocketSelector    bool
	KernelLineDependsOnEdgeChoice      bool
	DFPatternRestatesOrientation       bool
	FullHMixesWeakSockets              bool
	NoNativeDescentFullToOrient        bool
	NoNativeIncidenceFunctor           bool
	NoNativeR3SectorLedger             bool
	NoPhysicalParticleAssignment       bool
	NoGenerationCarrier                bool
	NoFlavorOrientation                bool
	NoIndividualYukawas                bool
	NoOfficialLedgerUpdate             bool
	NoNativeYukawaOperator             bool
	NoR4NativeYukawaTheorem            bool
	Verdict                            string
}

type Audit struct {
	ID             string
	NullEdge       NullEdgeMinimizationAudit
	ImageKernel    ImageKernelAudit
	NonCircularity NonCircularityAudit
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	nullEdge := buildNullEdgeMinimizationAudit()
	if !nullEdge.YPlus1Zero || !nullEdge.MinimalRankSevenCandidate || !nullEdge.SelectsHPlusCandidate || nullEdge.NativePrinciple {
		return Audit{}, fmt.Errorf("null-edge minimization promoted incorrectly: %s", FormatNullEdgeMinimization(nullEdge))
	}

	imageKernel := buildImageKernelAudit()
	if !imageKernel.ReconstructsKernel || !imageKernel.QuotientIsHPlusP1 || imageKernel.SelectsFrameNonCircularly {
		return Audit{}, fmt.Errorf("image/kernel audit promoted incorrectly: %s", FormatImageKernel(imageKernel))
	}

	noncircular := buildNonCircularityAudit()
	if noncircular.CanDefineHPlusFromKernelWithoutPriorFrame || !noncircular.EdgeSupportAssumesFrame || !noncircular.RequiresVariationalMinimality || noncircular.NativeSelectorFunctional {
		return Audit{}, fmt.Errorf("noncircularity leak: %s", FormatNonCircularity(noncircular))
	}

	freeze := FreezeAudit{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreeze, SupportOperatorNEffReproduced},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}

	firewalls := Firewalls{
		Enforced: true, NotNativeR3: true, AlphaStillSealed: true,
		NoNativeHiggsOrientationSource: true, NoNativeMinimalNullEdgeOrientation: true,
		NoNativeVariationalMinimality: true, NoNonCircularWeakSocketSelector: true,
		KernelLineDependsOnEdgeChoice: true, DFPatternRestatesOrientation: true,
		FullHMixesWeakSockets: true, NoNativeDescentFullToOrient: true,
		NoNativeIncidenceFunctor: true, NoNativeR3SectorLedger: true,
		NoPhysicalParticleAssignment: true, NoGenerationCarrier: true,
		NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true,
		NoR4NativeYukawaTheorem: true, Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:             AuditID,
		NullEdge:       nullEdge,
		ImageKernel:    imageKernel,
		NonCircularity: noncircular,
		Freeze:         freeze,
		Firewalls:      firewalls,
		Truth:          "Gate 894 source-types the Higgs/weak socket frame by a minimal null-edge candidate: Y_+1=0 removes the right neutral puncture e_+ tensor P_1 and leaves h_+ tensor P_1 as the forced left kernel of the rank-seven lepto-color preserving edge support. This weakens the orientation seal, but the kernel line still depends on the chosen edge support and no noncircular variational/minimality theorem is certified.",
		Final:          "The branch remains R3_DUALSEAL_MINIMAL_NULL_EDGE_ORIENTATION_CANDIDATE_NOT_NATIVE. The missing object is a noncircular WeakSocketSelectorFunctional or native variational MinimalNullEdgeOrientationPrinciple selecting the rank-seven edge support and its unique left kernel without assuming h_+ / h_- beforehand.",
	}, nil
}

func buildNullEdgeMinimizationAudit() NullEdgeMinimizationAudit {
	return NullEdgeMinimizationAudit{
		MissingEdge: MissingEdge, RightPuncture: RightPuncture, LeftKernel: LeftKernel,
		YPlus1Zero: true, AmbientRightRank: 8, ActiveRightRank: HRMinRank, RankReduction: 1,
		MinimalRankSevenCandidate: true, SelectsHPlusCandidate: true, NativePrinciple: false,
		Supports: []string{SupportMinimalNullEdgeSelectsHPlusCandidate, SupportPunctureKernelPairSourceTypesWeakSocket, SupportMinimalRankSevenEdgeDomain},
		Failures: []string{FailureNoNativeMinimalNullEdgeOrientation, FailureNoNativeHiggsOrientationSource},
	}
}

func buildImageKernelAudit() ImageKernelAudit {
	return ImageKernelAudit{
		ActiveEdges:  []string{EdgePiPlus3, EdgePiMinus3, EdgePiMinus1},
		ImageTargets: []string{"h_+ tensor P_3", "h_- tensor P_3", "h_- tensor P_1"},
		HLeftRank:    HLeftRank, ImageRank: HRMinRank, Kernel: LeftKernel, KernelRank: KernelRank,
		QuotientIsHPlusP1: true, ReconstructsKernel: true, SelectsFrameNonCircularly: false,
		Supports: []string{SupportImageYHasThreeActiveTargets, SupportHPlusKernelEqualsLeftQuotientCandidate, SupportHiggsOrientationSealWeakenedNotRemoved},
		Failures: []string{FailureKernelLineDependsOnEdgeSupportChoice, FailureDFPatternRestatesOrientation, FailureNoNonCircularWeakSocketSelector},
	}
}

func buildNonCircularityAudit() NonCircularityAudit {
	return NonCircularityAudit{
		CanDefineHPlusFromKernelWithoutPriorFrame: false, EdgeSupportAssumesFrame: true, RequiresVariationalMinimality: true,
		MissingObject:            "WeakSocketSelectorFunctional or MinimalNullEdgeOrientationPrinciple",
		NativeSelectorFunctional: false,
		Supports:                 []string{SupportMissingTheoremIsSelectorFunctional, SupportHiggsOrientationSealWeakenedNotRemoved},
		Failures:                 []string{FailureNoNativeVariationalMinimalityTheorem, FailureNoNonCircularWeakSocketSelector, FailureDFPatternRestatesOrientation},
	}
}

func FormatNullEdgeMinimization(n NullEdgeMinimizationAudit) string {
	return fmt.Sprintf("null_edge_minimization(missing=%s right_puncture=%s left_kernel=%s y_plus1_zero=%t ambient_right_rank=%d active_right_rank=%d rank_reduction=%d minimal_rank7=%t selects_h_plus=%t native=%t supports=%s failures=%s)", n.MissingEdge, n.RightPuncture, n.LeftKernel, n.YPlus1Zero, n.AmbientRightRank, n.ActiveRightRank, n.RankReduction, n.MinimalRankSevenCandidate, n.SelectsHPlusCandidate, n.NativePrinciple, strings.Join(n.Supports, ","), strings.Join(n.Failures, ","))
}

func FormatImageKernel(i ImageKernelAudit) string {
	return fmt.Sprintf("image_kernel(active_edges=%s image_targets=%s H_L_rank=%d image_rank=%d kernel=%s kernel_rank=%d quotient_is_h_plus_p1=%t reconstructs=%t noncircular_select=%t supports=%s failures=%s)", strings.Join(i.ActiveEdges, ","), strings.Join(i.ImageTargets, ","), i.HLeftRank, i.ImageRank, i.Kernel, i.KernelRank, i.QuotientIsHPlusP1, i.ReconstructsKernel, i.SelectsFrameNonCircularly, strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func FormatNonCircularity(n NonCircularityAudit) string {
	return fmt.Sprintf("noncircularity(h_plus_from_kernel_without_prior_frame=%t edge_support_assumes_frame=%t requires_variational_minimality=%t missing=%s native=%t supports=%s failures=%s)", n.CanDefineHPlusFromKernelWithoutPriorFrame, n.EdgeSupportAssumesFrame, n.RequiresVariationalMinimality, n.MissingObject, n.NativeSelectorFunctional, strings.Join(n.Supports, ","), strings.Join(n.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t no_higgs_orientation=%t no_null_edge=%t no_variational=%t no_selector=%t kernel_depends=%t DF_restates=%t full_H_mixes=%t no_descent=%t no_incidence=%t no_r3_ledger=%t no_physical=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.NoNativeHiggsOrientationSource, f.NoNativeMinimalNullEdgeOrientation, f.NoNativeVariationalMinimality, f.NoNonCircularWeakSocketSelector, f.KernelLineDependsOnEdgeChoice, f.DFPatternRestatesOrientation, f.FullHMixesWeakSockets, f.NoNativeDescentFullToOrient, f.NoNativeIncidenceFunctor, f.NoNativeR3SectorLedger, f.NoPhysicalParticleAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate893Inherited, StatusNullEdgeMinimizationAudited, StatusImageKernelReconstructed, StatusNonCircularityAudited, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportMinimalNullEdgeSelectsHPlusCandidate, SupportPunctureKernelPairSourceTypesWeakSocket, SupportHPlusKernelEqualsLeftQuotientCandidate, SupportImageYHasThreeActiveTargets, SupportMinimalRankSevenEdgeDomain, SupportHiggsOrientationSealWeakenedNotRemoved, SupportMissingTheoremIsSelectorFunctional, SupportOperatorNEffReproduced,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeHiggsOrientationSource, FailureNoNativeMinimalNullEdgeOrientation, FailureNoNativeVariationalMinimalityTheorem, FailureKernelLineDependsOnEdgeSupportChoice, FailureNoNonCircularWeakSocketSelector, FailureDFPatternRestatesOrientation, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailurePostOrientationNotFullAF, FailureNoNativeIncidenceFunctor, FailureNoNativeR3SectorLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeHiggsOrientationSource, FailureNoNativeMinimalNullEdgeOrientation, FailureNoNativeVariationalMinimalityTheorem, FailureNoNonCircularWeakSocketSelector, FailureKernelLineDependsOnEdgeSupportChoice, FailureDFPatternRestatesOrientation, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeIncidenceFunctor, FailureNoNativeR3SectorLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.NoNativeHiggsOrientationSource &&
		f.NoNativeMinimalNullEdgeOrientation && f.NoNativeVariationalMinimality && f.NoNonCircularWeakSocketSelector &&
		f.KernelLineDependsOnEdgeChoice && f.DFPatternRestatesOrientation && f.FullHMixesWeakSockets &&
		f.NoNativeDescentFullToOrient && f.NoNativeIncidenceFunctor && f.NoNativeR3SectorLedger &&
		f.NoPhysicalParticleAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas &&
		f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4NativeYukawaTheorem && f.Verdict == StatusFirewallVerdict
}

func containsAll(haystack []string, needles []string) bool {
	seen := map[string]bool{}
	for _, h := range haystack {
		seen[h] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}

func near(a, b float64) bool { return math.Abs(a-b) < 5e-15 }
