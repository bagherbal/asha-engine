package edgetohphiquotient

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 404 Registry Audit — Canonical Edge-to-H_phi Quotient / Contact-Edge Intertwiner Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 404 tests whether the finite one-form edge-slot space has a native quotient/intertwiner onto the four-real scalar carrier `H_phi` such that the induced operator `Q^T Delta_edge Q` becomes the irreducible contact quartic `q4`. It promotes only quotients selected by one-form support, `J`, first-order compatibility, scalar branch data, or contact/scalar response. Arbitrary four-mode projections and q4 companion placement are quarantined.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## q4 target\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Quotient arena\n\n```text\n")
	b.WriteString(FormatArena(a.Arena))
	b.WriteString("\n```\n\n")
	b.WriteString("## Quotient/intertwiner candidate table\n\n```text\n")
	b.WriteString(FormatSieve(a.Sieve))
	b.WriteString("\n```\n\n")
	b.WriteString("## Identity / impact audit\n\n```text\n")
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
