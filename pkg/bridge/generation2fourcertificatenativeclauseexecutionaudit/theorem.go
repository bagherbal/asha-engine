package generation2fourcertificatenativeclauseexecutionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE940-GENERATION2FOURCERTIFICATENATIVECLAUSEEXECUTIONAUDIT"
	theoremName = "Gate 940: FourCertificate Native Clause Execution and FailureLocalization Audit"
)

func Generation2FourCertificateNativeClauseExecutionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 940 four-certificate execution audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		failuresWithCertificates := append([]string{}, a.Failures...)
		failuresWithCertificates = append(failuresWithCertificates, certificateFailures(a.Certificates)...)
		supportsWithCertificates := append([]string{}, a.Supports...)
		supportsWithCertificates = append(supportsWithCertificates, certificateSupports(a.Certificates)...)
		checks := []theorem.Check{
			{Name: "inherits Gate 939 origin-rooted collapse", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "executes exactly four native-promotion certificates", Passed: len(a.Certificates) == 4, Detail: FormatCertificates(a.Certificates)},
			{Name: "does not grant native R3 unless all certificates pass", Passed: !a.FullNativeEligible && !allNativeCertified(a.Certificates), Detail: a.Verdict},
			{Name: "localizes partial and blocked clauses", Passed: countStatus(a.Certificates, CertificatePartialSupport) == 2 && countStatus(a.Certificates, CertificateBlocked) == 2, Detail: FormatCertificates(a.Certificates)},
			{Name: "S_split and orientation remain primary pressure points", Passed: containsAll(a.PrimaryPressure, []string{"Certificate II: S_split native response-parameter source", "Certificate IV: lawful finite one-form/spontaneous orientation or full A_F descent"}), Detail: stringsJoin(a.PrimaryPressure)},
			{Name: "native-promotion firewalls are preserved", Passed: containsAll(failuresWithCertificates, append(Failures(), certificateFailures(a.Certificates)...)), Detail: stringsJoin(failuresWithCertificates)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatDiagnostics(a.DiagnosticValues), FormatCertificates(a.Certificates), a.Final, NextGate}
		for _, c := range a.Certificates {
			notes = append(notes, "CERTIFICATE: "+c.Name, c.Question, c.RequiredTheorem, string(c.Status), c.Localization)
			notes = append(notes, c.PassMarkers...)
			notes = append(notes, c.SupportMarkers...)
			notes = append(notes, c.FailureMarkers...)
		}
		notes = append(notes, supportsWithCertificates...)
		notes = append(notes, failuresWithCertificates...)
		notes = append(notes, a.PrimaryPressure...)
		notes = append(notes, a.RetiredNonBlockers...)
		notes = append(notes, a.R4Boundary...)
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
