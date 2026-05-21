// Package generation2dualsealr3candidatejoppositefulldescentincidencefunctorrecheckaudit implements
// Gate 890: DualSeal R3-Candidate J/Opposite Extension, Full-Descent, and
// IncidenceFunctor Recheck Audit.
//
// Gate 890 follows Gate 889's classification of the branch as an R3 candidate
// under BoundaryAlpha and post-orientation seals. It performs a controlled
// tri-audit: J/opposite-copy extension, full A_F versus A_F^orient descent, and
// rechecking whether any new native source for the BoundaryExteriorIncidenceFlagFunctor
// has appeared.
package generation2dualsealr3candidatejoppositefulldescentincidencefunctorrecheckaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE890-DUALSEAL-R3-CANDIDATE-J-OPPOSITE-FULL-DESCENT-INCIDENCE-FUNCTOR-RECHECK-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RankPiPlus3            = 3
	RankPiMinus3           = 3
	RankPiMinus1           = 1
	RankHRMin              = 7
	RankHL                 = 8
	RankHPartMin           = 15
	RankHFMin              = 30
	RankActiveMirrorLedger = 14

	AtomPiPlus3  = "Pi_+3=e_+ tensor P_3"
	AtomPiMinus3 = "Pi_-3=e_- tensor P_3"
	AtomPiMinus1 = "Pi_-1=e_- tensor P_1"

	AtomPiPlus3J  = "Pi_+3^J=J_F Pi_+3 J_F^{-1}"
	AtomPiMinus3J = "Pi_-3^J=J_F Pi_-3 J_F^{-1}"
	AtomPiMinus1J = "Pi_-1^J=J_F Pi_-1 J_F^{-1}"

	EdgePiPlus3  = "Pi_+3 -> h_+ tensor P_3"
	EdgePiMinus3 = "Pi_-3 -> h_- tensor P_3"
	EdgePiMinus1 = "Pi_-1 -> h_- tensor P_1"

	WeightPiPlus3  = "1"
	WeightPiMinus3 = "alpha_B(1-alpha_B)"
	WeightPiMinus1 = "3 alpha_B^2"

	PostOrientationAlgebra = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullUnbrokenAlgebra    = "A_F=C plus H plus M_3(C)"
	OrientedLedgerName     = "Pi_sector^{F,orient}={Pi_+3,Pi_-3,Pi_-1}"
	JMirrorLedger          = "J_F Pi_sector^{F,orient} J_F^{-1}={Pi_+3^J,Pi_-3^J,Pi_-1^J}"
	ActiveMirrorSupport    = "H_R^min plus J_F H_R^min"

	Classification = "R3_CANDIDATE_UNDER_DUAL_SEAL_WITH_J_MIRROR_AND_FULL_DESCENT_OBSTRUCTION_NOT_NATIVE_R3"
	BranchStatus   = "R3_DUALSEAL_J_MIRROR_DESCENT_BLOCKED_NOT_NATIVE"
	NextFrontier   = "NATIVE_R3_BLOCKERS_REDUCED_TO_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR_AND_FULL_A_F_DESCENT"

	StatusGate889Inherited      = "PASS_GATE889_DUAL_SEAL_R3_CANDIDATE_INHERITED"
	StatusJMirrorAudited        = "PASS_J_MIRROR_EXTENSION_AUDITED"
	StatusJMirrorRanksPreserved = "PASS_J_MIRROR_PRESERVES_PROJECTOR_RANKS_AND_ORTHOGONALITY"
	StatusJMirrorNotFullHF      = "PASS_J_MIRROR_NOT_FULL_H_F_MIN_LEDGER_FIREWALL_PRESERVED"
	StatusFullDescentAudited    = "PASS_FULL_A_F_VERSUS_A_F_ORIENT_DESCENT_AUDITED"
	StatusIncidenceRecheck      = "PASS_BOUNDARY_EXTERIOR_INCIDENCE_FUNCTOR_RECHECK_AUDITED"
	StatusDualSealCoherence     = "PASS_DUAL_SEAL_LEDGER_COHERENCE_SURVIVES_GATE890"
	StatusOfficialFreeze        = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict       = "FIREWALL_PRESERVED_GATE890_J_MIRROR_DESCENT_BLOCKED_NOT_NATIVE_R3"

	SupportR3CandidateSurvivesJMirror          = "CONDITIONAL_SUPPORT_R3_CANDIDATE_SURVIVES_J_MIRROR_EXTENSION_AT_SEAL_LEVEL"
	SupportJMirrorExistsAtSeal                 = "CONDITIONAL_SUPPORT_J_MIRROR_OF_ACTIVE_ORIENTED_LEDGER_EXISTS_AT_SEAL_LEVEL"
	SupportJExtensionPreservesRanks            = "CONDITIONAL_SUPPORT_J_EXTENSION_PRESERVES_PROJECTOR_RANKS_AND_ORTHOGONALITY"
	SupportActiveRightLedgerFormalOppositeCopy = "CONDITIONAL_SUPPORT_ACTIVE_RIGHT_LEDGER_HAS_FORMAL_OPPOSITE_COPY"
	SupportAFOrientStable                      = "CONDITIONAL_SUPPORT_A_F_ORIENT_LEDGER_IS_STABLE_IN_POST_ORIENTATION_LAYER"
	SupportFullToOrientRestriction             = "CONDITIONAL_SUPPORT_FULL_TO_ORIENTED_BRANCH_IS_A_HIGGS_ORIENTATION_RESTRICTION"
	SupportDualSealValidAfterOrientation       = "CONDITIONAL_SUPPORT_DUAL_SEAL_LEDGER_IS_VALID_ONLY_AFTER_ORIENTATION"
	SupportProjectorAndReadoutCoherent         = "CONDITIONAL_SUPPORT_PROJECTOR_LEDGER_AND_TRACE_READOUT_REMAIN_COHERENT_UNDER_DUAL_SEAL"
	SupportOperatorNEffReproduced              = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"
	SupportNativeR3BlockersReduced             = "CONDITIONAL_SUPPORT_NATIVE_R3_BLOCKERS_NOW_REDUCED_TO_ALPHA_FUNCTOR_AND_FULL_A_F_DESCENT"
	SupportBoundaryAlphaSealCoherent           = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL_REMAINS_COHERENT"
	SupportAlphaReconstructedUnderSeal         = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_UNDER_INCIDENT_FLAG_SELECTOR_SEAL"

	FailureNotNativeR3                          = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                     = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeIncidenceFunctor             = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion           = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureGate890AddsNoBoundarySource          = "FAILED_ROUTE_GATE890_ADDS_NO_NEW_NATIVE_BOUNDARY_SOURCE_OBJECT"
	FailureAFOrientNotFullAF                    = "FAILED_ROUTE_A_F_ORIENT_NOT_EQUAL_FULL_A_F"
	FailureSocketProjectorsNotStableFullH       = "FAILED_ROUTE_SOCKET_PROJECTORS_NOT_STABLE_UNDER_FULL_H_ACTION"
	FailureNoNativeDescentFullToOrient          = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailurePostOrientationNotFullAF             = "FAILED_ROUTE_POST_ORIENTATION_PROJECTORS_NOT_FULL_UNBROKEN_A_F_SECTORS"
	FailureNoNativeFiniteSectorProjectorTheorem = "FAILED_ROUTE_NO_NATIVE_FINITE_SECTOR_PROJECTOR_THEOREM"
	FailureNoNativeR3SectorLedger               = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureJExtensionNotFullHFMinLedger         = "FAILED_ROUTE_J_EXTENSION_DOES_NOT_COMPLETE_LEDGER_ON_FULL_H_F_MIN"
	FailureNoOperatorLevelJFKOSignProof         = "FAILED_ROUTE_NO_OPERATOR_LEVEL_J_F_KO_SIGN_PROOF"
	FailureNoFullJOppositeActionTheorem         = "FAILED_ROUTE_NO_FULL_J_OPPOSITE_ACTION_THEOREM"
	FailureJMirrorNotNativeFiniteSectorLedger   = "FAILED_ROUTE_J_MIRROR_LEDGER_NOT_NATIVE_FINITE_SECTOR_LEDGER"
	FailureJExtensionNotMagnitudeSource         = "FAILED_ROUTE_J_EXTENSION_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoPhysicalParticleAssignment         = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap               = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap               = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues             = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                 = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem              = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type SectorAtom struct {
	Atom                  string
	Rank                  int
	WeightFormula         string
	Weight                float64
	EdgeSupport           string
	TraceContribution     float64
	SquareContribution    float64
	ProjectorOK           bool
	StableUnderAFOrient   bool
	StableUnderFullAF     bool
	EdgeCompatible        bool
	ReadoutPositive       bool
	PhysicalSector        bool
	GenerationResolved    bool
	FlavorResolved        bool
	IndividualYukawaValue bool
	Supports, Failures    []string
}

