# ShareChat Revamp — UI/UX Specification

**Version:** 1.0  
**Date:** 2026-04-14  
**Author:** Aisha (Design)  
**Status:** Ready for Engineering Handoff

---

## Table of Contents

1. [Design Tokens](#1-design-tokens)
2. [Typography System](#2-typography-system)
3. [Component Hierarchy](#3-component-hierarchy)
4. [Layout Wireframes](#4-layout-wireframes)
5. [Component Specifications](#5-component-specifications)
6. [Interaction Specifications](#6-interaction-specifications)
7. [Responsive Behavior](#7-responsive-behavior)
8. [Accessibility Specification](#8-accessibility-specification)
9. [Animation System](#9-animation-system)

---

## 1. Design Tokens

All components reference tokens. Never hardcode raw values in component code.

### 1.1 Primitive Tokens

```css
/* ── Color Primitives ── */
--color-black:       #000000;
--color-near-black:  #0a0a0a;
--color-gray-950:    #111111;
--color-gray-900:    #1a1a1a;
--color-gray-800:    #222222;
--color-gray-700:    #333333;
--color-gray-500:    rgba(255,255,255,0.35);
--color-gray-400:    rgba(255,255,255,0.50);
--color-gray-300:    rgba(255,255,255,0.60);
--color-white:       #ffffff;

--color-orange-500:  #FF6B35;
--color-pink-500:    #FF3CAC;
--color-purple-500:  #784BA0;
--color-blue-500:    #2B86C5;
--color-cyan-400:    #00D9F5;
--color-mint-400:    #00F5A0;
--color-gold-400:    #FFD200;
--color-gold-500:    #F7971E;
--color-red-500:     #FF4757;

/* ── Opacity Primitives ── */
--opacity-glass-bg:     0.06;
--opacity-glass-border: 0.08;
--opacity-overlay-dark: 0.7;
--opacity-overlay-mid:  0.4;

/* ── Space Primitives (4px grid) ── */
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
--space-24: 96px;

/* ── Radius Primitives ── */
--radius-2:  8px;
--radius-3:  12px;
--radius-4:  16px;
--radius-6:  24px;
--radius-8:  32px;
--radius-full: 9999px;

/* ── Duration Primitives ── */
--dur-instant: 80ms;
--dur-fast:    150ms;
--dur-normal:  250ms;
--dur-slow:    400ms;
--dur-slower:  600ms;
--dur-slowest: 800ms;

/* ── Easing Primitives ── */
--ease-default: cubic-bezier(0.25, 0.46, 0.45, 0.94);
--ease-spring:  cubic-bezier(0.34, 1.56, 0.64, 1.0);
--ease-in:      cubic-bezier(0.4, 0, 1, 1);
--ease-out:     cubic-bezier(0, 0, 0.2, 1);

/* ── Shadow Primitives ── */
--shadow-sm:  0 2px 8px rgba(0,0,0,0.3);
--shadow-md:  0 8px 24px rgba(0,0,0,0.4);
--shadow-lg:  0 16px 48px rgba(0,0,0,0.5);
--shadow-glow-accent: 0 8px 32px rgba(255,107,53,0.3);
--shadow-glow-pink:   0 8px 32px rgba(255,60,172,0.3);
```

### 1.2 Semantic Tokens

```css
/* ── Backgrounds ── */
--bg-page:      var(--color-near-black);   /* Base page background */
--bg-surface:   var(--color-gray-950);     /* Cards, panels */
--bg-elevated:  var(--color-gray-900);     /* Modals, dropdowns */
--bg-glass:     rgba(255,255,255,0.06);    /* Glassmorphic overlays */
--bg-overlay:   rgba(0,0,0,0.70);         /* Dimming overlays */

/* ── Borders ── */
--border-default: rgba(255,255,255,0.08);
--border-subtle:  rgba(255,255,255,0.05);
--border-focus:   var(--color-orange-500);

/* ── Text ── */
--text-primary:   var(--color-white);
--text-secondary: rgba(255,255,255,0.60);
--text-tertiary:  rgba(255,255,255,0.35);
--text-disabled:  rgba(255,255,255,0.25);
--text-inverse:   var(--color-near-black);

/* ── Brand / Accent ── */
--color-primary:    var(--color-orange-500);
--color-primary-glow: rgba(255,107,53,0.30);
--gradient-brand:   linear-gradient(135deg, #FF6B35, #FF3CAC);
--gradient-vibrant: linear-gradient(135deg, #FF3CAC, #784BA0, #2B86C5);
--gradient-cool:    linear-gradient(135deg, #00F5A0, #00D9F5);
--gradient-warm:    linear-gradient(135deg, #F7971E, #FFD200);

/* ── Status ── */
--color-success: #00F5A0;
--color-warning: #FFD200;
--color-error:   #FF4757;
--color-info:    #00D9F5;

/* ── Spacing Aliases ── */
--spacing-xs:  var(--space-1);   /* 4px  */
--spacing-sm:  var(--space-2);   /* 8px  */
--spacing-md:  var(--space-4);   /* 16px */
--spacing-lg:  var(--space-6);   /* 24px */
--spacing-xl:  var(--space-8);   /* 32px */
--spacing-2xl: var(--space-12);  /* 48px */

/* ── Radius Aliases ── */
--radius-sm:  var(--radius-3);   /* 12px */
--radius-md:  var(--radius-4);   /* 16px */
--radius-lg:  var(--radius-6);   /* 24px */
--radius-xl:  var(--radius-8);   /* 32px */
--radius-pill: var(--radius-full);

/* ── Transition Aliases ── */
--transition-fast:   all var(--dur-fast) var(--ease-default);
--transition-normal: all var(--dur-normal) var(--ease-default);
--transition-slow:   all var(--dur-slow) var(--ease-default);
```

### 1.3 Component Tokens

```css
/* ── Sidebar ── */
--sidebar-width:       72px;
--sidebar-bg:          var(--bg-surface);
--sidebar-border:      var(--border-default);
--sidebar-icon-size:   22px;
--sidebar-avatar-size: 36px;

/* ── Right Panel ── */
--right-panel-width:   280px;
--right-panel-bg:      var(--bg-surface);

/* ── Reel Card ── */
--reel-max-width:      420px;
--reel-aspect-ratio:   9/16;
--reel-radius:         var(--radius-lg);

/* ── Story Ring ── */
--story-ring-size:     56px;
--story-ring-border:   2.5px;

/* ── Action Button ── */
--action-btn-size:     44px;   /* Touch target minimum */
--action-icon-size:    24px;
--action-label-size:   11px;

/* ── Navigation Arrow ── */
--nav-arrow-size:      44px;
--nav-arrow-bg:        var(--bg-glass);
--nav-arrow-border:    var(--border-default);

/* ── Toast ── */
--toast-width:   max-content;
--toast-max-w:   320px;
--toast-padding: var(--space-3) var(--space-4);
--toast-radius:  var(--radius-pill);
--toast-bg:      var(--bg-elevated);

/* ── Modal Panel ── */
--panel-max-height: 80vh;
--panel-radius:     var(--radius-xl) var(--radius-xl) 0 0;
--panel-bg:         var(--bg-elevated);
--panel-handle-w:   40px;
--panel-handle-h:   4px;
```

---

## 2. Typography System

Font loading order: `Space Grotesk` (display) → `Inter` (UI) → system fallbacks.

```
Font stack (display): 'Space Grotesk', -apple-system, BlinkMacSystemFont, sans-serif
Font stack (UI):      'Inter', -apple-system, BlinkMacSystemFont, sans-serif
```

### 2.1 Type Scale

| Token            | Size / Line-H / Track | Weight | Font        | Usage                          |
|------------------|-----------------------|--------|-------------|--------------------------------|
| `--type-hero`    | clamp(48px,7vw,88px) / 1.0 / -0.03em | 700 | Space Grotesk | Hero headline                |
| `--type-display` | clamp(32px,4vw,52px) / 1.1 / -0.02em | 700 | Space Grotesk | Section titles               |
| `--type-title`   | clamp(24px,3vw,36px) / 1.2 / -0.01em | 600 | Space Grotesk | Card/modal titles            |
| `--type-heading` | 20px / 28px / -0.01em | 600 | Inter | Widget headings, subheadings       |
| `--type-body-lg` | 16px / 24px / 0em     | 400 | Inter | Primary body text                  |
| `--type-body`    | 14px / 22px / 0em     | 400 | Inter | Default UI body                    |
| `--type-body-sm` | 13px / 20px / 0em     | 400 | Inter | Secondary body, comments           |
| `--type-label`   | 12px / 16px / 0.01em  | 500 | Inter | Labels, tags, badges               |
| `--type-caption` | 11px / 16px / 0.01em  | 400 | Inter | Timestamps, meta, counts           |
| `--type-overline`| 11px / 16px / 0.15em  | 500 | Inter | Section overlines (ALL CAPS)       |

---

## 3. Component Hierarchy

```
App Root
├── GlobalLayout
│   ├── Sidebar (left, 72px)
│   │   ├── SidebarLogo
│   │   ├── SidebarNav
│   │   │   └── SidebarNavItem (× 6)
│   │   ├── SidebarDivider
│   │   ├── SidebarNavItem [settings]
│   │   ├── SidebarNavItem [analytics]
│   │   └── SidebarAvatar
│   ├── MainContent (flex-1)
│   │   ├── FeedHeader (sticky)
│   │   │   ├── StoriesBar
│   │   │   │   └── StoryRing (× n)
│   │   │   └── FeedTabs [optional]
│   │   └── ReelFeed (scroll-snap container)
│   │       └── ReelCard (× n)
│   │           ├── VideoPlayer
│   │           │   ├── Loader
│   │           │   ├── PlayOverlay
│   │           │   ├── GradientOverlay
│   │           │   ├── LiveComments
│   │           │   ├── CreatorInfo
│   │           │   ├── ActionSidebar
│   │           │   │   ├── LikeButton
│   │           │   │   ├── CommentButton
│   │           │   │   ├── ShareButton
│   │           │   │   ├── SaveButton
│   │           │   │   └── MoreButton
│   │           │   └── ProgressBar
│   │           └── ReelDescription [desktop only]
│   └── RightPanel (280px, desktop only)
│       ├── UserProfileCard
│       │   ├── Avatar
│       │   ├── UserName
│       │   └── FollowButton
│       ├── SuggestedUsers
│       │   └── SuggestedUserRow (× 5)
│       │       ├── Avatar
│       │       ├── UserInfo
│       │       └── FollowButton
│       └── TrendingTags
│           └── TrendingTagRow (× 6)
├── NavigationArrows (fixed, right of feed)
│   ├── ArrowUp
│   └── ArrowDown
├── CommentsPanel (slide-up modal)
│   ├── PanelHandle
│   ├── PanelHeader
│   ├── CommentList
│   │   └── CommentRow (× n)
│   └── CommentInput
└── SharePanel (slide-up modal)
    ├── PanelHandle
    ├── PanelHeader
    ├── ShareGrid
    │   └── ShareOption (× 6)
    └── CopyLinkRow
```

---

## 4. Layout Wireframes

### 4.1 Desktop Layout (≥ 1100px)

```
┌──────────────────────────────────────────────────────────────────────┐
│ 72px │           Feed (flex-1, center-aligned)          │   280px   │
│      │                                                   │           │
│  [S] │   ┌─── Stories ──────────────────────────────┐   │  [User]   │
│      │   │ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ (h-scroll)      │   │  [Sugg]   │
│  [⊞] │   └──────────────────────────────────────────┘   │  [Trend]  │
│  [⊙] │                                                   │           │
│  [▷] │   ┌──────────────────────────────────────────┐   │           │
│  [✉] │   │                                          │   │           │
│  [🔔]│   │   ┌──── ReelCard (max 420px) ───────┐   │   │    ↑      │
│  [+] │   │   │                                 │   │   │    ↓      │
│      │   │   │      [VIDEO 9:16]                │   │   │           │
│  ─── │   │   │                                 │   │   │           │
│  [⚙] │   │   │  [live comments]                │   │   │           │
│  [📊]│   │   │  ─────────────────────────────  │   │   │           │
│      │   │   │  Creator ●  Caption  [🎵]       │   │   │           │
│  [●] │   │   │                      [❤]  231   │   │   │           │
│      │   │   └─────────────────────────────────┘   │   │           │
│      │   │                                          │   │           │
│      │   └──────────────────────────────────────────┘   │           │
└──────────────────────────────────────────────────────────────────────┘
                         ↑ arrows float right of feed ↑
```

### 4.2 Mobile Layout (< 768px)

```
┌─────────────────────────┐
│  Full-screen video feed  │
│  (no sidebar/right panel)│
│                          │
│  [VIDEO 9:16 100vh]      │
│                          │
│  [live comments]         │
│                          │
│  Creator ●  Caption      │
│                  [❤] 231 │
│                  [💬] 18 │
│                  [↗] Shr │
│                  [🔖] Sv │
│  ──────────────────────  │
│  ████████░░░░░░░░░░░░░   │  ← progress bar
│                          │
│ ┌──┬──┬──┬──┬──┐        │
│ │🏠│🔍│ + │✉ │👤│        │  ← bottom nav
│ └──┴──┴──┴──┴──┘        │
└─────────────────────────┘
```

### 4.3 Comments Panel (slide-up)

```
┌─────────────────────────┐
│░░░░░░░░░░░░░░░░░░░░░░░░░│  ← dim overlay (tap to close)
│░░░░░░░░░░░░░░░░░░░░░░░░░│
│░░░░░░░░░░░░░░░░░░░░░░░░░│
│                          │
│  ┌────────────────────┐  │
│  │     ────           │  │  ← drag handle
│  │  Comments (18)  ✕  │  │  ← header
│  │  ─────────────────  │  │
│  │  [●] User   12m    │  │
│  │      Great video!  │  │
│  │  [●] User2  1h     │  │  ← comment rows (scrollable)
│  │      👏 🔥         │  │
│  │        ·  ·  ·     │  │
│  │  ─────────────────  │  │
│  │  [●] [  Type...  ] │  │  ← input row (sticky bottom)
│  └────────────────────┘  │
└─────────────────────────┘
```

### 4.4 Share Panel (slide-up)

```
┌─────────────────────────┐
│░░░░░░░░░░░░░░░░░░░░░░░░░│
│                          │
│  ┌────────────────────┐  │
│  │       ────         │  │  ← handle
│  │   Share  ·  ·  ·   │  │
│  │  ─────────────────  │  │
│  │  [WA] [TG] [IG][TW] │  │
│  │  [FB] [Copy link]   │  │
│  │  ─────────────────  │  │
│  │  🔗 [url…]  [Copy] │  │  ← copy row
│  └────────────────────┘  │
└─────────────────────────┘
```

---

## 5. Component Specifications

### 5.1 ReelCard

**States:** Loading → Partial (video buffering) → Complete → Paused → Error

```
State: Loading
  - Video element hidden
  - Spinner visible (centered, 48px, gradient stroke)
  - Background: --bg-surface

State: Partial (buffering mid-play)
  - Video visible with opacity 0.6
  - Small spinner overlay (24px, bottom-right of video)
  - Progress bar paused/pulsing

State: Complete (playing)
  - Video at full opacity
  - Play overlay hidden
  - Progress bar advancing
  - Live comments cycling

State: Paused
  - Video paused
  - Play overlay visible (opacity 1, transition: var(--transition-fast))
  - Icon: ▶ (48px, white, centered)
  - Live comments frozen

State: Error
  - Video hidden
  - Error icon (⚠ 32px, --color-warning)
  - Copy: "Could not load video"
  - Sub-copy: "Tap to retry"
  - Retry button (--btn-secondary)
```

**Visual structure:**

| Region                    | Size                  | Token Reference                     |
|---------------------------|-----------------------|-------------------------------------|
| Container                 | max-width: 420px; aspect: 9/16 | --reel-max-width, --reel-aspect-ratio |
| Border radius             | 24px                  | --radius-lg                         |
| Gradient overlay          | linear-gradient(0deg, rgba(0,0,0,0.85) 0%, transparent 50%) | Bottom half only |
| Creator name              | 14px / 500 / white    | --type-body                         |
| Caption text              | 13px / 400 / white    | --type-body-sm                      |
| Music badge               | 11px / Inter          | --type-caption                      |
| Progress bar height       | 2px                   | Fills width at bottom               |

**Interaction:**
- Single tap: toggle play/pause
- Double tap: trigger like animation (heart burst + floating hearts)
- Swipe up/down (mobile): advance to next/previous reel
- Arrow keys (desktop): advance to next/previous reel

---

### 5.2 ActionSidebar

Vertical strip of action buttons pinned to the right side of the video.

**Specs:**

| Button   | Icon       | Touch target | Label font     | State change           |
|----------|------------|--------------|----------------|------------------------|
| Like     | ❤          | 44×44px      | --type-caption | Filled red + burst anim |
| Comment  | 💬         | 44×44px      | --type-caption | Panel slides up        |
| Share    | ↗          | 44×44px      | --type-caption | Share panel slides up  |
| Save     | 🔖         | 44×44px      | --type-caption | Filled + scale anim    |
| More     | ···        | 44×44px      | —              | Context menu (future)  |

- Gap between buttons: `--space-5` (20px)
- Icon size: 24px
- Label: below icon, `--type-caption`, `--text-secondary`
- Active state: icon `scale(1.2)` over `80ms var(--ease-spring)`
- Like button liked state: color `--color-error`, `scale(1.3)` pulse then `scale(1.0)`

---

### 5.3 LikeButton (detailed micro-interaction)

**Interaction timeline:**

```
t=0ms      User taps ❤
t=0ms      Heart icon: scale(1.3), color → --color-error
t=80ms     Heart icon: scale(1.0)
t=0ms      8× floating hearts spawn at tap point
  t=0ms    heart[0]: opacity 1 → 0, translateY(0 → -160px), scale(0.6 → 1.0), duration 1300ms
  t=100ms  heart[1]: same, offset X ±random(20px)
  t=200ms  heart[2..7]: each 100ms staggered
t=0ms      Count +1 (optimistic update), animate: translateY(-4px → 0, 150ms)
t=800ms    All floating hearts removed from DOM
```

**Double-tap trigger (large heart burst):**
```
Center of video: large ❤ (80px) appears
  opacity: 0 → 1 → 0  (0–200ms rise, 200–800ms fade)
  scale: 0.5 → 1.2 → 1.0
  Easing: var(--ease-spring)
```

---

### 5.4 StoryRing

**States:** Empty, Unseen, Seen, Loading, Error

```
State: Unseen
  Border: 2.5px solid gradient-brand (orange → pink)
  Avatar: 48px circle, no gap

State: Seen
  Border: 2.5px solid --border-default
  Avatar: 48px circle

State: Loading (uploading own story)
  Border: 2.5px animated gradient sweep (spin)
  Avatar: 48px, opacity 0.7

State: Own (add story)
  Avatar: 48px circle
  Plus icon overlay: 18px, bottom-right, --bg-elevated border

State: Error
  Border: 2.5px solid --color-error
  Small ⚠ overlay at bottom-right
```

**Dimensions:**
- Outer ring: 56px × 56px
- Avatar inside ring: 48px (4px gap each side)
- Username label: 10px, max-width 56px, overflow ellipsis, centered below
- Touch target: full 56px ring + 24px label area = ~80px tall
- Horizontal gap between rings: `--space-3` (12px)

---

### 5.5 SidebarNavItem

**States:** Default, Hover, Active, Badged

```
State: Default
  Icon: 22px, --text-secondary
  Background: transparent
  Width: 44px, Height: 44px, border-radius: --radius-md

State: Hover
  Background: --bg-glass
  Icon: --text-primary
  Transition: var(--transition-fast)

State: Active
  Background: rgba(255,107,53,0.15)
  Icon: --color-primary (orange)
  Left border accent: 2px solid --color-primary (optional)

State: Badged (notifications)
  Badge: 8px × 8px circle, --color-error, top-right of icon
  Pulse animation: scale 1.0 → 1.3 → 1.0, 2s infinite
```

---

### 5.6 FollowButton

**States:** Default (unfollowed), Loading, Following, Error

```
State: Default
  Label: "Follow"
  Border: 1px solid --color-primary
  Color: --color-primary
  Padding: 6px 16px
  Border-radius: --radius-pill
  Font: 13px / 500

State: Loading
  Label: spinner (12px, --color-primary)
  Pointer-events: none
  Opacity: 0.7

State: Following
  Label: "Following"
  Background: --bg-glass
  Border: 1px solid --border-default
  Color: --text-secondary

State: Error
  Border: 1px solid --color-error
  Color: --color-error
  Label: "Retry"
```

---

### 5.7 CommentsPanel

**States:** Closed, Opening (slide-in), Open, Closing (slide-out)

```
State: Closed
  transform: translateY(100%)
  visibility: hidden

State: Opening
  transform: translateY(100%) → translateY(0)
  Duration: var(--dur-slow) — 400ms
  Easing: var(--ease-out)
  Overlay: opacity 0 → 0.7, same duration

State: Open
  transform: translateY(0)
  max-height: 80vh
  Overflow: auto (comment list only)

State: Closing (tap overlay or ✕)
  transform: translateY(0) → translateY(100%)
  Duration: var(--dur-normal) — 250ms
  Easing: var(--ease-in)
```

**Comment Row:**

| Element   | Spec                                                       |
|-----------|------------------------------------------------------------|
| Avatar    | 32px circle, --radius-full                                 |
| Username  | 13px / 500 / --text-primary                                |
| Timestamp | 11px / --text-tertiary                                     |
| Body      | 13px / 400 / --text-secondary, line-height 1.5             |
| Like icon | 14px ❤, right-aligned, toggleable, --text-tertiary → --color-error |

**Input Row (pinned to panel bottom):**

```
Layout: flex row, gap: --space-2
  - Avatar: 32px
  - Input: flex-1, background --bg-glass, border --border-default
            padding: 10px 16px, border-radius: --radius-pill
            placeholder: "Add a comment…"
  - Post button: text "Post", --color-primary, 13px/600
                 disabled when input empty (opacity 0.3, pointer-events: none)
```

---

### 5.8 Toast

**States:** Hidden, Enter, Visible, Exit

```
State: Enter
  transform: translateY(20px) → translateY(0)
  opacity: 0 → 1
  Duration: var(--dur-fast) — 150ms

State: Visible
  Auto-dismiss after 3000ms

State: Exit
  opacity: 1 → 0
  transform: translateY(0) → translateY(-8px)
  Duration: var(--dur-fast) — 150ms
```

**Layout:**
```
Background:   var(--bg-elevated)
Border:       1px solid var(--border-default)
Border-radius: var(--radius-pill)
Padding:      10px 16px
Font:         13px / 500 / --text-primary
Shadow:       var(--shadow-md)
Position:     fixed, bottom: 80px (above bottom nav on mobile), center-x
Z-index:      1000
```

---

### 5.9 ProgressBar

```
Container:
  Position: absolute, bottom: 0, left: 0, right: 0
  Height: 2px
  Background: rgba(255,255,255,0.20)

Fill:
  Background: var(--gradient-brand)
  Height: 100%
  Width: driven by video.currentTime / video.duration * 100%
  Transition: none (real-time, no CSS transition — JS updates width)

Scrub interaction (future):
  On hover: height expands to 4px (transition: height var(--dur-fast))
  Cursor: pointer
  On drag: seek video to position
```

---

## 6. Interaction Specifications

### 6.1 Scroll Snap Feed

```
Container:
  overflow-y: scroll
  scroll-snap-type: y mandatory
  -webkit-overflow-scrolling: touch

Each ReelCard:
  scroll-snap-align: start
  height: 100vh (mobile) / max(100vh, auto) (desktop)
```

**On snap settle:**
1. Pause previous video
2. Play current video (start from 0)
3. Preload next video (`preload="auto"`)
4. Set distant videos to `preload="none"`
5. Start live comments for current card
6. Update URL/history state (future)

---

### 6.2 Keyboard Navigation

| Key           | Action                                    |
|---------------|-------------------------------------------|
| `↓` / `j`     | Next reel                                 |
| `↑` / `k`     | Previous reel                             |
| `Space`       | Play / Pause                              |
| `l`           | Like current reel                         |
| `c`           | Open comments                             |
| `s`           | Share                                     |
| `Escape`      | Close open panel                          |
| `m`           | Mute / unmute                             |

Focus management:
- Feed container has `tabIndex="0"` to capture key events
- When a panel is open, focus moves into panel and is trapped there
- On panel close, focus returns to the action button that triggered it

---

### 6.3 Touch Gesture Map (Mobile)

| Gesture                   | Region          | Action                         |
|---------------------------|-----------------|--------------------------------|
| Single tap                | Video           | Play / Pause                   |
| Double tap                | Video           | Like + heart animation         |
| Swipe up                  | Video           | Next reel                      |
| Swipe down                | Video           | Previous reel                  |
| Swipe up (panel edge)     | Bottom 60px     | Open comments panel            |
| Long press                | Video           | Show options (future)          |
| Pinch                     | Video           | No-op (preserve default zoom)  |
| Tap                       | Overlay dim     | Close open panel               |
| Drag handle               | Panel handle    | Drag to dismiss panel          |

Swipe threshold: 60px vertical movement to register as swipe intent.

---

### 6.4 Video Playback

```
Auto-play policy:
  - First reel: muted autoplay (browser policy compliance)
  - User-initiated scroll: unmuted play
  - Background tab: pause all videos
  - Focus return: resume current video

Preload strategy:
  visible reel:   preload="auto", play()
  next reel:      preload="auto", paused
  reel after next: preload="metadata"
  all others:     preload="none"
```

---

## 7. Responsive Behavior

### 7.1 Breakpoints

| Token     | Value  | Description                                |
|-----------|--------|--------------------------------------------|
| `--bp-sm` | 480px  | Small mobile (landscape edge)              |
| `--bp-md` | 768px  | Tablet portrait / large mobile             |
| `--bp-lg` | 1100px | Desktop (right panel appears)              |
| `--bp-xl` | 1400px | Wide desktop (max content width)           |

### 7.2 Layout Changes by Breakpoint

**< 480px (Mobile S)**
- Full-screen video: 100vw × 100vh
- ActionSidebar: pinned right, 44px touch targets, no labels
- CreatorInfo: limited to 2 lines, truncated
- Stories bar hidden (above-fold priority = video)
- Bottom nav visible (56px height)

**480px – 767px (Mobile L)**
- Same as Mobile S with slight padding relaxation
- Stories bar becomes visible (horizontal scroll, 80px height)
- Caption expanded to 3 lines

**768px – 1099px (Tablet)**
- Sidebar appears (72px left)
- Video centered, max-width 360px
- Right panel hidden
- Bottom nav hidden
- ActionSidebar remains on video right edge

**≥ 1100px (Desktop)**
- Sidebar (72px) + centered feed + right panel (280px)
- Reel max-width 420px
- Navigation arrows appear (fixed, right of reel column)
- ReelDescription section visible below video
- ActionSidebar labels visible

**≥ 1400px (Wide)**
- Feed column gains extra breathing room
- Max content width 1280px, centered

### 7.3 Component-Level Responsive Rules

| Component         | Mobile                     | Tablet                     | Desktop                    |
|-------------------|----------------------------|----------------------------|----------------------------|
| Sidebar           | Hidden                     | 72px, icons only           | 72px, icons only           |
| RightPanel        | Hidden                     | Hidden                     | 280px visible              |
| BottomNav         | 5-tab, 56px                | Hidden                     | Hidden                     |
| StoriesBar        | Hidden or above feed       | Above feed                 | Above feed                 |
| ActionSidebar     | Right of video, no labels  | Right of video, no labels  | Right of video + labels    |
| NavigationArrows  | Hidden (swipe instead)     | Hidden                     | Visible, right of feed     |
| ReelDescription   | Hidden                     | Hidden                     | Below video                |
| CommentsPanel     | Slide-up 80vh              | Slide-up 80vh              | Slide-up 80vh              |
| Toast position    | bottom: 76px (above nav)   | bottom: 32px               | bottom: 32px               |

---

## 8. Accessibility Specification

### 8.1 Color Contrast Requirements

All foreground/background pairs must meet WCAG 2.1 AA (4.5:1 for text, 3:1 for UI).

| Foreground            | Background             | Ratio  | Level |
|-----------------------|------------------------|--------|-------|
| --text-primary (#fff) | --bg-page (#0a0a0a)    | 19.1:1 | AAA   |
| --text-secondary (60%)| --bg-page              | 5.9:1  | AA    |
| --text-tertiary (35%) | --bg-page              | 3.2:1  | AA (UI only) |
| --color-primary (#FF6B35) | --bg-page          | 4.6:1  | AA    |
| --color-error (#FF4757) | --bg-elevated        | 4.5:1  | AA    |
| --color-success (#00F5A0) | --bg-elevated      | 8.1:1  | AAA   |

Non-color indicators required (don't rely on color alone):
- Error state: red border + ⚠ icon + descriptive text
- Success state: green tint + ✓ icon + confirmation text
- Like state: filled icon shape (not just color change)

### 8.2 Focus Management

```
Focus ring:
  outline: 2px solid var(--border-focus)  /* --color-primary */
  outline-offset: 2px
  border-radius: matches element radius

Suppression (mouse only):
  :focus:not(:focus-visible) { outline: none; }

Trap boundaries:
  CommentsPanel: trap focus within when open
  SharePanel: trap focus within when open
  Restore: return focus to trigger button on close
```

### 8.3 ARIA Roles and Properties

| Component       | Role / Properties                                                    |
|-----------------|----------------------------------------------------------------------|
| Sidebar nav     | `<nav aria-label="Main navigation">`                                 |
| SidebarNavItem  | `<button aria-label="[name]" aria-current="page"` when active       |
| ReelFeed        | `role="feed" aria-busy="true/false"`                                 |
| ReelCard        | `role="article" aria-label="[creator] - [caption]"`                 |
| VideoPlayer     | `<video aria-label="[caption]">` + `<track kind="captions">`        |
| LikeButton      | `aria-label="Like [count] likes" aria-pressed="true/false"`         |
| CommentsPanel   | `role="dialog" aria-modal="true" aria-label="Comments"`             |
| SharePanel      | `role="dialog" aria-modal="true" aria-label="Share"`                |
| Toast           | `role="status" aria-live="polite" aria-atomic="true"`               |
| ProgressBar     | `role="progressbar" aria-valuenow aria-valuemin="0" aria-valuemax="100"` |
| FollowButton    | `aria-pressed="true/false"` `aria-label="Follow [username]"`        |

### 8.4 Screen Reader Announcements

```
Live regions:
  New live comment arrives:   silent (aria-live="off") — decorative
  Like count changes:         polite — "Liked. [n] likes."
  Reel advances:              polite — "Now playing: [creator] - [caption]"
  Toast notification:         polite — toast text
  Panel opens:                assertive — "Comments panel open"
  Panel closes:               polite — "Panel closed"
  Error state:                assertive — error message text
```

### 8.5 Reduced Motion

```css
@media (prefers-reduced-motion: reduce) {
  /* Disable floating hearts */
  .heart-float { display: none; }

  /* Replace smooth scroll-snap with instant */
  .feed-scroll { scroll-behavior: auto; }

  /* Disable live comment animations */
  .live-comment { animation: none; opacity: 0.7; }

  /* Collapse transitions to instant */
  * { transition-duration: 0.01ms !important;
      animation-duration: 0.01ms !important; }

  /* Keep progress bar but static (no animation on fill) */
}
```

---

## 9. Animation System

### 9.1 Motion Tokens

```css
/* Durations (semantic) */
--motion-instant:  var(--dur-instant);   /* 80ms  — haptic feedback */
--motion-micro:    var(--dur-fast);      /* 150ms — hover states, toggles */
--motion-interact: var(--dur-normal);    /* 250ms — button clicks, small UI */
--motion-enter:    var(--dur-slow);      /* 400ms — panel slide-in, page transitions */
--motion-exit:     var(--dur-normal);    /* 250ms — dismissals (always shorter than enter) */
--motion-brand:    var(--dur-slower);    /* 600ms — brand moments, like burst */
```

### 9.2 Named Keyframes

```css
/* Fade up entrance */
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

/* Floating hearts */
@keyframes floatUp {
  0%   { opacity: 1; transform: translateY(0) scale(0.6); }
  60%  { opacity: 0.8; transform: translateY(-120px) scale(1.1); }
  100% { opacity: 0; transform: translateY(-160px) scale(0.8); }
}

/* Heart burst (double-tap center) */
@keyframes heartBurst {
  0%   { opacity: 0; transform: translate(-50%,-50%) scale(0.5); }
  30%  { opacity: 1; transform: translate(-50%,-50%) scale(1.2); }
  70%  { opacity: 1; transform: translate(-50%,-50%) scale(1.0); }
  100% { opacity: 0; transform: translate(-50%,-50%) scale(1.0); }
}

/* Badge pulse */
@keyframes pulseDot {
  0%, 100% { transform: scale(1);   opacity: 1; }
  50%       { transform: scale(1.3); opacity: 0.7; }
}

/* Live comment scroll */
@keyframes liveComment {
  0%   { opacity: 0; transform: translateY(16px); }
  15%  { opacity: 1; transform: translateY(0); }
  80%  { opacity: 1; transform: translateY(0); }
  100% { opacity: 0; transform: translateY(-8px); }
}

/* Shimmer (loading state) */
@keyframes shimmer {
  0%   { background-position: -400px 0; }
  100% { background-position: 400px 0; }
}

/* Progress fill */
@keyframes progressFill {
  from { width: 0%; }
  to   { width: 100%; }
}

/* Panel slide-up */
@keyframes slideUp {
  from { transform: translateY(100%); }
  to   { transform: translateY(0); }
}

/* Toast enter */
@keyframes toastEnter {
  from { opacity: 0; transform: translateX(-50%) translateY(12px); }
  to   { opacity: 1; transform: translateX(-50%) translateY(0); }
}
```

### 9.3 Stagger Rules

When multiple elements animate in together (e.g., story rings, feature cards):

```
Base delay:   0ms
Stagger step: 60ms per item
Max stagger:  6 items (360ms total) — beyond that, no additional delay
```

### 9.4 Interaction Feedback Summary

| Trigger             | Element          | Animation                               | Duration |
|---------------------|------------------|-----------------------------------------|----------|
| Tap like            | Heart icon       | scale 1.3 → 1.0                        | 80ms     |
| Tap like            | Floating hearts  | floatUp × 8, staggered 100ms each      | 1300ms   |
| Double tap          | Center heart     | heartBurst                              | 800ms    |
| Tap save            | Bookmark icon    | scale 1.2 → 1.0                        | 150ms    |
| Tap comment         | Comment icon     | scale 1.1 → 1.0 + panel slides up     | 250ms    |
| Panel open          | Panel + overlay  | slideUp + opacity 0→0.7                | 400ms    |
| Panel close         | Panel + overlay  | slideDown + opacity 0.7→0              | 250ms    |
| Toast show          | Toast            | toastEnter                              | 150ms    |
| Toast dismiss       | Toast            | opacity 1→0 + translateY -8px          | 150ms    |
| Scroll to new reel  | Video            | crossfade / instant snap               | 0ms (snap) |
| Sidebar nav active  | Nav icon + bg    | opacity + background                   | 150ms    |
| Story ring tap      | Ring border      | scale 0.95 → 1.0                       | 80ms     |

---

## Appendix A — z-index Stack

```
z-index layer map (lowest to highest):

10    Sidebar (left nav)
20    RightPanel
30    FeedHeader (sticky)
40    NavigationArrows (fixed)
50    PanelOverlay (dim backdrop)
60    CommentsPanel / SharePanel
70    Toast
80    Tooltip (future)
```

---

## Appendix B — Icon Set

All icons use a consistent 24px viewBox from Lucide or Phosphor (outline style, 1.5px stroke).

| Function     | Icon name     | Active state             |
|--------------|---------------|--------------------------|
| Home         | `home`        | `home-fill` (bold)       |
| Search       | `search`      | —                        |
| Create       | `plus-circle` | —                        |
| Chat         | `message-circle` | + badge dot           |
| Notifications| `bell`        | `bell-fill` + badge dot  |
| Profile      | Avatar image  | Ring border              |
| Like         | `heart`       | `heart-fill` red         |
| Comment      | `message-circle` | —                     |
| Share        | `send`        | —                        |
| Save         | `bookmark`    | `bookmark-fill`          |
| More         | `more-vertical` | —                      |
| Settings     | `settings`    | —                        |
| Analytics    | `bar-chart-2` | —                        |
| Close        | `x`           | —                        |
| Play         | `play`        | —                        |
| Mute         | `volume-x`    | `volume-2` (unmuted)     |

---

*End of specification. Questions? Ping design before implementing edge cases.*
