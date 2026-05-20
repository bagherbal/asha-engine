// Package generation2covariantphasespacecliffordsourcerouterandyukawagapaccelerationaudit implements
// Gate 824: Covariant Phase-Space Clifford Source-Router and Yukawa-Gap Acceleration Audit.
//
// The gate treats the covariant phase-space Clifford idea as a source-router, not as
// an accepted theorem. Every master claim is routed through the current missing-object
// ledger: TraceMagnitudeOperatorSeal, GenerationOperatorSeal, BoundaryToTraceMagnitudeRestMap,
// BoundaryFNRestPressureMap, FlavorOrientationReadoutSeal, RealChiralityAirlock, and
// TrialityRealFormDescentMap.
package generation2covariantphasespacecliffordsourcerouterandyukawagapaccelerationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE824-COVARIANT-PHASE-SPACE-CLIFFORD-SOURCE-ROUTER-YUKAWA-GAP-ACCELERATION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate823Inherited       = "PASS_GATE823_DATA_REQUIRED_STATUS_INHERITED"
	StatusSourcePackageDefined   = "PASS_COVARIANT_PHASE_SPACE_CLIFFORD_SOURCE_PACKAGE_DEFINED"
	StatusMasterClaimsRouted     = "PASS_MASTER_CLAIM_INVENTORY_ROUTED"
	StatusRealFormPrecheck       = "PASS_REAL_FORM_CHIRALITY_PRECHECK_EXECUTED"
	StatusTrialityRouted         = "PASS_TRIALITY_SOURCE_ROUTER_AUDITED"
	StatusFockInventoryAudited   = "PASS_FOCK_PATI_SALAM_INVENTORY_AUDITED"
	StatusHiggsMassBridgeAudited = "PASS_HIGGS_MASS_BRIDGE_AUDITED"
	StatusBoundaryFNRouted       = "PASS_BOUNDARY_FN_SOURCE_ROUTER_AUDITED"
	StatusLagrangianSeparated    = "PASS_LAGRANGIAN_HIERARCHY_LANE_SEPARATED"
	StatusDarkSectorRouted       = "PASS_DARK_SECTOR_CLAIM_ROUTED"
	StatusOutputTableDefined     = "PASS_REQUIRED_OUTPUT_TABLE_DEFINED"
	StatusOutcomesDefined        = "PASS_FAST_BRANCH_OUTCOMES_DEFINED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	SupportContainerOnly        = "CONDITIONAL_SUPPORT_PHASE_SPACE_CLIFFORD_PACKAGE_MAY_SUPPLY_STATE_GAUGE_EDGE_CONTAINER"
	SupportNextDependsOnTrace   = "CONDITIONAL_SUPPORT_USEFUL_NEXT_STEP_DEPENDS_ON_WHETHER_POSITIVE_TRACE_READOUT_EXISTS"
	SupportBoundaryFNFastRoute  = "CONDITIONAL_SUPPORT_BOUNDARY_FN_MAP_REMAINS_FASTEST_ROUTE_IF_NO_H_F_SPECTRA_ARE_FOUND"
	SupportTrialityAirlock      = "CONDITIONAL_SUPPORT_TRIALITY_CAN_REMAIN_AIRLOCKED_GENERATION_SEARCH_GEOMETRY"
	SupportFockStateContainer   = "CONDITIONAL_SUPPORT_FOCK_STRUCTURE_MAY_SUPPLY_STATE_AND_GAUGE_CONTAINER"
	SupportMassBridgeEdgeOnly   = "CONDITIONAL_SUPPORT_MASS_BRIDGE_REINFORCES_EDGE_TEMPLATE_NOT_YUKAWA_VALUE_SOURCE"
	SupportBoundaryFNContainer  = "CONDITIONAL_SUPPORT_PHASE_SPACE_PACKAGE_MAY_SUPPLY_CONTAINER_FOR_BOUNDARY_TO_TRACE_MAP_IF_EXPLICIT_READOUT_EXISTS"
	SupportGaugeGravitySeparate = "CONDITIONAL_SUPPORT_GAUGE_GRAVITY_HIERARCHY_MAY_BE_DISTINCT_STRUCTURAL_BRANCH"
	SupportNuRSeparate          = "CONDITIONAL_SUPPORT_NU_R_DARK_CANDIDATE_IS_SEPARATE_BRANCH_IF_TYPED"

	FailureContainerNotYukawa          = "FAILED_ROUTE_CONTAINER_STRUCTURE_NOT_YUKAWA_TRACE_MAGNITUDE_THEOREM"
	FailureTrialityNoReadout           = "FAILED_ROUTE_TRIALITY_NOT_GENERATION_THEOREM_WITHOUT_READOUT"
	FailureTrialityNoGenerationOp      = "FAILED_ROUTE_TRIALITY_CLAIM_DOES_NOT_SOURCE_GENERATION_OPERATOR_WITHOUT_READOUT_PACKAGE"
	FailureTrialityNoNEff              = "FAILED_ROUTE_TRIALITY_DOES_NOT_UPDATE_N_EFF_WITHOUT_TRACE_MAGNITUDE_MAP"
	FailureChiralityNotYukawa          = "FAILED_ROUTE_CHIRALITY_NOT_YUKAWA_VALUE_SOURCE"
	FailureNaiveRealChirality          = "FAILED_ROUTE_NAIVE_REAL_CHIRALITY_PROJECTORS_INVALID_IF_OMEGA_SQUARED_MINUS_ONE"
	FailureChiralityAirlockNotYukawa   = "FAILED_ROUTE_CHIRALITY_AIRLOCK_NOT_YUKAWA_VALUE_THEOREM"
	FailureMassBridgeNotHierarchy      = "FAILED_ROUTE_MASS_BRIDGE_NOT_HIERARCHY_BREAKING_OPERATOR"
	FailureMassBridgeNoValues          = "FAILED_ROUTE_MASS_BRIDGE_TEMPLATE_DOES_NOT_DERIVE_YUKAWA_VALUES"
	FailureMassBridgeNoRest            = "FAILED_ROUTE_HIGGS_LEFT_RIGHT_EDGE_DOES_NOT_SOURCE_REST_PRESSURE"
	FailureStateInventoryNoTrace       = "FAILED_ROUTE_STATE_INVENTORY_NOT_YUKAWA_TRACE_MAGNITUDE_OPERATOR"
	FailureSixteenStatesNoEigenvalues  = "FAILED_ROUTE_16_FERMION_STATES_NOT_YUKAWA_EIGENVALUE_THEOREM"
	FailureBoundaryNoPositiveSpectra   = "FAILED_ROUTE_NO_BOUNDARY_TO_TRACE_MAGNITUDE_RESTMAP_WITHOUT_POSITIVE_SPECTRA"
	FailureCoeffNotTraceTheorem        = "FAILED_ROUTE_COEFFICIENT_STRUCTURE_NOT_TRACE_ATOM_THEOREM_BY_ITSELF"
	FailureLagrangianHierarchyNoNEff   = "FAILED_ROUTE_LAGRANGIAN_HIERARCHY_NOT_N_EFF_SOURCE"
	FailureForceHierarchyNotYukawa     = "FAILED_ROUTE_FORCE_HIERARCHY_NOT_YUKAWA_TRACE_MAGNITUDE_THEOREM"
	FailureAlphaOverAlphaGNoDeltaN     = "FAILED_ROUTE_ALPHA_OVER_ALPHA_G_DOES_NOT_SOURCE_N_EFF_MINUS_THREE"
	FailureDarkNoNuLedger              = "FAILED_ROUTE_DARK_SECTOR_CLAIM_DOES_NOT_SOURCE_N_EFF_WITHOUT_NEUTRINO_TRACE_LEDGER"
	FailureNoCYukawaUpdate             = "FAILED_ROUTE_GATE824_MUST_NOT_UPDATE_C_YUKAWA_UNLESS_POSITIVE_SPECTRA_OR_TRACE_MAP_ARE_CERTIFIED"
	FailureCHiggsLevelB                = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	FailureNoPositiveTraceReadoutFound = "FAILED_ROUTE_NO_POSITIVE_TRACE_MAGNITUDE_READOUT_FOUND_IN_PHASE_SPACE_SOURCE_PACKAGE"
	FailureNoBoundaryMapFound          = "FAILED_ROUTE_NO_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_SUPPLIED_BY_MASTER_PACKAGE"

	OutcomePartialContainer = "OUTCOME_B_PARTIAL_SOURCE_CONTAINER_GAUGE_STATE_EDGE_ONLY"
	NextGatePartial         = "Gate 825 — BoundaryToTraceMagnitudeRestMap Construction Candidate and Rest-Carrier Source Audit"
	StatusFirewallGate824   = "FIREWALL_PRESERVED_GATE824_COVARIANT_PHASE_SPACE_SOURCE_ROUTER_BOUNDARY"
)

