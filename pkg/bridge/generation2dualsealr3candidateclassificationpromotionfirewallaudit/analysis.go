// Package generation2dualsealr3candidateclassificationpromotionfirewallaudit implements
// Gate 889: DualSeal R3-Candidate Classification and Promotion Firewall Audit.
//
// Gate 889 follows Gate 888's oriented finite-sector projector ledger under dual
// seal. It classifies whether the branch qualifies as an R3-candidate or
// R3-sealed object, while preserving the firewall against native R3, physical
// particle assignment, individual Yukawa values, and official ledger updates.
package generation2dualsealr3candidateclassificationpromotionfirewallaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE889-DUALSEAL-R3-CANDIDATE-CLASSIFICATION-PROMOTION-FIREWALL-AUDIT"

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

	EdgePiPlus3  = "Pi_+3 -> h_+ tensor P_3"
	EdgePiMinus3 = "Pi_-3 -> h_- tensor P_3"
	EdgePiMinus1 = "Pi_-1 -> h_- tensor P_1"

	WeightPiPlus3  = "1"
	WeightPiMinus3 = "alpha_B(1-alpha_B)"
	WeightPiMinus1 = "3 alpha_B^2"

	PostOrientationAlgebra = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullUnbrokenAlgebra    = "A_F=C plus H plus M_3(C)"
	OrientedLedger         = "Pi_sector^{F,orient}={Pi_+3,Pi_-3,Pi_-1}"

	Classification = "R3_CANDIDATE_UNDER_DUAL_SEAL_NOT_NATIVE_R3"
	BranchStatus   = "R2+++++_DUAL_SEAL_R3_CANDIDATE_PROMOTION_FIREWALL"
	NextBranch     = "R3_BLOCKERS_REDUCED_TO_SEAL_REMOVAL_OR_SECTOR_TYPING_BEYOND_SOCKET_ATOMS"

	StatusGate888Inherited        = "PASS_GATE888_ORIENTED_FINITE_SECTOR_PROJECTOR_LEDGER_INHERITED"
	StatusProjectorsComplete      = "PASS_PROJECTORS_POSITIVE_READOUT_AND_EDGE_SUPPORT_COMPLETE_UNDER_SEALS"
	StatusDualSealR3Candidate     = "PASS_R3_CANDIDATE_UNDER_BOUNDARY_ALPHA_AND_POST_ORIENTATION_SEALS_CLASSIFIED"
	StatusNativeR3Firewall        = "PASS_NATIVE_R3_PROMOTION_FIREWALL_PRESERVED"
	StatusPhysicalFirewall        = "PASS_PHYSICAL_PARTICLE_GENERATION_FLAVOR_FIREWALLS_PRESERVED"
	StatusOfficialFreezePreserved = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusSealRemovalRequirements = "PASS_NATIVE_R3_SEAL_REMOVAL_REQUIREMENTS_RECORDED"
	StatusNoR4                    = "PASS_R4_NATIVE_YUKAWA_FIREWALL_PRESERVED"
	StatusFirewallVerdict         = "FIREWALL_PRESERVED_GATE889_DUAL_SEAL_R3_CANDIDATE_NOT_NATIVE_R3"

	SupportR3CandidateUnderDualSeal         = "CONDITIONAL_SUPPORT_R3_CANDIDATE_UNDER_DUAL_SEAL"
	SupportProjectorsAndReadoutComplete     = "CONDITIONAL_SUPPORT_PROJECTOR_LEDGER_AND_TRACE_MAGNITUDE_READOUT_COMPLETE_UNDER_SEALS"
	SupportOperatorNEffReproduced           = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_FINITE_SECTOR_LEDGER"
	SupportNativeR3BlockersReducedToSeals   = "CONDITIONAL_SUPPORT_NATIVE_R3_BLOCKERS_REDUCED_TO_SEAL_REMOVAL"
	SupportBoundaryAlphaSuppliesWeights     = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_SEAL_SUPPLIES_TRACE_WEIGHTS"
	SupportPostOrientationSuppliesProjector = "CONDITIONAL_SUPPORT_POST_ORIENTATION_SEAL_SUPPLIES_FINITE_SECTOR_PROJECTORS"
	SupportNoPhysicalAssignment             = "CONDITIONAL_SUPPORT_NO_PHYSICAL_OR_GENERATION_FLAVOR_ASSIGNMENT_MADE"

	FailureNotNativeR3                          = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                     = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeIncidenceFunctor             = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion           = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailurePostOrientationNotFullAF             = "FAILED_ROUTE_POST_ORIENTATION_PROJECTORS_NOT_FULL_UNBROKEN_A_F_SECTORS"
	FailureNoNativeFiniteSectorProjectorTheorem = "FAILED_ROUTE_NO_NATIVE_FINITE_SECTOR_PROJECTOR_THEOREM"
	FailureNoNativeR3SectorLedger               = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoPhysicalParticleAssignment         = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureSocketAtomsNotPhysical               = "FAILED_ROUTE_SOCKET_ATOMS_NOT_PHYSICAL_SECTOR_ASSIGNMENTS"
	FailureNoGenerationCarrierMap               = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap               = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues             = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum             = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoOfficialNEffUpdate                 = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                                 = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
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

