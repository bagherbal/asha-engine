// Package generation2observedelectroweakpreflight implements Gate 506:
// Observed Electroweak Comparator Airlock Preflight.
//
// Gate 505 proved that the electroweak bridge adapter can execute on explicit
// fake inputs without native promotion.  Gate 506 is the observed-data airlock
// before any real electroweak comparator may run.  It defines the required
// metadata, validates fail-closed cases, and deliberately imports no observed
// numerical values.  It is a preflight theorem, not a physics prediction.
package generation2observedelectroweakpreflight

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticelectroweakmatchingadapter"
)

const (
	AuditID = "GATE506-OBSERVED-ELECTROWEAK-COMPARATOR-AIRLOCK-PREFLIGHT-AUDIT"

	StatusGate505SyntheticAdapterInherited = "CONDITIONAL_SUPPORT_GATE505_SYNTHETIC_ADAPTER_INHERITED"
	StatusObservedEWPreflightPolicyDefined = "CONDITIONAL_SUPPORT_OBSERVED_ELECTROWEAK_PREFLIGHT_POLICY_DEFINED"
	StatusObservedEWSchemaAccepted         = "CONDITIONAL_SUPPORT_OBSERVED_ELECTROWEAK_SCHEMA_ACCEPTED"
	StatusObservedEWPreflightValidated     = "CONDITIONAL_SUPPORT_OBSERVED_ELECTROWEAK_PREFLIGHT_VALIDATED"
	StatusObservedEWBridgeAirlockAccepted  = "CONDITIONAL_SUPPORT_OBSERVED_ELECTROWEAK_BRIDGE_AIRLOCK_ACCEPTED"
	StatusObservedEWNoNumericImport        = "CONDITIONAL_SUPPORT_NO_OBSERVED_ELECTROWEAK_NUMBERS_IMPORTED"
	StatusObservedEWNoAdapterExecuted      = "CONDITIONAL_SUPPORT_OBSERVED_ELECTROWEAK_ADAPTER_NOT_EXECUTED_BY_DEFAULT"
	StatusGate507ObservedEWAdapterRedirect = "CONDITIONAL_SUPPORT_GATE507_OBSERVED_ELECTROWEAK_ADAPTER_FILE_RUN_REDIRECT_DEFINED"

	StatusFailedSwitchClosed                  = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_EMPIRICAL_SWITCH_CLOSED"
	StatusFailedMissingVEV                    = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_MISSING_VEV_ROW"
	StatusFailedMissingGaugeCoupling          = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_MISSING_GAUGE_COUPLING_ROWS"
	StatusFailedMissingScaleScheme            = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_MISSING_SCALE_OR_SCHEME_METADATA"
	StatusFailedMissingSourceUncertainty      = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_MISSING_SOURCE_OR_UNCERTAINTY_METADATA"
	StatusFailedObservedMassAsNativeInput     = "FAILED_ROUTE_OBSERVED_WZ_MASSES_USED_AS_NATIVE_INPUT_REJECTED"
	StatusFailedWeakAngleNativePromotion      = "FAILED_ROUTE_OBSERVED_WEAK_ANGLE_NATIVE_PROMOTION_REJECTED"
	StatusFailedKappaPromotion                = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_KAPPA_PROMOTION_REJECTED"
	StatusFailedNativePromotion               = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_NATIVE_PROMOTION_REJECTED"
	StatusFailedNumericalAdapterNotRun        = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_NUMERICAL_ADAPTER_NOT_RUN_IN_PREFLIGHT"
	StatusFailedNoNativeElectroweakPrediction = "FAILED_ROUTE_OBSERVED_ELECTROWEAK_PREFLIGHT_NO_NATIVE_PREDICTION"

	StatusFirewallObservedNumbersNotImported = "FIREWALL_PRESERVED_OBSERVED_ELECTROWEAK_NUMBERS_NOT_IMPORTED"
	StatusFirewallNativeWriteBlocked         = "FIREWALL_BLOCKED_OBSERVED_ELECTROWEAK_NATIVE_REGISTRY_WRITE"
)

