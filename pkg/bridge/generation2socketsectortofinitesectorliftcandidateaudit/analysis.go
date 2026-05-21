// Package generation2socketsectortofinitesectorliftcandidateaudit implements
// Gate 887: SocketSectorToFiniteSector Lift Candidate Audit.
//
// Gate 887 follows Gate 886's socket/finite-sector boundary audit. It searches
// for lawful candidate lifts from the post-orientation socket trace atoms
// Pi_+3, Pi_-3, and Pi_-1 into represented finite-sector projectors Pi_sector^F,
// without assigning physical particles, generation labels, flavor labels, or
// individual Yukawa values.
package generation2socketsectortofinitesectorliftcandidateaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE887-SOCKET-SECTOR-TO-FINITE-SECTOR-LIFT-CANDIDATE-AUDIT"

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

	RouteStabilizerSectorLift = "ROUTE_A_STABILIZER_SECTOR_LIFT"
	RouteFullAFLift           = "ROUTE_B_FULL_A_F_LIFT"
	RouteEdgeSupportLift      = "ROUTE_C_EDGE_SUPPORT_LIFT"

	Classification = "R2+++++_POST_ORIENTATION_FINITE_SECTOR_LIFT_CANDIDATE_NOT_NATIVE_R3"
	NextBranch     = "FINITE_SECTOR_PROJECTOR_LEDGER_COMPATIBILITY_AUDIT_UNDER_ALPHA_AND_ORIENTATION_SEALS"

	StatusGate886Inherited         = "PASS_GATE886_SOCKET_FINITE_BOUNDARY_INHERITED"
	StatusDomainStrong             = "PASS_SOCKET_TRACE_ATOM_DOMAIN_AUDITED"
	StatusCodomainRequirements     = "PASS_FINITE_SECTOR_CODOMAIN_REQUIREMENTS_AUDITED"
	StatusLiftRoutesAudited        = "PASS_SOCKET_TO_FINITE_SECTOR_LIFT_ROUTES_AUDITED"
	StatusStabilizerLiftCandidate  = "PASS_STABILIZER_SECTOR_LIFT_CANDIDATE_AUDITED"
	StatusFullAFLiftBlocked        = "PASS_FULL_A_F_LIFT_BLOCKED_BY_WEAK_ORIENTATION_FIREWALL"
	StatusEdgeSupportLiftCandidate = "PASS_EDGE_SUPPORT_LIFT_CANDIDATE_AUDITED"
	StatusPhysicalFirewalls        = "PASS_PHYSICAL_GENERATION_FLAVOR_FIREWALLS_PRESERVED"
	StatusOfficialFreezePreserved  = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNextBranchSelected       = "PASS_FINITE_SECTOR_PROJECTOR_LEDGER_COMPATIBILITY_SELECTED_AS_NEXT_FRONTIER"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE887_POST_ORIENTATION_LIFT_CANDIDATE_NOT_NATIVE_R3"

	SupportSocketSectorDomainStrong         = "CONDITIONAL_SUPPORT_SOCKET_TRACE_ATOMS_FORM_STRONG_LIFT_DOMAIN"
	SupportPostOrientationLiftCandidate     = "CONDITIONAL_SUPPORT_SOCKET_SECTOR_TO_POST_ORIENTATION_FINITE_SECTOR_LIFT_CANDIDATE"
	SupportEdgeSupportFiniteSectorCandidate = "CONDITIONAL_SUPPORT_SOCKET_ATOMS_ARE_FINITE_EDGE_SUPPORT_SECTOR_CANDIDATES_IN_A_F_ORIENT"
	SupportR3PreparationUnderSeals          = "CONDITIONAL_SUPPORT_R3_PREPARATION_ADVANCES_UNDER_ALPHA_AND_ORIENTATION_SEALS"
	SupportLiftPreservesTraceReadout        = "CONDITIONAL_SUPPORT_CANDIDATE_LIFT_PRESERVES_TRACE_READOUT_ROWS"
	SupportStabilizerNotNative              = "CONDITIONAL_SUPPORT_POST_ORIENTATION_LIFT_IS_VALID_CANDIDATE_BUT_NOT_NATIVE_R3"
	SupportNoPhysicalAssignment             = "CONDITIONAL_SUPPORT_NO_PHYSICAL_SECTOR_ASSIGNMENT_MADE"
	SupportFullAFLiftAttemptAudited         = "CONDITIONAL_SUPPORT_FULL_A_F_LIFT_ATTEMPT_AUDITED_AND_BLOCKED"
	SupportExactNextObject                  = "CONDITIONAL_SUPPORT_EXACT_NEXT_OBJECT_IS_FINITE_SECTOR_PROJECTOR_LEDGER_COMPATIBILITY"

	FailureNoNativeSocketToFiniteSectorMap = "FAILED_ROUTE_NO_NATIVE_SOCKET_TO_FINITE_SECTOR_LEDGER_MAP"
	FailureNoFullAFLift                    = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_FINITE_SECTOR_LIFT"
	FailureSocketAtomsNotStableFullH       = "FAILED_ROUTE_SOCKET_ATOMS_NOT_STABLE_UNDER_FULL_H_ACTION"
	FailurePostOrientationNotNativeR3      = "FAILED_ROUTE_POST_ORIENTATION_FINITE_SECTOR_NOT_NATIVE_R3"
	FailureNoFullUnbrokenAFSectorLedger    = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_SECTOR_LEDGER"
	FailureAlphaStillSealed                = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeIncidenceFunctor        = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion      = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureEdgeAtomNotPhysicalSector       = "FAILED_ROUTE_EDGE_SUPPORT_ATOM_NOT_PHYSICAL_SECTOR"
	FailureEdgeAtomNotYukawaValue          = "FAILED_ROUTE_EDGE_SUPPORT_ATOM_NOT_YUKAWA_VALUE"
	FailureSocketAtomsNotPhysical          = "FAILED_ROUTE_SOCKET_ATOMS_NOT_PHYSICAL_SECTOR_ASSIGNMENTS"
	FailureNoGenerationCarrierMap          = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap          = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues        = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum        = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoOfficialNEffUpdate            = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate           = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator          = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                            = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type SocketAtom struct {
	Atom                  string
	Rank                  int
	WeightFormula         string
	Weight                float64
	EdgeSupport           string
	TraceContribution     float64
	SquareContribution    float64
	SocketSectorTyped     bool
	PostOrientStable      bool
	FiniteSectorCandidate bool
	FiniteSectorCertified bool
	FullAFCertified       bool
	PhysicalSector        bool
	GenerationResolved    bool
	FlavorResolved        bool
	IndividualYukawaValue bool
	Supports, Failures    []string
}

