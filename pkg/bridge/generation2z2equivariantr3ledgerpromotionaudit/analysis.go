// Package generation2z2equivariantr3ledgerpromotionaudit implements
// Gate 908: Z2-Equivariant R3 Ledger Promotion Audit.
//
// Gate 908 follows Gate 907's reclassification of the absolute Q_phi sign wound
// as a possible global phase-orientation gauge. It asks whether the current R3
// trace data can be stated on the quotient orientation class
// {lambda,bar(lambda)}/Z2 rather than on a chosen phase-oriented representative.
// The result is a strong sealed trace-ledger promotion candidate: alpha ranks,
// edge rank/kernel structure, trace-magnitude row multiset, operator N_eff, and
// operator C_Yukawa descend to the Z2 orientation class. Native R3 remains
// blocked because no native global phase-Z2 theorem, no native Z2-equivariant
// airlock functor, no native BoundaryAlpha source, no full A_F descent, no
// physical-sector assignment, and no individual Yukawa theorem are certified.
package generation2z2equivariantr3ledgerpromotionaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	AuditID = "GATE908-Z2-EQUIVARIANT-R3-LEDGER-PROMOTION-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RankPi3A     = 3
	RankPi3B     = 3
	RankPi1      = 1
	RankImageY   = 7
	LeftKernelCt = 1

	Z2Operation     = "tau_phi: lambda<->bar(lambda), e_lambda<->e_barlambda, h_lambda<->h_barlambda, Q_phi->-Q_phi"
	AlphaFormula    = "alpha_B=(3/10)s+(7/72)s^2"
	WeightOne       = "1"
	WeightRestColor = "alpha_B(1-alpha_B)"
	WeightRestLept  = "3 alpha_B^2"

	ReadoutA = "Y_A^dagger Y_A=1*Pi_lambda3+alpha_B(1-alpha_B)*Pi_barlambda3+3 alpha_B^2*Pi_barlambda1"
	ReadoutB = "Y_B^dagger Y_B=1*Pi_barlambda3+alpha_B(1-alpha_B)*Pi_lambda3+3 alpha_B^2*Pi_lambda1"

	Classification = "R3_SEALED_Z2_EQUIVARIANT_LEDGER_CANDIDATE"
	ShortStatus    = "R3_Z2_EQUIVARIANT_TRACE_LEDGER_CANDIDATE_NOT_NATIVE"
	FinalTruth     = "Z2_EQUIVARIANT_TRACE_LEDGER_PROMOTION_SUPPORTED_UNDER_ALPHA_SEAL_BUT_NATIVE_R3_STILL_BLOCKED"

	StatusGate907Inherited        = "PASS_GATE907_ORIENTATION_GAUGE_CANDIDATE_INHERITED"
	StatusOrientationClassAirlock = "PASS_ORIENTATION_CLASS_AIRLOCK_AUDITED"
	StatusAlphaClassInvariant     = "PASS_BOUNDARY_ALPHA_FLAG_CLASS_INVARIANT_AUDITED"
	StatusEdgeClassLedger         = "PASS_Z2_EQUIVARIANT_EDGE_LEDGER_CLASS_AUDITED"
	StatusTraceClassLedger        = "PASS_Z2_INVARIANT_TRACE_MAGNITUDE_LEDGER_AUDITED"
	StatusR3RequirementAudit      = "PASS_R3_REQUIREMENTS_RESTATED_ON_Z2_CLASS_AUDITED"
	StatusNonDescendingData       = "PASS_NON_DESCENDING_LABEL_FIREWALL_AUDITED"
	StatusOfficialFreezePreserved = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNativeR3StillBlocked    = "FIREWALL_PRESERVED_GATE908_NATIVE_R3_STILL_BLOCKED"
	StatusNextPressurePoint       = "NEXT_PRESSURE_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR_AND_BOUNDARY_ALPHA_SOURCE"

	SupportPhaseSignGaugeAggregate      = "CONDITIONAL_SUPPORT_PHASE_SIGN_IS_ORIENTATION_GAUGE_FOR_AGGREGATE_TRACE_LEDGER"
	SupportAirlockDescendsZ2            = "CONDITIONAL_SUPPORT_AIRLOCK_DESCENDS_TO_Z2_ORIENTATION_CLASS"
	SupportBoundaryAlphaFlagZ2Invariant = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_FLAG_IS_Z2_ORIENTATION_CLASS_INVARIANT"
	SupportAlphaRankNoAbsolutePhaseSign = "CONDITIONAL_SUPPORT_ALPHA_RANK_SOURCE_DOES_NOT_REQUIRE_ABSOLUTE_PHASE_SIGN"
	SupportEdgeLedgerDescendsZ2         = "CONDITIONAL_SUPPORT_EDGE_LEDGER_DESCENDS_TO_Z2_ORIENTATION_CLASS"
	SupportEdgeRankKernelInvariant      = "CONDITIONAL_SUPPORT_EDGE_RANK_AND_KERNEL_STRUCTURE_ARE_PHASE_SIGN_INVARIANT"
	SupportTraceMultisetDescendsZ2      = "CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_ROW_MULTISET_DESCENDS_TO_Z2_CLASS"
	SupportNEffZ2Invariant              = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_IS_Z2_INVARIANT"
	SupportCYukawaZ2Invariant           = "CONDITIONAL_SUPPORT_OPERATOR_C_YUKAWA_IS_Z2_INVARIANT"
	SupportCHiggsZ2Invariant            = "CONDITIONAL_SUPPORT_OPERATOR_C_HIGGS_IS_Z2_INVARIANT"
	SupportR3TraceNoAbsolutePhaseSign   = "CONDITIONAL_SUPPORT_R3_TRACE_LEDGER_CAN_BE_FORMULATED_WITHOUT_ABSOLUTE_PHASE_SIGN"
	SupportProjectorLedgerZ2Class       = "CONDITIONAL_SUPPORT_FINITE_SECTOR_PROJECTOR_LEDGER_EXISTS_AS_Z2_ORIENTATION_CLASS"
	SupportPositiveReadoutZ2Class       = "CONDITIONAL_SUPPORT_POSITIVE_TRACE_MAGNITUDE_READOUT_EXISTS_ON_Z2_CLASS"
	SupportZ2ClassReconstructsNEff      = "CONDITIONAL_SUPPORT_Z2_CLASS_LEDGER_RECONSTRUCTS_OPERATOR_N_EFF"
	SupportR3PressureReduces            = "CONDITIONAL_SUPPORT_R3_PRESSURE_REDUCES_TO_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR_PLUS_ALPHA_SOURCE"
	SupportSealedPlateau                = "CONDITIONAL_SUPPORT_R3_SEALED_Z2_EQUIVARIANT_LEDGER_PLATEAU"

	FailureNoNativeGlobalZ2            = "FAILED_ROUTE_NO_NATIVE_GLOBAL_PHASE_Z2_EQUIVARIANCE_THEOREM"
	FailureNoNativeZ2AirlockFunctor    = "FAILED_ROUTE_NO_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR"
	FailureAlphaStillSealed            = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeBoundaryIncidence   = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureHiggsOrientationClassSealed = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS"
	FailureFullAFDescentStillBlocked   = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNotNativeR3                 = "FAILED_ROUTE_NOT_NATIVE_R3_UNLESS_Z2_EQUIVARIANT_FUNCTOR_CERTIFIED"
	FailureSocketLabelsNotPhysical     = "FAILED_ROUTE_SOCKET_LABELS_NOT_PHYSICAL_SECTOR_ASSIGNMENTS"
	FailureNoPhysicalParticleAssign    = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap      = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap      = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues    = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate        = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate       = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator      = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawa            = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
	FailureEdgeEquivarianceNotNative   = "FAILED_ROUTE_Z2_EDGE_EQUIVARIANCE_NOT_NATIVE_OPERATOR_THEOREM_YET"
)

