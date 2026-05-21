// Package generation2fullafdescenthiggsorientationsourceobstructionaudit implements
// Gate 891: Full A_F Descent and HiggsOrientation Source Obstruction Audit.
//
// Gate 891 follows Gate 890's dual-seal/J-mirror classification. It asks whether
// the post-orientation finite-sector projector ledger can descend from
// A_F^orient=C_R+C_H+M_3(C) to the full unbroken A_F=C+H+M_3(C), and whether
// the Higgs/weak socket orientation has a certified native source.
package generation2fullafdescenthiggsorientationsourceobstructionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE891-FULL-AF-DESCENT-HIGGS-ORIENTATION-SOURCE-OBSTRUCTION-AUDIT"

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
	RankHFMin    = 30

	AtomPiPlus3  = "Pi_+3=e_+ tensor P_3"
	AtomPiMinus3 = "Pi_-3=e_- tensor P_3"
	AtomPiMinus1 = "Pi_-1=e_- tensor P_1"

	EdgePiPlus3  = "Pi_+3 -> h_+ tensor P_3"
	EdgePiMinus3 = "Pi_-3 -> h_- tensor P_3"
	EdgePiMinus1 = "Pi_-1 -> h_- tensor P_1"

	PostOrientationAlgebra = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullUnbrokenAlgebra    = "A_F=C plus H plus M_3(C)"
	WeakStabilizer         = "C_H=Stab_H(h_+ plus h_-)"
	WeakFrame              = "C_L^2=h_+ plus h_-"

	Classification = "R3_DUALSEAL_J_MIRROR_FULL_A_F_DESCENT_BLOCKED_BY_HIGGS_ORIENTATION_SEAL"
	ShortStatus    = "R3_CANDIDATE_UNDER_ALPHA_AND_ORIENTATION_SEALS_DESCENT_BLOCKED"
	NextFrontier   = "HIGGS_ORIENTATION_SOURCE_THEOREM_OR_NATIVE_FULL_A_F_DESCENT_REQUIRED"

	StatusGate890Inherited          = "PASS_GATE890_DUALSEAL_J_MIRROR_STATUS_INHERITED"
	StatusFullAFStabilityAudited    = "PASS_FULL_A_F_STABILITY_AUDITED"
	StatusFullHMixingDetected       = "PASS_FULL_H_ACTION_MIXES_WEAK_SOCKET_FRAME"
	StatusStabilizerIdentified      = "PASS_HIGGS_ORIENTED_STABILIZER_IDENTIFIED"
	StatusOrientationSourcesAudited = "PASS_HIGGS_ORIENTATION_SOURCE_CANDIDATES_AUDITED"
	StatusDescentVerdict            = "PASS_FULL_DESCENT_VERDICT_RECORDED"
	StatusOfficialFreeze            = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict           = "FIREWALL_PRESERVED_GATE891_DESCENT_BLOCKED_NOT_NATIVE_R3"

	SupportAFOrientStable                     = "CONDITIONAL_SUPPORT_A_F_ORIENT_LEDGER_IS_STABLE_IN_POST_ORIENTATION_LAYER"
	SupportFullToOrientRestriction            = "CONDITIONAL_SUPPORT_FULL_TO_ORIENTED_BRANCH_IS_A_HIGGS_ORIENTATION_RESTRICTION"
	SupportDualSealValidAfterOrientation      = "CONDITIONAL_SUPPORT_DUAL_SEAL_LEDGER_IS_VALID_ONLY_AFTER_ORIENTATION"
	SupportAFOrientIsStabilizer               = "CONDITIONAL_SUPPORT_A_F_ORIENT_IS_STABILIZER_OF_HIGGS_WEAK_SOCKET_FRAME"
	SupportStabilizerSourceTyped              = "CONDITIONAL_SUPPORT_STABILIZER_SOURCE_TYPED_BUT_NOT_NATIVE_DESCENT"
	SupportOrientationSourceCandidatesAudited = "CONDITIONAL_SUPPORT_ORIENTATION_SOURCE_CANDIDATES_AUDITED"
	SupportR3CandidateRequiresOrientationSeal = "CONDITIONAL_SUPPORT_R3_CANDIDATE_REQUIRES_SPONTANEOUS_OR_ORIENTED_WEAK_FRAME_SEAL"
	SupportProjectorLedgerCoherentUnderOrient = "CONDITIONAL_SUPPORT_PROJECTOR_LEDGER_AND_TRACE_READOUT_REMAIN_COHERENT_UNDER_ORIENTATION_SEAL"
	SupportOperatorNEffReproduced             = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"

	FailureNotNativeR3                          = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                     = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureAFOrientNotFullAF                    = "FAILED_ROUTE_A_F_ORIENT_NOT_EQUAL_FULL_A_F"
	FailureFullHMixesWeakSockets                = "FAILED_ROUTE_FULL_H_ACTION_MIXES_H_PLUS_H_MINUS"
	FailureSocketProjectorsNotStableFullAF      = "FAILED_ROUTE_SOCKET_PROJECTORS_NOT_FULL_A_F_STABLE"
	FailureSocketProjectorsNotStableFullH       = "FAILED_ROUTE_SOCKET_PROJECTORS_NOT_STABLE_UNDER_FULL_H_ACTION"
	FailureNoNativeDescentFullToOrient          = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureStabilizerNotFullNativeAF            = "FAILED_ROUTE_STABILIZER_NOT_FULL_NATIVE_A_F"
	FailureNoNativeHiggsOrientationSource       = "FAILED_ROUTE_NO_NATIVE_HIGGS_ORIENTATION_SOURCE_CERTIFIED"
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

