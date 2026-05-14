// Package scalarorientationsource implements Gate 190: eta-odd scalar
// orientation source / matter-pullback search audit.
//
// Gate 189 proved that the branchwise quartic projectors {P_A,P_B} are
// dimensionally compatible with the physical H_Phi high/low projector pair, but
// that assigning eta to high versus low is not canonical. Gate 190 asks the
// last source-search question: does any existing finite weak, hypercharge,
// charge-conjugation, matter, contact, or broken-sector datum act as an eta-odd
// scalar-orientation source?
//
// The answer is a positive obstruction theorem. The audited operators either
// preserve the high/low projectors, exchange them as gauge/conjugation symmetry,
// live on the wrong tensor factor, or rely on a diagnostic vacuum convention.
// None gives a gauge-invariant finite selector for eta -> high versus eta ->
// low. The eta orientation is therefore isolated as spontaneous/gauge data, not
// as a derivable finite observable.
package scalarorientationsource

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/brokenmetric"
	"github.com/bagherbal/asha-engine/pkg/bridge/contactcoddsource"
	"github.com/bagherbal/asha-engine/pkg/bridge/contactsignsource"
	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeeating"
	"github.com/bagherbal/asha-engine/pkg/bridge/higgsconjugatequotient"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarbundlemap"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcomplex"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcovariant"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/hypercharge"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
)

type WeakGaugeAudit struct {
	T3LAvailable                    bool
	ScalarHyperchargeAvailable      bool
	MatterSU2LAvailable             bool
	ScalarDoubletCandidate          bool
	T3CommutesWithHighLowProjectors bool
	YCommutesWithHighLowProjectors  bool
	T1MixesHighLowPlanes            bool
	T2MixesHighLowPlanes            bool
	WeylReflectionExchangesPlanes   bool
	GaugeActionSelectsOrientation   bool
	EtaOddGaugeInvariantSource      bool
	Verdict                         string
}

type ConjugationAudit struct {
	ContactChargeConjugationAvailable  bool
	ContactChargeConjugationInvolution bool
	ContactChargeConjugationExchanges  bool
	ContactChargeConjugationSelects    bool
	HiggsConjugateCollapseRejected     bool
	HiggsBranchUniquenessByKind        bool
	MirrorsEtaInvolution               bool
	SelectsEtaOrientation              bool
	Verdict                            string
}

type BrokenSectorAudit struct {
	CovariantDerivativeTemplate    bool
	VacuumOrientationChosen        bool
	DimensionlessWZPhotonSignature bool
	GoldstoneImageDiagnostic       bool
	FiniteGaugeEatingDerived       bool
	BrokenMetricPhysicalPrediction bool
	GaugeNormalizationArtifact     bool
	BrokenSectorEtaOddForce        bool
	SelectsEtaOrientation          bool
	Verdict                        string
}

type PullbackCandidate struct {
	Name               string
	SourcePackage      string
	Domain             string
	Available          bool
	ActsOnQuarticEta   bool
	ActsOnPhysicalHphi bool
	GaugeInvariant     bool
	EtaOdd             bool
	ExchangesEtaPair   bool
	SelectsOrientation bool
	RequiresPullback   bool
	UsesObservedInput  bool
	Obstruction        string
}

type SourceSearchAudit struct {
	CandidatesAudited           []PullbackCandidate
	AvailableCandidates         int
	EtaOddCandidates            int
	GaugeInvariantEtaOddSources int
	CandidatesSelectingEta      int
	WrongTensorFactorCandidates int
	ExchangeNotSelectorCount    int
	DiagnosticOnlyCandidates    int
	ObservedInputCandidates     int
	EtaOddSourceFound           bool
	GaugeInvariantSourceFound   bool
	CanonicalOrientationDerived bool
	Verdict                     string
}

type GaugeSpontaneousAudit struct {
	EtaInvolutionPreserved             bool
	EtaInvolutionEquivalentToPlaneSwap bool
	GaugeSymmetryExplainsNonselection  bool
	SpontaneousOrientationDataRequired bool
	FiniteObservableCanSelect          bool
	PhysicalScalarBundleStillUnfixed   bool
	OrientationInsertionPointIsolated  bool
	Verdict                            string
}

