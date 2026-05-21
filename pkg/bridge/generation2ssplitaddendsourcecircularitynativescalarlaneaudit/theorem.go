package generation2ssplitaddendsourcecircularitynativescalarlaneaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE945-GENERATION2SSPLITADDENDSOURCECIRCULARITYNATIVESCALARLANEAUDIT"
	theoremName = "Gate 945: S_split Addend Source, Circularity, and Native ScalarLane Audit"
)

func Generation2SSplitAddendSourceCircularityNativeScalarLaneAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 945 addend source audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		addendNotes := []string{}
		for _, addend := range a.Addends {
			addendNotes = append(addendNotes, FormatAddend(addend))
		}
		truthNotes := []string{}
		for _, row := range a.TruthTable {
			truthNotes = append(truthNotes, FormatTruthRow(row))
		}
		allSupports := appendAll(a.Supports, allAddendSupports(a.Addends), a.ScalarLane.Supports, a.Certificate.Supports)
		allFailures := appendAll(a.Failures, allAddendFailures(a.Addends), a.ScalarLane.Failures, a.Certificate.Failures)
		checks := []theorem.Check{
			{Name: "inherits Gate 944 bridge/history scalar status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "keeps S_split expression explicit", Passed: a.Expression == "S_split=(R_3-1)+lambda(Lambda_12)" && a.Value == Ssplit, Detail: a.Expression},
			{Name: "audits two scalar addends without native certification", Passed: len(a.Addends) == 2 && a.Addends[0].ScalarAddendTyped && a.Addends[1].ScalarAddendTyped && !a.Addends[0].NativeH72Scalar && !a.Addends[1].NativeH72Scalar, Detail: stringsJoin(addendNotes)},
			{Name: "detects R3-minus-one circularity risk", Passed: a.Addends[0].Name == "R_3-1" && a.Addends[0].CircularRisk, Detail: FormatAddend(a.Addends[0])},
			{Name: "keeps lambda(Lambda_12) as bridge/history scalar", Passed: a.Addends[1].Name == "lambda(Lambda_12)" && a.Addends[1].BridgeHistoryScalar && !a.Addends[1].NativeH72Scalar, Detail: FormatAddend(a.Addends[1])},
			{Name: "requires common H72 scalar lane before native sum", Passed: a.ScalarLane.CanonicalAddition && !a.ScalarLane.BothAddendsNative && !a.ScalarLane.SsplitNative, Detail: FormatScalarLane(a.ScalarLane)},
			{Name: "truth table marks current bridge/history and circular-risk rows", Passed: containsCurrentCircularOrBridgeRow(a.TruthTable), Detail: stringsJoin(truthNotes)},
			{Name: "keeps Certificate II blocked by addend sources", Passed: a.Certificate.TransportLayerStrong && a.Certificate.CentralH72Compatible && !a.Certificate.AddendSourceNative && !a.Certificate.CertificatePassed, Detail: FormatCertificate(a.Certificate)},
			{Name: "preserves native-R3 and addend-source firewalls", Passed: containsAll(allFailures, Failures()), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Expression, FormatScalarLane(a.ScalarLane), FormatCertificate(a.Certificate), a.Final, NextGate}
		notes = append(notes, addendNotes...)
		notes = append(notes, truthNotes...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
