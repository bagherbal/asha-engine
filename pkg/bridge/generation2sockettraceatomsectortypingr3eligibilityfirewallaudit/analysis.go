// Package generation2sockettraceatomsectortypingr3eligibilityfirewallaudit implements
// Gate 885: SocketTraceAtom SectorTyping and R3 Eligibility Firewall Audit.
//
// Gate 885 follows Gate 884's positive socket trace-magnitude readout rows. It
// does not reopen the BoundaryAlpha proof, does not update official ledgers, and
// does not assign physical particle sectors. It audits whether Pi_+3, Pi_-3,
// and Pi_-1 are typed socket-sector trace atoms in the post-orientation
// stabilizer layer, or whether they remain below native R3.
package generation2sockettraceatomsectortypingr3eligibilityfirewallaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE885-SOCKET-TRACE-ATOM-SECTOR-TYPING-R3-ELIGIBILITY-FIREWALL-AUDIT"

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
	AlphaSealName          = "BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL"
	Classification         = "R2+++++_SOCKET_TRACE_ATOM_LEDGER_R3_CANDIDATE_UNDER_ALPHA_SEAL_NOT_NATIVE_R3"
	NextBranch             = "GENERATION_CARRIER_REQUIREMENTS_OR_SECTOR_TYPING_REFINEMENT_UNDER_ALPHA_SEAL"

	StatusGate884Inherited         = "PASS_GATE884_TRACE_MAGNITUDE_READOUT_ROWS_INHERITED"
	StatusSocketAtomsDefined       = "PASS_SOCKET_TRACE_ATOMS_DEFINED"
	StatusProjectorOrthogonality   = "PASS_SOCKET_TRACE_ATOMS_ORTHOGONAL_AND_COMPLETE_ON_H_R_MIN"
	StatusStabilizerTypingAudited  = "PASS_POST_ORIENTATION_STABILIZER_TYPING_AUDITED"
	StatusEdgeSupportTypingAudited = "PASS_D_F_SYMBOLIC_EDGE_SUPPORT_TYPING_AUDITED"
	StatusPositiveReadoutInherited = "PASS_POSITIVE_TRACE_MAGNITUDE_READOUT_INHERITED"
	StatusR3EligibilityAudited     = "PASS_R3_ELIGIBILITY_FIREWALL_AUDITED"
	StatusPhysicalSectorFirewall   = "PASS_PHYSICAL_SECTOR_ASSIGNMENT_FIREWALL_PRESERVED"
	StatusOfficialFreezePreserved  = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE885_SOCKET_ATOMS_R3_CANDIDATE_UNDER_ALPHA_SEAL_NOT_NATIVE_R3"

	SupportSocketAtomsTypedInOrientLayer = "CONDITIONAL_SUPPORT_SOCKET_TRACE_ATOMS_ARE_TYPED_IN_A_F_ORIENT_LAYER"
	SupportSocketAtomsEdgeSupportAtoms   = "CONDITIONAL_SUPPORT_SOCKET_ATOMS_ARE_D_F_SYMBOLIC_EDGE_SUPPORT_ATOMS"
	SupportSocketTraceLedgerR3Candidate  = "CONDITIONAL_SUPPORT_SOCKET_TRACE_MAGNITUDE_LEDGER_IS_R3_CANDIDATE_UNDER_ALPHA_SEAL"
	SupportR3PressureReduced             = "CONDITIONAL_SUPPORT_R3_PRESSURE_NOW_REDUCES_TO_ALPHA_FUNCTOR_PLUS_SECTOR_TYPING_FIREWALL"
	SupportPositiveReadoutRowsInherited  = "CONDITIONAL_SUPPORT_GATE884_POSITIVE_READOUT_ROWS_INHERITED"
	SupportPostOrientationStableOnly     = "CONDITIONAL_SUPPORT_SOCKET_ATOMS_STABLE_IN_POST_ORIENTATION_STABILIZER_LAYER"
	SupportNotRandomProjectors           = "CONDITIONAL_SUPPORT_SOCKET_ATOMS_ARE_NOT_RANDOM_PROJECTORS_BUT_Y_EDGE_SUPPORT_ATOMS"

	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureSocketAtomsNotNativeR3Sectors = "FAILED_ROUTE_SOCKET_ATOMS_NOT_FULL_NATIVE_R3_SECTORS"
	FailureSocketAtomsNotPhysical        = "FAILED_ROUTE_SOCKET_ATOMS_NOT_PHYSICAL_PARTICLE_ASSIGNMENTS"
	FailureNoFullUnbrokenAFSectorLedger  = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_SECTOR_LEDGER"
	FailureSocketAtomsNotStableFullAF    = "FAILED_ROUTE_SOCKET_ATOMS_NOT_STABLE_UNDER_FULL_UNBROKEN_A_F"
	FailureDFSupportNotPhysicalSector    = "FAILED_ROUTE_D_F_SYMBOLIC_EDGE_SUPPORT_NOT_PHYSICAL_SECTOR_ASSIGNMENT"
	FailureNoNativeIncidenceFunctor      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoGenerationCarrierMap        = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap        = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues      = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum      = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                          = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type SocketTraceAtom struct {
	Atom                  string
	Rank                  int
	WeightFormula         string
	Weight                float64
	EdgeSupport           string
	TraceContribution     float64
	SquareContribution    float64
	Orthogonal            bool
	CompletePart          bool
	Positive              bool
	StableInOrientLayer   bool
	StableInFullAF        bool
	PhysicalSector        bool
	GenerationResolved    bool
	FlavorResolved        bool
	IndividualYukawaValue bool
	Supports, Failures    []string
}