type SocketProjector struct {
	Atom                  string
	Rank                  int
	EdgeSupport           string
	StableUnderAFOrient   bool
	StableUnderFullAF     bool
	StableUnderFullH      bool
	PhysicalSector        bool
	GenerationResolved    bool
	FlavorResolved        bool
	IndividualYukawaValue bool
	Supports, Failures    []string
}

type StabilizerAudit struct {
	OrientedAlgebra        string
	FullAlgebra            string
	WeakFrame              string
	Stabilizer             string
	PreservesHPlusHMinus   bool
	PreservesProjectors    bool
	IsFullH                bool
	IsFullAF               bool
	NativeDescentCertified bool
	Supports, Failures     []string
}

type OrientationSourceCandidate struct {
	Name      string
	Audited   bool
	Plausible bool
	Certified bool
	Reason    string
	Supports  []string
	Failures  []string
}

type OrientationSourceAudit struct {
	Candidates               []OrientationSourceCandidate
	AnyNativeSourceCertified bool
	StrongestCandidate       string
	RequiresOrientationSeal  bool
	Supports, Failures       []string
}

type FullDescentAudit struct {
	Projectors               []SocketProjector
	AFOrientStable           bool
	FullAFStable             bool
	GenericHMixesWeakSockets bool
	WeakFrameFullHInvariant  bool
	NativeDescentCertified   bool
	Outcome                  string
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
	Enforced                       bool
	NotNativeR3                    bool
	AlphaStillSealed               bool
	NoNativeIncidenceFunctor       bool
	AFOrientNotFullAF              bool
	FullHMixesWeakSockets          bool
	NoNativeDescentFullToOrient    bool
	NoNativeHiggsOrientationSource bool
	PostOrientationNotFullAF       bool
	NoNativeFiniteSectorProjector  bool
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
	ID                 string
	Projectors         []SocketProjector
	Stabilizer         StabilizerAudit
	OrientationSources OrientationSourceAudit
	Descent            FullDescentAudit
	Freeze             OfficialFreeze
	Firewalls          Firewalls
	Truth              string
	Final              string
}

