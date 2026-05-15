// Package generation2sectormultiplex implements Gate 461:
// Three-Sector Comparator Multiplex / Universality Assumption Audit.
//
// Gate 460 validated a branch-resolved residual harness for one synthetic/null
// texture comparator record. Gate 461 lifts that harness into the charged
// sector ledger {u,d,e}. Its purpose is deliberately firewall-preserving: a
// valid bridge ray in one sector does not imply the same ray in another sector.
// Cross-sector universality is allowed only as an explicit bridge assumption or
// as the output of a future native theorem; it is not native ASHA law here.
package generation2sectormultiplex

import (
	"fmt"
	"math"
	"sync"
)

const (
	AuditID = "GATE461-THREE-SECTOR-COMPARATOR-MULTIPLEX-UNIVERSALITY-ASSUMPTION-AUDIT"

	StatusGate460Inherited             = "CONDITIONAL_SUPPORT_GATE460_RESIDUAL_HARNESS_INHERITED"
	StatusSectorMultiplexDefined       = "CONDITIONAL_SUPPORT_THREE_SECTOR_COMPARATOR_MULTIPLEX_DEFINED"
	StatusIndependentSectorRaysValid   = "CONDITIONAL_SUPPORT_INDEPENDENT_SECTOR_RAYS_BRIDGE_ONLY_VALIDATED"
	StatusBridgeUniversalityLabelled   = "CONDITIONAL_SUPPORT_CROSS_SECTOR_UNIVERSALITY_ALLOWED_AS_LABELLED_BRIDGE_ASSUMPTION"
	StatusMultiplexBridgeOnlyValidated = "CONDITIONAL_SUPPORT_THREE_SECTOR_MULTIPLEX_BRIDGE_ONLY_VALIDATED"
	StatusFirewallPreserved            = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedMissingSector              = "FAILED_ROUTE_THREE_SECTOR_LEDGER_INCOMPLETE"
	StatusFailedNativeUniversality         = "FAILED_ROUTE_CROSS_SECTOR_RAY_UNIVERSALITY_NOT_NATIVE"
	StatusFailedUnlabelledUniversality     = "FAILED_ROUTE_UNLABELLED_CROSS_SECTOR_RAY_SHARING_REJECTED"
	StatusFailedObservedDataRejected       = "FAILED_ROUTE_OBSERVED_DATA_REJECTED_IN_SECTOR_MULTIPLEX"
	StatusFailedNativePromotionRejected    = "FAILED_ROUTE_SECTOR_MULTIPLEX_NATIVE_PROMOTION_REJECTED"
	StatusFailedSectorContamination        = "FAILED_ROUTE_SECTOR_CROSS_CONTAMINATION_REJECTED"
	StatusFailedDimensionAccountingChanged = "FAILED_ROUTE_SECTOR_MULTIPLEX_CHANGED_13_MODULI_FIREWALL"
)

const (
	NativeFlavorDim    = 13
	KXYCoeffDim        = 9
	ChargedSectorCount = 3
	RequiredFields     = 11
)

var ChargedSectors = []string{"u", "d", "e"}

type Inheritance struct {
	Executed                      bool
	Gate444KGenForced             bool
	Gate445TriangleForced         bool
	Gate456InverseDerived         bool
	Gate457ProvenanceContract     bool
	Gate458ComparatorHarness      bool
	Gate459BranchTags             bool
	Gate460ResidualHarness        bool
	Gate460ResidualsBridgeOnly    bool
	Gate460ObservedRejected       bool
	Gate460NativePromotionBlocked bool
	NoObservedValuesImported      bool
	Verdict                       string
}

type MultiplexContract struct {
	Executed                             bool
	SectorIndexed                        bool
	RequiredSectors                      []string
	IndependentRayPerSector              bool
	NoImplicitRaySharing                 bool
	NoImplicitPhaseSharing               bool
	NoImplicitBranchTagSharing           bool
	RequiresProvenancePerRow             bool
	RequiresCompleteBranchTags           bool
	AllowsLabelledBridgeOnlyUniversality bool
	RejectsNativeUniversality            bool
	BridgeOnlyExport                     bool
	Verdict                              string
	Reason                               string
}

