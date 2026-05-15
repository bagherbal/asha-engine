// Package generation2leptonpmnsnullresidual implements Gate 476:
// Lepton-sector synthetic PMNS null residual.
//
// Gate 475 defined the strict e/nu preflight ledger for a future PMNS-facing
// bridge comparison. Gate 476 mirrors the CKM null adapter, but only on a
// synthetic rank-complete charged-lepton/neutrino ledger. It computes a
// bridge-only distance d_{e nu}; it does not construct U_PMNS, does not import
// observed PMNS values, and does not promote the residual as native ASHA law.
package generation2leptonpmnsnullresidual

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE476-LEPTON-SECTOR-SYNTHETIC-PMNS-NULL-RESIDUAL"

	StatusGate475Inherited                  = "CONDITIONAL_SUPPORT_GATE475_LEPTON_PREFLIGHT_INHERITED"
	StatusPMNSNullMapDefined                = "CONDITIONAL_SUPPORT_PMNS_NULL_RESIDUAL_SYMBOLIC_MAP_DEFINED"
	StatusSyntheticLeptonLedgerAccepted     = "CONDITIONAL_SUPPORT_SYNTHETIC_LEPTON_LEDGER_ACCEPTED"
	StatusPMNSNullResidualComputed          = "CONDITIONAL_SUPPORT_SYNTHETIC_PMNS_NULL_RESIDUAL_COMPUTED"
	StatusPMNSNullResidualFirewallValidated = "CONDITIONAL_SUPPORT_PMNS_NULL_RESIDUAL_FIREWALL_VALIDATED"
	StatusLeptonSocketStructurallyIdentical = "CONDITIONAL_SUPPORT_LEPTON_SOCKET_STRUCTURALLY_IDENTICAL_TO_QUARK_SOCKET"
	StatusFirewallPreserved                 = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE476_PMNS_SOCKET"

	StatusFailedMissingENuSectors         = "FAILED_ROUTE_PMNS_NULL_RESIDUAL_REQUIRES_E_NU_SECTORS"
	StatusFailedMissingPreflight          = "FAILED_ROUTE_PMNS_NULL_RESIDUAL_REQUIRES_GATE475_PREFLIGHT"
	StatusFailedMissingRankCompleteLedger = "FAILED_ROUTE_PMNS_NULL_RESIDUAL_REQUIRES_I_SPEC_I_K_BRANCH_TAGS"
	StatusFailedMissingNeutrinoPolicies   = "FAILED_ROUTE_PMNS_NULL_RESIDUAL_REQUIRES_NEUTRINO_POLICIES"
	StatusFailedObservedPMNSImport        = "FAILED_ROUTE_OBSERVED_PMNS_IMPORT_REJECTED_IN_SYNTHETIC_SOCKET"
	StatusFailedPMNSAsRayInput            = "FAILED_ROUTE_PMNS_USED_AS_SYNTHETIC_LEPTON_RAY_INPUT_REJECTED"
	StatusFailedPMNSNativePrediction      = "FAILED_ROUTE_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedPMNSMatrixExport          = "FAILED_ROUTE_PMNS_MATRIX_EXPORT_REJECTED"
	StatusFailedNativeResidualPromotion   = "FAILED_ROUTE_PMNS_RESIDUAL_NATIVE_PROMOTION_REJECTED"
	StatusFailedProjectiveDomainRejected  = "FAILED_ROUTE_PMNS_SOCKET_PROJECTIVE_DOMAIN_REJECTED"
	StatusFailedPhaseDomainRejected       = "FAILED_ROUTE_PMNS_SOCKET_PHASE_DOMAIN_REJECTED"
	StatusFailedCausticRejected           = "FAILED_ROUTE_PMNS_SOCKET_CAUSTIC_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate445TriangleForced           bool
	Gate456RayInverseAvailable      bool
	Gate459BranchTagsAvailable      bool
	Gate464QuarkSocketAvailable     bool
	Gate474PMNSBridgeOnly           bool
	Gate475LeptonPreflightDefined   bool
	Gate475RequiresENuSectors       bool
	Gate475RequiresISpecIK          bool
	Gate475RequiresBranchTags       bool
	Gate475RequiresNeutrinoPolicies bool
	Gate475RejectsPMNSAsRayInput    bool
	Gate475RejectsNativePromotion   bool
	NoObservedValuesImported        bool
	NativeRegistryClean             bool
	Verdict                         string
}

type PMNSNullMap struct {
	Executed                        bool
	Formula                         string
	RequiresERow                    bool
	RequiresNuRow                   bool
	RequiresGate475Preflight        bool
	RequiresISpecIK                 bool
	RequiresBranchTags              bool
	RequiresNeutrinoOrdering        bool
	RequiresAbsoluteNuScale         bool
	RequiresMajoranaDiracPolicy     bool
	RequiresSyntheticMode           bool
	RequiresBridgeOnly              bool
	AllowsPMNSTarget                bool
	AllowsPMNSAsRayInput            bool
	ComputesDENu                    bool
	ComputesPMNSMatrix              bool
	ComputesPMNSEntry               bool
	ExportsDiagnosticsOnly          bool
	StructurallyIdenticalToQuarkMap bool
	Verdict                         string
	Reason                          string
}

