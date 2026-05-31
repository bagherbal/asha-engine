from __future__ import annotations

from pathlib import Path
from math import log10
import html

from .style import PALETTE, STYLE, FigureStyle
from .readme_quarks import FigureContract, L, S_SPLIT


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
  <linearGradient id="goldLine" x1="0" x2="1"><stop offset="0%" stop-color="#7D6434" stop-opacity="0.20"/><stop offset="50%" stop-color="{p.gold}" stop-opacity="0.96"/><stop offset="100%" stop-color="#FFF2B8" stop-opacity="0.43"/></linearGradient>
  <linearGradient id="cyanLine" x1="0" x2="1"><stop offset="0%" stop-color="#114B5B" stop-opacity="0.20"/><stop offset="52%" stop-color="{p.cyan}" stop-opacity="0.92"/><stop offset="100%" stop-color="#D9FDFF" stop-opacity="0.42"/></linearGradient>
  <linearGradient id="warmLine" x1="0" x2="1"><stop offset="0%" stop-color="#552015" stop-opacity="0.20"/><stop offset="52%" stop-color="{p.warm}" stop-opacity="0.92"/><stop offset="100%" stop-color="#FFD5C4" stop-opacity="0.42"/></linearGradient>
  <filter id="softGlow" x="-180%" y="-180%" width="460%" height="460%"><feGaussianBlur stdDeviation="5" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <filter id="bigGlow" x="-220%" y="-220%" width="540%" height="540%"><feGaussianBlur stdDeviation="15" result="blur"/><feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge></filter>
  <marker id="arrow" markerWidth="10" markerHeight="10" refX="7" refY="3" orient="auto"><path d="M0,0 L0,6 L8,3 z" fill="{p.platinum}" fill-opacity="0.7"/></marker>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 62px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 34px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 25px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 2.5px; fill: {p.muted}; }}
    .tiny {{ font-family: {style.label_font}; font-size: 14px; letter-spacing: 1.4px; fill: #8B93A5; }}
    .caption {{ font-family: {style.label_font}; font-size: 19px; fill: #AAB3C2; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{w-96}" height="{h-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{w/2:.0f}" y="132" text-anchor="middle">ASHA README · QUARK SECTOR · BOUNDARY LANE STACK</text>
<text class="title" x="{w/2:.0f}" y="214" text-anchor="middle">{esc(title)}</text>
<text class="mathSmall" x="{w/2:.0f}" y="278" text-anchor="middle">top scalar-aligned · bottom contact-depth · shape laws transported at MZ</text>
''']


def svg_close(footer: str, check: str, style: FigureStyle = STYLE) -> str:
    p = PALETTE
    w = style.width
    return f'''
<text class="tiny" x="{w-220}" y="2240" text-anchor="end">{esc(check)}</text>
<text class="tiny" x="{w-220}" y="2270" text-anchor="end">{esc(footer)}</text>
</svg>'''


def boxed(parts: list[str], x: float, y: float, width: float, height: float, title: str, subtitle: str, stroke: str, fill_opacity: float = 0.50) -> None:
    parts.append(f'<rect x="{x}" y="{y}" width="{width}" height="{height}" rx="34" fill="#071018" fill-opacity="{fill_opacity}" stroke="{stroke}" stroke-opacity="0.38"/>')
    parts.append(f'<text class="label" x="{x+34}" y="{y+50}">{esc(title)}</text>')
    parts.append(f'<text class="tiny" x="{x+34}" y="{y+82}">{esc(subtitle)}</text>')


def mass_ladder(parts: list[str], masses: dict[str, float], x0: float, x1: float, y: float, title: str, color: str, marker: str) -> None:
    vals = {k: log10(v) for k, v in masses.items()}
    mn, mx = min(vals.values()), max(vals.values())
    parts.append(f'<text class="label" x="{x0}" y="{y-105}">{esc(title)}</text>')
    parts.append(f'<line x1="{x0}" y1="{y}" x2="{x1}" y2="{y}" stroke="{color}" stroke-opacity="0.30" stroke-width="3"/>')
    for name, val in masses.items():
        t = (log10(val) - mn) / (mx - mn)
        x = x0 + t * (x1 - x0)
        r = 24 if name.endswith("t") or name.endswith("b") else 17
        parts.append(f'<circle cx="{x:.1f}" cy="{y}" r="{r}" fill="{color}" fill-opacity="0.26" stroke="{color}" stroke-opacity="0.90" filter="url(#softGlow)"/>')
        parts.append(f'<text class="mathSmall" x="{x:.1f}" y="{y-40}" text-anchor="middle" fill="{color}">{esc(name)}</text>')
        parts.append(f'<text class="tiny" x="{x:.1f}" y="{y+55}" text-anchor="middle">{val:.6g} GeV</text>')
    parts.append(f'<text class="tiny" x="{x1}" y="{y+98}" text-anchor="end">log10 mass ladder · {esc(marker)}</text>')


def render_quarks(c: FigureContract, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    p = PALETTE
    q = c.quantities
    top = q["top_lane"]
    bottom = q["bottom_lane"]
    up = q["up_sector_shapes"]
    down = q["down_sector_shapes"]
    parts = svg_open(c.title, "Top and bottom quark anchor lanes plus exact up/down sector shape laws from the README.", style)

    # Central firewall spine: this is a boundary/RG transport board, not native derivation.
    parts.append(f'<line x1="900" y1="340" x2="900" y2="2070" stroke="{p.obsidian}" stroke-width="18" stroke-opacity="0.72"/>')
    parts.append(f'<line x1="900" y1="340" x2="900" y2="2070" stroke="{p.platinum}" stroke-width="1.2" stroke-opacity="0.12"/>')
    parts.append(f'<text class="tiny" x="930" y="2030" transform="rotate(-90 930 2030)">RG TRANSPORT FIREWALL · MZ BOUNDARY VALUES</text>')

    # Left: up/top lane. Right: down/bottom lane.
    boxed(parts, 150, 405, 690, 520, "UP SECTOR · TOP ANCHOR", "scalar-aligned lane; no universal depth penalty", p.cyan)
    boxed(parts, 960, 405, 690, 520, "DOWN SECTOR · BOTTOM ANCHOR", "contact-depth lane; boundary value before RG flow", p.gold)

    # Formula rings: action anchor -> Yukawa -> mass.
    for cx, cy, col, label, A, yval, mass, formula, source in [
        (360, 650, p.cyan, "top", top["A_t"], top["y_t"], top["m_t_GeV"], top["formula"], top["coefficient_source"]),
        (1170, 650, p.gold, "bottom", bottom["A_b"], bottom["y_b"], bottom["m_b_GeV"], bottom["formula"], bottom["coefficient_source"]),
    ]:
        parts.append(f'<circle cx="{cx}" cy="{cy}" r="115" fill="{col}" fill-opacity="0.09" stroke="{col}" stroke-opacity="0.62" filter="url(#bigGlow)"/>')
        parts.append(f'<text class="math" x="{cx}" y="{cy-42}" text-anchor="middle" fill="{col}">A_{label[0]}</text>')
        parts.append(f'<text class="mathSmall" x="{cx}" y="{cy+4}" text-anchor="middle">{A:.9f}</text>')
        parts.append(f'<text class="tiny" x="{cx}" y="{cy+46}" text-anchor="middle">{esc(formula)}</text>')
        parts.append(f'<path d="M {cx+120} {cy} C {cx+210} {cy-80} {cx+290} {cy-80} {cx+380} {cy}" stroke="{col}" stroke-opacity="0.34" stroke-width="3" fill="none" marker-end="url(#arrow)"/>')
        parts.append(f'<circle cx="{cx+455}" cy="{cy}" r="72" fill="#050508" stroke="{col}" stroke-opacity="0.55" filter="url(#softGlow)"/>')
        parts.append(f'<text class="mathSmall" x="{cx+455}" y="{cy-12}" text-anchor="middle">y={yval:.6f}</text>')
        parts.append(f'<text class="tiny" x="{cx+455}" y="{cy+28}" text-anchor="middle">m={mass:.6g} GeV</text>')
        parts.append(f'<text class="tiny" x="{cx}" y="{cy+173}" text-anchor="middle" fill="{col}">{esc(source)}</text>')

    # Correct a too-wide top lane arrow visual by masking at center firewall and drawing local formula paths.
    parts.append(f'<rect x="850" y="462" width="100" height="408" fill="{p.abyss}" fill-opacity="0.92"/>')
    parts.append(f'<line x1="900" y1="340" x2="900" y2="2070" stroke="{p.obsidian}" stroke-width="18" stroke-opacity="0.72"/>')
    parts.append(f'<line x1="900" y1="340" x2="900" y2="2070" stroke="{p.platinum}" stroke-width="1.2" stroke-opacity="0.12"/>')

    # Shape-law panels.
    boxed(parts, 150, 1010, 690, 430, "UP SHAPE LAWS", "K_u and D_u fix relative hierarchy after top anchor", p.cyan, 0.42)
    parts.append(f'<text class="math" x="495" y="1160" text-anchor="middle">Ku = 8/9 − S + 4S²</text>')
    parts.append(f'<text class="mathSmall" x="495" y="1215" text-anchor="middle" fill="{p.cyan}">{up["K_u"]:.12f}</text>')
    parts.append(f'<text class="math" x="495" y="1305" text-anchor="middle">Du = √(2π)/4 − 4S + 4S²</text>')
    parts.append(f'<text class="mathSmall" x="495" y="1360" text-anchor="middle" fill="{p.cyan}">{up["D_u"]:.12f}</text>')

    boxed(parts, 960, 1010, 690, 430, "DOWN SHAPE LAWS", "K_d and D_d fix relative hierarchy after bottom anchor", p.gold, 0.42)
    parts.append(f'<text class="math" x="1305" y="1160" text-anchor="middle">Kd = 3/4 − (12/7)S</text>')
    parts.append(f'<text class="mathSmall" x="1305" y="1215" text-anchor="middle" fill="{p.gold}">{down["K_d"]:.12f}</text>')
    parts.append(f'<text class="math" x="1305" y="1305" text-anchor="middle">Dd = −1 + 1/72 − 4S²</text>')
    parts.append(f'<text class="mathSmall" x="1305" y="1360" text-anchor="middle" fill="{p.gold}">{down["D_d"]:.12f}</text>')

    mass_ladder(parts, up["solved_masses_GeV"], 185, 810, 1660, "UP-TYPE MASS LADDER", p.cyan, "m_u < m_c < m_t")
    mass_ladder(parts, down["solved_masses_GeV"], 990, 1615, 1660, "DOWN-TYPE MASS LADDER", p.gold, "m_d < m_s < m_b")

    # Shared constants and status chips.
    parts.append(f'<rect x="310" y="1880" width="1180" height="190" rx="44" fill="#050508" fill-opacity="0.72" stroke="{p.platinum}" stroke-opacity="0.20"/>')
    parts.append(f'<text class="mathSmall" x="900" y="1948" text-anchor="middle">L = 1/(8π) = {L:.12f}     ·     S = {S_SPLIT:.16f}</text>')
    parts.append(f'<text class="caption" x="900" y="2015" text-anchor="middle">Locked as MZ-boundary formulas with RG transport; not promoted to a native quark theorem.</text>')

    parts.append(svg_close("BOUNDARY: QUARK FORMULAS REQUIRE RG TRANSPORT AND SOURCE-TYPED COEFFICIENTS.", "EXACT CHECK: At, Ab, Ku, Du, Kd, Dd MATCH README FORMULAS", style))
    Path(out_path).write_text("".join(parts), encoding="utf-8")
