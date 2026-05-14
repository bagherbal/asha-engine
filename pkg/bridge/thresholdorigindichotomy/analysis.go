// Package thresholdorigindichotomy implements Gate 179: threshold-origin
// dichotomy / new-sector versus continuum-decoupling bridge audit.
//
// Gate 178 proved that no currently derived object is a complete finite
// threshold operator. Gate 179 asks the next structural question: if
// non-universal threshold deformations are needed, where could they lawfully
// originate? The gate separates three possibilities:
//
//  1. existing finite spectral anchors, made physical by a missing continuum
//     decoupling bridge;
//  2. genuinely new finite sectors with representation-complete heavy modes;
//  3. externally fitted phenomenological vectors, which are useful witnesses
//     but forbidden as theorem input.
//
// The gate derives no thresholds. It records a dichotomy theorem: the present
// finite data are exhausted as threshold sources, and any future threshold
// theorem must either add a continuum bridge for existing spectra or derive new
// finite sectors.
package thresholdorigindichotomy

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitethresholdoperator"
)

type OriginClass string

const (
	ExistingFiniteBridge OriginClass = "existing-finite-continuum-bridge"
	NewFiniteSector      OriginClass = "new-finite-sector"
	SchemeNormalization  OriginClass = "scheme-or-normalization"
	PhenomenologicalFit  OriginClass = "phenomenological-fit"
)

type OriginBranch struct {
	Name                               string
	Class                              OriginClass
	Inputs                             []string
	RequiredNewObjects                 []string
	UsesCurrentFiniteData              bool
	RequiresContinuumBridge            bool
	RequiresNewFiniteSector            bool
	UsesObservedComparison             bool
	CanInPrincipleGiveNonUniversal     bool
	CurrentlyDerived                   bool
	PromotesGate177Repair              bool
	StrictNullityReductionAllowed      bool
	ConditionalNullityReductionAllowed bool
	AdmissibleFutureProgram            bool
	Verdict                            string
}

type ContinuumBridgeAudit struct {
	CandidateExistingAnchors         []string
	OrientedFourCycleRequired        bool
	PrincipalBundleRequired          bool
	ChernWeilNormalizationRequired   bool
	ContinuumTraceConventionRequired bool
	LocalFieldMapRequired            bool
	PhysicalMassUnitRequired         bool
	ActivationPredicateRequired      bool
	DecouplingLawRequired            bool
	GaugeRepresentationRowsRequired  bool
	AllRequiredObjectsPresent        bool
	BridgeDerived                    bool
	CanPromoteGate177Repair          bool
	Verdict                          string
}

type NewSectorAudit struct {
	RequiredFeatures                      []string
	KnownFiniteInventoryExhausted         bool
	DerivedNewSectors                     int
	RepresentationCompleteHeavyMultiplets int
	AnomalyCompatibleHeavyMultiplets      int
	VacuumCompatibleHeavyMultiplets       int
	CanonicalMassSpectrumCount            int
	CanGenerateNonUniversalDeltaB         bool
	CanPromoteGate177Repair               bool
	Verdict                               string
}

type RejectedOrigin struct {
	Name              string
	Reason            string
	UsesObservedInput bool
	CanServeAsTheorem bool
	PreservesFirewall bool
}

type DichotomyAudit struct {
	PreviousGate178NoGo                 bool
	CurrentFiniteThresholdExhausted     bool
	BranchesAudited                     int
	ContinuumBranches                   int
	NewSectorBranches                   int
	FitBranches                         int
	SchemeBranches                      int
	SurvivingProgrammaticOrigins        []string
	ObservedFitRejectedAsOrigin         bool
	SchemeOnlyRejectedAsThresholdOrigin bool
	DichotomyCompleteAtCurrentStage     bool
	ThresholdOriginDerived              bool
	Gate177RepairPromoted               bool
	Verdict                             string
}

type FirewallAudit struct {
	UsesObservedInputForDerivation bool
	NonUniversalDeltaBDerived      bool
	ThresholdOperatorDerived       bool
	ThresholdCorrectedBetaDerived  bool
	PhysicalConstantsDerived       bool
	BoundaryScaleDerived           bool
	StrictNullityBefore            int
	StrictNullityAfter             int
	ConditionalNullityBefore       int
	ConditionalNullityAfter        int
	SealedClaims                   []string
	OpenRequirements               []string
	RecommendedNextGate            string
	Verdict                        string
}

