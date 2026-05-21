// Package generation2k7minustrialityintertwinerconstructionr4stabilizationaudit implements
// Gate 955: K7Minus Triality Intertwiner Construction and R4 GenerationCarrier Stabilization Audit.
//
// This gate attempts the full GenerationCarrierMap construction at the
// unoriented aggregate level. It deliberately separates an abstract C3 action
// model on a three-dimensional carrier from a native triality restriction and
// from a typed R3 tracebody intertwiner. The former can pass algebraic sanity
// checks; the latter two remain uncertified unless a canonical operator and
// typed action map are actually supplied.
package generation2k7minustrialityintertwinerconstructionr4stabilizationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID         = "GATE955-GENERATION2K7MINUSTRIALITYINTERTWINERCONSTRUCTIONR4STABILIZATIONAUDIT"
	InheritedStatus = "R4_K7_MINUS_TRIALITY_COUPLING_CANDIDATE_NO_INTERTWINER"
	Verdict         = "K7_MINUS_ADMITS_ABSTRACT_ORDER_THREE_ACTION_MODEL_BUT_NO_NATIVE_TRIALITY_RESTRICTION_OR_R3_TRACEBODY_INTERTWINER_CERTIFIED"
	Classification  = "R4_GENERATION_ACTION_MODEL_SUPPORTED_NATIVE_INTERTWINER_MISSING"
	ShortStatus     = "R4_ABSTRACT_K7_MINUS_C3_ACTION_NO_GENERATION_MAP"
	NextGate        = "NEXT_GATE956_R3_TRACEBODY_INTERTWINER_REPAIR_OR_ALTERNATIVE_GENERATION_CARRIER_AUDIT"
)

const (
	K7Dim      = 7
	K7PlusDim  = 4
	K7MinusDim = 3

	TraceRowRankA = 3
	TraceRowRankB = 3
	TraceRowRankC = 1

	AlphaB  = 0.0003878958469680527
	NEff    = 3.002327375081808
	CYukawa = 0.9992248096922658
	CHiggs  = 1.037220510866514
	Tol     = 1e-12
)

type AuditItem struct {
	Name      string
	Status    string
	Supports  []string
	Firewalls []string
}

type C3ActionModel struct {
	Matrix                 [3][3]float64
	OrderThree             bool
	Nontrivial             bool
	Trace                  float64
	Determinant            float64
	MetricPreserving       bool
	OrbitSpan              int
	CanonicalTrialityInput bool
	NativeRestriction      bool
}

type IntertwinerAttempt struct {
	CandidateName                string
	Target                       string
	Certified                    bool
	ArbitraryBasisChoiceRequired bool
	UsesR3RowsAsGenerationLabels bool
	PreservesDualSeal            bool
	UsesFlavorBacksolve          bool
	UsesObservedYukawaOrMassData bool
	UsesCKMPMNSInput             bool
	FlavorOrientationCertified   bool
	IndividualYukawaCertified    bool
	PhysicalAssignmentCertified  bool
}

type Analysis struct {
	AuditID                     string
	Inherited                   string
	Verdict                     string
	Classification              string
	ShortStatus                 string
	NativeR3                    bool
	DualSealRequired            bool
	GenerationCarrierCertified  bool
	FlavorOrientationCertified  bool
	IndividualYukawaCertified   bool
	PhysicalAssignmentCertified bool
	OfficialLedgerUpdate        bool
	K7MinusDimension            int
	K7PlusDimension             int
	TraceRowRanks               []int
	Action                      C3ActionModel
	Intertwiner                 IntertwinerAttempt
	Items                       []AuditItem
	Supports                    []string
	Failures                    []string
	Final                       string
}