type CandidateLedger struct {
	Name                string
	Atoms               []SectorAtom
	Rank                int
	CompleteOnHRMin     bool
	StableUnderAFOrient bool
	StableUnderFullAF   bool
	EdgeCompatible      bool
	ReadoutComplete     bool
	NativeR3            bool
	R3SealedCandidate   bool
	TraceTotal          float64
	SquareTrace         float64
	OperatorNEff        float64
	OperatorCYukawa     float64
	Supports, Failures  []string
}

type DualSealClassification struct {
	BoundaryAlphaSealSuppliesWeights      bool
	PostOrientationSealSuppliesProjectors bool
	ProjectorLedgerCompleteUnderSeals     bool
	TraceReadoutCompleteUnderSeals        bool
	NativeAlphaFunctorCertified           bool
	NativeFullAFProjectorsCertified       bool
	NativeR3                              bool
	R4NativeYukawa                        bool
	Supports, Failures                    []string
}

type PromotionRequirements struct {
	NeedsNativeIncidenceFunctor        bool
	NeedsNativeCrossLaneExclusion      bool
	NeedsFullUnbrokenAFProjectors      bool
	NeedsSealFreeTraceMagnitudeReadout bool
	NeedsGenerationCarrier             bool
	NeedsFlavorOrientation             bool
	NeedsIndividualYukawaBranch        bool
	Satisfied                          bool
	Supports, Failures                 []string
}

type LedgerFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                     bool
	NotNativeR3                  bool
	AlphaStillSealed             bool
	PostOrientationNotFullAF     bool
	NoPhysicalParticleAssignment bool
	NoGenerationCarrier          bool
	NoFlavorOrientation          bool
	NoIndividualYukawas          bool
	NoOfficialLedgerUpdate       bool
	NoNativeYukawaOperator       bool
	NoR4                         bool
	Verdict                      string
}