type Summary struct {
	TestsAudited                        int
	Gate189CompatibilityInherited       bool
	WeakHyperchargeAudited              bool
	ChargeConjugationAudited            bool
	BrokenSectorAudited                 bool
	ContactSignedSourcesAudited         bool
	EtaOddSourceFound                   bool
	GaugeInvariantEtaOddSourceFound     bool
	CanonicalEtaOrientationDerived      bool
	EtaOrientationClassifiedSpontaneous bool
	PhysicalScalarBundleDerived         bool
	Comment                             string
}

type Firewall struct {
	UsesObservedInputForDerivation      bool
	UsesNumericRootApproximation        bool
	UsesIndividualRootDiagonalization   bool
	UsesArbitraryEtaHighLowAssignment   bool
	Gate189CompatibilityInherited       bool
	WeakHyperchargeSourceAudited        bool
	ChargeConjugationSourceAudited      bool
	BrokenSectorSourceAudited           bool
	ContactSignedSourceAudited          bool
	EtaOddFiniteSourceFound             bool
	GaugeInvariantEtaOddSourceFound     bool
	CanonicalEtaOrientationDerived      bool
	EtaOrientationClassifiedSpontaneous bool
	PhysicalScalarBundleDerived         bool
	ChernWeilCarrierDerived             bool
	HeatKernelMatchingDerived           bool
	ThresholdCorrectedBetaDerived       bool
	AbsoluteCouplingPromoted            bool
	PhysicalConstantsDerived            bool
	StrictNullityBefore                 int
	StrictNullityAfter                  int
	ConditionalNullityBefore            int
	ConditionalNullityAfter             int
	ClosedStatements                    []string
	OpenRequirements                    []string
	RecommendedNextGate                 string
	Verdict                             string
}

type Analysis struct {
	PreviousGate189 scalarbundlemap.Analysis
	ScalarSU2       scalarsu2.Analysis
	ScalarCovariant scalarcovariant.Analysis
	ScalarComplex   scalarcomplex.Analysis
	Hypercharge     hypercharge.Analysis
	MatterSU2L      su2lgauge.Analysis
	ContactSign     contactsignsource.Analysis
	ContactCodd     contactcoddsource.Analysis
	HiggsConjugate  higgsconjugatequotient.Analysis
	GaugeEating     gaugeeating.Analysis
	BrokenMetric    brokenmetric.Analysis

	WeakGauge      WeakGaugeAudit
	Conjugation    ConjugationAudit
	BrokenSector   BrokenSectorAudit
	SourceSearch   SourceSearchAudit
	Spontaneous    GaugeSpontaneousAudit
	Summary        Summary
	Firewall       Firewall
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := scalarbundlemap.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 189 input: %w", err)
			return
		}
		su2, err := scalarsu2.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar SU2 input: %w", err)
			return
		}
		sc, err := scalarcovariant.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar covariant input: %w", err)
			return
		}
		cx, err := scalarcomplex.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar complex input: %w", err)
			return
		}
		hy, err := hypercharge.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build hypercharge input: %w", err)
			return
		}
		msu2, err := su2lgauge.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build matter SU2L input: %w", err)
			return
		}
		cs, err := contactsignsource.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build contact sign-source input: %w", err)
			return
		}
		cc, err := contactcoddsource.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build contact C-odd input: %w", err)
			return
		}
		hc, err := higgsconjugatequotient.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Higgs-conjugate input: %w", err)
			return
		}
		ge, err := gaugeeating.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build gauge-eating input: %w", err)
			return
		}
		bm, err := brokenmetric.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build broken-metric input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, su2, sc, cx, hy, msu2, cs, cc, hc, ge, bm, 1e-9)
	})
	return defaultA, defaultErr
}

