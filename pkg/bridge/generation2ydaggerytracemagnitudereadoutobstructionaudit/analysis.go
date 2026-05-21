// Package generation2ydaggerytracemagnitudereadoutobstructionaudit implements
// Gate 864: Y^dagger Y TraceMagnitude Readout Obstruction Audit.
//
// Gate 864 follows Gate 863's post-orientation finite-triple seal
// classification. The branch now has an oriented stabilizer algebra, a minimal
// 15/30 carrier, a scalar operator-valued edge socket matrix Y, color
// centrality, and a socket-character seal. This gate audits the natural
// positive right-module readout Y^dagger Y and asks whether it can reproduce the
// aggregate response table without inserting the required socket magnitudes by
// hand. It certifies the carrier shape and positivity of the candidate, while
// preserving the obstruction that the socket magnitudes and alpha_B source are
// not derived.
package generation2ydaggerytracemagnitudereadoutobstructionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE864-Y-DAGGER-Y-TRACE-MAGNITUDE-READOUT-OBSTRUCTION-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	TopRank        = 3
	RestColorRank  = 3
	RestLeptonRank = 1
	ActiveRank     = 7
	PunctureRank   = 1
	LeftKernelRank = 1

	Classification = "Y_DAGGER_Y_TRACE_MAGNITUDE_READOUT_OBSTRUCTION"
	R2Status       = "R2+++++_Y_DAGGER_Y_READOUT_OBSTRUCTION"

	StatusGate863Inherited    = "PASS_GATE863_POST_ORIENTATION_FINITE_TRIPLE_SEAL_INHERITED"
	StatusYdaggerYConstructed = "PASS_Y_DAGGER_Y_POSITIVE_RIGHT_READOUT_CONSTRUCTED"
	StatusCorrectSupport      = "PASS_Y_DAGGER_Y_HAS_CORRECT_ACTIVE_SUPPORT"
	StatusPositive            = "PASS_Y_DAGGER_Y_IS_POSITIVE_ON_H_R_MIN"
	StatusPunctureAbsent      = "PASS_PUNCTURE_CELL_REMAINS_ABSENT"
	StatusLeftKernelExcluded  = "PASS_LEFT_KERNEL_DOES_NOT_ENTER_RIGHT_TRACE_READOUT"
	StatusRequiredWeights     = "PASS_REQUIRED_SOCKET_MAGNITUDES_FOR_H_AGG_MATCH_COMPUTED"
	StatusConditionalTrace    = "PASS_CONDITIONAL_TRACE_AND_SQUARE_TRACE_MATCH_H_AGG_IF_VALUES_INSERTED"
	StatusReadoutObstructed   = "PASS_TRACE_MAGNITUDE_READOUT_OBSTRUCTION_IDENTIFIED"
	StatusR3PressureReduced   = "PASS_R3_PRESSURE_REDUCED_TO_SOCKET_MAGNITUDE_SOURCE"
	StatusLedgerFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed  = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict     = "FIREWALL_PRESERVED_GATE864_Y_DAGGER_Y_NOT_R3"

	SupportYdaggerYCandidate       = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_IS_THE_CORRECT_POSITIVE_READOUT_CANDIDATE"
	SupportCarrierWiseYes          = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_HAS_CORRECT_CARRIER_SHAPE"
	SupportRequiredSocketValues    = "CONDITIONAL_SUPPORT_REQUIRED_SOCKET_MAGNITUDES_MATCH_AGGREGATE_TABLE_IF_SET"
	SupportTraceMatchesIfInserted  = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_REPRODUCES_H_AGG_IF_SOCKET_VALUES_INSERTED"
	SupportR3PressureSocketSource  = "CONDITIONAL_SUPPORT_R3_PRESSURE_REDUCES_TO_SOCKET_MAGNITUDE_SOURCE"
	SupportPunctureKernelPreserved = "CONDITIONAL_SUPPORT_PUNCTURE_AND_LEFT_KERNEL_FIREWALLS_PRESERVED_IN_RIGHT_READOUT"

	FailureYMagnitudesNotDerived    = "FAILED_ROUTE_Y_SOCKET_MAGNITUDES_NOT_DERIVED"
	FailureAlphaStillSealed         = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureOnlyIfValuesInserted     = "FAILED_ROUTE_Y_DAGGER_Y_REPRODUCES_H_AGG_ONLY_IF_SOCKET_VALUES_INSERTED"
	FailureReadoutNotNative         = "FAILED_ROUTE_TRACE_MAGNITUDE_READOUT_NOT_NATIVE_WITHOUT_SOCKET_MAGNITUDE_SOURCE"
	FailureYToAggNoTheorem          = "FAILED_ROUTE_NO_Y_DAGGER_Y_TO_H_AGG_TRACE_MAGNITUDE_THEOREM"
	FailureSocketValuesRestateTable = "FAILED_ROUTE_SOCKET_VALUE_ASSIGNMENT_RESTATES_AGGREGATE_TABLE"
	FailureNoNumericalYukawa        = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureYNotObservedYukawa       = "FAILED_ROUTE_Y_SOCKET_VALUES_NOT_OBSERVED_YUKAWA_VALUES"
	FailureNoSectorTraceMagnitude   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate    = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoFullUnbrokenAF         = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF        = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple     = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNoR3                     = "FAILED_ROUTE_Y_DAGGER_Y_READOUT_OBSTRUCTION_NOT_R3"
	FailureNoR4                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type SocketWeight struct {
	Cell, Support, SymbolicMagnitude string
	Rank                             int
	RequiredMagnitude                float64
	Derived, InsertedForMatch        bool
}