type ClaimStatus string

const (
	StatusNativeTheorem      ClaimStatus = "native theorem"
	StatusCertifiedObject    ClaimStatus = "already certified ASHA object"
	StatusAirlockedCandidate ClaimStatus = "airlocked candidate"
	StatusStructural         ClaimStatus = "structural resonance"
	StatusExternalSealed     ClaimStatus = "external/sealed input"
	StatusUnsupported        ClaimStatus = "unsupported or blocked route"
)

type MissingObject string

const (
	MissingTraceMagnitude     MissingObject = "TraceMagnitudeOperatorSeal"
	MissingGenerationOperator MissingObject = "GenerationOperatorSeal"
	MissingBoundaryMap        MissingObject = "BoundaryToTraceMagnitudeRestMap"
	MissingBoundaryFNMap      MissingObject = "BoundaryFNRestPressureMap"
	MissingFlavorOrientation  MissingObject = "FlavorOrientationReadoutSeal"
	MissingRealChirality      MissingObject = "RealChiralityAirlock"
	MissingTrialityDescent    MissingObject = "TrialityRealFormDescentMap"
	MissingNeutrinoLedger     MissingObject = "NeutrinoTraceLedgerSeal"
)

type Inherited struct {
	NEff, DeltaN, S, P, M2, AlphaB, DeltaNBFN, NEffBFN   float64
	Gate823Status, LiteralSectorStatus, BoundaryFNStatus string
	Verdicts, Supports, Failures                         []string
}

