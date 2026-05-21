// Package generation2sectortraceledgermapcandidatesourceandcarriercompatibilityaudit
// implements Gate 832: SectorTraceLedgerMap Candidate Source and
// Carrier-Compatibility Audit.
//
// Gate 832 follows Gate 831's R2++ / R3 firewall.  Gate 831 established that
// the aggregate trace carrier
//
//	I_3 ⊕ (P_1 ⊕ P_3)
//
// is not a K7 projector theorem and not an R3 sector ledger.  Gate 832 audits
// the next lawful source candidate for a sector ledger: the finite internal
// spectral-triple algebra A_F=C ⊕ H ⊕ M_3(C).  The gate asks whether finite
// sector projector candidates from A_F are compatible with the aggregate
// carrier and whether they supply trace magnitudes.  The expected honest
// result is source-typing plus obstruction: A_F is the strongest sector
// projector source, but no SectorTraceLedgerMap, color/Fock triplet bridge,
// aggregate-to-K7 map, or sector trace-magnitude readout is certified.
package generation2sectortraceledgermapcandidatesourceandcarriercompatibilityaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE832-SECTOR-TRACE-LEDGER-MAP-CANDIDATE-SOURCE-CARRIER-COMPATIBILITY-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	TopBlockDim        = 3
	FockP1Rank         = 1
	FockP3Rank         = 3
	RestBlockDim       = 4
	AggregateAtomCount = 7
	K7Dim              = 7

	StatusGate831Inherited                    = "PASS_GATE831_R2_PLUS_PLUS_R3_FIREWALL_INHERITED"
	StatusFiniteSectorCandidatesAudited       = "PASS_FINITE_SECTOR_PROJECTOR_CANDIDATES_AUDITED"
	StatusAggregateCarrierCompatibilityTested = "PASS_AGGREGATE_CARRIER_COMPATIBILITY_TESTED"
	StatusK7SevenCountRouteAudited            = "PASS_K7_SEVEN_COUNT_ROUTE_AUDITED"
	StatusDualTripletBridgeRouteAudited       = "PASS_DUAL_TRIPLET_BRIDGE_ROUTE_AUDITED"
	StatusSectorProjectorVsMagnitudeAudited   = "PASS_SECTOR_PROJECTOR_VS_TRACE_MAGNITUDE_FIREWALL_AUDITED"
	StatusSectorLedgerMissingObjectSharpened  = "PASS_SECTOR_LEDGER_MISSING_OBJECT_SHARPENED"
	StatusNoLedgerUpdates                     = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusPhysicalFirewalls                   = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate832                     = "FIREWALL_PRESERVED_GATE832_SECTOR_TRACE_LEDGER_MAP_SOURCE_AND_CARRIER_COMPATIBILITY_OBSTRUCTION"
	StatusR2PlusPlusRetained                  = "PASS_R2_PLUS_PLUS_STATUS_RETAINED_NOT_R3"
	StatusNoObservedDataUsed                  = "PASS_NO_OBSERVED_YUKAWA_MASS_CKM_PMNS_DATA_USED"

	SupportAFStrongestSectorProjectorSource    = "CONDITIONAL_SUPPORT_A_F_IS_STRONGEST_SECTOR_PROJECTOR_SOURCE"
	SupportFiniteSectorProjectorCandidates     = "CONDITIONAL_SUPPORT_FINITE_ALGEBRA_SUPPLIES_SECTOR_PROJECTOR_CANDIDATES"
	SupportSectorLedgerRequiresExtraReadoutMap = "CONDITIONAL_SUPPORT_SECTOR_TRACE_LEDGER_REQUIRES_EXTRA_READOUT_MAP"
	SupportAggregateCarrierStillR2PlusPlus     = "CONDITIONAL_SUPPORT_AGGREGATE_CARRIER_REMAINS_R2_PLUS_PLUS"
	SupportSectorProjectorMapMissingObject     = "CONDITIONAL_SUPPORT_SECTOR_PROJECTOR_MAP_IS_NOW_SEPARATE_MISSING_OBJECT"
	SupportTraceMagnitudeReadoutMissingObject  = "CONDITIONAL_SUPPORT_SECTOR_TRACE_MAGNITUDE_READOUT_IS_SEPARATE_MISSING_OBJECT"
	SupportSevenCountResonanceOnly             = "CONDITIONAL_SUPPORT_SEVEN_COUNT_RESONANCE_REMAINS_RESONANCE_ONLY"
	SupportDualTripletSeparation               = "CONDITIONAL_SUPPORT_DUAL_TRIPLET_SOURCE_TYPES_REMAIN_DISTINCT"

	FailureNoSectorTraceLedgerMap                     = "FAILED_ROUTE_NO_SECTOR_TRACE_LEDGER_MAP_CERTIFIED"
	FailureSectorProjectorsDoNotSupplyMagnitudes      = "FAILED_ROUTE_SECTOR_PROJECTORS_DO_NOT_SUPPLY_TRACE_MAGNITUDES"
	FailureAggregateNotCompatibleWithSectorProjectors = "FAILED_ROUTE_AGGREGATE_CARRIER_NOT_COMPATIBLE_WITH_FINITE_SECTOR_PROJECTORS"
	FailureNoTypedAggregateToSectorMap                = "FAILED_ROUTE_NO_TYPED_AGGREGATE_CARRIER_TO_FINITE_SECTOR_PROJECTOR_MAP"
	FailureNoCarrierCompatibilityTheorem              = "FAILED_ROUTE_NO_CARRIER_COMPATIBILITY_THEOREM"
	FailureNoSectorMagnitudeReadout                   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureSevenAtomAggregateNotK7                    = "FAILED_ROUTE_SEVEN_ATOM_AGGREGATE_NOT_K7_WITHOUT_TYPED_CARRIER_MAP"
	FailureNoAggregateToK7Map                         = "FAILED_ROUTE_NO_TYPED_AGGREGATE_CARRIER_TO_K7_MAP"
	FailureColorTripletToFockTripletNotCertified      = "FAILED_ROUTE_COLOR_TRIPLET_TO_FOCK_TRIPLET_MAP_NOT_CERTIFIED"
	FailureFiniteAlgebraDoesNotIdentifyTriplets       = "FAILED_ROUTE_FINITE_ALGEBRA_DOES_NOT_IDENTIFY_TOP_TRIPLET_WITH_FOCK_TRIPLET"
	FailureSectorProjectorNotYukawaValue              = "FAILED_ROUTE_SECTOR_PROJECTOR_NOT_YUKAWA_TRACE_VALUE"
	FailureAFSectorIdempotentNotYukawaMagnitude       = "FAILED_ROUTE_A_F_SECTOR_IDEMPOTENT_NOT_YUKAWA_MAGNITUDE"
	FailureAggregateOperatorNotR3                     = "FAILED_ROUTE_AGGREGATE_OPERATOR_NOT_R3"
	FailureR2NotR3                                    = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNoNativeYukawaOperator                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureAlphaStillSealed                           = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoBoundaryAlphaMap                         = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureNoNEffUpdate                               = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                            = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit                        = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoPMNSCKM                                  = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoSMParticleAssignment                     = "FAILED_ROUTE_NO_STANDARD_MODEL_PARTICLE_ASSIGNMENT_FROM_AGGREGATE_ATOMS"
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

