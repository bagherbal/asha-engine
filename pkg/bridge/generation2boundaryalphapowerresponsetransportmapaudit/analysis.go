// Package generation2boundaryalphapowerresponsetransportmapaudit implements
// Gate 867: BoundaryAlpha Power-Response TransportMap Audit.
//
// Gate 867 follows Gate 866's socket-rank source re-entry for alpha_B. The
// numerators 3 and 7 are no longer floating dimension coincidences: 3 is the
// rank of the dominant color socket Pi_top=e_+ tensor P_3 and 7 is the rank of
// the active punctured right edge-domain H_R^min. Gate 867 audits the sharper
// remaining wound: why the same boundary split coordinate s enters the dominant
// socket lane linearly and the active right-domain lane quadratically. The
// verdict is intentionally conservative: the socket-rank power-response shape
// is coherent, but no BoundaryAlphaPowerResponseTransportMap is certified.
package generation2boundaryalphapowerresponsetransportmapaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE867-BOUNDARY-ALPHA-POWER-RESPONSE-TRANSPORTMAP-AUDIT"

	AlphaB       = 0.0003878958469680527
	SBoundary    = 0.0012924448188162962
	OfficialNEff = 3.0023273474722147

	PiTopRank            = 3
	HRambientRank        = 8
	BoundaryPairRank     = 2
	LinearDenominator    = HRambientRank + BoundaryPairRank
	HRminRank            = 7
	Lambda4V8Rank        = 70
	H72BoundaryRank      = 2
	QuadraticDenominator = Lambda4V8Rank + H72BoundaryRank

	Classification = "BOUNDARY_ALPHA_POWER_RESPONSE_TRANSPORTMAP_OBSTRUCTION_GIVEN_SOCKET_RANK_SOURCE_TYPING"
	R2Status       = "R2+++++_BOUNDARY_ALPHA_POWER_RESPONSE_TRANSPORT_OBSTRUCTION"

	StatusGate866Inherited        = "PASS_GATE866_SOCKET_RANK_ALPHA_SOURCE_TYPING_INHERITED"
	StatusLinearLaneAudited       = "PASS_LINEAR_DOMINANT_SOCKET_LANE_AUDITED"
	StatusQuadraticLaneAudited    = "PASS_QUADRATIC_ACTIVE_RIGHT_DOMAIN_LANE_AUDITED"
	StatusSharedSTransportDefined = "PASS_SHARED_S_SPLIT_TRANSPORT_REQUIREMENT_DEFINED"
	StatusPowerOrderAudited       = "PASS_LINEAR_VS_QUADRATIC_RESPONSE_ORDER_AUDITED"
	StatusDenominatorsReaudited   = "PASS_BOUNDARY_AUGMENTED_DENOMINATORS_REAUDITED"
	StatusAlphaReconstructed      = "PASS_ALPHA_B_POWER_RESPONSE_SHAPE_RECONSTRUCTED"
	StatusTransportFirewall       = "PASS_BOUNDARY_ALPHA_POWER_RESPONSE_TRANSPORT_FIREWALL_ENFORCED"
	StatusLedgerFrozen            = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed      = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusNextWoundIdentified     = "PASS_NEXT_WOUND_IDENTIFIED_AS_BOUNDARY_ALPHA_POWER_RESPONSE_TRANSPORT_LAW"
	StatusFirewallVerdict         = "FIREWALL_PRESERVED_GATE867_POWER_RESPONSE_SHAPE_NOT_TRANSPORT_THEOREM"

	SupportAlphaPowerResponseShape       = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_SOCKET_RANK_POWER_RESPONSE_SHAPE"
	SupportLinearFirstOrderCandidate     = "CONDITIONAL_SUPPORT_LINEAR_LANE_IS_DOMINANT_SOCKET_FIRST_ORDER_RESPONSE_CANDIDATE"
	SupportQuadraticSecondOrderCandidate = "CONDITIONAL_SUPPORT_QUADRATIC_LANE_IS_ACTIVE_RIGHT_DOMAIN_SECOND_ORDER_RESPONSE_CANDIDATE"
	SupportDenominatorsTypedDomains      = "CONDITIONAL_SUPPORT_DENOMINATORS_ARE_TYPED_BOUNDARY_AUGMENTED_DOMAINS"
	SupportWoundReducesToTransportLaw    = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_WOUND_NOW_REDUCES_TO_POWER_RESPONSE_TRANSPORT_LAW"
	SupportSocketRankShapeInherited      = "CONDITIONAL_SUPPORT_SOCKET_RANK_SOURCE_TYPING_INHERITED_FROM_GATE866"
	SupportSharedCoordinatePressure      = "CONDITIONAL_SUPPORT_SAME_S_SPLIT_COORDINATE_FEEDS_TWO_TYPED_RESPONSE_LANE_CANDIDATES"
	SupportLinearCoefficientFromPiTop    = "CONDITIONAL_SUPPORT_3_OVER_10_FROM_PI_TOP_OVER_AMBIENT_RIGHT_RECTANGLE_PLUS_BOUNDARY"
	SupportQuadraticCoefficientFromHRMin = "CONDITIONAL_SUPPORT_7_OVER_72_FROM_H_R_MIN_OVER_AUGMENTED_72_CHAMBER"

	FailureNoPowerResponseTransport       = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_POWER_RESPONSE_TRANSPORTMAP_CERTIFIED"
	FailureNoLinearTransport              = "FAILED_ROUTE_NO_TYPED_S_SPLIT_TO_DOMINANT_SOCKET_LINEAR_TRANSPORT"
	FailureNoQuadraticTransport           = "FAILED_ROUTE_NO_TYPED_S_SPLIT_SQUARED_TO_ACTIVE_RIGHT_DOMAIN_TRANSPORT"
	FailureLinearVsQuadraticNotDerived    = "FAILED_ROUTE_LINEAR_VS_QUADRATIC_RESPONSE_ORDER_NOT_DERIVED"
	FailureSameSNotTransportedToBoth      = "FAILED_ROUTE_SAME_S_SPLIT_COORDINATE_NOT_LAWFULLY_TRANSPORTED_INTO_BOTH_SOCKET_RANK_DOMAINS"
	FailureSocketRankRatiosNotActivation  = "FAILED_ROUTE_SOCKET_RANK_RATIOS_NOT_ACTIVATION_THEOREM"
	FailureDenominatorTypingNotActivation = "FAILED_ROUTE_DENOMINATOR_TYPING_NOT_ACTIVATION_THEOREM"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeAlphaSource            = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureNoBoundaryAlphaTransport       = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_THEOREM"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNumericalYukawa              = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoSectorTraceMagnitude         = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoFullUnbrokenAF               = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF              = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple           = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNotR3                          = "FAILED_ROUTE_NOT_R3_POWER_RESPONSE_TRANSPORT_OBSTRUCTION"
	FailureNotR4                          = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB, SBoundary, OfficialNEff float64
	OfficialFrozen                  bool
	AlphaNative, PowerTransport     bool
	R3, R4                          bool
}

