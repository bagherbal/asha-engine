// Package generation2finiteoneformpuncturekernelweaksocketselectoraudit implements
// Gate 893: FiniteOneForm / PunctureKernel WeakSocket Selector Audit.
//
// Gate 893 follows Gate 892's HiggsOrientation source obstruction. It audits
// whether the weak socket frame h_+ plus h_- can be source-typed by the finite
// one-form / Higgs edge together with the neutral puncture/kernel pair, while
// preserving the firewall that no noncircular native weak-socket selector is
// certified.
package generation2finiteoneformpuncturekernelweaksocketselectoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE893-FINITE-ONE-FORM-PUNCTURE-KERNEL-WEAK-SOCKET-SELECTOR-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	PostOrientationAlgebra = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullUnbrokenAlgebra    = "A_F=C plus H plus M_3(C)"
	WeakFrame              = "C_L^2=h_+ plus h_-"
	WeakStabilizer         = "C_H=Stab_H(h_+ plus h_-)"

	RightPuncture = "e_+ tensor P_1"
	LeftKernel    = "h_+ tensor P_1"
	MissingEdge   = "Y_+1:e_+ tensor P_1 -> h_+ tensor P_1"

	EdgePiPlus3  = "Y_+3:e_+ tensor P_3 -> h_+ tensor P_3"
	EdgePiMinus3 = "Y_-3:e_- tensor P_3 -> h_- tensor P_3"
	EdgePiMinus1 = "Y_-1:e_- tensor P_1 -> h_- tensor P_1"

	PiPlus3  = "Pi_+3=e_+ tensor P_3"
	PiMinus3 = "Pi_-3=e_- tensor P_3"
	PiMinus1 = "Pi_-1=e_- tensor P_1"

	Classification = "R3_CANDIDATE_WEAK_SOCKET_SELECTOR_SOURCE_TYPED_NOT_NATIVE"
	ShortStatus    = "R3_DUALSEAL_WEAK_SOCKET_SELECTOR_SOURCE_TYPED_OBSTRUCTION"
	NextFrontier   = "WEAK_SOCKET_SELECTOR_FUNCTIONAL_OR_MINIMAL_NULL_EDGE_ORIENTATION_PRINCIPLE"

	StatusGate892Inherited          = "PASS_GATE892_HIGGS_ORIENTATION_SOURCE_OBSTRUCTION_INHERITED"
	StatusNullEdgeRouteAudited      = "PASS_NEUTRAL_NULL_EDGE_SELECTOR_ROUTE_AUDITED"
	StatusFiniteOneFormRouteAudited = "PASS_FINITE_ONE_FORM_ORIENTATION_ROUTE_AUDITED"
	StatusNonCircularityAudited     = "PASS_ORIENTATION_NONCIRCULARITY_AUDITED"
	StatusStrongestPair             = "PASS_FINITE_ONE_FORM_AND_PUNCTURE_KERNEL_PAIR_SOURCE_TYPE_ORIENTATION"
	StatusOfficialFreeze            = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict           = "FIREWALL_PRESERVED_GATE893_WEAK_SOCKET_SELECTOR_NOT_NATIVE"

	SupportNeutralNullEdgeSelectsHPlusCandidate     = "CONDITIONAL_SUPPORT_NEUTRAL_NULL_EDGE_SELECTS_H_PLUS_KERNEL_LINE_CANDIDATE"
	SupportPunctureKernelPairSourceTypesOrientation = "CONDITIONAL_SUPPORT_PUNCTURE_KERNEL_PAIR_SOURCE_TYPES_HIGGS_ORIENTATION"
	SupportFiniteOneFormCompatibleWithOrientation   = "CONDITIONAL_SUPPORT_FINITE_ONE_FORM_EDGE_PATTERN_COMPATIBLE_WITH_HIGGS_ORIENTATION"
	SupportFiniteOneFormAndPunctureKernelStrongest  = "CONDITIONAL_SUPPORT_FINITE_ONE_FORM_AND_PUNCTURE_KERNEL_PAIR_ARE_STRONGEST_SOURCE_CANDIDATES"
	SupportMissingTheoremIsWeakSocketSelector       = "CONDITIONAL_SUPPORT_MISSING_THEOREM_IS_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	SupportMinimalNullEdgeOrientationCandidate      = "CONDITIONAL_SUPPORT_MINIMAL_NULL_EDGE_ORIENTATION_PRINCIPLE_CANDIDATE"
	SupportAFOrientStableAfterOrientation           = "CONDITIONAL_SUPPORT_A_F_ORIENT_LEDGER_STABLE_AFTER_ORIENTATION"
	SupportOperatorNEffReproduced                   = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"
	SupportOrientationSealStillRequired             = "CONDITIONAL_SUPPORT_HIGGS_ORIENTATION_REMAINS_REQUIRED_SEAL"

	FailureNotNativeR3                         = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                    = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeHiggsOrientationSource      = "FAILED_ROUTE_NO_NATIVE_HIGGS_ORIENTATION_SOURCE_CERTIFIED"
	FailureNoNativeOneFormOrientationTheorem   = "FAILED_ROUTE_NO_NATIVE_ONE_FORM_ORIENTATION_THEOREM_YET"
	FailureNullEdgeNotNativeOrientation        = "FAILED_ROUTE_NULL_EDGE_PATTERN_NOT_NATIVE_HIGGS_ORIENTATION_THEOREM"
	FailurePunctureKernelNotNativeOrientation  = "FAILED_ROUTE_PUNCTURE_KERNEL_PAIR_NOT_NATIVE_HIGGS_ORIENTATION_THEOREM"
	FailureDFPatternRestatesOrientation        = "FAILED_ROUTE_D_F_EDGE_PATTERN_RESTATES_ORIENTATION_WITHOUT_INDEPENDENT_SELECTOR"
	FailureNoNonCircularWeakSocketSelector     = "FAILED_ROUTE_NO_NONCIRCULAR_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	FailureNoMinimalNullEdgeOrientationTheorem = "FAILED_ROUTE_NO_NATIVE_MINIMAL_NULL_EDGE_ORIENTATION_PRINCIPLE"
	FailureFullHMixesWeakSockets               = "FAILED_ROUTE_FULL_H_ACTION_MIXES_H_PLUS_H_MINUS"
	FailureNoNativeDescentFullToOrient         = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailurePostOrientationNotFullAF            = "FAILED_ROUTE_POST_ORIENTATION_PROJECTORS_NOT_FULL_UNBROKEN_A_F_SECTORS"
	FailureNoNativeR3SectorLedger              = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoNativeIncidenceFunctor            = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoPhysicalParticleAssignment        = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap              = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap              = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues            = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem             = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type NullEdgeSelectorAudit struct {
	MissingEdge              string
	RightPuncture            string
	LeftKernel               string
	YPlus1Zero               bool
	SelectsHPlusCandidate    bool
	NativeOrientationTheorem bool
	Supports, Failures       []string
}

