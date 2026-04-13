# ShareChat Revamp — UI Specification
**Version:** 1.0  
**Date:** 2026-04-13  
**Author:** Aisha (Senior Design Engineer)  
**Audience:** Frontend engineers implementing the ShareChat revamp

---

## Table of Contents

1. [Design Tokens](#1-design-tokens)
2. [Typography Scale](#2-typography-scale)
3. [Spacing System](#3-spacing-system)
4. [Component Hierarchy](#4-component-hierarchy)
5. [Component Specs](#5-component-specs)
   - [5.1 Sidebar Navigation](#51-sidebar-navigation)
   - [5.2 Reel Card](#52-reel-card)
   - [5.3 Action Button Strip](#53-action-button-strip)
   - [5.4 Comments Modal](#54-comments-modal)
   - [5.5 Share Modal](#55-share-modal)
   - [5.6 Floating Hearts Animation](#56-floating-hearts-animation)
   - [5.7 Toast Notification](#57-toast-notification)
   - [5.8 Stories Bar](#58-stories-bar)
   - [5.9 Right Sidebar (Suggestions)](#59-right-sidebar-suggestions)
   - [5.10 Video Progress Bar](#510-video-progress-bar)
6. [Layout Wireframes](#6-layout-wireframes)
   - [6.1 Feed View — Desktop](#61-feed-view--desktop)
   - [6.2 Feed View — Mobile](#62-feed-view--mobile)
   - [6.3 Marketing Page — Desktop](#63-marketing-page--desktop)
7. [Responsive Behavior](#7-responsive-behavior)
8. [Interaction Specifications](#8-interaction-specifications)
9. [Animation Catalogue](#9-animation-catalogue)
10. [Accessibility Specification](#10-accessibility-specification)

---

## 1. Design Tokens

All values are defined as CSS custom properties on `:root`. Engineers must reference tokens, never hardcode raw values.

### 1.1 Color — Primitive Tokens

| Token | Value | Swatch |
|---|---|---|
| `--color-black` | `#0a0a0a` | Pure black base |
| `--color-dark-1` | `#111111` | Very dark grey |
| `--color-dark-2` | `#1a1a1a` | Card surface |
| `--color-dark-3` | `#222222` | Elevated surface |
| `--color-white` | `#ffffff` | Pure white |
| `--color-orange-500` | `#FF6B35` | Primary accent |
| `--color-pink-500` | `#FF3CAC` | Gradient terminus |
| `--color-purple-500` | `#784BA0` | Gradient midpoint |
| `--color-blue-500` | `#2B86C5` | Gradient terminus |
| `--color-mint-500` | `#00F5A0` | Success/live accent |
| `--color-cyan-500` | `#00D9F5` | Info accent |
| `--color-gold-500` | `#F7971E` | Premium/gold |
| `--color-gold-400` | `#FFD200` | Premium/gold light |

### 1.2 Color — Semantic Tokens

| Token | Maps To | Usage |
|---|---|---|
| `--bg-primary` | `#0a0a0a` | Page/app background |
| `--bg-secondary` | `#111111` | Sidebar, panels |
| `--bg-card` | `#1a1a1a` | Card, modal surfaces |
| `--bg-glass` | `rgba(255,255,255,0.06)` | Glassmorphism fills |
| `--border-glass` | `rgba(255,255,255,0.08)` | Glassmorphism borders |
| `--border-subtle` | `rgba(255,255,255,0.12)` | Dividers, input borders |
| `--text-primary` | `#ffffff` | Body copy, headings |
| `--text-secondary` | `rgba(255,255,255,0.60)` | Meta, labels, captions |
| `--text-tertiary` | `rgba(255,255,255,0.35)` | Placeholder, disabled |
| `--color-primary` | `#FF6B35` | CTAs, active states, like |
| `--color-primary-glow` | `rgba(255,107,53,0.30)` | Primary button shadow |
| `--color-success` | `#00F5A0` | Live badge, success state |
| `--color-error` | `#FF4D4D` | Error states, validation |
| `--color-warning` | `#FFD200` | Warnings, gold badges |

### 1.3 Gradient Tokens

| Token | Value | Usage |
|---|---|---|
| `--gradient-primary` | `linear-gradient(135deg, #FF6B35, #FF3CAC)` | Primary CTAs, accent highlights |
| `--gradient-rainbow` | `linear-gradient(135deg, #FF3CAC, #784BA0, #2B86C5)` | Story rings, decorative |
| `--gradient-success` | `linear-gradient(135deg, #00F5A0, #00D9F5)` | Live badge, success fills |
| `--gradient-gold` | `linear-gradient(135deg, #F7971E, #FFD200)` | Premium/verified badges |
| `--gradient-overlay` | `linear-gradient(to top, rgba(0,0,0,0.85) 0%, transparent 60%)` | Video overlay scrim |

### 1.4 Border Radius Tokens

| Token | Value | Usage |
|---|---|---|
| `--radius-xs` | `6px` | Tags, chips, inline badges |
| `--radius-sm` | `12px` | Buttons, small cards |
| `--radius-md` | `16px` | Cards, modals, inputs |
| `--radius-lg` | `24px` | Large cards, bottom sheets |
| `--radius-xl` | `32px` | Phone frames, hero cards |
| `--radius-full` | `9999px` | Avatars, pill buttons, story rings |

### 1.5 Elevation / Shadow Tokens

| Token | Value | Usage |
|---|---|---|
| `--shadow-sm` | `0 2px 8px rgba(0,0,0,0.4)` | Small cards, inputs |
| `--shadow-md` | `0 8px 24px rgba(0,0,0,0.5)` | Modals, dropdowns |
| `--shadow-lg` | `0 16px 48px rgba(0,0,0,0.6)` | Full-screen overlays |
| `--shadow-glow-primary` | `0 0 20px rgba(255,107,53,0.4)` | Liked hearts, active CTAs |
| `--shadow-glow-glass` | `0 8px 32px rgba(0,0,0,0.3)` | Glass cards |

### 1.6 Transition Tokens

| Token | Value | Usage |
|---|---|---|
| `--duration-fast` | `150ms` | Hover, toggle, micro-states |
| `--duration-normal` | `250ms` | Button press, color shifts |
| `--duration-slow` | `400ms` | Modal open/close, panel slide |
| `--ease-default` | `cubic-bezier(0.25, 0.46, 0.45, 0.94)` | Standard ease-out |
| `--ease-spring` | `cubic-bezier(0.34, 1.56, 0.64, 1.00)` | Spring/bounce (like button) |
| `--ease-linear` | `linear` | Progress bars, spinners |

---

## 2. Typography Scale

Font stack: `'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif`  
Display font: `'Space Grotesk', sans-serif` (headings, hero copy)  
Loading strategy: preconnect to `fonts.googleapis.com`, display=swap (FOUT accepted).

| Token | Size | Line Height | Letter Spacing | Weight | Usage |
|---|---|---|---|---|---|
| `--text-display` | 56px | 64px | -0.02em | 800 | Hero headline (index) |
| `--text-heading-lg` | 36px | 44px | -0.02em | 700 | Section headings |
| `--text-heading-md` | 24px | 32px | -0.01em | 700 | Card titles, modal headings |
| `--text-heading-sm` | 18px | 26px | -0.01em | 600 | Sidebar section labels |
| `--text-body-lg` | 16px | 24px | 0 | 400 | Body copy, descriptions |
| `--text-body-md` | 14px | 20px | 0 | 400 | Labels, captions, meta |
| `--text-body-sm` | 12px | 16px | 0.01em | 400 | Timestamps, counts, tags |
| `--text-label` | 11px | 14px | 0.04em | 600 | Sidebar nav labels, badges |

**Mobile adjustments:**
- `--text-display` → 32px on `max-width: 768px`
- `--text-heading-lg` → 24px on `max-width: 768px`

---

## 3. Spacing System

Base unit: **4px**. All spacing values are multiples of 4.

| Token | Value | Common Use |
|---|---|---|
| `--space-1` | `4px` | Icon-to-label gap, tight padding |
| `--space-2` | `8px` | Inline element gaps |
| `--space-3` | `12px` | Small padding, button icon gap |
| `--space-4` | `16px` | Standard padding, list item gap |
| `--space-5` | `20px` | Medium spacing |
| `--space-6` | `24px` | Section inner padding |
| `--space-8` | `32px` | Card padding, section gaps |
| `--space-10` | `40px` | Large section padding |
| `--space-12` | `48px` | Section vertical margin |
| `--space-16` | `64px` | Hero padding, major sections |

---

## 4. Component Hierarchy

```
App
├── feed.html (Feed View)
│   ├── SidebarNav
│   │   ├── Logo
│   │   └── NavItem × 9
│   ├── FeedContainer
│   │   └── ReelCard × N
│   │       ├── VideoPlayer
│   │       │   └── VideoLoader (spinner)
│   │       ├── VideoProgressBar
│   │       ├── VideoOverlay (scrim gradient)
│   │       ├── CreatorInfo
│   │       │   ├── CreatorAvatar
│   │       │   ├── CreatorName
│   │       │   └── FollowButton
│   │       ├── ReelDescription
│   │       │   ├── HashtagList
│   │       │   └── MusicInfo
│   │       ├── ActionButtonStrip
│   │       │   ├── LikeButton
│   │       │   ├── CommentButton
│   │       │   ├── ShareButton
│   │       │   ├── SaveButton
│   │       │   └── MoreButton
│   │       ├── FloatingHearts (conditional)
│   │       └── LiveComments
│   ├── RightSidebar (≥1100px)
│   │   ├── SuggestedProfile × 5
│   │   └── TrendingTag × 6
│   ├── NavigationArrows
│   ├── CommentsModal (conditional)
│   │   ├── CommentList
│   │   │   └── CommentItem × N
│   │   └── CommentInput
│   ├── ShareModal (conditional)
│   │   ├── SharePlatformButton × 6
│   │   └── CopyLinkRow
│   └── ToastNotification (conditional)
│
└── index.html (Marketing View)
    ├── Navbar
    │   ├── Logo
    │   ├── NavLinks
    │   └── CTAButton
    ├── HeroSection
    ├── BeforeAfterSection
    ├── PhoneShowcaseSection
    ├── FeaturesSection
    ├── ComparisonSection
    ├── ImpactSection
    └── CTASection
```

---

## 5. Component Specs

### 5.1 Sidebar Navigation

**Dimensions:** 72px wide, 100vh tall, fixed position left.  
**Background:** `--bg-secondary` with `border-right: 1px solid --border-glass`.

#### States

| State | Description |
|---|---|
| **Empty** | Not applicable — nav items are always present |
| **Loading** | Skeleton shimmer on nav item icons (200ms stagger) |
| **Error** | Not applicable — static nav |
| **Default** | All items rendered, no active state |
| **Active Item** | Selected item has icon tinted `--color-primary`, background pill `--bg-glass`, label visible |

#### Anatomy

```
┌──────────┐
│  [Logo]  │  ← 40×40px logo mark, centered
├──────────┤
│  [Icon]  │  ← NavItem (repeated × 9)
│  label   │
├──────────┤
│  [Avatar]│  ← User avatar (bottom, 32×32px circle)
└──────────┘
```

#### NavItem Spec

- **Size:** 56px tall × 100% wide
- **Icon:** 24×24px, `--text-secondary`
- **Label:** `--text-label` (11px, 600 weight), below icon, hidden unless active or hover
- **Touch target:** Full 56px height, full sidebar width (≥44×44px)
- **Hover:** `background: --bg-glass`, icon color → `--text-primary`, transition: `--duration-fast`
- **Active:** `background: --bg-glass`, icon color → `--color-primary`, left border `2px solid --color-primary`

#### Keyboard Behavior

| Key | Action |
|---|---|
| `Tab` | Focus next nav item |
| `Shift+Tab` | Focus previous nav item |
| `Enter` / `Space` | Activate nav item |
| `ArrowUp` / `ArrowDown` | Move between items |

**ARIA:** `<nav aria-label="Main navigation">`, each item is `<a role="link" aria-current="page">` when active.

---

### 5.2 Reel Card

**Dimensions:** 100vw × 100dvh (full screen on mobile), max 480px wide centered on desktop.  
**Background:** `--bg-primary` behind video.

#### States

| State | Visual |
|---|---|
| **Empty** | Black screen — not shown; first reel preloads immediately |
| **Loading** | Centered spinner (`--color-primary`, 40px, 1s rotation), video area is black |
| **Error** | Centered icon + "Video unavailable" copy in `--text-secondary`, retry button |
| **Partial** | Video buffering mid-play: spinner overlay appears at center, progress bar pauses |
| **Complete** | Video playing, all overlays visible |

#### Layer Stack (bottom to top)

```
[ Video Element ]              z-index: 1
[ Gradient Scrim Overlay ]     z-index: 2  — covers bottom 60% of frame
[ Video Progress Bar ]         z-index: 3  — 3px, top of card
[ Creator Info ]               z-index: 4  — bottom-left
[ Reel Description ]           z-index: 4  — above creator info
[ Live Comments ]              z-index: 5  — floating, left side
[ Action Button Strip ]        z-index: 6  — right side, vertically centered
[ Floating Hearts ]            z-index: 7  — conditional, center-right
[ Play/Pause Overlay ]         z-index: 8  — full-card, tap-through after 800ms
```

#### Video Element

- `object-fit: cover`, `width: 100%`, `height: 100%`
- `autoplay`, `muted`, `loop`, `playsinline` attributes
- `preload="auto"` for current ±1 videos, `preload="none"` for rest

#### Creator Info (bottom-left)

- Position: `bottom: 80px`, `left: 16px`, `right: 96px` (leaves room for action strip)
- **Avatar:** 40×40px circle, border `2px solid rgba(255,255,255,0.3)`
- **Creator name:** `--text-body-lg`, 600 weight, `--text-primary`, truncated at 1 line
- **Follow button:** Pill, 60×26px, `border: 1.5px solid --text-primary`, `--text-sm`, appears when not following
- **Music info:** Row with spinning vinyl icon (16px) + marquee text at `--text-body-sm`
- **Tags:** `#hashtag` in `--color-primary`, `--text-body-sm`, max 3 tags visible

#### Scrim Overlay

- `background: --gradient-overlay`
- Covers bottom 60% of video frame
- Ensures all overlay text meets 4.5:1 contrast ratio against video content

---

### 5.3 Action Button Strip

**Position:** Right side of ReelCard, vertically centered, `right: 12px`.  
**Layout:** Vertical stack, `gap: --space-6` between buttons.

#### Button Anatomy

Each action button is a column: icon (32×32px) + count label below.

```
  ┌──────┐
  │  ♡   │  ← 32×32 icon, --text-primary
  └──────┘
  1.2K     ← --text-body-sm, --text-secondary
```

#### Individual Button States (all 5)

**Like Button**

| State | Visual |
|---|---|
| Empty | Heart outline, `--text-primary`, count "0" hidden |
| Loading | Not applicable |
| Error | Heart outline unchanged (optimistic UI, silent retry) |
| Unliked | Heart outline `♡`, white |
| Liked | Heart filled `♥`, `--color-primary` + `--shadow-glow-primary`, count increments, `actB` bounce animation fires |

- Double-tap on video → triggers like + floating hearts
- Single tap on button → toggle like

**Comment Button**

| State | Visual |
|---|---|
| Default | Speech bubble icon, `--text-primary` |
| Active (modal open) | Icon tinted `--color-primary` |
| Pressed | Scale 0.85, `--duration-fast` |

- Tap → opens CommentsModal

**Share Button**

| State | Visual |
|---|---|
| Default | Arrow-up-from-box icon, `--text-primary` |
| Pressed | Scale 0.85, `--duration-fast` |

- Tap → opens ShareModal

**Save Button**

| State | Visual |
|---|---|
| Unsaved | Bookmark outline, `--text-primary` |
| Saved | Bookmark filled, `--color-primary` |
| Pressed | Scale 0.85 → 1.1 spring rebound |

**More Button**

| State | Visual |
|---|---|
| Default | Ellipsis `⋯`, `--text-secondary` |
| Pressed | Scale 0.85 |

- Tap → opens bottom sheet (spec TBD in v1.1)

#### Minimum touch target

All buttons: 44×44px touch area (visual icon may be smaller).

---

### 5.4 Comments Modal

**Type:** Slide-up bottom sheet on mobile; centered overlay on desktop.  
**Trigger:** CommentButton tap.  
**Backdrop:** `rgba(0,0,0,0.7)` with `backdrop-filter: blur(4px)`.

#### States

| State | Visual |
|---|---|
| **Empty** | "No comments yet. Be the first!" with an illustration, comment input visible |
| **Loading** | 3-row skeleton shimmer (avatar circle + two lines each) |
| **Error** | "Couldn't load comments" + retry CTA |
| **Partial** | Comments load in batches; "Load more" at bottom, spinner inline |
| **Complete** | All comments visible, input ready |

#### Anatomy

```
┌─────────────────────────────────┐
│  ━━━  (drag handle)             │
│  Comments  (1,234)    [✕]       │ ← heading + count + close
├─────────────────────────────────┤
│ [Avatar] Creator        2h      │ ← CommentItem
│          Comment text here...   │
│          [♡ 42]  [Reply]        │
│ ─────────────────────────────── │
│ [Avatar] ...                    │
│         ...                     │
├─────────────────────────────────┤
│ [Avatar] [Write a comment... ] [→]│ ← CommentInput
└─────────────────────────────────┘
```

#### Modal Dimensions

- **Mobile:** Full width, 70vh max-height, `border-radius: --radius-lg --radius-lg 0 0`
- **Desktop:** 480px wide, 560px max-height, `border-radius: --radius-lg`, centered

#### CommentItem Spec

- Avatar: 32×32px circle
- Name: `--text-body-md` 600 weight, `--text-primary`
- Timestamp: `--text-body-sm`, `--text-tertiary`
- Comment text: `--text-body-md`, `--text-secondary`
- Like icon: 14px heart, count in `--text-body-sm`
- Reply label: `--text-body-sm`, `--text-secondary`

#### CommentInput Spec

- Height: 48px
- Background: `--bg-glass`, border `--border-subtle`
- Placeholder: "Add a comment...", `--text-tertiary`
- Send button: `--color-primary` when text present, `--text-tertiary` when empty
- On submit: optimistically prepend comment, clear input, show success toast

#### Dismiss

- Tap backdrop → close
- Drag handle swipe-down → close (spring back if threshold not met)
- `Escape` key → close

#### Keyboard / ARIA

- `role="dialog"`, `aria-labelledby="comments-heading"`, `aria-modal="true"`
- Focus trap: `Tab` cycles through interactive elements within modal
- On open: focus first comment or input if empty
- On close: return focus to trigger button
- `aria-live="polite"` on comment list for new comment announcements

---

### 5.5 Share Modal

**Type:** Center overlay.  
**Trigger:** ShareButton tap.

#### States

| State | Visual |
|---|---|
| Empty | N/A (static list of platforms) |
| Loading | N/A |
| Error | "Share failed" toast if deep-link fails |
| Default | Platform grid + copy link row |
| Link Copied | "Copied!" badge for 2s replaces copy icon |

#### Anatomy

```
┌────────────────────────┐
│  Share to       [✕]    │
├────────────────────────┤
│  [WA] [TG] [IG] [TW]  │ ← platform icons, 56×56px circles
│  [FB] [More]           │
├────────────────────────┤
│  https://sha.re/xyz  [Copy]│ ← link row
└────────────────────────┘
```

#### Platform Button Spec

- 56×56px circle, background: brand color at 15% opacity
- Icon: 28×28px, brand color
- Label: `--text-body-sm`, `--text-secondary`, below icon
- Hover: background opacity → 25%, scale 1.05, `--duration-fast`
- Tap: opens native share or deep-link

Platforms: WhatsApp (#25D366), Telegram (#2AABEE), Instagram (gradient), Twitter/X (#1DA1F2), Facebook (#1877F2), More (system share sheet)

#### Copy Link Row

- Input-like display (not editable): `--bg-glass`, `--radius-sm`, truncated URL
- "Copy" button: `--color-primary`, becomes "Copied!" for 2s on success
- Link uses shortened format `sha.re/{id}`

---

### 5.6 Floating Hearts Animation

**Trigger:** Like button tap OR double-tap on video.  
**Duration:** 1.3s–1.8s per heart.  
**Count:** 8 hearts spawned simultaneously.

#### Heart Specs

- Size: 20px–32px (randomized per heart)
- Color: `--color-primary` (#FF6B35) with `--shadow-glow-primary`
- Start position: near the like button (right side, vertically centered)
- End position: 120–180px upward from start (randomized x drift ±30px)
- Stagger: 0ms, 50ms, 100ms, 150ms, 200ms, 250ms, 300ms, 350ms
- Rotation: random ±20deg during flight
- Opacity: 1 → 0 (fades in last 40% of animation)
- Easing: `--ease-default` for vertical travel

#### Keyframe Reference

```
@keyframes floatUp {
  0%   { transform: translateY(0)   rotate(0deg);   opacity: 1;   }
  60%  { transform: translateY(-80px) rotate(10deg); opacity: 0.8; }
  100% { transform: translateY(-160px) rotate(-10deg); opacity: 0; }
}
```

#### Accessibility

- `aria-hidden="true"` on all heart elements (purely decorative)
- No reduced-motion consideration needed (sub-500ms, decorative)
- Provide non-animated feedback: like count increments, button turns red

---

### 5.7 Toast Notification

**Position:** `bottom: 80px`, horizontally centered.  
**Width:** Max 320px, auto height.

#### States

| State | Visual |
|---|---|
| Hidden | `opacity: 0`, `transform: translateY(16px)` |
| Entering | Animate to `opacity: 1`, `translateY(0)`, 250ms ease-out |
| Visible | Shown for 2000ms |
| Exiting | Animate to `opacity: 0`, `translateY(8px)`, 200ms ease-in |

#### Variants

| Variant | Background | Icon |
|---|---|---|
| Success | `rgba(0,245,160,0.15)` + `border: 1px solid --color-success` | ✓ green |
| Error | `rgba(255,77,77,0.15)` + `border: 1px solid --color-error` | ✕ red |
| Info | `--bg-glass` + `--border-glass` | ℹ blue |

- Text: `--text-body-md`, `--text-primary`
- Padding: `--space-3 --space-5`
- Border radius: `--radius-sm`
- `aria-live="assertive"` for errors, `aria-live="polite"` for success/info

---

### 5.8 Stories Bar

**Context:** Appears in phone mockup on index.html marketing page and at top of feed (planned).  
**Layout:** Horizontal scroll, `gap: --space-3`, no scrollbar visible.

#### Story Ring States

| State | Visual |
|---|---|
| Unseen | Gradient ring `--gradient-rainbow`, 2px, 2px gap |
| Seen | `--text-tertiary` ring, 1.5px |
| Loading | Shimmer placeholder circle |
| Playing | Animated arc progress around avatar |

#### Story Item Spec

- Avatar: 52×52px circle
- Ring: 3px border with 2px transparent gap between ring and avatar
- Label: `--text-body-sm`, `--text-secondary`, centered below, max 1 line truncated 8 chars
- Touch target: 64×80px (avatar + label area)

---

### 5.9 Right Sidebar (Suggestions)

**Dimensions:** 280px wide, 100vh, fixed right.  
**Visibility:** Only at viewport ≥1100px.  
**Background:** `--bg-secondary`, `border-left: 1px solid --border-glass`.

#### Suggested Profile Card

```
[Avatar 40px] Creator Name      [Follow]
              @handle
              123K followers
```

- Avatar: 40×40px circle
- Name: `--text-body-md`, 600 weight, `--text-primary`
- Handle: `--text-body-sm`, `--text-tertiary`
- Follower count: `--text-body-sm`, `--text-secondary`
- Follow button: Pill, 64×28px, `--gradient-primary` background, `--text-primary` 12px

**Follow button states:**
- Default: gradient background, "Follow"
- Following: `--bg-glass` border, "Following", with checkmark
- Hover (default): brightness 1.1

#### Trending Tags Section

- Section label: `--text-heading-sm`, `--text-secondary`
- Tag chip: `--bg-glass` fill, `--border-glass`, `--radius-xs`, `--text-body-sm`, `--text-secondary`
- Hashtag text in `--color-primary`
- Tag count in `--text-tertiary`

---

### 5.10 Video Progress Bar

**Position:** `top: 0`, full width of reel card.  
**Height:** 3px.

#### States

| State | Visual |
|---|---|
| Loading | No bar visible |
| Playing | Orange fill animating from 0%→100% over 8s linear |
| Paused | Fill stops at current position |
| Complete | Instant reset to 0% on next video |

- Fill color: `--color-primary`
- Track: `rgba(255,255,255,0.2)`
- Animation: `progress 8s linear forwards` (restarts on video change)

---

## 6. Layout Wireframes

### 6.1 Feed View — Desktop (≥1100px)

```
┌────────────────────────────────────────────────────────┐
│ SIDEBAR (72px)    │  FEED (flex 1, max 480px centered) │ RIGHT SIDEBAR (280px) │
│                   │                                     │                       │
│  [Logo]           │  ┌──────────────────────────────┐  │  [Profile suggest ×5] │
│                   │  │ [Video Progress Bar top 3px]  │  │                       │
│  [Home] ← active  │  │                               │  │  Trending Tags        │
│  [Discover]       │  │                               │  │  #tag1  12.4K         │
│  [Create +]       │  │   VIDEO CONTENT (cover)       │  │  #tag2  8.9K          │
│  [Notifications]  │  │                               │  │  #tag3  6.1K          │
│  [Messages]       │  │                               │  │                       │
│  [Profile]        │  │                               │  │                       │
│  [Saved]          │  │                               │  │                       │
│  [Settings]       │  │                               │  │                       │
│                   │  │  [Creator info] bottom-left   │  │                       │
│  [Avatar]         │  │  [Hashtags + music]           │  │                       │
│  bottom           │  │              [♡][💬][↑][🔖][⋯]│  │                       │
│                   │  └──────────────────────────────┘  │                       │
│                   │  (scroll-snap, next reel below)     │                       │
└────────────────────────────────────────────────────────┘

Navigation arrows: fixed right of feed, above/below center
  [↑] 40×40px circle, --bg-glass
  [↓] 40×40px circle, --bg-glass
```

**Grid values:**
- Left sidebar: `72px` fixed
- Feed column: `1fr`, centered, `max-width: 480px` with `margin: auto`
- Right sidebar: `280px` fixed
- Gap between sidebar and feed: `--space-8`

### 6.2 Feed View — Mobile (<768px)

```
┌─────────────────────────┐
│ [Video Progress Bar]    │
│                         │
│                         │
│   VIDEO (full screen)   │
│   100vw × 100dvh        │
│                         │
│                         │
│                         │
│ [Creator info]  [♡]     │
│ [#tags] [♫ music] [💬]  │
│                  [↑]    │
│                  [🔖]   │
│                  [⋯]    │
└─────────────────────────┘
No sidebars. No navigation arrows.
Bottom system nav bar: 50px safe area padding.
```

**Mobile-specific adjustments:**
- Action strip: `right: 12px`, `bottom: calc(80px + env(safe-area-inset-bottom))`
- Creator info: `bottom: calc(80px + env(safe-area-inset-bottom))`
- Touch targets: all ≥44×44px

### 6.3 Marketing Page — Desktop (index.html)

```
┌────────────────────────────────────────────┐
│  NAVBAR  (fixed, height 72px)              │
│  Logo    Home  Discover  Features  [CTA]   │
├────────────────────────────────────────────┤
│                 HERO                        │
│    [Eyebrow badge]                         │
│    Headline 56px / 800 weight              │
│    Sub-copy 18px                           │
│    [Primary CTA]  [Secondary CTA]          │
│    Stats row: 325M · 80M · 45 langs        │
│                 [Phone mockup]             │
├────────────────────────────────────────────┤
│              BEFORE / AFTER                │
│    Old Design   |   Reimagined             │
│    (phone mock) |   (phone mock)           │
├────────────────────────────────────────────┤
│             PHONE SHOWCASE                 │
│    3 rotating phone frames (tabs)          │
├────────────────────────────────────────────┤
│              FEATURES GRID                 │
│    [Icon] Title        [Icon] Title        │
│    Description         Description         │
├────────────────────────────────────────────┤
│           COMPARISON TABLE                 │
│    Feature | Current | Reimagined          │
│    ...      |  ✗      |  ✓                │
├────────────────────────────────────────────┤
│              IMPACT METRICS                │
│    +40%       +65%      -30%     ×2        │
│    Engagement Completion  Drop   Creation  │
├────────────────────────────────────────────┤
│                 CTA                        │
│    "Ready to reimagine?"  [Button]         │
└────────────────────────────────────────────┘

Max content width: 1200px, centered, 24px horizontal padding.
```

---

## 7. Responsive Behavior

### Breakpoints

| Name | Min Width | Description |
|---|---|---|
| `mobile` | 0px | Single column, no sidebars, full-screen video |
| `tablet` | 768px | Sidebars may appear at reduced width |
| `desktop` | 1100px | Full 3-column layout with both sidebars |
| `wide` | 1400px | Max content width caps, centers layout |

### Component Visibility by Breakpoint

| Component | Mobile (<768px) | Tablet (768-1100px) | Desktop (≥1100px) |
|---|---|---|---|
| Left Sidebar | Hidden | Hidden | Visible (72px) |
| Right Sidebar | Hidden | Hidden | Visible (280px) |
| Navigation Arrows | Hidden | Hidden | Visible |
| Video Progress Bar | Visible | Visible | Visible |
| Creator Info | Simplified (no follow) | Full | Full |
| Reel Description | Truncated (2 lines) | Full | Full |
| Live Comments | Hidden | Visible | Visible |

### Layout Shift Rules

- No layout shift (CLS) on video load — reserve space with aspect ratio containers
- Skeleton loaders for all async content areas prevent reflow
- Font sizes do not change on orientation change (use `dvh` for height units)

---

## 8. Interaction Specifications

### 8.1 Video Playback Interactions

| Gesture / Input | Action | Feedback |
|---|---|---|
| Single tap | Toggle play/pause | Play icon overlays center for 800ms then fades, progress bar pauses/resumes |
| Double tap | Like + float hearts | 8 hearts animate, like button activates, count increments |
| Swipe up / Arrow ↓ | Next reel | Smooth scroll-snap, new video autoplays, progress bar resets |
| Swipe down / Arrow ↑ | Previous reel | Smooth scroll-snap, video plays from beginning |
| Scroll past end | Loop to first | Seamless wrap (confirm behavior in v1.1) |

### 8.2 Follow Button Interaction

1. User taps "Follow"
2. Optimistically: button text → "Following", style switches to outlined
3. API call in background
4. On success: persisted state
5. On failure: revert to "Follow", show error toast
6. If already following: tap → "Unfollow?" confirmation (300ms delay before reverting to avoid accidental untaps)

### 8.3 Comment Submission

1. Input has focus, text is non-empty → send button activates (`--color-primary`)
2. User taps send
3. Optimistic: new comment prepends to list with avatar + text + "Just now", input clears
4. API call
5. On success: comment timestamp updates to real time
6. On failure: comment marked with retry icon, inline error text

### 8.4 Share Platform Tap

1. Platform icon taps → immediate haptic (if supported, `navigator.vibrate(10)`)
2. Native deep link opens (WhatsApp: `https://wa.me/?text=`, etc.)
3. Modal closes after 300ms delay
4. If link copy: clipboard write → button shows "Copied!" for 2000ms

### 8.5 Keyboard Navigation (Desktop)

| Key | Action |
|---|---|
| `↑` / `k` | Previous reel |
| `↓` / `j` | Next reel |
| `Space` | Toggle play/pause |
| `L` | Like |
| `C` | Open comments |
| `S` | Save |
| `Escape` | Close modal |
| `Tab` | Move focus forward through sidebar items |

---

## 9. Animation Catalogue

| Name | Trigger | Duration | Easing | Properties |
|---|---|---|---|---|
| `fadeUp` | Page section enter viewport | 800ms | ease-out | `opacity 0→1`, `translateY 20px→0` |
| `floatUp` (hearts) | Like action | 1300–1800ms | ease-default | `translateY 0→-160px`, `opacity 1→0`, `rotate ±20deg` |
| `actB` (button bounce) | Like/Save activate | 400ms | spring | `scale 1→0.85→1.2→1` |
| `spin` (vinyl) | Always while playing | 3s infinite | linear | `rotate 360deg` |
| `pulse` | Live badge | 2s infinite | ease-in-out | `opacity 1→0.4→1` |
| `float` (decorative) | Page load | 6s infinite | ease-in-out | `translateY 0→-10px→0` |
| `progress` (bar) | Video play | 8s | linear | `width 0→100%` |
| `lc` (live comment) | Simulated stream | 4s | ease-out | `translateY 30px→0`, `opacity 0→0.9→0` |
| Modal slide-in (mobile) | Modal open | 400ms | ease-default | `translateY 100%→0` |
| Modal fade-in (desktop) | Modal open | 250ms | ease-out | `opacity 0→1`, `scale 0.95→1` |
| Modal exit | Modal close | 200ms | ease-in | Reverse of enter |
| Toast enter | Toast show | 250ms | ease-out | `translateY 16px→0`, `opacity 0→1` |
| Toast exit | Auto-dismiss | 200ms | ease-in | `translateY 0→8px`, `opacity 1→0` |

**Reduced motion:** For users with `prefers-reduced-motion: reduce`, disable `floatUp`, `actB`, `float`, `lc`. Keep `fadeUp` but reduce to `opacity` only (no translate). Keep `progress` bar.

---

## 10. Accessibility Specification

### 10.1 Color Contrast

All text/background pairs must meet WCAG 2.1 AA (4.5:1 for normal text, 3:1 for large text).

| Foreground | Background | Ratio | Status |
|---|---|---|---|
| `--text-primary` (#fff) | `--bg-primary` (#0a0a0a) | 19.1:1 | ✓ AAA |
| `--text-secondary` (rgba 60%) | `--bg-primary` | 8.2:1 | ✓ AAA |
| `--text-tertiary` (rgba 35%) | `--bg-primary` | 4.1:1 | ⚠ Fails AA |
| `--text-primary` (#fff) | `--bg-card` (#1a1a1a) | 14.7:1 | ✓ AAA |
| `--color-primary` (#FF6B35) | `--bg-primary` | 3.9:1 | ⚠ Fails AA body text |
| `--text-primary` (#fff) | `--color-primary` (#FF6B35) | 3.2:1 | ✓ Large text / UI |

**Action items:**
1. `--text-tertiary` should only be used for decorative/non-essential text (timestamps, counts). Never for primary labels.
2. `--color-primary` should not be used as body text color on dark backgrounds. Use it for icons and button labels only at sizes ≥18px.
3. Increase `--text-secondary` opacity to 0.70 for text paired with `--bg-card` backgrounds to achieve 4.5:1.

### 10.2 Focus Indicators

- All interactive elements: `outline: 2px solid --color-primary`, `outline-offset: 2px`
- No `outline: none` unless a custom focus style is provided
- Focus ring must be visible over the video background (add white outline underneath: `outline: 2px solid white, 4px solid --color-primary`)

### 10.3 Screen Reader Announcements

| Event | `aria-live` | Message |
|---|---|---|
| Like toggled | `polite` | "Video liked" / "Like removed" |
| Follow toggled | `polite` | "@username followed" / "Unfollowed" |
| Comment posted | `polite` | "Comment posted" |
| Link copied | `polite` | "Link copied to clipboard" |
| Modal opened | (focus moves, no announcement needed) | — |
| Error toast | `assertive` | Error message text |

### 10.4 ARIA Roles Reference

| Component | Role | Key Attributes |
|---|---|---|
| Sidebar Nav | `navigation` | `aria-label="Main navigation"` |
| NavItem | `link` | `aria-current="page"` when active |
| ReelCard | `region` | `aria-label="Reel by {creator}"` |
| Video | native `<video>` | `aria-label="{description}"` |
| Like Button | `button` | `aria-pressed="true/false"`, `aria-label="Like, {count} likes"` |
| Comment Button | `button` | `aria-label="Comments, {count} comments"` |
| Share Button | `button` | `aria-label="Share"` |
| Save Button | `button` | `aria-pressed="true/false"`, `aria-label="Save reel"` |
| Comments Modal | `dialog` | `aria-modal="true"`, `aria-labelledby="comments-title"` |
| Share Modal | `dialog` | `aria-modal="true"`, `aria-labelledby="share-title"` |
| Toast | `status` or `alert` | `aria-live="polite"` / `"assertive"` |
| Progress Bar | `progressbar` | `aria-valuenow`, `aria-valuemin="0"`, `aria-valuemax="100"` |

### 10.5 Touch Targets

- Minimum touch target: **44×44px** on all interactive elements
- Spacing between adjacent touch targets: **≥8px** (no accidental double-tap)
- Action button strip: buttons are 40px icons with 4px gap → ensure wrapper div is 44px tall

### 10.6 Motion and Animation

- All animations >250ms must have a reduced-motion alternative
- The floating hearts (1.3–1.8s) use `@media (prefers-reduced-motion: reduce)` to skip entirely
- Progress bar is exempt (conveys time information, not decorative)
- Modal transitions: replace slide/scale with instant show/hide under reduced motion

---

*This document is a living spec. Version 1.1 will cover: More button bottom sheet, notification panel, search view, and profile page.*