type PowerLane struct {
	Name, NumeratorSource, DenominatorSource, Domain, Interpretation string
	Numerator, Denominator, Power                                    int
	Coefficient, Contribution                                        float64
	SocketRankSourced, DenominatorTyped                              bool
	TransportCertified, ResponseOrderDerived                         bool
	Supports, Failures                                               []string
}

type AlphaPowerResponse struct {
	Expression, SocketRankExpression string
	S                                float64
	Linear, Quadratic                PowerLane
	ReconstructedAlpha               float64
	ShapeCoherent, Native            bool
	TransportMapCertified            bool
	Supports, Failures               []string
}

type SharedCoordinateAudit struct {
	Coordinate, LinearCodomain, QuadraticCodomain string
	SameCoordinateUsed                            bool
	BothCodomainsTyped                            bool
	TransportIntoBothCertified                    bool
	PowerOrderDerived                             bool
	Supports, Failures                            []string
}

type DenominatorAudit struct {
	LinearDenominator, QuadraticDenominator int
	LinearMeaning, QuadraticMeaning         string
	TypedBoundaryAugmentedDomains           bool
	ActivationTheoremCertified              bool
	Supports, Failures                      []string
}

type TransportObstruction struct {
	SocketRankPowerShapeCoherent bool
	LinearTransportCertified     bool
	QuadraticTransportCertified  bool
	PowerOrderDerived            bool
	TransportMapCertified        bool
	AlphaNative                  bool
	RemainingWound, NextGate     string
	Supports, Failures           []string
}

type R3Assessment struct {
	YDagYReadoutCarrierReady, SocketMagnitudeTransferTyped bool
	AlphaNative, BoundaryPowerTransportNative              bool
	SectorTraceMagnitudeReadout                            bool
	EligibleForR3, EligibleForR4                           bool
	Supports, Failures                                     []string
}

