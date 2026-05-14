package asha

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

type Status string

const (
	StatusNative        Status = "native/audited"
	StatusBridge        Status = "bridge-required"
	StatusQuarantined   Status = "quarantined-axiom"
	StatusEnvironmental Status = "environmental"
	StatusFailedRoute   Status = "failed-route"
	StatusDerived       Status = "derived-runtime"
)

type Scenario string

const (
	ScenarioAll               Scenario = "all"
	ScenarioNative            Scenario = "native"
	ScenarioHiggs             Scenario = "higgs"
	ScenarioFamily            Scenario = "family"
	ScenarioDarkStableThermal Scenario = "dark-stable-thermal"
	ScenarioCosmology         Scenario = "cosmology"
	ScenarioEnvironment       Scenario = "environment"
	ScenarioCI                Scenario = "ci"
)

func ParseScenario(s string) (Scenario, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	switch Scenario(s) {
	case ScenarioAll, ScenarioNative, ScenarioHiggs, ScenarioFamily, ScenarioDarkStableThermal, ScenarioCosmology, ScenarioEnvironment, ScenarioCI:
		return Scenario(s), nil
	case "dark", "dark-sector", "thermal-majorana":
		return ScenarioDarkStableThermal, nil
	case "env", "phenomenology", "environmental", "cosmo-dark", "vacuum-fate":
		return ScenarioEnvironment, nil
	case "kxy", "flavor":
		return ScenarioFamily, nil
	default:
		return "", fmt.Errorf("unknown scenario %q", s)
	}
}

type Rational struct {
	Num int64 `json:"num"`
	Den int64 `json:"den"`
}

func NewRational(num, den int64) Rational {
	if den == 0 {
		panic("asha: zero denominator")
	}
	if den < 0 {
		num, den = -num, -den
	}
	g := gcd(abs(num), abs(den))
	return Rational{Num: num / g, Den: den / g}
}

func (r Rational) Float64() float64 { return float64(r.Num) / float64(r.Den) }
func (r Rational) String() string {
	if r.Den == 1 {
		return fmt.Sprintf("%d", r.Num)
	}
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}
func (r Rational) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Num     int64   `json:"num"`
		Den     int64   `json:"den"`
		Decimal float64 `json:"decimal"`
		Text    string  `json:"text"`
	}{r.Num, r.Den, r.Float64(), r.String()})
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	}
	return a
}

type Quantity struct {
	Symbol  string  `json:"symbol"`
	Value   float64 `json:"value,omitempty"`
	Text    string  `json:"text,omitempty"`
	Unit    string  `json:"unit,omitempty"`
	Formula string  `json:"formula,omitempty"`
	Status  Status  `json:"status"`
	Note    string  `json:"note,omitempty"`
}

type Boundary struct {
	Name    string `json:"name"`
	Formula string `json:"formula"`
	Status  Status `json:"status"`
	Meaning string `json:"meaning"`
}

type Check struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
	Detail string `json:"detail"`
}

type Section struct {
	Name       string     `json:"name"`
	Summary    string     `json:"summary"`
	Quantities []Quantity `json:"quantities,omitempty"`
	Boundaries []Boundary `json:"boundaries,omitempty"`
	Checks     []Check    `json:"checks,omitempty"`
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
