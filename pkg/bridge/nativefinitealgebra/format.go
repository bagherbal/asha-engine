package nativefinitealgebra

import (
	"fmt"
	"strings"
)

func FormatModes(rows []ModeRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s[index=%d kind=%s projector=%s]", r.Name, r.Index, r.Kind, r.Projector))
	}
	return strings.Join(parts, "; ")
}

func FormatSplit(a SplitAudit) string {
	return fmt.Sprintf("carrier=%q genDim=%d dimC(SC)=%d dimR(SC)=%d split=%d⊕%d projection=%t fockBookkeeping=%t fullParticleProjection=%t verdict=%s :: %s", a.Carrier, a.GeneratorDimension, a.ComplexifiedCarrierDimensionC, a.ComplexifiedCarrierDimensionR, a.LeptonLikeGeneratorCount, a.ColorLikeGeneratorCount, a.ModeLevelProjectionExists, a.ExtendsToFockBookkeeping, a.FullParticleSpeciesProjection, a.Verdict, a.ProjectionFormula)
}

func FormatCommutant(a CommutantAudit) string {
	return fmt.Sprintf("search=%q EndDimC=%d commutant=%q commDimC=%d Cpreflight=%t M3preflight=%t physicalColor=%t liftedToSC=%t maxOnSC=%t verdict=%s rows=[%s]", a.SearchSpace, a.NaiveFullEndDimensionC, a.ModeProjectionCommutant, a.ModeProjectionCommutantDimensionC, a.ComplexSingletPreflight, a.ColorMatrixAlgebraPreflight, a.M3CDerivedAsPhysicalColorGauge, a.LiftToFullExteriorRepresentation, a.MaximalAlgebraOnFullSC, a.Verdict, strings.Join(a.AmbiguityRows, " | "))
}

func FormatContact(a ContactIntegrationAudit) string {
	return fmt.Sprintf("lieInput=%q u1C=%t su2Lie=%t su2ToH=%t doubletProjection=%t leftH=%t closure=%t H=%t verdict=%s", a.DerivedLieInput, a.U1ComplexSummandPreflight, a.SU2LieAlgebraAvailable, a.SU2ToQuaternionHModuleDerived, a.DoubletProjectionDerived, a.LeftQuaternionicActionDerived, a.AssociativeClosureComputed, a.HGenerated, a.Verdict)
}

func FormatAlgebra(a AlgebraVerdictAudit) string {
	return fmt.Sprintf("importedConnes=%t CplusM3=%t H=%t exactCplusHplusM3=%t faithfulSC=%t opposite=%t orderOneReady=%t majoranaSieveReady=%t verdict=%s", a.ConnesAlgebraImported, a.CPlusM3Preflight, a.QuaternionicHDerived, a.ExactCPlusHPlusM3Derived, a.FaithfulRepresentationOnSC, a.OppositeAlgebraActionDerived, a.OrderOneCalculusReady, a.MajoranaSieveReady, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("importedConnes=%t insertedSM=%t insertedMatrices=%t insertedMassData=%t bGapMass=%t orderOne=%t smAlgClaim=%t polluted=%t verdict=%s", a.ImportedConnesAlgebra, a.InsertedSMGaugeGroup, a.InsertedGaugeMatrices, a.InsertedYukawaOrMassData, a.BGapPromotedToMass, a.ClaimedOrderOne, a.ClaimedSMAlgebraDerivation, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("split=%t CplusM3=%t u1C=%t H=%t exactSMAlg=%t orderOneReady=%t status=%s next=%q comment=%q", a.SplitDerived, a.CPlusM3Preflight, a.U1ComplexPreflight, a.QuaternionicHDerived, a.ExactSMAlgebraDerived, a.OrderOneReady, a.Status, a.NextGate, a.Comment)
}