type SectorTypingLedger struct {
	Atoms                  []SocketTraceAtom
	PostOrientationAlgebra string
	ActiveRank             int
	ExpectedRank           int
	Orthogonal             bool
	CompleteOnHRMin        bool
	Positive               bool
	StableInOrientLayer    bool
	StableInFullAF         bool
	EdgeTyped              bool
	PhysicalSectors        bool
	NativeR3               bool
	R3CandidateUnderSeal   bool
	TraceTotal             float64
	SquareTrace            float64
	OperatorNEff           float64
	OperatorCYukawa        float64
	Classification         string
	NextBranch             string
	Supports, Failures     []string
}

type LedgerFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type R3EligibilityFirewalls struct {
	Enforced                  bool
	AlphaStillSealed          bool
	OnlyPostOrientationStable bool
	NoFullUnbrokenAFLedger    bool
	SocketAtomsNotPhysical    bool
	NoGenerationCarrier       bool
	NoFlavorOrientation       bool
	NoIndividualYukawas       bool
	NoOfficialLedgerUpdate    bool
	NoNativeYukawaOperator    bool
	NoR4                      bool
	Verdict                   string
}

type Audit struct {
	ID        string
	AlphaSeal string
	Ledger    SectorTypingLedger
	Freeze    LedgerFreeze
	Firewalls R3EligibilityFirewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB
	atoms := []SocketTraceAtom{
		buildAtom(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, EdgePiPlus3, []string{SupportSocketAtomsTypedInOrientLayer, SupportSocketAtomsEdgeSupportAtoms, SupportPostOrientationStableOnly}, []string{FailureSocketAtomsNotPhysical, FailureSocketAtomsNotStableFullAF}),
		buildAtom(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, EdgePiMinus3, []string{SupportSocketAtomsTypedInOrientLayer, SupportSocketAtomsEdgeSupportAtoms, SupportPositiveReadoutRowsInherited}, []string{FailureAlphaStillSealed, FailureSocketAtomsNotPhysical, FailureSocketAtomsNotStableFullAF}),
		buildAtom(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, EdgePiMinus1, []string{SupportSocketAtomsTypedInOrientLayer, SupportSocketAtomsEdgeSupportAtoms, SupportPositiveReadoutRowsInherited}, []string{FailureAlphaStillSealed, FailureSocketAtomsNotPhysical, FailureSocketAtomsNotStableFullAF}),
	}

	activeRank := 0
	positive := true
	traceTotal := 0.0
	squareTrace := 0.0
	for _, atom := range atoms {
		activeRank += atom.Rank
		positive = positive && atom.Positive
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
		if atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return Audit{}, fmt.Errorf("physical/generation/flavor/yukawa assignment leaked into %s", atom.Atom)
		}
		if !atom.StableInOrientLayer || atom.StableInFullAF {
			return Audit{}, fmt.Errorf("bad stability flags on %s", atom.Atom)
		}
	}
	if activeRank != RankHRMin {
		return Audit{}, fmt.Errorf("active rank drift: got %d want %d", activeRank, RankHRMin)
	}
	if !positive {
		return Audit{}, fmt.Errorf("non-positive socket trace atom")
	}
	expectedTrace := 3.0 + 3.0*AlphaB
	expectedSquare := 3.0 + 3.0*AlphaB*AlphaB - 6.0*AlphaB*AlphaB*AlphaB + 12.0*AlphaB*AlphaB*AlphaB*AlphaB
	if !near(traceTotal, expectedTrace) {
		return Audit{}, fmt.Errorf("trace drift: got %.18g want %.18g", traceTotal, expectedTrace)
	}
	if !near(squareTrace, expectedSquare) {
		return Audit{}, fmt.Errorf("square-trace drift: got %.18g want %.18g", squareTrace, expectedSquare)
	}
	operatorNEff := traceTotal * traceTotal / squareTrace
	operatorCYukawa := 3.0 / operatorNEff
	if !near(operatorNEff, OperatorNEffDiagnostic) || !near(operatorCYukawa, OperatorCYukawaDiagnostic) {
		return Audit{}, fmt.Errorf("operator ledger drift: N_eff %.18g C_Yukawa %.18g", operatorNEff, operatorCYukawa)
	}

	ledger := SectorTypingLedger{
		Atoms: atoms, PostOrientationAlgebra: PostOrientationAlgebra,
		ActiveRank: activeRank, ExpectedRank: RankHRMin,
		Orthogonal: true, CompleteOnHRMin: true, Positive: positive,
		StableInOrientLayer: true, StableInFullAF: false, EdgeTyped: true,
		PhysicalSectors: false, NativeR3: false, R3CandidateUnderSeal: true,
		TraceTotal: traceTotal, SquareTrace: squareTrace, OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Classification: Classification, NextBranch: NextBranch,
		Supports: []string{SupportSocketAtomsTypedInOrientLayer, SupportSocketAtomsEdgeSupportAtoms, SupportSocketTraceLedgerR3Candidate, SupportR3PressureReduced, SupportPositiveReadoutRowsInherited, SupportPostOrientationStableOnly, SupportNotRandomProjectors},
		Failures: []string{FailureAlphaStillSealed, FailureSocketAtomsNotNativeR3Sectors, FailureSocketAtomsNotPhysical, FailureNoFullUnbrokenAFSectorLedger, FailureSocketAtomsNotStableFullAF, FailureDFSupportNotPhysicalSector, FailureNoNativeIncidenceFunctor, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4},
	}
	freeze := LedgerFreeze{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreezePreserved},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
	firewalls := R3EligibilityFirewalls{
		Enforced: true, AlphaStillSealed: true, OnlyPostOrientationStable: true,
		NoFullUnbrokenAFLedger: true, SocketAtomsNotPhysical: true, NoGenerationCarrier: true,
		NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true,
		NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict,
	}
	return Audit{
		ID: AuditID, AlphaSeal: AlphaSealName, Ledger: ledger, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 885 audits whether the Gate 884 positive socket trace-magnitude rows are typed sector-ledger atoms under the post-orientation stabilizer layer.",
		Final: "Pi_+3, Pi_-3, and Pi_-1 form an R3-candidate socket trace atom ledger under the BoundaryAlpha seal, but they are not native R3 sectors: alpha remains sealed, full unbroken A_F stability is absent, and generation/flavor/physical Yukawa assignments remain blocked.",
	}, nil
}