type Analysis struct {
	PreviousGate178 finitethresholdoperator.Analysis
	Branches        []OriginBranch
	Continuum       ContinuumBridgeAudit
	NewSector       NewSectorAudit
	Rejected        []RejectedOrigin
	Dichotomy       DichotomyAudit
	Firewall        FirewallAudit
	TruthStatement  string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := finitethresholdoperator.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev finitethresholdoperator.Analysis) (Analysis, error) {
	if prev.Firewall.StrictNullityAfter != 3 || prev.Firewall.ConditionalNullityAfter != 2 {
		return Analysis{}, fmt.Errorf("Gate 179 requires Gate 178 to leave strict nullity 3 and conditional nullity 2")
	}
	if prev.Firewall.ThresholdOperatorDerived || prev.Firewall.NonUniversalDeltaBDerived || prev.Firewall.ThresholdCorrectedBetaDerived || prev.Firewall.Gate177RepairPromoted {
		return Analysis{}, fmt.Errorf("Gate 179 requires Gate 178 threshold firewall to remain closed")
	}
	if prev.Firewall.UsesObservedInputForDerivation || prev.DeltaBWitness.CanBePromotedToFiniteOperator {
		return Analysis{}, fmt.Errorf("Gate 179 refuses non-quarantined observed input")
	}
	if !prev.Requirements.NoCandidateHasAllPieces || prev.Requirements.FiniteDerivedThresholdOps != 0 {
		return Analysis{}, fmt.Errorf("Gate 179 requires Gate 178 to prove no complete current threshold operator")
	}

	continuum := buildContinuumBridgeAudit(prev)
	newSector := buildNewSectorAudit()
	branches := buildOriginBranches(continuum, newSector)
	rejected := buildRejectedOrigins()
	dichotomy := auditDichotomy(prev, branches)
	firewall := FirewallAudit{
		UsesObservedInputForDerivation: false,
		NonUniversalDeltaBDerived:      false,
		ThresholdOperatorDerived:       false,
		ThresholdCorrectedBetaDerived:  false,
		PhysicalConstantsDerived:       false,
		BoundaryScaleDerived:           false,
		StrictNullityBefore:            prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:             prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:       prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:        prev.Firewall.ConditionalNullityAfter,
		SealedClaims: []string{
			"current finite spectra alone determine non-universal threshold beta rows",
			"Gate 177's fitted Δb vector is a finite threshold origin",
			"a normalization or scheme convention is equivalent to a sector-specific decoupling spectrum",
			"new heavy sectors may be postulated without representation, anomaly, vacuum, and mass checks",
		},
		OpenRequirements: []string{
			"continuum decoupling bridge for existing finite spectral anchors",
			"or a newly derived finite heavy sector with representation-complete beta rows",
			"physical mass unit or activation predicate",
			"matching law connecting finite spectra to Δb_i",
		},
		RecommendedNextGate: "Gate 180 — continuum decoupling bridge axiom inventory / finite heat-kernel matching preflight",
		Verdict:             "The threshold origin is not derived. Future work must choose between an existing-spectrum continuum bridge and genuinely new finite heavy sectors; observed-fit vectors remain quarantined witnesses.",
	}
	truth := "Gate 179 converts the threshold failure into an origin dichotomy. Gate 178 exhausted all currently derived finite threshold candidates. Therefore non-universal thresholds, if they exist, must come either from a new continuum-decoupling bridge acting on existing finite spectra, or from new finite heavy sectors not yet present in the algebra. Phenomenological Δb vectors and normalization conventions are comparison tools, not threshold origins. No nullity reduction is allowed."
	return Analysis{PreviousGate178: prev, Branches: branches, Continuum: continuum, NewSector: newSector, Rejected: rejected, Dichotomy: dichotomy, Firewall: firewall, TruthStatement: truth}, nil
}

func buildContinuumBridgeAudit(prev finitethresholdoperator.Analysis) ContinuumBridgeAudit {
	anchors := make([]string, 0)
	for _, c := range prev.Candidates {
		if c.FiniteSpectrum && c.ExactFiniteData {
			anchors = append(anchors, c.Name)
		}
	}
	return ContinuumBridgeAudit{
		CandidateExistingAnchors:         anchors,
		OrientedFourCycleRequired:        true,
		PrincipalBundleRequired:          true,
		ChernWeilNormalizationRequired:   true,
		ContinuumTraceConventionRequired: true,
		LocalFieldMapRequired:            true,
		PhysicalMassUnitRequired:         true,
		ActivationPredicateRequired:      true,
		DecouplingLawRequired:            true,
		GaugeRepresentationRowsRequired:  true,
		AllRequiredObjectsPresent:        false,
		BridgeDerived:                    false,
		CanPromoteGate177Repair:          false,
		Verdict:                          "existing spectra can only become thresholds after a finite-to-continuum field, mass, activation, and decoupling bridge is derived",
	}
}

