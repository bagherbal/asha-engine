// Package generation2spectralcutoffrenormalizationairlock implements Gate 514:
// Spectral Cutoff and Renormalization Airlock Comparator.
//
// Gate 513 proved a native stripped heat-kernel hierarchy after f4Λ4, f2Λ2,
// and f0 were factored away. Gate 514 deliberately does not try to compute a
// physical cutoff, Newton constant, or cosmological constant. It defines the
// fail-closed bridge schema required before any continuum comparator can be
// attached to the spectral-action normalization channels.
package generation2spectralcutoffrenormalizationairlock

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2spectralmomenthierarchyairlock"
)

const (
	AuditID = "GATE514-SPECTRAL-CUTOFF-RENORMALIZATION-AIRLOCK-COMPARATOR"

	StatusGate513Inherited                    = "CONDITIONAL_SUPPORT_GATE513_SPECTRAL_MOMENT_HIERARCHY_INHERITED"
	StatusCutoffRenormalizationSchemaDefined  = "CONDITIONAL_SUPPORT_SPECTRAL_CUTOFF_RENORMALIZATION_SCHEMA_DEFINED"
	StatusRedactedBridgeSchemaAccepted        = "CONDITIONAL_SUPPORT_REDACTED_BRIDGE_COMPARATOR_SCHEMA_ACCEPTED"
	StatusRequiredRowsEnumerated              = "CONDITIONAL_SUPPORT_REQUIRED_CUTOFF_MOMENT_RENORMALIZATION_ROWS_ENUMERATED"
	StatusMomentProductRowsQuarantined        = "CONDITIONAL_SUPPORT_F2LAMBDA2_AND_F4LAMBDA4_ROWS_QUARANTINED"
	StatusVacuumSubtractionPolicyFailClosed   = "CONDITIONAL_SUPPORT_VACUUM_SUBTRACTION_POLICY_FAIL_CLOSED"
	StatusPlanckNewtonCosmologyRowsBridgeOnly = "CONDITIONAL_SUPPORT_PLANCK_NEWTON_COSMOLOGY_ROWS_BRIDGE_ONLY"
	StatusComparatorRejectsNativePromotion    = "CONDITIONAL_SUPPORT_COMPARATOR_REJECTS_NATIVE_PROMOTION"
	StatusNoNumericalAdapterExecuted          = "CONDITIONAL_SUPPORT_NO_NUMERICAL_GRAVITY_COSMOLOGY_ADAPTER_EXECUTED"

	StatusFailedLambdaNativeSelectionRejected       = "FAILED_ROUTE_LAMBDA_CUTOFF_NATIVE_SELECTION_REJECTED"
	StatusFailedF2F4NativeSelectionRejected         = "FAILED_ROUTE_F2_F4_PROFILE_MOMENTS_NATIVE_SELECTION_REJECTED"
	StatusFailedMomentProductsNativePromotion       = "FAILED_ROUTE_F2LAMBDA2_F4LAMBDA4_NATIVE_PROMOTION_REJECTED"
	StatusFailedPlanckMatchingNativePromotion       = "FAILED_ROUTE_PLANCK_MATCHING_NATIVE_PROMOTION_REJECTED"
	StatusFailedVacuumSubtractionNativeSelection    = "FAILED_ROUTE_VACUUM_SUBTRACTION_NATIVE_SELECTION_REJECTED"
	StatusFailedNewtonAndCosmoStillNotDerived       = "FAILED_ROUTE_NEWTON_AND_COSMOLOGICAL_CONSTANT_STILL_NOT_DERIVED"
	StatusFailedObservedComparatorNotImported       = "FAILED_ROUTE_OBSERVED_GRAVITY_COSMOLOGY_COMPARATOR_NOT_IMPORTED"
	StatusFirewallNoEmpiricalGravityCosmologyImport = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_OR_DARK_ENERGY_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked                = "FIREWALL_BLOCKED_CUTOFF_RENORMALIZATION_NATIVE_WRITE"
)

type Inheritance struct {
	Executed                          bool
	Gate513Inherited                  bool
	StrippedHierarchyNative           bool
	A2OverA0Ratio                     float64
	A4OverA0Ratio                     float64
	A4OverA2Ratio                     float64
	Gate513F2Selected                 bool
	Gate513F4Selected                 bool
	Gate513CutoffSelected             bool
	Gate513NewtonDerived              bool
	Gate513CosmologicalDerived        bool
	Gate513NativeNormalizationBlocked bool
	Verdict                           string
	Reason                            string
}