type FiniteSectorCandidateAudit struct {
	Algebra                          string
	CandidateProjectorSource         string
	StrongestLawfulSource            bool
	SectorProjectorCandidatesAudited bool
	SuppliesGaugeInternalCarriers    bool
	SuppliesYukawaMagnitudes         bool
	SuppliesTraceMagnitudeReadout    bool
	UsesObservedYukawaData           bool
	Verdicts, Supports, Failures     []string
}

type AggregateCarrierCompatibilityAudit struct {
	DomainCarrier                  string
	TargetProjectors               string
	CarrierMapCertified            bool
	CarrierCompatibilityCertified  bool
	FiniteAlgebraCommutationProven bool
	NonCircular                    bool
	Compatible                     bool
	Verdicts, Supports, Failures   []string
}

type K7RouteAudit struct {
	AggregateExpression                    string
	TopAtoms, RestAtoms, TotalAtoms, K7Dim int
	CountMatchesK7                         bool
	TypedMapCertified                      bool
	ProjectorIdentityCertified             bool
	PromotedToK7                           bool
	Verdicts, Supports, Failures           []string
}

type DualTripletBridgeAudit struct {
	TopTripletRole, FockTripletRole string
	TopTripletDim, FockTripletDim   int
	FiniteAlgebraRouteAudited       bool
	MoritaRouteAudited              bool
	TraceRepresentationRouteAudited bool
	TypedBridgeCertified            bool
	TripletsIdentified              bool
	Verdicts, Supports, Failures    []string
}

type SectorProjectorMagnitudeAudit struct {
	SectorProjectorsFoundAsCandidates bool
	ProjectorsAreCarriers             bool
	ProjectorsAreMagnitudes           bool
	PositiveTraceAtomsDerived         bool
	ReadoutMapCertified               bool
	YukawaValuesDerived               bool
	Verdicts, Supports, Failures      []string
}

