# Low Level Design: CONTRIBUTING.md for ShareChat Revamp

**Document type:** Low Level Design (LLD)
**Feature:** Add `CONTRIBUTING.md` to the `sharechat-revamp` repository
**Branch:** `KAN-128-write-low-level-design-lld-document`
**Author:** Sarah (Staff Engineer, Architecture & Code Review)
**Date:** 2026-04-14
**Status:** Draft

---

## 1. Overview

This document defines the low-level design for adding a `CONTRIBUTING.md` file to the `sharechat-revamp` project. The file serves as the authoritative contributor handbook: covering environment setup, branching conventions, the pull request lifecycle, and code-style rules specific to this codebase.

The target project is a **static HTML prototype** (no build pipeline, no backend, no package manager currently). Design decisions must account for this constraint while leaving clear seams for future tooling adoption.

---

## 2. Scope & Goals

| Goal | In scope | Out of scope |
|---|---|---|
| Contributor onboarding | Yes | CI/CD pipeline setup |
| Branching convention | Yes | Git hooks automation |
| PR & review process | Yes | Issue triage workflow |
| Code style rules | Yes | Linter/formatter tooling config |
| Setup instructions | Yes | Deployment runbooks |

---

## 3. Repository Package Structure

Current state and the additions this feature introduces:

```
sharechat-revamp/
├── index.html              # Landing / marketing page
├── feed.html               # Web app (Instagram-style feed)
├── UI_SPEC.md              # Design system reference (existing)
├── CONTRIBUTING.md         # ← NEW: contributor handbook
├── .github/
│   ├── PULL_REQUEST_TEMPLATE.md   # ← NEW: standard PR description template
│   └── ISSUE_TEMPLATE/
│       └── bug_report.md          # ← NEW: bug report template
└── .gitignore              # ← NEW: basic gitignore for editors/OS files
```

**Rationale for `.github/` directory:**
GitHub auto-discovers `PULL_REQUEST_TEMPLATE.md` and issue templates from `.github/`. Centralising them here keeps the root clean and aligns with GitHub conventions, so contributors find them without being told where to look.

---

## 4. Data Models

Because this project has no persistent backend, "data models" in this context describe the **conceptual entities** that the contribution workflow operates on.

### 4.1 Contributor

```
Contributor {
  github_handle : string       // unique identifier
  role          : enum(contributor | maintainer | reviewer)
  cla_signed    : boolean      // required before first merge
}
```

### 4.2 Branch

```
Branch {
  name          : string       // format: <type>/<ticket-id>-<slug>
  base          : string       // always: main
  type          : enum(feat | fix | docs | refactor | chore | test)
  ticket_id     : string?      // e.g. KAN-128; nullable for hotfixes
  slug          : string       // kebab-case, max 40 chars
  status        : enum(open | merged | abandoned)
}
```

Branch name regex enforced in PR title check:
```
^(feat|fix|docs|refactor|chore|test)\/([A-Z]+-\d+-)?[a-z0-9-]{3,40}$
```

Examples:
- `feat/KAN-131-video-autoplay-loop`
- `fix/KAN-129-mobile-overflow-clip`
- `docs/contributing-guide`

### 4.3 Pull Request

```
PullRequest {
  id            : string       // GitHub PR number
  title         : string       // follows: <type>(<scope>): <summary>
  branch        : Branch
  author        : Contributor
  state         : enum(draft | ready | changes_requested | approved | merged | closed)
  reviewers     : Contributor[]  // min 1 required for merge
  labels        : string[]
  linked_ticket : string?      // e.g. "Closes KAN-128"
  created_at    : datetime
  merged_at     : datetime?
}
```

### 4.4 Commit

```
Commit {
  sha           : string
  message       : string       // follows Conventional Commits spec
  type          : enum(feat | fix | docs | style | refactor | test | chore)
  scope         : string?      // e.g. feed, landing, design-system
  breaking      : boolean      // BREAKING CHANGE footer triggers major version
  body          : string?
  footer        : string?
}
```

Commit message format:
```
<type>(<scope>): <imperative summary, ≤72 chars>

[optional body — wrap at 72 chars]

[optional footer: Closes KAN-XXX | BREAKING CHANGE: ...]
```