func BuildDefault() (Audit, error) {
	projectors := []SocketProjector{
		buildProjector(AtomPiPlus3, RankPiPlus3, EdgePiPlus3),
		buildProjector(AtomPiMinus3, RankPiMinus3, EdgePiMinus3),
		buildProjector(AtomPiMinus1, RankPiMinus1, EdgePiMinus1),
	}
	if !projectorsOK(projectors) {
		return Audit{}, fmt.Errorf("projector firewall leak: %s", FormatProjectors(projectors))
	}

	stabilizer := buildStabilizerAudit()
	if !stabilizer.PreservesHPlusHMinus || stabilizer.IsFullH || stabilizer.IsFullAF || stabilizer.NativeDescentCertified {
		return Audit{}, fmt.Errorf("stabilizer promoted incorrectly: %s", FormatStabilizer(stabilizer))
	}

	orientationSources := buildOrientationSourceAudit()
	if orientationSources.AnyNativeSourceCertified || !orientationSources.RequiresOrientationSeal {
		return Audit{}, fmt.Errorf("orientation source promoted incorrectly: %s", FormatOrientationSources(orientationSources))
	}

	descent := buildFullDescentAudit(projectors)
	if descent.FullAFStable || descent.WeakFrameFullHInvariant || descent.NativeDescentCertified || !descent.GenericHMixesWeakSockets {
		return Audit{}, fmt.Errorf("descent promoted incorrectly: %s", FormatDescent(descent))
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
		Enforced: true, NotNativeR3: true, AlphaStillSealed: true, NoNativeIncidenceFunctor: true,
		AFOrientNotFullAF: true, FullHMixesWeakSockets: true, NoNativeDescentFullToOrient: true,
		NoNativeHiggsOrientationSource: true, PostOrientationNotFullAF: true, NoNativeFiniteSectorProjector: true,
		NoPhysicalParticleAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true,
		NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true,
		NoR4NativeYukawaTheorem: true, Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID: AuditID, Projectors: projectors, Stabilizer: stabilizer, OrientationSources: orientationSources,
		Descent: descent, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 891 confirms that the oriented finite-sector ledger is stable only after Higgs/weak orientation. Generic full quaternionic action mixes h_+ and h_-, so the current R3 candidate cannot descend from A_F^orient to full A_F without a native orientation-source theorem.",
		Final: "The branch is classified as R3_DUALSEAL_J_MIRROR_FULL_A_F_DESCENT_BLOCKED_BY_HIGGS_ORIENTATION_SEAL. Native R3 remains blocked by the BoundaryExteriorIncidenceFlagFunctor and by the absence of a native full-A_F descent / HiggsOrientationSource theorem. No physical sector assignment, generation/flavor split, individual Yukawa value, official ledger update, or R4 native Yukawa theorem is permitted.",
	}, nil
}

func buildProjector(atom string, rank int, edge string) SocketProjector {
	return SocketProjector{
		Atom: atom, Rank: rank, EdgeSupport: edge,
		StableUnderAFOrient: true, StableUnderFullAF: false, StableUnderFullH: false,
		PhysicalSector: false, GenerationResolved: false, FlavorResolved: false, IndividualYukawaValue: false,
		Supports: []string{SupportAFOrientStable, SupportProjectorLedgerCoherentUnderOrient},
		Failures: []string{FailureFullHMixesWeakSockets, FailureSocketProjectorsNotStableFullAF, FailurePostOrientationNotFullAF, FailureNoPhysicalParticleAssignment},
	}
}

func buildStabilizerAudit() StabilizerAudit {
	return StabilizerAudit{
		OrientedAlgebra: PostOrientationAlgebra, FullAlgebra: FullUnbrokenAlgebra, WeakFrame: WeakFrame, Stabilizer: WeakStabilizer,
		PreservesHPlusHMinus: true, PreservesProjectors: true, IsFullH: false, IsFullAF: false, NativeDescentCertified: false,
		Supports: []string{SupportAFOrientIsStabilizer, SupportAFOrientStable, SupportFullToOrientRestriction, SupportDualSealValidAfterOrientation, SupportStabilizerSourceTyped},
		Failures: []string{FailureStabilizerNotFullNativeAF, FailureAFOrientNotFullAF, FailureNoNativeDescentFullToOrient},
	}
}