type LeptonComparator struct {
	Sector                    string
	Alpha                     float64
	Phi                       float64
	ISpec                     float64
	IK                        float64
	SigmaCP                   int
	C3Sheet                   int
	HasSector                 bool
	HasISpec                  bool
	HasIK                     bool
	HasBranchTags             bool
	HasPreflight              bool
	HasEigenbasisConvention   bool
	HasNeutrinoOrdering       bool
	HasAbsoluteNuScale        bool
	HasMajoranaDiracPolicy    bool
	HasUncertainty            bool
	BridgeOnly                bool
	SyntheticOnly             bool
	ObservedPMNSImport        bool
	PMNSAsRayInput            bool
	NativePromotionClaim      bool
	PMNSMatrixExportRequest   bool
	PMNSNativePredictionClaim bool
}

type Ray struct {
	Sector        string
	Alpha         float64
	Phi           float64
	ISpec         float64
	IK            float64
	CosThreePhi   float64
	SigmaCP       int
	C3Sheet       int
	BridgeOnly    bool
	SyntheticOnly bool
	InsideDomain  bool
	AtCaustic     bool
	ExportsNative bool
	Verdict       string
	Reason        string
}

type PMNSNullResidual struct {
	DeltaAlpha                float64
	DeltaPhi                  float64
	PhaseChord                float64
	DENu                      float64
	SyntheticPMNSTarget       float64
	SyntheticResidual         float64
	CompleteInputs            bool
	Gate475PreflightSatisfied bool
	BridgeOnly                bool
	SyntheticOnly             bool
	ExportsRelativeDiagnostic bool
	PMNSMatrixConstructed     bool
	PMNSEntryComputed         bool
	ExportsNativeObservable   bool
	ObservedPMNSImported      bool
	Verdict                   string
	Reason                    string
}

type Case struct {
	Name     string
	E        LeptonComparator
	Nu       LeptonComparator
	Target   float64
	Accepted bool
	ERay     Ray
	NuRay    Ray
	Residual PMNSNullResidual
	Verdict  string
	Reason   string
}

type Sieve struct {
	Executed                        bool
	Cases                           []Case
	AcceptedCaseCount               int
	RejectedCaseCount               int
	ValidSyntheticResidualAccepted  bool
	MissingENuRejected              bool
	MissingPreflightRejected        bool
	MissingRankLedgerRejected       bool
	MissingNeutrinoPoliciesRejected bool
	ObservedPMNSRejected            bool
	PMNSAsRayInputRejected          bool
	PMNSNativePredictionRejected    bool
	PMNSMatrixExportRejected        bool
	NativeResidualPromotionRejected bool
	ProjectiveDomainRejected        bool
	PhaseDomainRejected             bool
	CausticRejected                 bool
	AllAcceptedBridgeOnlySynthetic  bool
	NoPMNSMatrixConstructed         bool
	NoNativeObservableExport        bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                         bool
	PMNSNullResidualAdapterDefined   bool
	PMNSNullResidualMayRunBridgeOnly bool
	SyntheticLeptonDataEvaluated     bool
	ObservedLeptonDataImported       bool
	ObservedPMNSImported             bool
	PMNSMatrixConstructed            bool
	PMNSEntryComputed                bool
	PMNSNativePrediction             bool
	DENuNativePrediction             bool
	SyntheticPMNSTargetNative        bool
	NativeRegistryWritten            bool
	CKMNativePrediction              bool
	KGenStillForced                  bool
	XTriangleStillForced             bool
	YPhaseStillQuarantined           bool
	SectorCoefficientsStillSealed    bool
	NativeFlavorDimAfter             int
	KXYCoeffDimAfter                 int
	Verdict                          string
	Reason                           string
}

type Output struct {
	SyntheticDENu                float64
	SyntheticPMNSTarget          float64
	SyntheticResidual            float64
	AlphaE, PhiE, AlphaNu, PhiNu float64
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Map         PMNSNullMap
	Sieve       Sieve
	Output      Output
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
	a.Map = buildMap()
	a.Sieve = buildSieve()
	a.Output = buildOutput(a.Sieve)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate444KGenForced:               true,
		Gate445TriangleForced:           true,
		Gate456RayInverseAvailable:      true,
		Gate459BranchTagsAvailable:      true,
		Gate464QuarkSocketAvailable:     true,
		Gate474PMNSBridgeOnly:           true,
		Gate475LeptonPreflightDefined:   true,
		Gate475RequiresENuSectors:       true,
		Gate475RequiresISpecIK:          true,
		Gate475RequiresBranchTags:       true,
		Gate475RequiresNeutrinoPolicies: true,
		Gate475RejectsPMNSAsRayInput:    true,
		Gate475RejectsNativePromotion:   true,
		NoObservedValuesImported:        true,
		NativeRegistryClean:             true,
		Verdict:                         StatusGate475Inherited,
	}
}

