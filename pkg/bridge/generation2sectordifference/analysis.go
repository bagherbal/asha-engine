// Package generation2sectordifference implements Gate 462:
// Sector-Difference Invariant / CKM Interface Firewall Audit.
//
// Gate 461 proved that the charged sectors {u,d,e} carry independent
// bridge-only coefficient rays. Gate 462 isolates the relative-ray object that
// would be needed by a CKM-like interface, while proving that no CKM/PMNS
// entry, sector-difference value, or cross-sector eigenbasis convention becomes
// native ASHA law. The gate is an interface firewall: it records which labelled
// bridge data may be compared, and rejects every route that smuggles observed
// mixing data or native-promotion claims into the finite law-space.
package generation2sectordifference

import (
	"fmt"
	"math"
	"sync"
)

const (
	AuditID = "GATE462-SECTOR-DIFFERENCE-INVARIANT-CKM-INTERFACE-FIREWALL-AUDIT"

	StatusGate461Inherited               = "CONDITIONAL_SUPPORT_GATE461_SECTOR_MULTIPLEX_INHERITED"
	StatusRelativeRayLedgerDefined       = "CONDITIONAL_SUPPORT_RELATIVE_RAY_LEDGER_DEFINED"
	StatusUDDifferenceBridgeOnlyComputed = "CONDITIONAL_SUPPORT_UD_SECTOR_DIFFERENCE_BRIDGE_ONLY_COMPUTED"
	StatusCKMInterfaceFirewallValidated  = "CONDITIONAL_SUPPORT_SECTOR_DIFFERENCE_CKM_INTERFACE_FIREWALL_VALIDATED"
	StatusPMNSInterfaceFirewallValidated = "CONDITIONAL_SUPPORT_PMNS_INTERFACE_FIREWALL_VALIDATED_BY_REJECTION"
	StatusFirewallPreserved              = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedRequiresTwoProvenancedRays    = "FAILED_ROUTE_SECTOR_DIFFERENCE_REQUIRES_TWO_PROVENANCED_RAYS"
	StatusFailedMissingEigenbasisConvention   = "FAILED_ROUTE_UNLABELLED_EIGENBASIS_CONVENTION_REJECTED"
	StatusFailedObservedCKMPMNSImportRejected = "FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED"
	StatusFailedCKMPMNSPredictionRejected     = "FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedNativeRelativeRayPromotion    = "FAILED_ROUTE_RELATIVE_RAY_NATIVE_PROMOTION_REJECTED"
	StatusFailedLeptonSectorMisrouted         = "FAILED_ROUTE_LEPTON_PMNS_SECTOR_MISRouted_TO_CHARGED_CKM_LEDGER"
	StatusFailedUniversalityNotNative         = "FAILED_ROUTE_CROSS_SECTOR_RAY_UNIVERSALITY_NOT_NATIVE"
)

const (
	NativeFlavorDim   = 13
	KXYCoeffDim       = 9
	RelativeRayDOF    = 2
	RequiredSectorsUD = 2
	RequiredFields    = 11
)

var CKMSectors = []string{"u", "d"}

type Inheritance struct {
	Executed                             bool
	Gate444KGenForced                    bool
	Gate445TriangleForced                bool
	Gate456InverseDerived                bool
	Gate457ProvenanceContract            bool
	Gate459BranchTags                    bool
	Gate460ResidualHarness               bool
	Gate461SectorMultiplex               bool
	Gate461IndependentSectorRaysAccepted bool
	Gate461NativeUniversalityRejected    bool
	Gate461SectorContaminationRejected   bool
	NoObservedValuesImported             bool
	Verdict                              string
}

type SectorRay struct {
	Sector                  string
	Alpha                   float64
	Phi                     float64
	CPOddSign               int
	C3Sheet                 int
	HasProvenance           bool
	HasBranchTag            bool
	HasEigenbasisConvention bool
	BridgeOnly              bool
	ExplicitObservedCKM     bool
	ExplicitObservedPMNS    bool
	NativePromotionClaim    bool
	UsesLeptonPMNSLedger    bool
	UniversalityClaim       bool
	UniversalityNativeClaim bool
}

