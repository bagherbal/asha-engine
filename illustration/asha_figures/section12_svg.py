from __future__ import annotations

from pathlib import Path
import html
from math import sin, cos

from .style import PALETTE, STYLE, FigureStyle
from .readme_gravity import FigureContract


def esc(s: object) -> str:
    return html.escape(str(s), quote=True)


def svg_open(title: str, desc: str, style: FigureStyle = STYLE) -> list[str]:
    p = PALETTE
    w, h = style.width, style.height
    return [f'''<svg xmlns="http://www.w3.org/2000/svg" width="{w}" height="{h}" viewBox="0 0 {w} {h}" role="img" aria-labelledby="title desc">
<title id="title">{esc(title)}</title>
<desc id="desc">{esc(desc)}</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="72%"><stop offset="0%" stop-color="#10101B"/><stop offset="62%" stop-color="{p.abyss}"/><stop offset="100%" stop-color="#010103"/></radialGradient>
  <linearGradient id="metricBridge" x1="0" x2="1"><stop offset="0%" stop-color="{p.cyan}" stop-opacity="0.25"/><stop offset="48%" stop-color="{p.platinum}" stop-opacity="0.86"/><stop offset="100%" stop-color="{p.gold}" stop-opacity="0.38"/></linearGradient>
  <linearGradient id="firewall" x1="0" x2="1"><stop offset="0%" stop-color="#050508"/><stop offset="50%" stop-color="#15151E"/><stop offset="100%" stop-color="#050508"/></linearGradient>
  <radialGradient id="curvatureGlow" cx="50%" cy="50%" r="60%"><stop offset="0%" stop-color="{p.gold}" stop-opacity="0.35"/><stop offset="60%" stop-color="{p.cyan}" stop-opacity="0.08"/><stop offset="100%" stop-color="{p.cyan}" stop-opacity="0"/></radialGradient>
  <filter id="softGlow" x="-180%" y="-180%" width="460%" height="460%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-220%" y="-220%" width="540%" height="540%"><feGaussianBlur stdDeviation="15" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="10" markerHeight="10" refX="7" refY="3" orient="auto"><path d="M0,0 L0,6 L8,3 z" fill="{p.platinum}" fill-opacity="0.75"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 58px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 34px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 25px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 2.5px; fill: {p.muted}; }}
    .tiny {{ font-family: {style.label_font}; font-size: 14px; letter-spacing: 1.35px; fill: #8B93A5; }}
    .caption {{ font-family: {style.label_font}; font-size: 19px; fill: #AAB3C2; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{w-96}" height="{h-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README · GRAVITY · SOURCE-TYPED LOW-ENERGY BOUNDARY</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
<text class="mathSmall" x="{w/2:.0f}" y="278" text-anchor="middle">flat projection theorem · conditional Einstein-Hilbert lane · MP and Lambda wounds</text>
''']


def boxed(parts: list[str], x: float, y: float, width: float, height: float, title: str, subtitle: str, stroke: str, fill: str = "#071018", opacity: float = 0.42) -> None:
    parts.append(f'<rect x="{x}" y="{y}" width="{width}" height="{height}" rx="34" fill="{fill}" fill-opacity="{opacity}" stroke="{stroke}" stroke-opacity="0.38"/>')
    parts.append(f'<text class="label" x="{x+34}" y="{y+50}">{esc(title)}</text>')
    parts.append(f'<text class="tiny" x="{x+34}" y="{y+82}">{esc(subtitle)}</text>')


def component_axis(parts: list[str], x: float, y: float, label: str, color: str, warm: bool = False) -> None:
    parts.append(f'<line x1="{x}" y1="{y+68}" x2="{x}" y2="{y-68}" stroke="{color}" stroke-opacity="0.65" stroke-width="3" filter="url(#softGlow)"/>')
    parts.append(f'<circle cx="{x}" cy="{y-68}" r="9" fill="{color}" fill-opacity="0.90" filter="url(#softGlow)"/>')
    parts.append(f'<circle cx="{x}" cy="{y+68}" r="7" fill="{color}" fill-opacity="0.38"/>')
    parts.append(f'<text class="tiny" x="{x}" y="{y+102}" text-anchor="middle" fill="{color}">{esc(label)}</text>')
    if warm:
        parts.append(f'<circle cx="{x}" cy="{y}" r="42" fill="{color}" fill-opacity="0.07" stroke="{color}" stroke-opacity="0.25"/>')


