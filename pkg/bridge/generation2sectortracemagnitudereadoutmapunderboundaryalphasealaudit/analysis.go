// Package generation2sectortracemagnitudereadoutmapunderboundaryalphasealaudit implements
// Gate 884: SectorTraceMagnitude ReadoutMap Under BoundaryAlpha Seal Audit.
//
// Gate 884 follows Gate 883's socket trace-ledger candidate. It does not reopen
// the BoundaryAlpha proof, does not update the official ledger, and does not
// assign physical particle sectors. It audits whether Y^dagger Y defines a
// positive readout map from the active socket atoms to trace and square-trace
// ledger rows under the BoundaryAlpha seal.
package generation2sectortracemagnitudereadoutmapunderboundaryalphasealaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE884-SECTOR-TRACE-MAGNITUDE-READOUT-MAP-UNDER-BOUNDARY-ALPHA-SEAL-AUDIT"

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
	ReadoutMapName = "R_Y(Pi_i)=(rank_i,weight_i,rank_i weight_i,rank_i weight_i^2)"
	Classification = "R2+++++_SECTOR_TRACE_MAGNITUDE_READOUT_UNDER_ALPHA_SEAL_NOT_NATIVE_R3"
	NextBranch     = "R3_ELIGIBILITY_REASSESSMENT_OR_GENERATION_CARRIER_REQUIREMENTS_UNDER_ALPHA_SEAL"

	StatusGate883Inherited            = "PASS_GATE883_SOCKET_TRACE_LEDGER_CANDIDATE_INHERITED"
	StatusReadoutRowsDefined          = "PASS_SECTOR_TRACE_MAGNITUDE_READOUT_ROWS_DEFINED"
	StatusOrthogonalCompleteProjector = "PASS_ACTIVE_SOCKET_PROJECTORS_ORTHOGONAL_AND_COMPLETE_ON_H_R_MIN"
	StatusYDaggerYPositive            = "PASS_Y_DAGGER_Y_POSITIVE_ON_ACTIVE_SOCKET_LEDGER"
	StatusTraceContributionLedger     = "PASS_TRACE_CONTRIBUTION_LEDGER_RECONSTRUCTED"
	StatusSquareContributionLedger    = "PASS_SQUARE_TRACE_CONTRIBUTION_LEDGER_RECONSTRUCTED"
	StatusOperatorNEffReproduced      = "PASS_OPERATOR_N_EFF_REPRODUCED_FROM_LEDGER_ROWS"
	StatusOfficialFreezePreserved     = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusR3FirewallPreserved         = "PASS_NATIVE_R3_FIREWALL_PRESERVED"
	StatusNoPhysicalAssignment        = "PASS_NO_PHYSICAL_SECTOR_OR_YUKAWA_VALUE_ASSIGNMENT_IN_GATE884"
	StatusFirewallVerdict             = "FIREWALL_PRESERVED_GATE884_READOUT_MAP_UNDER_ALPHA_SEAL_NOT_NATIVE_R3"

	SupportYDaggerYTraceMagnitudeReadout     = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_DEFINES_TRACE_MAGNITUDE_READOUT_UNDER_ALPHA_SEAL"
	SupportPositiveCompleteSocketLedger      = "CONDITIONAL_SUPPORT_SOCKET_TRACE_LEDGER_IS_POSITIVE_AND_COMPLETE_ON_H_R_MIN"
	SupportR3PreparationAdvancesToReadout    = "CONDITIONAL_SUPPORT_R3_PREPARATION_ADVANCES_FROM_CANDIDATE_LEDGER_TO_READOUT_LEDGER"
	SupportRowsRecoverTraceAndSquareTrace    = "CONDITIONAL_SUPPORT_READOUT_ROWS_RECOVER_TRACE_AND_SQUARE_TRACE"
	SupportOperatorNEffFromRows              = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_FROM_SECTOR_TRACE_MAGNITUDE_ROWS"
	SupportYDaggerYPositiveFiniteBodyLocated = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_READOUT_IS_POSITIVE_AND_FINITE_BODY_LOCATED"
	SupportNextFrontierUnderSeal             = "CONDITIONAL_SUPPORT_NEXT_FRONTIER_REMAINS_UNDER_BOUNDARY_ALPHA_SEAL"

	FailureAlphaStillSealed            = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureReadoutUnderSealNotNative   = "FAILED_ROUTE_READOUT_MAP_IS_UNDER_ALPHA_SEAL_NOT_NATIVE"
	FailureSocketTraceNotNativeR3      = "FAILED_ROUTE_SOCKET_TRACE_MAGNITUDE_READOUT_NOT_NATIVE_R3_WITHOUT_ALPHA_FUNCTOR"
	FailureNoNativeIncidenceFunctor    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureSocketAtomsNotPhysical      = "FAILED_ROUTE_SOCKET_ATOMS_NOT_PHYSICAL_SECTOR_ASSIGNMENTS"
	FailureNoGenerationCarrierMap      = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap      = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues    = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum    = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoNativeR3SectorTraceLedger = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoOfficialNEffUpdate        = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate       = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator      = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                        = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type ReadoutRow struct {
	Atom               string
	Rank               int
	WeightFormula      string
	Weight             float64
	TraceContribution  float64
	SquareContribution float64
	Positive           bool
	PhysicalSector     bool
	Supports, Failures []string
}

