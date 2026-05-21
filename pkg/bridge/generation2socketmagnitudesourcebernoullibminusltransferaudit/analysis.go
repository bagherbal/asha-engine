// Package generation2socketmagnitudesourcebernoullibminusltransferaudit implements
// Gate 865: SocketMagnitude Source and Bernoulli/B-L Transfer Audit.
//
// Gate 865 follows Gate 864's Y^dagger Y trace-magnitude readout obstruction.
// Gate 864 showed that Y^dagger Y is the correct positive right-module carrier
// but that the socket magnitudes were not derived. Gate 865 audits whether the
// required magnitudes can at least be source-typed, given sealed alpha_B, by the
// same punctured socket response law: dominant identity normalization plus a
// B-L trace-zero rest transfer. It does not derive alpha_B, observed Yukawa
// values, a sector trace ledger, or any official ledger update.
package generation2socketmagnitudesourcebernoullibminusltransferaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE865-SOCKET-MAGNITUDE-SOURCE-BERNOULLI-BMINUSL-TRANSFER-AUDIT"

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

	Classification = "SOCKET_MAGNITUDE_SOURCE_TYPING_GIVEN_ALPHA_B_AND_B_MINUS_L_TRANSFER"
	R2Status       = "R2+++++_SOCKET_MAGNITUDE_TRANSFER_SOURCE_TYPING"

	StatusGate864Inherited    = "PASS_GATE864_Y_DAGGER_Y_READOUT_OBSTRUCTION_INHERITED"
	StatusRequiredWeights     = "PASS_REQUIRED_SOCKET_MAGNITUDES_AUDITED"
	StatusDominantNorm        = "PASS_DOMINANT_IDENTITY_NORMALIZATION_AUDITED"
	StatusRestTransfer        = "PASS_REST_B_MINUS_L_TRANSFER_MAGNITUDES_RECONSTRUCTED"
	StatusBernoulliShape      = "PASS_REST_COLOR_BERNOULLI_ACTIVATION_COMPLEMENT_SHAPE_AUDITED"
	StatusTripletQuadratic    = "PASS_REST_LEPTON_TRIPLET_QUADRATIC_TRANSFER_SHAPE_AUDITED"
	StatusYdaggerYEqualsHAgg  = "PASS_Y_DAGGER_Y_EQUALS_H_AGG_GIVEN_TRANSFER_MAGNITUDES"
	StatusTracePreservation   = "PASS_REST_TRACE_PRESERVATION_REPRODUCED"
	StatusSquareTrace         = "PASS_REST_AND_TOTAL_SQUARE_TRACE_REPRODUCED"
	StatusNonCircularity      = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusWoundReducedToAlpha = "PASS_REMAINING_WOUND_REDUCED_TO_ALPHA_B_SOURCE"
	StatusLedgerFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed  = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict     = "FIREWALL_PRESERVED_GATE865_SOCKET_MAGNITUDES_NOT_R3"

	SupportGivenAlphaSocketMagnitudes = "CONDITIONAL_SUPPORT_GIVEN_ALPHA_SOCKET_MAGNITUDES_ARE_SOURCE_TYPED_BY_B_MINUS_L_TRANSFER"
	SupportRestColorBernoulli         = "CONDITIONAL_SUPPORT_REST_COLOR_MAGNITUDE_IS_ALPHA_TIMES_ONE_MINUS_ALPHA"
	SupportRestLeptonTripletQuadratic = "CONDITIONAL_SUPPORT_REST_LEPTON_MAGNITUDE_IS_TRIPLET_MULTIPLICITY_TIMES_ALPHA_SQUARED"
	SupportDominantNormalization      = "CONDITIONAL_SUPPORT_DOMINANT_COLOR_SOCKET_IS_RELATIVE_IDENTITY_NORMALIZATION"
	SupportYdaggerYEqualsHAgg         = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_EQUALS_H_AGG_GIVEN_TRANSFER_MAGNITUDES"
	SupportTraceConservation          = "CONDITIONAL_SUPPORT_B_MINUS_L_TRANSFER_PRESERVES_REST_TRACE_GIVEN_ALPHA"
	SupportSocketWoundToAlpha         = "CONDITIONAL_SUPPORT_SOCKET_MAGNITUDE_WOUND_COLLAPSES_TO_ALPHA_B_SOURCE"

	FailureAlphaStillSealed         = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureSocketNotNativeNoAlpha   = "FAILED_ROUTE_SOCKET_MAGNITUDE_SOURCE_NOT_NATIVE_WITHOUT_ALPHA_ACTIVATION_THEOREM"
	FailureDominantNormNotTopYukawa = "FAILED_ROUTE_DOMINANT_NORMALIZATION_NOT_TOP_YUKAWA_THEOREM"
	FailureSocketRestatesTransfer   = "FAILED_ROUTE_SOCKET_MAGNITUDE_ASSIGNMENT_STILL_RESTATES_TRANSFER_LAW_WITHOUT_INDEPENDENT_SOURCE"
	FailureNoNativeTransferTheorem  = "FAILED_ROUTE_B_MINUS_L_TRANSFER_LAW_NOT_NATIVE_TRACE_MAGNITUDE_THEOREM"
	FailureNoNonCircularYSource     = "FAILED_ROUTE_NO_NONCIRCULAR_SOCKET_MAGNITUDE_SOURCE_INDEPENDENT_OF_H_AGG"
	FailureNoNumericalYukawa        = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureYNotObservedYukawa       = "FAILED_ROUTE_Y_SOCKET_VALUES_NOT_OBSERVED_YUKAWA_VALUES"
	FailureNoSectorTraceMagnitude   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate    = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoFullUnbrokenAF         = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF        = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple     = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNoR3                     = "FAILED_ROUTE_R3_NOT_ALLOWED_UNTIL_SOCKET_MAGNITUDES_ARE_NONCIRCULARLY_SOURCED"
	FailureNoR4                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type SocketMagnitude struct {
	Cell, Support, SymbolicMagnitude, SourceTyping string
	Rank                                           int
	Magnitude                                      float64
	GivenAlpha, SourceTyped, NativeDerived         bool
}