type ComparatorRow struct {
	Key                     string
	Symbol                  string
	Role                    string
	Required                bool
	Dimension               string
	Value                   *float64
	Source                  string
	Scale                   string
	Scheme                  string
	Uncertainty             string
	BridgeOnly              bool
	EmpiricalImport         bool
	NativePromotionRejected bool
}

type Schema struct {
	Executed                  bool
	Rows                      []ComparatorRow
	RequiredRowCount          int
	AcceptedRedactedRows      int
	NumericalRows             int
	EmpiricalRows             int
	RowsBridgeOnly            bool
	RowsRejectNativePromotion bool
	AllMetadataComplete       bool
	Verdict                   string
	Reason                    string
}

type PreflightCase struct {
	Name     string
	Rows     []ComparatorRow
	Execute  bool
	Accepted bool
	Reason   string
}

type Preflight struct {
	Executed                     bool
	Cases                        []PreflightCase
	AcceptedCases                int
	RejectedCases                int
	RejectedNativePromotionCases int
	RejectedNumericalCases       int
	RejectedMissingMetadataCases int
	RejectedExecutionCases       int
	Verdict                      string
	Reason                       string
}

type Airlock struct {
	Executed                        bool
	LambdaCutoffSelected            bool
	F2MomentSelected                bool
	F4MomentSelected                bool
	F2LambdaProductSeparated        bool
	F4LambdaProductSeparated        bool
	PlanckMatchingNative            bool
	NewtonConstantDerived           bool
	CosmologicalConstantDerived     bool
	VacuumSubtractionSelectedNative bool
	RenormalizationSchemeNative     bool
	ObservedComparatorImported      bool
	NumericalAdapterExecuted        bool
	NativeNormalizationWrite        bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                         bool
	NewtonConstantImported           bool
	PlanckScaleImported              bool
	CutoffLambdaImported             bool
	F2MomentImported                 bool
	F4MomentImported                 bool
	F2LambdaProductImported          bool
	F4LambdaProductImported          bool
	CosmologicalConstantImported     bool
	DarkEnergyImported               bool
	VacuumSubtractionImported        bool
	ObservedComparatorImported       bool
	NativeCutoffRenormalizationWrite bool
	Verdict                          string
	Reason                           string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      Schema
	Preflight   Preflight
	Airlock     Airlock
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g513, err := generation2spectralmomenthierarchyairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate513 spectral moment hierarchy: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g513)
	a.Schema = buildSchema()
	a.Preflight = buildPreflight(a.Schema.Rows)
	a.Airlock = buildAirlock()
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g513 generation2spectralmomenthierarchyairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate513Inherited:                  g513.Hierarchy.DimensionlessCombinatoric && g513.Airlock.NativeNormalizationWrite == false,
		StrippedHierarchyNative:           g513.Hierarchy.DimensionlessCombinatoric,
		A2OverA0Ratio:                     g513.Hierarchy.A2OverA0AfterFactoring,
		A4OverA0Ratio:                     g513.Hierarchy.A4OverA0AfterFactoring,
		A4OverA2Ratio:                     g513.Hierarchy.A4OverA2AfterFactoring,
		Gate513F2Selected:                 g513.Airlock.F2MomentSelected,
		Gate513F4Selected:                 g513.Airlock.F4MomentSelected,
		Gate513CutoffSelected:             g513.Airlock.CutoffLambdaSelected,
		Gate513NewtonDerived:              g513.Airlock.NewtonConstantDerived,
		Gate513CosmologicalDerived:        g513.Airlock.CosmologicalConstantDerived,
		Gate513NativeNormalizationBlocked: !g513.Airlock.NativeNormalizationWrite,
		Verdict:                           StatusGate513Inherited,
		Reason:                            "Gate514 inherits Gate513's stripped native heat-kernel prefactor hierarchy, but also inherits the failure to select f2, f4, Λ, Newton normalization, cosmological normalization, or vacuum subtraction.",
	}
}

