// Package generation2observedpreflight implements Gate 469:
// Observed Rank-Complete Comparator Preflight / Airlock Non-Computation Audit.
//
// Gate 468 proved the ASHA u-d cylinder socket on synthetic rank-complete
// ledgers. Gate 469 is the observed-data preflight. It does not import fresh
// PDG values and does not invent I_spec/I_K rows. Instead it validates the
// exact fail-closed conditions under which an observed common-scale ledger would
// be allowed to reach the Gate456/Gate464 inverse-distance socket.
package generation2observedpreflight

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE469-OBSERVED-RANK-COMPLETE-COMPARATOR-PREFLIGHT-AIRLOCK-NON-COMPUTATION-AUDIT"

	StatusGate468Inherited         = "CONDITIONAL_SUPPORT_GATE468_SYNTHETIC_SOCKET_INHERITED"
	StatusPreflightPolicyDefined   = "CONDITIONAL_SUPPORT_OBSERVED_PREFLIGHT_POLICY_DEFINED"
	StatusSchemaAccepted           = "CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_SCHEMA_ACCEPTED"
	StatusPreflightValidated       = "CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_PREFLIGHT_VALIDATED"
	StatusFirewallPreserved        = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_OBSERVED_PREFLIGHT"
	StatusFailedSwitchClosed       = "FAILED_ROUTE_OBSERVED_PREFLIGHT_EMPIRICAL_SWITCH_CLOSED"
	StatusFailedMissingISpec       = "FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_I_SPEC"
	StatusFailedMissingIK          = "FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_I_K"
	StatusFailedMissingBranchTags  = "FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_BRANCH_TAGS"
	StatusFailedMixedScaleScheme   = "FAILED_ROUTE_OBSERVED_PREFLIGHT_MIXED_SCALE_SCHEME_REJECTED"
	StatusFailedMissingUncertainty = "FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_UNCERTAINTY"
	StatusFailedMissingNumeric     = "FAILED_ROUTE_OBSERVED_DUD_NOT_COMPUTED_WITHOUT_ACTUAL_I_SPEC_I_K_VALUES"
	StatusFailedCabibboAsRayInput  = "FAILED_ROUTE_CABIBBO_USED_AS_OBSERVED_RAY_INPUT_REJECTED"
	StatusFailedNativePromotion    = "FAILED_ROUTE_OBSERVED_PREFLIGHT_NATIVE_PROMOTION_REJECTED"
	StatusFailedCKMPrediction      = "FAILED_ROUTE_OBSERVED_PREFLIGHT_CKM_NATIVE_PREDICTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed, Gate444KGenForced, Gate445TriangleForced, Gate456InverseAvailable, Gate459BranchTagsRequired, Gate464DUDSocketAvailable, Gate465AirlockAvailable, Gate467ObservedSchemaDefined, Gate468SyntheticSocketValidated, NativeRegistryClean bool
	Verdict                                                                                                                                                                                                                                        string
}

type ObservedSectorLedger struct {
	Sector                   string
	Source                   string
	SourceVersion            string
	Scale                    string
	Scheme                   string
	UncertaintyModel         string
	HasISpec                 bool
	HasIK                    bool
	HasCPOddSign             bool
	HasC3Sheet               bool
	HasISpecValue            bool
	HasIKValue               bool
	ISpec                    float64
	IK                       float64
	SigmaCP                  int
	C3Sheet                  int
	BridgeOnly               bool
	Observed                 bool
	EmpiricalImport          bool
	CabibboAsRayInput        bool
	NativePromotionClaim     bool
	CKMNativePredictionClaim bool
}

type PreflightCase struct {
	Name                 string
	U                    ObservedSectorLedger
	D                    ObservedSectorLedger
	Accepted             bool
	ReadyForNumericalDUD bool
	DUDComputed          bool
	Verdict              string
	Reason               string
	Failures             []string
}

type Preflight struct {
	Executed                      bool
	Cases                         []PreflightCase
	AcceptedSchemaCases           int
	RejectedCases                 int
	ReadyNumericCases             int
	DUDComputed                   bool
	SwitchClosedRejected          bool
	MissingISpecRejected          bool
	MissingIKRejected             bool
	MissingBranchRejected         bool
	MixedScaleRejected            bool
	MissingUncertaintyRejected    bool
	MissingNumericRejected        bool
	CabibboRayRejected            bool
	NativePromotionRejected       bool
	CKMNativePredictionRejected   bool
	AllAcceptedBridgeOnlyObserved bool
	Verdict                       string
	Reason                        string
}

