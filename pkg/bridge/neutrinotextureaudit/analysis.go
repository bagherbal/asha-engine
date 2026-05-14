// Package neutrinotextureaudit implements Gate 232:
// Neutrino flavor texture audit / NeutrinoTextureSeal activation.
//
// Gate 231 activated the IntermediateBreakingSeal and showed that an order-one
// Type-I seesaw at M_int gives m_nu≈91 eV, while the atmospheric scale requires
// a sealed Dirac Yukawa y_nu≈0.023 for the third generation. Gate 232 asks a
// narrower question: if M_R is degenerate at the sealed intermediate scale and
// the third Dirac mass is fixed by the atmospheric scale, can simple
// three-generation Dirac mass textures reproduce the solar/atmospheric ratio?
//
// The audit is intentionally phenomenological. It introduces a
// NeutrinoTextureSeal, compares standard hierarchy proxies, and refuses to
// derive PMNS angles, Majorana matrices, or finite flavor textures.
package neutrinotextureaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/intermediatebreakingseesaw"
)

const (
	AuditID = "GATE232-NEUTRINO-FLAVOR-TEXTURE-SEAL-AUDIT"

	StatusTextureSealActivated       = "NEUTRINO_TEXTURE_SEAL_ACTIVATED_PHENOMENOLOGICALLY"
	StatusSMMassProxyFailed          = "FAILED_ROUTE_SM_MASS_PROXY_TEXTURES_TOO_HIERARCHICAL"
	StatusGenerationQuadraticSupport = "CONDITIONAL_SUPPORT_TEXTURE_RESONANCE_GENERATION_QUADRATIC"
	StatusFiniteTextureNotDerived    = "FAILED_ROUTE_FINITE_NEUTRINO_TEXTURE_DERIVATION"
	StatusPMNSNotification           = "PMNS_AND_MASS_ORDERING_NOT_DERIVED"
)

const (
	solarDeltaM2EV2       = 7.5e-5
	atmosphericDeltaM2EV2 = 2.5e-3

	// Empirical masses used only to define comparison proxies. They are not
	// finite-derived and they do not enter the previous Gate-231 seesaw scale.
	meGeV   = 0.00051099895
	mmuGeV  = 0.1056583755
	mtauGeV = 1.77686

	muGeV = 0.00216
	mcGeV = 1.27
	mtGeV = 172.56

	mdGeV = 0.00467
	msGeV = 0.093
	mbGeV = 4.18
)

type Gate231Snapshot struct {
	Gate231Inherited                 bool
	IntermediateSealActive           bool
	SmallYukawaSeesawConditionallyOK bool
	FiniteNeutrinoMatrixDerived      bool
	MIntGeV                          float64
	VEVGeV                           float64
	MDirac3GeV                       float64
	YNu3                             float64
	AtmosphericTargetEV              float64
	OrderOneMassEV                   float64
	TruthStatement                   string
}

type NeutrinoTextureSeal struct {
	Name                     string
	AxiomID                  string
	Active                   bool
	PhenomenologicalBoundary bool
	FiniteDerived            bool
	Assumption               string
	DegenerateMRGeV          float64
	MDirac3GeV               float64
	YNu3                     float64
	DerivesPMNS              bool
	DerivesMassOrdering      bool
	DerivesMajoranaMatrix    bool
	DerivesDiracMatrix       bool
	Verdict                  string
}

type ExperimentalRatio struct {
	SolarDeltaM2EV2       float64
	AtmosphericDeltaM2EV2 float64
	Ratio                 float64
	ToleranceRelative     float64
}

type TextureCandidate struct {
	Name                  string
	Kind                  string
	Source                string
	DiracMassesGeV        [3]float64
	Yukawas               [3]float64
	ActiveMassesEV        [3]float64
	RatioM2ToM3           float64
	RatioError            float64
	RatioWithinTolerance  bool
	UsesEmpiricalSMMasses bool
	UsesGenerationIndex   bool
	FiniteDerived         bool
	Comment               string
}