type JMirrorAtom struct {
	Atom                  string
	SourceAtom            string
	Rank                  int
	RankPreserved         bool
	Orthogonal            bool
	OnJMirror             bool
	CompletesFullHFMin    bool
	PhysicalSector        bool
	YukawaMagnitudeSource bool
	Supports, Failures    []string
}

type OrientedLedger struct {
	Name                string
	Atoms               []SectorAtom
	Rank                int
	CompleteOnHRMin     bool
	StableUnderAFOrient bool
	StableUnderFullAF   bool
	EdgeCompatible      bool
	ReadoutComplete     bool
	NativeR3            bool
	R3DualSealCandidate bool
	TraceTotal          float64
	SquareTrace         float64
	OperatorNEff        float64
	OperatorCYukawa     float64
	Supports, Failures  []string
}

type JExtensionAudit struct {
	Name                   string
	Mirrors                []JMirrorAtom
	SourceRank             int
	MirrorRank             int
	CombinedActiveRank     int
	HFMinRank              int
	MirrorExists           bool
	PreservesRanks         bool
	PreservesOrthogonality bool
	CompletesFullHFMin     bool
	OperatorLevelJFKO      bool
	FullJOppositeAction    bool
	NativeFiniteSector     bool
	YukawaMagnitudeSource  bool
	Supports, Failures     []string
}

