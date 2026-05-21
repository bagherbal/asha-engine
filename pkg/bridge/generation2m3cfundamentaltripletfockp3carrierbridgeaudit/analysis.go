// Package generation2m3cfundamentaltripletfockp3carrierbridgeaudit
// implements Gate 833: M_3(C) Fundamental Triplet / Fock P_3
// Carrier-Bridge Audit.
//
// Gate 833 follows Gate 832's SectorTraceLedgerMap source/carrier obstruction.
// Gate 832 source-typed A_F=C plus H plus M_3(C) as the strongest finite-sector
// projector candidate source, but found no map from the R2++ aggregate carrier
// I_3 plus (P_1 plus P_3) to finite sector projectors.  Gate 833 audits the
// nearest concrete candidate bridge: the M_3(C) fundamental color triplet versus
// the Fock/projective P_3 triplet.  It allows the carrier-shape resonance C^3
// versus P_3 W to be recorded, but refuses to identify them without a canonical
// intertwiner, representation action, Morita bridge, or trace-representation
// theorem.  It also preserves the firewall that a carrier bridge is not a trace
// magnitude readout and not a Yukawa theorem.
package generation2m3cfundamentaltripletfockp3carrierbridgeaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE833-M3C-FUNDAMENTAL-TRIPLET-FOCK-P3-CARRIER-BRIDGE-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	TopBlockDim        = 3
	M3FundamentalDim   = 3
	M3MatrixUnitCount  = 9
	FockWDim           = 4
	FockP1Rank         = 1
	FockP3Rank         = 3
	RestBlockDim       = 4
	AggregateAtomCount = 7
	K7Dim              = 7

	StatusGate832Inherited               = "PASS_GATE832_SECTOR_LEDGER_SOURCE_CARRIER_OBSTRUCTION_INHERITED"
	StatusM3FundamentalTripletAudited    = "PASS_M3C_FUNDAMENTAL_TRIPLET_CARRIER_AUDITED"
	StatusFockP3TripletAudited           = "PASS_FOCK_P3_TRIPLET_CARRIER_AUDITED"
	StatusCarrierShapeComparisonAudited  = "PASS_M3C_FUNDAMENTAL_AND_FOCK_P3_CARRIER_SHAPE_COMPARISON_AUDITED"
	StatusIntertwinerRoutesAudited       = "PASS_INTERTWINER_MORITA_TRACE_REPRESENTATION_ROUTES_AUDITED"
	StatusTopI3CompatibilityAudited      = "PASS_TOP_I3_CARRIER_COMPATIBILITY_AUDITED"
	StatusSectorProjectorMapStillMissing = "PASS_SECTOR_PROJECTOR_MAP_REMAINS_MISSING_OBJECT"
	StatusMagnitudeFirewallPreserved     = "PASS_CARRIER_BRIDGE_NOT_TRACE_MAGNITUDE_FIREWALL_PRESERVED"
	StatusNoLedgerUpdates                = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusPhysicalFirewalls              = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusR2PlusPlusRetained             = "PASS_R2_PLUS_PLUS_STATUS_RETAINED_NOT_R3"
	StatusNoObservedDataUsed             = "PASS_NO_OBSERVED_YUKAWA_MASS_CKM_PMNS_DATA_USED"
	StatusFirewallGate833                = "FIREWALL_PRESERVED_GATE833_M3C_FUNDAMENTAL_TRIPLET_FOCK_P3_CARRIER_BRIDGE_OBSTRUCTION"

	SupportM3SuppliesCanonicalFundamentalTriplet = "CONDITIONAL_SUPPORT_M3C_SUPPLIES_CANONICAL_FUNDAMENTAL_TRIPLET_CARRIER"
	SupportM3MatrixUnitsSourceCarrierProjectors  = "CONDITIONAL_SUPPORT_M3C_MATRIX_UNITS_SOURCE_COLOR_CARRIER_PROJECTORS"
	SupportFockP3SuppliesRankThreeCarrier        = "CONDITIONAL_SUPPORT_FOCK_P3_SUPPLIES_RANK_THREE_PROJECTIVE_TRIPLET_CARRIER"
	SupportCarrierShapesMatchDimensionThree      = "CONDITIONAL_SUPPORT_M3C_FUNDAMENTAL_AND_FOCK_P3_HAVE_MATCHING_DIMENSION_THREE_CARRIER_SHAPE"
	SupportFormalIsomorphismsExist               = "CONDITIONAL_SUPPORT_FORMAL_C3_TO_P3W_ISOMORPHISMS_EXIST"
	SupportBridgeRequiresCanonicalIntertwiner    = "CONDITIONAL_SUPPORT_CARRIER_BRIDGE_REQUIRES_CANONICAL_INTERTWINER_OR_ACTION"
	SupportSectorProjectorMapStillNext           = "CONDITIONAL_SUPPORT_SECTOR_PROJECTOR_MAP_REMAINS_NEXT_MISSING_OBJECT"
	SupportCarrierBridgeNotMagnitude             = "CONDITIONAL_SUPPORT_CARRIER_BRIDGE_WOULD_STILL_NOT_BE_TRACE_MAGNITUDE_READOUT"
	SupportAggregateCarrierStillR2PlusPlus       = "CONDITIONAL_SUPPORT_AGGREGATE_CARRIER_REMAINS_R2_PLUS_PLUS"
	SupportDualTripletSeparation                 = "CONDITIONAL_SUPPORT_DUAL_TRIPLET_SOURCE_TYPES_REMAIN_DISTINCT"

	FailureNoCanonicalM3ToFockP3Intertwiner           = "FAILED_ROUTE_NO_CANONICAL_M3C_TO_FOCK_P3_INTERTWINER_CERTIFIED"
	FailureM3ColorTripletNotIdentifiedWithFockP3      = "FAILED_ROUTE_M3C_COLOR_TRIPLET_NOT_IDENTIFIED_WITH_FOCK_P3_TRIPLET"
	FailureNoM3ActionOnFockP3                         = "FAILED_ROUTE_NO_M3C_ACTION_ON_FOCK_P3W_CERTIFIED"
	FailureNoFockP3RepresentationOfM3                 = "FAILED_ROUTE_NO_FOCK_P3_REPRESENTATION_OF_M3C_CERTIFIED"
	FailureNoIntertwiningLaw                          = "FAILED_ROUTE_NO_INTERTWINING_LAW_BETWEEN_M3C_FUNDAMENTAL_AND_FOCK_P3"
	FailureNoMoritaBridge                             = "FAILED_ROUTE_NO_MORITA_BIMODULE_BRIDGE_CERTIFIED"
	FailureNoTraceRepresentationBridge                = "FAILED_ROUTE_NO_TRACE_REPRESENTATION_BRIDGE_CERTIFIED"
	FailureTopI3NotCarrierCompatibleWithM3Fundamental = "FAILED_ROUTE_TOP_I3_NOT_CARRIER_COMPATIBLE_WITH_M3C_FUNDAMENTAL_TRIPLET"
	FailureTopI3NotIdentifiedWithFockP3               = "FAILED_ROUTE_TOP_I3_NOT_IDENTIFIED_WITH_FOCK_P3_TRIPLET"
	FailureShapeMatchOnlyNotTypedBridge               = "FAILED_ROUTE_DIMENSION_THREE_SHAPE_MATCH_NOT_TYPED_CARRIER_BRIDGE"
	FailureCarrierBridgeNotSectorLedger               = "FAILED_ROUTE_CARRIER_BRIDGE_NOT_SECTOR_TRACE_LEDGER"
	FailureCarrierBridgeNotTraceMagnitude             = "FAILED_ROUTE_CARRIER_BRIDGE_NOT_TRACE_MAGNITUDE_READOUT"
	FailureNoSectorProjectorMap                       = "FAILED_ROUTE_NO_SECTOR_PROJECTOR_MAP_CERTIFIED"
	FailureNoSectorTraceLedgerMap                     = "FAILED_ROUTE_NO_SECTOR_TRACE_LEDGER_MAP_CERTIFIED"
	FailureNoSectorMagnitudeReadout                   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureSectorProjectorsDoNotSupplyMagnitudes      = "FAILED_ROUTE_SECTOR_PROJECTORS_DO_NOT_SUPPLY_TRACE_MAGNITUDES"
	FailureAggregateOperatorNotR3                     = "FAILED_ROUTE_AGGREGATE_OPERATOR_NOT_R3"
	FailureR2NotR3                                    = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNoNativeYukawaOperator                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureAlphaStillSealed                           = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoBoundaryAlphaMap                         = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureNoNEffUpdate                               = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                            = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit                        = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoPMNSCKM                                  = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoSMParticleAssignment                     = "FAILED_ROUTE_NO_STANDARD_MODEL_PARTICLE_ASSIGNMENT_FROM_TRIPLET_CARRIER_SHAPE"
)

