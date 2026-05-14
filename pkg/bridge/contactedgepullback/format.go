package contactedgepullback

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 405 Registry Audit — Contact-to-Edge Natural Transformation / Pullback Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 405 reverses the Gate-404 quotient arrow. Instead of pushing the one-form edge graph down to `H_phi`, it asks whether the exact contact quartic primary `Q[x]/(q4)` has a native pullback or natural transformation into the J-doubled one-form edge ledger `Omega^1_D(A_F)`. It promotes only a typed map selected by existing ASHA structures and quarantines companion matrices, root placements, chosen edge bases, and arbitrary injections.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## q4 target\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Pullback arena\n\n```text\n")
	b.WriteString(FormatArena(a.Arena))
	b.WriteString("\n```\n\n")
	b.WriteString("## Pullback / natural transformation sieve\n\n```text\n")
	b.WriteString(FormatSieve(a.Sieve))
	b.WriteString("\n```\n\n")
	b.WriteString("## Impact audit\n\n```text\n")
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