type YDaggerYReadout struct {
	Expression, TargetExpression                                                                string
	Positive, CorrectActiveSupport, PunctureAbsent, LeftKernelExcluded, RightModuleReadout      bool
	CarrierWiseMatch, MagnitudeWiseMatch, SocketMagnitudesDerived, RequiresInsertedSocketValues bool
	TraceIfMatched, SquareTraceIfMatched, OperatorNEffIfMatched, OfficialNEff                   float64
	Weights                                                                                     []SocketWeight
	Supports, Failures                                                                          []string
}

type Obstruction struct {
	CarrierShapePasses, PositivityPasses, PunctureFirewallPasses bool
	YSocketMagnitudeSourceMissing, AlphaSourceMissing            bool
	TraceReadoutNative, NonCircularMagnitudes                    bool
	NextMissingObject, NextGate                                  string
	Supports, Failures                                           []string
}

type R3Assessment struct {
	PostOrientationFiniteTripleSeal, YDaggerYCandidate, CarrierWiseYes bool
	MagnitudeWiseNo, SocketMagnitudesDerived, AlphaNative              bool
	SectorTraceLedgerPresent, EligibleForR3, EligibleForR4             bool
	Supports, Failures                                                 []string
}

type Impact struct {
	Classification, Status                                                                   string
	YdaggerYPositiveCandidate, CarrierShapeSolved, MagnitudeSourceSolved, NativeTraceReadout bool
	AlphaNative, YukawaMagnitudes, SectorTraceReadout                                        bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4         bool
}

type Firewalls struct {
	Enforced                                                                                               bool
	YMagnitudesNotDerived, AlphaStillSealed, OnlyIfValuesInserted, ReadoutNotNative, YToAggNoTheorem       bool
	SocketValuesRestateTable, NoNumericalYukawa, YNotObservedYukawa, NoSectorTraceMagnitude                bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoFullUnbrokenAF, AForientNotFullAF, NoNativeFiniteTriple bool
	NoR3, NoR4                                                                                             bool
	Verdict                                                                                                string
}

