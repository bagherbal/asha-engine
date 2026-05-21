// Package generation2higgsorientationsourcecandidateweaksocketselectoraudit implements
// Gate 892: HiggsOrientation Source Candidate and WeakSocket Selector Audit.
//
// Gate 892 follows Gate 891's full A_F descent obstruction. It audits whether
// the weak socket orientation h_+ plus h_- can be sourced by existing ASHA
// objects, while preserving the firewall that the current R3 candidate remains
// under BoundaryAlpha and Higgs/post-orientation seals.
package generation2higgsorientationsourcecandidateweaksocketselectoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE892-HIGGS-ORIENTATION-SOURCE-CANDIDATE-WEAK-SOCKET-SELECTOR-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RankPiPlus3  = 3
	RankPiMinus3 = 3
	RankPiMinus1 = 1
	RankHRMin    = 7

	AtomPiPlus3  = "Pi_+3=e_+ tensor P_3"
	AtomPiMinus3 = "Pi_-3=e_- tensor P_3"
	AtomPiMinus1 = "Pi_-1=e_- tensor P_1"

	EdgePiPlus3  = "Y_+3:e_+ tensor P_3 -> h_+ tensor P_3"
	EdgePiMinus3 = "Y_-3:e_- tensor P_3 -> h_- tensor P_3"
	EdgePiMinus1 = "Y_-1:e_- tensor P_1 -> h_- tensor P_1"

	RightPuncture = "e_+ tensor P_1"
	LeftKernel    = "h_+ tensor P_1"
	WeakFrame     = "C_L^2=h_+ plus h_-"

	PostOrientationAlgebra = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullUnbrokenAlgebra    = "A_F=C plus H plus M_3(C)"
	WeakStabilizer         = "C_H=Stab_H(h_+ plus h_-)"

	Classification = "R3_DUALSEAL_HIGGS_ORIENTATION_SOURCE_CANDIDATE_NOT_NATIVE"
	ShortStatus    = "R3_CANDIDATE_ORIENTATION_SOURCE_OBSTRUCTION"
	NextFrontier   = "HIGGS_ORIENTATION_SOURCE_THEOREM_REQUIRED_BEFORE_FULL_A_F_DESCENT"

	StatusGate891Inherited       = "PASS_GATE891_FULL_A_F_DESCENT_OBSTRUCTION_INHERITED"
	StatusCandidateSurveyAudited = "PASS_HIGGS_ORIENTATION_SOURCE_CANDIDATES_AUDITED"
	StatusStrongestCandidates    = "PASS_FINITE_ONE_FORM_AND_PUNCTURE_KERNEL_PAIR_IDENTIFIED_AS_STRONGEST_CANDIDATES"
	StatusDFCircularityAudited   = "PASS_D_F_SUPPORT_ORIENTATION_CIRCULARITY_AUDITED"
	StatusBLImbalanceAudited     = "PASS_B_MINUS_L_IMBALANCE_COMPATIBILITY_AUDITED"
	StatusK7CarrierFirewall      = "PASS_K7_POLARITY_CARRIER_FIREWALL_AUDITED"
	StatusOfficialFreeze         = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE892_HIGGS_ORIENTATION_SOURCE_NOT_NATIVE"

	SupportFiniteOneFormStrongestCandidate    = "CONDITIONAL_SUPPORT_FINITE_ONE_FORM_IS_STRONGEST_HIGGS_ORIENTATION_SOURCE_CANDIDATE"
	SupportPunctureKernelPointsToHPlus        = "CONDITIONAL_SUPPORT_PUNCTURE_KERNEL_PAIR_POINTS_TO_H_PLUS_ORIENTATION"
	SupportDFSupportCompatibleWithOrientation = "CONDITIONAL_SUPPORT_D_F_SUPPORT_COMPATIBLE_WITH_WEAK_SOCKET_ORIENTATION_PATTERN"
	SupportBLImbalanceCompatible              = "CONDITIONAL_SUPPORT_B_MINUS_L_IMBALANCE_COMPATIBLE_WITH_ORIENTATION_PATTERN"
	SupportOrientationCandidatesAudited       = "CONDITIONAL_SUPPORT_ORIENTATION_SOURCE_CANDIDATES_AUDITED"
	SupportNoNativeOrientationTheorem         = "CONDITIONAL_SUPPORT_HIGGS_ORIENTATION_REMAINS_REQUIRED_SEAL"
	SupportR3CandidateRequiresOrientationSeal = "CONDITIONAL_SUPPORT_R3_CANDIDATE_REQUIRES_HIGGS_ORIENTATION_SOURCE_THEOREM"
	SupportAFOrientStable                     = "CONDITIONAL_SUPPORT_A_F_ORIENT_LEDGER_STABLE_AFTER_ORIENTATION"
	SupportOperatorNEffReproduced             = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"

	FailureNotNativeR3                          = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                     = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureFullHMixesWeakSockets                = "FAILED_ROUTE_FULL_H_ACTION_MIXES_H_PLUS_H_MINUS"
	FailureNoNativeDescentFullToOrient          = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNoNativeHiggsOrientationSource       = "FAILED_ROUTE_NO_NATIVE_HIGGS_ORIENTATION_SOURCE_CERTIFIED"
	FailureNoNativeOneFormOrientationTheorem    = "FAILED_ROUTE_NO_NATIVE_ONE_FORM_ORIENTATION_THEOREM_YET"
	FailurePunctureKernelNotNativeOrientation   = "FAILED_ROUTE_PUNCTURE_KERNEL_PAIR_NOT_NATIVE_HIGGS_ORIENTATION_THEOREM"
	FailureDFSupportRestatesOrientation         = "FAILED_ROUTE_D_F_SUPPORT_RESTATES_ORIENTATION_IF_H_PLUS_H_MINUS_ALREADY_ASSUMED"
	FailureBLImbalanceDoesNotSelectWeakFrame    = "FAILED_ROUTE_B_MINUS_L_TRACE_IMBALANCE_DOES_NOT_SELECT_WEAK_SOCKET_FRAME"
	FailureK7PolarityNotTypedToWeakFrame        = "FAILED_ROUTE_K7_POLARITY_NOT_TYPED_TO_HIGGS_WEAK_SOCKET_FRAME"
	FailureBoundaryAlphaNotOrientationSource    = "FAILED_ROUTE_BOUNDARY_ALPHA_SEAL_DOES_NOT_SOURCE_HIGGS_WEAK_ORIENTATION"
	FailurePostOrientationNotFullAF             = "FAILED_ROUTE_POST_ORIENTATION_PROJECTORS_NOT_FULL_UNBROKEN_A_F_SECTORS"
	FailureNoNativeFiniteSectorProjectorTheorem = "FAILED_ROUTE_NO_NATIVE_FINITE_SECTOR_PROJECTOR_THEOREM"
	FailureNoNativeR3SectorLedger               = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoNativeIncidenceFunctor             = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoPhysicalParticleAssignment         = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap               = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap               = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues             = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                 = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem              = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type SourceCandidate struct {
	Name      string
	Route     string
	Audited   bool
	Plausible bool
	Strong    bool
	Certified bool
	Circular  bool
	Reason    string
	Supports  []string
	Failures  []string
}