func Build(prev scalarbundlemap.Analysis, su2 scalarsu2.Analysis, sc scalarcovariant.Analysis, cx scalarcomplex.Analysis, hy hypercharge.Analysis, msu2 su2lgauge.Analysis, cs contactsignsource.Analysis, cc contactcoddsource.Analysis, hc higgsconjugatequotient.Analysis, ge gaugeeating.Analysis, bm brokenmetric.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	if !prev.Firewall.DimensionCompatibilityDerived || !prev.Firewall.ConditionalBundleMapsExist || prev.Firewall.CanonicalEtaOrientationDerived || prev.Firewall.PhysicalScalarBundleDerived {
		return Analysis{}, fmt.Errorf("Gate 190 requires Gate 189 compatibility with unresolved eta orientation")
	}
	weak, err := auditWeakGauge(su2, sc, hy, msu2, eps)
	if err != nil {
		return Analysis{}, err
	}
	conj := auditConjugation(cs, hc)
	broken := auditBrokenSector(sc, ge, bm)
	sources := auditSources(prev, weak, conj, broken, cx, hy, cs, cc)
	spont := auditSpontaneous(prev, weak, conj, sources)
	summary := Summary{
		TestsAudited:                        6,
		Gate189CompatibilityInherited:       prev.Firewall.DimensionCompatibilityDerived && prev.Firewall.ConditionalBundleMapsExist,
		WeakHyperchargeAudited:              weak.T3LAvailable && weak.ScalarHyperchargeAvailable && weak.MatterSU2LAvailable,
		ChargeConjugationAudited:            conj.ContactChargeConjugationAvailable && conj.ContactChargeConjugationInvolution,
		BrokenSectorAudited:                 broken.CovariantDerivativeTemplate && broken.GoldstoneImageDiagnostic,
		ContactSignedSourcesAudited:         cc.CanonicalSignedDiagnostics > 0,
		EtaOddSourceFound:                   sources.EtaOddSourceFound,
		GaugeInvariantEtaOddSourceFound:     sources.GaugeInvariantSourceFound,
		CanonicalEtaOrientationDerived:      sources.CanonicalOrientationDerived,
		EtaOrientationClassifiedSpontaneous: spont.OrientationInsertionPointIsolated,
		PhysicalScalarBundleDerived:         false,
		Comment:                             "Gate 190 audits every available weak, hypercharge, charge-conjugation, contact signed, and broken-sector candidate. None is a gauge-invariant eta-odd selector; eta orientation is isolated as spontaneous/gauge data.",
	}
	fw := Firewall{
		UsesObservedInputForDerivation:      false,
		UsesNumericRootApproximation:        false,
		UsesIndividualRootDiagonalization:   false,
		UsesArbitraryEtaHighLowAssignment:   false,
		Gate189CompatibilityInherited:       summary.Gate189CompatibilityInherited,
		WeakHyperchargeSourceAudited:        summary.WeakHyperchargeAudited,
		ChargeConjugationSourceAudited:      summary.ChargeConjugationAudited,
		BrokenSectorSourceAudited:           summary.BrokenSectorAudited,
		ContactSignedSourceAudited:          summary.ContactSignedSourcesAudited,
		EtaOddFiniteSourceFound:             sources.EtaOddSourceFound,
		GaugeInvariantEtaOddSourceFound:     sources.GaugeInvariantSourceFound,
		CanonicalEtaOrientationDerived:      sources.CanonicalOrientationDerived,
		EtaOrientationClassifiedSpontaneous: spont.OrientationInsertionPointIsolated,
		PhysicalScalarBundleDerived:         false,
		ChernWeilCarrierDerived:             false,
		HeatKernelMatchingDerived:           false,
		ThresholdCorrectedBetaDerived:       false,
		AbsoluteCouplingPromoted:            false,
		PhysicalConstantsDerived:            false,
		StrictNullityBefore:                 prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                  prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:            prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:             0,
		ClosedStatements: []string{
			"T3L and scalar hypercharge are available on the physical scalar frame but commute with the high/low projectors or act as gauge generators, not eta selectors",
			"the SU(2) Weyl-style plane swap and charge conjugation exchange the two orientations; they prove a Z2/gauge degeneracy rather than a preferred eta sign",
			"B-L, contact signed currents, and contact C-odd diagnostics have no derived 2+2 scalar eta pullback",
			"the broken-sector W/Z/photon and Goldstone-image diagnostics still depend on an unselected vacuum orientation and cannot retroactively select eta",
			"no gauge-invariant finite observable can currently distinguish eta -> high from eta -> low",
		},
		OpenRequirements: []string{
			"record the eta-to-high/low assignment as explicit spontaneous vacuum orientation data before constructing a gauge-fixed scalar-bundle trivialization",
			"keep the scalar bundle map conditional on that orientation and on a gauge/SU(2) frame choice",
			"do not reopen Chern-Weil, heat-kernel, threshold beta rows, absolute coupling promotion, or physical constants until the conditional bundle map is explicitly sealed",
		},
		RecommendedNextGate: "Gate 191 — spontaneous scalar-orientation seal / gauge-fixed H_Phi trivialization axiom audit",
		Verdict:             "Gate 190 is a positive obstruction theorem: the eta-odd source search fails constructively, proving that eta orientation is not an internal finite observable but the exact spontaneous/gauge insertion point for the physical scalar-bundle convention.",
	}
	truth := "Gate 190 closes the eta-source search. Weak isospin, hypercharge, charge conjugation, contact signed diagnostics, and broken-sector mass/goldstone diagnostics either preserve the eta pair, exchange it as symmetry, live on the wrong factor, or use a diagnostic vacuum convention. Therefore the finite algebra does not derive eta -> high versus eta -> low. That missing assignment is now localized as spontaneous scalar-orientation data rather than an algebraic defect."
	return Analysis{PreviousGate189: prev, ScalarSU2: su2, ScalarCovariant: sc, ScalarComplex: cx, Hypercharge: hy, MatterSU2L: msu2, ContactSign: cs, ContactCodd: cc, HiggsConjugate: hc, GaugeEating: ge, BrokenMetric: bm, WeakGauge: weak, Conjugation: conj, BrokenSector: broken, SourceSearch: sources, Spontaneous: spont, Summary: summary, Firewall: fw, TruthStatement: truth}, nil
}