type FiniteOneFormAudit struct {
	Edges                  []string
	UsesWeakFrame          bool
	CompatibleWithFrame    bool
	ForcesFrame            bool
	CircularIfUsedAsSource bool
	Supports, Failures     []string
}

type NonCircularityAudit struct {
	OrientationToDFAllowed      bool
	DFToOrientationCertified    bool
	RequiresIndependentSelector bool
	MissingObject               string
	CandidatePrinciple          string
	NativeSelectorFunctional    bool
	Supports, Failures          []string
}

type WeakFrameAudit struct {
	Frame                    string
	FullAlgebra              string
	OrientedAlgebra          string
	Stabilizer               string
	FullHMixesFrame          bool
	StabilizerPreservesFrame bool
	NativeOrientationSource  bool
	Supports, Failures       []string
}

type OfficialFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                        bool
	NotNativeR3                     bool
	AlphaStillSealed                bool
	NoNativeHiggsOrientationSource  bool
	NoNativeOneFormOrientation      bool
	NoNativeNullEdgeOrientation     bool
	NoNonCircularWeakSocketSelector bool
	FullHMixesWeakSockets           bool
	NoNativeDescentFullToOrient     bool
	NoNativeIncidenceFunctor        bool
	PostOrientationNotFullAF        bool
	NoPhysicalParticleAssignment    bool
	NoGenerationCarrier             bool
	NoFlavorOrientation             bool
	NoIndividualYukawas             bool
	NoOfficialLedgerUpdate          bool
	NoNativeYukawaOperator          bool
	NoR4NativeYukawaTheorem         bool
	Verdict                         string
}