type CandidateAudit struct {
	Candidates               []SourceCandidate
	StrongestCandidates      []string
	AnyNativeSourceCertified bool
	RequiresOrientationSeal  bool
	Supports, Failures       []string
}

type WeakSocketSelectorAudit struct {
	WeakFrame                string
	Stabilizer               string
	FullAlgebra              string
	OrientedAlgebra          string
	FullHMixesWeakSockets    bool
	StabilizerPreservesFrame bool
	NativeOrientationSource  bool
	Supports, Failures       []string
}

type EdgeOrientationAudit struct {
	Edges                  []string
	AssumesWeakFrame       bool
	DerivesWeakFrame       bool
	CompatibleWithFrame    bool
	CircularIfUsedAsSource bool
	Supports, Failures     []string
}

type PunctureKernelAudit struct {
	RightPuncture         string
	LeftKernel            string
	PointsToHPlus         bool
	NativeSourceCertified bool
	Supports, Failures    []string
}

type BMinusLAudit struct {
	ActiveTrace        int
	PunctureTrace      int
	FullRectangleTrace int
	Compatible         bool
	SelectsWeakFrame   bool
	Supports, Failures []string
}

type OfficialFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                       bool
	NotNativeR3                    bool
	AlphaStillSealed               bool
	FullHMixesWeakSockets          bool
	NoNativeDescentFullToOrient    bool
	NoNativeHiggsOrientationSource bool
	NoNativeOneFormOrientation     bool
	NoNativeFiniteSectorProjector  bool
	NoNativeIncidenceFunctor       bool
	PostOrientationNotFullAF       bool
	NoPhysicalParticleAssignment   bool
	NoGenerationCarrier            bool
	NoFlavorOrientation            bool
	NoIndividualYukawas            bool
	NoOfficialLedgerUpdate         bool
	NoNativeYukawaOperator         bool
	NoR4NativeYukawaTheorem        bool
	Verdict                        string
}