func BuildDefault() (Analysis, error) {
	action := BuildAbstractC3ActionModel()
	intertwiner := IntertwinerAttempt{
		CandidateName:                "K7MinusTrialityTracebodyIntertwiner",
		Target:                       "dual-sealed aggregate R3 tracebody",
		Certified:                    false,
		ArbitraryBasisChoiceRequired: true,
		UsesR3RowsAsGenerationLabels: true,
		PreservesDualSeal:            true,
		UsesFlavorBacksolve:          false,
		UsesObservedYukawaOrMassData: false,
		UsesCKMPMNSInput:             false,
		FlavorOrientationCertified:   false,
		IndividualYukawaCertified:    false,
		PhysicalAssignmentCertified:  false,
	}
	items := DefaultItems(action, intertwiner)
	if len(items) != 8 {
		return Analysis{}, fmt.Errorf("expected 8 audit items, got %d", len(items))
	}
	a := Analysis{
		AuditID:                     AuditID,
		Inherited:                   InheritedStatus,
		Verdict:                     Verdict,
		Classification:              Classification,
		ShortStatus:                 ShortStatus,
		NativeR3:                    false,
		DualSealRequired:            true,
		GenerationCarrierCertified:  false,
		FlavorOrientationCertified:  false,
		IndividualYukawaCertified:   false,
		PhysicalAssignmentCertified: false,
		OfficialLedgerUpdate:        false,
		K7MinusDimension:            K7MinusDim,
		K7PlusDimension:             K7PlusDim,
		TraceRowRanks:               []int{TraceRowRankA, TraceRowRankB, TraceRowRankC},
		Action:                      action,
		Intertwiner:                 intertwiner,
		Items:                       items,
		Supports:                    Supports(),
		Failures:                    Failures(),
		Final:                       "Gate 955 attempts the full K7^-/triality/R3 tracebody GenerationCarrierMap construction. It finds that K7^- can host an abstract nontrivial order-three C3 action preserving the three-dimensional carrier model, but no native triality restriction to K7^- and no typed K7MinusTrialityTracebodyIntertwiner are certified. R4 generation carrier stabilization therefore fails at the native-intertwiner layer; flavor and individual Yukawa claims remain firewalled.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if !a.DualSealRequired || a.NativeR3 || a.OfficialLedgerUpdate {
		return fmt.Errorf("Gate 955 must preserve R3 dual seal and avoid native R3/official update overclaim")
	}
	if a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		return fmt.Errorf("Gate 955 overclaimed R4 downstream theorem status")
	}
	if a.K7MinusDimension != 3 || a.K7PlusDimension != 4 || K7Dim != a.K7MinusDimension+a.K7PlusDimension {
		return fmt.Errorf("bad K7 Hodge-polarity dimensions")
	}
	if len(a.TraceRowRanks) != 3 || a.TraceRowRanks[0] != 3 || a.TraceRowRanks[1] != 3 || a.TraceRowRanks[2] != 1 {
		return fmt.Errorf("R3 tracebody must remain aggregate 3,3,1 rows")
	}
	if !a.Action.OrderThree || !a.Action.Nontrivial || !a.Action.MetricPreserving || a.Action.OrbitSpan != 3 {
		return fmt.Errorf("abstract C3 action sanity checks should pass to make the construction attempt meaningful: %#v", a.Action)
	}
	if a.Action.CanonicalTrialityInput || a.Action.NativeRestriction {
		return fmt.Errorf("native triality restriction was overclaimed")
	}
	if a.Intertwiner.Certified || a.Intertwiner.FlavorOrientationCertified || a.Intertwiner.IndividualYukawaCertified || a.Intertwiner.PhysicalAssignmentCertified {
		return fmt.Errorf("intertwiner/flavor/Yukawa certification overclaimed")
	}
	if !a.Intertwiner.ArbitraryBasisChoiceRequired || !a.Intertwiner.UsesR3RowsAsGenerationLabels || !a.Intertwiner.PreservesDualSeal {
		return fmt.Errorf("intertwiner obstruction flags not set correctly: %#v", a.Intertwiner)
	}
	if a.Intertwiner.UsesFlavorBacksolve || a.Intertwiner.UsesObservedYukawaOrMassData || a.Intertwiner.UsesCKMPMNSInput {
		return fmt.Errorf("noncircularity failed: forbidden empirical/flavor input used")
	}
	return nil
}