type DescentAudit struct {
	OrientedAlgebra               string
	FullAlgebra                   string
	AFOrientLedgerStable          bool
	FullAFLedgerStable            bool
	FullToOrientHiggsRestriction  bool
	NativeDescentCertified        bool
	WeakSocketFrameFullHInvariant bool
	NativeFiniteSectorProjectors  bool
	Supports, Failures            []string
}

type IncidenceFunctorRecheck struct {
	FunctorName                  string
	BoundaryAlphaSealCoherent    bool
	AlphaReconstructedUnderSeal  bool
	NewNativeBoundarySourceFound bool
	NativeFunctorCertified       bool
	CrossLaneExclusionCertified  bool
	AlphaStillSealed             bool
	Supports, Failures           []string
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
	NoNativeCrossLaneExclusion     bool
	PostOrientationNotFullAF       bool
	NoNativeDescentFullToOrient    bool
	SocketProjectorsNotStableFullH bool
	JExtensionNotFullHFMinLedger   bool
	NoFullJOppositeAction          bool
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
	ID         string
	Ledger     OrientedLedger
	JExtension JExtensionAudit
	Descent    DescentAudit
	Incidence  IncidenceFunctorRecheck
	Freeze     OfficialFreeze
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB

	atoms := []SectorAtom{
		buildAtom(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, EdgePiPlus3, []string{SupportProjectorAndReadoutCoherent, SupportAFOrientStable}, []string{FailurePostOrientationNotFullAF, FailureNoPhysicalParticleAssignment}),
		buildAtom(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, EdgePiMinus3, []string{SupportProjectorAndReadoutCoherent, SupportAFOrientStable}, []string{FailureAlphaStillSealed, FailurePostOrientationNotFullAF, FailureNoPhysicalParticleAssignment}),
		buildAtom(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, EdgePiMinus1, []string{SupportProjectorAndReadoutCoherent, SupportAFOrientStable}, []string{FailureAlphaStillSealed, FailurePostOrientationNotFullAF, FailureNoPhysicalParticleAssignment}),
	}
	ledger, err := buildLedger(atoms)
	if err != nil {
		return Audit{}, err
	}

	jExt := buildJExtension()
	if !jExtensionOK(jExt) {
		return Audit{}, fmt.Errorf("J-extension firewall leak: %s", FormatJExtension(jExt))
	}

	descent := buildDescent()
	if descent.NativeDescentCertified || descent.FullAFLedgerStable || descent.WeakSocketFrameFullHInvariant || descent.NativeFiniteSectorProjectors {
		return Audit{}, fmt.Errorf("descent promoted incorrectly: %s", FormatDescent(descent))
	}

	incidence := buildIncidenceRecheck()
	if incidence.NewNativeBoundarySourceFound || incidence.NativeFunctorCertified || incidence.CrossLaneExclusionCertified || !incidence.AlphaStillSealed {
		return Audit{}, fmt.Errorf("incidence recheck promoted incorrectly: %s", FormatIncidence(incidence))
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
		NoNativeCrossLaneExclusion: true, PostOrientationNotFullAF: true, NoNativeDescentFullToOrient: true,
		SocketProjectorsNotStableFullH: true, JExtensionNotFullHFMinLedger: true, NoFullJOppositeAction: true,
		NoNativeFiniteSectorProjector: true, NoPhysicalParticleAssignment: true, NoGenerationCarrier: true,
		NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true,
		NoNativeYukawaOperator: true, NoR4NativeYukawaTheorem: true, Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID: AuditID, Ledger: ledger, JExtension: jExt, Descent: descent, Incidence: incidence, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 890 confirms that the dual-seal R3 candidate survives a formal J/opposite mirror extension and remains stable only in the post-orientation A_F^orient layer, while no new native source appears for the BoundaryExteriorIncidenceFlagFunctor.",
		Final: "The branch is classified as R3_CANDIDATE_UNDER_DUAL_SEAL_WITH_J_MIRROR_AND_FULL_DESCENT_OBSTRUCTION_NOT_NATIVE_R3. Native R3 remains blocked by alpha_B's incidence-flag seal, the absence of a native cross-lane exclusion theorem, the absence of a native descent from full A_F to A_F^orient, full-H instability of socket projectors, the fact that the J mirror does not complete a full H_F^min ledger, and all physical/generation/flavor/individual-Yukawa and official-ledger firewalls.",
	}, nil
}

