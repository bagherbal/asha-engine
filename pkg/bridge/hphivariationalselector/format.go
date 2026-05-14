package hphivariationalselector

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 408 Registry Audit — H_phi Variational Functional / Canonical Coefficient Selector Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 408 audits whether a native scalar action, Hessian, one-form kinetic trace, or quaternionic invariant trace selects a unique non-pair-degenerate element from the full `End_R(H_phi)` capacity discovered in Gate 407. It rejects arbitrary source terms and empirical Yukawa inputs.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Functional ledger\n\n```text\n")
	b.WriteString(FormatLedger(a.Ledger))
	for _, f := range a.Ledger.Functionals {
		b.WriteString("\n")
		b.WriteString(FormatFunctional(f))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Selector candidates\n\n```text\n")
	for i, s := range a.Selectors {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(FormatSelector(s))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Variational outcome\n\n```text\n")
	b.WriteString(FormatOutcome(a.Outcome))
	b.WriteString("\n```\n\n")
	b.WriteString("## Moduli impact\n\n```text\n")
	b.WriteString(FormatImpact(a.Impact))
	b.WriteString("\n```\n\n")
	b.WriteString("## Firewall status\n\n```text\n")
	b.WriteString(FormatFirewall(a.Firewall))
	b.WriteString("\n```\n\n")
	b.WriteString("## Statuses\n\n```text\n")
	b.WriteString(strings.Join(Statuses(a), "\n"))
	b.WriteString("\n```\n\n")
	b.WriteString("## Conclusion\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n## Next gate\n\n```text\n")
	b.WriteString(FormatNext(a.Next))
	b.WriteString("\n```\n")
	return b.String()
}
