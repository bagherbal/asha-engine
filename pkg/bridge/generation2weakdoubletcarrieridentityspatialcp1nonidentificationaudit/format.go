package generation2weakdoubletcarrieridentityspatialcp1nonidentificationaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate575Audit) string {
	return fmt.Sprintf("split=%t commBL=%t imHAction=%t fstCarrier=%t physicalWeak=%t flavorEW=%t required=%q verdict=%q", a.SealedCP1SplitExists, a.CommutesWithBMinusL, a.CarriesImHAction, a.PartOfFiniteWeakCarrier, a.CanBePhysicalWeakPlane, a.DerivesFlavorOrEWObservedData, a.AdditionalTheoremRequired, a.Verdict)
}

func FormatFiniteAlgebra(a FiniteAlgebraAudit) string {
	return fmt.Sprintf("A=%q C=%q H=%q M3=%q preU=%q gauge=%q weakSocket=%q imH=%q structural=%t dynamics=%t verdict=%q", a.Algebra, a.ComplexSummand, a.QuaternionicSummand, a.ColorSummand, a.PreUnimodularUnitary, a.UnimodularGaugeGroup, a.WeakSocketSource, a.ImHLieAlgebra, a.StructuralOnly, a.AbsoluteDynamicsDerived, a.Verdict)
}

func FormatWeakFermions(a WeakFermionCarrierInventory) string {
	items := []string{}
	for _, c := range a.Carriers {
		items = append(items, fmt.Sprintf("%s(dim=%d,copies=%d,weak=%s,right=%s,H=%s,color=%s)", c.Name, c.ComplexDimension, c.WeakDoubletCopies, c.WeakModule, c.RightModule, c.HAction, c.ColorBehavior))
	}
	return fmt.Sprintf("carriers=[%s] LL=%t QL=%t HLL=%t HQL=%t qColor=%d colorSource=%q colorIsWeak=%t cp1Used=%t finite=%t verdict=%q", strings.Join(items, "; "), a.LLPresent, a.QLPresent, a.HActsOnLL, a.HActsOnQL, a.QLColorMultiplicity, a.ColorActionSource, a.ColorIsWeakStructure, a.SealedSpatialCP1Used, a.FiniteWeakDoubletsAvailable, a.Verdict)
}

func FormatScalarDoublet(a ScalarDoubletInventory) string {
	return fmt.Sprintf("name=%q carrier=%q dimC=%d dimR=%d source=%q H=%t oneForm=%t separateW=%t separateUperp=%t cp1Used=%t dynamics=%t verdict=%q", a.CarrierName, a.Carrier, a.ComplexDimension, a.RealDimension, a.Source, a.HActionStructural, a.FromFiniteOneFormLane, a.SeparateFromWSpatial, a.SeparateFromUperp, a.SealedSpatialCP1Used, a.NumericalHiggsDynamicsDerived, a.Verdict)
}

func FormatSealedCompare(a SealedSpatialCP1ComparisonAudit) string {
	return fmt.Sprintf("seal=%q uperp=%q split=%t AF=%t DF=%t J=%t grading=%t firstOrder=%t oneForm=%t weakCarrier=%t verdict=%q", a.SealName, a.UperpDescription, a.CP1SplitExistsAlgebraically, a.AppearsInAFRepresentation, a.AppearsInDFEdges, a.AppearsInJ, a.AppearsInGrading, a.AppearsInFirstOrder, a.AppearsInOneFormHiggsLane, a.IsFiniteWeakCarrier, a.Verdict)
}

func FormatWeakCount(a WeakDoubletCountAudit) string {
	return fmt.Sprintf("lepton=%d quarkColors=%d quarkDoublets=%d total=%d su2Index=%q pattern=%q fromColor=%t fromSpatialCP1=%t verdict=%q", a.LeptonWeakDoublets, a.QuarkColorCopies, a.QuarkWeakDoublets, a.TotalWeakDoublets, a.Gate298SU2Index, a.OnePlusThreePattern, a.ComesFromColorMultiplicity, a.ComesFromSpatialCP1Selection, a.Verdict)
}

func FormatEdgeLane(a EdgeLaneRelationAudit) string {
	return fmt.Sprintf("edges=%v canonical=%t sealedSelector=%t uperp=%t hphi=%t firstOrder=%t verdict=%q", a.Edges, a.CanonicalEdgesReconfirmed, a.UsesSealedSpatialSelector, a.UsesUperpCarrier, a.UsesHPhiScalarLane, a.FirstOrderCompatible, a.Verdict)
}

func FormatNonIdentity(a NonIdentificationAudit) string {
	return fmt.Sprintf("uperp=Hphi?%t uperp=LL?%t uperp=QL?%t uperp=ImH?%t distinct=%v required=%q certified=%t verdict=%q", a.UperpEqualsHPhi, a.UperpEqualsLL, a.UperpEqualsQL, a.UperpEqualsImH, a.DistinctCarriers, a.NewFunctorRequired, a.Certified, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("weakPlane=%t weakIsospinFromCP1=%t WZPhoton=%t masses=%t generation=%t yukawa=%t ckmPmns=%t observed=%t gate564565=%t k7Time=%t verdict=%q", a.PhysicalWeakPlaneDerived, a.WeakIsospinDerivedFromCP1, a.WZPhotonDynamicsDerived, a.MassesDerived, a.GenerationHierarchy, a.YukawaTexture, a.CKMPMNS, a.ObservedFlavorData, a.Gate564565Preserved, a.K7TimePreserved, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("weakSocket=%q carriers=%v Hphi=%t cp1WeakCarrier=%t onePlusThreeColor=%t physicalData=%t required=%q verdict=%q", a.WeakSocketLocation, a.ActualWeakDoubletCarriers, a.HPhiIsScalarWeakDoublet, a.SealedSpatialCP1IsWeakCarrier, a.WeakDoubletOnePlusThreeFromColor, a.DerivesPhysicalWeakFlavorEWData, a.AdditionalTheoremRequired, a.Verdict)
}