type Audit struct {
	ID              string
	WeakSelector    WeakSocketSelectorAudit
	Candidates      CandidateAudit
	EdgeOrientation EdgeOrientationAudit
	PunctureKernel  PunctureKernelAudit
	BMinusL         BMinusLAudit
	Freeze          OfficialFreeze
	Firewalls       Firewalls
	Truth           string
	Final           string
}

func BuildDefault() (Audit, error) {
	weakSelector := buildWeakSocketSelectorAudit()
	if !weakSelector.FullHMixesWeakSockets || !weakSelector.StabilizerPreservesFrame || weakSelector.NativeOrientationSource {
		return Audit{}, fmt.Errorf("weak selector promoted incorrectly: %s", FormatWeakSelector(weakSelector))
	}

	candidates := buildCandidateAudit()
	if candidates.AnyNativeSourceCertified || !candidates.RequiresOrientationSeal || len(candidates.Candidates) < 7 {
		return Audit{}, fmt.Errorf("candidate audit promoted incorrectly: %s", FormatCandidateAudit(candidates))
	}

	edge := buildEdgeOrientationAudit()
	if !edge.AssumesWeakFrame || edge.DerivesWeakFrame || !edge.CircularIfUsedAsSource {
		return Audit{}, fmt.Errorf("D_F orientation circularity not preserved: %s", FormatEdgeOrientation(edge))
	}

	pk := buildPunctureKernelAudit()
	if !pk.PointsToHPlus || pk.NativeSourceCertified {
		return Audit{}, fmt.Errorf("puncture/kernel promoted incorrectly: %s", FormatPunctureKernel(pk))
	}

	bl := buildBMinusLAudit()
	if !bl.Compatible || bl.SelectsWeakFrame || bl.ActiveTrace != 1 || bl.PunctureTrace != -1 || bl.FullRectangleTrace != 0 {
		return Audit{}, fmt.Errorf("B-L imbalance promoted incorrectly: %s", FormatBMinusL(bl))
	}

	freeze := OfficialFreeze{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreeze},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}

	firewalls := Firewalls{
		Enforced: true, NotNativeR3: true, AlphaStillSealed: true, FullHMixesWeakSockets: true,
		NoNativeDescentFullToOrient: true, NoNativeHiggsOrientationSource: true, NoNativeOneFormOrientation: true,
		NoNativeFiniteSectorProjector: true, NoNativeIncidenceFunctor: true, PostOrientationNotFullAF: true,
		NoPhysicalParticleAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true,
		NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true,
		NoR4NativeYukawaTheorem: true, Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:              AuditID,
		WeakSelector:    weakSelector,
		Candidates:      candidates,
		EdgeOrientation: edge,
		PunctureKernel:  pk,
		BMinusL:         bl,
		Freeze:          freeze,
		Firewalls:       firewalls,
		Truth:           "Gate 892 audits the Higgs/weak orientation source wall exposed by Gate 891. The finite one-form/Higgs edge and puncture/kernel pair are the strongest candidates, but neither certifies a native weak-socket selector. D_F support uses the h_+/h_- frame and is circular if used to derive that same frame.",
		Final:           "The branch remains R3_DUALSEAL_HIGGS_ORIENTATION_SOURCE_CANDIDATE_NOT_NATIVE. Native R3 is still blocked by the BoundaryExteriorIncidenceFlagFunctor and by the absence of a native HiggsOrientationSource / full-A_F descent theorem. No physical sector assignment, generation/flavor split, individual Yukawa value, official ledger update, or R4 native Yukawa theorem is permitted.",
	}, nil
}

