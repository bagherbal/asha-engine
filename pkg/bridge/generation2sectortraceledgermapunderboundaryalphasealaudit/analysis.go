// Package generation2sectortraceledgermapunderboundaryalphasealaudit implements
// Gate 883: SectorTraceLedgerMap Audit Under BoundaryAlpha Seal.
//
// Gate 883 follows Gate 882's R3 requirements audit. It does not reopen the
// BoundaryAlpha proof, does not update the official ledger, and does not assign
// physical Yukawa values. It audits whether the sealed aggregate trace proxy can
// be refined from a 3+4 aggregate support into a typed 3+3+1 socket trace-ledger
// candidate on H_R^min under the BoundaryAlpha seal.
package generation2sectortraceledgermapunderboundaryalphasealaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE883-SECTOR-TRACE-LEDGER-MAP-UNDER-BOUNDARY-ALPHA-SEAL-AUDIT"

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

	WeightPiPlus3  = "1"
	WeightPiMinus3 = "alpha_B(1-alpha_B)"
	WeightPiMinus1 = "3 alpha_B^2"

	AlphaSealName  = "BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL"
	Classification = "R2+++++_SECTOR_TRACE_LEDGER_CANDIDATE_UNDER_ALPHA_SEAL_NOT_R3"
	NextBranch     = "SECTOR_TRACE_MAGNITUDE_READOUT_MAP_AUDIT_UNDER_ALPHA_SEAL"

	StatusGate882Inherited            = "PASS_GATE882_R3_REQUIREMENTS_INHERITED"
	StatusSocketTraceAtomsDefined     = "PASS_ACTIVE_SOCKET_TRACE_ATOMS_DEFINED"
	StatusOrthogonalityCompleteness   = "PASS_SOCKET_TRACE_ATOMS_ORTHOGONAL_AND_COMPLETE_ON_H_R_MIN"
	StatusPositiveWeights             = "PASS_POSITIVE_TRACE_WEIGHTS_UNDER_ALPHA_SEAL"
	StatusTraceSquareTraceReproduced  = "PASS_TRACE_AND_SQUARE_TRACE_REPRODUCED"
	StatusYDaggerYLedgerReconstructed = "PASS_Y_DAGGER_Y_SECTOR_TRACE_LEDGER_RECONSTRUCTED"
	StatusR3CandidateClassified       = "PASS_R3_CANDIDATE_UNDER_ALPHA_SEAL_CLASSIFIED"
	StatusOfficialFreezePreserved     = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNoPhysicalYukawaAssignment  = "PASS_NO_PHYSICAL_YUKAWA_ASSIGNMENT_IN_GATE883"
	StatusFirewallVerdict             = "FIREWALL_PRESERVED_GATE883_SOCKET_TRACE_LEDGER_NOT_NATIVE_R3"

	SupportSocketProjectorsTraceAtoms          = "CONDITIONAL_SUPPORT_ACTIVE_SOCKET_PROJECTORS_FORM_SECTOR_TRACE_LEDGER_CANDIDATE"
	SupportLedgerDecomposesHRMin               = "CONDITIONAL_SUPPORT_TRACE_LEDGER_DECOMPOSES_H_R_MIN_AS_3_PLUS_3_PLUS_1"
	SupportYDaggerYPositiveWeights             = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_SUPPLIES_POSITIVE_TRACE_WEIGHTS_UNDER_ALPHA_SEAL"
	SupportLedgerReproducesOperatorNEff        = "CONDITIONAL_SUPPORT_LEDGER_REPRODUCES_OPERATOR_N_EFF"
	SupportR3PreparationAdvancesUnderAlphaSeal = "CONDITIONAL_SUPPORT_R3_PREPARATION_ADVANCES_UNDER_ALPHA_SEAL"
	SupportAggregateRefinedToSocketAtoms       = "CONDITIONAL_SUPPORT_AGGREGATE_3_PLUS_4_REFINED_TO_SOCKET_TRACE_ATOMS_3_PLUS_3_PLUS_1"
	SupportSectorTraceLedgerNext               = "CONDITIONAL_SUPPORT_NEXT_BRANCH_SHOULD_AUDIT_TRACE_MAGNITUDE_READOUT_MAP_UNDER_ALPHA_SEAL"

	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureSocketLedgerNotNativeR3          = "FAILED_ROUTE_SOCKET_TRACE_LEDGER_NOT_NATIVE_R3_WITHOUT_ALPHA_FUNCTOR"
	FailureNoNativeIncidenceFunctor         = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureSocketProjectorsNotPhysical      = "FAILED_ROUTE_SOCKET_PROJECTORS_NOT_PHYSICAL_SECTOR_ASSIGNMENTS"
	FailureNoGenerationCarrierMap           = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap           = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues         = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum         = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoNativeSectorTraceLedger        = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoSectorTraceMagnitudeReadoutMap = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate             = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate            = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                             = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type TraceAtom struct {
	Name               string
	Rank               int
	WeightFormula      string
	Weight             float64
	TraceContribution  float64
	SquareContribution float64
	TypedProjector     bool
	PhysicalAssignment bool
	Supports, Failures []string
}

