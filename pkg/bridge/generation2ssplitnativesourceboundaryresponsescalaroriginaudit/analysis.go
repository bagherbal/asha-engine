// Package generation2ssplitnativesourceboundaryresponsescalaroriginaudit implements
// Gate 941: S_split Native Source and BoundaryResponse Scalar Origin Audit.
//
// Gate 941 follows Gate 940's failure localization. It traces the scalar
// S_split back to the earlier augmented-chamber/defect-trace rail and audits
// whether it can be promoted from bridge scalar to native boundary-response
// scalar. The result is deliberately conservative: the origin is source-typed,
// the uniform B2 insertion is strengthened, but Certificate II is not passed
// because the native transport theorem remains missing.
package generation2ssplitnativesourceboundaryresponsescalaroriginaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE941-GENERATION2SSPLITNATIVESOURCEBOUNDARYRESPONSESCALARORIGINAUDIT"
	InheritedStatus = "R3_NATIVE_CLAUSE_EXECUTION_PARTIAL_PASS_NOT_NATIVE_R3"
	Verdict         = "S_SPLIT_SOURCE_TRACED_TO_AUGMENTED_CHAMBER_DEFECT_SPLIT_AND_UNIFORM_B2_RESPONSE_PARAMETER_CANDIDATE_BUT_NATIVE_TRANSPORT_THEOREM_MISSING"
	Classification  = "R3_S_SPLIT_SOURCE_PARTIALLY_TRACED_NATIVE_TRANSPORT_BLOCKED"
	ShortStatus     = "R3_S_SPLIT_ORIGIN_TRACED_TO_AUGMENTED_CHAMBER_BUT_NOT_NATIVE"
	NextGate        = "NEXT_PRESSURE_GATE942_AUGMENTED_CHAMBER_DEFECT_SPLIT_TO_BOUNDARY_PAIR_RESPONSE_TRANSPORT_AUDIT"
)

const (
	Ssplit      = 0.0012924448188162962
	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281
	AlphaB      = 0.0003878958469680527
	H72Rank     = 72
)

type SourceStatus string

const (
	SourceBridgeStrongNotNative SourceStatus = "BRIDGE_STRONG_NOT_NATIVE"
	SourceNativeBlocked         SourceStatus = "NATIVE_TRANSPORT_BLOCKED"
)

type ScalarOrigin struct {
	Name           string
	Expression     string
	Status         SourceStatus
	Supports       []string
	Failures       []string
	Interpretation string
}

type TransportMap struct {
	Name              string
	Domain            string
	Codomain          string
	Candidate         string
	NativeCertified   bool
	CertificatePassed bool
	Supports          []string
	Failures          []string
}

type BoundaryResponseUse struct {
	Response       string
	UniformFactors []string
	AlphaLinear    float64
	AlphaQuadratic float64
	AlphaTotal     float64
}

type Analysis struct {
	AuditID           string
	Inherited         string
	Verdict           string
	Classification    string
	ShortStatus       string
	Ssplit            float64
	Origin            ScalarOrigin
	Transport         TransportMap
	Response          BoundaryResponseUse
	CertificateIIPass bool
	Supports          []string
	Failures          []string
	Final             string
}

func BuildDefault() (Analysis, error) {
	origin := DefaultOrigin()
	transport := DefaultTransportMap()
	response := DefaultResponseUse()
	if response.AlphaTotal != AlphaB {
		return Analysis{}, fmt.Errorf("alpha total changed: got %.19g want %.19g", response.AlphaTotal, AlphaB)
	}
	if transport.CertificatePassed || transport.NativeCertified {
		return Analysis{}, fmt.Errorf("Gate 941 must not pass Certificate II without native S_split transport")
	}
	return Analysis{
		AuditID:           AuditID,
		Inherited:         InheritedStatus,
		Verdict:           Verdict,
		Classification:    Classification,
		ShortStatus:       ShortStatus,
		Ssplit:            Ssplit,
		Origin:            origin,
		Transport:         transport,
		Response:          response,
		CertificateIIPass: false,
		Supports:          Supports(),
		Failures:          Failures(),
		Final:             "Gate 941 traces S_split to the augmented-chamber defect split and source-types its uniform B2 response insertion. Certificate II is weakened but not passed: the native augmented-chamber-to-B2 transport theorem remains missing, so native R3 is not granted.",
	}, nil
}

func DefaultOrigin() ScalarOrigin {
	return ScalarOrigin{
		Name:       "augmented chamber defect-trace split",
		Expression: "D_base=(7/72)S_split with S_split=(R_3-1)+lambda(Lambda_12)",
		Status:     SourceBridgeStrongNotNative,
		Supports: []string{
			"CONDITIONAL_SUPPORT_S_SPLIT_HAS_PRIOR_AUGMENTED_CHAMBER_DEFECT_TRACE_ORIGIN",
			"CONDITIONAL_SUPPORT_S_SPLIT_IS_NOT_NEWLY_INSERTED_FOR_ALPHA_BRANCH",
			"CONDITIONAL_SUPPORT_S_SPLIT_ALREADY_COUPLES_TO_72_CHAMBER_RESPONSE_STRUCTURE",
			"CONDITIONAL_SUPPORT_S_SPLIT_SOURCE_IS_COMPATIBLE_WITH_BOUNDARY_AUGMENTED_H72_LANE",
		},
		Failures: []string{
			"FAILED_ROUTE_AUGMENTED_CHAMBER_DEFECT_TRACE_ORIGIN_NOT_YET_NATIVE_B2_RESPONSE_TRANSPORT",
		},
		Interpretation: "S_split is inherited from the augmented chamber / defect-trace rail rather than introduced ad hoc in the alpha branch.",
	}
}