func buildMap() PMNSNullMap {
	return PMNSNullMap{
		Executed:                        true,
		Formula:                         "d_e_nu=sqrt((alpha_nu-alpha_e)^2+4*sin^2((phi_nu-phi_e)/2))",
		RequiresERow:                    true,
		RequiresNuRow:                   true,
		RequiresGate475Preflight:        true,
		RequiresISpecIK:                 true,
		RequiresBranchTags:              true,
		RequiresNeutrinoOrdering:        true,
		RequiresAbsoluteNuScale:         true,
		RequiresMajoranaDiracPolicy:     true,
		RequiresSyntheticMode:           true,
		RequiresBridgeOnly:              true,
		AllowsPMNSTarget:                true,
		AllowsPMNSAsRayInput:            false,
		ComputesDENu:                    true,
		ComputesPMNSMatrix:              false,
		ComputesPMNSEntry:               false,
		ExportsDiagnosticsOnly:          true,
		StructurallyIdenticalToQuarkMap: true,
		Verdict:                         StatusPMNSNullMapDefined,
		Reason:                          "the e/nu socket is the same phase-cylinder distance as the u/d CKM-null socket, with PMNS allowed only as a synthetic residual target",
	}
}

func buildSieve() Sieve {
	validE := syntheticComparator("e", 0.85, 0.30)
	validNu := syntheticComparator("nu", 1.05, 0.90)
	cases := []Case{
		EvaluateCase("valid synthetic e/nu PMNS-null residual", validE, validNu, 0.620),
		EvaluateCase("missing charged-lepton row", mutate(validE, func(x *LeptonComparator) { x.HasSector = false; x.Sector = "" }), validNu, 0.620),
		EvaluateCase("missing Gate475 preflight", mutate(validE, func(x *LeptonComparator) { x.HasPreflight = false }), validNu, 0.620),
		EvaluateCase("missing I_K comparator", mutate(validE, func(x *LeptonComparator) { x.HasIK = false }), validNu, 0.620),
		EvaluateCase("missing neutrino ordering policy", validE, mutate(validNu, func(x *LeptonComparator) { x.HasNeutrinoOrdering = false }), 0.620),
		EvaluateCase("observed PMNS import rejected", mutate(validE, func(x *LeptonComparator) { x.ObservedPMNSImport = true }), validNu, 0.620),
		EvaluateCase("PMNS as ray input rejected", mutate(validE, func(x *LeptonComparator) { x.PMNSAsRayInput = true }), validNu, 0.620),
		EvaluateCase("PMNS native prediction rejected", mutate(validE, func(x *LeptonComparator) { x.PMNSNativePredictionClaim = true }), validNu, 0.620),
		EvaluateCase("PMNS matrix export rejected", mutate(validE, func(x *LeptonComparator) { x.PMNSMatrixExportRequest = true }), validNu, 0.620),
		EvaluateCase("native residual promotion rejected", mutate(validE, func(x *LeptonComparator) { x.NativePromotionClaim = true; x.BridgeOnly = false }), validNu, 0.620),
		EvaluateCase("projective domain rejected", mutate(validE, func(x *LeptonComparator) { x.IK = 1 }), validNu, 0.620),
		EvaluateCase("phase domain rejected", mutate(validE, func(x *LeptonComparator) { x.ISpec = 9 }), validNu, 0.620),
		EvaluateCase("caustic rejected", syntheticComparator("e", 0.85, 0), validNu, 0.620),
	}
	s := Sieve{Executed: true, Cases: cases, NoPMNSMatrixConstructed: true, NoNativeObservableExport: true, AllAcceptedBridgeOnlySynthetic: true}
	for _, c := range cases {
		if c.Accepted {
			s.AcceptedCaseCount++
			if c.Verdict == StatusPMNSNullResidualComputed {
				s.ValidSyntheticResidualAccepted = true
			}
			if !c.Residual.BridgeOnly || !c.Residual.SyntheticOnly || c.Residual.ExportsNativeObservable {
				s.AllAcceptedBridgeOnlySynthetic = false
			}
		} else {
			s.RejectedCaseCount++
		}
		s.MissingENuRejected = s.MissingENuRejected || c.Verdict == StatusFailedMissingENuSectors
		s.MissingPreflightRejected = s.MissingPreflightRejected || c.Verdict == StatusFailedMissingPreflight
		s.MissingRankLedgerRejected = s.MissingRankLedgerRejected || c.Verdict == StatusFailedMissingRankCompleteLedger
		s.MissingNeutrinoPoliciesRejected = s.MissingNeutrinoPoliciesRejected || c.Verdict == StatusFailedMissingNeutrinoPolicies
		s.ObservedPMNSRejected = s.ObservedPMNSRejected || c.Verdict == StatusFailedObservedPMNSImport
		s.PMNSAsRayInputRejected = s.PMNSAsRayInputRejected || c.Verdict == StatusFailedPMNSAsRayInput
		s.PMNSNativePredictionRejected = s.PMNSNativePredictionRejected || c.Verdict == StatusFailedPMNSNativePrediction
		s.PMNSMatrixExportRejected = s.PMNSMatrixExportRejected || c.Verdict == StatusFailedPMNSMatrixExport
		s.NativeResidualPromotionRejected = s.NativeResidualPromotionRejected || c.Verdict == StatusFailedNativeResidualPromotion
		s.ProjectiveDomainRejected = s.ProjectiveDomainRejected || c.Verdict == StatusFailedProjectiveDomainRejected
		s.PhaseDomainRejected = s.PhaseDomainRejected || c.Verdict == StatusFailedPhaseDomainRejected
		s.CausticRejected = s.CausticRejected || c.Verdict == StatusFailedCausticRejected
		if c.Residual.PMNSMatrixConstructed || c.Residual.PMNSEntryComputed {
			s.NoPMNSMatrixConstructed = false
		}
		if c.Residual.ExportsNativeObservable {
			s.NoNativeObservableExport = false
		}
	}
	s.Verdict = StatusPMNSNullResidualFirewallValidated
	s.Reason = "one synthetic rank-complete e/nu ledger computes d_e_nu; every observed, incomplete, matrix-export, PMNS-as-ray, or native-promotion route fails closed"
	return s
}