type RealFormPrecheck struct {
	ClRealForm, Pseudoscalar, OmegaSquared string
	NaiveProjectorsIdempotent              bool
	RequiresAirlock                        bool
	Airlock                                string
	Verdicts, Supports, Failures           []string
}

type ClaimAudit struct {
	Claim                         string
	CurrentLayer                  string
	CertifiedStatus               ClaimStatus
	MissingObjects                []MissingObject
	RelevanceToNEff               string
	RelevanceToBoundaryFNMap      string
	RelevanceToGenerationOperator string
	RelevanceToCHiggs             string
	Verdict                       string
}

type SourceRouter struct {
	Claims                       []ClaimAudit
	Verdicts, Supports, Failures []string
}

type BoundaryFNAudit struct {
	TargetChain                  []string
	SuppliedByPackage            []string
	StillMissing                 []MissingObject
	CanMoveBeyondPartialR2       bool
	Verdicts, Supports, Failures []string
}

type FastOutcome struct {
	Outcome                      string
	NextGate                     string
	Reason                       string
	Verdicts, Supports, Failures []string
}

type Impact struct {
	OfficialNEff, OfficialCYukawa, OfficialCHiggs float64
	CanUpdate                                     bool
	Reason                                        string
	Verdicts, Supports, Failures                  []string
}

type Firewalls struct {
	Enforced                                                                            bool
	ContainerNotSpectrum, StatesNotValues, TrialityNotGeneration, ChiralityNotHierarchy bool
	MassBridgeNotValues, LabelsNotTraceAtoms, LagrangianNotNEff, DarkNotNuLedger        bool
	NoCYukawaUpdate                                                                     bool
	Verdict                                                                             string
}