def draw_phase_octave(parts: list[str], x0: float, y0: float, p) -> None:
    parts.append(f'<ellipse cx="{x0+300}" cy="{y0+175}" rx="330" ry="195" fill="{p.cyan}" fill-opacity="0.035" stroke="{p.cyan}" stroke-opacity="0.22"/>')
    parts.append(f'<text class="mathSmall" x="{x0+300}" y="{y0+42}" text-anchor="middle">V8 = X4 ⊕ P4</text>')
    labels = ["x0", "x1", "x2", "x3", "p0", "p1", "p2", "p3"]
    for i, lab in enumerate(labels):
        xx = x0 + 70 + i*66
        color = p.warm if lab in ("x0", "p0") else (p.cyan if lab.startswith("x") else p.frost)
        component_axis(parts, xx, y0+185, lab, color, warm=(lab in ("x0", "p0")))
    parts.append(f'<text class="tiny" x="{x0+300}" y="{y0+345}" text-anchor="middle">eta_ASHA = eta_1,3 ⊕ (-I4) · signature (1,7)</text>')
    parts.append(f'<text class="tiny" x="{x0+300}" y="{y0+376}" text-anchor="middle">Omega = Σ dp_mu ∧ dx^mu</text>')


def draw_minkowski_target(parts: list[str], x0: float, y0: float, p) -> None:
    parts.append(f'<rect x="{x0}" y="{y0}" width="520" height="390" rx="42" fill="#050508" fill-opacity="0.62" stroke="{p.platinum}" stroke-opacity="0.28"/>')
    parts.append(f'<text class="mathSmall" x="{x0+260}" y="{y0+62}" text-anchor="middle">X4, eta_1,3</text>')
    # draw a mini light cone
    cx, cy = x0+260, y0+210
    parts.append(f'<line x1="{cx}" y1="{cy+116}" x2="{cx}" y2="{cy-132}" stroke="{p.warm}" stroke-opacity="0.78" stroke-width="3" filter="url(#softGlow)"/>')
    parts.append(f'<line x1="{cx-156}" y1="{cy}" x2="{cx+156}" y2="{cy}" stroke="{p.cyan}" stroke-opacity="0.55" stroke-width="3"/>')
    parts.append(f'<path d="M {cx} {cy-122} L {cx-132} {cy+82} M {cx} {cy-122} L {cx+132} {cy+82} M {cx} {cy+122} L {cx-132} {cy-82} M {cx} {cy+122} L {cx+132} {cy-82}" stroke="{p.platinum}" stroke-opacity="0.28" stroke-width="2" fill="none"/>')
    parts.append(f'<ellipse cx="{cx}" cy="{cy}" rx="134" ry="74" fill="none" stroke="{p.gold}" stroke-opacity="0.24" stroke-dasharray="7 8"/>')
    parts.append(f'<text class="tiny" x="{x0+260}" y="{y0+345}" text-anchor="middle">g_X = Pi_X^* eta_ASHA Pi_X = eta_1,3</text>')


def draw_curved_grid(parts: list[str], x0: float, y0: float, width: float, height: float, p) -> None:
    cx, cy = x0 + width/2, y0 + height/2
    parts.append(f'<ellipse cx="{cx}" cy="{cy}" rx="470" ry="210" fill="url(#curvatureGlow)" opacity="0.62"/>')
    # Curved grid families.
    for j in range(9):
        yy = y0 + 55 + j*(height-110)/8
        d = 42 * sin((j-4)/4)
        parts.append(f'<path d="M {x0+70} {yy:.1f} C {x0+width*0.34} {yy+d:.1f} {x0+width*0.66} {yy-d:.1f} {x0+width-70} {yy:.1f}" stroke="{p.cyan}" stroke-opacity="0.16" stroke-width="1.8" fill="none"/>')
    for i in range(10):
        xx = x0 + 70 + i*(width-140)/9
        d = 58 * cos((i-4.5)/4)
        parts.append(f'<path d="M {xx:.1f} {y0+55} C {xx-d:.1f} {y0+height*0.35} {xx+d:.1f} {y0+height*0.65} {xx:.1f} {y0+height-55}" stroke="{p.platinum}" stroke-opacity="0.13" stroke-width="1.5" fill="none"/>')
    parts.append(f'<circle cx="{cx}" cy="{cy}" r="72" fill="{p.obsidian}" fill-opacity="0.72" stroke="{p.gold}" stroke-opacity="0.48" filter="url(#bigGlow)"/>')
    parts.append(f'<text class="math" x="{cx}" y="{cy+8}" text-anchor="middle" fill="{p.gold}">R</text>')