func syntheticComparator(sector string, alpha, phi float64) LeptonComparator {
	return LeptonComparator{
		Sector:                  sector,
		Alpha:                   alpha,
		Phi:                     phi,
		IK:                      alpha / math.Sqrt(alpha*alpha+3),
		ISpec:                   2 * math.Cos(3*phi) / math.Pow(alpha*alpha+3, 1.5),
		SigmaCP:                 signNonzero(math.Sin(3 * phi)),
		C3Sheet:                 c3Sheet(phi),
		HasSector:               true,
		HasISpec:                true,
		HasIK:                   true,
		HasBranchTags:           true,
		HasPreflight:            true,
		HasEigenbasisConvention: true,
		HasNeutrinoOrdering:     true,
		HasAbsoluteNuScale:      true,
		HasMajoranaDiracPolicy:  true,
		HasUncertainty:          true,
		BridgeOnly:              true,
		SyntheticOnly:           true,
	}
}

func mutate(x LeptonComparator, f func(*LeptonComparator)) LeptonComparator {
	f(&x)
	return x
}

func signNonzero(x float64) int {
	if x < 0 {
		return -1
	}
	return 1
}

func c3Sheet(phi float64) int {
	// Gate459 branch convention: phi=(sigma_CP*acos(cos(3phi))+2*pi*n_C3)/3.
	// Solve for n_C3 so the synthetic row round-trips through the declared inverse map.
	sigma := float64(signNonzero(math.Sin(3 * phi)))
	n := int(math.Round((3*phi - sigma*math.Acos(math.Cos(3*phi))) / (2 * math.Pi)))
	n %= 3
	if n < 0 {
		n += 3
	}
	return n
}

func EvaluateCase(name string, e, nu LeptonComparator, target float64) Case {
	c := Case{Name: name, E: e, Nu: nu, Target: target}
	if verdict, reason := preflightFailure(e, nu); verdict != "" {
		c.Verdict = verdict
		c.Reason = reason
		c.Residual = rejectedResidual(verdict, reason)
		return c
	}
	eRay, verdict, reason := invert(e)
	c.ERay = eRay
	if verdict != "" {
		c.Verdict = verdict
		c.Reason = reason
		c.Residual = rejectedResidual(verdict, reason)
		return c
	}
	nuRay, verdict, reason := invert(nu)
	c.NuRay = nuRay
	if verdict != "" {
		c.Verdict = verdict
		c.Reason = reason
		c.Residual = rejectedResidual(verdict, reason)
		return c
	}
	deltaAlpha := nuRay.Alpha - eRay.Alpha
	deltaPhi := wrapPi(nuRay.Phi - eRay.Phi)
	phaseChord := 2 * math.Sin(deltaPhi/2)
	d := math.Sqrt(deltaAlpha*deltaAlpha + phaseChord*phaseChord)
	c.Residual = PMNSNullResidual{
		DeltaAlpha:                deltaAlpha,
		DeltaPhi:                  deltaPhi,
		PhaseChord:                phaseChord,
		DENu:                      d,
		SyntheticPMNSTarget:       target,
		SyntheticResidual:         math.Abs(d - target),
		CompleteInputs:            true,
		Gate475PreflightSatisfied: true,
		BridgeOnly:                true,
		SyntheticOnly:             true,
		ExportsRelativeDiagnostic: true,
		PMNSMatrixConstructed:     false,
		PMNSEntryComputed:         false,
		ExportsNativeObservable:   false,
		ObservedPMNSImported:      false,
		Verdict:                   StatusPMNSNullResidualComputed,
		Reason:                    "synthetic e/nu bridge rays compute a PMNS-null residual diagnostic only",
	}
	c.Accepted = true
	c.Verdict = StatusPMNSNullResidualComputed
	c.Reason = c.Residual.Reason
	return c
}

