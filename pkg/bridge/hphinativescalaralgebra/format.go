package hphinativescalaralgebra

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 407 Registry Audit — H_phi-Native Scalar Selector Algebra / Pair-Degeneracy Closure Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 407 stops forcing the contact quartic `q4` into the scalar lane and audits only the native endomorphism algebra of `H_phi`: quaternionic weak-module actions, the pair-degenerate scalar response, and the canonical one-form/Yukawa edge quotient. It separates generic algebraic capacity from a canonical selected scalar theorem.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Native generator ledger\n\n```text\n")
	b.WriteString(FormatLedger(a.Ledger))
	for _, g := range a.Ledger.Generators {
		b.WriteString("\n")
		b.WriteString(FormatGenerator(g))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Generated scalar algebras\n\n```text\n")
	for i, c := range a.Closures {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(FormatClosure(c))
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
