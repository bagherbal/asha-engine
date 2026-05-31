from __future__ import annotations

import html
from pathlib import Path
from math import log10
from .style import PALETTE, STYLE, FigureStyle
from .contact_depth import ContactDepthGeometry


def esc(s: str) -> str:
    return html.escape(s, quote=True)


def render_contact_depth_svg(geometry: ContactDepthGeometry, out_path: str | Path,
                             style: FigureStyle = STYLE) -> None:
    out_path = Path(out_path)
    p = PALETTE
    width, height = style.width, style.height

    def color_for(role: str) -> str:
        return {
            "time": p.warm,
            "energy": p.gold,
            "space_event": p.cyan,
            "space_response": p.glass,
        }.get(role, p.platinum)

    # Log-scaled visual bar widths for W_Q values. They are tiny on a linear scale,
    # so the plot explicitly labels the logarithmic measurement.
    weights = geometry.w_q
    logs = [-log10(w) for w in weights]
    min_log, max_log = min(logs), max(logs)
    def bar_width(w: float) -> float:
        lv = -log10(w)
        return 90 + (lv - min_log) / (max_log - min_log) * 215

    parts: list[str] = []
    parts.append(f'''<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}" role="img" aria-labelledby="title desc">
<title id="title">ASHA README Figure — Contact Seven and Depth Triple</title>
<desc id="desc">Visual theorem map for V8 choosing x0 as observer reference, leaving V7_contact = R p0 plus three phase planes, then Q_contact^3 with N_Q diagonal entries 1/3, 1/2, 2/3 and W_Q = exp(-4 pi N_Q).</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="42%" r="75%">
    <stop offset="0%" stop-color="#12121F"/>
    <stop offset="62%" stop-color="{p.abyss}"/>
    <stop offset="100%" stop-color="#010103"/>
  </radialGradient>
  <linearGradient id="glassStroke" x1="0%" x2="100%">
    <stop offset="0%" stop-color="{p.cyan}" stop-opacity="0.22"/>
    <stop offset="50%" stop-color="{p.platinum}" stop-opacity="0.65"/>
    <stop offset="100%" stop-color="{p.gold}" stop-opacity="0.22"/>
  </linearGradient>
  <linearGradient id="goldBar" x1="0%" x2="100%">
    <stop offset="0%" stop-color="#7B5F21" stop-opacity="0.4"/>
    <stop offset="100%" stop-color="#FFE7A3" stop-opacity="0.95"/>
  </linearGradient>
  <filter id="softGlow" x="-200%" y="-200%" width="500%" height="500%">
    <feGaussianBlur stdDeviation="5" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <filter id="bigGlow" x="-250%" y="-250%" width="600%" height="600%">
    <feGaussianBlur stdDeviation="16" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 62px; letter-spacing: 0.4px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 19px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 34px; fill: {p.platinum}; }}
    .mathSmall {{ font-family: {style.title_font}; font-size: 27px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 18px; letter-spacing: 2.8px; fill: {p.muted}; text-transform: uppercase; }}
    .caption {{ font-family: {style.label_font}; font-size: 22px; fill: #AAB3C2; }}
    .tiny {{ font-family: {style.label_font}; font-size: 15px; letter-spacing: 1.4px; fill: #8B93A5; }}
    .number {{ font-family: {style.title_font}; font-size: 30px; fill: {p.gold}; }}
  ]]></style>
  <marker id="arrow" markerWidth="12" markerHeight="12" refX="10" refY="6" orient="auto" markerUnits="strokeWidth">
    <path d="M2,2 L10,6 L2,10" fill="none" stroke="{p.platinum}" stroke-opacity="0.55" stroke-width="1.5"/>
  </marker>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{width-96}" height="{height-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
<text class="subtitle" x="{width/2:.1f}" y="132" text-anchor="middle">ASHA README VISUAL ATLAS · NEXT REQUIRED FIGURE</text>
<text class="title" x="{width/2:.1f}" y="214" text-anchor="middle">Contact Seven and the Depth Triple</text>
<text class="math" x="{width/2:.1f}" y="282" text-anchor="middle">V8 → choose x0 → V7_contact → Q_contact^3</text>
''')

    # Column frames.
    frames = [
        (105, 390, 420, 1480, "1 · MEASUREMENT OCTAVE", "V8 = X4 ⊕ P4"),
        (620, 390, 540, 1480, "2 · CONTACT CARRIER", "V7 = R p0 ⊕ Π1 ⊕ Π2 ⊕ Π3"),
        (1260, 390, 435, 1480, "3 · DEPTH OPERATOR", "N_Q and W_Q"),
    ]
    for x, y, w, h, label, math in frames:
        parts.append(f'<rect x="{x}" y="{y}" width="{w}" height="{h}" rx="34" fill="#090912" fill-opacity="0.62" stroke="url(#glassStroke)" stroke-opacity="0.28" stroke-width="1.1"/>\n')
        parts.append(f'<text class="label" x="{x+34}" y="{y+58}">{esc(label)}</text>\n')
        parts.append(f'<text class="mathSmall" x="{x+34}" y="{y+102}">{esc(math)}</text>\n')

    # Basis phase-pair rails in V8.
    parts.append('<g id="v8_basis">\n')
    for i in range(4):
        a = geometry.basis_points[2*i]
        b = geometry.basis_points[2*i+1]
        stroke = p.warm if i == 0 else p.cyan
        parts.append(f'<line x1="{a.x:.1f}" y1="{a.y:.1f}" x2="{b.x:.1f}" y2="{b.y:.1f}" stroke="{stroke}" stroke-opacity="0.30" stroke-width="2.2"/>\n')
        parts.append(f'<text class="tiny" x="{(a.x+b.x)/2:.1f}" y="{a.y-45:.1f}" text-anchor="middle">PAIR {i}</text>\n')
    for point in geometry.basis_points:
        color = color_for(point.color_role)
        opacity = 1.0 if point.in_contact else 0.95
        r = 17 if point.key == "x0" else 14
        parts.append(f'<circle cx="{point.x:.1f}" cy="{point.y:.1f}" r="{r+12}" fill="{color}" opacity="0.115" filter="url(#bigGlow)"/>\n')
        parts.append(f'<circle cx="{point.x:.1f}" cy="{point.y:.1f}" r="{r}" fill="{color}" fill-opacity="{opacity}" stroke="{p.platinum}" stroke-opacity="0.55" stroke-width="0.8" filter="url(#softGlow)"/>\n')
        parts.append(f'<text class="mathSmall" x="{point.x:.1f}" y="{point.y+47:.1f}" text-anchor="middle" fill="{color}">{esc(point.label)}</text>\n')
    # Obsidian seal around x0 as selected observer reference.
    x0 = next(bp for bp in geometry.basis_points if bp.key == "x0")
    parts.append(f'<circle cx="{x0.x:.1f}" cy="{x0.y:.1f}" r="42" fill="none" stroke="#050508" stroke-opacity="0.92" stroke-width="9"/>\n')
    parts.append(f'<text class="tiny" x="{x0.x:.1f}" y="{x0.y-62:.1f}" text-anchor="middle" fill="{p.warm}">SELECTED REFERENCE</text>\n')
    parts.append('</g>\n')

    # Arrow from V8 to contact carrier.
    parts.append(f'<path d="M 535 1130 C 570 1130 585 1130 610 1130" fill="none" stroke="{p.platinum}" stroke-opacity="0.48" stroke-width="2.2" marker-end="url(#arrow)"/>\n')
    parts.append(f'<text class="tiny" x="572" y="1088" text-anchor="middle">REMOVE x0 FROM CARRIER</text>\n')

    # Contact carrier central p0 and phase planes.
    p0x, p0y = 895, 660
    parts.append(f'<circle cx="{p0x}" cy="{p0y}" r="32" fill="{p.gold}" fill-opacity="0.94" stroke="{p.platinum}" stroke-opacity="0.75" stroke-width="1" filter="url(#softGlow)"/>\n')
    parts.append(f'<text class="math" x="{p0x}" y="{p0y+74}" text-anchor="middle" fill="{p.gold}">R p0</text>\n')
    parts.append(f'<text class="tiny" x="{p0x}" y="{p0y-58}" text-anchor="middle">ENERGY RESPONSE AXIS</text>\n')

    for plane in geometry.phase_planes:
        parts.append(f'<ellipse cx="{plane.cx:.1f}" cy="{plane.cy:.1f}" rx="{plane.rx:.1f}" ry="{plane.ry:.1f}" fill="#0B1519" fill-opacity="0.62" stroke="{p.cyan}" stroke-opacity="0.42" stroke-width="1.3"/>\n')
        parts.append(f'<ellipse cx="{plane.cx:.1f}" cy="{plane.cy:.1f}" rx="{plane.rx*0.78:.1f}" ry="{plane.ry*0.52:.1f}" fill="none" stroke="{p.glass}" stroke-opacity="0.23" stroke-width="1"/>\n')
        parts.append(f'<circle cx="{plane.cx-82:.1f}" cy="{plane.cy:.1f}" r="17" fill="{p.cyan}" filter="url(#softGlow)"/>\n')
        parts.append(f'<circle cx="{plane.cx+82:.1f}" cy="{plane.cy:.1f}" r="17" fill="{p.glass}" filter="url(#softGlow)"/>\n')
        xlab = plane.basis[0].replace("1", "¹").replace("2", "²").replace("3", "³")
        plab = plane.basis[1].replace("1", "₁").replace("2", "₂").replace("3", "₃")
        parts.append(f'<text class="mathSmall" x="{plane.cx-82:.1f}" y="{plane.cy+52:.1f}" text-anchor="middle">{esc(xlab)}</text>\n')
        parts.append(f'<text class="mathSmall" x="{plane.cx+82:.1f}" y="{plane.cy+52:.1f}" text-anchor="middle">{esc(plab)}</text>\n')
        parts.append(f'<text class="math" x="{plane.cx:.1f}" y="{plane.cy-20:.1f}" text-anchor="middle" fill="{p.platinum}">{esc(plane.label)}</text>\n')
        parts.append(f'<text class="tiny" x="{plane.cx:.1f}" y="{plane.cy+8:.1f}" text-anchor="middle">span({esc(plane.basis[0])},{esc(plane.basis[1])})</text>\n')
        parts.append(f'<path d="M {p0x} {p0y+38} C {p0x} {(p0y+plane.cy)/2:.1f} {plane.cx} {(p0y+plane.cy)/2:.1f} {plane.cx} {plane.cy-plane.ry:.1f}" fill="none" stroke="{p.gold}" stroke-opacity="0.13" stroke-width="1.2"/>\n')

    # Arrow from contact carrier to depth operator.
    parts.append(f'<path d="M 1168 1130 C 1200 1130 1225 1130 1252 1130" fill="none" stroke="{p.platinum}" stroke-opacity="0.48" stroke-width="2.2" marker-end="url(#arrow)"/>\n')
    parts.append(f'<text class="tiny" x="1210" y="1088" text-anchor="middle">CENTERED DEPTH RULE</text>\n')

    # Depth operator and logarithmic weights.
    dx = 1312
    dy0 = 630
    row_gap = 300
    for idx, plane in enumerate(geometry.phase_planes):
        y = dy0 + idx * row_gap
        frac = geometry.n_q[idx]
        wval = geometry.w_q[idx]
        bw = bar_width(wval)
        parts.append(f'<circle cx="{dx}" cy="{y}" r="32" fill="{p.gold}" fill-opacity="0.18" stroke="{p.gold}" stroke-opacity="0.72" stroke-width="1.2" filter="url(#softGlow)"/>\n')
        parts.append(f'<text class="mathSmall" x="{dx}" y="{y+9}" text-anchor="middle" fill="{p.gold}">{esc(plane.label)}</text>\n')
        parts.append(f'<text class="mathSmall" x="{dx+58}" y="{y-18}" fill="{p.platinum}">N = {esc(frac)}</text>\n')
        parts.append(f'<rect x="{dx+58}" y="{y+10}" width="{bw:.1f}" height="18" rx="9" fill="url(#goldBar)" opacity="0.92"/>\n')
        parts.append(f'<text class="tiny" x="{dx+58}" y="{y+58}" fill="{p.muted}">W = {wval:.9f}</text>\n')
        parts.append(f'<text class="tiny" x="{dx+58}" y="{y+84}" fill="{p.muted}">exp(-4π·{esc(frac)}) · log bar</text>\n')
    parts.append(f'<text class="math" x="1475" y="1545" text-anchor="middle" fill="{p.gold}">diag(1/3, 1/2, 2/3)</text>\n')
    parts.append(f'<text class="caption" x="1475" y="1592" text-anchor="middle">broadcast depth weights for matter sockets</text>\n')

    # Footer validation.
    parts.append(f'''<g transform="translate(165 2030)">
  <text class="label" x="0" y="0">SOURCE-TYPED READING</text>
  <text class="caption" x="0" y="48">Carrier/depth grammar: x0 is a reference choice; V7_contact is exactly seven directions;</text>
  <text class="caption" x="0" y="90">Q_contact^3 is the three spatial phase-plane depth triple.</text>
  <text class="caption" x="0" y="134">No physical flavor, mass eigenvalues, CKM/PMNS, or observed hierarchy is claimed here.</text>
</g>
<text class="tiny" x="{width-165}" y="2240" text-anchor="end">EXACT CHECK: 8 BASIS DIRECTIONS · x0 REMOVED · 7 CONTACT DIRECTIONS · 3 DEPTH LAYERS</text>
<text class="tiny" x="{width-165}" y="2270" text-anchor="end">N_Q = DIAG(1/3, 1/2, 2/3) · W_Q = EXP(-4π N_Q)</text>
</svg>''')

    out_path.write_text("".join(parts), encoding="utf-8")