func BuildAbstractC3ActionModel() C3ActionModel {
	// This is the strongest algebraic *model* of a threefold action on a 3D
	// real carrier: the cyclic permutation representation. It is deliberately
	// marked non-native because no canonical triality operator restricted from
	// the ASHA Cl(1,7)/K7 construction is supplied by this gate.
	m := [3][3]float64{
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	m2 := matMul(m, m)
	m3 := matMul(m2, m)
	return C3ActionModel{
		Matrix:                 m,
		OrderThree:             matClose(m3, identity3(), Tol),
		Nontrivial:             !matClose(m, identity3(), Tol),
		Trace:                  trace3(m),
		Determinant:            det3(m),
		MetricPreserving:       preservesNegativeMetric(m),
		OrbitSpan:              orbitSpan(m, [3]float64{1, 2, 4}),
		CanonicalTrialityInput: false,
		NativeRestriction:      false,
	}
}

func DefaultItems(action C3ActionModel, intertwiner IntertwinerAttempt) []AuditItem {
	return []AuditItem{
		{
			Name:   "triality restriction to K7-minus",
			Status: "BLOCKED_NO_NATIVE_RESTRICTION_OPERATOR",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONGEST_DIMENSION_THREE_CARRIER_SHAPE",
				"CONDITIONAL_SUPPORT_TRIALITY_RESTRICTION_TEST_IS_CORRECT_FIRST_CONSTRUCTION_TEST",
			},
			Firewalls: []string{
				"FAILED_ROUTE_TRIALITY_DOES_NOT_CANONICALLY_RESTRICT_TO_K7_MINUS_IN_CURRENT_CERTIFICATE",
				"FAILED_ROUTE_NO_NATIVE_TRIALITY_RESTRICTION_OPERATOR_ON_K7_MINUS",
			},
		},
		{
			Name:   "abstract threefold action sanity test",
			Status: statusIf(action.OrderThree && action.Nontrivial, "ABSTRACT_C3_MODEL_PASSES_ALGEBRAIC_SHAPE", "ABSTRACT_C3_MODEL_FAILED"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ACTION_HAS_ORDER_THREE_SHAPE",
				"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ACTION_IS_NONTRIVIAL",
				"CONDITIONAL_SUPPORT_ABSTRACT_ACTION_HAS_TRACE_ZERO_DETERMINANT_ONE_SIGNATURE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_ABSTRACT_C3_ACTION_IS_NOT_NATIVE_TRIALITY_RESTRICTION",
				"FAILED_ROUTE_TRIALITY_ACTION_NOT_CANONICAL",
			},
		},
		{
			Name:   "metric and Hodge compatibility",
			Status: statusIf(action.MetricPreserving, "ABSTRACT_ACTION_PRESERVES_K7_MINUS_BILINEAR_MODEL", "ABSTRACT_ACTION_BREAKS_BILINEAR_MODEL"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_ABSTRACT_ACTION_PRESERVES_K7_MINUS_BILINEAR_STRUCTURE",
				"CONDITIONAL_SUPPORT_ABSTRACT_ACTION_STAYS_WITHIN_K7_MINUS_MODEL_SPACE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_BILINEAR_COMPATIBILITY_OF_ABSTRACT_MODEL_NOT_NATIVE_HODGE_POLARITY_THEOREM",
				"FAILED_ROUTE_NO_CERTIFIED_COMMUTATOR_WITH_NATIVE_K7_HODGE_POLARITY_OPERATOR",
			},
		},
		{
			Name:   "orbit-span generation-carrier test",
			Status: statusIf(action.OrbitSpan == 3, "ABSTRACT_ORBITS_CAN_SPAN_THREE_SLOTS", "ABSTRACT_ORBITS_DO_NOT_SPAN"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ORBITS_CAN_SPAN_THREE_GENERATION_SLOTS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_ORBIT_SPAN_OF_ABSTRACT_ACTION_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_ONLY_DIMENSION_THREE_PLUS_MODEL_ACTION_AVAILABLE_NO_NATIVE_GENERATING_ORBIT",
			},
		},
		{
			Name:   "R3 tracebody intertwiner",
			Status: "BLOCKED_NO_TYPED_INTERTWINER",
			Supports: []string{
				"CONDITIONAL_SUPPORT_REQUIRED_OBJECT_IS_K7_MINUS_TRIALITY_TRACEBODY_INTERTWINER",
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_TARGET_ONLY_AS_DUALSEALED_AGGREGATE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER",
				"FAILED_ROUTE_INTERTWINER_ONLY_EXISTS_AFTER_ARBITRARY_BASIS_CHOICE",
				"FAILED_ROUTE_INTERTWINER_USES_R3_ROWS_AS_GENERATION_LABELS",
			},
		},
		{
			Name:   "noncircularity",
			Status: statusIf(!intertwiner.UsesFlavorBacksolve && !intertwiner.UsesObservedYukawaOrMassData && !intertwiner.UsesCKMPMNSInput, "NONEMPIRICAL_INPUT_FIREWALL_PRESERVED", "FORBIDDEN_FLAVOR_INPUT_USED"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_GENERATION_CARRIER_CONSTRUCTION_ATTEMPT_IS_NONEMPIRICAL_AND_NONCIRCULAR",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_FLAVOR_FORMULA_BACKSOLVE_USED_BUT_GENERATION_MAP_STILL_MISSING",
			},
		},
		{
			Name:   "R3 dual-seal preservation",
			Status: "DUAL_SEAL_PRESERVED",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R4_GENERATION_CARRIER_ATTEMPT_PRESERVES_R3_DUALSEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
			},
		},
		{
			Name:   "flavor firewall",
			Status: "FLAVOR_ORIENTATION_STILL_MISSING",
			Supports: []string{
				"CONDITIONAL_SUPPORT_IF_GENERATION_CARRIER_EVER_CERTIFIES_NEXT_WOUND_IS_FLAVOR_ORIENTATION_MAP",
			},
			Firewalls: []string{
				"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONGEST_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ACTION_HAS_ORDER_THREE_SHAPE",
		"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ACTION_IS_NONTRIVIAL",
		"CONDITIONAL_SUPPORT_ABSTRACT_ACTION_PRESERVES_K7_MINUS_BILINEAR_STRUCTURE",
		"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ORBITS_CAN_SPAN_THREE_GENERATION_SLOTS",
		"CONDITIONAL_SUPPORT_REQUIRED_OBJECT_IS_K7_MINUS_TRIALITY_TRACEBODY_INTERTWINER",
		"CONDITIONAL_SUPPORT_GENERATION_CARRIER_CONSTRUCTION_ATTEMPT_IS_NONEMPIRICAL_AND_NONCIRCULAR",
		"CONDITIONAL_SUPPORT_R4_GENERATION_CARRIER_ATTEMPT_PRESERVES_R3_DUALSEAL",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_TRIALITY_DOES_NOT_CANONICALLY_RESTRICT_TO_K7_MINUS_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_NO_NATIVE_TRIALITY_RESTRICTION_OPERATOR_ON_K7_MINUS",
		"FAILED_ROUTE_ABSTRACT_C3_ACTION_IS_NOT_NATIVE_TRIALITY_RESTRICTION",
		"FAILED_ROUTE_TRIALITY_ACTION_NOT_CANONICAL",
		"FAILED_ROUTE_ORBIT_SPAN_OF_ABSTRACT_ACTION_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER",
		"FAILED_ROUTE_INTERTWINER_ONLY_EXISTS_AFTER_ARBITRARY_BASIS_CHOICE",
		"FAILED_ROUTE_INTERTWINER_USES_R3_ROWS_AS_GENERATION_LABELS",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
		"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func statusIf(ok bool, yes, no string) string {
	if ok {
		return yes
	}
	return no
}

func identity3() [3][3]float64 {
	return [3][3]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
}

func matMul(a, b [3][3]float64) [3][3]float64 {
	var out [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func matClose(a, b [3][3]float64, tol float64) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(a[i][j]-b[i][j]) > tol {
				return false
			}
		}
	}
	return true
}

func trace3(m [3][3]float64) float64 { return m[0][0] + m[1][1] + m[2][2] }

func det3(m [3][3]float64) float64 {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

func transpose3(m [3][3]float64) [3][3]float64 {
	return [3][3]float64{{m[0][0], m[1][0], m[2][0]}, {m[0][1], m[1][1], m[2][1]}, {m[0][2], m[1][2], m[2][2]}}
}

func negativeMetric3() [3][3]float64 {
	return [3][3]float64{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}}
}

func preservesNegativeMetric(m [3][3]float64) bool {
	left := matMul(transpose3(m), matMul(negativeMetric3(), m))
	return matClose(left, negativeMetric3(), Tol)
}

func matVec(m [3][3]float64, v [3]float64) [3]float64 {
	return [3]float64{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2],
	}
}