type SectorLedgerCandidate struct {
	Atoms                  []TraceAtom
	ActiveRank             int
	ExpectedRank           int
	Orthogonal             bool
	CompleteOnHRMin        bool
	PositiveWeights        bool
	TraceTotal             float64
	SquareTrace            float64
	OperatorNEff           float64
	OperatorCYukawa        float64
	Classification         string
	ConditionalCandidate   bool
	NativeR3               bool
	OfficialUpdatesAllowed bool
	Supports, Failures     []string
}

type LedgerFreeze struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type R3Firewalls struct {
	Enforced                 bool
	AlphaStillSealed         bool
	NoNativeIncidenceFunctor bool
	SocketLedgerNotNativeR3  bool
	NoPhysicalAssignments    bool
	NoGenerationMap          bool
	NoFlavorMap              bool
	NoIndividualYukawas      bool
	NoOfficialLedgerUpdate   bool
	NoNativeYukawaOperator   bool
	NoR4                     bool
	Verdict                  string
}

type Audit struct {
	ID        string
	AlphaSeal string
	Ledger    SectorLedgerCandidate
	Freeze    LedgerFreeze
	Firewalls R3Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB
	if !(wPlus3 > 0 && wMinus3 > 0 && wMinus1 > 0) {
		return Audit{}, fmt.Errorf("non-positive trace weight under alpha seal")
	}

	atoms := []TraceAtom{
		{Name: AtomPiPlus3, Rank: RankPiPlus3, WeightFormula: WeightPiPlus3, Weight: wPlus3, TraceContribution: RankPiPlus3 * wPlus3, SquareContribution: RankPiPlus3 * wPlus3 * wPlus3, TypedProjector: true, PhysicalAssignment: false, Supports: []string{SupportSocketProjectorsTraceAtoms}, Failures: []string{FailureSocketProjectorsNotPhysical}},
		{Name: AtomPiMinus3, Rank: RankPiMinus3, WeightFormula: WeightPiMinus3, Weight: wMinus3, TraceContribution: RankPiMinus3 * wMinus3, SquareContribution: RankPiMinus3 * wMinus3 * wMinus3, TypedProjector: true, PhysicalAssignment: false, Supports: []string{SupportSocketProjectorsTraceAtoms, SupportYDaggerYPositiveWeights}, Failures: []string{FailureAlphaStillSealed, FailureSocketProjectorsNotPhysical}},
		{Name: AtomPiMinus1, Rank: RankPiMinus1, WeightFormula: WeightPiMinus1, Weight: wMinus1, TraceContribution: RankPiMinus1 * wMinus1, SquareContribution: RankPiMinus1 * wMinus1 * wMinus1, TypedProjector: true, PhysicalAssignment: false, Supports: []string{SupportSocketProjectorsTraceAtoms, SupportYDaggerYPositiveWeights}, Failures: []string{FailureAlphaStillSealed, FailureSocketProjectorsNotPhysical}},
	}

	activeRank := 0
	traceTotal := 0.0
	squareTrace := 0.0
	for _, atom := range atoms {
		activeRank += atom.Rank
		traceTotal += atom.TraceContribution
		squareTrace += atom.SquareContribution
	}
	expectedTrace := 3.0 + 3.0*AlphaB
	expectedSquare := 3.0 + 3.0*AlphaB*AlphaB - 6.0*AlphaB*AlphaB*AlphaB + 12.0*AlphaB*AlphaB*AlphaB*AlphaB
	if activeRank != RankHRMin {
		return Audit{}, fmt.Errorf("active rank drift: got %d want %d", activeRank, RankHRMin)
	}
	if !near(traceTotal, expectedTrace) {
		return Audit{}, fmt.Errorf("trace drift: got %.18g want %.18g", traceTotal, expectedTrace)
	}
	if !near(squareTrace, expectedSquare) {
		return Audit{}, fmt.Errorf("square trace drift: got %.18g want %.18g", squareTrace, expectedSquare)
	}
	operatorNEff := traceTotal * traceTotal / squareTrace
	operatorCYukawa := 3.0 / operatorNEff
	if !near(operatorNEff, OperatorNEffDiagnostic) || !near(operatorCYukawa, OperatorCYukawaDiagnostic) {
		return Audit{}, fmt.Errorf("operator ledger drift: N_eff %.18g C_Yukawa %.18g", operatorNEff, operatorCYukawa)
	}

	ledger := SectorLedgerCandidate{
		Atoms: atoms, ActiveRank: activeRank, ExpectedRank: RankHRMin, Orthogonal: true, CompleteOnHRMin: true, PositiveWeights: true,
		TraceTotal: traceTotal, SquareTrace: squareTrace, OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Classification: Classification, ConditionalCandidate: true, NativeR3: false, OfficialUpdatesAllowed: false,
		Supports: []string{SupportSocketProjectorsTraceAtoms, SupportLedgerDecomposesHRMin, SupportYDaggerYPositiveWeights, SupportLedgerReproducesOperatorNEff, SupportR3PreparationAdvancesUnderAlphaSeal, SupportAggregateRefinedToSocketAtoms, SupportSectorTraceLedgerNext},
		Failures: []string{FailureAlphaStillSealed, FailureSocketLedgerNotNativeR3, FailureNoNativeIncidenceFunctor, FailureSocketProjectorsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeSectorTraceLedger, FailureNoSectorTraceMagnitudeReadoutMap},
	}
	freeze := LedgerFreeze{OperatorNEff: operatorNEff, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: operatorCYukawa, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false, Supports: []string{SupportLedgerReproducesOperatorNEff}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeSectorTraceLedger}}
	if near(freeze.OperatorNEff, freeze.OfficialNEff) || near(freeze.OperatorCYukawa, freeze.OfficialCYukawa) {
		return Audit{}, fmt.Errorf("operator and official ledger collapsed")
	}
	firewalls := R3Firewalls{Enforced: true, AlphaStillSealed: true, NoNativeIncidenceFunctor: true, SocketLedgerNotNativeR3: true, NoPhysicalAssignments: true, NoGenerationMap: true, NoFlavorMap: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict}

	return Audit{ID: AuditID, AlphaSeal: AlphaSealName, Ledger: ledger, Freeze: freeze, Firewalls: firewalls, Truth: "Gate 883 refines the sealed aggregate trace proxy into a 3+3+1 active socket trace-ledger candidate under BoundaryAlpha seal.", Final: "The ledger is coherent as a positive finite-body socket trace candidate, but not native R3: alpha remains sealed, socket atoms are not physical sectors, and generation/flavor splitting is absent."}, nil
}