type Analysis struct {
	Inherited  Inherited
	RealForm   RealFormPrecheck
	Router     SourceRouter
	BoundaryFN BoundaryFNAudit
	Outcome    FastOutcome
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func M2(s float64) float64        { return PBoundary * s * s }
func AlphaB(s float64) float64    { return (3.0/10.0)*s + M2(s) }
func DeltaNBFN(s float64) float64 { return 6.0 * AlphaB(s) }
func NEffBFN(s float64) float64   { return 3.0 + DeltaNBFN(s) }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	deltaBFN := DeltaNBFN(SBoundary)
	nEffBFN := NEffBFN(SBoundary)
	if alpha <= 0 || m2 <= 0 || nEffBFN <= 3 {
		return Analysis{}, fmt.Errorf("invalid inherited boundary-FN scales: alpha=%g M2=%g N_eff_BFN=%g", alpha, m2, nEffBFN)
	}
	inherited := Inherited{
		NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, AlphaB: alpha, DeltaNBFN: deltaBFN, NEffBFN: nEffBFN,
		Gate823Status:       "DATA_REQUIRED_EXTERNAL_LOW_SCALE_YUKAWA_RATIO_LEDGER",
		LiteralSectorStatus: "frozen; no sector decision without convention-locked ledger",
		BoundaryFNStatus:    "strengthened partial R2 aggregate concentration candidate; not R3/R4",
		Verdicts:            []string{StatusGate823Inherited},
		Supports:            []string{SupportBoundaryFNFastRoute},
		Failures:            []string{FailureNoPositiveTraceReadoutFound},
	}
	realForm := RealFormPrecheck{
		ClRealForm: "Cl(1,7) ≅ Mat(16,R)", Pseudoscalar: "omega=e0e1...e7", OmegaSquared: "-1",
		NaiveProjectorsIdempotent: false, RequiresAirlock: true, Airlock: "ComplexChiralityAirlock: gamma_chi=i omega, gamma_chi^2=1",
		Verdicts: []string{StatusRealFormPrecheck},
		Supports: []string{"CONDITIONAL_SUPPORT_CHIRALITY_PROJECTORS_REQUIRE_COMPLEX_OR_EQUIVALENT_AIRLOCK_IN_ACTIVE_CL17_BOARD"},
		Failures: []string{FailureNaiveRealChirality, FailureChiralityAirlockNotYukawa},
	}
	router := SourceRouter{
		Claims:   buildClaimAudits(),
		Verdicts: []string{StatusSourcePackageDefined, StatusMasterClaimsRouted, StatusTrialityRouted, StatusFockInventoryAudited, StatusHiggsMassBridgeAudited, StatusLagrangianSeparated, StatusDarkSectorRouted, StatusOutputTableDefined},
		Supports: []string{SupportContainerOnly, SupportNextDependsOnTrace, SupportTrialityAirlock, SupportFockStateContainer, SupportMassBridgeEdgeOnly, SupportGaugeGravitySeparate, SupportNuRSeparate},
		Failures: []string{FailureContainerNotYukawa, FailureTrialityNoReadout, FailureChiralityNotYukawa, FailureMassBridgeNotHierarchy, FailureStateInventoryNoTrace, FailureLagrangianHierarchyNoNEff, FailureDarkNoNuLedger},
	}
	boundary := BoundaryFNAudit{
		TargetChain:            []string{"s,p,5/3,color 3,boundary-pair 2", "alpha_B", "1+3 rest simplex", "positive trace atoms", "N_eff", "C_Yukawa"},
		SuppliedByPackage:      []string{"state/gauge/edge container candidates", "possible language for phase-space boundary carrier search"},
		StillMissing:           []MissingObject{MissingTraceMagnitude, MissingBoundaryMap, MissingBoundaryFNMap, MissingGenerationOperator, MissingFlavorOrientation},
		CanMoveBeyondPartialR2: false,
		Verdicts:               []string{StatusBoundaryFNRouted},
		Supports:               []string{SupportBoundaryFNContainer, SupportBoundaryFNFastRoute},
		Failures:               []string{FailureBoundaryNoPositiveSpectra, FailureCoeffNotTraceTheorem, FailureNoBoundaryMapFound},
	}
	outcome := FastOutcome{
		Outcome:  OutcomePartialContainer,
		NextGate: NextGatePartial,
		Reason:   "the covariant phase-space package routes many state/gauge/edge claims, but supplies no positive H_f spectra and no BoundaryToTraceMagnitudeRestMap",
		Verdicts: []string{StatusOutcomesDefined},
		Supports: []string{SupportContainerOnly, SupportBoundaryFNFastRoute},
		Failures: []string{FailureNoPositiveTraceReadoutFound, FailureNoBoundaryMapFound},
	}
	impact := Impact{
		OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, CanUpdate: false,
		Reason:   "no certified positive spectra or trace-magnitude rest map; C_Yukawa remains frozen",
		Verdicts: []string{"PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"},
		Failures: []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB},
	}
	firewalls := Firewalls{Enforced: true, ContainerNotSpectrum: true, StatesNotValues: true, TrialityNotGeneration: true, ChiralityNotHierarchy: true, MassBridgeNotValues: true, LabelsNotTraceAtoms: true, LagrangianNotNEff: true, DarkNotNuLedger: true, NoCYukawaUpdate: true, Verdict: StatusFirewallGate824}
	analysis := Analysis{
		Inherited: inherited, RealForm: realForm, Router: router, BoundaryFN: boundary, Outcome: outcome, Impact: impact, Firewalls: firewalls,
		Truth: "Gate 824 treats the covariant phase-space Clifford idea as a source-router: useful as state/gauge/edge container language, not as a certified Yukawa trace-magnitude source.",
		Final: "Outcome B: partial source container. No H_f spectra, BoundaryToTraceMagnitudeRestMap, GenerationOperatorSeal, or FlavorOrientationReadoutSeal is certified; C_Yukawa remains unchanged.",
	}
	return analysis, nil
}