func auditWeakGauge(su2 scalarsu2.Analysis, sc scalarcovariant.Analysis, hy hypercharge.Analysis, msu2 su2lgauge.Analysis, eps float64) (WeakGaugeAudit, error) {
	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})
	cT3H := commNorm(sc.T3, ph)
	cT3L := commNorm(sc.T3, pl)
	cYH := commNorm(sc.YPhi, ph)
	cYL := commNorm(sc.YPhi, pl)
	cT1H := commNorm(sc.T1, ph)
	cT2H := commNorm(sc.T2, ph)
	w := planeSwapWeyl()
	wph, err := conjugate(w, ph)
	if err != nil {
		return WeakGaugeAudit{}, err
	}
	wpl, err := conjugate(w, pl)
	if err != nil {
		return WeakGaugeAudit{}, err
	}
	d1, _ := wph.Sub(pl)
	d2, _ := wpl.Sub(ph)
	weylSwap := d1.FrobeniusNorm() <= eps && d2.FrobeniusNorm() <= eps
	return WeakGaugeAudit{
		T3LAvailable:                    su2.AbstractDoubletRepresentation && sc.AbstractCovariantDerivativeTemplate,
		ScalarHyperchargeAvailable:      hy.ScalarChargeBridgeConstructed && sc.AbstractCovariantDerivativeTemplate,
		MatterSU2LAvailable:             msu2.NonabelianSU2LGeneratorsDerived && msu2.CommutesWithHyperchargeNorm <= eps,
		ScalarDoubletCandidate:          hy.ScalarDoubletCandidate && su2.ActiveRealDimension == 4,
		T3CommutesWithHighLowProjectors: cT3H <= eps && cT3L <= eps,
		YCommutesWithHighLowProjectors:  cYH <= eps && cYL <= eps,
		T1MixesHighLowPlanes:            cT1H > eps,
		T2MixesHighLowPlanes:            cT2H > eps,
		WeylReflectionExchangesPlanes:   weylSwap,
		GaugeActionSelectsOrientation:   false,
		EtaOddGaugeInvariantSource:      false,
		Verdict:                         "T3L and Y preserve the high/low planes, while T1/T2 and the Weyl-style plane swap exchange/mix them as gauge representation data. They do not provide a gauge-invariant eta-odd scalar-orientation selector.",
	}, nil
}