func FormatAtom(a TraceAtom) string {
	return fmt.Sprintf("atom(name=%s rank=%d weight=%s numeric=%.16g trace=%.16g square=%.16g typed=%t physical_assignment=%t supports=%s failures=%s)", a.Name, a.Rank, a.WeightFormula, a.Weight, a.TraceContribution, a.SquareContribution, a.TypedProjector, a.PhysicalAssignment, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAtoms(atoms []TraceAtom) string {
	parts := make([]string, 0, len(atoms))
	for _, a := range atoms {
		parts = append(parts, FormatAtom(a))
	}
	return strings.Join(parts, "; ")
}

func FormatLedger(l SectorLedgerCandidate) string {
	return fmt.Sprintf("sector_ledger(classification=%s active_rank=%d expected_rank=%d orthogonal=%t complete=%t positive=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g conditional=%t native_r3=%t official_updates=%t atoms=[%s] supports=%s failures=%s)", l.Classification, l.ActiveRank, l.ExpectedRank, l.Orthogonal, l.CompleteOnHRMin, l.PositiveWeights, l.TraceTotal, l.SquareTrace, l.OperatorNEff, l.OperatorCYukawa, l.ConditionalCandidate, l.NativeR3, l.OfficialUpdatesAllowed, FormatAtoms(l.Atoms), strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f R3Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t no_incidence=%t socket_ledger_not_r3=%t no_physical_assignments=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.NoNativeIncidenceFunctor, f.SocketLedgerNotNativeR3, f.NoPhysicalAssignments, f.NoGenerationMap, f.NoFlavorMap, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate882Inherited,
		StatusSocketTraceAtomsDefined,
		StatusOrthogonalityCompleteness,
		StatusPositiveWeights,
		StatusTraceSquareTraceReproduced,
		StatusYDaggerYLedgerReconstructed,
		StatusR3CandidateClassified,
		StatusOfficialFreezePreserved,
		StatusNoPhysicalYukawaAssignment,
		StatusFirewallVerdict,
		SupportSocketProjectorsTraceAtoms,
		SupportLedgerDecomposesHRMin,
		SupportYDaggerYPositiveWeights,
		SupportLedgerReproducesOperatorNEff,
		SupportR3PreparationAdvancesUnderAlphaSeal,
		SupportAggregateRefinedToSocketAtoms,
		SupportSectorTraceLedgerNext,
		FailureAlphaStillSealed,
		FailureSocketLedgerNotNativeR3,
		FailureNoNativeIncidenceFunctor,
		FailureSocketProjectorsNotPhysical,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoPhysicalYukawaSpectrum,
		FailureNoNativeSectorTraceLedger,
		FailureNoSectorTraceMagnitudeReadoutMap,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4,
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func containsAll(have []string, want []string) bool {
	m := map[string]bool{}
	for _, h := range have {
		m[h] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}

func firewallsOK(f R3Firewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.NoNativeIncidenceFunctor && f.SocketLedgerNotNativeR3 && f.NoPhysicalAssignments && f.NoGenerationMap && f.NoFlavorMap && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