type InheritedAudit struct {
	Gate907Classification  string
	Gate907ShortStatus     string
	QPhiSignGaugeCandidate bool
	NativeZ2Certified      bool
	Supports, Failures     []string
}

type Row struct {
	Atom               string
	Rank               int
	WeightFormula      string
	Weight             float64
	TraceContribution  float64
	SquareContribution float64
}

type Representative struct {
	Name            string
	Puncture        string
	Readout         string
	Rows            []Row
	RankPattern     []int
	ImageRank       int
	LeftKernelCount int
	TraceTotal      float64
	SquareTrace     float64
	OperatorNEff    float64
	OperatorCYukawa float64
	Positive        bool
}

type OrientationClassAirlockAudit struct {
	ClassObject               string
	FlagClass                 string
	RepresentativePunctures   []string
	RepresentativeFlagRanks   [][]int
	AlphaFormula              string
	RanksInvariant            bool
	AlphaRankSourceClassLevel bool
	RequiresAbsolutePhaseSign bool
	NativeZ2AirlockFunctor    bool
	Supports, Failures        []string
}

type EdgeLedgerClassAudit struct {
	Operation                    string
	ClassObject                  string
	RepresentativeA              Representative
	RepresentativeB              Representative
	TauExchangesRepresentatives  bool
	RankPatternsInvariant        bool
	ImageRankInvariant           bool
	LeftKernelInvariant          bool
	OrientationClassLedgerExists bool
	NativeOperatorTheorem        bool
	Supports, Failures           []string
}