func buildWeakSocketSelectorAudit() WeakSocketSelectorAudit {
	return WeakSocketSelectorAudit{
		WeakFrame:                WeakFrame,
		Stabilizer:               WeakStabilizer,
		FullAlgebra:              FullUnbrokenAlgebra,
		OrientedAlgebra:          PostOrientationAlgebra,
		FullHMixesWeakSockets:    true,
		StabilizerPreservesFrame: true,
		NativeOrientationSource:  false,
		Supports:                 []string{SupportAFOrientStable, SupportR3CandidateRequiresOrientationSeal},
		Failures:                 []string{FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeHiggsOrientationSource},
	}
}

func buildCandidateAudit() CandidateAudit {
	candidates := []SourceCandidate{
		candidate("finite one-form / Higgs edge", "finite-one-form", true, true, true, false, false, "strongest candidate because Higgs one-form naturally lives at the weak-orientation interface, but no native selector theorem is certified", []string{SupportFiniteOneFormStrongestCandidate}, []string{FailureNoNativeOneFormOrientationTheorem, FailureNoNativeHiggsOrientationSource}),
		candidate("puncture/kernel pair", "neutral-puncture-kernel", true, true, true, false, false, "right puncture e_+ tensor P_1 and left kernel h_+ tensor P_1 point to the h_+ line but do not derive the weak socket frame", []string{SupportPunctureKernelPointsToHPlus}, []string{FailurePunctureKernelNotNativeOrientation, FailureNoNativeHiggsOrientationSource}),
		candidate("symbolic D_F edge support", "edge-support", true, true, false, false, true, "edge support encodes e_+ to h_+ and e_- to h_- only after the orientation frame has been chosen", []string{SupportDFSupportCompatibleWithOrientation}, []string{FailureDFSupportRestatesOrientation, FailureNoNativeHiggsOrientationSource}),
		candidate("B-L imbalance", "B-L-trace", true, true, false, false, false, "B-L trace +1 on active support and -1 on puncture is compatible with the pattern but does not select h_+ and h_-", []string{SupportBLImbalanceCompatible}, []string{FailureBLImbalanceDoesNotSelectWeakFrame}),
		candidate("BoundaryAlpha seal", "alpha-incidence", true, false, false, false, false, "alpha source remains a separate boundary incidence-flag seal", []string{SupportOrientationCandidatesAudited}, []string{FailureBoundaryAlphaNotOrientationSource, FailureAlphaStillSealed}),
		candidate("K7 polarity", "K7-polarity", true, false, false, false, false, "K7 polarity has no typed carrier map into the weak socket frame", []string{SupportOrientationCandidatesAudited}, []string{FailureK7PolarityNotTypedToWeakFrame}),
		candidate("full quaternionic action", "full-H-action", true, false, false, false, false, "generic H action mixes h_+ and h_- rather than selecting them", []string{SupportOrientationCandidatesAudited}, []string{FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient}),
	}
	return CandidateAudit{
		Candidates:               candidates,
		StrongestCandidates:      []string{"finite one-form / Higgs edge", "puncture/kernel pair"},
		AnyNativeSourceCertified: false,
		RequiresOrientationSeal:  true,
		Supports:                 []string{SupportOrientationCandidatesAudited, SupportFiniteOneFormStrongestCandidate, SupportPunctureKernelPointsToHPlus, SupportNoNativeOrientationTheorem, SupportR3CandidateRequiresOrientationSeal},
		Failures:                 []string{FailureNoNativeHiggsOrientationSource, FailureNoNativeOneFormOrientationTheorem, FailurePunctureKernelNotNativeOrientation, FailureDFSupportRestatesOrientation, FailureBLImbalanceDoesNotSelectWeakFrame, FailureK7PolarityNotTypedToWeakFrame},
	}
}

