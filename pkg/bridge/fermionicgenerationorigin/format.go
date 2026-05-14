package fermionicgenerationorigin

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 409 Registry Audit — Fermionic Matter-Carrier Origin / Nontrivial Generation Representation Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 409 is not an empirical Yukawa seal. It pivots after the H_phi scalar flavor-blindness result and audits whether the fermionic matter carrier itself derives a nontrivial generation representation before Yukawa amplitudes are inserted.\n\n")
	b.WriteString("## Why this is not Gate 409 as surrender seal\n\n")
	b.WriteString("The gate forbids observed masses, CKM/PMNS data, scalar q4 promotion, tau_eta insertion, and manual N=diag(0,1,2). It searches the fermion carrier, primitive ideals, commutant, bilinears, and dynamic generation-source candidates instead.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Fermionic carrier inventory\n\n```text\n")
	b.WriteString(FormatInventory(a.Inventory))
	for _, c := range a.Inventory.Carriers {
		b.WriteString("\n")
		b.WriteString(FormatCarrier(c))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Primitive idempotent table\n\n```text\n")
	b.WriteString(FormatIdempotentAudit(a.Idempotents))
	for _, c := range a.Idempotents.Candidates {
		b.WriteString("\n")
		b.WriteString(FormatIdempotent(c))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Commutant / centralizer result\n\n```text\n")
	b.WriteString(FormatCommutant(a.Commutant))
	b.WriteString("\n```\n\n")
	b.WriteString("## Triality from fermion side\n\n```text\n")
	b.WriteString(FormatTriality(a.Triality))
	b.WriteString("\n```\n\n")
	b.WriteString("## Fermionic bilinear operator table\n\n```text\n")
	b.WriteString(FormatBilinearAudit(a.Bilinears))
	for _, f := range a.Bilinears.Families {
		b.WriteString("\n")
		b.WriteString(FormatBilinear(f))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Dynamic generation-source table\n\n```text\n")
	b.WriteString(FormatSourceAudit(a.Sources))
	for _, s := range a.Sources.Sources {
		b.WriteString("\n")
		b.WriteString(FormatSource(s))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## CKM / PMNS capacity result\n\n```text\n")
	b.WriteString(FormatCKM(a.CKM))
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