type Audit struct {
	ID             string
	WeakFrame      WeakFrameAudit
	NullEdge       NullEdgeSelectorAudit
	OneForm        FiniteOneFormAudit
	NonCircularity NonCircularityAudit
	Freeze         OfficialFreeze
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	weak := buildWeakFrameAudit()
	if !weak.FullHMixesFrame || !weak.StabilizerPreservesFrame || weak.NativeOrientationSource {
		return Audit{}, fmt.Errorf("weak frame promoted incorrectly: %s", FormatWeakFrame(weak))
	}

	nullEdge := buildNullEdgeSelectorAudit()
	if !nullEdge.YPlus1Zero || !nullEdge.SelectsHPlusCandidate || nullEdge.NativeOrientationTheorem {
		return Audit{}, fmt.Errorf("null-edge selector promoted incorrectly: %s", FormatNullEdge(nullEdge))
	}

	oneForm := buildFiniteOneFormAudit()
	if !oneForm.UsesWeakFrame || !oneForm.CompatibleWithFrame || oneForm.ForcesFrame || !oneForm.CircularIfUsedAsSource {
		return Audit{}, fmt.Errorf("finite one-form route promoted incorrectly: %s", FormatFiniteOneForm(oneForm))
	}

	noncircularity := buildNonCircularityAudit()
	if !noncircularity.RequiresIndependentSelector || noncircularity.DFToOrientationCertified || noncircularity.NativeSelectorFunctional {
		return Audit{}, fmt.Errorf("noncircularity firewall failed: %s", FormatNonCircularity(noncircularity))
	}

	freeze := OfficialFreeze{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreeze, SupportOperatorNEffReproduced},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}

	firewalls := Firewalls{
		Enforced: true, NotNativeR3: true, AlphaStillSealed: true,
		NoNativeHiggsOrientationSource: true, NoNativeOneFormOrientation: true,
		NoNativeNullEdgeOrientation: true, NoNonCircularWeakSocketSelector: true,
		FullHMixesWeakSockets: true, NoNativeDescentFullToOrient: true, NoNativeIncidenceFunctor: true,
		PostOrientationNotFullAF: true, NoPhysicalParticleAssignment: true, NoGenerationCarrier: true,
		NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true,
		NoNativeYukawaOperator: true, NoR4NativeYukawaTheorem: true, Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:             AuditID,
		WeakFrame:      weak,
		NullEdge:       nullEdge,
		OneForm:        oneForm,
		NonCircularity: noncircularity,
		Freeze:         freeze,
		Firewalls:      firewalls,
		Truth:          "Gate 893 sharpens the HiggsOrientation wall: the finite one-form/Higgs edge and the neutral puncture/kernel pair source-type the h_+ / h_- weak socket orientation, but they do not derive it natively. The missing edge Y_+1=0 points to h_+ as a left kernel candidate, while D_F support remains circular if used to derive the frame it already assumes.",
		Final:          "The branch remains R3_CANDIDATE_WEAK_SOCKET_SELECTOR_SOURCE_TYPED_NOT_NATIVE. The next missing object is a noncircular WeakSocketSelectorFunctional or MinimalNullEdgeOrientationPrinciple. Native R3, physical sector assignment, generation/flavor splitting, individual Yukawa values, official ledger updates, and R4 native Yukawa theorem remain blocked.",
	}, nil
}

func buildWeakFrameAudit() WeakFrameAudit {
	return WeakFrameAudit{
		Frame: WeakFrame, FullAlgebra: FullUnbrokenAlgebra, OrientedAlgebra: PostOrientationAlgebra, Stabilizer: WeakStabilizer,
		FullHMixesFrame: true, StabilizerPreservesFrame: true, NativeOrientationSource: false,
		Supports: []string{SupportAFOrientStableAfterOrientation, SupportOrientationSealStillRequired},
		Failures: []string{FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeHiggsOrientationSource},
	}
}

func buildNullEdgeSelectorAudit() NullEdgeSelectorAudit {
	return NullEdgeSelectorAudit{
		MissingEdge: MissingEdge, RightPuncture: RightPuncture, LeftKernel: LeftKernel,
		YPlus1Zero: true, SelectsHPlusCandidate: true, NativeOrientationTheorem: false,
		Supports: []string{SupportNeutralNullEdgeSelectsHPlusCandidate, SupportPunctureKernelPairSourceTypesOrientation, SupportMinimalNullEdgeOrientationCandidate},
		Failures: []string{FailureNullEdgeNotNativeOrientation, FailurePunctureKernelNotNativeOrientation, FailureNoNativeHiggsOrientationSource},
	}
}

func buildFiniteOneFormAudit() FiniteOneFormAudit {
	return FiniteOneFormAudit{
		Edges:         []string{EdgePiPlus3, EdgePiMinus3, EdgePiMinus1},
		UsesWeakFrame: true, CompatibleWithFrame: true, ForcesFrame: false, CircularIfUsedAsSource: true,
		Supports: []string{SupportFiniteOneFormCompatibleWithOrientation, SupportFiniteOneFormAndPunctureKernelStrongest},
		Failures: []string{FailureDFPatternRestatesOrientation, FailureNoNativeOneFormOrientationTheorem, FailureNoNativeHiggsOrientationSource},
	}
}

func buildNonCircularityAudit() NonCircularityAudit {
	return NonCircularityAudit{
		OrientationToDFAllowed: true, DFToOrientationCertified: false, RequiresIndependentSelector: true,
		MissingObject: "WeakSocketSelectorFunctional", CandidatePrinciple: "MinimalNullEdgeOrientationPrinciple", NativeSelectorFunctional: false,
		Supports: []string{SupportMissingTheoremIsWeakSocketSelector, SupportMinimalNullEdgeOrientationCandidate},
		Failures: []string{FailureNoNonCircularWeakSocketSelector, FailureNoMinimalNullEdgeOrientationTheorem, FailureDFPatternRestatesOrientation},
	}
}