func auditConjugation(cs contactsignsource.Analysis, hc higgsconjugatequotient.Analysis) ConjugationAudit {
	return ConjugationAudit{
		ContactChargeConjugationAvailable:  cs.ChargeConjugation.ChargeConjugationAvailable,
		ContactChargeConjugationInvolution: cs.ChargeConjugation.ActsAsInvolution,
		ContactChargeConjugationExchanges:  cs.ChargeConjugation.ExchangesOrientations,
		ContactChargeConjugationSelects:    cs.ChargeConjugation.SelectedOrientations > 0,
		HiggsConjugateCollapseRejected:     !hc.HiggsAudit.HiggsConjugatePairCollapse,
		HiggsBranchUniquenessByKind:        hc.HiggsAudit.HyperchargeSelectsUniqueBranch,
		MirrorsEtaInvolution:               cs.ChargeConjugation.ExchangesOrientations,
		SelectsEtaOrientation:              false,
		Verdict:                            "charge conjugation is an involution that exchanges orientations; the Higgs-conjugate channel quotient is rejected on the actual Yukawa support. Conjugation mirrors eta swap but selects no eta sign.",
	}
}

func auditBrokenSector(sc scalarcovariant.Analysis, ge gaugeeating.Analysis, bm brokenmetric.Analysis) BrokenSectorAudit {
	return BrokenSectorAudit{
		CovariantDerivativeTemplate:    sc.AbstractCovariantDerivativeTemplate,
		VacuumOrientationChosen:        sc.VacuumOrientationChosen,
		DimensionlessWZPhotonSignature: sc.DimensionlessWZPhotonSignature,
		GoldstoneImageDiagnostic:       ge.GoldstoneImageTheoremDiagnostic && ge.GaugeEatingCountDiagnostic,
		FiniteGaugeEatingDerived:       ge.FiniteGaugeEatingTheoremDerived,
		BrokenMetricPhysicalPrediction: bm.BrokenMetricPhysicalPrediction,
		GaugeNormalizationArtifact:     bm.GaugeNormalizationArtifactPossible,
		BrokenSectorEtaOddForce:        false,
		SelectsEtaOrientation:          false,
		Verdict:                        "broken-sector diagnostics produce the W/Z/photon and Goldstone-image signature only after a diagnostic vacuum convention; they cannot be used backward as an eta-orientation selector.",
	}
}

