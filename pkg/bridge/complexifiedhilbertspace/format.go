package complexifiedhilbertspace

import (
	"fmt"
	"strings"
)

func FormatComplexification(a ComplexificationAudit) string {
	return fmt.Sprintf("source=%q derived=%t externalAdded=%t dimR_before=%d dimC_after=%d dimR_after=%d fixedReal=%d imag=%d particleSemantics=%t candidate=%t verdict=%s :: %s", a.SourceRealCarrier, a.DerivedByComplexification, a.ExternalStatesAdded, a.RealDimensionBefore, a.ComplexDimensionAfter, a.RealDimensionAfter, a.FixedRealHalfDimension, a.ImaginaryHalfDimension, a.ParticleAntiparticleSemantics, a.ParticleAntiparticleCandidate, a.Verdict, a.Formula)
}

func FormatJ(a AntiLinearJAudit) string {
	return fmt.Sprintf("candidate=%q antiLinear=%t J2=%+d Jgamma=%+d JD(if imposed)=%+d repConjugate=%t physicalChargeConj=%t KOderived=%t candidateOnly=%t verdict=%s", a.CandidateName, a.AntiLinear, a.J2Sign, a.JGammaSign, a.JDSignIfImposed, a.ExchangesRepresentationWithConjugate, a.PhysicalChargeConjugationDerived, a.KOConventionDerived, a.CandidateOnly, a.Verdict)
}

func FormatAlgebra(a FiniteAlgebraSearchAudit) string {
	return fmt.Sprintf("principle=%q importedConnes=%t LieInput=%q UEApreflight=%t explicitMatrices=%t contactRep=%t colorSplit=%t M3C=%t H=%t C=%t maximalAssoc=%t faithfulDoubled=%t opposite=%t orderOneReady=%t verdict=%s rows=[%s]", a.SearchPrinciple, a.ImportedConnesAlgebra, a.DerivedLieAlgebraInput, a.UniversalEnvelopingPreflight, a.ExplicitGaugeMatricesAvailable, a.ContactPreservingRepresentationAvailable, a.ColorLeptonSplitAvailable, a.ColorM3CDerived, a.QuaternionHDerived, a.ComplexSummandDerived, a.MaximalAssociativeSubalgebraDerived, a.FaithfulDoubledRepresentation, a.OppositeAlgebraActionDerived, a.OrderOneReady, a.Verdict, strings.Join(a.CandidateRows, "; "))
}

func FormatMajorana(a MajoranaBilinearAudit) string {
	return fmt.Sprintf("doubled=%t neutralParticle=%v neutralConjugate=%v capacity=%t slots=%d rhNuDerived=%t kinAllowed=%t neutralGaugeOK=%t gradingDerived=%t orderOneDerived=%t verdict=%s", a.DoubledSpaceAvailable, a.NeutralParticleStates, a.NeutralConjugateStates, a.NeutralBilinearCapacity, a.TotallyNeutralSlotCount, a.RHNeutrinoSlotDerived, a.MajoranaTermKinematicallyAllowed, a.GaugeInvariantIfNeutral, a.GradingCompatibilityDerived, a.OrderOneCompatibilityDerived, a.Verdict)
}

func FormatBGap(a BGapIdentificationAudit) string {
	return fmt.Sprintf("available=%t Bgap=%.12g slotExists=%t dimensionless=%t diagnosticInserted=%t canonicalMajorana=%t promotedMass=%t selectsRHNu=%t needsAlgebra=%t needsScaleMap=%t verdict=%s", a.BGapAvailable, a.BGap, a.CandidateMajoranaSlotExists, a.BGapDimensionless, a.BGapInsertedAsDiagnostic, a.BGapCanonicalMajoranaEntry, a.BGapPromotedToMass, a.BGapSelectsRHNeutrino, a.RequiresAlgebraRepresentation, a.RequiresScaleMap, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("externalAntiparticles=%t importedConnes=%t continuumMass=%t VEV=%t MB=%t MInt=%t MStar=%t BgapMass=%t MajoranaClaim=%t orderOneClaim=%t PMNSYukawa=%t polluted=%t :: %s", a.ExternalAntiparticlesAdded, a.ConnesAlgebraImported, a.ContinuumMassInserted, a.VEVInserted, a.MBInserted, a.MIntInserted, a.MStarInserted, a.BGapPromotedToMass, a.MajoranaMassClaimed, a.OrderOneClaimed, a.PMNSOrYukawaClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("complexification=%t J=%t nativeAlgebra=%t MajoranaCapacity=%t BgapMajorana=%t fullTriple=%t next=%q status=%q :: %s", a.ComplexificationDerived, a.AntiLinearJAvailable, a.NativeAlgebraDerived, a.MajoranaCapacity, a.BGapMajoranaIdentified, a.FullSpectralTripleDerived, a.NextGate, a.Status, a.Comment)
}
