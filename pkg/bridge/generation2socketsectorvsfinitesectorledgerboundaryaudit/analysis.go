// Package generation2socketsectorvsfinitesectorledgerboundaryaudit implements
// Gate 886: SocketSector vs FiniteSector Ledger Boundary Audit.
//
// Gate 886 follows Gate 885's typing of the active socket trace atoms in the
// post-orientation stabilizer layer. It audits the boundary between socket-sector
// trace atoms and true finite-sector trace atoms, asking what extra map is
// required to lift the R3-candidate ledger from A_F^orient socket support to a
// represented finite-sector ledger.
package generation2socketsectorvsfinitesectorledgerboundaryaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE886-SOCKET-SECTOR-VS-FINITE-SECTOR-LEDGER-BOUNDARY-AUDIT"

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
	FiniteSectorLedger     = "Pi_sector^F"
	MissingLiftMap         = "SocketSectorToFiniteSectorMap"
	MissingSigmaMap        = "Sigma_sector:{Pi_+3,Pi_-3,Pi_-1}->Pi_sector^F"

	AlphaSealName  = "BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL"
	Classification = "R2+++++_SOCKET_SECTOR_LEDGER_CANDIDATE_NO_FINITE_SECTOR_LIFT_NOT_R3"
	NextBranch     = "SOCKET_SECTOR_TO_FINITE_SECTOR_LEDGER_MAP_AUDIT"

	StatusGate885Inherited         = "PASS_GATE885_SOCKET_TRACE_ATOM_LEDGER_INHERITED"
	StatusSocketSectorStatus       = "PASS_SOCKET_SECTOR_STATUS_AUDITED"
	StatusFiniteSectorLiftAudited  = "PASS_FINITE_SECTOR_LIFT_REQUIREMENTS_AUDITED"
	StatusFullAFFirewall           = "PASS_FULL_UNBROKEN_A_F_SECTOR_FIREWALL_PRESERVED"
	StatusPhysicalSectorFirewall   = "PASS_PHYSICAL_SECTOR_ASSIGNMENT_FIREWALL_PRESERVED"
	StatusGenerationFlavorFirewall = "PASS_GENERATION_FLAVOR_FIREWALLS_PRESERVED"
	StatusOfficialFreezePreserved  = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNextBranchSelected       = "PASS_SOCKET_TO_FINITE_SECTOR_MAP_SELECTED_AS_NEXT_FRONTIER"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE886_SOCKET_SECTORS_NOT_NATIVE_FINITE_SECTORS"

	SupportSocketLedgerPostOrientationCandidate = "CONDITIONAL_SUPPORT_SOCKET_TRACE_ATOMS_FORM_POST_ORIENTATION_SECTOR_LEDGER_CANDIDATE"
	SupportSocketAtomsEdgeAndReadoutStable      = "CONDITIONAL_SUPPORT_SOCKET_ATOMS_ARE_EDGE_SUPPORT_AND_READOUT_STABLE_UNDER_A_F_ORIENT"
	SupportR3FrontierRequiresLiftMap            = "CONDITIONAL_SUPPORT_R3_FRONTIER_NOW_REQUIRES_SOCKET_TO_FINITE_SECTOR_MAP"
	SupportSocketSectorsTypedNotFinite          = "CONDITIONAL_SUPPORT_SOCKET_SECTORS_ARE_TYPED_BUT_NOT_FINITE_SECTORS"
	SupportBoundaryClarified                    = "CONDITIONAL_SUPPORT_BOUNDARY_BETWEEN_SOCKET_SECTOR_AND_FINITE_SECTOR_CLARIFIED"
	SupportNoPhysicalAssignment                 = "CONDITIONAL_SUPPORT_PHYSICAL_SECTOR_ASSIGNMENT_REMAINS_BLOCKED"

	FailureNoSocketToFiniteSectorMap    = "FAILED_ROUTE_NO_SOCKET_SECTOR_TO_FINITE_SECTOR_LEDGER_MAP"
	FailureNoFullUnbrokenAFSectorLedger = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_SECTOR_LEDGER"
	FailureAlphaStillSealed             = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureSocketLedgerNotNativeR3      = "FAILED_ROUTE_SOCKET_LEDGER_NOT_NATIVE_R3"
	FailureSocketAtomsNotStableFullAF   = "FAILED_ROUTE_SOCKET_ATOMS_NOT_STABLE_UNDER_FULL_UNBROKEN_A_F"
	FailurePostOrientNotNativeFinite    = "FAILED_ROUTE_POST_ORIENTATION_STABILIZER_SECTOR_NOT_NATIVE_FINITE_SECTOR"
	FailureNoPhysicalSectorAssignment   = "FAILED_ROUTE_NO_PHYSICAL_SECTOR_ASSIGNMENT"
	FailureNoGenerationCarrierMap       = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap       = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues     = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum     = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoOfficialNEffUpdate         = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate        = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator       = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                         = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type SocketSectorAtom struct {
	Atom                  string
	Rank                  int
	WeightFormula         string
	Weight                float64
	EdgeSupport           string
	TraceContribution     float64
	SquareContribution    float64
	SocketSectorTyped     bool
	FiniteSectorCertified bool
	StableInOrientLayer   bool
	StableInFullAF        bool
	PhysicalSector        bool
	GenerationResolved    bool
	FlavorResolved        bool
	IndividualYukawaValue bool
	Supports, Failures    []string
}