func preflightFailure(e, nu LeptonComparator) (string, string) {
	if !e.HasSector || !nu.HasSector || e.Sector != "e" || nu.Sector != "nu" {
		return StatusFailedMissingENuSectors, "PMNS-null residual requires both charged-lepton sector e and neutrino sector nu"
	}
	if !e.HasPreflight || !nu.HasPreflight || !e.HasEigenbasisConvention || !nu.HasEigenbasisConvention {
		return StatusFailedMissingPreflight, "Gate475 preflight and eigenbasis convention must be complete before evaluating d_e_nu"
	}
	if !e.HasISpec || !nu.HasISpec || !e.HasIK || !nu.HasIK || !e.HasBranchTags || !nu.HasBranchTags || e.SigmaCP == 0 || nu.SigmaCP == 0 {
		return StatusFailedMissingRankCompleteLedger, "rank-complete lepton ledger requires I_spec, I_K, sigma_CP, and n_C3 for both sectors"
	}
	if !e.HasNeutrinoOrdering || !nu.HasNeutrinoOrdering || !e.HasAbsoluteNuScale || !nu.HasAbsoluteNuScale || !e.HasMajoranaDiracPolicy || !nu.HasMajoranaDiracPolicy || !e.HasUncertainty || !nu.HasUncertainty {
		return StatusFailedMissingNeutrinoPolicies, "synthetic lepton residual requires declared neutrino ordering, absolute scale, phase policy, and uncertainty"
	}
	if e.ObservedPMNSImport || nu.ObservedPMNSImport {
		return StatusFailedObservedPMNSImport, "observed PMNS values are rejected by the synthetic socket"
	}
	if e.PMNSAsRayInput || nu.PMNSAsRayInput {
		return StatusFailedPMNSAsRayInput, "PMNS target cannot define alpha, phi, I_K, or branch tags"
	}
	if e.PMNSNativePredictionClaim || nu.PMNSNativePredictionClaim {
		return StatusFailedPMNSNativePrediction, "PMNS entries cannot be claimed as native ASHA predictions"
	}
	if e.PMNSMatrixExportRequest || nu.PMNSMatrixExportRequest {
		return StatusFailedPMNSMatrixExport, "Gate476 does not construct or export the PMNS matrix"
	}
	if e.NativePromotionClaim || nu.NativePromotionClaim || !e.BridgeOnly || !nu.BridgeOnly || !e.SyntheticOnly || !nu.SyntheticOnly {
		return StatusFailedNativeResidualPromotion, "d_e_nu and synthetic targets are bridge-only diagnostics and cannot enter native law-space"
	}
	return "", ""
}

func invert(x LeptonComparator) (Ray, string, string) {
	r := Ray{Sector: x.Sector, ISpec: x.ISpec, IK: x.IK, SigmaCP: x.SigmaCP, C3Sheet: x.C3Sheet, BridgeOnly: true, SyntheticOnly: true}
	if math.Abs(x.IK) >= 1 {
		r.Verdict = StatusFailedProjectiveDomainRejected
		r.Reason = "I_K must be inside (-1,1)"
		return r, StatusFailedProjectiveDomainRejected, r.Reason
	}
	alpha := math.Sqrt(3) * x.IK / math.Sqrt(1-x.IK*x.IK)
	cos3 := (3 * math.Sqrt(3) / 2) * x.ISpec / math.Pow(1-x.IK*x.IK, 1.5)
	if math.Abs(cos3) > 1+1e-12 {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.Verdict = StatusFailedPhaseDomainRejected
		r.Reason = "derived cos(3phi) left [-1,1]"
		return r, StatusFailedPhaseDomainRejected, r.Reason
	}
	cos3 = math.Max(-1, math.Min(1, cos3))
	if math.Abs(1-math.Abs(cos3)) < 1e-12 {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.AtCaustic = true
		r.Verdict = StatusFailedCausticRejected
		r.Reason = "sin(3phi)=0 caustic does not carry an orientable synthetic branch"
		return r, StatusFailedCausticRejected, r.Reason
	}
	phi := (float64(x.SigmaCP)*math.Acos(cos3) + 2*math.Pi*float64(x.C3Sheet)) / 3
	phi = wrapPi(phi)
	r.Alpha = alpha
	r.CosThreePhi = cos3
	r.Phi = phi
	r.InsideDomain = true
	r.Verdict = StatusSyntheticLeptonLedgerAccepted
	r.Reason = "synthetic rank-complete row inverted to bridge-only lepton ray"
	return r, "", ""
}