type TransferLaw struct {
	Expression, BMinusLExpression, YdaggerYExpression, HAggExpression  string
	DominantNormalization, RestBMinusLTransfer, BernoulliComplement    bool
	TripletQuadraticTransfer, TraceZeroRedistribution, TracePreserving bool
	YdaggerYEqualsHAggGivenMagnitudes, NonCircular, Native             bool
	Magnitudes                                                         []SocketMagnitude
	Supports, Failures                                                 []string
}

type TraceLedger struct {
	RestTrace, RestTraceExpected               float64
	RestSquareTrace, RestSquareTraceExpected   float64
	TotalTrace, TotalSquareTrace, OperatorNEff float64
	OfficialNEff                               float64
	TracePreserved, SquareTraceMatches, Frozen bool
}

type Obstruction struct {
	LayerA_GivenAlphaSourceTyped, LayerB_AlphaDerived bool
	SocketMagnitudeNative, NonCircularSource          bool
	RemainingWound, NextGate                          string
	Supports, Failures                                []string
}

type R3Assessment struct {
	YdaggerYReadoutCarrier, GivenAlphaMagnitudesSourceTyped   bool
	AlphaNative, NativeSocketMagnitudeSource                  bool
	SectorTraceMagnitudeReadout, EligibleForR3, EligibleForR4 bool
	Supports, Failures                                        []string
}

type Impact struct {
	Classification, Status                                                           string
	CarrierSolved, GivenAlphaMagnitudeTypingSolved, NativeMagnitudeSourceSolved      bool
	AlphaNative, YukawaMagnitudes, SectorTraceReadout                                bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                                            bool
	AlphaStillSealed, SocketNotNativeNoAlpha, DominantNormNotTopYukawa, SocketRestatesTransfer, NoNativeTransferTheorem bool
	NoNonCircularYSource, NoNumericalYukawa, YNotObservedYukawa, NoSectorTraceMagnitude                                 bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoFullUnbrokenAF, AForientNotFullAF, NoNativeFiniteTriple              bool
	NoR3, NoR4                                                                                                          bool
	Verdict                                                                                                             string
}

