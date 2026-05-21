// Package generation2ssplitaddendsourcecircularitynativescalarlaneaudit implements
// Gate 945: S_split Addend Source, Circularity, and Native ScalarLane Audit.
//
// Gate 945 follows Gate 944's honesty result that S_split is H72-compatible
// and B2-descent-ready, but not native-sourced. It audits the expression
// S_split=(R_3-1)+lambda(Lambda_12), separates transport from addend origin,
// and blocks native R3 promotion when R_3-1 risks circular input use or when
// lambda(Lambda_12) remains bridge/history sourced.
package generation2ssplitaddendsourcecircularitynativescalarlaneaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE945-GENERATION2SSPLITADDENDSOURCECIRCULARITYNATIVESCALARLANEAUDIT"
	InheritedStatus = "R3_S_SPLIT_SOURCE_REMAINS_BRIDGE_HISTORY_SCALAR"
	Verdict         = "S_SPLIT_ADDENDS_AUDITED_R3_MINUS_ONE_AND_LAMBDA_LAMBDA12_REMAIN_BRIDGE_HISTORY_SCALARS_AND_R3_MINUS_ONE_ROUTE_IS_CIRCULAR_WITHOUT_INDEPENDENT_SOURCE"
	Classification  = "R3_S_SPLIT_ADDEND_SOURCE_AUDIT_BLOCKS_NATIVE_SCALAR_PROMOTION"
	ShortStatus     = "R3_S_SPLIT_NATIVE_SOURCE_BLOCKED_BY_ADDEND_ORIGIN_AND_CIRCULARITY"
	NextGate        = "NEXT_PRESSURE_GATE946_NONCIRCULAR_S_SPLIT_REPLACEMENT_AND_FINITE_SCALAR_PROXY_AUDIT"
)

const (
	Ssplit = 0.0012924448188162962
)

type SourceState string

const (
	SourceNative        SourceState = "NATIVE_H72_SCALAR"
	SourceBridgeHistory SourceState = "BRIDGE_HISTORY_SCALAR"
	SourceCircularRisk  SourceState = "CIRCULAR_INPUT_RISK"
)

type AddendAudit struct {
	Name                string
	ExpressionRole      string
	ScalarAddendTyped   bool
	NativeH72Scalar     bool
	BridgeHistoryScalar bool
	CircularRisk        bool
	RequiresCertificate string
	Supports            []string
	Failures            []string
}

type ScalarLaneAudit struct {
	LaneName          string
	CanonicalAddition bool
	BothAddendsNative bool
	SsplitNative      bool
	Status            string
	Supports          []string
	Failures          []string
}

type SourceTruthRow struct {
	R3MinusOne SourceState
	Lambda12   SourceState
	Outcome    string
	Current    bool
}

type CertificateIIStatus struct {
	TransportLayerStrong bool
	CentralH72Compatible bool
	AddendSourceNative   bool
	CertificatePassed    bool
	Status               string
	Supports             []string
	Failures             []string
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	Expression     string
	Value          float64
	Addends        []AddendAudit
	ScalarLane     ScalarLaneAudit
	TruthTable     []SourceTruthRow
	Certificate    CertificateIIStatus
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	addends := DefaultAddends()
	lane := DefaultScalarLane(addends)
	truth := DefaultTruthTable()
	cert := DefaultCertificateIIStatus()
	if len(addends) != 2 {
		return Analysis{}, fmt.Errorf("expected two S_split addends, got %d", len(addends))
	}
	for _, a := range addends {
		if !a.ScalarAddendTyped {
			return Analysis{}, fmt.Errorf("addend %s must be scalar typed", a.Name)
		}
		if a.NativeH72Scalar {
			return Analysis{}, fmt.Errorf("Gate 945 must not native-certify addend %s", a.Name)
		}
	}
	if lane.SsplitNative || lane.BothAddendsNative {
		return Analysis{}, fmt.Errorf("S_split scalar lane must remain non-native: %#v", lane)
	}
	if !containsCurrentCircularOrBridgeRow(truth) {
		return Analysis{}, fmt.Errorf("truth table must mark current bridge/circular row")
	}
	if cert.AddendSourceNative || cert.CertificatePassed {
		return Analysis{}, fmt.Errorf("Certificate II must remain blocked by addend source: %#v", cert)
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Expression:     "S_split=(R_3-1)+lambda(Lambda_12)",
		Value:          Ssplit,
		Addends:        addends,
		ScalarLane:     lane,
		TruthTable:     truth,
		Certificate:    cert,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 945 prevents a hidden circular native-R3 proof: transport of S_split to B2 is already strongly supported, but the addends R_3-1 and lambda(Lambda_12) are not native H72 scalars. R_3-1 is especially dangerous unless it has an independent noncircular source, so Certificate II remains blocked by addend origin and circularity.",
	}, nil
}

