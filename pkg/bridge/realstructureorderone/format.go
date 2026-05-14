package realstructureorderone

import (
	"fmt"
	"strings"
)

func FormatJ(a JConstructionAudit) string {
	return fmt.Sprintf("candidate=%s constructed=%t dim=%d J2Residual=%.3g J2=%+d commutesGamma=%t Jgamma=%+d antiunitaryDerived=%t physicalChargeConjugation=%t doubled=%t verdict=%s :: %s", a.CandidateName, a.Constructed, a.Dimension, a.J2Residual, a.J2Sign, a.CommutesWithOccupationGamma, a.JGammaSign, a.AntiUnitaryComplexPartDerived, a.PhysicalChargeConjugation, a.ParticleAntiparticleDoubling, a.Verdict, a.Formula)
}

func FormatKO(a KOSignAudit) string {
	return fmt.Sprintf("J2=%+d Jgamma=%+d JD(if imposed)=%+d candidate=%q conventionDerived=%t JDgeneric=%t JDunit=%t JDbgapi=%t params=%d->%d promoted=%t verdict=%s", a.J2Epsilon, a.JGammaEpsilon, a.JDSignIfImposed, a.CandidateKODimension, a.KOConventionDerived, a.JDCommutesForGenericM, a.JDCommutesForUnitBlock, a.JDCommutesForBGapIBlock, a.FreeParametersBefore, a.FreeParametersAfterJ, a.PromotedKOTheorem, a.Verdict)
}

func FormatJReality(a JRealitySieve) string {
	return fmt.Sprintf("initial=%d orbits=%d afterJ=%d reduction=%.3f canonicalBlock=%t colorWeakSplit=%t physicalChirality=%t verdict=%s :: %s", a.InitialParameters, a.OrbitsUnderComplement, a.ParametersAfterReality, a.ReductionFraction, a.CanonicalBlockSelected, a.ColorWeakSubblocksDerived, a.PhysicalChiralityDerived, a.Verdict, a.ConstraintFormula)
}

func FormatAlgebra(a AlgebraRepresentationAudit) string {
	return fmt.Sprintf("candidates=%d faithfulTotal=%t physicalSMAlgebra=%t diagonalTried=%t BLTried=%t nontrivialOneForms=%t verdict=%s rows=[%s]", a.AlgebraCandidatesAudited, a.FaithfulTotalRepresentation, a.PhysicalSMAlgebraDerived, a.DiagonalOccupationAlgebraTried, a.BLAlgebraTried, a.NontrivialOneFormsDerived, a.Verdict, strings.Join(a.RepresentationRows, "; "))
}

func FormatOrderOne(a OrderOnePreflight) string {
	return fmt.Sprintf("formula=%q testable=%t fullAlgebra=%t Jcanonical=%t nonVacuous=%t BLallowed=%d fullDiagAllowed=%d JRealityAllowed=%d splitsColorWeak=%t verified=%t vacuous=%t promotable=%t verdict=%s", a.OrderOneFormula, a.TestableWithCurrentData, a.FullAlgebraRepresentation, a.RealStructureCanonical, a.NonVacuousCommutatorsAvailable, a.ProvisionalDiagonalBLAllowed, a.ProvisionalFullDiagonalAllowed, a.JRealityAllowed, a.SplitsColorWeakSubblocks, a.OrderOneVerified, a.OrderOneVacuous, a.PromotableFiniteDirac, a.Verdict)
}

func FormatBGap(a BGapMajoranaSieve) string {
	return fmt.Sprintf("available=%t Bgap=%.12g sterileCandidates=%d neutralMasks=%v rhNuSlot=%t doubled=%t majoranaSpace=%t diagnosticScalar=%t canonicalMajorana=%t forcedNeutral=%t promotedMass=%t broaderHilbert=%t verdict=%s", a.BGapAvailable, a.BGap, a.SterileVacuumCandidates, a.CandidateNeutralMasks, a.RightHandedNeutrinoSlotDerived, a.ParticleAntiparticleDoubling, a.MajoranaBilinearSpaceAvailable, a.BGapAllowedAsDiagnosticScalar, a.BGapCanonicalMajoranaEntry, a.BGapForcedToNeutralSector, a.BGapPromotedToMajoranaMass, a.RequiresBroaderHilbertSpace, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("continuumMass=%t VEV=%t MB=%t Mstar=%t observedMass=%t BgapMass=%t DFfit=%t KOtheorem=%t orderOneClaim=%t PMNSYukawa=%t polluted=%t :: %s", a.ContinuumMassInserted, a.VEVInserted, a.MBInserted, a.MStarInserted, a.ObservedFermionMassInserted, a.BGapPromotedToMass, a.DFChosenByFit, a.KOClaimedAsTheorem, a.OrderOneClaimed, a.PMNSOrYukawaClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("J=%t KOsigns=%t Jreduces=%t orderOne=%t BgapMajorana=%t canonicalDF=%t next=%q status=%q :: %s", a.CandidateJAvailable, a.CandidateKOSignsComputed, a.JRealityReducesParameters, a.OrderOneDerived, a.BGapMajoranaPlacement, a.CanonicalDFDerived, a.NextGate, a.Status, a.Comment)
}
