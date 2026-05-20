package generation2internalobstructionsealclosurepivot

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate656Inheritance) string {
	return fmt.Sprintf("halfTrace=%t clue=%t numerator=%t noHalfMap=%t no144=%t no72=%t noStressK7=%t noStress=%t noHistory=%t noScalarFlavor=%t claimsStress=%t claims144=%t claims72=%t claimsHistory=%t claimsScalarFlavor=%t firewall=%t verdict=%q", x.HalfTraceAudited, x.HalfTraceTypedClue, x.FanoNumeratorStrengthened, x.NoNativeHalfTraceMap, x.NoSevenOver144Theorem, x.NoSevenOver72Theorem, x.NoBoundaryStressFromK7, x.NoBoundaryStressDerived, x.NoHistoryLoopSource, x.NoScalarFlavorMap, x.ClaimsBoundaryStress, x.ClaimsSevenOver144, x.ClaimsSevenOver72, x.ClaimsHistoryLoopUnit, x.ClaimsScalarFlavor, x.Gate656Firewall, x.Verdict)
}

func FormatClosure(x RouteClosureAudit) string {
	return fmt.Sprintf("lane=%q internal=%t boundaryFailed=%t physicsBlocked=%t requiresPsi=%t seal=%q boundary=%q class=%q verdict=%q", x.Lane, x.InternalTheoremPathMature, x.BoundaryInterfaceFailed, x.PhysicsPromotionBlocked, x.FutureUseRequiresExplicitPsi, x.SealName, x.BoundaryStatus, x.Classification, x.Verdict)
}

func FormatActive(x ActiveSealLedger) string {
	parts := make([]string, 0, len(x.Seals))
	for _, s := range x.Seals {
		parts = append(parts, fmt.Sprintf("#%d %s active=%t formula=%q status=%q requires=%q next=%q", s.Rank, s.Name, s.Active, s.Formula, s.Status, s.Requires, s.NextUse))
	}
	return fmt.Sprintf("activeCount=%d primary=%q xi=%.15g absLambda=%.15g r3=%.15g L=%.15g verdict=%q seals=%s", x.ActiveCount, x.Primary, x.XiBoundary, x.AbsLambda, x.R3Minus1, x.HistoryLoopUnit, x.Verdict, strings.Join(parts, "; "))
}

func FormatInactive(x InactiveLaneAudit) string {
	parts := make([]string, 0, len(x.Lanes))
	for _, l := range x.Lanes {
		parts = append(parts, fmt.Sprintf("%s active=%t class=%q reactivate=%q reason=%q", l.Name, l.Active, l.Classification, l.ReactivateOnlyIf, l.Reason))
	}
	return fmt.Sprintf("fano=%t half=%t k7=%t hodge=%t splitG2=%t verdict=%q lanes=%s", x.FanoHitchinInactive, x.HalfTraceInactive, x.K7TraceInactive, x.HodgeK7W7Inactive, x.SplitG2Inactive, x.Verdict, strings.Join(parts, "; "))
}

func FormatRanking(x NextActionRanking) string {
	parts := make([]string, 0, len(x.Actions))
	for _, a := range x.Actions {
		parts = append(parts, fmt.Sprintf("#%d %s actionable=%t reason=%q touches=%s", a.Rank, a.Path, a.Actionable, a.Reason, strings.Join(a.Touches, ",")))
	}
	return fmt.Sprintf("primary=%q secondary=%q k7Low=%t verdict=%q actions=%s", x.PrimaryPath, x.SecondaryPath, x.K7BoundaryLow, x.Verdict, strings.Join(parts, "; "))
}

func FormatStrategic(x StrategicVerdict) string {
	return fmt.Sprintf("pivot=%q stopFano=%t returnTransport=%t boundary=%t scalar=%t history=%t flavor=%t k7Blocked=%t verdict=%q", x.RecommendedPivot, x.StopFanoBoundaryLane, x.ReturnToTransport, x.BoundaryStressLive, x.ScalarMatchingLive, x.HistoryLoopLive, x.FlavorOrientationLive, x.K7BoundaryBlocked, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("boundaryDerived=%t scalarRG=%t higgs=%t gauge=%t flavor=%t spacetime=%t splitG2=%t seven72=%t fanoBoundary=%t verdict=%q", x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGDerived, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerived, x.ClaimsPhysicalSpacetime, x.ClaimsSplitG2, x.ClaimsSevenOver72Theorem, x.ClaimsFanoBoundaryInterface, x.Verdict)
}