type Audit struct {
	ID          string
	Ledger      Ledger
	Transfer    TransferLaw
	Trace       TraceLedger
	Obstruction Obstruction
	R3          R3Assessment
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func RestColorWeight(alpha float64) float64  { return alpha * (1 - alpha) }
func RestLeptonWeight(alpha float64) float64 { return 3 * alpha * alpha }
func RestTrace(alpha float64) float64        { return 3*RestColorWeight(alpha) + RestLeptonWeight(alpha) }
func RestSquareTrace(alpha float64) float64 {
	rc := RestColorWeight(alpha)
	rl := RestLeptonWeight(alpha)
	return 3*rc*rc + rl*rl
}
func TotalTrace(alpha float64) float64       { return 3 + RestTrace(alpha) }
func TotalSquareTrace(alpha float64) float64 { return 3 + RestSquareTrace(alpha) }
func OperatorNEff(alpha float64) float64 {
	return TotalTrace(alpha) * TotalTrace(alpha) / TotalSquareTrace(alpha)
}

func BuildDefault() (Audit, error) {
	magnitudes := []SocketMagnitude{
		{Cell: "e_+ tensor P_3", Support: "dominant color socket", SymbolicMagnitude: "|y_+3|^2", SourceTyping: "relative identity normalization", Rank: TopRank, Magnitude: 1, GivenAlpha: true, SourceTyped: true, NativeDerived: false},
		{Cell: "e_- tensor P_3", Support: "rest color socket", SymbolicMagnitude: "|y_-3|^2", SourceTyping: "alpha_B(1-alpha_B) Bernoulli activation/complement from rest B-L transfer", Rank: RestColorRank, Magnitude: RestColorWeight(AlphaB), GivenAlpha: true, SourceTyped: true, NativeDerived: false},
		{Cell: "e_- tensor P_1", Support: "rest lepton singleton", SymbolicMagnitude: "|y_-1|^2", SourceTyping: "3 alpha_B^2 triplet-multiplicity quadratic transfer", Rank: RestLeptonRank, Magnitude: RestLeptonWeight(AlphaB), GivenAlpha: true, SourceTyped: true, NativeDerived: false},
	}
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Transfer: TransferLaw{
			Expression:            "|y_+3|^2=1, |y_-3|^2=alpha_B(1-alpha_B), |y_-1|^2=3alpha_B^2",
			BMinusLExpression:     "H_rest/T=alpha_B P_3+alpha_B^2(3P_1-P_3)=alpha_B P_3-3alpha_B^2(B-L)",
			YdaggerYExpression:    "Y^dagger Y=1(e_+ tensor P_3)+alpha_B(1-alpha_B)(e_- tensor P_3)+3alpha_B^2(e_- tensor P_1)",
			HAggExpression:        "H_agg/T=I_{e_+ tensor P_3} plus [alpha_B P_3-3alpha_B^2(B-L)]_{e_- tensor W}",
			DominantNormalization: true, RestBMinusLTransfer: true, BernoulliComplement: true, TripletQuadraticTransfer: true,
			TraceZeroRedistribution: true, TracePreserving: true, YdaggerYEqualsHAggGivenMagnitudes: true, NonCircular: false, Native: false,
			Magnitudes: magnitudes,
			Supports:   []string{StatusRequiredWeights, StatusDominantNorm, StatusRestTransfer, StatusBernoulliShape, StatusTripletQuadratic, StatusYdaggerYEqualsHAgg, SupportGivenAlphaSocketMagnitudes, SupportRestColorBernoulli, SupportRestLeptonTripletQuadratic, SupportDominantNormalization, SupportYdaggerYEqualsHAgg, SupportTraceConservation},
			Failures:   []string{FailureAlphaStillSealed, FailureSocketNotNativeNoAlpha, FailureSocketRestatesTransfer, FailureNoNativeTransferTheorem, FailureNoNonCircularYSource, FailureDominantNormNotTopYukawa},
		},
		Trace: TraceLedger{
			RestTrace: RestTrace(AlphaB), RestTraceExpected: 3 * AlphaB, RestSquareTrace: RestSquareTrace(AlphaB), RestSquareTraceExpected: 3*AlphaB*AlphaB - 6*math.Pow(AlphaB, 3) + 12*math.Pow(AlphaB, 4),
			TotalTrace: TotalTrace(AlphaB), TotalSquareTrace: TotalSquareTrace(AlphaB), OperatorNEff: OperatorNEff(AlphaB), OfficialNEff: OfficialNEff,
			TracePreserved: true, SquareTraceMatches: true, Frozen: true,
		},
		Obstruction: Obstruction{
			LayerA_GivenAlphaSourceTyped: true, LayerB_AlphaDerived: false, SocketMagnitudeNative: false, NonCircularSource: false,
			RemainingWound: "S_split -> alpha_B / BoundaryAlphaActivationMap", NextGate: "Gate 866 — Alpha_B Source Reassessment / BoundaryAlphaActivationMap Audit",
			Supports: []string{StatusWoundReducedToAlpha, SupportSocketWoundToAlpha, SupportGivenAlphaSocketMagnitudes},
			Failures: []string{FailureAlphaStillSealed, FailureSocketNotNativeNoAlpha, FailureNoNativeTransferTheorem, FailureNoNonCircularYSource},
		},
		R3: R3Assessment{
			YdaggerYReadoutCarrier: true, GivenAlphaMagnitudesSourceTyped: true, AlphaNative: false, NativeSocketMagnitudeSource: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false,
			Supports: []string{SupportGivenAlphaSocketMagnitudes, SupportYdaggerYEqualsHAgg, SupportSocketWoundToAlpha},
			Failures: []string{FailureAlphaStillSealed, FailureSocketNotNativeNoAlpha, FailureNoSectorTraceMagnitude, FailureNoR3, FailureNoR4},
		},
		Impact:    Impact{Classification: Classification, Status: R2Status, CarrierSolved: true, GivenAlphaMagnitudeTypingSolved: true, NativeMagnitudeSourceSolved: false, AlphaNative: false, YukawaMagnitudes: false, SectorTraceReadout: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false},
		Firewalls: Firewalls{Enforced: true, AlphaStillSealed: true, SocketNotNativeNoAlpha: true, DominantNormNotTopYukawa: true, SocketRestatesTransfer: true, NoNativeTransferTheorem: true, NoNonCircularYSource: true, NoNumericalYukawa: true, YNotObservedYukawa: true, NoSectorTraceMagnitude: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NoR3: true, NoR4: true, Verdict: StatusFirewallVerdict},
		Truth:     "Gate 865 source-types the required Y^dagger Y socket magnitudes by the B-L rest-transfer law given sealed alpha_B, but does not derive alpha_B or the transfer law natively.",
		Final:     "The Y^dagger Y carrier and given-alpha socket magnitude pattern now align; the remaining noncircular wound is the native source of alpha_B / BoundaryAlphaActivationMap.",
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
	if !a.Transfer.DominantNormalization || !a.Transfer.RestBMinusLTransfer || !a.Transfer.BernoulliComplement || !a.Transfer.TripletQuadraticTransfer || !a.Transfer.TraceZeroRedistribution || !a.Transfer.TracePreserving || !a.Transfer.YdaggerYEqualsHAggGivenMagnitudes || a.Transfer.NonCircular || a.Transfer.Native {
		return err("transfer law malformed or overpromoted")
	}
	if len(a.Transfer.Magnitudes) != 3 || !magnitudesOK(a.Transfer.Magnitudes) {
		return err("socket magnitudes malformed")
	}
	if !near(a.Trace.RestTrace, a.Trace.RestTraceExpected) || !near(a.Trace.RestTrace, 3*AlphaB) || !near(a.Trace.RestSquareTrace, a.Trace.RestSquareTraceExpected) || !near(a.Trace.TotalTrace, TotalTrace(AlphaB)) || !near(a.Trace.TotalSquareTrace, TotalSquareTrace(AlphaB)) || !near(a.Trace.OperatorNEff, OperatorNEff(AlphaB)) || !a.Trace.TracePreserved || !a.Trace.SquareTraceMatches || !a.Trace.Frozen {
		return err("trace ledger mismatch")
	}
	if !a.Obstruction.LayerA_GivenAlphaSourceTyped || a.Obstruction.LayerB_AlphaDerived || a.Obstruction.SocketMagnitudeNative || a.Obstruction.NonCircularSource || a.Obstruction.RemainingWound == "" {
		return err("obstruction malformed")
	}
	if !a.R3.YdaggerYReadoutCarrier || !a.R3.GivenAlphaMagnitudesSourceTyped || a.R3.AlphaNative || a.R3.NativeSocketMagnitudeSource || a.R3.SectorTraceMagnitudeReadout || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		return err("R3 assessment malformed or overpromoted")
	}
	if !a.Impact.CarrierSolved || !a.Impact.GivenAlphaMagnitudeTypingSolved || a.Impact.NativeMagnitudeSourceSolved || a.Impact.AlphaNative || a.Impact.YukawaMagnitudes || a.Impact.SectorTraceReadout || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact malformed or overpromoted")
	}
	if !firewallsOK(a.Firewalls) {
		return err("firewalls malformed")
	}
	return nil
}