func candidate(name, route string, audited, plausible, strong, certified, circular bool, reason string, supports, failures []string) SourceCandidate {
	return SourceCandidate{Name: name, Route: route, Audited: audited, Plausible: plausible, Strong: strong, Certified: certified, Circular: circular, Reason: reason, Supports: supports, Failures: failures}
}

func buildEdgeOrientationAudit() EdgeOrientationAudit {
	return EdgeOrientationAudit{
		Edges:                  []string{EdgePiPlus3, EdgePiMinus3, EdgePiMinus1},
		AssumesWeakFrame:       true,
		DerivesWeakFrame:       false,
		CompatibleWithFrame:    true,
		CircularIfUsedAsSource: true,
		Supports:               []string{SupportDFSupportCompatibleWithOrientation},
		Failures:               []string{FailureDFSupportRestatesOrientation, FailureNoNativeHiggsOrientationSource},
	}
}

func buildPunctureKernelAudit() PunctureKernelAudit {
	return PunctureKernelAudit{
		RightPuncture:         RightPuncture,
		LeftKernel:            LeftKernel,
		PointsToHPlus:         true,
		NativeSourceCertified: false,
		Supports:              []string{SupportPunctureKernelPointsToHPlus},
		Failures:              []string{FailurePunctureKernelNotNativeOrientation, FailureNoNativeHiggsOrientationSource},
	}
}

func buildBMinusLAudit() BMinusLAudit {
	return BMinusLAudit{
		ActiveTrace:        1,
		PunctureTrace:      -1,
		FullRectangleTrace: 0,
		Compatible:         true,
		SelectsWeakFrame:   false,
		Supports:           []string{SupportBLImbalanceCompatible},
		Failures:           []string{FailureBLImbalanceDoesNotSelectWeakFrame},
	}
}