func DefaultAddends() []AddendAudit {
	return []AddendAudit{
		{
			Name:                "R_3-1",
			ExpressionRole:      "dimensionless R3 deviation addend in S_split",
			ScalarAddendTyped:   true,
			NativeH72Scalar:     false,
			BridgeHistoryScalar: true,
			CircularRisk:        true,
			RequiresCertificate: "independent noncircular native H72 scalar derivation of R_3-1",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_MINUS_ONE_HAS_SCALAR_DEVIATION_TYPE",
				"CONDITIONAL_SUPPORT_R3_MINUS_ONE_CAN_BE_USED_ONLY_IF_PRIOR_INDEPENDENTLY_CERTIFIED",
				"CONDITIONAL_SUPPORT_R3_MINUS_ONE_SOURCE_REQUIRES_NONCIRCULAR_ORIGIN",
			},
			Failures: []string{
				"FAILED_ROUTE_R3_MINUS_ONE_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
				"FAILED_ROUTE_R3_MINUS_ONE_AS_INPUT_TO_R3_PROMOTION_IS_POTENTIALLY_CIRCULAR",
				"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_R3_DEVIATION_SOURCE_CERTIFIED",
			},
		},
		{
			Name:                "lambda(Lambda_12)",
			ExpressionRole:      "history/scale scalar addend in S_split",
			ScalarAddendTyped:   true,
			NativeH72Scalar:     false,
			BridgeHistoryScalar: true,
			CircularRisk:        false,
			RequiresCertificate: "finite H72 chamber scalar certificate for lambda(Lambda_12)",
			Supports: []string{
				"CONDITIONAL_SUPPORT_LAMBDA_LAMBDA12_HAS_SCALAR_ADDEND_TYPE",
				"CONDITIONAL_SUPPORT_LAMBDA_LAMBDA12_CAN_ADD_TO_R3_MINUS_ONE_IF_BOTH_LIVE_IN_SAME_SCALAR_LANE",
				"CONDITIONAL_SUPPORT_LAMBDA_LAMBDA12_SOURCE_REQUIRES_H72_SCALAR_CERTIFICATE",
			},
			Failures: []string{
				"FAILED_ROUTE_LAMBDA_LAMBDA12_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
				"FAILED_ROUTE_NO_FINITE_H72_CHAMBER_SOURCE_FOR_LAMBDA_LAMBDA12",
				"FAILED_ROUTE_LAMBDA_LAMBDA12_REMAINS_BRIDGE_HISTORY_SCALAR",
			},
		},
	}
}

func DefaultScalarLane(addends []AddendAudit) ScalarLaneAudit {
	bothNative := true
	for _, a := range addends {
		bothNative = bothNative && a.NativeH72Scalar
	}
	return ScalarLaneAudit{
		LaneName:          "Scal(H72) central scalar response lane",
		CanonicalAddition: true,
		BothAddendsNative: bothNative,
		SsplitNative:      bothNative,
		Status:            "CANONICAL_ADDITION_AVAILABLE_BUT_ADDEND_SOURCES_NOT_NATIVE",
		Supports: []string{
			"CONDITIONAL_SUPPORT_S_SPLIT_ADDITION_LAW_IS_CANONICAL_IF_BOTH_ADDENDS_ARE_H72_SCALARS",
			"CONDITIONAL_SUPPORT_S_SPLIT_CAN_BE_H72_SCALAR_ONLY_IF_R3_MINUS_ONE_AND_LAMBDA_LAMBDA12_ARE_H72_SCALARS",
			"CONDITIONAL_SUPPORT_NATIVE_S_SPLIT_REQUIRES_COMMON_H72_SCALAR_LANE_CERTIFICATE",
		},
		Failures: []string{
			"FAILED_ROUTE_NO_COMMON_NATIVE_H72_SCALAR_LANE_CERTIFICATE_FOR_BOTH_ADDENDS",
			"FAILED_ROUTE_CANONICAL_ADDITION_DOES_NOT_NATIVE_CERTIFY_ADDEND_SOURCES",
			"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_IF_EITHER_ADDEND_IS_BRIDGE_HISTORY",
		},
	}
}

func DefaultTruthTable() []SourceTruthRow {
	return []SourceTruthRow{
		{R3MinusOne: SourceNative, Lambda12: SourceNative, Outcome: "S_split native H72 scalar possible", Current: false},
		{R3MinusOne: SourceNative, Lambda12: SourceBridgeHistory, Outcome: "S_split bridge scalar", Current: false},
		{R3MinusOne: SourceBridgeHistory, Lambda12: SourceNative, Outcome: "S_split bridge scalar", Current: false},
		{R3MinusOne: SourceBridgeHistory, Lambda12: SourceBridgeHistory, Outcome: "S_split bridge/history scalar", Current: true},
		{R3MinusOne: SourceCircularRisk, Lambda12: SourceBridgeHistory, Outcome: "S_split cannot support native R3 promotion without independent source", Current: true},
	}
}

