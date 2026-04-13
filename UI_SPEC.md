# ShareChat Revamp — UI/UX Specification
**Version:** 1.0  
**Date:** 2026-04-14  
**Author:** Aisha (Design Engineering)  
**Status:** Ready for Engineering Handoff

---

## Table of Contents

1. [Design Tokens](#1-design-tokens)
2. [Typography System](#2-typography-system)
3. [Spacing & Layout Grid](#3-spacing--layout-grid)
4. [Component Hierarchy](#4-component-hierarchy)
5. [Page Layouts & Wireframes](#5-page-layouts--wireframes)
6. [Component Specifications](#6-component-specifications)
7. [Interaction & Animation Spec](#7-interaction--animation-spec)
8. [Responsive Behavior](#8-responsive-behavior)
9. [Accessibility Spec](#9-accessibility-spec)
10. [State Definitions (Five-State Model)](#10-state-definitions-five-state-model)

---

## 1. Design Tokens

All tokens are defined as CSS custom properties on `:root`. Engineers reference tokens only — never raw values.

### 1.1 Primitive Tokens

```css
/* ─── Color Primitives ─── */
--color-black-100: #0a0a0a;   /* true OLED black */
--color-black-90:  #111111;
--color-black-80:  #1a1a1a;
--color-black-70:  #222222;
--color-white:     #ffffff;

--color-orange-500: #FF6B35;
--color-pink-500:   #FF3CAC;
--color-purple-500: #784BA0;
--color-blue-500:   #2B86C5;
--color-mint-500:   #00F5A0;
--color-cyan-500:   #00D9F5;
--color-gold-400:   #FFD200;
--color-gold-500:   #F7971E;
--color-red-500:    #FF4444;
--color-green-500:  #00C853;

/* ─── Alpha Primitives ─── */
--alpha-glass-bg:     rgba(255, 255, 255, 0.06);
--alpha-glass-border: rgba(255, 255, 255, 0.08);
--alpha-glass-hover:  rgba(255, 255, 255, 0.10);
--alpha-overlay-dark: rgba(0, 0, 0, 0.60);
--alpha-overlay-mid:  rgba(0, 0, 0, 0.40);

/* ─── Spacing Primitives (4px grid) ─── */
--space-1:  4px;
--space-2:  8px;
--space-3:  12px;
--space-4:  16px;
--space-5:  20px;
--space-6:  24px;
--space-8:  32px;
--space-10: 40px;
--space-12: 48px;
--space-16: 64px;
--space-20: 80px;
--space-30: 120px;

/* ─── Border Radius Primitives ─── */
--radius-2:  8px;
--radius-3:  12px;
--radius-4:  16px;
--radius-6:  24px;
--radius-8:  32px;
--radius-full: 9999px;

/* ─── Shadow Primitives ─── */
--shadow-glow-sm: 0 0 8px rgba(255, 107, 53, 0.30);
--shadow-glow-md: 0 8px 32px rgba(255, 107, 53, 0.25);
--shadow-glow-lg: 0 12px 40px rgba(255, 107, 53, 0.35);
--shadow-card:    0 4px 24px rgba(0, 0, 0, 0.40);
--shadow-modal:   0 24px 64px rgba(0, 0, 0, 0.70);

/* ─── Duration Primitives ─── */
--duration-fast:   150ms;
--duration-normal: 250ms;
--duration-slow:   400ms;
--duration-xslow:  600ms;

/* ─── Easing Primitives ─── */
--ease-standard: cubic-bezier(0.25, 0.46, 0.45, 0.94);
--ease-spring:   cubic-bezier(0.34, 1.56, 0.64, 1.00);
--ease-exit:     cubic-bezier(0.55, 0.00, 1.00, 0.45);
--ease-enter:    cubic-bezier(0.00, 0.55, 0.45, 1.00);
```

### 1.2 Semantic Tokens

```css
/* ─── Backgrounds ─── */
--bg-body:     var(--color-black-100);
--bg-surface:  var(--color-black-90);
--bg-card:     var(--color-black-80);
--bg-elevated: var(--color-black-70);
--bg-glass:    var(--alpha-glass-bg);

/* ─── Text ─── */
--text-primary:   var(--color-white);
--text-secondary: rgba(255, 255, 255, 0.60);
--text-tertiary:  rgba(255, 255, 255, 0.35);
--text-disabled:  rgba(255, 255, 255, 0.20);
--text-inverse:   var(--color-black-100);
--text-accent:    var(--color-orange-500);
--text-link:      var(--color-orange-500);
--text-danger:    var(--color-red-500);
--text-success:   var(--color-green-500);

/* ─── Borders ─── */
--border-default: var(--alpha-glass-border);
--border-hover:   rgba(255, 255, 255, 0.16);
--border-focus:   var(--color-orange-500);
--border-error:   var(--color-red-500);

/* ─── Brand Gradients ─── */
--gradient-primary: linear-gradient(135deg, var(--color-orange-500), var(--color-pink-500));
--gradient-aurora:  linear-gradient(135deg, var(--color-pink-500), var(--color-purple-500), var(--color-blue-500));
--gradient-fresh:   linear-gradient(135deg, var(--color-mint-500), var(--color-cyan-500));
--gradient-gold:    linear-gradient(135deg, var(--color-gold-500), var(--color-gold-400));

/* ─── Spacing Aliases ─── */
--spacing-xs:  var(--space-1);   /*  4px */
--spacing-sm:  var(--space-2);   /*  8px */
--spacing-md:  var(--space-4);   /* 16px */
--spacing-lg:  var(--space-6);   /* 24px */
--spacing-xl:  var(--space-8);   /* 32px */
--spacing-2xl: var(--space-12);  /* 48px */
--spacing-3xl: var(--space-16);  /* 64px */

/* ─── Radius Aliases ─── */
--radius-sm:  var(--radius-3);    /* 12px */
--radius-md:  var(--radius-4);    /* 16px */
--radius-lg:  var(--radius-6);    /* 24px */
--radius-xl:  var(--radius-8);    /* 32px */
--radius-pill: var(--radius-full); /* 9999px */

/* ─── Transition Alias ─── */
--transition-default: all var(--duration-normal) var(--ease-standard);
--transition-fast:    all var(--duration-fast)   var(--ease-standard);
--transition-spring:  all var(--duration-slow)   var(--ease-spring);
```

### 1.3 Component Tokens

```css
/* ─── Button ─── */
--btn-height-sm:    36px;
--btn-height-md:    44px;
--btn-height-lg:    52px;
--btn-padding-x-sm: var(--space-3);   /* 12px */
--btn-padding-x-md: var(--space-6);   /* 24px */
--btn-padding-x-lg: var(--space-8);   /* 32px */
--btn-radius:       var(--radius-pill);
--btn-font-weight:  600;

/* ─── Card ─── */
--card-bg:      var(--bg-card);
--card-border:  var(--border-default);
--card-radius:  var(--radius-lg);
--card-padding: var(--spacing-lg);

/* ─── Nav ─── */
--nav-height:         64px;
--nav-sidebar-width:  72px;
--nav-right-width:    320px;

/* ─── Avatar ─── */
--avatar-xs: 24px;
--avatar-sm: 32px;
--avatar-md: 40px;
--avatar-lg: 48px;
--avatar-xl: 64px;

/* ─── Reel ─── */
--reel-max-width:    400px;
--reel-action-width: 56px;
```

---

## 2. Typography System

**Font Stack:**
```
Display / Headlines: 'Space Grotesk', -apple-system, system-ui, sans-serif
Body / UI:          'Inter', -apple-system, system-ui, sans-serif
```

**Loading strategy:** `font-display: swap` to prevent FOIT. Preload both fonts at `<head>` via `<link rel="preload">`.

### Type Scale

| Token | Size / Line-height | Letter-spacing | Weight | Usage |
|---|---|---|---|---|
| `--type-display` | 88px / 96px | -0.02em | 800 | Hero headline (fluid: clamp(48px, 7vw, 88px)) |
| `--type-h1` | 56px / 64px | -0.015em | 700 | Page titles |
| `--type-h2` | 40px / 48px | -0.01em | 700 | Section titles (fluid: clamp(28px, 4vw, 44px)) |
| `--type-h3` | 28px / 36px | -0.008em | 600 | Card titles, modal headings |
| `--type-h4` | 20px / 28px | -0.005em | 600 | Sub-section labels |
| `--type-body-lg` | 18px / 28px | 0 | 400 | Lead paragraphs, descriptions |
| `--type-body-md` | 16px / 24px | 0 | 400 | Primary body text |
| `--type-body-sm` | 14px / 20px | 0 | 400 | Secondary body, captions |
| `--type-label` | 13px / 16px | 0.02em | 500 | UI labels, nav items |
| `--type-caption` | 12px / 16px | 0.01em | 400 | Timestamps, helper text |
| `--type-micro` | 11px / 14px | 0.03em | 500 | Badges, tags (UPPERCASE) |

**Gradient Text Pattern:**
```css
.text-gradient {
  background: var(--gradient-primary);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
```
Use gradient text only for H1/H2 display content. Never on body text or interactive labels.

---

## 3. Spacing & Layout Grid

### 3.1 Baseline Grid

All spacing uses the **4px base unit**. Only multiples of 4 are permitted.

```
4  8  12  16  20  24  32  40  48  64  80  120
```

**Rule:** Content padding within sections = minimum `--spacing-lg` (24px). Section vertical padding = `--spacing-3xl` (64px) at mobile, `120px` at desktop.

### 3.2 Layout Breakpoints

| Name | Width | Behavior |
|---|---|---|
| `xs` | 0–639px | Single column, mobile-first. No sidebars. |
| `sm` | 640–767px | Single column. Touch-optimized. |
| `md` | 768–1023px | Two-column possible. Right sidebar hidden. |
| `lg` | 1024–1279px | Full layout. Right sidebar at 280px. |
| `xl` | 1280px+ | Max-width content centered. Right sidebar at 320px. |

### 3.3 Content Max-widths

| Context | Max-width |
|---|---|
| Hero content | 900px |
| Section prose content | 1000px |
| Feed center column | 520px |
| Reel card | 400px |
| Modal / drawer | 480px |

### 3.4 Touch Targets

All interactive elements must meet a minimum touch target of **44×44px**. On mobile:
- Buttons: `min-height: 44px`
- Nav icons: wrapped in 44×44px hit area (padding applied to `<a>` or `<button>`, not icon)
- Action buttons (like, share): 48×48px

---

## 4. Component Hierarchy

```
App
├── Navigation
│   ├── NavBar (desktop horizontal)
│   └── SidebarNav (desktop vertical, feed.html)
│       ├── NavLogo
│       ├── NavItem (×7)
│       └── NavUserAvatar
│
├── Pages
│   ├── LandingPage (index.html)
│   │   ├── HeroSection
│   │   │   ├── HeroBadge
│   │   │   ├── HeroHeadline
│   │   │   ├── HeroSubtitle
│   │   │   ├── HeroCtaGroup
│   │   │   │   ├── Button (primary)
│   │   │   │   └── Button (secondary)
│   │   │   └── HeroStats (×3)
│   │   ├── BeforeAfterSection
│   │   │   ├── BeforeCard
│   │   │   └── AfterCard
│   │   ├── PhoneShowcaseSection
│   │   │   └── PhoneMockup (×3)
│   │   │       └── FeedScreen
│   │   ├── FeaturesSection
│   │   │   └── FeatureCard (×6)
│   │   ├── ComparisonSection
│   │   │   └── ComparisonTable
│   │   │       └── ComparisonRow (×N)
│   │   ├── ImpactSection
│   │   │   └── ImpactMetricCard (×4)
│   │   └── CtaSection
│   │       └── CtaCard
│   │
│   └── FeedPage (feed.html)
│       ├── SidebarNav (left, 72px)
│       ├── FeedCenter
│       │   └── ReelCard (×N, scroll-snap)
│       │       ├── VideoContainer
│       │       │   ├── VideoElement
│       │       │   ├── ReelLoader
│       │       │   └── ReelPlayOverlay
│       │       ├── ReelOverlayInfo
│       │       │   ├── CreatorInfo
│       │       │   ├── Description
│       │       │   └── MusicTag
│       │       ├── ReelProgressBar
│       │       └── ReelActions
│       │           ├── ActionButton (like)
│       │           ├── ActionButton (comment)
│       │           ├── ActionButton (share)
│       │           └── ActionButton (bookmark)
│       └── RightSidebar (desktop, 320px)
│           ├── CreatorProfileCard
│           ├── SuggestedCreators
│           └── TrendingTags
│
└── Overlays
    ├── CommentsPanel (slide-up drawer)
    │   ├── CommentList
    │   │   └── CommentItem (×N)
    │   └── CommentInput
    ├── ShareModal
    │   └── ShareGrid (×6 platforms)
    ├── ToastNotification
    ├── LiveCommentBubble (animated)
    └── HeartBurst (double-tap animation)
```

---

## 5. Page Layouts & Wireframes

### 5.1 Landing Page — Desktop (1280px+)

```
┌──────────────────────────────────────────────────────────────────┐
│  NAV: [Logo]        [Vision] [Experience] [Features] [Impact]  [CTA btn] │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│                     HERO SECTION                                 │
│              [Badge: "Major Redesign 2026"]                      │
│                                                                  │
│         The Future of                                            │
│         Social in India            ←── gradient text            │
│                                                                  │
│    [Subtitle: 325M users…]                                       │
│                                                                  │
│    [▶ View Experience]  [Learn More]                             │
│                                                                  │
│    ████ 325M users  ████ 2.5B shares  ████ 15 languages          │
│                                                                  │
├──────────────────────────────────────────────────────────────────┤
│                   BEFORE / AFTER                                 │
│    ┌──────────────────┐    ┌──────────────────┐                  │
│    │   BEFORE         │    │   AFTER          │                  │
│    │  [white/light]   │    │  [dark/glass]    │                  │
│    │  feed mockup     │    │  feed mockup     │                  │
│    └──────────────────┘    └──────────────────┘                  │
├──────────────────────────────────────────────────────────────────┤
│               PHONE SHOWCASE (3 phones floating)                 │
│    ┌─────┐          ┌─────┐          ┌─────┐                     │
│    │Feed │          │Reels│          │Disc.│                     │
│    │     │          │     │          │     │                     │
│    └─────┘          └─────┘          └─────┘                     │
├──────────────────────────────────────────────────────────────────┤
│                   FEATURES (2×3 grid)                            │
│   ┌──────────┐  ┌──────────┐  ┌──────────┐                      │
│   │ 🎬 Video │  │ 🎨 Tools │  │ 🤖 AI    │                      │
│   └──────────┘  └──────────┘  └──────────┘                      │
│   ┌──────────┐  ┌──────────┐  ┌──────────┐                      │
│   │ 💰 Earn  │  │ 🌐 Lang  │  │ ✨ Live  │                      │
│   └──────────┘  └──────────┘  └──────────┘                      │
├──────────────────────────────────────────────────────────────────┤
│                   COMPARISON TABLE                               │
│   Feature              │  Current  │  Revamped                   │
│   ─────────────────────┼───────────┼──────────                   │
│   Vertical Video       │    ✗      │    ✓                        │
│   …                                                              │
├──────────────────────────────────────────────────────────────────┤
│                   IMPACT METRICS (4 cards)                       │
│   ┌────────┐  ┌────────┐  ┌────────┐  ┌────────┐                │
│   │ +40%   │  │ +65%   │  │ +25%   │  │  2×    │                │
│   │ Time   │  │ Video  │  │Creator │  │  Ad    │                │
│   └────────┘  └────────┘  └────────┘  └────────┘                │
├──────────────────────────────────────────────────────────────────┤
│                   CTA CARD (gradient bg)                         │
│              "Ready to build the future?"                        │
│                [Start Building]                                  │
└──────────────────────────────────────────────────────────────────┘
```

### 5.2 Landing Page — Mobile (≤639px)

```
┌────────────────────────┐
│ NAV: [Logo]   [≡ menu] │
├────────────────────────┤
│     HERO               │
│  [Badge]               │
│  The Future            │
│  of Social             │
│  in India              │
│                        │
│  [▶ View Experience]   │
│  [Learn More]          │
│                        │
│  325M   2.5B   15L     │
├────────────────────────┤
│   BEFORE / AFTER       │
│  ┌────────────────┐    │
│  │   BEFORE       │    │
│  └────────────────┘    │
│  ┌────────────────┐    │
│  │   AFTER        │    │
│  └────────────────┘    │
├────────────────────────┤
│  PHONE (1 at a time,   │
│  horizontal scroll)    │
│  ← ┌─────┐ →          │
│     │Feed │            │
│     └─────┘            │
├────────────────────────┤
│  FEATURES (1 column)   │
│  ┌──────────────────┐  │
│  │ 🎬 Vertical Video│  │
│  └──────────────────┘  │
│  ┌──────────────────┐  │
│  │ 🎨 Creator Tools │  │
│  └──────────────────┘  │
│  …                     │
└────────────────────────┘
```

### 5.3 Feed Page — Desktop (1280px+)

```
┌──────────────────────────────────────────────────────────────────────┐
│                                                                      │
│  ┌──────┐  ┌───────────────────────────────┐  ┌────────────────┐   │
│  │ NAV  │  │       FEED CENTER              │  │  RIGHT SIDEBAR │   │
│  │ 72px │  │       max 520px               │  │     320px      │   │
│  │      │  │ ┌───────────────────────────┐ │  │                │   │
│  │  🏠  │  │ │     REEL CARD             │ │  │ [Creator Card] │   │
│  │  🔍  │  │ │   ┌───────────────────┐   │ │  │                │   │
│  │  🎬  │  │ │   │   VIDEO (9:16)    │   │ │  │ Suggested:     │   │
│  │  💬  │  │ │   │                   │   │ │  │ [User 1]       │   │
│  │  🔔  │  │ │   │  [overlay info]   │   │ │  │ [User 2]       │   │
│  │  ✚   │  │ │   │  [creator name]   │   │ │  │ [User 3]       │   │
│  │  👤  │  │ │   │  [description]    │   │ │  │                │   │
│  │      │  │ │   └───────────────────┘   │ │  │ Trending:      │   │
│  │  ⚙   │  │ │   [▓▓▓▓▓▓░░░░] progress  │ │  │ #Cricket       │   │
│  └──────┘  │ │                    [♥ ▼]  │ │  │ #Bollywood     │   │
│            │ │   actions: ♥ 💬 ↗ 🔖     │ │  └────────────────┘   │
│            │ └───────────────────────────┘ │                       │
│            └───────────────────────────────┘                       │
└──────────────────────────────────────────────────────────────────────┘
```

### 5.4 Feed Page — Mobile (≤767px)

```
┌────────────────────────────┐
│  FULL-SCREEN REEL          │
│  (100vw × 100dvh)          │
│                            │
│  ┌──────────────────────┐  │
│  │    VIDEO (full)      │  │
│  │                      │  │
│  │  [Creator] [Follow]  │  │
│  │  [Description...]    │  │
│  │  [♫ Song name]       │  │
│  │                      │  │
│  │              [♥]     │  │
│  │              [💬]    │  │
│  │              [↗]     │  │
│  │              [🔖]    │  │
│  └──────────────────────┘  │
│  [▓▓▓▓▓▓░░░░░░] progress   │
└────────────────────────────┘
```

### 5.5 Comments Panel (Slide-up Drawer)

```
┌────────────────────────────┐
│  // video visible above    │
├────────────────────────────┤
│ ╔══════════════════════╗   │
│ ║  Comments (142)  [×] ║   │
│ ║──────────────────────║   │
│ ║ [avatar] @user  12m  ║   │
│ ║ Great video! 🔥      ║   │
│ ║    [Reply]  ♥ 24      ║   │
│ ║──────────────────────║   │
│ ║ [avatar] @user2 5m   ║   │
│ ║ Amazing content      ║   │
│ ║    [Reply]  ♥ 8       ║   │
│ ╠══════════════════════╣   │
│ ║ [avatar] Add comment…║   │
│ ║                 [→ ] ║   │
│ ╚══════════════════════╝   │
└────────────────────────────┘
```

### 5.6 Share Modal

```
┌─────────────────────────────┐
│  Share to…              [×] │
├─────────────────────────────┤
│  ┌────┐  ┌────┐  ┌────┐    │
│  │ WA │  │ TG │  │ IG │    │
│  │WhAp│  │Tele│  │Inst│    │
│  └────┘  └────┘  └────┘    │
│  ┌────┐  ┌────┐  ┌────┐    │
│  │ TW │  │ 🔗 │  │ …  │    │
│  │Twit│  │Copy│  │More│    │
│  └────┘  └────┘  └────┘    │
└─────────────────────────────┘
```

---

## 6. Component Specifications

### 6.1 Button

**Variants:** `primary` | `secondary` | `ghost` | `danger`  
**Sizes:** `sm` (36px) | `md` (44px) | `lg` (52px)

#### Visual States

| State | Primary | Secondary | Ghost |
|---|---|---|---|
| **Default** | Gradient bg (`--gradient-primary`), white text | Glass bg, white text, glass border | Transparent, `--text-secondary` |
| **Hover** | +4px lift (`translateY(-2px)`), glow shadow `--shadow-glow-md` | `--bg-elevated`, border brightens | `--text-primary`, underline |
| **Active/Press** | `scale(0.97)`, shadow removed | `scale(0.97)` | `scale(0.97)` |
| **Focus-visible** | 2px outline, `--border-focus`, 2px offset | Same | Same |
| **Disabled** | 40% opacity, `cursor: not-allowed`, no hover | Same | Same |
| **Loading** | Spinner replaces label, width locked | Same | n/a |

```
Anatomy:
┌─────────────────────────────┐
│  [leading-icon?] [Label] [trailing-icon?] │
└─────────────────────────────┘
   ↑ btn-padding-x                ↑ btn-padding-x
   ↑ height: btn-height-md (44px)
```

**Transition spec:** `transform var(--duration-fast) var(--ease-standard), box-shadow var(--duration-fast) var(--ease-standard), background var(--duration-normal) var(--ease-standard)`

**Accessibility:**
- `<button>` element (never `<div>` or `<a>` styled as button unless navigating)
- `aria-disabled="true"` when disabled (not HTML `disabled` alone)
- `aria-busy="true"` + `role="status"` spinner when loading
- Minimum 44×44px hit target

---

### 6.2 FeatureCard

```
┌─────────────────────────────────┐
│  ┌────────────┐                 │
│  │  ICON 48px │  gradient bg    │
│  └────────────┘                 │
│                                 │
│  Feature Title  ← type-h4      │
│                                 │
│  Feature description text.      │
│  Two to three lines max.        │
│  ← type-body-sm, text-secondary │
└─────────────────────────────────┘
```

**States:**
- **Default:** `--bg-card`, `--border-default`, `--radius-lg`
- **Hover:** `border-color: --border-hover`, translateY(-4px), `--shadow-card`
- **Focus-visible:** `--border-focus` outline
- **Loading:** Shimmer skeleton replaces icon + text
- **Empty:** Not applicable (always populated from static data)
- **Error:** Not applicable

**Dimensions:**
- Min-height: 180px
- Padding: `--card-padding` (24px)
- Icon container: 48×48px, `--radius-sm`, gradient background

**Transition:** `transform var(--duration-normal) var(--ease-standard), border-color var(--duration-fast) var(--ease-standard), box-shadow var(--duration-normal) var(--ease-standard)`

---

### 6.3 ReelCard

Primary feed unit. Full-screen on mobile, constrained (400px max-width) on desktop.

```
┌──────────────────────────────┐  ← reel-max-width: 400px
│                              │
│      VIDEO CONTENT           │  aspect-ratio: 9/16
│                              │
│   ┌──────────────────────┐   │
│   │ LOADER SPINNER       │   │  ← shown during buffering
│   └──────────────────────┘   │
│                              │
│   ┌──────────────────────┐   │
│   │  PLAY OVERLAY  ▶     │   │  ← shown when paused
│   └──────────────────────┘   │
│                              │
│  ┌─────────────────────────┐ │
│  │ gradient overlay (bottom│ │  ← linear-gradient(transparent → #000 80%)
│  │                         │ │
│  │ @creator_name  [Follow] │ │  ← type-label, pill button
│  │ Caption text…           │ │  ← type-body-sm, truncate 2 lines
│  │ ♫ Original Audio Name   │ │  ← type-caption with music icon
│  └─────────────────────────┘ │
│                         [♥]  │  ← right-side action column
│                         [💬] │
│                         [↗]  │
│                         [🔖] │
│  [▓▓▓▓▓▓▓▓░░░░░░░░░░░░░░]  │  ← progress bar, 3px height
└──────────────────────────────┘
```

**Five States:**

| State | Appearance | Notes |
|---|---|---|
| **Loading** | Dark bg + centered spinner animation | Show after 150ms of buffering to avoid flash |
| **Playing** | Video visible, overlay fades after 2s of no interaction | Double-tap triggers heart burst |
| **Paused** | Play overlay visible (56px icon, 50% opacity bg disk) | Tap to resume |
| **Error** | Icon + "Video unavailable" + retry button centered | Log error to telemetry |
| **Complete** | Loop restarts automatically OR next reel shown | No explicit visual state needed |

**Gestures (mobile):**
- **Single tap:** Toggle play/pause
- **Double-tap:** Trigger like + heart burst animation
- **Swipe up/down:** Next/previous reel (100dvh scroll-snap unit)
- **Long-press:** Open context menu (Save, Report, Share)

---

### 6.4 ActionButton (Like / Comment / Share / Bookmark)

```
  ┌──────┐
  │  ♥   │  ← icon, 24px
  │ 12.4K│  ← count, type-caption
  └──────┘
  48×48px touch target
```

**States:**
- **Default:** `--text-secondary` icon
- **Hover:** `--text-primary`, scale(1.1)
- **Active/Pressed:** scale(0.92), then spring back
- **Liked:** Heart filled, color `--color-orange-500`, scale pulse animation
- **Loading:** Spinner, icon hidden

**Like animation sequence:**
1. Tap: `scale(0.8)` (80ms, `--ease-exit`)
2. Release: `scale(1.3)` (120ms, `--ease-spring`)
3. Settle: `scale(1.0)` (150ms, `--ease-standard`)
4. Color transitions to `--color-orange-500` concurrently (250ms)

**Heart Burst (double-tap):** 8 hearts animate outward from tap point. Each: `opacity 0→1→0`, `translateY: 0 → -80px to -160px` (randomized), `scale: 0.5→1.0→0.5`, stagger 30–60ms between each, duration 800ms total.

---

### 6.5 SidebarNav (Feed Page)

```
┌──────┐  ← 72px wide, full viewport height, fixed
│  [●] │  ← logo, 32px
│      │
│  [🏠]│  ← NavItem, 40px icon, 44×44px hit area
│  [🔍]│
│  [🎬]│
│  [💬]│
│  [🔔]│  ← badge on notification
│      │
│  [✚] │  ← create button, accent gradient
│      │
│ ─────│  ← divider
│ [av] │  ← user avatar, 32px
│ [⚙]  │
└──────┘
```

**NavItem States:**
- **Default:** Icon `--text-tertiary`, no background
- **Hover:** Icon `--text-secondary`, `--bg-glass` background disk (40px circle)
- **Active/Current:** Icon `--text-primary`, accent gradient disk
- **Notification badge:** 8px circle, `--color-orange-500`, top-right of icon
- **Tooltip (desktop):** Label appears as tooltip on hover, 200ms delay

**Keyboard nav:** Arrow keys cycle through items. `Enter`/`Space` activates. `Tab` for focus management.

---

### 6.6 CommentsPanel

Slide-up drawer covering ~70% of screen height. Behind: video continues playing at 30% opacity.

```
Entry: translateY(100%) → translateY(0), duration 350ms, --ease-enter
Exit:  translateY(0) → translateY(100%), duration 250ms, --ease-exit
Backdrop: opacity 0 → 0.4, duration 200ms
```

**Focus trap:** When open, Tab cycles within panel only. Escape closes. Focus returns to trigger button on close.

**ARIA:** `role="dialog"`, `aria-modal="true"`, `aria-labelledby="comments-title"`.

**Comment Item:**
```
[avatar 32px]  @username  12m ago
               Comment text, up to 3 lines before "show more"
               [Reply]  [♥] 24
```

---

### 6.7 ShareModal

Sheet modal, centered on desktop, bottom sheet on mobile.

**Entry desktop:** `scale(0.95) + opacity(0)` → `scale(1) + opacity(1)`, 250ms, `--ease-spring`  
**Entry mobile:** `translateY(100%)` → `translateY(0)`, 350ms, `--ease-enter`  
**Backdrop:** `rgba(0,0,0,0.6)`, click-outside to dismiss.

**Share grid:** 3×2 grid of platform icons. Each: 56×56px icon container, `--radius-lg`. Label below, `type-micro`. Hover: scale(1.05), `--bg-elevated`.

**Platforms:** WhatsApp, Telegram, Instagram, Twitter/X, Copy Link, More.

**ARIA:** `role="dialog"`, `aria-label="Share video"`. Platform buttons have `aria-label="Share to [Platform Name]"`.

---

### 6.8 ToastNotification

```
┌────────────────────────────────┐
│  [✓]  Link copied to clipboard │
└────────────────────────────────┘
```

**Position:** Bottom-center, 24px from bottom edge.  
**Entry:** `translateY(100%) + opacity(0)` → resting, 300ms `--ease-spring`  
**Auto-dismiss:** After 2500ms, exit `translateY(8px) + opacity(0)`, 200ms `--ease-exit`  
**Max-width:** 320px  
**Stacking:** Multiple toasts stack vertically, newest on top, older shift down.

**ARIA:** `role="status"`, `aria-live="polite"`, `aria-atomic="true"`.

---

### 6.9 PhoneMockup

Decorative component. Displays UI screenshots within a device frame.

- **Frame:** `--border-default`, `--radius-8` (32px), proportional to real device aspect
- **Screen aspect:** 19.5:9
- **Animation:** `float` keyframe — subtle translateY oscillation (±10px), 3s duration, `ease-in-out`, infinite
- **Stagger:** Middle phone offset by 1.5s delay, right phone by 0.8s delay

**States:** Static (no empty/loading/error — purely presentational)

---

### 6.10 HeroStats Bar

Three stats displayed horizontally with a filled progress bar each.

```
[███████░░░░] 325M Monthly Users
[████████████] 2.5B Monthly Shares  
[██████░░░░░] 15 Languages
```

- Bar: height 6px, `--radius-pill`, background `--bg-elevated`, fill `--gradient-primary`
- Animation on entry: fill from 0% → target width, 800ms, `--ease-standard`, staggered 150ms between bars
- Label: `type-label`, `--text-primary` for value, `--text-secondary` for description

---

## 7. Interaction & Animation Spec

### 7.1 Animation Principles

- **Purposeful:** Every animation communicates state change or provides spatial context.
- **Performant:** All animations use `transform` and `opacity` only (GPU-composited). No `width`/`height`/`top`/`left` animations.
- **Reducible:** All animations > 250ms have a `@media (prefers-reduced-motion: reduce)` override that sets duration to 1ms or swaps motion for a simple fade.

### 7.2 Keyframe Definitions

```css
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}
/* Usage: scroll-reveal entries. Duration 400ms, stagger 100ms */

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50%       { transform: translateY(-10px); }
}
/* Usage: PhoneMockup decorative. Duration 3s, infinite */

@keyframes pulseOpacity {
  0%, 100% { opacity: 1; }
  50%       { opacity: 0.4; }
}
/* Usage: Live indicator badge. Duration 1.5s, infinite */

@keyframes progress {
  from { width: 0%; }
  to   { width: 100%; }
}
/* Usage: Video progress bar. Duration = video duration */

@keyframes spinCW {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}
/* Usage: Loading spinners. Duration 800ms, infinite, linear */

@keyframes shimmer {
  0%   { background-position: -200% center; }
  100% { background-position:  200% center; }
}
/* Usage: Skeleton loading. Duration 1.5s, infinite */

@keyframes liveComment {
  0%   { opacity: 0; transform: translateY(0) scale(0.8); }
  15%  { opacity: 1; transform: translateY(-10px) scale(1); }
  80%  { opacity: 1; transform: translateY(-60px); }
  100% { opacity: 0; transform: translateY(-80px); }
}
/* Usage: Floating live comments. Duration 3.5s */

@keyframes heartBurst {
  0%   { opacity: 1; transform: translateY(0) scale(0.5) rotate(var(--r)); }
  50%  { opacity: 1; transform: translateY(var(--dy)) scale(1.0) rotate(var(--r)); }
  100% { opacity: 0; transform: translateY(var(--dy-max)) scale(0.5) rotate(var(--r)); }
}
/* Usage: Heart burst. Each heart has CSS vars --r (rotation), --dy, --dy-max */

@keyframes actionPulse {
  0%   { transform: scale(1); }
  40%  { transform: scale(1.3); }
  70%  { transform: scale(0.95); }
  100% { transform: scale(1); }
}
/* Usage: Like button activation. Duration 400ms */
```

### 7.3 Scroll Reveal (Intersection Observer)

- **Trigger:** Element intersects at 15% threshold (rootMargin: `0px 0px -50px 0px`)
- **Animation:** `fadeUp`, 500ms, `--ease-standard`
- **Stagger:** Sibling elements delay +100ms each (max 400ms total stagger)
- **Once only:** `unobserve()` after first trigger — elements don't re-animate on scroll back

### 7.4 Video Scroll-Snap

```
Container:
  scroll-snap-type: y mandatory;
  overflow-y: scroll;
  height: 100dvh;

Each reel:
  scroll-snap-align: start;
  height: 100dvh;
  scroll-snap-stop: always;  /* prevent skipping on fast swipe */
```

**Autoplay policy:** Muted autoplay on entry. Pause when out of viewport (IntersectionObserver).

### 7.5 Focus Management

- Focus ring: `outline: 2px solid var(--border-focus)`, `outline-offset: 2px`
- Remove `:focus` outline for mouse users (`:focus:not(:focus-visible)`)
- Keyboard shortcut hints displayed in UI where applicable (e.g., `↑↓` for feed nav)

---

## 8. Responsive Behavior

### 8.1 Navigation

| Breakpoint | Behavior |
|---|---|
| `xs`–`md` | Horizontal nav hidden; hamburger menu (future implementation) |
| `lg`+ | Full horizontal nav bar with all links |

### 8.2 Feed Layout

| Breakpoint | Left Sidebar | Center Feed | Right Sidebar |
|---|---|---|---|
| `xs`–`sm` | Hidden | Full-width (100vw) | Hidden |
| `md` | Visible (72px) | Max 520px, centered | Hidden |
| `lg` | Visible (72px) | Max 520px, centered | Visible (280px) |
| `xl`+ | Visible (72px) | Max 520px, centered | Visible (320px) |

### 8.3 Phone Showcase (Landing)

| Breakpoint | Layout |
|---|---|
| `xs`–`sm` | Single phone, horizontal scroll carousel |
| `md` | Two phones side by side |
| `lg`+ | Three phones, staggered vertical offsets |

### 8.4 Feature Cards Grid

| Breakpoint | Columns |
|---|---|
| `xs` | 1 |
| `sm` | 1 |
| `md` | 2 |
| `lg`+ | 3 |

### 8.5 Impact Metrics

| Breakpoint | Columns |
|---|---|
| `xs`–`sm` | 2 |
| `md`+ | 4 |

### 8.6 Typography Fluid Sizing

```css
/* Hero headline */
font-size: clamp(40px, 7vw, 88px);

/* Section h2 */
font-size: clamp(28px, 4vw, 44px);

/* Body text: no clamp, fixed 16px / 14px for readability */
```

---

## 9. Accessibility Spec

### 9.1 Color Contrast

All text/background combinations must meet WCAG 2.2 AA (4.5:1 for body text, 3:1 for large text ≥18px bold or ≥24px).

| Foreground | Background | Ratio | Level |
|---|---|---|---|
| `--text-primary` (#fff) | `--bg-body` (#0a0a0a) | 19.6:1 | AAA |
| `--text-secondary` (rgba 60%) | `--bg-body` | ~7.2:1 | AA |
| `--text-tertiary` (rgba 35%) | `--bg-body` | ~3.8:1 | Fails AA for small text |
| `--color-orange-500` (#FF6B35) | `--bg-body` | 3.9:1 | AA large text only |
| White text on `--gradient-primary` | — | ~4.8:1 (at lightest point) | AA |

**Note:** `--text-tertiary` must not be used for informational text. Restrict to decorative/placeholder only. Orange (#FF6B35) must not carry critical information alone without supporting non-color indicator.

### 9.2 Keyboard Interaction Map

| Component | Keys | Behavior |
|---|---|---|
| NavBar / SidebarNav | `Tab` | Move between items |
| NavItem | `Enter` | Activate |
| Button | `Enter`, `Space` | Activate |
| ReelCard (focused) | `Space` | Play/Pause |
| ReelCard (focused) | `↑` / `↓` | Previous / Next reel |
| ReelCard (focused) | `l` | Toggle like |
| ReelCard (focused) | `c` | Open comments |
| ReelCard (focused) | `s` | Open share |
| CommentsPanel | `Escape` | Close, return focus to trigger |
| CommentsPanel | `Tab` | Cycle within panel (trapped) |
| ShareModal | `Escape` | Close, return focus |
| ShareModal | `Tab` | Cycle platforms |

### 9.3 ARIA Landmarks & Roles

```html
<!-- Landing Page Structure -->
<header role="banner">…</header>
<nav aria-label="Main navigation">…</nav>
<main id="main-content">
  <section aria-labelledby="hero-heading">…</section>
  <section aria-labelledby="features-heading">…</section>
  <!-- etc -->
</main>

<!-- Feed Page -->
<nav aria-label="App navigation">…</nav>
<main aria-label="Video feed">
  <article aria-label="Video by @creator_name">
    <video aria-label="Video content" muted playsinline>…</video>
    …
  </article>
</main>
<aside aria-label="Suggested creators">…</aside>

<!-- Dialogs -->
<div role="dialog" aria-modal="true" aria-labelledby="comments-title">…</div>
```

### 9.4 Live Regions

```html
<!-- Toast notifications -->
<div role="status" aria-live="polite" aria-atomic="true" class="toast-region">
  <!-- Injected by JS: "Link copied to clipboard" -->
</div>

<!-- Like count updates -->
<!-- Do NOT put like counts in aria-live — too chatty. -->
<!-- Instead: announce once on action: "Video liked. 12,401 likes." -->
<div role="status" aria-live="polite" class="sr-only" id="action-announce"></div>

<!-- Video playback state -->
<!-- Announce when video pauses/resumes via aria-live on the player container -->
```

### 9.5 Reduced Motion

```css
@media (prefers-reduced-motion: reduce) {
  *,
  *::before,
  *::after {
    animation-duration: 1ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 1ms !important;
    scroll-behavior: auto !important;
  }

  /* Keep float animation entirely off for decorative phone mockups */
  .phone-frame {
    animation: none !important;
  }

  /* Disable heart burst entirely */
  .heart-burst-container {
    display: none !important;
  }
}
```

### 9.6 Alternative Text

- All icon-only buttons: `aria-label="[action name]"`
- Decorative images/gradients: `aria-hidden="true"`
- Creator avatars: `alt="@[username]'s profile photo"`
- Video: `aria-label="Video by @[creator], [description truncated to 100 chars]"`
- For users who prefer no autoplay: provide a visible play button overlay by default; autoplay only if no `prefers-reduced-motion` and user hasn't disabled it.

---

## 10. State Definitions (Five-State Model)

Each component that handles data or user interaction is defined across five explicit states. Engineers must implement all five.

### 10.1 Feed / ReelCard States

| State | Visual | Notes |
|---|---|---|
| **Empty** | "No more videos" illustration, centered. CTA: "Explore more creators." | Shown when feed runs out of content |
| **Loading** | Dark bg + centered spinner (32px). Delay 150ms before showing. | Do not flash spinner for fast loads |
| **Error** | Warning icon + "Couldn't load this video" + "Retry" button | Log error; retry fetches the same reel |
| **Partial** | Video plays; overlay metadata "Loading…" shimmer placeholders | Creator name/description loads async |
| **Complete** | Full video + overlay info visible | Happy path |

### 10.2 CommentsPanel States

| State | Visual |
|---|---|
| **Empty** | Centered icon + "No comments yet. Be the first!" |
| **Loading** | 3× skeleton comment rows (shimmer) |
| **Error** | "Couldn't load comments" + retry link |
| **Partial** | First 10 comments visible + "Load more" at bottom |
| **Complete** | All comments visible, pagination complete |

### 10.3 Like Button State

| State | Visual |
|---|---|
| **Default (unliked)** | Outlined heart, `--text-secondary` |
| **Optimistic (liked)** | Filled heart, `--color-orange-500`, pulse animation |
| **Error (like failed)** | Revert to unliked + toast "Couldn't like. Try again." |
| **Loading** | Spinner in place of icon (if first like takes >500ms) |
| **Disabled** | Outlined heart, `--text-disabled`, no interaction |

### 10.4 SidebarNav States

| State | Visual |
|---|---|
| **Default** | All icons `--text-tertiary` |
| **Hover** | Icon `--text-secondary`, glass disk background |
| **Active** | Icon `--text-primary`, accent gradient disk |
| **Notification** | Orange badge dot (8px) top-right of icon |
| **Collapsed** | Not applicable (sidebar is fixed width) |

### 10.5 FeatureCard States

| State | Visual |
|---|---|
| **Default** | Flat glass card, icon, text |
| **Hover** | Float up 4px, glow border, shadow |
| **Focus** | Orange outline ring |
| **Loading** | Shimmer skeleton (icon placeholder + 2 text rows) |
| **Error** | Not applicable (static content) |

---

## Appendix A — CSS Variable Reference (Quick Lookup)

```
Semantic Token           → Value
──────────────────────────────────────────────────
--bg-body                → #0a0a0a
--bg-surface             → #111111
--bg-card                → #1a1a1a
--bg-glass               → rgba(255,255,255,0.06)

--text-primary           → #ffffff
--text-secondary         → rgba(255,255,255,0.60)
--text-tertiary          → rgba(255,255,255,0.35)
--text-accent            → #FF6B35

--border-default         → rgba(255,255,255,0.08)
--border-focus           → #FF6B35
--border-error           → #FF4444

--gradient-primary       → linear-gradient(135deg,#FF6B35,#FF3CAC)
--gradient-aurora        → linear-gradient(135deg,#FF3CAC,#784BA0,#2B86C5)

--radius-sm              → 12px
--radius-md              → 16px
--radius-lg              → 24px
--radius-xl              → 32px
--radius-pill            → 9999px

--spacing-xs / sm / md / lg / xl / 2xl / 3xl
                         → 4 / 8 / 16 / 24 / 32 / 48 / 64px

--duration-fast          → 150ms
--duration-normal        → 250ms
--duration-slow          → 400ms

--ease-standard          → cubic-bezier(0.25,0.46,0.45,0.94)
--ease-spring            → cubic-bezier(0.34,1.56,0.64,1.00)
--ease-exit              → cubic-bezier(0.55,0.00,1.00,0.45)
```

---

## Appendix B — Font Loading Snippet

```html
<!-- In <head>, before any CSS -->
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link rel="preload" as="style"
  href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Space+Grotesk:wght@600;700&display=swap">
<link rel="stylesheet"
  href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Space+Grotesk:wght@600;700&display=swap"
  media="print" onload="this.media='all'">
<noscript>
  <link rel="stylesheet" href="…same href…">
</noscript>
```

This pattern (Richard Rutter's FOUT-safe async load) prevents render blocking while ensuring fonts display on JS-disabled clients.

---

*End of specification. Questions → tag @aisha in design review.*