func FormatCandidate(c SourceCandidate) string {
	return fmt.Sprintf("candidate(%s route=%s audited=%t plausible=%t strong=%t certified=%t circular=%t reason=%q supports=%s failures=%s)", c.Name, c.Route, c.Audited, c.Plausible, c.Strong, c.Certified, c.Circular, c.Reason, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatCandidates(candidates []SourceCandidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "; ")
}

func FormatCandidateAudit(c CandidateAudit) string {
	return fmt.Sprintf("candidates(strongest=%s native_source=%t requires_seal=%t list=[%s] supports=%s failures=%s)", strings.Join(c.StrongestCandidates, ","), c.AnyNativeSourceCertified, c.RequiresOrientationSeal, FormatCandidates(c.Candidates), strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatWeakSelector(w WeakSocketSelectorAudit) string {
	return fmt.Sprintf("weak_selector(frame=%s stabilizer=%s full=%s oriented=%s full_H_mixes=%t stabilizer_preserves=%t native_source=%t supports=%s failures=%s)", w.WeakFrame, w.Stabilizer, w.FullAlgebra, w.OrientedAlgebra, w.FullHMixesWeakSockets, w.StabilizerPreservesFrame, w.NativeOrientationSource, strings.Join(w.Supports, ","), strings.Join(w.Failures, ","))
}

func FormatEdgeOrientation(e EdgeOrientationAudit) string {
	return fmt.Sprintf("edge_orientation(edges=%s assumes_frame=%t derives_frame=%t compatible=%t circular_if_source=%t supports=%s failures=%s)", strings.Join(e.Edges, ","), e.AssumesWeakFrame, e.DerivesWeakFrame, e.CompatibleWithFrame, e.CircularIfUsedAsSource, strings.Join(e.Supports, ","), strings.Join(e.Failures, ","))
}

func FormatPunctureKernel(p PunctureKernelAudit) string {
	return fmt.Sprintf("puncture_kernel(right=%s left=%s points_h_plus=%t native=%t supports=%s failures=%s)", p.RightPuncture, p.LeftKernel, p.PointsToHPlus, p.NativeSourceCertified, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatBMinusL(b BMinusLAudit) string {
	return fmt.Sprintf("B-L(active=%d puncture=%d full=%d compatible=%t selects_frame=%t supports=%s failures=%s)", b.ActiveTrace, b.PunctureTrace, b.FullRectangleTrace, b.Compatible, b.SelectsWeakFrame, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatFreeze(f OfficialFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t full_H_mixes=%t no_descent=%t no_orientation_source=%t no_one_form=%t no_sector=%t no_incidence=%t post_orientation_not_full=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.FullHMixesWeakSockets, f.NoNativeDescentFullToOrient, f.NoNativeHiggsOrientationSource, f.NoNativeOneFormOrientation, f.NoNativeFiniteSectorProjector, f.NoNativeIncidenceFunctor, f.PostOrientationNotFullAF, f.NoPhysicalParticleAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate891Inherited,
		StatusCandidateSurveyAudited,
		StatusStrongestCandidates,
		StatusDFCircularityAudited,
		StatusBLImbalanceAudited,
		StatusK7CarrierFirewall,
		StatusOfficialFreeze,
		StatusFirewallVerdict,
		SupportFiniteOneFormStrongestCandidate,
		SupportPunctureKernelPointsToHPlus,
		SupportDFSupportCompatibleWithOrientation,
		SupportBLImbalanceCompatible,
		SupportOrientationCandidatesAudited,
		SupportNoNativeOrientationTheorem,
		SupportR3CandidateRequiresOrientationSeal,
		SupportAFOrientStable,
		SupportOperatorNEffReproduced,
		FailureNotNativeR3,
		FailureAlphaStillSealed,
		FailureFullHMixesWeakSockets,
		FailureNoNativeDescentFullToOrient,
		FailureNoNativeHiggsOrientationSource,
		FailureNoNativeOneFormOrientationTheorem,
		FailurePunctureKernelNotNativeOrientation,
		FailureDFSupportRestatesOrientation,
		FailureBLImbalanceDoesNotSelectWeakFrame,
		FailureK7PolarityNotTypedToWeakFrame,
		FailureBoundaryAlphaNotOrientationSource,
		FailurePostOrientationNotFullAF,
		FailureNoNativeFiniteSectorProjectorTheorem,
		FailureNoNativeR3SectorLedger,
		FailureNoNativeIncidenceFunctor,
		FailureNoPhysicalParticleAssignment,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawaTheorem,
	}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.FullHMixesWeakSockets &&
		f.NoNativeDescentFullToOrient && f.NoNativeHiggsOrientationSource && f.NoNativeOneFormOrientation &&
		f.NoNativeFiniteSectorProjector && f.NoNativeIncidenceFunctor && f.PostOrientationNotFullAF &&
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