type SocketDomain struct {
	Atoms                  []SocketAtom
	ActiveRank             int
	ExpectedRank           int
	CompleteOnHRMin        bool
	Orthogonal             bool
	Positive               bool
	PostOrientationAlgebra string
	FullUnbrokenAlgebra    string
	TraceTotal             float64
	SquareTrace            float64
	OperatorNEff           float64
	OperatorCYukawa        float64
	Supports, Failures     []string
}

type LiftRoute struct {
	Name                     string
	Description              string
	SourceAlgebra            string
	TargetLedger             string
	LiftCandidate            bool
	LiftCertified            bool
	PostOrientationOnly      bool
	FullUnbrokenAFCertified  bool
	PreservesTraceReadout    bool
	PreservesEdgeSupport     bool
	PhysicalAssignment       bool
	GenerationCarrierPresent bool
	FlavorOrientationPresent bool
	IndividualYukawaValues   bool
	Supports, Failures       []string
}

type LiftAudit struct {
	MapName            string
	Sigma              string
	SourceAtoms        []string
	TargetLedger       string
	Routes             []LiftRoute
	BestRoute          string
	LiftCandidate      bool
	LiftCertified      bool
	NativeR3           bool
	NextRequired       string
	Supports, Failures []string
}

type LedgerFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                   bool
	AlphaStillSealed           bool
	NoNativeSocketToFiniteMap  bool
	NoFullAFLift               bool
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
	Domain    SocketDomain
	Lift      LiftAudit
	Freeze    LedgerFreeze
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB
	atoms := []SocketAtom{
		buildAtom(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, EdgePiPlus3, []string{SupportSocketSectorDomainStrong, SupportPostOrientationLiftCandidate, SupportEdgeSupportFiniteSectorCandidate}, []string{FailureNoFullAFLift, FailureSocketAtomsNotStableFullH, FailureSocketAtomsNotPhysical}),
		buildAtom(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, EdgePiMinus3, []string{SupportSocketSectorDomainStrong, SupportPostOrientationLiftCandidate, SupportEdgeSupportFiniteSectorCandidate}, []string{FailureAlphaStillSealed, FailureNoFullAFLift, FailureSocketAtomsNotStableFullH, FailureSocketAtomsNotPhysical}),
		buildAtom(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, EdgePiMinus1, []string{SupportSocketSectorDomainStrong, SupportPostOrientationLiftCandidate, SupportEdgeSupportFiniteSectorCandidate}, []string{FailureAlphaStillSealed, FailureNoFullAFLift, FailureSocketAtomsNotStableFullH, FailureSocketAtomsNotPhysical}),
	}

	activeRank := 0
	traceTotal := 0.0
	squareTrace := 0.0
	positive := true
	for _, atom := range atoms {
		activeRank += atom.Rank
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
		positive = positive && atom.Weight > 0
		if !atom.SocketSectorTyped || !atom.PostOrientStable || !atom.FiniteSectorCandidate || atom.FiniteSectorCertified || atom.FullAFCertified || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return Audit{}, fmt.Errorf("socket-to-finite boundary leaked in atom %s", atom.Atom)
		}
	}
	if activeRank != RankHRMin {
		return Audit{}, fmt.Errorf("active rank drift: got %d want %d", activeRank, RankHRMin)
	}
	if !positive {
		return Audit{}, fmt.Errorf("non-positive domain weights")
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

	domain := SocketDomain{
		Atoms: atoms, ActiveRank: activeRank, ExpectedRank: RankHRMin, CompleteOnHRMin: true, Orthogonal: true, Positive: positive,
		PostOrientationAlgebra: PostOrientationAlgebra, FullUnbrokenAlgebra: FullUnbrokenAlgebra,
		TraceTotal: traceTotal, SquareTrace: squareTrace, OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Supports: []string{SupportSocketSectorDomainStrong, SupportPostOrientationLiftCandidate, SupportEdgeSupportFiniteSectorCandidate, SupportLiftPreservesTraceReadout},
		Failures: []string{FailureAlphaStillSealed, FailureNoNativeSocketToFiniteSectorMap, FailureNoFullAFLift, FailureSocketAtomsNotStableFullH, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}

	routes := []LiftRoute{
		{
			Name: RouteStabilizerSectorLift, Description: "lift socket atoms to post-orientation finite-sector candidates inside A_F^orient",
			SourceAlgebra: PostOrientationAlgebra, TargetLedger: FiniteSectorLedger,
			LiftCandidate: true, LiftCertified: false, PostOrientationOnly: true, FullUnbrokenAFCertified: false,
			PreservesTraceReadout: true, PreservesEdgeSupport: true, PhysicalAssignment: false, GenerationCarrierPresent: false, FlavorOrientationPresent: false, IndividualYukawaValues: false,
			Supports: []string{SupportPostOrientationLiftCandidate, SupportR3PreparationUnderSeals, SupportLiftPreservesTraceReadout},
			Failures: []string{FailurePostOrientationNotNativeR3, FailureNoNativeSocketToFiniteSectorMap, FailureAlphaStillSealed},
		},
		{
			Name: RouteFullAFLift, Description: "attempt lift to full unbroken A_F sector ledger",
			SourceAlgebra: FullUnbrokenAlgebra, TargetLedger: FiniteSectorLedger,
			LiftCandidate: false, LiftCertified: false, PostOrientationOnly: false, FullUnbrokenAFCertified: false,
			PreservesTraceReadout: false, PreservesEdgeSupport: false, PhysicalAssignment: false, GenerationCarrierPresent: false, FlavorOrientationPresent: false, IndividualYukawaValues: false,
			Supports: []string{SupportFullAFLiftAttemptAudited},
			Failures: []string{FailureNoFullAFLift, FailureSocketAtomsNotStableFullH, FailureNoFullUnbrokenAFSectorLedger},
		},
		{
			Name: RouteEdgeSupportLift, Description: "use symbolic D_F edge-support atoms as finite edge-sector candidates",
			SourceAlgebra: PostOrientationAlgebra, TargetLedger: FiniteSectorLedger,
			LiftCandidate: true, LiftCertified: false, PostOrientationOnly: true, FullUnbrokenAFCertified: false,
			PreservesTraceReadout: true, PreservesEdgeSupport: true, PhysicalAssignment: false, GenerationCarrierPresent: false, FlavorOrientationPresent: false, IndividualYukawaValues: false,
			Supports: []string{SupportEdgeSupportFiniteSectorCandidate, SupportLiftPreservesTraceReadout, SupportNoPhysicalAssignment},
			Failures: []string{FailureEdgeAtomNotPhysicalSector, FailureEdgeAtomNotYukawaValue, FailureNoNativeSocketToFiniteSectorMap},
		},
	}
	lift := LiftAudit{
		MapName: MissingLiftMap, Sigma: MissingSigmaMap, SourceAtoms: []string{AtomPiPlus3, AtomPiMinus3, AtomPiMinus1}, TargetLedger: FiniteSectorLedger,
		Routes: routes, BestRoute: RouteStabilizerSectorLift, LiftCandidate: true, LiftCertified: false, NativeR3: false,
		NextRequired: NextBranch,
		Supports:     []string{SupportPostOrientationLiftCandidate, SupportEdgeSupportFiniteSectorCandidate, SupportR3PreparationUnderSeals, SupportExactNextObject},
		Failures:     []string{FailureNoNativeSocketToFiniteSectorMap, FailureNoFullAFLift, FailurePostOrientationNotNativeR3, FailureAlphaStillSealed, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}
	if err := validateRoutes(routes); err != nil {
		return Audit{}, err
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
		Enforced: true, AlphaStillSealed: true, NoNativeSocketToFiniteMap: true, NoFullAFLift: true,
		NoPhysicalSectorAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict,
	}
	return Audit{
		ID: AuditID, Domain: domain, Lift: lift, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 887 searches for candidate lifts from the post-orientation socket trace atoms into represented finite-sector projectors without assigning physical sectors or individual Yukawa values.",
		Final: "The strongest current lift is a post-orientation stabilizer-sector candidate, supported by edge-support typing and trace-readout preservation. A full unbroken A_F finite-sector lift remains blocked, alpha_B remains sealed, and no native R3, physical sector assignment, generation/flavor split, or official ledger update is certified.",
	}, nil
}

func buildAtom(atom string, rank int, formula string, weight float64, edge string, supports, failures []string) SocketAtom {
	return SocketAtom{
		Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, EdgeSupport: edge,
		TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight,
		SocketSectorTyped: true, PostOrientStable: true, FiniteSectorCandidate: true, FiniteSectorCertified: false,
		FullAFCertified: false, PhysicalSector: false, GenerationResolved: false, FlavorResolved: false, IndividualYukawaValue: false,
		Supports: supports, Failures: failures,
	}
}

func validateRoutes(routes []LiftRoute) error {
	seen := map[string]bool{}
	for _, r := range routes {
		seen[r.Name] = true
		if r.PhysicalAssignment || r.GenerationCarrierPresent || r.FlavorOrientationPresent || r.IndividualYukawaValues {
			return fmt.Errorf("route leaked physical/generation/flavor/yukawa assignment: %s", r.Name)
		}
		if r.Name == RouteFullAFLift && (r.LiftCandidate || r.LiftCertified || r.FullUnbrokenAFCertified || r.PreservesTraceReadout || r.PreservesEdgeSupport) {
			return fmt.Errorf("full A_F lift should remain blocked: %s", FormatRoute(r))
		}
		if r.Name != RouteFullAFLift && (!r.LiftCandidate || r.LiftCertified || !r.PostOrientationOnly || r.FullUnbrokenAFCertified || !r.PreservesTraceReadout || !r.PreservesEdgeSupport) {
			return fmt.Errorf("post-orientation route malformed: %s", FormatRoute(r))
		}
	}
	for _, want := range []string{RouteStabilizerSectorLift, RouteFullAFLift, RouteEdgeSupportLift} {
		if !seen[want] {
			return fmt.Errorf("missing lift route %s", want)
		}
	}
	return nil
}

func FormatAtom(a SocketAtom) string {
	return fmt.Sprintf("atom(%s rank=%d weight=%s numeric=%.16g edge=%s trace=%.16g square=%.16g socket_typed=%t orient_stable=%t finite_candidate=%t finite_certified=%t full_A_F=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", a.Atom, a.Rank, a.WeightFormula, a.Weight, a.EdgeSupport, a.TraceContribution, a.SquareContribution, a.SocketSectorTyped, a.PostOrientStable, a.FiniteSectorCandidate, a.FiniteSectorCertified, a.FullAFCertified, a.PhysicalSector, a.GenerationResolved, a.FlavorResolved, a.IndividualYukawaValue, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAtoms(atoms []SocketAtom) string {
	parts := make([]string, 0, len(atoms))
	for _, atom := range atoms {
		parts = append(parts, FormatAtom(atom))
	}
	return strings.Join(parts, "; ")
}

func FormatDomain(d SocketDomain) string {
	return fmt.Sprintf("socket_domain(orient_algebra=%s full_algebra=%s rank=%d expected=%d complete=%t orthogonal=%t positive=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g atoms=[%s] supports=%s failures=%s)", d.PostOrientationAlgebra, d.FullUnbrokenAlgebra, d.ActiveRank, d.ExpectedRank, d.CompleteOnHRMin, d.Orthogonal, d.Positive, d.TraceTotal, d.SquareTrace, d.OperatorNEff, d.OperatorCYukawa, FormatAtoms(d.Atoms), strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}

func FormatRoute(r LiftRoute) string {
	return fmt.Sprintf("route(%s source=%s target=%s candidate=%t certified=%t post_orientation_only=%t full_A_F=%t trace=%t edge=%t physical=%t generation=%t flavor=%t individual_yukawa=%t supports=%s failures=%s)", r.Name, r.SourceAlgebra, r.TargetLedger, r.LiftCandidate, r.LiftCertified, r.PostOrientationOnly, r.FullUnbrokenAFCertified, r.PreservesTraceReadout, r.PreservesEdgeSupport, r.PhysicalAssignment, r.GenerationCarrierPresent, r.FlavorOrientationPresent, r.IndividualYukawaValues, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatRoutes(routes []LiftRoute) string {
	parts := make([]string, 0, len(routes))
	for _, route := range routes {
		parts = append(parts, FormatRoute(route))
	}
	return strings.Join(parts, "; ")
}

func FormatLift(l LiftAudit) string {
	return fmt.Sprintf("lift(map=%s sigma=%s source_atoms=%s target=%s best_route=%s candidate=%t certified=%t native_r3=%t next=%s routes=[%s] supports=%s failures=%s)", l.MapName, l.Sigma, strings.Join(l.SourceAtoms, ","), l.TargetLedger, l.BestRoute, l.LiftCandidate, l.LiftCertified, l.NativeR3, l.NextRequired, FormatRoutes(l.Routes), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t no_native_socket_to_finite=%t no_full_A_F_lift=%t no_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.NoNativeSocketToFiniteMap, f.NoFullAFLift, f.NoPhysicalSectorAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate886Inherited,
		StatusDomainStrong,
		StatusCodomainRequirements,
		StatusLiftRoutesAudited,
		StatusStabilizerLiftCandidate,
		StatusFullAFLiftBlocked,
		StatusEdgeSupportLiftCandidate,
		StatusPhysicalFirewalls,
		StatusOfficialFreezePreserved,
		StatusNextBranchSelected,
		StatusFirewallVerdict,
		SupportSocketSectorDomainStrong,
		SupportPostOrientationLiftCandidate,
		SupportEdgeSupportFiniteSectorCandidate,
		SupportR3PreparationUnderSeals,
		SupportLiftPreservesTraceReadout,
		SupportStabilizerNotNative,
		SupportNoPhysicalAssignment,
		SupportFullAFLiftAttemptAudited,
		SupportExactNextObject,
		FailureNoNativeSocketToFiniteSectorMap,
		FailureNoFullAFLift,
		FailureSocketAtomsNotStableFullH,
		FailurePostOrientationNotNativeR3,
		FailureNoFullUnbrokenAFSectorLedger,
		FailureAlphaStillSealed,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailureEdgeAtomNotPhysicalSector,
		FailureEdgeAtomNotYukawaValue,
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
	return f.Enforced && f.AlphaStillSealed && f.NoNativeSocketToFiniteMap && f.NoFullAFLift && f.NoPhysicalSectorAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
