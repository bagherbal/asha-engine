// Package generation2r3sectorledgerrequirementsunderboundaryalphasealaudit implements
// Gate 882: R3 SectorLedger Requirements Under BoundaryAlpha Seal Audit.
//
// Gate 882 follows Gate 881's closure of the conditional Yukawa trace-proxy
// branch. It does not reopen the alpha proof and does not attempt individual
// Yukawa values. It audits what the mature conditional trace proxy already
// supplies, what a native R3 sector trace ledger still requires, and which
// blockers must be cleared before official ledger promotion is allowed.
package generation2r3sectorledgerrequirementsunderboundaryalphasealaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE882-R3-SECTOR-LEDGER-REQUIREMENTS-UNDER-BOUNDARY-ALPHA-SEAL-AUDIT"

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	AlphaSealName         = "BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL"
	TraceProxy            = "CONDITIONAL_YUKAWA_TRACE_PROXY"
	Classification        = "R3_SECTOR_LEDGER_REQUIREMENTS_UNDER_BOUNDARY_ALPHA_SEAL"
	R2Status              = "R2+++++_R3_PREPARATION_UNDER_BOUNDARY_ALPHA_SEAL_NOT_R3"
	NextRecommendedBranch = "SECTOR_TRACE_LEDGER_MAP_AUDIT_UNDER_ALPHA_SEAL"

	BlockerBoundaryIncidenceFunctor = "BoundaryExteriorIncidenceFlagFunctor"
	BlockerSectorTraceLedgerMap     = "SectorTraceLedgerMap"
	BlockerSectorTraceMagnitudeMap  = "SectorTraceMagnitudeReadoutMap"
	BlockerGenerationCarrierMap     = "GenerationCarrierMap"
	BlockerFlavorOrientationMap     = "FlavorOrientationMap"

	StatusGate881Inherited              = "PASS_GATE881_BRANCH_CLOSURE_INHERITED"
	StatusConditionalSuppliesAudited    = "PASS_R2_PLUS_PLUS_PLUS_PLUS_PLUS_SUPPLIES_AUDITED"
	StatusR3RequirementsAudited         = "PASS_R3_REQUIREMENTS_AUDITED"
	StatusBlockersRanked                = "PASS_R3_BLOCKERS_RANKED"
	StatusUnderSealCandidateClassified  = "PASS_R3_CANDIDATE_UNDER_BOUNDARY_ALPHA_SEAL_CLASSIFIED"
	StatusSectorTraceLedgerNextSelected = "PASS_SECTOR_TRACE_LEDGER_MAP_SELECTED_AS_NEXT_BRANCH"
	StatusOfficialFreezePreserved       = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNoIndividualYukawaJump        = "PASS_NO_INDIVIDUAL_YUKAWA_JUMP_IN_GATE882"
	StatusNoAlphaProofLoop              = "PASS_NO_ALPHA_PROOF_LOOP_IN_GATE882"
	StatusFirewallVerdict               = "FIREWALL_PRESERVED_GATE882_R3_PREPARATION_NOT_NATIVE_R3"

	SupportR3PreparationUnderSeal      = "CONDITIONAL_SUPPORT_R3_PREPARATION_CAN_PROCEED_UNDER_ALPHA_SEAL"
	SupportAggregateTraceProxyInput    = "CONDITIONAL_SUPPORT_AGGREGATE_TRACE_PROXY_SUPPLIES_R3_INPUT_CANDIDATE"
	SupportYDaggerYPositiveFiniteBody  = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_READOUT_IS_POSITIVE_AND_FINITE_BODY_LOCATED"
	SupportOperatorNEffDiagnostic      = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_IS_VALID_DIAGNOSTIC_TRACE_PROXY"
	SupportPostOrientationFiniteTriple = "CONDITIONAL_SUPPORT_POST_ORIENTATION_FINITE_TRIPLE_SEAL_AVAILABLE"
	SupportBoundaryIncidencePrimary    = "CONDITIONAL_SUPPORT_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR_IS_HIGHEST_R3_BLOCKER"
	SupportSectorTraceLedgerSecond     = "CONDITIONAL_SUPPORT_SECTOR_TRACE_LEDGER_MAP_IS_SECOND_R3_BLOCKER"
	SupportNextBranchSectorLedger      = "CONDITIONAL_SUPPORT_NEXT_BRANCH_SHOULD_AUDIT_SECTOR_TRACE_LEDGER_MAP_UNDER_ALPHA_SEAL"

	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeIncidenceFunctor      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion    = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoNativeR3SectorLedger        = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureAggregateProxyNotSectorLedger = "FAILED_ROUTE_AGGREGATE_TRACE_PROXY_NOT_SECTOR_LEDGER"
	FailureNoSectorTraceLedgerMap        = "FAILED_ROUTE_NO_SECTOR_TRACE_LEDGER_MAP"
	FailureNoSectorTraceMagnitudeMap     = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoGenerationCarrierMap        = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap        = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues      = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoPhysicalYukawaSpectrum      = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                          = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type ConditionalSupplies struct {
	BoundaryAlphaSeal               bool
	PostOrientationFiniteTripleSeal bool
	SymbolicDFEdgeMatrix            bool
	YDaggerYPositiveReadout         bool
	AggregateHAgg                   bool
	OperatorNEff                    float64
	OperatorCYukawa                 float64
	OperatorCHiggs                  float64
	DiagnosticOnly                  bool
	Supports, Failures              []string
}