type Impact struct {
	Classification, Status                                                            boolString
	AlphaSourceTypingSharpened, NativeAlphaSourceSolved, BoundaryPowerTransportSolved bool
	SocketMagnitudeSourceNative, SectorTraceReadout                                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4  bool
}

type boolString string

type Firewalls struct {
	Enforced                                                                                                   bool
	NoPowerResponseTransport, NoLinearTransport, NoQuadraticTransport, LinearVsQuadraticNotDerived             bool
	SameSNotTransportedToBoth, SocketRankRatiosNotActivation, DenominatorTypingNotActivation, AlphaStillSealed bool
	NoNativeAlphaSource, NoBoundaryAlphaTransport, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate                 bool
	NoNumericalYukawa, NoSectorTraceMagnitude, NoFullUnbrokenAF, AForientNotFullAF                             bool
	NoNativeFiniteTriple, NotR3, NotR4                                                                         bool
	Verdict                                                                                                    string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	Alpha        AlphaPowerResponse
	SharedS      SharedCoordinateAudit
	Denominators DenominatorAudit
	Obstruction  TransportObstruction
	R3           R3Assessment
	Impact       Impact
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func LinearCoefficient() float64           { return float64(PiTopRank) / float64(LinearDenominator) }
func QuadraticCoefficient() float64        { return float64(HRminRank) / float64(QuadraticDenominator) }
func LinearContribution(s float64) float64 { return LinearCoefficient() * s }
func QuadraticContribution(s float64) float64 {
	return QuadraticCoefficient() * s * s
}
func ReconstructedAlpha(s float64) float64 { return LinearContribution(s) + QuadraticContribution(s) }

func BuildDefault() (Audit, error) {
	linear := PowerLane{
		Name: "linear dominant socket response lane", NumeratorSource: "rank(Pi_top=e_+ tensor P_3)", DenominatorSource: "dim(H_R^ambient)+dim(B_2)", Domain: "H_R^ambient plus boundary pair", Interpretation: "first-order boundary response candidate of the dominant color socket over the ambient right lepto-color rectangle plus boundary pair", Numerator: PiTopRank, Denominator: LinearDenominator, Power: 1, Coefficient: LinearCoefficient(), Contribution: LinearContribution(SBoundary), SocketRankSourced: true, DenominatorTyped: true, TransportCertified: false, ResponseOrderDerived: false,
		Supports: []string{StatusLinearLaneAudited, SupportLinearFirstOrderCandidate, SupportLinearCoefficientFromPiTop},
		Failures: []string{FailureNoLinearTransport, FailureLinearVsQuadraticNotDerived, FailureSocketRankRatiosNotActivation},
	}
	quadratic := PowerLane{
		Name: "quadratic active right-domain response lane", NumeratorSource: "rank(H_R^min)", DenominatorSource: "dim(Lambda^4 V_8)+dim(B_2)", Domain: "H_72 augmented chamber", Interpretation: "second-order boundary response candidate of the active punctured right edge-domain over the augmented 72-chamber", Numerator: HRminRank, Denominator: QuadraticDenominator, Power: 2, Coefficient: QuadraticCoefficient(), Contribution: QuadraticContribution(SBoundary), SocketRankSourced: true, DenominatorTyped: true, TransportCertified: false, ResponseOrderDerived: false,
		Supports: []string{StatusQuadraticLaneAudited, SupportQuadraticSecondOrderCandidate, SupportQuadraticCoefficientFromHRMin},
		Failures: []string{FailureNoQuadraticTransport, FailureLinearVsQuadraticNotDerived, FailureSocketRankRatiosNotActivation},
	}
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, SBoundary: SBoundary, OfficialNEff: OfficialNEff, OfficialFrozen: true, AlphaNative: false, PowerTransport: false, R3: false, R4: false},
		Alpha: AlphaPowerResponse{
			Expression: "alpha_B = (3/10)s + (7/72)s^2", SocketRankExpression: "alpha_B = rank(Pi_top)/(8+2) s + rank(H_R^min)/(70+2) s^2", S: SBoundary, Linear: linear, Quadratic: quadratic, ReconstructedAlpha: ReconstructedAlpha(SBoundary), ShapeCoherent: true, Native: false, TransportMapCertified: false,
			Supports: []string{StatusGate866Inherited, StatusAlphaReconstructed, SupportAlphaPowerResponseShape, SupportSocketRankShapeInherited, SupportWoundReducesToTransportLaw},
			Failures: []string{FailureNoPowerResponseTransport, FailureNoLinearTransport, FailureNoQuadraticTransport, FailureLinearVsQuadraticNotDerived, FailureAlphaStillSealed},
		},
		SharedS:      SharedCoordinateAudit{Coordinate: "S_split=s", LinearCodomain: "H_R^ambient plus B_2 / dominant socket lane", QuadraticCodomain: "H_72 / active punctured right-domain lane", SameCoordinateUsed: true, BothCodomainsTyped: true, TransportIntoBothCertified: false, PowerOrderDerived: false, Supports: []string{StatusSharedSTransportDefined, SupportSharedCoordinatePressure}, Failures: []string{FailureSameSNotTransportedToBoth, FailureNoPowerResponseTransport, FailureLinearVsQuadraticNotDerived}},
		Denominators: DenominatorAudit{LinearDenominator: LinearDenominator, QuadraticDenominator: QuadraticDenominator, LinearMeaning: "8+2 = ambient right lepto-color rectangle plus boundary pair", QuadraticMeaning: "70+2 = Lambda^4 V_8 chamber plus boundary pair", TypedBoundaryAugmentedDomains: true, ActivationTheoremCertified: false, Supports: []string{StatusDenominatorsReaudited, SupportDenominatorsTypedDomains}, Failures: []string{FailureDenominatorTypingNotActivation, FailureNoPowerResponseTransport}},
		Obstruction:  TransportObstruction{SocketRankPowerShapeCoherent: true, LinearTransportCertified: false, QuadraticTransportCertified: false, PowerOrderDerived: false, TransportMapCertified: false, AlphaNative: false, RemainingWound: "BoundaryAlphaPowerResponseTransportMap: why s appears linearly in the dominant socket lane and s^2 in the active right-domain lane", NextGate: "Gate 868 — BoundaryAlpha Response Functional Source Audit", Supports: []string{StatusNextWoundIdentified, SupportWoundReducesToTransportLaw, SupportAlphaPowerResponseShape}, Failures: []string{FailureNoPowerResponseTransport, FailureNoLinearTransport, FailureNoQuadraticTransport, FailureLinearVsQuadraticNotDerived, FailureAlphaStillSealed}},
		R3:           R3Assessment{YDagYReadoutCarrierReady: true, SocketMagnitudeTransferTyped: true, AlphaNative: false, BoundaryPowerTransportNative: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false, Supports: []string{SupportWoundReducesToTransportLaw, SupportSocketRankShapeInherited}, Failures: []string{FailureAlphaStillSealed, FailureNoSectorTraceMagnitude, FailureNotR3, FailureNotR4}},
		Impact:       Impact{Classification: boolString(Classification), Status: boolString(R2Status), AlphaSourceTypingSharpened: true, NativeAlphaSourceSolved: false, BoundaryPowerTransportSolved: false, SocketMagnitudeSourceNative: false, SectorTraceReadout: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false},
		Firewalls:    Firewalls{Enforced: true, NoPowerResponseTransport: true, NoLinearTransport: true, NoQuadraticTransport: true, LinearVsQuadraticNotDerived: true, SameSNotTransportedToBoth: true, SocketRankRatiosNotActivation: true, DenominatorTypingNotActivation: true, AlphaStillSealed: true, NoNativeAlphaSource: true, NoBoundaryAlphaTransport: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNumericalYukawa: true, NoSectorTraceMagnitude: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NotR3: true, NotR4: true, Verdict: StatusFirewallVerdict},
		Truth:        "Gate 867 keeps Gate 866's finite-body numerators but audits the actual missing law: why S_split feeds a first-order dominant socket response and a second-order active right-domain response. The socket-rank power-response shape is coherent, but no transport principle derives the s and s^2 powers.",
		Final:        "The alpha wound has moved from coefficient source typing to a BoundaryAlphaPowerResponseTransportMap: rank ratios over typed domains do not yet form an activation law.",
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if a.ID != AuditID || a.Ledger.AlphaB != AlphaB || !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.PowerTransport || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger malformed or overpromoted")
	}
	if a.Alpha.Linear.Numerator != PiTopRank || a.Alpha.Linear.Denominator != LinearDenominator || a.Alpha.Linear.Power != 1 || !near(a.Alpha.Linear.Coefficient, 0.3) || !a.Alpha.Linear.SocketRankSourced || !a.Alpha.Linear.DenominatorTyped || a.Alpha.Linear.TransportCertified || a.Alpha.Linear.ResponseOrderDerived {
		return err("linear power-response lane malformed")
	}
	if a.Alpha.Quadratic.Numerator != HRminRank || a.Alpha.Quadratic.Denominator != QuadraticDenominator || a.Alpha.Quadratic.Power != 2 || !near(a.Alpha.Quadratic.Coefficient, 7.0/72.0) || !a.Alpha.Quadratic.SocketRankSourced || !a.Alpha.Quadratic.DenominatorTyped || a.Alpha.Quadratic.TransportCertified || a.Alpha.Quadratic.ResponseOrderDerived {
		return err("quadratic power-response lane malformed")
	}
	if !a.Alpha.ShapeCoherent || a.Alpha.Native || a.Alpha.TransportMapCertified || !near(a.Alpha.ReconstructedAlpha, AlphaB) || !containsAll(a.Alpha.Supports, []string{StatusGate866Inherited, SupportAlphaPowerResponseShape}) || !containsAll(a.Alpha.Failures, []string{FailureNoPowerResponseTransport, FailureLinearVsQuadraticNotDerived, FailureAlphaStillSealed}) {
		return err("alpha power response malformed or overpromoted")
	}
	if !a.SharedS.SameCoordinateUsed || !a.SharedS.BothCodomainsTyped || a.SharedS.TransportIntoBothCertified || a.SharedS.PowerOrderDerived || !containsAll(a.SharedS.Failures, []string{FailureSameSNotTransportedToBoth, FailureNoPowerResponseTransport}) {
		return err("shared S_split audit malformed")
	}
	if !a.Denominators.TypedBoundaryAugmentedDomains || a.Denominators.ActivationTheoremCertified || !containsAll(a.Denominators.Failures, []string{FailureDenominatorTypingNotActivation}) {
		return err("denominator audit malformed")
	}
	if !a.Obstruction.SocketRankPowerShapeCoherent || a.Obstruction.TransportMapCertified || a.Obstruction.LinearTransportCertified || a.Obstruction.QuadraticTransportCertified || a.Obstruction.PowerOrderDerived || a.Obstruction.AlphaNative {
		return err("transport obstruction overpromoted")
	}
	if !a.R3.YDagYReadoutCarrierReady || !a.R3.SocketMagnitudeTransferTyped || a.R3.AlphaNative || a.R3.BoundaryPowerTransportNative || a.R3.SectorTraceMagnitudeReadout || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		return err("R3 assessment overpromoted")
	}
	if a.Impact.NativeAlphaSourceSolved || a.Impact.BoundaryPowerTransportSolved || a.Impact.SocketMagnitudeSourceNative || a.Impact.SectorTraceReadout || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
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
	return f.Enforced && f.NoPowerResponseTransport && f.NoLinearTransport && f.NoQuadraticTransport && f.LinearVsQuadraticNotDerived && f.SameSNotTransportedToBoth && f.SocketRankRatiosNotActivation && f.DenominatorTypingNotActivation && f.AlphaStillSealed && f.NoNativeAlphaSource && f.NoBoundaryAlphaTransport && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NotR3 && f.NotR4 && f.Verdict == StatusFirewallVerdict
}

