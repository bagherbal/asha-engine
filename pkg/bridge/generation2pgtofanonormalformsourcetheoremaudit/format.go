package generation2pgtofanonormalformsourcetheoremaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate653Inheritance) string {
	return fmt.Sprintf("gate653=%t normalForm=%t positive=%t negative=%t mixed=%t closed=%t pgBasisFree=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t gate652Finite=%t gate652Full=%t verdict=%q", x.FanoHitchinIdentityInherited, x.NormalFormInherited, x.SymbolicPositive, x.SymbolicNegative, x.SymbolicMixedZero, x.InternalMechanismClosed, x.PGToFanoAlreadyBasisFree, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate653FirewallPreserved, x.Gate652FiniteSourceVisible, x.Gate652FullSourceTheorem, x.Verdict)
}

func FormatSupport(x PGSupportDecompositionAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s degree=%s norm=%.6g expected=%s residual=%.3g verdict=%q", r.Component, r.Degree, r.Norm, r.Expected, r.Residual, r.Verdict))
	}
	return fmt.Sprintf("Omega+++=0:%t Omega++-:%t Omega+--=0:%t Omega---:%t lambda21+03=%t residual=%.3g verdict=%q rows=%s", x.OmegaPPPZero, x.OmegaPPMNonzero, x.OmegaPMMZero, x.OmegaMMMNonzero, x.ReducesToLambda21Plus03, x.Residual, x.Verdict, strings.Join(parts, "; "))
}

func FormatBVolume(x NegativeVolumeSourceAudit) string {
	return fmt.Sprintf("expr=%q beta=%.6g sign=%d so3Vol=%t residual=%.3g basisIndependent=%t verdict=%q", x.BExpression, x.Beta, x.OrientationSign, x.SO3VolumeCovariant, x.ResidualAgainstVolMinus, x.BasisIndependentVolume, x.Verdict)
}

func FormatAMap(x AMapSourceAudit) string {
	return fmt.Sprintf("map=%q domain=%q codomain=%q rank=%d alpha=%.6g adj=%q isometry=%t selfDual=%t imageDim=%d wedge=%t residual=%.3g verdict=%q", x.MapName, x.Domain, x.Codomain, x.Rank, x.ScaleAlpha, x.FAdjointF, x.IsometryUpToScale, x.ImageInSelfDualForms, x.ImageDimension, x.WedgeOrthonormal, x.Residual, x.Verdict)
}

func FormatQuaternionic(x QuaternionicSourceAudit) string {
	return fmt.Sprintf("endomorphisms=%t identity=%q jResidual=%.3g wedgeResidual=%.3g orientation=%q triple=%t verdict=%q", x.FormsDefineEndomorphisms, x.JIdentity, x.JIdentityResidual, x.WedgeIdentityResidual, x.OrientationConvention, x.QuaternionicTriple, x.Verdict)
}

func FormatGauge(x SO3GaugeCovarianceAudit) string {
	return fmt.Sprintf("eta=%q omega=%q Ainv=%t Bvol=%t Feq=%t normalForm=%t arbitrary=%t verdict=%q", x.EtaRotation, x.OmegaRotation, x.AInvariant, x.BVolumeInvariant, x.FMapEquivariant, x.NormalFormGaugeCovariant, x.BasisArbitrary, x.Verdict)
}

func FormatRoutes(x RouteSourceAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s support=%.3g B=%.3g F=%.3g quat=%.3g gauge=%t same=%t", r.RouteName, r.SupportResidual, r.BVolumeResidual, r.FMapResidual, r.QuaternionicResidual, r.GaugeControlled, r.ReducesToSameSource))
	}
	return fmt.Sprintf("all=%t samePG=%t routeOnly=%t verdict=%q rows=%s", x.AllRoutesReduce, x.SamePGSourcePackage, x.RouteDependentOnly, x.Verdict, strings.Join(parts, "; "))
}

func FormatSourceTheorem(x SourceTheoremReadiness) string {
	return fmt.Sprintf("pgForces=%t gaugeControlled=%t basisFree=%t gate653=%t sourced=%t gap=%q theorem=%q verdict=%q", x.PGForcesFanoNormalForm, x.GaugeControlledSource, x.BasisFreeSourceTheorem, x.Gate653ImplicationAvailable, x.InternalMechanismSourced, x.RemainingGap, x.CandidateTheorem, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