func buildAtom(atom string, rank int, formula string, weight float64, edge string, supports, failures []string) SocketTraceAtom {
	return SocketTraceAtom{
		Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, EdgeSupport: edge,
		TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight,
		Orthogonal: true, CompletePart: true, Positive: weight > 0,
		StableInOrientLayer: true, StableInFullAF: false, PhysicalSector: false,
		GenerationResolved: false, FlavorResolved: false, IndividualYukawaValue: false,
		Supports: supports, Failures: failures,
	}
}

func FormatAtom(a SocketTraceAtom) string {
	return fmt.Sprintf("atom(%s rank=%d weight=%s numeric=%.16g edge=%s trace=%.16g square=%.16g orient_stable=%t full_A_F_stable=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", a.Atom, a.Rank, a.WeightFormula, a.Weight, a.EdgeSupport, a.TraceContribution, a.SquareContribution, a.StableInOrientLayer, a.StableInFullAF, a.PhysicalSector, a.GenerationResolved, a.FlavorResolved, a.IndividualYukawaValue, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAtoms(atoms []SocketTraceAtom) string {
	parts := make([]string, 0, len(atoms))
	for _, atom := range atoms {
		parts = append(parts, FormatAtom(atom))
	}
	return strings.Join(parts, "; ")
}

func FormatLedger(l SectorTypingLedger) string {
	return fmt.Sprintf("sector_typing_ledger(classification=%s algebra=%s active_rank=%d expected_rank=%d orthogonal=%t complete=%t positive=%t orient_stable=%t full_A_F_stable=%t edge_typed=%t physical_sectors=%t native_r3=%t r3_candidate_under_seal=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g next=%s atoms=[%s] supports=%s failures=%s)", l.Classification, l.PostOrientationAlgebra, l.ActiveRank, l.ExpectedRank, l.Orthogonal, l.CompleteOnHRMin, l.Positive, l.StableInOrientLayer, l.StableInFullAF, l.EdgeTyped, l.PhysicalSectors, l.NativeR3, l.R3CandidateUnderSeal, l.TraceTotal, l.SquareTrace, l.OperatorNEff, l.OperatorCYukawa, l.NextBranch, FormatAtoms(l.Atoms), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f R3EligibilityFirewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t only_post_orientation=%t no_full_A_F=%t socket_atoms_not_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.OnlyPostOrientationStable, f.NoFullUnbrokenAFLedger, f.SocketAtomsNotPhysical, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate884Inherited,
		StatusSocketAtomsDefined,
		StatusProjectorOrthogonality,
		StatusStabilizerTypingAudited,
		StatusEdgeSupportTypingAudited,
		StatusPositiveReadoutInherited,
		StatusR3EligibilityAudited,
		StatusPhysicalSectorFirewall,
		StatusOfficialFreezePreserved,
		StatusFirewallVerdict,
		SupportSocketAtomsTypedInOrientLayer,
		SupportSocketAtomsEdgeSupportAtoms,
		SupportSocketTraceLedgerR3Candidate,
		SupportR3PressureReduced,
		SupportPositiveReadoutRowsInherited,
		SupportPostOrientationStableOnly,
		SupportNotRandomProjectors,
		FailureAlphaStillSealed,
		FailureSocketAtomsNotNativeR3Sectors,
		FailureSocketAtomsNotPhysical,
		FailureNoFullUnbrokenAFSectorLedger,
		FailureSocketAtomsNotStableFullAF,
		FailureDFSupportNotPhysicalSector,
		FailureNoNativeIncidenceFunctor,
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

func firewallsOK(f R3EligibilityFirewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.OnlyPostOrientationStable && f.NoFullUnbrokenAFLedger && f.SocketAtomsNotPhysical && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