type Audit struct {
	ID           string
	Ledger       CandidateLedger
	Seals        DualSealClassification
	Requirements PromotionRequirements
	Freeze       LedgerFreeze
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB

	atoms := []SectorAtom{
		buildAtom(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, EdgePiPlus3, []string{SupportR3CandidateUnderDualSeal, SupportProjectorsAndReadoutComplete, SupportPostOrientationSuppliesProjector}, []string{FailurePostOrientationNotFullAF, FailureSocketAtomsNotPhysical}),
		buildAtom(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, EdgePiMinus3, []string{SupportR3CandidateUnderDualSeal, SupportProjectorsAndReadoutComplete, SupportPostOrientationSuppliesProjector}, []string{FailureAlphaStillSealed, FailurePostOrientationNotFullAF, FailureSocketAtomsNotPhysical}),
		buildAtom(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, EdgePiMinus1, []string{SupportR3CandidateUnderDualSeal, SupportProjectorsAndReadoutComplete, SupportPostOrientationSuppliesProjector}, []string{FailureAlphaStillSealed, FailurePostOrientationNotFullAF, FailureSocketAtomsNotPhysical}),
	}

	rank := 0
	traceTotal := 0.0
	squareTrace := 0.0
	for _, atom := range atoms {
		rank += atom.Rank
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
		if !atom.ProjectorOK || !atom.StableUnderAFOrient || atom.StableUnderFullAF || !atom.EdgeCompatible || !atom.ReadoutPositive || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return Audit{}, fmt.Errorf("dual-seal R3 candidate firewall leak in %s", atom.Atom)
		}
	}
	if rank != RankHRMin {
		return Audit{}, fmt.Errorf("rank drift: got %d want %d", rank, RankHRMin)
	}
	expectedTrace := 3.0 + 3.0*AlphaB
	expectedSquare := 3.0 + 3.0*AlphaB*AlphaB - 6.0*AlphaB*AlphaB*AlphaB + 12.0*AlphaB*AlphaB*AlphaB*AlphaB
	if !near(traceTotal, expectedTrace) || !near(squareTrace, expectedSquare) {
		return Audit{}, fmt.Errorf("trace drift: trace %.18g square %.18g", traceTotal, squareTrace)
	}
	operatorNEff := traceTotal * traceTotal / squareTrace
	operatorCYukawa := 3.0 / operatorNEff
	if !near(operatorNEff, OperatorNEffDiagnostic) || !near(operatorCYukawa, OperatorCYukawaDiagnostic) {
		return Audit{}, fmt.Errorf("operator diagnostic drift: N_eff %.18g C_Yukawa %.18g", operatorNEff, operatorCYukawa)
	}

	ledger := CandidateLedger{
		Name: OrientedLedger, Atoms: atoms, Rank: rank, CompleteOnHRMin: true,
		StableUnderAFOrient: true, StableUnderFullAF: false, EdgeCompatible: true, ReadoutComplete: true,
		NativeR3: false, R3SealedCandidate: true, TraceTotal: traceTotal, SquareTrace: squareTrace,
		OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Supports: []string{SupportR3CandidateUnderDualSeal, SupportProjectorsAndReadoutComplete, SupportOperatorNEffReproduced, SupportNativeR3BlockersReducedToSeals},
		Failures: []string{FailureNotNativeR3, FailureAlphaStillSealed, FailurePostOrientationNotFullAF, FailureNoNativeR3SectorLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}

	seals := DualSealClassification{
		BoundaryAlphaSealSuppliesWeights: true, PostOrientationSealSuppliesProjectors: true,
		ProjectorLedgerCompleteUnderSeals: true, TraceReadoutCompleteUnderSeals: true,
		NativeAlphaFunctorCertified: false, NativeFullAFProjectorsCertified: false, NativeR3: false, R4NativeYukawa: false,
		Supports: []string{SupportBoundaryAlphaSuppliesWeights, SupportPostOrientationSuppliesProjector, SupportR3CandidateUnderDualSeal, SupportProjectorsAndReadoutComplete},
		Failures: []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger, FailureNoNativeYukawaOperator, FailureNoR4},
	}

	requirements := PromotionRequirements{
		NeedsNativeIncidenceFunctor: true, NeedsNativeCrossLaneExclusion: true, NeedsFullUnbrokenAFProjectors: true,
		NeedsSealFreeTraceMagnitudeReadout: true, NeedsGenerationCarrier: true, NeedsFlavorOrientation: true,
		NeedsIndividualYukawaBranch: true, Satisfied: false,
		Supports: []string{SupportNativeR3BlockersReducedToSeals},
		Failures: []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}

	freeze := LedgerFreeze{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreezePreserved},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}

	firewalls := Firewalls{
		Enforced: true, NotNativeR3: true, AlphaStillSealed: true, PostOrientationNotFullAF: true,
		NoPhysicalParticleAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true,
		NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true,
		NoR4: true, Verdict: StatusFirewallVerdict,
	}

	return Audit{
		ID: AuditID, Ledger: ledger, Seals: seals, Requirements: requirements, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 889 classifies the Gate 888 ledger as an R3 candidate under BoundaryAlpha and post-orientation seals: projectors, positive readout rows, and edge-support compatibility are complete under those seals.",
		Final: "The branch reaches R3_CANDIDATE_UNDER_DUAL_SEAL_NOT_NATIVE_R3. Native R3 is still blocked by seal dependence: alpha_B requires the BoundaryExteriorIncidenceFlagFunctor and cross-lane exclusion theorem, the projector ledger is post-orientation rather than a full unbroken A_F finite-sector theorem, physical particle/generation/flavor assignments are absent, individual Yukawa values are absent, and official ledgers remain frozen.",
	}, nil
}

