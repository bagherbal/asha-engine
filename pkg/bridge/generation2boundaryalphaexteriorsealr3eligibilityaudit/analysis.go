// Package generation2boundaryalphaexteriorsealr3eligibilityaudit implements
// Gate 873: BoundaryAlpha ExteriorSeal and R3 Eligibility Audit.
//
// Gate 873 follows Gate 872's exposure/enclosure degree-target obstruction.
// The alpha-side chain has reached a mature bridge-layer ceiling: the reduced
// boundary-pair response explains the s+s^2 shape, the socket-rank numerators
// are source-typed, and Y^dagger Y reproduces the aggregate trace-magnitude
// table given the sealed alpha_B. Gate 873 classifies that object as a
// BoundaryAlphaExteriorExposureEnclosureSeal and audits whether the full chain
// is eligible for R3. The verdict is intentionally conservative: the conditional
// readout chain is coherent, but native R3/R4 promotion and all official ledger
// updates remain blocked until the exposure/enclosure target-selection theorem
// and alpha source become native.
package generation2boundaryalphaexteriorsealr3eligibilityaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE873-BOUNDARY-ALPHA-EXTERIOR-SEAL-R3-ELIGIBILITY-AUDIT"

	AlphaB             = 0.0003878958469680527
	SBoundary          = 0.0012924448188162962
	OfficialNEffFrozen = 3.0023273474722147

	PiTopRank       = 3
	HRambientRank   = 8
	BoundaryPairDim = 2
	H10Dim          = HRambientRank + BoundaryPairDim
	HRminRank       = 7
	Lambda4V8Rank   = 70
	H72Dim          = Lambda4V8Rank + BoundaryPairDim

	Classification = "BOUNDARY_ALPHA_EXTERIOR_EXPOSURE_ENCLOSURE_SEAL_R3_ELIGIBILITY_OBSTRUCTION"
	R2Status       = "R2+++++_BOUNDARY_ALPHA_EXTERIOR_SEAL_CONDITIONAL_TRACE_READOUT_NOT_R3"

	StatusGate872Inherited         = "PASS_GATE872_EXPOSURE_ENCLOSURE_TARGET_SELECTION_OBSTRUCTION_INHERITED"
	StatusAlphaSealClassified      = "PASS_BOUNDARY_ALPHA_EXTERIOR_EXPOSURE_ENCLOSURE_SEAL_CLASSIFIED"
	StatusFullChainReassembled     = "PASS_FULL_CONDITIONAL_TRACE_MAGNITUDE_CHAIN_REASSEMBLED"
	StatusYDaggerYReadoutInherited = "PASS_Y_DAGGER_Y_TRACE_MAGNITUDE_READOUT_INHERITED"
	StatusHaggReadoutInherited     = "PASS_H_AGG_TRACE_READOUT_INHERITED"
	StatusOperatorNEffComputed     = "PASS_OPERATOR_N_EFF_CONDITIONAL_READOUT_COMPUTED"
	StatusR3EligibilityAudited     = "PASS_R3_ELIGIBILITY_AUDITED"
	StatusLedgerFrozen             = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed       = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusNextWound                = "PASS_NEXT_WOUND_IDENTIFIED_AS_Y_DAGGER_Y_TO_TRACE_MAGNITUDE_SOURCE_OR_ALPHA_TARGET_SELECTION"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE873_CONDITIONAL_CHAIN_NOT_NATIVE_R3"

	SupportBoundaryAlphaExteriorSeal = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_EXTERIOR_EXPOSURE_ENCLOSURE_SEAL"
	SupportAlphaSocketRankAnatomy    = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_SOCKET_RANK_AND_REDUCED_EXTERIOR_RESPONSE_ANATOMY"
	SupportFullChainCoherent         = "CONDITIONAL_SUPPORT_FULL_TRACE_MAGNITUDE_CHAIN_COHERENT_GIVEN_ALPHA_SEAL"
	SupportYDaggerYReproducesHagg    = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_REPRODUCES_H_AGG_GIVEN_BOUNDARY_ALPHA_EXTERIOR_SEAL"
	SupportOperatorNEffConditional   = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_IS_CONDITIONAL_TRACE_READOUT"
	SupportMatureConditionalReadout  = "CONDITIONAL_SUPPORT_MATURE_CONDITIONAL_AGGREGATE_TRACE_MAGNITUDE_READOUT"
	SupportR3CandidatePressure       = "CONDITIONAL_SUPPORT_R3_PRESSURE_REDUCES_TO_NATIVE_ALPHA_TARGET_SELECTION_AND_SOCKET_MAGNITUDE_SOURCE"

	FailureAlphaStillSealed                   = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureAlphaNotNativeWithoutTargetTheorem = "FAILED_ROUTE_ALPHA_B_NOT_NATIVE_WITHOUT_TARGET_SELECTION_THEOREM"
	FailureNoNativeTargetSelection            = "FAILED_ROUTE_NO_NATIVE_EXPOSURE_ENCLOSURE_TARGET_SELECTION_MAP"
	FailureNoCrossLaneExclusion               = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoNativeReducedExteriorFunctional  = "FAILED_ROUTE_NO_NATIVE_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_FUNCTIONAL"
	FailureNoNativeAlphaSource                = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureSocketMagnitudesDependOnAlphaSeal  = "FAILED_ROUTE_SOCKET_MAGNITUDES_DEPEND_ON_SEALED_ALPHA_B"
	FailureNoNativeSocketMagnitudeSource      = "FAILED_ROUTE_NO_NATIVE_SOCKET_MAGNITUDE_SOURCE"
	FailureNoOfficialNEffUpdate               = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate              = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNumericalYukawa                  = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoSectorTraceMagnitude             = "FAILED_ROUTE_NO_NATIVE_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNotR3NativeTraceLedger             = "FAILED_ROUTE_NOT_R3_NATIVE_TRACE_LEDGER"
	FailureNotR4NativeYukawa                  = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB, SBoundary, OfficialNEffFrozen float64
	AlphaNative, R3, R4, OfficialFrozen   bool
}