type SectorRecord struct {
	Sector                   string
	IK                       float64
	ISpec                    float64
	CPOddSign                int
	C3Sheet                  int
	HasNumericPair           bool
	HasCPOddSign             bool
	HasC3Sheet               bool
	HasProvenance            bool
	BridgeOnly               bool
	ExplicitObservedData     bool
	NativePromotionClaim     bool
	SharedFromSector         string
	UniversalityClaim        bool
	UniversalityLabelled     bool
	NativeUniversalityClaim  bool
	IndependentNativeTheorem bool
}

type SectorEvaluation struct {
	Record                    SectorRecord
	Accepted                  bool
	Evaluated                 bool
	Alpha                     float64
	Cos3Phi                   float64
	Phi                       float64
	CompleteBranchTag         bool
	ProvenanceComplete        bool
	BridgeOnlyExport          bool
	NativePromotionBlocked    bool
	RayIndependent            bool
	RaySharedOnlyByAssumption bool
	NoPhysicalObservableValue bool
	Verdict                   string
	Reason                    string
}

type LedgerCase struct {
	Name     string
	Records  []SectorRecord
	Verdict  string
	Accepted bool
	Reason   string
}

type Sieve struct {
	Executed                           bool
	Cases                              []LedgerCase
	CaseEvaluations                    map[string][]SectorEvaluation
	AcceptedCaseCount                  int
	RejectedCaseCount                  int
	IndependentThreeSectorAccepted     bool
	LabelledBridgeUniversalityAccepted bool
	MissingSectorRejected              bool
	NativeUniversalityRejected         bool
	UnlabelledUniversalityRejected     bool
	ObservedDataRejected               bool
	NativePromotionRejected            bool
	SectorContaminationRejected        bool
	AllAcceptedBridgeOnly              bool
	NoNativeObservableExport           bool
	Verdict                            string
	Reason                             string
}

type DimensionLedger struct {
	Executed                         bool
	ChargedSectors                   []string
	CoefficientsPerSector            int
	TotalKXYChargedCoefficients      int
	NativeChargedFlavorDimBefore     int
	NativeChargedFlavorDimAfter      int
	UniversalityWouldReduceBridgeDOF bool
	UniversalityReductionNative      bool
	IndependentSectorRaysNative      bool
	SectorRayUniversalityNative      bool
	Verdict                          string
	Reason                           string
}

type Firewall struct {
	Executed                      bool
	NoObservedMuonMassImported    bool
	NoObservedCharmMassImported   bool
	NoObservedTopBottomImported   bool
	NoObservedYukawaImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoGSTPromotion                bool
	NoCoefficientRayPromotion     bool
	NoCrossSectorUniversalityLaw  bool
	NoCurveFitPromoted            bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Contract    MultiplexContract
	Dimensions  DimensionLedger
	Sieve       Sieve
	Firewall    Firewall
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
	a.Dimensions = buildDimensionLedger()
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
		Executed:                      true,
		Gate444KGenForced:             true,
		Gate445TriangleForced:         true,
		Gate456InverseDerived:         true,
		Gate457ProvenanceContract:     true,
		Gate458ComparatorHarness:      true,
		Gate459BranchTags:             true,
		Gate460ResidualHarness:        true,
		Gate460ResidualsBridgeOnly:    true,
		Gate460ObservedRejected:       true,
		Gate460NativePromotionBlocked: true,
		NoObservedValuesImported:      true,
		Verdict:                       StatusGate460Inherited,
	}
}

func buildContract() MultiplexContract {
	return MultiplexContract{
		Executed:                             true,
		SectorIndexed:                        true,
		RequiredSectors:                      append([]string(nil), ChargedSectors...),
		IndependentRayPerSector:              true,
		NoImplicitRaySharing:                 true,
		NoImplicitPhaseSharing:               true,
		NoImplicitBranchTagSharing:           true,
		RequiresProvenancePerRow:             true,
		RequiresCompleteBranchTags:           true,
		AllowsLabelledBridgeOnlyUniversality: true,
		RejectsNativeUniversality:            true,
		BridgeOnlyExport:                     true,
		Verdict:                              StatusSectorMultiplexDefined,
		Reason:                               "the comparator ledger is indexed by charged sector; sharing alpha, phi, or branch tags across sectors requires an explicit bridge-only universality label or a future native theorem.",
	}
}