type R3Requirements struct {
	TypedSectorProjectors    bool
	SectorTraceAtoms         bool
	PositiveReadoutMap       bool
	SectorLedgerConsistency  bool
	NonCircularAlphaSource   bool
	GenerationFlavorFirewall bool
	NativeR3Ready            bool
	Supports, Failures       []string
}

type Blocker struct {
	Name, Description  string
	Priority           int
	Native             bool
	BlocksR3           bool
	BlocksOfficial     bool
	Supports, Failures []string
}

type Eligibility struct {
	Classification         string
	ConditionalCandidate   bool
	NativeR3               bool
	NativeR4               bool
	OfficialUpdatesAllowed bool
	NextBranch             string
	Supports, Failures     []string
}

type OfficialLedger struct {
	OperatorNEff, OfficialNEff       float64
	OperatorCYukawa, OfficialCYukawa float64
	OperatorCHiggs, OfficialCHiggs   float64
	Frozen, DiagnosticOnly           bool
	CanUpdate                        bool
	Supports, Failures               []string
}

type Firewalls struct {
	Enforced                      bool
	AlphaStillSealed              bool
	NoNativeIncidenceFunctor      bool
	AggregateProxyNotSectorLedger bool
	NoSectorTraceLedgerMap        bool
	NoGenerationFlavorTheorem     bool
	NoIndividualYukawaValues      bool
	NoOfficialLedgerUpdate        bool
	NoNativeYukawaOperator        bool
	NoR4                          bool
	Verdict                       string
}