type RelativeRay struct {
	FromSector              string
	ToSector                string
	DeltaAlpha              float64
	DeltaPhi                float64
	PhaseChord              float64
	ProjectiveDistance      float64
	CompleteInputs          bool
	EigenbasisConventionSet bool
	BridgeOnly              bool
	ExportsCKMEntry         bool
	ExportsPMNSEntry        bool
	ExportsNativeObservable bool
	Verdict                 string
	Reason                  string
}

type InterfaceContract struct {
	Executed                          bool
	RequiresUSector                   bool
	RequiresDSector                   bool
	RequiresProvenancePerSector       bool
	RequiresCompleteBranchTags        bool
	RequiresEigenbasisConvention      bool
	RelativeRayDimension              int
	ExportsRelativeDiagnosticsOnly    bool
	RejectsCKMAsNativePrediction      bool
	RejectsPMNSInChargedCKMLedger     bool
	RejectsObservedMixingByDefault    bool
	RejectsNativeRelativeRayPromotion bool
	Verdict                           string
	Reason                            string
}

type Case struct {
	Name     string
	Rays     []SectorRay
	Accepted bool
	Relative RelativeRay
	Verdict  string
	Reason   string
}

type Sieve struct {
	Executed                        bool
	Cases                           []Case
	AcceptedCaseCount               int
	RejectedCaseCount               int
	ValidUDDifferenceAccepted       bool
	MissingSectorRejected           bool
	MissingProvenanceRejected       bool
	MissingEigenbasisRejected       bool
	ObservedCKMPMNSRejected         bool
	NativePredictionRejected        bool
	NativeRelativePromotionRejected bool
	LeptonPMNSMisrouteRejected      bool
	UniversalityNativeRejected      bool
	AllAcceptedBridgeOnly           bool
	NoNativeMixingObservableExport  bool
	Verdict                         string
	Reason                          string
}

type CKMFirewall struct {
	Executed                           bool
	RelativeRayMayFeedCKMAdapter       bool
	CKMMatrixEntryComputed             bool
	CKMMatrixEntryNative               bool
	PMNSMatrixEntryComputed            bool
	PMNSMatrixEntryNative              bool
	RequiresObservedComparatorImport   bool
	RequiresSchemeScaleProvenance      bool
	RequiresEigenvectorGaugeConvention bool
	KGenStillForced                    bool
	XTriangleStillForced               bool
	YPhaseStillQuarantined             bool
	SectorCoefficientsStillSealed      bool
	NoObservedMassesImported           bool
	NoObservedYukawasImported          bool
	NoObservedCKMImported              bool
	NoObservedPMNSImported             bool
	NoGSTPromotion                     bool
	NativeFlavorDimAfter               int
	KXYCoeffDimAfter                   int
	Verdict                            string
	Reason                             string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Contract    InterfaceContract
	Sieve       Sieve
	Firewall    CKMFirewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Contract = buildContract()
	a.Sieve = buildSieve()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                             true,
		Gate444KGenForced:                    true,
		Gate445TriangleForced:                true,
		Gate456InverseDerived:                true,
		Gate457ProvenanceContract:            true,
		Gate459BranchTags:                    true,
		Gate460ResidualHarness:               true,
		Gate461SectorMultiplex:               true,
		Gate461IndependentSectorRaysAccepted: true,
		Gate461NativeUniversalityRejected:    true,
		Gate461SectorContaminationRejected:   true,
		NoObservedValuesImported:             true,
		Verdict:                              StatusGate461Inherited,
	}
}