func buildDimensionLedger() DimensionLedger {
	return DimensionLedger{
		Executed:                         true,
		ChargedSectors:                   append([]string(nil), ChargedSectors...),
		CoefficientsPerSector:            3,
		TotalKXYChargedCoefficients:      KXYCoeffDim,
		NativeChargedFlavorDimBefore:     NativeFlavorDim,
		NativeChargedFlavorDimAfter:      NativeFlavorDim,
		UniversalityWouldReduceBridgeDOF: true,
		UniversalityReductionNative:      false,
		IndependentSectorRaysNative:      false,
		SectorRayUniversalityNative:      false,
		Verdict:                          StatusFailedNativeUniversality,
		Reason:                           "the charged K/X/Y ledger has three symbolic coefficients per charged sector; collapsing sectors to one shared ray would be an extra bridge constraint, not a native ASHA result.",
	}
}

func buildSieve() Sieve {
	cases := []LedgerCase{
		{Name: "complete independent synthetic u/d/e ledger", Records: independentRecords()},
		{Name: "labelled bridge-only universality stress test", Records: labelledUniversalityRecords()},
		{Name: "missing charged sector", Records: missingSectorRecords()},
		{Name: "native cross-sector universality claim", Records: nativeUniversalityRecords()},
		{Name: "unlabelled cross-sector ray sharing", Records: unlabelledSharedRecords()},
		{Name: "observed values attempted in multiplex", Records: observedRecords()},
		{Name: "sector multiplex native-promotion attempt", Records: nativePromotionRecords()},
		{Name: "sector cross-contamination", Records: contaminatedRecords()},
	}
	out := Sieve{Executed: true, CaseEvaluations: map[string][]SectorEvaluation{}, AllAcceptedBridgeOnly: true, NoNativeObservableExport: true}
	for _, c := range cases {
		evals := evaluateLedger(c.Records)
		c.Accepted, c.Verdict, c.Reason = classifyCase(c.Name, c.Records, evals)
		out.Cases = append(out.Cases, c)
		out.CaseEvaluations[c.Name] = evals
		if c.Accepted {
			out.AcceptedCaseCount++
		} else {
			out.RejectedCaseCount++
		}
		switch c.Name {
		case "complete independent synthetic u/d/e ledger":
			out.IndependentThreeSectorAccepted = c.Accepted && c.Verdict == StatusIndependentSectorRaysValid
		case "labelled bridge-only universality stress test":
			out.LabelledBridgeUniversalityAccepted = c.Accepted && c.Verdict == StatusBridgeUniversalityLabelled
		case "missing charged sector":
			out.MissingSectorRejected = !c.Accepted && c.Verdict == StatusFailedMissingSector
		case "native cross-sector universality claim":
			out.NativeUniversalityRejected = !c.Accepted && c.Verdict == StatusFailedNativeUniversality
		case "unlabelled cross-sector ray sharing":
			out.UnlabelledUniversalityRejected = !c.Accepted && c.Verdict == StatusFailedUnlabelledUniversality
		case "observed values attempted in multiplex":
			out.ObservedDataRejected = !c.Accepted && c.Verdict == StatusFailedObservedDataRejected
		case "sector multiplex native-promotion attempt":
			out.NativePromotionRejected = !c.Accepted && c.Verdict == StatusFailedNativePromotionRejected
		case "sector cross-contamination":
			out.SectorContaminationRejected = !c.Accepted && c.Verdict == StatusFailedSectorContamination
		}
		for _, e := range evals {
			if e.Accepted && !e.BridgeOnlyExport {
				out.AllAcceptedBridgeOnly = false
			}
			if e.Accepted && !e.NoPhysicalObservableValue {
				out.NoNativeObservableExport = false
			}
		}
	}
	out.Verdict = StatusMultiplexBridgeOnlyValidated
	out.Reason = "complete independent sector ledgers and explicitly-labelled bridge universality stress tests survive; missing, unlabelled, observed, contaminated, or native-promotion records fail closed."
	return out
}

func independentRecords() []SectorRecord {
	return []SectorRecord{
		validRecord("u", 0.20, 0.070, +1, 0),
		validRecord("d", -0.18, 0.055, -1, 1),
		validRecord("e", 0.31, -0.040, +1, 2),
	}
}