func buildOrientationSourceAudit() OrientationSourceAudit {
	candidates := []OrientationSourceCandidate{
		candidate("finite one-form / Higgs edge", true, true, false, "plausible orientation carrier but no native selection theorem"),
		candidate("D_F symbolic support", true, true, false, "post-orientation edge support depends on h_+/h_- frame"),
		candidate("left kernel h_+ tensor P_1", true, true, false, "orientation-relative kernel singleton, not a native source"),
		candidate("right puncture e_+ tensor P_1", true, true, false, "puncture types the branch but does not derive weak frame"),
		candidate("B-L imbalance", true, true, false, "supports socket polarity but not quaternionic descent"),
		candidate("BoundaryAlpha seal", true, false, false, "alpha response remains separate incidence-flag seal"),
		candidate("K7 polarity", true, false, false, "no typed map from K7 polarity to weak socket orientation"),
	}
	return OrientationSourceAudit{
		Candidates: candidates, AnyNativeSourceCertified: false, StrongestCandidate: "finite one-form / Higgs edge plus D_F symbolic support", RequiresOrientationSeal: true,
		Supports: []string{SupportOrientationSourceCandidatesAudited, SupportR3CandidateRequiresOrientationSeal},
		Failures: []string{FailureNoNativeHiggsOrientationSource, FailureNoNativeDescentFullToOrient, FailurePostOrientationNotFullAF},
	}
}

func candidate(name string, audited, plausible, certified bool, reason string) OrientationSourceCandidate {
	failures := []string{FailureNoNativeHiggsOrientationSource}
	if !plausible {
		failures = append(failures, FailureNoNativeDescentFullToOrient)
	}
	return OrientationSourceCandidate{Name: name, Audited: audited, Plausible: plausible, Certified: certified, Reason: reason, Supports: []string{SupportOrientationSourceCandidatesAudited}, Failures: failures}
}

func buildFullDescentAudit(projectors []SocketProjector) FullDescentAudit {
	return FullDescentAudit{
		Projectors: projectors, AFOrientStable: true, FullAFStable: false, GenericHMixesWeakSockets: true,
		WeakFrameFullHInvariant: false, NativeDescentCertified: false, Outcome: "OUTCOME_B_C_STABILIZER_SOURCE_TYPED_BUT_FULL_DESCENT_BLOCKED_BY_HIGGS_ORIENTATION_SEAL",
		Supports: []string{SupportAFOrientStable, SupportFullToOrientRestriction, SupportDualSealValidAfterOrientation, SupportR3CandidateRequiresOrientationSeal},
		Failures: []string{FailureFullHMixesWeakSockets, FailureSocketProjectorsNotStableFullH, FailureSocketProjectorsNotStableFullAF, FailureNoNativeDescentFullToOrient, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger},
	}
}