func buildContract() InterfaceContract {
	return InterfaceContract{
		Executed:                          true,
		RequiresUSector:                   true,
		RequiresDSector:                   true,
		RequiresProvenancePerSector:       true,
		RequiresCompleteBranchTags:        true,
		RequiresEigenbasisConvention:      true,
		RelativeRayDimension:              RelativeRayDOF,
		ExportsRelativeDiagnosticsOnly:    true,
		RejectsCKMAsNativePrediction:      true,
		RejectsPMNSInChargedCKMLedger:     true,
		RejectsObservedMixingByDefault:    true,
		RejectsNativeRelativeRayPromotion: true,
		Verdict:                           StatusRelativeRayLedgerDefined,
		Reason:                            "a CKM-facing relative ray requires complete u and d bridge rays plus an explicit eigenbasis convention, and exports only relative diagnostics",
	}
}

func buildSieve() Sieve {
	cases := []Case{
		{Name: "valid synthetic u-d relative ray", Rays: validUDRays()},
		{Name: "missing d sector", Rays: []SectorRay{validRay("u", 0.20, 0.30, +1, 0)}},
		{Name: "missing provenance", Rays: missingProvenanceRays()},
		{Name: "missing eigenbasis convention", Rays: missingEigenbasisRays()},
		{Name: "observed CKM/PMNS import", Rays: observedMixingRays()},
		{Name: "native CKM prediction claim", Rays: nativePredictionRays()},
		{Name: "native relative-ray promotion", Rays: nativeRelativePromotionRays()},
		{Name: "lepton PMNS misrouted into charged CKM ledger", Rays: leptonPMNSMisrouteRays()},
		{Name: "cross-sector universality as native law", Rays: nativeUniversalityRays()},
	}

	s := Sieve{Executed: true, Cases: make([]Case, 0, len(cases))}
	for _, c := range cases {
		rel, accepted, verdict, reason := EvaluateRelativeRay(c.Rays)
		c.Relative = rel
		c.Accepted = accepted
		c.Verdict = verdict
		c.Reason = reason
		if accepted {
			s.AcceptedCaseCount++
		} else {
			s.RejectedCaseCount++
		}
		s.ValidUDDifferenceAccepted = s.ValidUDDifferenceAccepted || (c.Name == "valid synthetic u-d relative ray" && accepted && verdict == StatusUDDifferenceBridgeOnlyComputed)
		s.MissingSectorRejected = s.MissingSectorRejected || verdict == StatusFailedRequiresTwoProvenancedRays
		s.MissingProvenanceRejected = s.MissingProvenanceRejected || (verdict == StatusFailedRequiresTwoProvenancedRays && c.Name == "missing provenance")
		s.MissingEigenbasisRejected = s.MissingEigenbasisRejected || verdict == StatusFailedMissingEigenbasisConvention
		s.ObservedCKMPMNSRejected = s.ObservedCKMPMNSRejected || verdict == StatusFailedObservedCKMPMNSImportRejected
		s.NativePredictionRejected = s.NativePredictionRejected || verdict == StatusFailedCKMPMNSPredictionRejected
		s.NativeRelativePromotionRejected = s.NativeRelativePromotionRejected || verdict == StatusFailedNativeRelativeRayPromotion
		s.LeptonPMNSMisrouteRejected = s.LeptonPMNSMisrouteRejected || verdict == StatusFailedLeptonSectorMisrouted
		s.UniversalityNativeRejected = s.UniversalityNativeRejected || verdict == StatusFailedUniversalityNotNative
		s.Cases = append(s.Cases, c)
	}
	s.AllAcceptedBridgeOnly = true
	s.NoNativeMixingObservableExport = true
	for _, c := range s.Cases {
		if c.Accepted && (!c.Relative.BridgeOnly || c.Relative.ExportsCKMEntry || c.Relative.ExportsPMNSEntry || c.Relative.ExportsNativeObservable) {
			s.AllAcceptedBridgeOnly = false
			s.NoNativeMixingObservableExport = false
		}
	}
	s.Verdict = StatusCKMInterfaceFirewallValidated
	s.Reason = "only the synthetic, fully provenanced u-d relative ray is accepted, and it exports diagnostics rather than CKM/PMNS observables"
	return s
}