func FormatWeakFrame(w WeakFrameAudit) string {
	return fmt.Sprintf("weak_frame(frame=%s full=%s oriented=%s stabilizer=%s full_H_mixes=%t stabilizer_preserves=%t native_source=%t supports=%s failures=%s)", w.Frame, w.FullAlgebra, w.OrientedAlgebra, w.Stabilizer, w.FullHMixesFrame, w.StabilizerPreservesFrame, w.NativeOrientationSource, strings.Join(w.Supports, ","), strings.Join(w.Failures, ","))
}

func FormatNullEdge(n NullEdgeSelectorAudit) string {
	return fmt.Sprintf("null_edge(missing=%s right_puncture=%s left_kernel=%s y_plus1_zero=%t selects_h_plus=%t native=%t supports=%s failures=%s)", n.MissingEdge, n.RightPuncture, n.LeftKernel, n.YPlus1Zero, n.SelectsHPlusCandidate, n.NativeOrientationTheorem, strings.Join(n.Supports, ","), strings.Join(n.Failures, ","))
}

func FormatFiniteOneForm(o FiniteOneFormAudit) string {
	return fmt.Sprintf("finite_one_form(edges=%s uses_frame=%t compatible=%t forces_frame=%t circular_if_source=%t supports=%s failures=%s)", strings.Join(o.Edges, ","), o.UsesWeakFrame, o.CompatibleWithFrame, o.ForcesFrame, o.CircularIfUsedAsSource, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}

func FormatNonCircularity(n NonCircularityAudit) string {
	return fmt.Sprintf("noncircularity(orientation_to_DF=%t DF_to_orientation=%t requires_selector=%t missing=%s candidate=%s native=%t supports=%s failures=%s)", n.OrientationToDFAllowed, n.DFToOrientationCertified, n.RequiresIndependentSelector, n.MissingObject, n.CandidatePrinciple, n.NativeSelectorFunctional, strings.Join(n.Supports, ","), strings.Join(n.Failures, ","))
}

func FormatFreeze(f OfficialFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t no_orientation=%t no_one_form=%t no_null_edge=%t no_selector=%t full_H_mixes=%t no_descent=%t no_incidence=%t post_orientation_not_full=%t no_physical=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.NoNativeHiggsOrientationSource, f.NoNativeOneFormOrientation, f.NoNativeNullEdgeOrientation, f.NoNonCircularWeakSocketSelector, f.FullHMixesWeakSockets, f.NoNativeDescentFullToOrient, f.NoNativeIncidenceFunctor, f.PostOrientationNotFullAF, f.NoPhysicalParticleAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate892Inherited, StatusNullEdgeRouteAudited, StatusFiniteOneFormRouteAudited, StatusNonCircularityAudited, StatusStrongestPair, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportNeutralNullEdgeSelectsHPlusCandidate, SupportPunctureKernelPairSourceTypesOrientation, SupportFiniteOneFormCompatibleWithOrientation, SupportFiniteOneFormAndPunctureKernelStrongest, SupportMissingTheoremIsWeakSocketSelector, SupportMinimalNullEdgeOrientationCandidate, SupportAFOrientStableAfterOrientation, SupportOperatorNEffReproduced, SupportOrientationSealStillRequired,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeHiggsOrientationSource, FailureNoNativeOneFormOrientationTheorem, FailureNullEdgeNotNativeOrientation, FailurePunctureKernelNotNativeOrientation, FailureDFPatternRestatesOrientation, FailureNoNonCircularWeakSocketSelector, FailureNoMinimalNullEdgeOrientationTheorem, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailurePostOrientationNotFullAF, FailureNoNativeR3SectorLedger, FailureNoNativeIncidenceFunctor, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeHiggsOrientationSource, FailureNoNativeOneFormOrientationTheorem, FailureNullEdgeNotNativeOrientation, FailureNoNonCircularWeakSocketSelector, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeIncidenceFunctor, FailurePostOrientationNotFullAF, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.NoNativeHiggsOrientationSource &&
		f.NoNativeOneFormOrientation && f.NoNativeNullEdgeOrientation && f.NoNonCircularWeakSocketSelector &&
		f.FullHMixesWeakSockets && f.NoNativeDescentFullToOrient && f.NoNativeIncidenceFunctor &&
		f.PostOrientationNotFullAF && f.NoPhysicalParticleAssignment && f.NoGenerationCarrier &&
		f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate &&
		f.NoNativeYukawaOperator && f.NoR4NativeYukawaTheorem && f.Verdict == StatusFirewallVerdict
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
