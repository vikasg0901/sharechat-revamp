# Low Level Design: Add MIT LICENSE File
**Ticket:** KAN-124  
**Author:** Sarah (Staff Engineer)  
**Date:** 2026-04-14  
**Status:** Draft

---

## 1. Overview

Add a `LICENSE` file to the root of the `sharechat-revamp` repository containing the OSI-approved MIT license text. This is a pure file-system change — no runtime code, no data model, no API surface, no build pipeline impact.

---

## 2. Scope & Constraints

| Dimension | Decision |
|-----------|----------|
| File count | 1 new file (`LICENSE`) |
| Location | Repository root |
| Format | Plain text (no extension), UNIX line endings (`\n`) |
| License | MIT (SPDX identifier: `MIT`) |
| Copyright year | 2026 |
| Runtime impact | None |
| Build impact | None |
| Test impact | None |

Out of scope: adding license headers to individual source files (`index.html`, `feed.html`). That is a separate, subsequent task if required.

---

## 3. Repository Package Structure

Current layout:
```
sharechat-revamp/
├── index.html       # Landing / marketing page
├── feed.html        # Web app (Instagram-style feed)
└── CLAUDE.md        # Agent task context (not shipped)
```

Post-change layout:
```
sharechat-revamp/
├── LICENSE          # ← NEW: MIT license, plain text
├── index.html
├── feed.html
└── CLAUDE.md
```

### 3.1 File Naming Rationale

`LICENSE` (no extension) is the de-facto standard recognized by:
- GitHub — renders automatically on the repository home page and shows the SPDX badge
- `npm`, `pip`, `cargo`, and other package registries — auto-detected for package metadata
- REUSE / SPDX tooling — compliant with `REUSE spec v3.0` without additional configuration

Alternative names (`LICENSE.txt`, `LICENSE.md`) are acceptable but reduce automatic detection coverage. `LICENSE` is preferred.

---

## 4. Data Model

The MIT license file is a structured text artifact. Its schema is defined by the [OSI MIT License template](https://opensource.org/license/mit):

```
MIT License

Copyright (c) <YEAR> <COPYRIGHT HOLDER>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

### 4.1 Template Variable Binding

| Template variable | Value | Source |
|-------------------|-------|--------|
| `<YEAR>` | `2026` | Current year at time of authoring |
| `<COPYRIGHT HOLDER>` | `Vikas Gupta` | Git user config (`git config user.name`) |

> **Note:** If the project is owned by an organization rather than an individual, `<COPYRIGHT HOLDER>` should be updated to the organization name (e.g., `ShareChat Inc.`). Clarify with the ticket requester before merging.

---

## 5. DB Schema

Not applicable. This change introduces no database entities, migrations, or schema modifications.

---

## 6. API Contracts

Not applicable. This is a static file addition with no API surface.

GitHub's License API (`GET /repos/{owner}/{repo}/license`) will automatically detect and return the license after the file is committed. No action required on our side to enable this.

---

## 7. Sequence Diagram

### 7.1 Author Flow

```
Author              Git                  GitHub
  |                  |                     |
  |-- git checkout --branch KAN-124 -----> |
  |                  |                     |
  |-- Write LICENSE -|                     |
  |   (MIT text)     |                     |
  |                  |                     |
  |-- git add LICENSE|                     |
  |-- git commit -m "Add MIT LICENSE" ---> |
  |                  |                     |
  |-- git push origin KAN-124 -----------> |
  |                  |                     |
  |                  |<-- PR created ------|
  |                  |    (review gate)    |
  |                  |                     |
  |<-- PR approved --|                     |
  |                  |                     |
  |-- git merge to main ----------------> |
  |                  |                     |
  |                  |   GitHub detects    |
  |                  |   LICENSE, updates  |
  |                  |   repo metadata     |
```

### 7.2 GitHub License Detection Flow (automatic, post-merge)

```
GitHub crawler
  |
  |-- Scan repo root for LICENSE file
  |-- Parse SPDX identifier: MIT
  |-- Set repo.license = { spdx_id: "MIT", name: "MIT License" }
  |-- Render license badge on repo page
```

---

## 8. Implementation Pattern

### 8.1 File Content

Exact bytes to write to `LICENSE` (UTF-8, LF line endings):

```
MIT License

Copyright (c) 2026 Vikas Gupta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

### 8.2 Commit Message Convention

```
Add MIT LICENSE

Adds the OSI-approved MIT license to the repository root.
GitHub will auto-detect SPDX identifier MIT and update repo metadata.

Refs: KAN-124
```

### 8.3 Git Workflow

```bash
# On branch KAN-124-write-low-level-design-lld-document (already checked out)
# After LLD is approved, the LICENSE file is created separately or in the same branch.

git add LICENSE
git commit -m "Add MIT LICENSE

Adds the OSI-approved MIT license to the repository root.

Refs: KAN-124"

git push origin KAN-124-write-low-level-design-lld-document
```

---

## 9. Validation Checklist

| Check | How | Expected |
|-------|-----|----------|
| File exists at root | `ls LICENSE` | File present |
| SPDX text intact | `grep "MIT License" LICENSE` | Match found |
| Copyright line correct | `grep "Copyright" LICENSE` | `Copyright (c) 2026 Vikas Gupta` |
| No trailing whitespace | `grep -P "\s+$" LICENSE` | No output |
| Line endings are LF | `file LICENSE` | `ASCII text` (not CRLF) |
| GitHub detection | View repo on GitHub post-merge | License badge shows "MIT" |

---

## 10. Risks & Mitigations

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| Wrong copyright holder name | Low | Medium | Confirm with ticket requester; default to git user |
| CRLF line endings on Windows | Medium | Low | Use `\n` explicitly; configure `.gitattributes` if needed |
| License text modified from OSI template | Low | High | Copy verbatim from `https://opensource.org/license/mit`; do not paraphrase |
| Applying MIT to code that has third-party dependencies under incompatible licenses | Low | High | Out of scope for this ticket; project is currently pure HTML with no npm deps — low risk today |

---

## 11. Alternatives Considered

| Option | Decision | Reason |
|--------|----------|--------|
| Apache 2.0 | Rejected | More verbose; MIT is the conventional choice for front-end projects |
| GPL-3.0 | Rejected | Copyleft; restricts downstream consumers unnecessarily for a UI prototype |
| No license (status quo) | Rejected | All-rights-reserved by default; blocks open-source use |
| LICENSE.txt | Rejected | Less canonical; GitHub detection rate lower than bare `LICENSE` |

---

## 12. Open Questions

1. **Copyright holder:** Should this be `Vikas Gupta` (individual) or an organization name (e.g., `ShareChat Inc.`)? Needs confirmation before implementation.
2. **Source file headers:** Should `index.html` and `feed.html` receive `<!-- SPDX-License-Identifier: MIT -->` comment headers? Not required for MIT compliance, but good practice. Recommend deferring to a follow-up ticket.

---

## 13. Implementation Estimate

| Task | Effort |
|------|--------|
| Create `LICENSE` file | < 5 min |
| Commit and push | < 5 min |
| Open PR and get review | Depends on reviewer availability |
| **Total engineering effort** | **~10 min** |