func buildClaimAudits() []ClaimAudit {
	return []ClaimAudit{
		claim("Cl(1,7) phase-space board", "kinematic Clifford container", StatusCertifiedObject, []MissingObject{MissingTraceMagnitude, MissingBoundaryMap}, "home/container only", "may host boundary map language, does not supply it", "does not define Y_f on generation space", "no update", FailureContainerNotYukawa),
		claim("Fock 16-state inventory", "fermion state inventory", StatusStructural, []MissingObject{MissingGenerationOperator, MissingTraceMagnitude}, "no trace atoms", "not a rest map", "slots/labels only", "no update", FailureSixteenStatesNoEigenvalues),
		claim("Pati-Salam/u(4) structure", "gauge-label organization", StatusStructural, []MissingObject{MissingTraceMagnitude, MissingGenerationOperator}, "no H_f spectra", "not a boundary-to-rest trace readout", "may label states but not Y_f values", "no update", FailureStateInventoryNoTrace),
		claim("SO(8)/D4 triality", "complex/airlocked triality candidate", StatusAirlockedCandidate, []MissingObject{MissingTrialityDescent, MissingGenerationOperator, MissingTraceMagnitude}, "no N_eff readout", "not a BoundaryFN map", "three types are not generation copies", "no update", FailureTrialityNoReadout),
		claim("chirality restriction", "real-form-sensitive chirality lane", StatusAirlockedCandidate, []MissingObject{MissingRealChirality, MissingGenerationOperator}, "not trace magnitude", "not rest pressure", "not Y_f source", "no update", FailureChiralityNotYukawa),
		claim("Higgs mass bridge", "finite one-form / left-right edge template", StatusCertifiedObject, []MissingObject{MissingGenerationOperator, MissingTraceMagnitude}, "edge existence only", "not alpha/beta/q source", "requires Y_f operator", "no update", FailureMassBridgeNoValues),
		claim("finite spectral triple Yukawa edges", "sector/gauge/chirality edge skeleton", StatusCertifiedObject, []MissingObject{MissingGenerationOperator, MissingTraceMagnitude}, "trace-form template only", "color factor and trace shape, no rest map", "where not what", "no update", FailureContainerNotYukawa),
		claim("boundary-FN coefficients", "bridge-layer coefficient prior", StatusStructural, []MissingObject{MissingBoundaryMap, MissingTraceMagnitude}, "strong aggregate closure candidate", "missing positive trace readout", "not generation operator", "no update", FailureCoeffNotTraceTheorem),
		claim("1+3 rest simplex", "partial R2 aggregate concentration model", StatusStructural, []MissingObject{MissingBoundaryMap, MissingTraceMagnitude}, "positive concentration candidate", "missing trace atom map", "not generation-space operator", "no update", FailureBoundaryNoPositiveSpectra),
		claim("four-term Lagrangian", "dynamical sector inventory", StatusStructural, []MissingObject{MissingTraceMagnitude}, "not N_eff source", "not boundary-FN map", "not Y_f spectra", "no update", FailureLagrangianHierarchyNoNEff),
		claim("gauge/gravity hierarchy", "separate force-hierarchy lane", StatusStructural, []MissingObject{MissingTraceMagnitude}, "does not source Delta_N", "not rest pressure", "not Y_f", "no update", FailureForceHierarchyNotYukawa),
		claim("nu_R dark-sector candidate", "possible dark-sector branch", StatusAirlockedCandidate, []MissingObject{MissingNeutrinoLedger, MissingTraceMagnitude}, "needs neutrino trace ledger", "not rest map", "not generation operator", "no update", FailureDarkNoNuLedger),
	}
}