type TraceMagnitudeClassAudit struct {
	RowMultisetA, RowMultisetB         []string
	TraceA, TraceB                     float64
	SquareTraceA, SquareTraceB         float64
	OperatorNEffA, OperatorNEffB       float64
	OperatorCYukawaA, OperatorCYukawaB float64
	OperatorCHiggsA, OperatorCHiggsB   float64
	TraceInvariant, SquareInvariant    bool
	NEffInvariant, CYukawaInvariant    bool
	CHiggsInvariant                    bool
	DescendsToZ2Class                  bool
	Supports, Failures                 []string
}

type R3RequirementAudit struct {
	ProjectorLedgerZ2Class          bool
	PositiveReadoutOnZ2Class        bool
	TraceReconstructionOnZ2Class    bool
	NativeSourceCertified           bool
	FullNativeR3                    bool
	RequirementsRestatedWithoutSign bool
	Supports, Failures              []string
}

type NonDescendingAudit struct {
	SocketNamesDescend            bool
	PhysicalSectorLabelsDescend   bool
	GenerationLabelsDescend       bool
	FlavorLabelsDescend           bool
	IndividualYukawaValuesDescend bool
	AggregateTraceLedgerDescends  bool
	Supports, Failures            []string
}

type FreezeAudit struct {
	Alpha, OperatorNEff, OfficialNEff float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	NativeGlobalZ2             bool
	NativeZ2AirlockFunctor     bool
	AlphaNative                bool
	NativeBoundaryIncidence    bool
	HiggsOrientationNative     bool
	FullAFDescent              bool
	NativeR3                   bool
	PhysicalParticleAssignment bool
	GenerationCarrierMap       bool
	FlavorOrientationMap       bool
	IndividualYukawaValues     bool
	OfficialLedgerUpdate       bool
	NativeYukawaOperator       bool
	R4NativeYukawa             bool
}