func Statuses() []string {
	return []string{
		StatusGate866Inherited, StatusLinearLaneAudited, StatusQuadraticLaneAudited, StatusSharedSTransportDefined, StatusPowerOrderAudited, StatusDenominatorsReaudited, StatusAlphaReconstructed, StatusTransportFirewall, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusNextWoundIdentified, StatusFirewallVerdict,
		SupportAlphaPowerResponseShape, SupportLinearFirstOrderCandidate, SupportQuadraticSecondOrderCandidate, SupportDenominatorsTypedDomains, SupportWoundReducesToTransportLaw, SupportSocketRankShapeInherited, SupportSharedCoordinatePressure, SupportLinearCoefficientFromPiTop, SupportQuadraticCoefficientFromHRMin,
		FailureNoPowerResponseTransport, FailureNoLinearTransport, FailureNoQuadraticTransport, FailureLinearVsQuadraticNotDerived, FailureSameSNotTransportedToBoth, FailureSocketRankRatiosNotActivation, FailureDenominatorTypingNotActivation, FailureAlphaStillSealed, FailureNoNativeAlphaSource, FailureNoBoundaryAlphaTransport, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNumericalYukawa, FailureNoSectorTraceMagnitude, FailureNoFullUnbrokenAF, FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNotR3, FailureNotR4,
	}
}