type Ledger struct {
	S, AlphaB                                            float64
	AggregateCarrier                                     string
	OperatorNEff, OfficialNEff                           float64
	TopBlockDim, RestBlockDim, AggregateAtomCount, K7Dim int
	R2PlusPlusConsolidated                               bool
	R3SectorLedgerCertified                              bool
	AlphaSealed                                          bool
}

type M3FundamentalTripletAudit struct {
	Algebra, FundamentalCarrier, MatrixUnitCarrier string
	Dim, MatrixUnitCount, IdentityTrace            int
	CanonicalCarrier, MatrixUnitsAudited           bool
	SuppliesCarrierProjectors                      bool
	SuppliesYukawaMagnitudes                       bool
	UsesObservedYukawaData                         bool
	Verdicts, Supports, Failures                   []string
}

type FockP3TripletAudit struct {
	Source, Carrier, Selector    string
	WDim, P1Rank, P3Rank         int
	RankThreeCarrier             bool
	BMinusLSelectorTyped         bool
	M3ActionCertified            bool
	RepresentsM3                 bool
	Verdicts, Supports, Failures []string
}

type CarrierShapeBridgeAudit struct {
	LeftCarrier, RightCarrier                           string
	LeftDim, RightDim                                   int
	DimensionMatches, FormalIsomorphismsExist           bool
	CanonicalIntertwinerCertified, TypedBridgeCertified bool
	TripletsIdentified, NonCircular                     bool
	Verdicts, Supports, Failures                        []string
}