func buildNewSectorAudit() NewSectorAudit {
	return NewSectorAudit{
		RequiredFeatures: []string{
			"finite carrier not already counted in baseline fields",
			"canonical gauge representation under SU(3)×SU(2)×U(1)",
			"finite mass or activation scale",
			"decoupling/matching rule",
			"beta-index row contribution Δb_i",
			"anomaly and vacuum compatibility checks",
		},
		KnownFiniteInventoryExhausted:         true,
		DerivedNewSectors:                     0,
		RepresentationCompleteHeavyMultiplets: 0,
		AnomalyCompatibleHeavyMultiplets:      0,
		VacuumCompatibleHeavyMultiplets:       0,
		CanonicalMassSpectrumCount:            0,
		CanGenerateNonUniversalDeltaB:         false,
		CanPromoteGate177Repair:               false,
		Verdict:                               "new sectors are a valid future origin class, but none are currently derived or representation-complete",
	}
}

func buildOriginBranches(c ContinuumBridgeAudit, n NewSectorAudit) []OriginBranch {
	return []OriginBranch{
		{
			Name: "existing finite spectra plus continuum decoupling bridge", Class: ExistingFiniteBridge,
			Inputs:                c.CandidateExistingAnchors,
			RequiredNewObjects:    []string{"finite-to-continuum local field map", "physical mass unit", "activation predicate", "decoupling/matching law", "gauge representation rows"},
			UsesCurrentFiniteData: true, RequiresContinuumBridge: true, CanInPrincipleGiveNonUniversal: true,
			CurrentlyDerived: false, PromotesGate177Repair: false, StrictNullityReductionAllowed: false, ConditionalNullityReductionAllowed: false, AdmissibleFutureProgram: true,
			Verdict: "surviving origin branch, but bridge-required and not yet a theorem",
		},
		{
			Name: "new finite heavy sector", Class: NewFiniteSector,
			Inputs: []string{"none currently derived"}, RequiredNewObjects: n.RequiredFeatures,
			RequiresNewFiniteSector: true, CanInPrincipleGiveNonUniversal: true,
			CurrentlyDerived: false, PromotesGate177Repair: false, StrictNullityReductionAllowed: false, ConditionalNullityReductionAllowed: false, AdmissibleFutureProgram: true,
			Verdict: "surviving origin branch only if new algebraic sectors are derived and pass compatibility checks",
		},
		{
			Name: "normalization or renormalization-scheme convention", Class: SchemeNormalization,
			Inputs: []string{"absolute trace convention", "scheme convention"}, RequiredNewObjects: []string{"sector-specific decoupling data, if intended as threshold"},
			CanInPrincipleGiveNonUniversal: false, CurrentlyDerived: false, PromotesGate177Repair: false, AdmissibleFutureProgram: false,
			Verdict: "can move an overall normalization or convention, but cannot itself produce non-universal threshold origins",
		},
		{
			Name: "Gate 177 fitted Δb vector", Class: PhenomenologicalFit,
			Inputs: []string{"external M_Z comparison", "minimum-norm criterion"}, UsesObservedComparison: true, CanInPrincipleGiveNonUniversal: true,
			CurrentlyDerived: false, PromotesGate177Repair: false, AdmissibleFutureProgram: false,
			Verdict: "comparison witness only; forbidden as theorem-level threshold origin",
		},
	}
}

func buildRejectedOrigins() []RejectedOrigin {
	return []RejectedOrigin{
		{Name: "observed-fit Δb as finite source", Reason: "imports external comparison values and a minimization criterion; no finite operator is supplied", UsesObservedInput: true, CanServeAsTheorem: false, PreservesFirewall: true},
		{Name: "universal threshold or scheme shift", Reason: "sector differences cancel universal shifts, so it cannot repair non-universal running", CanServeAsTheorem: false, PreservesFirewall: true},
		{Name: "baseline scalar row as threshold", Reason: "already counted as a baseline field contribution and has no heavy activation law", CanServeAsTheorem: false, PreservesFirewall: true},
	}
}

func auditDichotomy(prev finitethresholdoperator.Analysis, branches []OriginBranch) DichotomyAudit {
	d := DichotomyAudit{PreviousGate178NoGo: prev.Requirements.NoCandidateHasAllPieces && !prev.Firewall.ThresholdOperatorDerived, CurrentFiniteThresholdExhausted: true}
	for _, b := range branches {
		d.BranchesAudited++
		switch b.Class {
		case ExistingFiniteBridge:
			d.ContinuumBranches++
		case NewFiniteSector:
			d.NewSectorBranches++
		case PhenomenologicalFit:
			d.FitBranches++
		case SchemeNormalization:
			d.SchemeBranches++
		}
		if b.AdmissibleFutureProgram {
			d.SurvivingProgrammaticOrigins = append(d.SurvivingProgrammaticOrigins, b.Name)
		}
		if b.CurrentlyDerived || b.PromotesGate177Repair || b.StrictNullityReductionAllowed || b.ConditionalNullityReductionAllowed {
			d.ThresholdOriginDerived = true
			d.Gate177RepairPromoted = true
		}
	}
	d.ObservedFitRejectedAsOrigin = d.FitBranches == 1
	d.SchemeOnlyRejectedAsThresholdOrigin = d.SchemeBranches == 1
	d.DichotomyCompleteAtCurrentStage = d.PreviousGate178NoGo && d.CurrentFiniteThresholdExhausted && d.ContinuumBranches == 1 && d.NewSectorBranches == 1 && len(d.SurvivingProgrammaticOrigins) == 2 && d.ObservedFitRejectedAsOrigin && d.SchemeOnlyRejectedAsThresholdOrigin && !d.ThresholdOriginDerived
	d.Verdict = "At the current stage, lawful threshold origins are exhausted into two open program branches: existing finite spectra plus a continuum-decoupling bridge, or genuinely new finite heavy sectors. Neither is derived."
	return d
}