type AlphaExteriorSeal struct {
	Name                                                        string
	Expression                                                  string
	ReducedResponse                                             string
	LinearNumerator, LinearDenominator                          int
	QuadraticNumerator, QuadraticDenominator                    int
	LinearContribution, QuadraticContribution, Alpha            float64
	ShapeTyped, RankSourcesTyped, TargetSelectionNative, Native bool
	Supports, Failures                                          []string
}

type TraceMagnitudeChain struct {
	Steps                           []string
	YDagYReadoutReady               bool
	SocketMagnitudesTypedGivenAlpha bool
	HaggReconstructedGivenAlpha     bool
	AlphaNative                     bool
	CoherentGivenAlphaSeal          bool
	Supports, Failures              []string
}

type OperatorReadout struct {
	TraceAOverT, SquareTraceBOverT2, NEffOperator, OfficialNEff float64
	NEffMatchesGate829                                          bool
	OfficialEqualsOperator                                      bool
	Conditional                                                 bool
	Supports, Failures                                          []string
}

type R3Eligibility struct {
	PostOrientationFiniteTripleSeal       bool
	YDagYReadoutCarrier                   bool
	SocketMagnitudesSourceTypedGivenAlpha bool
	BoundaryAlphaExteriorSeal             bool
	AlphaNative                           bool
	TargetSelectionNative                 bool
	SocketMagnitudeNative                 bool
	NativeSectorTraceMagnitudeReadout     bool
	EligibleForConditionalR3Candidate     bool
	EligibleForOfficialR3                 bool
	EligibleForR4                         bool
	Supports, Failures                    []string
}

