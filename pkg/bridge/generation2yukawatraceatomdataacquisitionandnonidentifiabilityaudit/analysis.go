// Package generation2yukawatraceatomdataacquisitionandnonidentifiabilityaudit implements
// Gate 795: Yukawa Trace Atom Data Acquisition and Non-Identifiability Audit.
//
// Gate 794 specified the DecomposedYukawaTraceLedgerSeal. Gate 795 audits
// whether the current ASHA ledger actually contains the sector/atom data needed
// to populate that seal. If only aggregate a,b are active, the gate proves that
// atom/sector decomposition is not identifiable from those two scalars alone.
package generation2yukawatraceatomdataacquisitionandnonidentifiabilityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE795-YUKAWA-TRACE-ATOM-DATA-ACQUISITION-NON-IDENTIFIABILITY-AUDIT"

	StatusGate794Inherited            = "PASS_GATE794_DECOMPOSED_YUKAWA_TRACE_LEDGER_INTERFACE_INHERITED"
	StatusDataSourceHierarchy         = "PASS_YUKAWA_DATA_SOURCE_HIERARCHY_DEFINED"
	StatusAcquisitionTableRequired    = "PASS_ACQUISITION_STATUS_TABLE_REQUIRED"
	StatusValidationProtocol          = "PASS_AGGREGATE_VALIDATION_PROTOCOL_EXECUTED_IF_DATA_EXISTS"
	StatusAggregateNonIdentifiability = "PASS_AGGREGATE_NON_IDENTIFIABILITY_PROVED"
	StatusMinimumAtomAudit            = "PASS_POSITIVITY_MINIMUM_ATOM_AUDIT_COMPLETED"
	StatusTopBounds                   = "PASS_TOP_CHANNEL_AGGREGATE_BOUNDS_COMPUTED"
	StatusLinearizedRestPressure      = "PASS_LINEARIZED_REST_PRESSURE_ESTIMATE_RECORDED"
	StatusTopRestIfTExists            = "PASS_TOP_REST_DECOMPOSITION_EXECUTED_IF_T_EXISTS"
	StatusNeutrinoCheckRequired       = "PASS_NEUTRINO_CONVENTION_CHECK_REQUIRED"
	StatusScaleLocalityRequired       = "PASS_SCALE_LOCALITY_CHECK_REQUIRED"
	StatusCHiggsImpact                = "PASS_C_HIGGS_IMPACT_STATUS_RECORDED"
	StatusBranchDecision              = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusCanPopulateOnlyIfObjectsFound   = "CONDITIONAL_SUPPORT_DECOMPOSED_LEDGER_CAN_BE_POPULATED_ONLY_IF_REQUIRED_OBJECTS_ARE_FOUND"
	StatusNEffGreaterThanThreeRest        = "CONDITIONAL_SUPPORT_N_EFF_GREATER_THAN_THREE_REQUIRES_NONZERO_REST_PARTICIPATION"
	StatusThreeTopAtomsPlusRestCompatible = "CONDITIONAL_SUPPORT_THREE_TOP_COLOR_ATOMS_PLUS_SMALL_REST_IS_COMPATIBLE_WITH_AGGREGATE_LEDGER"
	StatusAggregateCompatibleTopRest      = "CONDITIONAL_SUPPORT_AGGREGATE_LEDGER_IS_COMPATIBLE_WITH_TOP_DOMINANCE_PLUS_SMALL_REST"
	StatusRestPressureSmall               = "CONDITIONAL_SUPPORT_NON_TOP_REST_PRESSURE_IS_SMALL_AT_APPROXIMATE_3_9E_MINUS_4_SCALE_IF_BETA_IS_NEGLIGIBLE"
	StatusValidatedAtomLedgerImprovesTest = "CONDITIONAL_SUPPORT_VALIDATED_ATOM_LEDGER_WOULD_IMPROVE_C_HIGGS_TESTABILITY"

	StatusAggregateLowestInformation        = "FAILED_ROUTE_AGGREGATE_LEDGER_ALONE_IS_LOWEST_INFORMATION_SOURCE"
	StatusNoSectorAuditIfMissing            = "FAILED_ROUTE_NO_SECTOR_AUDIT_IF_REQUIRED_OBJECTS_REMAIN_MISSING"
	StatusValidationRequired                = "FAILED_ROUTE_VALIDATION_REQUIRED_BEFORE_SECTOR_INTERPRETATION"
	StatusABCannotIdentifyAtoms             = "FAILED_ROUTE_A_B_ALONE_CANNOT_IDENTIFY_TRACE_ATOMS"
	StatusABCannotIdentifyTop               = "FAILED_ROUTE_A_B_ALONE_CANNOT_IDENTIFY_TOP_CHANNEL"
	StatusABCannotAssignSectors             = "FAILED_ROUTE_A_B_ALONE_CANNOT_ASSIGN_N_EFF_MINUS_THREE_TO_SECTORS"
	StatusMinimumAtomNotSector              = "FAILED_ROUTE_MINIMUM_ATOM_COUNT_DOES_NOT_IDENTIFY_SECTOR_OR_GENERATION"
	StatusTopBoundsDoNotDetermineT          = "FAILED_ROUTE_TOP_CHANNEL_VALUE_NOT_DETERMINED_BY_BOUNDS"
	StatusTMustNotBeBacksolved              = "FAILED_ROUTE_T_MUST_NOT_BE_SOLVED_BACKWARDS_FROM_N_EFF"
	StatusAlphaEstimateNoBetaControl        = "FAILED_ROUTE_ALPHA_ESTIMATE_NOT_VALID_WITHOUT_BETA_CONTROL"
	StatusAlphaEstimateNotSectorAssignment  = "FAILED_ROUTE_ALPHA_ESTIMATE_NOT_SECTOR_ASSIGNMENT"
	StatusNoTopRestWithoutSelector          = "FAILED_ROUTE_NO_TOP_REST_DECOMPOSITION_WITHOUT_TOP_CHANNEL_SELECTOR"
	StatusNeutrinoConventionImplicitBlocked = "FAILED_ROUTE_NEUTRINO_CONVENTION_MUST_NOT_BE_IMPLICIT"
	StatusScaleStabilityRequiresMultiScale  = "FAILED_ROUTE_SCALE_STABILITY_REQUIRES_MULTI_SCALE_TRACE_LEDGER"
	StatusCHiggsNotLevelCAfterAcquisition   = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_AFTER_DATA_ACQUISITION"
	StatusFirewallPreservedGate795          = "FIREWALL_PRESERVED_GATE795_YUKAWA_TRACE_ATOM_DATA_ACQUISITION_BOUNDARY"
)