func claim(name, layer string, status ClaimStatus, missing []MissingObject, nEff, bfn, gen, cHiggs, verdict string) ClaimAudit {
	return ClaimAudit{Claim: name, CurrentLayer: layer, CertifiedStatus: status, MissingObjects: missing, RelevanceToNEff: nEff, RelevanceToBoundaryFNMap: bfn, RelevanceToGenerationOperator: gen, RelevanceToCHiggs: cHiggs, Verdict: verdict}
}

func FormatInherited(i Inherited) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g Delta_N_BFN=%.16g N_eff_BFN=%.16g status=%s", i.NEff, i.DeltaN, i.S, i.P, i.M2, i.AlphaB, i.DeltaNBFN, i.NEffBFN, i.BoundaryFNStatus)
}

func FormatRealForm(r RealFormPrecheck) string {
	return fmt.Sprintf("%s; %s; omega^2=%s; naive projectors idempotent=%v; requires airlock=%v; airlock=%s", r.ClRealForm, r.Pseudoscalar, r.OmegaSquared, r.NaiveProjectorsIdempotent, r.RequiresAirlock, r.Airlock)
}

func FormatClaims(claims []ClaimAudit) string {
	var b strings.Builder
	for _, c := range claims {
		missing := make([]string, 0, len(c.MissingObjects))
		for _, m := range c.MissingObjects {
			missing = append(missing, string(m))
		}
		b.WriteString(fmt.Sprintf("%s | layer=%s | status=%s | missing=%s | N_eff=%s | BoundaryFN=%s | Generation=%s | C_Higgs=%s | verdict=%s\n", c.Claim, c.CurrentLayer, c.CertifiedStatus, strings.Join(missing, ","), c.RelevanceToNEff, c.RelevanceToBoundaryFNMap, c.RelevanceToGenerationOperator, c.RelevanceToCHiggs, c.Verdict))
	}
	return strings.TrimSpace(b.String())
}

func FormatBoundaryFN(b BoundaryFNAudit) string {
	missing := make([]string, 0, len(b.StillMissing))
	for _, m := range b.StillMissing {
		missing = append(missing, string(m))
	}
	return fmt.Sprintf("target=%s supplied=%s missing=%s canMoveBeyondPartialR2=%v", strings.Join(b.TargetChain, " -> "), strings.Join(b.SuppliedByPackage, "; "), strings.Join(missing, ","), b.CanMoveBeyondPartialR2)
}

func FormatOutcome(o FastOutcome) string {
	return fmt.Sprintf("%s; next=%s; reason=%s", o.Outcome, o.NextGate, o.Reason)
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("official N_eff=%.16g C_Yukawa=%.16g C_Higgs=%.16g canUpdate=%v reason=%s", i.OfficialNEff, i.OfficialCYukawa, i.OfficialCHiggs, i.CanUpdate, i.Reason)
}

func Statuses() []string {
	return []string{
		StatusGate823Inherited, StatusSourcePackageDefined, StatusMasterClaimsRouted, StatusRealFormPrecheck, StatusTrialityRouted, StatusFockInventoryAudited, StatusHiggsMassBridgeAudited, StatusBoundaryFNRouted, StatusLagrangianSeparated, StatusDarkSectorRouted, StatusOutputTableDefined, StatusOutcomesDefined, StatusPhysicalFirewalls,
		SupportContainerOnly, SupportNextDependsOnTrace, SupportBoundaryFNFastRoute, SupportTrialityAirlock, SupportFockStateContainer, SupportMassBridgeEdgeOnly, SupportBoundaryFNContainer, SupportGaugeGravitySeparate, SupportNuRSeparate,
		FailureContainerNotYukawa, FailureTrialityNoReadout, FailureTrialityNoGenerationOp, FailureTrialityNoNEff, FailureChiralityNotYukawa, FailureNaiveRealChirality, FailureChiralityAirlockNotYukawa, FailureMassBridgeNotHierarchy, FailureMassBridgeNoValues, FailureMassBridgeNoRest, FailureStateInventoryNoTrace, FailureSixteenStatesNoEigenvalues, FailureBoundaryNoPositiveSpectra, FailureCoeffNotTraceTheorem, FailureLagrangianHierarchyNoNEff, FailureForceHierarchyNotYukawa, FailureAlphaOverAlphaGNoDeltaN, FailureDarkNoNuLedger, FailureNoCYukawaUpdate, FailureCHiggsLevelB, FailureNoPositiveTraceReadoutFound, FailureNoBoundaryMapFound, StatusFirewallGate824,
	}
}