type Audit struct {
	ID           string
	Supplies     ConditionalSupplies
	Requirements R3Requirements
	Blockers     []Blocker
	Eligibility  Eligibility
	Ledger       OfficialLedger
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func BuildDefault() (Audit, error) {
	computedCYukawa := 3.0 / OperatorNEffDiagnostic
	if !near(computedCYukawa, OperatorCYukawaDiagnostic) {
		return Audit{}, fmt.Errorf("operator C_Yukawa drift: got %.18g want %.18g", computedCYukawa, OperatorCYukawaDiagnostic)
	}
	if near(OperatorNEffDiagnostic, OfficialNEffFrozen) || near(OperatorCYukawaDiagnostic, OfficialCYukawaFrozen) || near(OperatorCHiggsDiagnostic, OfficialCHiggsFrozen) {
		return Audit{}, fmt.Errorf("operator and official ledgers collapsed")
	}

	supplies := ConditionalSupplies{
		BoundaryAlphaSeal: true, PostOrientationFiniteTripleSeal: true, SymbolicDFEdgeMatrix: true,
		YDaggerYPositiveReadout: true, AggregateHAgg: true,
		OperatorNEff: OperatorNEffDiagnostic, OperatorCYukawa: OperatorCYukawaDiagnostic, OperatorCHiggs: OperatorCHiggsDiagnostic,
		DiagnosticOnly: true,
		Supports:       []string{SupportR3PreparationUnderSeal, SupportAggregateTraceProxyInput, SupportYDaggerYPositiveFiniteBody, SupportOperatorNEffDiagnostic, SupportPostOrientationFiniteTriple},
		Failures:       []string{FailureAlphaStillSealed, FailureAggregateProxyNotSectorLedger, FailureNoNativeR3SectorLedger},
	}
	requirements := R3Requirements{
		TypedSectorProjectors: false, SectorTraceAtoms: false, PositiveReadoutMap: false, SectorLedgerConsistency: false,
		NonCircularAlphaSource: false, GenerationFlavorFirewall: true, NativeR3Ready: false,
		Supports: []string{SupportR3PreparationUnderSeal, SupportAggregateTraceProxyInput},
		Failures: []string{FailureAlphaStillSealed, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureNoSectorTraceLedgerMap, FailureNoSectorTraceMagnitudeMap, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoNativeR3SectorLedger},
	}
	blockers := []Blocker{
		{Name: BlockerBoundaryIncidenceFunctor, Priority: 1, Description: "native BoundaryExterior incidence-flag selector and cross-lane exclusion for alpha_B", Native: false, BlocksR3: true, BlocksOfficial: true, Supports: []string{SupportBoundaryIncidencePrimary}, Failures: []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed}},
		{Name: BlockerSectorTraceLedgerMap, Priority: 2, Description: "map aggregate trace-proxy input into typed sector projectors and trace atoms", Native: false, BlocksR3: true, BlocksOfficial: true, Supports: []string{SupportSectorTraceLedgerSecond}, Failures: []string{FailureNoSectorTraceLedgerMap, FailureAggregateProxyNotSectorLedger}},
		{Name: BlockerSectorTraceMagnitudeMap, Priority: 3, Description: "positive sector readout map, beyond aggregate Y^dagger Y", Native: false, BlocksR3: true, BlocksOfficial: true, Failures: []string{FailureNoSectorTraceMagnitudeMap}},
		{Name: BlockerGenerationCarrierMap, Priority: 4, Description: "generation carrier and non-physical flavor firewall before individual values", Native: false, BlocksR3: true, BlocksOfficial: false, Failures: []string{FailureNoGenerationCarrierMap}},
		{Name: BlockerFlavorOrientationMap, Priority: 5, Description: "flavor orientation and individual Yukawa splitting firewall", Native: false, BlocksR3: true, BlocksOfficial: false, Failures: []string{FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}},
	}
	eligibility := Eligibility{
		Classification: R2Status, ConditionalCandidate: true, NativeR3: false, NativeR4: false, OfficialUpdatesAllowed: false, NextBranch: NextRecommendedBranch,
		Supports: []string{SupportR3PreparationUnderSeal, SupportAggregateTraceProxyInput, SupportNextBranchSectorLedger},
		Failures: []string{FailureAlphaStillSealed, FailureNoNativeR3SectorLedger, FailureAggregateProxyNotSectorLedger, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoR4},
	}
	ledger := OfficialLedger{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{SupportOperatorNEffDiagnostic},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeR3SectorLedger},
	}
	firewalls := Firewalls{Enforced: true, AlphaStillSealed: true, NoNativeIncidenceFunctor: true, AggregateProxyNotSectorLedger: true, NoSectorTraceLedgerMap: true, NoGenerationFlavorTheorem: true, NoIndividualYukawaValues: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict}

	return Audit{ID: AuditID, Supplies: supplies, Requirements: requirements, Blockers: blockers, Eligibility: eligibility, Ledger: ledger, Firewalls: firewalls, Truth: "Gate 882 audits R3 requirements under the frozen BoundaryAlpha seal: the conditional trace proxy is a real R3 input candidate, but not a native sector trace ledger.", Final: "Next lawful branch is SectorTraceLedgerMap under the alpha seal; individual physical Yukawa values and official ledger updates remain blocked."}, nil
}

