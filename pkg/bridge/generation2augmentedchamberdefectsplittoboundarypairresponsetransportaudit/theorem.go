package generation2augmentedchamberdefectsplittoboundarypairresponsetransportaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE942-GENERATION2AUGMENTEDCHAMBERDEFECTSPLITTOBOUNDARYPAIRRESPONSETRANSPORTAUDIT"
	theoremName = "Gate 942: AugmentedChamberDefectSplit to BoundaryPair Response Transport Audit"
)

func Generation2AugmentedChamberDefectSplitToBoundaryPairResponseTransportAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 942 H72-to-B2 transport audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, a.Carrier.Supports, a.Interface.Supports, a.UniformInsertion.Supports, a.Identification.Supports)
		allFailures := appendAll(a.Failures, a.Carrier.Failures, a.Interface.Failures, a.UniformInsertion.Failures, a.Identification.Failures)
		checks := []theorem.Check{
			{Name: "inherits Gate 941 S_split origin status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "validates shared H72 carrier ranks", Passed: a.Carrier.Lambda4Rank+a.Carrier.BoundaryRank == a.Carrier.TotalRank && a.Carrier.TotalRank == H72Rank, Detail: FormatCarrier(a.Carrier)},
			{Name: "detects common 7/72 defect and alpha-quadratic normalization", Passed: a.Carrier.DefectLane == "D_base=(7/72)S_split" && a.Carrier.AlphaQuadraticLane == "alpha_quad=(7/72)S_split^2", Detail: FormatCarrier(a.Carrier)},
			{Name: "finds B2 as shared boundary interface of H72", Passed: a.Interface.Boundary == "B2=<b1,b2>" && a.Interface.Status == TransportStronglySourceTyped, Detail: FormatInterface(a.Interface)},
			{Name: "preserves uniform S_split insertion and exterior square origin", Passed: len(a.UniformInsertion.BoundaryFactors) == 2 && a.UniformInsertion.UniformScalar == "S_split", Detail: FormatUniformInsertion(a.UniformInsertion)},
			{Name: "keeps Certificate II blocked without native descent", Passed: !a.CertificateIIPassed && !a.Identification.CertificateIIPassed && !a.Identification.NativeCertified, Detail: FormatIdentification(a.Identification)},
			{Name: "preserves H72 descent firewalls", Passed: containsAll(allFailures, Failures()) && containsAll(allFailures, a.Identification.Failures), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatCarrier(a.Carrier), FormatInterface(a.Interface), FormatUniformInsertion(a.UniformInsertion), FormatIdentification(a.Identification), a.CertificateIIStatus, a.Final, NextGate}
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
