// Package generation2boundaryalphasocketranksourcereentryaudit implements
// Gate 866: BoundaryAlpha SocketRank Source Re-entry Audit.
//
// Gate 866 follows Gate 865's collapse of the Y^dagger Y trace-magnitude
// wound to the source of alpha_B. It re-audits alpha_B through the post-
// orientation finite-triple seal: the linear 3/10 coefficient is source-typed
// by the dominant socket rank over the ambient right lepto-color rectangle plus
// boundary pair, and the quadratic 7/72 coefficient is source-typed by the
// minimal active right edge-domain rank over the augmented H_72 chamber. This
// is a stronger source typing than the earlier abstract dimension resonance,
// but it is still not a native BoundaryAlphaTransportMap theorem.
package generation2boundaryalphasocketranksourcereentryaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE866-BOUNDARY-ALPHA-SOCKET-RANK-SOURCE-REENTRY-AUDIT"

	AlphaB       = 0.0003878958469680527
	SBoundary    = 0.0012924448188162962
	OfficialNEff = 3.0023273474722147

	PiTopRank              = 3
	HRambientRank          = 8
	BoundaryPairRank       = 2
	LinearDenominator      = HRambientRank + BoundaryPairRank
	HRminRank              = 7
	Lambda4V8Rank          = 70
	H72BoundaryRank        = 2
	QuadraticDenominator   = Lambda4V8Rank + H72BoundaryRank
	K7ContactVacuumRank    = 7
	AmbientCarrierFullRank = 16
	MinimalCarrierRank     = 15
	MinimalFullCarrierRank = 30

	Classification = "BOUNDARY_ALPHA_SOCKET_RANK_SOURCE_REENTRY_GIVEN_POST_ORIENTATION_FINITE_TRIPLE_SEAL"
	R2Status       = "R2+++++_BOUNDARY_ALPHA_SOCKET_RANK_REENTRY_SOURCE_TYPING"

	StatusGate865Inherited   = "PASS_GATE865_SOCKET_MAGNITUDE_ALPHA_WOUND_INHERITED"
	StatusFiniteSocketRanks  = "PASS_FINITE_SOCKET_RANKS_3_AND_7_AUDITED"
	StatusLinearThreeTenths  = "PASS_LINEAR_THREE_TENTHS_REINTERPRETED_AS_DOMINANT_SOCKET_OVER_AMBIENT_RIGHT_RECTANGLE_PLUS_BOUNDARY"
	StatusQuadraticSeven72   = "PASS_QUADRATIC_SEVEN_SEVENTY_TWO_REINTERPRETED_AS_ACTIVE_RIGHT_MODULE_OVER_H72"
	StatusAlphaReconstructed = "PASS_ALPHA_B_RECONSTRUCTED_FROM_SOCKET_RANK_SOURCE_CANDIDATES"
	StatusDualSevenFirewall  = "PASS_DUAL_SEVEN_FIREWALL_AUDITED"
	StatusTransportFirewall  = "PASS_BOUNDARY_ALPHA_TRANSPORT_FIREWALL_REENFORCED"
	StatusLedgerFrozen       = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusRemainingWound     = "PASS_REMAINING_WOUND_IDENTIFIED_AS_BOUNDARY_ALPHA_TRANSPORT_MAP"
	StatusFirewallVerdict    = "FIREWALL_PRESERVED_GATE866_SOCKET_RANK_RATIOS_NOT_ACTIVATION_THEOREM"

	SupportAlphaCoefficientsHaveSocketRankSource = "CONDITIONAL_SUPPORT_ALPHA_B_COEFFICIENTS_HAVE_POST_ORIENTATION_FINITE_TRIPLE_SOCKET_RANK_SOURCE"
	SupportThreeSourceIsPiTopRank                = "CONDITIONAL_SUPPORT_3_SOURCE_IS_PI_TOP_RANK"
	SupportSevenSourceIsHRMinRank                = "CONDITIONAL_SUPPORT_7_SOURCE_IS_H_R_MIN_ACTIVE_EDGE_DOMAIN_RANK"
	SupportLinearLaneSocketRank                  = "CONDITIONAL_SUPPORT_LINEAR_LANE_IS_DOMINANT_COLOR_SOCKET_OVER_AMBIENT_RIGHT_RECTANGLE_PLUS_BOUNDARY_PAIR"
	SupportQuadraticLaneSocketRank               = "CONDITIONAL_SUPPORT_QUADRATIC_LANE_IS_ACTIVE_PUNCTURED_RIGHT_EDGE_DOMAIN_OVER_H72"
	SupportSocketMagnitudePressureToAlpha        = "CONDITIONAL_SUPPORT_SOCKET_MAGNITUDE_SOURCE_PRESSURE_REDUCES_TO_BOUNDARY_ALPHA_TRANSPORT"
	SupportFiniteTripleBranchReentersAlpha       = "CONDITIONAL_SUPPORT_POST_ORIENTATION_FINITE_TRIPLE_BRANCH_REENTERS_ALPHA_SOURCE_TYPING"

	FailureSocketRankRatiosNotNativeActivation = "FAILED_ROUTE_SOCKET_RANK_RATIOS_NOT_NATIVE_ACTIVATION_MAP"
	FailureNoHRMinToK7Map                      = "FAILED_ROUTE_NO_TYPED_H_R_MIN_TO_K7_MAP"
	FailureNoBoundaryAlphaTransport            = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_THEOREM"
	FailureNoBoundaryAlphaActivation           = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_ACTIVATION_THEOREM"
	FailureAlphaStillSealed                    = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureDimensionRatioNotActivation         = "FAILED_ROUTE_DIMENSION_RATIO_NOT_ACTIVATION_MAP"
	FailureDualSevenNotIdentified              = "FAILED_ROUTE_H_R_MIN_SEVEN_NOT_IDENTIFIED_WITH_K7_CONTACT_VACUUM"
	FailureNoNativeAlphaSource                 = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureNoOfficialNEffUpdate                = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNumericalYukawa                   = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoSectorTraceMagnitude              = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoFullUnbrokenAF                    = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF                   = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple                = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNotR3                               = "FAILED_ROUTE_NOT_R3_SOCKET_RANK_SOURCE_TYPING_ONLY"
	FailureNotR4                               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB, SBoundary, OfficialNEff float64
	OfficialFrozen                  bool
	AlphaNative, BoundaryTransport  bool
	R3, R4                          bool
}

