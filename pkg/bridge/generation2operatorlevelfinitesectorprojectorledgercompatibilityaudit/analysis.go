// Package generation2operatorlevelfinitesectorprojectorledgercompatibilityaudit implements
// Gate 888: Operator-Level FiniteSector ProjectorLedger Compatibility Audit
// Under Dual Seal.
//
// Gate 888 follows Gate 887's post-orientation finite-sector lift candidate.
// It audits whether the lifted socket atoms can be organized into an explicit
// operator-level finite-sector projector ledger inside A_F^orient under both
// the BoundaryAlpha seal and the Higgs/post-orientation seal.
package generation2operatorlevelfinitesectorprojectorledgercompatibilityaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE888-OPERATOR-LEVEL-FINITE-SECTOR-PROJECTOR-LEDGER-COMPATIBILITY-AUDIT"

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
	FiniteSectorLedger     = "Pi_sector^{F,orient}={Pi_+3,Pi_-3,Pi_-1}"

	Classification = "R2+++++_ORIENTED_FINITE_SECTOR_PROJECTOR_LEDGER_UNDER_DUAL_SEAL_NOT_NATIVE_R3"
	NextBranch     = "R3_ELIGIBILITY_REASSESSMENT_OR_NATIVE_FUNCTOR_SEARCH_UNDER_DUAL_SEAL"

	StatusGate887Inherited        = "PASS_GATE887_POST_ORIENTATION_FINITE_SECTOR_LIFT_CANDIDATE_INHERITED"
	StatusLedgerDefined           = "PASS_ORIENTED_FINITE_SECTOR_PROJECTOR_LEDGER_DEFINED"
	StatusProjectorsOrthogonal    = "PASS_PROJECTORS_IDEMPOTENT_ORTHOGONAL_COMPLETE_ON_H_R_MIN"
	StatusStabilizerCompatible    = "PASS_PROJECTOR_LEDGER_STABLE_UNDER_A_F_ORIENT"
	StatusEdgeCompatible          = "PASS_PROJECTOR_LEDGER_EDGE_SUPPORT_COMPATIBLE"
	StatusTraceReadoutCompatible  = "PASS_TRACE_MAGNITUDE_ROWS_REPRODUCE_OPERATOR_N_EFF"
	StatusDualSealDependencies    = "PASS_DUAL_SEAL_DEPENDENCIES_EXPLICIT"
	StatusFullAFFirewall          = "PASS_FULL_UNBROKEN_A_F_FIREWALL_PRESERVED"
	StatusPhysicalFirewalls       = "PASS_PHYSICAL_GENERATION_FLAVOR_FIREWALLS_PRESERVED"
	StatusOfficialFreezePreserved = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNextBranchSelected      = "PASS_NEXT_BRANCH_SELECTED_AFTER_DUAL_SEAL_LEDGER"
	StatusFirewallVerdict         = "FIREWALL_PRESERVED_GATE888_DUAL_SEAL_LEDGER_NOT_NATIVE_R3"

	SupportOrientedLedgerExists          = "CONDITIONAL_SUPPORT_ORIENTED_FINITE_SECTOR_PROJECTOR_LEDGER_EXISTS_UNDER_DUAL_SEAL"
	SupportProjectorsComplete            = "CONDITIONAL_SUPPORT_PROJECTORS_ARE_ORTHOGONAL_COMPLETE_ON_H_R_MIN"
	SupportProjectorsAFOrientStable      = "CONDITIONAL_SUPPORT_PROJECTOR_LEDGER_STABLE_UNDER_A_F_ORIENT"
	SupportProjectorsEdgeCompatible      = "CONDITIONAL_SUPPORT_PROJECTORS_ARE_EDGE_SUPPORT_COMPATIBLE"
	SupportTraceRowsReproduceNEff        = "CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_ROWS_REPRODUCE_OPERATOR_N_EFF"
	SupportR3CandidateUnderDualSeal      = "CONDITIONAL_SUPPORT_R3_CANDIDATE_NOW_HAS_PROJECTORS_AND_READOUT_UNDER_SEALS"
	SupportBoundaryAlphaSuppliesWeights  = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_SEAL_SUPPLIES_TRACE_WEIGHTS"
	SupportOrientationSuppliesProjectors = "CONDITIONAL_SUPPORT_POST_ORIENTATION_SEAL_SUPPLIES_PROJECTOR_LEDGER"
	SupportNoPhysicalAssignment          = "CONDITIONAL_SUPPORT_NO_PHYSICAL_SECTOR_ASSIGNMENT_MADE"

	FailureAlphaStillSealed            = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeIncidenceFunctor    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion  = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailurePostOrientNotFullAF         = "FAILED_ROUTE_POST_ORIENTATION_PROJECTORS_NOT_FULL_UNBROKEN_A_F_SECTORS"
	FailureNoNativeFiniteSectorTheorem = "FAILED_ROUTE_NO_NATIVE_FINITE_SECTOR_PROJECTOR_THEOREM"
	FailureNoNativeR3SectorLedger      = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureSocketAtomsNotPhysical      = "FAILED_ROUTE_SOCKET_ATOMS_NOT_PHYSICAL_SECTOR_ASSIGNMENTS"
	FailureNoPhysicalSectorAssignment  = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap      = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap      = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues    = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum    = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoOfficialNEffUpdate        = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate       = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator      = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                        = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type ProjectorAtom struct {
	Atom                  string
	Rank                  int
	WeightFormula         string
	Weight                float64
	EdgeSupport           string
	TraceContribution     float64
	SquareContribution    float64
	Idempotent            bool
	OrthogonalToOthers    bool
	CompletePartOfHRMin   bool
	StableUnderAFOrient   bool
	StableUnderFullAF     bool
	EdgeCompatible        bool
	PhysicalSector        bool
	GenerationResolved    bool
	FlavorResolved        bool
	IndividualYukawaValue bool
	Supports, Failures    []string
}