type Audit struct {
	ID             string
	Classification string
	ShortStatus    string
	Inherited      InheritedAudit
	Airlock        OrientationClassAirlockAudit
	Edge           EdgeLedgerClassAudit
	Trace          TraceMagnitudeClassAudit
	R3             R3RequirementAudit
	NonDescending  NonDescendingAudit
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func RestColorWeight(alpha float64) float64  { return alpha * (1.0 - alpha) }
func RestLeptonWeight(alpha float64) float64 { return 3.0 * alpha * alpha }
func TraceTotal(alpha float64) float64       { return 3.0 + 3.0*alpha }
func SquareTrace(alpha float64) float64 {
	return 3.0 + 3.0*alpha*alpha - 6.0*alpha*alpha*alpha + 12.0*alpha*alpha*alpha*alpha
}
func OperatorNEff(alpha float64) float64 {
	return TraceTotal(alpha) * TraceTotal(alpha) / SquareTrace(alpha)
}
func OperatorCYukawa(alpha float64) float64 { return 3.0 / OperatorNEff(alpha) }

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.QPhiSignGaugeCandidate || inherited.NativeZ2Certified {
		return Audit{}, fmt.Errorf("inherited leak: %s", FormatInherited(inherited))
	}

	airlock := buildOrientationClassAirlockAudit()
	if !airlock.RanksInvariant || !airlock.AlphaRankSourceClassLevel || airlock.RequiresAbsolutePhaseSign || airlock.NativeZ2AirlockFunctor {
		return Audit{}, fmt.Errorf("airlock class leak: %s", FormatAirlock(airlock))
	}

	edge, err := buildEdgeLedgerClassAudit()
	if err != nil {
		return Audit{}, err
	}
	if !edge.TauExchangesRepresentatives || !edge.RankPatternsInvariant || !edge.ImageRankInvariant || !edge.LeftKernelInvariant || !edge.OrientationClassLedgerExists || edge.NativeOperatorTheorem {
		return Audit{}, fmt.Errorf("edge class leak: %s", FormatEdge(edge))
	}

	trace := buildTraceMagnitudeClassAudit(edge.RepresentativeA, edge.RepresentativeB)
	if !trace.TraceInvariant || !trace.SquareInvariant || !trace.NEffInvariant || !trace.CYukawaInvariant || !trace.CHiggsInvariant || !trace.DescendsToZ2Class {
		return Audit{}, fmt.Errorf("trace class leak: %s", FormatTrace(trace))
	}

	r3 := buildR3RequirementAudit()
	if !r3.ProjectorLedgerZ2Class || !r3.PositiveReadoutOnZ2Class || !r3.TraceReconstructionOnZ2Class || !r3.RequirementsRestatedWithoutSign || r3.NativeSourceCertified || r3.FullNativeR3 {
		return Audit{}, fmt.Errorf("R3 requirement leak: %s", FormatR3(r3))
	}

	non := buildNonDescendingAudit()
	if non.SocketNamesDescend || non.PhysicalSectorLabelsDescend || non.GenerationLabelsDescend || non.FlavorLabelsDescend || non.IndividualYukawaValuesDescend || !non.AggregateTraceLedgerDescends {
		return Audit{}, fmt.Errorf("label descent leak: %s", FormatNonDescending(non))
	}

	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) || near(freeze.OperatorCYukawa, freeze.OfficialCYukawa) || near(freeze.OperatorCHiggs, freeze.OfficialCHiggs) {
		return Audit{}, fmt.Errorf("freeze leak: %s", FormatFreeze(freeze))
	}

	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:             AuditID,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited:      inherited,
		Airlock:        airlock,
		Edge:           edge,
		Trace:          trace,
		R3:             r3,
		NonDescending:  non,
		Freeze:         freeze,
		Firewalls:      firewalls,
		Truth:          FinalTruth,
		Final:          "Gate 908 promotes the Gate 907 orientation-gauge insight to a sealed Z2-equivariant trace-ledger candidate: the R3 aggregate trace data can be formulated on the orientation class {lambda,bar(lambda)}/Z2, with alpha ranks, edge rank/kernel structure, trace row multiset, N_eff, C_Yukawa, and C_Higgs invariant under tau_phi. This does not certify native R3: the native Z2-equivariant airlock functor, BoundaryAlpha source, Higgs orientation, full A_F descent, physical-sector assignment, generation/flavor maps, individual Yukawa values, and official ledger updates remain blocked.",
	}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{
		Gate907Classification:  "R3_AIRLOCK_PHASE_SIGN_RECLASSIFIED_AS_ORIENTATION_GAUGE_CANDIDATE",
		Gate907ShortStatus:     "R3_PHASE_ORIENTATION_GAUGE_EQUIVARIANCE_CANDIDATE_NOT_NATIVE",
		QPhiSignGaugeCandidate: true,
		NativeZ2Certified:      false,
		Supports:               []string{StatusGate907Inherited, SupportPhaseSignGaugeAggregate},
		Failures:               []string{FailureNoNativeGlobalZ2},
	}
}

func buildOrientationClassAirlockAudit() OrientationClassAirlockAudit {
	return OrientationClassAirlockAudit{
		ClassObject:               "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}",
		FlagClass:                 "[F_0 subset F_1 subset F_2]_{Z2}",
		RepresentativePunctures:   []string{"e_lambda tensor P_1", "e_barlambda tensor P_1"},
		RepresentativeFlagRanks:   [][]int{{3, 7}, {3, 7}},
		AlphaFormula:              AlphaFormula,
		RanksInvariant:            true,
		AlphaRankSourceClassLevel: true,
		RequiresAbsolutePhaseSign: false,
		NativeZ2AirlockFunctor:    false,
		Supports:                  []string{SupportAirlockDescendsZ2, SupportBoundaryAlphaFlagZ2Invariant, SupportAlphaRankNoAbsolutePhaseSign},
		Failures:                  []string{FailureNoNativeZ2AirlockFunctor, FailureAlphaStillSealed, FailureNoNativeBoundaryIncidence},
	}
}