type TextureAudit struct {
	Candidates                  []TextureCandidate
	Best                        TextureCandidate
	BestStandardSMMassProxy     TextureCandidate
	BestGenerationIndexProxy    TextureCandidate
	AnySMMassProxySupported     bool
	AnyGenerationProxySupported bool
	RequiredM2DiracGeV          float64
	RequiredY2                  float64
	RequiredM2OverM3Dirac       float64
	RequiredPowerIndexLaw       float64
	Verdict                     string
}

type MatrixObstruction struct {
	RightHandedNeutrinoFieldsDerived bool
	DegenerateMajoranaMatrixDerived  bool
	DiracTextureDerived              bool
	PMNSMatrixDerived                bool
	CPPhasesDerived                  bool
	MassOrderingDerived              bool
	ThreeActiveEigenvaluesDerived    bool
	OnlyRatioPreflightAvailable      bool
	Verdict                          string
}

type FirewallAudit struct {
	UsesGate231SealedScale        bool
	ActivatesNeutrinoTextureSeal  bool
	ClaimsFiniteTexture           bool
	ClaimsFinitePMNS              bool
	ClaimsFiniteMajoranaMatrix    bool
	TunesToObservedMixingAngles   bool
	UsesObservedRatioAsTargetOnly bool
	ReopensIntermediateDynamics   bool
	ReopensPatiSalam              bool
	FiniteCorePolluted            bool
	Verdict                       string
}

type Summary struct {
	NeutrinoTextureSealActive  bool
	SMProxyTextureSupported    bool
	GenerationTextureSupported bool
	FiniteTextureDerived       bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	Gate231  Gate231Snapshot
	Seal     NeutrinoTextureSeal
	Ratio    ExperimentalRatio
	Audit    TextureAudit
	Matrix   MatrixObstruction
	Firewall FirewallAudit
	Summary  Summary

	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g231, err := intermediatebreakingseesaw.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 231 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g231)
	})
	return defaultA, defaultErr
}