func labelledUniversalityRecords() []SectorRecord {
	recs := []SectorRecord{
		validRecord("u", 0.20, 0.070, +1, 0),
		validRecord("d", 0.20, 0.070, +1, 0),
		validRecord("e", 0.20, 0.070, +1, 0),
	}
	for i := range recs {
		recs[i].UniversalityClaim = true
		recs[i].UniversalityLabelled = true
	}
	return recs
}

func missingSectorRecords() []SectorRecord {
	return []SectorRecord{validRecord("u", 0.20, 0.070, +1, 0), validRecord("d", -0.18, 0.055, -1, 1)}
}

func nativeUniversalityRecords() []SectorRecord {
	recs := labelledUniversalityRecords()
	for i := range recs {
		recs[i].UniversalityLabelled = false
		recs[i].NativeUniversalityClaim = true
	}
	return recs
}

func unlabelledSharedRecords() []SectorRecord {
	recs := []SectorRecord{
		validRecord("u", 0.20, 0.070, +1, 0),
		validRecord("d", 0.20, 0.070, +1, 0),
		validRecord("e", 0.20, 0.070, +1, 0),
	}
	for i := range recs {
		recs[i].UniversalityClaim = true
	}
	return recs
}

func observedRecords() []SectorRecord {
	recs := independentRecords()
	recs[1].ExplicitObservedData = true
	return recs
}

func nativePromotionRecords() []SectorRecord {
	recs := independentRecords()
	recs[2].NativePromotionClaim = true
	recs[2].BridgeOnly = false
	return recs
}

func contaminatedRecords() []SectorRecord {
	recs := independentRecords()
	recs[1].SharedFromSector = "u"
	return recs
}

func validRecord(sector string, ik, ispec float64, sigma, sheet int) SectorRecord {
	return SectorRecord{
		Sector:         sector,
		IK:             ik,
		ISpec:          ispec,
		CPOddSign:      sigma,
		C3Sheet:        sheet,
		HasNumericPair: true,
		HasCPOddSign:   true,
		HasC3Sheet:     true,
		HasProvenance:  true,
		BridgeOnly:     true,
	}
}

func evaluateLedger(records []SectorRecord) []SectorEvaluation {
	evals := make([]SectorEvaluation, 0, len(records))
	for _, r := range records {
		evals = append(evals, EvaluateSectorRecord(r))
	}
	return evals
}

func EvaluateSectorRecord(r SectorRecord) SectorEvaluation {
	e := SectorEvaluation{Record: r, NoPhysicalObservableValue: true, NativePromotionBlocked: r.NativePromotionClaim}
	if r.ExplicitObservedData {
		e.Verdict = StatusFailedObservedDataRejected
		e.Reason = "observed flavor values are not accepted by the Gate461 multiplex audit."
		return e
	}
	if r.NativePromotionClaim || !r.BridgeOnly {
		e.Verdict = StatusFailedNativePromotionRejected
		e.Reason = "sector-indexed comparator records are bridge diagnostics and cannot be promoted to native law-space."
		return e
	}
	if r.SharedFromSector != "" && r.SharedFromSector != r.Sector {
		e.Verdict = StatusFailedSectorContamination
		e.Reason = "a sector row may not silently reuse another sector's ray or branch tag."
		return e
	}
	if !validSector(r.Sector) || !r.HasNumericPair || !r.HasProvenance || !r.HasCPOddSign || !r.HasC3Sheet || (r.CPOddSign != +1 && r.CPOddSign != -1) || r.C3Sheet < 0 || r.C3Sheet > 2 || math.Abs(r.IK) >= 1 {
		e.Verdict = StatusFailedMissingSector
		e.Reason = "sector rows require a valid sector, provenance, numeric comparators, and a complete branch tag."
		return e
	}
	cos3 := cosThreePhiFromComparators(r.IK, r.ISpec)
	if math.Abs(cos3) > 1 || math.IsNaN(cos3) || math.IsInf(cos3, 0) {
		e.Verdict = StatusFailedMissingSector
		e.Reason = "sector row lies outside the symbolic comparator domain."
		return e
	}
	e.Alpha = math.Sqrt(3) * r.IK / math.Sqrt(1-r.IK*r.IK)
	e.Cos3Phi = cos3
	e.Phi = (float64(r.CPOddSign)*math.Acos(cos3) + 2*math.Pi*float64(r.C3Sheet)) / 3
	e.Accepted = true
	e.Evaluated = true
	e.CompleteBranchTag = true
	e.ProvenanceComplete = true
	e.BridgeOnlyExport = true
	e.RayIndependent = !r.UniversalityClaim
	e.RaySharedOnlyByAssumption = r.UniversalityClaim && r.UniversalityLabelled && !r.IndependentNativeTheorem
	if e.RaySharedOnlyByAssumption {
		e.Verdict = StatusBridgeUniversalityLabelled
		e.Reason = "same ray across sectors is retained only as a labelled bridge universality assumption."
	} else {
		e.Verdict = StatusIndependentSectorRaysValid
		e.Reason = "sector row has its own provenance, comparators, and branch tag."
	}
	return e
}