type Impact struct {
	CanPromoteToR3, CanPromoteToR4                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	CurrentLevel, NextMissingObject, NextGate        string
	Reason                                           string
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced                                             bool
	AFProjectorsNotMagnitudes, AggregateNotSectorLedger  bool
	NoCarrierMap, NoSectorLedgerMap, NoMagnitudeReadout  bool
	DualTripletSeparated, SevenNotK7                     bool
	AlphaSealed, NoBoundaryAlphaMap                      bool
	NotR3, NotR4                                         bool
	NoNEffUpdate, NoCYukawaUpdate                        bool
	NoObservedYukawaFit, NoPMNSCKM, NoParticleAssignment bool
	Verdict                                              string
}

type Analysis struct {
	Ledger        Ledger
	FiniteSector  FiniteSectorCandidateAudit
	Compatibility AggregateCarrierCompatibilityAudit
	K7Route       K7RouteAudit
	DualTriplet   DualTripletBridgeAudit
	Magnitude     SectorProjectorMagnitudeAudit
	Impact        Impact
	Firewalls     Firewalls
	Truth         string
	Final         string
}

func BuildDefault() (Analysis, error) {
	if TopBlockDim+RestBlockDim != AggregateAtomCount {
		return Analysis{}, fmt.Errorf("aggregate atom count mismatch: %d+%d != %d", TopBlockDim, RestBlockDim, AggregateAtomCount)
	}
	if FockP1Rank+FockP3Rank != RestBlockDim {
		return Analysis{}, fmt.Errorf("rest block rank mismatch: %d+%d != %d", FockP1Rank, FockP3Rank, RestBlockDim)
	}
	if AggregateAtomCount != K7Dim {
		return Analysis{}, fmt.Errorf("expected seven-count resonance mismatch: aggregate=%d K7=%d", AggregateAtomCount, K7Dim)
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

	finite := FiniteSectorCandidateAudit{
		Algebra:                          "A_F = C plus H plus M_3(C)",
		CandidateProjectorSource:         "finite internal spectral-triple algebra sector idempotents / representation projectors",
		StrongestLawfulSource:            true,
		SectorProjectorCandidatesAudited: true,
		SuppliesGaugeInternalCarriers:    true,
		SuppliesYukawaMagnitudes:         false,
		SuppliesTraceMagnitudeReadout:    false,
		UsesObservedYukawaData:           false,
		Verdicts:                         []string{StatusFiniteSectorCandidatesAudited, StatusNoObservedDataUsed},
		Supports:                         []string{SupportAFStrongestSectorProjectorSource, SupportFiniteSectorProjectorCandidates, SupportSectorLedgerRequiresExtraReadoutMap},
		Failures:                         []string{FailureSectorProjectorsDoNotSupplyMagnitudes, FailureSectorProjectorNotYukawaValue, FailureAFSectorIdempotentNotYukawaMagnitude},
	}

	compatibility := AggregateCarrierCompatibilityAudit{
		DomainCarrier:                  "I_3 plus (P_1 plus P_3)",
		TargetProjectors:               "Pi_sector candidates sourced from A_F",
		CarrierMapCertified:            false,
		CarrierCompatibilityCertified:  false,
		FiniteAlgebraCommutationProven: false,
		NonCircular:                    true,
		Compatible:                     false,
		Verdicts:                       []string{StatusAggregateCarrierCompatibilityTested},
		Supports:                       []string{SupportAggregateCarrierStillR2PlusPlus, SupportSectorProjectorMapMissingObject},
		Failures:                       []string{FailureAggregateNotCompatibleWithSectorProjectors, FailureNoTypedAggregateToSectorMap, FailureNoCarrierCompatibilityTheorem, FailureNoSectorTraceLedgerMap},
	}

	k7 := K7RouteAudit{
		AggregateExpression:        "I_3 plus (P_1 plus P_3)",
		TopAtoms:                   TopBlockDim,
		RestAtoms:                  RestBlockDim,
		TotalAtoms:                 AggregateAtomCount,
		K7Dim:                      K7Dim,
		CountMatchesK7:             true,
		TypedMapCertified:          false,
		ProjectorIdentityCertified: false,
		PromotedToK7:               false,
		Verdicts:                   []string{StatusK7SevenCountRouteAudited},
		Supports:                   []string{SupportSevenCountResonanceOnly},
		Failures:                   []string{FailureSevenAtomAggregateNotK7, FailureNoAggregateToK7Map},
	}

	dual := DualTripletBridgeAudit{
		TopTripletRole:                  "I_3 top/dominant trace participation block",
		FockTripletRole:                 "P_3 Fock/projective B-L selector triplet",
		TopTripletDim:                   TopBlockDim,
		FockTripletDim:                  FockP3Rank,
		FiniteAlgebraRouteAudited:       true,
		MoritaRouteAudited:              true,
		TraceRepresentationRouteAudited: true,
		TypedBridgeCertified:            false,
		TripletsIdentified:              false,
		Verdicts:                        []string{StatusDualTripletBridgeRouteAudited},
		Supports:                        []string{SupportDualTripletSeparation},
		Failures:                        []string{FailureColorTripletToFockTripletNotCertified, FailureFiniteAlgebraDoesNotIdentifyTriplets},
	}

	magnitude := SectorProjectorMagnitudeAudit{
		SectorProjectorsFoundAsCandidates: true,
		ProjectorsAreCarriers:             true,
		ProjectorsAreMagnitudes:           false,
		PositiveTraceAtomsDerived:         false,
		ReadoutMapCertified:               false,
		YukawaValuesDerived:               false,
		Verdicts:                          []string{StatusSectorProjectorVsMagnitudeAudited},
		Supports:                          []string{SupportTraceMagnitudeReadoutMissingObject, SupportSectorLedgerRequiresExtraReadoutMap},
		Failures:                          []string{FailureSectorProjectorsDoNotSupplyMagnitudes, FailureNoSectorMagnitudeReadout, FailureSectorProjectorNotYukawaValue, FailureNoNativeYukawaOperator},
	}

	impact := Impact{
		CanPromoteToR3:    false,
		CanPromoteToR4:    false,
		CanUpdateNEff:     false,
		CanUpdateCYukawa:  false,
		CanUpdateCHiggs:   false,
		CurrentLevel:      "R2++ aggregate trace operator, not R3 sector ledger",
		NextMissingObject: "SectorProjectorMap plus SectorTraceMagnitudeReadoutMap; alpha_B native source remains separately blocked",
		NextGate:          "Gate 833 — Finite Algebra Sector Projector / Aggregate Carrier Map Audit, or direct Sector TraceMagnitude Readout Obstruction Audit if no carrier map emerges",
		Reason:            "A_F is a lawful source of finite-internal sector projector candidates, but no typed map connects the aggregate carrier to those projectors and no readout map converts projectors into positive trace magnitudes.",
		Verdicts:          []string{StatusGate831Inherited, StatusSectorLedgerMissingObjectSharpened, StatusR2PlusPlusRetained, StatusNoLedgerUpdates},
		Supports:          []string{SupportAggregateCarrierStillR2PlusPlus, SupportSectorProjectorMapMissingObject, SupportTraceMagnitudeReadoutMissingObject},
		Failures:          []string{FailureAlphaStillSealed, FailureNoBoundaryAlphaMap, FailureNoSectorTraceLedgerMap, FailureNoSectorMagnitudeReadout, FailureAggregateOperatorNotR3, FailureR2NotR3, FailureNoNativeYukawaOperator, FailureNoNEffUpdate, FailureNoCYukawaUpdate},
	}

	firewalls := Firewalls{
		Enforced:                  true,
		AFProjectorsNotMagnitudes: true,
		AggregateNotSectorLedger:  true,
		NoCarrierMap:              true,
		NoSectorLedgerMap:         true,
		NoMagnitudeReadout:        true,
		DualTripletSeparated:      true,
		SevenNotK7:                true,
		AlphaSealed:               true,
		NoBoundaryAlphaMap:        true,
		NotR3:                     true,
		NotR4:                     true,
		NoNEffUpdate:              true,
		NoCYukawaUpdate:           true,
		NoObservedYukawaFit:       true,
		NoPMNSCKM:                 true,
		NoParticleAssignment:      true,
		Verdict:                   StatusFirewallGate832,
	}

	truth := "Gate 832 source-types A_F=C plus H plus M_3(C) as the strongest lawful finite-sector projector candidate source, but blocks promotion from R2++ aggregate carrier to R3 sector ledger.  Sector projectors, even if lawful as carriers, are not trace magnitudes; no aggregate-carrier compatibility map, no color-triplet-to-Fock-triplet bridge, no aggregate-to-K7 theorem, and no sector trace-magnitude readout are certified."
	final := "Outcome: successful source-and-carrier obstruction.  The missing object splits into SectorProjectorMap and SectorTraceMagnitudeReadoutMap, while alpha_B remains a separate sealed bridge response."

	return Analysis{Ledger: ledger, FiniteSector: finite, Compatibility: compatibility, K7Route: k7, DualTriplet: dual, Magnitude: magnitude, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}, nil
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

func FormatFiniteSector(a FiniteSectorCandidateAudit) string {
	return fmt.Sprintf("algebra=%q source=%q strongest=%t candidates=%t carriers=%t magnitudes=%t readout=%t observed_data=%t verdicts=%s supports=%s failures=%s", a.Algebra, a.CandidateProjectorSource, a.StrongestLawfulSource, a.SectorProjectorCandidatesAudited, a.SuppliesGaugeInternalCarriers, a.SuppliesYukawaMagnitudes, a.SuppliesTraceMagnitudeReadout, a.UsesObservedYukawaData, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatCompatibility(a AggregateCarrierCompatibilityAudit) string {
	return fmt.Sprintf("domain=%q target=%q map=%t compatibility=%t commutation=%t noncircular=%t compatible=%t verdicts=%s supports=%s failures=%s", a.DomainCarrier, a.TargetProjectors, a.CarrierMapCertified, a.CarrierCompatibilityCertified, a.FiniteAlgebraCommutationProven, a.NonCircular, a.Compatible, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatK7Route(a K7RouteAudit) string {
	return fmt.Sprintf("aggregate=%q atoms=%d+%d=%d K7=%d count_match=%t typed_map=%t projector_identity=%t promoted=%t verdicts=%s supports=%s failures=%s", a.AggregateExpression, a.TopAtoms, a.RestAtoms, a.TotalAtoms, a.K7Dim, a.CountMatchesK7, a.TypedMapCertified, a.ProjectorIdentityCertified, a.PromotedToK7, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatDualTriplet(a DualTripletBridgeAudit) string {
	return fmt.Sprintf("top=%q dim=%d fock=%q dim=%d finite_algebra=%t morita=%t trace_rep=%t bridge=%t identified=%t verdicts=%s supports=%s failures=%s", a.TopTripletRole, a.TopTripletDim, a.FockTripletRole, a.FockTripletDim, a.FiniteAlgebraRouteAudited, a.MoritaRouteAudited, a.TraceRepresentationRouteAudited, a.TypedBridgeCertified, a.TripletsIdentified, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatMagnitude(a SectorProjectorMagnitudeAudit) string {
	return fmt.Sprintf("sector_projector_candidates=%t carriers=%t magnitudes=%t positive_atoms=%t readout=%t yukawa_values=%t verdicts=%s supports=%s failures=%s", a.SectorProjectorsFoundAsCandidates, a.ProjectorsAreCarriers, a.ProjectorsAreMagnitudes, a.PositiveTraceAtomsDerived, a.ReadoutMapCertified, a.YukawaValuesDerived, strings.Join(a.Verdicts, ","), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("promote_R3=%t promote_R4=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t current=%q next_missing=%q next_gate=%q reason=%q verdicts=%s supports=%s failures=%s", i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CurrentLevel, i.NextMissingObject, i.NextGate, i.Reason, strings.Join(i.Verdicts, ","), strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func Statuses() []string {
	return []string{
		StatusGate831Inherited,
		StatusFiniteSectorCandidatesAudited,
		StatusAggregateCarrierCompatibilityTested,
		StatusK7SevenCountRouteAudited,
		StatusDualTripletBridgeRouteAudited,
		StatusSectorProjectorVsMagnitudeAudited,
		StatusSectorLedgerMissingObjectSharpened,
		StatusNoLedgerUpdates,
		StatusPhysicalFirewalls,
		StatusFirewallGate832,
		StatusR2PlusPlusRetained,
		StatusNoObservedDataUsed,
		SupportAFStrongestSectorProjectorSource,
		SupportFiniteSectorProjectorCandidates,
		SupportSectorLedgerRequiresExtraReadoutMap,
		SupportAggregateCarrierStillR2PlusPlus,
		SupportSectorProjectorMapMissingObject,
		SupportTraceMagnitudeReadoutMissingObject,
		SupportSevenCountResonanceOnly,
		SupportDualTripletSeparation,
		FailureNoSectorTraceLedgerMap,
		FailureSectorProjectorsDoNotSupplyMagnitudes,
		FailureAggregateNotCompatibleWithSectorProjectors,
		FailureNoTypedAggregateToSectorMap,
		FailureNoCarrierCompatibilityTheorem,
		FailureNoSectorMagnitudeReadout,
		FailureSevenAtomAggregateNotK7,
		FailureNoAggregateToK7Map,
		FailureColorTripletToFockTripletNotCertified,
		FailureFiniteAlgebraDoesNotIdentifyTriplets,
		FailureSectorProjectorNotYukawaValue,
		FailureAFSectorIdempotentNotYukawaMagnitude,
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