type IntertwinerRouteAudit struct {
	FiniteAlgebraRouteAudited, MoritaRouteAudited, TraceRepresentationRouteAudited bool
	ActionLawCertified, CommutationLawCertified, RepresentationLawCertified        bool
	CanonicalIntertwinerCertified, CarrierBridgeCertified                          bool
	Verdicts, Supports, Failures                                                   []string
}

type TopI3CompatibilityAudit struct {
	TopBlock, M3Identity, FockP3Identity                string
	TopDim, M3Dim, FockP3Dim                            int
	IdentityShapeMatches                                bool
	SameAsM3FundamentalCertified, SameAsFockP3Certified bool
	TopToM3ToP3ChainCertified                           bool
	Verdicts, Supports, Failures                        []string
}

type SectorImpact struct {
	CarrierShapeSupport, CarrierBridgeCertified, SectorProjectorMapCertified bool
	SectorTraceLedgerCertified, TraceMagnitudeReadoutCertified               bool
	CanPromoteToR3, CanPromoteToR4                                           bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs                         bool
	CurrentLevel, NextMissingObject, NextGate, Reason                        string
	Verdicts, Supports, Failures                                             []string
}

type Firewalls struct {
	Enforced                                                    bool
	ShapeNotBridge, NoM3P3Intertwiner, NoM3ActionOnP3           bool
	TopI3Separated, CarrierBridgeNotSectorLedger                bool
	NoSectorProjectorMap, NoSectorLedgerMap, NoMagnitudeReadout bool
	AlphaSealed, NoBoundaryAlphaMap                             bool
	NotR3, NotR4                                                bool
	NoNEffUpdate, NoCYukawaUpdate                               bool
	NoObservedYukawaFit, NoPMNSCKM, NoParticleAssignment        bool
	Verdict                                                     string
}