---

## 5. DB Schema

This project has no database. This section documents the **file-system schema** — the structural contract that files in this repo must conform to — which serves an analogous role to a DB schema in terms of integrity constraints.

### 5.1 File Naming Rules

| File type | Convention | Example |
|---|---|---|
| HTML pages | `kebab-case.html` | `feed.html`, `404.html` |
| Markdown docs | `SCREAMING_SNAKE.md` (root), `kebab-case.md` (`.github/`) | `CONTRIBUTING.md`, `bug_report.md` |
| Assets (future) | `kebab-case.<ext>` | `hero-bg.webp` |
| CSS files (future) | `kebab-case.css` | `design-tokens.css` |

### 5.2 CSS Custom Property Schema

The design system is defined as CSS custom properties in both HTML files. The canonical schema:

```
--bg-primary      : #0a0a0a
--bg-secondary    : #111111
--bg-card         : #1a1a1a
--bg-glass        : rgba(255,255,255,0.06)
--border-glass    : rgba(255,255,255,0.08)
--text-primary    : #ffffff
--text-secondary  : rgba(255,255,255,0.6)
--text-tertiary   : rgba(255,255,255,0.35)
--accent          : #FF6B35
--accent-glow     : rgba(255,107,53,0.3)
--gradient-1      : linear-gradient(135deg, #FF6B35, #FF3CAC)
--gradient-2      : linear-gradient(135deg, #FF3CAC, #784BA0, #2B86C5)
--gradient-3      : linear-gradient(135deg, #00F5A0, #00D9F5)
--gradient-gold   : linear-gradient(135deg, #F7971E, #FFD200)
--radius-sm       : 12px
--radius-md       : 16px
--radius-lg       : 24px
--radius-xl       : 32px
--transition      : all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94)
```

**Constraint:** any PR that introduces a new CSS color or spacing value not in this schema MUST either (a) add it as a new custom property in `:root`, or (b) document the deviation in the PR description with a reason. Hard-coded hex values in component styles are a merge blocker.

---

## 6. API Contracts

Since the project is a static frontend with no backend API, "API contracts" in this context describe the **interfaces between the contribution workflow components** — what each component accepts and returns.

### 6.1 GitHub PR Template Contract

The PR template enforces a structured input contract. Every PR description MUST satisfy:

```
## Summary
- <bullet: what changed and why>

## Type of change
- [ ] Bug fix (non-breaking)
- [ ] New feature (non-breaking)
- [ ] Breaking change
- [ ] Documentation update

## Checklist
- [ ] CSS uses design tokens from :root (no raw hex values)
- [ ] Tested on Chrome/Firefox/Safari desktop
- [ ] Tested on iOS Safari / Android Chrome (if UI changed)
- [ ] Accessibility: keyboard nav and focus states verified
- [ ] No console errors or warnings

## Screenshots / recordings
<!-- Required for any UI change -->

## Related issue
Closes KAN-XXX
```

**Validation rules (reviewed by maintainer, not automated):**
- `Summary` section: min 1 bullet, max 10 bullets
- `Screenshots` section: required if any `.html` file is modified
- `Related issue`: required unless type is `docs` or `chore`

### 6.2 Branch Protection Rules (GitHub Settings Contract)

These are the settings that MUST be configured on the `main` branch by a repo admin. They form a hard contract that enforces the workflow:

```yaml
branch: main
protection:
  required_pull_request_reviews:
    required_approving_review_count: 1
    dismiss_stale_reviews: true
    require_code_owner_reviews: false
  required_status_checks:
    strict: true               # branch must be up-to-date before merge
    contexts: []               # no CI yet; expand when CI is added
  enforce_admins: false        # maintainers can bypass in emergencies
  allow_force_pushes: false
  allow_deletions: false
  required_linear_history: true  # squash or rebase only; no merge commits
```

### 6.3 Commit Linting Contract (Future — when CI is added)

When a CI pipeline is introduced, commit messages will be validated against this contract:

```
Input:  git commit message string
Output: PASS | FAIL(reason)

Rules:
  - type    ∈ {feat, fix, docs, style, refactor, test, chore}
  - subject length ≤ 72 characters
  - subject not ending with "."
  - no ALL-CAPS subject
  - if breaking change: BREAKING CHANGE footer required
```