type Audit struct {
	ID          string
	Ledger      Ledger
	Readout     YDaggerYReadout
	Obstruction Obstruction
	R3          R3Assessment
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func RestColorWeight(alpha float64) float64  { return alpha * (1 - alpha) }
func RestLeptonWeight(alpha float64) float64 { return 3 * alpha * alpha }
func TraceIfMatched(alpha float64) float64 {
	return 3 + 3*alpha
}
func SquareTraceIfMatched(alpha float64) float64 {
	return 3 + 3*alpha*alpha - 6*alpha*alpha*alpha + 12*math.Pow(alpha, 4)
}
func OperatorNEffIfMatched(alpha float64) float64 {
	return TraceIfMatched(alpha) * TraceIfMatched(alpha) / SquareTraceIfMatched(alpha)
}

func BuildDefault() (Audit, error) {
	weights := []SocketWeight{
		{Cell: "e_+ tensor P_3", Support: "dominant color socket", SymbolicMagnitude: "|y_+3|^2", Rank: TopRank, RequiredMagnitude: 1.0, Derived: false, InsertedForMatch: true},
		{Cell: "e_- tensor P_3", Support: "rest color socket", SymbolicMagnitude: "|y_-3|^2", Rank: RestColorRank, RequiredMagnitude: RestColorWeight(AlphaB), Derived: false, InsertedForMatch: true},
		{Cell: "e_- tensor P_1", Support: "rest lepton socket", SymbolicMagnitude: "|y_-1|^2", Rank: RestLeptonRank, RequiredMagnitude: RestLeptonWeight(AlphaB), Derived: false, InsertedForMatch: true},
	}
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Readout: YDaggerYReadout{
			Expression:                   "Y^dagger Y=|y_+3|^2(e_+ tensor P_3)+|y_-3|^2(e_- tensor P_3)+|y_-1|^2(e_- tensor P_1)",
			TargetExpression:             "H_agg/T=1(e_+ tensor P_3)+alpha_B(1-alpha_B)(e_- tensor P_3)+3alpha_B^2(e_- tensor P_1)",
			Positive:                     true,
			CorrectActiveSupport:         true,
			PunctureAbsent:               true,
			LeftKernelExcluded:           true,
			RightModuleReadout:           true,
			CarrierWiseMatch:             true,
			MagnitudeWiseMatch:           false,
			SocketMagnitudesDerived:      false,
			RequiresInsertedSocketValues: true,
			TraceIfMatched:               TraceIfMatched(AlphaB),
			SquareTraceIfMatched:         SquareTraceIfMatched(AlphaB),
			OperatorNEffIfMatched:        OperatorNEffIfMatched(AlphaB),
			OfficialNEff:                 OfficialNEff,
			Weights:                      weights,
			Supports: []string{
				StatusYdaggerYConstructed, StatusCorrectSupport, StatusPositive, StatusPunctureAbsent, StatusLeftKernelExcluded,
				StatusRequiredWeights, StatusConditionalTrace, SupportYdaggerYCandidate, SupportCarrierWiseYes,
				SupportRequiredSocketValues, SupportTraceMatchesIfInserted, SupportPunctureKernelPreserved,
			},
			Failures: []string{FailureYMagnitudesNotDerived, FailureAlphaStillSealed, FailureOnlyIfValuesInserted, FailureReadoutNotNative, FailureYToAggNoTheorem, FailureSocketValuesRestateTable},
		},
		Obstruction: Obstruction{
			CarrierShapePasses: true, PositivityPasses: true, PunctureFirewallPasses: true,
			YSocketMagnitudeSourceMissing: true, AlphaSourceMissing: true, TraceReadoutNative: false, NonCircularMagnitudes: false,
			NextMissingObject: "SocketMagnitudeSource(|y_+3|^2, |y_-3|^2, |y_-1|^2)",
			NextGate:          "Gate 865 — SocketMagnitude Source / Alpha_B Compatibility Audit",
			Supports:          []string{StatusReadoutObstructed, StatusR3PressureReduced, SupportR3PressureSocketSource},
			Failures:          []string{FailureYMagnitudesNotDerived, FailureAlphaStillSealed, FailureSocketValuesRestateTable, FailureReadoutNotNative, FailureNoSectorTraceMagnitude},
		},
		R3: R3Assessment{
			PostOrientationFiniteTripleSeal: true, YDaggerYCandidate: true, CarrierWiseYes: true, MagnitudeWiseNo: true,
			SocketMagnitudesDerived: false, AlphaNative: false, SectorTraceLedgerPresent: false, EligibleForR3: false, EligibleForR4: false,
			Supports: []string{SupportYdaggerYCandidate, SupportCarrierWiseYes, SupportR3PressureSocketSource},
			Failures: []string{FailureYMagnitudesNotDerived, FailureNoSectorTraceMagnitude, FailureAlphaStillSealed, FailureNoR3, FailureNoR4},
		},
		Impact: Impact{
			Classification: Classification, Status: R2Status, YdaggerYPositiveCandidate: true, CarrierShapeSolved: true,
			MagnitudeSourceSolved: false, NativeTraceReadout: false, AlphaNative: false, YukawaMagnitudes: false, SectorTraceReadout: false,
			CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false,
		},
		Firewalls: Firewalls{
			Enforced: true, YMagnitudesNotDerived: true, AlphaStillSealed: true, OnlyIfValuesInserted: true, ReadoutNotNative: true, YToAggNoTheorem: true,
			SocketValuesRestateTable: true, NoNumericalYukawa: true, YNotObservedYukawa: true, NoSectorTraceMagnitude: true,
			NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true,
			NoR3: true, NoR4: true, Verdict: StatusFirewallVerdict,
		},
		Truth: "Y^dagger Y is the correct positive carrier-shape candidate on H_R^min, but it reproduces H_agg/T only after inserting socket magnitudes.",
		Final: "Gate 864 reduces the R3 wound to the missing source of |y_+3|^2, |y_-3|^2, and |y_-1|^2; it does not derive alpha_B or permit any ledger update.",
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if a.ID != AuditID || a.Ledger.AlphaB != AlphaB || !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger malformed or overpromoted")
	}
	if !a.Readout.Positive || !a.Readout.CorrectActiveSupport || !a.Readout.PunctureAbsent || !a.Readout.LeftKernelExcluded || !a.Readout.RightModuleReadout || !a.Readout.CarrierWiseMatch || a.Readout.MagnitudeWiseMatch || a.Readout.SocketMagnitudesDerived || !a.Readout.RequiresInsertedSocketValues {
		return err("readout malformed or overpromoted")
	}
	if len(a.Readout.Weights) != 3 || !weightsOK(a.Readout.Weights) {
		return err("required socket magnitude ledger malformed")
	}
	if !near(a.Readout.TraceIfMatched, TraceIfMatched(AlphaB)) || !near(a.Readout.SquareTraceIfMatched, SquareTraceIfMatched(AlphaB)) || !near(a.Readout.OperatorNEffIfMatched, OperatorNEffIfMatched(AlphaB)) {
		return err("conditional trace reconstruction mismatch")
	}
	if !a.Obstruction.CarrierShapePasses || !a.Obstruction.PositivityPasses || !a.Obstruction.PunctureFirewallPasses || !a.Obstruction.YSocketMagnitudeSourceMissing || !a.Obstruction.AlphaSourceMissing || a.Obstruction.TraceReadoutNative || a.Obstruction.NonCircularMagnitudes {
		return err("obstruction malformed")
	}
	if !a.R3.PostOrientationFiniteTripleSeal || !a.R3.YDaggerYCandidate || !a.R3.CarrierWiseYes || !a.R3.MagnitudeWiseNo || a.R3.SocketMagnitudesDerived || a.R3.AlphaNative || a.R3.SectorTraceLedgerPresent || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		return err("R3 assessment malformed or overpromoted")
	}
	if a.Impact.Classification != Classification || a.Impact.Status != R2Status || !a.Impact.YdaggerYPositiveCandidate || !a.Impact.CarrierShapeSolved || a.Impact.MagnitudeSourceSolved || a.Impact.NativeTraceReadout || a.Impact.AlphaNative || a.Impact.YukawaMagnitudes || a.Impact.SectorTraceReadout || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact malformed or overpromoted")
	}
	if !firewallsOK(a.Firewalls) {
		return err("firewalls not preserved")
	}
	return nil
}