func FormatSupplies(s ConditionalSupplies) string {
	return fmt.Sprintf("supplies(alpha_seal=%t post_orientation=%t symbolic_DF=%t YdagY=%t Hagg=%t Neff=%.16g CYukawa=%.16g CHiggs=%.16g diagnostic=%t supports=%s failures=%s)", s.BoundaryAlphaSeal, s.PostOrientationFiniteTripleSeal, s.SymbolicDFEdgeMatrix, s.YDaggerYPositiveReadout, s.AggregateHAgg, s.OperatorNEff, s.OperatorCYukawa, s.OperatorCHiggs, s.DiagnosticOnly, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatRequirements(r R3Requirements) string {
	return fmt.Sprintf("requirements(sector_projectors=%t trace_atoms=%t positive_readout=%t consistency=%t noncircular_alpha=%t flavor_firewall=%t native_r3=%t supports=%s failures=%s)", r.TypedSectorProjectors, r.SectorTraceAtoms, r.PositiveReadoutMap, r.SectorLedgerConsistency, r.NonCircularAlphaSource, r.GenerationFlavorFirewall, r.NativeR3Ready, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatBlocker(b Blocker) string {
	return fmt.Sprintf("blocker(priority=%d name=%s native=%t blocks_r3=%t blocks_official=%t desc=%q supports=%s failures=%s)", b.Priority, b.Name, b.Native, b.BlocksR3, b.BlocksOfficial, b.Description, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatBlockers(bs []Blocker) string {
	parts := make([]string, 0, len(bs))
	for _, b := range bs {
		parts = append(parts, FormatBlocker(b))
	}
	return strings.Join(parts, "; ")
}

func FormatEligibility(e Eligibility) string {
	return fmt.Sprintf("eligibility(classification=%s conditional_candidate=%t native_r3=%t native_r4=%t official_updates=%t next=%s supports=%s failures=%s)", e.Classification, e.ConditionalCandidate, e.NativeR3, e.NativeR4, e.OfficialUpdatesAllowed, e.NextBranch, strings.Join(e.Supports, ","), strings.Join(e.Failures, ","))
}

func FormatLedger(l OfficialLedger) string {
	return fmt.Sprintf("ledger(operator_N_eff=%.16g official_N_eff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", l.OperatorNEff, l.OfficialNEff, l.OperatorCYukawa, l.OfficialCYukawa, l.OperatorCHiggs, l.OfficialCHiggs, l.Frozen, l.DiagnosticOnly, l.CanUpdate, strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t no_incidence=%t aggregate_not_sector=%t no_sector_map=%t no_generation_flavor=%t no_individual_yukawa=%t no_official_update=%t no_yukawa_theorem=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.NoNativeIncidenceFunctor, f.AggregateProxyNotSectorLedger, f.NoSectorTraceLedgerMap, f.NoGenerationFlavorTheorem, f.NoIndividualYukawaValues, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate881Inherited,
		StatusConditionalSuppliesAudited,
		StatusR3RequirementsAudited,
		StatusBlockersRanked,
		StatusUnderSealCandidateClassified,
		StatusSectorTraceLedgerNextSelected,
		StatusOfficialFreezePreserved,
		StatusNoIndividualYukawaJump,
		StatusNoAlphaProofLoop,
		StatusFirewallVerdict,
		SupportR3PreparationUnderSeal,
		SupportAggregateTraceProxyInput,
		SupportYDaggerYPositiveFiniteBody,
		SupportOperatorNEffDiagnostic,
		SupportPostOrientationFiniteTriple,
		SupportBoundaryIncidencePrimary,
		SupportSectorTraceLedgerSecond,
		SupportNextBranchSectorLedger,
		FailureAlphaStillSealed,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailureNoNativeR3SectorLedger,
		FailureAggregateProxyNotSectorLedger,
		FailureNoSectorTraceLedgerMap,
		FailureNoSectorTraceMagnitudeMap,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoPhysicalYukawaSpectrum,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4,
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func containsAll(have []string, want []string) bool {
	m := map[string]bool{}
	for _, h := range have {
		m[h] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.AlphaStillSealed && f.NoNativeIncidenceFunctor && f.AggregateProxyNotSectorLedger && f.NoSectorTraceLedgerMap && f.NoGenerationFlavorTheorem && f.NoIndividualYukawaValues && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