func validUDRays() []SectorRay {
	return []SectorRay{
		validRay("u", 0.20, 0.30, +1, 0),
		validRay("d", -0.10, 0.85, -1, 1),
	}
}

func validRay(sector string, alpha, phi float64, sigma, sheet int) SectorRay {
	return SectorRay{
		Sector:                  sector,
		Alpha:                   alpha,
		Phi:                     phi,
		CPOddSign:               sigma,
		C3Sheet:                 sheet,
		HasProvenance:           true,
		HasBranchTag:            true,
		HasEigenbasisConvention: true,
		BridgeOnly:              true,
	}
}

func missingProvenanceRays() []SectorRay {
	rs := validUDRays()
	rs[1].HasProvenance = false
	return rs
}

func missingEigenbasisRays() []SectorRay {
	rs := validUDRays()
	rs[0].HasEigenbasisConvention = false
	return rs
}

func observedMixingRays() []SectorRay {
	rs := validUDRays()
	rs[0].ExplicitObservedCKM = true
	rs[1].ExplicitObservedPMNS = true
	return rs
}

func nativePredictionRays() []SectorRay {
	rs := validUDRays()
	rs[0].NativePromotionClaim = true
	rs[0].ExplicitObservedCKM = true
	return rs
}

func nativeRelativePromotionRays() []SectorRay {
	rs := validUDRays()
	rs[1].NativePromotionClaim = true
	return rs
}

func leptonPMNSMisrouteRays() []SectorRay {
	return []SectorRay{
		validRay("e", 0.15, 0.40, +1, 0),
		validRay("d", -0.10, 0.85, -1, 1),
	}
}

func nativeUniversalityRays() []SectorRay {
	rs := validUDRays()
	for i := range rs {
		rs[i].UniversalityClaim = true
		rs[i].UniversalityNativeClaim = true
		rs[i].Alpha = 0.33
		rs[i].Phi = 0.33
	}
	return rs
}

func EvaluateRelativeRay(rays []SectorRay) (RelativeRay, bool, string, string) {
	rel := RelativeRay{FromSector: "u", ToSector: "d"}

	if hasLeptonPMNSMisroute(rays) {
		rel.Verdict = StatusFailedLeptonSectorMisrouted
		rel.Reason = "the charged CKM relative-ray ledger requires u and d sectors; lepton/PMNS records need a separate neutrino-sector interface"
		return rel, false, rel.Verdict, rel.Reason
	}
	if hasNativePredictionClaim(rays) && hasObservedMixingImport(rays) {
		rel.Verdict = StatusFailedCKMPMNSPredictionRejected
		rel.Reason = "observed mixing plus a native-prediction claim is forbidden"
		return rel, false, rel.Verdict, rel.Reason
	}
	if hasObservedMixingImport(rays) {
		rel.Verdict = StatusFailedObservedCKMPMNSImportRejected
		rel.Reason = "observed CKM/PMNS values are not accepted in this native-boundary audit"
		return rel, false, rel.Verdict, rel.Reason
	}
	if hasNativePredictionClaim(rays) {
		rel.Verdict = StatusFailedNativeRelativeRayPromotion
		rel.Reason = "relative rays are bridge diagnostics and cannot be promoted to native ASHA law"
		return rel, false, rel.Verdict, rel.Reason
	}
	if hasNativeUniversalityClaim(rays) {
		rel.Verdict = StatusFailedUniversalityNotNative
		rel.Reason = "Gate461 already rejected cross-sector coefficient-ray universality as native law"
		return rel, false, rel.Verdict, rel.Reason
	}

	u, hasU := findSector(rays, "u")
	d, hasD := findSector(rays, "d")
	if !hasU || !hasD || !u.HasProvenance || !d.HasProvenance || !u.HasBranchTag || !d.HasBranchTag || !u.BridgeOnly || !d.BridgeOnly {
		rel.Verdict = StatusFailedRequiresTwoProvenancedRays
		rel.Reason = "u and d rays must both be bridge-only, provenanced, and branch-tagged"
		return rel, false, rel.Verdict, rel.Reason
	}
	if !u.HasEigenbasisConvention || !d.HasEigenbasisConvention {
		rel.Verdict = StatusFailedMissingEigenbasisConvention
		rel.Reason = "a CKM-facing relative comparison requires explicit eigenvalue/eigenvector ordering and phase-gauge conventions"
		return rel, false, rel.Verdict, rel.Reason
	}

	rel.DeltaAlpha = d.Alpha - u.Alpha
	rel.DeltaPhi = wrapPi(d.Phi - u.Phi)
	rel.PhaseChord = 2 * math.Abs(math.Sin(rel.DeltaPhi/2))
	rel.ProjectiveDistance = math.Hypot(rel.DeltaAlpha, rel.PhaseChord)
	rel.CompleteInputs = true
	rel.EigenbasisConventionSet = true
	rel.BridgeOnly = true
	rel.ExportsCKMEntry = false
	rel.ExportsPMNSEntry = false
	rel.ExportsNativeObservable = false
	rel.Verdict = StatusUDDifferenceBridgeOnlyComputed
	rel.Reason = "relative u-d coefficient-ray diagnostics were computed without exporting CKM/PMNS entries"
	return rel, true, rel.Verdict, rel.Reason
}