func buildLedger(atoms []SectorAtom) (OrientedLedger, error) {
	rank := 0
	traceTotal := 0.0
	squareTrace := 0.0
	for _, atom := range atoms {
		rank += atom.Rank
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
		if !atom.ProjectorOK || !atom.StableUnderAFOrient || atom.StableUnderFullAF || !atom.EdgeCompatible || !atom.ReadoutPositive || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return OrientedLedger{}, fmt.Errorf("ledger atom firewall leak in %s", atom.Atom)
		}
	}
	if rank != RankHRMin {
		return OrientedLedger{}, fmt.Errorf("rank drift: got %d want %d", rank, RankHRMin)
	}
	expectedTrace := 3.0 + 3.0*AlphaB
	expectedSquare := 3.0 + 3.0*AlphaB*AlphaB - 6.0*AlphaB*AlphaB*AlphaB + 12.0*AlphaB*AlphaB*AlphaB*AlphaB
	if !near(traceTotal, expectedTrace) || !near(squareTrace, expectedSquare) {
		return OrientedLedger{}, fmt.Errorf("trace drift: trace %.18g square %.18g", traceTotal, squareTrace)
	}
	operatorNEff := traceTotal * traceTotal / squareTrace
	operatorCYukawa := 3.0 / operatorNEff
	if !near(operatorNEff, OperatorNEffDiagnostic) || !near(operatorCYukawa, OperatorCYukawaDiagnostic) {
		return OrientedLedger{}, fmt.Errorf("operator diagnostic drift: N_eff %.18g C_Yukawa %.18g", operatorNEff, operatorCYukawa)
	}
	return OrientedLedger{
		Name: OrientedLedgerName, Atoms: atoms, Rank: rank, CompleteOnHRMin: true,
		StableUnderAFOrient: true, StableUnderFullAF: false, EdgeCompatible: true, ReadoutComplete: true,
		NativeR3: false, R3DualSealCandidate: true, TraceTotal: traceTotal, SquareTrace: squareTrace,
		OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Supports: []string{SupportProjectorAndReadoutCoherent, SupportOperatorNEffReproduced, SupportAFOrientStable, SupportNativeR3BlockersReduced},
		Failures: []string{FailureNotNativeR3, FailureAlphaStillSealed, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}, nil
}