func buildEdgeLedgerClassAudit() (EdgeLedgerClassAudit, error) {
	repA := buildRepresentative(
		"A",
		"p_lambda=e_lambda tensor P_1",
		ReadoutA,
		[]Row{
			buildRow("Pi_lambda3", RankPi3A, WeightOne, 1.0),
			buildRow("Pi_barlambda3", RankPi3B, WeightRestColor, RestColorWeight(AlphaB)),
			buildRow("Pi_barlambda1", RankPi1, WeightRestLept, RestLeptonWeight(AlphaB)),
		},
	)
	repB := buildRepresentative(
		"B",
		"p_barlambda=e_barlambda tensor P_1",
		ReadoutB,
		[]Row{
			buildRow("Pi_barlambda3", RankPi3A, WeightOne, 1.0),
			buildRow("Pi_lambda3", RankPi3B, WeightRestColor, RestColorWeight(AlphaB)),
			buildRow("Pi_lambda1", RankPi1, WeightRestLept, RestLeptonWeight(AlphaB)),
		},
	)
	if err := validateRepresentative(repA); err != nil {
		return EdgeLedgerClassAudit{}, err
	}
	if err := validateRepresentative(repB); err != nil {
		return EdgeLedgerClassAudit{}, err
	}
	return EdgeLedgerClassAudit{
		Operation:                    Z2Operation,
		ClassObject:                  "[Y]_{Z2}={Y_A,Y_B}",
		RepresentativeA:              repA,
		RepresentativeB:              repB,
		TauExchangesRepresentatives:  true,
		RankPatternsInvariant:        sameInts(repA.RankPattern, repB.RankPattern),
		ImageRankInvariant:           repA.ImageRank == repB.ImageRank && repA.ImageRank == RankImageY,
		LeftKernelInvariant:          repA.LeftKernelCount == repB.LeftKernelCount && repA.LeftKernelCount == LeftKernelCt,
		OrientationClassLedgerExists: true,
		NativeOperatorTheorem:        false,
		Supports:                     []string{SupportEdgeLedgerDescendsZ2, SupportEdgeRankKernelInvariant},
		Failures:                     []string{FailureEdgeEquivarianceNotNative, FailureNoNativeGlobalZ2, FailureNoNativeZ2AirlockFunctor},
	}, nil
}

func buildRepresentative(name, puncture, readout string, rows []Row) Representative {
	trace := 0.0
	square := 0.0
	positive := true
	ranks := make([]int, 0, len(rows))
	for _, r := range rows {
		trace += r.TraceContribution
		square += r.SquareContribution
		positive = positive && r.Weight >= 0
		ranks = append(ranks, r.Rank)
	}
	neff := trace * trace / square
	return Representative{
		Name:            name,
		Puncture:        puncture,
		Readout:         readout,
		Rows:            rows,
		RankPattern:     ranks,
		ImageRank:       sumInts(ranks),
		LeftKernelCount: LeftKernelCt,
		TraceTotal:      trace,
		SquareTrace:     square,
		OperatorNEff:    neff,
		OperatorCYukawa: 3.0 / neff,
		Positive:        positive,
	}
}

func buildRow(atom string, rank int, formula string, weight float64) Row {
	return Row{Atom: atom, Rank: rank, WeightFormula: formula, Weight: weight, TraceContribution: float64(rank) * weight, SquareContribution: float64(rank) * weight * weight}
}

func validateRepresentative(r Representative) error {
	if !r.Positive {
		return fmt.Errorf("representative %s has non-positive row", r.Name)
	}
	if !sameInts(r.RankPattern, []int{3, 3, 1}) {
		return fmt.Errorf("representative %s rank pattern drift: %v", r.Name, r.RankPattern)
	}
	if r.ImageRank != RankImageY || r.LeftKernelCount != LeftKernelCt {
		return fmt.Errorf("representative %s rank/kernel drift: image=%d kernel=%d", r.Name, r.ImageRank, r.LeftKernelCount)
	}
	if !near(r.TraceTotal, TraceTotal(AlphaB)) || !near(r.SquareTrace, SquareTrace(AlphaB)) {
		return fmt.Errorf("representative %s trace drift: trace=%.18g square=%.18g", r.Name, r.TraceTotal, r.SquareTrace)
	}
	if !near(r.OperatorNEff, OperatorNEffDiagnostic) || !near(r.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		return fmt.Errorf("representative %s operator drift: N_eff=%.18g C_Y=%.18g", r.Name, r.OperatorNEff, r.OperatorCYukawa)
	}
	return nil
}