func weightsOK(w []SocketWeight) bool {
	want := map[string]float64{
		"e_+ tensor P_3": 1.0,
		"e_- tensor P_3": RestColorWeight(AlphaB),
		"e_- tensor P_1": RestLeptonWeight(AlphaB),
	}
	ranks := map[string]int{"e_+ tensor P_3": TopRank, "e_- tensor P_3": RestColorRank, "e_- tensor P_1": RestLeptonRank}
	for _, x := range w {
		v, ok := want[x.Cell]
		if !ok || !near(x.RequiredMagnitude, v) || ranks[x.Cell] != x.Rank || x.Derived || !x.InsertedForMatch {
			return false
		}
	}
	return true
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.YMagnitudesNotDerived && f.AlphaStillSealed && f.OnlyIfValuesInserted && f.ReadoutNotNative && f.YToAggNoTheorem && f.SocketValuesRestateTable && f.NoNumericalYukawa && f.YNotObservedYukawa && f.NoSectorTraceMagnitude && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NoR3 && f.NoR4 && f.Verdict == StatusFirewallVerdict
}

func Statuses() []string {
	return []string{
		StatusGate863Inherited, StatusYdaggerYConstructed, StatusCorrectSupport, StatusPositive, StatusPunctureAbsent,
		StatusLeftKernelExcluded, StatusRequiredWeights, StatusConditionalTrace, StatusReadoutObstructed, StatusR3PressureReduced,
		StatusLedgerFrozen, StatusNoObservedDataUsed, StatusFirewallVerdict,
		SupportYdaggerYCandidate, SupportCarrierWiseYes, SupportRequiredSocketValues, SupportTraceMatchesIfInserted,
		SupportR3PressureSocketSource, SupportPunctureKernelPreserved,
		FailureYMagnitudesNotDerived, FailureAlphaStillSealed, FailureOnlyIfValuesInserted, FailureReadoutNotNative,
		FailureYToAggNoTheorem, FailureSocketValuesRestateTable, FailureNoNumericalYukawa, FailureYNotObservedYukawa,
		FailureNoSectorTraceMagnitude, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoFullUnbrokenAF,
		FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNoR3, FailureNoR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatWeights(w []SocketWeight) string {
	parts := make([]string, 0, len(w))
	for _, x := range w {
		parts = append(parts, fmt.Sprintf("%s rank=%d symbolic=%s required=%.16g derived=%t inserted_for_match=%t", x.Cell, x.Rank, x.SymbolicMagnitude, x.RequiredMagnitude, x.Derived, x.InsertedForMatch))
	}
	return strings.Join(parts, " | ")
}

func FormatReadout(r YDaggerYReadout) string {
	return fmt.Sprintf("expr=%s target=%s positive=%t support=%t carrier_match=%t magnitude_match=%t magnitudes_derived=%t requires_inserted=%t trace_if_matched=%.16g square_trace_if_matched=%.16g operator_N_eff_if_matched=%.16g official_N_eff=%.16g weights=[%s] supports=%s failures=%s", r.Expression, r.TargetExpression, r.Positive, r.CorrectActiveSupport, r.CarrierWiseMatch, r.MagnitudeWiseMatch, r.SocketMagnitudesDerived, r.RequiresInsertedSocketValues, r.TraceIfMatched, r.SquareTraceIfMatched, r.OperatorNEffIfMatched, r.OfficialNEff, FormatWeights(r.Weights), strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("carrier_shape=%t positivity=%t puncture_firewall=%t y_magnitude_source_missing=%t alpha_source_missing=%t native_readout=%t noncircular_magnitudes=%t next=%s gate=%s supports=%s failures=%s", o.CarrierShapePasses, o.PositivityPasses, o.PunctureFirewallPasses, o.YSocketMagnitudeSourceMissing, o.AlphaSourceMissing, o.TraceReadoutNative, o.NonCircularMagnitudes, o.NextMissingObject, o.NextGate, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("post_orientation_seal=%t ydagger_candidate=%t carrier_yes=%t magnitude_no=%t socket_magnitudes_derived=%t alpha_native=%t sector_trace_ledger=%t R3=%t R4=%t supports=%s failures=%s", r.PostOrientationFiniteTripleSeal, r.YDaggerYCandidate, r.CarrierWiseYes, r.MagnitudeWiseNo, r.SocketMagnitudesDerived, r.AlphaNative, r.SectorTraceLedgerPresent, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s ydagger_candidate=%t carrier_shape_solved=%t magnitude_source_solved=%t native_readout=%t alpha_native=%t yukawa_magnitudes=%t sector_trace=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t R3=%t R4=%t", i.Classification, i.Status, i.YdaggerYPositiveCandidate, i.CarrierShapeSolved, i.MagnitudeSourceSolved, i.NativeTraceReadout, i.AlphaNative, i.YukawaMagnitudes, i.SectorTraceReadout, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t y_magnitudes_not_derived=%t alpha_sealed=%t only_if_values_inserted=%t readout_not_native=%t y_to_agg_no_theorem=%t values_restate_table=%t no_numerical_yukawa=%t y_not_observed_yukawa=%t no_trace_magnitude=%t no_N_eff_update=%t no_C_updates=%t no_full_AF=%t AForient_not_full_AF=%t no_native_finite_triple=%t no_R3=%t no_R4=%t verdict=%s", f.Enforced, f.YMagnitudesNotDerived, f.AlphaStillSealed, f.OnlyIfValuesInserted, f.ReadoutNotNative, f.YToAggNoTheorem, f.SocketValuesRestateTable, f.NoNumericalYukawa, f.YNotObservedYukawa, f.NoSectorTraceMagnitude, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoFullUnbrokenAF, f.AForientNotFullAF, f.NoNativeFiniteTriple, f.NoR3, f.NoR4, f.Verdict)
}

func containsAll(haystack, needles []string) bool {
	set := make(map[string]bool, len(haystack))
	for _, h := range haystack {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}
