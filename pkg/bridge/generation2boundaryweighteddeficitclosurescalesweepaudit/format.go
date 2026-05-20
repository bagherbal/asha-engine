package generation2boundaryweighteddeficitclosurescalesweepaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate661Inheritance) string {
	return fmt.Sprintf("inherited=%t Ksum=%.15g W72=%.15g E72=%.15g lambda12Only=%t formulaCircular=%t no7=%t noTransport=%t noEndpoint=%t firewall=%t verdict=%q", x.ClosureInherited, x.KSum, x.W72, x.E72, x.Lambda12OnlyComputed, x.FormulaLiftCircular, x.NoNativeSevenOver72, x.NoNativeTransport, x.NoIndependentEndpoint, x.FirewallPreserved, x.Verdict)
}

func FormatSeed(x TransportSeed) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g t12=%.15g t13=%.15g t23=%.15g tgeom=%.15g g1=%.12g gy=%.12g g2=%.12g g3=%.12g lambdaMZ=%.15g init=%d verdict=%q", x.Mu0GeV, x.Lambda12GeV, x.T12, x.T13, x.T23, x.TGeom, x.G1MZ, x.GYMZ, x.G2MZ, x.G3MZ, x.LambdaMZ, len(x.InitialVector), x.Verdict)
}

func FormatScaleRow(x ScaleSweepRow) string {
	return fmt.Sprintf("%s mu=%.12g t=%.12g |lambda|=%.15g rEW=%.15g pair=%.15g W_EW=%.15g E_EW=%.15g W_pair=%.15g E_pair=%.15g gauge=%q", x.Name, x.MuGeV, x.T, x.AbsLambda, x.GaugeResidualEWMean, x.PairResidual, x.W72EWMean, x.E72EWMean, x.W72Pair, x.E72Pair, x.GaugeDefinition)
}

func FormatScaleSweep(x ScaleSweepAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatScaleRow(r))
	}
	return fmt.Sprintf("rows=%d bestEW=%q bestEWResidual=%.15g bestPair=%q bestPairResidual=%.15g lambda12MinEW=%t lambda12MinPair=%t verdict=%q ledger=[%s]", len(x.Rows), x.BestEWMeanScale, x.BestEWMeanResidual, x.BestPairScale, x.BestPairResidual, x.Lambda12UniquelyMinimalEW, x.Lambda12UniquelyMinimalPair, x.Verdict, strings.Join(rows, "; "))
}

func FormatLocalRow(x LocalPerturbationRow) string {
	return fmt.Sprintf("dlog=%.3g mu=%.12g |lambda|=%.15g rEW=%.15g W72=%.15g E72=%.15g abs=%.15g", x.DeltaLog, x.MuGeV, x.AbsLambda, x.GaugeResidualEWMean, x.W72, x.E72, x.AbsE72)
}

func FormatLocalPerturbation(x LocalPerturbationAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatLocalRow(r))
	}
	return fmt.Sprintf("rows=%d minDLog=%.15g minAbs=%.15g width1e-4=%.15g slope=%.15g selectsLambda12=%t verdict=%q ledger=[%s]", len(x.Rows), x.MinimumDeltaLog, x.MinimumAbsResidual, x.Threshold1eMinus4Width, x.FiniteDifferenceSlope, x.LocalGridSelectsLambda12, x.Verdict, strings.Join(rows, "; "))
}

func FormatWeight(x WeightSensitivityAudit) string {
	return fmt.Sprintf("wBest=%.15g diff7over72=%.15g wBestOrient=%.15g diffOrient=%.15g exactResidual=%.15g orientResidual=%.15g exactNear=%t orientNear=%t verdict=%q", x.WBestExact, x.WBestExactMinus7Over72, x.WBestOrientation, x.WBestOrientationMinus7Over72, x.ExactCandidateResidual, x.OrientationCandidateResidual, x.ExactWeightNear7Over72, x.OrientationWeightNear7Over72, x.Verdict)
}

func FormatJacobian(x InputJacobianAudit) string {
	return fmt.Sprintf("dE/dkappaE=%.15g dE/dAbsLambda=%.15g dE/dR3=%.15g dk/dlambdaRuntime=%.15g dk/dlambdaProxy=%.15g dk/dL=%.15g L=%.15g rho=%.15g notes=[%s] verdict=%q", x.DE_DKappaE, x.DE_DAbsLambda, x.DE_DR3Minus1, x.DKappa_DLambdaRuntime, x.DKappa_DLambdaProxy, x.DKappa_DL, x.L, x.RhoLambdaMatch, strings.Join(x.Notes, "; "), x.Verdict)
}

func FormatOrientation(x OrientationScaleAudit) string {
	return fmt.Sprintf("kappaExact=%.15g kappaOrient=%.15g exactE=%.15g orientE=%.15g wBest=%.15g wBestOrient=%.15g weightShift=%.15g amplification=%.15g verdict=%q", x.KappaEExact, x.KappaEOrientation, x.ExactE72AtLambda12, x.OrientationE72AtLambda12, x.ExactWBest, x.OrientationWBest, x.BestWeightShift, x.ClosureResidualAmplification, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsScale=%t claims7=%t claimsUncertainty=%t claimsTransport=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsGauge=%t claimsCKM=%t verdict=%q", x.ClaimsNativeScaleSelection, x.ClaimsNativeSevenOver72Theorem, x.ClaimsFullUncertaintyPropagation, x.ClaimsNativeTransportTheorem, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
