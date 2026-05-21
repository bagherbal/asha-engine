// Package generation2r3dualsealbridgetheoremfinalboundaryr4preconditionaudit implements
// Gate 949: R3 DualSeal Bridge-Theorem Final Boundary and R4 Precondition Audit.
//
// Gate 949 follows Gate 948's dual-sealed R3 tracebridge boundary. It freezes
// the bridge-theorem classification and defines the only lawful way later R4
// generation/flavor work may proceed: under explicit ScalarSourceSeal and
// PostOrientationSeal firewalls.
package generation2r3dualsealbridgetheoremfinalboundaryr4preconditionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE949-GENERATION2R3DUALSEALBRIDGETHEOREFINALBOUNDARYR4PRECONDITIONAUDIT"
	InheritedStatus = "R3_TRACEBRIDGE_TEST_PASSED_DUAL_SEALED_NOT_NATIVE"
	Verdict         = "R3_TRACEBRIDGE_FINALIZED_AS_SCALAR_SOURCE_SEALED_AND_POST_ORIENTATION_SEALED_BRIDGE_THEOREM_CANDIDATE_NOT_NATIVE"
	Classification  = "R3_DUALSEAL_TRACEBRIDGE_THEOREM_CANDIDATE_FINALIZED_NOT_NATIVE"
	ShortStatus     = "R3_DUALSEAL_TRACEBRIDGE_FINAL_BOUNDARY"
	NextGate        = "NEXT_GATE950_R4_GENERATIONCARRIER_PRECONDITION_AUDIT_UNDER_R3_DUALSEAL"
)

const (
	Ssplit  = 0.0012924448188162962
	AlphaB  = 0.0003878958469680527
	NEff    = 3.002327375081808
	CYukawa = 0.9992248096922658
	CHiggs  = 1.037220510866514
)

type DualSealBoundary struct {
	TracebridgeTestPassed bool
	ClosureFactored       bool
	ScalarSourceSeal      bool
	PostOrientationSeal   bool
	NativeR3              bool
	OfficialLedgerUpdate  bool
	PhysicalAssignment    bool
	GenerationCarrier     bool
	FlavorOrientation     bool
	IndividualYukawa      bool
	R4NativeSpectrum      bool
	Status                string
}

type R4PreconditionPolicy struct {
	MayProceedUnderSeal       bool
	RequiresScalarSourceSeal  bool
	RequiresPostOrientSeal    bool
	AllowsGenerationAudit     bool
	AllowsFlavorAudit         bool
	AllowsYukawaPrecondition  bool
	AllowsPhysicalAssignment  bool
	AllowsOfficialLedger      bool
	AllowsNativeSpectrumClaim bool
	Warnings                  []string
}

type AuditItem struct {
	Name      string
	Status    string
	Supports  []string
	Firewalls []string
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	Ssplit         float64
	AlphaB         float64
	NEff           float64
	CYukawa        float64
	CHiggs         float64
	Boundary       DualSealBoundary
	Policy         R4PreconditionPolicy
	Items          []AuditItem
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	boundary := DefaultBoundary()
	policy := DefaultPolicy()
	items := DefaultItems()
	if !boundary.TracebridgeTestPassed || !boundary.ClosureFactored || !boundary.ScalarSourceSeal || !boundary.PostOrientationSeal {
		return Analysis{}, fmt.Errorf("Gate 949 requires test-passed closure-factored dual-sealed tracebridge: %#v", boundary)
	}
	if boundary.NativeR3 || boundary.OfficialLedgerUpdate || boundary.PhysicalAssignment || boundary.GenerationCarrier || boundary.FlavorOrientation || boundary.IndividualYukawa || boundary.R4NativeSpectrum {
		return Analysis{}, fmt.Errorf("Gate 949 must not grant native R3 or downstream R4/physical claims: %#v", boundary)
	}
	if !policy.MayProceedUnderSeal || !policy.RequiresScalarSourceSeal || !policy.RequiresPostOrientSeal {
		return Analysis{}, fmt.Errorf("Gate 949 R4 policy must require explicit dual seals: %#v", policy)
	}
	if policy.AllowsPhysicalAssignment || policy.AllowsOfficialLedger || policy.AllowsNativeSpectrumClaim {
		return Analysis{}, fmt.Errorf("Gate 949 R4 policy must not allow physical/official/native-spectrum claims: %#v", policy)
	}
	if len(items) != 6 {
		return Analysis{}, fmt.Errorf("expected six audit items, got %d", len(items))
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Ssplit:         Ssplit,
		AlphaB:         AlphaB,
		NEff:           NEff,
		CYukawa:        CYukawa,
		CHiggs:         CHiggs,
		Boundary:       boundary,
		Policy:         policy,
		Items:          items,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 949 freezes the R3 tracebridge as a test-passed, Z2-equivariant, closure-factored, scalar-source-sealed, post-orientation-sealed bridge theorem candidate. Native R3 is not granted. Later R4 generation/flavor work may proceed only under explicit ScalarSourceSeal(S_split) and PostOrientationSeal(A_F^orient) warnings.",
	}, nil
}