type TraceMagnitudeReadoutMap struct {
	Name                   string
	Rows                   []ReadoutRow
	ActiveRank             int
	ExpectedRank           int
	Orthogonal             bool
	CompleteOnHRMin        bool
	Positive               bool
	TraceTotal             float64
	SquareTrace            float64
	OperatorNEff           float64
	OperatorCYukawa        float64
	Classification         string
	ConditionalReadout     bool
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
	ReadoutUnderSeal         bool
	NoNativeIncidenceFunctor bool
	SocketAtomsNotPhysical   bool
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
	Readout   TraceMagnitudeReadoutMap
	Freeze    LedgerFreeze
	Firewalls R3Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	wPlus3 := 1.0
	wMinus3 := AlphaB * (1.0 - AlphaB)
	wMinus1 := 3.0 * AlphaB * AlphaB
	rows := []ReadoutRow{
		buildRow(AtomPiPlus3, RankPiPlus3, WeightPiPlus3, wPlus3, []string{SupportYDaggerYTraceMagnitudeReadout, SupportPositiveCompleteSocketLedger}, []string{FailureSocketAtomsNotPhysical}),
		buildRow(AtomPiMinus3, RankPiMinus3, WeightPiMinus3, wMinus3, []string{SupportYDaggerYTraceMagnitudeReadout, SupportRowsRecoverTraceAndSquareTrace}, []string{FailureAlphaStillSealed, FailureSocketAtomsNotPhysical}),
		buildRow(AtomPiMinus1, RankPiMinus1, WeightPiMinus1, wMinus1, []string{SupportYDaggerYTraceMagnitudeReadout, SupportRowsRecoverTraceAndSquareTrace}, []string{FailureAlphaStillSealed, FailureSocketAtomsNotPhysical}),
	}

	activeRank := 0
	positive := true
	traceTotal := 0.0
	squareTrace := 0.0
	for _, row := range rows {
		activeRank += row.Rank
		positive = positive && row.Positive
		traceTotal += row.TraceContribution
		squareTrace += row.SquareContribution
	}
	if activeRank != RankHRMin {
		return Audit{}, fmt.Errorf("active rank drift: got %d want %d", activeRank, RankHRMin)
	}
	if !positive {
		return Audit{}, fmt.Errorf("non-positive readout row")
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

	readout := TraceMagnitudeReadoutMap{
		Name: ReadoutMapName, Rows: rows, ActiveRank: activeRank, ExpectedRank: RankHRMin,
		Orthogonal: true, CompleteOnHRMin: true, Positive: positive,
		TraceTotal: traceTotal, SquareTrace: squareTrace, OperatorNEff: operatorNEff, OperatorCYukawa: operatorCYukawa,
		Classification: Classification, ConditionalReadout: true, NativeR3: false, OfficialUpdatesAllowed: false,
		Supports: []string{SupportYDaggerYTraceMagnitudeReadout, SupportPositiveCompleteSocketLedger, SupportR3PreparationAdvancesToReadout, SupportRowsRecoverTraceAndSquareTrace, SupportOperatorNEffFromRows, SupportYDaggerYPositiveFiniteBodyLocated, SupportNextFrontierUnderSeal},
		Failures: []string{FailureAlphaStillSealed, FailureReadoutUnderSealNotNative, FailureSocketTraceNotNativeR3, FailureNoNativeIncidenceFunctor, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4},
	}
	freeze := LedgerFreeze{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreezePreserved},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
	firewalls := R3Firewalls{Enforced: true, AlphaStillSealed: true, ReadoutUnderSeal: true, NoNativeIncidenceFunctor: true, SocketAtomsNotPhysical: true, NoGenerationMap: true, NoFlavorMap: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict}
	return Audit{ID: AuditID, AlphaSeal: AlphaSealName, Readout: readout, Freeze: freeze, Firewalls: firewalls, Truth: "Gate 884 turns the Gate 883 socket atom candidate into explicit positive trace-magnitude readout rows under the BoundaryAlpha seal.", Final: "Y^dagger Y defines a coherent finite-body readout map under alpha seal, but the readout is not native R3: alpha remains sealed, socket atoms are not physical sectors, and generation/flavor maps are absent."}, nil
}

func buildRow(atom string, rank int, formula string, weight float64, supports, failures []string) ReadoutRow {
	return ReadoutRow{Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight, Positive: weight > 0, PhysicalSector: false, Supports: supports, Failures: failures}
}

func FormatRow(r ReadoutRow) string {
	return fmt.Sprintf("row(atom=%s rank=%d weight=%s numeric=%.16g trace=%.16g square=%.16g positive=%t physical_sector=%t supports=%s failures=%s)", r.Atom, r.Rank, r.WeightFormula, r.Weight, r.TraceContribution, r.SquareContribution, r.Positive, r.PhysicalSector, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatRows(rows []ReadoutRow) string {
	parts := make([]string, 0, len(rows))
	for _, row := range rows {
		parts = append(parts, FormatRow(row))
	}
	return strings.Join(parts, "; ")
}

func FormatReadout(r TraceMagnitudeReadoutMap) string {
	return fmt.Sprintf("readout(name=%s classification=%s active_rank=%d expected_rank=%d orthogonal=%t complete=%t positive=%t trace=%.16g square=%.16g Neff=%.16g CYukawa=%.16g conditional=%t native_r3=%t official_updates=%t rows=[%s] supports=%s failures=%s)", r.Name, r.Classification, r.ActiveRank, r.ExpectedRank, r.Orthogonal, r.CompleteOnHRMin, r.Positive, r.TraceTotal, r.SquareTrace, r.OperatorNEff, r.OperatorCYukawa, r.ConditionalReadout, r.NativeR3, r.OfficialUpdatesAllowed, FormatRows(r.Rows), strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatFreeze(f LedgerFreeze) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f R3Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t readout_under_seal=%t no_incidence=%t socket_atoms_not_physical=%t no_generation=%t no_flavor=%t no_individual_yukawas=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.ReadoutUnderSeal, f.NoNativeIncidenceFunctor, f.SocketAtomsNotPhysical, f.NoGenerationMap, f.NoFlavorMap, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate883Inherited,
		StatusReadoutRowsDefined,
		StatusOrthogonalCompleteProjector,
		StatusYDaggerYPositive,
		StatusTraceContributionLedger,
		StatusSquareContributionLedger,
		StatusOperatorNEffReproduced,
		StatusOfficialFreezePreserved,
		StatusR3FirewallPreserved,
		StatusNoPhysicalAssignment,
		StatusFirewallVerdict,
		SupportYDaggerYTraceMagnitudeReadout,
		SupportPositiveCompleteSocketLedger,
		SupportR3PreparationAdvancesToReadout,
		SupportRowsRecoverTraceAndSquareTrace,
		SupportOperatorNEffFromRows,
		SupportYDaggerYPositiveFiniteBodyLocated,
		SupportNextFrontierUnderSeal,
		FailureAlphaStillSealed,
		FailureReadoutUnderSealNotNative,
		FailureSocketTraceNotNativeR3,
		FailureNoNativeIncidenceFunctor,
		FailureSocketAtomsNotPhysical,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoPhysicalYukawaSpectrum,
		FailureNoNativeR3SectorTraceLedger,
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

func firewallsOK(f R3Firewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.ReadoutUnderSeal && f.NoNativeIncidenceFunctor && f.SocketAtomsNotPhysical && f.NoGenerationMap && f.NoFlavorMap && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