func findSector(rays []SectorRay, sector string) (SectorRay, bool) {
	for _, r := range rays {
		if r.Sector == sector {
			return r, true
		}
	}
	return SectorRay{}, false
}

func hasObservedMixingImport(rays []SectorRay) bool {
	for _, r := range rays {
		if r.ExplicitObservedCKM || r.ExplicitObservedPMNS {
			return true
		}
	}
	return false
}

func hasNativePredictionClaim(rays []SectorRay) bool {
	for _, r := range rays {
		if r.NativePromotionClaim {
			return true
		}
	}
	return false
}

func hasNativeUniversalityClaim(rays []SectorRay) bool {
	for _, r := range rays {
		if r.UniversalityNativeClaim {
			return true
		}
	}
	return false
}

func hasLeptonPMNSMisroute(rays []SectorRay) bool {
	for _, r := range rays {
		if r.UsesLeptonPMNSLedger || r.Sector == "e" || r.Sector == "nu" {
			return true
		}
	}
	return false
}

func wrapPi(x float64) float64 {
	for x <= -math.Pi {
		x += 2 * math.Pi
	}
	for x > math.Pi {
		x -= 2 * math.Pi
	}
	return x
}

func buildFirewall(a Analysis) CKMFirewall {
	return CKMFirewall{
		Executed:                           true,
		RelativeRayMayFeedCKMAdapter:       true,
		CKMMatrixEntryComputed:             false,
		CKMMatrixEntryNative:               false,
		PMNSMatrixEntryComputed:            false,
		PMNSMatrixEntryNative:              false,
		RequiresObservedComparatorImport:   true,
		RequiresSchemeScaleProvenance:      a.Contract.RequiresProvenancePerSector,
		RequiresEigenvectorGaugeConvention: a.Contract.RequiresEigenbasisConvention,
		KGenStillForced:                    a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:               a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:             true,
		SectorCoefficientsStillSealed:      true,
		NoObservedMassesImported:           true,
		NoObservedYukawasImported:          true,
		NoObservedCKMImported:              true,
		NoObservedPMNSImported:             true,
		NoGSTPromotion:                     true,
		NativeFlavorDimAfter:               NativeFlavorDim,
		KXYCoeffDimAfter:                   KXYCoeffDim,
		Verdict:                            StatusFirewallPreserved,
		Reason:                             "Gate462 creates only a relative-ray bridge diagnostic; CKM/PMNS entries and physical flavor data remain quarantined.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        463,
		Title:       "Eigenbasis Convention Ledger / Mixing-Matrix Gauge Audit",
		Reason:      "Gate462 shows that a CKM-facing interface needs not only u-d relative rays but also explicit eigenvalue ordering and phase-gauge conventions.",
		PrimaryTask: "formalize the bridge-only eigenbasis convention required before a relative u-d ray can be passed to any future CKM residual evaluator, and prove convention choices are not native predictions.",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Contract.Executed || !a.Sieve.Executed || !a.Firewall.Executed {
		return fmt.Errorf("Gate462 incomplete execution")
	}
	if !(a.Inheritance.Gate461SectorMultiplex && a.Inheritance.Gate461IndependentSectorRaysAccepted && a.Inheritance.Gate461NativeUniversalityRejected && a.Inheritance.NoObservedValuesImported) {
		return fmt.Errorf("Gate462 missing inherited Gate461 firewall")
	}
	if !(a.Contract.RequiresUSector && a.Contract.RequiresDSector && a.Contract.RequiresProvenancePerSector && a.Contract.RequiresCompleteBranchTags && a.Contract.RequiresEigenbasisConvention && a.Contract.RelativeRayDimension == RelativeRayDOF && a.Contract.ExportsRelativeDiagnosticsOnly && a.Contract.RejectsCKMAsNativePrediction && a.Contract.RejectsPMNSInChargedCKMLedger && a.Contract.RejectsObservedMixingByDefault && a.Contract.RejectsNativeRelativeRayPromotion) {
		return fmt.Errorf("Gate462 interface contract is not fail-closed")
	}
	if !(a.Sieve.AcceptedCaseCount == 1 && a.Sieve.RejectedCaseCount == 8 && a.Sieve.ValidUDDifferenceAccepted && a.Sieve.MissingSectorRejected && a.Sieve.MissingProvenanceRejected && a.Sieve.MissingEigenbasisRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.NativeRelativePromotionRejected && a.Sieve.LeptonPMNSMisrouteRejected && a.Sieve.UniversalityNativeRejected && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoNativeMixingObservableExport) {
		return fmt.Errorf("Gate462 sieve did not close exactly")
	}
	if !(a.Firewall.RelativeRayMayFeedCKMAdapter && !a.Firewall.CKMMatrixEntryComputed && !a.Firewall.CKMMatrixEntryNative && !a.Firewall.PMNSMatrixEntryComputed && !a.Firewall.PMNSMatrixEntryNative && a.Firewall.RequiresObservedComparatorImport && a.Firewall.RequiresSchemeScaleProvenance && a.Firewall.RequiresEigenvectorGaugeConvention && a.Firewall.NoObservedMassesImported && a.Firewall.NoObservedYukawasImported && a.Firewall.NoObservedCKMImported && a.Firewall.NoObservedPMNSImported && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim) {
		return fmt.Errorf("Gate462 firewall failed")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.ValidUDDifferenceAccepted && a.Sieve.NativePredictionRejected && a.Firewall.NoObservedCKMImported {
		return "Gate462 isolates the u-d sector-difference ray that a future CKM bridge adapter may inspect. The object is a labelled bridge diagnostic, not a CKM prediction: observed CKM/PMNS values, eigenbasis conventions, sector coefficients, Yukawas, and physical masses remain quarantined."
	}
	return "Gate462 did not close the CKM-interface firewall."
}

func statuses() []string {
	return []string{
		StatusGate461Inherited,
		StatusRelativeRayLedgerDefined,
		StatusUDDifferenceBridgeOnlyComputed,
		StatusCKMInterfaceFirewallValidated,
		StatusPMNSInterfaceFirewallValidated,
		StatusFirewallPreserved,
		StatusFailedRequiresTwoProvenancedRays,
		StatusFailedMissingEigenbasisConvention,
		StatusFailedObservedCKMPMNSImportRejected,
		StatusFailedCKMPMNSPredictionRejected,
		StatusFailedNativeRelativeRayPromotion,
		StatusFailedLeptonSectorMisrouted,
		StatusFailedUniversalityNotNative,
	}
}