type Inheritance struct {
	Executed                    bool
	Gate505AuditDefined         bool
	SyntheticAdapterExecuted    bool
	SyntheticOnly               bool
	Gate505ObservedDataImported bool
	Gate505NativeDataImported   bool
	Gate505NativeWriteBlocked   bool
	Gate506RedirectDefined      bool
	Verdict                     string
	Reason                      string
}

type ObservedRow struct {
	Name                 string
	Present              bool
	ValueProvided        bool
	Source               string
	SourceVersion        string
	Scale                string
	Scheme               string
	Uncertainty          string
	BridgeOnly           bool
	Observed             bool
	EmpiricalImport      bool
	NativePromotionClaim bool
	NativeInputClaim     bool
	KappaPromotionClaim  bool
	FormulaOutputOnly    bool
}

type ObservedLedger struct {
	Name                 string
	Rows                 []ObservedRow
	EmpiricalImport      bool
	BridgeOnly           bool
	NativeRegistryWrite  bool
	NumericalAdapterRun  bool
	ObservedValuesLoaded bool
}

type PreflightCase struct {
	Name                     string
	Ledger                   ObservedLedger
	AcceptedSchema           bool
	ReadyForNumericalAdapter bool
	NumericalAdapterRun      bool
	ObservedNumbersImported  bool
	Verdict                  string
	Reason                   string
	Failures                 []string
}

type Preflight struct {
	Executed                          bool
	Cases                             []PreflightCase
	AcceptedSchemaCases               int
	RejectedCases                     int
	ReadyForNumericalAdapterCases     int
	NumericalAdapterRun               bool
	ObservedNumbersImported           bool
	SwitchClosedRejected              bool
	MissingVEVRejected                bool
	MissingGaugeCouplingRejected      bool
	MissingScaleSchemeRejected        bool
	MissingSourceUncertaintyRejected  bool
	ObservedMassAsNativeInputRejected bool
	WeakAngleNativePromotionRejected  bool
	KappaPromotionRejected            bool
	NativePromotionRejected           bool
	AllAcceptedBridgeOnlyObserved     bool
	Verdict                           string
	Reason                            string
}

type Firewall struct {
	Executed                        bool
	ObservedVEVImported             bool
	ObservedGaugeCouplingsImported  bool
	ObservedWeakAngleImported       bool
	ObservedWMassImported           bool
	ObservedZMassImported           bool
	ObservedNumbersImported         bool
	NumericalAdapterExecuted        bool
	NativeVEVWritten                bool
	NativeGaugeCouplingWritten      bool
	NativeWeakAngleWritten          bool
	NativeWZMassWritten             bool
	NativeKappaWritten              bool
	NativeRegistryWritten           bool
	NativeElectroweakPredictionMade bool
	Verdict                         string
	Reason                          string
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
	Preflight   Preflight
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
	g505, err := generation2syntheticelectroweakmatchingadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate505 synthetic electroweak adapter: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g505)
	a.Preflight = buildPreflight()
	a.Firewall = buildFirewall(a.Preflight)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g505 generation2syntheticelectroweakmatchingadapter.Analysis) Inheritance {
	nativeBlocked := !g505.Firewall.NativeVEVWritten && !g505.Firewall.NativeGaugeCouplingWritten && !g505.Firewall.NativeWeakAngleWritten && !g505.Firewall.NativeWZMassWritten && !g505.Firewall.NativeKappaWritten && !g505.Firewall.SyntheticOutputWrittenNative
	return Inheritance{
		Executed:                    true,
		Gate505AuditDefined:         true,
		SyntheticAdapterExecuted:    g505.Adapter.Executed,
		SyntheticOnly:               g505.Adapter.SyntheticOnly,
		Gate505ObservedDataImported: g505.Adapter.ObservedDataImported || g505.Firewall.ObservedVEVImported || g505.Firewall.ObservedGaugeCouplingsImported || g505.Firewall.ObservedWeakAngleImported || g505.Firewall.ObservedWMassImported || g505.Firewall.ObservedZMassImported,
		Gate505NativeDataImported:   g505.Adapter.NativeDataImported,
		Gate505NativeWriteBlocked:   nativeBlocked,
		Gate506RedirectDefined:      g505.Next.Gate == 506,
		Verdict:                     StatusGate505SyntheticAdapterInherited,
		Reason:                      "Gate505 ran only a fake v=2, g2=3, gY=4 dry-run and explicitly redirected to an observed electroweak comparator airlock before any real data may be used.",
	}
}