func DefaultCertificateIIStatus() CertificateIIStatus {
	return CertificateIIStatus{
		TransportLayerStrong: true,
		CentralH72Compatible: true,
		AddendSourceNative:   false,
		CertificatePassed:    false,
		Status:               "TRANSPORT_STRONG_ADDEND_SOURCE_BLOCKED",
		Supports: []string{
			"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_LAYER_ALREADY_STRONGLY_SUPPORTED",
			"CONDITIONAL_SUPPORT_NATIVE_S_SPLIT_REQUIRES_COMMON_H72_SCALAR_LANE_CERTIFICATE",
			"CONDITIONAL_SUPPORT_CERTIFICATE_II_NOW_REDUCED_TO_NATIVE_ADDEND_SOURCE",
		},
		Failures: []string{
			"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
			"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED_BECAUSE_ADDEND_SOURCES_NOT_NATIVE",
			"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_LAYER_ALREADY_STRONGLY_SUPPORTED",
		"CONDITIONAL_SUPPORT_R3_MINUS_ONE_HAS_SCALAR_DEVIATION_TYPE",
		"CONDITIONAL_SUPPORT_LAMBDA_LAMBDA12_HAS_SCALAR_ADDEND_TYPE",
		"CONDITIONAL_SUPPORT_S_SPLIT_ADDITION_LAW_IS_CANONICAL_IF_BOTH_ADDENDS_ARE_H72_SCALARS",
		"CONDITIONAL_SUPPORT_NATIVE_S_SPLIT_REQUIRES_COMMON_H72_SCALAR_LANE_CERTIFICATE",
		"CONDITIONAL_SUPPORT_CERTIFICATE_II_NOW_REDUCED_TO_NATIVE_ADDEND_SOURCE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_R3_MINUS_ONE_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
		"FAILED_ROUTE_R3_MINUS_ONE_AS_INPUT_TO_R3_PROMOTION_IS_POTENTIALLY_CIRCULAR",
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_R3_DEVIATION_SOURCE_CERTIFIED",
		"FAILED_ROUTE_LAMBDA_LAMBDA12_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
		"FAILED_ROUTE_NO_FINITE_H72_CHAMBER_SOURCE_FOR_LAMBDA_LAMBDA12",
		"FAILED_ROUTE_NO_COMMON_NATIVE_H72_SCALAR_LANE_CERTIFICATE_FOR_BOTH_ADDENDS",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED_BECAUSE_ADDEND_SOURCES_NOT_NATIVE",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func FormatAddend(a AddendAudit) string {
	return fmt.Sprintf("%s: role=%s scalar=%v nativeH72=%v bridgeHistory=%v circularRisk=%v requires=%s", a.Name, a.ExpressionRole, a.ScalarAddendTyped, a.NativeH72Scalar, a.BridgeHistoryScalar, a.CircularRisk, a.RequiresCertificate)
}

func FormatScalarLane(l ScalarLaneAudit) string {
	return fmt.Sprintf("%s: addition=%v bothAddendsNative=%v ssplitNative=%v status=%s", l.LaneName, l.CanonicalAddition, l.BothAddendsNative, l.SsplitNative, l.Status)
}

func FormatTruthRow(r SourceTruthRow) string {
	flag := ""
	if r.Current {
		flag = " current"
	}
	return fmt.Sprintf("R3-1=%s lambda=%s -> %s%s", r.R3MinusOne, r.Lambda12, r.Outcome, flag)
}

func FormatCertificate(c CertificateIIStatus) string {
	return fmt.Sprintf("%s; transport=%v centralH72=%v addendNative=%v certificatePassed=%v", c.Status, c.TransportLayerStrong, c.CentralH72Compatible, c.AddendSourceNative, c.CertificatePassed)
}

func allAddendSupports(addends []AddendAudit) []string {
	var out []string
	for _, a := range addends {
		out = append(out, a.Supports...)
	}
	return out
}

func allAddendFailures(addends []AddendAudit) []string {
	var out []string
	for _, a := range addends {
		out = append(out, a.Failures...)
	}
	return out
}

func containsCurrentCircularOrBridgeRow(rows []SourceTruthRow) bool {
	bridge := false
	circular := false
	for _, r := range rows {
		if r.Current && r.R3MinusOne == SourceBridgeHistory && r.Lambda12 == SourceBridgeHistory {
			bridge = true
		}
		if r.Current && r.R3MinusOne == SourceCircularRisk {
			circular = true
		}
	}
	return bridge && circular
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