func FormatLane(l PowerLane) string {
	return fmt.Sprintf("%s: %s/%s = %d/%d = %.15g, contribution=%.16g, power=s^%d, transportCertified=%t, responseOrderDerived=%t", l.Name, l.NumeratorSource, l.DenominatorSource, l.Numerator, l.Denominator, l.Coefficient, l.Contribution, l.Power, l.TransportCertified, l.ResponseOrderDerived)
}

func FormatAlpha(a AlphaPowerResponse) string {
	return fmt.Sprintf("%s | %s | s=%.16g, alpha=%.16g, reconstructed=%.16g, shapeCoherent=%t, native=%t, transportMapCertified=%t, linear=(%s), quadratic=(%s)", a.Expression, a.SocketRankExpression, a.S, AlphaB, a.ReconstructedAlpha, a.ShapeCoherent, a.Native, a.TransportMapCertified, FormatLane(a.Linear), FormatLane(a.Quadratic))
}

func FormatSharedS(s SharedCoordinateAudit) string {
	return fmt.Sprintf("coordinate=%s, linearCodomain=%s, quadraticCodomain=%s, sameCoordinateUsed=%t, bothCodomainsTyped=%t, transportIntoBothCertified=%t, powerOrderDerived=%t", s.Coordinate, s.LinearCodomain, s.QuadraticCodomain, s.SameCoordinateUsed, s.BothCodomainsTyped, s.TransportIntoBothCertified, s.PowerOrderDerived)
}