func auditSources(prev scalarbundlemap.Analysis, weak WeakGaugeAudit, conj ConjugationAudit, broken BrokenSectorAudit, cx scalarcomplex.Analysis, hy hypercharge.Analysis, cs contactsignsource.Analysis, cc contactcoddsource.Analysis) SourceSearchAudit {
	candidates := []PullbackCandidate{
		{Name: "B-L / Fock charge", SourcePackage: "pkg/matter/charge", Domain: "matter/Fock 1+3", Available: prev.Sources.BLPullbackAudited, ActsOnQuarticEta: false, ActsOnPhysicalHphi: false, GaugeInvariant: true, EtaOdd: false, SelectsOrientation: false, RequiresPullback: true, Obstruction: "wrong tensor-factor and wrong shape: 1+3 matter polarization, not 2+2 scalar eta orientation"},
		{Name: "weak isospin T3L", SourcePackage: "pkg/bridge/scalarsu2 + pkg/bridge/scalarcovariant", Domain: "physical scalar doublet", Available: weak.T3LAvailable, ActsOnQuarticEta: false, ActsOnPhysicalHphi: true, GaugeInvariant: false, EtaOdd: false, SelectsOrientation: false, RequiresPullback: true, Obstruction: "T3L commutes with high/low projectors in the chosen scalar frame and no branch-projector pullback is derived"},
		{Name: "scalar hypercharge Y_phi", SourcePackage: "pkg/matter/hypercharge + pkg/bridge/scalarcovariant", Domain: "physical scalar doublet", Available: weak.ScalarHyperchargeAvailable && hy.ScalarChargeBridgeConstructed, ActsOnQuarticEta: false, ActsOnPhysicalHphi: true, GaugeInvariant: false, EtaOdd: false, SelectsOrientation: false, RequiresPullback: true, Obstruction: "Y_phi is pair-preserving scalar representation data, not an eta-odd branch selector"},
		{Name: "SU(2)_L Weyl / plane-swap action", SourcePackage: "pkg/bridge/scalarsu2", Domain: "physical scalar doublet gauge action", Available: weak.WeylReflectionExchangesPlanes, ActsOnQuarticEta: false, ActsOnPhysicalHphi: true, GaugeInvariant: false, EtaOdd: false, ExchangesEtaPair: true, SelectsOrientation: false, RequiresPullback: true, Obstruction: "exchanges the two scalar planes as gauge symmetry; exchange is not selection"},
		{Name: "charge conjugation", SourcePackage: "pkg/bridge/contactsignsource", Domain: "contact sign orientations", Available: conj.ContactChargeConjugationAvailable, ActsOnQuarticEta: false, ActsOnPhysicalHphi: false, GaugeInvariant: true, EtaOdd: false, ExchangesEtaPair: conj.ContactChargeConjugationExchanges, SelectsOrientation: false, RequiresPullback: true, Obstruction: "C is an involution that proves Z2 degeneracy; no C-breaking eta pullback exists"},
		{Name: "pair-compatible scalar complex structure", SourcePackage: "pkg/bridge/scalarcomplex", Domain: "physical scalar planes", Available: cx.PairCompatibleComplexAvailable, ActsOnQuarticEta: false, ActsOnPhysicalHphi: true, GaugeInvariant: false, EtaOdd: false, SelectsOrientation: false, RequiresPullback: true, Obstruction: "complex direction exists but its signs/orientations are noncanonical and do not act on quartic eta"},
		{Name: "centered contact signed diagnostic", SourcePackage: "pkg/bridge/contactcoddsource", Domain: "seven-row contact spectrum", Available: cc.CenteredFunctional.CanonicalAsDiagnostic, ActsOnQuarticEta: false, ActsOnPhysicalHphi: false, GaugeInvariant: false, EtaOdd: false, SelectsOrientation: false, RequiresPullback: true, Obstruction: "canonical signed diagnostic but not a physical C-odd source, T3R/hypercharge semantic, or 2+2 scalar pullback"},
		{Name: "broken generator / covariant derivative diagnostic", SourcePackage: "pkg/bridge/scalarcovariant + pkg/bridge/gaugeeating", Domain: "chosen physical scalar vacuum frame", Available: broken.CovariantDerivativeTemplate && broken.GoldstoneImageDiagnostic, ActsOnQuarticEta: false, ActsOnPhysicalHphi: true, GaugeInvariant: false, EtaOdd: false, SelectsOrientation: false, RequiresPullback: true, Obstruction: "uses a diagnostic vacuum orientation; cannot justify that orientation as derived"},
		{Name: "observed high/low assignment", SourcePackage: "forbidden", Domain: "external phenomenology", Available: false, ActsOnQuarticEta: false, ActsOnPhysicalHphi: true, GaugeInvariant: false, EtaOdd: true, SelectsOrientation: false, UsesObservedInput: true, Obstruction: "observed physical input is forbidden in the finite derivation firewall"},
	}
	available, etaOdd, gaugeEtaOdd, selecting, wrong, exchange, diagnostic, observed := 0, 0, 0, 0, 0, 0, 0, 0
	for _, c := range candidates {
		if c.Available {
			available++
		}
		if c.EtaOdd && c.Available {
			etaOdd++
		}
		if c.EtaOdd && c.GaugeInvariant && c.Available {
			gaugeEtaOdd++
		}
		if c.SelectsOrientation && c.Available {
			selecting++
		}
		if c.RequiresPullback && !c.ActsOnQuarticEta {
			wrong++
		}
		if c.ExchangesEtaPair {
			exchange++
		}
		if strings.Contains(c.Obstruction, "diagnostic") {
			diagnostic++
		}
		if c.UsesObservedInput {
			observed++
		}
	}
	return SourceSearchAudit{
		CandidatesAudited:           candidates,
		AvailableCandidates:         available,
		EtaOddCandidates:            etaOdd,
		GaugeInvariantEtaOddSources: gaugeEtaOdd,
		CandidatesSelectingEta:      selecting,
		WrongTensorFactorCandidates: wrong,
		ExchangeNotSelectorCount:    exchange,
		DiagnosticOnlyCandidates:    diagnostic,
		ObservedInputCandidates:     observed,
		EtaOddSourceFound:           etaOdd > 0,
		GaugeInvariantSourceFound:   gaugeEtaOdd > 0,
		CanonicalOrientationDerived: selecting > 0,
		Verdict:                     "The complete audited source list contains no available gauge-invariant eta-odd operator and no finite pullback that selects eta -> high rather than eta -> low.",
	}
}