type RankSource struct {
	Name, NumeratorSource, DenominatorSource, Domain, Interpretation string
	Numerator, Denominator, Power                                    int
	Coefficient                                                      float64
	SocketRankSourced, NativeTransport                               bool
	Supports, Failures                                               []string
}

type AlphaFormula struct {
	Expression, SocketRankExpression                               string
	S, LinearCoefficient, QuadraticCoefficient, ReconstructedAlpha float64
	LinearLane, QuadraticLane                                      RankSource
	ReconstructsAlpha, Native, TransportMapCertified               bool
	Supports, Failures                                             []string
}

type DualSevenFirewall struct {
	HRMinRank, K7Rank                          int
	SameInteger, Identified, TypedMapCertified bool
	HRMinMeaning, K7Meaning, Verdict           string
	Supports, Failures                         []string
}

type TransportObstruction struct {
	DimensionRatiosCoherent, SocketRankRatiosCoherent bool
	ActivationMapCertified, AlphaNative               bool
	RemainingWound, NextGate                          string
	Supports, Failures                                []string
}

type R3Assessment struct {
	SocketMagnitudePressureReducedToAlpha bool
	AlphaNative, BoundaryTransportNative  bool
	SectorTraceMagnitudeReadout           bool
	EligibleForR3, EligibleForR4          bool
	Supports, Failures                    []string
}

type Impact struct {
	Classification, Status                                                           string
	AlphaCoefficientSourceTypingUpgraded, NativeAlphaSourceSolved                    bool
	BoundaryTransportSolved, SocketMagnitudeSourceNative, SectorTraceReadout         bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                              bool
	SocketRankRatiosNotActivation, NoHRMinToK7Map, NoBoundaryAlphaTransport, AlphaStillSealed             bool
	DimensionRatioNotActivation, DualSevenNotIdentified, NoNativeAlphaSource, NoOfficialNEffUpdate        bool
	NoCYukawaCHiggsUpdate, NoNumericalYukawa, NoSectorTraceMagnitude, NoFullUnbrokenAF, AForientNotFullAF bool
	NoNativeFiniteTriple, NotR3, NotR4                                                                    bool
	Verdict                                                                                               string
}