---

## 7. Sequence Diagrams

### 7.1 Standard Feature Contribution Flow

```
Contributor          GitHub                  Maintainer
     │                  │                        │
     │  fork repo        │                        │
     │─────────────────▶│                        │
     │                  │                        │
     │  git checkout -b feat/KAN-XXX-slug         │
     │─────────────────▶│                        │
     │                  │                        │
     │  [implement changes]                       │
     │                  │                        │
     │  git commit -m "feat(feed): ..."           │
     │─────────────────▶│                        │
     │                  │                        │
     │  git push origin feat/KAN-XXX-slug         │
     │─────────────────▶│                        │
     │                  │                        │
     │  open PR → fill template                   │
     │─────────────────▶│                        │
     │                  │  notify reviewer        │
     │                  │───────────────────────▶│
     │                  │                        │
     │                  │     review + comments   │
     │                  │◀───────────────────────│
     │                  │                        │
     │  address feedback │                        │
     │─────────────────▶│                        │
     │                  │  re-request review      │
     │                  │───────────────────────▶│
     │                  │                        │
     │                  │     approve PR          │
     │                  │◀───────────────────────│
     │                  │                        │
     │  squash & merge  │                        │
     │─────────────────▶│                        │
     │                  │  delete branch          │
     │                  │───────────────────────▶│
```

### 7.2 Hotfix Flow (bypasses normal branch naming)

```
Maintainer           GitHub
     │                  │
     │  git checkout -b fix/critical-<slug> main
     │─────────────────▶│
     │                  │
     │  [minimal targeted fix]
     │                  │
     │  open PR, label: hotfix
     │─────────────────▶│
     │                  │
     │  self-review or fast peer review (1 approve)
     │─────────────────▶│
     │                  │
     │  squash & merge to main
     │─────────────────▶│
```

### 7.3 Documentation-Only Change Flow

```
Contributor          GitHub                  Maintainer
     │                  │                        │
     │  branch: docs/<slug>                       │
     │─────────────────▶│                        │
     │                  │                        │
     │  edit .md files only                       │
     │─────────────────▶│                        │
     │                  │                        │
     │  PR: no screenshot required               │
     │─────────────────▶│                        │
     │                  │  notify reviewer        │
     │                  │───────────────────────▶│
     │                  │                        │
     │                  │     approve (light review)
     │                  │◀───────────────────────│
     │                  │                        │
     │  merge           │                        │
     │─────────────────▶│                        │
```

---

## 8. Implementation Patterns

### 8.1 HTML Code Style Rules

These rules apply to all `*.html` files and will be documented verbatim in `CONTRIBUTING.md`:

**Indentation**
- 2-space indentation (no tabs)
- Inline `<style>` blocks: same 2-space indent
- Inline `<script>` blocks: same 2-space indent

**CSS organisation within `<style>` blocks**
```
1. :root { custom properties }
2. CSS reset / base
3. Layout components (in DOM order)
4. Micro-interactions / animations (@keyframes at end)
5. Media queries (grouped at end, smallest breakpoint up)
```

**Commenting convention**
```css
/* ==================== SECTION NAME ==================== */
/* Sub-section comment */
```

**JavaScript patterns (inline scripts)**
- `const` / `let` only; no `var`
- Arrow functions for callbacks
- `querySelectorAll` + `forEach` over `getElementsBy*`
- Event delegation where >3 similar listeners would exist
- No external script dependencies; self-contained behaviour only

**Accessibility baseline (enforced at review)**
- Every interactive element has `aria-label` or visible text
- Focus styles visible (`outline: 2px solid var(--accent)` minimum)
- Color contrast: text on background ≥ 4.5:1 (WCAG AA)
- Video elements have `aria-label` describing content

### 8.2 CONTRIBUTING.md Document Structure

The file will be authored with these sections in order:

```
1. Welcome & project context       (~50 words)
2. Code of conduct                 (link to GitHub community guidelines)
3. Prerequisites                   (browser, git, text editor)
4. Local setup                     (clone → open in browser; no build step)
5. Branch naming                   (table: type → pattern → example)
6. Commit messages                 (Conventional Commits cheatsheet)
7. Opening a pull request          (checklist reference → .github/PULL_REQUEST_TEMPLATE.md)
8. Code review etiquette           (reviewer/author responsibilities)
9. Code style quick reference      (link to UI_SPEC.md; inline summary)
10. Getting help                   (how to reach maintainers)
```

### 8.3 Template Patterns

**PR Template structure pattern (`.github/PULL_REQUEST_TEMPLATE.md`)**

The template uses GitHub's checkbox syntax which GitHub renders as interactive checkboxes on the PR form. All checkboxes are unchecked by default — the author must check them before requesting review. This is a social contract enforced at review time, not by automation.

**Issue Template structure pattern (`.github/ISSUE_TEMPLATE/bug_report.md`)**

YAML front matter controls GitHub's "New issue" form:
```yaml
---
name: Bug report
about: File a bug against feed.html or index.html
title: '[BUG] '
labels: bug
assignees: ''
---
```

Sections:
1. **Describe the bug** — one-sentence summary
2. **Steps to reproduce** — numbered list
3. **Expected behavior** — what should happen
4. **Actual behavior** — what happens
5. **Environment** — browser, OS, device type
6. **Screenshots** — if applicable

### 8.4 `.gitignore` Pattern

The `.gitignore` should cover common editor and OS artifacts for a no-build-tooling static project:

```gitignore
# OS
.DS_Store
Thumbs.db

# Editors
.vscode/
.idea/
*.swp
*.swo
*~

# Temporary files
*.tmp
*.log
```

**Pattern reasoning:** a static HTML project has no `node_modules`, no build output, no `.env`. A minimal ignore file avoids false confidence — contributors won't think they're protected from accidentally committing secrets that aren't listed. When a build pipeline is added, the `.gitignore` must be expanded in the same PR.

---

## 9. Edge Cases & Decision Log

| Decision | Options considered | Chosen | Reason |
|---|---|---|---|
| Squash vs rebase vs merge commit | All three | Squash | Keeps `main` log readable; one feature = one commit; aligns with linear history rule |
| Require CLA | CLA Assistant bot vs honour system | Honour system | Project is a prototype; CLA tooling adds friction not proportional to risk |
| Screenshot requirement | Always required vs UI-change-only | UI-change-only | Docs/chore PRs have no visual output to screenshot |
| Commit lint automation | Pre-commit hook vs CI vs manual | Manual (document only) | No tooling in repo yet; adding `husky` is a separate task |
| PR template location | Root vs `.github/` | `.github/` | GitHub auto-detects; root stays uncluttered |
| `required_linear_history` | Enabled vs disabled | Enabled | Prevents merge commits; easier bisect/revert |

---

## 10. Open Questions

1. **CLA requirement** — Does the project expect external contributors? If yes, a CLA bot (e.g., CLA Assistant) should be added; scope is a separate KAN ticket.
2. **CI pipeline** — When will automated checks be introduced? The branch protection `contexts: []` must be updated when first CI job is added.
3. **CODEOWNERS file** — Should a `.github/CODEOWNERS` be added to auto-assign reviewers? Recommend yes once team ownership is stable.
4. **Versioning / changelog** — No `CHANGELOG.md` or semantic versioning defined yet. Should this CONTRIBUTING.md reference a versioning strategy?

---

## 11. Implementation Checklist

- [ ] Create `CONTRIBUTING.md` at repo root (per Section 8.2 structure)
- [ ] Create `.github/PULL_REQUEST_TEMPLATE.md` (per Section 6.1)
- [ ] Create `.github/ISSUE_TEMPLATE/bug_report.md` (per Section 8.3)
- [ ] Create `.gitignore` (per Section 8.4)
- [ ] Configure branch protection rules on `main` (per Section 6.2) — requires repo admin
- [ ] Link `UI_SPEC.md` from CONTRIBUTING.md code style section
- [ ] PR this LLD for review before implementation begins

---

*This LLD covers the design decisions behind adding `CONTRIBUTING.md`. The actual file contents will be authored in the implementation PR, using this document as the specification.*