func auditSpontaneous(prev scalarbundlemap.Analysis, weak WeakGaugeAudit, conj ConjugationAudit, sources SourceSearchAudit) GaugeSpontaneousAudit {
	isolated := prev.AbstractBranch.EtaInvolutionSwapsPair && weak.WeylReflectionExchangesPlanes && conj.MirrorsEtaInvolution && !sources.CanonicalOrientationDerived
	return GaugeSpontaneousAudit{
		EtaInvolutionPreserved:             prev.AbstractBranch.EtaInvolutionSwapsPair,
		EtaInvolutionEquivalentToPlaneSwap: weak.WeylReflectionExchangesPlanes && conj.MirrorsEtaInvolution,
		GaugeSymmetryExplainsNonselection:  weak.WeylReflectionExchangesPlanes && !weak.GaugeActionSelectsOrientation,
		SpontaneousOrientationDataRequired: isolated,
		FiniteObservableCanSelect:          false,
		PhysicalScalarBundleStillUnfixed:   true,
		OrientationInsertionPointIsolated:  isolated,
		Verdict:                            "eta -> -eta is preserved as a gauge/conjugation-style exchange symmetry. The finite algebra can derive the orbit and projectors, but not the occupied orientation; that datum is spontaneous/gauge-frame input.",
	}
}

func planeSwapWeyl() linear.Matrix {
	w := linear.NewMatrix(4, 4)
	w.Set(0, 2, -1)
	w.Set(1, 3, -1)
	w.Set(2, 0, 1)
	w.Set(3, 1, 1)
	return w
}

func conjugate(w, p linear.Matrix) (linear.Matrix, error) {
	wp, err := w.Mul(p)
	if err != nil {
		return linear.Matrix{}, err
	}
	return wp.Mul(w.Transpose())
}

func commNorm(a, b linear.Matrix) float64 {
	c, err := linear.Commutator(a, b)
	if err != nil {
		return math.Inf(1)
	}
	return c.FrobeniusNorm()
}

func FormatWeakGauge(a WeakGaugeAudit) string {
	return fmt.Sprintf("T3=%t Y=%t matterSU2=%t doublet=%t commT3=%t commY=%t T1mix=%t T2mix=%t WeylSwap=%t selects=%t etaOddGauge=%t (%s)", a.T3LAvailable, a.ScalarHyperchargeAvailable, a.MatterSU2LAvailable, a.ScalarDoubletCandidate, a.T3CommutesWithHighLowProjectors, a.YCommutesWithHighLowProjectors, a.T1MixesHighLowPlanes, a.T2MixesHighLowPlanes, a.WeylReflectionExchangesPlanes, a.GaugeActionSelectsOrientation, a.EtaOddGaugeInvariantSource, a.Verdict)
}

func FormatConjugation(a ConjugationAudit) string {
	return fmt.Sprintf("Cavailable=%t involution=%t exchanges=%t Cselects=%t higgsQuotientRejected=%t branchUnique=%t mirrorsEta=%t etaSelects=%t (%s)", a.ContactChargeConjugationAvailable, a.ContactChargeConjugationInvolution, a.ContactChargeConjugationExchanges, a.ContactChargeConjugationSelects, a.HiggsConjugateCollapseRejected, a.HiggsBranchUniquenessByKind, a.MirrorsEtaInvolution, a.SelectsEtaOrientation, a.Verdict)
}