func completeRedactedLedger() ObservedLedger {
	rows := []ObservedRow{
		completeRow("Higgs vacuum expectation value v", true, false, false, true),
		completeRow("SU(2)_L gauge coupling g2", true, false, false, true),
		completeRow("U(1)_Y gauge coupling gY", true, false, false, true),
		completeRow("weak mixing angle sin^2(theta_W)", true, false, false, false),
		completeRow("W boson comparator mass", true, false, false, false),
		completeRow("Z boson comparator mass", true, false, false, false),
	}
	return ObservedLedger{Name: "redacted observed electroweak comparator schema", Rows: rows, EmpiricalImport: true, BridgeOnly: true, NativeRegistryWrite: false, NumericalAdapterRun: false, ObservedValuesLoaded: false}
}

func completeRow(name string, present, valueProvided, nativeInput, requiresExplicit bool) ObservedRow {
	_ = requiresExplicit
	return ObservedRow{Name: name, Present: present, ValueProvided: valueProvided, Source: "external observed electroweak reference ledger", SourceVersion: "redacted preflight record; no numerical values imported at Gate506", Scale: "declared electroweak reference scale", Scheme: "declared renormalization scheme", Uncertainty: "declared uncertainty model", BridgeOnly: true, Observed: true, EmpiricalImport: true, NativePromotionClaim: false, NativeInputClaim: nativeInput, KappaPromotionClaim: false, FormulaOutputOnly: !nativeInput}
}

func cloneLedger(l ObservedLedger) ObservedLedger {
	out := l
	out.Rows = append([]ObservedRow(nil), l.Rows...)
	return out
}

func mutateLedger(l ObservedLedger, fn func(*ObservedLedger)) ObservedLedger {
	l = cloneLedger(l)
	fn(&l)
	return l
}
func mutateRow(l ObservedLedger, name string, fn func(*ObservedRow)) ObservedLedger {
	l = cloneLedger(l)
	for i := range l.Rows {
		if l.Rows[i].Name == name {
			fn(&l.Rows[i])
		}
	}
	return l
}