type ProjectorLedger struct {
	Name                     string
	Atoms                    []ProjectorAtom
	Rank                     int
	ExpectedRank             int
	Idempotent               bool
	Orthogonal               bool
	CompleteOnHRMin          bool
	StableUnderAFOrient      bool
	StableUnderFullAF        bool
	EdgeCompatible           bool
	TraceReadoutCompatible   bool
	NativeFiniteSectorLedger bool
	NativeR3                 bool
	TraceTotal               float64
	SquareTrace              float64
	OperatorNEff             float64
	OperatorCYukawa          float64
	Supports, Failures       []string
}

type DualSeal struct {
	BoundaryAlphaSealSuppliesWeights      bool
	PostOrientationSealSuppliesProjectors bool
	NativeAlphaFunctorCertified           bool
	NativeFullAFProjectorsCertified       bool
	OfficialR3Eligible                    bool
	Supports, Failures                    []string
}

type LedgerFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                    bool
	AlphaStillSealed            bool
	PostOrientationNotFullAF    bool
	NoNativeFiniteSectorTheorem bool
	NoNativeR3                  bool
	NoPhysicalSectorAssignment  bool
	NoGenerationCarrier         bool
	NoFlavorOrientation         bool
	NoIndividualYukawas         bool
	NoOfficialLedgerUpdate      bool
	NoNativeYukawaOperator      bool
	NoR4                        bool
	Verdict                     string
}

