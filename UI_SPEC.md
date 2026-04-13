# ShareChat Revamp — UI/UX Specification

**Version:** 1.0  
**Date:** 2026-04-14  
**Author:** Aisha (Design)  
**Audience:** Frontend engineers  

---

## Table of Contents

1. [Design Tokens](#1-design-tokens)
2. [Typography System](#2-typography-system)
3. [Component Hierarchy](#3-component-hierarchy)
4. [Page Layouts & Wireframes](#4-page-layouts--wireframes)
5. [Component Specifications](#5-component-specifications)
6. [Interaction & Animation Specs](#6-interaction--animation-specs)
7. [Responsive Behavior](#7-responsive-behavior)
8. [Accessibility Specification](#8-accessibility-specification)

---

## 1. Design Tokens

### 1.1 Primitive Color Tokens

```
// Neutrals
black-0:    #0a0a0a
black-100:  #111111
black-200:  #1a1a1a
black-300:  rgba(255,255,255,0.06)   // glass fill
black-400:  rgba(255,255,255,0.08)   // glass border
black-500:  rgba(255,255,255,0.15)   // subtle border

// Text
white-100:  #ffffff
white-60:   rgba(255,255,255,0.60)
white-35:   rgba(255,255,255,0.35)
white-15:   rgba(255,255,255,0.15)

// Brand
orange-500: #FF6B35
orange-glow: rgba(255,107,53,0.30)
pink-500:   #FF3CAC
purple-500: #784BA0
blue-500:   #2B86C5
teal-400:   #00F5A0
cyan-400:   #00D9F5
gold-400:   #F7971E
gold-500:   #FFD200

// Status
green-500:  #22c55e
```

### 1.2 Semantic Color Tokens

| Token | Primitive | Usage |
|---|---|---|
| `color-bg-primary` | `black-0` | Page / app background |
| `color-bg-secondary` | `black-100` | Nav, panel backgrounds |
| `color-bg-card` | `black-200` | Cards, feed items |
| `color-bg-glass` | `black-300` | Glass-morphism fills |
| `color-border-glass` | `black-400` | Glass-morphism borders |
| `color-border-subtle` | `black-500` | Dividers |
| `color-text-primary` | `white-100` | Headlines, labels |
| `color-text-secondary` | `white-60` | Subtext, descriptions |
| `color-text-tertiary` | `white-35` | Timestamps, metadata |
| `color-text-disabled` | `white-15` | Disabled states |
| `color-accent` | `orange-500` | CTAs, active states, links |
| `color-accent-glow` | `orange-glow` | Glow shadows on accent |
| `color-success` | `green-500` | Confirmation, checkmarks |

### 1.3 Gradient Tokens

```
gradient-brand:  linear-gradient(135deg, #FF6B35, #FF3CAC)
gradient-aurora: linear-gradient(135deg, #FF3CAC, #784BA0, #2B86C5)
gradient-teal:   linear-gradient(135deg, #00F5A0, #00D9F5)
gradient-gold:   linear-gradient(135deg, #F7971E, #FFD200)
```

### 1.4 Spacing Tokens (4px grid)

| Token | Value | Common use |
|---|---|---|
| `space-1` | 4px | Icon gap, tight metadata |
| `space-2` | 8px | Button padding-y, chip gap |
| `space-3` | 12px | Story gap, inline gap |
| `space-4` | 16px | Base padding, post actions gap |
| `space-5` | 20px | Reel sidebar gap |
| `space-6` | 24px | Section padding-x, card gap |
| `space-8` | 32px | Nav logo gap, CTA card padding |
| `space-10` | 40px | Sidebar section gap |
| `space-12` | 48px | Hero CTA margin |
| `space-16` | 64px | Feature grid gap |
| `space-20` | 80px | Hero stats margin-top |
| `space-24` | 96px | — |
| `space-30` | 120px | Section vertical padding |

### 1.5 Border Radius Tokens

| Token | Value | Usage |
|---|---|---|
| `radius-xs` | 4px | Discover grid items |
| `radius-sm` | 12px | Feature icons, tags |
| `radius-md` | 16px | Post media, discover search |
| `radius-lg` | 24px | Cards, feature cards, panels |
| `radius-xl` | 32px | CTA card |
| `radius-full` | 9999px | Pills, avatars, buttons |

### 1.6 Shadow Tokens

```
shadow-card:    0 40px 80px rgba(0,0,0,0.50)
shadow-phone:   0 60px 120px rgba(0,0,0,0.60), 0 0 160px rgba(255,107,53,0.08)
shadow-btn:     0 8px 32px rgba(255,107,53,0.25)
shadow-btn-hover: 0 12px 40px rgba(255,107,53,0.35)
shadow-glow:    0 0 8px rgba(255,107,53,0.30)
```

### 1.7 Animation Duration Tokens

| Token | Value | Usage |
|---|---|---|
| `duration-fast` | 150ms | Hover color transitions |
| `duration-normal` | 300ms | Transform, opacity, panels |
| `duration-slow` | 500ms | Phone hover/rotate |
| `duration-story` | 8000ms | Story / reel progress bar |
| `duration-live` | 4500ms | Live comment lifetime |

### 1.8 Easing Tokens

```
ease-default: cubic-bezier(0.25, 0.46, 0.45, 0.94)   // smooth decelerate
ease-spring:  cubic-bezier(0.34, 1.56, 0.64, 1.00)    // slight overshoot (hearts, likes)
ease-linear:  linear                                    // progress bars, spins
```

---

## 2. Typography System

### 2.1 Font Stacks

```
font-display: 'Space Grotesk', -apple-system, sans-serif   // headings, logo, numbers
font-body:    'Inter', -apple-system, sans-serif            // all body text
```

Loading strategy: Google Fonts `display=swap` — use FOUT, no layout shift.

### 2.2 Scale

| Token | Size | Line-height | Letter-spacing | Weight | Usage |
|---|---|---|---|---|---|
| `text-hero` | clamp(48px, 7vw, 88px) | 1.05 | -0.03em | 700 | Hero h1 |
| `text-section` | clamp(32px, 4vw, 52px) | 1.1 | -0.02em | 700 | Section titles |
| `text-heading-xl` | 56px | 1.0 | — | 700 | Impact metrics |
| `text-heading-lg` | 36px | 1.1 | — | 700 | Stat values |
| `text-heading-md` | 24px | 1.3 | — | 600 | CTA card h2, clamp(28px,4vw,44px) |
| `text-heading-sm` | 20px | 1.3 | — | 600 | Feature titles, feed logo |
| `text-body-lg` | 18px | 1.6 | — | 400 | Hero paragraph |
| `text-body-md` | 16px | 1.6 | — | 400 | Section desc, CTA copy |
| `text-body-sm` | 14px | 1.7 | — | 400 | Feature desc, nav links |
| `text-ui-md` | 15px | — | — | 600 | Buttons |
| `text-ui-sm` | 13px | — | — | 500–600 | Post name, reel creator |
| `text-ui-xs` | 12px | — | — | 500–600 | Chips, post follow, tags |
| `text-meta` | 11px | 1.5 | — | 400–500 | Timestamps, reel desc, counts |
| `text-label` | 12px | — | 0.15em | 600 | Section tags (uppercase) |
| `text-caption` | 10px | — | 0.1em | 400–600 | Story names, discover views |

---

## 3. Component Hierarchy

```
App
├── LandingPage (index.html)
│   ├── Nav
│   │   ├── Logo
│   │   ├── NavLinks
│   │   └── NavCTA (button)
│   ├── HeroSection
│   │   ├── HeroBadge
│   │   ├── HeroHeading (gradient text)
│   │   ├── HeroSubtext
│   │   ├── HeroCTA (primary + secondary buttons)
│   │   └── HeroStats (stat items ×3)
│   ├── PhoneSection
│   │   └── PhoneShowcase
│   │       ├── PhoneFrame (Feed)
│   │       │   └── FeedScreen (mini)
│   │       │       ├── FeedHeader
│   │       │       ├── StoriesBar
│   │       │       └── FeedPost ×2
│   │       ├── PhoneFrame (Reels) [center, featured]
│   │       │   └── ReelsScreen (mini)
│   │       │       ├── ReelItem
│   │       │       ├── ReelContent (creator + desc + music)
│   │       │       └── ReelSidebar (action buttons)
│   │       └── PhoneFrame (Discover)
│   │           └── DiscoverScreen (mini)
│   │               ├── DiscoverSearch
│   │               ├── DiscoverChips
│   │               └── DiscoverGrid
│   ├── FeaturesSection
│   │   └── FeatureCard ×N
│   ├── ComparisonSection
│   │   └── ComparisonTable
│   │       └── ComparisonRow ×N
│   ├── BeforeAfterSection
│   │   ├── BACard (before)
│   │   └── BACard (after)
│   ├── ImpactSection
│   │   └── ImpactMetric ×N
│   └── CTASection
│       └── CTACard
│
└── FeedApp (feed.html)
    ├── Sidebar (left)
    │   ├── SidebarLogo
    │   ├── SidebarNav (icon links)
    │   └── SidebarProfile
    ├── FeedCenter
    │   └── FeedScroll
    │       └── VideoCard ×N (snap-scrolled)
    │           ├── VideoElement
    │           ├── VideoGradient
    │           ├── LiveComments overlay
    │           ├── VideoInfo (creator + caption + music)
    │           └── VideoActions (like, comment, share, save, more)
    ├── RightSidebar
    │   ├── UserProfile card
    │   ├── SuggestionsList
    │   └── TrendsList
    ├── CommentsPanel (overlay)
    │   ├── PanelHeader
    │   ├── CommentList
    │   │   └── CommentItem ×N
    │   └── CommentInput
    ├── SharePanel (overlay)
    │   ├── PanelHeader
    │   ├── ShareUrl (copy link)
    │   └── ShareOptions (WhatsApp, etc.)
    ├── NavArrows (keyboard hint)
    └── Toast
```

---

## 4. Page Layouts & Wireframes

### 4.1 Landing Page — Desktop (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│ NAV  [Logo]        [Features] [Feed] [Impact] [Download]  [CTA btn] │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│                    [Badge: Live Redesign ●]                         │
│                                                                     │
│              ShareChat,                                             │
│              Reimagined.  (gradient on "Reimagined")                │
│                                                                     │
│       Your community deserves a beautiful home.                     │
│                                                                     │
│            [▶ Open Live Feed]    [↓ See Mockups]                   │
│                                                                     │
│       200M Users     50M Videos     98% Satisfaction               │
│                                                                     │
├─────────────────────────────────────────────────────────────────────┤
│                  PHONE SHOWCASE (3 phones, perspective)             │
│                                                                     │
│   [Phone Feed]    [Phone Reels ★ featured]    [Phone Discover]     │
│                                                                     │
├─────────────────────────────────────────────────────────────────────┤
│  FEATURES (3-col grid)                                              │
│  ┌───────────┐  ┌───────────┐  ┌───────────┐                       │
│  │ 🔥 Reels  │  │ ✨ Discover│  │ 🤝 Connect│                       │
│  └───────────┘  └───────────┘  └───────────┘                       │
├─────────────────────────────────────────────────────────────────────┤
│  BEFORE / AFTER  (2-col)                                            │
│  ┌──────────────────┐  ┌──────────────────┐                        │
│  │  OLD (faded)     │  │  NEW (vivid)     │                        │
│  └──────────────────┘  └──────────────────┘                        │
├─────────────────────────────────────────────────────────────────────┤
│  IMPACT METRICS  (centered row)                                     │
│     2× faster    85% DAU↑    4.8★ rating    3× shares             │
├─────────────────────────────────────────────────────────────────────┤
│                    CTA CARD (gradient bg)                           │
│              Ready to experience ShareChat?                         │
│                    [Open Live Feed →]                               │
└─────────────────────────────────────────────────────────────────────┘
```

### 4.2 Feed App — Desktop (≥1024px)

```
┌──────────┬──────────────────────────┬──────────────────────────────┐
│ SIDEBAR  │       FEED CENTER        │       RIGHT PANEL            │
│ 240px    │       ~480px (auto)      │       320px                  │
│          │                          │                              │
│ [Logo]   │  ┌────────────────────┐  │  ┌──────────────────────┐   │
│          │  │                    │  │  │ [Avatar] You          │   │
│ [Home]   │  │    Video #1        │  │  │ username · Follow     │   │
│ [Search] │  │    (full-height)   │  │  └──────────────────────┘   │
│ [Create] │  │                    │  │                              │
│ [Reels]  │  │  [Live Comments]   │  │  Suggestions for you         │
│ [Profile]│  │                    │  │  ┌──────────────────────┐   │
│          │  │  [Creator info]    │  │  │ [Av] Name   [Follow] │   │
│ [Avatar] │  │  [Caption]         │  │  │ [Av] Name   [Follow] │   │
│          │  │  [♫ music]         │  │  │ [Av] Name   [Follow] │   │
│          │  │                    │  │  └──────────────────────┘   │
│          │  │ ❤️ 💬 ➤ 🔖 ···      │  │                              │
│          │  └────────────────────┘  │  Trending                    │
│ [↑] [↓]  │      ← snap scroll →    │  • #DanceChallenge           │
│          │                          │  • IPL 2026                  │
└──────────┴──────────────────────────┴──────────────────────────────┘
```

### 4.3 Feed App — Mobile (<768px)

```
┌─────────────────────────────────┐
│     [← 🔴 LIVE]   ShareChat  [≡]│  ← sticky top bar (no left sidebar)
├─────────────────────────────────┤
│                                 │
│         Video (full-bleed)      │
│                                 │
│   Live comment flows up ↑       │
│                                 │
│  [Creator] [Follow]         ❤️  │
│  Caption text                💬 │
│  ♫ Trending Sound            ➤ │
│                              🔖 │
│                              ···│
├─────────────────────────────────┤
│  [●] [●] [●] [●] [●]  ← bottom nav
└─────────────────────────────────┘
```

### 4.4 Comments Panel (slide-up on mobile, slide-in on desktop)

```
Desktop (right-of-center overlay):
┌─────────────────────────────────┐
│ Comments (42)              [✕]  │
├─────────────────────────────────┤
│ [Av] Creator Name               │
│       Pinned comment text       │
│       2h ago · ❤ 1.2k           │
├─────────────────────────────────┤
│ [Av] User123                    │
│       This is amazing!          │
│       5m ago · ❤ 24             │
│  ...                            │
├─────────────────────────────────┤
│ [😎]  Add a comment...  [Post]  │
└─────────────────────────────────┘
```

### 4.5 Share Panel

```
┌─────────────────────────────────┐
│ Share                      [✕]  │
├─────────────────────────────────┤
│ https://sharechat.com/v/xxx     │
│                       [Copy ✓]  │
├─────────────────────────────────┤
│ [WhatsApp] [Instagram] [Twitter]│
│ [Telegram] [Copy Link]          │
└─────────────────────────────────┘
```

---

## 5. Component Specifications

### 5.1 Nav (Landing Page)

**States:** default (transparent) → scrolled (frosted glass)

| Property | Value |
|---|---|
| Position | `fixed` top 0, full width, z-index 100 |
| Padding | 16px 32px (`space-4` `space-8`) |
| Background (default) | transparent |
| Background (scrolled) | `rgba(10,10,10,0.80)` + `backdrop-filter: blur(20px)` |
| Border-bottom (scrolled) | `1px solid color-border-glass` |
| Transition | background `duration-normal ease-default` |

**Logo:**
- Font: `font-display`, 22px, weight 700
- Icon box: 32×32px, `radius-sm`, `gradient-brand`
- Letter inside: "S", 16px, weight 800, white

**Nav Links:**
- Gap: `space-8` (32px)
- Font: `text-body-sm`, `color-text-secondary`
- Hover: `color-text-primary`, `duration-fast`
- Active: `color-text-primary`, weight 500

**Accessibility:**
- `<nav role="navigation" aria-label="Main">`
- Active link: `aria-current="page"`
- Focus: 2px solid `color-accent` outline, offset 2px

---

### 5.2 Button System

#### Primary Button
```
padding:          16px 36px
border-radius:    radius-full
background:       gradient-brand
color:            white
font:             text-ui-md (15px/600)
shadow:           shadow-btn
transition:       transform + shadow duration-normal ease-default

hover:   translateY(-2px), shadow-btn-hover
active:  translateY(0px), scale(0.98)
focus:   2px solid white outline, offset 2px
disabled: opacity 0.4, cursor not-allowed
```

#### Secondary Button
```
padding:          16px 36px
border-radius:    radius-full
border:           1px solid color-border-glass
background:       color-bg-glass
backdrop-filter:  blur(20px)
color:            white
font:             text-ui-md (15px/500)

hover:   background rgba(255,255,255,0.10), translateY(-2px)
active:  translateY(0px)
focus:   2px solid color-accent outline, offset 2px
```

#### White Button (CTA on gradient bg)
```
padding:          16px 40px
border-radius:    radius-full
background:       white
color:            black-0 (#0a0a0a)
font:             text-ui-md (15px/600)

hover:   translateY(-2px), box-shadow 0 12px 40px rgba(0,0,0,0.30)
```

#### Follow Pill (inline on posts / reels)
```
padding:          4px 14px
border-radius:    radius-full
border:           1px solid color-accent
color:            color-accent
font:             text-ui-xs (12px/600)
background:       transparent

active (Following): color color-text-secondary, border-color color-border-subtle
transition:       color + border-color duration-fast
```

---

### 5.3 StoryRing

**Five states:**

| State | Visual |
|---|---|
| Empty / Add | Dashed border, `color-border-subtle`, plus icon |
| Unwatched | `gradient-brand` ring, 2px padding |
| Watched | `color-border-subtle` ring (grey) |
| Live | `gradient-brand` ring + pulsing dot |
| Loading | Shimmer animation on ring and avatar |

```
Outer ring:   56×56px, radius-full
Padding:      2px (creates border effect)
Inner avatar: 52×52px, radius-full, border 2px solid color-bg-primary (gap trick)
Label:        text-caption (10px), color-text-tertiary, max-width 56px, ellipsis
```

---

### 5.4 FeedPost (landing page mini)

```
Container:    padding 0 16px 20px, border-bottom 1px solid rgba(255,255,255,0.04)
Header:       flex row, gap space-3 (12px), margin-bottom space-3
  Avatar:     36×36px, radius-full, gradient-aurora
  Name:       text-ui-sm (13px/600), color-text-primary
  Time:       text-caption (11px), color-text-tertiary
  Follow pill: right-aligned
Media:        aspect-ratio 4/5, radius-md, overflow hidden
  Gradient overlay: bottom 50%, linear transparent→rgba(0,0,0,0.7)
  Play button: 48×48px circle, rgba(255,255,255,0.20) + blur(10px)
  Tag:        bottom-left, pill, 11px/500, rgba(255,255,255,0.15) + blur
Actions:      flex row, gap 20px, margin-top space-3
  Like / Comment / Share / Save:
    Icon:     20×20px SVG
    Count:    text-meta (12px), color-text-secondary
```

---

### 5.5 VideoCard (feed.html)

**Five states:**

| State | Visual |
|---|---|
| Empty | Dark bg, animated gradient shimmer |
| Loading | Spinner centered, semi-transparent overlay |
| Error | Error icon + "Video unavailable" + retry button |
| Paused | Play icon overlay visible, progress frozen |
| Playing | No overlay, live comments flowing, progress bar animating |

```
Container:    full viewport height, scroll-snap-align start
Video:        width 100%, height 100%, object-fit cover
              autoplay, muted, loop, playsInline
Gradient:     absolute inset 0, transparent 40% → rgba(0,0,0,0.80)
Live Comments:
  Position:  absolute, above VideoInfo
  Each item: slide-in from left, fade-out after 4500ms
  Max shown: 3 at once, FIFO removal
VideoInfo:   absolute bottom 80px, left 16px, right 60px
  Creator:   avatar 32×32 + name text-ui-sm/600 + Follow pill
  Caption:   text-meta (12px/1.5), color rgba(255,255,255,0.85)
  Music:     ♫ icon (spinning) + text-caption, color-text-secondary
VideoActions: absolute right 12px, bottom 100px
  Buttons:   vertical stack, gap 20px
  Each btn:  icon-circle 36×36px rgba(255,255,255,0.10) + blur + count label
  Like btn:  on tap → hearts float animation (8 hearts, staggered)
ProgressBar: absolute bottom 0, height 3px, rgba(255,255,255,0.15)
  Fill:      white, animates left-to-right over video duration
```

**Keyboard interactions:** ↑↓ arrows navigate videos, Space toggles play/pause.

---

### 5.6 Heart Float Animation

Triggered on like tap. 8 heart elements created, appended to VideoCard, then removed after animation.

```
Timing:    each heart 0.6s + staggered delay (0 to 0.4s random offset)
Path:      translateY(-120px to -200px random) + translateX(±30px random)
Opacity:   1 → 0 (fade during last 30% of animation)
Scale:     1.0 → 0.6 → 1.2 → 0 (spring feel)
Easing:    ease-spring for scale, ease-linear for translateY
Remove:    after animationend to keep DOM clean
```

---

### 5.7 CommentsPanel

**States:** closed (off-screen) → open (slide in)

```
Overlay:      fixed inset 0, rgba(0,0,0,0.50), z-index 50
Panel box:    fixed bottom 0, width 100% (mobile) / 420px right-anchored (desktop)
              radius-lg top corners only (mobile), radius-lg all (desktop)
              background color-bg-card, border 1px color-border-glass
              max-height 70vh (mobile), 80vh (desktop), overflow-y auto

Open animation:   translateY(100%) → translateY(0), duration-normal ease-default (mobile)
                  translateX(100%) → translateX(0), desktop

Header:       padding 16px 24px, border-bottom 1px color-border-glass
  Title:      text-heading-sm (20px/600)
  Count:      color-text-secondary
  Close (✕): 24×24px touch target, color-text-secondary

CommentItem:
  Avatar:     36×36px, radius-full
  Name:       text-ui-sm (13px/600), color-accent for self
  Text:       text-body-sm (14px/1.5)
  Meta:       text-meta (11px), color-text-tertiary, "time · ❤ count"

Input bar:    sticky bottom, padding 12px 16px, background color-bg-secondary
  Input:      flex-1, background color-bg-glass, radius-full, padding 10px 16px
  Post btn:   color-accent, font-weight 600
```

**Accessibility:**
- `role="dialog"`, `aria-modal="true"`, `aria-label="Comments"`
- Focus trap: Tab cycles within panel
- `aria-live="polite"` on comment list for new additions
- Escape key closes panel, returns focus to trigger button

---

### 5.8 Toast Notification

```
Position:     fixed bottom 24px, left 50%, transform translateX(-50%)
              z-index 200
Padding:      12px 24px
Border-radius: radius-full
Background:   color-bg-secondary, border 1px color-border-glass
Font:         text-ui-xs (12px/500), color-text-primary
Backdrop:     blur(20px)

Show animation:   opacity 0 + translateY(8px) → opacity 1 + translateY(0)
                  duration: duration-normal
Auto-dismiss: 2000ms after show
Hide animation:   opacity 1 → 0, duration-fast

aria-live="assertive" role="status"
```

---

### 5.9 DiscoverSearch Bar

**States:** idle → focused → has-query → loading → results

```
Container:    margin 0 16px 20px
              padding 12px 16px
              radius-md
              background color-bg-glass, border 1px color-border-glass
              display flex, align-items center, gap space-3

Idle:         🔍 icon (color-text-tertiary) + placeholder text-body-sm (color-text-tertiary)
Focus:        border-color color-accent, outline none
              transition: border-color duration-fast
Loading:      spinner replaces search icon, pulses
Has-query:    ✕ clear button appears right side

aria-label="Search content"
role="searchbox"
```

---

### 5.10 FilterChip

```
Padding:          8px 16px
Border-radius:    radius-full
Font:             text-ui-xs (12px/500)
Background:       color-bg-glass
Border:           1px solid color-border-glass
Color:            color-text-secondary

Active state:
  Background:     gradient-brand
  Border:         transparent
  Color:          white

Hover (inactive):
  Border-color:   color-border-subtle
  Color:          color-text-primary
  duration:       duration-fast

role="tab" within role="tablist"
aria-selected="true/false"
```

---

### 5.11 SidebarNav (feed.html)

```
Width:        240px (desktop) / hidden (mobile, replaced by bottom nav)
Padding:      24px 16px
Gap between items: space-2 (8px)

NavItem:
  Padding:    12px 16px
  Radius:     radius-md
  Gap:        space-3 (12px)
  Font:       text-body-sm (14px/500)
  Color:      color-text-secondary
  Icon:       24×24px

  Active:
    Background:   color-bg-glass
    Color:        color-text-primary
    Icon:         filled variant

  Hover:
    Background:   rgba(255,255,255,0.04)
    Color:        color-text-primary
    duration:     duration-fast

  Focus:    2px solid color-accent inset outline

Keyboard: Tab cycles items, Enter activates
aria-label="Feed navigation"
role="navigation"
aria-current="page" on active item
```

---

### 5.12 RightSidebar — Suggestion Item

```
Container:    display flex, align-items center, gap space-3, padding 8px 0
Avatar:       36×36px, radius-full, background color-bg-card, font-size 18px (emoji)
Info:         flex-1
  Name:       text-ui-sm (13px/600), color-text-primary
  Meta:       text-caption (11px), color-text-tertiary
Follow toggle:
  Idle:       "Follow", color-accent, text-ui-xs (12px/600)
  Active:     "Following", color-text-secondary
  Transition: color duration-fast
```

---

### 5.13 PhoneFrame (landing page showcase)

```
Dimensions:   320×640px
Border-radius: 40px
Background:   black-0
Border:       2px solid rgba(255,255,255,0.10)
Box-shadow:   shadow-phone

Default:      transform rotateY(±5deg)
Hover:        transform rotateY(0deg) scale(1.02), duration-slow ease-default
Center phone: transform rotateY(0deg) scale(1.05), elevated z-index
Center hover: scale(1.07)

Notch:
  Position:   absolute top 0, centered
  Size:       120×28px
  Radius:     0 0 20px 20px (bottom corners)
  Background: black-0
```

---

## 6. Interaction & Animation Specs

### 6.1 Scroll Behavior

| Scroll container | Behavior |
|---|---|
| Landing page | `scroll-behavior: smooth` |
| Feed videos | `scroll-snap-type: y mandatory`, `scroll-snap-align: start` |
| Stories bar | horizontal, `scrollbar-width: none` |
| Filter chips | horizontal, `scrollbar-width: none` |
| Reels screen | `scroll-snap-type: y mandatory`, no scrollbar |

### 6.2 Video Auto-play on Scroll

```
Trigger:       IntersectionObserver threshold 0.8
Enter viewport: play(), start live comments, start progress bar
Exit viewport:  pause(), stop live comments, reset progress bar
Current index: tracked in `cur` variable, used for keyboard nav
```

### 6.3 Like Button Micro-animation

```
1. User taps ❤️
2. Icon scale: 1.0 → 1.4 (ease-spring, 150ms) → 1.0 (ease-default, 150ms)
3. Count increments with fade-up: count slides up 4px + fades in
4. 8 heart particles spawn (see § 5.6)
5. Icon color: color-text-secondary → #FF3CAC (pink), transition 100ms
```

### 6.4 Navigation Arrow (keyboard hint)

```
[↑] [↓] buttons — desktop only
Position:     fixed bottom-right, z-index 50
Size:         40×40px each, radius-full
Background:   color-bg-glass, border color-border-glass
Hover:        background color-bg-card, duration-fast
Click / keyboard ↑↓: scroll to prev/next video
```

### 6.5 Page-Level Entrance Animations

```
fadeUp keyframe:   opacity 0 + translateY(20px) → opacity 1 + translateY(0)
Applied with stagger delays on hero elements:
  HeroBadge:    0.0s delay, 0.8s duration
  HeroHeading:  0.1s delay, 0.8s duration
  HeroSubtext:  0.2s delay, 0.8s duration
  HeroCTA:      0.3s delay, 0.8s duration
  HeroStats:    0.5s delay, 0.8s duration

Reduced-motion override (@prefers-reduced-motion: reduce):
  All transition durations: max 100ms
  All translateY animations: remove translate, keep opacity only
  Heart float animation: disabled, only count increments
```

### 6.6 Live Comment Stream

```
New comment every:  2500ms + random 0–2000ms
Slide-in:           translateX(-20px) + opacity 0 → translateX(0) + opacity 1
                    duration-fast (150ms)
Lifetime:           4500ms
Fade-out:           opacity 1 → 0, duration-normal (300ms)
Max visible:        3; oldest removed immediately when 4th added
```

### 6.7 Progress Bar (Reel / Story)

```
Height:     3px
Background: rgba(255,255,255,0.15)
Fill:       white, left-to-right
Duration:   matched to video length (default 8s fallback)
Easing:     linear
Resets:     on video end / manual nav
```

---

## 7. Responsive Behavior

### 7.1 Breakpoints

| Token | Value | Layout change |
|---|---|---|
| `bp-sm` | 640px | Single column, stacked hero CTAs |
| `bp-md` | 768px | Two columns, feed sidebar appears |
| `bp-lg` | 1024px | Full desktop layout, right panel shows |
| `bp-xl` | 1280px | Max content width centered |

### 7.2 Landing Page (index.html)

| Element | Mobile (<640px) | Tablet (640–1023px) | Desktop (≥1024px) |
|---|---|---|---|
| Hero h1 | clamp(48px min) | ~64px | 88px max |
| Hero CTAs | stacked column | row | row |
| Hero Stats | 2-col grid | row (3 items) | row (3 items), gap 48px |
| Phone Showcase | single phone | 2 phones | 3 phones, perspective |
| Features grid | 1 col | 2 col | 3 col auto-fit 300px |
| Comparison table | scroll-x | full width | full width |
| Before/After | 1 col | 2 col | 2 col |
| Impact metrics | 2-col grid | row wrap | row, gap 64px |
| Nav links | hidden (hamburger) | hidden | visible |

### 7.3 Feed App (feed.html)

| Element | Mobile (<768px) | Desktop (≥768px) |
|---|---|---|
| Left sidebar | hidden | visible, 240px fixed |
| Right panel | hidden | visible, 320px fixed |
| Feed center | full width | auto (fills between sidebars) |
| Bottom nav | visible (5 icons) | hidden |
| Comments panel | slide-up sheet, 70vh | side panel, right |
| Share panel | slide-up sheet | centered modal |
| Nav arrows | hidden | visible, fixed bottom-right |
| Video actions | right edge overlay | right edge overlay |

**Touch targets (mobile):** All interactive elements minimum 44×44px.

### 7.4 Feed Video Sizing

```
Mobile:   width 100vw, height 100dvh (dynamic viewport height for mobile browser chrome)
Desktop:  width ~480px centered, height 100vh
          max-width enforced by feed-center container
```

---

## 8. Accessibility Specification

### 8.1 Color Contrast Compliance

| Foreground | Background | Ratio | Level |
|---|---|---|---|
| `color-text-primary` (#fff) | `color-bg-primary` (#0a0a0a) | 19.5:1 | AAA |
| `color-text-secondary` (rgba white 60%) | `color-bg-primary` | 5.8:1 | AA |
| `color-text-tertiary` (rgba white 35%) | `color-bg-primary` | 3.2:1 | AA Large only |
| `color-accent` (#FF6B35) | `color-bg-primary` | 4.6:1 | AA |
| white | `gradient-brand` (#FF6B35) | 3.5:1 | ⚠ AA Large only — use weight ≥600 |
| black | white button | 19.5:1 | AAA |

**Non-color indicators required:**
- Like state: icon color change AND count change (not color alone)
- Active tab: underline or background fill AND color change
- Error state: icon + border + text message (not border color alone)
- Follow/Following: text label change, not just color

### 8.2 Keyboard Navigation Map

| Context | Key | Action |
|---|---|---|
| Feed | `↑` / `↓` | Previous / next video |
| Feed | `Space` | Play / pause current video |
| Feed | `L` | Like current video |
| Feed | `C` | Open comments |
| Feed | `S` | Open share |
| Comments panel | `Escape` | Close panel, return focus to trigger |
| Comments panel | `Tab` | Cycle within panel (trapped) |
| Comments panel | `Enter` | Submit comment (in input) |
| Share panel | `Escape` | Close panel |
| Nav | `Tab` | Cycle nav links |
| Chips / tabs | `→` / `←` | Navigate chip options |
| Chips / tabs | `Enter` / `Space` | Activate chip |

### 8.3 ARIA Roles and Properties

```html
<!-- Feed app shell -->
<main role="main" aria-label="Video feed">
<nav role="navigation" aria-label="App navigation">
<aside role="complementary" aria-label="Suggestions">

<!-- Each video -->
<article aria-label="Video by {creator}: {caption}" aria-posinset="{n}" aria-setsize="-1">
  <video aria-label="Video content" title="{caption}">

<!-- Action buttons -->
<button aria-label="Like video. {count} likes" aria-pressed="{true|false}">
<button aria-label="Open comments. {count} comments">
<button aria-label="Share video">
<button aria-label="Save video" aria-pressed="{true|false}">

<!-- Comments panel -->
<div role="dialog" aria-modal="true" aria-label="Comments" aria-describedby="cmtC">
<ul role="list" aria-live="polite" aria-label="Comment list">

<!-- Share panel -->
<div role="dialog" aria-modal="true" aria-label="Share this video">

<!-- Toast -->
<div role="status" aria-live="assertive">

<!-- Story rings -->
<button aria-label="{name}'s story" aria-pressed="{watched}">

<!-- Filter chips -->
<div role="tablist" aria-label="Content filters">
<button role="tab" aria-selected="{true|false}">
```

### 8.4 Focus Management

```
Focus indicators:
  All interactive:  outline: 2px solid color-accent, outline-offset: 2px
  Exception on dark bg overlays: outline: 2px solid white
  Never: outline: none without a visible replacement

Focus traps:
  CommentsPanel:  trap when open, release on close
  SharePanel:     trap when open, release on close

Focus restoration:
  On panel close: return focus to the button that opened it

Skip link:
  <a href="#feed" class="skip-link">Skip to feed</a>
  Visible on focus only (translate into view)
```

### 8.5 Screen Reader Announcements

| Event | Announcement | Region |
|---|---|---|
| New video scrolled into view | "{creator}: {caption}" | `aria-live="polite"` |
| Video liked | "Video liked. {count} likes." | `aria-live="polite"` |
| Comment posted | "Comment posted." | `aria-live="assertive"` |
| Link copied | "Link copied to clipboard." | `aria-live="assertive"` |
| Follow toggled | "Now following {name}" / "Unfollowed {name}" | `aria-live="polite"` |
| Panel opened | Focus moves to panel; `aria-label` announces panel name | — |

### 8.6 Reduced Motion

```css
@media (prefers-reduced-motion: reduce) {
  /* Hero entrance: remove translate, keep opacity fade */
  .hero-badge, .hero h1, .hero p, .hero-cta, .hero-stats {
    animation: fadeIn 0.4s ease both;  /* opacity only */
  }

  /* Phone hover: no 3D transform */
  .phone-frame, .phone-frame:hover {
    transform: none;
    transition: none;
  }

  /* Heart float: disabled */
  .heart-particle { display: none; }

  /* All transition durations clamped */
  * { transition-duration: 100ms !important; }

  /* Progress bar: static full width (no animation) */
  .reel-progress-bar { animation: none; width: 100%; }

  /* Live comments: no slide-in, just appear/disappear */
  .live-comment { transition: opacity 0.1s; }
}
```

---

## Appendix A — Design Decisions & Rationale

| Decision | Rationale |
|---|---|
| Dark-first theme | Reduces eye strain in video consumption context; enables vibrant gradient accents without washing out |
| Glass-morphism fills | Creates depth hierarchy without heavy shadows; suits dark video backgrounds |
| `gradient-brand` on CTAs only | Accent gradient used sparingly to preserve attention economy |
| `Space Grotesk` for display, `Inter` for body | Space Grotesk gives a modern, distinct brand identity; Inter is maximum legibility at small sizes |
| Vertical snap-scroll feed | Industry-standard short-video pattern (TikTok/Reels); eliminates partial-video state problem |
| 4px grid | Ensures pixel-perfect spacing at all DPR scales |
| Emoji avatars as placeholder | Avoids broken image states in demo; maps to "empty" avatar state |

---

## Appendix B — Open Issues for Engineering Review

| # | Issue | Owner |
|---|---|---|
| B1 | `color-text-tertiary` (white 35%) fails AA at small body text — avoid for informational content, metadata only | Design → confirm with a11y audit |
| B2 | Live comment interval uses `setInterval` without cleanup on unmount — memory leak in SPA context | Frontend |
| B3 | `scroll-snap` on feed — iOS Safari requires `-webkit-overflow-scrolling: touch` fallback | Frontend |
| B4 | No loading skeleton shown between video loads — add shimmer skeleton state for VideoCard | Design |
| B5 | `backdrop-filter: blur()` not supported in Firefox <103 — add fallback solid background | Frontend |
| B6 | Skip-link not yet implemented in feed.html — required for keyboard-only users | Frontend |
| B7 | Comment input lacks character limit UI — define max chars (280?) and show count near limit | Design |
