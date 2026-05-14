package orientededgeincidence

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 403 Registry Audit — Oriented Edge-Incidence Boundary Operator Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 403 tests whether the finite one-form Dirac edge graph becomes a canonical `q4` selector after upgrading undirected adjacency to a signed chiral boundary operator `d`. The gate audits `d^T d` / `d^†d`, Majorana orientation, J-doubling, and possible four-dimensional quotients without importing Yukawa amplitudes, observed masses, CKM/PMNS data, or manual `q4` placement.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## q4 target\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Oriented boundary arena\n\n```text\n")
	b.WriteString(FormatArena(a.Arena))
	b.WriteString("\n```\n\n")
	b.WriteString("## Boundary candidate table\n\n```text\n")
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
