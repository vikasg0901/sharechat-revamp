# Contributing to ShareChat — Reimagined

This is a **static HTML design prototype** — two files, no build step, no package manager.
Contributions should stay within that constraint unless the scope is explicitly expanded.

---

## Table of Contents

1. [What this project is](#1-what-this-project-is)
2. [File structure](#2-file-structure)
3. [Running locally](#3-running-locally)
4. [Design tokens](#4-design-tokens)
5. [Coding conventions](#5-coding-conventions)
6. [Making a change](#6-making-a-change)
7. [Submitting a pull request](#7-submitting-a-pull-request)
8. [Out of scope](#8-out-of-scope)

---

## 1. What this project is

A high-fidelity browser prototype demonstrating a redesign of the ShareChat short-video feed.
It is intended for **design review and stakeholder sign-off**, not production deployment.
Expect inline styles, embedded scripts, and hard-coded sample data — all intentional for
a zero-dependency demo.

---

## 2. File structure

```
index.html   — Landing / marketing page
feed.html    — Interactive video feed (the core prototype)
CONTRIBUTING.md
```

All CSS lives in `<style>` blocks inside each file.
All JavaScript lives in `<script>` blocks at the bottom of each file.
There are no external asset files to maintain.

---

## 3. Running locally

Open either file directly in a browser:

```
open index.html
open feed.html
```

For auto-reload on save, use any static file server — for example:

```bash
# Python (built-in)
python -m http.server 8000

# Node (npx, no install required)
npx serve .
```

No install step. No build step. If it opens in a browser, it runs.

---

## 4. Design tokens

Tokens are defined as CSS custom properties in the `:root` block near the top of each file.
**Always use a token — never hardcode a raw value.**

### Color tokens (defined in both files)

| Token | Value | Role |
|---|---|---|
| `--bg-primary` | `#0a0a0a` | Page background |
| `--bg-secondary` | `#111111` | Elevated surfaces |
| `--bg-card` | `#1a1a1a` | Card backgrounds |
| `--bg-glass` | `rgba(255,255,255,0.06)` | Glassmorphism fills |
| `--border-glass` | `rgba(255,255,255,0.08)` | Glassmorphism strokes |
| `--text-primary` | `#ffffff` | Body copy, headings |
| `--text-secondary` | `rgba(255,255,255,0.6)` | Supporting labels |
| `--text-tertiary` | `rgba(255,255,255,0.35)` | Disabled / placeholder |
| `--accent` | `#FF6B35` | Primary brand accent |
| `--accent-glow` | `rgba(255,107,53,0.3)` | Glow effects on accent |

### Gradient tokens

| Token | Use |
|---|---|
| `--gradient-1` | Primary CTA backgrounds (`#FF6B35 → #FF3CAC`) |
| `--gradient-2` | Secondary / hero backgrounds |
| `--gradient-3` | Success / positive states |
| `--gradient-gold` | Premium / reward indicators |

### Spacing

Use multiples of **4px**. Common values: `4 8 12 16 24 32 48 64px`.
Do not use arbitrary values like `13px` or `22px`.

### Border radius tokens

| Token | Value | Use |
|---|---|---|
| `--radius-sm` | `12px` | Chips, badges |
| `--radius-md` | `16px` | Cards |
| `--radius-lg` | `24px` | Panels, modals |
| `--radius-xl` | `32px` | Large containers |

### Motion token

`--transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94)`

Use this for all interactive state changes (hover, active, focus). For orchestrated
animations (floating hearts, toasts) use `300–500ms` with the same easing.

### Typography

Font stack: `'Inter', -apple-system, sans-serif`
Weights in use: `300 400 500 600 700 800 900` — loaded from Google Fonts.

---

## 5. Coding conventions

### CSS

- Styles go inside the `<style>` block of the relevant file. Do not add external `.css` files.
- Organize sections with the comment pattern used throughout: `/* ===== SECTION NAME ===== */`
- Reference tokens via `var(--token-name)`. No raw hex values or pixel values outside of token definitions.
- Mobile-first: base styles target small viewports; `@media (min-width: N)` rules layer on top.
- Touch targets: minimum `44 × 44px` on interactive elements (buttons, icons, nav items).

### JavaScript

- Scripts go in the `<script>` block at the bottom of the file.
- No frameworks, no modules, no transpilation. Vanilla ES2020 is fine (`const`, `let`,
  arrow functions, optional chaining, template literals).
- Group related functionality under clearly labeled comment blocks.
- Use `addEventListener` — do not use inline `onclick` attributes on new elements.

### HTML

- Semantic elements where meaningful (`<nav>`, `<main>`, `<article>`, `<button>`).
- Every `<img>` must have an `alt` attribute. Decorative images use `alt=""`.
- Interactive custom elements need `role` and `tabindex` attributes so they are keyboard-reachable.

---

## 6. Making a change

1. **Identify the file.** Landing page edits go in `index.html`; feed/app edits go in `feed.html`.
2. **Locate the CSS section.** Find the comment block for the component you are changing.
3. **Use existing tokens.** Check the `:root` block before adding any new value.
4. **Add a new token if needed.** Add it to `:root` in the file where it is first used;
   add the same token to the other file if it belongs to the shared system.
5. **Test all five states** for any component you touch: empty, loading, error, partial, complete.
6. **Verify at 375px, 768px, and 1280px** viewport widths before considering the change done.
7. **Check keyboard accessibility.** Tab through any new interactive element; confirm visible focus ring.

---

## 7. Submitting a pull request

- Branch from `main`. Branch name format: `KAN-{ticket}-short-description`.
- Keep the PR scope tight — one feature or fix per PR.
- In the PR description, include:
  - **What changed** and which file(s).
  - **Screenshot or screen recording** at mobile and desktop widths.
  - **Token changes**, if any (new tokens added, values modified).
- Request review from at least one team member before merging.

---

## 8. Out of scope

The following are **not** part of this prototype and should not be added without
explicit agreement:

- npm / package.json / node_modules
- CSS preprocessors (Sass, PostCSS)
- JavaScript bundlers (Webpack, Vite, Rollup)
- Automated tests
- Linters or formatters enforced via tooling
- CI pipelines
- Backend or API integration

If you believe one of these is now needed, open a discussion in the PR or a GitHub issue
before implementing it — adding tooling changes the maintenance contract for every contributor.