func magnitudesOK(ms []SocketMagnitude) bool {
	if len(ms) != 3 {
		return false
	}
	return ms[0].Cell == "e_+ tensor P_3" && ms[0].Rank == TopRank && near(ms[0].Magnitude, 1) && ms[0].SourceTyped && !ms[0].NativeDerived &&
		ms[1].Cell == "e_- tensor P_3" && ms[1].Rank == RestColorRank && near(ms[1].Magnitude, RestColorWeight(AlphaB)) && ms[1].SourceTyped && !ms[1].NativeDerived &&
		ms[2].Cell == "e_- tensor P_1" && ms[2].Rank == RestLeptonRank && near(ms[2].Magnitude, RestLeptonWeight(AlphaB)) && ms[2].SourceTyped && !ms[2].NativeDerived
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.SocketNotNativeNoAlpha && f.DominantNormNotTopYukawa && f.SocketRestatesTransfer && f.NoNativeTransferTheorem && f.NoNonCircularYSource && f.NoNumericalYukawa && f.YNotObservedYukawa && f.NoSectorTraceMagnitude && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NoR3 && f.NoR4 && f.Verdict == StatusFirewallVerdict
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }

func containsAll(have, want []string) bool {
	set := map[string]bool{}
	for _, h := range have {
		set[h] = true
	}
	for _, w := range want {
		if !set[w] {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatMagnitudes(ms []SocketMagnitude) string {
	parts := make([]string, 0, len(ms))
	for _, m := range ms {
		parts = append(parts, fmt.Sprintf("%s[%s,%s,rank=%d,mag=%.16g,given_alpha=%t,source_typed=%t,native=%t]", m.Cell, m.SymbolicMagnitude, m.SourceTyping, m.Rank, m.Magnitude, m.GivenAlpha, m.SourceTyped, m.NativeDerived))
	}
	return strings.Join(parts, "; ")
}

func FormatTransfer(t TransferLaw) string {
	return fmt.Sprintf("expr=%s | B-L=%s | YdaggerY=%s | Hagg=%s | dominant_norm=%t rest_transfer=%t bernoulli=%t triplet_quadratic=%t trace_preserving=%t YdaggerY_equals_Hagg_given_magnitudes=%t native=%t noncircular=%t | magnitudes={%s} | failures=%s", t.Expression, t.BMinusLExpression, t.YdaggerYExpression, t.HAggExpression, t.DominantNormalization, t.RestBMinusLTransfer, t.BernoulliComplement, t.TripletQuadraticTransfer, t.TracePreserving, t.YdaggerYEqualsHAggGivenMagnitudes, t.Native, t.NonCircular, FormatMagnitudes(t.Magnitudes), strings.Join(t.Failures, ","))
}

func FormatTrace(t TraceLedger) string {
	return fmt.Sprintf("rest_trace=%.16g expected=%.16g rest_square=%.16g expected=%.16g total_trace=%.16g total_square=%.16g operator_N_eff=%.16g official_N_eff=%.16g trace_preserved=%t square_trace_matches=%t frozen=%t", t.RestTrace, t.RestTraceExpected, t.RestSquareTrace, t.RestSquareTraceExpected, t.TotalTrace, t.TotalSquareTrace, t.OperatorNEff, t.OfficialNEff, t.TracePreserved, t.SquareTraceMatches, t.Frozen)
}

func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("given_alpha_source_typed=%t alpha_derived=%t socket_native=%t noncircular=%t remaining_wound=%s next=%s failures=%s", o.LayerA_GivenAlphaSourceTyped, o.LayerB_AlphaDerived, o.SocketMagnitudeNative, o.NonCircularSource, o.RemainingWound, o.NextGate, strings.Join(o.Failures, ","))
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("YdaggerY_carrier=%t given_alpha_magnitudes_source_typed=%t alpha_native=%t native_socket_source=%t sector_trace=%t R3=%t R4=%t failures=%s", r.YdaggerYReadoutCarrier, r.GivenAlphaMagnitudesSourceTyped, r.AlphaNative, r.NativeSocketMagnitudeSource, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s carrier_solved=%t given_alpha_typing=%t native_magnitude_source=%t alpha_native=%t yukawa_magnitudes=%t sector_trace=%t update_N_eff=%t update_CYukawa=%t update_CHiggs=%t promote_R3=%t promote_R4=%t", i.Classification, i.Status, i.CarrierSolved, i.GivenAlphaMagnitudeTypingSolved, i.NativeMagnitudeSourceSolved, i.AlphaNative, i.YukawaMagnitudes, i.SectorTraceReadout, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t alpha_sealed=%t socket_not_native_no_alpha=%t dominant_norm_not_top=%t restates_transfer=%t no_native_transfer=%t no_noncircular_y_source=%t no_yukawa=%t no_sector_trace=%t no_update_N_eff=%t no_update_C=%t no_full_AF=%t AForient_not_full=%t no_native_finite_triple=%t no_R3=%t no_R4=%t verdict=%s", f.Enforced, f.AlphaStillSealed, f.SocketNotNativeNoAlpha, f.DominantNormNotTopYukawa, f.SocketRestatesTransfer, f.NoNativeTransferTheorem, f.NoNonCircularYSource, f.NoNumericalYukawa, f.NoSectorTraceMagnitude, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoFullUnbrokenAF, f.AForientNotFullAF, f.NoNativeFiniteTriple, f.NoR3, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate864Inherited, StatusRequiredWeights, StatusDominantNorm, StatusRestTransfer, StatusBernoulliShape, StatusTripletQuadratic,
		StatusYdaggerYEqualsHAgg, StatusTracePreservation, StatusSquareTrace, StatusNonCircularity, StatusWoundReducedToAlpha,
		StatusLedgerFrozen, StatusNoObservedDataUsed, StatusFirewallVerdict,
		SupportGivenAlphaSocketMagnitudes, SupportRestColorBernoulli, SupportRestLeptonTripletQuadratic, SupportDominantNormalization,
		SupportYdaggerYEqualsHAgg, SupportTraceConservation, SupportSocketWoundToAlpha,
		FailureAlphaStillSealed, FailureSocketNotNativeNoAlpha, FailureDominantNormNotTopYukawa, FailureSocketRestatesTransfer,
		FailureNoNativeTransferTheorem, FailureNoNonCircularYSource, FailureNoNumericalYukawa, FailureYNotObservedYukawa,
		FailureNoSectorTraceMagnitude, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoFullUnbrokenAF,
		FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNoR3, FailureNoR4,
	}
}
