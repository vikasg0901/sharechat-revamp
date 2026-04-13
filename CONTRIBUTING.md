# Contributing to ShareChat Revamp

Thank you for your interest in contributing to ShareChat Revamp — a dark-themed, Instagram-style social video web app. This guide will help you get up and running quickly and ensure your contributions land smoothly.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Report a Bug](#how-to-report-a-bug)
- [How to Request a Feature](#how-to-request-a-feature)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Design System](#design-system)
- [Coding Standards](#coding-standards)
- [Commit Message Format](#commit-message-format)
- [Pull Request Process](#pull-request-process)
- [Accessibility Requirements](#accessibility-requirements)

---

## Code of Conduct

Be respectful and constructive in all interactions. We welcome contributors of all backgrounds and experience levels. Harassment, discrimination, or hostile behavior will not be tolerated.

---

## How to Report a Bug

Before filing a new issue, search existing issues to avoid duplicates.

When opening a bug report, include:

- **Environment** — browser name and version, OS, viewport size
- **Steps to reproduce** — numbered, minimal steps that reliably trigger the bug
- **Expected behavior** — what should have happened
- **Actual behavior** — what actually happened
- **Screenshots or screen recording** — attach if the bug is visual or interactive

---

## How to Request a Feature

Open a GitHub issue with the label `enhancement`. Describe:

- The user problem you are trying to solve
- The proposed solution or interaction
- Any design references or mockups

For large changes, open a discussion issue first so the approach can be agreed on before work begins.

---

## Development Setup

The project uses plain HTML, CSS, and JavaScript — no build step required.

```bash
# 1. Clone the repository
git clone https://github.com/<org>/sharechat-revamp.git
cd sharechat-revamp

# 2. Open in a browser
#    Any static server works. Example with Python:
python -m http.server 8080
#    Then navigate to http://localhost:8080
```

For live reload during development, use a tool such as [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) (VS Code extension) or [browser-sync](https://browsersync.io/).

---

## Project Structure

```
sharechat-revamp/
├── index.html        # Landing / marketing page
├── feed.html         # Web app — Instagram-style video feed
├── 404.html          # Not-found error page
├── CONTRIBUTING.md   # This file
└── README.md         # Project overview
```

All styles are **inline** within each HTML file using CSS custom properties defined in `:root`. There is intentionally no external stylesheet — keep it that way to preserve the zero-build-step constraint.

---

## Design System

All design decisions must follow the established token set. Do not introduce hard-coded color or spacing values.

### Color Tokens

| Token | Value | Usage |
|---|---|---|
| `--bg-primary` | `#0a0a0a` | Page background |
| `--bg-secondary` | `#111111` | Section backgrounds |
| `--bg-card` | `#1a1a1a` | Card and panel surfaces |
| `--bg-glass` | `rgba(255,255,255,0.06)` | Frosted-glass overlays |
| `--border-glass` | `rgba(255,255,255,0.08)` | Glass element borders |
| `--text-primary` | `#ffffff` | Headings, primary copy |
| `--text-secondary` | `rgba(255,255,255,0.6)` | Secondary copy |
| `--text-tertiary` | `rgba(255,255,255,0.35)` | Placeholder, disabled |
| `--accent` | `#FF6B35` | Brand accent, CTAs |
| `--accent-glow` | `rgba(255,107,53,0.3)` | Accent glow effects |

### Gradients

| Token | Direction & Stops |
|---|---|
| `--gradient-1` | `135deg` · `#FF6B35` → `#FF3CAC` |
| `--gradient-2` | `135deg` · `#FF3CAC` → `#784BA0` → `#2B86C5` |
| `--gradient-3` | `135deg` · `#00F5A0` → `#00D9F5` |
| `--gradient-gold` | `135deg` · `#F7971E` → `#FFD200` |

### Border Radius

| Token | Value |
|---|---|
| `--radius-sm` | `12px` |
| `--radius-md` | `16px` |
| `--radius-lg` | `24px` |
| `--radius-xl` | `32px` |

### Typography

- **Display / headings:** `Space Grotesk`, weights 400–700
- **Body / UI:** `Inter`, weights 300–900
- Use `clamp()` for fluid heading sizes (e.g. `clamp(48px, 7vw, 88px)`)

### Motion

Use the shared easing function for all transitions:

```css
--transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
```

Prefer `transform` and `opacity` for animated properties to stay on the compositor thread.

---

## Coding Standards

### HTML

- Use semantic elements: `<nav>`, `<main>`, `<section>`, `<article>`, `<button>`, `<a>`.
- Never use `<div>` or `<span>` as interactive elements. An action is a `<button>`; navigation is an `<a>`.
- Every image must have a meaningful `alt` attribute. Use `alt=""` for decorative images.
- Every `<input>` must have an associated `<label>` (visible or `sr-only`).
- Set explicit `width` and `height` on images to prevent Cumulative Layout Shift.

### CSS

- Define all new design values as CSS custom properties in `:root` — never inline a raw hex or pixel value.
- Follow mobile-first ordering: base styles first, then `@media (min-width: ...)` overrides.
- Standard breakpoints: `640px` (sm) · `768px` (md) · `1024px` (lg) · `1280px` (xl).
- Avoid `!important`. If you need it, the selector specificity model is broken — fix that instead.

### JavaScript

- Keep scripts self-contained within the `<script>` tag at the bottom of each HTML file.
- Use `const` and `let`; never `var`.
- Prefer event delegation over attaching listeners to every individual element.
- Debounce input handlers that trigger expensive operations (300 ms).
- Do not add third-party JS libraries without team agreement — keep the zero-dependency constraint.

---

## Commit Message Format

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.

```
<type>(<scope>): <short summary>

[optional body]

[optional footer]
```

**Types:** `feat` · `fix` · `refactor` · `style` · `docs` · `test` · `chore`

**Scope examples:** `feed` · `index` · `nav` · `404` · `a11y` · `animation`

**Examples:**

```
feat(feed): add floating hearts animation on like
fix(feed): make action buttons clickable on mobile
docs: add CONTRIBUTING.md
style(index): tighten hero heading letter-spacing
a11y(feed): add aria-label to mute toggle button
```

- Keep the summary under 72 characters.
- Use the imperative mood: "add", not "added" or "adds".
- Reference the Kanban ticket at the end of the body if applicable: `Closes KAN-130`.

---

## Pull Request Process

1. **Branch** off `main` using the naming convention `KAN-<ticket>-<short-slug>`.
2. **Keep PRs focused** — one logical change per PR. Do not bundle unrelated fixes.
3. **Self-review** before requesting review: check the diff, run the checklist below.
4. **Fill in the PR description** — what changed, why, and how to test it.
5. **Request review** from at least one team member.
6. **Squash merge** into `main` once approved.

### Pre-submission Checklist

- [ ] Tested in Chrome, Firefox, and Safari (latest stable)
- [ ] Tested at mobile (375px), tablet (768px), and desktop (1280px) widths
- [ ] No new hard-coded color or spacing values — design tokens used throughout
- [ ] Interactive elements are keyboard-navigable (Tab, Enter, Escape where applicable)
- [ ] Visible focus indicator present on all interactive elements
- [ ] No `console.error` or `console.warn` output in the browser
- [ ] Images have `alt` text; decorative images use `alt=""`
- [ ] Motion respects `prefers-reduced-motion` if the change adds animation

---

## Accessibility Requirements

All contributions must meet WCAG 2.1 AA as a minimum bar.

- **Color contrast:** text on background must meet 4.5:1 (normal text) or 3:1 (large text / UI components).
- **Focus management:** modals, drawers, and overlays must trap focus and return it to the trigger on close.
- **Keyboard navigation:** every interaction available by pointer must also be available by keyboard.
- **Screen reader:** test with NVDA (Windows) or VoiceOver (macOS) before marking UI work complete.
- **Reduced motion:** wrap CSS animations in `@media (prefers-reduced-motion: no-preference)` so users who opt out of motion are not affected.

---

Questions? Open a discussion issue or reach out in the project's Slack channel.
