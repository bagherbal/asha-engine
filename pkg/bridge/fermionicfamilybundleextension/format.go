package fermionicfamilybundleextension

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 410 Registry Audit — Fermionic Representation Extension / Nontrivial Family Bundle Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 410 audits whether advanced representation extensions already present in ASHA replace the trivial fermionic generation multiplicity with a nontrivial family bundle. It is not an empirical Yukawa seal and does not promote new axioms.\n\n")
	b.WriteString("## Prior boundary inherited\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Extension candidate table\n\n```text\n")
	b.WriteString(FormatExtensionAudit(a.Extensions))
	for _, c := range a.Extensions.Candidates {
		b.WriteString("\n")
		b.WriteString(FormatCandidate(c))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Family bundle construction result\n\n```text\n")
	b.WriteString(FormatFamilyBundle(a.FamilyBundle))
	b.WriteString("\n```\n\n")
	b.WriteString("## KO / twisted spectral triple audit\n\n```text\n")
	b.WriteString(FormatKOTwist(a.KOTwist))
	b.WriteString("\n```\n\n")
	b.WriteString("## Modular nontracial / KMS audit\n\n```text\n")
	b.WriteString(FormatModularKMS(a.ModularKMS))
	b.WriteString("\n```\n\n")
	b.WriteString("## Primitive ideal extension audit\n\n```text\n")
	b.WriteString(FormatPrimitiveIdeals(a.PrimitiveIdeals))
	b.WriteString("\n```\n\n")
	b.WriteString("## Noncommuting texture capacity\n\n```text\n")
	b.WriteString(FormatNoncommuting(a.Noncommuting))
	b.WriteString("\n```\n\n")
	b.WriteString("## Moduli impact table\n\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	for _, s := range a.Moduli.Scenarios {
		b.WriteString("\n")
		b.WriteString(FormatScenario(s))
	}
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