func orbitSpan(m [3][3]float64, v [3]float64) int {
	v1 := matVec(m, v)
	v2 := matVec(matMul(m, m), v)
	cols := [3][3]float64{
		{v[0], v1[0], v2[0]},
		{v[1], v1[1], v2[1]},
		{v[2], v1[2], v2[2]},
	}
	if math.Abs(det3(cols)) > Tol {
		return 3
	}
	// Degenerate fallback: enough for this finite audit.
	if math.Abs(v[0])+math.Abs(v[1])+math.Abs(v[2]) > Tol {
		return 1
	}
	return 0
}

func ItemSupports(items []AuditItem) []string {
	var out []string
	for _, it := range items {
		out = append(out, it.Supports...)
	}
	return out
}

func ItemFailures(items []AuditItem) []string {
	var out []string
	for _, it := range items {
		out = append(out, it.Firewalls...)
	}
	return out
}

func appendAll(parts ...[]string) []string {
	var out []string
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}

func containsAll(hay []string, needles []string) bool {
	set := map[string]bool{}
	for _, h := range hay {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}

func stringsJoin(v []string) string { return strings.Join(v, "; ") }

func ItemNotes(items []AuditItem) []string {
	notes := make([]string, 0, len(items))
	for _, it := range items {
		notes = append(notes, it.Name+" => "+it.Status)
	}
	return notes
}