func buildSchema() Schema {
	rows := requiredRows()
	req, redacted, numeric, empirical := 0, 0, 0, 0
	bridgeOnly, rejectNative, meta := true, true, true
	for _, r := range rows {
		if r.Required {
			req++
		}
		if r.Value == nil {
			redacted++
		} else {
			numeric++
		}
		if r.EmpiricalImport {
			empirical++
		}
		bridgeOnly = bridgeOnly && r.BridgeOnly
		rejectNative = rejectNative && r.NativePromotionRejected
		meta = meta && r.Source != "" && r.Scale != "" && r.Scheme != "" && r.Uncertainty != "" && r.Dimension != ""
	}
	return Schema{
		Executed: true, Rows: rows, RequiredRowCount: req, AcceptedRedactedRows: redacted, NumericalRows: numeric, EmpiricalRows: empirical,
		RowsBridgeOnly: bridgeOnly, RowsRejectNativePromotion: rejectNative, AllMetadataComplete: meta,
		Verdict: strings.Join([]string{StatusCutoffRenormalizationSchemaDefined, StatusRequiredRowsEnumerated, StatusRedactedBridgeSchemaAccepted, StatusMomentProductRowsQuarantined, StatusPlanckNewtonCosmologyRowsBridgeOnly}, ";"),
		Reason:  "the comparator schema enumerates the minimum external bridge rows needed for gravity/cosmology normalization while keeping every value redacted, bridge-only, and rejected for native promotion.",
	}
}

func requiredRows() []ComparatorRow {
	return []ComparatorRow{
		redacted("cutoff_lambda", "Λ", "spectral cutoff scale; required for a0/a2 physical normalization", "mass"),
		redacted("spectral_moment_f2", "f₂", "profile moment multiplying the Einstein-Hilbert a2 channel", "dimensionless moment"),
		redacted("spectral_moment_f4", "f₄", "profile moment multiplying the cosmological a0 channel", "dimensionless moment"),
		redacted("spectral_moment_f0", "f₀", "profile/contact moment for a4; may be symbolic but not a low-energy gravity normalization", "dimensionless moment"),
		redacted("f2_lambda_squared", "f₂Λ²", "bridge product needed before Newton normalization can be compared", "mass²"),
		redacted("f4_lambda_fourth", "f₄Λ⁴", "bridge product needed before cosmological-volume normalization can be compared", "mass⁴"),
		redacted("planck_normalization", "M_P or G", "external convention for mapping a2 coefficient to Newton normalization", "mass or inverse mass²"),
		redacted("cosmological_constant", "Λ_cosmo", "external comparator for vacuum/cosmological normalization", "mass⁴ or curvature"),
		redacted("vacuum_subtraction", "δρ_vac", "renormalization/subtraction prescription for the volume channel", "scheme"),
		redacted("renormalization_scheme", "scheme", "scale, regulator, subtraction, and boundary convention metadata", "metadata"),
	}
}

func redacted(key, symbol, role, dimension string) ComparatorRow {
	return ComparatorRow{Key: key, Symbol: symbol, Role: role, Required: true, Dimension: dimension, Value: nil, Source: "redacted-gate514-preflight", Scale: "redacted", Scheme: "redacted", Uncertainty: "redacted", BridgeOnly: true, EmpiricalImport: false, NativePromotionRejected: true}
}

func buildPreflight(validRows []ComparatorRow) Preflight {
	cases := []PreflightCase{
		{Name: "accepted redacted bridge schema", Rows: validRows, Accepted: true, Reason: "all required rows are present, redacted, bridge-only, metadata-complete, and reject native promotion"},
		{Name: "missing cutoff lambda", Rows: without(validRows, "cutoff_lambda"), Accepted: false, Reason: "Λ row is required"},
		{Name: "numerical lambda by default", Rows: withNumber(validRows, "cutoff_lambda", 1), Accepted: false, Reason: "numerical cutoff values are forbidden in the default native audit"},
		{Name: "f2 native promotion attempted", Rows: withNativePromotion(validRows, "spectral_moment_f2"), Accepted: false, Reason: "f2 profile moment cannot be promoted to native data"},
		{Name: "f4 native promotion attempted", Rows: withNativePromotion(validRows, "spectral_moment_f4"), Accepted: false, Reason: "f4 profile moment cannot be promoted to native data"},
		{Name: "Planck matching native promotion attempted", Rows: withNativePromotion(validRows, "planck_normalization"), Accepted: false, Reason: "Planck/Newton matching is bridge/environmental"},
		{Name: "vacuum subtraction native promotion attempted", Rows: withNativePromotion(validRows, "vacuum_subtraction"), Accepted: false, Reason: "vacuum subtraction is not selected by the finite trace"},
		{Name: "missing scheme metadata", Rows: withMissingScheme(validRows, "renormalization_scheme"), Accepted: false, Reason: "renormalization metadata must be explicit"},
		{Name: "numerical adapter execution requested", Rows: validRows, Execute: true, Accepted: false, Reason: "Gate514 is a preflight airlock and executes no numerical gravity/cosmology adapter"},
	}
	accepted, rejected, nativeRejects, numericalRejects, metadataRejects, executionRejects := 0, 0, 0, 0, 0, 0
	for _, c := range cases {
		if c.Accepted {
			accepted++
			continue
		}
		rejected++
		if strings.Contains(c.Reason, "promot") {
			nativeRejects++
		}
		if strings.Contains(c.Reason, "numerical") || strings.Contains(c.Reason, "values") {
			numericalRejects++
		}
		if strings.Contains(c.Reason, "metadata") {
			metadataRejects++
		}
		if c.Execute {
			executionRejects++
		}
	}
	return Preflight{Executed: true, Cases: cases, AcceptedCases: accepted, RejectedCases: rejected, RejectedNativePromotionCases: nativeRejects, RejectedNumericalCases: numericalRejects, RejectedMissingMetadataCases: metadataRejects, RejectedExecutionCases: executionRejects, Verdict: strings.Join([]string{StatusComparatorRejectsNativePromotion, StatusVacuumSubtractionPolicyFailClosed, StatusNoNumericalAdapterExecuted}, ";"), Reason: "Gate514 preflight accepts only the redacted bridge schema and rejects missing rows, numerical defaults, native-promotion attempts, incomplete metadata, and adapter execution."}
}