func DefaultBoundary() DualSealBoundary {
	return DualSealBoundary{
		TracebridgeTestPassed: true,
		ClosureFactored:       true,
		ScalarSourceSeal:      true,
		PostOrientationSeal:   true,
		NativeR3:              false,
		OfficialLedgerUpdate:  false,
		PhysicalAssignment:    false,
		GenerationCarrier:     false,
		FlavorOrientation:     false,
		IndividualYukawa:      false,
		R4NativeSpectrum:      false,
		Status:                "TEST_PASSED_CLOSURE_FACTORED_DUAL_SEALED_TRACEBRIDGE_NOT_NATIVE",
	}
}

func DefaultPolicy() R4PreconditionPolicy {
	warnings := []string{
		"depends on scalar-source-sealed R3 tracebridge",
		"depends on post-orientation stabilizer layer",
		"not native R3",
		"not official physical spectrum",
	}
	return R4PreconditionPolicy{
		MayProceedUnderSeal:       true,
		RequiresScalarSourceSeal:  true,
		RequiresPostOrientSeal:    true,
		AllowsGenerationAudit:     true,
		AllowsFlavorAudit:         true,
		AllowsYukawaPrecondition:  true,
		AllowsPhysicalAssignment:  false,
		AllowsOfficialLedger:      false,
		AllowsNativeSpectrumClaim: false,
		Warnings:                  warnings,
	}
}