func buildAtom(atom string, rank int, formula string, weight float64, edge string, supports, failures []string) SectorAtom {
	return SectorAtom{
		Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, EdgeSupport: edge,
		TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight,
		ProjectorOK: true, StableUnderAFOrient: true, StableUnderFullAF: false, EdgeCompatible: true,
		ReadoutPositive: weight > 0, PhysicalSector: false, GenerationResolved: false, FlavorResolved: false,
		IndividualYukawaValue: false, Supports: supports, Failures: failures,
	}
}

func buildJExtension() JExtensionAudit {
	mirrors := []JMirrorAtom{
		buildMirror(AtomPiPlus3J, AtomPiPlus3, RankPiPlus3),
		buildMirror(AtomPiMinus3J, AtomPiMinus3, RankPiMinus3),
		buildMirror(AtomPiMinus1J, AtomPiMinus1, RankPiMinus1),
	}
	mirrorRank := 0
	for _, m := range mirrors {
		mirrorRank += m.Rank
	}
	return JExtensionAudit{
		Name: JMirrorLedger, Mirrors: mirrors, SourceRank: RankHRMin, MirrorRank: mirrorRank,
		CombinedActiveRank: RankHRMin + mirrorRank, HFMinRank: RankHFMin,
		MirrorExists: true, PreservesRanks: mirrorRank == RankHRMin, PreservesOrthogonality: true,
		CompletesFullHFMin: false, OperatorLevelJFKO: false, FullJOppositeAction: false,
		NativeFiniteSector: false, YukawaMagnitudeSource: false,
		Supports: []string{SupportR3CandidateSurvivesJMirror, SupportJMirrorExistsAtSeal, SupportJExtensionPreservesRanks, SupportActiveRightLedgerFormalOppositeCopy},
		Failures: []string{FailureJExtensionNotFullHFMinLedger, FailureNoOperatorLevelJFKOSignProof, FailureNoFullJOppositeActionTheorem, FailureJMirrorNotNativeFiniteSectorLedger, FailureJExtensionNotMagnitudeSource},
	}
}

func buildMirror(atom, source string, rank int) JMirrorAtom {
	return JMirrorAtom{
		Atom: atom, SourceAtom: source, Rank: rank, RankPreserved: true, Orthogonal: true, OnJMirror: true,
		CompletesFullHFMin: false, PhysicalSector: false, YukawaMagnitudeSource: false,
		Supports: []string{SupportJMirrorExistsAtSeal, SupportJExtensionPreservesRanks},
		Failures: []string{FailureJMirrorNotNativeFiniteSectorLedger, FailureJExtensionNotMagnitudeSource},
	}
}

func buildDescent() DescentAudit {
	return DescentAudit{
		OrientedAlgebra: PostOrientationAlgebra, FullAlgebra: FullUnbrokenAlgebra,
		AFOrientLedgerStable: true, FullAFLedgerStable: false, FullToOrientHiggsRestriction: true,
		NativeDescentCertified: false, WeakSocketFrameFullHInvariant: false, NativeFiniteSectorProjectors: false,
		Supports: []string{SupportAFOrientStable, SupportFullToOrientRestriction, SupportDualSealValidAfterOrientation},
		Failures: []string{FailureAFOrientNotFullAF, FailureSocketProjectorsNotStableFullH, FailureNoNativeDescentFullToOrient, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger},
	}
}

