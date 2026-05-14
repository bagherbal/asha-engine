package u1orientation

import (
	"fmt"
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ChiralOrientationalAbelianSourceSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CHIRAL-ORIENTATIONAL-ABELIAN-SOURCE-SEARCH"
	const name = "chiral / orientational abelian kinetic-source search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build chiral/orientational source search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 77 non-factorized input", Passed: a.Previous.NonFactorizedActionDerived && a.Previous.LocalNonzeroCorrelation, Detail: fmt.Sprintf("local correlation exists, but signed total=%.3e", a.Previous.SignedBLContactMoment)},
			{Name: "natural orientation probes", Passed: a.NaturalProbeCount >= 8, Detail: formatProbes(a.Probes, true)},
			{Name: "natural signed sources", Passed: a.NaturalNonzeroSources > 0, Detail: fmt.Sprintf("natural nonzero sources=%d, best |signed|=%.3e; all natural probes cancel", a.NaturalNonzeroSources, a.BestNaturalAbsSigned)},
			{Name: "kind-level cancellation ledger", Passed: len(a.Kinds) == 4, Detail: formatKinds(a.Kinds)},
			{Name: "non-canonical selector firewall", Passed: a.ArbitraryNonzeroSources > 0, Detail: "nonzero sources can be manufactured by selectors such as up-only, but these are rejected unless derived by a finite theorem"},
			{Name: "canonical orientational source", Passed: a.CanonicalSourceDerived, Detail: "not derived; chiral/parity/scalar orientation does not break the cancellation"},
			{Name: "full U(1) Hessian", Passed: a.FullU1HessianDerived, Detail: "not derived; no physical U(1)_Y coupling follows"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, v, or measured coupling was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatProbes(xs []Probe, canonicalOnly bool) string {
	parts := make([]string, 0, len(xs))
	for _, p := range xs {
		if canonicalOnly && !p.Canonical {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s(signed=%.3e, abs=%.3e, cancels=%v)", p.Name, p.Signed, p.Absolute, p.Cancels))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func formatKinds(xs []KindMoment) string {
	parts := make([]string, 0, len(xs))
	total := 0.0
	for _, k := range xs {
		total += k.Signed
		parts = append(parts, fmt.Sprintf("%s(signed=%.10f, abs=%.10f)", k.Kind, k.Signed, k.Absolute))
	}
	parts = append(parts, fmt.Sprintf("total=%.3e", total))
	parts = append(parts, fmt.Sprintf("cancels=%v", math.Abs(total) < 1e-10))
	return "[" + strings.Join(parts, "; ") + "]"
}