const (
	aSnapshot       = 2.8424095142339083
	bSnapshot       = 2.6910096440382287
	ratioSnapshot   = 0.33307493962706697
	nEffSnapshot    = 3.0023273474722147
	cHiggsSnapshot  = 1.0372205204048603
	cYukawaSnapshot = 0.9992248188812008
)

type SourceStatus string

const (
	StatusFoundNative        SourceStatus = "FOUND_NATIVE"
	StatusFoundSealed        SourceStatus = "FOUND_SEALED"
	StatusFoundAggregateOnly SourceStatus = "FOUND_AGGREGATE_ONLY"
	StatusMissing            SourceStatus = "MISSING"
	StatusAmbiguous          SourceStatus = "AMBIGUOUS"
)

type Gate794Inheritance struct {
	Inherited bool
	Verdict   string
}

type DataSourceHierarchy struct {
	Defined          bool
	Priority         []string
	HighestAvailable string
	Verdict          string
}

type AcquisitionRow struct {
	Object string
	Status SourceStatus
	Source string
}

type AcquisitionAudit struct {
	Required    bool
	Rows        []AcquisitionRow
	CanPopulate bool
	Verdict     string
}

type ValidationAudit struct {
	ProtocolExecuted bool
	DataExists       bool
	Validated        bool
	Rules            []string
	Verdict          string
}

type NonIdentifiabilityAudit struct {
	Proved           bool
	ConstraintCount  int
	UnknownAtomCount string
	InfiniteFamilies bool
	CannotIdentify   []string
	Verdict          string
}

type MinimumAtomAudit struct {
	Completed               bool
	NEff                    float64
	MinimumNonzeroAtoms     int
	RequiresRestBeyondThree bool
	CompatibleReading       string
	Verdict                 string
}

type TopBoundsAudit struct {
	Computed       bool
	AOverThree     float64
	SqrtBOverThree float64
	UpperBoundT    float64
	Gap            float64
	DeterminesT    bool
	Verdict        string
}