func FormatDenominators(d DenominatorAudit) string {
	return fmt.Sprintf("linear=%d (%s), quadratic=%d (%s), typedBoundaryAugmentedDomains=%t, activationTheoremCertified=%t", d.LinearDenominator, d.LinearMeaning, d.QuadraticDenominator, d.QuadraticMeaning, d.TypedBoundaryAugmentedDomains, d.ActivationTheoremCertified)
}

func FormatObstruction(o TransportObstruction) string {
	return fmt.Sprintf("shapeCoherent=%t, linearTransport=%t, quadraticTransport=%t, powerOrderDerived=%t, transportMapCertified=%t, alphaNative=%t, remaining=%s, next=%s", o.SocketRankPowerShapeCoherent, o.LinearTransportCertified, o.QuadraticTransportCertified, o.PowerOrderDerived, o.TransportMapCertified, o.AlphaNative, o.RemainingWound, o.NextGate)
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("YdagYReady=%t, socketMagnitudeTransferTyped=%t, alphaNative=%t, boundaryPowerTransportNative=%t, sectorTraceReadout=%t, eligibleR3=%t, eligibleR4=%t", r.YDagYReadoutCarrierReady, r.SocketMagnitudeTransferTyped, r.AlphaNative, r.BoundaryPowerTransportNative, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s, status=%s, alphaSourceTypingSharpened=%t, nativeAlphaSolved=%t, boundaryPowerTransportSolved=%t, canUpdateNEff=%t, canPromoteR3=%t", i.Classification, i.Status, i.AlphaSourceTypingSharpened, i.NativeAlphaSourceSolved, i.BoundaryPowerTransportSolved, i.CanUpdateNEff, i.CanPromoteToR3)
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g, s=%.16g, officialNEff=%.16g, frozen=%t, alphaNative=%t, powerTransport=%t, R3=%t, R4=%t", l.AlphaB, l.SBoundary, l.OfficialNEff, l.OfficialFrozen, l.AlphaNative, l.PowerTransport, l.R3, l.R4)
}

func FormatFirewalls(f Firewalls) string {
	parts := []string{}
	if f.NoPowerResponseTransport {
		parts = append(parts, FailureNoPowerResponseTransport)
	}
	if f.NoLinearTransport {
		parts = append(parts, FailureNoLinearTransport)
	}
	if f.NoQuadraticTransport {
		parts = append(parts, FailureNoQuadraticTransport)
	}
	if f.LinearVsQuadraticNotDerived {
		parts = append(parts, FailureLinearVsQuadraticNotDerived)
	}
	if f.AlphaStillSealed {
		parts = append(parts, FailureAlphaStillSealed)
	}
	if f.NotR3 {
		parts = append(parts, FailureNotR3)
	}
	return fmt.Sprintf("enforced=%t, verdict=%s, failures=[%s]", f.Enforced, f.Verdict, strings.Join(parts, "; "))
}