type Audit struct {
	ID          string
	Ledger      Ledger
	Alpha       AlphaFormula
	DualSeven   DualSevenFirewall
	Obstruction TransportObstruction
	R3          R3Assessment
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func LinearCoefficient() float64           { return float64(PiTopRank) / float64(LinearDenominator) }
func QuadraticCoefficient() float64        { return float64(HRminRank) / float64(QuadraticDenominator) }
func ReconstructedAlpha(s float64) float64 { return LinearCoefficient()*s + QuadraticCoefficient()*s*s }

func BuildDefault() (Audit, error) {
	linear := RankSource{
		Name: "linear dominant socket lane", NumeratorSource: "rank(Pi_top=e_+ tensor P_3)", DenominatorSource: "dim(H_R^ambient)+dim(B_2)", Domain: "H_R^ambient plus boundary pair", Interpretation: "dominant color socket response over ambient right lepto-color rectangle plus boundary pair", Numerator: PiTopRank, Denominator: LinearDenominator, Power: 1, Coefficient: LinearCoefficient(), SocketRankSourced: true, NativeTransport: false,
		Supports: []string{SupportThreeSourceIsPiTopRank, SupportLinearLaneSocketRank}, Failures: []string{FailureSocketRankRatiosNotNativeActivation, FailureNoBoundaryAlphaTransport},
	}
	quadratic := RankSource{
		Name: "quadratic active right edge-domain lane", NumeratorSource: "rank(H_R^min)", DenominatorSource: "dim(Lambda^4 V_8 plus B_2)=dim(H_72)", Domain: "H_72 augmented chamber", Interpretation: "active punctured right edge-domain response over augmented 72-chamber", Numerator: HRminRank, Denominator: QuadraticDenominator, Power: 2, Coefficient: QuadraticCoefficient(), SocketRankSourced: true, NativeTransport: false,
		Supports: []string{SupportSevenSourceIsHRMinRank, SupportQuadraticLaneSocketRank}, Failures: []string{FailureSocketRankRatiosNotNativeActivation, FailureNoBoundaryAlphaTransport, FailureNoHRMinToK7Map},
	}
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, SBoundary: SBoundary, OfficialNEff: OfficialNEff, OfficialFrozen: true, AlphaNative: false, BoundaryTransport: false, R3: false, R4: false},
		Alpha: AlphaFormula{
			Expression: "alpha_B = (3/10)s + (7/72)s^2", SocketRankExpression: "alpha_B = rank(Pi_top)/(dim(H_R^ambient)+2) s + rank(H_R^min)/dim(H_72) s^2",
			S: SBoundary, LinearCoefficient: LinearCoefficient(), QuadraticCoefficient: QuadraticCoefficient(), ReconstructedAlpha: ReconstructedAlpha(SBoundary), LinearLane: linear, QuadraticLane: quadratic, ReconstructsAlpha: true, Native: false, TransportMapCertified: false,
			Supports: []string{StatusGate865Inherited, StatusFiniteSocketRanks, StatusLinearThreeTenths, StatusQuadraticSeven72, StatusAlphaReconstructed, SupportAlphaCoefficientsHaveSocketRankSource, SupportFiniteTripleBranchReentersAlpha},
			Failures: []string{FailureSocketRankRatiosNotNativeActivation, FailureNoBoundaryAlphaTransport, FailureNoBoundaryAlphaActivation, FailureAlphaStillSealed, FailureDimensionRatioNotActivation},
		},
		DualSeven:   DualSevenFirewall{HRMinRank: HRminRank, K7Rank: K7ContactVacuumRank, SameInteger: true, Identified: false, TypedMapCertified: false, HRMinMeaning: "active punctured right edge-domain rank from post-orientation finite-triple seal", K7Meaning: "Boolean-octonionic contact-vacuum/contact carrier dimension", Verdict: StatusDualSevenFirewall, Supports: []string{StatusDualSevenFirewall, SupportSevenSourceIsHRMinRank}, Failures: []string{FailureNoHRMinToK7Map, FailureDualSevenNotIdentified}},
		Obstruction: TransportObstruction{DimensionRatiosCoherent: true, SocketRankRatiosCoherent: true, ActivationMapCertified: false, AlphaNative: false, RemainingWound: "BoundaryAlphaTransportMap / S_split -> alpha_B", NextGate: "Gate 867 — BoundaryAlpha Transport Functional Source Audit", Supports: []string{StatusRemainingWound, SupportSocketMagnitudePressureToAlpha, SupportAlphaCoefficientsHaveSocketRankSource}, Failures: []string{FailureSocketRankRatiosNotNativeActivation, FailureNoBoundaryAlphaTransport, FailureNoBoundaryAlphaActivation, FailureAlphaStillSealed}},
		R3:          R3Assessment{SocketMagnitudePressureReducedToAlpha: true, AlphaNative: false, BoundaryTransportNative: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false, Supports: []string{SupportSocketMagnitudePressureToAlpha, SupportFiniteTripleBranchReentersAlpha}, Failures: []string{FailureAlphaStillSealed, FailureNoSectorTraceMagnitude, FailureNotR3, FailureNotR4}},
		Impact:      Impact{Classification: Classification, Status: R2Status, AlphaCoefficientSourceTypingUpgraded: true, NativeAlphaSourceSolved: false, BoundaryTransportSolved: false, SocketMagnitudeSourceNative: false, SectorTraceReadout: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false},
		Firewalls:   Firewalls{Enforced: true, SocketRankRatiosNotActivation: true, NoHRMinToK7Map: true, NoBoundaryAlphaTransport: true, AlphaStillSealed: true, DimensionRatioNotActivation: true, DualSevenNotIdentified: true, NoNativeAlphaSource: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNumericalYukawa: true, NoSectorTraceMagnitude: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NotR3: true, NotR4: true, Verdict: StatusFirewallVerdict},
		Truth:       "Gate 866 re-enters the alpha_B source problem through finite-triple socket ranks: 3 becomes rank(Pi_top) and 7 becomes rank(H_R^min). This strengthens source typing, but socket rank ratios still do not constitute a native activation map.",
		Final:       "The alpha coefficients now have post-orientation finite-triple socket-rank source candidates; the remaining lawful wound is still the BoundaryAlphaTransportMap S_split -> alpha_B.",
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if a.ID != AuditID || a.Ledger.AlphaB != AlphaB || !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.BoundaryTransport || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger malformed or overpromoted")
	}
	if a.Alpha.LinearLane.Numerator != PiTopRank || a.Alpha.LinearLane.Denominator != LinearDenominator || !near(a.Alpha.LinearLane.Coefficient, 0.3) || !a.Alpha.LinearLane.SocketRankSourced || a.Alpha.LinearLane.NativeTransport {
		return err("linear socket-rank lane malformed")
	}
	if a.Alpha.QuadraticLane.Numerator != HRminRank || a.Alpha.QuadraticLane.Denominator != QuadraticDenominator || !near(a.Alpha.QuadraticLane.Coefficient, 7.0/72.0) || !a.Alpha.QuadraticLane.SocketRankSourced || a.Alpha.QuadraticLane.NativeTransport {
		return err("quadratic socket-rank lane malformed")
	}
	if !a.Alpha.ReconstructsAlpha || a.Alpha.Native || a.Alpha.TransportMapCertified || !near(a.Alpha.ReconstructedAlpha, AlphaB) || !containsAll(a.Alpha.Supports, []string{StatusGate865Inherited, SupportAlphaCoefficientsHaveSocketRankSource}) || !containsAll(a.Alpha.Failures, []string{FailureSocketRankRatiosNotNativeActivation, FailureNoBoundaryAlphaTransport, FailureAlphaStillSealed}) {
		return err("alpha reconstruction malformed or overpromoted")
	}
	if !a.DualSeven.SameInteger || a.DualSeven.Identified || a.DualSeven.TypedMapCertified || !containsAll(a.DualSeven.Failures, []string{FailureNoHRMinToK7Map, FailureDualSevenNotIdentified}) {
		return err("dual-seven firewall malformed")
	}
	if !a.Obstruction.DimensionRatiosCoherent || !a.Obstruction.SocketRankRatiosCoherent || a.Obstruction.ActivationMapCertified || a.Obstruction.AlphaNative || !containsAll(a.Obstruction.Failures, []string{FailureSocketRankRatiosNotNativeActivation, FailureNoBoundaryAlphaTransport, FailureNoBoundaryAlphaActivation}) {
		return err("transport obstruction malformed")
	}
	if !a.R3.SocketMagnitudePressureReducedToAlpha || a.R3.AlphaNative || a.R3.BoundaryTransportNative || a.R3.SectorTraceMagnitudeReadout || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		return err("R3 assessment overpromoted")
	}
	if a.Impact.NativeAlphaSourceSolved || a.Impact.BoundaryTransportSolved || a.Impact.SocketMagnitudeSourceNative || a.Impact.SectorTraceReadout || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact overpromoted")
	}
	if !firewallsOK(a.Firewalls) {
		return err("firewalls not enforced")
	}
	return nil
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func containsAll(hay []string, needles []string) bool {
	m := make(map[string]bool, len(hay))
	for _, h := range hay {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.SocketRankRatiosNotActivation && f.NoHRMinToK7Map && f.NoBoundaryAlphaTransport && f.AlphaStillSealed && f.DimensionRatioNotActivation && f.DualSevenNotIdentified && f.NoNativeAlphaSource && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NotR3 && f.NotR4 && f.Verdict == StatusFirewallVerdict
}

func Statuses() []string {
	return []string{
		StatusGate865Inherited, StatusFiniteSocketRanks, StatusLinearThreeTenths, StatusQuadraticSeven72, StatusAlphaReconstructed, StatusDualSevenFirewall, StatusTransportFirewall, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusRemainingWound, StatusFirewallVerdict,
		SupportAlphaCoefficientsHaveSocketRankSource, SupportThreeSourceIsPiTopRank, SupportSevenSourceIsHRMinRank, SupportLinearLaneSocketRank, SupportQuadraticLaneSocketRank, SupportSocketMagnitudePressureToAlpha, SupportFiniteTripleBranchReentersAlpha,
		FailureSocketRankRatiosNotNativeActivation, FailureNoHRMinToK7Map, FailureNoBoundaryAlphaTransport, FailureNoBoundaryAlphaActivation, FailureAlphaStillSealed, FailureDimensionRatioNotActivation, FailureDualSevenNotIdentified, FailureNoNativeAlphaSource, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNumericalYukawa, FailureNoSectorTraceMagnitude, FailureNoFullUnbrokenAF, FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNotR3, FailureNotR4,
	}
}

func FormatRankSource(r RankSource) string {
	return fmt.Sprintf("%s: %s/%s = %d/%d = %.15g, power=s^%d, socketRankSourced=%t, nativeTransport=%t", r.Name, r.NumeratorSource, r.DenominatorSource, r.Numerator, r.Denominator, r.Coefficient, r.Power, r.SocketRankSourced, r.NativeTransport)
}

func FormatAlpha(a AlphaFormula) string {
	return fmt.Sprintf("%s | %s | s=%.16g, alpha=%.16g, reconstructed=%.16g, linear=(%s), quadratic=(%s), native=%t, transportMapCertified=%t", a.Expression, a.SocketRankExpression, a.S, AlphaB, a.ReconstructedAlpha, FormatRankSource(a.LinearLane), FormatRankSource(a.QuadraticLane), a.Native, a.TransportMapCertified)
}

func FormatDualSeven(d DualSevenFirewall) string {
	return fmt.Sprintf("HRminRank=%d (%s), K7Rank=%d (%s), sameInteger=%t, identified=%t, typedMapCertified=%t", d.HRMinRank, d.HRMinMeaning, d.K7Rank, d.K7Meaning, d.SameInteger, d.Identified, d.TypedMapCertified)
}

func FormatObstruction(o TransportObstruction) string {
	return fmt.Sprintf("dimensionRatiosCoherent=%t, socketRankRatiosCoherent=%t, activationMapCertified=%t, alphaNative=%t, remaining=%s, next=%s", o.DimensionRatiosCoherent, o.SocketRankRatiosCoherent, o.ActivationMapCertified, o.AlphaNative, o.RemainingWound, o.NextGate)
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("socketMagnitudePressureReducedToAlpha=%t, alphaNative=%t, boundaryTransportNative=%t, sectorTraceReadout=%t, eligibleR3=%t, eligibleR4=%t", r.SocketMagnitudePressureReducedToAlpha, r.AlphaNative, r.BoundaryTransportNative, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s, status=%s, sourceTypingUpgraded=%t, nativeAlphaSolved=%t, boundaryTransportSolved=%t, canUpdateNEff=%t, canPromoteR3=%t", i.Classification, i.Status, i.AlphaCoefficientSourceTypingUpgraded, i.NativeAlphaSourceSolved, i.BoundaryTransportSolved, i.CanUpdateNEff, i.CanPromoteToR3)
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g, s=%.16g, officialNEff=%.16g, frozen=%t, alphaNative=%t, boundaryTransport=%t, R3=%t, R4=%t", l.AlphaB, l.SBoundary, l.OfficialNEff, l.OfficialFrozen, l.AlphaNative, l.BoundaryTransport, l.R3, l.R4)
}

func FormatFirewalls(f Firewalls) string {
	parts := []string{}
	if f.SocketRankRatiosNotActivation {
		parts = append(parts, FailureSocketRankRatiosNotNativeActivation)
	}
	if f.NoHRMinToK7Map {
		parts = append(parts, FailureNoHRMinToK7Map)
	}
	if f.NoBoundaryAlphaTransport {
		parts = append(parts, FailureNoBoundaryAlphaTransport)
	}
	if f.AlphaStillSealed {
		parts = append(parts, FailureAlphaStillSealed)
	}
	if f.NotR3 {
		parts = append(parts, FailureNotR3)
	}
	if f.NotR4 {
		parts = append(parts, FailureNotR4)
	}
	return fmt.Sprintf("enforced=%t, verdict=%s, failures=[%s]", f.Enforced, f.Verdict, strings.Join(parts, "; "))
}