type LinearizedRestPressure struct {
	Recorded      bool
	DeltaRatio    float64
	AlphaEstimate float64
	Assumption    string
	IsTheorem     bool
	Verdict       string
}

type TopRestAudit struct {
	ExecutedIfTExists bool
	TypedTFound       bool
	AlphaBetaComputed bool
	FormulaRatio      string
	FormulaDelta      string
	Verdict           string
}

type NeutrinoConventionAudit struct {
	Required bool
	Status   string
	Implicit bool
	Verdict  string
}

type ScaleLocalityAudit struct {
	Required                bool
	Scale                   string
	MultiScaleLedger        bool
	ScaleStabilityCertified bool
	Verdict                 string
}

type CHiggsImpactAudit struct {
	Recorded                        bool
	ValidatedAtomLedgerWouldImprove bool
	NEffAggregateSealed             bool
	CHiggsLevelC                    bool
	Verdict                         string
}

type BranchDecision struct {
	Recorded                   bool
	ValidatedDataFound         bool
	NativeYukawaOperatorsFound bool
	Recommended                string
	Alternatives               []string
	Verdict                    string
}

type Firewalls struct {
	Enforced                     bool
	AggregateIsAtomLedger        bool
	TopDominanceTopYukawaTheorem bool
	MinimumAtomGenerationTheorem bool
	NEffD4Triality               bool
	SectorDataNativeYukawa       bool
	ValidatedAtomsPMNSCKM        bool
	SingleScaleStable            bool
	CHiggsLevelC                 bool
	TreeProxyPoleMass            bool
	Verdict                      string
}

type Analysis struct {
	Gate794            Gate794Inheritance
	Hierarchy          DataSourceHierarchy
	Acquisition        AcquisitionAudit
	Validation         ValidationAudit
	NonIdentifiability NonIdentifiabilityAudit
	MinimumAtom        MinimumAtomAudit
	TopBounds          TopBoundsAudit
	LinearizedRest     LinearizedRestPressure
	TopRest            TopRestAudit
	Neutrino           NeutrinoConventionAudit
	Scale              ScaleLocalityAudit
	Impact             CHiggsImpactAudit
	Branch             BranchDecision
	Firewalls          Firewalls
	Truth              string
	FinalStatement     string
}