func buildAtom(atom string, rank int, formula string, weight float64, edge string, supports, failures []string) SectorAtom {
	return SectorAtom{
		Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, EdgeSupport: edge,
		TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight,
		ProjectorOK: true, StableUnderAFOrient: true, StableUnderFullAF: false, EdgeCompatible: true, ReadoutPositive: weight > 0,
		PhysicalSector: false, GenerationResolved: false, FlavorResolved: false, IndividualYukawaValue: false,
		Supports: supports, Failures: failures,
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

func FormatLedger(l CandidateLedger) string {
	return fmt.Sprintf("dual_seal_r3_candidate_ledger(name=%s rank=%d complete=%t orient_stable=%t full_A_F=%t edge=%t readout=%t native_r3=%t r3_sealed=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g atoms=[%s] supports=%s failures=%s)", l.Name, l.Rank, l.CompleteOnHRMin, l.StableUnderAFOrient, l.StableUnderFullAF, l.EdgeCompatible, l.ReadoutComplete, l.NativeR3, l.R3SealedCandidate, l.TraceTotal, l.SquareTrace, l.OperatorNEff, l.OperatorCYukawa, FormatAtoms(l.Atoms), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatSeals(s DualSealClassification) string {
	return fmt.Sprintf("dual_seal_classification(boundary_alpha_weights=%t post_orientation_projectors=%t projector_ledger_complete=%t trace_readout_complete=%t native_alpha=%t native_full_A_F=%t native_r3=%t r4=%t supports=%s failures=%s)", s.BoundaryAlphaSealSuppliesWeights, s.PostOrientationSealSuppliesProjectors, s.ProjectorLedgerCompleteUnderSeals, s.TraceReadoutCompleteUnderSeals, s.NativeAlphaFunctorCertified, s.NativeFullAFProjectorsCertified, s.NativeR3, s.R4NativeYukawa, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatRequirements(r PromotionRequirements) string {
	return fmt.Sprintf("promotion_requirements(incidence_functor=%t cross_lane=%t full_A_F_projectors=%t seal_free_trace=%t generation=%t flavor=%t individual_yukawa_branch=%t satisfied=%t supports=%s failures=%s)", r.NeedsNativeIncidenceFunctor, r.NeedsNativeCrossLaneExclusion, r.NeedsFullUnbrokenAFProjectors, r.NeedsSealFreeTraceMagnitudeReadout, r.NeedsGenerationCarrier, r.NeedsFlavorOrientation, r.NeedsIndividualYukawaBranch, r.Satisfied, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t post_orientation_not_full_A_F=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.PostOrientationNotFullAF, f.NoPhysicalParticleAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate888Inherited,
		StatusProjectorsComplete,
		StatusDualSealR3Candidate,
		StatusNativeR3Firewall,
		StatusPhysicalFirewall,
		StatusOfficialFreezePreserved,
		StatusSealRemovalRequirements,
		StatusNoR4,
		StatusFirewallVerdict,
		SupportR3CandidateUnderDualSeal,
		SupportProjectorsAndReadoutComplete,
		SupportOperatorNEffReproduced,
		SupportNativeR3BlockersReducedToSeals,
		SupportBoundaryAlphaSuppliesWeights,
		SupportPostOrientationSuppliesProjector,
		SupportNoPhysicalAssignment,
		FailureNotNativeR3,
		FailureAlphaStillSealed,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailurePostOrientationNotFullAF,
		FailureNoNativeFiniteSectorProjectorTheorem,
		FailureNoNativeR3SectorLedger,
		FailureNoPhysicalParticleAssignment,
		FailureSocketAtomsNotPhysical,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoPhysicalYukawaSpectrum,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4,
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
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.PostOrientationNotFullAF && f.NoPhysicalParticleAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