def source_card(parts: list[str], x: float, y: float, title: str, body: str, status: str, color: str, wound: bool = False) -> None:
    dash = ' stroke-dasharray="10 9"' if wound else ''
    fill = "url(#firewall)" if wound else "#050508"
    parts.append(f'<rect x="{x}" y="{y}" width="620" height="205" rx="34" fill="{fill}" fill-opacity="0.78" stroke="{color}" stroke-opacity="0.47"{dash} filter="url(#softGlow)"/>')
    parts.append(f'<text class="mathSmall" x="{x+42}" y="{y+58}" fill="{color}">{esc(title)}</text>')
    parts.append(f'<text class="caption" x="{x+42}" y="{y+108}">{esc(body)}</text>')
    parts.append(f'<text class="tiny" x="{x+42}" y="{y+164}">{esc(status)}</text>')


def render_gravity(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p = PALETTE
    q = c.quantities
    parts = svg_open(c.title, "Gravity metric projection theorem, conditional low-energy action, and source-type firewalls for M_P and Lambda.", style)

    # Theorem-level projection
    boxed(parts, 130, 390, 1540, 570, "THEOREM-LEVEL FLAT METRIC PROJECTION", "phase-space measurement carrier projects to ordinary Lorentzian event metric", p.platinum, 0.34)
    draw_phase_octave(parts, 190, 500, p)
    parts.append(f'<path d="M 820 695 C 945 650 1005 650 1115 695" stroke="url(#metricBridge)" stroke-width="8" fill="none" marker-end="url(#arrow)" filter="url(#softGlow)"/>')
    parts.append(f'<text class="mathSmall" x="965" y="638" text-anchor="middle" fill="{p.platinum}">Pi_X</text>')
    draw_minkowski_target(parts, 1130, 510, p)
    parts.append(f'<text class="mathSmall" x="900" y="900" text-anchor="middle">(V8, eta_1,7, Omega, Pi_X) -> (X4, eta_1,3)</text>')

    # Conditional gravitational action
    boxed(parts, 130, 1050, 1540, 610, "LOW-ENERGY GRAVITY LANE", "conditional on smooth Lorentzian metric and standard low-energy assumptions", p.gold, 0.40)
    draw_curved_grid(parts, 230, 1170, 740, 340, p)
    parts.append(f'<rect x="1040" y="1190" width="540" height="305" rx="38" fill="#050508" fill-opacity="0.68" stroke="{p.gold}" stroke-opacity="0.34"/>')
    parts.append(f'<text class="math" x="1310" y="1272" text-anchor="middle">S_grav</text>')
    parts.append(f'<text class="mathSmall" x="1310" y="1340" text-anchor="middle">= (M_P^2 / 2) integral</text>')
    parts.append(f'<text class="mathSmall" x="1310" y="1390" text-anchor="middle">(R - 2 Lambda) sqrt(-g) d^4x</text>')
    parts.append(f'<text class="tiny" x="1310" y="1460" text-anchor="middle">Einstein-Hilbert form is admitted after low-energy assumptions</text>')
    parts.append(f'<path d="M 900 1340 C 990 1320 1005 1320 1040 1340" stroke="{p.gold}" stroke-opacity="0.48" stroke-width="5" fill="none" marker-end="url(#arrow)"/>')

    # Firewall cards
    boxed(parts, 130, 1750, 1540, 360, "SOURCE-TYPE FIREWALLS", "strong statement: projection theorem exists; Planck stiffness and vacuum residual remain wounds", p.warm, 0.35)
    parts.append(f'<rect x="190" y="1844" width="1420" height="86" rx="30" fill="url(#firewall)" stroke="{p.obsidian}" stroke-opacity="0.95" stroke-width="10"/>')
    parts.append(f'<text class="caption" x="900" y="1898" text-anchor="middle">OBSIDIAN FIREWALL: metric projection != native M_P theorem != native Lambda theorem</text>')
    source_card(parts, 210, 1980, "M_P^2", q["source_types"]["M_P^2"], "physical filling / Planck-stiffness seal", p.gold, False)
    source_card(parts, 970, 1980, "Lambda", q["source_types"]["Lambda"], "unresolved wound / vacuum-boundary residual", p.warm, True)

    parts.append(f'<text class="tiny" x="1580" y="2255" text-anchor="end">EXACT CHECK: ACTIVE SIGNATURE (1,7) · PROJECTED SIGNATURE (1,3) · EINSTEIN-HILBERT FORM SOURCE-TYPED</text>')
    parts.append(f'<text class="tiny" x="1580" y="2288" text-anchor="end">BOUNDARY: NO NATIVE M_P^2 DERIVATION · NO NATIVE COSMOLOGICAL CONSTANT DERIVATION</text>')
    parts.append('</svg>')
    Path(out_path).write_text("".join(parts), encoding="utf-8")