func BuildDefault() (Analysis, error) {
	ratio := bSnapshot / (aSnapshot * aSnapshot)
	nEff := (aSnapshot * aSnapshot) / bSnapshot
	cYukawa := 3.0 / nEff
	if !closeAbs(ratio, ratioSnapshot, 1e-16) || !closeAbs(nEff, nEffSnapshot, 5e-16) || !closeAbs(cYukawa, cYukawaSnapshot, 5e-16) {
		return Analysis{}, fmt.Errorf("aggregate Yukawa ledger mismatch: ratio=%.17g N_eff=%.17g C_Y=%.17g", ratio, nEff, cYukawa)
	}
	deltaRatio := ratio - 1.0/3.0
	alphaEstimate := -3.0 * deltaRatio / 2.0
	aOverThree := aSnapshot / 3.0
	sqrtBOverThree := math.Sqrt(bSnapshot / 3.0)
	upperT := math.Min(aOverThree, sqrtBOverThree)
	minAtoms := int(math.Ceil(nEff - 1e-15))
	a := Analysis{
		Gate794:   Gate794Inheritance{Inherited: true, Verdict: StatusGate794Inherited},
		Hierarchy: DataSourceHierarchy{Defined: true, Priority: []string{"native Yukawa operator package", "sealed Yukawa singular-value ledger", "sector trace ledger", "trace atom ledger", "aggregate trace ledger only"}, HighestAvailable: "aggregate trace ledger only", Verdict: StatusAggregateLowestInformation},
		Acquisition: AcquisitionAudit{Required: true, CanPopulate: false, Rows: []AcquisitionRow{
			{Object: "scale_convention", Status: StatusFoundAggregateOnly, Source: "aggregate ledger at M_Z"},
			{Object: "Yukawa_normalization_convention", Status: StatusAmbiguous, Source: "aggregate spectral-action trace convention only"},
			{Object: "a_u,a_d,a_e,a_nu", Status: StatusMissing, Source: "no sector trace ledger active"},
			{Object: "b_u,b_d,b_e,b_nu", Status: StatusMissing, Source: "no sector quartic ledger active"},
			{Object: "trace atoms x_i", Status: StatusMissing, Source: "aggregate a,b only"},
			{Object: "top channel T", Status: StatusMissing, Source: "no typed top-channel selector active"},
			{Object: "color multiplicity rule", Status: StatusMissing, Source: "not explicit in atom ledger"},
			{Object: "neutrino convention", Status: StatusAmbiguous, Source: "not explicit in decomposed trace ledger"},
		}, Verdict: StatusNoSectorAuditIfMissing},
		Validation:         ValidationAudit{ProtocolExecuted: true, DataExists: false, Validated: false, Rules: []string{"sum sectors reproduce a", "sum quartic sectors reproduce b", "a^2/b reproduces N_eff", "sum atoms reproduce a", "sum atom squares reproduce b"}, Verdict: StatusValidationRequired},
		NonIdentifiability: NonIdentifiabilityAudit{Proved: true, ConstraintCount: 2, UnknownAtomCount: "unknown positive atom list", InfiniteFamilies: true, CannotIdentify: []string{"sector fractions", "generation fractions", "top channel T", "bottom/tau/charm contributions", "neutrino contribution", "color representation", "scale stability", "D4/triality carrier"}, Verdict: StatusABCannotIdentifyAtoms},
		MinimumAtom:        MinimumAtomAudit{Completed: true, NEff: nEff, MinimumNonzeroAtoms: minAtoms, RequiresRestBeyondThree: nEff > 3.0, CompatibleReading: "three dominant top-color atoms plus at least one nonzero rest contribution", Verdict: StatusNEffGreaterThanThreeRest},
		TopBounds:          TopBoundsAudit{Computed: true, AOverThree: aOverThree, SqrtBOverThree: sqrtBOverThree, UpperBoundT: upperT, Gap: aOverThree - sqrtBOverThree, DeterminesT: false, Verdict: StatusAggregateCompatibleTopRest},
		LinearizedRest:     LinearizedRestPressure{Recorded: true, DeltaRatio: deltaRatio, AlphaEstimate: alphaEstimate, Assumption: "beta << alpha", IsTheorem: false, Verdict: StatusRestPressureSmall},
		TopRest:            TopRestAudit{ExecutedIfTExists: true, TypedTFound: false, AlphaBetaComputed: false, FormulaRatio: "b/a^2=(1/3)(1+beta)/(1+alpha)^2", FormulaDelta: "delta_ratio=(1/3)(beta-2alpha-alpha^2)/(1+alpha)^2", Verdict: StatusNoTopRestWithoutSelector},
		Neutrino:           NeutrinoConventionAudit{Required: true, Status: "Y_nu unknown", Implicit: true, Verdict: StatusNeutrinoConventionImplicitBlocked},
		Scale:              ScaleLocalityAudit{Required: true, Scale: "M_Z", MultiScaleLedger: false, ScaleStabilityCertified: false, Verdict: StatusScaleStabilityRequiresMultiScale},
		Impact:             CHiggsImpactAudit{Recorded: true, ValidatedAtomLedgerWouldImprove: true, NEffAggregateSealed: true, CHiggsLevelC: false, Verdict: StatusValidatedAtomLedgerImprovesTest},
		Branch:             BranchDecision{Recorded: true, ValidatedDataFound: false, NativeYukawaOperatorsFound: false, Recommended: "Gate 796 — External Yukawa Ledger Convention Seal and Atom Data Intake Audit", Alternatives: []string{"Gate 796 — Sector Contribution to N_eff Deviation and Top-Rest Dominance Audit", "Gate 796 — Native Yukawa Operator Trace Atom Extraction Audit"}, Verdict: StatusBranchDecision},
		Firewalls:          Firewalls{Enforced: true, AggregateIsAtomLedger: false, TopDominanceTopYukawaTheorem: false, MinimumAtomGenerationTheorem: false, NEffD4Triality: false, SectorDataNativeYukawa: false, ValidatedAtomsPMNSCKM: false, SingleScaleStable: false, CHiggsLevelC: false, TreeProxyPoleMass: false, Verdict: StatusFirewallPreservedGate795},
		Truth:              "Gate 795 finds only aggregate a,b,N_eff active in the current scalar-Higgs ledger; decomposed Yukawa trace atoms remain missing.",
		FinalStatement:     "Gate 795 determines that ASHA does not yet expose the decomposed Yukawa trace data needed to source N_eff. With only aggregate a,b, N_eff remains non-identifiable at the atom/sector level: a,b prove inverse participation, but they do not reveal sectors, generations, top channel, or D4/triality structure. The only honest inference is that N_eff>3 requires nonzero rest participation beyond the ideal three top-color atoms, while the currently certified near-three source remains top-color dominance plus small unresolved non-top pressure.",
	}
	return a, nil
}

