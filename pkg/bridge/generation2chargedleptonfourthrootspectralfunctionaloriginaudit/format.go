package generation2chargedleptonfourthrootspectralfunctionaloriginaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedTypeAudit) string {
	return fmt.Sprintf("BFlavExpr=%q BFlav=%.15g primary=%q theorem=%q gate595Verdict=%q verdict=%q", a.BFlavExpression, a.BFlav, a.PrimaryObstruction, a.RequiredTheorem, a.Gate595Verdict, a.Verdict)
}

func FormatRootFunctional(a RootFunctionalSpec) string {
	return fmt.Sprintf("input=%q eigens=%q root=%q fourier=%q wall=%q fourthRoot=%t chamber=%t observed=%t native=%t envWellDefined=%t verdict=%q", a.Input, a.EigenvalueConvention, a.RootCoordinates, a.KoideFourierForm, a.ChamberWallCoordinate, a.RequiresFourthRoot, a.RequiresOrderedChamber, a.RequiresObservedLedger, a.NativePresent, a.EnvironmentalWellDefined, a.Verdict)
}

func FormatNativeSpectralOperation(a NativeSpectralOperation) string {
	return fmt.Sprintf("op=%q admits=%t nativeLane=%t suppliesEpsilon=%t reason=%q verdict=%q", a.Operation, a.CurrentASHAAdmits, a.NativeForCurrentLane, a.SuppliesEpsilonHE, a.Reason, a.Verdict)
}

func FormatNativeSpectralAudit(a NativeSpectralAudit) string {
	items := make([]string, 0, len(a.Operations))
	for _, op := range a.Operations {
		items = append(items, FormatNativeSpectralOperation(op))
	}
	return fmt.Sprintf("poly=%t detLogPf=%t heatKernel=%t zetaEta=%t fractional=%t fourthRootTrace=%t orderedChamber=%t fourierWall=%t ops=[%s] verdict=%q", a.PolynomialAdmissible, a.DeterminantLogPfaffianAdmissible, a.HeatKernelAdmissible, a.ZetaEtaLanePresent, a.FractionalPowersNative, a.FourthRootTraceNative, a.OrderedChamberNative, a.FourierWallNative, strings.Join(items, "; "), a.Verdict)
}

func FormatRoute(a RouteAudit) string {
	return fmt.Sprintf("name=%q mechanism=%q status=%q native=%t seal=%t reason=%q verdict=%q", a.Name, a.Mechanism, a.Status, a.NativePromotion, a.BridgeSeal, a.Reason, a.Verdict)
}

func FormatRoutes(a RouteComparison) string {
	items := make([]string, 0, len(a.Routes))
	for _, r := range a.Routes {
		items = append(items, FormatRoute(r))
	}
	return fmt.Sprintf("closest=%q anyNative=%t routes=[%s] verdict=%q", a.ClosestLawfulRoute, a.AnyNativeRoute, strings.Join(items, "; "), a.Verdict)
}

func FormatSeal(a MinimalSeal) string {
	return fmt.Sprintf("name=%q components=%q mayEnterBFlav=%t native=%t reason=%q verdict=%q", a.Name, strings.Join(a.Components, ","), a.MayEnterBFlav, a.NativeLaw, a.Reason, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("fitsConstants=%t koide=%t chargedLeptonMasses=%t pmns=%t ckm=%t yukawas=%t neutrinos=%t flavorTexture=%t carrier=%t selector=%t bFlavNative=%t gate352=%t verdict=%q", a.FitsNewConstants, a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesCKM, a.DerivesYukawas, a.DerivesNeutrinos, a.DerivesFlavorTexture, a.AddsCarrier, a.AddsSelector, a.PromotesBFlavZero, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("epsilonEnv=%t nativeFourthRoot=%t closest=%q seal=%q bFlavEnv=%t theorem=%q decision=%q verdict=%q", a.EpsilonEnvironmentalWellDefined, a.NativeFourthRootPresent, a.ClosestPromotionRoute, a.MinimalSealName, a.BFlavStillEnvironmental, a.RequiredTheorem, a.Decision, a.Verdict)
}