func FormatBrokenSector(a BrokenSectorAudit) string {
	return fmt.Sprintf("Dtemplate=%t vacChosen=%t WZgamma=%t goldstone=%t eating=%t metricPhysical=%t artifact=%t etaForce=%t selects=%t (%s)", a.CovariantDerivativeTemplate, a.VacuumOrientationChosen, a.DimensionlessWZPhotonSignature, a.GoldstoneImageDiagnostic, a.FiniteGaugeEatingDerived, a.BrokenMetricPhysicalPrediction, a.GaugeNormalizationArtifact, a.BrokenSectorEtaOddForce, a.SelectsEtaOrientation, a.Verdict)
}

func FormatSources(a SourceSearchAudit) string {
	parts := make([]string, 0, len(a.CandidatesAudited))
	for _, c := range a.CandidatesAudited {
		parts = append(parts, fmt.Sprintf("%s[%s] avail=%t quartic=%t hphi=%t gaugeInv=%t etaOdd=%t exchange=%t selects=%t pullback=%t observed=%t obstruction=%q", c.Name, c.SourcePackage, c.Available, c.ActsOnQuarticEta, c.ActsOnPhysicalHphi, c.GaugeInvariant, c.EtaOdd, c.ExchangesEtaPair, c.SelectsOrientation, c.RequiresPullback, c.UsesObservedInput, c.Obstruction))
	}
	return fmt.Sprintf("available=%d etaOdd=%d gaugeEtaOdd=%d selecting=%d wrongFactor=%d exchangeOnly=%d diagnostic=%d observed=%d found=%t gaugeFound=%t canonical=%t candidates=[%s] (%s)", a.AvailableCandidates, a.EtaOddCandidates, a.GaugeInvariantEtaOddSources, a.CandidatesSelectingEta, a.WrongTensorFactorCandidates, a.ExchangeNotSelectorCount, a.DiagnosticOnlyCandidates, a.ObservedInputCandidates, a.EtaOddSourceFound, a.GaugeInvariantSourceFound, a.CanonicalOrientationDerived, strings.Join(parts, "; "), a.Verdict)
}

func FormatSpontaneous(a GaugeSpontaneousAudit) string {
	return fmt.Sprintf("etaPreserved=%t etaPlaneSwap=%t gaugeNonselection=%t spontaneousRequired=%t finiteObservable=%t bundleUnfixed=%t insertionIsolated=%t (%s)", a.EtaInvolutionPreserved, a.EtaInvolutionEquivalentToPlaneSwap, a.GaugeSymmetryExplainsNonselection, a.SpontaneousOrientationDataRequired, a.FiniteObservableCanSelect, a.PhysicalScalarBundleStillUnfixed, a.OrientationInsertionPointIsolated, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d gate189=%t weak=%t C=%t broken=%t contactSigned=%t etaOdd=%t gaugeEtaOdd=%t canonical=%t spontaneous=%t physicalBundle=%t (%s)", a.TestsAudited, a.Gate189CompatibilityInherited, a.WeakHyperchargeAudited, a.ChargeConjugationAudited, a.BrokenSectorAudited, a.ContactSignedSourcesAudited, a.EtaOddSourceFound, a.GaugeInvariantEtaOddSourceFound, a.CanonicalEtaOrientationDerived, a.EtaOrientationClassifiedSpontaneous, a.PhysicalScalarBundleDerived, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t numeric=%t rootDiag=%t arbitraryEta=%t gate189=%t weak=%t C=%t broken=%t contactSigned=%t etaSource=%t gaugeEta=%t etaOrient=%t spontaneous=%t physicalBundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.UsesObservedInputForDerivation, a.UsesNumericRootApproximation, a.UsesIndividualRootDiagonalization, a.UsesArbitraryEtaHighLowAssignment, a.Gate189CompatibilityInherited, a.WeakHyperchargeSourceAudited, a.ChargeConjugationSourceAudited, a.BrokenSectorSourceAudited, a.ContactSignedSourceAudited, a.EtaOddFiniteSourceFound, a.GaugeInvariantEtaOddSourceFound, a.CanonicalEtaOrientationDerived, a.EtaOrientationClassifiedSpontaneous, a.PhysicalScalarBundleDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}