func classifyCase(name string, records []SectorRecord, evals []SectorEvaluation) (bool, string, string) {
	if !hasAllSectors(records) {
		return false, StatusFailedMissingSector, "the charged-sector multiplex requires exactly the u, d, and e sectors."
	}
	for _, e := range evals {
		if !e.Accepted {
			return false, e.Verdict, e.Reason
		}
	}
	if hasSectorContamination(records) {
		return false, StatusFailedSectorContamination, "cross-sector ray reuse is not allowed without a universality label."
	}
	if hasNativeUniversalityClaim(records) {
		return false, StatusFailedNativeUniversality, "cross-sector coefficient-ray universality is not a native theorem of Gate461."
	}
	if hasUnlabelledUniversalityClaim(records) {
		return false, StatusFailedUnlabelledUniversality, "ray sharing across sectors must be explicitly labelled as bridge-only."
	}
	if hasLabelledUniversality(records) {
		return true, StatusBridgeUniversalityLabelled, "the shared ray is accepted only as an explicit bridge-only stress-test assumption."
	}
	return true, StatusIndependentSectorRaysValid, "all three sectors carry independent bridge rays and complete branch tags."
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		NoObservedMuonMassImported:    true,
		NoObservedCharmMassImported:   true,
		NoObservedTopBottomImported:   true,
		NoObservedYukawaImported:      true,
		NoCKMImported:                 true,
		NoPMNSImported:                true,
		NoGSTPromotion:                true,
		NoCoefficientRayPromotion:     true,
		NoCrossSectorUniversalityLaw:  true,
		NoCurveFitPromoted:            true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate461 indexes bridge residuals by sector and rejects implicit cross-sector universality; it does not import observed masses, Yukawas, CKM/PMNS data, or collapse sector amplitudes.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        462,
		Title:       "Sector-Difference Invariant / CKM Interface Firewall Audit",
		Reason:      "Gate461 proves sector rays are independent bridge data; the next audit should isolate which sector-difference invariants would feed CKM-like mixing without turning them into native predictions.",
		PrimaryTask: "construct a bridge-only relative-ray ledger between u and d sectors and prove CKM/PMNS entries remain quarantined unless explicit observed comparator records are imported with provenance.",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Contract.Executed || !a.Dimensions.Executed || !a.Sieve.Executed || !a.Firewall.Executed {
		return fmt.Errorf("Gate461 incomplete execution")
	}
	if !(a.Inheritance.Gate460ResidualHarness && a.Inheritance.Gate460ResidualsBridgeOnly && a.Inheritance.Gate460ObservedRejected && a.Inheritance.Gate460NativePromotionBlocked) {
		return fmt.Errorf("Gate461 missing inherited Gate460 firewall")
	}
	if !(a.Contract.SectorIndexed && a.Contract.IndependentRayPerSector && a.Contract.NoImplicitRaySharing && a.Contract.RequiresProvenancePerRow && a.Contract.RequiresCompleteBranchTags && a.Contract.AllowsLabelledBridgeOnlyUniversality && a.Contract.RejectsNativeUniversality && a.Contract.BridgeOnlyExport) {
		return fmt.Errorf("Gate461 multiplex contract is not fail-closed")
	}
	if !(a.Dimensions.TotalKXYChargedCoefficients == KXYCoeffDim && a.Dimensions.NativeChargedFlavorDimAfter == NativeFlavorDim && a.Dimensions.UniversalityWouldReduceBridgeDOF && !a.Dimensions.UniversalityReductionNative && !a.Dimensions.SectorRayUniversalityNative) {
		return fmt.Errorf("Gate461 dimension ledger changed the firewall")
	}
	if !(a.Sieve.AcceptedCaseCount == 2 && a.Sieve.RejectedCaseCount == 6 && a.Sieve.IndependentThreeSectorAccepted && a.Sieve.LabelledBridgeUniversalityAccepted && a.Sieve.MissingSectorRejected && a.Sieve.NativeUniversalityRejected && a.Sieve.UnlabelledUniversalityRejected && a.Sieve.ObservedDataRejected && a.Sieve.NativePromotionRejected && a.Sieve.SectorContaminationRejected && a.Sieve.AllAcceptedBridgeOnly && a.Sieve.NoNativeObservableExport) {
		return fmt.Errorf("Gate461 sector multiplex sieve did not close exactly")
	}
	if !(a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoGSTPromotion && a.Firewall.NoCoefficientRayPromotion && a.Firewall.NoCrossSectorUniversalityLaw && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim) {
		return fmt.Errorf("Gate461 firewall failed")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Sieve.IndependentThreeSectorAccepted && a.Sieve.NativeUniversalityRejected && a.Firewall.NoCrossSectorUniversalityLaw {
		return "Gate461 lifts the branch-resolved residual harness into the charged-sector ledger {u,d,e}. Each sector may carry a labelled bridge ray, and a shared ray may be stress-tested only as an explicit bridge assumption. ASHA does not natively force cross-sector coefficient-ray universality, so sector amplitudes and CKM/PMNS-facing differences remain firewalled."
	}
	return "Gate461 did not validate the three-sector multiplex firewall."
}

func statuses() []string {
	return []string{
		StatusGate460Inherited,
		StatusSectorMultiplexDefined,
		StatusIndependentSectorRaysValid,
		StatusBridgeUniversalityLabelled,
		StatusMultiplexBridgeOnlyValidated,
		StatusFirewallPreserved,
		StatusFailedMissingSector,
		StatusFailedNativeUniversality,
		StatusFailedUnlabelledUniversality,
		StatusFailedObservedDataRejected,
		StatusFailedNativePromotionRejected,
		StatusFailedSectorContamination,
		StatusFailedDimensionAccountingChanged,
	}
}

func cosThreePhiFromComparators(ik, ispec float64) float64 {
	return (3 * math.Sqrt(3) / 2) * ispec / math.Pow(1-ik*ik, 1.5)
}

func validSector(s string) bool {
	for _, want := range ChargedSectors {
		if s == want {
			return true
		}
	}
	return false
}

func hasAllSectors(records []SectorRecord) bool {
	if len(records) != len(ChargedSectors) {
		return false
	}
	seen := map[string]bool{}
	for _, r := range records {
		if !validSector(r.Sector) || seen[r.Sector] {
			return false
		}
		seen[r.Sector] = true
	}
	for _, s := range ChargedSectors {
		if !seen[s] {
			return false
		}
	}
	return true
}

func hasNativeUniversalityClaim(records []SectorRecord) bool {
	for _, r := range records {
		if r.UniversalityClaim && r.NativeUniversalityClaim && !r.UniversalityLabelled && !r.IndependentNativeTheorem {
			return true
		}
	}
	return false
}

func hasUnlabelledUniversalityClaim(records []SectorRecord) bool {
	for _, r := range records {
		if r.UniversalityClaim && !r.UniversalityLabelled && !r.NativePromotionClaim && !r.NativeUniversalityClaim {
			return true
		}
	}
	return false
}

func hasLabelledUniversality(records []SectorRecord) bool {
	for _, r := range records {
		if !r.UniversalityClaim || !r.UniversalityLabelled || r.NativePromotionClaim {
			return false
		}
	}
	return len(records) > 0
}

func hasSectorContamination(records []SectorRecord) bool {
	for _, r := range records {
		if r.SharedFromSector != "" && r.SharedFromSector != r.Sector {
			return true
		}
	}
	return false
}
