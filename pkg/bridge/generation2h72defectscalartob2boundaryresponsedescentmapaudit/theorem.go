package generation2h72defectscalartob2boundaryresponsedescentmapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE943-GENERATION2H72DEFECTSCALARTOB2BOUNDARYRESPONSEDESCENTMAPAUDIT"
	theoremName = "Gate 943: H72 DefectScalar to B2 BoundaryResponse DescentMap Audit"
)

func Generation2H72DefectScalarToB2BoundaryResponseDescentMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 943 H72 defect scalar descent audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, a.Projection.Supports, a.Restriction.Supports, a.Insertion.Supports, a.Relation.Supports, a.Certificate.Supports)
		allFailures := appendAll(a.Failures, a.Projection.Failures, a.Restriction.Failures, a.Insertion.Failures, a.Relation.Failures, a.Certificate.Failures)
		checks := []theorem.Check{
			{Name: "inherits Gate 942 transport-interface status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "validates H72 direct-sum boundary projection", Passed: a.Projection.Canonical && Lambda4V8Rank+B2Rank == H72Rank, Detail: FormatProjection(a.Projection)},
			{Name: "central scalar restriction forces s=S_split on B2", Passed: a.Restriction.SameScalarValue && a.Restriction.ForcesSEquals && !a.Restriction.NativeSourceKnown, Detail: FormatRestriction(a.Restriction)},
			{Name: "boundary insertion is uniform and needs no second transport", Passed: a.Insertion.Uniform && !a.Insertion.SecondTransportReq, Detail: FormatInsertion(a.Insertion)},
			{Name: "keeps H72 defect and alpha quadratic relation on shared normalization", Passed: a.Relation.SharedChamber == "H72" && a.Relation.SharedWeight == "7/72", Detail: FormatRelation(a.Relation)},
			{Name: "supports transport component but keeps Certificate II open", Passed: a.Certificate.TransportComponentSupported && !a.Certificate.NativeSourceCertified && !a.Certificate.CertificatePassed, Detail: FormatCertificate(a.Certificate)},
			{Name: "preserves native-source firewalls", Passed: containsAll(allFailures, Failures()) && containsAll(allFailures, a.Certificate.Failures), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatProjection(a.Projection), FormatRestriction(a.Restriction), FormatInsertion(a.Insertion), FormatRelation(a.Relation), FormatCertificate(a.Certificate), a.Final, NextGate}
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