func DefaultItems() []AuditItem {
	return []AuditItem{
		{
			Name:   "inherited validated R3 tracebridge",
			Status: "TEST_PASSED_DIAGNOSTIC_BRIDGE_VALUES",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_TEST_PASSED",
				"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_CLOSURE_FACTORED",
				"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA_B_UNDER_SEAL",
				"CONDITIONAL_SUPPORT_TRACE_ROWS_RECONSTRUCT_OPERATOR_DIAGNOSTICS_UNDER_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
		{
			Name:   "ScalarSourceSeal",
			Status: "FINALIZED_SEAL",
			Supports: []string{
				"CONDITIONAL_SUPPORT_SCALAR_SOURCE_SEAL_FINALIZED",
			},
			Firewalls: []string{
				"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
				"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
			},
		},
		{
			Name:   "PostOrientationSeal",
			Status: "FINALIZED_SEAL",
			Supports: []string{
				"CONDITIONAL_SUPPORT_POST_ORIENTATION_SEAL_FINALIZED",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
				"FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM",
			},
		},
		{
			Name:   "R4 precondition permission under explicit dual seal",
			Status: "ALLOWED_ONLY_AS_PRECONDITION_AUDITS_UNDER_DUAL_SEAL",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R4_WORK_MAY_PROCEED_ONLY_UNDER_EXPLICIT_DUAL_SEAL",
				"CONDITIONAL_SUPPORT_NATIVE_R3_AND_R4_FLAVOR_REMAIN_SEPARATED",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
			},
		},
		{
			Name:   "forbidden downstream claims",
			Status: "BLOCKED_UNLESS_NEW_NATIVE_GATES_CERTIFY_MISSING_OBJECTS",
			Firewalls: []string{
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_R3_TRACEBRIDGE_NOT_NATIVE_THEOREM",
			},
		},
		{
			Name:   "final R3 branch boundary",
			Status: "FROZEN_DUALSEAL_BRIDGE_THEOREM_BOUNDARY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_NATIVE_R3_AND_R4_FLAVOR_REMAIN_SEPARATED",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_TEST_PASSED",
		"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_CLOSURE_FACTORED",
		"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA_B_UNDER_SEAL",
		"CONDITIONAL_SUPPORT_TRACE_ROWS_RECONSTRUCT_OPERATOR_DIAGNOSTICS_UNDER_SEAL",
		"CONDITIONAL_SUPPORT_SCALAR_SOURCE_SEAL_FINALIZED",
		"CONDITIONAL_SUPPORT_POST_ORIENTATION_SEAL_FINALIZED",
		"CONDITIONAL_SUPPORT_R4_WORK_MAY_PROCEED_ONLY_UNDER_EXPLICIT_DUAL_SEAL",
		"CONDITIONAL_SUPPORT_NATIVE_R3_AND_R4_FLAVOR_REMAIN_SEPARATED",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_NOT_NATIVE_THEOREM",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func FormatBoundary(b DualSealBoundary) string {
	return fmt.Sprintf("%s: testPassed=%v closureFactored=%v scalarSeal=%v postOrientationSeal=%v nativeR3=%v officialLedger=%v physical=%v generation=%v flavor=%v individualYukawa=%v r4NativeSpectrum=%v", b.Status, b.TracebridgeTestPassed, b.ClosureFactored, b.ScalarSourceSeal, b.PostOrientationSeal, b.NativeR3, b.OfficialLedgerUpdate, b.PhysicalAssignment, b.GenerationCarrier, b.FlavorOrientation, b.IndividualYukawa, b.R4NativeSpectrum)
}

func FormatPolicy(p R4PreconditionPolicy) string {
	return fmt.Sprintf("R4Policy: mayProceedUnderSeal=%v requiresScalarSeal=%v requiresPostOrientSeal=%v allowsGenerationAudit=%v allowsFlavorAudit=%v allowsYukawaPrecondition=%v allowsPhysicalAssignment=%v allowsOfficialLedger=%v allowsNativeSpectrum=%v warnings=[%s]", p.MayProceedUnderSeal, p.RequiresScalarSourceSeal, p.RequiresPostOrientSeal, p.AllowsGenerationAudit, p.AllowsFlavorAudit, p.AllowsYukawaPrecondition, p.AllowsPhysicalAssignment, p.AllowsOfficialLedger, p.AllowsNativeSpectrumClaim, strings.Join(p.Warnings, ";"))
}

func FormatItem(i AuditItem) string {
	return fmt.Sprintf("%s: status=%s supports=[%s] firewalls=[%s]", i.Name, i.Status, strings.Join(i.Supports, ","), strings.Join(i.Firewalls, ","))
}

func ItemNotes(items []AuditItem) []string {
	out := make([]string, 0, len(items))
	for _, i := range items {
		out = append(out, FormatItem(i))
	}
	return out
}

func ItemSupports(items []AuditItem) []string {
	var out []string
	for _, i := range items {
		out = append(out, i.Supports...)
	}
	return out
}

func ItemFailures(items []AuditItem) []string {
	var out []string
	for _, i := range items {
		out = append(out, i.Firewalls...)
	}
	return out
}

func containsAll(haystack, needles []string) bool {
	set := map[string]struct{}{}
	for _, h := range haystack {
		set[h] = struct{}{}
	}
	for _, n := range needles {
		if _, ok := set[n]; !ok {
			return false
		}
	}
	return true
}

func appendAll(items ...[]string) []string {
	var out []string
	for _, group := range items {
		out = append(out, group...)
	}
	return out
}

func stringsJoin(items []string) string { return strings.Join(items, " | ") }