func FormatBranch(b OriginBranch) string {
	return fmt.Sprintf("%s[%s](current=%t,contBridge=%t,newSector=%t,obs=%t,nonunivPossible=%t,derived=%t,promotes=%t,strictNullity=%t,future=%t): %s", b.Name, b.Class, b.UsesCurrentFiniteData, b.RequiresContinuumBridge, b.RequiresNewFiniteSector, b.UsesObservedComparison, b.CanInPrincipleGiveNonUniversal, b.CurrentlyDerived, b.PromotesGate177Repair, b.StrictNullityReductionAllowed, b.AdmissibleFutureProgram, b.Verdict)
}

func FormatBranches(xs []OriginBranch) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatBranch(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatContinuum(c ContinuumBridgeAudit) string {
	return fmt.Sprintf("anchors=%d fourCycle=%t bundle=%t cernWeil=%t trace=%t fieldMap=%t mass=%t activation=%t decoupling=%t repRows=%t allPresent=%t derived=%t promotes=%t: %s", len(c.CandidateExistingAnchors), c.OrientedFourCycleRequired, c.PrincipalBundleRequired, c.ChernWeilNormalizationRequired, c.ContinuumTraceConventionRequired, c.LocalFieldMapRequired, c.PhysicalMassUnitRequired, c.ActivationPredicateRequired, c.DecouplingLawRequired, c.GaugeRepresentationRowsRequired, c.AllRequiredObjectsPresent, c.BridgeDerived, c.CanPromoteGate177Repair, c.Verdict)
}

func FormatNewSector(n NewSectorAudit) string {
	return fmt.Sprintf("features=[%s] inventoryExhausted=%t newSectors=%d repComplete=%d anomaly=%d vacuum=%d mass=%d nonuniv=%t promotes=%t: %s", strings.Join(n.RequiredFeatures, "; "), n.KnownFiniteInventoryExhausted, n.DerivedNewSectors, n.RepresentationCompleteHeavyMultiplets, n.AnomalyCompatibleHeavyMultiplets, n.VacuumCompatibleHeavyMultiplets, n.CanonicalMassSpectrumCount, n.CanGenerateNonUniversalDeltaB, n.CanPromoteGate177Repair, n.Verdict)
}

func FormatRejected(xs []RejectedOrigin) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(obs=%t,theorem=%t,firewall=%t): %s", x.Name, x.UsesObservedInput, x.CanServeAsTheorem, x.PreservesFirewall, x.Reason))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatDichotomy(d DichotomyAudit) string {
	return fmt.Sprintf("prevNoGo=%t exhausted=%t branches=%d continuum=%d newSector=%d fit=%d scheme=%d surviving=[%s] fitRejected=%t schemeRejected=%t complete=%t originDerived=%t promotes=%t: %s", d.PreviousGate178NoGo, d.CurrentFiniteThresholdExhausted, d.BranchesAudited, d.ContinuumBranches, d.NewSectorBranches, d.FitBranches, d.SchemeBranches, strings.Join(d.SurvivingProgrammaticOrigins, "; "), d.ObservedFitRejectedAsOrigin, d.SchemeOnlyRejectedAsThresholdOrigin, d.DichotomyCompleteAtCurrentStage, d.ThresholdOriginDerived, d.Gate177RepairPromoted, d.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("obsInput=%t Δb=%t thresholdOp=%t thresholdBeta=%t constants=%t boundary=%t strict=%d->%d conditional=%d->%d sealed=[%s] open=[%s] next=%s: %s", f.UsesObservedInputForDerivation, f.NonUniversalDeltaBDerived, f.ThresholdOperatorDerived, f.ThresholdCorrectedBetaDerived, f.PhysicalConstantsDerived, f.BoundaryScaleDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.ConditionalNullityBefore, f.ConditionalNullityAfter, strings.Join(f.SealedClaims, "; "), strings.Join(f.OpenRequirements, "; "), f.RecommendedNextGate, f.Verdict)
}