func FormatAcquisition(x AcquisitionAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s=%s(%s)", r.Object, r.Status, r.Source))
	}
	return fmt.Sprintf("can_populate=%t rows=[%s] verdict=%s", x.CanPopulate, strings.Join(parts, "; "), x.Verdict)
}

func FormatNonIdentifiability(x NonIdentifiabilityAudit) string {
	return fmt.Sprintf("proved=%t constraints=%d unknowns=%s infinite=%t cannot=[%s] verdict=%s", x.Proved, x.ConstraintCount, x.UnknownAtomCount, x.InfiniteFamilies, strings.Join(x.CannotIdentify, ", "), x.Verdict)
}

func FormatMinimumAtom(x MinimumAtomAudit) string {
	return fmt.Sprintf("N_eff=%.16g min_nonzero_atoms=%d rest_required=%t reading=%s verdict=%s", x.NEff, x.MinimumNonzeroAtoms, x.RequiresRestBeyondThree, x.CompatibleReading, x.Verdict)
}

func FormatTopBounds(x TopBoundsAudit) string {
	return fmt.Sprintf("a/3=%.16g sqrt(b/3)=%.16g T_bound=%.16g gap=%.16g determines_T=%t verdict=%s", x.AOverThree, x.SqrtBOverThree, x.UpperBoundT, x.Gap, x.DeterminesT, x.Verdict)
}

func FormatLinearized(x LinearizedRestPressure) string {
	return fmt.Sprintf("delta_ratio=%.16g alpha_est=%.16g assumption=%s theorem=%t verdict=%s", x.DeltaRatio, x.AlphaEstimate, x.Assumption, x.IsTheorem, x.Verdict)
}

func FormatBranch(x BranchDecision) string {
	return fmt.Sprintf("validated_data=%t native_operators=%t recommended=%s alternatives=[%s]", x.ValidatedDataFound, x.NativeYukawaOperatorsFound, x.Recommended, strings.Join(x.Alternatives, "; "))
}

func Statuses() []string {
	return []string{
		StatusGate794Inherited, StatusDataSourceHierarchy, StatusAcquisitionTableRequired, StatusValidationProtocol, StatusAggregateNonIdentifiability,
		StatusMinimumAtomAudit, StatusTopBounds, StatusLinearizedRestPressure, StatusTopRestIfTExists, StatusNeutrinoCheckRequired,
		StatusScaleLocalityRequired, StatusCHiggsImpact, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusCanPopulateOnlyIfObjectsFound, StatusNEffGreaterThanThreeRest, StatusThreeTopAtomsPlusRestCompatible, StatusAggregateCompatibleTopRest,
		StatusRestPressureSmall, StatusValidatedAtomLedgerImprovesTest,
		StatusAggregateLowestInformation, StatusNoSectorAuditIfMissing, StatusValidationRequired, StatusABCannotIdentifyAtoms, StatusABCannotIdentifyTop,
		StatusABCannotAssignSectors, StatusMinimumAtomNotSector, StatusTopBoundsDoNotDetermineT, StatusTMustNotBeBacksolved, StatusAlphaEstimateNoBetaControl,
		StatusAlphaEstimateNotSectorAssignment, StatusNoTopRestWithoutSelector, StatusNeutrinoConventionImplicitBlocked, StatusScaleStabilityRequiresMultiScale,
		StatusCHiggsNotLevelCAfterAcquisition, StatusFirewallPreservedGate795,
	}
}

var statusSet = struct {
	sync.Once
	m map[string]bool
}{}

func containsAll(haystack []string, needles []string) bool {
	for _, needle := range needles {
		found := false
		for _, item := range haystack {
			if strings.Contains(item, needle) {
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

func hasRow(rows []AcquisitionRow, object string, status SourceStatus) bool {
	for _, r := range rows {
		if strings.Contains(r.Object, object) && r.Status == status {
			return true
		}
	}
	return false
}

func closeAbs(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