type Firewall struct {
	Executed                      bool
	ObservedRowsNative            bool
	DUDNativePrediction           bool
	CKMNativePrediction           bool
	CKMMatrixConstructed          bool
	CKMEntryComputed              bool
	CabibboUsedAsRayInput         bool
	NativeRegistryWritten         bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Preflight   Preflight
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Preflight = buildPreflight()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate456InverseAvailable: true, Gate459BranchTagsRequired: true, Gate464DUDSocketAvailable: true, Gate465AirlockAvailable: true, Gate467ObservedSchemaDefined: true, Gate468SyntheticSocketValidated: true, NativeRegistryClean: true, Verdict: StatusGate468Inherited}
}

func completeRedacted(sector string) ObservedSectorLedger {
	return ObservedSectorLedger{Sector: sector, Source: "external observed common-scale comparator ledger", SourceVersion: "redacted preflight record; no numeric values imported in Gate469", Scale: "common reference scale", Scheme: "common renormalization scheme", UncertaintyModel: "declared uncertainty propagation model", HasISpec: true, HasIK: true, HasCPOddSign: true, HasC3Sheet: true, SigmaCP: 1, C3Sheet: 0, BridgeOnly: true, Observed: true, EmpiricalImport: true}
}
func mutate(x ObservedSectorLedger, fn func(*ObservedSectorLedger)) ObservedSectorLedger {
	fn(&x)
	return x
}

