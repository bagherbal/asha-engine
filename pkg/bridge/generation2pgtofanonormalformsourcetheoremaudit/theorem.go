package generation2pgtofanonormalformsourcetheoremaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2PGToFanoNormalFormSourceTheoremAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 654 — P_G-to-Fano Normal-Form Source Theorem Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate654 P_G-to-Fano source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate653 Fano/Hitchin identity and firewall", Passed: a.Inherited.FanoHitchinIdentityInherited && a.Inherited.NormalFormInherited && a.Inherited.SymbolicPositive && a.Inherited.SymbolicNegative && a.Inherited.SymbolicMixedZero && a.Inherited.InternalMechanismClosed && !a.Inherited.PGToFanoAlreadyBasisFree && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate653FirewallPreserved && a.Inherited.Gate652FiniteSourceVisible && !a.Inherited.Gate652FullSourceTheorem, Detail: FormatInherited(a.Inherited)},
			{Name: "audit P_G pullback support decomposition", Passed: a.Support.OmegaPPPZero && a.Support.OmegaPPMNonzero && a.Support.OmegaPMMZero && a.Support.OmegaMMMNonzero && a.Support.ReducesToLambda21Plus03 && a.Support.Residual < tol && len(a.Support.Rows) == 4, Detail: FormatSupport(a.Support)},
			{Name: "certify negative volume-form source", Passed: near(a.BVolume.Beta, unit) && a.BVolume.OrientationSign == 1 && a.BVolume.SO3VolumeCovariant && a.BVolume.ResidualAgainstVolMinus < tol && a.BVolume.BasisIndependentVolume, Detail: FormatBVolume(a.BVolume)},
			{Name: "audit A as K7-minus to two-forms source map", Passed: a.AMap.Rank == minusDim && near(a.AMap.ScaleAlpha, unit) && a.AMap.IsometryUpToScale && a.AMap.ImageInSelfDualForms && a.AMap.ImageDimension == minusDim && a.AMap.WedgeOrthonormal && a.AMap.Residual < tol, Detail: FormatAMap(a.AMap)},
			{Name: "audit quaternionic/Fano two-form triple source", Passed: a.Quaternionic.FormsDefineEndomorphisms && a.Quaternionic.QuaternionicTriple && a.Quaternionic.JIdentityResidual < tol && a.Quaternionic.WedgeIdentityResidual < tol, Detail: FormatQuaternionic(a.Quaternionic)},
			{Name: "audit SO(3) gauge covariance of normal form", Passed: a.Gauge.AInvariant && a.Gauge.BVolumeInvariant && a.Gauge.FMapEquivariant && a.Gauge.NormalFormGaugeCovariant && !a.Gauge.BasisArbitrary, Detail: FormatGauge(a.Gauge)},
			{Name: "audit route source independence", Passed: a.Routes.AllRoutesReduce && a.Routes.SamePGSourcePackage && !a.Routes.RouteDependentOnly && len(a.Routes.Rows) == 3, Detail: FormatRoutes(a.Routes)},
			{Name: "sharpen P_G-to-Fano source theorem", Passed: a.SourceTheorem.PGForcesFanoNormalForm && a.SourceTheorem.GaugeControlledSource && !a.SourceTheorem.BasisFreeSourceTheorem && a.SourceTheorem.Gate653ImplicationAvailable && a.SourceTheorem.InternalMechanismSourced, Detail: FormatSourceTheorem(a.SourceTheorem)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate654Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate654 sources the Gate653 normal form from the finite P_G/Fano package up to scale and SO(3) gauge, while preserving the stronger basis-free source theorem as a still-open target.")
		if !strings.Contains(a.SourceTheorem.Verdict, StatusNoFullBasisFreePGToFanoTheorem) {
			notes = append(notes, "WARNING_MISSING_BASIS_FREE_SOURCE_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