func Build(g231 intermediatebreakingseesaw.Analysis) (Analysis, error) {
	snap := snapshotFromGate231(g231)
	if !snap.Gate231Inherited || !snap.IntermediateSealActive || snap.MIntGeV <= 0 || snap.MDirac3GeV <= 0 {
		return Analysis{}, fmt.Errorf("Gate 232 requires Gate 231 sealed intermediate scale and atmospheric Dirac mass")
	}
	seal := activateSeal(snap)
	ratio := buildRatio()
	audit := auditTextures(seal, ratio)
	matrix := auditMatrix()
	firewall := auditFirewall(snap, seal, audit, matrix)
	summary := summarize(seal, audit, matrix)
	truth := buildTruth(snap, seal, ratio, audit, matrix)
	return Analysis{Gate231: snap, Seal: seal, Ratio: ratio, Audit: audit, Matrix: matrix, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate231(a intermediatebreakingseesaw.Analysis) Gate231Snapshot {
	return Gate231Snapshot{
		Gate231Inherited:                 a.Summary.Status != "" && a.Seal.Active,
		IntermediateSealActive:           a.Seal.Active,
		SmallYukawaSeesawConditionallyOK: a.Summary.SmallYukawaSeesawConditionallyOK,
		FiniteNeutrinoMatrixDerived:      a.Summary.NeutrinoMatrixDerived,
		MIntGeV:                          a.Seal.ScaleGeV,
		VEVGeV:                           a.Input.VEVGeV,
		MDirac3GeV:                       a.Compute.DiracMassForAtmosphericGeV,
		YNu3:                             a.Compute.YukawaForAtmosphericScale,
		AtmosphericTargetEV:              a.Bounds.AtmosphericScaleEV,
		OrderOneMassEV:                   a.Compute.OrderOneMassEV,
		TruthStatement:                   a.TruthStatement,
	}
}

func activateSeal(g Gate231Snapshot) NeutrinoTextureSeal {
	return NeutrinoTextureSeal{
		Name:                     "NeutrinoTextureSeal",
		AxiomID:                  "SEAL-NEUTRINO-TEXTURE-GATE232",
		Active:                   true,
		PhenomenologicalBoundary: true,
		FiniteDerived:            false,
		Assumption:               "A three-generation Dirac neutrino texture and a degenerate right-handed Majorana scale exist at M_int; only ratio-level preflights are permitted.",
		DegenerateMRGeV:          g.MIntGeV,
		MDirac3GeV:               g.MDirac3GeV,
		YNu3:                     g.YNu3,
		DerivesPMNS:              false,
		DerivesMassOrdering:      false,
		DerivesMajoranaMatrix:    false,
		DerivesDiracMatrix:       false,
		Verdict:                  StatusTextureSealActivated,
	}
}

func buildRatio() ExperimentalRatio {
	return ExperimentalRatio{
		SolarDeltaM2EV2:       solarDeltaM2EV2,
		AtmosphericDeltaM2EV2: atmosphericDeltaM2EV2,
		Ratio:                 math.Sqrt(solarDeltaM2EV2 / atmosphericDeltaM2EV2),
		ToleranceRelative:     0.25,
	}
}

func auditTextures(seal NeutrinoTextureSeal, r ExperimentalRatio) TextureAudit {
	m3 := seal.MDirac3GeV
	v := intermediatebreakingseesaw.FormatSeal // keep package imported in old Go versions? no-op impossible
	_ = v
	candidates := []TextureCandidate{
		candidateFromMassProxy("charged-lepton direct", "SM_MASS_PROXY", "mD_i ∝ (m_e,m_mu,m_tau)", [3]float64{meGeV, mmuGeV, mtauGeV}, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromMassProxy("up-quark direct", "SM_MASS_PROXY", "mD_i ∝ (m_u,m_c,m_t)", [3]float64{muGeV, mcGeV, mtGeV}, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromMassProxy("down-quark direct", "SM_MASS_PROXY", "mD_i ∝ (m_d,m_s,m_b)", [3]float64{mdGeV, msGeV, mbGeV}, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromMassProxyPower("charged-lepton square-root", "SM_MASS_PROXY_SQRT", "mD_i ∝ sqrt(m_e,m_mu,m_tau)", [3]float64{meGeV, mmuGeV, mtauGeV}, 0.5, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromMassProxyPower("up-quark square-root", "SM_MASS_PROXY_SQRT", "mD_i ∝ sqrt(m_u,m_c,m_t)", [3]float64{muGeV, mcGeV, mtGeV}, 0.5, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromMassProxyPower("down-quark square-root", "SM_MASS_PROXY_SQRT", "mD_i ∝ sqrt(m_d,m_s,m_b)", [3]float64{mdGeV, msGeV, mbGeV}, 0.5, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromRatios("generation-index linear", "GENERATION_INDEX", "mD_i ∝ i", [3]float64{1.0 / 3.0, 2.0 / 3.0, 1}, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromRatios("generation-index quadratic", "GENERATION_INDEX", "mD_i ∝ i^2", [3]float64{1.0 / 9.0, 4.0 / 9.0, 1}, m3, seal.DegenerateMRGeV, seal.YNu3, r),
		candidateFromRatios("generation-index cubic", "GENERATION_INDEX", "mD_i ∝ i^3", [3]float64{1.0 / 27.0, 8.0 / 27.0, 1}, m3, seal.DegenerateMRGeV, seal.YNu3, r),
	}
	sort.SliceStable(candidates, func(i, j int) bool { return candidates[i].RatioError < candidates[j].RatioError })
	best := candidates[0]
	bestSM := TextureCandidate{RatioError: math.Inf(1)}
	bestGen := TextureCandidate{RatioError: math.Inf(1)}
	anySM := false
	anyGen := false
	for _, c := range candidates {
		if c.UsesEmpiricalSMMasses && c.RatioError < bestSM.RatioError {
			bestSM = c
		}
		if c.UsesGenerationIndex && c.RatioError < bestGen.RatioError {
			bestGen = c
		}
		if c.UsesEmpiricalSMMasses && c.RatioWithinTolerance {
			anySM = true
		}
		if c.UsesGenerationIndex && c.RatioWithinTolerance {
			anyGen = true
		}
	}
	requiredDiracRatio := math.Sqrt(r.Ratio)
	requiredM2 := m3 * requiredDiracRatio
	requiredPower := math.Log(requiredDiracRatio) / math.Log(2.0/3.0)
	verdict := StatusSMMassProxyFailed
	if anyGen {
		verdict = StatusGenerationQuadraticSupport
	}
	return TextureAudit{Candidates: candidates, Best: best, BestStandardSMMassProxy: bestSM, BestGenerationIndexProxy: bestGen, AnySMMassProxySupported: anySM, AnyGenerationProxySupported: anyGen, RequiredM2DiracGeV: requiredM2, RequiredY2: seal.YNu3 * requiredDiracRatio, RequiredM2OverM3Dirac: requiredDiracRatio, RequiredPowerIndexLaw: requiredPower, Verdict: verdict}
}

func candidateFromMassProxy(name, kind, source string, masses [3]float64, m3, mR, y3 float64, r ExperimentalRatio) TextureCandidate {
	return candidateFromMassProxyPower(name, kind, source, masses, 1, m3, mR, y3, r)
}

func candidateFromMassProxyPower(name, kind, source string, masses [3]float64, power, m3, mR, y3 float64, r ExperimentalRatio) TextureCandidate {
	ratios := [3]float64{}
	denom := math.Pow(masses[2], power)
	for i := range masses {
		ratios[i] = math.Pow(masses[i], power) / denom
	}
	c := candidateFromRatios(name, kind, source, ratios, m3, mR, y3, r)
	c.UsesEmpiricalSMMasses = true
	c.Comment = "Empirical SM mass proxy only; not finite-derived."
	return c
}

func candidateFromRatios(name, kind, source string, ratios [3]float64, m3, mR, y3 float64, r ExperimentalRatio) TextureCandidate {
	var md, y, mn [3]float64
	for i := range ratios {
		md[i] = m3 * ratios[i]
		y[i] = y3 * ratios[i]
		mn[i] = md[i] * md[i] / mR * 1e9
	}
	ratio := mn[1] / mn[2]
	err := math.Abs(ratio-r.Ratio) / r.Ratio
	return TextureCandidate{
		Name:                  name,
		Kind:                  kind,
		Source:                source,
		DiracMassesGeV:        md,
		Yukawas:               y,
		ActiveMassesEV:        mn,
		RatioM2ToM3:           ratio,
		RatioError:            err,
		RatioWithinTolerance:  err <= r.ToleranceRelative,
		UsesEmpiricalSMMasses: false,
		UsesGenerationIndex:   kind == "GENERATION_INDEX",
		FiniteDerived:         false,
		Comment:               "Phenomenological texture proxy; no finite flavor theorem.",
	}
}

func auditMatrix() MatrixObstruction {
	return MatrixObstruction{
		RightHandedNeutrinoFieldsDerived: false,
		DegenerateMajoranaMatrixDerived:  false,
		DiracTextureDerived:              false,
		PMNSMatrixDerived:                false,
		CPPhasesDerived:                  false,
		MassOrderingDerived:              false,
		ThreeActiveEigenvaluesDerived:    false,
		OnlyRatioPreflightAvailable:      true,
		Verdict:                          StatusFiniteTextureNotDerived,
	}
}

func auditFirewall(g Gate231Snapshot, seal NeutrinoTextureSeal, audit TextureAudit, matrix MatrixObstruction) FirewallAudit {
	return FirewallAudit{
		UsesGate231SealedScale:        g.IntermediateSealActive && seal.DegenerateMRGeV == g.MIntGeV,
		ActivatesNeutrinoTextureSeal:  seal.Active && seal.PhenomenologicalBoundary && !seal.FiniteDerived,
		ClaimsFiniteTexture:           matrix.DiracTextureDerived,
		ClaimsFinitePMNS:              matrix.PMNSMatrixDerived,
		ClaimsFiniteMajoranaMatrix:    matrix.DegenerateMajoranaMatrixDerived,
		TunesToObservedMixingAngles:   false,
		UsesObservedRatioAsTargetOnly: true,
		ReopensIntermediateDynamics:   false,
		ReopensPatiSalam:              false,
		FiniteCorePolluted:            false,
		Verdict:                       "FIREWALLS_PRESERVED_TEXTURE_PREFLIGHT_ONLY",
	}
}

func summarize(seal NeutrinoTextureSeal, audit TextureAudit, matrix MatrixObstruction) Summary {
	parts := []string{StatusTextureSealActivated}
	if !audit.AnySMMassProxySupported {
		parts = append(parts, StatusSMMassProxyFailed)
	}
	if audit.AnyGenerationProxySupported {
		parts = append(parts, StatusGenerationQuadraticSupport)
	}
	if !matrix.DiracTextureDerived {
		parts = append(parts, StatusFiniteTextureNotDerived, StatusPMNSNotification)
	}
	return Summary{
		NeutrinoTextureSealActive:  seal.Active,
		SMProxyTextureSupported:    audit.AnySMMassProxySupported,
		GenerationTextureSupported: audit.AnyGenerationProxySupported,
		FiniteTextureDerived:       matrix.DiracTextureDerived,
		Status:                     strings.Join(parts, ";"),
		NextGate:                   "derive or seal PMNS/rank structure; do not infer exact neutrino masses from ratio-level texture proxies",
		Comment:                    "Gate 232 finds that direct SM mass proxies are much too hierarchical, while a simple quadratic generation-index Dirac texture gives a ratio close to the solar/atmospheric scale. This is conditional phenomenology, not a finite flavor theorem.",
	}
}

func buildTruth(g Gate231Snapshot, seal NeutrinoTextureSeal, r ExperimentalRatio, audit TextureAudit, matrix MatrixObstruction) string {
	return fmt.Sprintf("Gate 232 activates %s at M_R=%.12e GeV and fixes mD3=%.8g GeV from the Gate-231 atmospheric-scale preflight. The target solar/atmospheric mass ratio is r=%.8g. Direct SM mass proxies fail: best SM proxy %q gives m2/m3=%.8g. A quadratic generation-index texture mD∝i^2 gives m2/m3=%.8g, within %.1f%% tolerance, but no finite Dirac texture, Majorana matrix, PMNS matrix, CP phase, or mass ordering is derived. The required Dirac ratio is mD2/mD3=%.8g (mD2≈%.8g GeV, y2≈%.8g).",
		seal.Name, seal.DegenerateMRGeV, seal.MDirac3GeV, r.Ratio, audit.BestStandardSMMassProxy.Name, audit.BestStandardSMMassProxy.RatioM2ToM3, audit.BestGenerationIndexProxy.RatioM2ToM3, r.ToleranceRelative*100, audit.RequiredM2OverM3Dirac, audit.RequiredM2DiracGeV, audit.RequiredY2)
}

func FormatGate231(g Gate231Snapshot) string {
	return fmt.Sprintf("inherited=%t intermediateSeal=%t smallYukawaOK=%t finiteMatrix=%t M_int=%.12e v=%.5f mD3=%.10g y3=%.10g atmTarget=%.6g orderOneMass=%.10g eV",
		g.Gate231Inherited, g.IntermediateSealActive, g.SmallYukawaSeesawConditionallyOK, g.FiniteNeutrinoMatrixDerived, g.MIntGeV, g.VEVGeV, g.MDirac3GeV, g.YNu3, g.AtmosphericTargetEV, g.OrderOneMassEV)
}

func FormatSeal(s NeutrinoTextureSeal) string {
	return fmt.Sprintf("name=%s axiom=%s active=%t phenomenological=%t finiteDerived=%t M_R=%.12e mD3=%.10g y3=%.10g derivesPMNS=%t derivesOrdering=%t derivesMajorana=%t derivesDirac=%t verdict=%s assumption=%q",
		s.Name, s.AxiomID, s.Active, s.PhenomenologicalBoundary, s.FiniteDerived, s.DegenerateMRGeV, s.MDirac3GeV, s.YNu3, s.DerivesPMNS, s.DerivesMassOrdering, s.DerivesMajoranaMatrix, s.DerivesDiracMatrix, s.Verdict, s.Assumption)
}

func FormatRatio(r ExperimentalRatio) string {
	return fmt.Sprintf("solarDm2=%.8g atmosphericDm2=%.8g targetRatio=sqrt(solar/atm)=%.10g tolerance=%.1f%%", r.SolarDeltaM2EV2, r.AtmosphericDeltaM2EV2, r.Ratio, r.ToleranceRelative*100)
}

func FormatCandidate(c TextureCandidate) string {
	return fmt.Sprintf("name=%q kind=%s source=%q mD=[%.6g,%.6g,%.6g]GeV y=[%.6g,%.6g,%.6g] mnu=[%.6g,%.6g,%.6g]eV ratio=%.10g relErr=%.4g withinTol=%t SMmass=%t genIndex=%t finiteDerived=%t comment=%q",
		c.Name, c.Kind, c.Source, c.DiracMassesGeV[0], c.DiracMassesGeV[1], c.DiracMassesGeV[2], c.Yukawas[0], c.Yukawas[1], c.Yukawas[2], c.ActiveMassesEV[0], c.ActiveMassesEV[1], c.ActiveMassesEV[2], c.RatioM2ToM3, c.RatioError, c.RatioWithinTolerance, c.UsesEmpiricalSMMasses, c.UsesGenerationIndex, c.FiniteDerived, c.Comment)
}

func FormatAudit(a TextureAudit) string {
	return fmt.Sprintf("candidates=%d best=%s bestSM=%s bestGen=%s anySM=%t anyGen=%t required mD2/mD3=%.10g mD2=%.10gGeV y2=%.10g indexPower=%.10g verdict=%s",
		len(a.Candidates), FormatCandidate(a.Best), FormatCandidate(a.BestStandardSMMassProxy), FormatCandidate(a.BestGenerationIndexProxy), a.AnySMMassProxySupported, a.AnyGenerationProxySupported, a.RequiredM2OverM3Dirac, a.RequiredM2DiracGeV, a.RequiredY2, a.RequiredPowerIndexLaw, a.Verdict)
}

func FormatMatrix(m MatrixObstruction) string {
	return fmt.Sprintf("RHfields=%t degenerateMajorana=%t DiracTexture=%t PMNS=%t CP=%t ordering=%t eigenvalues=%t ratioOnly=%t verdict=%s",
		m.RightHandedNeutrinoFieldsDerived, m.DegenerateMajoranaMatrixDerived, m.DiracTextureDerived, m.PMNSMatrixDerived, m.CPPhasesDerived, m.MassOrderingDerived, m.ThreeActiveEigenvaluesDerived, m.OnlyRatioPreflightAvailable, m.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate231Scale=%t textureSeal=%t finiteTexture=%t finitePMNS=%t finiteMajorana=%t tunesMixing=%t ratioTargetOnly=%t reopenIntermediate=%t reopenPS=%t polluted=%t verdict=%s",
		f.UsesGate231SealedScale, f.ActivatesNeutrinoTextureSeal, f.ClaimsFiniteTexture, f.ClaimsFinitePMNS, f.ClaimsFiniteMajoranaMatrix, f.TunesToObservedMixingAngles, f.UsesObservedRatioAsTargetOnly, f.ReopensIntermediateDynamics, f.ReopensPatiSalam, f.FiniteCorePolluted, f.Verdict)
}