func without(rows []ComparatorRow, key string) []ComparatorRow {
	out := make([]ComparatorRow, 0, len(rows))
	for _, r := range rows {
		if r.Key != key {
			out = append(out, r)
		}
	}
	return out
}
func withNumber(rows []ComparatorRow, key string, value float64) []ComparatorRow {
	out := cloneRows(rows)
	for i := range out {
		if out[i].Key == key {
			out[i].Value = &value
			out[i].EmpiricalImport = true
		}
	}
	return out
}
func withNativePromotion(rows []ComparatorRow, key string) []ComparatorRow {
	out := cloneRows(rows)
	for i := range out {
		if out[i].Key == key {
			out[i].NativePromotionRejected = false
			out[i].BridgeOnly = false
		}
	}
	return out
}
func withMissingScheme(rows []ComparatorRow, key string) []ComparatorRow {
	out := cloneRows(rows)
	for i := range out {
		if out[i].Key == key {
			out[i].Scheme = ""
		}
	}
	return out
}
func cloneRows(rows []ComparatorRow) []ComparatorRow {
	out := make([]ComparatorRow, len(rows))
	copy(out, rows)
	return out
}

func buildAirlock() Airlock {
	return Airlock{Executed: true, LambdaCutoffSelected: false, F2MomentSelected: false, F4MomentSelected: false, F2LambdaProductSeparated: false, F4LambdaProductSeparated: false, PlanckMatchingNative: false, NewtonConstantDerived: false, CosmologicalConstantDerived: false, VacuumSubtractionSelectedNative: false, RenormalizationSchemeNative: false, ObservedComparatorImported: false, NumericalAdapterExecuted: false, NativeNormalizationWrite: false, Verdict: strings.Join([]string{StatusFailedLambdaNativeSelectionRejected, StatusFailedF2F4NativeSelectionRejected, StatusFailedMomentProductsNativePromotion, StatusFailedPlanckMatchingNativePromotion, StatusFailedVacuumSubtractionNativeSelection, StatusFailedNewtonAndCosmoStillNotDerived}, ";"), Reason: "the airlock defines what external data would be needed, but selects none of it natively and executes no matching calculation."}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, Verdict: strings.Join([]string{StatusFailedObservedComparatorNotImported, StatusFirewallNoEmpiricalGravityCosmologyImport, StatusFirewallNativeWriteBlocked}, ";"), Reason: "Gate514 imports no numerical Newton constant, Planck scale, cutoff, f2/f4 moments, moment products, cosmological constant, dark-energy value, or vacuum-subtraction prescription, and writes no native cutoff/renormalization data."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No new physical gravity/cosmology normalization is added. Gate513's stripped heat-kernel ratios remain the last native result in this lane."},
		BridgeEntries:        []string{"A fail-closed bridge schema is defined for Λ, f2, f4, f0, f2Λ², f4Λ⁴, Planck/Newton matching, cosmological comparison, vacuum subtraction, and renormalization metadata.", "Only redacted, metadata-complete, bridge-only rows are accepted by default."},
		EnvironmentalEntries: []string{"Any numerical cutoff, profile moment, Planck/Newton normalization, cosmological comparator, vacuum subtraction, or renormalization scheme."},
		FailedRoutes:         []string{"Selecting Λ from the finite trace.", "Selecting f2 or f4 from the finite trace.", "Promoting f2Λ² or f4Λ⁴ to native products.", "Promoting Planck/Newton matching or vacuum subtraction to native ASHA data.", "Executing observed gravity/cosmology matching by default."},
		OpenTheorems:         []string{"A native regulator/profile selector, if one exists.", "A finite spectral principle selecting vacuum subtraction or renormalization boundary conditions.", "A bridge-only numerical gravity/cosmology adapter that consumes this schema without native writes."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 515, Title: "Bridge-Only Gravity/Cosmology Adapter Dry-Run", Reason: "Gate514 now defines the fail-closed comparator schema; the next safe step is a synthetic dry-run that proves any future numerical adapter remains bridge-only.", PrimaryTask: "Execute a synthetic, explicitly fake Λ/f2/f4/vacuum-subtraction adapter to test formula plumbing, residual reporting, and native-write blocking without importing observed gravity or cosmology data."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate513Inherited && a.Inheritance.StrippedHierarchyNative && nearly(a.Inheritance.A2OverA0Ratio, 1.0/12.0, 1e-12) && nearly(a.Inheritance.A4OverA0Ratio, 1.0/360.0, 1e-12) && nearly(a.Inheritance.A4OverA2Ratio, 1.0/30.0, 1e-12) && !a.Inheritance.Gate513F2Selected && !a.Inheritance.Gate513F4Selected && !a.Inheritance.Gate513CutoffSelected && !a.Inheritance.Gate513NewtonDerived && !a.Inheritance.Gate513CosmologicalDerived && a.Inheritance.Gate513NativeNormalizationBlocked, "Gate514 inheritance invalid"},
		{a.Schema.Executed && a.Schema.RequiredRowCount == 10 && a.Schema.AcceptedRedactedRows == 10 && a.Schema.NumericalRows == 0 && a.Schema.EmpiricalRows == 0 && a.Schema.RowsBridgeOnly && a.Schema.RowsRejectNativePromotion && a.Schema.AllMetadataComplete, "Gate514 schema invalid"},
		{a.Preflight.Executed && a.Preflight.AcceptedCases == 1 && a.Preflight.RejectedCases == 8 && a.Preflight.RejectedNativePromotionCases >= 2 && a.Preflight.RejectedNumericalCases >= 2 && a.Preflight.RejectedMissingMetadataCases == 1 && a.Preflight.RejectedExecutionCases == 1, "Gate514 preflight invalid"},
		{a.Airlock.Executed && !a.Airlock.LambdaCutoffSelected && !a.Airlock.F2MomentSelected && !a.Airlock.F4MomentSelected && !a.Airlock.F2LambdaProductSeparated && !a.Airlock.F4LambdaProductSeparated && !a.Airlock.PlanckMatchingNative && !a.Airlock.NewtonConstantDerived && !a.Airlock.CosmologicalConstantDerived && !a.Airlock.VacuumSubtractionSelectedNative && !a.Airlock.RenormalizationSchemeNative && !a.Airlock.ObservedComparatorImported && !a.Airlock.NumericalAdapterExecuted && !a.Airlock.NativeNormalizationWrite, "Gate514 airlock invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.F2LambdaProductImported && !a.Firewall.F4LambdaProductImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.VacuumSubtractionImported && !a.Firewall.ObservedComparatorImported && !a.Firewall.NativeCutoffRenormalizationWrite, "Gate514 firewall invalid"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate513Inherited,
		StatusCutoffRenormalizationSchemaDefined,
		StatusRequiredRowsEnumerated,
		StatusRedactedBridgeSchemaAccepted,
		StatusMomentProductRowsQuarantined,
		StatusPlanckNewtonCosmologyRowsBridgeOnly,
		StatusComparatorRejectsNativePromotion,
		StatusVacuumSubtractionPolicyFailClosed,
		StatusNoNumericalAdapterExecuted,
		StatusFailedLambdaNativeSelectionRejected,
		StatusFailedF2F4NativeSelectionRejected,
		StatusFailedMomentProductsNativePromotion,
		StatusFailedPlanckMatchingNativePromotion,
		StatusFailedVacuumSubtractionNativeSelection,
		StatusFailedNewtonAndCosmoStillNotDerived,
		StatusFailedObservedComparatorNotImported,
		StatusFirewallNoEmpiricalGravityCosmologyImport,
		StatusFirewallNativeWriteBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 514 does not derive gravity or cosmology normalization. It constructs the explicit fail-closed bridge airlock for the quantities Gate 513 proved missing: Λ, f₂, f₄, f₂Λ², f₄Λ⁴, Planck/Newton matching, cosmological comparison, vacuum subtraction, and renormalization metadata. The only accepted default state is a redacted, bridge-only schema; all numerical imports, native promotions, and adapter executions are blocked."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate513=%t; stripped hierarchy=%t; a2/a0=%.12g; a4/a0=%.12g; a4/a2=%.12g; f2 selected=%t; f4 selected=%t; Λ selected=%t; G derived=%t; Λ_cosmo derived=%t; native normalization blocked=%t", x.Gate513Inherited, x.StrippedHierarchyNative, x.A2OverA0Ratio, x.A4OverA0Ratio, x.A4OverA2Ratio, x.Gate513F2Selected, x.Gate513F4Selected, x.Gate513CutoffSelected, x.Gate513NewtonDerived, x.Gate513CosmologicalDerived, x.Gate513NativeNormalizationBlocked)
}
func FormatSchema(x Schema) string {
	return fmt.Sprintf("required rows=%d; accepted redacted rows=%d; numerical rows=%d; empirical rows=%d; bridge-only=%t; reject native promotion=%t; metadata complete=%t", x.RequiredRowCount, x.AcceptedRedactedRows, x.NumericalRows, x.EmpiricalRows, x.RowsBridgeOnly, x.RowsRejectNativePromotion, x.AllMetadataComplete)
}
func FormatPreflight(x Preflight) string {
	return fmt.Sprintf("cases=%d; accepted=%d; rejected=%d; native-promotion rejects=%d; numerical rejects=%d; metadata rejects=%d; execution rejects=%d", len(x.Cases), x.AcceptedCases, x.RejectedCases, x.RejectedNativePromotionCases, x.RejectedNumericalCases, x.RejectedMissingMetadataCases, x.RejectedExecutionCases)
}
func FormatAirlock(x Airlock) string {
	return fmt.Sprintf("Λ selected=%t; f2 selected=%t; f4 selected=%t; f2Λ² separated=%t; f4Λ4 separated=%t; Planck native=%t; G derived=%t; Λ_cosmo derived=%t; subtraction native=%t; renormalization native=%t; observed imported=%t; adapter executed=%t; native write=%t", x.LambdaCutoffSelected, x.F2MomentSelected, x.F4MomentSelected, x.F2LambdaProductSeparated, x.F4LambdaProductSeparated, x.PlanckMatchingNative, x.NewtonConstantDerived, x.CosmologicalConstantDerived, x.VacuumSubtractionSelectedNative, x.RenormalizationSchemeNative, x.ObservedComparatorImported, x.NumericalAdapterExecuted, x.NativeNormalizationWrite)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("G imported=%t; Planck imported=%t; Λ imported=%t; f2 imported=%t; f4 imported=%t; f2Λ² imported=%t; f4Λ4 imported=%t; Λ_cosmo imported=%t; dark energy imported=%t; subtraction imported=%t; observed comparator imported=%t; native write=%t", x.NewtonConstantImported, x.PlanckScaleImported, x.CutoffLambdaImported, x.F2MomentImported, x.F4MomentImported, x.F2LambdaProductImported, x.F4LambdaProductImported, x.CosmologicalConstantImported, x.DarkEnergyImported, x.VacuumSubtractionImported, x.ObservedComparatorImported, x.NativeCutoffRenormalizationWrite)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 514 Registry Audit — Spectral Cutoff and Renormalization Airlock Comparator\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Comparator schema\n\n" + a.Schema.Reason + "\n\n```text\n" + FormatSchema(a.Schema) + "\n```\n\n")
	b.WriteString("Required rows:\n\n")
	for _, r := range a.Schema.Rows {
		b.WriteString(fmt.Sprintf("- `%s` (`%s`): %s — %s\n", r.Key, r.Symbol, r.Role, r.Dimension))
	}
	b.WriteString("\n## Preflight rejection ledger\n\n" + a.Preflight.Reason + "\n\n```text\n" + FormatPreflight(a.Preflight) + "\n```\n\n")
	b.WriteString("## Cutoff and renormalization airlock\n\n" + a.Airlock.Reason + "\n\n```text\n" + FormatAirlock(a.Airlock) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate515 should be:\n\n```text\nGate 515 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
func nearly(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