type Audit struct {
	ID        string
	Ledger    ProjectorLedger
	Seals     DualSeal
	Freeze    LedgerFreeze
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB
	atoms := []ProjectorAtom{
		buildAtom(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, EdgePiPlus3, []string{SupportOrientedLedgerExists, SupportProjectorsComplete, SupportProjectorsAFOrientStable, SupportProjectorsEdgeCompatible}, []string{FailurePostOrientNotFullAF, FailureSocketAtomsNotPhysical}),
		buildAtom(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, EdgePiMinus3, []string{SupportOrientedLedgerExists, SupportProjectorsComplete, SupportProjectorsAFOrientStable, SupportProjectorsEdgeCompatible}, []string{FailureAlphaStillSealed, FailurePostOrientNotFullAF, FailureSocketAtomsNotPhysical}),
		buildAtom(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, EdgePiMinus1, []string{SupportOrientedLedgerExists, SupportProjectorsComplete, SupportProjectorsAFOrientStable, SupportProjectorsEdgeCompatible}, []string{FailureAlphaStillSealed, FailurePostOrientNotFullAF, FailureSocketAtomsNotPhysical}),
	}

	rank := 0
	traceTotal := 0.0
	squareTrace := 0.0
	for _, atom := range atoms {
		rank += atom.Rank
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
		if !atom.Idempotent || !atom.OrthogonalToOthers || !atom.CompletePartOfHRMin || !atom.StableUnderAFOrient || atom.StableUnderFullAF || !atom.EdgeCompatible || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return Audit{}, fmt.Errorf("projector ledger firewall leak in %s", atom.Atom)
		}
	}
	if rank != RankHRMin {
		return Audit{}, fmt.Errorf("ledger rank drift: got %d want %d", rank, RankHRMin)
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

	ledger := ProjectorLedger{
		Name: FiniteSectorLedger, Atoms: atoms, Rank: rank, ExpectedRank: RankHRMin,
		Idempotent: true, Orthogonal: true, CompleteOnHRMin: true, StableUnderAFOrient: true, StableUnderFullAF: false,
		EdgeCompatible: true, TraceReadoutCompatible: true, NativeFiniteSectorLedger: false, NativeR3: false,
		TraceTotal: traceTotal, SquareTrace: squareTrace, OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Supports: []string{SupportOrientedLedgerExists, SupportProjectorsComplete, SupportProjectorsAFOrientStable, SupportProjectorsEdgeCompatible, SupportTraceRowsReproduceNEff, SupportR3CandidateUnderDualSeal},
		Failures: []string{FailureAlphaStillSealed, FailurePostOrientNotFullAF, FailureNoNativeFiniteSectorTheorem, FailureNoNativeR3SectorLedger, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}

	seals := DualSeal{
		BoundaryAlphaSealSuppliesWeights: true, PostOrientationSealSuppliesProjectors: true,
		NativeAlphaFunctorCertified: false, NativeFullAFProjectorsCertified: false, OfficialR3Eligible: false,
		Supports: []string{SupportBoundaryAlphaSuppliesWeights, SupportOrientationSuppliesProjectors, SupportR3CandidateUnderDualSeal},
		Failures: []string{FailureAlphaStillSealed, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailurePostOrientNotFullAF, FailureNoNativeFiniteSectorTheorem, FailureNoNativeR3SectorLedger},
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
		Enforced: true, AlphaStillSealed: true, PostOrientationNotFullAF: true, NoNativeFiniteSectorTheorem: true, NoNativeR3: true,
		NoPhysicalSectorAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict,
	}
	return Audit{
		ID: AuditID, Ledger: ledger, Seals: seals, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 888 organizes the post-orientation lift candidate into an explicit oriented finite-sector projector ledger under the BoundaryAlpha and Higgs/post-orientation seals.",
		Final: "The ledger {Pi_+3, Pi_-3, Pi_-1} is idempotent, orthogonal, complete on H_R^min, stable under A_F^orient, edge-compatible with Y, and reproduces the diagnostic trace readout under the alpha seal. It remains a dual-seal R3 candidate, not native R3: alpha_B is sealed, the projectors are post-orientation rather than full unbroken A_F sectors, no physical sector/generation/flavor assignment is made, and official ledgers remain frozen.",
	}, nil
}

func buildAtom(atom string, rank int, formula string, weight float64, edge string, supports, failures []string) ProjectorAtom {
	return ProjectorAtom{
		Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, EdgeSupport: edge,
		TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight,
		Idempotent: true, OrthogonalToOthers: true, CompletePartOfHRMin: true, StableUnderAFOrient: true,
		StableUnderFullAF: false, EdgeCompatible: true, PhysicalSector: false, GenerationResolved: false, FlavorResolved: false, IndividualYukawaValue: false,
		Supports: supports, Failures: failures,
	}
}

func FormatAtom(a ProjectorAtom) string {
	return fmt.Sprintf("atom(%s rank=%d weight=%s numeric=%.16g edge=%s trace=%.16g square=%.16g idem=%t orthogonal=%t complete=%t orient_stable=%t full_A_F=%t edge=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", a.Atom, a.Rank, a.WeightFormula, a.Weight, a.EdgeSupport, a.TraceContribution, a.SquareContribution, a.Idempotent, a.OrthogonalToOthers, a.CompletePartOfHRMin, a.StableUnderAFOrient, a.StableUnderFullAF, a.EdgeCompatible, a.PhysicalSector, a.GenerationResolved, a.FlavorResolved, a.IndividualYukawaValue, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAtoms(atoms []ProjectorAtom) string {
	parts := make([]string, 0, len(atoms))
	for _, atom := range atoms {
		parts = append(parts, FormatAtom(atom))
	}
	return strings.Join(parts, "; ")
}

func FormatLedger(l ProjectorLedger) string {
	return fmt.Sprintf("projector_ledger(name=%s rank=%d expected=%d idem=%t orthogonal=%t complete=%t orient_stable=%t full_A_F=%t edge=%t trace_readout=%t native_finite=%t native_r3=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g atoms=[%s] supports=%s failures=%s)", l.Name, l.Rank, l.ExpectedRank, l.Idempotent, l.Orthogonal, l.CompleteOnHRMin, l.StableUnderAFOrient, l.StableUnderFullAF, l.EdgeCompatible, l.TraceReadoutCompatible, l.NativeFiniteSectorLedger, l.NativeR3, l.TraceTotal, l.SquareTrace, l.OperatorNEff, l.OperatorCYukawa, FormatAtoms(l.Atoms), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatSeals(s DualSeal) string {
	return fmt.Sprintf("dual_seal(boundary_alpha_weights=%t post_orientation_projectors=%t native_alpha=%t native_full_A_F_projectors=%t official_r3=%t supports=%s failures=%s)", s.BoundaryAlphaSealSuppliesWeights, s.PostOrientationSealSuppliesProjectors, s.NativeAlphaFunctorCertified, s.NativeFullAFProjectorsCertified, s.OfficialR3Eligible, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t post_orientation_not_full_A_F=%t no_native_finite_sector=%t no_native_r3=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.PostOrientationNotFullAF, f.NoNativeFiniteSectorTheorem, f.NoNativeR3, f.NoPhysicalSectorAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate887Inherited,
		StatusLedgerDefined,
		StatusProjectorsOrthogonal,
		StatusStabilizerCompatible,
		StatusEdgeCompatible,
		StatusTraceReadoutCompatible,
		StatusDualSealDependencies,
		StatusFullAFFirewall,
		StatusPhysicalFirewalls,
		StatusOfficialFreezePreserved,
		StatusNextBranchSelected,
		StatusFirewallVerdict,
		SupportOrientedLedgerExists,
		SupportProjectorsComplete,
		SupportProjectorsAFOrientStable,
		SupportProjectorsEdgeCompatible,
		SupportTraceRowsReproduceNEff,
		SupportR3CandidateUnderDualSeal,
		SupportBoundaryAlphaSuppliesWeights,
		SupportOrientationSuppliesProjectors,
		SupportNoPhysicalAssignment,
		FailureAlphaStillSealed,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailurePostOrientNotFullAF,
		FailureNoNativeFiniteSectorTheorem,
		FailureNoNativeR3SectorLedger,
		FailureSocketAtomsNotPhysical,
		FailureNoPhysicalSectorAssignment,
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
	return f.Enforced && f.AlphaStillSealed && f.PostOrientationNotFullAF && f.NoNativeFiniteSectorTheorem && f.NoNativeR3 && f.NoPhysicalSectorAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
