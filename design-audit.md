# ShareChat Revamp — Design System Audit

## CSS Custom Properties (Design Tokens)

All tokens are defined in `index.html:9–28` via `:root`.

### Color Tokens

| Token | Value | Usage |
|---|---|---|
| `--bg-primary` | `#0a0a0a` | Page background, phone frame bg |
| `--bg-secondary` | `#111111` | Defined but unused directly |
| `--bg-card` | `#1a1a1a` | Card backgrounds, story avatars |
| `--bg-glass` | `rgba(255,255,255,0.06)` | Glass surfaces (nav, buttons, chips) |
| `--border-glass` | `rgba(255,255,255,0.08)` | Glass borders throughout |
| `--text-primary` | `#ffffff` | Body text |
| `--text-secondary` | `rgba(255,255,255,0.6)` | Subtitles, labels, metadata |
| `--text-tertiary` | `rgba(255,255,255,0.35)` | Timestamps, muted labels |
| `--accent` | `#FF6B35` | Brand orange — primary interactive color |
| `--accent-glow` | `rgba(255,107,53,0.3)` | Glow shadows on accent elements |

> **Note:** `feed.html` hardcodes `#FF6B35` and `#0a0a0a` directly (does not import `:root` vars). These are identical values but not token-referenced — a gap to close.

### Gradient Tokens

| Token | Value | Purpose |
|---|---|---|
| `--gradient-1` | `linear-gradient(135deg, #FF6B35, #FF3CAC)` | Primary brand gradient (buttons, story rings, stat values, CTAs) |
| `--gradient-2` | `linear-gradient(135deg, #FF3CAC, #784BA0, #2B86C5)` | Secondary — avatars, before/after accents |
| `--gradient-3` | `linear-gradient(135deg, #00F5A0, #00D9F5)` | Teal/green accent — story rings, discover grid |
| `--gradient-gold` | `linear-gradient(135deg, #F7971E, #FFD200)` | Gold — story rings, awards |

Additional one-off gradients used inline in discover grid tiles:
- `#667eea → #764ba2`, `#f093fb → #f5576c`, `#4facfe → #00f2fe`, `#43e97b → #38f9d7`, `#fa709a → #fee140`

### Border Radius Tokens

| Token | Value | Used On |
|---|---|---|
| `--radius-sm` | `12px` | Feature icons, sidebar items |
| `--radius-md` | `16px` | Post media, search bars, reel containers |
| `--radius-lg` | `24px` | Feature cards, comparison table, BA cards |
| `--radius-xl` | `32px` | CTA card |
| `100px` (hardcoded) | `100px` | All pill buttons, chips, follow badges, tags |
| `40px` (hardcoded) | `40px` | Phone frame outer border |
| `50%` | `50%` | Avatars, icon circles, dots |

### Transition Token

```css
--transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94)
```

Used on cards, buttons, nav links. `feed.html` also uses `transition: all 0.2s` and `transition: all 0.35s cubic-bezier(0.32, 0.72, 0, 1)` (panel open) inline.

---

## Typography

### Font Families

| Family | Weights Loaded | Role |
|---|---|---|
| **Inter** | 300, 400, 500, 600, 700, 800, 900 | Body copy, UI labels, buttons |
| **Space Grotesk** | 400, 500, 600, 700 | Display headings, logo, stat values |

`feed.html` loads only Inter (weights 400–800). Space Grotesk not loaded there.

### Type Scale (index.html)

| Element | Size | Weight | Font | Notes |
|---|---|---|---|---|
| Hero `h1` | `clamp(48px, 7vw, 88px)` | 700 | Space Grotesk | -0.03em tracking, 1.05 line-height |
| Section title | `clamp(32px, 4vw, 52px)` | 700 | Space Grotesk | -0.02em tracking |
| CTA card `h2` | `clamp(28px, 4vw, 44px)` | 700 | Space Grotesk | |
| Impact value | `56px` | 700 | Space Grotesk | Gradient text, line-height 1 |
| Stat value | `36px` | 700 | Space Grotesk | Gradient text |
| Feature title | `20px` | 600 | Space Grotesk | |
| Body (hero `p`) | `18px` | 400 | Inter | 1.6 line-height |
| Body (section) | `16px` | 400 | Inter | 1.6 line-height |
| Body (cards) | `14px` | 400 | Inter | 1.7 line-height |
| Post name | `13px` | 600 | Inter | |
| Section tag | `12px` | 600 | Inter | uppercase, 0.15em tracking |
| Stat label | `13px` | 400 | Inter | uppercase, 0.1em tracking |
| Chips / tags | `12px` | 500 | Inter | |
| Timestamps | `11px` | 400 | Inter | `--text-tertiary` |
| Micro labels | `10px` | 400 | Inter | Story names, discover views |

---

## Spacing System

No spacing tokens defined. Key values observed:

| Context | Value |
|---|---|
| Section padding (all major sections) | `120px 24px` |
| Section max-widths | `1200px` (features, phone), `1000px` (comparison, B/A) |
| Section header bottom margin | `80px` |
| Feature card padding | `40px 32px` |
| Features grid gap | `24px` |
| Nav padding | `16px 32px` |
| Nav link gap | `32px` |
| Hero CTA gap | `16px` |
| Hero stat gap | `48px` |
| Impact metric gap | `64px` |
| Phone showcase gap | `40px` |
| Stories gap | `12px` |
| Post padding bottom | `20px` |
| Feed icon gap | `16px` |
| Reel sidebar action gap | `20px` |

---

## Button Styles

### `.btn-primary`
```css
padding: 16px 36px;
border-radius: 100px;
background: var(--gradient-1);
color: white;
font-size: 15px;
font-weight: 600;
box-shadow: 0 8px 32px rgba(255,107,53,0.25);
/* hover: translateY(-2px), box-shadow 0 12px 40px rgba(255,107,53,0.35) */
```

### `.btn-secondary`
```css
padding: 16px 36px;
border-radius: 100px;
border: 1px solid var(--border-glass);
background: var(--bg-glass);
color: white;
font-size: 15px;
font-weight: 500;
backdrop-filter: blur(20px);
/* hover: rgba(255,255,255,0.1) bg, translateY(-2px) */
```

### `.btn-white`
```css
padding: 16px 40px;
border-radius: 100px;
background: white;
color: #0a0a0a;
font-size: 15px;
font-weight: 600;
/* hover: translateY(-2px), box-shadow 0 12px 40px rgba(0,0,0,0.3) */
```

### `.post-follow` / follow buttons
```css
font-size: 12px;
font-weight: 600;
color: var(--accent);
border: 1px solid var(--accent);
padding: 4px 14px;
border-radius: 100px;
background: transparent;
```

### `feed.html` panel submit / share buttons
```css
background: #FF6B35;
border-radius: 100px;
padding: 8px 18px;
color: #fff;
font-size: 12px;
font-weight: 600;
```

---

## Nav Structure

### Landing Page Nav (`index.html:902–915`)
```
[Logo: S icon + "ShareChat"] ← left
[Vision | Experience | Features | Impact] ← center links
[btn-primary "Let's Build This"] ← right
```
- Fixed top, `backdrop-filter: blur(20px)`, `background: rgba(10,10,10,0.8)`
- Mobile (`≤768px`): `.nav-links` hidden, logo + CTA only

### App Sidebar Nav (`feed.html`)
**Left sidebar (72px, always visible desktop):**
```
[S] logo (accent color)
[🏠] Home (active — accent left bar indicator)
[🔍] Search
[🎬] Reels
[💬] Chat + badge dot
[🔔] Notifications + badge dot
[➕] Create
— spacer —
[📈] Analytics
[⚙️] Settings
[V] Avatar (gradient circle)
```

**Right sidebar (280px, visible at ≥1100px):**
```
Profile row (avatar + name + handle + Switch)
"Suggested for you" → See All
  [suggestions list]
"Trending Tags" → More
  [tag list]
Footer (About · Help · …)
```

**Mobile (≤768px):** Both sidebars hidden, nav-arrows hidden. Full-width feed.

**Nav items active state:**
```css
.sidebar-item.active { background: rgba(255,107,53,0.1); }
.sidebar-item.active::before {
  /* 3px left accent bar */
  background: #FF6B35;
  border-radius: 0 3px 3px 0;
}
```

---

## Layout Patterns

### Glassmorphism Card
```css
background: var(--bg-glass);        /* rgba(255,255,255,0.06) */
border: 1px solid var(--border-glass); /* rgba(255,255,255,0.08) */
backdrop-filter: blur(20px);
border-radius: var(--radius-lg);    /* 24px */
```
Used on: nav, hero-badge, btn-secondary, feature cards, comparison header, discover search

### Gradient Text
```css
background: var(--gradient-1);
-webkit-background-clip: text;
-webkit-text-fill-color: transparent;
background-clip: text;
```
Used on: hero h1 accent word, section titles, stat/impact values

### Scroll-Snap Feed
```css
/* Container */
scroll-snap-type: y mandatory;
overflow-y: scroll;
scrollbar-width: none;

/* Items */
scroll-snap-align: start;
scroll-snap-stop: always;
height: 100vh; /* or 100dvh */
```

### Phone Frame
```css
width: 320px; height: 640px;
border-radius: 40px;
background: #0a0a0a;
border: 2px solid rgba(255,255,255,0.1);
box-shadow: 0 40px 80px rgba(0,0,0,0.5), 0 0 120px rgba(255,107,53,0.05);
perspective: 1200px; /* on parent */
transform: rotateY(-5deg); /* left/right tilt */
```