func DefaultTransportMap() TransportMap {
	return TransportMap{
		Name:              "AugmentedChamberDefectSplitToBoundaryPairResponseTransportTheorem",
		Domain:            "augmented chamber defect split scalar S_split",
		Codomain:          "scalar parameter s of reduced rank-two boundary-pair response R_B(s)",
		Candidate:         "T_s(S_split)=s with uniform insertion into (1+s b1)(1+s b2)",
		NativeCertified:   false,
		CertificatePassed: false,
		Supports: []string{
			"CONDITIONAL_SUPPORT_REQUIRED_TRANSPORT_MAP_IS_NOW_EXPLICIT",
			"CONDITIONAL_SUPPORT_S_SPLIT_TO_B2_PARAMETER_MAP_HAS_AUGMENTED_CHAMBER_SOURCE_CANDIDATE",
			"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_WOUND_REDUCED_TO_AUGMENTED_CHAMBER_TO_B2_RESPONSE_MAP",
		},
		Failures: []string{
			"FAILED_ROUTE_NO_NATIVE_AUGMENTED_CHAMBER_TO_B2_RESPONSE_TRANSPORT_THEOREM",
			"FAILED_ROUTE_NO_NATIVE_S_SPLIT_TO_B2_RESPONSE_PARAMETER_MAP",
			"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
		},
	}
}

func DefaultResponseUse() BoundaryResponseUse {
	return BoundaryResponseUse{
		Response: "R_B(S_split)=(1+S_split b1)(1+S_split b2)-1",
		UniformFactors: []string{
			"1+S_split b1",
			"1+S_split b2",
		},
		AlphaLinear:    AlphaLinear,
		AlphaQuadratic: AlphaQuad,
		AlphaTotal:     AlphaLinear + AlphaQuad,
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_S_SPLIT_HAS_SCALAR_RESPONSE_PARAMETER_TYPE",
		"CONDITIONAL_SUPPORT_S_SPLIT_IS_DIMENSIONLESS_AND_CAN_MULTIPLY_BOUNDARY_GENERATORS",
		"CONDITIONAL_SUPPORT_S_SPLIT_IS_NOT_A_PROJECTOR_OR_TARGET_RANK",
		"CONDITIONAL_SUPPORT_S_SPLIT_HAS_PRIOR_AUGMENTED_CHAMBER_DEFECT_TRACE_ORIGIN",
		"CONDITIONAL_SUPPORT_S_SPLIT_IS_NOT_NEWLY_INSERTED_FOR_ALPHA_BRANCH",
		"CONDITIONAL_SUPPORT_S_SPLIT_ALREADY_COUPLES_TO_72_CHAMBER_RESPONSE_STRUCTURE",
		"CONDITIONAL_SUPPORT_S_SPLIT_SOURCE_IS_COMPATIBLE_WITH_BOUNDARY_AUGMENTED_H72_LANE",
		"CONDITIONAL_SUPPORT_BOUNDARY_PAIR_SYMMETRY_FORCES_UNIFORM_SCALAR_INSERTION",
		"CONDITIONAL_SUPPORT_S_SPLIT_IS_THE_ONLY_AVAILABLE_INHERITED_BOUNDARY_SPLIT_SCALAR",
		"CONDITIONAL_SUPPORT_S_SPLIT_UNIFORM_INSERTION_IS_SOURCE_TYPED",
		"CONDITIONAL_SUPPORT_S_SPLIT_SQUARED_TERM_ARISES_FROM_BOUNDARY_PAIR_MULTIPLICATION",
		"CONDITIONAL_SUPPORT_REQUIRED_TRANSPORT_MAP_IS_NOW_EXPLICIT",
		"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_WOUND_REDUCED_TO_AUGMENTED_CHAMBER_TO_B2_RESPONSE_MAP",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM",
		"FAILED_ROUTE_SCALAR_TYPE_COMPATIBILITY_NOT_NATIVE_TRANSPORT_THEOREM",
		"FAILED_ROUTE_AUGMENTED_CHAMBER_DEFECT_TRACE_ORIGIN_NOT_YET_NATIVE_B2_RESPONSE_TRANSPORT",
		"FAILED_ROUTE_UNIQUENESS_OF_S_SPLIT_AS_BOUNDARY_SCALAR_NOT_NATIVE_CERTIFIED",
		"FAILED_ROUTE_NO_NATIVE_AUGMENTED_CHAMBER_TO_B2_RESPONSE_TRANSPORT_THEOREM",
		"FAILED_ROUTE_NO_NATIVE_S_SPLIT_TO_B2_RESPONSE_PARAMETER_MAP",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func FormatOrigin(o ScalarOrigin) string {
	return fmt.Sprintf("%s: %s [%s]", o.Name, o.Expression, o.Status)
}

func FormatTransport(t TransportMap) string {
	return fmt.Sprintf("%s: %s -> %s; %s; native=%v certificateII=%v", t.Name, t.Domain, t.Codomain, t.Candidate, t.NativeCertified, t.CertificatePassed)
}

func FormatResponse(r BoundaryResponseUse) string {
	return fmt.Sprintf("%s; factors=%s; alpha_linear=%.19g alpha_quad=%.19g alpha_B=%.19g", r.Response, strings.Join(r.UniformFactors, ","), r.AlphaLinear, r.AlphaQuadratic, r.AlphaTotal)
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
	for _, item := range items {
		out = append(out, item...)
	}
	return out
}