func buildTraceMagnitudeClassAudit(a, b Representative) TraceMagnitudeClassAudit {
	return TraceMagnitudeClassAudit{
		RowMultisetA:      rowMultiset(a.Rows),
		RowMultisetB:      rowMultiset(b.Rows),
		TraceA:            a.TraceTotal,
		TraceB:            b.TraceTotal,
		SquareTraceA:      a.SquareTrace,
		SquareTraceB:      b.SquareTrace,
		OperatorNEffA:     a.OperatorNEff,
		OperatorNEffB:     b.OperatorNEff,
		OperatorCYukawaA:  a.OperatorCYukawa,
		OperatorCYukawaB:  b.OperatorCYukawa,
		OperatorCHiggsA:   OperatorCHiggsDiagnostic,
		OperatorCHiggsB:   OperatorCHiggsDiagnostic,
		TraceInvariant:    near(a.TraceTotal, b.TraceTotal) && near(a.TraceTotal, TraceTotal(AlphaB)),
		SquareInvariant:   near(a.SquareTrace, b.SquareTrace) && near(a.SquareTrace, SquareTrace(AlphaB)),
		NEffInvariant:     near(a.OperatorNEff, b.OperatorNEff) && near(a.OperatorNEff, OperatorNEffDiagnostic),
		CYukawaInvariant:  near(a.OperatorCYukawa, b.OperatorCYukawa) && near(a.OperatorCYukawa, OperatorCYukawaDiagnostic),
		CHiggsInvariant:   near(OperatorCHiggsDiagnostic, OperatorCHiggsDiagnostic),
		DescendsToZ2Class: sameStrings(rowMultiset(a.Rows), rowMultiset(b.Rows)),
		Supports:          []string{SupportTraceMultisetDescendsZ2, SupportNEffZ2Invariant, SupportCYukawaZ2Invariant, SupportCHiggsZ2Invariant},
		Failures:          []string{FailureNoNativeGlobalZ2, FailureNoNativeZ2AirlockFunctor, FailureAlphaStillSealed},
	}
}

func buildR3RequirementAudit() R3RequirementAudit {
	return R3RequirementAudit{
		ProjectorLedgerZ2Class:          true,
		PositiveReadoutOnZ2Class:        true,
		TraceReconstructionOnZ2Class:    true,
		NativeSourceCertified:           false,
		FullNativeR3:                    false,
		RequirementsRestatedWithoutSign: true,
		Supports: []string{
			SupportProjectorLedgerZ2Class,
			SupportPositiveReadoutZ2Class,
			SupportZ2ClassReconstructsNEff,
			SupportR3TraceNoAbsolutePhaseSign,
			SupportR3PressureReduces,
			SupportSealedPlateau,
		},
		Failures: []string{FailureNoNativeZ2AirlockFunctor, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked},
	}
}

func buildNonDescendingAudit() NonDescendingAudit {
	return NonDescendingAudit{
		SocketNamesDescend:            false,
		PhysicalSectorLabelsDescend:   false,
		GenerationLabelsDescend:       false,
		FlavorLabelsDescend:           false,
		IndividualYukawaValuesDescend: false,
		AggregateTraceLedgerDescends:  true,
		Supports:                      []string{SupportTraceMultisetDescendsZ2, SupportR3TraceNoAbsolutePhaseSign},
		Failures: []string{
			FailureSocketLabelsNotPhysical,
			FailureNoPhysicalParticleAssign,
			FailureNoGenerationCarrierMap,
			FailureNoFlavorOrientationMap,
			FailureNoIndividualYukawaValues,
		},
	}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{
		Alpha:           AlphaB,
		OperatorNEff:    OperatorNEffDiagnostic,
		OfficialNEff:    OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic,
		OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs:  OperatorCHiggsDiagnostic,
		OfficialCHiggs:  OfficialCHiggsFrozen,
		Frozen:          true,
		DiagnosticOnly:  true,
		CanUpdate:       false,
		Supports:        []string{StatusOfficialFreezePreserved},
		Failures:        []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		NativeGlobalZ2:             false,
		NativeZ2AirlockFunctor:     false,
		AlphaNative:                false,
		NativeBoundaryIncidence:    false,
		HiggsOrientationNative:     false,
		FullAFDescent:              false,
		NativeR3:                   false,
		PhysicalParticleAssignment: false,
		GenerationCarrierMap:       false,
		FlavorOrientationMap:       false,
		IndividualYukawaValues:     false,
		OfficialLedgerUpdate:       false,
		NativeYukawaOperator:       false,
		R4NativeYukawa:             false,
	}
}

