package contacteigenoperatorreconstruction

import "strings"

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 406 Registry Audit — Contact-Eigenoperator Internal Reconstruction / q4 Lives Only in Contact Sector\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 406 stops trying to force the contact quartic `q4` into `H_phi` or the one-form edge ledger. It reconstructs `q4` internally as a contact-sector eigenoperator and classifies whether it should remain a contact-only spectral invariant under the current functorial inventory.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Internal contact q4 operator\n\n```text\n")
	b.WriteString(FormatContactQ4(a.ContactQ4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Contact algebra / centralizer audit\n\n```text\n")
	b.WriteString(FormatContactAlgebra(a.ContactAlgebra))
	b.WriteString("\n```\n\n")
	b.WriteString("## Classification sieve\n\n```text\n")
	b.WriteString(FormatClassification(a.Classification))
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