func buildPreflight() Preflight {
	base := completeRedactedLedger()
	cases := []PreflightCase{
		EvaluateCase("complete redacted observed schema, no numeric values", base),
		EvaluateCase("empirical switch closed rejected", mutateLedger(base, func(l *ObservedLedger) { l.EmpiricalImport = false })),
		EvaluateCase("missing VEV row rejected", mutateRow(base, "Higgs vacuum expectation value v", func(r *ObservedRow) { r.Present = false })),
		EvaluateCase("missing SU2 coupling row rejected", mutateRow(base, "SU(2)_L gauge coupling g2", func(r *ObservedRow) { r.Present = false })),
		EvaluateCase("missing scale metadata rejected", mutateRow(base, "U(1)_Y gauge coupling gY", func(r *ObservedRow) { r.Scale = "" })),
		EvaluateCase("missing uncertainty metadata rejected", mutateRow(base, "W boson comparator mass", func(r *ObservedRow) { r.Uncertainty = "" })),
		EvaluateCase("observed W mass as native input rejected", mutateRow(base, "W boson comparator mass", func(r *ObservedRow) { r.NativeInputClaim = true })),
		EvaluateCase("weak angle native promotion rejected", mutateRow(base, "weak mixing angle sin^2(theta_W)", func(r *ObservedRow) { r.NativePromotionClaim = true })),
		EvaluateCase("kappa promotion rejected", mutateRow(base, "U(1)_Y gauge coupling gY", func(r *ObservedRow) { r.KappaPromotionClaim = true })),
		EvaluateCase("ledger native registry write rejected", mutateLedger(base, func(l *ObservedLedger) { l.NativeRegistryWrite = true })),
		EvaluateCase("numeric adapter run rejected in preflight", mutateLedger(base, func(l *ObservedLedger) { l.NumericalAdapterRun = true })),
	}
	p := Preflight{Executed: true, Cases: cases, Verdict: strings.Join([]string{StatusObservedEWPreflightPolicyDefined, StatusObservedEWSchemaAccepted, StatusObservedEWPreflightValidated, StatusObservedEWBridgeAirlockAccepted, StatusObservedEWNoNumericImport, StatusObservedEWNoAdapterExecuted}, ";"), Reason: "The preflight accepts only a redacted, bridge-only observed electroweak schema and rejects every missing-metadata or native-promotion route; it does not import numbers and does not run matching."}
	p.AllAcceptedBridgeOnlyObserved = true
	for _, c := range cases {
		if c.AcceptedSchema {
			p.AcceptedSchemaCases++
			if !c.Ledger.BridgeOnly {
				p.AllAcceptedBridgeOnlyObserved = false
			}
			for _, r := range c.Ledger.Rows {
				if r.Present && (!r.BridgeOnly || !r.Observed) {
					p.AllAcceptedBridgeOnlyObserved = false
				}
			}
		} else {
			p.RejectedCases++
		}
		if c.ReadyForNumericalAdapter {
			p.ReadyForNumericalAdapterCases++
		}
		p.NumericalAdapterRun = p.NumericalAdapterRun || c.NumericalAdapterRun
		p.ObservedNumbersImported = p.ObservedNumbersImported || c.ObservedNumbersImported
		for _, f := range c.Failures {
			switch f {
			case StatusFailedSwitchClosed:
				p.SwitchClosedRejected = true
			case StatusFailedMissingVEV:
				p.MissingVEVRejected = true
			case StatusFailedMissingGaugeCoupling:
				p.MissingGaugeCouplingRejected = true
			case StatusFailedMissingScaleScheme:
				p.MissingScaleSchemeRejected = true
			case StatusFailedMissingSourceUncertainty:
				p.MissingSourceUncertaintyRejected = true
			case StatusFailedObservedMassAsNativeInput:
				p.ObservedMassAsNativeInputRejected = true
			case StatusFailedWeakAngleNativePromotion:
				p.WeakAngleNativePromotionRejected = true
			case StatusFailedKappaPromotion:
				p.KappaPromotionRejected = true
			case StatusFailedNativePromotion:
				p.NativePromotionRejected = true
			}
		}
	}
	return p
}