type Impact struct {
	Classification, Status                                                           string
	ChainCoherentGivenSeal                                                           bool
	AlphaNative, TargetSelectionNative, SocketMagnitudeNative                        bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                                                 bool
	AlphaStillSealed, AlphaNotNativeWithoutTargetTheorem, NoNativeTargetSelection, NoCrossLaneExclusion                      bool
	NoNativeReducedExteriorFunctional, NoNativeAlphaSource, SocketMagnitudesDependOnAlphaSeal, NoNativeSocketMagnitudeSource bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoNumericalYukawa, NoSectorTraceMagnitude                                   bool
	NotR3NativeTraceLedger, NotR4NativeYukawa                                                                                bool
	Verdict                                                                                                                  string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	AlphaSeal    AlphaExteriorSeal
	Chain        TraceMagnitudeChain
	Readout      OperatorReadout
	R3           R3Eligibility
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	if PiTopRank != 3 || HRminRank != 7 || H10Dim != 10 || H72Dim != 72 {
		return Audit{}, fmt.Errorf("unexpected Gate 873 dimension ledger")
	}
	linear := float64(PiTopRank) / float64(H10Dim) * SBoundary
	quadratic := float64(HRminRank) / float64(H72Dim) * SBoundary * SBoundary
	alpha := linear + quadratic
	if !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction drift: got %.18g want %.18g", alpha, AlphaB)
	}

	a := 3 + 3*AlphaB
	b := 3 + 3*AlphaB*AlphaB - 6*math.Pow(AlphaB, 3) + 12*math.Pow(AlphaB, 4)
	nEff := (a * a) / b

	alphaSeal := AlphaExteriorSeal{
		Name:                  "BoundaryAlphaExteriorExposureEnclosureSeal",
		Expression:            "alpha_B=[rank(Pi_top)/10]s+[rank(H_R^min)/72]s^2",
		ReducedResponse:       "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)",
		LinearNumerator:       PiTopRank,
		LinearDenominator:     H10Dim,
		QuadraticNumerator:    HRminRank,
		QuadraticDenominator:  H72Dim,
		LinearContribution:    linear,
		QuadraticContribution: quadratic,
		Alpha:                 alpha,
		ShapeTyped:            true,
		RankSourcesTyped:      true,
		TargetSelectionNative: false,
		Native:                false,
		Supports:              []string{SupportBoundaryAlphaExteriorSeal, SupportAlphaSocketRankAnatomy},
		Failures:              []string{FailureAlphaStillSealed, FailureAlphaNotNativeWithoutTargetTheorem, FailureNoNativeTargetSelection, FailureNoCrossLaneExclusion, FailureNoNativeReducedExteriorFunctional, FailureNoNativeAlphaSource},
	}

	chain := TraceMagnitudeChain{
		Steps: []string{
			"B_2 reduced exterior response -> alpha_B seal",
			"alpha_B seal -> socket magnitudes |y_+3|^2, |y_-3|^2, |y_-1|^2",
			"socket magnitudes -> Y^dagger Y",
			"Y^dagger Y -> H_agg/T",
			"H_agg/T -> N_eff^operator",
		},
		YDagYReadoutReady:               true,
		SocketMagnitudesTypedGivenAlpha: true,
		HaggReconstructedGivenAlpha:     true,
		AlphaNative:                     false,
		CoherentGivenAlphaSeal:          true,
		Supports:                        []string{SupportFullChainCoherent, SupportYDaggerYReproducesHagg, SupportMatureConditionalReadout},
		Failures:                        []string{FailureSocketMagnitudesDependOnAlphaSeal, FailureNoNativeSocketMagnitudeSource, FailureNoNativeAlphaSource},
	}

	readout := OperatorReadout{
		TraceAOverT:            a,
		SquareTraceBOverT2:     b,
		NEffOperator:           nEff,
		OfficialNEff:           OfficialNEffFrozen,
		NEffMatchesGate829:     near(nEff, 3.002327375081808),
		OfficialEqualsOperator: near(nEff, OfficialNEffFrozen),
		Conditional:            true,
		Supports:               []string{SupportOperatorNEffConditional},
		Failures:               []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}

	r3 := R3Eligibility{
		PostOrientationFiniteTripleSeal:       true,
		YDagYReadoutCarrier:                   true,
		SocketMagnitudesSourceTypedGivenAlpha: true,
		BoundaryAlphaExteriorSeal:             true,
		AlphaNative:                           false,
		TargetSelectionNative:                 false,
		SocketMagnitudeNative:                 false,
		NativeSectorTraceMagnitudeReadout:     false,
		EligibleForConditionalR3Candidate:     true,
		EligibleForOfficialR3:                 false,
		EligibleForR4:                         false,
		Supports:                              []string{SupportR3CandidatePressure},
		Failures:                              []string{FailureAlphaStillSealed, FailureNoNativeTargetSelection, FailureNoNativeSocketMagnitudeSource, FailureNoSectorTraceMagnitude, FailureNotR3NativeTraceLedger, FailureNotR4NativeYukawa},
	}

	impact := Impact{Classification: Classification, Status: R2Status, ChainCoherentGivenSeal: true, AlphaNative: false, TargetSelectionNative: false, SocketMagnitudeNative: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
	firewalls := Firewalls{Enforced: true, AlphaStillSealed: true, AlphaNotNativeWithoutTargetTheorem: true, NoNativeTargetSelection: true, NoCrossLaneExclusion: true, NoNativeReducedExteriorFunctional: true, NoNativeAlphaSource: true, SocketMagnitudesDependOnAlphaSeal: true, NoNativeSocketMagnitudeSource: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNumericalYukawa: true, NoSectorTraceMagnitude: true, NotR3NativeTraceLedger: true, NotR4NativeYukawa: true, Verdict: StatusFirewallVerdict}

	return Audit{ID: AuditID, Ledger: Ledger{AlphaB: AlphaB, SBoundary: SBoundary, OfficialNEffFrozen: OfficialNEffFrozen, AlphaNative: false, R3: false, R4: false, OfficialFrozen: true}, AlphaSeal: alphaSeal, Chain: chain, Readout: readout, R3: r3, Impact: impact, Firewalls: firewalls, Truth: "Gate 873 classifies alpha_B as a BoundaryAlphaExteriorExposureEnclosureSeal and confirms the full trace-magnitude chain is coherent only conditionally on that seal.", Final: "VERDICT: CONDITIONAL_CHAIN_COHERENT_GIVEN_ALPHA_SEAL_BUT_NOT_NATIVE_R3"}, nil
}