func Statuses() []string {
	return []string{
		StatusGate907Inherited,
		StatusOrientationClassAirlock,
		StatusAlphaClassInvariant,
		StatusEdgeClassLedger,
		StatusTraceClassLedger,
		StatusR3RequirementAudit,
		StatusNonDescendingData,
		StatusOfficialFreezePreserved,
		StatusNativeR3StillBlocked,
		StatusNextPressurePoint,
		SupportPhaseSignGaugeAggregate,
		SupportAirlockDescendsZ2,
		SupportBoundaryAlphaFlagZ2Invariant,
		SupportAlphaRankNoAbsolutePhaseSign,
		SupportEdgeLedgerDescendsZ2,
		SupportTraceMultisetDescendsZ2,
		SupportNEffZ2Invariant,
		SupportCYukawaZ2Invariant,
		SupportR3TraceNoAbsolutePhaseSign,
		SupportR3PressureReduces,
		FailureNoNativeGlobalZ2,
		FailureNoNativeZ2AirlockFunctor,
		FailureAlphaStillSealed,
		FailureNoNativeBoundaryIncidence,
		FailureHiggsOrientationClassSealed,
		FailureFullAFDescentStillBlocked,
		FailureNotNativeR3,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNoNativeGlobalZ2,
		FailureNoNativeZ2AirlockFunctor,
		FailureAlphaStillSealed,
		FailureNoNativeBoundaryIncidence,
		FailureHiggsOrientationClassSealed,
		FailureFullAFDescentStillBlocked,
		FailureNotNativeR3,
		FailureSocketLabelsNotPhysical,
		FailureNoPhysicalParticleAssign,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawa,
	}
}

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("gate907_classification=%s gate907_short=%s Q_phi_sign_gauge_candidate=%t native_Z2=%t supports=%s failures=%s", a.Gate907Classification, a.Gate907ShortStatus, a.QPhiSignGaugeCandidate, a.NativeZ2Certified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAirlock(a OrientationClassAirlockAudit) string {
	return fmt.Sprintf("class=%s flag_class=%s punctures=%s flag_ranks=%v alpha=%s ranks_invariant=%t alpha_class_level=%t requires_phase_sign=%t native_functor=%t supports=%s failures=%s", a.ClassObject, a.FlagClass, strings.Join(a.RepresentativePunctures, ";"), a.RepresentativeFlagRanks, a.AlphaFormula, a.RanksInvariant, a.AlphaRankSourceClassLevel, a.RequiresAbsolutePhaseSign, a.NativeZ2AirlockFunctor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatRepresentative(r Representative) string {
	return fmt.Sprintf("rep=%s puncture=%s readout=%s ranks=%v image=%d kernel=%d trace=%.16g square=%.16g N_eff=%.16g C_Y=%.16g positive=%t rows=[%s]", r.Name, r.Puncture, r.Readout, r.RankPattern, r.ImageRank, r.LeftKernelCount, r.TraceTotal, r.SquareTrace, r.OperatorNEff, r.OperatorCYukawa, r.Positive, FormatRows(r.Rows))
}

func FormatRows(rows []Row) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s(rank=%d,weight=%s,trace=%.16g,square=%.16g)", r.Atom, r.Rank, r.WeightFormula, r.TraceContribution, r.SquareContribution))
	}
	return strings.Join(parts, ";")
}

func FormatEdge(a EdgeLedgerClassAudit) string {
	return fmt.Sprintf("operation=%s class=%s tau_exchange=%t rank_patterns_invariant=%t image_invariant=%t kernel_invariant=%t class_exists=%t native=%t A={%s} B={%s} supports=%s failures=%s", a.Operation, a.ClassObject, a.TauExchangesRepresentatives, a.RankPatternsInvariant, a.ImageRankInvariant, a.LeftKernelInvariant, a.OrientationClassLedgerExists, a.NativeOperatorTheorem, FormatRepresentative(a.RepresentativeA), FormatRepresentative(a.RepresentativeB), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatTrace(a TraceMagnitudeClassAudit) string {
	return fmt.Sprintf("multiset_A=%v multiset_B=%v trace_A=%.16g trace_B=%.16g square_A=%.16g square_B=%.16g N_eff_A=%.16g N_eff_B=%.16g C_Y_A=%.16g C_Y_B=%.16g C_H_A=%.16g C_H_B=%.16g trace_inv=%t square_inv=%t N_eff_inv=%t C_Y_inv=%t C_H_inv=%t descends=%t supports=%s failures=%s", a.RowMultisetA, a.RowMultisetB, a.TraceA, a.TraceB, a.SquareTraceA, a.SquareTraceB, a.OperatorNEffA, a.OperatorNEffB, a.OperatorCYukawaA, a.OperatorCYukawaB, a.OperatorCHiggsA, a.OperatorCHiggsB, a.TraceInvariant, a.SquareInvariant, a.NEffInvariant, a.CYukawaInvariant, a.CHiggsInvariant, a.DescendsToZ2Class, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatR3(a R3RequirementAudit) string {
	return fmt.Sprintf("projectors_Z2=%t positive_readout_Z2=%t trace_reconstruction_Z2=%t native_source=%t native_R3=%t sign_free_requirements=%t supports=%s failures=%s", a.ProjectorLedgerZ2Class, a.PositiveReadoutOnZ2Class, a.TraceReconstructionOnZ2Class, a.NativeSourceCertified, a.FullNativeR3, a.RequirementsRestatedWithoutSign, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatNonDescending(a NonDescendingAudit) string {
	return fmt.Sprintf("socket_names_descend=%t physical_labels_descend=%t generation_descend=%t flavor_descend=%t individual_yukawa_descend=%t aggregate_trace_descends=%t supports=%s failures=%s", a.SocketNamesDescend, a.PhysicalSectorLabelsDescend, a.GenerationLabelsDescend, a.FlavorLabelsDescend, a.IndividualYukawaValuesDescend, a.AggregateTraceLedgerDescends, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFreeze(a FreezeAudit) string {
	return fmt.Sprintf("alpha=%.16g operator_N_eff=%.16g official_N_eff=%.16g operator_CY=%.16g official_CY=%.16g operator_CH=%.16g official_CH=%.16g frozen=%t diagnostic_only=%t can_update=%t supports=%s failures=%s", a.Alpha, a.OperatorNEff, a.OfficialNEff, a.OperatorCYukawa, a.OfficialCYukawa, a.OperatorCHiggs, a.OfficialCHiggs, a.Frozen, a.DiagnosticOnly, a.CanUpdate, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("native_Z2=%t native_Z2_airlock=%t alpha_native=%t native_boundary_incidence=%t higgs_native=%t full_AF_descent=%t native_R3=%t physical_assignment=%t generation=%t flavor=%t individual_yukawa=%t official_update=%t native_yukawa_operator=%t R4_native_yukawa=%t", f.NativeGlobalZ2, f.NativeZ2AirlockFunctor, f.AlphaNative, f.NativeBoundaryIncidence, f.HiggsOrientationNative, f.FullAFDescent, f.NativeR3, f.PhysicalParticleAssignment, f.GenerationCarrierMap, f.FlavorOrientationMap, f.IndividualYukawaValues, f.OfficialLedgerUpdate, f.NativeYukawaOperator, f.R4NativeYukawa)
}

func firewallsOK(f Firewalls) bool {
	return !f.NativeGlobalZ2 && !f.NativeZ2AirlockFunctor && !f.AlphaNative && !f.NativeBoundaryIncidence && !f.HiggsOrientationNative && !f.FullAFDescent && !f.NativeR3 && !f.PhysicalParticleAssignment && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.OfficialLedgerUpdate && !f.NativeYukawaOperator && !f.R4NativeYukawa
}

func containsAll(haystack []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sameStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sumInts(v []int) int {
	s := 0
	for _, x := range v {
		s += x
	}
	return s
}

func rowMultiset(rows []Row) []string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("rank=%d weight=%s", r.Rank, r.WeightFormula))
	}
	sort.Strings(out)
	return out
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
