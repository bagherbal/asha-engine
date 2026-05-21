package generation2ssplitnativesourceboundaryresponsescalaroriginaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE941-GENERATION2SSPLITNATIVESOURCEBOUNDARYRESPONSESCALARORIGINAUDIT"
	theoremName = "Gate 941: S_split Native Source and BoundaryResponse Scalar Origin Audit"
)

func Generation2SSplitNativeSourceBoundaryResponseScalarOriginAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 941 S_split origin audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, a.Origin.Supports, a.Transport.Supports)
		allFailures := appendAll(a.Failures, a.Origin.Failures, a.Transport.Failures)
		checks := []theorem.Check{
			{Name: "inherits Gate 940 Certificate II localization", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "traces S_split to augmented chamber defect split", Passed: a.Origin.Name == "augmented chamber defect-trace split" && a.Origin.Status == SourceBridgeStrongNotNative, Detail: FormatOrigin(a.Origin)},
			{Name: "uses S_split as scalar B2 response parameter candidate", Passed: a.Transport.Candidate == "T_s(S_split)=s with uniform insertion into (1+s b1)(1+s b2)", Detail: FormatTransport(a.Transport)},
			{Name: "keeps Certificate II blocked without native transport", Passed: !a.CertificateIIPass && !a.Transport.CertificatePassed && !a.Transport.NativeCertified, Detail: a.Verdict},
			{Name: "preserves alpha reconstruction values under bridge use", Passed: a.Response.AlphaLinear == AlphaLinear && a.Response.AlphaQuadratic == AlphaQuad && a.Response.AlphaTotal == AlphaB, Detail: FormatResponse(a.Response)},
			{Name: "preserves native transport firewalls", Passed: containsAll(allFailures, Failures()) && containsAll(allFailures, a.Transport.Failures), Detail: stringsJoin(allFailures)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatOrigin(a.Origin), FormatTransport(a.Transport), FormatResponse(a.Response), a.Final, NextGate}
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func stringsJoin(items []string) string {
	out := ""
	for i, item := range items {
		if i > 0 {
			out += " | "
		}
		out += item
	}
	return out
}
