package protectedconnection

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ProtectedCarrierOperatorBFContactConnectionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-PROTECTED-CARRIER-BF-CONNECTION-SEARCH"
	const name = "protected-carrier operator and BF/contact connection search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build protected-carrier connection audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 88 protected metric input", Passed: a.ProtectedMetric.AbstractEuclideanMetricAvailable && !a.ProtectedMetric.FiniteProtectedConnectionDerived, Detail: fmt.Sprintf("protected dim=%d, abstract I3 metric exists, finite connection derived=%v", a.ProtectedDimension, a.ProtectedMetric.FiniteProtectedConnectionDerived)},
			{Name: "abstract protected operator spaces", Passed: a.AbstractOperatorSpaceExists && a.AbstractSO3ConnectionExists, Detail: fmt.Sprintf("End(R^3) dim=%d, skew so(3) connection dim=%d, symmetric metric-deformation dim=%d", a.EndomorphismDimension, a.SkewConnectionDimension, a.SymmetricMetricDeformationDimension)},
			{Name: "abstract operators are not finite dynamics", Passed: !a.AbstractOperatorSpaceCanonical, Detail: "operator spaces exist for every 3D vector space; a finite action must select a specific operator/connection"},
			{Name: "diagonal generation spurion available", Passed: a.DiagonalSpurionAvailable, Detail: fmt.Sprintf("bridge-level eigenvalues=%s", FormatSlice(a.DiagonalSpurionEigenvalues))},
			{Name: "diagonal spurion is not protected BF connection", Passed: !a.DiagonalSpurionIntrinsicToProtected && !a.DiagonalSpurionReducesO3, Detail: "it is a useful Higgs/contact anisotropy bridge, not an intrinsic protected-carrier connection form"},
			{Name: "protected contact curvature restriction", Passed: a.ContactCurvatureFlatOnProtected, Detail: fmt.Sprintf("operators=%d, max norm=%s, span rank=%d", a.ContactCurvatureOperators, FormatScientific(a.ContactCurvatureMaxNorm), a.ContactCurvatureSpanRank)},
			{Name: "active-sector curvature remains nonzero", Passed: a.ActiveCurvatureNonzero, Detail: fmt.Sprintf("active max norm=%s, active span rank=%d", FormatScientific(a.ActiveCurvatureMaxNorm), a.ActiveCurvatureSpanRank)},
			{Name: "intrinsic protected operator selected", Passed: a.IntrinsicProtectedOperatorDerived, Detail: "not derived; current finite BF/contact data does not select a nonzero operator on the protected carrier"},
			{Name: "O(3) freedom reduced or proven gauge", Passed: a.O3FreedomReduced || a.O3FreedomProvenGauge, Detail: "still open; no protected connection/operator reduces the O(3) frame freedom yet"},
		}, Notes: []string{a.TruthStatement, "candidate sources: " + Join(a.CandidateOperatorSources), "rejected shortcuts: " + Join(a.RejectedShortcuts), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