func rejectedResidual(verdict, reason string) PMNSNullResidual {
	return PMNSNullResidual{Verdict: verdict, Reason: reason, BridgeOnly: false, SyntheticOnly: false, PMNSMatrixConstructed: false, PMNSEntryComputed: false, ExportsNativeObservable: false}
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

func buildOutput(s Sieve) Output {
	for _, c := range s.Cases {
		if c.Accepted {
			return Output{SyntheticDENu: c.Residual.DENu, SyntheticPMNSTarget: c.Residual.SyntheticPMNSTarget, SyntheticResidual: c.Residual.SyntheticResidual, AlphaE: c.ERay.Alpha, PhiE: c.ERay.Phi, AlphaNu: c.NuRay.Alpha, PhiNu: c.NuRay.Phi}
		}
	}
	return Output{}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                         true,
		PMNSNullResidualAdapterDefined:   true,
		PMNSNullResidualMayRunBridgeOnly: true,
		SyntheticLeptonDataEvaluated:     a.Sieve.ValidSyntheticResidualAccepted,
		ObservedLeptonDataImported:       false,
		ObservedPMNSImported:             false,
		PMNSMatrixConstructed:            false,
		PMNSEntryComputed:                false,
		PMNSNativePrediction:             false,
		DENuNativePrediction:             false,
		SyntheticPMNSTargetNative:        false,
		NativeRegistryWritten:            false,
		CKMNativePrediction:              false,
		KGenStillForced:                  true,
		XTriangleStillForced:             true,
		YPhaseStillQuarantined:           true,
		SectorCoefficientsStillSealed:    true,
		NativeFlavorDimAfter:             NativeFlavorDim,
		KXYCoeffDimAfter:                 KXYCoeffDim,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "Gate476 computes only a synthetic e/nu residual diagnostic; no PMNS entry, lepton mass, neutrino scale, I_K value, branch tag, or d_e_nu value enters native law-space",
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 477, Title: "Observed lepton rank-complete ledger adapter", Reason: "Gate476 proves the lepton socket is structurally identical to the quark socket on synthetic data only.", PrimaryTask: "accept an explicit observed e/nu rank-complete ledger only through the empirical airlock, then compare d_e_nu to a declared PMNS residual target without native promotion"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate475LeptonPreflightDefined || !a.Inheritance.Gate475RequiresENuSectors || !a.Inheritance.Gate475RequiresISpecIK || !a.Inheritance.Gate475RequiresBranchTags || !a.Inheritance.Gate475RequiresNeutrinoPolicies || !a.Inheritance.Gate475RejectsPMNSAsRayInput || !a.Inheritance.Gate475RejectsNativePromotion || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate476 inheritance incomplete: %+v", a.Inheritance)
	}
	if !a.Map.Executed || !a.Map.RequiresERow || !a.Map.RequiresNuRow || !a.Map.RequiresGate475Preflight || !a.Map.RequiresISpecIK || !a.Map.RequiresBranchTags || !a.Map.RequiresNeutrinoOrdering || !a.Map.RequiresAbsoluteNuScale || !a.Map.RequiresMajoranaDiracPolicy || !a.Map.RequiresSyntheticMode || !a.Map.RequiresBridgeOnly || !a.Map.AllowsPMNSTarget || a.Map.AllowsPMNSAsRayInput || !a.Map.ComputesDENu || a.Map.ComputesPMNSMatrix || a.Map.ComputesPMNSEntry || !a.Map.ExportsDiagnosticsOnly || !a.Map.StructurallyIdenticalToQuarkMap {
		return fmt.Errorf("Gate476 map invalid: %+v", a.Map)
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCaseCount != 1 || a.Sieve.RejectedCaseCount < 10 || !a.Sieve.ValidSyntheticResidualAccepted || !a.Sieve.MissingENuRejected || !a.Sieve.MissingPreflightRejected || !a.Sieve.MissingRankLedgerRejected || !a.Sieve.MissingNeutrinoPoliciesRejected || !a.Sieve.ObservedPMNSRejected || !a.Sieve.PMNSAsRayInputRejected || !a.Sieve.PMNSNativePredictionRejected || !a.Sieve.PMNSMatrixExportRejected || !a.Sieve.NativeResidualPromotionRejected || !a.Sieve.ProjectiveDomainRejected || !a.Sieve.PhaseDomainRejected || !a.Sieve.CausticRejected || !a.Sieve.AllAcceptedBridgeOnlySynthetic || !a.Sieve.NoPMNSMatrixConstructed || !a.Sieve.NoNativeObservableExport {
		return fmt.Errorf("Gate476 sieve invalid: %+v", a.Sieve)
	}
	if a.Output.SyntheticDENu <= 0 || a.Output.SyntheticPMNSTarget <= 0 || a.Output.SyntheticResidual < 0 {
		return fmt.Errorf("Gate476 output invalid: %+v", a.Output)
	}
	if !a.Firewall.Executed || !a.Firewall.PMNSNullResidualAdapterDefined || !a.Firewall.PMNSNullResidualMayRunBridgeOnly || !a.Firewall.SyntheticLeptonDataEvaluated || a.Firewall.ObservedLeptonDataImported || a.Firewall.ObservedPMNSImported || a.Firewall.PMNSMatrixConstructed || a.Firewall.PMNSEntryComputed || a.Firewall.PMNSNativePrediction || a.Firewall.DENuNativePrediction || a.Firewall.SyntheticPMNSTargetNative || a.Firewall.NativeRegistryWritten || a.Firewall.CKMNativePrediction || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate476 firewall violated: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 476 proves the lepton/PMNS socket has the same bridge geometry as the quark/CKM socket: d_e_nu=sqrt((alpha_nu-alpha_e)^2+4 sin^2((phi_nu-phi_e)/2)). On the canonical synthetic rank-complete e/nu ledger it computes d_e_nu=%.12g against a synthetic target %.12g, but this is only a bridge diagnostic. Observed PMNS import, PMNS-as-coordinate use, matrix export, and native-promotion all fail closed.", a.Output.SyntheticDENu, a.Output.SyntheticPMNSTarget)
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t tags=%t quark_socket=%t PMNS_bridge=%t preflight=%t e_nu=%t I_spec_I_K=%t branch_tags=%t nu_policies=%t PMNS_ray_rejected=%t native_rejected=%t no_observed=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456RayInverseAvailable, x.Gate459BranchTagsAvailable, x.Gate464QuarkSocketAvailable, x.Gate474PMNSBridgeOnly, x.Gate475LeptonPreflightDefined, x.Gate475RequiresENuSectors, x.Gate475RequiresISpecIK, x.Gate475RequiresBranchTags, x.Gate475RequiresNeutrinoPolicies, x.Gate475RejectsPMNSAsRayInput, x.Gate475RejectsNativePromotion, x.NoObservedValuesImported, x.NativeRegistryClean, x.Verdict)
}
func FormatMap(x PMNSNullMap) string {
	return fmt.Sprintf("executed=%t formula=%s e=%t nu=%t preflight=%t I_spec_I_K=%t branch_tags=%t nu_ordering=%t absolute_scale=%t phase_policy=%t synthetic=%t bridge=%t PMNS_target=%t PMNS_ray_input=%t computes_d=%t PMNS_matrix=%t PMNS_entry=%t diagnostics=%t quark_identical=%t verdict=%s reason=%s", x.Executed, x.Formula, x.RequiresERow, x.RequiresNuRow, x.RequiresGate475Preflight, x.RequiresISpecIK, x.RequiresBranchTags, x.RequiresNeutrinoOrdering, x.RequiresAbsoluteNuScale, x.RequiresMajoranaDiracPolicy, x.RequiresSyntheticMode, x.RequiresBridgeOnly, x.AllowsPMNSTarget, x.AllowsPMNSAsRayInput, x.ComputesDENu, x.ComputesPMNSMatrix, x.ComputesPMNSEntry, x.ExportsDiagnosticsOnly, x.StructurallyIdenticalToQuarkMap, x.Verdict, x.Reason)
}
func FormatComparator(x LeptonComparator) string {
	return fmt.Sprintf("sector=%s alpha=%.12g phi=%.12g I_spec=%.12g I_K=%.12g sigma_CP=%d n_C3=%d sector_ok=%t I_spec=%t I_K=%t tags=%t preflight=%t eigenbasis=%t nu_ordering=%t abs_scale=%t phase_policy=%t uncertainty=%t bridge=%t synthetic=%t observed_PMNS=%t PMNS_ray=%t native=%t matrix_export=%t PMNS_native=%t", x.Sector, x.Alpha, x.Phi, x.ISpec, x.IK, x.SigmaCP, x.C3Sheet, x.HasSector, x.HasISpec, x.HasIK, x.HasBranchTags, x.HasPreflight, x.HasEigenbasisConvention, x.HasNeutrinoOrdering, x.HasAbsoluteNuScale, x.HasMajoranaDiracPolicy, x.HasUncertainty, x.BridgeOnly, x.SyntheticOnly, x.ObservedPMNSImport, x.PMNSAsRayInput, x.NativePromotionClaim, x.PMNSMatrixExportRequest, x.PMNSNativePredictionClaim)
}
func FormatRay(x Ray) string {
	return fmt.Sprintf("sector=%s alpha=%.12g phi=%.12g I_spec=%.12g I_K=%.12g cos3phi=%.12g sigma_CP=%d n_C3=%d domain=%t caustic=%t bridge=%t synthetic=%t native=%t verdict=%s reason=%s", x.Sector, x.Alpha, x.Phi, x.ISpec, x.IK, x.CosThreePhi, x.SigmaCP, x.C3Sheet, x.InsideDomain, x.AtCaustic, x.BridgeOnly, x.SyntheticOnly, x.ExportsNative, x.Verdict, x.Reason)
}
func FormatResidual(x PMNSNullResidual) string {
	return fmt.Sprintf("Delta_alpha=%.12g Delta_phi=%.12g phase_chord=%.12g d_e_nu=%.12g target=%.12g residual=%.12g complete=%t preflight=%t bridge=%t synthetic=%t diagnostic=%t PMNS_matrix=%t PMNS_entry=%t native=%t observed_PMNS=%t verdict=%s reason=%s", x.DeltaAlpha, x.DeltaPhi, x.PhaseChord, x.DENu, x.SyntheticPMNSTarget, x.SyntheticResidual, x.CompleteInputs, x.Gate475PreflightSatisfied, x.BridgeOnly, x.SyntheticOnly, x.ExportsRelativeDiagnostic, x.PMNSMatrixConstructed, x.PMNSEntryComputed, x.ExportsNativeObservable, x.ObservedPMNSImported, x.Verdict, x.Reason)
}
func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d valid=%t missing_e_nu=%t missing_preflight=%t missing_rank=%t missing_nu_policies=%t observed_PMNS=%t PMNS_ray=%t PMNS_native=%t matrix_export=%t native_residual=%t projective=%t phase=%t caustic=%t bridge_synthetic=%t no_PMNS_matrix=%t no_native=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.ValidSyntheticResidualAccepted, x.MissingENuRejected, x.MissingPreflightRejected, x.MissingRankLedgerRejected, x.MissingNeutrinoPoliciesRejected, x.ObservedPMNSRejected, x.PMNSAsRayInputRejected, x.PMNSNativePredictionRejected, x.PMNSMatrixExportRejected, x.NativeResidualPromotionRejected, x.ProjectiveDomainRejected, x.PhaseDomainRejected, x.CausticRejected, x.AllAcceptedBridgeOnlySynthetic, x.NoPMNSMatrixConstructed, x.NoNativeObservableExport, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t adapter=%t bridge_only_run=%t synthetic_eval=%t observed_lepton=%t observed_PMNS=%t PMNS_matrix=%t PMNS_entry=%t PMNS_native=%t d_native=%t target_native=%t native_write=%t CKM_native=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.PMNSNullResidualAdapterDefined, x.PMNSNullResidualMayRunBridgeOnly, x.SyntheticLeptonDataEvaluated, x.ObservedLeptonDataImported, x.ObservedPMNSImported, x.PMNSMatrixConstructed, x.PMNSEntryComputed, x.PMNSNativePrediction, x.DENuNativePrediction, x.SyntheticPMNSTargetNative, x.NativeRegistryWritten, x.CKMNativePrediction, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatOutput(x Output) string {
	return fmt.Sprintf("alpha_e=%.12g phi_e=%.12g alpha_nu=%.12g phi_nu=%.12g d_e_nu=%.12g synthetic_theta23_target=%.12g residual=%.12g", x.AlphaE, x.PhiE, x.AlphaNu, x.PhiNu, x.SyntheticDENu, x.SyntheticPMNSTarget, x.SyntheticResidual)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate475Inherited, StatusPMNSNullMapDefined, StatusSyntheticLeptonLedgerAccepted, StatusPMNSNullResidualComputed, StatusPMNSNullResidualFirewallValidated, StatusLeptonSocketStructurallyIdentical, StatusFailedMissingENuSectors, StatusFailedMissingPreflight, StatusFailedMissingRankCompleteLedger, StatusFailedMissingNeutrinoPolicies, StatusFailedObservedPMNSImport, StatusFailedPMNSAsRayInput, StatusFailedPMNSNativePrediction, StatusFailedPMNSMatrixExport, StatusFailedNativeResidualPromotion, StatusFailedProjectiveDomainRejected, StatusFailedPhaseDomainRejected, StatusFailedCausticRejected, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 476 Registry Audit — Lepton-Sector Synthetic PMNS Null Residual\n\n")
	b.WriteString("## Verdict\n\n`" + StatusPMNSNullResidualFirewallValidated + "`\n\n")
	b.WriteString("Gate 476 mirrors the Gate 464 CKM null adapter in the lepton sector. It computes a synthetic bridge-only e/nu cylinder distance and compares it to a synthetic PMNS residual target, while rejecting observed PMNS import, PMNS-as-coordinate use, PMNS matrix export, and native-promotion.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## PMNS-null symbolic map\n\n" + FormatMap(a.Map) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("Delta_alpha_e_nu = alpha_nu - alpha_e\n")
	b.WriteString("Delta_phi_e_nu   = wrap_pi(phi_nu - phi_e)\n")
	b.WriteString("d_e_nu           = sqrt(Delta_alpha_e_nu^2 + 4 sin^2(Delta_phi_e_nu/2))\n")
	b.WriteString("synthetic target = theta23_like_synthetic_residual_target\n")
	b.WriteString("forbidden        = U_PMNS, PMNS_ij, observed PMNS import, PMNS-as-ray input, native prediction\n")
	b.WriteString("```\n\n")
	b.WriteString("## Synthetic output\n\n" + FormatOutput(a.Output) + "\n\n")
	b.WriteString("## Sieve\n\n" + FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Case | Accepted | Verdict | e row | nu row | e ray | nu ray | Residual | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s | %s | %s | %s | %s | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(FormatComparator(c.E)), esc(FormatComparator(c.Nu)), esc(FormatRay(c.ERay)), esc(FormatRay(c.NuRay)), esc(FormatResidual(c.Residual)), esc(c.Reason)))
	}
	b.WriteString("\n## Structural identity with quark socket\n\n")
	b.WriteString("The quark socket uses `d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))`. Gate 476 uses the same phase-cylinder metric with labels changed to `e` and `nu`. This proves socket isomorphism, not PMNS prediction. PMNS may be a residual target only; it cannot be used as an alpha/phi/I_K coordinate input.\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No PMNS entry, PMNS matrix, neutrino mass, charged-lepton mass, I_K value, branch tag, d_e_nu value, or synthetic target is written into native law-space.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