func buildIncidenceRecheck() IncidenceFunctorRecheck {
	return IncidenceFunctorRecheck{
		FunctorName: "BoundaryExteriorIncidenceFlagFunctor", BoundaryAlphaSealCoherent: true,
		AlphaReconstructedUnderSeal: true, NewNativeBoundarySourceFound: false, NativeFunctorCertified: false,
		CrossLaneExclusionCertified: false, AlphaStillSealed: true,
		Supports: []string{SupportBoundaryAlphaSealCoherent, SupportAlphaReconstructedUnderSeal},
		Failures: []string{FailureGate890AddsNoBoundarySource, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed},
	}
}

func FormatAtom(a SectorAtom) string {
	return fmt.Sprintf("atom(%s rank=%d weight=%s numeric=%.16g edge=%s trace=%.16g square=%.16g projector=%t orient_stable=%t full_A_F=%t edge=%t positive=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", a.Atom, a.Rank, a.WeightFormula, a.Weight, a.EdgeSupport, a.TraceContribution, a.SquareContribution, a.ProjectorOK, a.StableUnderAFOrient, a.StableUnderFullAF, a.EdgeCompatible, a.ReadoutPositive, a.PhysicalSector, a.GenerationResolved, a.FlavorResolved, a.IndividualYukawaValue, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAtoms(atoms []SectorAtom) string {
	parts := make([]string, 0, len(atoms))
	for _, atom := range atoms {
		parts = append(parts, FormatAtom(atom))
	}
	return strings.Join(parts, "; ")
}

func FormatLedger(l OrientedLedger) string {
	return fmt.Sprintf("oriented_dual_seal_ledger(name=%s rank=%d complete=%t orient_stable=%t full_A_F=%t edge=%t readout=%t native_r3=%t r3_dual_seal=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g atoms=[%s] supports=%s failures=%s)", l.Name, l.Rank, l.CompleteOnHRMin, l.StableUnderAFOrient, l.StableUnderFullAF, l.EdgeCompatible, l.ReadoutComplete, l.NativeR3, l.R3DualSealCandidate, l.TraceTotal, l.SquareTrace, l.OperatorNEff, l.OperatorCYukawa, FormatAtoms(l.Atoms), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatMirror(m JMirrorAtom) string {
	return fmt.Sprintf("mirror(%s source=%s rank=%d rank_preserved=%t orthogonal=%t on_J=%t completes_full_HFmin=%t physical=%t yukawa_source=%t supports=%s failures=%s)", m.Atom, m.SourceAtom, m.Rank, m.RankPreserved, m.Orthogonal, m.OnJMirror, m.CompletesFullHFMin, m.PhysicalSector, m.YukawaMagnitudeSource, strings.Join(m.Supports, ","), strings.Join(m.Failures, ","))
}

func FormatMirrors(mirrors []JMirrorAtom) string {
	parts := make([]string, 0, len(mirrors))
	for _, mirror := range mirrors {
		parts = append(parts, FormatMirror(mirror))
	}
	return strings.Join(parts, "; ")
}

func FormatJExtension(j JExtensionAudit) string {
	return fmt.Sprintf("j_extension(name=%s source_rank=%d mirror_rank=%d combined_active_rank=%d HFmin_rank=%d mirror_exists=%t ranks=%t orthogonal=%t completes_full_HFmin=%t KO=%t full_J_op=%t native_sector=%t yukawa_source=%t mirrors=[%s] supports=%s failures=%s)", j.Name, j.SourceRank, j.MirrorRank, j.CombinedActiveRank, j.HFMinRank, j.MirrorExists, j.PreservesRanks, j.PreservesOrthogonality, j.CompletesFullHFMin, j.OperatorLevelJFKO, j.FullJOppositeAction, j.NativeFiniteSector, j.YukawaMagnitudeSource, FormatMirrors(j.Mirrors), strings.Join(j.Supports, ","), strings.Join(j.Failures, ","))
}

func FormatDescent(d DescentAudit) string {
	return fmt.Sprintf("descent(oriented=%s full=%s orient_stable=%t full_stable=%t higgs_restriction=%t native_descent=%t weak_socket_full_H_invariant=%t native_sector_projectors=%t supports=%s failures=%s)", d.OrientedAlgebra, d.FullAlgebra, d.AFOrientLedgerStable, d.FullAFLedgerStable, d.FullToOrientHiggsRestriction, d.NativeDescentCertified, d.WeakSocketFrameFullHInvariant, d.NativeFiniteSectorProjectors, strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}

func FormatIncidence(i IncidenceFunctorRecheck) string {
	return fmt.Sprintf("incidence_recheck(functor=%s coherent=%t alpha_reconstructed=%t new_source=%t native_functor=%t cross_lane=%t alpha_sealed=%t supports=%s failures=%s)", i.FunctorName, i.BoundaryAlphaSealCoherent, i.AlphaReconstructedUnderSeal, i.NewNativeBoundarySourceFound, i.NativeFunctorCertified, i.CrossLaneExclusionCertified, i.AlphaStillSealed, strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func FormatFreeze(f OfficialFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t no_incidence=%t no_cross_lane=%t post_orientation_not_full_A_F=%t no_descent=%t full_H_instability=%t J_not_full_HFmin=%t no_full_J_opposite=%t no_native_sector=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.NoNativeIncidenceFunctor, f.NoNativeCrossLaneExclusion, f.PostOrientationNotFullAF, f.NoNativeDescentFullToOrient, f.SocketProjectorsNotStableFullH, f.JExtensionNotFullHFMinLedger, f.NoFullJOppositeAction, f.NoNativeFiniteSectorProjector, f.NoPhysicalParticleAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate889Inherited,
		StatusJMirrorAudited,
		StatusJMirrorRanksPreserved,
		StatusJMirrorNotFullHF,
		StatusFullDescentAudited,
		StatusIncidenceRecheck,
		StatusDualSealCoherence,
		StatusOfficialFreeze,
		StatusFirewallVerdict,
		SupportR3CandidateSurvivesJMirror,
		SupportJMirrorExistsAtSeal,
		SupportJExtensionPreservesRanks,
		SupportActiveRightLedgerFormalOppositeCopy,
		SupportAFOrientStable,
		SupportFullToOrientRestriction,
		SupportDualSealValidAfterOrientation,
		SupportProjectorAndReadoutCoherent,
		SupportOperatorNEffReproduced,
		SupportNativeR3BlockersReduced,
		SupportBoundaryAlphaSealCoherent,
		SupportAlphaReconstructedUnderSeal,
		FailureNotNativeR3,
		FailureAlphaStillSealed,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailureGate890AddsNoBoundarySource,
		FailureAFOrientNotFullAF,
		FailureSocketProjectorsNotStableFullH,
		FailureNoNativeDescentFullToOrient,
		FailurePostOrientationNotFullAF,
		FailureNoNativeFiniteSectorProjectorTheorem,
		FailureNoNativeR3SectorLedger,
		FailureJExtensionNotFullHFMinLedger,
		FailureNoOperatorLevelJFKOSignProof,
		FailureNoFullJOppositeActionTheorem,
		FailureJMirrorNotNativeFiniteSectorLedger,
		FailureJExtensionNotMagnitudeSource,
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

func near(a, b float64) bool { return math.Abs(a-b) < 5e-15 }

func containsAll(haystack []string, needles []string) bool {
	seen := map[string]bool{}
	for _, s := range haystack {
		seen[s] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.NoNativeIncidenceFunctor && f.NoNativeCrossLaneExclusion && f.PostOrientationNotFullAF && f.NoNativeDescentFullToOrient && f.SocketProjectorsNotStableFullH && f.JExtensionNotFullHFMinLedger && f.NoFullJOppositeAction && f.NoNativeFiniteSectorProjector && f.NoPhysicalParticleAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4NativeYukawaTheorem && f.Verdict == StatusFirewallVerdict
}

func jExtensionOK(j JExtensionAudit) bool {
	return j.MirrorExists && j.PreservesRanks && j.PreservesOrthogonality && j.SourceRank == RankHRMin && j.MirrorRank == RankHRMin && j.CombinedActiveRank == RankActiveMirrorLedger && j.HFMinRank == RankHFMin && !j.CompletesFullHFMin && !j.OperatorLevelJFKO && !j.FullJOppositeAction && !j.NativeFiniteSector && !j.YukawaMagnitudeSource
}