### Story Ring
```css
width: 56px; height: 56px;
border-radius: 50%;
padding: 2px;
background: var(--gradient-1); /* colored ring */
/* inner avatar has border: 2px solid var(--bg-primary) to create gap */
```

### Modal / Bottom Sheet (feed.html)
```css
.panel-overlay {
  position: fixed; inset: 0; z-index: 500;
  background: rgba(0,0,0,0.6);
  opacity: 0; visibility: hidden;
  transition: all 0.3s;
}
.panel-box {
  background: #1a1a1a;
  border-radius: 16px;
  width: 420px; max-height: 70vh;
  transform: translateY(30px) scale(0.97); /* enter animation */
}
/* Mobile: width 100%, border-radius 20px 20px 0 0, position absolute bottom:0 */
```

### Toast Notification
```css
position: fixed; bottom: 40px; left: 50%;
transform: translateX(-50%) translateY(16px);
background: rgba(255,255,255,0.95);
color: #000;
padding: 10px 28px;
border-radius: 100px;
font-size: 13px; font-weight: 600;
/* .show state: opacity 1, translateY(0) */
```

---

## Reusable CSS Classes

| Class | Location | Description |
|---|---|---|
| `.gradient-text` | index.html | Gradient clip text utility |
| `.glow-border` | index.html | Pseudo-element gradient border on hover (via mask) |
| `.floating` | index.html | `float` keyframe animation (6s infinite) |
| `.section-tag` | index.html | Accent-colored uppercase category label (12px, 600, 0.15em) |
| `.section-title` | index.html | Space Grotesk display heading with fluid sizing |
| `.section-desc` | index.html | Muted body copy, max-width 480px, centered |
| `.hero-badge` | index.html | Glass pill with accent dot + text |
| `.chip` / `.chip.active` | index.html | Horizontal filter tab — glass default, gradient-1 when active |
| `.story-ring` | index.html | Gradient-bordered story avatar |
| `.feature-card` | index.html | Glass card with top gradient line on hover |
| `.comparison-cell .check` | index.html | `#22c55e` green checkmark |
| `.comparison-cell .cross` | index.html | `rgba(255,255,255,0.15)` dim cross |
| `.act` / `.act.liked` | feed.html | Action button with bounce animation on interact |
| `.live-comment` | feed.html | Pill-shaped animated comment bubble |
| `.reel-progress-fill` | feed.html | Orange video progress bar |
| `.panel-overlay.open` | feed.html | Modal open state |
| `.sidebar-item.active` | feed.html | Active nav item with left accent bar |

---

## Animation Inventory

| Name | Duration | Easing | Trigger | Effect |
|---|---|---|---|---|
| `fadeUp` | 0.8s | ease | page load (stagger 0–0.5s) | opacity 0→1, translateY 20px→0 |
| `pulse` | 2s | ease | continuous | opacity 1→0.4→1 (status dot) |
| `spin` | 3s / 0.8s | linear | continuous / load | 360° rotation (music note / spinner) |
| `progress` | 8s | linear | continuous | width 0→100% (reel progress bar) |
| `float` | 6s | ease-in-out | continuous | translateY ±10px (phone frames) |
| `shimmer` | — | — | defined, unused | background-position shift |
| `lc` | 4s | ease | interval (2.5s) | live comment: fade in, hold, fade+rise out |
| `hb` | 0.8s | ease | double-tap | heart burst: scale 0.2→1.1→0.95, fade out |
| `floatUp` | 1.3–1.8s | ease-out | like action | 8 hearts float upward with rotate + fade |
| `actB` | 0.4s | ease | button press | spring scale: 1→1.35→0.9→1 |

---

## Responsive Breakpoints

| Breakpoint | Behavior |
|---|---|
| `max-width: 768px` | Phone showcase stacks vertically; nav links hidden; comparison columns tighten; BA cards stack; phone frames lose 3D tilt; feed goes full-width borderless; sidebars + nav arrows hidden; panel becomes bottom sheet |
| `min-width: 1100px` | Right sidebar visible in feed.html |

---

## Gaps & Inconsistencies to Resolve

1. **Token adoption in feed.html** — `#0a0a0a`, `#1a1a1a`, `#FF6B35`, etc. are hardcoded. Should import and use CSS vars.
2. **Spacing system missing** — No `--space-*` tokens; all spacing is ad-hoc px values. Consider an 8px base grid.
3. **No shared CSS file** — All styles are inline `<style>` blocks. Both files duplicate base reset, scrollbar, and some component styles.
4. **Space Grotesk not loaded in feed.html** — Logo "S" and any display text there falls back to system-ui.
5. **`--bg-secondary` unused** — Token defined, never referenced.
6. **`shimmer` animation defined, never applied** — Likely intended for skeleton states.
7. **Mixed border-radius conventions** — Some elements use token (`var(--radius-md)`), others hardcode (`16px`, `12px`, `100px`).
8. **Discover grid tile gradients are all inline styles** — No class abstraction for the gradient tile pattern.