type SocketSectorLedger struct {
	Atoms                  []SocketSectorAtom
	PostOrientationAlgebra string
	FullUnbrokenAlgebra    string
	ActiveRank             int
	ExpectedRank           int
	Orthogonal             bool
	CompleteOnHRMin        bool
	Positive               bool
	SocketSectorTyped      bool
	FiniteSectorCertified  bool
	StableInOrientLayer    bool
	StableInFullAF         bool
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

type FiniteSectorLiftBoundary struct {
	MapName                  string
	Sigma                    string
	SourceAtoms              []string
	TargetLedger             string
	SourceSocketTyped        bool
	TargetCertified          bool
	LiftCertified            bool
	NativeR3                 bool
	FullUnbrokenAFCertified  bool
	PhysicalAssignment       bool
	GenerationCarrierPresent bool
	FlavorOrientationPresent bool
	NextRequiredObject       string
	Supports, Failures       []string
}

type LedgerFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type BoundaryFirewalls struct {
	Enforced                   bool
	AlphaStillSealed           bool
	NoSocketToFiniteSectorMap  bool
	NoFullUnbrokenAFLedger     bool
	NoPhysicalSectorAssignment bool
	NoGenerationCarrier        bool
	NoFlavorOrientation        bool
	NoIndividualYukawas        bool
	NoOfficialLedgerUpdate     bool
	NoNativeYukawaOperator     bool
	NoR4                       bool
	Verdict                    string
}

type Audit struct {
	ID        string
	AlphaSeal string
	Ledger    SocketSectorLedger
	Lift      FiniteSectorLiftBoundary
	Freeze    LedgerFreeze
	Firewalls BoundaryFirewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB
	atoms := []SocketSectorAtom{
		buildAtom(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, EdgePiPlus3, []string{SupportSocketLedgerPostOrientationCandidate, SupportSocketAtomsEdgeAndReadoutStable}, []string{FailureNoSocketToFiniteSectorMap, FailureNoPhysicalSectorAssignment, FailureSocketAtomsNotStableFullAF}),
		buildAtom(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, EdgePiMinus3, []string{SupportSocketLedgerPostOrientationCandidate, SupportSocketAtomsEdgeAndReadoutStable}, []string{FailureAlphaStillSealed, FailureNoSocketToFiniteSectorMap, FailureNoPhysicalSectorAssignment, FailureSocketAtomsNotStableFullAF}),
		buildAtom(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, EdgePiMinus1, []string{SupportSocketLedgerPostOrientationCandidate, SupportSocketAtomsEdgeAndReadoutStable}, []string{FailureAlphaStillSealed, FailureNoSocketToFiniteSectorMap, FailureNoPhysicalSectorAssignment, FailureSocketAtomsNotStableFullAF}),
	}

	activeRank := 0
	positive := true
	traceTotal := 0.0
	squareTrace := 0.0
	for _, atom := range atoms {
		activeRank += atom.Rank
		positive = positive && atom.Weight > 0
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
		if !atom.SocketSectorTyped || atom.FiniteSectorCertified || atom.StableInFullAF || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return Audit{}, fmt.Errorf("socket/finite/physical boundary leaked in %s", atom.Atom)
		}
		if !atom.StableInOrientLayer {
			return Audit{}, fmt.Errorf("post-orientation stability missing on %s", atom.Atom)
		}
	}
	if activeRank != RankHRMin {
		return Audit{}, fmt.Errorf("active rank drift: got %d want %d", activeRank, RankHRMin)
	}
	if !positive {
		return Audit{}, fmt.Errorf("non-positive socket-sector ledger")
	}
	expectedTrace := 3.0 + 3.0*AlphaB
	expectedSquare := 3.0 + 3.0*AlphaB*AlphaB - 6.0*AlphaB*AlphaB*AlphaB + 12.0*AlphaB*AlphaB*AlphaB*AlphaB
	if !near(traceTotal, expectedTrace) || !near(squareTrace, expectedSquare) {
		return Audit{}, fmt.Errorf("trace ledger drift: trace %.18g square %.18g", traceTotal, squareTrace)
	}
	operatorNEff := traceTotal * traceTotal / squareTrace
	operatorCYukawa := 3.0 / operatorNEff
	if !near(operatorNEff, OperatorNEffDiagnostic) || !near(operatorCYukawa, OperatorCYukawaDiagnostic) {
		return Audit{}, fmt.Errorf("operator diagnostics drift: N_eff %.18g C_Yukawa %.18g", operatorNEff, operatorCYukawa)
	}

	ledger := SocketSectorLedger{
		Atoms: atoms, PostOrientationAlgebra: PostOrientationAlgebra, FullUnbrokenAlgebra: FullUnbrokenAlgebra,
		ActiveRank: activeRank, ExpectedRank: RankHRMin, Orthogonal: true, CompleteOnHRMin: true, Positive: positive,
		SocketSectorTyped: true, FiniteSectorCertified: false, StableInOrientLayer: true, StableInFullAF: false,
		PhysicalSectors: false, NativeR3: false, R3CandidateUnderSeal: true,
		TraceTotal: traceTotal, SquareTrace: squareTrace, OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Classification: Classification, NextBranch: NextBranch,
		Supports: []string{SupportSocketLedgerPostOrientationCandidate, SupportSocketAtomsEdgeAndReadoutStable, SupportR3FrontierRequiresLiftMap, SupportSocketSectorsTypedNotFinite, SupportBoundaryClarified, SupportNoPhysicalAssignment},
		Failures: []string{FailureNoSocketToFiniteSectorMap, FailureNoFullUnbrokenAFSectorLedger, FailureAlphaStillSealed, FailureSocketLedgerNotNativeR3, FailureSocketAtomsNotStableFullAF, FailurePostOrientNotNativeFinite, FailureNoPhysicalSectorAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4},
	}
	lift := FiniteSectorLiftBoundary{
		MapName: MissingLiftMap, Sigma: MissingSigmaMap,
		SourceAtoms: []string{AtomPiPlus3, AtomPiMinus3, AtomPiMinus1}, TargetLedger: FiniteSectorLedger,
		SourceSocketTyped: true, TargetCertified: false, LiftCertified: false, NativeR3: false,
		FullUnbrokenAFCertified: false, PhysicalAssignment: false, GenerationCarrierPresent: false, FlavorOrientationPresent: false,
		NextRequiredObject: MissingSigmaMap,
		Supports:           []string{SupportR3FrontierRequiresLiftMap, SupportSocketSectorsTypedNotFinite, SupportBoundaryClarified},
		Failures:           []string{FailureNoSocketToFiniteSectorMap, FailureNoFullUnbrokenAFSectorLedger, FailurePostOrientNotNativeFinite, FailureNoPhysicalSectorAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}
	freeze := LedgerFreeze{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreezePreserved}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
	firewalls := BoundaryFirewalls{
		Enforced: true, AlphaStillSealed: true, NoSocketToFiniteSectorMap: true, NoFullUnbrokenAFLedger: true,
		NoPhysicalSectorAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict,
	}
	return Audit{
		ID: AuditID, AlphaSeal: AlphaSealName, Ledger: ledger, Lift: lift, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 886 audits the boundary between post-orientation socket-sector trace atoms and true represented finite-sector trace atoms.",
		Final: "Pi_+3, Pi_-3, and Pi_-1 are typed post-orientation socket-sector trace atoms and edge-support/readout-stable candidates, but no SocketSectorToFiniteSectorMap is certified; they remain below native R3, physical sector assignment, generation/flavor splitting, and official ledger update.",
	}, nil
}

func buildAtom(atom string, rank int, formula string, weight float64, edge string, supports, failures []string) SocketSectorAtom {
	return SocketSectorAtom{
		Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, EdgeSupport: edge,
		TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight,
		SocketSectorTyped: true, FiniteSectorCertified: false, StableInOrientLayer: true, StableInFullAF: false,
		PhysicalSector: false, GenerationResolved: false, FlavorResolved: false, IndividualYukawaValue: false,
		Supports: supports, Failures: failures,
	}
}

func FormatAtom(a SocketSectorAtom) string {
	return fmt.Sprintf("atom(%s rank=%d weight=%s numeric=%.16g edge=%s trace=%.16g square=%.16g socket_typed=%t finite_sector=%t orient_stable=%t full_A_F_stable=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", a.Atom, a.Rank, a.WeightFormula, a.Weight, a.EdgeSupport, a.TraceContribution, a.SquareContribution, a.SocketSectorTyped, a.FiniteSectorCertified, a.StableInOrientLayer, a.StableInFullAF, a.PhysicalSector, a.GenerationResolved, a.FlavorResolved, a.IndividualYukawaValue, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAtoms(atoms []SocketSectorAtom) string {
	parts := make([]string, 0, len(atoms))
	for _, atom := range atoms {
		parts = append(parts, FormatAtom(atom))
	}
	return strings.Join(parts, "; ")
}

func FormatLedger(l SocketSectorLedger) string {
	return fmt.Sprintf("socket_sector_ledger(classification=%s orient_algebra=%s full_algebra=%s active_rank=%d expected_rank=%d orthogonal=%t complete=%t positive=%t socket_typed=%t finite_sector_certified=%t orient_stable=%t full_A_F_stable=%t physical_sectors=%t native_r3=%t r3_candidate_under_seal=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g next=%s atoms=[%s] supports=%s failures=%s)", l.Classification, l.PostOrientationAlgebra, l.FullUnbrokenAlgebra, l.ActiveRank, l.ExpectedRank, l.Orthogonal, l.CompleteOnHRMin, l.Positive, l.SocketSectorTyped, l.FiniteSectorCertified, l.StableInOrientLayer, l.StableInFullAF, l.PhysicalSectors, l.NativeR3, l.R3CandidateUnderSeal, l.TraceTotal, l.SquareTrace, l.OperatorNEff, l.OperatorCYukawa, l.NextBranch, FormatAtoms(l.Atoms), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatLift(l FiniteSectorLiftBoundary) string {
	return fmt.Sprintf("finite_sector_lift(map=%s sigma=%s target=%s source_typed=%t target_certified=%t lift_certified=%t native_r3=%t full_A_F=%t physical=%t generation=%t flavor=%t next=%s source_atoms=%s supports=%s failures=%s)", l.MapName, l.Sigma, l.TargetLedger, l.SourceSocketTyped, l.TargetCertified, l.LiftCertified, l.NativeR3, l.FullUnbrokenAFCertified, l.PhysicalAssignment, l.GenerationCarrierPresent, l.FlavorOrientationPresent, l.NextRequiredObject, strings.Join(l.SourceAtoms, ","), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f BoundaryFirewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t no_socket_to_finite=%t no_full_A_F=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.NoSocketToFiniteSectorMap, f.NoFullUnbrokenAFLedger, f.NoPhysicalSectorAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate885Inherited,
		StatusSocketSectorStatus,
		StatusFiniteSectorLiftAudited,
		StatusFullAFFirewall,
		StatusPhysicalSectorFirewall,
		StatusGenerationFlavorFirewall,
		StatusOfficialFreezePreserved,
		StatusNextBranchSelected,
		StatusFirewallVerdict,
		SupportSocketLedgerPostOrientationCandidate,
		SupportSocketAtomsEdgeAndReadoutStable,
		SupportR3FrontierRequiresLiftMap,
		SupportSocketSectorsTypedNotFinite,
		SupportBoundaryClarified,
		SupportNoPhysicalAssignment,
		FailureNoSocketToFiniteSectorMap,
		FailureNoFullUnbrokenAFSectorLedger,
		FailureAlphaStillSealed,
		FailureSocketLedgerNotNativeR3,
		FailureSocketAtomsNotStableFullAF,
		FailurePostOrientNotNativeFinite,
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

func firewallsOK(f BoundaryFirewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.NoSocketToFiniteSectorMap && f.NoFullUnbrokenAFLedger && f.NoPhysicalSectorAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
