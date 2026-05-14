package asha

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type Format string

const (
	FormatText     Format = "text"
	FormatMarkdown Format = "markdown"
	FormatJSON     Format = "json"
)

func ParseFormat(s string) (Format, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	switch Format(s) {
	case FormatText, FormatMarkdown, FormatJSON:
		return Format(s), nil
	default:
		return "", fmt.Errorf("unknown format %q", s)
	}
}

func WriteReport(w io.Writer, r Report, f Format) error {
	switch f {
	case FormatJSON:
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		return enc.Encode(r)
	case FormatMarkdown:
		_, err := io.WriteString(w, Markdown(r))
		return err
	case FormatText:
		_, err := io.WriteString(w, Text(r))
		return err
	default:
		return fmt.Errorf("unsupported format %q", f)
	}
}

func Text(r Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "ASHA runtime board — %s (Gate %d)\n", r.Metadata.RuntimeVersion, r.Metadata.LatestGate)
	fmt.Fprintf(&b, "Scenario: %s\n", r.Scenario)
	fmt.Fprintf(&b, "Source: %s\n\n", r.Metadata.Source)
	for _, sec := range r.Sections {
		fmt.Fprintf(&b, "== %s ==\n%s\n", sec.Name, sec.Summary)
		if len(sec.Quantities) > 0 {
			tw := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
			fmt.Fprintln(tw, "SYMBOL\tVALUE\tSTATUS\tFORMULA/NOTE")
			for _, q := range sec.Quantities {
				fmt.Fprintf(tw, "%s\t%s\t%s\t%s %s\n", q.Symbol, qValue(q), q.Status, q.Formula, q.Note)
			}
			tw.Flush()
		}
		if len(sec.Boundaries) > 0 {
			for _, bd := range sec.Boundaries {
				fmt.Fprintf(&b, "BOUNDARY: %s | %s | %s\n", bd.Name, bd.Formula, bd.Meaning)
			}
		}
		b.WriteByte('\n')
	}
	fmt.Fprintf(&b, "Checks:\n")
	for _, c := range r.Checks {
		mark := "PASS"
		if !c.Passed {
			mark = "FAIL"
		}
		fmt.Fprintf(&b, "- %s %s: %s\n", mark, c.Name, c.Detail)
	}
	fmt.Fprintf(&b, "\nVerdict: %s\n", r.Verdict)
	return b.String()
}

func Markdown(r Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# ASHA Runtime Board\n\n")
	fmt.Fprintf(&b, "- Runtime: `%s`\n- Latest gate: `%d`\n- Scenario: `%s`\n- Source: `%s`\n\n", r.Metadata.RuntimeVersion, r.Metadata.LatestGate, r.Scenario, r.Metadata.Source)
	for _, sec := range r.Sections {
		fmt.Fprintf(&b, "## %s\n\n%s\n\n", sec.Name, sec.Summary)
		if len(sec.Quantities) > 0 {
			b.WriteString("| Symbol | Value | Status | Formula / note |\n|---|---:|---|---|\n")
			for _, q := range sec.Quantities {
				fmt.Fprintf(&b, "| `%s` | %s | `%s` | %s %s |\n", q.Symbol, qValue(q), q.Status, esc(q.Formula), esc(q.Note))
			}
			b.WriteByte('\n')
		}
		if len(sec.Boundaries) > 0 {
			b.WriteString("### Boundaries\n\n")
			for _, bd := range sec.Boundaries {
				fmt.Fprintf(&b, "- **%s:** `%s` — %s\n", bd.Name, bd.Formula, bd.Meaning)
			}
			b.WriteByte('\n')
		}
	}
	b.WriteString("## Runtime checks\n\n")
	for _, c := range r.Checks {
		mark := "✅"
		if !c.Passed {
			mark = "❌"
		}
		fmt.Fprintf(&b, "- %s **%s** — %s\n", mark, c.Name, c.Detail)
	}
	fmt.Fprintf(&b, "\n## Verdict\n\n%s\n", r.Verdict)
	return b.String()
}

func qValue(q Quantity) string {
	if q.Text != "" {
		return q.Text
	}
	if q.Unit != "" {
		return fmt.Sprintf("%.12g %s", q.Value, q.Unit)
	}
	return fmtFloat(q.Value)
}

func fmtFloat(x float64) string { return fmt.Sprintf("%.12g", x) }
func esc(s string) string       { return strings.ReplaceAll(s, "|", "\\|") }