func EvaluateCase(name string, l ObservedLedger) PreflightCase {
	c := PreflightCase{Name: name, Ledger: l}
	failures := []string{}
	if !l.EmpiricalImport {
		failures = append(failures, StatusFailedSwitchClosed)
	}
	if l.NativeRegistryWrite || !l.BridgeOnly {
		failures = append(failures, StatusFailedNativePromotion)
	}
	if !hasRow(l, "Higgs vacuum expectation value v") {
		failures = append(failures, StatusFailedMissingVEV)
	}
	if !hasRow(l, "SU(2)_L gauge coupling g2") || !hasRow(l, "U(1)_Y gauge coupling gY") {
		failures = append(failures, StatusFailedMissingGaugeCoupling)
	}
	for _, r := range l.Rows {
		if !r.Present {
			continue
		}
		if r.Source == "" || r.SourceVersion == "" || r.Uncertainty == "" {
			failures = appendUnique(failures, StatusFailedMissingSourceUncertainty)
		}
		if r.Scale == "" || r.Scheme == "" {
			failures = appendUnique(failures, StatusFailedMissingScaleScheme)
		}
		if r.NativePromotionClaim || !r.BridgeOnly {
			if strings.Contains(r.Name, "weak mixing") {
				failures = appendUnique(failures, StatusFailedWeakAngleNativePromotion)
			}
			failures = appendUnique(failures, StatusFailedNativePromotion)
		}
		if r.KappaPromotionClaim {
			failures = appendUnique(failures, StatusFailedKappaPromotion)
		}
		if r.NativeInputClaim && (strings.Contains(r.Name, "W boson") || strings.Contains(r.Name, "Z boson")) {
			failures = appendUnique(failures, StatusFailedObservedMassAsNativeInput)
		}
	}
	if l.NumericalAdapterRun {
		failures = appendUnique(failures, StatusFailedNumericalAdapterNotRun)
	}
	c.Failures = failures
	c.AcceptedSchema = len(failures) == 0
	c.ReadyForNumericalAdapter = c.AcceptedSchema && allRequiredValuesProvided(l)
	c.NumericalAdapterRun = false
	c.ObservedNumbersImported = false
	if c.AcceptedSchema {
		c.Verdict = StatusObservedEWSchemaAccepted
		c.Reason = "Redacted observed electroweak schema is complete and bridge-only, but no numeric values are loaded, so no adapter runs."
	} else {
		c.Verdict = strings.Join(failures, ";")
		c.Reason = "Fail-closed observed electroweak preflight rejected this ledger before any numerical matching."
	}
	return c
}

func hasRow(l ObservedLedger, name string) bool {
	for _, r := range l.Rows {
		if r.Name == name && r.Present {
			return true
		}
	}
	return false
}
func allRequiredValuesProvided(l ObservedLedger) bool {
	required := map[string]bool{"Higgs vacuum expectation value v": false, "SU(2)_L gauge coupling g2": false, "U(1)_Y gauge coupling gY": false}
	for _, r := range l.Rows {
		if _, ok := required[r.Name]; ok && r.Present && r.ValueProvided {
			required[r.Name] = true
		}
	}
	for _, ok := range required {
		if !ok {
			return false
		}
	}
	return true
}
func appendUnique(xs []string, x string) []string {
	for _, y := range xs {
		if y == x {
			return xs
		}
	}
	return append(xs, x)
}

func buildFirewall(p Preflight) Firewall {
	return Firewall{Executed: true, ObservedNumbersImported: p.ObservedNumbersImported, NumericalAdapterExecuted: p.NumericalAdapterRun, Verdict: strings.Join([]string{StatusFirewallObservedNumbersNotImported, StatusFirewallNativeWriteBlocked}, ";"), Reason: "Gate506 imports no observed numerical electroweak values, runs no observed adapter, and writes no electroweak scale, coupling, angle, kappa, or W/Z mass into the native registry."}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No native electroweak VEV, gauge coupling, weak angle, kappa, or W/Z mass theorem is admitted at Gate506."},
		BridgeEntries:        []string{"An observed electroweak comparator preflight schema is admitted as bridge-only: VEV, g2, gY, weak angle, W/Z comparator rows require explicit source, version, scale, scheme, uncertainty, empirical-import switch, and no native-promotion claims.", "A complete redacted schema may pass the airlock, but it is not numerically executed because no observed values are loaded."},
		EnvironmentalEntries: []string{"Actual electroweak VEV, running/pole gauge couplings, weak angle, W/Z masses, kappa comparisons, and any matching residuals remain environmental bridge data."},
		FailedRoutes:         []string{StatusFailedSwitchClosed, StatusFailedMissingVEV, StatusFailedMissingGaugeCoupling, StatusFailedMissingScaleScheme, StatusFailedMissingSourceUncertainty, StatusFailedObservedMassAsNativeInput, StatusFailedWeakAngleNativePromotion, StatusFailedKappaPromotion, StatusFailedNativePromotion, StatusFailedNumericalAdapterNotRun, StatusFailedNoNativeElectroweakPrediction},
		OpenTheorems:         []string{"Gate507 may run an explicit observed electroweak data-file adapter only if the empirical switch is open and all rows carry scale/scheme/source/uncertainty metadata; outputs must remain bridge comparators.", "A separate native finite-action theorem would still be required to derive a nonzero Higgs ray, gauge couplings, kappa_U1, or W/Z mass matrix."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 507, Title: "Observed Electroweak Comparator File Adapter", Reason: "Gate506 defines the fail-closed observed electroweak airlock without importing values; the next safe step is an explicit bridge data-file run that computes comparator residuals while rejecting native promotion.", PrimaryTask: "load a tagged observed electroweak comparator ledger from file, compute only bridge-level tree residuals, preserve photon zero mode, and block every native electroweak registry write"}
}