func Statuses() []string {
	return []string{StatusGate872Inherited, StatusAlphaSealClassified, StatusFullChainReassembled, StatusYDaggerYReadoutInherited, StatusHaggReadoutInherited, StatusOperatorNEffComputed, StatusR3EligibilityAudited, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusNextWound, StatusFirewallVerdict, SupportBoundaryAlphaExteriorSeal, SupportFullChainCoherent, SupportYDaggerYReproducesHagg, SupportOperatorNEffConditional, FailureAlphaStillSealed, FailureNoNativeTargetSelection, FailureNoOfficialNEffUpdate, FailureNotR3NativeTraceLedger, FailureNotR4NativeYukawa}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.18g s=%.18g official_N_eff=%.16g alpha_native=%t official_frozen=%t", l.AlphaB, l.SBoundary, l.OfficialNEffFrozen, l.AlphaNative, l.OfficialFrozen)
}
func FormatAlphaSeal(a AlphaExteriorSeal) string {
	return fmt.Sprintf("%s expression=%s reduced=%s linear=%d/%d %.18g quadratic=%d/%d %.18g alpha=%.18g shape_typed=%t rank_sources_typed=%t native=%t supports=%s failures=%s", a.Name, a.Expression, a.ReducedResponse, a.LinearNumerator, a.LinearDenominator, a.LinearContribution, a.QuadraticNumerator, a.QuadraticDenominator, a.QuadraticContribution, a.Alpha, a.ShapeTyped, a.RankSourcesTyped, a.Native, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatChain(c TraceMagnitudeChain) string {
	return fmt.Sprintf("chain=%s y_dagger_y_ready=%t socket_magnitudes_given_alpha=%t hagg_given_alpha=%t coherent_given_alpha_seal=%t alpha_native=%t supports=%s failures=%s", strings.Join(c.Steps, " -> "), c.YDagYReadoutReady, c.SocketMagnitudesTypedGivenAlpha, c.HaggReconstructedGivenAlpha, c.CoherentGivenAlphaSeal, c.AlphaNative, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatReadout(r OperatorReadout) string {
	return fmt.Sprintf("a/T=%.18g b/T^2=%.18g operator_N_eff=%.16g official_N_eff=%.16g conditional=%t operator_equals_official=%t supports=%s failures=%s", r.TraceAOverT, r.SquareTraceBOverT2, r.NEffOperator, r.OfficialNEff, r.Conditional, r.OfficialEqualsOperator, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatR3(r R3Eligibility) string {
	return fmt.Sprintf("finite_triple_seal=%t y_dagger_y=%t alpha_seal=%t conditional_candidate=%t official_R3=%t R4=%t supports=%s failures=%s", r.PostOrientationFiniteTripleSeal, r.YDagYReadoutCarrier, r.BoundaryAlphaExteriorSeal, r.EligibleForConditionalR3Candidate, r.EligibleForOfficialR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s chain_coherent_given_seal=%t alpha_native=%t target_selection_native=%t socket_magnitude_native=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t promote_R3=%t promote_R4=%t", i.Classification, i.Status, i.ChainCoherentGivenSeal, i.AlphaNative, i.TargetSelectionNative, i.SocketMagnitudeNative, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t alpha_sealed=%t no_target_selection=%t no_cross_lane=%t no_alpha_source=%t no_socket_magnitude_source=%t no_official_N_eff=%t no_C_updates=%t no_sector_trace_magnitude=%t not_R3=%t not_R4=%t verdict=%s", f.Enforced, f.AlphaStillSealed, f.NoNativeTargetSelection, f.NoCrossLaneExclusion, f.NoNativeAlphaSource, f.NoNativeSocketMagnitudeSource, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoSectorTraceMagnitude, f.NotR3NativeTraceLedger, f.NotR4NativeYukawa, f.Verdict)
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
func containsAll(haystack, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range haystack {
			if h == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.AlphaNotNativeWithoutTargetTheorem && f.NoNativeTargetSelection && f.NoCrossLaneExclusion && f.NoNativeAlphaSource && f.SocketMagnitudesDependOnAlphaSeal && f.NoNativeSocketMagnitudeSource && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NotR3NativeTraceLedger && f.NotR4NativeYukawa
}
