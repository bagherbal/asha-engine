package generation2ssplitnativeh72scalarsourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE944-GENERATION2SSPLITNATIVEH72SCALARSOURCEAUDIT"
	theoremName = "Gate 944: S_split Native H72 Scalar Source Audit"
)

func Generation2SSplitNativeH72ScalarSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 944 S_split native H72 scalar source audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		componentNotes := []string{}
		for _, c := range a.Components {
			componentNotes = append(componentNotes, FormatComponent(c))
		}
		allSupports := appendAll(a.Supports, a.Expression.Supports, componentSupports(a.Components), a.Criteria.Supports, a.Certificate.Supports)
		allFailures := appendAll(a.Failures, a.Expression.Failures, componentFailures(a.Components), a.Criteria.Failures, a.Certificate.Failures)
		checks := []theorem.Check{
			{Name: "inherits Gate 943 source-native-status-open result", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "keeps S_split expression explicit and H72-compatible", Passed: a.Expression.Expression == "S_split=(R_3-1)+lambda(Lambda_12)" && a.Expression.Dimensionless && a.Expression.H72Compatible && !a.Expression.NativeCertified, Detail: FormatExpression(a.Expression)},
			{Name: "audits components without native certification", Passed: len(a.Components) == 2 && !a.Components[0].NativeCertified && !a.Components[1].NativeCertified, Detail: stringsJoin(componentNotes)},
			{Name: "localizes blocked native finite H72 chamber criteria", Passed: a.Criteria.Status == SourceNativeMissing && len(a.Criteria.Blocked) >= 3, Detail: FormatCriteria(a.Criteria)},
			{Name: "keeps Certificate II blocked by native source status", Passed: a.Certificate.DescentSupported && !a.Certificate.SourceNative && !a.Certificate.CertificatePassed, Detail: FormatCertificate(a.Certificate)},
			{Name: "preserves bridge-history scalar firewalls", Passed: containsAll(allFailures, Failures()) && containsAll(allFailures, a.Certificate.Failures), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatExpression(a.Expression), FormatCriteria(a.Criteria), FormatCertificate(a.Certificate), a.Final, NextGate}
		notes = append(notes, componentNotes...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