func FormatProjector(p SocketProjector) string {
	return fmt.Sprintf("projector(%s rank=%d edge=%s orient_stable=%t full_A_F=%t full_H=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", p.Atom, p.Rank, p.EdgeSupport, p.StableUnderAFOrient, p.StableUnderFullAF, p.StableUnderFullH, p.PhysicalSector, p.GenerationResolved, p.FlavorResolved, p.IndividualYukawaValue, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatProjectors(projectors []SocketProjector) string {
	parts := make([]string, 0, len(projectors))
	for _, p := range projectors {
		parts = append(parts, FormatProjector(p))
	}
	return strings.Join(parts, "; ")
}

func FormatStabilizer(s StabilizerAudit) string {
	return fmt.Sprintf("stabilizer(oriented=%s full=%s weak_frame=%s stabilizer=%s preserves_frame=%t preserves_projectors=%t is_full_H=%t is_full_A_F=%t native_descent=%t supports=%s failures=%s)", s.OrientedAlgebra, s.FullAlgebra, s.WeakFrame, s.Stabilizer, s.PreservesHPlusHMinus, s.PreservesProjectors, s.IsFullH, s.IsFullAF, s.NativeDescentCertified, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatCandidate(c OrientationSourceCandidate) string {
	return fmt.Sprintf("candidate(%s audited=%t plausible=%t certified=%t reason=%q supports=%s failures=%s)", c.Name, c.Audited, c.Plausible, c.Certified, c.Reason, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatCandidates(candidates []OrientationSourceCandidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "; ")
}

func FormatOrientationSources(o OrientationSourceAudit) string {
	return fmt.Sprintf("orientation_sources(strongest=%s native_source=%t requires_seal=%t candidates=[%s] supports=%s failures=%s)", o.StrongestCandidate, o.AnyNativeSourceCertified, o.RequiresOrientationSeal, FormatCandidates(o.Candidates), strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}

func FormatDescent(d FullDescentAudit) string {
	return fmt.Sprintf("full_descent(orient_stable=%t full_A_F_stable=%t generic_H_mixes=%t weak_frame_full_H_invariant=%t native_descent=%t outcome=%s projectors=[%s] supports=%s failures=%s)", d.AFOrientStable, d.FullAFStable, d.GenericHMixesWeakSockets, d.WeakFrameFullHInvariant, d.NativeDescentCertified, d.Outcome, FormatProjectors(d.Projectors), strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}

func FormatFreeze(f OfficialFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t no_incidence=%t AForient_not_full=%t full_H_mixes=%t no_descent=%t no_orientation_source=%t post_orientation_not_full=%t no_native_sector=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.NoNativeIncidenceFunctor, f.AFOrientNotFullAF, f.FullHMixesWeakSockets, f.NoNativeDescentFullToOrient, f.NoNativeHiggsOrientationSource, f.PostOrientationNotFullAF, f.NoNativeFiniteSectorProjector, f.NoPhysicalParticleAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate890Inherited,
		StatusFullAFStabilityAudited,
		StatusFullHMixingDetected,
		StatusStabilizerIdentified,
		StatusOrientationSourcesAudited,
		StatusDescentVerdict,
		StatusOfficialFreeze,
		StatusFirewallVerdict,
		SupportAFOrientStable,
		SupportFullToOrientRestriction,
		SupportDualSealValidAfterOrientation,
		SupportAFOrientIsStabilizer,
		SupportStabilizerSourceTyped,
		SupportOrientationSourceCandidatesAudited,
		SupportR3CandidateRequiresOrientationSeal,
		SupportProjectorLedgerCoherentUnderOrient,
		SupportOperatorNEffReproduced,
		FailureNotNativeR3,
		FailureAlphaStillSealed,
		FailureAFOrientNotFullAF,
		FailureFullHMixesWeakSockets,
		FailureSocketProjectorsNotStableFullAF,
		FailureSocketProjectorsNotStableFullH,
		FailureNoNativeDescentFullToOrient,
		FailureStabilizerNotFullNativeAF,
		FailureNoNativeHiggsOrientationSource,
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

func projectorsOK(projectors []SocketProjector) bool {
	if len(projectors) != 3 {
		return false
	}
	rank := 0
	for _, p := range projectors {
		rank += p.Rank
		if !p.StableUnderAFOrient || p.StableUnderFullAF || p.StableUnderFullH || p.PhysicalSector || p.GenerationResolved || p.FlavorResolved || p.IndividualYukawaValue {
			return false
		}
	}
	return rank == RankHRMin
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.NoNativeIncidenceFunctor && f.AFOrientNotFullAF && f.FullHMixesWeakSockets && f.NoNativeDescentFullToOrient && f.NoNativeHiggsOrientationSource && f.PostOrientationNotFullAF && f.NoNativeFiniteSectorProjector && f.NoPhysicalParticleAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4NativeYukawaTheorem && f.Verdict == StatusFirewallVerdict
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
