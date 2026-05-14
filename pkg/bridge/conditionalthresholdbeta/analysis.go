// Package conditionalthresholdbeta implements Gate 198: conditional threshold
// beta-row activation / decoupling scheme firewall audit.
//
// Gate 197 supplied formal fermion threshold symbols only after two quarantined
// inputs: the empirical Yukawa texture seal and the empirical VEV scale seal.
// Gate 198 asks what can now be said about one-loop threshold rows.  It builds
// the exact rational fermion beta-row ledger as conditional continuum/RG
// scaffolding, introduces an explicit decoupling-scheme seal for sharp
// step-function matching, and refuses to evaluate a physical RG flow or infer
// W/Z thresholds, boundary couplings, or continuum normalization.
package conditionalthresholdbeta

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/electroweakvevseal"
	"github.com/bagherbal/asha-engine/pkg/bridge/yukawaamplitudeseal"
)

type Rational struct {
	Num int64
	Den int64
}

func NewRational(num, den int64) Rational {
	if den == 0 {
		panic("zero denominator")
	}
	if den < 0 {
		num = -num
		den = -den
	}
	g := gcd(abs(num), den)
	return Rational{Num: num / g, Den: den / g}
}

func (r Rational) Add(s Rational) Rational { return NewRational(r.Num*s.Den+s.Num*r.Den, r.Den*s.Den) }
func (r Rational) Neg() Rational           { return Rational{Num: -r.Num, Den: r.Den} }
func (r Rational) Equal(s Rational) bool   { return r.Num == s.Num && r.Den == s.Den }
func (r Rational) String() string {
	if r.Den == 1 {
		return fmt.Sprintf("%d", r.Num)
	}
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

type BetaRow struct {
	Sector                            yukawaamplitudeseal.YukawaKind
	Generation                        int
	MassSymbol                        string
	ActivePredicate                   string
	U1YContribution                   Rational
	SU2LContribution                  Rational
	SU3CContribution                  Rational
	U1YSubtractionBelowThreshold      Rational
	SU2LSubtractionBelowThreshold     Rational
	SU3CSubtractionBelowThreshold     Rational
	ExactRational                     bool
	RequiresEmpiricalTextureSeal      bool
	RequiresVEVSeal                   bool
	RequiresContinuumDecouplingScheme bool
	UsesObservedMass                  bool
	NativeFiniteThresholdRowDerived   bool
	BrokenPhaseHalfDoubletBookkeeping bool
	Verdict                           string
}

type SectorRow struct {
	Sector                      yukawaamplitudeseal.YukawaKind
	Generations                 int
	U1YContribution             Rational
	SU2LContribution            Rational
	SU3CContribution            Rational
	SubtractionVector           string
	HyperchargeSource           string
	NonAbelianSource            string
	SplitRequiresBrokenPhaseVEV bool
	ExactRational               bool
	Verdict                     string
}

type BetaShiftLedger struct {
	Rows                               []BetaRow
	SectorRows                         []SectorRow
	SectorCount                        int
	Generations                        int
	ThresholdRows                      int
	AllRowsExactRational               bool
	AllRowsConditionalOnTexture        bool
	AllRowsConditionalOnVEV            bool
	AllRowsConditionalOnScheme         bool
	AnyObservedMassUsed                bool
	AnyNativeFiniteThresholdRowDerived bool
	ReconstructsFullFermionU1Y         bool
	ReconstructsFullFermionSU2L        bool
	ReconstructsFullFermionSU3C        bool
	FermionContributionU1Y             Rational
	FermionContributionSU2L            Rational
	FermionContributionSU3C            Rational
	OneLoopContinuumFormulaAssumed     bool
	NativeFiniteBetaTheoremDerived     bool
	Verdict                            string
}

type ContinuumDecouplingSchemeSeal struct {
	Name                             string
	AxiomID                          string
	ConditionalStatus                string
	SchemeFamily                     string
	SharpStepPredicate               string
	ExplicitBoundaryConvention       bool
	Quarantined                      bool
	RequiredForThresholdRows         bool
	DerivedFromFiniteAlgebra         bool
	SharpStepDerivedNatively         bool
	SmoothRegulatorSearched          bool
	SmoothRegulatorDerived           bool
	SchemeSelectedNatively           bool
	MSbarDerived                     bool
	MOMDerived                       bool
	TreeLevelContinuityEnforced      bool
	FiniteMatchingCorrectionsDerived bool
	DownstreamMustDeclareSeal        bool
	Verdict                          string
}

type LowEnergyDomainAudit struct {
	FermionThresholdSymbolsAvailable bool
	GaugeBosonThresholdsAvailable    bool
	WZThresholdsBlocked              bool
	DeepInfraredFlowDefined          bool
	RunToMZAllowed                   bool
	DomainRestriction                string
	BlockReason                      string
	Verdict                          string
}

type PiecewiseRGAssembly struct {
	BetaFunctionExpression           string
	RunningExpression                string
	ThresholdLogKernel               string
	TreeLevelMatchingCondition       string
	FiniteMatchingCorrectionTerm     string
	PiecewiseSymbolicTreeBuilt       bool
	EvaluatedNumerically             bool
	ThresholdOrderingKnown           bool
	BoundaryScaleDerived             bool
	BoundaryCouplingDerived          bool
	FiniteMatchingCorrectionsDerived bool
	EnforcesTreeLevelContinuity      bool
	SchemeDependentCorrectionsSealed bool
	Verdict                          string
}

type RGFirewallAudit struct {
	Gate197VEVSealInherited               bool
	FormalMassThresholdsAvailable         bool
	ContinuumDecouplingSchemeSealInserted bool
	ConditionalBetaRowsActivated          bool
	FiniteThresholdRowsDerived            bool
	NativeSmoothRegulatorDerived          bool
	SchemeConventionDerivedFromFinite     bool
	PiecewiseRGTreeConstructed            bool
	PhysicalRGFlowEvaluated               bool
	NumericalThresholdsKnown              bool
	ThresholdOrderingKnown                bool
	LowEnergyWZDomainKnown                bool
	GaugeBosonThresholdsDerived           bool
	AbsoluteBoundaryScaleDerived          bool
	AbsoluteBoundaryCouplingDerived       bool
	GaugeCouplingsDerived                 bool
	TopologicalEightPiSquaredImported     bool
	FiniteToContinuumScaleDerived         bool
	ObservedInputsImported                bool
	StrictNullityBefore                   int
	StrictNullityAfter                    int
	ConditionalSchemeNullityBefore        int
	ConditionalSchemeNullityAfter         int
	ConditionalBetaRowNullityBefore       int
	ConditionalBetaRowNullityAfter        int
	ConditionalRGEvaluationNullityBefore  int
	ConditionalRGEvaluationNullityAfter   int
	OpenRequirements                      []string
	RecommendedNextGate                   string
	Verdict                               string
}

type Summary struct {
	TestsAudited                      int
	Gate197Inherited                  bool
	ExactConditionalBetaRowsBuilt     bool
	DecouplingSchemeSealRecorded      bool
	TreeLevelContinuityAnswerRecorded bool
	PiecewiseRGScaffoldBuilt          bool
	WZLowEnergyFirewallPreserved      bool
	AbsoluteCouplingFirewallPreserved bool
	Comment                           string
}

type Analysis struct {
	PreviousGate197 electroweakvevseal.Analysis
	Ledger          BetaShiftLedger
	SchemeSeal      ContinuumDecouplingSchemeSeal
	Domain          LowEnergyDomainAudit
	Piecewise       PiecewiseRGAssembly
	Firewall        RGFirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := electroweakvevseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 197 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev electroweakvevseal.Analysis) (Analysis, error) {
	if !prev.Summary.EmpiricalVEVSealRecorded || !prev.Summary.FormalMassThresholdsActivated || prev.Firewall.ThresholdBetaRowsDerived || prev.Firewall.ThresholdCorrectedRGFlowDerived {
		return Analysis{}, fmt.Errorf("Gate 198 requires Gate 197 formal threshold symbols while beta rows/RG flow remain sealed")
	}
	if prev.Firewall.NumericalMassThresholdsAvailable || prev.Firewall.GaugeBosonThresholdsDerived || prev.Firewall.AbsoluteBoundaryCouplingDerived || prev.Firewall.FiniteToContinuumScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 198 refuses leaked numerical thresholds, W/Z thresholds, couplings, or continuum normalization")
	}

	ledger := buildBetaLedger(prev)
	scheme := buildSchemeSeal(ledger)
	domain := auditLowEnergyDomain(prev)
	piecewise := buildPiecewiseRG(ledger, scheme, domain)
	fw := auditFirewall(prev, ledger, scheme, domain, piecewise)
	summary := Summary{
		TestsAudited:                      6,
		Gate197Inherited:                  fw.Gate197VEVSealInherited && fw.FormalMassThresholdsAvailable,
		ExactConditionalBetaRowsBuilt:     ledger.ThresholdRows == 12 && ledger.AllRowsExactRational && ledger.ReconstructsFullFermionU1Y && ledger.ReconstructsFullFermionSU2L && ledger.ReconstructsFullFermionSU3C,
		DecouplingSchemeSealRecorded:      scheme.ExplicitBoundaryConvention && scheme.Quarantined && scheme.RequiredForThresholdRows && !scheme.DerivedFromFiniteAlgebra,
		TreeLevelContinuityAnswerRecorded: piecewise.EnforcesTreeLevelContinuity && piecewise.SchemeDependentCorrectionsSealed && !piecewise.FiniteMatchingCorrectionsDerived,
		PiecewiseRGScaffoldBuilt:          piecewise.PiecewiseSymbolicTreeBuilt && !piecewise.EvaluatedNumerically,
		WZLowEnergyFirewallPreserved:      domain.WZThresholdsBlocked && !domain.RunToMZAllowed && !domain.DeepInfraredFlowDefined,
		AbsoluteCouplingFirewallPreserved: !fw.AbsoluteBoundaryScaleDerived && !fw.AbsoluteBoundaryCouplingDerived && !fw.GaugeCouplingsDerived && fw.StrictNullityBefore == fw.StrictNullityAfter,
		Comment:                           "Gate 198 builds exact rational fermion beta-row bookkeeping under the VEV, texture, and continuum-decoupling scheme seals. Matching is tree-level continuous by convention; finite matching corrections, smooth regulators, W/Z thresholds, absolute couplings, and physical RG evaluation remain sealed.",
	}
	truth := "Gate 198 activates conditional threshold beta-row scaffolding, not a finite RG theorem. The exact one-loop representation rows for the 12 formal fermion thresholds reconstruct the full fermion contributions (b1,b2,b3)=(4,4,4), but they rely on an explicit continuum decoupling scheme seal and a broken-phase half-doublet bookkeeping convention. The engine enforces tree-level continuity at thresholds while sealing finite matching corrections. No W/Z threshold, numerical threshold ordering, M*, g_*², absolute gauge coupling, 8π² normalization, smooth regulator, or physical running prediction is derived."
	return Analysis{PreviousGate197: prev, Ledger: ledger, SchemeSeal: scheme, Domain: domain, Piecewise: piecewise, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func buildBetaLedger(prev electroweakvevseal.Analysis) BetaShiftLedger {
	sectorRows := []SectorRow{
		sectorRow(yukawaamplitudeseal.YukawaUp, NewRational(17, 30), NewRational(1, 2), NewRational(2, 3), "u_L/u_R hypercharges with color multiplicity", "colored Dirac pair plus half of the broken weak doublet"),
		sectorRow(yukawaamplitudeseal.YukawaDown, NewRational(1, 6), NewRational(1, 2), NewRational(2, 3), "d_L/d_R hypercharges with color multiplicity", "colored Dirac pair plus half of the broken weak doublet"),
		sectorRow(yukawaamplitudeseal.YukawaElectron, NewRational(1, 2), NewRational(1, 6), NewRational(0, 1), "e_L/e_R hypercharges", "leptonic Dirac pair plus half of the broken weak doublet"),
		sectorRow(yukawaamplitudeseal.YukawaNeutrino, NewRational(1, 10), NewRational(1, 6), NewRational(0, 1), "nu_L plus sterile/right-neutrino hypercharge zero", "leptonic Dirac pair plus half of the broken weak doublet"),
	}
	sort.Slice(sectorRows, func(i, j int) bool { return string(sectorRows[i].Sector) < string(sectorRows[j].Sector) })

	rows := make([]BetaRow, 0, len(sectorRows)*3)
	for _, sr := range sectorRows {
		for g := 1; g <= 3; g++ {
			mass := fmt.Sprintf("M_%s,%d", subscript(sr.Sector), g)
			rows = append(rows, BetaRow{
				Sector:                            sr.Sector,
				Generation:                        g,
				MassSymbol:                        mass,
				ActivePredicate:                   fmt.Sprintf("Theta(mu - %s)", mass),
				U1YContribution:                   sr.U1YContribution,
				SU2LContribution:                  sr.SU2LContribution,
				SU3CContribution:                  sr.SU3CContribution,
				U1YSubtractionBelowThreshold:      sr.U1YContribution.Neg(),
				SU2LSubtractionBelowThreshold:     sr.SU2LContribution.Neg(),
				SU3CSubtractionBelowThreshold:     sr.SU3CContribution.Neg(),
				ExactRational:                     true,
				RequiresEmpiricalTextureSeal:      true,
				RequiresVEVSeal:                   prev.Seal.ExplicitBoundaryData,
				RequiresContinuumDecouplingScheme: true,
				UsesObservedMass:                  false,
				NativeFiniteThresholdRowDerived:   false,
				BrokenPhaseHalfDoubletBookkeeping: true,
				Verdict:                           "exact rational one-loop row is available as conditional continuum bookkeeping; below-threshold subtraction is the negative row; no observed mass or native finite threshold row is derived",
			})
		}
	}
	sum1, sum2, sum3 := NewRational(0, 1), NewRational(0, 1), NewRational(0, 1)
	allExact, allTex, allVEV, allScheme := true, true, true, true
	anyObs, anyNative := false, false
	for _, r := range rows {
		sum1 = sum1.Add(r.U1YContribution)
		sum2 = sum2.Add(r.SU2LContribution)
		sum3 = sum3.Add(r.SU3CContribution)
		allExact = allExact && r.ExactRational
		allTex = allTex && r.RequiresEmpiricalTextureSeal
		allVEV = allVEV && r.RequiresVEVSeal
		allScheme = allScheme && r.RequiresContinuumDecouplingScheme
		anyObs = anyObs || r.UsesObservedMass
		anyNative = anyNative || r.NativeFiniteThresholdRowDerived
	}
	return BetaShiftLedger{
		Rows:                               rows,
		SectorRows:                         sectorRows,
		SectorCount:                        len(sectorRows),
		Generations:                        3,
		ThresholdRows:                      len(rows),
		AllRowsExactRational:               allExact,
		AllRowsConditionalOnTexture:        allTex,
		AllRowsConditionalOnVEV:            allVEV,
		AllRowsConditionalOnScheme:         allScheme,
		AnyObservedMassUsed:                anyObs,
		AnyNativeFiniteThresholdRowDerived: anyNative,
		ReconstructsFullFermionU1Y:         sum1.Equal(NewRational(4, 1)),
		ReconstructsFullFermionSU2L:        sum2.Equal(NewRational(4, 1)),
		ReconstructsFullFermionSU3C:        sum3.Equal(NewRational(4, 1)),
		FermionContributionU1Y:             sum1,
		FermionContributionSU2L:            sum2,
		FermionContributionSU3C:            sum3,
		OneLoopContinuumFormulaAssumed:     true,
		NativeFiniteBetaTheoremDerived:     false,
		Verdict:                            "The exact rational sector rows reconstruct the full fermion one-loop contribution (4,4,4). They are conditional continuum representation rows, not native finite threshold operators.",
	}
}

func sectorRow(kind yukawaamplitudeseal.YukawaKind, u1, su2, su3 Rational, hyper, nonab string) SectorRow {
	return SectorRow{
		Sector:                      kind,
		Generations:                 3,
		U1YContribution:             u1,
		SU2LContribution:            su2,
		SU3CContribution:            su3,
		SubtractionVector:           fmt.Sprintf("(-%s, -%s, -%s)", u1.String(), su2.String(), su3.String()),
		HyperchargeSource:           hyper,
		NonAbelianSource:            nonab,
		SplitRequiresBrokenPhaseVEV: true,
		ExactRational:               true,
		Verdict:                     "sector row is an exact rational contribution; separating up/down or electron/neutrino thresholds uses the VEV-broken mass basis and is not an unbroken-SU(2) finite theorem",
	}
}

func buildSchemeSeal(ledger BetaShiftLedger) ContinuumDecouplingSchemeSeal {
	return ContinuumDecouplingSchemeSeal{
		Name:                             "ContinuumDecouplingSchemeSeal",
		AxiomID:                          "CONDITIONAL-CONTINUUM-DECOUPLING-SCHEME-SEAL-G198",
		ConditionalStatus:                "CONDITIONAL_ON_CONTINUUM_DECOUPLING_SCHEME",
		SchemeFamily:                     "sharp step / tree-level continuity scaffold",
		SharpStepPredicate:               "active_f,g(mu)=Theta(mu-M_f,g)",
		ExplicitBoundaryConvention:       true,
		Quarantined:                      true,
		RequiredForThresholdRows:         ledger.ThresholdRows == 12,
		DerivedFromFiniteAlgebra:         false,
		SharpStepDerivedNatively:         false,
		SmoothRegulatorSearched:          true,
		SmoothRegulatorDerived:           false,
		SchemeSelectedNatively:           false,
		MSbarDerived:                     false,
		MOMDerived:                       false,
		TreeLevelContinuityEnforced:      true,
		FiniteMatchingCorrectionsDerived: false,
		DownstreamMustDeclareSeal:        true,
		Verdict:                          "The step function is admitted only as a continuum matching convention. The finite algebra does not select MSbar, MOM, a smooth regulator, or finite threshold matching corrections.",
	}
}

func auditLowEnergyDomain(prev electroweakvevseal.Analysis) LowEnergyDomainAudit {
	blocked := !prev.MassLedger.GaugeBosonMassesAvailable && prev.MassLedger.GaugeBosonMassBlockReason != ""
	return LowEnergyDomainAudit{
		FermionThresholdSymbolsAvailable: prev.MassLedger.AllFormalThresholdsAvailable,
		GaugeBosonThresholdsAvailable:    false,
		WZThresholdsBlocked:              blocked,
		DeepInfraredFlowDefined:          false,
		RunToMZAllowed:                   false,
		DomainRestriction:                "symbolic threshold RG scaffold is restricted to domains above the still-sealed electroweak gauge-boson thresholds; running to M_Z or deep IR is not defined",
		BlockReason:                      prev.MassLedger.GaugeBosonMassBlockReason,
		Verdict:                          "Fermion threshold symbols exist, but W/Z thresholds require physical gauge couplings and kinetic normalization. The low-energy electroweak flow domain remains formally bounded.",
	}
}

func buildPiecewiseRG(ledger BetaShiftLedger, scheme ContinuumDecouplingSchemeSeal, domain LowEnergyDomainAudit) PiecewiseRGAssembly {
	return PiecewiseRGAssembly{
		BetaFunctionExpression:           "b_i^cond(mu)=b_i^{gauge+scalar}+sum_{f,g} Theta(mu-M_{f,g}) c_{f,i}",
		RunningExpression:                "A_i(mu)=A_i(M*)+(1/8π²)∫_{ln mu}^{ln M*} b_i^cond(q)dln q; A_i=1/g_i²",
		ThresholdLogKernel:               "L(mu,M;M*)=∫_{ln mu}^{ln M*}Theta(q-M)dln q, left symbolic because ordering and domain are sealed",
		TreeLevelMatchingCondition:       "A_i(M_{f,g}^-)=A_i(M_{f,g}^+) under the declared sharp-step tree-level matching convention",
		FiniteMatchingCorrectionTerm:     "delta_i^match(M_{f,g}) is scheme-dependent and not finite-derived",
		PiecewiseSymbolicTreeBuilt:       ledger.ThresholdRows == 12 && scheme.ExplicitBoundaryConvention && domain.FermionThresholdSymbolsAvailable,
		EvaluatedNumerically:             false,
		ThresholdOrderingKnown:           false,
		BoundaryScaleDerived:             false,
		BoundaryCouplingDerived:          false,
		FiniteMatchingCorrectionsDerived: false,
		EnforcesTreeLevelContinuity:      scheme.TreeLevelContinuityEnforced,
		SchemeDependentCorrectionsSealed: !scheme.FiniteMatchingCorrectionsDerived,
		Verdict:                          "The engine assembles a symbolic piecewise RG tree and enforces tree-level continuity by convention. Finite matching corrections and threshold ordering are not derived.",
	}
}

func auditFirewall(prev electroweakvevseal.Analysis, ledger BetaShiftLedger, scheme ContinuumDecouplingSchemeSeal, domain LowEnergyDomainAudit, piece PiecewiseRGAssembly) RGFirewallAudit {
	return RGFirewallAudit{
		Gate197VEVSealInherited:               prev.Summary.EmpiricalVEVSealRecorded,
		FormalMassThresholdsAvailable:         prev.MassLedger.AllFormalThresholdsAvailable,
		ContinuumDecouplingSchemeSealInserted: scheme.ExplicitBoundaryConvention && scheme.Quarantined,
		ConditionalBetaRowsActivated:          ledger.ThresholdRows == 12 && ledger.AllRowsExactRational,
		FiniteThresholdRowsDerived:            false,
		NativeSmoothRegulatorDerived:          false,
		SchemeConventionDerivedFromFinite:     false,
		PiecewiseRGTreeConstructed:            piece.PiecewiseSymbolicTreeBuilt,
		PhysicalRGFlowEvaluated:               false,
		NumericalThresholdsKnown:              false,
		ThresholdOrderingKnown:                false,
		LowEnergyWZDomainKnown:                !domain.WZThresholdsBlocked,
		GaugeBosonThresholdsDerived:           false,
		AbsoluteBoundaryScaleDerived:          false,
		AbsoluteBoundaryCouplingDerived:       false,
		GaugeCouplingsDerived:                 false,
		TopologicalEightPiSquaredImported:     false,
		FiniteToContinuumScaleDerived:         false,
		ObservedInputsImported:                false,
		StrictNullityBefore:                   prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                    prev.Firewall.StrictNullityAfter,
		ConditionalSchemeNullityBefore:        1,
		ConditionalSchemeNullityAfter:         0,
		ConditionalBetaRowNullityBefore:       1,
		ConditionalBetaRowNullityAfter:        0,
		ConditionalRGEvaluationNullityBefore:  1,
		ConditionalRGEvaluationNullityAfter:   1,
		OpenRequirements: []string{
			"numerical Yukawa singular values and observed-free threshold ordering remain absent",
			"W/Z and scalar thresholds require physical gauge couplings and kinetic normalization",
			"finite matching corrections require a continuum subtraction/matching scheme beyond the sharp-step scaffold",
			"absolute boundary scale M* and boundary coupling g_*² remain underived",
			"finite-to-continuum normalization and 8π² action promotion remain sealed",
		},
		RecommendedNextGate: "Gate 199 — gauge-coupling boundary seal / symbolic RG evaluation firewall audit",
		Verdict:             "Conditional beta rows and a symbolic RG tree are admitted only under explicit texture, VEV, and decoupling-scheme seals. Physical running remains unevaluated.",
	}
}

func subscript(k yukawaamplitudeseal.YukawaKind) string {
	switch k {
	case yukawaamplitudeseal.YukawaUp:
		return "u"
	case yukawaamplitudeseal.YukawaDown:
		return "d"
	case yukawaamplitudeseal.YukawaElectron:
		return "e"
	case yukawaamplitudeseal.YukawaNeutrino:
		return "nu"
	default:
		return strings.TrimPrefix(string(k), "Y_")
	}
}

func FormatRationalVector(a, b, c Rational) string { return fmt.Sprintf("(%s,%s,%s)", a, b, c) }

func FormatSectorRow(r SectorRow) string {
	return fmt.Sprintf("%s c=%s subtract=%s gens=%d exact=%t brokenSplit=%t hyper=%s nonab=%s", r.Sector, FormatRationalVector(r.U1YContribution, r.SU2LContribution, r.SU3CContribution), r.SubtractionVector, r.Generations, r.ExactRational, r.SplitRequiresBrokenPhaseVEV, r.HyperchargeSource, r.NonAbelianSource)
}

func FormatBetaRow(r BetaRow) string {
	return fmt.Sprintf("%s gen=%d M=%s active=%s c=%s subtract=%s exact=%t texture=%t VEV=%t scheme=%t observed=%t native=%t halfDoublet=%t", r.Sector, r.Generation, r.MassSymbol, r.ActivePredicate, FormatRationalVector(r.U1YContribution, r.SU2LContribution, r.SU3CContribution), FormatRationalVector(r.U1YSubtractionBelowThreshold, r.SU2LSubtractionBelowThreshold, r.SU3CSubtractionBelowThreshold), r.ExactRational, r.RequiresEmpiricalTextureSeal, r.RequiresVEVSeal, r.RequiresContinuumDecouplingScheme, r.UsesObservedMass, r.NativeFiniteThresholdRowDerived, r.BrokenPhaseHalfDoubletBookkeeping)
}

func FormatLedger(l BetaShiftLedger) string {
	parts := make([]string, 0, len(l.SectorRows))
	for _, r := range l.SectorRows {
		parts = append(parts, FormatSectorRow(r))
	}
	sort.Strings(parts)
	return fmt.Sprintf("sectors=%d gens=%d rows=%d exact=%t texture=%t VEV=%t scheme=%t observed=%t native=%t sums=%s reconstruct=(%t,%t,%t) formula=%t finiteBeta=%t [%s]", l.SectorCount, l.Generations, l.ThresholdRows, l.AllRowsExactRational, l.AllRowsConditionalOnTexture, l.AllRowsConditionalOnVEV, l.AllRowsConditionalOnScheme, l.AnyObservedMassUsed, l.AnyNativeFiniteThresholdRowDerived, FormatRationalVector(l.FermionContributionU1Y, l.FermionContributionSU2L, l.FermionContributionSU3C), l.ReconstructsFullFermionU1Y, l.ReconstructsFullFermionSU2L, l.ReconstructsFullFermionSU3C, l.OneLoopContinuumFormulaAssumed, l.NativeFiniteBetaTheoremDerived, strings.Join(parts, "; "))
}

func FormatScheme(s ContinuumDecouplingSchemeSeal) string {
	return fmt.Sprintf("%s axiom=%s status=%s family=%s predicate=%s explicit=%t quarantine=%t required=%t finiteDerived=%t sharpNative=%t smoothSearched=%t smoothDerived=%t schemeNative=%t MSbar=%t MOM=%t continuity=%t finiteCorrections=%t downstream=%t", s.Name, s.AxiomID, s.ConditionalStatus, s.SchemeFamily, s.SharpStepPredicate, s.ExplicitBoundaryConvention, s.Quarantined, s.RequiredForThresholdRows, s.DerivedFromFiniteAlgebra, s.SharpStepDerivedNatively, s.SmoothRegulatorSearched, s.SmoothRegulatorDerived, s.SchemeSelectedNatively, s.MSbarDerived, s.MOMDerived, s.TreeLevelContinuityEnforced, s.FiniteMatchingCorrectionsDerived, s.DownstreamMustDeclareSeal)
}

func FormatDomain(d LowEnergyDomainAudit) string {
	return fmt.Sprintf("fermionSymbols=%t gaugeBoson=%t WZblocked=%t deepIR=%t runToMZ=%t domain=%s block=%s", d.FermionThresholdSymbolsAvailable, d.GaugeBosonThresholdsAvailable, d.WZThresholdsBlocked, d.DeepInfraredFlowDefined, d.RunToMZAllowed, d.DomainRestriction, d.BlockReason)
}

func FormatPiecewise(p PiecewiseRGAssembly) string {
	return fmt.Sprintf("beta=%s running=%s kernel=%s match=%s finiteCorrection=%s tree=%t eval=%t ordering=%t Mstar=%t gstar=%t finiteCorrections=%t continuity=%t correctionsSealed=%t", p.BetaFunctionExpression, p.RunningExpression, p.ThresholdLogKernel, p.TreeLevelMatchingCondition, p.FiniteMatchingCorrectionTerm, p.PiecewiseSymbolicTreeBuilt, p.EvaluatedNumerically, p.ThresholdOrderingKnown, p.BoundaryScaleDerived, p.BoundaryCouplingDerived, p.FiniteMatchingCorrectionsDerived, p.EnforcesTreeLevelContinuity, p.SchemeDependentCorrectionsSealed)
}

func FormatFirewall(f RGFirewallAudit) string {
	return fmt.Sprintf("g197=%t massSymbols=%t schemeSeal=%t condRows=%t finiteRows=%t smooth=%t schemeNative=%t tree=%t physicalRG=%t numeric=%t ordering=%t lowWZ=%t WZ=%t Mstar=%t gstar=%t gauge=%t 8pi2=%t contScale=%t observed=%t strict=%d->%d scheme=%d->%d rows=%d->%d rgEval=%d->%d next=%s", f.Gate197VEVSealInherited, f.FormalMassThresholdsAvailable, f.ContinuumDecouplingSchemeSealInserted, f.ConditionalBetaRowsActivated, f.FiniteThresholdRowsDerived, f.NativeSmoothRegulatorDerived, f.SchemeConventionDerivedFromFinite, f.PiecewiseRGTreeConstructed, f.PhysicalRGFlowEvaluated, f.NumericalThresholdsKnown, f.ThresholdOrderingKnown, f.LowEnergyWZDomainKnown, f.GaugeBosonThresholdsDerived, f.AbsoluteBoundaryScaleDerived, f.AbsoluteBoundaryCouplingDerived, f.GaugeCouplingsDerived, f.TopologicalEightPiSquaredImported, f.FiniteToContinuumScaleDerived, f.ObservedInputsImported, f.StrictNullityBefore, f.StrictNullityAfter, f.ConditionalSchemeNullityBefore, f.ConditionalSchemeNullityAfter, f.ConditionalBetaRowNullityBefore, f.ConditionalBetaRowNullityAfter, f.ConditionalRGEvaluationNullityBefore, f.ConditionalRGEvaluationNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d g197=%t rows=%t scheme=%t continuity=%t tree=%t WZ=%t coupling=%t :: %s", s.TestsAudited, s.Gate197Inherited, s.ExactConditionalBetaRowsBuilt, s.DecouplingSchemeSealRecorded, s.TreeLevelContinuityAnswerRecorded, s.PiecewiseRGScaffoldBuilt, s.WZLowEnergyFirewallPreserved, s.AbsoluteCouplingFirewallPreserved, s.Comment)
}