type Analysis struct {
	Ledger      Ledger
	M3Carrier   M3FundamentalTripletAudit
	FockP3      FockP3TripletAudit
	ShapeBridge CarrierShapeBridgeAudit
	Routes      IntertwinerRouteAudit
	TopI3       TopI3CompatibilityAudit
	Impact      SectorImpact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func BuildDefault() (Analysis, error) {
	if TopBlockDim != M3FundamentalDim || M3FundamentalDim != FockP3Rank {
		return Analysis{}, fmt.Errorf("triplet dimension mismatch: top=%d M3=%d P3=%d", TopBlockDim, M3FundamentalDim, FockP3Rank)
	}
	if M3MatrixUnitCount != M3FundamentalDim*M3FundamentalDim {
		return Analysis{}, fmt.Errorf("M3 matrix-unit count mismatch: %d != %d^2", M3MatrixUnitCount, M3FundamentalDim)
	}
	if FockP1Rank+FockP3Rank != RestBlockDim {
		return Analysis{}, fmt.Errorf("Fock selector rank mismatch: %d+%d != %d", FockP1Rank, FockP3Rank, RestBlockDim)
	}
	if TopBlockDim+RestBlockDim != AggregateAtomCount || AggregateAtomCount != K7Dim {
		return Analysis{}, fmt.Errorf("aggregate seven-count mismatch: top=%d rest=%d aggregate=%d K7=%d", TopBlockDim, RestBlockDim, AggregateAtomCount, K7Dim)
	}

	ledger := Ledger{
		S:                       SBoundary,
		AlphaB:                  AlphaB,
		AggregateCarrier:        "I_3 plus (P_1 plus P_3)",
		OperatorNEff:            OperatorNEff,
		OfficialNEff:            OfficialNEff,
		TopBlockDim:             TopBlockDim,
		RestBlockDim:            RestBlockDim,
		AggregateAtomCount:      AggregateAtomCount,
		K7Dim:                   K7Dim,
		R2PlusPlusConsolidated:  true,
		R3SectorLedgerCertified: false,
		AlphaSealed:             true,
	}

	m3 := M3FundamentalTripletAudit{
		Algebra:                   "M_3(C) subset A_F=C plus H plus M_3(C)",
		FundamentalCarrier:        "C^3_color fundamental module",
		MatrixUnitCarrier:         "E_ij matrix units, i,j=1..3",
		Dim:                       M3FundamentalDim,
		MatrixUnitCount:           M3MatrixUnitCount,
		IdentityTrace:             M3FundamentalDim,
		CanonicalCarrier:          true,
		MatrixUnitsAudited:        true,
		SuppliesCarrierProjectors: true,
		SuppliesYukawaMagnitudes:  false,
		UsesObservedYukawaData:    false,
		Verdicts:                  []string{StatusM3FundamentalTripletAudited, StatusNoObservedDataUsed},
		Supports:                  []string{SupportM3SuppliesCanonicalFundamentalTriplet, SupportM3MatrixUnitsSourceCarrierProjectors, SupportCarrierBridgeNotMagnitude},
		Failures:                  []string{FailureCarrierBridgeNotTraceMagnitude, FailureSectorProjectorsDoNotSupplyMagnitudes},
	}

	fock := FockP3TripletAudit{
		Source:               "Fock/projective B-L selector W=C^4 with P_1 plus P_3=I_4",
		Carrier:              "P_3 W projective triplet",
		Selector:             "B-L = -P_1 + (1/3)P_3",
		WDim:                 FockWDim,
		P1Rank:               FockP1Rank,
		P3Rank:               FockP3Rank,
		RankThreeCarrier:     true,
		BMinusLSelectorTyped: true,
		M3ActionCertified:    false,
		RepresentsM3:         false,
		Verdicts:             []string{StatusFockP3TripletAudited},
		Supports:             []string{SupportFockP3SuppliesRankThreeCarrier, SupportDualTripletSeparation},
		Failures:             []string{FailureNoM3ActionOnFockP3, FailureNoFockP3RepresentationOfM3, FailureCarrierBridgeNotTraceMagnitude},
	}

	shape := CarrierShapeBridgeAudit{
		LeftCarrier:                   "C^3_color from M_3(C)",
		RightCarrier:                  "P_3 W from Fock/projective B-L selector",
		LeftDim:                       M3FundamentalDim,
		RightDim:                      FockP3Rank,
		DimensionMatches:              true,
		FormalIsomorphismsExist:       true,
		CanonicalIntertwinerCertified: false,
		TypedBridgeCertified:          false,
		TripletsIdentified:            false,
		NonCircular:                   true,
		Verdicts:                      []string{StatusCarrierShapeComparisonAudited},
		Supports:                      []string{SupportCarrierShapesMatchDimensionThree, SupportFormalIsomorphismsExist, SupportBridgeRequiresCanonicalIntertwiner},
		Failures:                      []string{FailureNoCanonicalM3ToFockP3Intertwiner, FailureM3ColorTripletNotIdentifiedWithFockP3, FailureShapeMatchOnlyNotTypedBridge, FailureNoIntertwiningLaw},
	}

	routes := IntertwinerRouteAudit{
		FiniteAlgebraRouteAudited:       true,
		MoritaRouteAudited:              true,
		TraceRepresentationRouteAudited: true,
		ActionLawCertified:              false,
		CommutationLawCertified:         false,
		RepresentationLawCertified:      false,
		CanonicalIntertwinerCertified:   false,
		CarrierBridgeCertified:          false,
		Verdicts:                        []string{StatusIntertwinerRoutesAudited},
		Supports:                        []string{SupportBridgeRequiresCanonicalIntertwiner, SupportSectorProjectorMapStillNext},
		Failures:                        []string{FailureNoM3ActionOnFockP3, FailureNoIntertwiningLaw, FailureNoMoritaBridge, FailureNoTraceRepresentationBridge, FailureNoCanonicalM3ToFockP3Intertwiner},
	}

	top := TopI3CompatibilityAudit{
		TopBlock:                     "I_3 top/dominant aggregate trace block",
		M3Identity:                   "I_{C^3} identity on M_3(C) fundamental carrier",
		FockP3Identity:               "P_3 identity on P_3 W",
		TopDim:                       TopBlockDim,
		M3Dim:                        M3FundamentalDim,
		FockP3Dim:                    FockP3Rank,
		IdentityShapeMatches:         true,
		SameAsM3FundamentalCertified: false,
		SameAsFockP3Certified:        false,
		TopToM3ToP3ChainCertified:    false,
		Verdicts:                     []string{StatusTopI3CompatibilityAudited},
		Supports:                     []string{SupportCarrierShapesMatchDimensionThree, SupportDualTripletSeparation},
		Failures:                     []string{FailureTopI3NotCarrierCompatibleWithM3Fundamental, FailureTopI3NotIdentifiedWithFockP3, FailureShapeMatchOnlyNotTypedBridge},
	}

	impact := SectorImpact{
		CarrierShapeSupport:            true,
		CarrierBridgeCertified:         false,
		SectorProjectorMapCertified:    false,
		SectorTraceLedgerCertified:     false,
		TraceMagnitudeReadoutCertified: false,
		CanPromoteToR3:                 false,
		CanPromoteToR4:                 false,
		CanUpdateNEff:                  false,
		CanUpdateCYukawa:               false,
		CanUpdateCHiggs:                false,
		CurrentLevel:                   "R2++ aggregate trace operator with conditional dimension-three carrier-shape resonance; not R3 sector ledger",
		NextMissingObject:              "canonical M_3(C) fundamental -> Fock P_3 intertwiner/action, then SectorProjectorMap, then SectorTraceMagnitudeReadoutMap",
		NextGate:                       "Gate 834 — SectorProjectorMap Construction/Obstruction Audit only if a typed carrier bridge appears; otherwise external sector-ledger airlock remains the honest path",
		Reason:                         "M_3(C) supplies a canonical C^3 fundamental carrier and P_3 W is a rank-three Fock/projective carrier, but equal dimension and formal isomorphisms do not supply a canonical typed bridge, an action of M_3(C) on P_3 W, or a trace-magnitude readout.",
		Verdicts:                       []string{StatusGate832Inherited, StatusSectorProjectorMapStillMissing, StatusMagnitudeFirewallPreserved, StatusR2PlusPlusRetained, StatusNoLedgerUpdates},
		Supports:                       []string{SupportCarrierShapesMatchDimensionThree, SupportFormalIsomorphismsExist, SupportSectorProjectorMapStillNext, SupportCarrierBridgeNotMagnitude, SupportAggregateCarrierStillR2PlusPlus},
		Failures:                       []string{FailureNoCanonicalM3ToFockP3Intertwiner, FailureNoM3ActionOnFockP3, FailureTopI3NotCarrierCompatibleWithM3Fundamental, FailureCarrierBridgeNotSectorLedger, FailureNoSectorProjectorMap, FailureNoSectorTraceLedgerMap, FailureNoSectorMagnitudeReadout, FailureAggregateOperatorNotR3, FailureR2NotR3, FailureNoNativeYukawaOperator, FailureAlphaStillSealed, FailureNoBoundaryAlphaMap, FailureNoNEffUpdate, FailureNoCYukawaUpdate},
	}

	firewalls := Firewalls{
		Enforced:                     true,
		ShapeNotBridge:               true,
		NoM3P3Intertwiner:            true,
		NoM3ActionOnP3:               true,
		TopI3Separated:               true,
		CarrierBridgeNotSectorLedger: true,
		NoSectorProjectorMap:         true,
		NoSectorLedgerMap:            true,
		NoMagnitudeReadout:           true,
		AlphaSealed:                  true,
		NoBoundaryAlphaMap:           true,
		NotR3:                        true,
		NotR4:                        true,
		NoNEffUpdate:                 true,
		NoCYukawaUpdate:              true,
		NoObservedYukawaFit:          true,
		NoPMNSCKM:                    true,
		NoParticleAssignment:         true,
		Verdict:                      StatusFirewallGate833,
	}

	truth := "Gate 833 audits the nearest concrete triplet bridge: M_3(C)'s canonical C^3 fundamental carrier versus the Fock/projective P_3 W triplet.  It conditionally supports the matching carrier shape and the existence of formal C^3 isomorphisms, but blocks identification because no canonical intertwiner, no M_3(C) action on P_3 W, no Morita bridge, no trace-representation bridge, and no top-I_3 carrier compatibility theorem are certified."
	final := "Outcome: partial carrier-shape support plus obstruction.  The M_3(C) fundamental triplet and Fock P_3 triplet are dimension-compatible candidates only; there is no SectorProjectorMap, no sector trace-magnitude readout, no R3 promotion, and no native Yukawa theorem."

	return Analysis{Ledger: ledger, M3Carrier: m3, FockP3: fock, ShapeBridge: shape, Routes: routes, TopI3: top, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func containsAll(haystack []string, needles []string) bool {
	m := make(map[string]bool, len(haystack))
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

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.16g alpha_B=%.16g carrier=%q operator_N_eff=%.16g official_N_eff=%.16g dims(top=%d rest=%d total=%d K7=%d) R2++=%t R3=%t alpha_sealed=%t", l.S, l.AlphaB, l.AggregateCarrier, l.OperatorNEff, l.OfficialNEff, l.TopBlockDim, l.RestBlockDim, l.AggregateAtomCount, l.K7Dim, l.R2PlusPlusConsolidated, l.R3SectorLedgerCertified, l.AlphaSealed)
}

func FormatM3Carrier(a M3FundamentalTripletAudit) string {
	return fmt.Sprintf("algebra=%q carrier=%q matrix_units=%q dim=%d unit_count=%d trace_I=%d canonical=%t units_audited=%t carrier_projectors=%t magnitudes=%t observed_data=%t verdicts=%s supports=%s failures=%s", a.Algebra, a.FundamentalCarrier, a.MatrixUnitCarrier, a.Dim, a.MatrixUnitCount, a.IdentityTrace, a.CanonicalCarrier, a.MatrixUnitsAudited, a.SuppliesCarrierProjectors, a.SuppliesYukawaMagnitudes, a.UsesObservedYukawaData, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFockP3(a FockP3TripletAudit) string {
	return fmt.Sprintf("source=%q carrier=%q selector=%q Wdim=%d P1=%d P3=%d rank_three=%t B-L=%t M3_action=%t represents_M3=%t verdicts=%s supports=%s failures=%s", a.Source, a.Carrier, a.Selector, a.WDim, a.P1Rank, a.P3Rank, a.RankThreeCarrier, a.BMinusLSelectorTyped, a.M3ActionCertified, a.RepresentsM3, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatShapeBridge(a CarrierShapeBridgeAudit) string {
	return fmt.Sprintf("left=%q dim=%d right=%q dim=%d dimension_match=%t formal_isomorphisms=%t canonical_intertwiner=%t typed_bridge=%t identified=%t noncircular=%t verdicts=%s supports=%s failures=%s", a.LeftCarrier, a.LeftDim, a.RightCarrier, a.RightDim, a.DimensionMatches, a.FormalIsomorphismsExist, a.CanonicalIntertwinerCertified, a.TypedBridgeCertified, a.TripletsIdentified, a.NonCircular, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatRoutes(a IntertwinerRouteAudit) string {
	return fmt.Sprintf("finite_algebra=%t morita=%t trace_rep=%t action=%t commutation=%t representation=%t canonical_intertwiner=%t bridge=%t verdicts=%s supports=%s failures=%s", a.FiniteAlgebraRouteAudited, a.MoritaRouteAudited, a.TraceRepresentationRouteAudited, a.ActionLawCertified, a.CommutationLawCertified, a.RepresentationLawCertified, a.CanonicalIntertwinerCertified, a.CarrierBridgeCertified, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatTopI3(a TopI3CompatibilityAudit) string {
	return fmt.Sprintf("top=%q dim=%d M3=%q dim=%d P3=%q dim=%d shape=%t same_M3=%t same_P3=%t chain=%t verdicts=%s supports=%s failures=%s", a.TopBlock, a.TopDim, a.M3Identity, a.M3Dim, a.FockP3Identity, a.FockP3Dim, a.IdentityShapeMatches, a.SameAsM3FundamentalCertified, a.SameAsFockP3Certified, a.TopToM3ToP3ChainCertified, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatImpact(i SectorImpact) string {
	return fmt.Sprintf("shape_support=%t bridge=%t sector_projector_map=%t sector_ledger=%t magnitude_readout=%t promote_R3=%t promote_R4=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t current=%q next_missing=%q next_gate=%q reason=%q verdicts=%s supports=%s failures=%s", i.CarrierShapeSupport, i.CarrierBridgeCertified, i.SectorProjectorMapCertified, i.SectorTraceLedgerCertified, i.TraceMagnitudeReadoutCertified, i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CurrentLevel, i.NextMissingObject, i.NextGate, i.Reason, strings.Join(i.Verdicts, ","), strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func Statuses() []string {
	return []string{
		StatusGate832Inherited,
		StatusM3FundamentalTripletAudited,
		StatusFockP3TripletAudited,
		StatusCarrierShapeComparisonAudited,
		StatusIntertwinerRoutesAudited,
		StatusTopI3CompatibilityAudited,
		StatusSectorProjectorMapStillMissing,
		StatusMagnitudeFirewallPreserved,
		StatusNoLedgerUpdates,
		StatusPhysicalFirewalls,
		StatusR2PlusPlusRetained,
		StatusNoObservedDataUsed,
		StatusFirewallGate833,
		SupportM3SuppliesCanonicalFundamentalTriplet,
		SupportM3MatrixUnitsSourceCarrierProjectors,
		SupportFockP3SuppliesRankThreeCarrier,
		SupportCarrierShapesMatchDimensionThree,
		SupportFormalIsomorphismsExist,
		SupportBridgeRequiresCanonicalIntertwiner,
		SupportSectorProjectorMapStillNext,
		SupportCarrierBridgeNotMagnitude,
		SupportAggregateCarrierStillR2PlusPlus,
		SupportDualTripletSeparation,
		FailureNoCanonicalM3ToFockP3Intertwiner,
		FailureM3ColorTripletNotIdentifiedWithFockP3,
		FailureNoM3ActionOnFockP3,
		FailureNoFockP3RepresentationOfM3,
		FailureNoIntertwiningLaw,
		FailureNoMoritaBridge,
		FailureNoTraceRepresentationBridge,
		FailureTopI3NotCarrierCompatibleWithM3Fundamental,
		FailureTopI3NotIdentifiedWithFockP3,
		FailureShapeMatchOnlyNotTypedBridge,
		FailureCarrierBridgeNotSectorLedger,
		FailureCarrierBridgeNotTraceMagnitude,
		FailureNoSectorProjectorMap,
		FailureNoSectorTraceLedgerMap,
		FailureNoSectorMagnitudeReadout,
		FailureSectorProjectorsDoNotSupplyMagnitudes,
		FailureAggregateOperatorNotR3,
		FailureR2NotR3,
		FailureNoNativeYukawaOperator,
		FailureAlphaStillSealed,
		FailureNoBoundaryAlphaMap,
		FailureNoNEffUpdate,
		FailureNoCYukawaUpdate,
		FailureNoObservedYukawaFit,
		FailureNoPMNSCKM,
		FailureNoSMParticleAssignment,
	}
}