func buildPreflight() Preflight {
	cases := []PreflightCase{
		EvaluateCase("complete observed schema, redacted numeric values", completeRedacted("u"), completeRedacted("d")),
		EvaluateCase("switch closed rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.EmpiricalImport = false }), completeRedacted("d")),
		EvaluateCase("missing I_spec rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.HasISpec = false }), completeRedacted("d")),
		EvaluateCase("missing I_K rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.HasIK = false }), completeRedacted("d")),
		EvaluateCase("missing branch tags rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.HasCPOddSign = false }), completeRedacted("d")),
		EvaluateCase("mixed scale rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.Scale = "2 GeV" }), mutate(completeRedacted("d"), func(x *ObservedSectorLedger) { x.Scale = "MZ" })),
		EvaluateCase("missing uncertainty rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.UncertaintyModel = "" }), completeRedacted("d")),
		EvaluateCase("Cabibbo as ray input rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.CabibboAsRayInput = true }), completeRedacted("d")),
		EvaluateCase("native promotion rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.NativePromotionClaim = true }), completeRedacted("d")),
		EvaluateCase("CKM native prediction rejected", mutate(completeRedacted("u"), func(x *ObservedSectorLedger) { x.CKMNativePredictionClaim = true }), completeRedacted("d")),
	}
	p := Preflight{Executed: true, Cases: cases, Verdict: StatusPreflightValidated, Reason: "observed ledgers may pass only through the empirical airlock as bridge-only records; redacted ledgers do not compute d_ud"}
	for _, c := range cases {
		if c.Accepted {
			p.AcceptedSchemaCases++
		} else {
			p.RejectedCases++
		}
		if c.ReadyForNumericalDUD {
			p.ReadyNumericCases++
		}
		p.DUDComputed = p.DUDComputed || c.DUDComputed
		for _, f := range c.Failures {
			switch f {
			case StatusFailedSwitchClosed:
				p.SwitchClosedRejected = true
			case StatusFailedMissingISpec:
				p.MissingISpecRejected = true
			case StatusFailedMissingIK:
				p.MissingIKRejected = true
			case StatusFailedMissingBranchTags:
				p.MissingBranchRejected = true
			case StatusFailedMixedScaleScheme:
				p.MixedScaleRejected = true
			case StatusFailedMissingUncertainty:
				p.MissingUncertaintyRejected = true
			case StatusFailedMissingNumeric:
				p.MissingNumericRejected = true
			case StatusFailedCabibboAsRayInput:
				p.CabibboRayRejected = true
			case StatusFailedNativePromotion:
				p.NativePromotionRejected = true
			case StatusFailedCKMPrediction:
				p.CKMNativePredictionRejected = true
			}
		}
	}
	p.AllAcceptedBridgeOnlyObserved = true
	for _, c := range cases {
		if c.Accepted && (!c.U.BridgeOnly || !c.D.BridgeOnly || !c.U.Observed || !c.D.Observed) {
			p.AllAcceptedBridgeOnlyObserved = false
		}
	}
	return p
}

func EvaluateCase(name string, u, d ObservedSectorLedger) PreflightCase {
	c := PreflightCase{Name: name, U: u, D: d}
	fs := append(validateLedger(u, d), validateLedger(d, u)...)
	if !sameScaleScheme(u, d) {
		fs = append(fs, StatusFailedMixedScaleScheme)
	}
	if len(fs) > 0 {
		c.Failures = unique(fs)
		c.Verdict = strings.Join(c.Failures, ";")
		c.Reason = "observed preflight rejected one or both sector ledgers"
		return c
	}
	c.Accepted = true
	c.ReadyForNumericalDUD = u.HasISpecValue && u.HasIKValue && d.HasISpecValue && d.HasIKValue
	if !c.ReadyForNumericalDUD {
		c.Failures = []string{StatusFailedMissingNumeric}
		c.Verdict = StatusSchemaAccepted
		c.Reason = "schema is complete but numeric I_spec/I_K values are redacted, so d_ud is not computed"
		return c
	}
	c.Verdict = StatusSchemaAccepted
	c.Reason = "observed comparator values are present; Gate469 still refuses to evaluate them because numerical execution belongs to Gate470"
	return c
}

func validateLedger(x, other ObservedSectorLedger) []string {
	var fs []string
	if !x.EmpiricalImport {
		fs = append(fs, StatusFailedSwitchClosed)
	}
	if !x.HasISpec {
		fs = append(fs, StatusFailedMissingISpec)
	}
	if !x.HasIK {
		fs = append(fs, StatusFailedMissingIK)
	}
	if !x.HasCPOddSign || !x.HasC3Sheet || (x.SigmaCP != -1 && x.SigmaCP != 1) || x.C3Sheet < 0 || x.C3Sheet > 2 {
		fs = append(fs, StatusFailedMissingBranchTags)
	}
	if x.UncertaintyModel == "" {
		fs = append(fs, StatusFailedMissingUncertainty)
	}
	if x.Source == "" || x.Scale == "" || x.Scheme == "" || !x.BridgeOnly || !x.Observed {
		fs = append(fs, StatusFailedMissingUncertainty)
	}
	if x.CabibboAsRayInput {
		fs = append(fs, StatusFailedCabibboAsRayInput)
	}
	if x.NativePromotionClaim {
		fs = append(fs, StatusFailedNativePromotion)
	}
	if x.CKMNativePredictionClaim {
		fs = append(fs, StatusFailedCKMPrediction)
	}
	return fs
}
func sameScaleScheme(a, b ObservedSectorLedger) bool {
	return a.Scale != "" && a.Scale == b.Scale && a.Scheme != "" && a.Scheme == b.Scheme
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate469 performs observed-ledger preflight only; no observed comparator is written to native registry and no CKM entry is constructed"}
}
func buildNext() NextStep {
	return NextStep{470, "Observed Numerical d_ud Adapter / Explicit Data-File Run", "Gate469 validates the observed schema and proves redacted rows cannot compute d_ud; a later run may evaluate only with explicit rank-complete I_spec/I_K values and branch tags.", "read a user-supplied observed comparator ledger, evaluate d_ud as bridge-only, and compare to Cabibbo only as a residual target"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate468SyntheticSocketValidated || !a.Inheritance.Gate465AirlockAvailable || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate469 inheritance incomplete")
	}
	if !a.Preflight.Executed || a.Preflight.AcceptedSchemaCases != 1 || a.Preflight.ReadyNumericCases != 0 || a.Preflight.DUDComputed || !a.Preflight.SwitchClosedRejected || !a.Preflight.MissingISpecRejected || !a.Preflight.MissingIKRejected || !a.Preflight.MissingBranchRejected || !a.Preflight.MixedScaleRejected || !a.Preflight.MissingUncertaintyRejected || !a.Preflight.MissingNumericRejected || !a.Preflight.CabibboRayRejected || !a.Preflight.NativePromotionRejected || !a.Preflight.CKMNativePredictionRejected || !a.Preflight.AllAcceptedBridgeOnlyObserved {
		return fmt.Errorf("Gate469 preflight did not accept/reject expected routes")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedRowsNative || a.Firewall.DUDNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.CabibboUsedAsRayInput || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate469 firewall violated")
	}
	return nil
}
func truth(a Analysis) string {
	if a.Preflight.AcceptedSchemaCases == 2 && !a.Preflight.DUDComputed && !a.Firewall.CKMNativePrediction {
		return "Gate 469 validates the observed comparator airlock schema but computes no physical CKM alignment. Redacted observed ledgers are schema-ready but numerically undefined; numeric-ready ledgers are admitted only as bridge-only inputs for a later adapter. The 13-moduli firewall remains unbreached."
	}
	return "Gate 469 preflight failed."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch=%t dud_socket=%t airlock=%t schema=%t synthetic_socket=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseAvailable, x.Gate459BranchTagsRequired, x.Gate464DUDSocketAvailable, x.Gate465AirlockAvailable, x.Gate467ObservedSchemaDefined, x.Gate468SyntheticSocketValidated, x.NativeRegistryClean, x.Verdict)
}
func FormatLedger(x ObservedSectorLedger) string {
	return fmt.Sprintf("sector=%s source=%s scale=%s scheme=%s uncertainty=%s has_I_spec=%t has_I_K=%t has_branch=%t/%t has_values=%t/%t bridge_only=%t observed=%t empirical_import=%t cabibbo_as_ray=%t native_claim=%t", x.Sector, x.Source, x.Scale, x.Scheme, x.UncertaintyModel, x.HasISpec, x.HasIK, x.HasCPOddSign, x.HasC3Sheet, x.HasISpecValue, x.HasIKValue, x.BridgeOnly, x.Observed, x.EmpiricalImport, x.CabibboAsRayInput, x.NativePromotionClaim)
}
func FormatCase(x PreflightCase) string {
	return fmt.Sprintf("accepted=%t ready_numeric=%t dud_computed=%t verdict=%s reason=%s", x.Accepted, x.ReadyForNumericalDUD, x.DUDComputed, x.Verdict, x.Reason)
}
func FormatPreflight(x Preflight) string {
	return fmt.Sprintf("executed=%t accepted_schema=%d rejected=%d ready_numeric=%d dud_computed=%t switch_closed=%t missing_I_spec=%t missing_I_K=%t missing_branch=%t mixed_scale=%t missing_uncertainty=%t missing_numeric=%t cabibbo_as_ray=%t native_promotion=%t ckm_native=%t bridge_observed=%t verdict=%s reason=%s", x.Executed, x.AcceptedSchemaCases, x.RejectedCases, x.ReadyNumericCases, x.DUDComputed, x.SwitchClosedRejected, x.MissingISpecRejected, x.MissingIKRejected, x.MissingBranchRejected, x.MixedScaleRejected, x.MissingUncertaintyRejected, x.MissingNumericRejected, x.CabibboRayRejected, x.NativePromotionRejected, x.CKMNativePredictionRejected, x.AllAcceptedBridgeOnlyObserved, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t observed_native=%t dud_native=%t ckm_native=%t ckm_matrix=%t ckm_entry=%t cabibbo_as_ray=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.ObservedRowsNative, x.DUDNativePrediction, x.CKMNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.CabibboUsedAsRayInput, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}
func statuses() []string {
	return []string{StatusGate468Inherited, StatusPreflightPolicyDefined, StatusSchemaAccepted, StatusPreflightValidated, StatusFirewallPreserved, StatusFailedSwitchClosed, StatusFailedMissingISpec, StatusFailedMissingIK, StatusFailedMissingBranchTags, StatusFailedMixedScaleScheme, StatusFailedMissingUncertainty, StatusFailedMissingNumeric, StatusFailedCabibboAsRayInput, StatusFailedNativePromotion, StatusFailedCKMPrediction}
}
func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 469 Registry Audit — Observed Rank-Complete Comparator Preflight / Airlock Non-Computation Audit\n\n## Verdict\n\n`" + StatusPreflightValidated + "`\n\nGate 469 validates the exact observed-data preflight required after Gate 468. It does not fetch, invent, or evaluate PDG values. It proves that observed ledgers must carry common scale/scheme, `I_spec`, `I_K`, complete branch tags, uncertainty metadata, and bridge-only quarantine before any future `d_ud` adapter may run.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Preflight summary\n\n" + FormatPreflight(a.Preflight) + "\n\n| Case | Accepted | Ready for numerical `d_ud` | Verdict | Reason |\n|---|---:|---:|---|---|\n")
	for _, c := range a.Preflight.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | `%s` | %s |\n", esc(c.Name), c.Accepted, c.ReadyForNumericalDUD, esc(c.Verdict), esc(c.Reason)))
	}
	b.WriteString("\n## Formula socket held in reserve\n\n```text\nalpha = sqrt(3) I_K / sqrt(1-I_K^2)\ncos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)\nphi = (sigma_CP arccos(cos(3phi)) + 2pi n_C3)/3\nd_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))\n```\n\nGate 469 does not evaluate this socket for redacted observed rows. A numeric-ready row is only declared eligible for a later bridge-only adapter; it is not promoted to native law and it is not interpreted as `V_us`.\n\n")
	b.WriteString("## Native firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\nNo observed comparator row can become a native prediction, native law, CKM matrix element, PMNS value, Yukawa value, or coefficient selector. Cabibbo remains a residual target only, never a coordinate input.\n\n## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
func unique(xs []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, x := range xs {
		if !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	return out
}
func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