func truth(a Analysis) string {
	if a.Preflight.AcceptedSchemaCases == 1 && a.Preflight.RejectedCases >= 10 && !a.Preflight.NumericalAdapterRun && !a.Firewall.NativeRegistryWritten {
		return "Gate506 establishes the observed electroweak comparator airlock: a complete redacted schema can be accepted as bridge-only, but no observed numerical values are imported and no adapter runs by default.  Every route that treats observed VEV, couplings, weak angle, W/Z masses, or kappa as native ASHA output is rejected before computation."
	}
	return "Gate506 did not establish the observed electroweak comparator airlock."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.SyntheticAdapterExecuted && a.Inheritance.SyntheticOnly && !a.Inheritance.Gate505ObservedDataImported && !a.Inheritance.Gate505NativeDataImported && a.Inheritance.Gate505NativeWriteBlocked && a.Inheritance.Gate506RedirectDefined, "Gate505 inheritance invalid"},
		{a.Preflight.Executed && a.Preflight.AcceptedSchemaCases == 1 && a.Preflight.RejectedCases == 10 && a.Preflight.ReadyForNumericalAdapterCases == 0 && !a.Preflight.NumericalAdapterRun && !a.Preflight.ObservedNumbersImported && a.Preflight.AllAcceptedBridgeOnlyObserved, "preflight counts invalid"},
		{a.Preflight.SwitchClosedRejected && a.Preflight.MissingVEVRejected && a.Preflight.MissingGaugeCouplingRejected && a.Preflight.MissingScaleSchemeRejected && a.Preflight.MissingSourceUncertaintyRejected && a.Preflight.ObservedMassAsNativeInputRejected && a.Preflight.WeakAngleNativePromotionRejected && a.Preflight.KappaPromotionRejected && a.Preflight.NativePromotionRejected, "preflight rejection coverage incomplete"},
		{a.Firewall.Executed && !a.Firewall.ObservedNumbersImported && !a.Firewall.NumericalAdapterExecuted && !a.Firewall.NativeVEVWritten && !a.Firewall.NativeGaugeCouplingWritten && !a.Firewall.NativeWeakAngleWritten && !a.Firewall.NativeWZMassWritten && !a.Firewall.NativeKappaWritten && !a.Firewall.NativeRegistryWritten && !a.Firewall.NativeElectroweakPredictionMade, "firewall violation"},
		{a.Next.Gate == 507, "Gate507 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate505=%t synthetic_executed=%t synthetic_only=%t observed=%t native_input=%t native_blocked=%t redirect506=%t verdict=%s reason=%s", x.Gate505AuditDefined, x.SyntheticAdapterExecuted, x.SyntheticOnly, x.Gate505ObservedDataImported, x.Gate505NativeDataImported, x.Gate505NativeWriteBlocked, x.Gate506RedirectDefined, x.Verdict, x.Reason)
}
func FormatPreflight(x Preflight) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d ready_numeric=%d adapter_run=%t observed_numbers=%t switch=%t missing_v=%t missing_g=%t missing_scale=%t missing_uncertainty=%t mass_native=%t theta_native=%t kappa=%t native_promotion=%t bridge_only=%t verdict=%s reason=%s", x.Executed, x.AcceptedSchemaCases, x.RejectedCases, x.ReadyForNumericalAdapterCases, x.NumericalAdapterRun, x.ObservedNumbersImported, x.SwitchClosedRejected, x.MissingVEVRejected, x.MissingGaugeCouplingRejected, x.MissingScaleSchemeRejected, x.MissingSourceUncertaintyRejected, x.ObservedMassAsNativeInputRejected, x.WeakAngleNativePromotionRejected, x.KappaPromotionRejected, x.NativePromotionRejected, x.AllAcceptedBridgeOnlyObserved, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_numbers=%t adapter=%t native_v=%t native_gauge=%t native_theta=%t native_wz=%t native_kappa=%t native_registry=%t prediction=%t verdict=%s reason=%s", x.ObservedNumbersImported, x.NumericalAdapterExecuted, x.NativeVEVWritten, x.NativeGaugeCouplingWritten, x.NativeWeakAngleWritten, x.NativeWZMassWritten, x.NativeKappaWritten, x.NativeRegistryWritten, x.NativeElectroweakPredictionMade, x.Verdict, x.Reason)
}
func FormatCase(c PreflightCase) string {
	return fmt.Sprintf("%s accepted=%t ready=%t run=%t imported=%t failures=[%s] reason=%s", c.Name, c.AcceptedSchema, c.ReadyForNumericalAdapter, c.NumericalAdapterRun, c.ObservedNumbersImported, strings.Join(c.Failures, ";"), c.Reason)
}
func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 506 Registry Audit — Observed Electroweak Comparator Airlock Preflight\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range []string{StatusGate505SyntheticAdapterInherited, StatusObservedEWPreflightPolicyDefined, StatusObservedEWSchemaAccepted, StatusObservedEWPreflightValidated, StatusObservedEWBridgeAirlockAccepted, StatusObservedEWNoNumericImport, StatusObservedEWNoAdapterExecuted, StatusFailedSwitchClosed, StatusFailedMissingVEV, StatusFailedMissingGaugeCoupling, StatusFailedMissingScaleScheme, StatusFailedMissingSourceUncertainty, StatusFailedObservedMassAsNativeInput, StatusFailedWeakAngleNativePromotion, StatusFailedKappaPromotion, StatusFailedNativePromotion, StatusFailedNumericalAdapterNotRun, StatusFailedNoNativeElectroweakPrediction, StatusFirewallObservedNumbersNotImported, StatusFirewallNativeWriteBlocked, StatusGate507ObservedEWAdapterRedirect} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate505 verified the electroweak bridge adapter only on explicit fake inputs.  It imported no observed electroweak data and wrote no synthetic output into the native registry.  Gate506 therefore may define an observed-data airlock, but it may not silently run observed matching or promote any comparator value to a theorem.\n\n")
	b.WriteString("```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Observed electroweak preflight schema\n\n")
	b.WriteString("Required bridge rows:\n\n")
	for _, r := range completeRedactedLedger().Rows {
		b.WriteString(fmt.Sprintf("- `%s`: source/version, scale, scheme, uncertainty, bridge-only flag, empirical-import switch, and native-promotion rejection.\n", r.Name))
	}
	b.WriteString("\nThe accepted preflight case is redacted: row names and metadata policy are present, but no numerical VEV, coupling, weak-angle, W, or Z value is loaded.\n\n")
	b.WriteString("## Fail-closed cases\n\n")
	b.WriteString(fmt.Sprintf("```text\n%s\n```\n\n", FormatPreflight(a.Preflight)))
	for _, c := range a.Preflight.Cases {
		b.WriteString("- " + FormatCase(c) + "\n")
	}
	b.WriteString("\n## Firewall result\n\n")
	b.WriteString("No observed electroweak numbers are imported at Gate506.  No numerical adapter executes.  Observed W/Z masses may only be comparator outputs or external bridge rows; they cannot be native inputs.  The weak angle, kappa, VEV, and gauge couplings remain forbidden as native writes.\n\n")
	b.WriteString("```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString("Gate507 should be:\n\n```text\nGate 507 — Observed Electroweak Comparator File Adapter\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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
