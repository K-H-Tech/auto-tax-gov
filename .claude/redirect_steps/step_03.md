# Step 3: 200

**Timestamp:** 2025-11-26T16:42:31+03:30

**URL:** `https://sso.my.gov.ir/login`

## Request

**Method:** GET

### Request Headers

```http
Cookie: SESSION=MTg0Y2Y4NzctZTQxYi00YjYwLTkzNjMtZTU3Yjg4Y2Y4YjU2; f5avraaaaaaaaaaaaaaaa_session_=HGPHMFNKOBMOMLFAGEJHMPHDBFIEJJKKMMGKAMLLGNCICDJCEIPBEGMJPCODPONANAEDGCDDCMDFAFBFJAKAGKLKDPGKJMKHJAHLFKPELICGAJDMOGHOMINMBPEBDPLE
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Connection: keep-alive
Upgrade-Insecure-Requests: 1
```

## Response

**Status Code:** 200

**Status:** 200

### Response Headers

```http
X-Content-Type-Options: nosniff
X-Content-Type-Options: nosniff
Pragma: no-cache
Content-Type: text/html;charset=UTF-8
Access-Control-Expose-Headers: *
Content-Language: fa
Access-Control-Allow-Headers: Content-Type,Sec-Fetch-Site, Sec-Fetch-Mode, Sec-Fetch-Dest, Origin, Cross-Domain, Cache-Control, Referer, User-Agent, Host, Accept, Accept-Encoding, Accept-Language, Authorization, cachecontrol,X-HTTP-Method-Override
Access-Control-Allow-Origin: https://my.gov.ir
Access-Control-Max-Age: 3600
Set-Cookie: s2=51284839-a635-4d2c-ae22-135e3ee985fd; HttpOnly
X-Xss-Protection: 0
X-Xss-Protection: 1; mode=block
X-Powered-By: ITO
Server: ITO
Referrer-Policy: no-referrer-when-downgrade
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
Date: Wed, 26 Nov 2025 13:12:49 GMT
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Expires: 0
Connection: keep-alive
Access-Control-Allow-Methods: *
Vary: origin,access-control-request-method,access-control-request-headers,accept-encoding
X-Frame-Options: DENY
X-Frame-Options: DENY
```

### Captcha

**Captcha Key:** `cfe4a49f-0c94-4991-9465-3e3b45aa4df8`

**CSRF Token:** `yAtpHmp7GsEKyyxyDAxnFXTIEAYmsSd2_4TL6W_0UWedrJKP8Gpafw4eePMn8k9KOyFTcRHwPT9C0kNbmrH4iFuRaFeuzarp`

**Auth Header:** `Basic cHdhLnVzZXI6MDFhbnNraXludXd2czRzYnRvamE=`

**Captcha Image:** ![captcha](step_03_captcha.png)

**Base64 Data (first 100 chars):** `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAA8CAYAAACJvWQEAAApA0lEQVR4Xu1dB1xUx/be9PLyUl+S9/...`

### Response Body

```html
<!doctype html>
<html lang="fa" dir="rtl">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1.0, maximum-scale=1"
    />
    <!-- <link rel="stylesheet" href="../static/web/assets/css/common.css" /> -->
    <!-- <link rel="stylesheet" href="../static/web/assets/css/newApp.css" /> -->
    <!-- <link rel="stylesheet" href="../static/web/assets/css/tailwindcss-v4.css" /> -->

    <title>پنجره ملی خدمات دولت هوشمند - ورود یکپارچه</title>
    <link
      rel="apple-touch-icon"
      sizes="76x76"
      href="../static/web/assets/images/apple-touch-icon.png"
    />
    <link
      rel="icon"
      type="image/png"
      sizes="32x32"
      href="../static/web/assets/images/favicon-32x32.png"
    />
    <link
      rel="icon"
      type="image/png"
      sizes="16x16"
      href="../static/web/assets/images/favicon-16x16.png"
    />
    <link
      rel="mask-icon"
      href="../static/web/assets/images/safari-pinned-tab.svg"
      color="#5bbad5"
    />

    <meta name="description" content="«پنجره ملی خدمات دولت هوشمند» یک راه امن برای دسترسی آنلاین به خدمات دولتی از طریق یک درگاه واحد است. این درگاه، پلتفرم ملی یکپارچه ارائه خدمات دولتی و عمومی به مردم بوده و هدف آن ایجاد سهولت در ارائه این خدمات به شهروندان کشور است؛ به‌طوری‌که تنها با یک بار احرازهویت می‌توان خدمات موردنظر را دریافت نموده و یا به سامانه‌های متفاوت ارائه دهنده خدمات وارد شد" />

    <style>
  @font-face {
    font-family: peyda;
    src: url(../static/web/assets/fonts/peyda/woff/PeydaWebFaNum-Regular.woff);
  }

  body {
    font-family: peyda;
    user-select: none;
  }

  ::-webkit-scrollbar {
    display: none;
  }

  .form-background {
    background-image: url(../static/web/assets/images/new-my/form-bg1.png);
    background-size: contain;
    background-position: center;
    background-repeat: no-repeat;
  }

  .linearBorder {
    position: relative;

    &::before {
      content: "";
      position: absolute;
      pointer-events: none;
      inset: 0;
      padding: 1px;
      mask:
        linear-gradient(#000 0 0) content-box,
        linear-gradient(#000 0 0);
      mask-composite: exclude;
    }
  }

  .custom-select {
    position: relative;
  }

  .select-hidden {
    display: none;
  }

  .select-items {
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
  }

  .select-items div {
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    background: rgba(255, 255, 255, 0.1);
  }

  .select-items div:hover {
    background: rgba(255, 255, 255, 0.2);
  }

  .select-selected svg {
    transition: transform 0.3s;
  }

  .select-selected.select-arrow-active svg {
    transform: rotate(180deg);
  }

  .select-items.show {
    animation: slideDown 0.3s ease-out forwards;
  }

  @keyframes slideDown {
    from {
      opacity: 0;
      transform: translateY(-10px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .accordion_btn {
    font-size: 0.75rem;
    display: flex;
    gap: 0.75rem;
    cursor: pointer;
  }

  .accordion_btn > p {
    text-align: right;
  }

  .accordion_content.active {
    height: auto;
    margin-top: 0.5rem;
  }

  .accordion_content {
    height: 0;
    text-align: justify;
    font-size: 0.875rem;
    overflow: hidden;
    opacity: 0.7;
    transition-property: all;
    transition-timing-function: ease-in-out;
    transition-duration: 500ms;
  }

  .accordion_line {
    height: 1px;
    width: 100%;
    background-color: #5a6a8f;
    border-radius: 99999px;
    margin: 1rem 0;
  }

  @media (width >= 40rem) {
    .accordion_btn {
      font-size: 1.125rem;
    }
    .accordion_content {
      font-size: 1.125rem;
    }
  }

  /*toast classes*/
  #toast-wrapper {
    overflow: hidden;
  }
  #toast-wrapper::before {
    content: "";
    position: absolute;
    z-index: 100;
    opacity: 0.4;
    background-image: url("../static/web/assets/images/new-my/toast-line.png");
    background-position: right top;
    background-repeat: no-repeat;
    width: 100%;
    height: 100%;
  }
  #toast-wrapper[data-type="error"] .toast-title {
    color: #f04438;
  }
  #toast-wrapper[data-type="success"] .toast-title {
    color: #12b76a;
  }
  #toast-wrapper[data-type="warning"] .toast-title {
    color: #ffb800;
  }
  /*toast classes end*/
</style>
    <style>
  .otp-generator-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-flow: row-reverse;
    column-gap: 1rem;
  }

  .otp-generator-input {
    border-radius: 1rem;
    width: 64px;
    height: 64px;
    text-align: center;
    border: 1px solid #85888b;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 1.25rem;
    font-weight: 700;
  }

  @media only screen and (max-width: 640px) {
    .otp-generator-input {
      width: 40px;
      height: 40px;
    }
  }

  .otp-generator-input:focus {
    outline: none;
  }

  .errors-holder {
    min-width: 250px;
    color: #950000;
    font-size: 0.8rem;
  }

  .ltr {
    direction: ltr;
  }

  .rtl {
    direction: rtl;
  }

  div.number-box::before {
    content: "";
    /* override the gradient here */
    background-image: conic-gradient(#4469ad 20deg, transparent 120deg);
    background-color: white;
    box-shadow: inset -20px -20px 20px #0000008c;
    width: 150%;
    height: 150%;
    position: absolute;
    animation: rotate 2s linear infinite;
  }

  div.number-box::after {
    /*content: "";*/
    content: attr(data-after);
    width: 140px;
    height: 140px;
    background: linear-gradient(0deg, #ffffff 0%, #e3e3e3 100%);
    position: absolute;
    border-radius: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #122d46;
    font-size: larger;
    letter-spacing: 5px;
  }

  @keyframes rotate {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(-360deg);
    }
  }

  .custom-select-wrapper {
    flex: 1 1 0%;
    display: flex;
    justify-content: center;
    height: 2.5rem;
    border-bottom: 1px solid rgb(194, 196, 197);
    position: relative;
    text-align: center;
  }

  .custom-select {
    width: 100%;
    text-align: center;
    outline: none;
    display: flex;
    align-items: center;
  }

  .custom-select:after {
    content: "";
    background-image: url('data:image/svg+xml;utf8,<svg width="9" height="5" viewBox="0 0 10 7" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M9.53033 0.96967C9.82322 1.26256 9.82322 1.73744 9.53033 2.03033L5.53033 6.03033C5.23744 6.32322 4.76256 6.32322 4.46967 6.03033L0.46967 2.03033C0.176777 1.73744 0.176777 1.26256 0.46967 0.96967C0.762563 0.676777 1.23744 0.676777 1.53033 0.96967L5 4.43934L8.46967 0.96967C8.76256 0.676777 9.23744 0.676777 9.53033 0.96967Z" fill="%2385888B"/></svg>');
    background-repeat: no-repeat;
    background-size: 14px auto;
    position: absolute;
    left: 10px;
    top: 40%;
    width: 12px;
    height: 12px;
    pointer-events: none;
    transition: all 0.4s;
  }

  .custom-select.open:after {
    rotate: 180deg;
  }

  .select-options {
    width: 100%;
    position: absolute;
    top: calc(100% + 10px);
    left: 0;
    display: none;
    flex-direction: column;
    align-items: center;
    gap: 0.5rem;
    background-color: transparent;
    max-height: 180px;
    overflow: auto;
    z-index: 999;
    border: none;
    border-radius: 16px;
    box-shadow: 2px 2px 4px 1px #0000001a inset;
    backdrop-filter: blur(30px);
    padding: 10px 0;
  }

  .select-options.active {
    display: flex;
  }

  /*.select-option*/

  .select-option {
    width: 100%;
    display: flex;
    justify-content: center;
    gap: 20px;
    text-align: center;
    transition: 0.4s;
    position: relative;
    padding: 5px;
  }

  .select-option::after {
    content: "";
    position: absolute;
    bottom: 0;
    left: 25%;
    width: 50%;
    border-bottom: 1px solid #85888b;
  }

  .select-option:last-of-type::after {
    border-bottom: none;
  }

  .select-option:hover,
  .select-option:focus {
    background: #7c96c6;
    color: white;
  }
  .select-option.selected {
    font-weight: bold;
    background: #7c96c6;
    color: white;
  }

  .custom-select > div:first-child {
    cursor: pointer;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    padding-right: 10px;
  }

  .placeholder:where(.select-trigger-iranian, .select-trigger-foreigner) {
    color: rgb(194, 196, 197);
  }

  #frc-captcha {
    background: #f2f2f2;
    border: 1px solid #ffffff;
    box-shadow: 0 -4px 4px 0 #00000029 inset;
    border-radius: 12px;
  }

  .frc-text {
    font-weight: 500 !important;
    font-size: 1rem !important;
    color: #223557 !important;
  }

  .frc-icon {
    stroke: none !important;
    fill: none !important;
  }

  input:-webkit-autofill,
  input:-webkit-autofill:hover,
  input:-webkit-autofill:focus,
  input:-webkit-autofill:active {
    -webkit-box-shadow: 0 0 0 30px white inset !important;
  }

  .webkit-overflow-scrolling-touch {
    -webkit-overflow-scrolling: touch;
  }
</style>
    <style>
  /*! tailwindcss v4.1.17 | MIT License | https://tailwindcss.com */
  @layer properties;
  @layer theme, base, components, utilities;
  @layer theme {
    :root,
    :host {
      --font-sans:
        ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji",
        "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
      --font-mono:
        ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
        "Liberation Mono", "Courier New", monospace;
      --color-gray-400: oklch(70.7% 0.022 261.325);
      --color-neutral-50: oklch(98.5% 0 0);
      --color-neutral-500: oklch(55.6% 0 0);
      --color-neutral-900: oklch(20.5% 0 0);
      --color-black: #000;
      --color-white: #fff;
      --spacing: 0.25rem;
      --text-xs: 0.75rem;
      --text-xs--line-height: calc(1 / 0.75);
      --text-sm: 0.875rem;
      --text-sm--line-height: calc(1.25 / 0.875);
      --text-base: 1rem;
      --text-base--line-height: calc(1.5 / 1);
      --text-lg: 1.125rem;
      --text-lg--line-height: calc(1.75 / 1.125);
      --text-xl: 1.25rem;
      --text-xl--line-height: calc(1.75 / 1.25);
      --text-2xl: 1.5rem;
      --text-2xl--line-height: calc(2 / 1.5);
      --font-weight-normal: 400;
      --font-weight-medium: 500;
      --font-weight-semibold: 600;
      --font-weight-bold: 700;
      --font-weight-extrabold: 800;
      --radius-md: 0.375rem;
      --radius-lg: 0.5rem;
      --radius-2xl: 1rem;
      --radius-3xl: 1.5rem;
      --ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);
      --animate-spin: spin 1s linear infinite;
      --default-transition-duration: 150ms;
      --default-transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
      --default-font-family: var(--font-sans);
      --default-mono-font-family: var(--font-mono);
    }
  }
  @layer base {
    *,
    ::after,
    ::before,
    ::backdrop,
    ::file-selector-button {
      box-sizing: border-box;
      margin: 0;
      padding: 0;
      border: 0 solid;
    }
    html,
    :host {
      line-height: 1.5;
      -webkit-text-size-adjust: 100%;
      tab-size: 4;
      font-family: var(
        --default-font-family,
        ui-sans-serif,
        system-ui,
        sans-serif,
        "Apple Color Emoji",
        "Segoe UI Emoji",
        "Segoe UI Symbol",
        "Noto Color Emoji"
      );
      font-feature-settings: var(--default-font-feature-settings, normal);
      font-variation-settings: var(--default-font-variation-settings, normal);
      -webkit-tap-highlight-color: transparent;
    }
    hr {
      height: 0;
      color: inherit;
      border-top-width: 1px;
    }
    abbr:where([title]) {
      -webkit-text-decoration: underline dotted;
      text-decoration: underline dotted;
    }
    h1,
    h2,
    h3,
    h4,
    h5,
    h6 {
      font-size: inherit;
      font-weight: inherit;
    }
    a {
      color: inherit;
      -webkit-text-decoration: inherit;
      text-decoration: inherit;
    }
    b,
    strong {
      font-weight: bolder;
    }
    code,
    kbd,
    samp,
    pre {
      font-family: var(
        --default-mono-font-family,
        ui-monospace,
        SFMono-Regular,
        Menlo,
        Monaco,
        Consolas,
        "Liberation Mono",
        "Courier New",
        monospace
      );
      font-feature-settings: var(--default-mono-font-feature-settings, normal);
      font-variation-settings: var(
        --default-mono-font-variation-settings,
        normal
      );
      font-size: 1em;
    }
    small {
      font-size: 80%;
    }
    sub,
    sup {
      font-size: 75%;
      line-height: 0;
      position: relative;
      vertical-align: baseline;
    }
    sub {
      bottom: -0.25em;
    }
    sup {
      top: -0.5em;
    }
    table {
      text-indent: 0;
      border-color: inherit;
      border-collapse: collapse;
    }
    :-moz-focusring {
      outline: auto;
    }
    progress {
      vertical-align: baseline;
    }
    summary {
      display: list-item;
    }
    ol,
    ul,
    menu {
      list-style: none;
    }
    img,
    svg,
    video,
    canvas,
    audio,
    iframe,
    embed,
    object {
      display: block;
      vertical-align: middle;
    }
    img,
    video {
      max-width: 100%;
      height: auto;
    }
    button,
    input,
    select,
    optgroup,
    textarea,
    ::file-selector-button {
      font: inherit;
      font-feature-settings: inherit;
      font-variation-settings: inherit;
      letter-spacing: inherit;
      color: inherit;
      border-radius: 0;
      background-color: transparent;
      opacity: 1;
    }
    :where(select:is([multiple], [size])) optgroup {
      font-weight: bolder;
    }
    :where(select:is([multiple], [size])) optgroup option {
      padding-inline-start: 20px;
    }
    ::file-selector-button {
      margin-inline-end: 4px;
    }
    ::placeholder {
      opacity: 1;
    }
    @supports (not (-webkit-appearance: -apple-pay-button)) or
      (contain-intrinsic-size: 1px) {
      ::placeholder {
        color: currentcolor;
        @supports (color: color-mix(in lab, red, red)) {
          color: color-mix(in oklab, currentcolor 50%, transparent);
        }
      }
    }
    textarea {
      resize: vertical;
    }
    ::-webkit-search-decoration {
      -webkit-appearance: none;
    }
    ::-webkit-date-and-time-value {
      min-height: 1lh;
      text-align: inherit;
    }
    ::-webkit-datetime-edit {
      display: inline-flex;
    }
    ::-webkit-datetime-edit-fields-wrapper {
      padding: 0;
    }
    ::-webkit-datetime-edit,
    ::-webkit-datetime-edit-year-field,
    ::-webkit-datetime-edit-month-field,
    ::-webkit-datetime-edit-day-field,
    ::-webkit-datetime-edit-hour-field,
    ::-webkit-datetime-edit-minute-field,
    ::-webkit-datetime-edit-second-field,
    ::-webkit-datetime-edit-millisecond-field,
    ::-webkit-datetime-edit-meridiem-field {
      padding-block: 0;
    }
    ::-webkit-calendar-picker-indicator {
      line-height: 1;
    }
    :-moz-ui-invalid {
      box-shadow: none;
    }
    button,
    input:where([type="button"], [type="reset"], [type="submit"]),
    ::file-selector-button {
      appearance: button;
    }
    ::-webkit-inner-spin-button,
    ::-webkit-outer-spin-button {
      height: auto;
    }
    [hidden]:where(:not([hidden="until-found"])) {
      display: none !important;
    }
  }
  @layer utilities {
    .visible {
      visibility: visible;
    }
    .absolute {
      position: absolute;
    }
    .fixed {
      position: fixed;
    }
    .relative {
      position: relative;
    }
    .top-0 {
      top: calc(var(--spacing) * 0);
    }
    .top-1 {
      top: calc(var(--spacing) * 1);
    }
    .top-1\/2 {
      top: calc(1 / 2 * 100%);
    }
    .top-4 {
      top: calc(var(--spacing) * 4);
    }
    .top-10 {
      top: calc(var(--spacing) * 10);
    }
    .top-12 {
      top: calc(var(--spacing) * 12);
    }
    .top-\[-27\%\] {
      top: -27%;
    }
    .top-\[-42\%\] {
      top: -42%;
    }
    .top-full {
      top: 100%;
    }
    .-right-2\/3 {
      right: calc(calc(2 / 3 * 100%) * -1);
    }
    .-right-\[85vw\] {
      right: calc(85vw * -1);
    }
    .right-0 {
      right: calc(var(--spacing) * 0);
    }
    .right-4 {
      right: calc(var(--spacing) * 4);
    }
    .right-\[-2px\] {
      right: -2px;
    }
    .right-\[-290px\] {
      right: -290px;
    }
    .right-\[-379px\] {
      right: -379px;
    }
    .right-\[15px\] {
      right: 15px;
    }
    .-bottom-1 {
      bottom: calc(var(--spacing) * -1);
    }
    .-bottom-1\.5 {
      bottom: calc(var(--spacing) * -1.5);
    }
    .-bottom-1\/5 {
      bottom: calc(calc(1 / 5 * 100%) * -1);
    }
    .bottom-0 {
      bottom: calc(var(--spacing) * 0);
    }
    .bottom-\[-255px\] {
      bottom: -255px;
    }
    .bottom-\[-258px\] {
      bottom: -258px;
    }
    .bottom-full {
      bottom: 100%;
    }
    .-left-1\/2 {
      left: calc(calc(1 / 2 * 100%) * -1);
    }
    .-left-2\/3 {
      left: calc(calc(2 / 3 * 100%) * -1);
    }
    .-left-3\.5 {
      left: calc(var(--spacing) * -3.5);
    }
    .left-0 {
      left: calc(var(--spacing) * 0);
    }
    .left-1\/2 {
      left: calc(1 / 2 * 100%);
    }
    .left-\[-293px\] {
      left: -293px;
    }
    .left-\[-337px\] {
      left: -337px;
    }
    .left-\[70\%\] {
      left: 70%;
    }
    .z-0 {
      z-index: 0;
    }
    .z-10 {
      z-index: 10;
    }
    .z-50 {
      z-index: 50;
    }
    .z-\[-1\] {
      z-index: -1;
    }
    .z-\[-10\] {
      z-index: -10;
    }
    .z-\[1\] {
      z-index: 1;
    }
    .order-1 {
      order: 1;
    }
    .order-2 {
      order: 2;
    }
    .col-span-12 {
      grid-column: span 12 / span 12;
    }
    .float-left {
      float: left;
    }
    .container {
      width: 100%;
      @media (width >= 40rem) {
        max-width: 40rem;
      }
      @media (width >= 48rem) {
        max-width: 48rem;
      }
      @media (width >= 64rem) {
        max-width: 64rem;
      }
      @media (width >= 80rem) {
        max-width: 80rem;
      }
      @media (width >= 96rem) {
        max-width: 96rem;
      }
    }
    .m-auto {
      margin: auto;
    }
    .mx-2 {
      margin-inline: calc(var(--spacing) * 2);
    }
    .mx-6 {
      margin-inline: calc(var(--spacing) * 6);
    }
    .mx-10 {
      margin-inline: calc(var(--spacing) * 10);
    }
    .mx-20 {
      margin-inline: calc(var(--spacing) * 20);
    }
    .mx-auto {
      margin-inline: auto;
    }
    .my-4 {
      margin-block: calc(var(--spacing) * 4);
    }
    .my-8 {
      margin-block: calc(var(--spacing) * 8);
    }
    .mt-1 {
      margin-top: calc(var(--spacing) * 1);
    }
    .mt-2 {
      margin-top: calc(var(--spacing) * 2);
    }
    .mt-2\.5 {
      margin-top: calc(var(--spacing) * 2.5);
    }
    .mt-3 {
      margin-top: calc(var(--spacing) * 3);
    }
    .mt-4 {
      margin-top: calc(var(--spacing) * 4);
    }
    .mt-5 {
      margin-top: calc(var(--spacing) * 5);
    }
    .mt-6 {
      margin-top: calc(var(--spacing) * 6);
    }
    .mt-8 {
      margin-top: calc(var(--spacing) * 8);
    }
    .mt-9 {
      margin-top: calc(var(--spacing) * 9);
    }
    .mt-10 {
      margin-top: calc(var(--spacing) * 10);
    }
    .mt-11 {
      margin-top: calc(var(--spacing) * 11);
    }
    .mt-16 {
      margin-top: calc(var(--spacing) * 16);
    }
    .mt-32 {
      margin-top: calc(var(--spacing) * 32);
    }
    .mt-\[17px\] {
      margin-top: 17px;
    }
    .mt-\[40px\] {
      margin-top: 40px;
    }
    .mt-\[58px\] {
      margin-top: 58px;
    }
    .mt-\[60px\] {
      margin-top: 60px;
    }
    .mt-\[64px\] {
      margin-top: 64px;
    }
    .mt-\[66px\] {
      margin-top: 66px;
    }
    .mt-\[72px\] {
      margin-top: 72px;
    }
    .mt-\[78px\] {
      margin-top: 78px;
    }
    .mt-\[80px\] {
      margin-top: 80px;
    }
    .mt-\[94px\] {
      margin-top: 94px;
    }
    .mt-\[106px\] {
      margin-top: 106px;
    }
    .mr-4 {
      margin-right: calc(var(--spacing) * 4);
    }
    .mr-\[20px\] {
      margin-right: 20px;
    }
    .-mb-1 {
      margin-bottom: calc(var(--spacing) * -1);
    }
    .-mb-4 {
      margin-bottom: calc(var(--spacing) * -4);
    }
    .mb-2 {
      margin-bottom: calc(var(--spacing) * 2);
    }
    .mb-2\.5 {
      margin-bottom: calc(var(--spacing) * 2.5);
    }
    .mb-4 {
      margin-bottom: calc(var(--spacing) * 4);
    }
    .mb-6 {
      margin-bottom: calc(var(--spacing) * 6);
    }
    .mb-8 {
      margin-bottom: calc(var(--spacing) * 8);
    }
    .mb-23 {
      margin-bottom: calc(var(--spacing) * 23);
    }
    .mb-\[23px\] {
      margin-bottom: 23px;
    }
    .mb-\[44px\] {
      margin-bottom: 44px;
    }
    .mb-\[46px\] {
      margin-bottom: 46px;
    }
    .mb-\[80px\] {
      margin-bottom: 80px;
    }
    .mb-\[120px\] {
      margin-bottom: 120px;
    }
    .mb-\[174px\] {
      margin-bottom: 174px;
    }
    .mb-\[176px\] {
      margin-bottom: 176px;
    }
    .block {
      display: block;
    }
    .flex {
      display: flex;
    }
    .grid {
      display: grid;
    }
    .hidden {
      display: none;
    }
    .inline-block {
      display: inline-block;
    }
    .inline-flex {
      display: inline-flex;
    }
    .aspect-square {
      aspect-ratio: 1 / 1;
    }
    .h-0\.5 {
      height: calc(var(--spacing) * 0.5);
    }
    .h-4 {
      height: calc(var(--spacing) * 4);
    }
    .h-4\.5 {
      height: calc(var(--spacing) * 4.5);
    }
    .h-5 {
      height: calc(var(--spacing) * 5);
    }
    .h-6 {
      height: calc(var(--spacing) * 6);
    }
    .h-7 {
      height: calc(var(--spacing) * 7);
    }
    .h-\[1px\] {
      height: 1px;
    }
    .h-\[3px\] {
      height: 3px;
    }
    .h-\[20px\] {
      height: 20px;
    }
    .h-\[30px\] {
      height: 30px;
    }
    .h-\[34px\] {
      height: 34px;
    }
    .h-\[40px\] {
      height: 40px;
    }
    .h-\[50px\] {
      height: 50px;
    }
    .h-\[64px\] {
      height: 64px;
    }
    .h-\[80px\] {
      height: 80px;
    }
    .h-\[127px\] {
      height: 127px;
    }
    .h-\[150px\] {
      height: 150px;
    }
    .h-full {
      height: 100%;
    }
    .h-px {
      height: 1px;
    }
    .h-screen {
      height: 100vh;
    }
    .w-1\/2 {
      width: calc(1 / 2 * 100%);
    }
    .w-5 {
      width: calc(var(--spacing) * 5);
    }
    .w-6 {
      width: calc(var(--spacing) * 6);
    }
    .w-7 {
      width: calc(var(--spacing) * 7);
    }
    .w-100 {
      width: calc(var(--spacing) * 100);
    }
    .w-\[2px\] {
      width: 2px;
    }
    .w-\[4px\] {
      width: 4px;
    }
    .w-\[12vw\] {
      width: 12vw;
    }
    .w-\[20px\] {
      width: 20px;
    }
    .w-\[30px\] {
      width: 30px;
    }
    .w-\[34px\] {
      width: 34px;
    }
    .w-\[41px\] {
      width: 41px;
    }
    .w-\[50px\] {
      width: 50px;
    }
    .w-\[64px\] {
      width: 64px;
    }
    .w-\[85\%\] {
      width: 85%;
    }
    .w-\[90vw\] {
      width: 90vw;
    }
    .w-\[120px\] {
      width: 120px;
    }
    .w-\[150px\] {
      width: 150px;
    }
    .w-\[185\%\] {
      width: 185%;
    }
    .w-fit {
      width: fit-content;
    }
    .w-full {
      width: 100%;
    }
    .w-max {
      width: max-content;
    }
    .w-px {
      width: 1px;
    }
    .\!max-w-none {
      max-width: none !important;
    }
    .max-w-\[85\%\] {
      max-width: 85%;
    }
    .max-w-\[168px\] {
      max-width: 168px;
    }
    .max-w-\[170px\] {
      max-width: 170px;
    }
    .max-w-\[450px\] {
      max-width: 450px;
    }
    .max-w-\[620px\] {
      max-width: 620px;
    }
    .max-w-\[660px\] {
      max-width: 660px;
    }
    .min-w-\[60\%\] {
      min-width: 60%;
    }
    .min-w-\[110px\] {
      min-width: 110px;
    }
    .min-w-\[490px\] {
      min-width: 490px;
    }
    .flex-1 {
      flex: 1;
    }
    .flex-grow {
      flex-grow: 1;
    }
    .-translate-x-1\/2 {
      --tw-translate-x: calc(calc(1 / 2 * 100%) * -1);
      translate: var(--tw-translate-x) var(--tw-translate-y);
    }
    .-translate-y-1\/2 {
      --tw-translate-y: calc(calc(1 / 2 * 100%) * -1);
      translate: var(--tw-translate-x) var(--tw-translate-y);
    }
    .scale-\[1\.1\] {
      scale: 1.1;
    }
    .transform {
      transform: var(--tw-rotate-x,) var(--tw-rotate-y,) var(--tw-rotate-z,)
        var(--tw-skew-x,) var(--tw-skew-y,);
    }
    .animate-spin {
      animation: var(--animate-spin);
    }
    .cursor-pointer {
      cursor: pointer;
    }
    .list-disc {
      list-style-type: disc;
    }
    .grid-cols-3 {
      grid-template-columns: repeat(3, minmax(0, 1fr));
    }
    .grid-cols-12 {
      grid-template-columns: repeat(12, minmax(0, 1fr));
    }
    .flex-col {
      flex-direction: column;
    }
    .flex-row {
      flex-direction: row;
    }
    .flex-row-reverse {
      flex-direction: row-reverse;
    }
    .content-center {
      align-content: center;
    }
    .items-center {
      align-items: center;
    }
    .items-start {
      align-items: flex-start;
    }
    .justify-between {
      justify-content: space-between;
    }
    .justify-center {
      justify-content: center;
    }
    .justify-start {
      justify-content: flex-start;
    }
    .gap-1 {
      gap: calc(var(--spacing) * 1);
    }
    .gap-2 {
      gap: calc(var(--spacing) * 2);
    }
    .gap-3 {
      gap: calc(var(--spacing) * 3);
    }
    .gap-4 {
      gap: calc(var(--spacing) * 4);
    }
    .gap-5 {
      gap: calc(var(--spacing) * 5);
    }
    .gap-6 {
      gap: calc(var(--spacing) * 6);
    }
    .gap-8 {
      gap: calc(var(--spacing) * 8);
    }
    .gap-12 {
      gap: calc(var(--spacing) * 12);
    }
    .gap-\[60px\] {
      gap: 60px;
    }
    .space-y-1 {
      :where(& > :not(:last-child)) {
        --tw-space-y-reverse: 0;
        margin-block-start: calc(
          calc(var(--spacing) * 1) * var(--tw-space-y-reverse)
        );
        margin-block-end: calc(
          calc(var(--spacing) * 1) * calc(1 - var(--tw-space-y-reverse))
        );
      }
    }
    .space-y-6 {
      :where(& > :not(:last-child)) {
        --tw-space-y-reverse: 0;
        margin-block-start: calc(
          calc(var(--spacing) * 6) * var(--tw-space-y-reverse)
        );
        margin-block-end: calc(
          calc(var(--spacing) * 6) * calc(1 - var(--tw-space-y-reverse))
        );
      }
    }
    .overflow-hidden {
      overflow: hidden;
    }
    .overflow-visible {
      overflow: visible;
    }
    .overflow-x-hidden {
      overflow-x: hidden;
    }
    .overflow-y-auto {
      overflow-y: auto;
    }
    .overscroll-x-contain {
      overscroll-behavior-x: contain;
    }
    .overscroll-y-contain {
      overscroll-behavior-y: contain;
    }
    .rounded {
      border-radius: 0.25rem;
    }
    .rounded-2xl {
      border-radius: var(--radius-2xl);
    }
    .rounded-3xl {
      border-radius: var(--radius-3xl);
    }
    .rounded-\[10px\] {
      border-radius: 10px;
    }
    .rounded-full {
      border-radius: calc(infinity * 1px);
    }
    .rounded-lg {
      border-radius: var(--radius-lg);
    }
    .rounded-md {
      border-radius: var(--radius-md);
    }
    .rounded-t-full {
      border-top-left-radius: calc(infinity * 1px);
      border-top-right-radius: calc(infinity * 1px);
    }
    .border {
      border-style: var(--tw-border-style);
      border-width: 1px;
    }
    .border-0 {
      border-style: var(--tw-border-style);
      border-width: 0px;
    }
    .border-2 {
      border-style: var(--tw-border-style);
      border-width: 2px;
    }
    .border-\[3px\] {
      border-style: var(--tw-border-style);
      border-width: 3px;
    }
    .border-b {
      border-bottom-style: var(--tw-border-style);
      border-bottom-width: 1px;
    }
    .border-none {
      --tw-border-style: none;
      border-style: none;
    }
    .border-\[\#7C96C6\] {
      border-color: #7c96c6;
    }
    .border-\[\#8092C1\] {
      border-color: #8092c1;
    }
    .border-\[\#85888B\] {
      border-color: #85888b;
    }
    .border-\[\#F5F5F5\] {
      border-color: #f5f5f5;
    }
    .border-\[\#c4c4c4\] {
      border-color: #c4c4c4;
    }
    .border-\[\#e3e3e3\] {
      border-color: #e3e3e3;
    }
    .border-white {
      border-color: var(--color-white);
    }
    .bg-\[\#3F619D\] {
      background-color: #3f619d;
    }
    .bg-\[\#304A79\] {
      background-color: #304a79;
    }
    .bg-\[\#4469AD\] {
      background-color: #4469ad;
    }
    .bg-\[\#A1B4D680\] {
      background-color: #a1b4d680;
    }
    .bg-\[\#C0DEFE\] {
      background-color: #c0defe;
    }
    .bg-\[\#DADBDC\] {
      background-color: #dadbdc;
    }
    .bg-\[\#ECF0F7\] {
      background-color: #ecf0f7;
    }
    .bg-\[\#ECF4FF1A\] {
      background-color: #ecf4ff1a;
    }
    .bg-\[\#F4F6FB\] {
      background-color: #f4f6fb;
    }
    .bg-\[\#bababa\] {
      background-color: #bababa;
    }
    .bg-gray-400 {
      background-color: var(--color-gray-400);
    }
    .bg-transparent {
      background-color: transparent;
    }
    .bg-white {
      background-color: var(--color-white);
    }
    .bg-white\/20 {
      background-color: color-mix(in srgb, #fff 20%, transparent);
      @supports (color: color-mix(in lab, red, red)) {
        background-color: color-mix(
          in oklab,
          var(--color-white) 20%,
          transparent
        );
      }
    }
    .bg-gradient-to-l {
      --tw-gradient-position: to left in oklab;
      background-image: linear-gradient(var(--tw-gradient-stops));
    }
    .bg-gradient-to-r {
      --tw-gradient-position: to right in oklab;
      background-image: linear-gradient(var(--tw-gradient-stops));
    }
    .bg-gradient-to-t {
      --tw-gradient-position: to top in oklab;
      background-image: linear-gradient(var(--tw-gradient-stops));
    }
    .bg-\[linear-gradient\(201\.67deg\,\#304A79_34\.36\%\,\#4469AD_95\.69\%\)\] {
      background-image: linear-gradient(
        201.67deg,
        #304a79 34.36%,
        #4469ad 95.69%
      );
    }
    .bg-\[linear-gradient\(262\.38deg\,rgba\(68\,105\,173\,0\.35\)52\.04\%\,rgba\(48\,74\,121\,0\)93\.87\%\)\] {
      background-image: linear-gradient(
        262.38deg,
        rgba(68, 105, 173, 0.35) 52.04%,
        rgba(48, 74, 121, 0) 93.87%
      );
    }
    .bg-\[linear-gradient\(360deg\,\#304A79_8\.85\%\,rgba\(68\,105\,173\,0\.35\)100\%\)\] {
      background-image: linear-gradient(
        360deg,
        #304a79 8.85%,
        rgba(68, 105, 173, 0.35) 100%
      );
    }
    .from-\[\#C2C4C5\] {
      --tw-gradient-from: #c2c4c5;
      --tw-gradient-stops: var(
        --tw-gradient-via-stops,
        var(--tw-gradient-position),
        var(--tw-gradient-from) var(--tw-gradient-from-position),
        var(--tw-gradient-to) var(--tw-gradient-to-position)
      );
    }
    .from-white {
      --tw-gradient-from: var(--color-white);
      --tw-gradient-stops: var(
        --tw-gradient-via-stops,
        var(--tw-gradient-position),
        var(--tw-gradient-from) var(--tw-gradient-from-position),
        var(--tw-gradient-to) var(--tw-gradient-to-position)
      );
    }
    .to-\[\#e3e3e3\] {
      --tw-gradient-to: #e3e3e3;
      --tw-gradient-stops: var(
        --tw-gradient-via-stops,
        var(--tw-gradient-position),
        var(--tw-gradient-from) var(--tw-gradient-from-position),
        var(--tw-gradient-to) var(--tw-gradient-to-position)
      );
    }
    .to-white {
      --tw-gradient-to: var(--color-white);
      --tw-gradient-stops: var(
        --tw-gradient-via-stops,
        var(--tw-gradient-position),
        var(--tw-gradient-from) var(--tw-gradient-from-position),
        var(--tw-gradient-to) var(--tw-gradient-to-position)
      );
    }
    .bg-center {
      background-position: center;
    }
    .bg-no-repeat {
      background-repeat: no-repeat;
    }
    .fill-none {
      fill: none;
    }
    .fill-white {
      fill: var(--color-white);
    }
    .stroke-\[\#7C96C6\] {
      stroke: #7c96c6;
    }
    .stroke-\[\#9CBDFF\] {
      stroke: #9cbdff;
    }
    .stroke-\[\#9cbdff\] {
      stroke: #9cbdff;
    }
    .stroke-\[\#304A79\] {
      stroke: #304a79;
    }
    .stroke-\[\#85888B\] {
      stroke: #85888b;
    }
    .stroke-\[\#AAACAE\] {
      stroke: #aaacae;
    }
    .stroke-2 {
      stroke-width: 2;
    }
    .p-2 {
      padding: calc(var(--spacing) * 2);
    }
    .p-2\.5 {
      padding: calc(var(--spacing) * 2.5);
    }
    .p-4 {
      padding: calc(var(--spacing) * 4);
    }
    .p-5 {
      padding: calc(var(--spacing) * 5);
    }
    .px-2 {
      padding-inline: calc(var(--spacing) * 2);
    }
    .px-3 {
      padding-inline: calc(var(--spacing) * 3);
    }
    .px-4 {
      padding-inline: calc(var(--spacing) * 4);
    }
    .px-6 {
      padding-inline: calc(var(--spacing) * 6);
    }
    .px-8 {
      padding-inline: calc(var(--spacing) * 8);
    }
    .px-10 {
      padding-inline: calc(var(--spacing) * 10);
    }
    .px-\[17px\] {
      padding-inline: 17px;
    }
    .px-\[72px\] {
      padding-inline: 72px;
    }
    .px-\[81px\] {
      padding-inline: 81px;
    }
    .px-\[92px\] {
      padding-inline: 92px;
    }
    .py-1 {
      padding-block: calc(var(--spacing) * 1);
    }
    .py-1\.5 {
      padding-block: calc(var(--spacing) * 1.5);
    }
    .py-2 {
      padding-block: calc(var(--spacing) * 2);
    }
    .py-4 {
      padding-block: calc(var(--spacing) * 4);
    }
    .py-12 {
      padding-block: calc(var(--spacing) * 12);
    }
    .pt-2 {
      padding-top: calc(var(--spacing) * 2);
    }
    .pt-12 {
      padding-top: calc(var(--spacing) * 12);
    }
    .pt-17 {
      padding-top: calc(var(--spacing) * 17);
    }
    .pt-\[65px\] {
      padding-top: 65px;
    }
    .pr-2 {
      padding-right: calc(var(--spacing) * 2);
    }
    .pr-4 {
      padding-right: calc(var(--spacing) * 4);
    }
    .pr-4\.5 {
      padding-right: calc(var(--spacing) * 4.5);
    }
    .pr-5 {
      padding-right: calc(var(--spacing) * 5);
    }
    .\!pb-0 {
      padding-bottom: calc(var(--spacing) * 0) !important;
    }
    .pb-2 {
      padding-bottom: calc(var(--spacing) * 2);
    }
    .pb-3 {
      padding-bottom: calc(var(--spacing) * 3);
    }
    .pb-4 {
      padding-bottom: calc(var(--spacing) * 4);
    }
    .pb-10 {
      padding-bottom: calc(var(--spacing) * 10);
    }
    .pb-11 {
      padding-bottom: calc(var(--spacing) * 11);
    }
    .pb-20 {
      padding-bottom: calc(var(--spacing) * 20);
    }
    .pl-10 {
      padding-left: calc(var(--spacing) * 10);
    }
    .text-center {
      text-align: center;
    }
    .text-justify {
      text-align: justify;
    }
    .text-right {
      text-align: right;
    }
    .text-base {
      font-size: var(--text-base);
      line-height: var(--tw-leading, var(--text-base--line-height));
    }
    .text-lg {
      font-size: var(--text-lg);
      line-height: var(--tw-leading, var(--text-lg--line-height));
    }
    .text-sm {
      font-size: var(--text-sm);
      line-height: var(--tw-leading, var(--text-sm--line-height));
    }
    .text-xl {
      font-size: var(--text-xl);
      line-height: var(--tw-leading, var(--text-xl--line-height));
    }
    .text-xs {
      font-size: var(--text-xs);
      line-height: var(--tw-leading, var(--text-xs--line-height));
    }
    .text-\[12px\] {
      font-size: 12px;
    }
    .text-\[16px\] {
      font-size: 16px;
    }
    .text-\[18px\] {
      font-size: 18px;
    }
    .text-\[20px\] {
      font-size: 20px;
    }
    .text-\[23px\] {
      font-size: 23px;
    }
    .text-\[24px\] {
      font-size: 24px;
    }
    .text-\[32px\] {
      font-size: 32px;
    }
    .text-\[40px\] {
      font-size: 40px;
    }
    .font-bold {
      --tw-font-weight: var(--font-weight-bold);
      font-weight: var(--font-weight-bold);
    }
    .font-extrabold {
      --tw-font-weight: var(--font-weight-extrabold);
      font-weight: var(--font-weight-extrabold);
    }
    .font-medium {
      --tw-font-weight: var(--font-weight-medium);
      font-weight: var(--font-weight-medium);
    }
    .font-normal {
      --tw-font-weight: var(--font-weight-normal);
      font-weight: var(--font-weight-normal);
    }
    .font-semibold {
      --tw-font-weight: var(--font-weight-semibold);
      font-weight: var(--font-weight-semibold);
    }
    .text-nowrap {
      text-wrap: nowrap;
    }
    .text-\[\#2E90FA\] {
      color: #2e90fa;
    }
    .text-\[\#3D5F9C\] {
      color: #3d5f9c;
    }
    .text-\[\#3E96F5\] {
      color: #3e96f5;
    }
    .text-\[\#5D5F61\] {
      color: #5d5f61;
    }
    .text-\[\#7C96C6\] {
      color: #7c96c6;
    }
    .text-\[\#9CBDFF\] {
      color: #9cbdff;
    }
    .text-\[\#304A79\] {
      color: #304a79;
    }
    .text-\[\#787A7D\] {
      color: #787a7d;
    }
    .text-\[\#4469AD\] {
      color: #4469ad;
    }
    .text-\[\#85888B\] {
      color: #85888b;
    }
    .text-\[\#424446\] {
      color: #424446;
    }
    .text-\[\#797979\] {
      color: #797979;
    }
    .text-\[\#AAACAE\] {
      color: #aaacae;
    }
    .text-\[\#DDEAFF\] {
      color: #ddeaff;
    }
    .text-\[\#ECF4FF\] {
      color: #ecf4ff;
    }
    .text-black {
      color: var(--color-black);
    }
    .text-neutral-50 {
      color: var(--color-neutral-50);
    }
    .text-neutral-500 {
      color: var(--color-neutral-500);
    }
    .text-neutral-900 {
      color: var(--color-neutral-900);
    }
    .text-white {
      color: var(--color-white);
    }
    .no-underline {
      text-decoration-line: none;
    }
    .underline {
      text-decoration-line: underline;
    }
    .opacity-0 {
      opacity: 0%;
    }
    .opacity-30 {
      opacity: 30%;
    }
    .opacity-50 {
      opacity: 50%;
    }
    .opacity-60 {
      opacity: 60%;
    }
    .opacity-70 {
      opacity: 70%;
    }
    .opacity-100 {
      opacity: 100%;
    }
    .shadow-\[-2px_-2px_16px_0px_\#00000014\] {
      --tw-shadow: -2px -2px 16px 0px var(--tw-shadow-color, #00000014);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[-2px_-2px_16px_0px_rgba\(0\,0\,0\,0\.08\)\] {
      --tw-shadow: -2px -2px 16px 0px
        var(--tw-shadow-color, rgba(0, 0, 0, 0.08));
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[-4px_4px_10px_0px_\#00000080\] {
      --tw-shadow: -4px 4px 10px 0px var(--tw-shadow-color, #00000080);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[0px_3px_7px_0px_\#00000040\] {
      --tw-shadow: 0px 3px 7px 0px var(--tw-shadow-color, #00000040);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[1px_2px_2px\] {
      --tw-shadow: 1px 2px 2px var(--tw-shadow-color, currentcolor);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[1px_4px_2px\] {
      --tw-shadow: 1px 4px 2px var(--tw-shadow-color, currentcolor);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[4px_4px_8px_0px_\#00000040\] {
      --tw-shadow: 4px 4px 8px 0px var(--tw-shadow-color, #00000040);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[4px_4px_10px_0px_rgba\(0\,0\,0\,0\.1\)\] {
      --tw-shadow: 4px 4px 10px 0px var(--tw-shadow-color, rgba(0, 0, 0, 0.1));
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-lg {
      --tw-shadow:
        0 10px 15px -3px var(--tw-shadow-color, rgb(0 0 0 / 0.1)),
        0 4px 6px -4px var(--tw-shadow-color, rgb(0 0 0 / 0.1));
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .inset-shadow-\[0_0_32px_0px_\#00000090\] {
      --tw-inset-shadow: inset 0 0 32px 0px
        var(--tw-inset-shadow-color, #00000090);
      box-shadow:
        var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
        var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
    }
    .shadow-\[\#3F619D\] {
      --tw-shadow-color: #3f619d;
      @supports (color: color-mix(in lab, red, red)) {
        --tw-shadow-color: color-mix(
          in oklab,
          #3f619d var(--tw-shadow-alpha),
          transparent
        );
      }
    }
    .shadow-\[\#85888B\] {
      --tw-shadow-color: #85888b;
      @supports (color: color-mix(in lab, red, red)) {
        --tw-shadow-color: color-mix(
          in oklab,
          #85888b var(--tw-shadow-alpha),
          transparent
        );
      }
    }
    .drop-shadow-\[-10px_-5px_32px_0px_\#00000066\] {
      --tw-drop-shadow-size: drop-shadow(
        -10px -5px 32px 0px var(--tw-drop-shadow-color, #00000066)
      );
      --tw-drop-shadow: var(--tw-drop-shadow-size);
      filter: var(--tw-blur,) var(--tw-brightness,) var(--tw-contrast,)
        var(--tw-grayscale,) var(--tw-hue-rotate,) var(--tw-invert,)
        var(--tw-saturate,) var(--tw-sepia,) var(--tw-drop-shadow,);
    }
    .filter {
      filter: var(--tw-blur,) var(--tw-brightness,) var(--tw-contrast,)
        var(--tw-grayscale,) var(--tw-hue-rotate,) var(--tw-invert,)
        var(--tw-saturate,) var(--tw-sepia,) var(--tw-drop-shadow,);
    }
    .backdrop-filter {
      -webkit-backdrop-filter: var(--tw-backdrop-blur,)
        var(--tw-backdrop-brightness,) var(--tw-backdrop-contrast,)
        var(--tw-backdrop-grayscale,) var(--tw-backdrop-hue-rotate,)
        var(--tw-backdrop-invert,) var(--tw-backdrop-opacity,)
        var(--tw-backdrop-saturate,) var(--tw-backdrop-sepia,);
      backdrop-filter: var(--tw-backdrop-blur,) var(--tw-backdrop-brightness,)
        var(--tw-backdrop-contrast,) var(--tw-backdrop-grayscale,)
        var(--tw-backdrop-hue-rotate,) var(--tw-backdrop-invert,)
        var(--tw-backdrop-opacity,) var(--tw-backdrop-saturate,)
        var(--tw-backdrop-sepia,);
    }
    .transition-all {
      transition-property: all;
      transition-timing-function: var(
        --tw-ease,
        var(--default-transition-timing-function)
      );
      transition-duration: var(
        --tw-duration,
        var(--default-transition-duration)
      );
    }
    .transition-colors {
      transition-property:
        color, background-color, border-color, outline-color,
        text-decoration-color, fill, stroke, --tw-gradient-from,
        --tw-gradient-via, --tw-gradient-to;
      transition-timing-function: var(
        --tw-ease,
        var(--default-transition-timing-function)
      );
      transition-duration: var(
        --tw-duration,
        var(--default-transition-duration)
      );
    }
    .transition-opacity {
      transition-property: opacity;
      transition-timing-function: var(
        --tw-ease,
        var(--default-transition-timing-function)
      );
      transition-duration: var(
        --tw-duration,
        var(--default-transition-duration)
      );
    }
    .duration-200 {
      --tw-duration: 200ms;
      transition-duration: 200ms;
    }
    .duration-500 {
      --tw-duration: 500ms;
      transition-duration: 500ms;
    }
    .ease-in-out {
      --tw-ease: var(--ease-in-out);
      transition-timing-function: var(--ease-in-out);
    }
    .outline-none {
      --tw-outline-style: none;
      outline-style: none;
    }
    .not-disabled\:cursor-pointer {
      &:not(*:disabled) {
        cursor: pointer;
      }
    }
    .not-in-data-\[loading\=true\]\:block {
      &:not(:where(*[data-loading="true"]) *) {
        display: block;
      }
    }
    .group-hover\/tooltip\:opacity-100 {
      &:is(:where(.group\/tooltip):hover *) {
        @media (hover: hover) {
          opacity: 100%;
        }
      }
    }
    .group-\[\&\[data-active-form\=application\]\]\:block {
      &:is(:where(.group)[data-active-form="application"] *) {
        display: block;
      }
    }
    .group-\[\&\[data-active-form\=one-time-pass\]\]\:block {
      &:is(:where(.group)[data-active-form="one-time-pass"] *) {
        display: block;
      }
    }
    .peer-focus\:stroke-\[\#304A79\] {
      &:is(:where(.peer):focus ~ *) {
        stroke: #304a79;
      }
    }
    .peer-data-\[disabled\=true\]\:flex {
      &:is(:where(.peer)[data-disabled="true"] ~ *) {
        display: flex;
      }
    }
    .peer-data-\[loading\=true\]\:block {
      &:is(:where(.peer)[data-loading="true"] ~ *) {
        display: block;
      }
    }
    .placeholder\:text-\[\#7C96C6\] {
      &::placeholder {
        color: #7c96c6;
      }
    }
    .placeholder\:text-\[\#C2C4C5\] {
      &::placeholder {
        color: #c2c4c5;
      }
    }
    .before\:absolute {
      &::before {
        content: var(--tw-content);
        position: absolute;
      }
    }
    .before\:top-0 {
      &::before {
        content: var(--tw-content);
        top: calc(var(--spacing) * 0);
      }
    }
    .before\:left-0 {
      &::before {
        content: var(--tw-content);
        left: calc(var(--spacing) * 0);
      }
    }
    .before\:h-full {
      &::before {
        content: var(--tw-content);
        height: 100%;
      }
    }
    .before\:w-full {
      &::before {
        content: var(--tw-content);
        width: 100%;
      }
    }
    .before\:rounded-2xl {
      &::before {
        content: var(--tw-content);
        border-radius: var(--radius-2xl);
      }
    }
    .before\:rounded-3xl {
      &::before {
        content: var(--tw-content);
        border-radius: var(--radius-3xl);
      }
    }
    .before\:rounded-full {
      &::before {
        content: var(--tw-content);
        border-radius: calc(infinity * 1px);
      }
    }
    .before\:bg-\[linear-gradient\(83\.19deg\,\#A1B4D6_5\.3\%\,\#4469AD_49\.08\%\,\#A1B4D6_94\.64\%\)\] {
      &::before {
        content: var(--tw-content);
        background-image: linear-gradient(
          83.19deg,
          #a1b4d6 5.3%,
          #4469ad 49.08%,
          #a1b4d6 94.64%
        );
      }
    }
    .before\:bg-\[linear-gradient\(88\.93deg\,\#304A79_0\.91\%\,\#7C96C6_98\.92\%\)\] {
      &::before {
        content: var(--tw-content);
        background-image: linear-gradient(
          88.93deg,
          #304a79 0.91%,
          #7c96c6 98.92%
        );
      }
    }
    .before\:bg-\[linear-gradient\(193\.24deg\,rgba\(61\,95\,156\,0\)41\.8\%\,\#A1B4D6_101\.77\%\)\] {
      &::before {
        content: var(--tw-content);
        background-image: linear-gradient(
          193.24deg,
          rgba(61, 95, 156, 0) 41.8%,
          #a1b4d6 101.77%
        );
      }
    }
    .before\:bg-\[linear-gradient\(360deg\,rgba\(61\,95\,156\,0\)0\%\,\#A1B4D6_100\%\)\] {
      &::before {
        content: var(--tw-content);
        background-image: linear-gradient(
          360deg,
          rgba(61, 95, 156, 0) 0%,
          #a1b4d6 100%
        );
      }
    }
    .before\:content-\[\'\'\] {
      &::before {
        --tw-content: "";
        content: var(--tw-content);
      }
    }
    .hover\:bg-white\/20 {
      &:hover {
        @media (hover: hover) {
          background-color: color-mix(in srgb, #fff 20%, transparent);
          @supports (color: color-mix(in lab, red, red)) {
            background-color: color-mix(
              in oklab,
              var(--color-white) 20%,
              transparent
            );
          }
        }
      }
    }
    .focus\:border-\[\#304A79\] {
      &:focus {
        border-color: #304a79;
      }
    }
    .focus\:outline-none {
      &:focus {
        --tw-outline-style: none;
        outline-style: none;
      }
    }
    .disabled\:bg-\[\#DADBDC\] {
      &:disabled {
        background-color: #dadbdc;
      }
    }
    .disabled\:text-\[\#85888B\] {
      &:disabled {
        color: #85888b;
      }
    }
    .in-has-\[\#register-foreigner-form\.hidden\]\:hidden {
      :where(*:has(*:is(#register-foreigner-form.hidden))) & {
        display: none;
      }
    }
    .in-has-\[\#register-iranian-form\.hidden\]\:hidden {
      :where(*:has(*:is(#register-iranian-form.hidden))) & {
        display: none;
      }
    }
    .in-data-\[loading\=true\]\:block {
      :where(*[data-loading="true"]) & {
        display: block;
      }
    }
    .in-data-\[selected\=true\]\:block {
      :where(*[data-selected="true"]) & {
        display: block;
      }
    }
    .in-data-\[visible\=false\]\:block {
      :where(*[data-visible="false"]) & {
        display: block;
      }
    }
    .in-data-\[visible\=false\]\:hidden {
      :where(*[data-visible="false"]) & {
        display: none;
      }
    }
    .in-data-\[visible\=true\]\:block {
      :where(*[data-visible="true"]) & {
        display: block;
      }
    }
    .in-data-\[visible\=true\]\:hidden {
      :where(*[data-visible="true"]) & {
        display: none;
      }
    }
    .data-\[disabled\=true\]\:opacity-10 {
      &[data-disabled="true"] {
        opacity: 10%;
      }
    }
    .data-\[disabled\=true\]\:blur-\[1px\] {
      &[data-disabled="true"] {
        --tw-blur: blur(1px);
        filter: var(--tw-blur,) var(--tw-brightness,) var(--tw-contrast,)
          var(--tw-grayscale,) var(--tw-hue-rotate,) var(--tw-invert,)
          var(--tw-saturate,) var(--tw-sepia,) var(--tw-drop-shadow,);
      }
    }
    .data-\[loading\=true\]\:opacity-10 {
      &[data-loading="true"] {
        opacity: 10%;
      }
    }
    .data-\[loading\=true\]\:blur-\[1px\] {
      &[data-loading="true"] {
        --tw-blur: blur(1px);
        filter: var(--tw-blur,) var(--tw-brightness,) var(--tw-contrast,)
          var(--tw-grayscale,) var(--tw-hue-rotate,) var(--tw-invert,)
          var(--tw-saturate,) var(--tw-sepia,) var(--tw-drop-shadow,);
      }
    }
    .data-\[selected\=true\]\:border {
      &[data-selected="true"] {
        border-style: var(--tw-border-style);
        border-width: 1px;
      }
    }
    .data-\[selected\=true\]\:border-\[\#C7D2E6\] {
      &[data-selected="true"] {
        border-color: #c7d2e6;
      }
    }
    .data-\[selected\=true\]\:bg-linear-270 {
      &[data-selected="true"] {
        --tw-gradient-position: 270deg;
        @supports (background-image: linear-gradient(in lab, red, red)) {
          --tw-gradient-position: 270deg in oklab;
        }
        background-image: linear-gradient(var(--tw-gradient-stops));
      }
    }
    .data-\[selected\=true\]\:from-\[\#ECF0F7\] {
      &[data-selected="true"] {
        --tw-gradient-from: #ecf0f7;
        --tw-gradient-stops: var(
          --tw-gradient-via-stops,
          var(--tw-gradient-position),
          var(--tw-gradient-from) var(--tw-gradient-from-position),
          var(--tw-gradient-to) var(--tw-gradient-to-position)
        );
      }
    }
    .data-\[selected\=true\]\:to-white {
      &[data-selected="true"] {
        --tw-gradient-to: var(--color-white);
        --tw-gradient-stops: var(
          --tw-gradient-via-stops,
          var(--tw-gradient-position),
          var(--tw-gradient-from) var(--tw-gradient-from-position),
          var(--tw-gradient-to) var(--tw-gradient-to-position)
        );
      }
    }
    .data-\[selected\=true\]\:opacity-100 {
      &[data-selected="true"] {
        opacity: 100%;
      }
    }
    .data-\[selected\=true\]\:shadow-\[0px_4px_11px_0px_\#0000001A\] {
      &[data-selected="true"] {
        --tw-shadow: 0px 4px 11px 0px var(--tw-shadow-color, #0000001a);
        box-shadow:
          var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
          var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
      }
    }
    .data-\[type\=error\]\:bg-\[linear-gradient\(215deg\,\#FFE3E2_2\.43\%\,\#FFFFFF_34\.3\%\)\] {
      &[data-type="error"] {
        background-image: linear-gradient(215deg, #ffe3e2 2.43%, #ffffff 34.3%);
      }
    }
    .data-\[type\=success\]\:bg-\[linear-gradient\(215deg\,\#98EDB3_2\.43\%\,\#FFFFFF_34\.3\%\)\] {
      &[data-type="success"] {
        background-image: linear-gradient(215deg, #98edb3 2.43%, #ffffff 34.3%);
      }
    }
    .data-\[type\=warning\]\:bg-\[linear-gradient\(215deg\,\#FFEDC2_2\.43\%\,\#FFFFFF_34\.3\%\)\] {
      &[data-type="warning"] {
        background-image: linear-gradient(215deg, #ffedc2 2.43%, #ffffff 34.3%);
      }
    }
    .sm\:right-11 {
      @media (width >= 40rem) {
        right: calc(var(--spacing) * 11);
      }
    }
    .sm\:mt-4 {
      @media (width >= 40rem) {
        margin-top: calc(var(--spacing) * 4);
      }
    }
    .sm\:mt-\[48px\] {
      @media (width >= 40rem) {
        margin-top: 48px;
      }
    }
    .sm\:mt-\[56px\] {
      @media (width >= 40rem) {
        margin-top: 56px;
      }
    }
    .sm\:mt-\[64px\] {
      @media (width >= 40rem) {
        margin-top: 64px;
      }
    }
    .sm\:mr-0 {
      @media (width >= 40rem) {
        margin-right: calc(var(--spacing) * 0);
      }
    }
    .sm\:mb-4 {
      @media (width >= 40rem) {
        margin-bottom: calc(var(--spacing) * 4);
      }
    }
    .sm\:mb-\[40px\] {
      @media (width >= 40rem) {
        margin-bottom: 40px;
      }
    }
    .sm\:mb-\[56px\] {
      @media (width >= 40rem) {
        margin-bottom: 56px;
      }
    }
    .sm\:max-w-\[70\%\] {
      @media (width >= 40rem) {
        max-width: 70%;
      }
    }
    .sm\:px-11 {
      @media (width >= 40rem) {
        padding-inline: calc(var(--spacing) * 11);
      }
    }
    .sm\:text-center {
      @media (width >= 40rem) {
        text-align: center;
      }
    }
    .sm\:text-2xl {
      @media (width >= 40rem) {
        font-size: var(--text-2xl);
        line-height: var(--tw-leading, var(--text-2xl--line-height));
      }
    }
    .sm\:text-sm {
      @media (width >= 40rem) {
        font-size: var(--text-sm);
        line-height: var(--tw-leading, var(--text-sm--line-height));
      }
    }
    .sm\:text-\[18px\] {
      @media (width >= 40rem) {
        font-size: 18px;
      }
    }
    .sm\:text-\[24px\] {
      @media (width >= 40rem) {
        font-size: 24px;
      }
    }
    .sm\:font-normal {
      @media (width >= 40rem) {
        --tw-font-weight: var(--font-weight-normal);
        font-weight: var(--font-weight-normal);
      }
    }
    .md\:top-\[-26\%\] {
      @media (width >= 48rem) {
        top: -26%;
      }
    }
    .md\:right-\[-73\%\] {
      @media (width >= 48rem) {
        right: -73%;
      }
    }
    .md\:right-\[-80\%\] {
      @media (width >= 48rem) {
        right: -80%;
      }
    }
    .md\:bottom-\[-28\%\] {
      @media (width >= 48rem) {
        bottom: -28%;
      }
    }
    .md\:bottom-\[-29\%\] {
      @media (width >= 48rem) {
        bottom: -29%;
      }
    }
    .md\:left-\[-66\%\] {
      @media (width >= 48rem) {
        left: -66%;
      }
    }
    .md\:left-\[-70\%\] {
      @media (width >= 48rem) {
        left: -70%;
      }
    }
    .md\:order-none {
      @media (width >= 48rem) {
        order: 0;
      }
    }
    .md\:col-span-6 {
      @media (width >= 48rem) {
        grid-column: span 6 / span 6;
      }
    }
    .md\:col-span-12 {
      @media (width >= 48rem) {
        grid-column: span 12 / span 12;
      }
    }
    .md\:mt-4 {
      @media (width >= 48rem) {
        margin-top: calc(var(--spacing) * 4);
      }
    }
    .md\:mb-\[28\%\] {
      @media (width >= 48rem) {
        margin-bottom: 28%;
      }
    }
    .md\:mb-\[29\%\] {
      @media (width >= 48rem) {
        margin-bottom: 29%;
      }
    }
    .md\:h-15 {
      @media (width >= 48rem) {
        height: calc(var(--spacing) * 15);
      }
    }
    .md\:h-full {
      @media (width >= 48rem) {
        height: 100%;
      }
    }
    .md\:h-screen {
      @media (width >= 48rem) {
        height: 100vh;
      }
    }
    .md\:w-10\/12 {
      @media (width >= 48rem) {
        width: calc(10 / 12 * 100%);
      }
    }
    .md\:w-full {
      @media (width >= 48rem) {
        width: 100%;
      }
    }
    .md\:max-w-\[70\%\] {
      @media (width >= 48rem) {
        max-width: 70%;
      }
    }
    .md\:justify-center {
      @media (width >= 48rem) {
        justify-content: center;
      }
    }
    .md\:gap-6 {
      @media (width >= 48rem) {
        gap: calc(var(--spacing) * 6);
      }
    }
    .md\:rounded-l-3xl {
      @media (width >= 48rem) {
        border-top-left-radius: var(--radius-3xl);
        border-bottom-left-radius: var(--radius-3xl);
      }
    }
    .md\:rounded-r-3xl {
      @media (width >= 48rem) {
        border-top-right-radius: var(--radius-3xl);
        border-bottom-right-radius: var(--radius-3xl);
      }
    }
    .md\:border {
      @media (width >= 48rem) {
        border-style: var(--tw-border-style);
        border-width: 1px;
      }
    }
    .md\:px-8 {
      @media (width >= 48rem) {
        padding-inline: calc(var(--spacing) * 8);
      }
    }
    .md\:py-32 {
      @media (width >= 48rem) {
        padding-block: calc(var(--spacing) * 32);
      }
    }
    .md\:text-lg {
      @media (width >= 48rem) {
        font-size: var(--text-lg);
        line-height: var(--tw-leading, var(--text-lg--line-height));
      }
    }
    .md\:font-semibold {
      @media (width >= 48rem) {
        --tw-font-weight: var(--font-weight-semibold);
        font-weight: var(--font-weight-semibold);
      }
    }
    .md\:shadow-lg {
      @media (width >= 48rem) {
        --tw-shadow:
          0 10px 15px -3px var(--tw-shadow-color, rgb(0 0 0 / 0.1)),
          0 4px 6px -4px var(--tw-shadow-color, rgb(0 0 0 / 0.1));
        box-shadow:
          var(--tw-inset-shadow), var(--tw-inset-ring-shadow),
          var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow);
      }
    }
    .lg\:block {
      @media (width >= 64rem) {
        display: block;
      }
    }
    .lg\:flex {
      @media (width >= 64rem) {
        display: flex;
      }
    }
    .lg\:hidden {
      @media (width >= 64rem) {
        display: none;
      }
    }
    .lg\:w-4\/6 {
      @media (width >= 64rem) {
        width: calc(4 / 6 * 100%);
      }
    }
    .xl\:min-w-\[660px\] {
      @media (width >= 80rem) {
        min-width: 660px;
      }
    }
    .xl\:px-28 {
      @media (width >= 80rem) {
        padding-inline: calc(var(--spacing) * 28);
      }
    }
  }
  @property --tw-translate-x {
    syntax: "*";
    inherits: false;
    initial-value: 0;
  }
  @property --tw-translate-y {
    syntax: "*";
    inherits: false;
    initial-value: 0;
  }
  @property --tw-translate-z {
    syntax: "*";
    inherits: false;
    initial-value: 0;
  }
  @property --tw-rotate-x {
    syntax: "*";
    inherits: false;
  }
  @property --tw-rotate-y {
    syntax: "*";
    inherits: false;
  }
  @property --tw-rotate-z {
    syntax: "*";
    inherits: false;
  }
  @property --tw-skew-x {
    syntax: "*";
    inherits: false;
  }
  @property --tw-skew-y {
    syntax: "*";
    inherits: false;
  }
  @property --tw-space-y-reverse {
    syntax: "*";
    inherits: false;
    initial-value: 0;
  }
  @property --tw-border-style {
    syntax: "*";
    inherits: false;
    initial-value: solid;
  }
  @property --tw-gradient-position {
    syntax: "*";
    inherits: false;
  }
  @property --tw-gradient-from {
    syntax: "<color>";
    inherits: false;
    initial-value: #0000;
  }
  @property --tw-gradient-via {
    syntax: "<color>";
    inherits: false;
    initial-value: #0000;
  }
  @property --tw-gradient-to {
    syntax: "<color>";
    inherits: false;
    initial-value: #0000;
  }
  @property --tw-gradient-stops {
    syntax: "*";
    inherits: false;
  }
  @property --tw-gradient-via-stops {
    syntax: "*";
    inherits: false;
  }
  @property --tw-gradient-from-position {
    syntax: "<length-percentage>";
    inherits: false;
    initial-value: 0%;
  }
  @property --tw-gradient-via-position {
    syntax: "<length-percentage>";
    inherits: false;
    initial-value: 50%;
  }
  @property --tw-gradient-to-position {
    syntax: "<length-percentage>";
    inherits: false;
    initial-value: 100%;
  }
  @property --tw-font-weight {
    syntax: "*";
    inherits: false;
  }
  @property --tw-shadow {
    syntax: "*";
    inherits: false;
    initial-value: 0 0 #0000;
  }
  @property --tw-shadow-color {
    syntax: "*";
    inherits: false;
  }
  @property --tw-shadow-alpha {
    syntax: "<percentage>";
    inherits: false;
    initial-value: 100%;
  }
  @property --tw-inset-shadow {
    syntax: "*";
    inherits: false;
    initial-value: 0 0 #0000;
  }
  @property --tw-inset-shadow-color {
    syntax: "*";
    inherits: false;
  }
  @property --tw-inset-shadow-alpha {
    syntax: "<percentage>";
    inherits: false;
    initial-value: 100%;
  }
  @property --tw-ring-color {
    syntax: "*";
    inherits: false;
  }
  @property --tw-ring-shadow {
    syntax: "*";
    inherits: false;
    initial-value: 0 0 #0000;
  }
  @property --tw-inset-ring-color {
    syntax: "*";
    inherits: false;
  }
  @property --tw-inset-ring-shadow {
    syntax: "*";
    inherits: false;
    initial-value: 0 0 #0000;
  }
  @property --tw-ring-inset {
    syntax: "*";
    inherits: false;
  }
  @property --tw-ring-offset-width {
    syntax: "<length>";
    inherits: false;
    initial-value: 0px;
  }
  @property --tw-ring-offset-color {
    syntax: "*";
    inherits: false;
    initial-value: #fff;
  }
  @property --tw-ring-offset-shadow {
    syntax: "*";
    inherits: false;
    initial-value: 0 0 #0000;
  }
  @property --tw-blur {
    syntax: "*";
    inherits: false;
  }
  @property --tw-brightness {
    syntax: "*";
    inherits: false;
  }
  @property --tw-contrast {
    syntax: "*";
    inherits: false;
  }
  @property --tw-grayscale {
    syntax: "*";
    inherits: false;
  }
  @property --tw-hue-rotate {
    syntax: "*";
    inherits: false;
  }
  @property --tw-invert {
    syntax: "*";
    inherits: false;
  }
  @property --tw-opacity {
    syntax: "*";
    inherits: false;
  }
  @property --tw-saturate {
    syntax: "*";
    inherits: false;
  }
  @property --tw-sepia {
    syntax: "*";
    inherits: false;
  }
  @property --tw-drop-shadow {
    syntax: "*";
    inherits: false;
  }
  @property --tw-drop-shadow-color {
    syntax: "*";
    inherits: false;
  }
  @property --tw-drop-shadow-alpha {
    syntax: "<percentage>";
    inherits: false;
    initial-value: 100%;
  }
  @property --tw-drop-shadow-size {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-blur {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-brightness {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-contrast {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-grayscale {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-hue-rotate {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-invert {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-opacity {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-saturate {
    syntax: "*";
    inherits: false;
  }
  @property --tw-backdrop-sepia {
    syntax: "*";
    inherits: false;
  }
  @property --tw-duration {
    syntax: "*";
    inherits: false;
  }
  @property --tw-ease {
    syntax: "*";
    inherits: false;
  }
  @property --tw-content {
    syntax: "*";
    initial-value: "";
    inherits: false;
  }
  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
  @layer properties {
    @supports ((-webkit-hyphens: none) and (not (margin-trim: inline))) or
      ((-moz-orient: inline) and (not (color: rgb(from red r g b)))) {
      *,
      ::before,
      ::after,
      ::backdrop {
        --tw-translate-x: 0;
        --tw-translate-y: 0;
        --tw-translate-z: 0;
        --tw-rotate-x: initial;
        --tw-rotate-y: initial;
        --tw-rotate-z: initial;
        --tw-skew-x: initial;
        --tw-skew-y: initial;
        --tw-space-y-reverse: 0;
        --tw-border-style: solid;
        --tw-gradient-position: initial;
        --tw-gradient-from: #0000;
        --tw-gradient-via: #0000;
        --tw-gradient-to: #0000;
        --tw-gradient-stops: initial;
        --tw-gradient-via-stops: initial;
        --tw-gradient-from-position: 0%;
        --tw-gradient-via-position: 50%;
        --tw-gradient-to-position: 100%;
        --tw-font-weight: initial;
        --tw-shadow: 0 0 #0000;
        --tw-shadow-color: initial;
        --tw-shadow-alpha: 100%;
        --tw-inset-shadow: 0 0 #0000;
        --tw-inset-shadow-color: initial;
        --tw-inset-shadow-alpha: 100%;
        --tw-ring-color: initial;
        --tw-ring-shadow: 0 0 #0000;
        --tw-inset-ring-color: initial;
        --tw-inset-ring-shadow: 0 0 #0000;
        --tw-ring-inset: initial;
        --tw-ring-offset-width: 0px;
        --tw-ring-offset-color: #fff;
        --tw-ring-offset-shadow: 0 0 #0000;
        --tw-blur: initial;
        --tw-brightness: initial;
        --tw-contrast: initial;
        --tw-grayscale: initial;
        --tw-hue-rotate: initial;
        --tw-invert: initial;
        --tw-opacity: initial;
        --tw-saturate: initial;
        --tw-sepia: initial;
        --tw-drop-shadow: initial;
        --tw-drop-shadow-color: initial;
        --tw-drop-shadow-alpha: 100%;
        --tw-drop-shadow-size: initial;
        --tw-backdrop-blur: initial;
        --tw-backdrop-brightness: initial;
        --tw-backdrop-contrast: initial;
        --tw-backdrop-grayscale: initial;
        --tw-backdrop-hue-rotate: initial;
        --tw-backdrop-invert: initial;
        --tw-backdrop-opacity: initial;
        --tw-backdrop-saturate: initial;
        --tw-backdrop-sepia: initial;
        --tw-duration: initial;
        --tw-ease: initial;
        --tw-content: "";
      }
    }
  }
</style>
  </head>
  <body class="overflow-x-hidden">
    <svg width="0" height="0" style="opacity: 0">
  <defs>
    <!-- main-logo -->
    <clipPath id="main-logo">
      <rect width="66" height="56" fill="white" />
    </clipPath>

    <!-- qr-btn -->
    <filter
      id="qr-btn"
      x="-0.5"
      y="-3"
      width="69"
      height="69"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="4" dy="-4" />
      <feGaussianBlur stdDeviation="4" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 1 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_3769_168764"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-3" dy="3" />
      <feGaussianBlur stdDeviation="4" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.3 0"
      />
      <feBlend
        mode="normal"
        in2="effect1_dropShadow_3769_168764"
        result="effect2_dropShadow_3769_168764"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect2_dropShadow_3769_168764"
        result="shape"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-2" dy="2" />
      <feGaussianBlur stdDeviation="1.5" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.45 0"
      />
      <feBlend
        mode="normal"
        in2="shape"
        result="effect3_innerShadow_3769_168764"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="2" dy="-2" />
      <feGaussianBlur stdDeviation="1.5" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.4 0"
      />
      <feBlend
        mode="normal"
        in2="effect3_innerShadow_3769_168764"
        result="effect4_innerShadow_3769_168764"
      />
    </filter>
    <linearGradient
      id="paint0_linear_3769_168764"
      x1="29.2011"
      y1="10.4735"
      x2="19.8287"
      y2="30.4327"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint1_linear_3769_168764"
      x1="29.2011"
      y1="35.7743"
      x2="19.8287"
      y2="55.7334"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint2_linear_3769_168764"
      x1="54.5011"
      y1="10.4735"
      x2="45.1287"
      y2="30.4327"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint3_linear_3769_168764"
      x1="46.1895"
      y1="35.2831"
      x2="39.2437"
      y2="47.6094"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint4_linear_3769_168764"
      x1="42.0337"
      y1="48.5908"
      x2="38.9096"
      y2="55.2438"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint5_linear_3769_168764"
      x1="53.5337"
      y1="34.792"
      x2="50.4095"
      y2="41.445"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint6_linear_3769_168764"
      x1="53.5337"
      y1="48.5908"
      x2="50.4095"
      y2="55.2438"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>

    <!-- app-btn -->
    <filter
      id="app-btn"
      x="0.249841"
      y="-3.25016"
      width="52.5003"
      height="70.5003"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="3.86847" dy="-3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 1 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_3769_168762"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-3.86847" dy="3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.5 0"
      />
      <feBlend
        mode="normal"
        in2="effect1_dropShadow_3769_168762"
        result="effect2_dropShadow_3769_168762"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect2_dropShadow_3769_168762"
        result="shape"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-1.93424" dy="1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.77 0"
      />
      <feBlend
        mode="normal"
        in2="shape"
        result="effect3_innerShadow_3769_168762"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="1.93424" dy="-1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.45 0"
      />
      <feBlend
        mode="normal"
        in2="effect3_innerShadow_3769_168762"
        result="effect4_innerShadow_3769_168762"
      />
    </filter>
    <linearGradient
      id="paint0_linear_3769_168762"
      x1="37.7961"
      y1="12.2746"
      x2="11.6151"
      y2="46.2114"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>

    <!-- otp-btn -->
    <filter
      id="otp-btn"
      x="-2.75016"
      y="-3.25016"
      width="70.5003"
      height="70.5003"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="3.86847" dy="-3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 1 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_3876_137583"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-3.86847" dy="3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.5 0"
      />
      <feBlend
        mode="normal"
        in2="effect1_dropShadow_3876_137583"
        result="effect2_dropShadow_3876_137583"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect2_dropShadow_3876_137583"
        result="shape"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-1.93424" dy="1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.77 0"
      />
      <feBlend
        mode="normal"
        in2="shape"
        result="effect3_innerShadow_3876_137583"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="1.93424" dy="-1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.45 0"
      />
      <feBlend
        mode="normal"
        in2="effect3_innerShadow_3876_137583"
        result="effect4_innerShadow_3876_137583"
      />
    </filter>
    <linearGradient
      id="paint0_linear_3876_137583"
      x1="51.0579"
      y1="12.2746"
      x2="30.2299"
      y2="56.6286"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>

    <!-- grid-box -->
    <linearGradient
      id="grid-box"
      x1="235"
      y1="5.99999"
      x2="120"
      y2="133.5"
      gradientUnits="userSpaceOnUse"
    >
      <stop offset="0.0001" stop-color="#86C6FF" />
      <stop offset="0.203125" stop-color="#86C6FF" stop-opacity="0.5" />
      <stop offset="1" stop-color="#F7F7F7" stop-opacity="0" />
    </linearGradient>

    <!-- messenger-btn -->
    <filter
      id="messenger-btn"
      x="-3.25016"
      y="-0.250159"
      width="70.5003"
      height="64.5003"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="3.86847" dy="-3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 1 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_7670_204772"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-3.86847" dy="3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.5 0"
      />
      <feBlend
        mode="normal"
        in2="effect1_dropShadow_7670_204772"
        result="effect2_dropShadow_7670_204772"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect2_dropShadow_7670_204772"
        result="shape"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-1.93424" dy="1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.77 0"
      />
      <feBlend
        mode="normal"
        in2="shape"
        result="effect3_innerShadow_7670_204772"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="1.93424" dy="-1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.45 0"
      />
      <feBlend
        mode="normal"
        in2="effect3_innerShadow_7670_204772"
        result="effect4_innerShadow_7670_204772"
      />
    </filter>
    <linearGradient
      id="paint0_linear_7670_204772"
      x1="41.3344"
      y1="14.4324"
      x2="26.3236"
      y2="47.9148"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>
    <linearGradient
      id="paint1_linear_7670_204772"
      x1="52.5401"
      y1="28.3353"
      x2="41.0022"
      y2="52.9013"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>

    <!-- email-btn -->
    <filter
      id="email-btn"
      x="-3.25016"
      y="0.249841"
      width="70.5003"
      height="55.5003"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="3.86847" dy="-3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 1 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_3790_173457"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-3.86847" dy="3.86847" />
      <feGaussianBlur stdDeviation="4.19084" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.5 0"
      />
      <feBlend
        mode="normal"
        in2="effect1_dropShadow_3790_173457"
        result="effect2_dropShadow_3790_173457"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect2_dropShadow_3790_173457"
        result="shape"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-1.93424" dy="1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.77 0"
      />
      <feBlend
        mode="normal"
        in2="shape"
        result="effect3_innerShadow_3790_173457"
      />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="1.93424" dy="-1.93424" />
      <feGaussianBlur stdDeviation="1.19278" />
      <feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.45 0"
      />
      <feBlend
        mode="normal"
        in2="effect3_innerShadow_3790_173457"
        result="effect4_innerShadow_3790_173457"
      />
    </filter>
    <linearGradient
      id="paint0_linear_3790_173457"
      x1="50.5579"
      y1="14.7068"
      x2="40.0638"
      y2="47.8678"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#7C96C6" />
      <stop offset="1" stop-color="#4469AD" />
    </linearGradient>

    <!-- stroke-right -->
    <linearGradient
      id="stroke-right"
      x1="0.5"
      y1="1.5"
      x2="188.997"
      y2="1.5"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#C2C4C5" />
      <stop offset="1" stop-color="white" stop-opacity="0" />
    </linearGradient>

    <!-- stroke-left -->
    <linearGradient
      id="stroke-left"
      x1="-0.99734"
      y1="1.5"
      x2="187.5"
      y2="1.5"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="white" stop-opacity="0" />
      <stop offset="1" stop-color="#C2C4C5" />
    </linearGradient>

    <!-- sibche-logo -->
    <pattern
      id="pattern0_2842_101903"
      patternContentUnits="objectBoundingBox"
      width="1"
      height="1"
    >
      <use
        xlink:href="#image0_2842_101903"
        transform="matrix(0.0116279 0 0 0.010341 -2.94186 -0.0444444)"
      />
    </pattern>
    <image
      id="image0_2842_101903"
      width="339"
      height="101"
      preserveAspectRatio="none"
      xlink:href="data:image/png;base64,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"
    />

    <!-- sibche-title -->
    <pattern
      id="pattern0_2842_101973"
      patternContentUnits="objectBoundingBox"
      width="1"
      height="1"
    >
      <use
        xlink:href="#image0_2842_101973"
        transform="matrix(0.0023753 0 0 0.00724638 -0.106888 -1.99275)"
      />
    </pattern>
    <image
      id="image0_2842_101973"
      width="512"
      height="512"
      preserveAspectRatio="none"
      xlink:href="data:image/png;base64,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0mgkzb7ZxmAWmzU2weCFNew4BkwwTiAQAphwIYQAyb03kBt8MWYLW4DcGMJi9s2ADbED2MYGjLFnvMyMZrRLMyOp9+7azvv87h/vOVXVpZbUUldVL/p9P5+Wuk+dOvt5n3c77wFEREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREZPey/gkiIsOSRwaSwYFApvSHMDOAAMEiTTKAweAG0AxuBmYhcOPSRGQ7lAEQkYHLo4dIVHLnWDPa3GqLD7QiPhpmz60Eu2+6hoOTGScnqlZBJx1ypn9DXM/Rrrfi2nqOp3PHSSdOjlftw4fG8aHpWrbUsyoRuUnKAIjIQLRzrzSdU5fr+PhWtJccmbQXzo7hGIBpAIcAzAHINn7r2gjA4AAQiXDiwkr8l3fNVv73xrlE5GZU+ieIiGxVdLdmxNRTq/4Zl+r+mqPT2cdMVXEM8MNwxybx3pHi+pYUwd+JUDWgPVG1D/XPIyIiIiPSjp5dWo/3PLEQvy+S7yW54kycJL0dncxJxp6fG+Yko6fvn1rMX59HD/3bIiIiIkMW3e2ZtXjf6aX4U06e6ovtOcncbzLYb8bdIxlJzy+fWY4v6N8eERERGbKlRjz0xHz7+6LzVCqUkySLkn6MZFn6Lwv+gxBzd+dyk/9rveUT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i2B+b09JOwJpND3SAsyKm7fbfSDVGmyctl3lMru1DqmkTwDGWNQIBC9mycp1E2F+pYlHnlrLf+Ku6ewtE1WsZSFs1mAgIjtEGQCRXaaZe22+wfvW2/iGBw6ELzLgfgITQOeG9ehAFhxA6ER5kp0q+m5GYLuZgP4MhqMI+ACcROgd3bDhxJmL6/zbtbb/j6OT4a+nx8Jq92MR2U2UARDZpdrRs7U2D15Y9c+7bbLyNbdP4Dlw3olg4wQAEmYEEGL6hlvZDm+GUJbWt8mBsk8Byxb+QAARQAVow1vPtFA7eWqJvzk3xt+fq9mF8WpobliKiOw6ygCI7HLR3VoRYystHF1o8OWHav6Fh6crzwZwFMDhcr4UpA1FzHZnQLCNfQa3oHf+UE4oaxRIA8yWAFyut+LJc2v2p1NV/OlcDSfGK7ZeSdUSIrIHbDVREJFdILpb7qg2cs5crvNjc8eLD0xk/+j2CdwJYJbAjMGnQU7Qsu3e4BHAmhOrwXx5rW0XL6zy7wC87cA43j9VsUu1DE0FfZG9aZvpg4jspDx6iESllXNiLcfh9Rz3t6M9txrw7IkqHrptgrdXQhgDUDGgAiADYCSDmTngDgQnI82s3c7ZXG6FxZUcZ9uRTwD40HjGE9M1uzCW2UotoKVH+ET2B2UARPaZPHpwIotE5s4sEhmJisMqTlYcViERDPAQ2A5ADMFyI2IwxmAWs4AYAA8BUb33RURERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERERGTorH+CbE90NxJGwEgE9h3jALgZ3AzMQmDvZ3tddDd3ZP3Tt6L/OI2SASRgV/u/d77e78H6/r6K8ntWzL+XznsePUSiQl7//JT7V9r4ne6vG+cjgiFWsxC7066uHT1zIguA93+2XWbwShYGvlyR3eq6N7VcXXS36KhEZ7UecWCthfvaznuIcF8w3FXN/I5qyCaDASTzPOb1nNnlnOFcJJ+qwM9PVMPJySovjmW2lhnyvZwxOLmYf36WhVc32mxmBgcQCGujSKyvHSAAAAEbE/by79AzrZdhY1Au5y2XcbXv9etdR+8yi7waSBjNLDqQu7MFskWz2PO9DkvLoxloJGmh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ZjxwfpVffMdU+LrJMbs3A+ZI1szMAM8BNIiwvNzgyafr+M3bJ/jmmZpd2urTDiKjoAyA7Ih29KyRY2qxiQdWW/w4gz9gFqYBWwN4eqJqD8/VeGKiYiu1StjWi1LauVcakZOrbbtjve13R4YjJKcAeGa4NFG1J2ZrPDeW2dp215VHD2s55i7X8fGtPH5iMLs3ZGE65rxE4MnxKh8+OGaPTVZseT8M/tLMvVbPObvY4LMa0Z5F4rZgqIJcqGZ2arqG41NVe2a8gvrNBNqb0cy91oqcXGvj9mbkQSKMGb05VrGFyQoujVVsbVCZrnrbx9favG21xftz2mEQwczqtQxPTVd5frxiy7UMzVHtu8iNUAZAdoU8emAajoVm4DATzN7q9mGuJ48eWAydZwYf9n7ttOhuLMbH2e/7upnyurrV9ltEREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREREQG4v8HdiIZtkB/rc0AAAAASUVORK5CYII="
    />

    <!-- telephone -->
    <clipPath id="clip0_1494_69546">
      <rect
        width="24"
        height="24"
        fill="white"
        transform="translate(8.2085) rotate(20)"
      />
    </clipPath>

    <!-- mail -->
    <clipPath id="clip0_1494_69585">
      <rect
        width="24"
        height="24"
        fill="white"
        transform="translate(8.2085) rotate(20)"
      />
    </clipPath>

    <!-- fax -->
    <clipPath id="clip0_1495_69538">
      <rect
        width="24"
        height="24"
        fill="white"
        transform="translate(8.2085) rotate(20)"
      />
    </clipPath>

    <!-- iaps-logo -->
    <pattern
      id="pattern0_2842_102012"
      patternContentUnits="objectBoundingBox"
      width="1"
      height="1"
    >
      <use
        xlink:href="#image0_2842_102012"
        transform="translate(-3.21505 -0.129032) scale(0.0107527)"
      />
    </pattern>
    <clipPath id="clip0_2842_102012">
      <rect width="40" height="40" fill="white" />
    </clipPath>
    <image
      id="image0_2842_102012"
      width="400"
      height="107"
      preserveAspectRatio="none"
      xlink:href="data:image/png;base64,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"
    />

    <!-- iaps-title -->
    <pattern
      id="pattern0_2842_102062"
      patternContentUnits="objectBoundingBox"
      width="1"
      height="1"
    >
      <use
        xlink:href="#image0_2842_102062"
        transform="matrix(0.0034965 0 0 0.0153846 0 -0.569231)"
      />
    </pattern>
    <image
      id="image0_2842_102062"
      width="400"
      height="107"
      preserveAspectRatio="none"
      xlink:href="data:image/png;base64,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"
    />

    <!-- document -->
    <filter
      id="filter0_d_1480_70759"
      x="-8.74996e-05"
      y="0.00041759"
      width="40.2097"
      height="57.697"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="1.48317" dy="-0.697963" />
      <feGaussianBlur stdDeviation="1.39593" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.06 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_1480_70759"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect1_dropShadow_1480_70759"
        result="shape"
      />
    </filter>
    <filter
      id="filter1_d_1480_70759"
      x="0.697876"
      y="2.87952"
      width="37.4976"
      height="55.0796"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="0.872454" dy="0.872454" />
      <feGaussianBlur stdDeviation="0.741586" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.15 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_1480_70759"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect1_dropShadow_1480_70759"
        result="shape"
      />
    </filter>
    <filter
      id="filter2_d_1480_70759"
      x="14.1544"
      y="2.70046"
      width="23.0016"
      height="22.9352"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dx="-1.04695" dy="1.30868" />
      <feGaussianBlur stdDeviation="1.13419" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.15 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_1480_70759"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect1_dropShadow_1480_70759"
        result="shape"
      />
    </filter>
    <filter
      id="filter3_d_1480_70759"
      x="5.18129"
      y="25.3164"
      width="16.0739"
      height="3.1882"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dy="0.348982" />
      <feGaussianBlur stdDeviation="0.174491" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_1480_70759"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect1_dropShadow_1480_70759"
        result="shape"
      />
    </filter>
    <filter
      id="filter4_f_1480_70759"
      x="5.53707"
      y="25.416"
      width="15.6153"
      height="1.59382"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter5_f_1480_70759"
      x="5.57907"
      y="26.8828"
      width="15.3634"
      height="1.43366"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter6_f_1480_70759"
      x="5.62692"
      y="32.7461"
      width="23.7462"
      height="1.83015"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter7_f_1480_70759"
      x="5.59664"
      y="38.2988"
      width="23.7462"
      height="1.83015"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter8_f_1480_70759"
      x="6.03903"
      y="31.4004"
      width="23.2071"
      height="1.99616"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter9_f_1480_70759"
      x="6.00973"
      y="36.9531"
      width="23.2071"
      height="1.99616"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter10_d_1480_70759"
      x="5.1891"
      y="31.3867"
      width="24.5505"
      height="3.35031"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dy="0.348982" />
      <feGaussianBlur stdDeviation="0.174491" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_1480_70759"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect1_dropShadow_1480_70759"
        result="shape"
      />
    </filter>
    <filter
      id="filter11_d_1480_70759"
      x="5.15883"
      y="36.9375"
      width="24.5505"
      height="3.35031"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feColorMatrix
        in="SourceAlpha"
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
        result="hardAlpha"
      />
      <feOffset dy="0.348982" />
      <feGaussianBlur stdDeviation="0.174491" />
      <feComposite in2="hardAlpha" operator="out" />
      <feColorMatrix
        type="matrix"
        values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
      />
      <feBlend
        mode="normal"
        in2="BackgroundImageFix"
        result="effect1_dropShadow_1480_70759"
      />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="effect1_dropShadow_1480_70759"
        result="shape"
      />
    </filter>
    <filter
      id="filter12_f_1480_70759"
      x="5.5859"
      y="31.2089"
      width="24.1807"
      height="1.60163"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <filter
      id="filter13_f_1480_70759"
      x="5.55661"
      y="36.7617"
      width="24.1807"
      height="1.60163"
      filterUnits="userSpaceOnUse"
      color-interpolation-filters="sRGB"
    >
      <feFlood flood-opacity="0" result="BackgroundImageFix" />
      <feBlend
        mode="normal"
        in="SourceGraphic"
        in2="BackgroundImageFix"
        result="shape"
      />
      <feGaussianBlur
        stdDeviation="0.261736"
        result="effect1_foregroundBlur_1480_70759"
      />
    </filter>
    <linearGradient
      id="paint0_linear_1480_70759"
      x1="1.07286"
      y1="12.2924"
      x2="36.035"
      y2="15.4399"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#E1E1E1" />
      <stop offset="0.302349" stop-color="#F9F9F9" />
      <stop offset="0.608197" stop-color="#F2F2F2" />
      <stop offset="1" stop-color="#DADADA" />
    </linearGradient>
    <linearGradient
      id="paint1_linear_1480_70759"
      x1="14.9716"
      y1="4.71254"
      x2="16.8797"
      y2="55.5723"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="white" />
      <stop offset="0.442708" stop-color="#BCBCBC" />
      <stop offset="1" stop-color="#6B6B6B" />
    </linearGradient>
    <linearGradient
      id="paint2_linear_1480_70759"
      x1="14.9716"
      y1="4.71254"
      x2="18.112"
      y2="27.5767"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#F1F1F1" />
      <stop offset="1" stop-color="#D9D9DB" />
    </linearGradient>
    <linearGradient
      id="paint3_linear_1480_70759"
      x1="31.9925"
      y1="27.393"
      x2="8.67575"
      y2="11.3784"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#DFDFDF" />
      <stop offset="0.505208" stop-color="#F7F7F7" />
      <stop offset="0.96875" stop-color="#DFDFDF" />
    </linearGradient>
    <linearGradient
      id="paint4_linear_1480_70759"
      x1="13.1725"
      y1="24.877"
      x2="13.3557"
      y2="28.2263"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#9BAAF3" />
      <stop offset="0.411458" stop-color="#5273B0" />
      <stop offset="0.875" stop-color="#3D5F9C" />
    </linearGradient>
    <linearGradient
      id="paint5_linear_1480_70759"
      x1="17.4511"
      y1="31.2806"
      x2="17.4732"
      y2="34.1338"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#9BAAF3" />
      <stop offset="0.411458" stop-color="#5273B0" />
      <stop offset="0.875" stop-color="#3D5F9C" />
    </linearGradient>
    <linearGradient
      id="paint6_linear_1480_70759"
      x1="17.4208"
      y1="36.8313"
      x2="17.443"
      y2="39.6846"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color="#9BAAF3" />
      <stop offset="0.411458" stop-color="#5273B0" />
      <stop offset="0.875" stop-color="#3D5F9C" />
    </linearGradient>
  </defs>
</svg>
    <div class="flex items-center justify-center h-screen overflow-hidden">
      <div
        dir="rtl"
        class="w-full h-full overflow-visible bg-white flex items-center relative inset-shadow-[0_0_32px_0px_#00000090]"
      >
        <div class="flex">
  <div
    class="pt-17 h-full max-w-[170px] w-full min-w-[110px] lg:flex flex-col items-center hidden"
  >
    <svg
  width="60"
  height="40"
  viewBox="0 0 60 40"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  class="mb-23"
>
  <path
    d="M38.54 13.39C38.7276 13.39 38.9257 13.3782 39.1134 13.39C39.7911 13.4137 40.229 13.8638 40.2812 14.61C40.3125 15.1548 40.3125 15.7114 40.2812 16.2681C40.2395 17.0024 39.812 17.4643 39.176 17.5117C38.7068 17.5473 38.2376 17.5473 37.7684 17.5117C37.122 17.4643 36.6839 16.9314 36.6527 16.197C36.6318 15.7588 36.6422 15.3324 36.6527 14.8942C36.6631 13.9112 37.1115 13.4019 37.9665 13.39C38.1542 13.3782 38.3419 13.39 38.54 13.39Z"
    fill="#2A38D8"
  />
  <path
    d="M38.2232 31.3673C38.0125 31.3673 37.7927 31.3856 37.5821 31.3673C37.115 31.3125 36.7394 31.0018 36.6845 30.5358C36.6295 30.0607 36.6295 29.5673 36.6662 29.083C36.712 28.5896 37.06 28.2424 37.5546 28.1876C38.0034 28.1419 38.4705 28.1328 38.9193 28.1785C39.3772 28.2242 39.7436 28.4983 39.7986 28.9734C39.8627 29.4942 39.8627 30.0333 39.7986 30.5541C39.7344 31.0567 39.3589 31.3308 38.846 31.3765C38.6445 31.3856 38.4339 31.3673 38.2232 31.3673Z"
    fill="#2A38D8"
  />
  <path
    d="M23.1525 25.3692C22.9894 25.3692 22.8173 25.3692 22.6542 25.3692C21.9929 25.3485 21.5762 24.9129 21.54 24.1662C21.5218 23.7513 21.5218 23.3365 21.54 22.9217C21.5671 22.2475 21.9386 21.7808 22.5184 21.7186C22.917 21.6771 23.3337 21.6771 23.7323 21.7186C24.303 21.7705 24.6563 22.2061 24.7107 22.8698C24.7379 23.2639 24.7379 23.658 24.7198 24.0625C24.6835 24.9544 24.294 25.3692 23.5149 25.3796C23.3971 25.3692 23.2703 25.3692 23.1525 25.3692Z"
    fill="#2A38D8"
  />
  <path
    d="M36.6447 24.9053C36.6447 24.7184 36.6356 24.5211 36.6447 24.3341C36.6809 23.5864 37.0801 23.119 37.7331 23.0878C38.0597 23.0671 38.3953 23.0671 38.7309 23.0878C39.4112 23.119 39.8194 23.5967 39.8376 24.3757C39.8466 24.7703 39.8557 25.165 39.8285 25.57C39.7922 26.2555 39.4203 26.6813 38.8216 26.7436C38.4951 26.7747 38.1595 26.7644 37.8329 26.754C37.0347 26.7332 36.6809 26.3282 36.6537 25.4246C36.6447 25.2481 36.6537 25.0819 36.6447 24.9053Z"
    fill="#2A38D8"
  />
  <path
    d="M31.1455 22.8608C31.1455 22.6612 31.1455 22.4733 31.1455 22.2737C31.156 21.264 31.6068 20.7708 32.4979 20.7708C32.8438 20.7708 33.1898 20.7591 33.5462 20.7708C34.2801 20.7943 34.7623 21.3109 34.7938 22.1094C34.8148 22.5908 34.8148 23.0839 34.7938 23.5653C34.7623 24.3167 34.3745 24.7981 33.693 24.8803C33.2842 24.9273 32.8543 24.9273 32.4455 24.9038C31.5649 24.8568 31.1665 24.3637 31.156 23.3774C31.1455 23.213 31.1455 23.0369 31.1455 22.8608Z"
    fill="#2A38D8"
  />
  <path
    d="M44.8638 18.2687C44.8638 18.4691 44.8743 18.6603 44.8638 18.8607C44.8219 19.3797 44.4557 19.7622 43.8592 19.8078C43.3256 19.8533 42.7815 19.8533 42.2478 19.8078C41.6409 19.7622 41.2537 19.3615 41.2328 18.8242C41.2119 18.46 41.2223 18.0957 41.2328 17.7314C41.2537 17.1122 41.6723 16.6933 42.3734 16.6386C42.8233 16.6022 43.2837 16.6022 43.7337 16.6386C44.4557 16.6933 44.8533 17.094 44.8847 17.7223C44.8743 17.9045 44.8638 18.0866 44.8638 18.2687Z"
    fill="#2A38D8"
  />
  <path
    d="M33.1745 19.8195C32.904 19.8195 32.6218 19.8405 32.3513 19.8195C31.7046 19.7669 31.246 19.3989 31.1872 18.8207C31.1402 18.2844 31.1284 17.7377 31.1755 17.191C31.2342 16.5811 31.7516 16.1816 32.4336 16.1606C32.9393 16.1501 33.4567 16.1395 33.9623 16.1606C34.6443 16.1816 35.1499 16.5917 35.2205 17.2015C35.2675 17.6536 35.2793 18.1162 35.2558 18.5683C35.2087 19.441 34.7384 19.8195 33.7624 19.8405C33.5625 19.8405 33.3744 19.8405 33.1745 19.8405C33.1745 19.83 33.1745 19.8195 33.1745 19.8195Z"
    fill="#2A38D8"
  />
  <path
    d="M32.971 29.5253C32.7618 29.5253 32.5525 29.5358 32.3433 29.5253C31.6529 29.4937 31.1926 29.0727 31.1612 28.3779C31.1403 27.9253 31.1403 27.4622 31.1612 27.0095C31.1926 26.3253 31.6529 25.8727 32.3328 25.8516C32.7513 25.8411 33.1697 25.8411 33.5882 25.8516C34.3414 25.8727 34.7703 26.3148 34.8017 27.0832C34.8121 27.4516 34.8121 27.8201 34.8017 28.1885C34.7808 29.1043 34.3623 29.5253 33.4417 29.5358C33.2848 29.5358 33.1279 29.5253 32.971 29.5253Z"
    fill="#2A38D8"
  />
  <path
    d="M29.7623 26.0789C29.7623 26.2903 29.7714 26.5108 29.7623 26.7222C29.7256 27.2461 29.35 27.6413 28.8277 27.6689C28.4062 27.6965 27.9755 27.6965 27.5449 27.6689C26.9859 27.6321 26.6103 27.2461 26.5828 26.6855C26.5645 26.2719 26.5553 25.8583 26.5828 25.4447C26.6194 24.8657 27.0043 24.498 27.5907 24.4705C27.9297 24.4521 28.2596 24.4613 28.5986 24.4613C29.3866 24.4705 29.7623 24.8565 29.7714 25.6561C29.7714 25.8031 29.7714 25.941 29.7623 26.0789Z"
    fill="#2A38D8"
  />
  <path
    d="M43.0612 20.7893C43.2915 20.7893 43.5114 20.7774 43.7418 20.7893C44.3387 20.8247 44.81 21.2737 44.8519 21.9473C44.8937 22.5263 44.8937 23.129 44.8623 23.708C44.8204 24.3934 44.3806 24.8661 43.7732 24.9015C43.302 24.9251 42.8308 24.9251 42.37 24.9015C41.7103 24.8661 41.2495 24.3225 41.2286 23.5662C41.2181 23.1526 41.2181 22.739 41.2286 22.3254C41.2286 21.2974 41.6893 20.7774 42.6109 20.7656C42.747 20.7774 42.9041 20.7774 43.0612 20.7893Z"
    fill="#2A38D8"
  />
  <path
    d="M43.075 25.8543C43.3052 25.8543 43.525 25.8437 43.7553 25.8543C44.3833 25.8963 44.8335 26.3273 44.8649 26.9579C44.8963 27.452 44.8858 27.9355 44.8649 28.4295C44.8335 29.0392 44.4042 29.4701 43.7972 29.5122C43.3157 29.5437 42.8238 29.5437 42.3318 29.5122C41.7143 29.4807 41.2642 29.0182 41.2328 28.3875C41.2119 27.946 41.2223 27.515 41.2328 27.0736C41.2537 26.3483 41.7143 25.8753 42.4469 25.8438C42.6563 25.8438 42.8656 25.8543 43.075 25.8543Z"
    fill="#2A38D8"
  />
  <path
    d="M24.7232 9.44967C24.7232 9.65193 24.7324 9.845 24.7232 10.0473C24.6956 10.6265 24.3647 11.0034 23.7947 11.0494C23.3534 11.0862 22.903 11.0862 22.4709 11.0586C21.8918 11.0126 21.5516 10.6357 21.5333 10.0656C21.5241 9.67951 21.5241 9.30257 21.5333 8.91643C21.5516 8.28206 21.9286 7.89592 22.5721 7.85915C22.8662 7.84076 23.1512 7.84995 23.4454 7.84995C24.3555 7.85915 24.7324 8.2269 24.7324 9.12789C24.7232 9.23821 24.7232 9.34854 24.7232 9.44967Z"
    fill="#2A38D8"
  />
  <path
    d="M23.1118 12.4656C23.3209 12.4656 23.5391 12.4551 23.7482 12.4656C24.2846 12.5078 24.6846 12.9615 24.7119 13.584C24.7392 14.0588 24.7392 14.5335 24.7119 15.0083C24.6846 15.6519 24.2755 16.1056 23.7119 16.1372C23.3209 16.1583 22.9209 16.1583 22.53 16.1372C21.9663 16.1056 21.5936 15.6941 21.5481 15.0505C21.5209 14.5968 21.5209 14.1326 21.539 13.6789C21.5754 12.8982 21.9572 12.4867 22.6482 12.4656C22.8118 12.4656 22.9573 12.4656 23.1118 12.4656Z"
    fill="#2A38D8"
  />
  <path
    d="M23.1093 20.7578C22.8998 20.7578 22.6811 20.7788 22.4716 20.7578C21.9705 20.7157 21.597 20.3478 21.5605 19.7697C21.515 19.2126 21.515 18.645 21.5605 18.0984C21.597 17.5413 21.9796 17.1419 22.4716 17.0998C22.8907 17.0683 23.3189 17.0683 23.747 17.0998C24.3028 17.1419 24.6945 17.6254 24.7219 18.2771C24.7401 18.645 24.731 19.0129 24.7219 19.3808C24.7219 20.3268 24.3483 20.7578 23.5284 20.7683C23.3826 20.7683 23.246 20.7683 23.1093 20.7578Z"
    fill="#2A38D8"
  />
  <path
    d="M23.1049 6.45374C22.904 6.45374 22.7123 6.46425 22.5114 6.45374C21.9636 6.41169 21.5893 6.03329 21.5528 5.40261C21.5163 4.89807 21.5163 4.38301 21.5528 3.87847C21.5984 3.22677 21.9728 2.81683 22.5297 2.7853C22.9223 2.76428 23.324 2.76428 23.7166 2.7853C24.3009 2.81683 24.6935 3.28984 24.7209 3.97307C24.7391 4.34097 24.73 4.70886 24.7209 5.07676C24.7117 6.01226 24.3283 6.44323 23.5157 6.45374C23.3788 6.46425 23.2418 6.46425 23.1049 6.45374C23.1049 6.46425 23.1049 6.46425 23.1049 6.45374Z"
    fill="#2A38D8"
  />
  <path
    d="M23.3807 26.7781C23.611 26.7781 23.8308 26.7676 24.061 26.7781C24.6891 26.8201 25.1392 27.2511 25.1706 27.8818C25.202 28.3758 25.1915 28.8593 25.1706 29.3534C25.1392 29.963 24.71 30.394 24.1029 30.436C23.6214 30.4676 23.1295 30.4676 22.6375 30.436C22.0199 30.4045 21.5699 29.942 21.5385 29.3113C21.5175 28.8698 21.528 28.4389 21.5385 27.9974C21.5594 27.2721 22.0199 26.7991 22.7526 26.7676C22.962 26.7676 23.1713 26.7781 23.3807 26.7781Z"
    fill="#2193F7"
  />
  <path
    d="M18.3416 0.0105119C18.5719 0.0105119 18.7917 6.41402e-07 19.022 0.0105119C19.65 0.0525571 20.1001 0.48352 20.1315 1.1142C20.1629 1.60823 20.1524 2.09175 20.1315 2.58578C20.1001 3.19543 19.6709 3.62639 19.0638 3.66844C18.5824 3.69997 18.0904 3.69997 17.5984 3.66844C16.9809 3.63691 16.5308 3.17441 16.4994 2.54373C16.4785 2.10226 16.4889 1.67129 16.4994 1.22982C16.5203 0.504543 16.9809 0.0315339 17.7136 0C17.9229 0 18.1323 6.41402e-07 18.3416 0.0105119Z"
    fill="#2193F7"
  />
  <path
    d="M18.3426 4.62575C18.5729 4.62575 18.7927 4.61523 19.023 4.62575C19.651 4.66779 20.1011 5.09875 20.1325 5.72943C20.1639 6.22346 20.1534 6.70698 20.1325 7.20101C20.1011 7.81067 19.6719 8.24163 19.0648 8.28367C18.5833 8.31521 18.0914 8.31521 17.5994 8.28367C16.9819 8.25214 16.5318 7.78964 16.5004 7.15897C16.4794 6.71749 16.4899 6.28653 16.5004 5.84506C16.5213 5.11978 16.9819 4.64677 17.7146 4.61523C17.9239 4.61523 18.1332 4.62575 18.3426 4.62575Z"
    fill="#2193F7"
  />
  <path
    d="M18.3426 9.7006C18.5729 9.7006 18.7927 9.69141 19.023 9.7006C19.651 9.73739 20.1011 10.1145 20.1325 10.6664C20.1639 11.0986 20.1534 11.5217 20.1325 11.954C20.1011 12.4874 19.6719 12.8645 19.0648 12.9013C18.5833 12.9289 18.0914 12.9289 17.5994 12.9013C16.9819 12.8737 16.5318 12.469 16.5004 11.9172C16.4794 11.5309 16.4899 11.1538 16.5004 10.7675C16.5213 10.1329 16.9819 9.719 17.7146 9.69141C17.9239 9.69141 18.1332 9.7006 18.3426 9.7006Z"
    fill="#2193F7"
  />
  <path
    d="M18.3416 14.3191C18.5719 14.3191 18.7917 14.3086 19.022 14.3191C19.65 14.3612 20.1001 14.7921 20.1315 15.4228C20.1629 15.9168 20.1524 16.4003 20.1315 16.8944C20.1001 17.504 19.6709 17.935 19.0638 17.977C18.5824 18.0086 18.0904 18.0086 17.5984 17.977C16.9809 17.9455 16.5308 17.483 16.4994 16.8523C16.4785 16.4109 16.4889 15.9799 16.4994 15.5384C16.5203 14.8131 16.9809 14.3401 17.7136 14.3086C17.9229 14.3086 18.1323 14.3191 18.3416 14.3191Z"
    fill="#2193F7"
  />
  <path
    d="M18.3426 18.9304C18.5729 18.9304 18.7927 18.9199 19.023 18.9304C19.651 18.9725 20.1011 19.4034 20.1325 20.0341C20.1639 20.5281 20.1534 21.0117 20.1325 21.5057C20.1011 22.1154 19.6719 22.5463 19.0648 22.5884C18.5833 22.6199 18.0914 22.6199 17.5994 22.5884C16.9819 22.5568 16.5318 22.0943 16.5004 21.4637C16.4794 21.0222 16.4899 20.5912 16.5004 20.1497C16.5213 19.4245 16.9819 18.9515 17.7146 18.9199C17.9239 18.9199 18.1332 18.9304 18.3426 18.9304Z"
    fill="#2193F7"
  />
  <path
    d="M28.4197 19.8549C28.65 19.8549 28.8698 19.8457 29.1001 19.8549C29.7281 19.8917 30.1782 20.2688 30.2096 20.8206C30.241 21.2529 30.2306 21.676 30.2096 22.1083C30.1782 22.6417 29.7491 23.0188 29.142 23.0556C28.6605 23.0832 28.1685 23.0832 27.6766 23.0556C27.059 23.028 26.6089 22.6233 26.5775 22.0715C26.5566 21.6852 26.567 21.3081 26.5775 20.9218C26.5985 20.2872 27.059 19.8733 27.7917 19.8457C28.0011 19.8457 28.2104 19.8549 28.4197 19.8549Z"
    fill="#2193F7"
  />
  <path
    d="M28.4197 15.2397C28.65 15.2397 28.8698 15.2305 29.1001 15.2397C29.7281 15.2765 30.1782 15.6535 30.2096 16.2054C30.241 16.6377 30.2306 17.0607 30.2096 17.493C30.1782 18.0265 29.7491 18.4036 29.142 18.4404C28.6605 18.4679 28.1685 18.4679 27.6766 18.4404C27.059 18.4128 26.6089 18.0081 26.5775 17.4562C26.5566 17.0699 26.567 16.6929 26.5775 16.3066C26.5985 15.6719 27.059 15.2581 27.7917 15.2305C28.0011 15.2305 28.2104 15.2397 28.4197 15.2397Z"
    fill="#2193F7"
  />
  <path
    d="M28.4197 10.6244C28.65 10.6244 28.8698 10.6152 29.1001 10.6244C29.7281 10.6612 30.1782 11.0383 30.2096 11.5902C30.241 12.0224 30.2306 12.4455 30.2096 12.8778C30.1782 13.4112 29.7491 13.7883 29.142 13.8251C28.6605 13.8527 28.1685 13.8527 27.6766 13.8251C27.059 13.7975 26.6089 13.3928 26.5775 12.841C26.5566 12.4547 26.567 12.0776 26.5775 11.6913C26.5985 11.0567 27.059 10.6428 27.7917 10.6152C28.0011 10.6152 28.2104 10.6244 28.4197 10.6244Z"
    fill="#2193F7"
  />
  <path
    d="M28.4188 6.01051C28.649 6.01051 28.8689 6 29.0991 6.01051C29.7272 6.05256 30.1772 6.48352 30.2086 7.1142C30.24 7.60823 30.2296 8.09175 30.2086 8.58578C30.1772 9.19543 29.7481 9.62639 29.141 9.66844C28.6595 9.69997 28.1676 9.69997 27.6756 9.66844C27.058 9.63691 26.6079 9.17441 26.5765 8.54373C26.5556 8.10226 26.5661 7.6713 26.5765 7.22982C26.5975 6.50454 27.058 6.03153 27.7907 6C28.0001 6 28.2094 6.01051 28.4188 6.01051Z"
    fill="#2193F7"
  />
  <path
    d="M28.4197 1.39332C28.65 1.39332 28.8698 1.38281 29.1001 1.39332C29.7281 1.43537 30.1782 1.86633 30.2096 2.49701C30.241 2.99104 30.2306 3.47456 30.2096 3.96859C30.1782 4.57824 29.7491 5.00921 29.142 5.05125C28.6605 5.08279 28.1685 5.08279 27.6766 5.05125C27.059 5.01972 26.6089 4.55722 26.5775 3.92654C26.5566 3.48507 26.567 3.05411 26.5775 2.61263C26.5985 1.88735 27.059 1.41435 27.7917 1.38281C28.0011 1.38281 28.2104 1.39332 28.4197 1.39332Z"
    fill="#2193F7"
  />
  <path
    d="M33.4578 6.93239C33.6881 6.93239 33.9079 6.92187 34.1382 6.93239C34.7662 6.97443 35.2163 7.40539 35.2477 8.03606C35.2791 8.53008 35.2686 9.0136 35.2477 9.50762C35.2163 10.1173 34.7872 10.5483 34.1801 10.5903C33.6986 10.6218 33.2066 10.6218 32.7147 10.5903C32.0971 10.5588 31.647 10.0962 31.6156 9.46558C31.5947 9.02411 31.6051 8.59315 31.6156 8.15168C31.6365 7.42641 32.0971 6.95341 32.8298 6.92188C33.0391 6.92188 33.2485 6.93239 33.4578 6.93239Z"
    fill="#2193F7"
  />
  <path
    d="M28.4188 28.6199C28.649 28.6199 28.8689 28.6094 29.0991 28.6199C29.7272 28.6619 30.1772 29.0929 30.2086 29.7236C30.24 30.2176 30.2296 30.7011 30.2086 31.1952C30.1772 31.8048 29.7481 32.2358 29.141 32.2778C28.6595 32.3093 28.1676 32.3093 27.6756 32.2778C27.058 32.2463 26.6079 31.7838 26.5765 31.1531C26.5556 30.7116 26.5661 30.2807 26.5765 29.8392C26.5975 29.1139 27.058 28.6409 27.7907 28.6094C28.0001 28.6094 28.2094 28.6094 28.4188 28.6199Z"
    fill="#2193F7"
  />
  <path
    d="M43.534 6.93239C43.7643 6.93239 43.9841 6.92187 44.2144 6.93239C44.8424 6.97443 45.2925 7.40539 45.3239 8.03606C45.3553 8.53008 45.3448 9.0136 45.3239 9.50762C45.2925 10.1173 44.8633 10.5483 44.2562 10.5903C43.7747 10.6218 43.2828 10.6218 42.7908 10.5903C42.1733 10.5588 41.7232 10.0962 41.6918 9.46558C41.6708 9.02411 41.6813 8.59315 41.6918 8.15168C41.7127 7.42641 42.1733 6.95341 42.906 6.92188C43.1153 6.92188 43.3246 6.93239 43.534 6.93239Z"
    fill="#2193F7"
  />
  <path
    d="M43.534 11.5476C43.7643 11.5476 43.9841 11.5371 44.2144 11.5476C44.8424 11.5897 45.2925 12.0206 45.3239 12.6513C45.3553 13.1453 45.3448 13.6289 45.3239 14.1229C45.2925 14.7325 44.8633 15.1635 44.2562 15.2055C43.7747 15.2371 43.2828 15.2371 42.7908 15.2055C42.1733 15.174 41.7232 14.7115 41.6918 14.0808C41.6708 13.6394 41.6813 13.2084 41.6918 12.7669C41.7127 12.0417 42.1733 11.5686 42.906 11.5371C43.1153 11.5371 43.3246 11.5371 43.534 11.5476Z"
    fill="#2193F7"
  />
  <path
    d="M43.534 30.4675C43.7643 30.4675 43.9841 30.457 44.2144 30.4675C44.8424 30.5096 45.2925 30.9406 45.3239 31.5712C45.3553 32.0653 45.3448 32.5488 45.3239 33.0428C45.2925 33.6525 44.8633 34.0834 44.2562 34.1255C43.7747 34.157 43.2828 34.157 42.7908 34.1255C42.1733 34.0939 41.7232 33.6314 41.6918 33.0008C41.6708 32.5593 41.6813 32.1283 41.6918 31.6869C41.7127 30.9616 42.1733 30.4886 42.906 30.457C43.1153 30.457 43.3246 30.4675 43.534 30.4675Z"
    fill="#2193F7"
  />
  <path
    d="M38.4949 18.4714C38.7252 18.4714 38.945 18.4609 39.1753 18.4714C39.8033 18.5135 40.2534 18.9445 40.2848 19.5751C40.3162 20.0692 40.3058 20.5527 40.2848 21.0467C40.2534 21.6564 39.8243 22.0873 39.2172 22.1294C38.7357 22.1609 38.2437 22.1609 37.7518 22.1294C37.1342 22.0978 36.6841 21.6353 36.6527 21.0047C36.6318 20.5632 36.6422 20.1322 36.6527 19.6908C36.6736 18.9655 37.1342 18.4925 37.8669 18.4609C38.0763 18.4609 38.2856 18.4714 38.4949 18.4714Z"
    fill="#2193F7"
  />
  <path
    d="M33.4569 11.5476C33.6871 11.5476 33.9069 11.5371 34.1372 11.5476C34.7652 11.5897 35.2153 12.0206 35.2467 12.6513C35.2781 13.1453 35.2677 13.6289 35.2467 14.1229C35.2153 14.7325 34.7862 15.1635 34.1791 15.2055C33.6976 15.2371 33.2056 15.2371 32.7137 15.2055C32.0961 15.174 31.646 14.7115 31.6146 14.0808C31.5937 13.6394 31.6042 13.2084 31.6146 12.7669C31.6356 12.0417 32.0961 11.5686 32.8288 11.5371C33.0382 11.5371 33.2475 11.5371 33.4569 11.5476Z"
    fill="#2193F7"
  />
  <path
    d="M38.2436 12.4591C39.129 12.4591 39.8466 11.7359 39.8466 10.8438C39.8466 9.9517 39.129 9.22852 38.2436 9.22852C37.3584 9.22852 36.6406 9.9517 36.6406 10.8438C36.6406 11.7359 37.3584 12.4591 38.2436 12.4591Z"
    fill="#2193F7"
  />
  <path
    d="M14.5084 13.3848V15.086C14.5084 15.086 6.93289 21.0813 7.666 26.595C8.64349 33.9636 18.1943 34.312 27.0222 32.2828V33.8918C27.0222 33.8918 11.8712 36.8229 7.42163 30.8685C2.98224 24.904 9.57006 17.0947 14.5084 13.3848Z"
    fill="#2193F7"
  />
  <path
    d="M14.6978 8.30859V9.91787C14.6978 9.91787 0.435302 21.8017 2.83613 30.0234C5.23696 38.2452 18.3398 41.3399 38.4722 32.3858V33.5824C38.4722 33.5824 10.4048 47.0446 1.26949 33.995C1.26949 34.0053 -6.10594 24.3703 14.6978 8.30859Z"
    fill="#2193F7"
  />
  <path
    d="M32.5195 3.1573V5.58922C32.5195 5.58922 43.3976 2.49032 49.7625 7.31311C56.1275 12.1359 47.8197 20.3244 47.8197 20.3244V24.4597C47.8197 24.4597 63.7269 15.163 59.1835 6.37933C54.0834 -3.47148 32.7624 2.99312 32.5195 3.1573Z"
    fill="#2193F7"
  />
</svg>

    <ul
      id="menu"
      class="text-sm font-medium flex flex-col items-center gap-12 relative w-[12vw] max-w-[168px]"
    >
      <div
        id="indicator"
        class="absolute bg-[#304A79] w-[4px] h-6 right-0 transition-all duration-500 ease-in-out top-0"
      ></div>
      <li
        data-content="faq"
        class="text-neutral-900 flex flex-col items-center gap-4 cursor-pointer active"
      >
        <svg
          id="info"
          class="fill-white stroke-[#85888B]"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle cx="12" cy="12" r="10" stroke-width="1.5" />
          <path
            d="M10 9C10 7.89543 10.8954 7 12 7C13.1046 7 14 7.89543 14 9C14 9.39815 13.8837 9.76913 13.6831 10.0808C13.0854 11.0097 12 11.8954 12 13V13.5"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path
            d="M11.992 17H12.001"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <svg
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          class="hidden"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle
            cx="12"
            cy="12"
            r="10"
            fill="#3D5F9C"
            stroke="#3D5F9C"
            stroke-width="1.5"
          />
          <path
            d="M10 9C10 7.89543 10.8954 7 12 7C13.1046 7 14 7.89543 14 9C14 9.39815 13.8837 9.76913 13.6831 10.0808C13.0854 11.0097 12 11.8954 12 13V13.5"
            stroke="#ECF0F7"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path
            d="M11.992 17H12.001"
            stroke="#ECF0F7"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <p>سوالات متداول</p>
      </li>
      <!--      <li-->
      <!--        data-content="support"-->
      <!--        class="text-neutral-500 flex flex-col items-center gap-4 cursor-pointer"-->
      <!--      >-->
      <!--        <svg-->
      <!--          id="head_phone"-->
      <!--          width="24"-->
      <!--          height="24"-->
      <!--          viewBox="0 0 24 24"-->
      <!--          fill="none"-->
      <!--          class="stroke-[#85888B] fill-white"-->
      <!--          xmlns="http://www.w3.org/2000/svg"-->
      <!--        >-->
      <!--          <path-->
      <!--            d="M17 10.8045C17 10.4588 17 10.286 17.052 10.132C17.2032 9.68444 17.6018 9.51076 18.0011 9.32888C18.45 9.12442 18.6744 9.02219 18.8968 9.0042C19.1493 8.98378 19.4022 9.03818 19.618 9.15929C19.9041 9.31984 20.1036 9.62493 20.3079 9.87302C21.2513 11.0188 21.7229 11.5918 21.8955 12.2236C22.0348 12.7334 22.0348 13.2666 21.8955 13.7764C21.6438 14.6979 20.8485 15.4704 20.2598 16.1854C19.9587 16.5511 19.8081 16.734 19.618 16.8407C19.4022 16.9618 19.1493 17.0162 18.8968 16.9958C18.6744 16.9778 18.45 16.8756 18.0011 16.6711C17.6018 16.4892 17.2032 16.3156 17.052 15.868C17 15.714 17 15.5412 17 15.1955V10.8045Z"-->
      <!--            stroke-width="1.5"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M7 10.8046C7 10.3694 6.98778 9.97821 6.63591 9.6722C6.50793 9.5609 6.33825 9.48361 5.99891 9.32905C5.55001 9.12458 5.32556 9.02235 5.10316 9.00436C4.43591 8.9504 4.07692 9.40581 3.69213 9.87318C2.74875 11.019 2.27706 11.5919 2.10446 12.2237C1.96518 12.7336 1.96518 13.2668 2.10446 13.7766C2.3562 14.6981 3.15152 15.4705 3.74021 16.1856C4.11129 16.6363 4.46577 17.0475 5.10316 16.996C5.32556 16.978 5.55001 16.8757 5.99891 16.6713C6.33825 16.5167 6.50793 16.4394 6.63591 16.3281C6.98778 16.0221 7 15.631 7 15.1957V10.8046Z"-->
      <!--            stroke-width="1.5"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M5 9C5 5.68629 8.13401 3 12 3C15.866 3 19 5.68629 19 9"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="square"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M19 17V17.8C19 19.5673 17.2091 21 15 21H13"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--        </svg>-->
      <!--        <svg-->
      <!--          width="22"-->
      <!--          height="20"-->
      <!--          viewBox="0 0 22 20"-->
      <!--          fill="none"-->
      <!--          class="hidden"-->
      <!--          xmlns="http://www.w3.org/2000/svg"-->
      <!--        >-->
      <!--          <path-->
      <!--            d="M16 8.80445C16 8.45883 16 8.28603 16.052 8.13197C16.2032 7.68444 16.6018 7.51076 17.0011 7.32888C17.45 7.12442 17.6744 7.02219 17.8968 7.0042C18.1493 6.98378 18.4022 7.03818 18.618 7.15929C18.9041 7.31984 19.1036 7.62493 19.3079 7.87302C20.2513 9.01885 20.7229 9.59176 20.8955 10.2236C21.0348 10.7334 21.0348 11.2666 20.8955 11.7764C20.6438 12.6979 19.8485 13.4704 19.2598 14.1854C18.9587 14.5511 18.8081 14.734 18.618 14.8407C18.4022 14.9618 18.1493 15.0162 17.8968 14.9958C17.6744 14.9778 17.45 14.8756 17.0011 14.6711C16.6018 14.4892 16.2032 14.3156 16.052 13.868C16 13.714 16 13.5412 16 13.1955V8.80445Z"-->
      <!--            fill="#3D5F9C"-->
      <!--            stroke="#3D5F9C"-->
      <!--            stroke-width="1.5"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M6 8.80462C6 8.36937 5.98778 7.97821 5.63591 7.6722C5.50793 7.5609 5.33825 7.48361 4.99891 7.32905C4.55001 7.12458 4.32556 7.02235 4.10316 7.00436C3.43591 6.9504 3.07692 7.40581 2.69213 7.87318C1.74875 9.01901 1.27706 9.59193 1.10446 10.2237C0.965178 10.7336 0.965178 11.2668 1.10446 11.7766C1.3562 12.6981 2.15152 13.4705 2.74021 14.1856C3.11129 14.6363 3.46577 15.0475 4.10316 14.996C4.32556 14.978 4.55001 14.8757 4.99891 14.6713C5.33825 14.5167 5.50793 14.4394 5.63591 14.3281C5.98778 14.0221 6 13.631 6 13.1957V8.80462Z"-->
      <!--            fill="#3D5F9C"-->
      <!--            stroke="#3D5F9C"-->
      <!--            stroke-width="1.5"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M4 7C4 3.68629 7.13401 1 11 1C14.866 1 18 3.68629 18 7"-->
      <!--            stroke="#3D5F9C"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="square"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M18 15V15.8C18 17.5673 16.2091 19 14 19H12"-->
      <!--            stroke="#3D5F9C"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--        </svg>-->
      <!--        <p>پشتیبانی 24/7</p>-->
      <!--      </li>-->
      <li
        data-content="download"
        id="download-app-tab"
        class="text-neutral-500 flex flex-col items-center gap-4 cursor-pointer"
      >
        <svg
          class="stroke-[#85888B]"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M12 14.5L12 4.5M12 14.5C11.2998 14.5 9.99153 12.5057 9.5 12M12 14.5C12.7002 14.5 14.0085 12.5057 14.5 12"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M20 16.5C20 18.982 19.482 19.5 17 19.5H7C4.518 19.5 4 18.982 4 16.5"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <svg
          width="18"
          height="18"
          viewBox="0 0 18 18"
          fill="none"
          class="hidden"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M9 11.5L9 1.5M9 11.5C8.29977 11.5 6.99153 9.5057 6.5 9M9 11.5C9.70023 11.5 11.0085 9.5057 11.5 9"
            stroke="#3D5F9C"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M17 13.5C17 15.982 16.482 16.5 14 16.5H4C1.518 16.5 1 15.982 1 13.5"
            stroke="#3D5F9C"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>

        <p>دانلود اپلیکیشن</p>
      </li>
      <li
        data-content="privacy"
        class="text-neutral-500 flex flex-col items-center gap-4 cursor-pointer"
      >
        <svg
          class="stroke-[#85888B]"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M11.4706 22C7.47751 22 5.48098 22 4.24049 20.8284C3 19.6569 3 17.7712 3 14L3 10C3 6.22876 3 4.34315 4.24049 3.17157C5.48098 2 7.47752 2 11.4706 2L12.5294 2C16.5225 2 18.519 2 19.7595 3.17157C21 4.34315 21 6.22876 21 10M11.5 22H13"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path d="M8 7H16" stroke-width="1.5" stroke-linecap="round" />
          <path d="M8 12H13" stroke-width="1.5" stroke-linecap="round" />
          <path
            d="M17.5 18.5896L16.5978 21.7428C16.5572 21.9011 16.7139 22.0385 16.8659 21.9778L18.8514 21.1849C18.9468 21.1468 19.0532 21.1468 19.1486 21.1849L21.1531 21.9854C21.3014 22.0446 21.456 21.9149 21.4231 21.7589L20.6589 18.4911M22 15.9951C22 14.341 20.6569 13 19 13C17.3431 13 16 14.341 16 15.9951C16 17.6493 17.3431 18.9902 19 18.9902C20.6569 18.9902 22 17.6493 22 15.9951Z"
            stroke-width="1.5"
            stroke-linejoin="round"
          />
        </svg>
        <svg
          width="21"
          height="22"
          viewBox="0 0 21 22"
          fill="none"
          class="hidden"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M9.47059 21C5.47751 21 3.48098 21 2.24049 19.8284C1 18.6569 1 16.7712 1 13L1 9C1 5.22876 1 3.34315 2.24049 2.17157C3.48098 1 5.47752 1 9.47059 1L10.5294 1C14.5225 1 16.519 1 17.7595 2.17157C19 3.34315 19 5.22876 19 9M9.5 21H11"
            stroke="#3D5F9C"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path
            d="M6 6H14"
            stroke="#3D5F9C"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path
            d="M6 11H11"
            stroke="#3D5F9C"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <path
            d="M20 14.9951C20 16.6493 18.6569 17.9902 17 17.9902C15.3431 17.9902 14 16.6493 14 14.9951C14 13.341 15.3431 12 17 12C18.6569 12 20 13.341 20 14.9951Z"
            fill="#3D5F9C"
          />
          <path
            d="M15.5 17.5896L14.5978 20.7428C14.5572 20.9011 14.7139 21.0385 14.8659 20.9778L16.8514 20.1849C16.9468 20.1468 17.0532 20.1468 17.1486 20.1849L19.1531 20.9854C19.3014 21.0446 19.456 20.9149 19.4231 20.7589L18.6589 17.4911M20 14.9951C20 13.341 18.6569 12 17 12C15.3431 12 14 13.341 14 14.9951C14 16.6493 15.3431 17.9902 17 17.9902C18.6569 17.9902 20 16.6493 20 14.9951Z"
            stroke="#3D5F9C"
            stroke-width="1.5"
            stroke-linejoin="round"
          />
        </svg>
        <p>حریم خصوصی</p>
      </li>
      <!--      <li-->
      <!--        data-content="help"-->
      <!--        class="text-neutral-500 flex flex-col items-center gap-4 cursor-pointer"-->
      <!--      >-->
      <!--        <svg-->
      <!--          class="stroke-[#85888B]"-->
      <!--          width="24"-->
      <!--          height="24"-->
      <!--          viewBox="0 0 24 24"-->
      <!--          fill="none"-->
      <!--          xmlns="http://www.w3.org/2000/svg"-->
      <!--        >-->
      <!--          <circle cx="12" cy="12" r="10" stroke-width="1.5" />-->
      <!--          <path-->
      <!--            d="M11.992 15H12.001"-->
      <!--            stroke-width="2"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M12 12L12 8"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--        </svg>-->
      <!--        <svg-->
      <!--          width="20"-->
      <!--          height="20"-->
      <!--          viewBox="0 0 20 20"-->
      <!--          fill="none"-->
      <!--          class="hidden"-->
      <!--          xmlns="http://www.w3.org/2000/svg"-->
      <!--        >-->
      <!--          <circle cx="10" cy="10" r="10" fill="#3D5F9C" />-->
      <!--          <path-->
      <!--            d="M9.99199 13H10.001"-->
      <!--            stroke="white"-->
      <!--            stroke-width="2"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M10 10L10 6"-->
      <!--            stroke="white"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--        </svg>-->

      <!--        <p>راهنما سامانه</p>-->
      <!--      </li>-->
      <!--      <li-->
      <!--        data-content="language"-->
      <!--        class="text-neutral-500 flex flex-col items-center gap-4 cursor-pointer"-->
      <!--      >-->
      <!--        <svg-->
      <!--          class="stroke-[#85888B]"-->
      <!--          width="24"-->
      <!--          height="24"-->
      <!--          viewBox="0 0 24 24"-->
      <!--          fill="none"-->
      <!--          xmlns="http://www.w3.org/2000/svg"-->
      <!--        >-->
      <!--          <path-->
      <!--            d="M2 12C2 13.0519 2.18046 14.0617 2.51212 15M13.0137 9H21.5015M11 15H2.51212M21.5015 9C20.266 5.50442 16.9323 3 13.0137 3C14.6146 3 15.9226 6.76201 16.0091 11.5M21.5015 9C21.7803 9.78867 21.9522 10.6278 22 11.5M2.51212 15C3.74763 18.4956 7.08134 21 11 21C9.45582 21 8.18412 17.5 8.01831 13"-->
      <!--            stroke-width="1.5"-->
      <!--            stroke-linecap="round"-->
      <!--            stroke-linejoin="round"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M2 5.29734C2 4.19897 2 3.64979 2.18738 3.22389C2.3861 2.77223 2.72861 2.40921 3.15476 2.1986C3.55661 2 4.07478 2 5.11111 2H6C7.88562 2 8.82843 2 9.41421 2.62085C10 3.2417 10 4.24095 10 6.23944V8.49851C10 9.37026 10 9.80613 9.73593 9.95592C9.47186 10.1057 9.12967 9.86392 8.4453 9.38036L8.34103 9.30669C7.84086 8.95329 7.59078 8.77658 7.30735 8.68563C7.02392 8.59468 6.72336 8.59468 6.12223 8.59468H5.11111C4.07478 8.59468 3.55661 8.59468 3.15476 8.39608C2.72861 8.18547 2.3861 7.82245 2.18738 7.37079C2 6.94489 2 6.39571 2 5.29734Z"-->
      <!--            stroke-width="1.5"-->
      <!--          />-->
      <!--          <path-->
      <!--            d="M22 17.2973C22 16.199 22 15.6498 21.8126 15.2239C21.6139 14.7722 21.2714 14.4092 20.8452 14.1986C20.4434 14 19.9252 14 18.8889 14H18C16.1144 14 15.1716 14 14.5858 14.6209C14 15.2417 14 16.2409 14 18.2394V20.4985C14 21.3703 14 21.8061 14.2641 21.9559C14.5281 22.1057 14.8703 21.8639 15.5547 21.3804L15.659 21.3067C16.1591 20.9533 16.4092 20.7766 16.6926 20.6856C16.9761 20.5947 17.2766 20.5947 17.8778 20.5947H18.8889C19.9252 20.5947 20.4434 20.5947 20.8452 20.3961C21.2714 20.1855 21.6139 19.8225 21.8126 19.3708C22 18.9449 22 18.3957 22 17.2973Z"-->
      <!--            stroke-width="1.5"-->
      <!--          />-->
      <!--        </svg>-->
      <!--        <p>EN / FA</p>-->
      <!--      </li>-->
    </ul>
  </div>
  <div
    id="content-wrapper"
    class="max-w-[660px] w-1/2 min-w-[490px] xl:min-w-[660px] overflow-y-auto pb-10 bg-[#304A79] h-screen text-white lg:block hidden z-10"
  >
    <div id="faq-content">
      <h1
        class="flex items-center justify-center mt-[106px] text-[32px] font-bold text-neutral-50 mb-6"
      >
        ســــوالات متــــداول
      </h1>
      <p
        class="flex items-center justify-center font-medium text-lg text-neutral-50 opacity-60"
      >
        هر آنچه در مورد پنجره ملی و خدمات آن میخواهید بدانید:
      </p>
      <!-- tabs -->
      <section
        id="tabs-wrapper"
        class="flex items-center justify-center relative w-fit mx-auto gap-[60px] mt-9"
      >
        <button
          type="button"
          class="text-white opacity-100 cursor-pointer tab_btn"
        >
          عمومی
        </button>
        <button
          type="button"
          class="text-white opacity-70 cursor-pointer tab_btn"
        >
          تخصصی
        </button>
        <button
          type="button"
          class="text-white opacity-70 cursor-pointer tab_btn"
        >
          وزارت خانه‌ها
        </button>
        <button
          type="button"
          class="text-white opacity-70 cursor-pointer tab_btn"
        >
          سایر
        </button>
        <div
          id="line_indicator"
          class="h-0.5 bg-white absolute -bottom-1 left-0 transition-all duration-500 ease-in-out rounded-full"
        ></div>
      </section>
      <section class="px-[72px] mt-[78px]">
        <!-- faq -->
        <div
          class="faq-tab-content tab-content flex flex-col items-start block"
          data-category="614193452171d0294fea7f10"
        ></div>
        <!-- ................................................... -->
        <div
          class="faq-tab-content tab-content flex flex-col items-start hidden"
          data-category="646df7a1a8a9c325f21b957c"
        ></div>
        <div
          class="faq-tab-content tab-content flex flex-col items-start hidden"
          data-category="68be7493bf4e0571d7163b62"
        ></div>
        <div
          class="faq-tab-content tab-content flex flex-col items-start hidden"
          data-category="68be75035515b84e6e96da1c"
        ></div>
      </section>
    </div>
    <!-- ...........................24/7 support............................... -->
    <!--<div id="support-content">
      <section
        class="flex flex-col items-center justify-center mt-[106px] gap-6 text-white"
      >
        <p class="text-[32px] font-bold">پــــشتیــــبانی 24/7</p>
        <p class="text-lg font-medium opacity-60 text-center px-[92px]">
          کاربران گرامی می‌توانند در هر ساعت از شبانه‌روز نسبت به ثبت درخواست
          اقدام نموده و پاسخ کارشناسان را دریافت نمایند.
        </p>
      </section>
      <section
        class="px-[81px] mx-10 mt-11 before:bg-[linear-gradient(360deg,rgba(61,95,156,0)0%,#A1B4D6_100%)] bg-[linear-gradient(360deg,#304A79_8.85%,rgba(68,105,173,0.35)100%)] rounded-3xl linearBorder before:rounded-3xl pt-12 pb-11"
      >
        <form action="" class="space-y-6 flex flex-col items-center">
          <input
            class="bg-transparent outline-none border-b border-[#7C96C6] placeholder:text-[#7C96C6] text-sm font-medium focus:outline-none w-full pb-2 pr-2"
            placeholder="نام و نام خانوادگی"
          />
          <input
            class="bg-transparent outline-none border-b border-[#7C96C6] placeholder:text-[#7C96C6] text-sm font-medium focus:outline-none w-full pb-2 pr-2"
            placeholder="موضوع"
          />
          <div class="custom-select w-full text-[#7C96C6]">
            <select class="select-hidden" name="support-type">
              <option value="" disabled selected>نوع پشتیبانی</option>
              <option value="1">گزینه ۱</option>
              <option value="2">گزینه ۲</option>
              <option value="3">گزینه ۳</option>
              <option value="4">گزینه ۴</option>
            </select>
            <div
              class="select-selected flex justify-between items-center bg-transparent text-sm border-0 border-b border-[#7C96C6] text-[#7C96C6] pr-2 pb-2 cursor-pointer outline-none"
            >
              نوع پشتیبانی
              <svg
                class="w-6 h-6 fill-none stroke-[#9CBDFF] stroke-2"
                viewBox="0 0 24 24"
              >
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            </div>
            <div
              class="select-items hidden absolute top-full right-0 left-0 rounded-md overflow-hidden z-10"
            >
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="1"
              >
                گزینه ۱
              </div>
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="2"
              >
                گزینه ۲
              </div>
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="3"
              >
                گزینه ۳
              </div>
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="4"
              >
                گزینه ۴
              </div>
            </div>
          </div>
          <div class="custom-select w-full">
            <select class="select-hidden" name="priority-level">
              <option value="" disabled selected>درجه اهمیت</option>
              <option value="1">گزینه ۱</option>
              <option value="2">گزینه ۲</option>
              <option value="3">گزینه ۳</option>
              <option value="4">گزینه ۴</option>
            </select>
            <div
              class="select-selected flex justify-between items-center bg-transparent text-sm border-0 border-b border-[#7C96C6] text-[#7C96C6] pr-2 pb-2 cursor-pointer outline-none"
            >
              درجه اهمیت
              <svg
                class="w-6 h-6 fill-none stroke-[#9CBDFF] stroke-2"
                viewBox="0 0 24 24"
              >
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            </div>
            <div
              class="select-items hidden absolute top-full right-0 left-0 rounded-md overflow-hidden z-10"
            >
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="1"
              >
                گزینه ۱
              </div>
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="2"
              >
                گزینه ۲
              </div>
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="3"
              >
                گزینه ۳
              </div>
              <div
                class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                data-value="4"
              >
                گزینه ۴
              </div>
            </div>
          </div>
          <textarea
            name=""
            placeholder="توضیحات"
            class="w-full bg-[#ECF4FF1A] p-2 h-[127px] placeholder:text-[#7C96C6] text-sm outline-none focus:outline-none border-none"
          ></textarea>
          <button
            type="button"
            class="mt-16 w-full bg-white shadow-[4px_4px_8px_0px_#00000040] text-[#3D5F9C] py-2 flex items-center justify-center rounded-[10px] font-medium"
          >
            ارسال کد
          </button>
        </form>
      </section>
      <section class="px-10 mt-8 grid grid-cols-3 justify-between gap-4">
        <div
          class="before:bg-[linear-gradient(193.24deg,rgba(61,95,156,0)41.8%,#A1B4D6_101.77%)] linearBorder drop-shadow-[-10px_-5px_32px_0px_#00000066] rounded-2xl before:rounded-2xl p-4 bg-[linear-gradient(201.67deg,#304A79_34.36%,#4469AD_95.69%)]"
        >
          <div
            class="w-[34px] h-[34px] bg-[#4469AD] flex items-center justify-center rounded-full"
          >
            <svg th:replace="fragments/svgs :: telephone"></svg>
          </div>
          <p class="text-sm font-normal text-[#DDEAFF] mt-[17px]">تلفن</p>
          <p class="text-base font-normal text-white mt-1">021-91091516</p>
        </div>
        <div
          class="before:bg-[linear-gradient(193.24deg,rgba(61,95,156,0)41.8%,#A1B4D6_101.77%)] linearBorder drop-shadow-[-10px_-5px_32px_0px_#00000066] rounded-2xl before:rounded-2xl p-4 bg-[linear-gradient(201.67deg,#304A79_34.36%,#4469AD_95.69%)]"
        >
          <div
            class="w-[34px] h-[34px] bg-[#4469AD] flex items-center justify-center rounded-full"
          >
            <svg th:replace="fragments/svgs :: mail"></svg>
          </div>
          <p class="text-sm font-normal text-[#DDEAFF] mt-[17px]">ایمیل</p>
          <p class="text-base font-normal text-white mt-1">example@gmail.com</p>
        </div>
        <div
          class="before:bg-[linear-gradient(193.24deg,rgba(61,95,156,0)41.8%,#A1B4D6_101.77%)] linearBorder drop-shadow-[-10px_-5px_32px_0px_#00000066] rounded-2xl before:rounded-2xl p-4 bg-[linear-gradient(201.67deg,#304A79_34.36%,#4469AD_95.69%)]"
        >
          <div
            class="w-[34px] h-[34px] bg-[#4469AD] flex items-center justify-center rounded-full"
          >
            <svg th:replace="fragments/svgs :: fax"></svg>
          </div>
          <p class="text-sm font-normal text-[#DDEAFF] mt-[17px]">فکس</p>
          <p class="text-base font-normal text-white mt-1">021-91091516</p>
        </div>
      </section>
    </div> -->
    <!-- ..........................Download the application..................... -->
    <div id="download-content">
      <p
        class="text-[32px] font-bold flex items-center justify-center mt-[106px]"
      >
        دانــــلود اپلیــــکیشن
      </p>
      <p class="opacity-60 font-medium flex items-center justify-center">
        هر آنچه در مورد پنجره ملی و خدمات آن میخواهید بدانید:
      </p>
      <section class="mt-[72px] flex flex-col items-center gap-4">
        <a
          class="linearBorder justify-between min-w-[60%] h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
          href="https://file.my.gov.ir/static/app/MyGov.apk"
        >
          <div class="flex items-center gap-5 w-full">
            <svg
              width="32"
              height="36"
              viewBox="0 0 18 18"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M9 10.875L9 3.375M9 10.875C8.47483 10.875 7.49365 9.37927 7.125 9M9 10.875C9.52517 10.875 10.5064 9.37927 10.875 9"
                stroke="white"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
              <path
                d="M15 12.375C15 14.2365 14.6115 14.625 12.75 14.625H5.25C3.3885 14.625 3 14.2365 3 12.375"
                stroke="white"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>

            <div class="mt-2">
              <p class="text-sm">دریافت نسخه اندروید از</p>
              <p class="text-[23px] font-extrabold mt-1">دانلود مستقیم</p>
            </div>
          </div>
          <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
    fill="white"
  />
</svg>
        </a>
        <div class="h-[1px] bg-[#A1B4D680] min-w-[60%]"></div>
        <a
          class="linearBorder justify-between min-w-[60%] h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
          href="https://cafebazaar.ir/app/ir.gov.mygov/"
        >
          <div class="flex items-center gap-5 w-full">
            <svg
              width="32"
              height="36"
              viewBox="0 0 40 40"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M37.0207 14.7852C36.8435 15.0691 36.5766 15.3951 36.1752 15.7066C35.1133 16.5311 33.5512 16.9461 31.5197 16.9461C31.0337 16.9461 30.521 16.9224 29.9817 16.8748C23.0416 16.2624 16.6581 16.1791 11.5211 16.6338C7.26225 17.0108 4.11764 16.5147 2.66649 15.2371C2.57141 15.1533 2.48443 15.0663 2.40413 14.9764C2.39387 15.1914 2.40212 15.4161 2.43076 15.6514C2.56132 16.7238 2.69648 17.798 2.89345 18.8596C3.43674 21.7876 4.12326 24.6826 5.00854 27.5282C5.77108 29.9792 6.56081 32.4206 7.83888 34.6656C8.2942 35.4653 8.77654 36.2643 9.59851 36.7487C10.1999 37.1031 10.818 37.4413 11.4599 37.7127C14.0801 38.8206 16.863 39.0557 19.8271 39.1992C20.5632 39.1381 21.4615 39.0581 22.3607 38.9899C24.1892 38.8513 25.9659 38.4766 27.6935 37.8599C28.7314 37.4893 29.5615 36.8686 30.2226 35.9779C31.4036 34.3868 32.2264 32.6084 32.9761 30.7963C34.5585 26.9715 35.7433 23.0203 36.6121 18.9743C36.8856 17.7005 37.1443 16.4203 37.055 15.1072C37.0476 14.9981 37.036 14.8909 37.0207 14.7852Z"
                fill="white"
              />
              <path
                d="M35.9059 12.7396C34.7856 11.8063 33.4718 11.2313 32.0908 10.8727C30.3303 10.4156 28.534 10.0958 26.7515 9.72513C26.448 9.66201 26.2688 9.53706 26.1351 9.22194C25.4024 7.49481 24.676 5.76225 23.8566 4.07568C23.3603 3.05407 22.6893 2.12042 21.748 1.42848C21.2029 1.02781 20.5845 0.750019 19.9169 0.808573C19.1244 0.878103 18.4392 1.25361 17.847 1.81238C16.8226 2.77895 16.0672 3.93852 15.45 5.1802C14.7916 6.50497 14.2313 7.87826 13.6081 9.22105C13.5288 9.39205 13.3464 9.60093 13.1813 9.63248C12.1198 9.83527 11.0418 9.95507 9.98427 10.1748C8.00596 10.5857 6.06507 11.1189 4.29927 12.1546C3.41769 12.6717 2.80515 13.3303 2.54297 14.1722C2.66145 14.4016 2.82964 14.6133 3.04779 14.8053C4.35987 15.9606 7.43009 16.418 11.4706 16.0605C11.9549 16.0176 12.4505 15.9797 12.9561 15.9464C13.1818 12.901 15.0074 9.59839 15.0951 9.44057L16.0696 9.35364C16.0493 9.39031 14.0408 13.0517 13.7868 15.8963C17.3992 15.699 21.5038 15.733 25.9059 15.9958C25.7278 13.111 23.8462 9.41833 23.8462 9.41833L24.81 9.41692C24.8927 9.57778 26.5818 12.9734 26.7337 16.0474C27.8173 16.1186 28.9178 16.2032 30.0326 16.3015C31.8284 16.4601 34.2835 16.412 35.781 15.2836C36.5732 14.6867 36.7534 14.0492 36.7835 13.9161C36.5965 13.4736 36.3108 13.0771 35.9059 12.7396ZM16.1858 9.31248C16.1858 9.29395 16.1685 9.23537 16.1883 9.19508C16.8182 7.91445 17.489 6.66361 18.529 5.64991C19.6567 4.55076 20.6214 4.71238 21.5902 5.72275C22.5077 6.67948 23.103 7.86505 23.6495 9.06586C23.6863 9.14686 23.6861 9.24476 23.6989 9.31248C21.187 9.31248 18.6958 9.31248 16.1858 9.31248ZM25.143 9.58839C25.0127 9.44612 24.8375 9.32495 24.7593 9.15825C24.1703 7.90384 23.5844 6.64683 22.7088 5.55971C22.4101 5.18889 22.077 4.82649 21.6974 4.54375C20.453 3.61683 19.0625 3.7726 17.8741 4.91372C17.0429 5.71185 16.4346 6.66948 15.884 7.66718C15.6132 8.15795 15.3702 8.66465 15.13 9.17142C15.004 9.43707 14.8411 9.55872 14.4908 9.4024C14.7257 8.83622 14.9511 8.27003 15.1943 7.71155C15.8248 6.26354 16.5457 4.86383 17.4676 3.57623C17.9328 2.92639 18.4519 2.32881 19.1727 1.94294C19.9291 1.53805 20.623 1.66782 21.2756 2.17971C22.0568 2.79248 22.6418 3.58174 23.0647 4.45412C23.8402 6.05423 24.5288 7.69652 25.2512 9.32232C25.2716 9.36803 25.2644 9.42595 25.2701 9.47818C25.2277 9.51491 25.1853 9.55164 25.143 9.58839Z"
                fill="white"
              />
            </svg>

            <div class="mt-2">
              <p class="text-sm">دریافت نسخه اندروید از</p>
              <svg
                width="48"
                height="28"
                viewBox="0 0 48 28"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                class="mx-auto"
              >
                <path
                  d="M1.36795e-05 17.5018C0.0449225 17.3814 0.0954738 17.2624 0.133868 17.1404C0.474276 16.0578 1.52036 15.3084 2.76522 15.2957C4.2772 15.2802 5.7895 15.2912 7.30166 15.2899C8.14703 15.2892 8.49442 14.9794 8.49468 14.2252C8.49545 12.0433 8.48955 9.86143 8.49735 7.67959C8.50225 6.30926 9.54951 5.2622 11.0464 5.10692C12.5617 4.94974 14.0157 6.03825 14.0974 7.41533C14.1334 8.02144 14.1129 8.63041 14.1131 9.23808C14.114 11.2605 14.1288 13.2831 14.1081 15.3054C14.0827 17.7845 12.1706 19.8258 9.47691 20.2712C9.11029 20.3318 8.73109 20.3514 8.3575 20.3531C6.59151 20.3614 4.82544 20.3577 3.05941 20.3566C1.45566 20.3556 0.44995 19.6424 0.0627207 18.2363C0.0496796 18.189 0.0212531 18.1451 0 18.0996C1.19534e-05 17.9003 1.36795e-05 17.7011 1.36795e-05 17.5018Z"
                  fill="white"
                />
                <path
                  d="M33.918 11.5277C33.918 9.36596 33.8993 7.20403 33.9237 5.04251C33.9443 3.22128 36.0131 2.02554 37.8524 2.76898C38.9526 3.21369 39.5171 4.02083 39.5284 5.11059C39.5441 6.62468 39.5327 8.139 39.5328 9.65322C39.5328 12.0042 39.5327 14.3553 39.5328 16.7063C39.5329 17.5446 39.8504 17.8279 40.7893 17.828C42.2241 17.8283 43.6589 17.8193 45.0936 17.8307C46.814 17.8444 48.0439 18.9736 47.9989 20.478C47.9619 21.7145 46.8479 22.801 45.4788 22.8782C44.643 22.9253 43.8021 22.9015 42.9635 22.9026C41.7605 22.9041 40.5565 22.9264 39.3547 22.8915C36.724 22.8149 34.5404 21.1355 34.0292 18.8121C33.9482 18.4439 33.9227 18.0607 33.9208 17.6842C33.9107 15.6321 33.916 13.5799 33.916 11.5277L33.918 11.5277Z"
                  fill="white"
                />
                <path
                  d="M31.0587 16.6535C31.0587 17.859 31.0618 19.0645 31.0581 20.27C31.049 23.2151 28.6035 25.4339 25.3454 25.4458C23.458 25.4527 21.5704 25.4591 19.6832 25.4386C18.4313 25.4249 17.3701 24.6341 17.0617 23.5438C16.7494 22.4395 17.2811 21.2952 18.3728 20.7232C18.8327 20.4822 19.3303 20.3775 19.8605 20.3784C21.3175 20.3808 22.7744 20.3796 24.2314 20.3789C25.1124 20.3784 25.4417 20.0771 25.442 19.2721C25.4426 17.0903 25.4294 14.9083 25.4472 12.7266C25.4611 11.0294 27.1986 9.8447 28.9951 10.2801C30.191 10.57 31.0386 11.5383 31.0531 12.6784C31.0698 14.0033 31.057 15.3284 31.057 16.6535H31.0587Z"
                  fill="white"
                />
                <path
                  d="M22.5851 8.89035C22.5853 11.0114 22.6 13.1326 22.5809 15.2535C22.5642 17.1083 20.5362 18.326 18.6827 17.5981C17.6291 17.1844 17.0681 16.428 16.9764 15.3949C16.9562 15.1673 16.9686 14.9371 16.9686 14.7082C16.9684 10.6851 16.9663 6.66211 16.9698 2.63909C16.971 1.2851 17.9057 0.259774 19.3239 0.0375143C20.9972 -0.224741 22.5518 0.920163 22.5771 2.46753C22.6023 4.01069 22.5841 5.55442 22.5848 7.0979C22.585 7.69539 22.5848 8.29287 22.5851 8.89035Z"
                  fill="white"
                />
                <path
                  d="M28.2522 2.55469C29.8077 2.55407 31.0536 3.6739 31.057 5.07559C31.0603 6.4854 29.7941 7.62682 28.2357 7.61881C26.6926 7.61089 25.4404 6.46998 25.4443 5.07567C25.4483 3.67749 26.6982 2.5553 28.2522 2.55469Z"
                  fill="white"
                />
                <path
                  d="M36.7028 25.4602C36.7046 26.8583 35.4584 27.9924 33.9108 28.0011C32.3648 28.0098 31.0773 26.8477 31.086 25.4514C31.0947 24.0566 32.352 22.934 33.904 22.9355C35.4611 22.9371 36.7011 24.0556 36.7028 25.4602Z"
                  fill="white"
                />
              </svg>
            </div>
          </div>
          <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
    fill="white"
  />
</svg>
        </a>
        <div class="h-[1px] bg-[#A1B4D680] min-w-[60%]"></div>
        <a
          class="linearBorder justify-between min-w-[60%] h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
          href="https://sibche.com/applications/mygov/"
        >
          <div class="flex items-center gap-5 w-full">
            <svg
  width="32"
  height="36"
  viewBox="0 0 40 40"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
>
  <rect
    x="2.3999"
    width="35.2667"
    height="39.9245"
    fill="url(#pattern0_2842_101903)"
  />
</svg>

            <div class="mt-2">
              <p class="text-sm">دریافت نسخه iOS از</p>

              <svg
  width="98"
  height="33"
  viewBox="0 0 98 33"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  class="mx-auto"
>
  <rect width="98" height="32.1235" fill="url(#pattern0_2842_101973)" />
</svg>
            </div>
          </div>
          <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
    fill="white"
  />
</svg>
        </a>
        <div class="h-[1px] bg-[#A1B4D680] min-w-[60%]"></div>
        <a
          class="linearBorder justify-between min-w-[60%] h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
          href="https://iapps.ir/app/%D9%BE%D9%86%D8%AC%D8%B1%D9%87-%D9%85%D9%84%DB%8C-%D8%AE%D8%AF%D9%85%D8%A7%D8%AA-%D8%AF%D9%88%D9%84%D8%AA-%D9%87%D9%88%D8%B4%D9%85%D9%86%D8%AF-%D8%AF%D9%88%D9%84%D8%AA-%D9%85%D9%86-/964465353"
        >
          <div class="flex items-center gap-5 w-full">
            <svg
  width="32"
  height="36"
  viewBox="0 0 40 40"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
>
  <g clip-path="url(#clip0_2842_102012)">
    <rect
      y="-2"
      width="40"
      height="40"
      rx="8"
      fill="url(#pattern0_2842_102012)"
    />
  </g>
</svg>

            <div class="mt-2">
              <p class="text-sm">دریافت نسخه iOS از</p>
              <svg
  width="122"
  height="24"
  viewBox="0 0 122 24"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  class="mx-auto"
>
  <rect width="122" height="24" fill="url(#pattern0_2842_102062)" />
</svg>
            </div>
          </div>
          <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
    fill="white"
  />
</svg>
        </a>
        <div class="h-[1px] bg-[#A1B4D680] min-w-[60%]"></div>
        <a
          class="linearBorder justify-between min-w-[60%] h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
          href="https://myket.ir/app/ir.gov.mygov"
        >
          <div class="flex items-center gap-5 w-full">
            <svg
              width="32"
              height="36"
              viewBox="0 0 40 40"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M38.5888 8.38785C37.1551 5.97741 34.2879 3.2054 28.6728 3.2054C27.066 3.15183 25.4665 3.44691 23.9826 4.07069C22.4986 4.69446 21.1646 5.63239 20.0709 6.82106C18.9772 5.63239 17.6433 4.69446 16.1593 4.07069C14.6753 3.44691 13.0759 3.15183 11.4691 3.2054C3.22568 3.2054 0 9.35203 0 13.5703V27.5509C0 28.318 0.302085 29.0537 0.839804 29.5962C1.37752 30.1387 2.10683 30.4434 2.86728 30.4434C3.62773 30.4434 4.35703 30.1387 4.89475 29.5962C5.43247 29.0537 5.73456 28.318 5.73456 27.5509V13.6908C5.85403 12.6061 6.45137 8.99046 11.4691 8.99046C12.2095 8.93855 12.9525 9.04541 13.649 9.30401C14.3456 9.56261 14.9799 9.96708 15.5104 10.4908C16.0408 11.0145 16.4553 11.6456 16.7266 12.3426C16.9978 13.0396 17.1197 13.7866 17.0842 14.5345V26.2251C17.0842 26.9923 17.3863 27.728 17.924 28.2705C18.4617 28.8129 19.191 29.1177 19.9515 29.1177C20.7119 29.1177 21.4412 28.8129 21.9789 28.2705C22.5167 27.728 22.8187 26.9923 22.8187 26.2251V14.5345C22.7832 13.7866 22.9051 13.0396 23.1763 12.3426C23.4476 11.6456 23.8621 11.0145 24.3925 10.4908C24.923 9.96708 25.5573 9.56261 26.2539 9.30401C26.9504 9.04541 27.6934 8.93855 28.4338 8.99046C33.4515 8.99046 34.0489 12.4856 34.1684 13.6908V27.5509C34.1684 28.318 34.4705 29.0537 35.0082 29.5962C35.5459 30.1387 36.2752 30.4434 37.0357 30.4434C37.7961 30.4434 38.5254 30.1387 39.0631 29.5962C39.6008 29.0537 39.9029 28.318 39.9029 27.5509V13.5703C40.1419 13.3293 40.0224 10.9188 38.5888 8.38785Z"
                fill="white"
              />
              <path
                d="M12.0601 31.0488C12.8804 32.6118 14.1251 33.9066 15.6485 34.7817C17.1719 35.6568 18.9109 36.076 20.6619 35.9902C22.4088 36.0519 24.1383 35.6225 25.6573 34.7499C27.1763 33.8773 28.4249 32.5959 29.2637 31.0488C29.2637 31.0488 20.6619 38.4007 12.0601 31.0488Z"
                fill="white"
              />
            </svg>

            <div class="mt-2">
              <p class="text-sm">دریافت نسخه اندروید از</p>
              <svg
                width="94"
                height="30"
                viewBox="0 0 94 30"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                class="mx-auto"
              >
                <path
                  d="M14.3391 9.25714L11.7186 6.98047L9.09814 9.25714L11.7186 11.5338L14.3391 9.25714Z"
                  fill="white"
                />
                <path
                  d="M20.9758 9.26048L18.3553 6.83203L15.7349 9.26048L18.3553 11.5372L20.9758 9.26048Z"
                  fill="white"
                />
                <path
                  d="M65.6976 14.4189C65.6976 13.9636 65.5229 13.5083 65.5229 13.2047C65.3907 12.8627 65.3315 12.5027 65.3482 12.1423C65.3482 11.8387 65.1735 11.5351 65.1735 11.2316C65.1826 11.0662 65.1203 10.9038 64.9988 10.7762L61.3302 11.6869C61.3302 11.8387 61.5049 11.9905 61.5049 12.294C61.5049 12.5976 61.6796 12.7494 61.6796 13.0529C61.6796 13.3565 61.8543 13.66 61.8543 13.9636C61.8543 14.2671 62.029 14.5707 62.029 15.026C62.029 15.3296 62.2037 15.6332 62.2037 15.9367V16.6956C62.2037 17.1509 62.2037 17.6063 61.8543 17.758C61.7169 18.0292 61.4676 18.2458 61.1555 18.3652C60.9808 18.5169 60.6314 18.5169 60.4567 18.6687H52.9447C52.3466 18.6562 51.7555 18.5534 51.1977 18.3652C50.6387 18.1713 50.1103 17.9163 49.6255 17.6063C49.0316 17.2635 48.55 16.7929 48.2279 16.2403C47.6213 15.6903 47.0934 15.0788 46.6556 14.4189C45.6074 13.2047 44.734 12.1423 43.6858 11.0798C42.6376 10.0174 41.7641 8.80313 40.7159 7.74069L52.5953 3.03556L51.0231 0L37.3967 5.76757L36.6979 8.49958C37.3967 9.41025 38.2702 10.1691 38.9689 10.928C39.6677 11.6869 40.3665 12.5976 41.0653 13.2047C41.7641 13.9636 42.1135 14.5707 42.6376 15.1778C42.9477 15.6407 43.1273 16.1608 43.1617 16.6956C43.1533 17.0228 43.0315 17.3401 42.8123 17.6063C42.6938 17.7637 42.5398 17.8986 42.3597 18.003C42.1795 18.1073 41.9769 18.1789 41.7641 18.2134C41.2744 18.4196 40.736 18.5235 40.1918 18.5169C39.6677 18.5169 38.9689 18.6687 38.2702 18.6687H32.6798C32.585 18.679 32.4887 18.6706 32.398 18.6444C32.3073 18.6181 32.2246 18.5746 32.1558 18.5169C31.9811 18.5169 31.8064 18.3652 31.457 18.3652C31.2823 18.2134 31.1076 18.2134 30.9329 18.0616C30.8263 17.9894 30.7397 17.8974 30.6792 17.7923C30.6188 17.6873 30.5861 17.5719 30.5835 17.4545C30.4249 17.2271 30.3631 16.9585 30.4088 16.6956V12.7494H26.5654V17.1509C26.5553 17.5252 26.435 17.8909 26.2161 18.2134C25.9496 18.3984 25.6558 18.5515 25.3426 18.6687C24.894 18.8034 24.4179 18.8551 23.945 18.8205H5.07768C4.81491 18.7827 4.57138 18.6769 4.37889 18.5169C4.16311 18.344 3.98549 18.1382 3.8548 17.9098C3.69882 17.5731 3.63888 17.2086 3.6801 16.8474C3.65243 16.536 3.71238 16.2235 3.8548 15.9367C3.8548 15.6332 4.0295 15.1778 4.0295 14.8743C4.2042 14.5707 4.2042 14.1154 4.37889 13.8118C4.55359 13.5083 4.55359 13.2047 4.72828 12.9011L2.80662 12.4458L1.05964 11.8387C1.05964 11.9905 0.88495 12.1423 0.88495 12.4458C0.710253 12.7494 0.710251 13.2047 0.535554 13.66C0.360857 14.1154 0.360856 14.5707 0.186158 15.1778C0.0417092 15.6734 -0.0171636 16.1849 0.0114637 16.6956C-0.027171 17.1564 0.0320433 17.6194 0.186158 18.0616C0.360856 18.5169 0.360854 18.8205 0.710249 19.124L1.23434 20.0347C1.40904 20.3383 1.75843 20.4901 1.93313 20.6418C2.50488 21.0822 3.15352 21.4408 3.8548 21.7043C4.59342 21.8908 5.35688 21.9928 6.12587 22.0078H24.2944C25.9749 22.0465 27.6262 21.6214 29.0112 20.7936C29.3175 21.0472 29.6722 21.2526 30.0594 21.4007C30.4527 21.58 30.8616 21.7322 31.2823 21.8561C31.6731 21.9806 32.0898 22.0323 32.5051 22.0078H38.7942C39.6133 22.0277 40.4325 21.9769 41.24 21.8561C42.1135 21.7043 42.8123 21.5525 43.5111 21.4007C44.1712 21.1294 44.8128 20.8253 45.4327 20.4901C45.9371 20.0484 46.3507 19.5351 46.6556 18.9723C47.0767 19.4141 47.5446 19.8207 48.0532 20.1865C48.5773 20.4901 48.9267 20.9454 49.6255 21.0972C50.2234 21.3897 50.8726 21.5948 51.5471 21.7043C52.3521 21.8419 53.1736 21.8929 53.9929 21.8561H60.8061C61.5165 21.8291 62.2204 21.7272 62.9025 21.5525C63.6369 21.3466 64.2984 20.9808 64.8241 20.4901C65.3859 19.9992 65.8573 19.4361 66.2217 18.8205C66.5794 18.049 66.7575 17.2237 66.7458 16.392V15.3296C65.8723 15.1778 65.6976 14.7225 65.6976 14.4189Z"
                  fill="white"
                />
                <path
                  d="M60.2812 27.1687L62.9017 29.4454L65.5222 27.1687L62.9017 24.7402L60.2812 27.1687Z"
                  fill="white"
                />
                <path
                  d="M53.644 27.1673L56.2645 29.444L58.885 27.1673L56.2645 24.8906L53.644 27.1673Z"
                  fill="white"
                />
                <path
                  d="M92.0787 12.2947C91.5405 11.6782 90.8143 11.205 89.9824 10.9287C89.1061 10.5676 88.156 10.3612 87.1872 10.3216C86.2057 10.2684 85.2293 10.4804 84.3921 10.9287C83.6772 11.2989 83.0296 11.7592 82.4704 12.2947C81.9984 12.8674 81.589 13.4772 81.2475 14.116C80.9594 14.758 80.7258 15.4175 80.5487 16.0892C80.374 16.3927 80.374 16.848 80.1993 17.1516C80.0246 17.4552 80.0246 17.7587 79.8499 17.9105L79.3259 18.3658C79.257 18.4235 79.1743 18.467 79.0836 18.4933C78.9929 18.5195 78.8966 18.5278 78.8018 18.5176H77.5789C76.8753 18.5565 76.169 18.5053 75.4825 18.3658C75.1005 18.2789 74.7435 18.1238 74.4343 17.9105C74.228 17.5688 74.0528 17.2135 73.9102 16.848V2.27734H70.0669V17.4552C70.0826 18.0241 70.2005 18.5874 70.4163 19.1247C70.591 19.5801 70.9404 20.0354 71.1151 20.4907C71.5484 20.8302 72.0162 21.135 72.5127 21.4014C73.1611 21.7168 73.8724 21.9228 74.609 22.0085C75.3559 22.1404 76.1188 22.1914 76.8801 22.1603H78.8018C79.4566 22.1467 80.105 22.0442 80.7234 21.8567C81.3216 21.663 81.8598 21.3512 82.2957 20.9461C82.8599 21.2294 83.4433 21.4828 84.0427 21.7049C84.5318 21.9698 85.061 22.1741 85.615 22.3121C86.0631 22.467 86.5331 22.5691 87.0125 22.6156C87.5235 22.7408 88.0547 22.7921 88.5848 22.7674H89.4583C89.8806 22.7177 90.2924 22.6154 90.6812 22.4638C91.1047 22.2866 91.5133 22.0838 91.904 21.8567C92.2958 21.5358 92.6472 21.1797 92.9522 20.7943C93.3171 20.2864 93.5552 19.7176 93.651 19.1247C93.8754 18.3297 93.9926 17.5148 94.0004 16.6963C93.9812 15.8785 93.8641 15.0646 93.651 14.2678C93.0686 13.6467 92.5429 12.987 92.0787 12.2947ZM90.1571 18.3658C90.0222 18.638 89.8459 18.8932 89.633 19.1247C89.5498 19.2173 89.444 19.2926 89.3231 19.3451C89.2022 19.3976 89.0693 19.426 88.9342 19.4283H88.4101C88.0524 19.441 87.6955 19.3893 87.3619 19.2765C86.9686 19.0972 86.5597 18.945 86.139 18.8212C85.7896 18.6694 85.2656 18.5176 84.9162 18.3658C84.5668 18.2141 84.0427 17.9105 83.6933 17.7587L84.2174 16.3927L84.7415 15.0267C84.9291 14.6163 85.2289 14.2516 85.615 13.9643C85.9948 13.6601 86.4962 13.4968 87.0125 13.5089C87.5567 13.5023 88.0951 13.6063 88.5848 13.8125L89.633 14.7232C89.8077 15.0267 89.9824 15.482 90.1571 15.7856C90.2995 16.0723 90.3594 16.3849 90.3318 16.6963C90.401 17.2574 90.3416 17.8248 90.1571 18.3658Z"
                  fill="white"
                />
              </svg>
            </div>
          </div>
          <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
    fill="white"
  />
</svg>
        </a>
        <div class="h-[1px] bg-[#A1B4D680] min-w-[60%]"></div>
        <a
          class="linearBorder justify-between min-w-[60%] h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
          href="https://sibapp.com/applications/MyGov"
        >
          <div class="flex items-center gap-5 w-full">
            <svg
              width="32"
              height="36"
              viewBox="0 0 40 40"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M25.366 17.3762C27.0674 18.9186 28.2166 20.9761 28.6375 23.2337C29.0584 25.4913 28.7279 27.8247 27.6965 29.8766C33.3111 20.5543 31.1924 13.2448 31.1924 13.2448C19.1158 7.20648 12.3359 12.3973 12.3359 12.3973C17.1395 12.4212 21.7705 14.1907 25.366 17.3762Z"
                fill="white"
              />
              <path
                d="M11.7011 19.3863C13.2289 17.6632 15.2876 16.4985 17.5515 16.0764C19.8154 15.6543 22.1554 15.9989 24.2015 17.0558C14.9851 11.4412 7.6756 13.5599 7.6756 13.5599C1.74324 25.7424 6.82813 32.5223 6.82813 32.5223C6.71272 27.6816 8.45662 22.9806 11.7011 19.3863Z"
                fill="white"
              />
              <path
                d="M13.82 32.9456C12.1185 31.4033 10.9694 29.3458 10.5485 27.0882C10.1276 24.8306 10.4581 22.4972 11.4894 20.4453C5.98081 29.7676 7.99359 37.0771 7.99359 37.0771C20.0702 43.1154 26.85 37.9246 26.85 37.9246C22.0284 37.9998 17.3628 36.217 13.82 32.9456Z"
                fill="white"
              />
              <path
                d="M32.5706 17.7949C32.5664 22.6339 30.796 27.3049 27.5917 30.9309C26.0639 32.654 24.0052 33.8187 21.7413 34.2408C19.4774 34.6629 17.1374 34.3183 15.0913 33.2614C24.3077 38.876 31.6172 36.7573 31.6172 36.7573C37.6555 24.5748 32.5706 17.7949 32.5706 17.7949Z"
                fill="white"
              />
              <path
                d="M28.2271 0C19.5404 0.21187 19.0107 7.73325 19.0107 7.73325C27.6974 8.4748 28.2271 0 28.2271 0Z"
                fill="white"
              />
            </svg>

            <div class="mt-2">
              <p class="text-sm">دریافت نسخه iOS از</p>
              <svg
                width="118"
                height="28"
                viewBox="0 0 118 28"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                class="mx-auto"
              >
                <path
                  d="M30.7006 20.1428L28.3165 15.805C28.1053 15.4203 27.7805 15.0967 27.3786 14.8703C26.9767 14.644 26.5134 14.5239 26.0407 14.5234H1.02545C0.851144 14.5241 0.679873 14.5649 0.527678 14.6422C0.375483 14.7195 0.247347 14.8307 0.15529 14.9654C0.063233 15.1 0.0102651 15.2537 0.00134815 15.4121C-0.00756883 15.5704 0.0278602 15.7283 0.104311 15.8708L1.71179 18.7133C2.13703 19.4591 2.78015 20.0838 3.57042 20.5186C4.36069 20.9535 5.26749 21.1816 6.19106 21.1779H29.9781C30.1129 21.1788 30.2457 21.1485 30.3641 21.0898C30.4825 21.0311 30.5825 20.946 30.6547 20.8425C30.727 20.739 30.7691 20.6204 30.7771 20.498C30.7851 20.3756 30.7588 20.2534 30.7006 20.1428Z"
                  fill="white"
                />
                <path
                  d="M17.4983 23.4473H13.7596C13.5854 23.4472 13.4181 23.509 13.2933 23.6195C13.1685 23.73 13.096 23.8804 13.0913 24.0388V27.3249C13.096 27.4833 13.1685 27.6337 13.2933 27.7442C13.4181 27.8547 13.5854 27.9165 13.7596 27.9164H17.4983C17.6725 27.9165 17.8398 27.8547 17.9646 27.7442C18.0894 27.6337 18.1619 27.4833 18.1666 27.3249V24.0388C18.1619 23.8804 18.0894 23.73 17.9646 23.6195C17.8398 23.509 17.6725 23.4472 17.4983 23.4473Z"
                  fill="white"
                />
                <path
                  d="M116.999 14.5258H89.7625C89.6882 14.5277 89.6145 14.5128 89.548 14.4824C89.4815 14.4521 89.4243 14.4074 89.3814 14.3521C89.3385 14.2968 89.3113 14.2327 89.3021 14.1656C89.293 14.0984 89.3023 14.0303 89.3291 13.9672C89.7522 12.8747 89.966 11.7243 89.9612 10.566V7.16483C89.9612 6.96001 89.8718 6.76359 89.7126 6.61877C89.5534 6.47394 89.3375 6.39258 89.1123 6.39258H83.3326C83.1075 6.39258 82.8915 6.47394 82.7323 6.61877C82.5732 6.76359 82.4837 6.96001 82.4837 7.16483V10.6153C82.4742 11.6554 82.0133 12.65 81.2014 13.3825C80.3895 14.1149 79.2924 14.5259 78.1489 14.5258H75.6384C75.5662 14.5266 75.4949 14.5117 75.4303 14.4823C75.3658 14.4529 75.31 14.4098 75.2675 14.3567C75.2251 14.3036 75.1972 14.242 75.1863 14.1771C75.1754 14.1122 75.1818 14.0458 75.2049 13.9836C75.6078 12.9266 75.8153 11.8162 75.819 10.6974V7.24698C75.819 7.14417 75.7965 7.04239 75.7527 6.94761C75.7089 6.85283 75.6448 6.76696 75.564 6.69503C75.4833 6.6231 75.3875 6.56656 75.2824 6.52873C75.1774 6.4909 75.065 6.47255 74.952 6.47473H69.1182C68.8913 6.47469 68.6735 6.55552 68.5114 6.69989C68.3493 6.84426 68.2559 7.04067 68.2512 7.24698V10.6482C68.2489 11.1666 68.1337 11.6796 67.9123 12.1573C67.6909 12.6351 67.3677 13.0683 66.9614 13.4319C66.555 13.7954 66.0735 14.0822 65.5447 14.2756C65.0159 14.469 64.4502 14.5652 63.8803 14.5587H47.1192C46.6464 14.5591 46.1832 14.6792 45.7813 14.9056C45.3794 15.1319 45.0546 15.4556 44.8434 15.8403L42.4593 20.178C42.4011 20.2886 42.3747 20.4109 42.3828 20.5333C42.3908 20.6557 42.4329 20.7742 42.5052 20.8777C42.5774 20.9812 42.6774 21.0663 42.7958 21.125C42.9141 21.1837 43.0469 21.214 43.1817 21.2131H63.9164C65.6701 21.2188 67.4037 20.8745 68.9953 20.2045C70.5868 19.5345 71.9975 18.5552 73.1278 17.3355L74.7895 20.3587C74.924 20.602 75.1306 20.8063 75.3859 20.9486C75.6412 21.0908 75.9351 21.1654 76.2344 21.1639H78.0406C79.894 21.1715 81.7239 20.7865 83.3854 20.0393C85.047 19.2921 86.4944 18.2033 87.6132 16.859L89.311 20.3423C89.4365 20.6004 89.6414 20.82 89.901 20.9745C90.1606 21.129 90.4638 21.2119 90.774 21.2131H111.834C112.761 21.2128 113.67 20.9813 114.463 20.5439C115.256 20.1065 115.902 19.4799 116.331 18.7321L117.884 15.8731C117.965 15.7336 118.004 15.5777 118 15.4202C117.995 15.2628 117.946 15.1091 117.857 14.9739C117.768 14.8387 117.643 14.7265 117.492 14.648C117.342 14.5696 117.173 14.5275 116.999 14.5258Z"
                  fill="white"
                />
                <path
                  d="M86.0581 23.4473H75.9798C75.8057 23.4472 75.6384 23.509 75.5135 23.6195C75.3887 23.73 75.3162 23.8804 75.3115 24.0388V27.3249C75.3162 27.4833 75.3887 27.6337 75.5135 27.7442C75.6384 27.8547 75.8057 27.9165 75.9798 27.9164H86.0581C86.2323 27.9165 86.3996 27.8547 86.5244 27.7442C86.6492 27.6337 86.7217 27.4833 86.7264 27.3249V24.0388C86.7217 23.8804 86.6492 23.73 86.5244 23.6195C86.3996 23.509 86.2323 23.4472 86.0581 23.4473Z"
                  fill="white"
                />
                <path
                  d="M30.5918 23.4473H20.5135C20.3393 23.4472 20.172 23.509 20.0472 23.6195C19.9224 23.73 19.8499 23.8804 19.8452 24.0388V27.3249C19.8499 27.4833 19.9224 27.6337 20.0472 27.7442C20.172 27.8547 20.3393 27.9165 20.5135 27.9164H30.5918C30.766 27.9165 30.9333 27.8547 31.0581 27.7442C31.1829 27.6337 31.2554 27.4833 31.2601 27.3249V24.0388C31.2554 23.8804 31.1829 23.73 31.0581 23.6195C30.9333 23.509 30.766 23.4472 30.5918 23.4473Z"
                  fill="white"
                />
                <path
                  d="M61.1155 23.4475H57.3948C57.3079 23.4453 57.2214 23.459 57.1405 23.4877C57.0595 23.5165 56.9856 23.5597 56.9233 23.6148C56.861 23.67 56.8115 23.7359 56.7777 23.8088C56.7439 23.8816 56.7265 23.9599 56.7266 24.039V27.3252C56.7265 27.4042 56.7439 27.4825 56.7777 27.5554C56.8115 27.6282 56.861 27.6942 56.9233 27.7493C56.9856 27.8045 57.0595 27.8477 57.1405 27.8765C57.2214 27.9052 57.3079 27.9189 57.3948 27.9167H61.1155C61.2024 27.9189 61.2889 27.9052 61.3699 27.8765C61.4509 27.8477 61.5247 27.8045 61.587 27.7493C61.6493 27.6942 61.6988 27.6282 61.7326 27.5554C61.7664 27.4825 61.7838 27.4042 61.7838 27.3252V24.039C61.7838 23.9599 61.7664 23.8816 61.7326 23.8088C61.6988 23.7359 61.6493 23.67 61.587 23.6148C61.5247 23.5597 61.4509 23.5165 61.3699 23.4877C61.2889 23.459 61.2024 23.4453 61.1155 23.4475Z"
                  fill="white"
                />
                <path
                  d="M39.695 4.40979e-08H33.4998C33.3305 -5.94084e-05 33.1678 0.0599983 33.0463 0.167398C32.9249 0.274799 32.8543 0.421065 32.8496 0.575078V20.5385C32.8543 20.7099 32.9325 20.8729 33.0675 20.9926C33.2025 21.1123 33.3836 21.1793 33.5721 21.1793H39.6227C39.8095 21.1793 39.9887 21.1118 40.1208 20.9916C40.2529 20.8714 40.3271 20.7084 40.3271 20.5385V0.575078C40.3271 0.422558 40.2605 0.276284 40.142 0.168436C40.0234 0.0605879 39.8626 4.40979e-08 39.695 4.40979e-08Z"
                  fill="white"
                />
              </svg>
            </div>
          </div>
          <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
    fill="white"
  />
</svg>
        </a>
      </section>
    </div>
    <!-- ..........................Privacy..................................... -->
    <div id="privacy-content">
      <section class="px-10">
        <p
          class="text-[32px] font-bold flex items-center justify-center mt-[106px] mb-6"
        >
          حریــــم خــــصوصی
        </p>
        <p class="font-normal text-justify mb-8">
          اطلاعات مربوط به هر شخص، حریم خصوصی وی محسوب می‌شود. حفاظت و حراست از
          اطلاعات شخصی شهروندان در پنجره ملی خدمات دولت هوشمند، نه تنها موجب حفظ
          امنیت کاربران می‌شود، بلکه باعث اعتماد بیشتر و مشارکت آنها در
          فعالیت‌های جاری می‌گردد. هدف از این بیانیه، آگاه ساختن شما درباره ی
          نوع و نحوه ی استفاده از اطلاعاتی است که در هنگام استفاده از  پنجره ملی
          خدمات دولت هوشمند، از جانب شما دریافت می‌گردد. سازمان فناوری اطلاعات
          ایران خود را ملزم به رعایت حریم خصوصی همه  شهروندان و کاربران سامانه
          دانسته و آن دسته از اطلاعات کاربران را که فقط به منظور ارائه خدمات
          کفایت می‌کند، دریافت کرده و از انتشار آن یا در اختیار قرار دادن آن به
          دیگران خودداری مینماید. اطلاعات جمع آوری شده از بازدیدکنندگان در بخش
          نظرسنجی کارکنان دولت و سازمان ها، صرفا ً برای بهبود کیفیت خدمات و
          محتوای تارنما مورد استفاده قرار می‌گیرند.
        </p>
        <p class="font-bold text-lg mb-4">
          چگونگی جمع آوری و استفاده از اطلاعات کاربران
        </p>
        <p class="font-medium">
          لف: اطلاعاتی که شما خود در اختيار این سازمان قرار می‌دهيد، شامل موارد
          زيرهستند:
        </p>
        <p class="font-normal">
          در درگاه پنجره ملی خدمات دولت هوشمند، اقلام اطلاعاتی شامل شماره تلفن
          همراه، تاریخ تولد، کد پستی و کد ملی کاربران را دریافت مینماییم که از
          این اقلام، صرفا جهت احراز هویت کاربران استفاده خواهد شد.
        </p>
        <p class="font-medium">
          ب: برخی اطلاعات ديگر که به صورت خودکار از شما دريافت میشود شامل موارد
          زير می‌باشد:
        </p>
        <ul class="list-disc pr-4.5 mt-2">
          <li>
            دستگاهی که از طریق آن درگاه سازمان فناوری اطلاعات ایران را مشاهده
            می‌نمایید.
          </li>
          <li>نام و نسخه سیستم عامل وbrowser کامپیوتر شما.</li>
          <li>اطلاعات صفحات بازدید شده.</li>
          <li>تعداد بازدیدهای روزانه در درگاه.</li>
          <li>
            هدف ما از دریافت این اطلاعات استفاده از آنها در تحلیل عملکرد کاربران
            درگاه می باشد تا بتوانیم در خدمت رسانی بهتر عمل کنیم.
          </li>
        </ul>
        <p class="text-lg font-bold mb-4 mt-6">امنیت اطلاعات</p>
        <p class="font-normal text-justify">
          متعهدیم که امنیت اطلاعات شما را تضمین نماییم و برای جلوگیری از هر نوع
          دسترسی غیرمجاز و افشای اطلاعات شما از همه شیوه‌‌های لازم استفاده
          می‌کنیم تا امنیت اطلاعاتی را که به صورت آنلاین گردآوری می‌کنیم، حفظ
          شود. لازم به ذکر است در سامانه ما، ممکن است به سایت های دیگری لینک
          شوید، وقتی که شما از طریق این لینک‌ها از سامانه ما خارج می‌شوید، توجه
          داشته باشید که ما بر دیگر سایت ها کنترل نداریم و سازمان تعهدی بر حفظ
          حریم شخصی آنان در سایت مقصد نخواهد داشت و مراجعه کنندگان  میبایست به
          بیانیه حریم شخصی آن سایت ها مراجعه نمایند.
        </p>
        <p class="text-lg font-bold mt-8 mb-4">
          تغييرات در بيانيه سياست حفظ محرمانگی 
        </p>
        <p class="font-normal text-justify">
          ما این حق را برای خود محفوظ می‌دانیم که در هر زمانی این بیانیه را
          ویرایش نموده و تغییرات خود را اعمال نماییم. در این شرایط شما موظف
          هستید که آخرین تغییرات ما را در بیانیه از طریق سامانه مشاهده نموده و
          از آن مطلع شوید. 
        </p>
      </section>
    </div>
    <!-- .......................... System Guide...............................-->
    <!-- <div id="help-content">
      <section class="xl:px-28 px-10 mt-[106px]">
        <p class="text-[32px] font-bold flex items-center justify-center">
          راهنــــما سامــــانه
        </p>
        <p
          class="flex items-center justify-center font-medium text-lg text-center mt-6 opacity-60"
        >
          برای آشنایی سریع با امکانات سامانه، فرآیند ورود و پاسخ به پرسش‌های
          رایج، می‌توانید راهنمای جامع را در قالب PDF دریافت کنید.
        </p>
        <ul class="mt-[60px] px-6 list-disc">
          <li class="font-bold text-[#ECF4FF]">
            ایجاد و فعال‌سازی حساب کاربری
            <p class="font-normal mt-2">
              مراحل ثبت‌نام، تأیید هویت و فعال‌سازی اولیه
            </p>
          </li>
          <div class="w-full bg-white/20 h-[1px] my-4"></div>
          <li class="font-bold text-[#ECF4FF]">
            ورود امن به سامانه
            <p class="font-normal mt-2">
              توضیح شیوه‌های احراز هویت (رمز عبور، رمز یکبار مصرف) و نکات امنیتی
            </p>
          </li>
          <div class="w-full bg-white/20 h-[1px] my-4"></div>
          <li class="font-bold text-[#ECF4FF]">
            بازیابی گذرواژه و بازیابی دسترسی
            <p class="font-normal mt-2">
              راهنمای گام‌به‌گام برای بازیابی رمز عبور و رفع مشکل ورود
            </p>
          </li>
          <div class="w-full bg-white/20 h-[1px] my-4"></div>
          <li class="font-bold text-[#ECF4FF]">
            مدیریت اطلاعات شخصی
            <p class="font-normal mt-2">
              نحوه به‌روزرسانی پروفایل، شماره تماس و تنظیمات امنیتی
            </p>
          </li>
          <div class="w-full bg-white/20 h-[1px] my-4"></div>
          <li class="font-bold text-[#ECF4FF]">
            پرسش‌های متداول (FAQ)
            <p class="font-normal mt-2">
              پاسخ به رایج‌ترین سؤالات کاربران درباره خدمات و خطاهای احتمالی
            </p>
          </li>
          <div class="w-full bg-white/20 h-[1px] my-4"></div>
          <li class="font-bold text-[#ECF4FF]">
            تماس با پشتیبانی
            <p class="font-normal mt-2">
              روش‌های ارتباط با مرکز تماس، ایمیل و تیکت آنلاین
            </p>
          </li>
        </ul>
        <div
          class="w-full rounded-2xl before:bg-[linear-gradient(88.93deg,#304A79_0.91%,#7C96C6_98.92%)] bg-white/20 linearBorder before:rounded-2xl mt-[80px] py-4 px-8 flex items-center justify-between"
        >
          <div class="flex items-center gap-4">
            <svg th:replace="fragments/svgs :: document"></svg>

            <div class="text-[#ECF4FF] flex flex-col gap-1">
              <p class="font-bold">راهنمای جامع پنجره ملی</p>
              <p class="font-normal">فرمت فایل PDF</p>
            </div>
          </div>
          <svg
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M12 14.5L12 4.5M12 14.5C11.2998 14.5 9.99153 12.5057 9.5 12M12 14.5C12.7002 14.5 14.0085 12.5057 14.5 12"
              stroke="white"
              stroke-width="1.5"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M20 16.5C20 18.982 19.482 19.5 17 19.5H7C4.518 19.5 4 18.982 4 16.5"
              stroke="white"
              stroke-width="1.5"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>
      </section>
    </div> -->
  </div>
  <section
    id="drawer_wrapper"
    class="lg:hidden block shadow-[-4px_4px_10px_0px_#00000080] fixed z-50 top-0 transition-all duration-500 ease-in-out bg-[#304A79] -right-[85vw] w-[90vw] h-screen text-white"
  >
    <div class="w-full h-full">
      <div
        id="drawer_toggler"
        class="w-7 h-7 rounded-lg bg-[#304A79] border border-[#8092C1] flex items-center justify-center absolute -left-3.5 top-12"
      >
        <svg
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M5.80441 5.13807C6.06476 4.87772 6.06476 4.45561 5.80441 4.19526C5.54406 3.93491 5.12195 3.93491 4.8616 4.19526L1.52827 7.5286C1.26792 7.78894 1.26792 8.21106 1.52827 8.4714L4.8616 11.8047C5.12195 12.0651 5.54406 12.0651 5.80441 11.8047C6.06476 11.5444 6.06476 11.1223 5.80441 10.8619L2.94248 8L5.80441 5.13807Z"
            fill="#9CBDFF"
          />
          <path
            d="M11.1377 4.19526C10.8774 3.93491 10.4553 3.93491 10.1949 4.19526C9.93459 4.45561 9.93459 4.87772 10.1949 5.13807L13.0569 8L10.1949 10.8619C9.93459 11.1223 9.93459 11.5444 10.1949 11.8047C10.4553 12.0651 10.8774 12.0651 11.1377 11.8047L14.4711 8.4714C14.7314 8.21106 14.7314 7.78894 14.4711 7.5286L11.1377 4.19526Z"
            fill="#9CBDFF"
          />
        </svg>
      </div>
      <div id="drawer_top" class="mt-5 mr-4 pb-4 flex items-center relative">
        <svg
          width="37"
          height="24"
          viewBox="0 0 37 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M23.3048 8.09693C23.4182 8.09693 23.538 8.08977 23.6515 8.09693C24.0613 8.11126 24.3261 8.38341 24.3577 8.8346C24.3766 9.16404 24.3766 9.50065 24.3577 9.83726C24.3324 10.2813 24.0739 10.5606 23.6893 10.5892C23.4056 10.6107 23.1219 10.6107 22.8382 10.5892C22.4473 10.5606 22.1824 10.2383 22.1635 9.79429C22.1509 9.5293 22.1572 9.27147 22.1635 9.00648C22.1698 8.41205 22.441 8.10409 22.958 8.09693C23.0715 8.08977 23.185 8.09693 23.3048 8.09693Z"
            fill="white"
          />
          <path
            d="M23.1132 18.9679C22.9858 18.9679 22.8529 18.9789 22.7255 18.9679C22.4431 18.9347 22.216 18.7469 22.1828 18.4651C22.1495 18.1778 22.1495 17.8795 22.1717 17.5866C22.1994 17.2883 22.4098 17.0783 22.7089 17.0452C22.9803 17.0175 23.2627 17.012 23.5341 17.0396C23.811 17.0673 24.0326 17.233 24.0658 17.5203C24.1046 17.8353 24.1046 18.1612 24.0658 18.4762C24.027 18.78 23.8 18.9458 23.4898 18.9734C23.368 18.9789 23.2406 18.9679 23.1132 18.9679Z"
            fill="white"
          />
          <path
            d="M14.0004 15.3415C13.9018 15.3415 13.7977 15.3415 13.6991 15.3415C13.2992 15.329 13.0472 15.0656 13.0253 14.614C13.0144 14.3632 13.0144 14.1124 13.0253 13.8615C13.0417 13.4539 13.2663 13.1717 13.6169 13.134C13.858 13.109 14.11 13.109 14.351 13.134C14.6961 13.1654 14.9098 13.4288 14.9426 13.8302C14.9591 14.0685 14.9591 14.3068 14.9481 14.5513C14.9262 15.0907 14.6906 15.3415 14.2195 15.3478C14.1483 15.3415 14.0716 15.3415 14.0004 15.3415Z"
            fill="white"
          />
          <path
            d="M22.1587 15.0596C22.1587 14.9466 22.1532 14.8272 22.1587 14.7142C22.1806 14.262 22.422 13.9794 22.8169 13.9606C23.0143 13.948 23.2173 13.948 23.4202 13.9606C23.8316 13.9794 24.0784 14.2683 24.0894 14.7393C24.0949 14.978 24.1004 15.2166 24.0839 15.4615C24.062 15.876 23.8371 16.1335 23.4751 16.1712C23.2776 16.19 23.0747 16.1837 22.8772 16.1774C22.3945 16.1649 22.1806 15.92 22.1642 15.3736C22.1587 15.2668 22.1642 15.1664 22.1587 15.0596Z"
            fill="white"
          />
          <path
            d="M18.8335 13.8236C18.8335 13.7029 18.8335 13.5893 18.8335 13.4686C18.8398 12.858 19.1124 12.5598 19.6513 12.5598C19.8605 12.5598 20.0697 12.5527 20.2852 12.5598C20.729 12.574 21.0206 12.8864 21.0396 13.3692C21.0523 13.6603 21.0523 13.9585 21.0396 14.2495C21.0206 14.7039 20.786 14.995 20.374 15.0447C20.1267 15.0731 19.8668 15.0731 19.6196 15.0589C19.0871 15.0305 18.8462 14.7323 18.8398 14.136C18.8335 14.0366 18.8335 13.9301 18.8335 13.8236Z"
            fill="white"
          />
          <path
            d="M27.1287 11.0471C27.1287 11.1683 27.135 11.2839 27.1287 11.4051C27.1034 11.7189 26.8819 11.9502 26.5213 11.9777C26.1986 12.0053 25.8696 12.0053 25.5469 11.9777C25.1799 11.9502 24.9457 11.7079 24.9331 11.383C24.9204 11.1628 24.9268 10.9425 24.9331 10.7222C24.9457 10.3478 25.1988 10.0945 25.6228 10.0614C25.8949 10.0394 26.1733 10.0394 26.4454 10.0614C26.8819 10.0945 27.1224 10.3368 27.1414 10.7167C27.135 10.8269 27.1287 10.937 27.1287 11.0471Z"
            fill="white"
          />
          <path
            d="M20.0604 11.9835C19.8968 11.9835 19.7262 11.9962 19.5626 11.9835C19.1716 11.9517 18.8943 11.7292 18.8587 11.3795C18.8303 11.0553 18.8232 10.7247 18.8516 10.3941C18.8872 10.0253 19.2 9.78373 19.6124 9.77101C19.9182 9.76466 20.231 9.7583 20.5368 9.77101C20.9492 9.78373 21.2549 10.0317 21.2976 10.4004C21.326 10.6738 21.3331 10.9535 21.3189 11.2269C21.2905 11.7546 21.0061 11.9835 20.4159 11.9962C20.295 11.9962 20.1813 11.9962 20.0604 11.9962C20.0604 11.9899 20.0604 11.9835 20.0604 11.9835Z"
            fill="white"
          />
          <path
            d="M19.9373 17.8531C19.8108 17.8531 19.6843 17.8595 19.5578 17.8531C19.1403 17.834 18.862 17.5794 18.843 17.1593C18.8303 16.8856 18.8303 16.6056 18.843 16.3319C18.862 15.9182 19.1403 15.6445 19.5515 15.6317C19.8045 15.6254 20.0575 15.6254 20.3106 15.6317C20.766 15.6445 21.0254 15.9118 21.0443 16.3764C21.0507 16.5992 21.0507 16.822 21.0443 17.0448C21.0317 17.5985 20.7787 17.8531 20.222 17.8595C20.1271 17.8595 20.0322 17.8531 19.9373 17.8531Z"
            fill="white"
          />
          <path
            d="M17.9971 15.7705C17.9971 15.8984 18.0026 16.0317 17.9971 16.1596C17.9749 16.4763 17.7477 16.7153 17.4319 16.732C17.177 16.7487 16.9166 16.7487 16.6562 16.732C16.3182 16.7098 16.0911 16.4763 16.0745 16.1373C16.0634 15.8872 16.0578 15.6372 16.0745 15.3871C16.0966 15.0369 16.3293 14.8146 16.6839 14.798C16.8889 14.7868 17.0884 14.7924 17.2934 14.7924C17.7699 14.798 17.9971 15.0314 18.0026 15.5149C18.0026 15.6038 18.0026 15.6872 17.9971 15.7705Z"
            fill="white"
          />
          <path
            d="M26.0387 12.5709C26.178 12.5709 26.311 12.5638 26.4503 12.5709C26.8112 12.5924 27.0962 12.8639 27.1215 13.2712C27.1468 13.6213 27.1468 13.9857 27.1278 14.3358C27.1025 14.7503 26.8365 15.0361 26.4693 15.0575C26.1843 15.0718 25.8994 15.0718 25.6207 15.0575C25.2218 15.0361 24.9432 14.7074 24.9305 14.2501C24.9242 14 24.9242 13.7499 24.9305 13.4998C24.9305 12.8782 25.2091 12.5638 25.7664 12.5566C25.8487 12.5638 25.9437 12.5638 26.0387 12.5709Z"
            fill="white"
          />
          <path
            d="M26.047 15.6333C26.1863 15.6333 26.3192 15.627 26.4584 15.6333C26.8382 15.6587 27.1104 15.9193 27.1294 16.3007C27.1484 16.5994 27.142 16.8918 27.1294 17.1905C27.1104 17.5592 26.8508 17.8198 26.4837 17.8452C26.1926 17.8643 25.8951 17.8643 25.5977 17.8452C25.2242 17.8261 24.9521 17.5465 24.9331 17.1651C24.9204 16.8982 24.9268 16.6376 24.9331 16.3706C24.9457 15.932 25.2242 15.646 25.6673 15.627C25.7939 15.627 25.9204 15.6333 26.047 15.6333Z"
            fill="white"
          />
          <path
            d="M14.9502 5.7148C14.9502 5.83711 14.9557 5.95385 14.9502 6.07616C14.9335 6.42639 14.7334 6.65433 14.3887 6.68212C14.1219 6.70436 13.8495 6.70436 13.5883 6.68768C13.2381 6.65989 13.0324 6.43195 13.0213 6.08728C13.0157 5.85379 13.0157 5.62585 13.0213 5.39236C13.0324 5.00877 13.2603 4.77528 13.6494 4.75304C13.8273 4.74192 13.9996 4.74748 14.1775 4.74748C14.7278 4.75304 14.9557 4.97542 14.9557 5.52023C14.9502 5.58694 14.9502 5.65365 14.9502 5.7148Z"
            fill="white"
          />
          <path
            d="M13.9758 7.53799C14.1023 7.53799 14.2342 7.53161 14.3606 7.53799C14.685 7.56351 14.9269 7.83784 14.9434 8.21424C14.9599 8.50133 14.9599 8.78841 14.9434 9.0755C14.9269 9.46466 14.6795 9.73899 14.3386 9.75813C14.1023 9.77089 13.8604 9.77089 13.624 9.75813C13.2831 9.73899 13.0577 9.49018 13.0302 9.10102C13.0138 8.82669 13.0138 8.54599 13.0248 8.27166C13.0467 7.79956 13.2776 7.55075 13.6954 7.53799C13.7944 7.53799 13.8824 7.53799 13.9758 7.53799Z"
            fill="white"
          />
          <path
            d="M13.9738 12.5524C13.8471 12.5524 13.7149 12.5651 13.5882 12.5524C13.2852 12.5269 13.0593 12.3045 13.0373 11.9549C13.0097 11.618 13.0097 11.2748 13.0373 10.9443C13.0593 10.6074 13.2907 10.3659 13.5882 10.3405C13.8416 10.3214 14.1005 10.3214 14.3594 10.3405C14.6955 10.3659 14.9324 10.6583 14.9489 11.0523C14.9599 11.2748 14.9544 11.4973 14.9489 11.7197C14.9489 12.2918 14.723 12.5524 14.2272 12.5587C14.1391 12.5587 14.0564 12.5587 13.9738 12.5524Z"
            fill="white"
          />
          <path
            d="M13.9717 3.9016C13.8502 3.9016 13.7343 3.90796 13.6128 3.9016C13.2816 3.87618 13.0552 3.64736 13.0332 3.26601C13.0111 2.96092 13.0111 2.64947 13.0332 2.34439C13.0608 1.95031 13.2871 1.70243 13.6239 1.68336C13.8613 1.67065 14.1042 1.67065 14.3416 1.68336C14.6949 1.70243 14.9323 1.98845 14.9489 2.40159C14.9599 2.62405 14.9544 2.84651 14.9489 3.06897C14.9433 3.63465 14.7115 3.89525 14.2201 3.9016C14.1373 3.90796 14.0545 3.90796 13.9717 3.9016C13.9717 3.90796 13.9717 3.90796 13.9717 3.9016Z"
            fill="white"
          />
          <path
            d="M14.1384 16.1919C14.2776 16.1919 14.4105 16.1855 14.5498 16.1919C14.9295 16.2173 15.2017 16.4779 15.2207 16.8593C15.2397 17.158 15.2334 17.4504 15.2207 17.7491C15.2017 18.1178 14.9422 18.3784 14.5751 18.4038C14.2839 18.4229 13.9865 18.4229 13.689 18.4038C13.3156 18.3847 13.0434 18.1051 13.0244 17.7237C13.0117 17.4567 13.0181 17.1962 13.0244 16.9292C13.0371 16.4906 13.3156 16.2046 13.7586 16.1855C13.8852 16.1855 14.0118 16.1919 14.1384 16.1919Z"
            fill="#2193F7"
          />
          <path
            d="M11.0915 0.00635638C11.2307 0.00635638 11.3637 3.87845e-07 11.5029 0.00635638C11.8827 0.0317804 12.1548 0.292376 12.1738 0.673736C12.1928 0.972467 12.1865 1.26484 12.1738 1.56358C12.1548 1.93222 11.8953 2.19282 11.5282 2.21824C11.2371 2.23731 10.9396 2.23731 10.6421 2.21824C10.2687 2.19917 9.99651 1.91951 9.97752 1.53815C9.96486 1.2712 9.97119 1.0106 9.97752 0.743652C9.99018 0.305088 10.2687 0.019068 10.7117 0C10.8383 0 10.9649 3.87845e-07 11.0915 0.00635638Z"
            fill="#2193F7"
          />
          <path
            d="M11.0915 2.79737C11.2307 2.79737 11.3637 2.79102 11.5029 2.79737C11.8827 2.8228 12.1548 3.08339 12.1738 3.46475C12.1928 3.76348 12.1865 4.05586 12.1738 4.35459C12.1548 4.72324 11.8953 4.98383 11.5282 5.00926C11.2371 5.02833 10.9396 5.02833 10.6421 5.00926C10.2687 4.99019 9.99651 4.71053 9.97752 4.32917C9.96486 4.06221 9.97119 3.80162 9.97752 3.53467C9.99018 3.0961 10.2687 2.81008 10.7117 2.79102C10.8383 2.79102 10.9649 2.79737 11.0915 2.79737Z"
            fill="#2193F7"
          />
          <path
            d="M11.0915 5.86689C11.2307 5.86689 11.3637 5.86133 11.5029 5.86689C11.8827 5.88914 12.1548 6.11715 12.1738 6.45086C12.1928 6.71225 12.1865 6.96808 12.1738 7.22946C12.1548 7.55203 11.8953 7.78005 11.5282 7.80229C11.2371 7.81898 10.9396 7.81898 10.6421 7.80229C10.2687 7.78561 9.99651 7.5409 9.97752 7.20722C9.96486 6.97364 9.97119 6.74562 9.97752 6.51204C9.99018 6.12828 10.2687 5.87801 10.7117 5.86133C10.8383 5.86133 10.9649 5.86689 11.0915 5.86689Z"
            fill="#2193F7"
          />
          <path
            d="M11.0915 8.6587C11.2307 8.6587 11.3637 8.65234 11.5029 8.6587C11.8827 8.68412 12.1548 8.94472 12.1738 9.32608C12.1928 9.62481 12.1865 9.91719 12.1738 10.2159C12.1548 10.5846 11.8953 10.8452 11.5282 10.8706C11.2371 10.8897 10.9396 10.8897 10.6421 10.8706C10.2687 10.8515 9.99651 10.5719 9.97752 10.1905C9.96486 9.92354 9.97119 9.66295 9.97752 9.39599C9.99018 8.95743 10.2687 8.67141 10.7117 8.65234C10.8383 8.65234 10.9649 8.6587 11.0915 8.6587Z"
            fill="#2193F7"
          />
          <path
            d="M11.0915 11.4458C11.2307 11.4458 11.3637 11.4395 11.5029 11.4458C11.8827 11.4712 12.1548 11.7318 12.1738 12.1132C12.1928 12.4119 12.1865 12.7043 12.1738 13.003C12.1548 13.3717 11.8953 13.6323 11.5282 13.6577C11.2371 13.6768 10.9396 13.6768 10.6421 13.6577C10.2687 13.6386 9.99651 13.359 9.97752 12.9776C9.96486 12.7107 9.97119 12.4501 9.97752 12.1831C9.99018 11.7445 10.2687 11.4585 10.7117 11.4395C10.8383 11.4395 10.9649 11.4458 11.0915 11.4458Z"
            fill="#2193F7"
          />
          <path
            d="M17.1852 12.0056C17.3245 12.0056 17.4574 12 17.5967 12.0056C17.9764 12.0278 18.2486 12.2558 18.2676 12.5895C18.2866 12.8509 18.2802 13.1067 18.2676 13.3681C18.2486 13.6907 17.9891 13.9187 17.622 13.941C17.3308 13.9576 17.0333 13.9576 16.7359 13.941C16.3624 13.9243 16.0903 13.6796 16.0713 13.3459C16.0586 13.1123 16.0649 12.8843 16.0713 12.6507C16.0839 12.267 16.3624 12.0167 16.8055 12C16.9321 12 17.0587 12.0056 17.1852 12.0056Z"
            fill="#2193F7"
          />
          <path
            d="M17.1852 9.21455C17.3245 9.21455 17.4574 9.20898 17.5967 9.21455C17.9764 9.23679 18.2486 9.46481 18.2676 9.7985C18.2866 10.0599 18.2802 10.3157 18.2676 10.5771C18.2486 10.8997 17.9891 11.1277 17.622 11.1499C17.3308 11.1666 17.0333 11.1666 16.7359 11.1499C16.3624 11.1333 16.0903 10.8886 16.0713 10.5549C16.0586 10.3213 16.0649 10.0933 16.0713 9.85968C16.0839 9.47594 16.3624 9.22567 16.8055 9.20898C16.9321 9.20898 17.0587 9.21455 17.1852 9.21455Z"
            fill="#2193F7"
          />
          <path
            d="M17.1852 6.42353C17.3245 6.42353 17.4574 6.41797 17.5967 6.42353C17.9764 6.44578 18.2486 6.6738 18.2676 7.00749C18.2866 7.26888 18.2802 7.52471 18.2676 7.7861C18.2486 8.10866 17.9891 8.33669 17.622 8.35893C17.3308 8.37562 17.0333 8.37562 16.7359 8.35893C16.3624 8.34225 16.0903 8.09754 16.0713 7.76385C16.0586 7.53027 16.0649 7.30225 16.0713 7.06866C16.0839 6.68492 16.3624 6.43465 16.8055 6.41797C16.9321 6.41797 17.0587 6.42353 17.1852 6.42353Z"
            fill="#2193F7"
          />
          <path
            d="M17.1848 3.63526C17.324 3.63526 17.4569 3.62891 17.5962 3.63526C17.9759 3.66069 18.2481 3.92128 18.2671 4.30264C18.2861 4.60137 18.2797 4.89375 18.2671 5.19248C18.2481 5.56113 17.9886 5.82173 17.6215 5.84715C17.3303 5.86622 17.0329 5.86622 16.7354 5.84715C16.3619 5.82808 16.0898 5.54842 16.0708 5.16706C16.0581 4.90011 16.0645 4.63951 16.0708 4.37256C16.0834 3.93399 16.3619 3.64797 16.805 3.62891C16.9316 3.62891 17.0582 3.63526 17.1848 3.63526Z"
            fill="#2193F7"
          />
          <path
            d="M17.1852 0.842294C17.3245 0.842294 17.4574 0.835938 17.5967 0.842294C17.9764 0.867718 18.2486 1.12831 18.2676 1.50967C18.2866 1.8084 18.2802 2.10078 18.2676 2.39951C18.2486 2.76816 17.9891 3.02876 17.622 3.05418C17.3308 3.07325 17.0333 3.07325 16.7359 3.05418C16.3624 3.03511 16.0903 2.75545 16.0713 2.37409C16.0586 2.10714 16.0649 1.84654 16.0713 1.57959C16.0839 1.14103 16.3624 0.855005 16.8055 0.835938C16.9321 0.835938 17.0587 0.842294 17.1852 0.842294Z"
            fill="#2193F7"
          />
          <path
            d="M20.2316 4.1919C20.3709 4.1919 20.5038 4.18555 20.643 4.1919C21.0228 4.21733 21.295 4.47792 21.314 4.85927C21.3329 5.158 21.3266 5.45038 21.314 5.7491C21.295 6.11775 21.0355 6.37837 20.6684 6.40379C20.3772 6.42286 20.0797 6.42286 19.7822 6.40379C19.4088 6.38472 19.1366 6.10504 19.1177 5.72368C19.105 5.45673 19.1113 5.19614 19.1177 4.92919C19.1303 4.49063 19.4088 4.20461 19.8519 4.18555C19.9785 4.18555 20.105 4.1919 20.2316 4.1919Z"
            fill="#2193F7"
          />
          <path
            d="M17.1848 17.3052C17.324 17.3052 17.4569 17.2988 17.5962 17.3052C17.9759 17.3306 18.2481 17.5912 18.2671 17.9726C18.2861 18.2713 18.2797 18.5637 18.2671 18.8624C18.2481 19.2311 17.9886 19.4916 17.6215 19.5171C17.3303 19.5361 17.0329 19.5361 16.7354 19.5171C16.3619 19.498 16.0898 19.2183 16.0708 18.837C16.0581 18.57 16.0645 18.3094 16.0708 18.0425C16.0834 17.6039 16.3619 17.3179 16.805 17.2988C16.9316 17.2988 17.0582 17.2988 17.1848 17.3052Z"
            fill="#2193F7"
          />
          <path
            d="M26.3249 4.1919C26.4641 4.1919 26.5971 4.18555 26.7363 4.1919C27.1161 4.21733 27.3882 4.47792 27.4072 4.85927C27.4262 5.158 27.4199 5.45038 27.4072 5.7491C27.3882 6.11775 27.1287 6.37837 26.7616 6.40379C26.4705 6.42286 26.173 6.42286 25.8755 6.40379C25.5021 6.38472 25.2299 6.10504 25.2109 5.72368C25.1983 5.45673 25.2046 5.19614 25.2109 4.92919C25.2236 4.49063 25.5021 4.20461 25.9451 4.18555C26.0717 4.18555 26.1983 4.1919 26.3249 4.1919Z"
            fill="#2193F7"
          />
          <path
            d="M26.3249 6.98292C26.4641 6.98292 26.5971 6.97656 26.7363 6.98292C27.1161 7.00834 27.3882 7.26894 27.4072 7.6503C27.4262 7.94903 27.4199 8.24141 27.4072 8.54014C27.3882 8.90879 27.1287 9.16938 26.7616 9.19481C26.4705 9.21387 26.173 9.21387 25.8755 9.19481C25.5021 9.17574 25.2299 8.89607 25.2109 8.51471C25.1983 8.24776 25.2046 7.98717 25.2109 7.72021C25.2236 7.28165 25.5021 6.99563 25.9451 6.97656C26.0717 6.97656 26.1983 6.97656 26.3249 6.98292Z"
            fill="#2193F7"
          />
          <path
            d="M26.3249 18.4243C26.4641 18.4243 26.5971 18.418 26.7363 18.4243C27.1161 18.4497 27.3882 18.7103 27.4072 19.0917C27.4262 19.3904 27.4199 19.6828 27.4072 19.9815C27.3882 20.3502 27.1287 20.6108 26.7616 20.6362C26.4705 20.6553 26.173 20.6553 25.8755 20.6362C25.5021 20.6171 25.2299 20.3375 25.2109 19.9561C25.1983 19.6892 25.2046 19.4286 25.2109 19.1616C25.2236 18.7231 25.5021 18.437 25.9451 18.418C26.0717 18.418 26.1983 18.4243 26.3249 18.4243Z"
            fill="#2193F7"
          />
          <path
            d="M23.2775 11.1704C23.4168 11.1704 23.5497 11.1641 23.6889 11.1704C24.0687 11.1958 24.3409 11.4564 24.3599 11.8378C24.3788 12.1365 24.3725 12.4289 24.3599 12.7276C24.3409 13.0963 24.0814 13.3569 23.7143 13.3823C23.4231 13.4014 23.1256 13.4014 22.8281 13.3823C22.4547 13.3632 22.1825 13.0836 22.1636 12.7022C22.1509 12.4353 22.1572 12.1747 22.1636 11.9077C22.1762 11.4692 22.4547 11.1831 22.8978 11.1641C23.0244 11.1641 23.1509 11.1704 23.2775 11.1704Z"
            fill="#2193F7"
          />
          <path
            d="M20.2311 6.98292C20.3704 6.98292 20.5033 6.97656 20.6426 6.98292C21.0223 7.00834 21.2945 7.26894 21.3135 7.6503C21.3325 7.94903 21.3261 8.24141 21.3135 8.54014C21.2945 8.90879 21.035 9.16938 20.6679 9.19481C20.3767 9.21387 20.0792 9.21387 19.7818 9.19481C19.4083 9.17574 19.1362 8.89607 19.1172 8.51471C19.1045 8.24776 19.1108 7.98717 19.1172 7.72021C19.1298 7.28165 19.4083 6.99563 19.8514 6.97656C19.978 6.97656 20.1046 6.97656 20.2311 6.98292Z"
            fill="#2193F7"
          />
          <path
            d="M23.1256 7.53355C23.6609 7.53355 24.0949 7.09624 24.0949 6.55683C24.0949 6.01737 23.6609 5.58008 23.1256 5.58008C22.5903 5.58008 22.1562 6.01737 22.1562 6.55683C22.1562 7.09624 22.5903 7.53355 23.1256 7.53355Z"
            fill="#2193F7"
          />
          <path
            d="M8.77327 8.09375V9.12246C8.77327 9.12246 4.19245 12.7477 4.63576 16.0817C5.22683 20.5374 11.0021 20.7481 16.3402 19.5211V20.494C16.3402 20.494 7.17861 22.2664 4.48799 18.6659C1.80353 15.0592 5.78712 10.3371 8.77327 8.09375Z"
            fill="#2193F7"
          />
          <path
            d="M8.88764 5.02344V5.99654C8.88764 5.99654 0.263223 13.1825 1.71498 18.154C3.16674 23.1256 11.0899 24.9969 23.2637 19.5825V20.3061C23.2637 20.3061 6.29171 28.4464 0.767647 20.5556C0.767647 20.5618 -3.6922 14.7357 8.88764 5.02344Z"
            fill="#2193F7"
          />
          <path
            d="M19.6646 1.90914V3.37967C19.6646 3.37967 26.2424 1.50582 30.0912 4.42208C33.94 7.33836 28.9164 12.2898 28.9164 12.2898V14.7903C28.9164 14.7903 38.5354 9.16877 35.788 3.85745C32.704 -2.09917 19.8114 1.80986 19.6646 1.90914Z"
            fill="#2193F7"
          />
        </svg>

        <section
          class="flex items-center justify-between relative w-full mx-auto pr-4 pl-10 sm:px-11"
        >
          <button type="button" class="tab_btn" data-link-tab="faq">
            <svg
              width="18"
              height="18"
              viewBox="0 0 18 18"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <g clip-path="url(#clip0_330_168)">
                <circle
                  cx="9"
                  cy="9"
                  r="7.5"
                  stroke="white"
                  stroke-width="1.5"
                />
                <path
                  d="M7.5 6.75C7.5 5.92157 8.17157 5.25 9 5.25C9.82843 5.25 10.5 5.92157 10.5 6.75C10.5 7.04861 10.4127 7.32685 10.2623 7.56059C9.81405 8.25725 9 8.92157 9 9.75V10.125"
                  stroke="white"
                  stroke-width="1.5"
                  stroke-linecap="round"
                />
                <path
                  d="M8.994 12.75H9.00073"
                  stroke="white"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </g>
              <defs>
                <clipPath id="clip0_330_168">
                  <rect width="18" height="18" fill="white" />
                </clipPath>
              </defs>
            </svg>
          </button>
          <!--          <button class="tab_btn active" data-link-tab="support">-->
          <!--            <svg-->
          <!--              id="info"-->
          <!--              class="fill-white stroke-[#304A79]"-->
          <!--              width="24"-->
          <!--              height="24"-->
          <!--              viewBox="0 0 24 24"-->
          <!--              fill="none"-->
          <!--              xmlns="http://www.w3.org/2000/svg"-->
          <!--            >-->
          <!--              <circle cx="12" cy="12" r="10" stroke-width="1.5" />-->
          <!--              <path-->
          <!--                d="M10 9C10 7.89543 10.8954 7 12 7C13.1046 7 14 7.89543 14 9C14 9.39815 13.8837 9.76913 13.6831 10.0808C13.0854 11.0097 12 11.8954 12 13V13.5"-->
          <!--                stroke-width="1.5"-->
          <!--                stroke-linecap="round"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M11.992 17H12.001"-->
          <!--                stroke-width="2"-->
          <!--                stroke-linecap="round"-->
          <!--                stroke-linejoin="round"-->
          <!--              />-->
          <!--            </svg>-->
          <!--          </button>-->
          <!--          <button class="tab_btn">-->
          <!--            <svg-->
          <!--              id="head_phone"-->
          <!--              width="24"-->
          <!--              height="24"-->
          <!--              viewBox="0 0 24 24"-->
          <!--              fill="none"-->
          <!--              class="stroke-[#9cbdff]"-->
          <!--              xmlns="http://www.w3.org/2000/svg"-->
          <!--            >-->
          <!--              <path-->
          <!--                d="M17 10.8045C17 10.4588 17 10.286 17.052 10.132C17.2032 9.68444 17.6018 9.51076 18.0011 9.32888C18.45 9.12442 18.6744 9.02219 18.8968 9.0042C19.1493 8.98378 19.4022 9.03818 19.618 9.15929C19.9041 9.31984 20.1036 9.62493 20.3079 9.87302C21.2513 11.0188 21.7229 11.5918 21.8955 12.2236C22.0348 12.7334 22.0348 13.2666 21.8955 13.7764C21.6438 14.6979 20.8485 15.4704 20.2598 16.1854C19.9587 16.5511 19.8081 16.734 19.618 16.8407C19.4022 16.9618 19.1493 17.0162 18.8968 16.9958C18.6744 16.9778 18.45 16.8756 18.0011 16.6711C17.6018 16.4892 17.2032 16.3156 17.052 15.868C17 15.714 17 15.5412 17 15.1955V10.8045Z"-->
          <!--                stroke-width="1.5"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M7 10.8046C7 10.3694 6.98778 9.97821 6.63591 9.6722C6.50793 9.5609 6.33825 9.48361 5.99891 9.32905C5.55001 9.12458 5.32556 9.02235 5.10316 9.00436C4.43591 8.9504 4.07692 9.40581 3.69213 9.87318C2.74875 11.019 2.27706 11.5919 2.10446 12.2237C1.96518 12.7336 1.96518 13.2668 2.10446 13.7766C2.3562 14.6981 3.15152 15.4705 3.74021 16.1856C4.11129 16.6363 4.46577 17.0475 5.10316 16.996C5.32556 16.978 5.55001 16.8757 5.99891 16.6713C6.33825 16.5167 6.50793 16.4394 6.63591 16.3281C6.98778 16.0221 7 15.631 7 15.1957V10.8046Z"-->
          <!--                stroke-width="1.5"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M5 9C5 5.68629 8.13401 3 12 3C15.866 3 19 5.68629 19 9"-->
          <!--                stroke-width="1.5"-->
          <!--                stroke-linecap="square"-->
          <!--                stroke-linejoin="round"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M19 17V17.8C19 19.5673 17.2091 21 15 21H13"-->
          <!--                stroke-width="1.5"-->
          <!--                stroke-linecap="round"-->
          <!--                stroke-linejoin="round"-->
          <!--              />-->
          <!--            </svg>-->
          <!--          </button>-->
          <button type="button" class="tab_btn" data-link-tab="download-app">
            <svg
              class="stroke-[#9cbdff]"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M12 14.5L12 4.5M12 14.5C11.2998 14.5 9.99153 12.5057 9.5 12M12 14.5C12.7002 14.5 14.0085 12.5057 14.5 12"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
              <path
                d="M20 16.5C20 18.982 19.482 19.5 17 19.5H7C4.518 19.5 4 18.982 4 16.5"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
          <button type="button" class="tab_btn" data-link-tab="privacy">
            <svg
              class="stroke-[#9cbdff]"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M11.4706 22C7.47751 22 5.48098 22 4.24049 20.8284C3 19.6569 3 17.7712 3 14L3 10C3 6.22876 3 4.34315 4.24049 3.17157C5.48098 2 7.47752 2 11.4706 2L12.5294 2C16.5225 2 18.519 2 19.7595 3.17157C21 4.34315 21 6.22876 21 10M11.5 22H13"
                stroke-width="1.5"
                stroke-linecap="round"
              />
              <path d="M8 7H16" stroke-width="1.5" stroke-linecap="round" />
              <path d="M8 12H13" stroke-width="1.5" stroke-linecap="round" />
              <path
                d="M17.5 18.5896L16.5978 21.7428C16.5572 21.9011 16.7139 22.0385 16.8659 21.9778L18.8514 21.1849C18.9468 21.1468 19.0532 21.1468 19.1486 21.1849L21.1531 21.9854C21.3014 22.0446 21.456 21.9149 21.4231 21.7589L20.6589 18.4911M22 15.9951C22 14.341 20.6569 13 19 13C17.3431 13 16 14.341 16 15.9951C16 17.6493 17.3431 18.9902 19 18.9902C20.6569 18.9902 22 17.6493 22 15.9951Z"
                stroke-width="1.5"
                stroke-linejoin="round"
              />
            </svg>
          </button>
          <!--          <button class="tab_btn" data-link-tab="guide">-->
          <!--            <svg-->
          <!--              class="stroke-[#9cbdff]"-->
          <!--              width="24"-->
          <!--              height="24"-->
          <!--              viewBox="0 0 24 24"-->
          <!--              fill="none"-->
          <!--              xmlns="http://www.w3.org/2000/svg"-->
          <!--            >-->
          <!--              <circle cx="12" cy="12" r="10" stroke-width="1.5" />-->
          <!--              <path-->
          <!--                d="M11.992 15H12.001"-->
          <!--                stroke-width="2"-->
          <!--                stroke-linecap="round"-->
          <!--                stroke-linejoin="round"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M12 12L12 8"-->
          <!--                stroke-width="1.5"-->
          <!--                stroke-linecap="round"-->
          <!--                stroke-linejoin="round"-->
          <!--              />-->
          <!--            </svg>-->
          <!--          </button>-->
          <!--          <button class="tab_btn" data-link-tab="multi-lang">-->
          <!--            <svg-->
          <!--              class="stroke-[#9cbdff]"-->
          <!--              width="24"-->
          <!--              height="24"-->
          <!--              viewBox="0 0 24 24"-->
          <!--              fill="none"-->
          <!--              xmlns="http://www.w3.org/2000/svg"-->
          <!--            >-->
          <!--              <path-->
          <!--                d="M2 12C2 13.0519 2.18046 14.0617 2.51212 15M13.0137 9H21.5015M11 15H2.51212M21.5015 9C20.266 5.50442 16.9323 3 13.0137 3C14.6146 3 15.9226 6.76201 16.0091 11.5M21.5015 9C21.7803 9.78867 21.9522 10.6278 22 11.5M2.51212 15C3.74763 18.4956 7.08134 21 11 21C9.45582 21 8.18412 17.5 8.01831 13"-->
          <!--                stroke-width="1.5"-->
          <!--                stroke-linecap="round"-->
          <!--                stroke-linejoin="round"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M2 5.29734C2 4.19897 2 3.64979 2.18738 3.22389C2.3861 2.77223 2.72861 2.40921 3.15476 2.1986C3.55661 2 4.07478 2 5.11111 2H6C7.88562 2 8.82843 2 9.41421 2.62085C10 3.2417 10 4.24095 10 6.23944V8.49851C10 9.37026 10 9.80613 9.73593 9.95592C9.47186 10.1057 9.12967 9.86392 8.4453 9.38036L8.34103 9.30669C7.84086 8.95329 7.59078 8.77658 7.30735 8.68563C7.02392 8.59468 6.72336 8.59468 6.12223 8.59468H5.11111C4.07478 8.59468 3.55661 8.59468 3.15476 8.39608C2.72861 8.18547 2.3861 7.82245 2.18738 7.37079C2 6.94489 2 6.39571 2 5.29734Z"-->
          <!--                stroke-width="1.5"-->
          <!--              />-->
          <!--              <path-->
          <!--                d="M22 17.2973C22 16.199 22 15.6498 21.8126 15.2239C21.6139 14.7722 21.2714 14.4092 20.8452 14.1986C20.4434 14 19.9252 14 18.8889 14H18C16.1144 14 15.1716 14 14.5858 14.6209C14 15.2417 14 16.2409 14 18.2394V20.4985C14 21.3703 14 21.8061 14.2641 21.9559C14.5281 22.1057 14.8703 21.8639 15.5547 21.3804L15.659 21.3067C16.1591 20.9533 16.4092 20.7766 16.6926 20.6856C16.9761 20.5947 17.2766 20.5947 17.8778 20.5947H18.8889C19.9252 20.5947 20.4434 20.5947 20.8452 20.3961C21.2714 20.1855 21.6139 19.8225 21.8126 19.3708C22 18.9449 22 18.3957 22 17.2973Z"-->
          <!--                stroke-width="1.5"-->
          <!--              />-->
          <!--            </svg>-->
          <!--          </button>-->
          <div
            id="mobile_line_indicator"
            class="h-0.5 bg-white absolute -bottom-1.5 right-4 sm:right-11 transition-all duration-500 ease-in-out rounded-full"
          ></div>
        </section>
      </div>
      <section
        id="drawer_content"
        class="pt-[65px] mx-auto overflow-y-auto h-full pb-20 hidden"
      >
        <div
          class="tab-content flex flex-col items-center justify-center active"
          data-tab-name="faq"
        >
          <h1 class="flex items-center justify-center text-lg font-bold">
            ســــوالات متــــداول
          </h1>
          <p
            class="flex items-center justify-center text-sm opacity-60 font-medium mt-2 mb-8"
          >
            هر آنچه در مورد پنجره ملی و خدمات آن میخواهید بدانید:
          </p>
          <section
            id="mobile_tabs_wrapper"
            class="flex items-center justify-between relative w-full px-[17px] mt-9 text-[12px]"
          >
            <button
              type="button"
              class="text-white opacity-100 cursor-pointer mobile-tab-btn active-mobile-tab"
            >
              عمومی
            </button>
            <button
              type="button"
              class="text-white opacity-70 cursor-pointer mobile-tab-btn"
            >
              تخصصی
            </button>
            <button
              type="button"
              class="text-white opacity-70 text-nowrap cursor-pointer mobile-tab-btn"
            >
              وزارت خانه‌ها
            </button>
            <button
              type="button"
              class="text-white opacity-70 cursor-pointer mobile-tab-btn"
            >
              سایر
            </button>
            <div
              id="child_line_indicator"
              class="h-0.5 bg-white absolute -bottom-1 right-[15px] transition-all duration-500 ease-in-out rounded-full"
            ></div>
          </section>
          <section class="px-[17px] mt-9 w-full text-[12px]">
            <!-- faq -->
            <div
              class="faq-tab-content block mobile-tab-content active-mobile-tab flex flex-col items-start w-full"
              data-category="614193452171d0294fea7f10"
            ></div>
            <!-- ................................................... -->
            <div
              class="faq-tab-content hidden mobile-tab-content"
              data-category="646df7a1a8a9c325f21b957c"
            ></div>
            <div
              class="faq-tab-content hidden mobile-tab-content"
              data-category="68be7493bf4e0571d7163b62"
            ></div>
            <div
              class="faq-tab-content hidden mobile-tab-content"
              data-category="68be75035515b84e6e96da1c"
            ></div>
          </section>
        </div>
        <div class="hidden tab-content overflow-y-auto" data-tab-name="support">
          <div id="support-content overflow-y-auto">
            <section
              class="flex flex-col items-center justify-center gap-6 text-white"
            >
              <p class="flex items-center justify-center text-lg font-bold">
                پــــشتیــــبانی 24/7
              </p>
              <p
                class="flex items-center justify-center text-center text-sm opacity-60 font-medium mt-2 mb-8 px-3"
              >
                کاربران گرامی می‌توانند در هر ساعت از شبانه‌روز نسبت به ثبت
                درخواست اقدام نموده و پاسخ کارشناسان را دریافت نمایند.
              </p>
            </section>
            <section
              class="px-8 mx-6 mt-4 before:bg-[linear-gradient(360deg,rgba(61,95,156,0)0%,#A1B4D6_100%)] bg-[linear-gradient(360deg,#304A79_8.85%,rgba(68,105,173,0.35)100%)] rounded-3xl linearBorder before:rounded-3xl pt-12 pb-11"
            >
              <form action="#" class="space-y-6 flex flex-col items-center">
                <input
                  class="bg-transparent outline-none border-b border-[#7C96C6] placeholder:text-[#7C96C6] text-sm font-medium focus:outline-none w-full pb-2 pr-2"
                  placeholder="نام و نام خانوادگی"
                />
                <input
                  class="bg-transparent outline-none border-b border-[#7C96C6] placeholder:text-[#7C96C6] text-sm font-medium focus:outline-none w-full pb-2 pr-2"
                  placeholder="موضوع"
                />
                <div class="custom-select w-full">
                  <select class="select-hidden" name="support-type">
                    <option value="" disabled selected>نوع پشتیبانی</option>
                    <option value="1">گزینه ۱</option>
                    <option value="2">گزینه ۲</option>
                    <option value="3">گزینه ۳</option>
                    <option value="4">گزینه ۴</option>
                  </select>
                  <div
                    class="select-selected flex justify-between items-center bg-transparent text-sm border-0 border-b border-[#7C96C6] text-[#7C96C6] pr-2 pb-2 cursor-pointer outline-none"
                  >
                    نوع پشتیبانی
                    <svg
                      class="w-6 h-6 fill-none stroke-[#7C96C6] stroke-2"
                      viewBox="0 0 24 24"
                    >
                      <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                  </div>
                  <div
                    class="select-items hidden absolute top-full right-0 left-0 rounded-md overflow-hidden z-10"
                  >
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="1"
                    >
                      گزینه ۱
                    </div>
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="2"
                    >
                      گزینه ۲
                    </div>
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="3"
                    >
                      گزینه ۳
                    </div>
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="4"
                    >
                      گزینه ۴
                    </div>
                  </div>
                </div>
                <div class="custom-select w-full">
                  <select class="select-hidden" name="priority-level">
                    <option value="" disabled selected>درجه اهمیت</option>
                    <option value="1">گزینه ۱</option>
                    <option value="2">گزینه ۲</option>
                    <option value="3">گزینه ۳</option>
                    <option value="4">گزینه ۴</option>
                  </select>
                  <div
                    class="select-selected flex justify-between items-center bg-transparent text-sm border-0 border-b border-[#7C96C6] text-[#7C96C6] pr-2 pb-2 cursor-pointer outline-none"
                  >
                    درجه اهمیت
                    <svg
                      class="w-6 h-6 fill-none stroke-[#7C96C6] stroke-2"
                      viewBox="0 0 24 24"
                    >
                      <polyline points="6 9 12 15 18 9"></polyline>
                    </svg>
                  </div>
                  <div
                    class="select-items hidden absolute top-full right-0 left-0 rounded-md overflow-hidden z-10"
                  >
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="1"
                    >
                      گزینه ۱
                    </div>
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="2"
                    >
                      گزینه ۲
                    </div>
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="3"
                    >
                      گزینه ۳
                    </div>
                    <div
                      class="text-[#9CBDFF] p-2 cursor-pointer hover:bg-white/20 text-right"
                      data-value="4"
                    >
                      گزینه ۴
                    </div>
                  </div>
                </div>
                <textarea
                  name=""
                  id=""
                  placeholder="توضیحات"
                  class="w-full bg-[#ECF4FF1A] placeholder:text-[#7C96C6] p-2 h-[127px] text-sm outline-none focus:outline-none border-none"
                ></textarea>
                <button
                  type="button"
                  class="mt-2 w-full bg-white shadow-[4px_4px_8px_0px_#00000040] text-[#3D5F9C] py-2 flex items-center justify-center rounded-[10px] font-medium"
                >
                  ارسال کد
                </button>
              </form>
            </section>
            <section
              class="px-6 mt-2 grid grid-cols-3 justify-between w-full gap-2"
            >
              <div
                class="before:bg-[linear-gradient(193.24deg,rgba(61,95,156,0)41.8%,#A1B4D6_101.77%)] linearBorder drop-shadow-[-10px_-5px_32px_0px_#00000066] rounded-2xl before:rounded-2xl p-2 bg-[linear-gradient(201.67deg,#304A79_34.36%,#4469AD_95.69%)]"
              >
                <div
                  class="w-[30px] h-[30px] bg-[#4469AD] flex items-center justify-center rounded-full"
                >
                  <svg
  width="31"
  height="31"
  viewBox="0 0 31 31"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g clip-path="url(#clip0_1494_69546)">
    <path
      d="M7.74867 15.1244L10.0584 14.4891C10.7099 14.3099 11.3245 14.2546 11.9814 14.3131C12.7356 14.3802 13.1803 14.1172 13.4722 13.3153C14.2126 11.2812 18.9898 12.7748 18.1707 15.0254C17.8788 15.8273 18.0505 16.3147 18.6714 16.748C19.2165 17.1285 19.6505 17.5673 20.0313 18.1189L21.3923 20.0903C22.0081 20.9823 22.2909 21.408 22.3415 21.9398C22.3891 22.4401 22.2071 22.9403 21.843 23.9405C21.1885 25.7388 20.8612 26.6379 20.1479 27.0311C19.284 27.5073 17.7204 26.8703 16.7465 26.5158L7.1773 23.0329C6.20344 22.6784 4.615 22.21 4.24045 21.2412C3.94675 20.4816 4.27401 19.5824 4.92853 17.7842C5.29259 16.7839 5.47462 16.2838 5.83268 15.9311C6.21326 15.5562 6.70357 15.4119 7.74867 15.1244Z"
      stroke="#ECF0F7"
      stroke-width="1.5"
    />
    <path
      d="M15.5498 20.7631C15.172 21.801 14.0244 22.3362 12.9864 21.9584C11.9484 21.5807 11.4133 20.433 11.7911 19.395C12.1688 18.3571 13.3165 17.8219 14.3545 18.1997C15.3924 18.5775 15.9276 19.7251 15.5498 20.7631Z"
      stroke="#ECF0F7"
      stroke-width="1.5"
    />
    <path
      d="M13.4842 5.85602C12.1165 5.75881 11.0661 5.87343 10.0583 6.10622C8.46279 6.47478 7.35976 7.78392 6.98115 9.29675C6.82113 9.93613 7.20844 10.4107 7.83989 10.46C8.31962 10.4975 8.79887 10.5385 9.27856 10.5677C10.6814 10.653 11.1454 10.3048 11.7931 9.07276L13.4842 5.85602ZM13.4842 5.85602C16.8643 6.09628 20.2123 7.31484 22.956 9.30349M22.956 9.30349C24.0662 10.1081 24.7972 10.8711 25.4196 11.6973C26.4049 13.0052 26.4084 14.7171 25.726 16.1193C25.4376 16.712 24.8359 16.8265 24.3204 16.4585C23.9288 16.1788 23.5354 15.9022 23.1492 15.6162C22.0197 14.7797 21.888 14.2148 22.1838 12.8546L22.956 9.30349Z"
      stroke="#ECF0F7"
      stroke-width="1.5"
      stroke-linejoin="round"
    />
  </g>
</svg>
                </div>
                <p class="text-[12px] font-normal text-[#DDEAFF] mt-[17px]">
                  تلفن
                </p>
                <p class="text-sm font-normal text-white mt-1">021-91091516</p>
              </div>
              <div
                class="before:bg-[linear-gradient(193.24deg,rgba(61,95,156,0)41.8%,#A1B4D6_101.77%)] linearBorder drop-shadow-[-10px_-5px_32px_0px_#00000066] rounded-2xl before:rounded-2xl p-2 bg-[linear-gradient(201.67deg,#304A79_34.36%,#4469AD_95.69%)]"
              >
                <div
                  class="w-[30px] h-[30px] bg-[#4469AD] flex items-center justify-center rounded-full"
                >
                  <svg
  width="31"
  height="31"
  viewBox="0 0 31 31"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g clip-path="url(#clip0_1494_69585)">
    <path
      d="M11.8792 10.3809L14.0488 13.0216C15.3137 14.5612 15.958 14.7957 17.9166 14.4294L21.2761 13.8011"
      stroke="white"
      stroke-width="1.5"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
    <path
      d="M5.49376 13.3521C4.50672 16.2551 4.0132 17.7066 4.68777 19.1604C5.36234 20.6142 6.82812 21.1898 9.75969 22.341C11.5665 23.0505 13.372 23.7077 15.2121 24.3255C18.1978 25.3281 19.6906 25.8293 21.1419 25.1492C22.5931 24.4691 23.1481 23.04 24.258 20.1817C24.6149 19.2626 24.95 18.3419 25.2674 17.4084C26.2544 14.5054 26.7479 13.0539 26.0734 11.6001C25.3988 10.1463 23.933 9.57067 21.0014 8.41948C19.1947 7.70998 17.3891 7.05283 15.549 6.43495C12.5633 5.43244 11.0705 4.93118 9.61925 5.61127C8.168 6.29135 7.61305 7.7205 6.50314 10.5788C6.14627 11.4979 5.81113 12.4186 5.49376 13.3521Z"
      stroke="white"
      stroke-width="1.5"
      stroke-linejoin="round"
    />
  </g>
</svg>
                </div>
                <p class="text-[12px] font-normal text-[#DDEAFF] mt-[17px]">
                  ایمیل
                </p>
                <p
                  class="text-[12px] font-normal text-white mt-1 w-full overflow-hidden"
                >
                  example@gmail.com
                </p>
              </div>
              <div
                class="before:bg-[linear-gradient(193.24deg,rgba(61,95,156,0)41.8%,#A1B4D6_101.77%)] linearBorder drop-shadow-[-10px_-5px_32px_0px_#00000066] rounded-2xl before:rounded-2xl p-2 bg-[linear-gradient(201.67deg,#304A79_34.36%,#4469AD_95.69%)]"
              >
                <div
                  class="w-[30px] h-[30px] bg-[#4469AD] flex items-center justify-center rounded-full"
                >
                  <svg
  width="31"
  height="31"
  viewBox="0 0 31 31"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g clip-path="url(#clip0_1495_69538)">
    <path
      d="M13.4615 4.57316C12.3239 4.35936 11.4301 4.40629 10.6037 4.79745C8.81499 5.64404 8.04032 7.77243 6.49097 12.0292C4.94163 16.286 4.16695 18.4144 4.99297 20.2127C5.819 22.0109 7.92313 22.7767 12.1314 24.3084C16.3397 25.8401 18.4438 26.606 20.2325 25.7594C22.0211 24.9128 22.7958 22.7844 24.3451 18.5276C25.8945 14.2708 26.6691 12.1424 25.8431 10.3442C25.4615 9.51329 24.8069 8.90283 23.7981 8.33538"
      stroke="white"
      stroke-width="1.5"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
    <path
      d="M15.4255 7.94884C16.0603 7.64175 17.9718 6.21517 18.6297 6.45466M20.1239 9.65894C19.835 9.01563 19.2877 6.69415 18.6297 6.45466M18.6297 6.45466L15.8936 13.9722"
      stroke="white"
      stroke-width="1.5"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
    <path
      d="M23.7947 20.0394L19.166 18.3547C18.3747 18.0667 17.5123 18.5016 16.9092 19.0734C16.2541 19.6947 15.2406 20.1185 13.8415 19.6093C12.4424 19.1001 11.9385 18.124 11.8359 17.2269C11.7415 16.4012 11.3604 15.5137 10.5691 15.2257L5.94049 13.541"
      stroke="white"
      stroke-width="1.5"
      stroke-linejoin="round"
    />
  </g>
</svg>
                </div>
                <p class="text-[12px] font-normal text-[#DDEAFF] mt-[17px]">
                  فکس
                </p>
                <p class="text-sm font-normal text-white mt-1">021-91091516</p>
              </div>
            </section>
          </div>
        </div>
        <div class="hidden tab-content" data-tab-name="download-app">
          <p class="flex items-center justify-center text-lg font-bold">
            دانــــلود اپلیــــکیشن
          </p>
          <p
            class="flex items-center justify-center text-center text-sm opacity-60 font-medium mt-2 mb-8 px-3"
          >
            هر آنچه در مورد پنجره ملی و خدمات آن میخواهید بدانید:
          </p>
          <section class="flex flex-col items-center gap-4 px-10">
            <a
              class="linearBorder justify-between w-full h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
              href="https://file.my.gov.ir/static/app/MyGov.apk"
            >
              <div class="flex items-center gap-5 w-full">
                <svg
                  width="32"
                  height="36"
                  viewBox="0 0 18 18"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M9 10.875L9 3.375M9 10.875C8.47483 10.875 7.49365 9.37927 7.125 9M9 10.875C9.52517 10.875 10.5064 9.37927 10.875 9"
                    stroke="white"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                  <path
                    d="M15 12.375C15 14.2365 14.6115 14.625 12.75 14.625H5.25C3.3885 14.625 3 14.2365 3 12.375"
                    stroke="white"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>

                <div class="mt-2">
                  <p class="text-sm">دریافت نسخه اندروید از</p>
                  <p class="text-[23px] font-extrabold mt-1">دانلود مستقیم</p>
                </div>
              </div>
              <svg
                width="32"
                height="32"
                viewBox="0 0 32 32"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
                  fill="white"
                />
              </svg>
            </a>
            <div class="h-[1px] bg-[#A1B4D680] w-full"></div>
            <a
              class="linearBorder justify-between w-full h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
              href="https://cafebazaar.ir/app/ir.gov.mygov/"
            >
              <div class="flex items-center gap-5 w-full">
                <svg
                  width="32"
                  height="36"
                  viewBox="0 0 40 40"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M37.0207 14.7852C36.8435 15.0691 36.5766 15.3951 36.1752 15.7066C35.1133 16.5311 33.5512 16.9461 31.5197 16.9461C31.0337 16.9461 30.521 16.9224 29.9817 16.8748C23.0416 16.2624 16.6581 16.1791 11.5211 16.6338C7.26225 17.0108 4.11764 16.5147 2.66649 15.2371C2.57141 15.1533 2.48443 15.0663 2.40413 14.9764C2.39387 15.1914 2.40212 15.4161 2.43076 15.6514C2.56132 16.7238 2.69648 17.798 2.89345 18.8596C3.43674 21.7876 4.12326 24.6826 5.00854 27.5282C5.77108 29.9792 6.56081 32.4206 7.83888 34.6656C8.2942 35.4653 8.77654 36.2643 9.59851 36.7487C10.1999 37.1031 10.818 37.4413 11.4599 37.7127C14.0801 38.8206 16.863 39.0557 19.8271 39.1992C20.5632 39.1381 21.4615 39.0581 22.3607 38.9899C24.1892 38.8513 25.9659 38.4766 27.6935 37.8599C28.7314 37.4893 29.5615 36.8686 30.2226 35.9779C31.4036 34.3868 32.2264 32.6084 32.9761 30.7963C34.5585 26.9715 35.7433 23.0203 36.6121 18.9743C36.8856 17.7005 37.1443 16.4203 37.055 15.1072C37.0476 14.9981 37.036 14.8909 37.0207 14.7852Z"
                    fill="white"
                  />
                  <path
                    d="M35.9059 12.7396C34.7856 11.8063 33.4718 11.2313 32.0908 10.8727C30.3303 10.4156 28.534 10.0958 26.7515 9.72513C26.448 9.66201 26.2688 9.53706 26.1351 9.22194C25.4024 7.49481 24.676 5.76225 23.8566 4.07568C23.3603 3.05407 22.6893 2.12042 21.748 1.42848C21.2029 1.02781 20.5845 0.750019 19.9169 0.808573C19.1244 0.878103 18.4392 1.25361 17.847 1.81238C16.8226 2.77895 16.0672 3.93852 15.45 5.1802C14.7916 6.50497 14.2313 7.87826 13.6081 9.22105C13.5288 9.39205 13.3464 9.60093 13.1813 9.63248C12.1198 9.83527 11.0418 9.95507 9.98427 10.1748C8.00596 10.5857 6.06507 11.1189 4.29927 12.1546C3.41769 12.6717 2.80515 13.3303 2.54297 14.1722C2.66145 14.4016 2.82964 14.6133 3.04779 14.8053C4.35987 15.9606 7.43009 16.418 11.4706 16.0605C11.9549 16.0176 12.4505 15.9797 12.9561 15.9464C13.1818 12.901 15.0074 9.59839 15.0951 9.44057L16.0696 9.35364C16.0493 9.39031 14.0408 13.0517 13.7868 15.8963C17.3992 15.699 21.5038 15.733 25.9059 15.9958C25.7278 13.111 23.8462 9.41833 23.8462 9.41833L24.81 9.41692C24.8927 9.57778 26.5818 12.9734 26.7337 16.0474C27.8173 16.1186 28.9178 16.2032 30.0326 16.3015C31.8284 16.4601 34.2835 16.412 35.781 15.2836C36.5732 14.6867 36.7534 14.0492 36.7835 13.9161C36.5965 13.4736 36.3108 13.0771 35.9059 12.7396ZM16.1858 9.31248C16.1858 9.29395 16.1685 9.23537 16.1883 9.19508C16.8182 7.91445 17.489 6.66361 18.529 5.64991C19.6567 4.55076 20.6214 4.71238 21.5902 5.72275C22.5077 6.67948 23.103 7.86505 23.6495 9.06586C23.6863 9.14686 23.6861 9.24476 23.6989 9.31248C21.187 9.31248 18.6958 9.31248 16.1858 9.31248ZM25.143 9.58839C25.0127 9.44612 24.8375 9.32495 24.7593 9.15825C24.1703 7.90384 23.5844 6.64683 22.7088 5.55971C22.4101 5.18889 22.077 4.82649 21.6974 4.54375C20.453 3.61683 19.0625 3.7726 17.8741 4.91372C17.0429 5.71185 16.4346 6.66948 15.884 7.66718C15.6132 8.15795 15.3702 8.66465 15.13 9.17142C15.004 9.43707 14.8411 9.55872 14.4908 9.4024C14.7257 8.83622 14.9511 8.27003 15.1943 7.71155C15.8248 6.26354 16.5457 4.86383 17.4676 3.57623C17.9328 2.92639 18.4519 2.32881 19.1727 1.94294C19.9291 1.53805 20.623 1.66782 21.2756 2.17971C22.0568 2.79248 22.6418 3.58174 23.0647 4.45412C23.8402 6.05423 24.5288 7.69652 25.2512 9.32232C25.2716 9.36803 25.2644 9.42595 25.2701 9.47818C25.2277 9.51491 25.1853 9.55164 25.143 9.58839Z"
                    fill="white"
                  />
                </svg>

                <div class="mt-2">
                  <p class="text-sm">دریافت نسخه اندروید از</p>
                  <svg
                    width="48"
                    height="28"
                    viewBox="0 0 48 28"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    class="mx-auto"
                  >
                    <path
                      d="M1.36795e-05 17.5018C0.0449225 17.3814 0.0954738 17.2624 0.133868 17.1404C0.474276 16.0578 1.52036 15.3084 2.76522 15.2957C4.2772 15.2802 5.7895 15.2912 7.30166 15.2899C8.14703 15.2892 8.49442 14.9794 8.49468 14.2252C8.49545 12.0433 8.48955 9.86143 8.49735 7.67959C8.50225 6.30926 9.54951 5.2622 11.0464 5.10692C12.5617 4.94974 14.0157 6.03825 14.0974 7.41533C14.1334 8.02144 14.1129 8.63041 14.1131 9.23808C14.114 11.2605 14.1288 13.2831 14.1081 15.3054C14.0827 17.7845 12.1706 19.8258 9.47691 20.2712C9.11029 20.3318 8.73109 20.3514 8.3575 20.3531C6.59151 20.3614 4.82544 20.3577 3.05941 20.3566C1.45566 20.3556 0.44995 19.6424 0.0627207 18.2363C0.0496796 18.189 0.0212531 18.1451 0 18.0996C1.19534e-05 17.9003 1.36795e-05 17.7011 1.36795e-05 17.5018Z"
                      fill="white"
                    />
                    <path
                      d="M33.918 11.5277C33.918 9.36596 33.8993 7.20403 33.9237 5.04251C33.9443 3.22128 36.0131 2.02554 37.8524 2.76898C38.9526 3.21369 39.5171 4.02083 39.5284 5.11059C39.5441 6.62468 39.5327 8.139 39.5328 9.65322C39.5328 12.0042 39.5327 14.3553 39.5328 16.7063C39.5329 17.5446 39.8504 17.8279 40.7893 17.828C42.2241 17.8283 43.6589 17.8193 45.0936 17.8307C46.814 17.8444 48.0439 18.9736 47.9989 20.478C47.9619 21.7145 46.8479 22.801 45.4788 22.8782C44.643 22.9253 43.8021 22.9015 42.9635 22.9026C41.7605 22.9041 40.5565 22.9264 39.3547 22.8915C36.724 22.8149 34.5404 21.1355 34.0292 18.8121C33.9482 18.4439 33.9227 18.0607 33.9208 17.6842C33.9107 15.6321 33.916 13.5799 33.916 11.5277L33.918 11.5277Z"
                      fill="white"
                    />
                    <path
                      d="M31.0587 16.6535C31.0587 17.859 31.0618 19.0645 31.0581 20.27C31.049 23.2151 28.6035 25.4339 25.3454 25.4458C23.458 25.4527 21.5704 25.4591 19.6832 25.4386C18.4313 25.4249 17.3701 24.6341 17.0617 23.5438C16.7494 22.4395 17.2811 21.2952 18.3728 20.7232C18.8327 20.4822 19.3303 20.3775 19.8605 20.3784C21.3175 20.3808 22.7744 20.3796 24.2314 20.3789C25.1124 20.3784 25.4417 20.0771 25.442 19.2721C25.4426 17.0903 25.4294 14.9083 25.4472 12.7266C25.4611 11.0294 27.1986 9.8447 28.9951 10.2801C30.191 10.57 31.0386 11.5383 31.0531 12.6784C31.0698 14.0033 31.057 15.3284 31.057 16.6535H31.0587Z"
                      fill="white"
                    />
                    <path
                      d="M22.5851 8.89035C22.5853 11.0114 22.6 13.1326 22.5809 15.2535C22.5642 17.1083 20.5362 18.326 18.6827 17.5981C17.6291 17.1844 17.0681 16.428 16.9764 15.3949C16.9562 15.1673 16.9686 14.9371 16.9686 14.7082C16.9684 10.6851 16.9663 6.66211 16.9698 2.63909C16.971 1.2851 17.9057 0.259774 19.3239 0.0375143C20.9972 -0.224741 22.5518 0.920163 22.5771 2.46753C22.6023 4.01069 22.5841 5.55442 22.5848 7.0979C22.585 7.69539 22.5848 8.29287 22.5851 8.89035Z"
                      fill="white"
                    />
                    <path
                      d="M28.2522 2.55469C29.8077 2.55407 31.0536 3.6739 31.057 5.07559C31.0603 6.4854 29.7941 7.62682 28.2357 7.61881C26.6926 7.61089 25.4404 6.46998 25.4443 5.07567C25.4483 3.67749 26.6982 2.5553 28.2522 2.55469Z"
                      fill="white"
                    />
                    <path
                      d="M36.7028 25.4602C36.7046 26.8583 35.4584 27.9924 33.9108 28.0011C32.3648 28.0098 31.0773 26.8477 31.086 25.4514C31.0947 24.0566 32.352 22.934 33.904 22.9355C35.4611 22.9371 36.7011 24.0556 36.7028 25.4602Z"
                      fill="white"
                    />
                  </svg>
                </div>
              </div>
              <svg
                width="32"
                height="32"
                viewBox="0 0 32 32"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
                  fill="white"
                />
              </svg>
            </a>
            <div class="h-[1px] bg-[#A1B4D680] w-full"></div>
            <a
              class="linearBorder justify-between w-full h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
              href="https://sibche.com/applications/mygov/"
            >
              <div class="flex items-center gap-5 w-full">
                <svg
  width="32"
  height="36"
  viewBox="0 0 40 40"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
>
  <rect
    x="2.3999"
    width="35.2667"
    height="39.9245"
    fill="url(#pattern0_2842_101903)"
  />
</svg>

                <div class="mt-2">
                  <p class="text-sm">دریافت نسخه iOS از</p>

                  <svg
  width="98"
  height="33"
  viewBox="0 0 98 33"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  class="mx-auto"
>
  <rect width="98" height="32.1235" fill="url(#pattern0_2842_101973)" />
</svg>
                </div>
              </div>
              <svg
                width="32"
                height="32"
                viewBox="0 0 32 32"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
                  fill="white"
                />
              </svg>
            </a>
            <div class="h-[1px] bg-[#A1B4D680] w-full"></div>
            <a
              class="linearBorder justify-between w-full h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
              href="https://iapps.ir/app/%D9%BE%D9%86%D8%AC%D8%B1%D9%87-%D9%85%D9%84%DB%8C-%D8%AE%D8%AF%D9%85%D8%A7%D8%AA-%D8%AF%D9%88%D9%84%D8%AA-%D9%87%D9%88%D8%B4%D9%85%D9%86%D8%AF-%D8%AF%D9%88%D9%84%D8%AA-%D9%85%D9%86-/964465353"
            >
              <div class="flex items-center gap-5 w-full">
                <svg
  width="32"
  height="36"
  viewBox="0 0 40 40"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
>
  <g clip-path="url(#clip0_2842_102012)">
    <rect
      y="-2"
      width="40"
      height="40"
      rx="8"
      fill="url(#pattern0_2842_102012)"
    />
  </g>
</svg>

                <div class="mt-2">
                  <p class="text-sm">دریافت نسخه iOS از</p>
                  <svg
  width="122"
  height="24"
  viewBox="0 0 122 24"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  class="mx-auto"
>
  <rect width="122" height="24" fill="url(#pattern0_2842_102062)" />
</svg>
                </div>
              </div>
              <svg
                width="32"
                height="32"
                viewBox="0 0 32 32"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
                  fill="white"
                />
              </svg>
            </a>
            <div class="h-[1px] bg-[#A1B4D680] w-full"></div>
            <a
              class="linearBorder justify-between w-full h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
              href="https://myket.ir/app/ir.gov.mygov"
            >
              <div class="flex items-center gap-5 w-full">
                <svg
                  width="32"
                  height="36"
                  viewBox="0 0 40 40"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M38.5888 8.38785C37.1551 5.97741 34.2879 3.2054 28.6728 3.2054C27.066 3.15183 25.4665 3.44691 23.9826 4.07069C22.4986 4.69446 21.1646 5.63239 20.0709 6.82106C18.9772 5.63239 17.6433 4.69446 16.1593 4.07069C14.6753 3.44691 13.0759 3.15183 11.4691 3.2054C3.22568 3.2054 0 9.35203 0 13.5703V27.5509C0 28.318 0.302085 29.0537 0.839804 29.5962C1.37752 30.1387 2.10683 30.4434 2.86728 30.4434C3.62773 30.4434 4.35703 30.1387 4.89475 29.5962C5.43247 29.0537 5.73456 28.318 5.73456 27.5509V13.6908C5.85403 12.6061 6.45137 8.99046 11.4691 8.99046C12.2095 8.93855 12.9525 9.04541 13.649 9.30401C14.3456 9.56261 14.9799 9.96708 15.5104 10.4908C16.0408 11.0145 16.4553 11.6456 16.7266 12.3426C16.9978 13.0396 17.1197 13.7866 17.0842 14.5345V26.2251C17.0842 26.9923 17.3863 27.728 17.924 28.2705C18.4617 28.8129 19.191 29.1177 19.9515 29.1177C20.7119 29.1177 21.4412 28.8129 21.9789 28.2705C22.5167 27.728 22.8187 26.9923 22.8187 26.2251V14.5345C22.7832 13.7866 22.9051 13.0396 23.1763 12.3426C23.4476 11.6456 23.8621 11.0145 24.3925 10.4908C24.923 9.96708 25.5573 9.56261 26.2539 9.30401C26.9504 9.04541 27.6934 8.93855 28.4338 8.99046C33.4515 8.99046 34.0489 12.4856 34.1684 13.6908V27.5509C34.1684 28.318 34.4705 29.0537 35.0082 29.5962C35.5459 30.1387 36.2752 30.4434 37.0357 30.4434C37.7961 30.4434 38.5254 30.1387 39.0631 29.5962C39.6008 29.0537 39.9029 28.318 39.9029 27.5509V13.5703C40.1419 13.3293 40.0224 10.9188 38.5888 8.38785Z"
                    fill="white"
                  />
                  <path
                    d="M12.0601 31.0488C12.8804 32.6118 14.1251 33.9066 15.6485 34.7817C17.1719 35.6568 18.9109 36.076 20.6619 35.9902C22.4088 36.0519 24.1383 35.6225 25.6573 34.7499C27.1763 33.8773 28.4249 32.5959 29.2637 31.0488C29.2637 31.0488 20.6619 38.4007 12.0601 31.0488Z"
                    fill="white"
                  />
                </svg>

                <div class="mt-2">
                  <p class="text-sm">دریافت نسخه اندروید از</p>
                  <svg
                    width="94"
                    height="30"
                    viewBox="0 0 94 30"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    class="mx-auto"
                  >
                    <path
                      d="M14.3391 9.25714L11.7186 6.98047L9.09814 9.25714L11.7186 11.5338L14.3391 9.25714Z"
                      fill="white"
                    />
                    <path
                      d="M20.9758 9.26048L18.3553 6.83203L15.7349 9.26048L18.3553 11.5372L20.9758 9.26048Z"
                      fill="white"
                    />
                    <path
                      d="M65.6976 14.4189C65.6976 13.9636 65.5229 13.5083 65.5229 13.2047C65.3907 12.8627 65.3315 12.5027 65.3482 12.1423C65.3482 11.8387 65.1735 11.5351 65.1735 11.2316C65.1826 11.0662 65.1203 10.9038 64.9988 10.7762L61.3302 11.6869C61.3302 11.8387 61.5049 11.9905 61.5049 12.294C61.5049 12.5976 61.6796 12.7494 61.6796 13.0529C61.6796 13.3565 61.8543 13.66 61.8543 13.9636C61.8543 14.2671 62.029 14.5707 62.029 15.026C62.029 15.3296 62.2037 15.6332 62.2037 15.9367V16.6956C62.2037 17.1509 62.2037 17.6063 61.8543 17.758C61.7169 18.0292 61.4676 18.2458 61.1555 18.3652C60.9808 18.5169 60.6314 18.5169 60.4567 18.6687H52.9447C52.3466 18.6562 51.7555 18.5534 51.1977 18.3652C50.6387 18.1713 50.1103 17.9163 49.6255 17.6063C49.0316 17.2635 48.55 16.7929 48.2279 16.2403C47.6213 15.6903 47.0934 15.0788 46.6556 14.4189C45.6074 13.2047 44.734 12.1423 43.6858 11.0798C42.6376 10.0174 41.7641 8.80313 40.7159 7.74069L52.5953 3.03556L51.0231 0L37.3967 5.76757L36.6979 8.49958C37.3967 9.41025 38.2702 10.1691 38.9689 10.928C39.6677 11.6869 40.3665 12.5976 41.0653 13.2047C41.7641 13.9636 42.1135 14.5707 42.6376 15.1778C42.9477 15.6407 43.1273 16.1608 43.1617 16.6956C43.1533 17.0228 43.0315 17.3401 42.8123 17.6063C42.6938 17.7637 42.5398 17.8986 42.3597 18.003C42.1795 18.1073 41.9769 18.1789 41.7641 18.2134C41.2744 18.4196 40.736 18.5235 40.1918 18.5169C39.6677 18.5169 38.9689 18.6687 38.2702 18.6687H32.6798C32.585 18.679 32.4887 18.6706 32.398 18.6444C32.3073 18.6181 32.2246 18.5746 32.1558 18.5169C31.9811 18.5169 31.8064 18.3652 31.457 18.3652C31.2823 18.2134 31.1076 18.2134 30.9329 18.0616C30.8263 17.9894 30.7397 17.8974 30.6792 17.7923C30.6188 17.6873 30.5861 17.5719 30.5835 17.4545C30.4249 17.2271 30.3631 16.9585 30.4088 16.6956V12.7494H26.5654V17.1509C26.5553 17.5252 26.435 17.8909 26.2161 18.2134C25.9496 18.3984 25.6558 18.5515 25.3426 18.6687C24.894 18.8034 24.4179 18.8551 23.945 18.8205H5.07768C4.81491 18.7827 4.57138 18.6769 4.37889 18.5169C4.16311 18.344 3.98549 18.1382 3.8548 17.9098C3.69882 17.5731 3.63888 17.2086 3.6801 16.8474C3.65243 16.536 3.71238 16.2235 3.8548 15.9367C3.8548 15.6332 4.0295 15.1778 4.0295 14.8743C4.2042 14.5707 4.2042 14.1154 4.37889 13.8118C4.55359 13.5083 4.55359 13.2047 4.72828 12.9011L2.80662 12.4458L1.05964 11.8387C1.05964 11.9905 0.88495 12.1423 0.88495 12.4458C0.710253 12.7494 0.710251 13.2047 0.535554 13.66C0.360857 14.1154 0.360856 14.5707 0.186158 15.1778C0.0417092 15.6734 -0.0171636 16.1849 0.0114637 16.6956C-0.027171 17.1564 0.0320433 17.6194 0.186158 18.0616C0.360856 18.5169 0.360854 18.8205 0.710249 19.124L1.23434 20.0347C1.40904 20.3383 1.75843 20.4901 1.93313 20.6418C2.50488 21.0822 3.15352 21.4408 3.8548 21.7043C4.59342 21.8908 5.35688 21.9928 6.12587 22.0078H24.2944C25.9749 22.0465 27.6262 21.6214 29.0112 20.7936C29.3175 21.0472 29.6722 21.2526 30.0594 21.4007C30.4527 21.58 30.8616 21.7322 31.2823 21.8561C31.6731 21.9806 32.0898 22.0323 32.5051 22.0078H38.7942C39.6133 22.0277 40.4325 21.9769 41.24 21.8561C42.1135 21.7043 42.8123 21.5525 43.5111 21.4007C44.1712 21.1294 44.8128 20.8253 45.4327 20.4901C45.9371 20.0484 46.3507 19.5351 46.6556 18.9723C47.0767 19.4141 47.5446 19.8207 48.0532 20.1865C48.5773 20.4901 48.9267 20.9454 49.6255 21.0972C50.2234 21.3897 50.8726 21.5948 51.5471 21.7043C52.3521 21.8419 53.1736 21.8929 53.9929 21.8561H60.8061C61.5165 21.8291 62.2204 21.7272 62.9025 21.5525C63.6369 21.3466 64.2984 20.9808 64.8241 20.4901C65.3859 19.9992 65.8573 19.4361 66.2217 18.8205C66.5794 18.049 66.7575 17.2237 66.7458 16.392V15.3296C65.8723 15.1778 65.6976 14.7225 65.6976 14.4189Z"
                      fill="white"
                    />
                    <path
                      d="M60.2812 27.1687L62.9017 29.4454L65.5222 27.1687L62.9017 24.7402L60.2812 27.1687Z"
                      fill="white"
                    />
                    <path
                      d="M53.644 27.1673L56.2645 29.444L58.885 27.1673L56.2645 24.8906L53.644 27.1673Z"
                      fill="white"
                    />
                    <path
                      d="M92.0787 12.2947C91.5405 11.6782 90.8143 11.205 89.9824 10.9287C89.1061 10.5676 88.156 10.3612 87.1872 10.3216C86.2057 10.2684 85.2293 10.4804 84.3921 10.9287C83.6772 11.2989 83.0296 11.7592 82.4704 12.2947C81.9984 12.8674 81.589 13.4772 81.2475 14.116C80.9594 14.758 80.7258 15.4175 80.5487 16.0892C80.374 16.3927 80.374 16.848 80.1993 17.1516C80.0246 17.4552 80.0246 17.7587 79.8499 17.9105L79.3259 18.3658C79.257 18.4235 79.1743 18.467 79.0836 18.4933C78.9929 18.5195 78.8966 18.5278 78.8018 18.5176H77.5789C76.8753 18.5565 76.169 18.5053 75.4825 18.3658C75.1005 18.2789 74.7435 18.1238 74.4343 17.9105C74.228 17.5688 74.0528 17.2135 73.9102 16.848V2.27734H70.0669V17.4552C70.0826 18.0241 70.2005 18.5874 70.4163 19.1247C70.591 19.5801 70.9404 20.0354 71.1151 20.4907C71.5484 20.8302 72.0162 21.135 72.5127 21.4014C73.1611 21.7168 73.8724 21.9228 74.609 22.0085C75.3559 22.1404 76.1188 22.1914 76.8801 22.1603H78.8018C79.4566 22.1467 80.105 22.0442 80.7234 21.8567C81.3216 21.663 81.8598 21.3512 82.2957 20.9461C82.8599 21.2294 83.4433 21.4828 84.0427 21.7049C84.5318 21.9698 85.061 22.1741 85.615 22.3121C86.0631 22.467 86.5331 22.5691 87.0125 22.6156C87.5235 22.7408 88.0547 22.7921 88.5848 22.7674H89.4583C89.8806 22.7177 90.2924 22.6154 90.6812 22.4638C91.1047 22.2866 91.5133 22.0838 91.904 21.8567C92.2958 21.5358 92.6472 21.1797 92.9522 20.7943C93.3171 20.2864 93.5552 19.7176 93.651 19.1247C93.8754 18.3297 93.9926 17.5148 94.0004 16.6963C93.9812 15.8785 93.8641 15.0646 93.651 14.2678C93.0686 13.6467 92.5429 12.987 92.0787 12.2947ZM90.1571 18.3658C90.0222 18.638 89.8459 18.8932 89.633 19.1247C89.5498 19.2173 89.444 19.2926 89.3231 19.3451C89.2022 19.3976 89.0693 19.426 88.9342 19.4283H88.4101C88.0524 19.441 87.6955 19.3893 87.3619 19.2765C86.9686 19.0972 86.5597 18.945 86.139 18.8212C85.7896 18.6694 85.2656 18.5176 84.9162 18.3658C84.5668 18.2141 84.0427 17.9105 83.6933 17.7587L84.2174 16.3927L84.7415 15.0267C84.9291 14.6163 85.2289 14.2516 85.615 13.9643C85.9948 13.6601 86.4962 13.4968 87.0125 13.5089C87.5567 13.5023 88.0951 13.6063 88.5848 13.8125L89.633 14.7232C89.8077 15.0267 89.9824 15.482 90.1571 15.7856C90.2995 16.0723 90.3594 16.3849 90.3318 16.6963C90.401 17.2574 90.3416 17.8248 90.1571 18.3658Z"
                      fill="white"
                    />
                  </svg>
                </div>
              </div>
              <svg
                width="32"
                height="32"
                viewBox="0 0 32 32"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
                  fill="white"
                />
              </svg>
            </a>
            <div class="h-[1px] bg-[#A1B4D680] w-full"></div>
            <a
              class="linearBorder justify-between w-full h-[80px] before:bg-[linear-gradient(83.19deg,#A1B4D6_5.3%,#4469AD_49.08%,#A1B4D6_94.64%)] bg-[linear-gradient(262.38deg,rgba(68,105,173,0.35)52.04%,rgba(48,74,121,0)93.87%)] before:rounded-2xl rounded-2xl flex items-center px-4"
              href="https://sibapp.com/applications/MyGov"
            >
              <div class="flex items-center gap-5 w-full">
                <svg
                  width="32"
                  height="36"
                  viewBox="0 0 40 40"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M25.366 17.3762C27.0674 18.9186 28.2166 20.9761 28.6375 23.2337C29.0584 25.4913 28.7279 27.8247 27.6965 29.8766C33.3111 20.5543 31.1924 13.2448 31.1924 13.2448C19.1158 7.20648 12.3359 12.3973 12.3359 12.3973C17.1395 12.4212 21.7705 14.1907 25.366 17.3762Z"
                    fill="white"
                  />
                  <path
                    d="M11.7011 19.3863C13.2289 17.6632 15.2876 16.4985 17.5515 16.0764C19.8154 15.6543 22.1554 15.9989 24.2015 17.0558C14.9851 11.4412 7.6756 13.5599 7.6756 13.5599C1.74324 25.7424 6.82813 32.5223 6.82813 32.5223C6.71272 27.6816 8.45662 22.9806 11.7011 19.3863Z"
                    fill="white"
                  />
                  <path
                    d="M13.82 32.9456C12.1185 31.4033 10.9694 29.3458 10.5485 27.0882C10.1276 24.8306 10.4581 22.4972 11.4894 20.4453C5.98081 29.7676 7.99359 37.0771 7.99359 37.0771C20.0702 43.1154 26.85 37.9246 26.85 37.9246C22.0284 37.9998 17.3628 36.217 13.82 32.9456Z"
                    fill="white"
                  />
                  <path
                    d="M32.5706 17.7949C32.5664 22.6339 30.796 27.3049 27.5917 30.9309C26.0639 32.654 24.0052 33.8187 21.7413 34.2408C19.4774 34.6629 17.1374 34.3183 15.0913 33.2614C24.3077 38.876 31.6172 36.7573 31.6172 36.7573C37.6555 24.5748 32.5706 17.7949 32.5706 17.7949Z"
                    fill="white"
                  />
                  <path
                    d="M28.2271 0C19.5404 0.21187 19.0107 7.73325 19.0107 7.73325C27.6974 8.4748 28.2271 0 28.2271 0Z"
                    fill="white"
                  />
                </svg>

                <div class="mt-2">
                  <p class="text-sm">دریافت نسخه iOS از</p>
                  <svg
                    width="118"
                    height="28"
                    viewBox="0 0 118 28"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    class="mx-auto"
                  >
                    <path
                      d="M30.7006 20.1428L28.3165 15.805C28.1053 15.4203 27.7805 15.0967 27.3786 14.8703C26.9767 14.644 26.5134 14.5239 26.0407 14.5234H1.02545C0.851144 14.5241 0.679873 14.5649 0.527678 14.6422C0.375483 14.7195 0.247347 14.8307 0.15529 14.9654C0.063233 15.1 0.0102651 15.2537 0.00134815 15.4121C-0.00756883 15.5704 0.0278602 15.7283 0.104311 15.8708L1.71179 18.7133C2.13703 19.4591 2.78015 20.0838 3.57042 20.5186C4.36069 20.9535 5.26749 21.1816 6.19106 21.1779H29.9781C30.1129 21.1788 30.2457 21.1485 30.3641 21.0898C30.4825 21.0311 30.5825 20.946 30.6547 20.8425C30.727 20.739 30.7691 20.6204 30.7771 20.498C30.7851 20.3756 30.7588 20.2534 30.7006 20.1428Z"
                      fill="white"
                    />
                    <path
                      d="M17.4983 23.4473H13.7596C13.5854 23.4472 13.4181 23.509 13.2933 23.6195C13.1685 23.73 13.096 23.8804 13.0913 24.0388V27.3249C13.096 27.4833 13.1685 27.6337 13.2933 27.7442C13.4181 27.8547 13.5854 27.9165 13.7596 27.9164H17.4983C17.6725 27.9165 17.8398 27.8547 17.9646 27.7442C18.0894 27.6337 18.1619 27.4833 18.1666 27.3249V24.0388C18.1619 23.8804 18.0894 23.73 17.9646 23.6195C17.8398 23.509 17.6725 23.4472 17.4983 23.4473Z"
                      fill="white"
                    />
                    <path
                      d="M116.999 14.5258H89.7625C89.6882 14.5277 89.6145 14.5128 89.548 14.4824C89.4815 14.4521 89.4243 14.4074 89.3814 14.3521C89.3385 14.2968 89.3113 14.2327 89.3021 14.1656C89.293 14.0984 89.3023 14.0303 89.3291 13.9672C89.7522 12.8747 89.966 11.7243 89.9612 10.566V7.16483C89.9612 6.96001 89.8718 6.76359 89.7126 6.61877C89.5534 6.47394 89.3375 6.39258 89.1123 6.39258H83.3326C83.1075 6.39258 82.8915 6.47394 82.7323 6.61877C82.5732 6.76359 82.4837 6.96001 82.4837 7.16483V10.6153C82.4742 11.6554 82.0133 12.65 81.2014 13.3825C80.3895 14.1149 79.2924 14.5259 78.1489 14.5258H75.6384C75.5662 14.5266 75.4949 14.5117 75.4303 14.4823C75.3658 14.4529 75.31 14.4098 75.2675 14.3567C75.2251 14.3036 75.1972 14.242 75.1863 14.1771C75.1754 14.1122 75.1818 14.0458 75.2049 13.9836C75.6078 12.9266 75.8153 11.8162 75.819 10.6974V7.24698C75.819 7.14417 75.7965 7.04239 75.7527 6.94761C75.7089 6.85283 75.6448 6.76696 75.564 6.69503C75.4833 6.6231 75.3875 6.56656 75.2824 6.52873C75.1774 6.4909 75.065 6.47255 74.952 6.47473H69.1182C68.8913 6.47469 68.6735 6.55552 68.5114 6.69989C68.3493 6.84426 68.2559 7.04067 68.2512 7.24698V10.6482C68.2489 11.1666 68.1337 11.6796 67.9123 12.1573C67.6909 12.6351 67.3677 13.0683 66.9614 13.4319C66.555 13.7954 66.0735 14.0822 65.5447 14.2756C65.0159 14.469 64.4502 14.5652 63.8803 14.5587H47.1192C46.6464 14.5591 46.1832 14.6792 45.7813 14.9056C45.3794 15.1319 45.0546 15.4556 44.8434 15.8403L42.4593 20.178C42.4011 20.2886 42.3747 20.4109 42.3828 20.5333C42.3908 20.6557 42.4329 20.7742 42.5052 20.8777C42.5774 20.9812 42.6774 21.0663 42.7958 21.125C42.9141 21.1837 43.0469 21.214 43.1817 21.2131H63.9164C65.6701 21.2188 67.4037 20.8745 68.9953 20.2045C70.5868 19.5345 71.9975 18.5552 73.1278 17.3355L74.7895 20.3587C74.924 20.602 75.1306 20.8063 75.3859 20.9486C75.6412 21.0908 75.9351 21.1654 76.2344 21.1639H78.0406C79.894 21.1715 81.7239 20.7865 83.3854 20.0393C85.047 19.2921 86.4944 18.2033 87.6132 16.859L89.311 20.3423C89.4365 20.6004 89.6414 20.82 89.901 20.9745C90.1606 21.129 90.4638 21.2119 90.774 21.2131H111.834C112.761 21.2128 113.67 20.9813 114.463 20.5439C115.256 20.1065 115.902 19.4799 116.331 18.7321L117.884 15.8731C117.965 15.7336 118.004 15.5777 118 15.4202C117.995 15.2628 117.946 15.1091 117.857 14.9739C117.768 14.8387 117.643 14.7265 117.492 14.648C117.342 14.5696 117.173 14.5275 116.999 14.5258Z"
                      fill="white"
                    />
                    <path
                      d="M86.0581 23.4473H75.9798C75.8057 23.4472 75.6384 23.509 75.5135 23.6195C75.3887 23.73 75.3162 23.8804 75.3115 24.0388V27.3249C75.3162 27.4833 75.3887 27.6337 75.5135 27.7442C75.6384 27.8547 75.8057 27.9165 75.9798 27.9164H86.0581C86.2323 27.9165 86.3996 27.8547 86.5244 27.7442C86.6492 27.6337 86.7217 27.4833 86.7264 27.3249V24.0388C86.7217 23.8804 86.6492 23.73 86.5244 23.6195C86.3996 23.509 86.2323 23.4472 86.0581 23.4473Z"
                      fill="white"
                    />
                    <path
                      d="M30.5918 23.4473H20.5135C20.3393 23.4472 20.172 23.509 20.0472 23.6195C19.9224 23.73 19.8499 23.8804 19.8452 24.0388V27.3249C19.8499 27.4833 19.9224 27.6337 20.0472 27.7442C20.172 27.8547 20.3393 27.9165 20.5135 27.9164H30.5918C30.766 27.9165 30.9333 27.8547 31.0581 27.7442C31.1829 27.6337 31.2554 27.4833 31.2601 27.3249V24.0388C31.2554 23.8804 31.1829 23.73 31.0581 23.6195C30.9333 23.509 30.766 23.4472 30.5918 23.4473Z"
                      fill="white"
                    />
                    <path
                      d="M61.1155 23.4475H57.3948C57.3079 23.4453 57.2214 23.459 57.1405 23.4877C57.0595 23.5165 56.9856 23.5597 56.9233 23.6148C56.861 23.67 56.8115 23.7359 56.7777 23.8088C56.7439 23.8816 56.7265 23.9599 56.7266 24.039V27.3252C56.7265 27.4042 56.7439 27.4825 56.7777 27.5554C56.8115 27.6282 56.861 27.6942 56.9233 27.7493C56.9856 27.8045 57.0595 27.8477 57.1405 27.8765C57.2214 27.9052 57.3079 27.9189 57.3948 27.9167H61.1155C61.2024 27.9189 61.2889 27.9052 61.3699 27.8765C61.4509 27.8477 61.5247 27.8045 61.587 27.7493C61.6493 27.6942 61.6988 27.6282 61.7326 27.5554C61.7664 27.4825 61.7838 27.4042 61.7838 27.3252V24.039C61.7838 23.9599 61.7664 23.8816 61.7326 23.8088C61.6988 23.7359 61.6493 23.67 61.587 23.6148C61.5247 23.5597 61.4509 23.5165 61.3699 23.4877C61.2889 23.459 61.2024 23.4453 61.1155 23.4475Z"
                      fill="white"
                    />
                    <path
                      d="M39.695 4.40979e-08H33.4998C33.3305 -5.94084e-05 33.1678 0.0599983 33.0463 0.167398C32.9249 0.274799 32.8543 0.421065 32.8496 0.575078V20.5385C32.8543 20.7099 32.9325 20.8729 33.0675 20.9926C33.2025 21.1123 33.3836 21.1793 33.5721 21.1793H39.6227C39.8095 21.1793 39.9887 21.1118 40.1208 20.9916C40.2529 20.8714 40.3271 20.7084 40.3271 20.5385V0.575078C40.3271 0.422558 40.2605 0.276284 40.142 0.168436C40.0234 0.0605879 39.8626 4.40979e-08 39.695 4.40979e-08Z"
                      fill="white"
                    />
                  </svg>
                </div>
              </div>
              <svg
                width="32"
                height="32"
                viewBox="0 0 32 32"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M18.7072 22.0411C19.0977 21.6506 19.0977 21.0174 18.7072 20.6269L14.081 16.0007L18.7072 11.3744C19.0977 10.9839 19.0977 10.3507 18.7072 9.96021C18.3167 9.56969 17.6835 9.56969 17.293 9.96021L11.9596 15.2935C11.5691 15.6841 11.5691 16.3172 11.9596 16.7078L17.293 22.0411C17.6835 22.4316 18.3167 22.4316 18.7072 22.0411Z"
                  fill="white"
                />
              </svg>
            </a>
          </section>
        </div>
        <div class="hidden tab-content" data-tab-name="privacy">
          <section class="px-6 text-sm">
            <p class="flex items-center justify-center text-lg font-bold mb-4">
              حریــــم خــــصوصی
            </p>
            <p class="font-normal text-justify mb-8">
              اطلاعات مربوط به هر شخص، حریم خصوصی وی محسوب می‌شود. حفاظت و حراست
              از اطلاعات شخصی شهروندان در پنجره ملی خدمات دولت هوشمند، نه تنها
              موجب حفظ امنیت کاربران می‌شود، بلکه باعث اعتماد بیشتر و مشارکت
              آنها در فعالیت‌های جاری می‌گردد. هدف از این بیانیه، آگاه ساختن شما
              درباره ی نوع و نحوه ی استفاده از اطلاعاتی است که در هنگام استفاده
              از  پنجره ملی خدمات دولت هوشمند، از جانب شما دریافت می‌گردد.
              سازمان فناوری اطلاعات ایران خود را ملزم به رعایت حریم خصوصی همه
               شهروندان و کاربران سامانه دانسته و آن دسته از اطلاعات کاربران را
              که فقط به منظور ارائه خدمات کفایت می‌کند، دریافت کرده و از انتشار
              آن یا در اختیار قرار دادن آن به دیگران خودداری مینماید. اطلاعات
              جمع آوری شده از بازدیدکنندگان در بخش نظرسنجی کارکنان دولت و سازمان
              ها، صرفا ً برای بهبود کیفیت خدمات و محتوای تارنما مورد استفاده
              قرار می‌گیرند.
            </p>
            <p class="font-bold mb-4">
              چگونگی جمع آوری و استفاده از اطلاعات کاربران
            </p>
            <p class="font-medium">
              لف: اطلاعاتی که شما خود در اختيار این سازمان قرار می‌دهيد، شامل
              موارد زيرهستند:
            </p>
            <p class="font-normal">
              در درگاه پنجره ملی خدمات دولت هوشمند، اقلام اطلاعاتی شامل شماره
              تلفن همراه، تاریخ تولد، کد پستی و کد ملی کاربران را دریافت
              مینماییم که از این اقلام، صرفا جهت احراز هویت کاربران استفاده
              خواهد شد.
            </p>
            <p class="font-medium">
              ب: برخی اطلاعات ديگر که به صورت خودکار از شما دريافت میشود شامل
              موارد زير می‌باشد:
            </p>
            <ul class="list-disc pr-4.5 mt-2">
              <li>
                دستگاهی که از طریق آن درگاه سازمان فناوری اطلاعات ایران را
                مشاهده می‌نمایید.
              </li>
              <li>نام و نسخه سیستم عامل وbrowser کامپیوتر شما.</li>
              <li>اطلاعات صفحات بازدید شده.</li>
              <li>تعداد بازدیدهای روزانه در درگاه.</li>
              <li>
                هدف ما از دریافت این اطلاعات استفاده از آنها در تحلیل عملکرد
                کاربران درگاه می باشد تا بتوانیم در خدمت رسانی بهتر عمل کنیم.
              </li>
            </ul>
            <p class="font-bold mb-4 mt-6">امنیت اطلاعات</p>
            <p class="font-normal text-justify">
              متعهدیم که امنیت اطلاعات شما را تضمین نماییم و برای جلوگیری از هر
              نوع دسترسی غیرمجاز و افشای اطلاعات شما از همه شیوه‌‌های لازم
              استفاده می‌کنیم تا امنیت اطلاعاتی را که به صورت آنلاین گردآوری
              می‌کنیم، حفظ شود. لازم به ذکر است در سامانه ما، ممکن است به سایت
              های دیگری لینک شوید، وقتی که شما از طریق این لینک‌ها از سامانه ما
              خارج می‌شوید، توجه داشته باشید که ما بر دیگر سایت ها کنترل نداریم
              و سازمان تعهدی بر حفظ حریم شخصی آنان در سایت مقصد نخواهد داشت و
              مراجعه کنندگان  میبایست به بیانیه حریم شخصی آن سایت ها مراجعه
              نمایند.
            </p>
            <p class="font-bold mt-8 mb-4">
              تغييرات در بيانيه سياست حفظ محرمانگی 
            </p>
            <p class="font-normal text-justify">
              ما این حق را برای خود محفوظ می‌دانیم که در هر زمانی این بیانیه را
              ویرایش نموده و تغییرات خود را اعمال نماییم. در این شرایط شما موظف
              هستید که آخرین تغییرات ما را در بیانیه از طریق سامانه مشاهده نموده
              و از آن مطلع شوید. 
            </p>
          </section>
        </div>
        <div class="hidden tab-content" data-tab-name="guide">
          <section class="px-6">
            <p class="flex items-center justify-center text-lg font-bold mb-4">
              راهنــــما سامــــانه
            </p>
            <p
              class="flex items-center justify-center text-center text-sm opacity-60 font-medium mt-2 mb-8 px-3"
            >
              برای آشنایی سریع با امکانات سامانه، فرآیند ورود و پاسخ به پرسش‌های
              رایج، می‌توانید راهنمای جامع را در قالب PDF دریافت کنید.
            </p>
            <ul class="mt-10 px-6 list-disc text-sm">
              <li class="font-bold text-[#ECF4FF]">
                ایجاد و فعال‌سازی حساب کاربری
                <p class="font-normal mt-2">
                  مراحل ثبت‌نام، تأیید هویت و فعال‌سازی اولیه
                </p>
              </li>
              <div class="w-full bg-white/20 h-[1px] my-4"></div>
              <li class="font-bold text-[#ECF4FF]">
                ورود امن به سامانه
                <p class="font-normal mt-2">
                  توضیح شیوه‌های احراز هویت (رمز عبور، رمز یکبار مصرف) و نکات
                  امنیتی
                </p>
              </li>
              <div class="w-full bg-white/20 h-[1px] my-4"></div>
              <li class="font-bold text-[#ECF4FF]">
                بازیابی گذرواژه و بازیابی دسترسی
                <p class="font-normal mt-2">
                  راهنمای گام‌به‌گام برای بازیابی رمز عبور و رفع مشکل ورود
                </p>
              </li>
              <div class="w-full bg-white/20 h-[1px] my-4"></div>
              <li class="font-bold text-[#ECF4FF]">
                مدیریت اطلاعات شخصی
                <p class="font-normal mt-2">
                  نحوه به‌روزرسانی پروفایل، شماره تماس و تنظیمات امنیتی
                </p>
              </li>
              <div class="w-full bg-white/20 h-[1px] my-4"></div>
              <li class="font-bold text-[#ECF4FF]">
                پرسش‌های متداول (FAQ)
                <p class="font-normal mt-2">
                  پاسخ به رایج‌ترین سؤالات کاربران درباره خدمات و خطاهای احتمالی
                </p>
              </li>
              <div class="w-full bg-white/20 h-[1px] my-4"></div>
              <li class="font-bold text-[#ECF4FF]">
                تماس با پشتیبانی
                <p class="font-normal mt-2">
                  روش‌های ارتباط با مرکز تماس، ایمیل و تیکت آنلاین
                </p>
              </li>
            </ul>
            <div
              class="w-full rounded-2xl before:bg-[linear-gradient(88.93deg,#304A79_0.91%,#7C96C6_98.92%)] bg-white/20 linearBorder before:rounded-2xl mt-10 py-4 px-4 flex items-center justify-between text-sm"
            >
              <div class="flex items-center gap-4">
                <svg repalce="document"></svg>

                <div class="text-[#ECF4FF] flex flex-col gap-1">
                  <p class="font-bold">راهنمای جامع پنجره ملی</p>
                  <p class="font-normal">فرمت فایل PDF</p>
                </div>
              </div>
              <svg
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M12 14.5L12 4.5M12 14.5C11.2998 14.5 9.99153 12.5057 9.5 12M12 14.5C12.7002 14.5 14.0085 12.5057 14.5 12"
                  stroke="white"
                  stroke-width="1.5"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
                <path
                  d="M20 16.5C20 18.982 19.482 19.5 17 19.5H7C4.518 19.5 4 18.982 4 16.5"
                  stroke="white"
                  stroke-width="1.5"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
            </div>
          </section>
        </div>
        <div class="hidden tab-content">content-6</div>
      </section>
    </div>
  </section>
</div>

        <!-- .................................form wrapper...................................... -->
        <section
          id="form-wrapper"
          class="relative overflow-x-hidden w-full h-full flex flex-col items-center overflow-y-auto overscroll-y-contain overscroll-x-contain webkit-overflow-scrolling-touch pb-20 z-0"
        >
          <div
            class="relative md:max-w-[70%] max-w-[85%] mt-32 text-center mr-[20px] sm:mr-0 z-0 h-full"
          >
            <svg
  version="1.0"
  id="Layer_1"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  x="0px"
  y="0px"
  viewBox="0 0 500 500"
  enable-background="new 0 0 500 500"
  xml:space="preserve"
  class="absolute top-[-42%] md:top-[-26%] left-[-293px] md:left-[-70%] opacity-50 z-[-1]"
  width="497"
  height="489"
>
  <image
    overflow="visible"
    width="500"
    height="492"
    xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAHsCAYAAAAkU198AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJ
bWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdp
bj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6
eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDkuMS1jMDAyIDc5LmI3
YzY0Y2NmOSwgMjAyNC8wNy8xNi0xMjozOTowNCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRm
PSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNj
cmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8x
LjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6
c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHht
cDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDI2LjEgKFdpbmRvd3MpIiB4bXBNTTpJbnN0
YW5jZUlEPSJ4bXAuaWlkOkY3MjU1NkYzQkVGQzExRjBBMzQ1RDg0NjBEODEyREY2IiB4bXBNTTpE
b2N1bWVudElEPSJ4bXAuZGlkOkY3MjU1NkY0QkVGQzExRjBBMzQ1RDg0NjBEODEyREY2Ij4gPHht
cE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6RjcyNTU2RjFCRUZDMTFG
MEEzNDVEODQ2MEQ4MTJERjYiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6RjcyNTU2RjJCRUZD
MTFGMEEzNDVEODQ2MEQ4MTJERjYiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94
OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz5qwO29AAB33UlEQVR42uydB1xTVxuHs0OAMMUB
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
zUR+g0SLgAAAAABJRU5ErkJggg=="
  ></image>
</svg>
            <svg
  version="1.0"
  id="Layer_1"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  x="0px"
  y="0px"
  viewBox="0 0 500 500"
  enable-background="new 0 0 500 500"
  xml:space="preserve"
  class="absolute top-[-27%] right-[-379px] md:right-[-73%] opacity-50 z-[-1]"
  width="497"
  height="489"
>
  <image
    overflow="visible"
    width="500"
    height="492"
    xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAHsCAYAAAAkU198AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJ
bWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdp
bj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6
eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDkuMS1jMDAyIDc5LmI3
YzY0Y2NmOSwgMjAyNC8wNy8xNi0xMjozOTowNCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRm
PSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNj
cmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8x
LjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6
c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHht
cDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDI2LjEgKFdpbmRvd3MpIiB4bXBNTTpJbnN0
YW5jZUlEPSJ4bXAuaWlkOkU1RkU1MDE1QkVGQzExRjBCMzI5RTcyNzAxODE3OEE0IiB4bXBNTTpE
b2N1bWVudElEPSJ4bXAuZGlkOkU1RkU1MDE2QkVGQzExRjBCMzI5RTcyNzAxODE3OEE0Ij4gPHht
cE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6RTVGRTUwMTNCRUZDMTFG
MEIzMjlFNzI3MDE4MTc4QTQiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6RTVGRTUwMTRCRUZD
MTFGMEIzMjlFNzI3MDE4MTc4QTQiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94
OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz66T1kAAABrcElEQVR42uydB1wUx9+Hb6/f0auC
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"
  ></image>
</svg>
            <svg
  version="1.0"
  id="Layer_1"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  x="0px"
  y="0px"
  viewBox="0 0 500 500"
  enable-background="new 0 0 500 500"
  xml:space="preserve"
  class="absolute mb-[174px] md:mb-[28%] bottom-[-255px] md:bottom-[-28%] left-[-337px] md:left-[-66%] opacity-50 z-[-1]"
  width="497"
  height="428"
>
  <image
    overflow="visible"
    width="500"
    height="430"
    xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAGuCAYAAABmw/QbAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJ
bWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdp
bj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6
eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDkuMS1jMDAyIDc5LmI3
YzY0Y2NmOSwgMjAyNC8wNy8xNi0xMjozOTowNCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRm
PSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNj
cmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8x
LjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6
c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHht
cDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDI2LjEgKFdpbmRvd3MpIiB4bXBNTTpJbnN0
YW5jZUlEPSJ4bXAuaWlkOkI2QThCODI1QkVGQzExRjA5RkZEOEZBN0U1Njc5MEM1IiB4bXBNTTpE
b2N1bWVudElEPSJ4bXAuZGlkOkI2QThCODI2QkVGQzExRjA5RkZEOEZBN0U1Njc5MEM1Ij4gPHht
cE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6QjZBOEI4MjNCRUZDMTFG
MDlGRkQ4RkE3RTU2NzkwQzUiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6QjZBOEI4MjRCRUZD
MTFGMDlGRkQ4RkE3RTU2NzkwQzUiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94
OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz7+S5CeAAB9F0lEQVR42uydBXgURx+Hz3PJXYy4
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-tax-govhLj7OzsRipDevft27/+
5s3IK0zmf4ciMDIy5sydO6/XhIkTjdSp/LDU8T31iRPtoPe7DFpaWlQajQbBNwAouNAPIKG7QE40
HaNGjbJ88/ZdgqWllbkip/Obb751evT4cQSKyPWr+1xTU5O6bNlyC2NjYw6Foj7dLObMmeuKpP4e
SR06yiGcnJzGbfP0nICv5gAAoNhCx4+tuUFONC0eCxdOiI17dm/kyJFmipi+xYu/nvbufUaM9dCh
pjWt1759e43lK1b0t3dw6KxO5Td7zhxHJPVUJPUh6nwe7/Dy+ibu2fM4CwsLU/hVA4DiC/0UEvp0
yImmZ/6CBWNep6Q+dHScPE6xIi5nh6h798OsrKw61WX9du3bs1HFbrli5cr+aiZ1F8nldzv1lPru
PXu84l8lpvbo0cMYfs0AAEJXe6ysrfu/jI+PXLtunTvu5S5PjIyMNI+fOPnzXzf+Pm9sbGxQz++y
V6xcNWD6jBlm6lR+s2fPcVVHqfv67vZKf/M2C8m8C/yKAQCEDhB0NDHRQ/kcEnHt+klzc3MTeaRh
zBhbqydPY6MO+fnv1eJqNWgyHtwg+fnnX0YtWbp0gDpK3c5uklpI3cfH1ys3Lz9vwoSJ5vDrBQDl
EvpJEHrzQ6fTSZ7bt69KTEqO27p12zetW7fWaIn9duvWrc3x4yd+SUp+HTt37txGX/pvY9CGeez4
iTFr1q4drmZSd1EHqe/es8ebkLkF/GoBAIQO1IC5hYXp2XPnjsXGPbu+ecuWxYZt22o1x37MzMw6
oMp5w+uU1Kg//vxze7t27XSaattoW+S7UfdWfLd+/Wh1KrtZs2djqWfZTVJNqe/bv3/ni/j49wMG
DuwKv1QAUE6h+yOhT4WcaDnwODC2trZ9zwUF+z96/OTG3r37vhs+3KY3l8ulN2a7uEe6k5PzmJ9+
/mVX7LPnN1Bk7tWnT59meY4cd5S7dz969tp169Tq0a7ZlVLPVDWpI5nvehn/6n23bt06wy8UAJRX
6Pg5dGfICfnQu3fvjl5e3sv+uXkzJObR44gjR4/+vHDRIoe+fftadOzYUYfNZv/fA+D4mXAqolOn
TvqT7O2/2LR5s4ef/2GfJ09jb0REXAtas3btPOOOHVs1dSOExWL9Z2CRfv36tU5NS9+zbPnyoWoW
qTsTUleJ+dT37f/RB8k8Fcm8SR5Ng9FnAEB+Qj8WGBAQ6u/vdxpyQ37gMdiRILvZjBhh06tXrz76
+vpGenp6wo8fP5bw+eVcTU1NPAFLnkAgIBcVFTHQ+qVcrpYYfcZMSXn9/t69e0/+vnHjfmpq6pvK
EWKbllatWtGNjY2ZT548KZZ933roUJNVq1bbBgede3Tq1KkH6lJeM2fOct7h5bU3JSXlubhybFw8
eD+D8Bm5Gq/hdaST8CjEhGR4yGE8gQ+VQmG4urqMTUhIeNcU24XZ1gBAfuDfH8y2JmcqKipE0dHR
8XiRRsSGbdtq9+/fvyOfx2dzOBwqjpBLS0v5iYkJJejz4qzs7NzcnJzSlkifoWFbpria4OvO7dup
5eXl508cP/Ft+w5GvX19dvmpQ3kdP37s7OvXr+NQI6cdamRhmYsIobM/E7Ti8fdxWeHh1jiKchxU
KpUSHX0/Gh3LB/gVAoDyR+hH4B66YoKnQh08eLC2IqTFwsJC097eoc3nPre2tm6X/ubt2W/XrJkD
JafeQHQAAPKDT4LbXgqJUCgU6+npMYyNO7LknRZ8VbmMV1bxuc/v3Lnzfqqry5LlK1bYrVmzdjGU
HggdAAA51NWkysuAgKIVDJIolUYj0+k0BagjyRyhQFhjw+Lu3btZU11cVy9bvnzQ2nXrvoQSBKED
ANDyQudBNiik0MXFRUUVZApF7ldQyspKS007mdb6zPzdu3fSUKS+bv78BfO+XbN2DpSi+gGd4gBA
vg1qaFQrIFQqVSLy3NxcvrzTQqfTNUxNTfE0q69ql/rdbDu7iWvCwyP+JJNJVF8fHz8oTRA6AACA
OkfoeE5yhjaXy8zPyxPIMy3GHTvSu3TpqlfX9RNevcqY5jbVPSDwbBB6qYGk/jOUqPpECAAAAIAM
uFOcgUEbDbRw5Z2WLl266AsEgnrdmomKispwcXF2XbZs+Zy1a9cthRJVfZgsFh2EDgAAUA0MBkNM
o8m/U1w5vxwPaFPvCWXuRUW9cXVxtl+6bNl8JPXlUKKqC5VKJe3auesnEDoAyAk8gAmguLTv0J4r
FovlXkdqaWlx2rVt26Bx21Gk/tZFKvV1676FUlXNeuTnX349PGLkyAFwDx0A5ASvrEwAUldcTExM
2tDpNIG80yESi6hCkSivod9HkXo6lnpg4Nmb6CXLZ9eu7VC6qiPz3w8cODpkiKWJo4P9CBA6AMiJ
4uJiELoCU1RUrFFaVsZXgKTgsQoaNcwsknqqi7PzqMCzZ/9BL4VI6t5Qwqohcysra4tJdhNHpKam
loLQAUBO6BsYaDbHZB5A08DR0GCxmMwKBUgKfga9W2M3cu9eVJKLs9OwgMBzN8kkMm3Xrp2eUMpK
LfNjlpZW5kjmXyKZF+H3QegAICeQzPFkEgzICcUkKTk5taS0VBEG/sHnSJMMQXvv3r0UVxenkZLL
72SSYNfOnRCpK6XMDx6zsrQ0mTTJbjiSeYn0M+gUBwDyg0pEX4AC0qNHj94cDqe1AiQF19NNdikH
ST3Z2dnZ5uuvly5a991366GklU7mxy0tLTvb2U0cIytziNABQL7gqR9LIBsUEwqF3IZCJmsqQFLw
s/BN2rC4f/9esouL05eBZ89F4svvKFCHy+9KIPMDKDIfYmnZE98zT0tL+7+rRyB0AJAf5ZAFiotY
LGlwCRQkOU3e2eI+itQnOzpYBwQE3hSJhGIfHx/o/a4EMrefZDcSybyouvVA6AAgP7AsFO4eOo1G
I+P5wPHwpxKTiESkiooK9EckrrouHvMcVzbSdWtCuk28LWK7Yun3KCgcJnr8y26ITNRRYvyhoaEh
nc1mk4VCIendu3e4MSQ2MDBgVKn4yGibZLRtssw2JJes0T7KBUIJkqTj9/D6aM8UYr//LuV8PhWl
t0BbR5fE1eZqsVhsCrEtKiFXsUz9iSMl3BueRSRBiLYpyS9x5Xr4aQYaSheN+K5IpuzZxPdFMnUy
nXhPXFxczONwNOl0Ok2sraNDYjKZOFqvkFmfSfrvFLziehQ1+fmzZ+mrVq4ce/CPP/5CyRUePHjA
m8Vma5FV8/ELGunT7IYCYqES5VAhUyYUmQa3tJwFxPt0mTyuSx7JNgrrfesE/0DwCevl5X0Qydxs
kp3dl2lpqYWfLVCoUwFAPoSFRxwKDAi45+/vd0iR0tW6tT7NxMREA4sPS0lXV5eGZMp8/vxZiay4
sejRuhRj444c7OPPSR27Aa8bHx9fjIRE7ty5M1qfQkIVU8mHDx8kj+517tyFo6enR0P7ZBKSwhUf
B1emaN2ywsJCSnl5ebGmpmY5shsFLXgbFOOOxq2l9RiWUGlpaaGBvkFbcwsLA4FAUIDexmOgV6C0
lSQnJ70xN7dgtWnTBj9doIEaI7r5+fmaKF25aFMfCYkWVh6aiNanT5+FaH/ZqAGC04OfAy9GiwlR
wecSrzVkGh5ctE0R2mbWy5cvUy0sLPCgMIZoWww+n6+Djx2tUyQjZNx/woDYhrRRwSCOv4xoJOBj
7Y3Sp4G2cZ8QUB5xXCJiPQNCNHQiz6oWhPizdT3KMz6PV8DlcvV1dHTMcnJy3pSUlLxj4sSqltTF
RL5KG4B8Qtg0olElfc2WNiKJvBUT+con8l6DkDS5Dv4kE2XNkzlHhPVsdIlZLBY3F5WL2zQ3h7TU
1ILaWiwAAMivklG46VNzcrIFaCmUjdg1NLD/aOSqok5PfyN+9CimuC6VP37uHos8NvZpcWVUTiXj
15i3b9/xxGKRNHKWRkj4jQq8aSRnMZK1WLpfDocjWQfJllYlohGizx8Ts6UJie1JIiS0DYaOji4D
NU4YosrnBZmSaF4sEshE55LNoMZD+Z49e0uDgoIuRd29G4UqVTaxHbLMNgVEHcogFkk0hqMqtE18
bGIyiYzXZfJ4PD7ahpj0aYY9KlHZC6tcpZGmQ3J8SK7F8xcsWDRw4MBhy5ctm89gMNhEGtgy0mHL
yKW6chCRau4ATebxyngdTUxaHTz4x6WHDx682LJly1oNDTYXlRFFxX5vstE1WSZ/pK9l16FU0yCq
762Puoi/RvDVq4yM9++zs7NrHYsAhA4A8kNEUpx7tJ8FyxRFyE2SThz1FxUV1bStWveDrwSgxoG0
Yq2tH4Ls9sqzsjLxUqe0ogaEflpqSmlSUiKuSEvllf8ovR/LEAkJr3B0VtBc+0lKSnrr5DRl+PmQ
83ccHB3i9+/b5ws/UeUCHlsDAPmBozQWZINCQ1andHA0NATe3l72S5YsXbp+/YYfoPiVC4jQAUB+
4Et5QsgGhUVEUpN+RoZt27J9fHy9Hz9+dG3f3r1XHj16NOJcUPAtfPTeXl5b4VSACB0AgJrRJBZA
McEdzNgK1PhrFtq1a6cZHBxyJjExIR/LHL/34MGDFGcnp2GLFi1esH4DROoQoQMAACg3bAUKepjN
JPNWKBK/dOnSxXAUiW+T/Sw6+j6S+pThZ88F4bHfyV5eO7bAKQEROgAA1YMfYSqEbFBYyApSR0of
nWpSBg4caHTp8pXISxcvXq0q809Sj8ZSt1ng7j5v3779O1gs6PIBQgcA4HPCEEM2KCyKNFIctSk3
Zmlp2SEoOOTvwMCAM97eXjtqWhdL3dXF2cbF1XX2Di/vHXQ6Hc4MEDoAAFXAj1zB/KlAXWiy82TI
EMsOwSHnb1+8ePHMrp076zTcK76nPsnObqizs/OsX3/73QcidRA6AAD/RTpyFNBAGAwG2cTEhNNM
o5pRSYrTz6hJ6uohlpZGQcHBkWFhYadXrVzxfX2+GxPzMNXRwX7YuHHjpvn5+++BSB2EDgDAJ/DQ
nzAfeiNo164dw2bECK64LoPJKzeNPj7JZfag4NsR4eHBHu4L1gsE9b+bEBMTk+pgb29tbT3U5fcD
ByFSB6EDAECARx8TQDY0nPz8fKFIKBJLh5BVUciNPU8sLa2MzlXK/IKHh/u3FRUVDd4WitTTpkye
PNTW1tZt/48/7aRSqXAigtABQO2RTvIANLQCQyLX1tZm4zHZVfxQGzzmf6XMg26Fh4dfRDJf3hiZ
y0rdwX7SsLFjx7r9eegQROogdABQe3Bow4FsaEQGUqnibt27t2smoctO1CFvGtQpztLSsktQcPAD
JPMLCz0WNInMP0k9JnXKZEebr74a7eZ/+MheJpMJJyQIHQDUWujwG2wE+gYGmlwtLX2Saj8tUN6Q
hgWSuUlQcEhkWGjoGQ/3BSsqKpr+7g5xT32otbW105EjR0HqIHQAUFsUKQJUSnS0tRl6rVtrN1Nd
1uipL5tQ6PXCysqq67ngkPthYaGBHh7uKxvSAa6uPHoUk4Z7vw8bPtzZe+dObzgrQegAoI7g65/F
kA0NB4lKQ1NTszsx17iqIqxPw08i86Dg2ygyP+Xh3rwy/yT1R2n2k+yG2ts7TP9+0yZPODPlA4zl
DgDyA98/hy7CjYBKpXKzs7O0eTwePLYmkbk1knnQ7VAk84Ue7qtaQuZSYmJi0lxdnIdJZmlDbPf0
3ARnKAgdANQFPDIHzLbWOPBN244qfox16h9AyPwOkvnJlpa5lIcPH6Y5TZmMpR5JJpFJnp7bQOog
dABQC3BFzYdsaEQGisUkNltD1Z+ZwoMP1Xh71Mrautu5c5LI/ASS+Wp5yLyK1IcHBJ79m4MIDw87
wWKx8CBK+NZBGXE8NGLBj+NVVHMFgkb8PkRVGsB00qe+JzTi91NOvM8mtkcltint/yCUucrBIf5K
x+mXDnfHJxrXOH0UCpXKehUfH5uUlJQLQgcAoC7wSNCPpXFCF4koXC6XwWazSWVlZap8qBU1ROZY
5rdCQ68eXOjhsUmeMpeV+lRXl1H79v94aOSXXw5B5SQiRCogfbrNJNt/hEV8ziOEX0FIlkxIHf9O
dInPpZ0EsZwLSZ+eFhET38VLKenTI6EiGbmTiW2wifepMtKXDMWMkirW1dXVp1AoZa4uzuOePn2a
DUIHAKA2cIVVBNnQCKELRWIkdDoe2ESFhV5O+kxve2tr6+5nzwVFhl5FMl/osVkRZC4FT+gywmb4
qOYZZr/5EAqFJNxAPHz4yC/nL1y8hXvwP3nyRCmkDkIHAPlBIcFIcY1EjEeLoxKTszRpxzgOR4Ni
YmqqExkZKe+DpJKq6TwplflVLHMP981YRIqGIjUw6kNJSQlp1qyZS48dP/5zyPkLkUjqNkjqWcpQ
oQAAIB+4aNGGbGgMZFJzTcwSEhLiv8PLa9/o0aP7KNpRWw8dimQeTMjcQyFlruzweDzSrJkzl0VH
R19HUr/Zt29fAxA6AACfDy8bMGgIUI3Vm+FKx5HDh/12eG7fceTY8euKJHWJzM+eQzK/fGDRQixz
mN+nOaU+Z/YsGan3Uxip6+jq0lgsFgWEDgCKI3SojRWYAwd+/8Vz69atcpb6v1cgJJfZzwZFXr1y
Bcl84Q/KeklbmcB9M6RSP4+l3k/+Up8wYeLgFctXjKyoqBBBCQGAAhAWHnFm3rz58yEnGk7//gMG
PXkae79Vq1bN2hdh7rx5Himpae9GffWVRUsf4/IVK1zPnAk4MmDgwPbvMz5kHjrkt5VGg+5PLQ2b
xSYFBJ79KTUtPV6eUl+xcuXsZ89fRA0aNMi06mdwVgCA/BC2bt1aR19f3wA/rwvZUT/4fB6/bVvD
dhQKhWVkbNxJQ0NDgETX5Fcd8T36C+fPB3br2q3rkSNHb3h4uI9/FBPzlo27QjczxQiOhkaHHmZm
vfBl9itXLv++eNGiLRCZyyFS55WRZs+aufzosWM/oUg90sHefviTJ49btKPctm2ea2fOmrV0sqPD
6MePH6dU/Rx62AKAnAgIDPQZNGjw1JKSkrw6/hbxY24wnZWMa5HMOXQ6vV15eXlSM/WN+1fqfD6/
TEdHpwOXy+2Qk5OTVlpa+oHFYmmQmneCHRGVSm2tra3d4cKF8+uXL1u2E2QuX/AjkkeOHvtpyJAh
tg74kbbHLSP1bZ6eG6dPn7ESyXzo06dPX1W3DggdAOSEoaEhHhCFCb/DhkbofF6//v2tvL137pk1
c+ao3NycZonQZaVeUVFR4uTkvHj9hg1bfti8eezZs4ExSLbazSV1FKAXff31kuVfDBkyfML4cVOa
s9EC1FvqP2KpOzrYD3/czFL39Ny+cdr06auQzEcgmT+DEgAAQOUg7qFHN/c99Kq4u3ssTklNSxkx
YmTX5t7X8hUrPELOXwiE0lY8qZ8JCPwR31Pv169/m2aT+fbtG5Nfp+T07tPHAnIdAAAVFnr/gbGx
cdF6enot/sSOx8KFS99nfHgzZoxtjxYQ+lkobUWVeoBE6uhcbHKpg8wBAFAbBgwYOODe/ehILS0t
uex/4aJFS1+jUN3GxqbZInUQunJIPS09/VW/JpS65/YdUpmbQy4DAKDyDEL8/c/NG7hSlRcLFy5c
giL19OaK1EHoyiH10xKpv2kSqW/fvuP7pOTXOb17g8wBAFCbCH1Av4cPY25oa2vLtWOhh4dE6s1y
+R2ErkRSPxOwH0u9MZffJTJ/nZLVu3dvkDkAAOoDqjgHPI2NuyePe+hVcffwwFJPs7W17Q5CV0+Y
TKZE6qnpbxIaIvXtO3ZsQpH5m169e3eG3AQAQM2EPmDgk6exd1u6l3sNUndHUk9BUjcDoaux1E+f
wZF6Qv8BA+osdSTzzRKZ9+rVCXIRAAB1FDp+bC0KCV1h0kRE6m9tx45tksvvIHRljdTPVF5+HzDA
sLb1d+zwwjJ/iyJzkDkAACB0RcLd3eNrQuqNjtRB6MoqdRbp1Okz+7DUB9Qg9X9l3gtkDgAACP2e
ogn9P1K3bZzUQejKC+4oR0g9obpIXSYyN4XcAgAAhP409gESukJOBe3u7r64sZE6CF3ZI3VmtVLf
4eX1A9wzBwAA+K/QoxWlU1x1LHB3X4QfabO1bdgjbSB0VZH66X2paemv+vTpo71p8+b1ScnJ6SBz
AACAf4Xef9DT2LjoVgrw2FpNLFy4aDGqzNMtLa26gdDVEzqdTvrjz0NbUOOu+GHMo6fde/Qwaup9
UCCbAQBQVshkMkkgEIhFQqFCT0N28OCB34ODgn4KCg6+OXr06B5QcupHRUUF6f3790V4yl06nVEh
FAqLQOgAAACflF4pdJFI4ecVXb161e6zgYF7T58J+HvMmDFmUHbqhZf3zi2urq7Lvxo1qkPMwwcR
oaFhUQMHDmwLOQMAAED695L7Az0Fv+Quy779P67Jyc3LGDOmboPPwCV31ZB5YvLrNxYWvSS92YmO
cnvT37xNAKkDAACQ/tMpTqmuNu7bv/9bJPUPdZE6CF0FZJ6UjGRuYSL7Ppb6yVOn9iCpJ4LUAQAA
oSup0KVSz87JrTVSB6ErvczTkcyrfc5cIvWTIHUAAAClFrpE6vskkXrmmBrGfgehKyfe3ju3EjI3
qWk9LPUTlVJPAqkDAABCV1KhY/bu2/9NTVIHoSuhzHdKZJ5mbl6zzKUwGIxKqb8FqQMAAEJX6id2
CKl/qG6WNhC6ssrc3KQ+3/sk9XeJgwYNAqkDAKCWQr+vyCPF1VvqY8f2BKErJzt37ZJcZq+vzP8r
9ZMH3rx9lwxSBwAAInSllvq+/5M6CF1pZL6NkHnHxmxHck/9BEgdAAAQuvJLfe9/pY6E7o6EHgil
rfoy/xSpg9QBAFA/oQ9W5NnWGsqevXtXY6lbWVm1X7hw0UIk9HNQ2oor84TEpLSmkvknqTNIx0+c
wFJPQlI3hJwGAEDFhd7fJjbuWayenh5Z1Y4NSz0lNS0hODjkyKnTp09BaSukzD2xzHs28J55HaV+
lIjUQeoAAKguAwYMHHE/+sETLperksfn67t7aRmPLz527PgRKG3FYtcun0qZ9+xp0pz7IXq/H0f7
etWpUyc9yHkAAFQSa+uhwy9dvnJLlY9x3/4ff0hLf5M/dty47lDiCiJzH1+pzDu2xP6w1E8iqcc8
evzMwKBNq8+tR4aiAQBAWTEyMm4zfsL4rlFRUYm8sjJSWlpaLp/PFxgbG3NRJcgWI9BqIgqFUsbj
8YSZmZn8Dh06MNFrpp5ea00KhUzCa+BpWBECOoNOp1ApzKKCojL8P5vNFhd8/ChE22ToGxhooM0x
iV2L0ELh8/jZaD0ttD02el2BFqZMvcpACxUtOA1C4n8S8Zpcpf7F/2ugpZx4jf9Si4qKRFMmTxmx
avXqNbmIb1avWpyUlFSmxeXidfloobdgPS4m9kmTST+ZeE9IfC49NjFxvFSZY2IQC1lme2KZ1zwi
DzWI70nzjC+zL/wehVgYMtsgEd+V7rMMLSyZtFTN6wqZfJZ+TiMWAZGvIpkyq0wgYvr0GY62trYT
HR3sR7x48SKtpc51LPWjR4+dMDE17Ws/aZJNVlZmLggdAACVgcPhkA4fOepfXl5esmH9d2s7dDBq
g+rcQn19fRqdTmNhoSNZk4uKiwuYDAatoKCQoqfXSoNKpYq7dO3Kehb3rLB9h/as1q1ba6JV+T17
mhuibRokJyfno/f0dHV1tSJv3kwqKPgoHj9hYhe0jg5aJFJBfw1QAyIBNR700fb00GssAF2cLBkZ
6ROiEhLv82VEwSAEIhUKg5AJRlhWVspHHzBYDCYNNShKWuvrt7KxsbFOTUlJfhgTk6ihocGsKpwW
QCpNESFDEXFMFBn5yqZJKnR8nHj+by1ikTaKRDL58ZHIAw3iPalky4n3STLvU6o0kEhVZF8h0/AQ
VyN2kczxSF1IJb4rlGmI/aezJWq4UdGKxTt2bF8cFxeX0tLnu0Tqx44fMTExGeBgP2k4aqDmQy0A
AIDSo6WlRbp6NfRsYmLS027du2ORkuh0OhUtqN6lkGQXHIGjv2QajUYhqHMwQyb43MfNdXxfDBli
Nm7cuIGy7/n4+i598/ZdypejRnWDM0A9wVI/fSbgQOSt23+z2Wwq5AgAAEqNpqYWOTQsPOSvG39H
mZqaqlyPOLtJkywLCosKVqxcOa3qZ76796zEY7+PHTfOHM4E9ZV6QGCg744dXssgNwAAUObInBwW
Hn4uLDziNvqfrWrH5+Dg8NWHzKySrdu2LfrcOljquXn5WeNA6moLi8WiOTu79GzXrj0DcgMAACWV
ecS5sDCJzDVUTuaOjiMzPmQWTZ8+Y2Zt6/ru3l0p9fHjLeDMUE+kt5QAAACUS+ZcbqXMw8NVUuaO
jo7WSOb5dZG5lN0oUs/Oyfsw3MbGDM4QAAAAQAlkrkUOD48IIi6zq6jMs+ol838jdd/dK7JzcjPH
Q6QOAAAAKDJ4FLjwiIig0LBwlbxn7ug4eRiKzAumNUDm/0p99+4VuXn5SOoTesEZAwAAACheZK6l
peoyt8KX2adNnz6zsdvCkTq+pw5SBwAAABRN5uTwiGuqLHPrppK5FJ9PUofL7wAAAID84XK5hMzD
bmlqqp7MJ08mZD6t6WT+f1KfAFIHAAAAFEDmV0OxzDVVUeZDkcxz3aZNc2uuffj4+q7Lzf+YMwGk
DgAAAMhL5hHXJDK/ppoynzI0IzMzH8l8RnPva5eP74bcPCz1iSB1AAAAoEVlTkEyD0EyD0Myp6mc
zKdMwb3Zs5HMp7bUPpHUf0CReh5IHQAAAGi5yDziWjCS+VVVlPkUqczdpjm09L59fXd7vXufkdWz
Z88ucKYBAAAAzSvza9fPXg0NDUcyV7lxqQmZZ8pD5lJ2797j9TL+1ftu3buD1AEAAIBmk3kIknmo
il5mH1opc7dJ8k7L3n37NryIj39oaGjIgTMPAAAAaEqZ05HMQ69cDb2hmpG503Ak86ypbm4OipAe
PKX71m2es12nTv0Czj4AAACgSdDW1sYyj0Ayv8nhcFgqKPPRSOZ5U6e6TVa0tKH8psIZCAAAADSJ
zK9d++salrkmR1PlZO7k5DQWj82OZO4CpQ0AAACoqszJ167/dVEic01VlLnz2IzMrI9I5nZQ2gAA
AIAqy/z8latXVVXmI1FkXoxk7gylDQAAAKiqzEkSmV+RyJypgjLHHeBQZD7VDUobAAAAUFmZX7/+
V/DlK1f/4XA4qihz/GhavqsryBwAAABQUbhcLpb5ISTzCFXsze7k7IwHjckDmQMAAAAqi46ODvXC
xUuely5fOaCKkbmzs8sAJPNUV2RzKG0AAABAVaGePn1mTlh4+BZVlLmLi2s/JPMrSOYToagBAAAA
lYSM8PXdPe7AwYMuqtibfdCgQSYvXr76aepUty+htAEAAACVhE6nk9esWdtl9epvVHK6zr79+rUN
j7g2f9y48TAdKQAAAKC6TJ4yRW/x4q874ihd1Y6tV69eWr/99vsXvXr11oWSBgAAAFQaHR0dOoPB
ULmxwtu3b0/btHmzSd++fXWglAFF5X8CDACar4sG30/5FQAAAABJRU5ErkJggg=="
  ></image>
</svg>
            <svg
  version="1.0"
  id="Layer_1"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  x="0px"
  y="0px"
  viewBox="0 0 500 500"
  enable-background="new 0 0 500 500"
  xml:space="preserve"
  class="absolute mb-[176px] bottom-[-258px] md:mb-[29%] md:bottom-[-29%] right-[-290px] md:right-[-80%] opacity-50 z-[-1]"
  width="497"
  height="428"
>
  <image
    overflow="visible"
    width="500"
    height="430"
    xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAGuCAYAAABmw/QbAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJ
bWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdp
bj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6
eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDkuMS1jMDAyIDc5LmI3
YzY0Y2NmOSwgMjAyNC8wNy8xNi0xMjozOTowNCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRm
PSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNj
cmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8x
LjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6
c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHht
cDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDI2LjEgKFdpbmRvd3MpIiB4bXBNTTpJbnN0
YW5jZUlEPSJ4bXAuaWlkOjAyRDREMEI4QkVGRDExRjA4NDBFQjgzQjVBRkVFRkRGIiB4bXBNTTpE
b2N1bWVudElEPSJ4bXAuZGlkOjAyRDREMEI5QkVGRDExRjA4NDBFQjgzQjVBRkVFRkRGIj4gPHht
cE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6MDJENEQwQjZCRUZEMTFG
MDg0MEVCODNCNUFGRUVGREYiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6MDJENEQwQjdCRUZE
MTFGMDg0MEVCODNCNUFGRUVGREYiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94
OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz6dv2x7AABmB0lEQVR42uydBXgURx+Hzy/u7g5B
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
9bunPodK/Wn8CyL1uhWQ+VXcMwcAqCVeXt6m9++HjwqYPLmSZd7CKSEx6emqv/+eqYmR+Xulnp6R
4KuAjnJlUl9CpX5bLaU+53WkXrc8me8JLZX5FTMzM8gcAKB+ePv4mBw7fqLv0KFDm1bmdpuTyDzz
2fOk2bPnTEMuvyX19RvmpqVlJPr6NldQpC6VuqenOkp9ammk/tVXdd+U+V7IHACgvjRv3tw8Kvrx
0GHDhzer3Mj8a+e0dNrMvm42m8VGRr9F6dSrVOo0Um+uIKkvLpO6p6f63VOfPWfO1Li4eDr1qieD
yaQy333pEmQOAFBTmvn6WiYmJS+bVMm92anM0zMyEteuXafR98w/Rer0cT5IvXKZMGHiqKvXrkeE
7t0XduLkyVPoAAe+BMy2BhgmpqZaVlZWPDorGofDYSYlJhYWFRWVzqxFm16dXVx0yF/W23OdU+jU
nHSGLfq9uLi4gpLiYrGLq6sOHdmKTs7FYbO1SMShzWKxxIWFhYzsnByhhbk5T25atiJG2TSjdMYT
PbLwyH6EL7OyXnh4eOgtWLDwf6tWrri6ZMni/ZV1vA6Ojrbnzv176ezZs8HDhw2dSec5B++HloHV
QWvmfN+p06BePfxaXrp0Kb6607Bo8ZIVnTt3/nrjxg2LSflgyZUfISlb3IyMjEJDQ0Mun89nl1NO
mbRiEh8fn0/KIjc3N5eWRYmjo6PWnTt3cqytrXl2dnbatPx/4Jooed9nNCnp6elFZNt8+luo6DGR
/UmEQmH+sGHD5zk7O9fq1rVLi5MnT15EiQMQOvhsdHR1OXTuanJxETs5OWlTGVOh0wsVuXJKTE1N
6RSm7/QWo9OaPoqIyCMXUo69nZ1W5rNndO5ziZmpKZcKngqdS3BxcdUj60ru3b+Xm5SUVNSuXTsz
Po/PkZRdealN6YVUmwqdvMUm3xXY2tqWkLRYkAtcXlDQ6quVday1atd2CAnZevru3Tvbhw0dOgsy
/xSpB83p1KnzoJ49/FoQqSdU5/7p/eXff18w18XFxUcgEBTJSbZ0rnqRWFxm+f+m75XNTS+g9U5a
zIUiEZsUZIaYFDsJrWySMlpMKqQcWiEly0cmrhdIK53vlfOnyFzuN8SJjYt9IBQIc7r7+Y3q17dP
+7t374ajxAEAwAf4+usWzvTRtPXrN8xCM/vnSX3d+g1zHz58FG5gYKCNHKlcZs+eMyXuafyLr75q
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
QgAAAABJRU5ErkJggg=="
  ></image>
</svg>
            <div
              class="flex items-center md:justify-center gap-4 text-[#304A79]"
            >
              <svg
  width="66"
  height="56"
  viewBox="0 0 66 56"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g clip-path="url(#main-logo)">
    <path
      d="M45.5392 0.0078125C46.8063 0.0499319 47.8105 1.22011 47.7134 2.46389C46.3705 2.60075 45.3964 1.24814 45.5392 0.0078125Z"
      fill="#304A79"
    />
    <path
      d="M42.7899 0.799641C44.041 0.448735 45.5199 2.11371 44.6993 3.19606C43.4339 3.28694 42.6875 1.91507 42.7899 0.799641Z"
      fill="#304A79"
    />
    <path
      d="M39.7282 1.36931C40.9865 0.953568 42.3365 2.51836 41.7189 3.66396C40.4099 4.3271 39.3627 2.4991 39.7282 1.36931Z"
      fill="#304A79"
    />
    <path
      d="M36.4738 1.71207C37.718 1.54731 39.2215 3.29108 38.0286 4.2472C36.6485 4.6191 35.8544 2.78751 36.4738 1.71207Z"
      fill="#304A79"
    />
    <path
      d="M25.9703 2.92561C27.8515 2.5747 29.7751 2.75022 31.6775 2.60101C32.454 2.60101 33.2569 2.44316 34.0211 2.63954C34.5506 3.01848 34.4129 3.71671 34.3441 4.26758C31.7127 4.4237 29.0709 4.66061 26.4361 4.51498C25.6986 4.32724 26.0039 3.47462 25.9703 2.92561Z"
      fill="#304A79"
    />
    <path
      d="M51.0473 3.58984C52.3655 3.58984 53.1526 4.9091 53.2638 6.08459C51.8551 6.5564 50.8338 4.80958 51.0473 3.58984Z"
      fill="#304A79"
    />
    <path
      d="M48.4584 4.56838C49.1534 4.25082 49.7643 4.83677 50.1331 5.36134C50.5672 5.85083 50.4314 6.54387 50.4539 7.14032C49.1291 7.27717 47.9855 5.79648 48.4584 4.56838Z"
      fill="#304A79"
    />
    <path
      d="M45.3012 5.28359C46.6411 4.98876 47.9483 6.59049 47.2426 7.77993C45.8926 8.44122 44.7912 6.42733 45.3012 5.28359Z"
      fill="#304A79"
    />
    <path
      d="M42.0205 5.82207C43.2312 5.76586 44.2706 7.06944 43.8858 8.21371C43.18 9.21368 41.9199 8.26805 41.6658 7.39444C41.3504 6.85007 41.7791 6.30106 42.0205 5.82207Z"
      fill="#304A79"
    />
    <path
      d="M38.1981 6.49919C39.6665 5.9729 40.9811 8.58868 39.3929 9.19736C37.9792 9.48156 36.9875 7.20087 38.1981 6.49919Z"
      fill="#304A79"
    />
    <path
      d="M56.1138 6.91406C57.4503 7.0735 58.055 8.50357 57.9915 9.69992C56.5475 9.72982 55.9655 8.11759 56.1138 6.91406Z"
      fill="#304A79"
    />
    <path
      d="M16.9504 7.70177C17.764 7.08606 18.835 7.50353 19.7741 7.50353C24.6659 7.73512 29.5826 7.76316 34.4693 7.37717C35.4169 7.23102 36.0731 8.50868 35.3869 9.18246C33.2693 9.64205 31.0721 9.51583 28.9155 9.58944C24.8925 9.58944 20.8523 9.66491 16.848 9.20877C16.8411 8.70865 16.7546 8.1789 16.9504 7.70177Z"
      fill="#304A79"
    />
    <path
      d="M53.6469 8.05518C54.9546 7.82186 55.8122 9.34401 55.6117 10.4727C55.2587 11.1885 54.4963 10.5885 54.1629 10.171C53.631 9.59553 53.6117 8.79207 53.6469 8.05518Z"
      fill="#304A79"
    />
    <path
      d="M50.6008 9.10752C52.0126 8.43387 53.3609 10.5636 52.542 11.6812C51.9685 11.9868 51.3649 11.5514 51.0297 11.0987C50.502 10.5701 50.6395 9.7789 50.6008 9.10752Z"
      fill="#304A79"
    />
    <path
      d="M47.555 9.9174C48.6633 9.44904 49.775 11.0601 49.4396 12.0805C49.2507 12.863 48.0789 12.7612 47.6995 12.1962C47.1755 11.5981 46.7433 10.4666 47.555 9.9174Z"
      fill="#304A79"
    />
    <path
      d="M43.727 10.6133C45.204 10.1397 46.2365 12.2431 45.3041 13.2957C44.48 13.7185 43.4211 13.1659 43.2146 12.2887C43.081 11.701 43.0351 10.8659 43.727 10.6133Z"
      fill="#304A79"
    />
    <path
      d="M61.0498 11.0605C62.3222 11.2676 62.7157 13.0026 62.4016 14.0886C61.2633 13.6359 60.6863 12.1939 61.0498 11.0605Z"
      fill="#304A79"
    />
    <path
      d="M38.4363 11.7205C39.1546 11.7205 40.1323 11.3696 40.6317 12.0714C40.8946 12.5837 40.87 13.5083 40.1552 13.6118C38.3374 13.851 36.5057 13.9328 34.6844 14.1504C30.6255 14.261 26.5824 14.5925 22.5305 14.3258C20.5892 14.1732 18.648 14.2727 16.7068 13.9995C15.3885 13.8241 14.0102 13.9995 12.7361 13.5644C12.106 13.0539 12.4343 12.0293 13.1155 11.7276C13.4758 11.6907 13.8389 11.6907 14.1991 11.7276C17.1428 12.0609 20.1058 12.21 23.0652 12.2977C28.2007 12.3784 33.3449 12.3679 38.4363 11.7205Z"
      fill="#304A79"
    />
    <path
      d="M58.7381 13.2457C58.6305 12.8685 58.874 12.0704 59.3787 12.3037C60.367 12.8528 60.6899 14.13 60.4569 15.1721C59.3593 15.2703 58.8441 14.1632 58.7381 13.2457Z"
      fill="#304A79"
    />
    <path
      d="M56.1277 13.4816C57.633 12.9553 58.7943 15.6552 57.4442 16.4047C56.1718 16.1093 55.8771 14.5692 56.1277 13.4816Z"
      fill="#304A79"
    />
    <path
      d="M53.1069 14.3755C54.537 14.2294 55.2846 16.0334 54.7394 17.1614C54.1513 17.839 53.2959 17.1807 53.0046 16.5596C52.5986 15.872 52.7909 15.0562 53.1069 14.3755Z"
      fill="#304A79"
    />
    <path
      d="M49.1871 15.7671C49.8524 14.8302 51.1601 15.7074 51.2855 16.6056C51.4373 17.2144 51.2238 17.8179 50.9979 18.3741C50.6044 18.3741 50.2108 18.36 49.8173 18.3548C49.2047 17.6882 48.5693 16.618 49.1871 15.7671Z"
      fill="#304A79"
    />
    <path
      d="M62.834 15.7148C64.154 15.8675 64.4134 17.6446 64.0693 18.6849C63.4517 18.8814 63.1517 18.1902 62.9488 17.7415C62.7647 17.0815 62.834 16.3903 62.834 15.7148Z"
      fill="#304A79"
    />
    <path
      d="M45.1143 16.3645C45.9472 15.7364 47.0308 16.5399 47.1985 17.4171C47.5391 18.2451 46.8597 19.4381 45.8785 19.2012C44.6942 18.9328 44.1648 17.1996 45.1143 16.3645Z"
      fill="#304A79"
    />
    <path
      d="M40.6738 16.9207C41.7875 16.5189 42.968 17.5803 42.6874 18.7224C42.7456 19.861 41.0727 20.1997 40.4161 19.4241C39.6591 18.7785 39.7156 17.3593 40.6738 16.9207Z"
      fill="#304A79"
    />
    <path
      d="M61.0268 19.8872C60.4639 18.9942 60.088 17.6644 61.0515 16.8926C62.0663 17.6276 62.2693 18.896 62.031 20.0503C61.701 19.9976 61.3127 20.0942 61.0268 19.8872Z"
      fill="#304A79"
    />
    <path
      d="M13.7564 17.4187L13.947 17.2593C15.8882 17.0839 17.8295 17.5084 19.7707 17.5084C22.8181 17.5469 25.8644 17.8856 28.9139 17.6312C31.3157 17.4995 33.73 17.5715 36.123 17.2803C36.7935 17.3119 37.6301 17.0048 38.156 17.5574C38.8919 18.1083 38.3501 19.475 37.4395 19.403C35.1453 19.4609 32.874 19.7883 30.5798 19.7539C26.3443 19.7539 22.0701 19.9293 17.847 19.4767C16.5074 19.382 15.1274 19.5364 13.8251 19.1557C13.1829 18.7661 13.2075 17.8748 13.7564 17.4187Z"
      fill="#304A79"
    />
    <path
      d="M58.1152 18.2881C59.1123 18.0424 59.527 19.351 59.5411 20.1217C59.6664 20.7111 59.0117 21.7268 58.394 21.1742C57.4946 20.6192 57.2328 18.8757 58.1152 18.2881Z"
      fill="#304A79"
    />
    <path
      d="M54.7977 19.4039C56.2255 19.1319 56.739 21.3438 55.936 22.2109C55.313 22.8898 54.4006 22.0647 54.2453 21.395C54.1325 20.7357 53.9983 19.7024 54.7977 19.4039Z"
      fill="#304A79"
    />
    <path
      d="M50.8779 20.379C52.1467 19.9211 53.0185 21.8088 52.4132 22.7878C52.1396 23.4895 51.1055 23.6088 50.6307 23.0632C49.8967 22.3439 49.8825 20.8842 50.8779 20.379Z"
      fill="#304A79"
    />
    <path
      d="M64.2462 21.0258C64.5462 20.7462 64.8086 20.7754 65.0333 21.1135C65.4711 21.8521 65.461 22.7696 65.3457 23.5924C65.1409 24.2568 64.4226 23.7468 64.3045 23.3012C64.1511 22.552 64.1314 21.7818 64.2462 21.0258Z"
      fill="#304A79"
    />
    <path
      d="M45.6011 22.8553C45.4105 22.0062 46.0141 21.1203 46.9142 21.1133C47.3995 21.2519 47.7963 21.5887 48.2006 21.8658C48.1565 22.5799 48.483 23.5185 47.7788 24.0061C46.9282 24.7308 45.5394 23.9237 45.6011 22.8553Z"
      fill="#304A79"
    />
    <path
      d="M5.16154 21.8192C5.53217 21.5578 6.01574 21.7122 6.43034 21.7454C12.4747 22.8173 18.6284 23.0262 24.7503 23.2279C30.574 23.1945 36.4154 22.8769 42.2214 22.2946C43.155 22.1543 43.7868 23.3472 43.0635 23.9911C41.0319 25.0191 38.671 24.9805 36.4544 25.2192C33.4772 25.263 30.5001 25.5596 27.5246 25.3261C20.208 25.2156 12.877 24.8244 5.63161 23.7945C4.66741 23.8072 4.28087 22.23 5.16154 21.8192Z"
      fill="#304A79"
    />
    <path
      d="M61.9978 22.0879C63.572 22.1721 63.5616 24.3334 62.9119 25.3686C61.3537 25.1457 61.629 23.1738 61.9978 22.0879Z"
      fill="#304A79"
    />
    <path
      d="M59.2534 23.5276C59.7969 23.0398 60.3846 23.803 60.5258 24.3064C60.7887 25.0923 60.6862 26.4323 59.6704 26.5748C58.528 26.1485 58.4433 24.2767 59.2534 23.5276Z"
      fill="#304A79"
    />
    <path
      d="M55.5123 27.3038C54.7358 26.4441 55.0846 24.3056 56.5236 24.5951C57.6019 25.1407 57.3477 26.5548 57.0707 27.5039C56.5482 27.6074 55.8635 27.8758 55.5123 27.3038Z"
      fill="#304A79"
    />
    <path
      d="M51.7392 25.6714C52.1401 25.3328 52.7133 25.5557 53.0557 25.8679C53.8175 26.8012 53.5251 28.5748 52.1892 28.7872C50.7668 28.4976 50.5674 26.3818 51.7392 25.6714Z"
      fill="#304A79"
    />
    <path
      d="M65.0788 26.4199C65.3613 26.4901 65.933 26.3831 65.9083 26.8374C66.0847 27.8444 66.0443 28.9286 65.4424 29.797C64.5548 28.8713 64.8971 27.5409 65.0788 26.4199Z"
      fill="#304A79"
    />
    <path
      d="M-0.220454 26.7415C0.430724 26.6239 1.10142 26.5415 1.75767 26.6959C7.51905 27.8125 13.3548 28.5091 19.2184 28.7799C25.8733 28.9747 32.5476 28.9677 39.1743 28.3149C41.8674 28.1096 44.5321 27.6342 47.2205 27.3658C47.7163 27.336 48.4787 27.2482 48.5846 27.9096C48.8687 28.743 47.8945 29.2782 47.245 29.5394C44.649 30.4376 41.8854 30.6972 39.1641 30.9183C34.2687 31.2043 29.3609 31.4025 24.4636 31.0832C16.83 30.8919 9.21879 30.1768 1.68429 28.9431C1.10542 28.7799 0.272474 28.8659 -0.0151559 28.2168C-0.303454 27.7695 -0.261085 27.2414 -0.220454 26.7415Z"
      fill="#304A79"
    />
    <path
      d="M63.0372 27.5352C64.1154 28.0614 64.0731 29.479 63.6386 30.4245C63.4904 30.8332 62.9504 31.1262 62.6345 30.7087C62.0666 29.6755 62.4495 28.465 63.0372 27.5352Z"
      fill="#304A79"
    />
    <path
      d="M60.0512 29.0361C61.0076 28.7221 61.2636 29.913 61.1872 30.6027C61.1995 31.2869 60.7002 32.4378 59.8408 32.1097C58.952 31.308 59.1952 29.7571 60.0512 29.0361Z"
      fill="#304A79"
    />
    <path
      d="M55.6428 31.9784C55.6146 31.0135 56.2569 29.9468 57.3476 30.082C57.6599 30.7381 57.9864 31.4771 57.7288 32.2117C57.5684 32.7082 57.2557 33.3748 56.6293 33.3169C55.8968 33.3977 55.6251 32.5626 55.6428 31.9784Z"
      fill="#304A79"
    />
    <path
      d="M51.8593 31.6653C52.1646 31.0724 52.934 31.2881 53.4811 31.2109C53.714 31.7956 54.0825 32.4091 53.8954 33.0601C53.7038 33.5182 53.4636 33.9547 53.1789 34.3622C52.7006 34.2972 52.0723 34.5113 51.7477 34.0566C51.2236 33.3706 51.3841 32.3373 51.8593 31.6653Z"
      fill="#304A79"
    />
    <path
      d="M46.6653 34.3925C46.189 33.1225 47.4186 31.7381 48.7552 32.0505C49.4169 32.7891 49.6552 33.9978 48.9193 34.7715C48.3525 35.5626 46.9142 35.3785 46.6653 34.3925Z"
      fill="#304A79"
    />
    <path
      d="M65.5609 32.0955L65.7373 31.9902C66.1503 33.0236 65.9791 34.3867 65.1885 35.2112C64.6079 35.0849 64.8355 34.3182 64.809 33.899C64.9273 33.2744 65.0138 32.5007 65.5609 32.0955Z"
      fill="#304A79"
    />
    <path
      d="M2.1882 32.4001C2.89602 32.3374 3.60922 32.3764 4.30586 32.5158C10.1296 33.6264 16.0593 33.9842 21.9753 34.1386C25.9408 34.4052 29.9168 34.0911 33.8787 34.0228C35.5517 34.0333 37.2068 33.7 38.8817 33.772C39.4853 33.772 40.0094 34.3947 39.8307 34.9877C39.0384 35.907 37.7323 36.0702 36.5962 36.2351C30.6278 36.8862 24.6134 36.465 18.6379 36.1579C14.3475 35.807 10.0452 35.5439 5.78331 34.9298C4.54792 34.7369 3.21028 34.7719 2.11429 34.0947C2.12645 33.5335 2.15078 32.9651 2.1882 32.4001Z"
      fill="#304A79"
    />
    <path
      d="M41.4489 33.9714C41.7207 32.7433 43.6726 32.3433 44.2513 33.531C44.6695 34.4433 44.1702 35.3766 43.4784 35.9871L42.0401 36.0012C41.6271 35.4182 41.1118 34.7076 41.4489 33.9714Z"
      fill="#304A79"
    />
    <path
      d="M62.9701 33.5965C63.2136 33.3596 63.6053 33.3614 63.7131 33.7333C63.9319 34.6999 63.6795 35.6947 63.1837 36.5403C63.0072 36.5403 62.6418 36.5649 62.4619 36.5754C62.3181 36.0726 62.289 35.5442 62.3769 35.0288C62.4648 34.5134 62.6673 34.024 62.9698 33.5965H62.9701Z"
      fill="#304A79"
    />
    <path
      d="M59.2535 36.6538C59.3223 35.7398 59.8077 34.5099 60.9495 34.7749C61.3184 35.8661 61.1259 37.3644 59.963 37.8958C59.3559 37.8186 59.2006 37.1643 59.2535 36.6538Z"
      fill="#304A79"
    />
    <path
      d="M56.481 35.8887C56.8318 35.9022 57.1805 35.9492 57.5222 36.029C57.9776 36.7815 57.8169 37.771 57.337 38.4675C56.9928 38.9605 56.1474 39.2675 55.7769 38.6429C55.2458 37.7045 55.911 36.65 56.481 35.8887Z"
      fill="#304A79"
    />
    <path
      d="M52.325 37.0563C52.7661 36.9159 53.2374 37.0404 53.6891 37.0563C53.8056 37.6948 54.051 38.3896 53.7014 39.0072C53.3749 39.7475 52.2896 40.4953 51.6155 39.7088C51.1304 38.8247 51.4691 37.5966 52.325 37.0563Z"
      fill="#304A79"
    />
    <path
      d="M47.6186 37.9196C48.148 37.7844 48.8628 37.7442 49.1681 38.302C49.8952 39.5301 48.6916 41.316 47.2445 40.872C46.0143 40.1949 46.588 38.4809 47.6186 37.9196Z"
      fill="#304A79"
    />
    <path
      d="M7.06741 39.2151C7.90571 38.7431 8.87272 38.9537 9.77276 39.0397C14.0076 39.5577 18.2675 39.8476 22.5338 39.9081C26.1445 40.2257 29.7552 39.859 33.3676 39.8204C35.7536 39.6714 38.1448 39.5555 40.5203 39.2801C41.7232 39.1941 42.9395 38.8695 44.1345 39.1888C44.1907 39.5237 44.2072 39.8641 44.184 40.2028C43.4975 41.0345 42.3327 41.1449 41.3462 41.3765C36.7824 42.1502 32.1359 42.0993 27.5192 42.0659C24.0055 42.1238 20.5022 41.7869 16.9906 41.6992C13.9641 41.4535 10.9198 41.3325 7.91092 40.9255C7.08506 40.7957 7.1186 39.8466 7.06741 39.2151Z"
      fill="#304A79"
    />
    <path
      d="M62.5324 38.8866C62.9066 38.4234 63.2578 39.3744 63.09 39.6919C62.9136 40.4796 62.6623 41.4025 61.8759 41.7971C61.0358 41.0075 61.8141 39.5515 62.5324 38.8866Z"
      fill="#304A79"
    />
    <path
      d="M59.7528 40.0743C60.1163 39.876 60.4429 40.2146 60.4252 40.5777C60.5259 41.4547 60.0242 42.1987 59.5605 42.8847C59.2481 42.8952 58.934 42.9093 58.6249 42.9233C58.5103 41.8725 58.7415 40.6217 59.7528 40.0743Z"
      fill="#304A79"
    />
    <path
      d="M55.3601 42.0055C55.6884 41.4651 56.3008 40.9336 56.9731 41.2563C57.5378 42.4984 56.5708 44.3861 55.0548 44.0247C55.0372 43.3405 54.9083 42.5844 55.3601 42.0055Z"
      fill="#304A79"
    />
    <path
      d="M50.9692 45.0689C50.5262 43.8076 51.3568 41.9672 52.8928 42.2619C54.0875 43.3673 52.4823 45.5935 50.9692 45.0689Z"
      fill="#304A79"
    />
    <path
      d="M47.0553 43.3776C47.5159 43.0442 48.2907 42.9524 48.6877 43.4307C49.4889 44.6398 47.9042 46.4413 46.5788 45.8357C45.7159 45.1934 46.263 43.8478 47.0553 43.3776Z"
      fill="#304A79"
    />
    <path
      d="M41.6127 44.3233C42.1633 43.7339 43.0161 43.8934 43.7303 43.9478C44.5897 45.2512 43.2009 47.1057 41.7203 46.5986C40.7233 46.2456 40.8485 44.8562 41.6127 44.3233Z"
      fill="#304A79"
    />
    <path
      d="M36.7738 44.4821C37.3809 44.3436 38.2862 44.3278 38.5475 45.0277C39.2092 46.3856 37.2062 47.8627 36.0644 46.9012C35.1503 46.2225 35.7432 44.705 36.7738 44.4821Z"
      fill="#304A79"
    />
    <path
      d="M14.7762 45.9425C14.3138 45.3671 14.9232 44.518 15.6127 44.6197C17.7304 44.8232 19.8587 44.7495 21.9748 44.9706C23.8719 45.1035 25.785 44.9174 27.6804 45.0775C28.2627 45.3034 28.2592 45.9545 28.3068 46.481C27.878 46.6652 27.4562 46.9284 26.9761 46.9477C23.5755 47.0805 20.1835 46.7407 16.7829 46.688C16.0821 46.6074 15.1556 46.6741 14.7762 45.9425Z"
      fill="#304A79"
    />
    <path
      d="M31.264 44.8123C31.9488 44.5948 32.9017 44.5895 33.2459 45.3474C34.1013 46.614 32.12 48.1157 30.9871 47.2667C29.967 46.7578 30.2965 45.193 31.264 44.8123Z"
      fill="#304A79"
    />
    <path
      d="M57.2223 47.6717C57.2382 46.5103 57.8595 45.1206 59.1317 44.9121C59.3559 46.1069 58.5989 47.705 57.2223 47.6717Z"
      fill="#304A79"
    />
    <path
      d="M53.8944 48.7789C53.6773 47.4948 54.5403 45.9579 55.9892 46.0703C56.2161 47.3527 55.3427 48.8929 53.8944 48.7789Z"
      fill="#304A79"
    />
    <path
      d="M51.1718 47.1995C51.5862 46.8485 52.3012 47.131 52.3189 47.6872C52.3877 48.5328 51.7383 49.2573 51.0358 49.638C50.9394 49.6978 50.8314 49.7368 50.719 49.7527C50.6065 49.7686 50.4919 49.7609 50.3825 49.7302C50.2732 49.6995 50.1715 49.6464 50.0841 49.5743C49.9966 49.5022 49.9252 49.4128 49.8745 49.3117C49.7358 48.4276 50.3818 47.5644 51.1718 47.1995Z"
      fill="#304A79"
    />
    <path
      d="M45.3434 49.1826C45.9417 48.6704 46.5152 47.83 47.4223 48.0089C48.1512 48.0301 48.0259 48.9598 47.8795 49.4387C47.5052 50.2738 46.2241 51.0527 45.4671 50.2352C45.3628 49.8966 45.3734 49.5318 45.3434 49.1826Z"
      fill="#304A79"
    />
    <path
      d="M41.4929 48.8653C41.7499 48.7193 42.0439 48.6502 42.3395 48.6664C42.635 48.6826 42.9196 48.7833 43.1588 48.9566C43.8347 50.3232 41.9923 51.7969 40.7693 51.0285C39.9841 50.3337 40.7183 49.2022 41.4929 48.8653Z"
      fill="#304A79"
    />
    <path
      d="M36.202 49.3692C36.862 49.0043 37.8821 49.1393 38.1785 49.9147C38.4786 51.3024 36.4244 52.3708 35.4732 51.3323C34.8943 50.6416 35.4801 49.6797 36.202 49.3692Z"
      fill="#304A79"
    />
    <path
      d="M19.2941 50.2096C20.7242 49.9008 22.1936 50.1271 23.6425 50.0956C26.3196 50.157 28.9968 50.0447 31.6739 50.0956C32.5293 50.1061 32.4469 51.0237 32.3533 51.6167C29.9475 52.1289 27.4701 51.8558 25.0277 51.9167C23.1235 51.8448 21.165 52.029 19.3133 51.4886C19.3046 51.0586 19.2958 50.6341 19.2941 50.2096Z"
      fill="#304A79"
    />
    <path
      d="M52.8403 51.3803C53.2145 51.0135 53.7597 50.3277 54.3332 50.7505C54.5096 51.9504 53.3891 52.7767 52.4167 53.2065C52.1149 52.5575 52.3302 51.8628 52.8403 51.3803Z"
      fill="#304A79"
    />
    <path
      d="M48.6282 53.987C48.0564 52.7256 49.5846 51.4326 50.8147 51.573C51.1078 52.842 49.8819 54.0414 48.6282 53.987Z"
      fill="#304A79"
    />
    <path
      d="M44.1966 54.5239C44.066 53.1099 45.5131 52.1713 46.8172 52.4378C46.803 52.829 46.7961 53.2238 46.7873 53.6168C46.1166 54.3169 45.2042 55.0309 44.1966 54.5239Z"
      fill="#304A79"
    />
    <path
      d="M40.9669 53.0783C41.5934 52.8046 42.5164 53.2047 42.3946 53.9731C42.137 55.0537 40.3104 55.9766 39.6222 54.8047C39.5035 53.9872 40.2539 53.3257 40.9669 53.0783Z"
      fill="#304A79"
    />
    <path
      d="M29.8648 54.2177C30.9236 54.1125 31.9983 54.1844 33.0625 54.1213C34.2654 54.0932 35.4626 53.8178 36.6662 53.9459C37.6086 54.4073 36.6662 55.4599 35.9373 55.4757C34.4285 55.6109 32.9125 55.6265 31.4019 55.7653C30.643 55.7266 29.6848 56.0232 29.1361 55.3354C29.0194 54.8529 29.3035 54.2177 29.8648 54.2177Z"
      fill="#304A79"
    />
  </g>
</svg>
              <div class="space-y-1">
                <p class="sm:text-[24px] text-[20px] font-bold">
                  درگاه یکپارچه احراز هویت
                </p>
                <p class="sm:text-[18px] text-[16px] font-medium">
                  پنجره ملی خدمات دولت هوشمند
                </p>
              </div>
            </div>
            <p
              class="font-bold sm:font-normal sm:text-2xl text-base text-right sm:text-center text-black mt-[40px] sm:mt-[56px] mb-2.5 sm:mb-4"
            >
              درگاه یکپارچه دسترسی به تمامی خدمات دولت
            </p>

            <!--........................input mobile and captcha form............................-->
            <form
              method="POST"
              id="login-form"
              class="group mt-2.5 md:mt-4 flex-col"
              data-active-form="one-time-pass"
              novalidate
            >
              <p
                class="text-right text-[#424446] font-medium md:font-semibold hidden group-[&[data-active-form=one-time-pass]]:block"
              >
                ورود با کد یکبارمصرف
              </p>
              <p
                class="text-right text-[#424446] font-medium md:font-semibold hidden group-[&[data-active-form=application]]:block"
              >
                ورود با اپلیکیشن
              </p>
              <p
                class="text-[#5D5F61] mb-[80px] text-sm md:text-lg text-justify"
              >
                جهت ورود به سامانه و دسترسی به خدمات، شماره موبایل خود را به
                همراه کد امنیتی وارد بفرمایید.
              </p>
              <div class="flex flex-col w-full gap-8">
                <div class="relative">
                  <label>
                    <input
                      name="mobile"
                      id="mobile"
                      type="tel"
                      class="w-full focus:outline-none peer border-b pr-5 border-[#85888B] pb-2 placeholder:text-[#C2C4C5] focus:border-[#304A79] transition-colors duration-200 rtl"
                      placeholder="تلفن همراه"
                    />
                  </label>
                  <svg
  width="16"
  height="16"
  viewBox="0 0 16 16"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  class="absolute top-1"
>
  <path
    d="M3.33301 6.00065C3.33301 3.80076 3.33301 2.70082 4.01643 2.0174C4.69984 1.33398 5.79979 1.33398 7.99967 1.33398C10.1996 1.33398 11.2995 1.33398 11.9829 2.0174C12.6663 2.70082 12.6663 3.80076 12.6663 6.00065V10.0007C12.6663 12.2005 12.6663 13.3005 11.9829 13.9839C11.2995 14.6673 10.1996 14.6673 7.99967 14.6673C5.79979 14.6673 4.69984 14.6673 4.01643 13.9839C3.33301 13.3005 3.33301 12.2005 3.33301 10.0007V6.00065Z"
    stroke="#AAACAE"
    stroke-linecap="round"
  />
  <path
    d="M7.33301 12.668H8.66634"
    stroke="#AAACAE"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
  <path
    d="M6 1.33398L6.05934 1.68999C6.18792 2.46151 6.25221 2.84727 6.51679 3.08201C6.79278 3.32688 7.1841 3.33398 8 3.33398C8.8159 3.33398 9.20721 3.32688 9.48321 3.08201C9.74779 2.84727 9.81208 2.46151 9.94067 1.68999L10 1.33398"
    stroke="#AAACAE"
    stroke-linejoin="round"
  />
</svg>
                </div>
                <div
                  class="frc-captcha w-full !max-w-none !pb-0"
                  id="frc-captcha"
                  data-sitekey="FCMGEMUD2LR7ICLH"
                  data-start="auto"
                  data-callback="onFrcDone"
                  data-lang="fa"
                  style="display:none"
                ></div>
                <div
                  id="simple-captcha"
                  style="display:block"
                  class="relative"
                >
                  <label>
                    <input
                      name="captcha_value"
                      id="captcha-value"
                      type="tel"
                      class="w-full focus:outline-none focus:border-[#304A79] peer border-b pr-5 border-[#85888B] pb-2 placeholder:text-[#C2C4C5] rtl"
                      placeholder="کد امنیتی"
                    />
                  </label>

                  <svg
  width="17"
  height="16"
  viewBox="0 0 17 16"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  class="absolute top-1"
>
  <path
    d="M14.7984 5.46602C14.8333 6.16611 14.8333 6.99892 14.8333 7.99935C14.8333 10.9849 14.8333 12.4777 13.9058 13.4052C12.9783 14.3327 11.4856 14.3327 8.49999 14.3327C5.51443 14.3327 4.02165 14.3327 3.09415 13.4052C2.16666 12.4777 2.16666 10.9849 2.16666 7.99935C2.16666 5.01379 2.16666 3.52101 3.09415 2.59351C4.02165 1.66602 5.51443 1.66602 8.49999 1.66602C9.21461 1.66602 9.8437 1.66602 10.4 1.67873"
    stroke="#AAACAE"
    stroke-linecap="round"
  />
  <path
    d="M5.83334 7.66602C5.83334 7.66602 6.83334 7.66602 8.16668 9.99935C8.16668 9.99935 11.5392 3.88824 14.8333 2.66602"
    stroke="#AAACAE"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
</svg>

                  <div
                    class="absolute top-0 left-0 flex items-center gap-2 cursor-pointer"
                  >
                    <img
                      src=""
                      alt="captcha"
                      class="captcha-image reset-captcha w-[120px]"
                    />
                    <button type="button" class="voice-captcha cursor-pointer">
                      <svg
  width="25"
  height="24"
  viewBox="0 0 25 24"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M14.5 14.8135V9.18646C14.5 6.04126 14.5 4.46866 13.5747 4.0773C12.6494 3.68593 11.5603 4.79793 9.38232 7.02192C8.25439 8.17365 7.61085 8.42869 6.00604 8.42869C4.60257 8.42869 3.90084 8.42869 3.39675 8.77262C2.35035 9.48655 2.50852 10.882 2.50852 12C2.50852 13.118 2.35035 14.5134 3.39675 15.2274C3.90084 15.5713 4.60257 15.5713 6.00604 15.5713C7.61085 15.5713 8.25439 15.8264 9.38232 16.9781C11.5603 19.2021 12.6494 20.3141 13.5747 19.9227C14.5 19.5313 14.5 17.9587 14.5 14.8135Z"
    stroke="#8C8F92"
    stroke-width="1.5"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
  <path
    d="M17.5 9C18.1254 9.81968 18.5 10.8634 18.5 12C18.5 13.1366 18.1254 14.1803 17.5 15"
    stroke="#8C8F92"
    stroke-width="1.5"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
  <path
    d="M20.5 7C21.7508 8.36613 22.5 10.1057 22.5 12C22.5 13.8943 21.7508 15.6339 20.5 17"
    stroke="#8C8F92"
    stroke-width="1.5"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
</svg>
                    </button>
                    <div class="w-px h-4 bg-[#DADBDC] rounded-full"></div>
                    <button type="button" class="reset-captcha cursor-pointer">
                      <svg
  width="22"
  height="24"
  viewBox="0 0 22 24"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M14.1667 0.999756L14.7646 2.11753C15.1689 2.87322 15.371 3.25107 15.2374 3.41289C15.1037 3.57471 14.6635 3.44402 13.7831 3.18264C12.9029 2.92131 11.9684 2.78071 11 2.78071C5.75329 2.78071 1.5 6.90822 1.5 11.9998C1.5 13.6789 1.96262 15.2533 2.77093 16.6093M7.83333 22.9998L7.23536 21.882C6.83108 21.1263 6.62894 20.7484 6.7626 20.5866C6.89627 20.4248 7.33649 20.5555 8.21689 20.8169C9.09709 21.0782 10.0316 21.2188 11 21.2188C16.2467 21.2188 20.5 17.0913 20.5 11.9998C20.5 10.3206 20.0374 8.74623 19.2291 7.39023"
    stroke="#AAACAE"
    stroke-width="1.5"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
</svg>
                    </button>
                  </div>
                </div>
                <input type="hidden" class="captcha-key" name="captcha_key" />
              </div>

              <div class="w-full text-right pt-2 errors-holder">



              </div>

              <button
                id="submit-login-button"
                disabled
                data-loading="false"
                class="w-full h-[40px] py-1.5 rounded-[10px] text-white disabled:bg-[#DADBDC] disabled:text-[#85888B] not-disabled:cursor-pointer bg-[#3F619D] shadow-[4px_4px_8px_0px_#00000040] sm:mt-[64px] mt-[58px]"
              >
                <span class="hidden not-in-data-[loading=true]:block">
                  <span
                    class="hidden group-[&[data-active-form=one-time-pass]]:block"
                    >ارسال کد</span
                  >
                  <span
                    class="hidden group-[&[data-active-form=application]]:block"
                    >ارسال رمز یکبار مصرف</span
                  >
                </span>
                <div
                  class="hidden in-data-[loading=true]:block animate-spin rounded-full mx-auto before:absolute before:content-[''] before:top-0 before:left-0 before:w-full before:h-full before:rounded-full w-[20px] h-[20px] shadow-[#85888B] shadow-[1px_2px_2px]"
                ></div>
              </button>

              <div
                class="w-full flex items-center justify-center mt-8 sm:mt-[48px] mb-6"
              >
                <div
                  class="w-full h-px bg-gradient-to-r from-[#C2C4C5] to-white"
                ></div>
                <p class="mx-2">یا</p>
                <div
                  class="w-full h-px bg-gradient-to-l from-[#C2C4C5] to-white"
                ></div>
              </div>

              <div class="flex w-full gap-8 items-center justify-center">
                <div class="group/tooltip relative">
                  <button
                    type="button"
                    class="login-with-qr-btn peer w-[64px] h-[64px] rounded-lg cursor-pointer"
                    style="
                      background: linear-gradient(
                        199.93deg,
                        rgba(199, 210, 230, 0.5) 8.72%,
                        rgba(255, 255, 255, 0.5) 87.25%
                      );
                      box-shadow: 0px 5.4px 5.4px 0px #00000040;
                      backdrop-filter: blur(344.82763671875px);
                    "
                  >
                    <svg
  width="67"
  height="64"
  viewBox="0 0 67 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#qr-btn)">
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M10.5 11.3C10.5 10.0297 11.5298 9 12.8 9H28.9C30.1703 9 31.2 10.0297 31.2 11.3V27.3997C31.2 28.67 30.1703 29.6997 28.9 29.6997H12.8C11.5298 29.6997 10.5 28.67 10.5 27.3997V11.3ZM17.4 18.1999C17.4 16.9296 18.4298 15.8999 19.7 15.8999H22C23.2702 15.8999 24.3 16.9296 24.3 18.1999V20.4998C24.3 21.7701 23.2702 22.7998 22 22.7998H19.7C18.4298 22.7998 17.4 21.7701 17.4 20.4998V18.1999Z"
      fill="url(#paint0_linear_3769_168764)"
    />
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M10.5 36.6007C10.5 35.3305 11.5298 34.3008 12.8 34.3008H28.9C30.1703 34.3008 31.2 35.3305 31.2 36.6007V52.7005C31.2 53.9708 30.1703 55.0005 28.9 55.0005H12.8C11.5298 55.0005 10.5 53.9708 10.5 52.7005V36.6007ZM19.7 41.2007C18.4298 41.2007 17.4 42.2304 17.4 43.5007V45.8006C17.4 47.0709 18.4298 48.1006 19.7 48.1006H22C23.2702 48.1006 24.3 47.0709 24.3 45.8006V43.5007C24.3 42.2304 23.2702 41.2007 22 41.2007H19.7Z"
      fill="url(#paint1_linear_3769_168764)"
    />
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M38.1 9C36.8297 9 35.8 10.0297 35.8 11.3V27.3997C35.8 28.67 36.8297 29.6997 38.1 29.6997H54.2C55.4703 29.6997 56.5 28.67 56.5 27.3997V11.3C56.5 10.0297 55.4703 9 54.2 9H38.1ZM42.7 18.1999C42.7 16.9296 43.7297 15.8999 45 15.8999H47.3C48.5703 15.8999 49.6 16.9296 49.6 18.1999V20.4998C49.6 21.7701 48.5703 22.7998 47.3 22.7998H45C43.7297 22.7998 42.7 21.7701 42.7 20.4998V18.1999Z"
      fill="url(#paint2_linear_3769_168764)"
    />
    <path
      d="M38.1 34.3008C36.8297 34.3008 35.8 35.3305 35.8 36.6007V38.9007C35.8 40.171 36.8297 41.2007 38.1 41.2007H40.4V45.8006C40.4 47.0709 41.4297 48.1006 42.7 48.1006H45C46.2703 48.1006 47.3 47.0709 47.3 45.8006V41.2007C47.3 39.9304 46.2703 38.9007 45 38.9007H42.7V36.6007C42.7 35.3305 41.6703 34.3008 40.4 34.3008H38.1Z"
      fill="url(#paint3_linear_3769_168764)"
    />
    <path
      d="M38.1 48.0996C36.8297 48.0996 35.8 49.1293 35.8 50.3996V52.6995C35.8 53.9698 36.8297 54.9995 38.1 54.9995H40.4C41.6703 54.9995 42.7 53.9698 42.7 52.6995V50.3996C42.7 49.1293 41.6703 48.0996 40.4 48.0996H38.1Z"
      fill="url(#paint4_linear_3769_168764)"
    />
    <path
      d="M47.3 36.6007C47.3 35.3305 48.3297 34.3008 49.6 34.3008H51.9C53.1703 34.3008 54.2 35.3305 54.2 36.6007V38.9007C54.2 40.171 53.1703 41.2007 51.9 41.2007H49.6C48.3297 41.2007 47.3 40.171 47.3 38.9007V36.6007Z"
      fill="url(#paint5_linear_3769_168764)"
    />
    <path
      d="M49.6 48.0996C48.3297 48.0996 47.3 49.1293 47.3 50.3996V52.6995C47.3 53.9698 48.3297 54.9995 49.6 54.9995H51.9C53.1703 54.9995 54.2 53.9698 54.2 52.6995V50.3996C54.2 49.1293 53.1703 48.0996 51.9 48.0996H49.6Z"
      fill="url(#paint6_linear_3769_168764)"
    />
  </g>
</svg>
                  </button>
                  <div
                    class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                  >
                    <div class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"></div>
                    <p>ورود با QR کد</p>
                  </div>
                </div>

                <div class="group/tooltip relative">
                  <button
                    type="button"
                    data-activate-form="application"
                    data-toggleable="true"
                    class="login-with-app-btn w-[64px] h-[64px] rounded-lg cursor-pointer"
                    style="
                      background: linear-gradient(
                        199.93deg,
                        rgba(199, 210, 230, 0.5) 8.72%,
                        rgba(255, 255, 255, 0.5) 87.25%
                      );
                      box-shadow: 0px 5.4px 5.4px 0px #00000040;
                      backdrop-filter: blur(344.82763671875px);
                    "
                  >
                    <div
                      class="w-full flex gap-2 items-center justify-center hidden group-[&[data-active-form=one-time-pass]]:block"
                    >
                      <svg
  width="53"
  height="64"
  viewBox="0 0 53 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#app-btn)">
    <path
      d="M33.7454 9H19.2547C15.5328 9 12.5 11.9647 12.5 15.6128V48.3872C12.5 52.0309 15.5283 55 19.2547 55H33.743C37.4648 55 40.4976 52.0353 40.4976 48.3872L40.5 15.6128C40.5 11.9714 37.467 9 33.7454 9ZM26.5037 49.645C25.2092 49.645 24.157 48.6149 24.157 47.3477C24.157 46.0804 25.2092 45.0503 26.5037 45.0503C27.7981 45.0503 28.8503 46.0804 28.8503 47.3477C28.8456 48.6149 27.7981 49.645 26.5037 49.645ZM29.4393 15.864H23.5585C22.8887 15.864 22.3424 15.3292 22.3424 14.6734C22.3424 14.0177 22.8887 13.4829 23.5585 13.4829H29.4393C30.1091 13.4829 30.6554 14.0177 30.6554 14.6734C30.6554 15.3269 30.1163 15.864 29.4393 15.864Z"
      fill="url(#paint0_linear_3769_168762)"
    />
  </g>
</svg>
                    </div>
                    <div
                      class="w-full flex gap-2 items-center justify-center hidden group-[&[data-active-form=application]]:block"
                    >
                      <svg
  width="65"
  height="64"
  viewBox="0 0 65 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#otp-btn)">
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M39.5 36.999V47.3983H9.5V11.9999C9.5 11.2039 9.816 10.4419 10.378 9.87998C10.94 9.31599 11.7039 9 12.4999 9H36.4998C37.2959 9 38.0579 9.31599 38.6218 9.87998C39.1838 10.442 39.4998 11.2039 39.4998 11.9999V15.7997H18.0998C17.1978 15.7997 16.3338 16.1577 15.6958 16.7956C15.0578 17.4336 14.6999 18.2976 14.6999 19.1995V33.5988C14.6999 34.5007 15.0578 35.3667 15.6958 36.0026C16.3338 36.6406 17.1998 36.9986 18.1018 36.9966H20.2998C20.2998 36.9966 20.2858 39.4104 20.2998 39.9984C20.3038 40.1744 20.3278 40.3584 20.3798 40.5264C20.4278 40.6864 20.5078 40.8544 20.6038 40.9984C20.6918 41.1284 20.8198 41.2744 20.9538 41.3863C21.0718 41.4823 21.2478 41.5943 21.4118 41.6623C21.5438 41.7163 21.7558 41.7743 21.9298 41.7903C22.0978 41.8063 22.3038 41.7923 22.4518 41.7643C22.6218 41.7303 22.8198 41.6563 22.9438 41.5883C23.0978 41.5063 23.2618 41.3803 23.3738 41.2704C24.1578 40.5064 27.6458 36.9985 27.6458 36.9985L39.5 36.999ZM39.5 48.5983V51.9981C39.5 52.7941 39.184 53.5581 38.622 54.12C38.06 54.684 37.2961 55 36.5001 55L12.5002 54.998C11.7042 54.998 10.9421 54.682 10.3782 54.118C9.81616 53.556 9.5002 52.7921 9.5002 51.9961V48.5963H39.5002L39.5 48.5983ZM27.1481 35.799L22.5241 40.4227C22.3521 40.5947 22.0941 40.6467 21.8701 40.5527C21.6461 40.4607 21.5001 40.2407 21.5001 39.9988V35.795L18.1021 35.797C17.5182 35.797 16.9581 35.565 16.5461 35.153C16.1341 34.741 15.9021 34.1811 15.9021 33.5971L15.9001 19.1999C15.9001 18.6159 16.1321 18.0579 16.5441 17.6439C16.9561 17.2319 17.5161 17 18.1001 17H53.3C53.884 17 54.444 17.2319 54.856 17.6439C55.268 18.0559 55.5 18.6159 55.5 19.1999V33.5991C55.5 34.8131 54.516 35.799 53.3 35.799L27.1481 35.799ZM38.6178 26.3996L36.9658 27.4996C36.6898 27.6836 36.6158 28.0555 36.7998 28.3315C36.9838 28.6075 37.3558 28.6795 37.6318 28.4975L39.0998 27.5195V28.7975C39.0998 29.1295 39.3698 29.3975 39.6998 29.3975C40.0318 29.3975 40.2998 29.1295 40.2998 28.7975V27.5195L41.7678 28.4975C42.0438 28.6815 42.4158 28.6075 42.5998 28.3315C42.7838 28.0555 42.7078 27.6836 42.4338 27.4996L40.7838 26.3996L42.4338 25.2997C42.7098 25.1157 42.7838 24.7438 42.5998 24.4678C42.4158 24.1918 42.0438 24.1178 41.7678 24.3018L40.2998 25.2797V23.9998C40.2998 23.6698 40.0318 23.3998 39.6998 23.3998C39.3698 23.3998 39.0998 23.6698 39.0998 23.9998V25.2797L37.6318 24.3018C37.3558 24.1178 36.9838 24.1918 36.7998 24.4678C36.6158 24.7437 36.6918 25.1157 36.9658 25.2997L38.6178 26.3996ZM22.6179 26.3996L20.9659 27.4996C20.6899 27.6836 20.6159 28.0555 20.7999 28.3315C20.9839 28.6075 21.3559 28.6795 21.6319 28.4975L23.0999 27.5195V28.7975C23.0999 29.1295 23.3679 29.3975 23.6999 29.3975C24.0299 29.3975 24.2999 29.1295 24.2999 28.7975V27.5195L25.7679 28.4975C26.0439 28.6815 26.4159 28.6075 26.5999 28.3315C26.7839 28.0555 26.7099 27.6836 26.4339 27.4996L24.7819 26.3996L26.4339 25.2997C26.7099 25.1157 26.7839 24.7438 26.5999 24.4678C26.4159 24.1918 26.0439 24.1178 25.7679 24.3018L24.2999 25.2797V23.9998C24.2999 23.6698 24.0319 23.3998 23.6999 23.3998C23.3679 23.3998 23.0999 23.6698 23.0999 23.9998V25.2797L21.6319 24.3018C21.3559 24.1178 20.9839 24.1918 20.7999 24.4678C20.6159 24.7437 20.6919 25.1157 20.9659 25.2997L22.6179 26.3996ZM30.6179 26.3996L28.9659 27.4996C28.6899 27.6836 28.6159 28.0555 28.7998 28.3315C28.9838 28.6075 29.3558 28.6795 29.6318 28.4975L31.0998 27.5195V28.7975C31.0998 29.1295 31.3678 29.3975 31.6999 29.3975C32.0319 29.3975 32.2999 29.1295 32.2999 28.7975V27.5195L33.7679 28.4975C34.0439 28.6815 34.4159 28.6075 34.5999 28.3315C34.7839 28.0555 34.7099 27.6836 34.4339 27.4996L32.7819 26.3996L34.4339 25.2997C34.7099 25.1157 34.7839 24.7438 34.5999 24.4678C34.4159 24.1918 34.0439 24.1178 33.7679 24.3018L32.2999 25.2797V23.9998C32.2999 23.6698 32.0319 23.3998 31.6999 23.3998C31.3699 23.3998 31.0998 23.6698 31.0998 23.9998V25.2797L29.6318 24.3018C29.3558 24.1178 28.9839 24.1918 28.7998 24.4678C28.6158 24.7437 28.6899 25.1157 28.9659 25.2997L30.6179 26.3996ZM46.6178 26.3996L44.9678 27.4996C44.6918 27.6836 44.6178 28.0555 44.8018 28.3315C44.9858 28.6075 45.3578 28.6795 45.6338 28.4975L47.1018 27.5195V28.7975C47.1018 29.1295 47.3718 29.3975 47.7018 29.3975C48.0318 29.3975 48.3018 29.1295 48.3018 28.7975V27.5195L49.7698 28.4975C50.0458 28.6815 50.4178 28.6075 50.6018 28.3315C50.7858 28.0555 50.7098 27.6836 50.4358 27.4996L48.7838 26.3996L50.4358 25.2997C50.7118 25.1157 50.7858 24.7438 50.6018 24.4678C50.4178 24.1918 50.0458 24.1178 49.7698 24.3018L48.3018 25.2797V23.9998C48.3018 23.6698 48.0318 23.3998 47.7018 23.3998C47.3718 23.3998 47.1018 23.6698 47.1018 23.9998V25.2797L45.6338 24.3018C45.3578 24.1178 44.9858 24.1918 44.8018 24.4678C44.6178 24.7437 44.6938 25.1157 44.9678 25.2997L46.6178 26.3996Z"
      fill="url(#paint0_linear_3876_137583)"
    />
  </g>
</svg>
                    </div>
                  </button>
                  <div
                    class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                  >
                    <div class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"></div>
                    <p
                      class="hidden group-[&[data-active-form=one-time-pass]]:block"
                    >
                      ورود با اپلیکیشن
                    </p>
                    <p
                      class="hidden group-[&[data-active-form=application]]:block"
                    >
                      ورود با رمز یکبارمصرف
                    </p>
                  </div>
                </div>
              </div>
            </form>

            <!--........................OTP code login............................-->
            <form
              method="post"
              action="/login_otp"
              id="otp-form"
              class="hidden mt-2.5 md:mt-4"
              novalidate
            >
              <div class="w-full">
                <p class="text-[#5D5F61] text-sm md:text-lg text-justify">
                  کد فعالسازی <b class="text-[#4469AD]">5</b> رقمی به تلفن همراه
                  <b class="text-[#4469AD]" id="mobile-placeholder"></b> ارسال
                  شده است.
                </p>

                <div
                  id="change-phone"
                  class="w-full flex items-center justify-start gap-1 mt-1 cursor-pointer"
                >
                  <svg
  width="18"
  height="18"
  viewBox="0 0 18 18"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M4.15499 14.6406C3.69749 14.6406 3.26999 14.4831 2.96249 14.1906C2.57249 13.8231 2.38499 13.2681 2.45249 12.6681L2.72999 10.2381C2.78249 9.78065 3.05999 9.17315 3.38249 8.84315L9.53999 2.32565C11.0775 0.698147 12.6825 0.653147 14.31 2.19065C15.9375 3.72815 15.9825 5.33315 14.445 6.96065L8.28749 13.4781C7.97249 13.8156 7.38749 14.1306 6.92999 14.2056L4.51499 14.6181C4.38749 14.6256 4.27499 14.6406 4.15499 14.6406ZM11.9475 2.18315C11.37 2.18315 10.8675 2.54315 10.3575 3.08315L4.19999 9.60815C4.04999 9.76565 3.87749 10.1406 3.84749 10.3581L3.56999 12.7881C3.53999 13.0356 3.59999 13.2381 3.73499 13.3656C3.86999 13.4931 4.07249 13.5381 4.31999 13.5006L6.73499 13.0881C6.95249 13.0506 7.31249 12.8556 7.46249 12.6981L13.62 6.18065C14.55 5.19065 14.8875 4.27565 13.53 3.00065C12.93 2.42315 12.4125 2.18315 11.9475 2.18315Z"
    fill="#3E96F5"
  />
  <path
    d="M13.0048 8.21263C12.9898 8.21263 12.9673 8.21263 12.9523 8.21263C10.6123 7.98013 8.72984 6.20263 8.36984 3.87763C8.32484 3.57013 8.53484 3.28513 8.84234 3.23263C9.14984 3.18763 9.43484 3.39763 9.48734 3.70513C9.77234 5.52013 11.2423 6.91513 13.0723 7.09513C13.3798 7.12513 13.6048 7.40263 13.5748 7.71013C13.5373 7.99513 13.2898 8.21263 13.0048 8.21263Z"
    fill="#3E96F5"
  />
  <path
    d="M15.75 17.0625H2.25C1.9425 17.0625 1.6875 16.8075 1.6875 16.5C1.6875 16.1925 1.9425 15.9375 2.25 15.9375H15.75C16.0575 15.9375 16.3125 16.1925 16.3125 16.5C16.3125 16.8075 16.0575 17.0625 15.75 17.0625Z"
    fill="#3E96F5"
  />
</svg>
                  <p class="text-sm text-[#3E96F5]">ویرایش شماره موبایل</p>
                </div>

                <div
                  class="relative flex gap-4 items-start w-full overflow-hidden border border-white rounded-2xl p-4 mt-8 shadow-[-2px_-2px_16px_0px_#00000014] z-[1]"
                  style="
                    background: linear-gradient(
                      220deg,
                      #d5e9ff 10.47%,
                      #ffffff 29.17%
                    );
                  "
                >
                  <div class="absolute right-[-2px] top-0 z-[-10]">
                    <svg
  width="169"
  height="100"
  viewBox="0 0 169 100"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M215.5 23.5H192.5V0H191.5V23.5H168.5V0H167.5V23.5H144.5V0H143.5V23.5H120.5V0H119.5V23.5H96.5V0H95.5V23.5H72.5V0H71.5V23.5H48.5V0H47.5V23.5H24.5V0H23.5V23.5H0V24.5H23.5V47.5H0V48.5H23.5V71.5H0V72.5H23.5V95.5H0V96.5H23.5V119.5H0V120.5H23.5V143.5H0V144.5H23.5V167.5H0V168.5H23.5V192H24.5V168.5H47.5V192H48.5V168.5H71.5V192H72.5V168.5H95.5V192H96.5V168.5H119.5V192H120.5V168.5H143.5V192H144.5V168.5H167.5V192H168.5V168.5H191.5V192H192.5V168.5H215.5V192H216.5V168.5H240V167.5H216.5V144.5H240V143.5H216.5V120.5H240V119.5H216.5V96.5H240V95.5H216.5V72.5H240V71.5H216.5V48.5H240V47.5H216.5V24.5H240V23.5H216.5V0H215.5V23.5ZM215.5 144.5V167.5H192.5V144.5H215.5ZM191.5 144.5V167.5H168.5V144.5H191.5ZM167.5 144.5V167.5H144.5V144.5H167.5ZM143.5 144.5V167.5H120.5V144.5H143.5ZM119.5 144.5V167.5H96.5V144.5H119.5ZM95.5 144.5V167.5H72.5V144.5H95.5ZM71.5 144.5V167.5H48.5V144.5H71.5ZM47.5 144.5V167.5H24.5V144.5H47.5ZM215.5 120.5V143.5H192.5V120.5H215.5ZM191.5 120.5V143.5H168.5V120.5H191.5ZM167.5 120.5V143.5H144.5V120.5H167.5ZM143.5 120.5V143.5H120.5V120.5H143.5ZM119.5 120.5V143.5H96.5V120.5H119.5ZM95.5 120.5V143.5H72.5V120.5H95.5ZM71.5 120.5V143.5H48.5V120.5H71.5ZM47.5 120.5V143.5H24.5V120.5H47.5ZM215.5 96.5V119.5H192.5V96.5H215.5ZM191.5 96.5V119.5H168.5V96.5H191.5ZM167.5 96.5V119.5H144.5V96.5H167.5ZM143.5 96.5V119.5H120.5V96.5H143.5ZM119.5 96.5V119.5H96.5V96.5H119.5ZM95.5 96.5V119.5H72.5V96.5H95.5ZM71.5 96.5V119.5H48.5V96.5H71.5ZM47.5 96.5V119.5H24.5V96.5H47.5ZM215.5 72.5V95.5H192.5V72.5H215.5ZM191.5 72.5V95.5H168.5V72.5H191.5ZM167.5 72.5V95.5H144.5V72.5H167.5ZM143.5 72.5V95.5H120.5V72.5H143.5ZM119.5 72.5V95.5H96.5V72.5H119.5ZM95.5 72.5V95.5H72.5V72.5H95.5ZM71.5 72.5V95.5H48.5V72.5H71.5ZM47.5 72.5V95.5H24.5V72.5H47.5ZM215.5 48.5V71.5H192.5V48.5H215.5ZM191.5 48.5V71.5H168.5V48.5H191.5ZM167.5 48.5V71.5H144.5V48.5H167.5ZM143.5 48.5V71.5H120.5V48.5H143.5ZM119.5 48.5V71.5H96.5V48.5H119.5ZM95.5 48.5V71.5H72.5V48.5H95.5ZM71.5 48.5V71.5H48.5V48.5H71.5ZM47.5 48.5V71.5H24.5V48.5H47.5ZM215.5 24.5V47.5H192.5V24.5H215.5ZM191.5 24.5V47.5H168.5V24.5H191.5ZM167.5 24.5V47.5H144.5V24.5H167.5ZM143.5 24.5V47.5H120.5V24.5H143.5ZM119.5 24.5V47.5H96.5V24.5H119.5ZM95.5 24.5V47.5H72.5V24.5H95.5ZM71.5 24.5V47.5H48.5V24.5H71.5ZM47.5 24.5V47.5H24.5V24.5H47.5Z"
    fill="url(#grid-box)"
  />
</svg>
                  </div>
                  <div class="bg-[#C0DEFE] rounded-lg opacity-50 p-2 mt-1">
                    <svg
  width="32"
  height="32"
  viewBox="0 0 32 32"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M29.3332 16.0007C29.3332 8.63686 23.3636 2.66732 15.9998 2.66732C8.63604 2.66732 2.6665 8.63686 2.6665 16.0007C2.6665 23.3644 8.63604 29.334 15.9998 29.334C23.3636 29.334 29.3332 23.3644 29.3332 16.0007Z"
    stroke="#2E90FA"
    stroke-width="1.5"
  />
  <path
    d="M16.3226 22.666V15.9993C16.3226 15.3708 16.3226 15.0565 16.1273 14.8613C15.9321 14.666 15.6178 14.666 14.9893 14.666"
    stroke="#2E90FA"
    stroke-width="1.5"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
  <path
    d="M15.989 10.666H16.001"
    stroke="#2E90FA"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
</svg>
                  </div>
                  <div class="flex flex-col items-start">
                    <p class="text-lg text-[#2E90FA] font-semibold">
                      دریافت مجدد کد
                    </p>
                    <p class="text-right text-sm text-[#85888B]">
                      در صورت عدم دریافت کد فعال‌سازی از طریق پیامک، می‌توانید
                      آن را از طریق کد دستوری
                      <span class="text-[#2E90FA]"># 40* ‎* 4</span> یا
                      <span class="text-[#2E90FA]">ایمیل ملی</span>
                      دریافت نمایید.
                    </p>
                  </div>
                </div>

                <!-- ......................................otp inputs.........................................-->
                <div id="otp-generator" class="mt-[40px] sm:mt-[64px]"></div>
                <div id="frc-captcha-holder" class="mt-8"></div>
                <input type="hidden" id="password" name="password" />
                <input
                  type="hidden"
                  name="_csrf"
                  value="yAtpHmp7GsEKyyxyDAxnFXTIEAYmsSd2_4TL6W_0UWedrJKP8Gpafw4eePMn8k9KOyFTcRHwPT9C0kNbmrH4iFuRaFeuzarp"
                  header="X-CSRF-TOKEN"
                />
                <input
                  type="hidden"
                  name="username"
                  id="username"
                  class="username_final"
                  value=""
                />
                <button
                  disabled
                  id="otp-submit-btn"
                  class="mt-8 disabled:bg-[#DADBDC] disabled:text-[#85888B] not-disabled:cursor-pointer rounded-[10px] py-2 w-[85%] shadow-[4px_4px_8px_0px_#00000040] bg-[#3F619D] text-white"
                  data-init-text="ارسال مجدد کد فعال سازی"
                >
                  <!-- <p id="timer" class="flex-1 ltr">00 : 00</p> -->
                  <span
                    id="timer"
                    class="hidden not-in-data-[loading=true]:block flex-1 ltr"
                    data-finished-text="ارسال مجدد کد"
                    >00 : 00</span
                  >
                  <div
                    class="hidden in-data-[loading=true]:block animate-spin rounded-full mx-auto before:absolute before:content-[''] before:top-0 before:left-0 before:w-full before:h-full before:rounded-full w-[20px] h-[20px] shadow-[#85888B] shadow-[1px_2px_2px]"
                  ></div>
                </button>

                <div
                  class="w-full flex items-center justify-center mt-8 sm:mt-4 mb-2"
                >
                  <div
                    class="w-full h-px bg-gradient-to-r from-[#C2C4C5] to-white"
                  ></div>
                  <p class="mx-2">یا</p>
                  <div
                    class="w-full h-px bg-gradient-to-l from-[#C2C4C5] to-white"
                  ></div>
                </div>

                <div class="flex w-full gap-8 items-center justify-center">
                  <div class="group/tooltip relative">
                    <button
                      type="button"
                      id="send-otp-bale"
                      class="peer w-[64px] h-[64px] rounded-lg cursor-pointer"
                      data-loading="false"
                      style="
                        background: linear-gradient(
                          199.93deg,
                          rgba(199, 210, 230, 0.5) 8.72%,
                          rgba(255, 255, 255, 0.5) 87.25%
                        );
                        box-shadow: 0px 5.4px 5.4px 0px #00000040;
                        backdrop-filter: blur(344.82763671875px);
                      "
                    >
                      <svg
  width="64"
  height="64"
  viewBox="0 0 64 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  class="hidden not-in-data-[loading=true]:block"
>
  <g filter="url(#messenger-btn)">
    <path
      d="M40.6227 35.4793C39.4904 36.5892 38.1476 37.4623 36.673 38.0474C35.1984 38.6325 33.6216 38.9178 32.0351 38.8866C31.152 38.8866 30.2642 38.8866 29.3811 38.8866C29.2348 38.8771 29.0884 38.9012 28.9528 38.9569C28.8173 39.0126 28.6963 39.0985 28.5992 39.208C26.6075 41.302 24.5975 43.3822 22.6196 45.4669C22.4348 45.6846 22.2052 45.86 21.9465 45.9812C21.6877 46.1025 21.4059 46.1667 21.1201 46.1695C20.881 46.169 20.6444 46.1207 20.4243 46.0276C20.2041 45.9345 20.0048 45.7983 19.8381 45.6272C19.6714 45.4561 19.5406 45.2534 19.4535 45.0311C19.3664 44.8088 19.3247 44.5714 19.3308 44.3327C19.3308 42.5969 19.3308 40.8627 19.3308 39.13C19.3308 38.7994 19.2388 38.6708 18.9031 38.6065C13.9998 37.3483 11.0101 34.1568 9.55656 29.4225C9.32401 28.5281 9.13824 27.6223 9 26.7086L9 24.1922C9.08867 23.3912 9.2425 22.5988 9.45997 21.8227C10.6605 17.4695 13.2455 14.3148 17.5508 12.6755C18.7884 12.2174 20.0993 11.9887 21.4191 12.0004C24.8826 12.0004 28.3508 12.0004 31.8189 12.0004C36.6578 11.9775 40.2777 14.1265 42.7615 18.1537C44.3949 20.8651 45.0717 24.0445 44.6841 27.1847C44.2964 30.3249 42.8666 33.245 40.6227 35.4793Z"
      fill="url(#paint0_linear_7670_204772)"
    />
    <path
      d="M51.5656 28.7211C50.2878 27.5736 48.7106 26.811 47.0165 26.5215C47.0165 26.6455 46.9843 26.7189 46.9751 26.7878C46.8218 29.0041 46.1945 31.1618 45.1353 33.1156C42.2927 38.2586 37.9828 41.069 32.0078 41.1424C31.3777 41.1424 30.7475 41.1424 30.1174 41.1424C29.6022 41.1424 29.3722 41.5373 29.6344 41.9644C31.0879 44.3339 33.121 45.9548 35.9038 46.437C37.1075 46.5843 38.3204 46.6442 39.5329 46.6161C39.6723 46.6075 39.8119 46.6308 39.941 46.6841C40.0701 46.7374 40.1854 46.8193 40.278 46.9238C41.0876 47.7871 41.9155 48.632 42.7343 49.4861C43.3644 50.1427 43.99 50.8132 44.6385 51.4607C44.9128 51.7748 45.3006 51.9676 45.717 51.9969C46.1334 52.0262 46.5445 51.8895 46.8602 51.6168C47.0505 51.444 47.1988 51.23 47.2937 50.9913C47.3885 50.7525 47.4275 50.4953 47.4075 50.2392C47.4075 49.1509 47.4075 48.058 47.4075 46.9697C47.4075 46.5794 47.5133 46.4416 47.8997 46.2946C48.8535 45.9909 49.7676 45.5747 50.6227 45.0548C56.0411 41.3353 56.5148 33.111 51.5656 28.7211Z"
      fill="url(#paint1_linear_7670_204772)"
    />
  </g>
</svg>

                      <div
                        class="hidden in-data-[loading=true]:block animate-spin rounded-full mx-auto before:absolute before:content-[''] before:top-0 before:left-0 before:w-full before:h-full before:rounded-full w-[20px] h-[20px] shadow-[#85888B] shadow-[1px_2px_2px]"
                      ></div>
                    </button>
                    <div
                      class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                    >
                      <div
                        class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"
                      ></div>
                      <p>ارسال از طریق پیام رسان</p>
                    </div>
                  </div>

                  <div class="group/tooltip relative">
                    <button
                      type="button"
                      id="send-otp-email"
                      class="w-[64px] h-[64px] rounded-lg cursor-pointer"
                      data-loading="false"
                      style="
                        background: linear-gradient(
                          199.93deg,
                          rgba(199, 210, 230, 0.5) 8.72%,
                          rgba(255, 255, 255, 0.5) 87.25%
                        );
                        box-shadow: 0px 5.4px 5.4px 0px #00000040;
                        backdrop-filter: blur(344.82763671875px);
                      "
                    >
                      <div
                        class="w-full flex gap-2 items-center justify-center hidden not-in-data-[loading=true]:block"
                      >
                        <svg
  width="64"
  height="56"
  viewBox="0 0 64 56"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#email-btn)">
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M55 15.1451L39.5849 27.9997L55 40.8542V15.1451ZM54.2298 42.8557C53.8003 43.2542 53.2262 43.5 52.6005 43.5H11.3995C10.7759 43.5 10.2039 43.2563 9.77641 42.8619L26.0135 29.3217L28.4936 31.3909C29.4704 32.2045 30.7486 32.6093 32.0206 32.6031C33.2906 32.5969 34.5709 32.1839 35.5538 31.362L38.0008 29.3217L54.2298 42.8557ZM9 40.865L24.4299 27.9999L9 15.1348V40.865ZM9.77646 13.1381C10.2039 12.7416 10.7759 12.5 11.3995 12.5H52.6006C53.2283 12.5 53.8024 12.7457 54.2298 13.1443L37.3652 27.2089L37.3425 27.2275L34.2574 29.8005C33.6462 30.3106 32.8326 30.5667 32.0148 30.5708C31.1991 30.5749 30.3917 30.3271 29.7929 29.8295L26.6748 27.2296L26.6479 27.2069L9.77646 13.1381Z"
      fill="url(#paint0_linear_3790_173457)"
    />
  </g>
</svg>
                      </div>
                      <div
                        class="hidden in-data-[loading=true]:block animate-spin rounded-full mx-auto before:absolute before:content-[''] before:top-0 before:left-0 before:w-full before:h-full before:rounded-full w-[20px] h-[20px] shadow-[#85888B] shadow-[1px_2px_2px]"
                      ></div>
                    </button>
                    <div
                      class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                    >
                      <div
                        class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"
                      ></div>
                      <p>ارسال از طریق ایمیل</p>
                    </div>
                  </div>
                </div>
              </div>
            </form>
            <!--........................Application login........................-->
            <form
              method="post"
              action="/login_app"
              id="app-code-form"
              class="hidden"
              novalidate
            >
              <div
                class="relative flex flex-col items-center justify-center gap-4 w-full"
              >
                <div
                  class="flex items-start w-full max-w-[450px] rounded-2xl text-[#797979] text-justify p-4"
                >
                  <p>
                    کاربر گرامی، لطفاً اپلیکیشن «دولت من» را اجرا کرده و عدد زیر
                    را از بین گزینه‌های نمایش‌داده‌شده در اپ انتخاب کنید. در
                    صورت انتخاب صحیح، ورود شما به‌صورت خودکار انجام می‌شود. برای
                    نصب اپلیکیشن،
                    <a
                      href="#"
                      id="download-app"
                      class="text-[#2E90FA] underline"
                      style="color: #48acde"
                      >کلیک کنید.</a
                    >
                  </p>
                </div>

                <div
                  class="flex items-center justify-center bg-gradient-to-t from-white to-[#e3e3e3] border-2 border-[#e3e3e3] rounded-full shadow-[4px_4px_10px_0px_rgba(0,0,0,0.1)] p-2"
                >
                  <div
                    id="app-number-box"
                    class="number-box relative inline-flex items-center justify-center p-2.5 text-[40px] bg-white font-bold w-[150px] h-[150px] overflow-hidden rounded-full"
                    data-after=""
                  ></div>
                </div>

                <input
                  type="hidden"
                  name="_csrf"
                  value="yAtpHmp7GsEKyyxyDAxnFXTIEAYmsSd2_4TL6W_0UWedrJKP8Gpafw4eePMn8k9KOyFTcRHwPT9C0kNbmrH4iFuRaFeuzarp"
                  header="X-CSRF-TOKEN"
                />
                <input
                  type="hidden"
                  name="username"
                  class="username_final"
                  id="appcode-username-login"
                />
                <input
                  type="hidden"
                  name="password"
                  id="password-application-login"
                  class="password_final"
                />

                <p class="text-[#787A7D] font-semibold mt-4">
                  در انتظار تایید در اپلیکیشن
                </p>

                <div
                  class="w-full flex items-center justify-center mt-8 sm:mt-4 mb-2"
                >
                  <div
                    class="w-full h-px bg-gradient-to-r from-[#C2C4C5] to-white"
                  ></div>
                  <p class="mx-2">یا</p>
                  <div
                    class="w-full h-px bg-gradient-to-l from-[#C2C4C5] to-white"
                  ></div>
                </div>

                <div class="flex w-full gap-8 items-center justify-center">
                  <div class="group/tooltip relative">
                    <button
                      type="button"
                      class="login-with-qr-btn peer w-[64px] h-[64px] rounded-lg cursor-pointer"
                      style="
                        background: linear-gradient(
                          199.93deg,
                          rgba(199, 210, 230, 0.5) 8.72%,
                          rgba(255, 255, 255, 0.5) 87.25%
                        );
                        box-shadow: 0px 5.4px 5.4px 0px #00000040;
                        backdrop-filter: blur(344.82763671875px);
                      "
                    >
                      <svg
  width="67"
  height="64"
  viewBox="0 0 67 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#qr-btn)">
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M10.5 11.3C10.5 10.0297 11.5298 9 12.8 9H28.9C30.1703 9 31.2 10.0297 31.2 11.3V27.3997C31.2 28.67 30.1703 29.6997 28.9 29.6997H12.8C11.5298 29.6997 10.5 28.67 10.5 27.3997V11.3ZM17.4 18.1999C17.4 16.9296 18.4298 15.8999 19.7 15.8999H22C23.2702 15.8999 24.3 16.9296 24.3 18.1999V20.4998C24.3 21.7701 23.2702 22.7998 22 22.7998H19.7C18.4298 22.7998 17.4 21.7701 17.4 20.4998V18.1999Z"
      fill="url(#paint0_linear_3769_168764)"
    />
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M10.5 36.6007C10.5 35.3305 11.5298 34.3008 12.8 34.3008H28.9C30.1703 34.3008 31.2 35.3305 31.2 36.6007V52.7005C31.2 53.9708 30.1703 55.0005 28.9 55.0005H12.8C11.5298 55.0005 10.5 53.9708 10.5 52.7005V36.6007ZM19.7 41.2007C18.4298 41.2007 17.4 42.2304 17.4 43.5007V45.8006C17.4 47.0709 18.4298 48.1006 19.7 48.1006H22C23.2702 48.1006 24.3 47.0709 24.3 45.8006V43.5007C24.3 42.2304 23.2702 41.2007 22 41.2007H19.7Z"
      fill="url(#paint1_linear_3769_168764)"
    />
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M38.1 9C36.8297 9 35.8 10.0297 35.8 11.3V27.3997C35.8 28.67 36.8297 29.6997 38.1 29.6997H54.2C55.4703 29.6997 56.5 28.67 56.5 27.3997V11.3C56.5 10.0297 55.4703 9 54.2 9H38.1ZM42.7 18.1999C42.7 16.9296 43.7297 15.8999 45 15.8999H47.3C48.5703 15.8999 49.6 16.9296 49.6 18.1999V20.4998C49.6 21.7701 48.5703 22.7998 47.3 22.7998H45C43.7297 22.7998 42.7 21.7701 42.7 20.4998V18.1999Z"
      fill="url(#paint2_linear_3769_168764)"
    />
    <path
      d="M38.1 34.3008C36.8297 34.3008 35.8 35.3305 35.8 36.6007V38.9007C35.8 40.171 36.8297 41.2007 38.1 41.2007H40.4V45.8006C40.4 47.0709 41.4297 48.1006 42.7 48.1006H45C46.2703 48.1006 47.3 47.0709 47.3 45.8006V41.2007C47.3 39.9304 46.2703 38.9007 45 38.9007H42.7V36.6007C42.7 35.3305 41.6703 34.3008 40.4 34.3008H38.1Z"
      fill="url(#paint3_linear_3769_168764)"
    />
    <path
      d="M38.1 48.0996C36.8297 48.0996 35.8 49.1293 35.8 50.3996V52.6995C35.8 53.9698 36.8297 54.9995 38.1 54.9995H40.4C41.6703 54.9995 42.7 53.9698 42.7 52.6995V50.3996C42.7 49.1293 41.6703 48.0996 40.4 48.0996H38.1Z"
      fill="url(#paint4_linear_3769_168764)"
    />
    <path
      d="M47.3 36.6007C47.3 35.3305 48.3297 34.3008 49.6 34.3008H51.9C53.1703 34.3008 54.2 35.3305 54.2 36.6007V38.9007C54.2 40.171 53.1703 41.2007 51.9 41.2007H49.6C48.3297 41.2007 47.3 40.171 47.3 38.9007V36.6007Z"
      fill="url(#paint5_linear_3769_168764)"
    />
    <path
      d="M49.6 48.0996C48.3297 48.0996 47.3 49.1293 47.3 50.3996V52.6995C47.3 53.9698 48.3297 54.9995 49.6 54.9995H51.9C53.1703 54.9995 54.2 53.9698 54.2 52.6995V50.3996C54.2 49.1293 53.1703 48.0996 51.9 48.0996H49.6Z"
      fill="url(#paint6_linear_3769_168764)"
    />
  </g>
</svg>
                    </button>
                    <div
                      class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                    >
                      <div
                        class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"
                      ></div>
                      <p>ورود با QR کد</p>
                    </div>
                  </div>

                  <div class="group/tooltip relative">
                    <button
                      type="button"
                      data-activate-form="one-time-pass"
                      data-toggleable="false"
                      class="login-with-app-btn w-[64px] h-[64px] rounded-lg cursor-pointer"
                      style="
                        background: linear-gradient(
                          199.93deg,
                          rgba(199, 210, 230, 0.5) 8.72%,
                          rgba(255, 255, 255, 0.5) 87.25%
                        );
                        box-shadow: 0px 5.4px 5.4px 0px #00000040;
                        backdrop-filter: blur(344.82763671875px);
                      "
                    >
                      <div
                        class="w-full flex gap-2 items-center justify-center"
                      >
                        <svg
  width="65"
  height="64"
  viewBox="0 0 65 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#otp-btn)">
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M39.5 36.999V47.3983H9.5V11.9999C9.5 11.2039 9.816 10.4419 10.378 9.87998C10.94 9.31599 11.7039 9 12.4999 9H36.4998C37.2959 9 38.0579 9.31599 38.6218 9.87998C39.1838 10.442 39.4998 11.2039 39.4998 11.9999V15.7997H18.0998C17.1978 15.7997 16.3338 16.1577 15.6958 16.7956C15.0578 17.4336 14.6999 18.2976 14.6999 19.1995V33.5988C14.6999 34.5007 15.0578 35.3667 15.6958 36.0026C16.3338 36.6406 17.1998 36.9986 18.1018 36.9966H20.2998C20.2998 36.9966 20.2858 39.4104 20.2998 39.9984C20.3038 40.1744 20.3278 40.3584 20.3798 40.5264C20.4278 40.6864 20.5078 40.8544 20.6038 40.9984C20.6918 41.1284 20.8198 41.2744 20.9538 41.3863C21.0718 41.4823 21.2478 41.5943 21.4118 41.6623C21.5438 41.7163 21.7558 41.7743 21.9298 41.7903C22.0978 41.8063 22.3038 41.7923 22.4518 41.7643C22.6218 41.7303 22.8198 41.6563 22.9438 41.5883C23.0978 41.5063 23.2618 41.3803 23.3738 41.2704C24.1578 40.5064 27.6458 36.9985 27.6458 36.9985L39.5 36.999ZM39.5 48.5983V51.9981C39.5 52.7941 39.184 53.5581 38.622 54.12C38.06 54.684 37.2961 55 36.5001 55L12.5002 54.998C11.7042 54.998 10.9421 54.682 10.3782 54.118C9.81616 53.556 9.5002 52.7921 9.5002 51.9961V48.5963H39.5002L39.5 48.5983ZM27.1481 35.799L22.5241 40.4227C22.3521 40.5947 22.0941 40.6467 21.8701 40.5527C21.6461 40.4607 21.5001 40.2407 21.5001 39.9988V35.795L18.1021 35.797C17.5182 35.797 16.9581 35.565 16.5461 35.153C16.1341 34.741 15.9021 34.1811 15.9021 33.5971L15.9001 19.1999C15.9001 18.6159 16.1321 18.0579 16.5441 17.6439C16.9561 17.2319 17.5161 17 18.1001 17H53.3C53.884 17 54.444 17.2319 54.856 17.6439C55.268 18.0559 55.5 18.6159 55.5 19.1999V33.5991C55.5 34.8131 54.516 35.799 53.3 35.799L27.1481 35.799ZM38.6178 26.3996L36.9658 27.4996C36.6898 27.6836 36.6158 28.0555 36.7998 28.3315C36.9838 28.6075 37.3558 28.6795 37.6318 28.4975L39.0998 27.5195V28.7975C39.0998 29.1295 39.3698 29.3975 39.6998 29.3975C40.0318 29.3975 40.2998 29.1295 40.2998 28.7975V27.5195L41.7678 28.4975C42.0438 28.6815 42.4158 28.6075 42.5998 28.3315C42.7838 28.0555 42.7078 27.6836 42.4338 27.4996L40.7838 26.3996L42.4338 25.2997C42.7098 25.1157 42.7838 24.7438 42.5998 24.4678C42.4158 24.1918 42.0438 24.1178 41.7678 24.3018L40.2998 25.2797V23.9998C40.2998 23.6698 40.0318 23.3998 39.6998 23.3998C39.3698 23.3998 39.0998 23.6698 39.0998 23.9998V25.2797L37.6318 24.3018C37.3558 24.1178 36.9838 24.1918 36.7998 24.4678C36.6158 24.7437 36.6918 25.1157 36.9658 25.2997L38.6178 26.3996ZM22.6179 26.3996L20.9659 27.4996C20.6899 27.6836 20.6159 28.0555 20.7999 28.3315C20.9839 28.6075 21.3559 28.6795 21.6319 28.4975L23.0999 27.5195V28.7975C23.0999 29.1295 23.3679 29.3975 23.6999 29.3975C24.0299 29.3975 24.2999 29.1295 24.2999 28.7975V27.5195L25.7679 28.4975C26.0439 28.6815 26.4159 28.6075 26.5999 28.3315C26.7839 28.0555 26.7099 27.6836 26.4339 27.4996L24.7819 26.3996L26.4339 25.2997C26.7099 25.1157 26.7839 24.7438 26.5999 24.4678C26.4159 24.1918 26.0439 24.1178 25.7679 24.3018L24.2999 25.2797V23.9998C24.2999 23.6698 24.0319 23.3998 23.6999 23.3998C23.3679 23.3998 23.0999 23.6698 23.0999 23.9998V25.2797L21.6319 24.3018C21.3559 24.1178 20.9839 24.1918 20.7999 24.4678C20.6159 24.7437 20.6919 25.1157 20.9659 25.2997L22.6179 26.3996ZM30.6179 26.3996L28.9659 27.4996C28.6899 27.6836 28.6159 28.0555 28.7998 28.3315C28.9838 28.6075 29.3558 28.6795 29.6318 28.4975L31.0998 27.5195V28.7975C31.0998 29.1295 31.3678 29.3975 31.6999 29.3975C32.0319 29.3975 32.2999 29.1295 32.2999 28.7975V27.5195L33.7679 28.4975C34.0439 28.6815 34.4159 28.6075 34.5999 28.3315C34.7839 28.0555 34.7099 27.6836 34.4339 27.4996L32.7819 26.3996L34.4339 25.2997C34.7099 25.1157 34.7839 24.7438 34.5999 24.4678C34.4159 24.1918 34.0439 24.1178 33.7679 24.3018L32.2999 25.2797V23.9998C32.2999 23.6698 32.0319 23.3998 31.6999 23.3998C31.3699 23.3998 31.0998 23.6698 31.0998 23.9998V25.2797L29.6318 24.3018C29.3558 24.1178 28.9839 24.1918 28.7998 24.4678C28.6158 24.7437 28.6899 25.1157 28.9659 25.2997L30.6179 26.3996ZM46.6178 26.3996L44.9678 27.4996C44.6918 27.6836 44.6178 28.0555 44.8018 28.3315C44.9858 28.6075 45.3578 28.6795 45.6338 28.4975L47.1018 27.5195V28.7975C47.1018 29.1295 47.3718 29.3975 47.7018 29.3975C48.0318 29.3975 48.3018 29.1295 48.3018 28.7975V27.5195L49.7698 28.4975C50.0458 28.6815 50.4178 28.6075 50.6018 28.3315C50.7858 28.0555 50.7098 27.6836 50.4358 27.4996L48.7838 26.3996L50.4358 25.2997C50.7118 25.1157 50.7858 24.7438 50.6018 24.4678C50.4178 24.1918 50.0458 24.1178 49.7698 24.3018L48.3018 25.2797V23.9998C48.3018 23.6698 48.0318 23.3998 47.7018 23.3998C47.3718 23.3998 47.1018 23.6698 47.1018 23.9998V25.2797L45.6338 24.3018C45.3578 24.1178 44.9858 24.1918 44.8018 24.4678C44.6178 24.7437 44.6938 25.1157 44.9678 25.2997L46.6178 26.3996Z"
      fill="url(#paint0_linear_3876_137583)"
    />
  </g>
</svg>
                      </div>
                    </button>
                    <div
                      class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                    >
                      <div
                        class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"
                      ></div>
                      <p>ورود با رمز یکبارمصرف</p>
                    </div>
                  </div>
                </div>
              </div>
            </form>

            <!--........................QR code login............................-->
            <form
              method="post"
              action="/login_otp"
              id="qr-form"
              class="mt-3 flex-col hidden"
              novalidate
            >
              <p
                class="text-[#5D5F61] mb-[80px] text-sm md:text-lg text-justify"
              >
                ورود به سامانه با اسکن QR کد
              </p>

              <div
                class="relative mx-auto w-fit bg-[#ECF0F7] p-4 rounded-2xl shadow-[0px_3px_7px_0px_#00000040]"
              >
                <img
                  class="data-[disabled=true]:opacity-10 data-[disabled=true]:blur-[1px] data-[loading=true]:opacity-10 data-[loading=true]:blur-[1px] peer"
                  src="../static/web/assets/images/new-my/QrCodeIcon.png"
                  id="qr-code-image"
                  alt="QR Placeholder"
                  data-disabled="false"
                  data-loading="true"
                />

                <div
                  class="hidden peer-data-[loading=true]:block absolute left-1/2 -translate-x-1/2 top-1/2 -translate-y-1/2 animate-spin rounded-full mx-auto before:absolute before:content-[''] before:top-0 before:left-0 before:w-full before:h-full before:rounded-full w-[50px] h-[50px] shadow-[#3F619D] shadow-[1px_4px_2px]"
                ></div>

                <span
                  id="retry-otp"
                  class="absolute left-1/2 -translate-x-1/2 top-1/2 -translate-y-1/2 flex-col items-center gap-8 mt-6 hidden peer-data-[disabled=true]:flex"
                >
                  <svg
  width="40"
  height="40"
  viewBox="0 0 40 40"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  id="qr-reset-btn"
  class="cursor-pointer"
>
  <path
    d="M25.2771 1.66602L26.2737 3.52897C26.9475 4.78846 27.2844 5.4182 27.0616 5.6879C26.8389 5.9576 26.1052 5.73979 24.6379 5.30415C23.1709 4.86861 21.6133 4.63427 19.9993 4.63427C11.2548 4.63427 4.16602 11.5135 4.16602 19.9994C4.16602 22.798 4.93705 25.4219 6.28423 27.6819M14.7216 38.3327L13.725 36.4697C13.0512 35.2102 12.7143 34.5805 12.937 34.3108C13.1598 34.0411 13.8935 34.2589 15.3608 34.6946C16.8278 35.1301 18.3854 35.3644 19.9993 35.3644C28.7439 35.3644 35.8327 28.4852 35.8327 19.9994C35.8327 17.2007 35.0616 14.5768 33.7145 12.3168"
    stroke="#304A79"
    stroke-width="3"
    stroke-linecap="round"
    stroke-linejoin="round"
  />
</svg>

                  <p
                    class="text-[18px] text-[#304A79]"
                  >دریافت مجدد</p>
                </span>
              </div>

              <input type="hidden" name="qr-code-key" id="qr-code-key" />

              <input
                type="hidden"
                name="_csrf"
                value="yAtpHmp7GsEKyyxyDAxnFXTIEAYmsSd2_4TL6W_0UWedrJKP8Gpafw4eePMn8k9KOyFTcRHwPT9C0kNbmrH4iFuRaFeuzarp"
                header="X-CSRF-TOKEN"
              />

              <input type="hidden" id="qr-password" name="password" />
              <input
                type="hidden"
                name="username"
                id="qr-username"
                class="username_final"
                value=""
              />

              <div
                id="qr-timer"
                class="my-8 text-[24px] ltr"
                style="margin-top: 1rem"
              >
                00:00
              </div>
              <div class="flex items-center justify-center gap-3 mb-[23px]">
                <svg
  width="188"
  height="2"
  viewBox="0 0 188 2"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path d="M0.5 1C74.1128 1 114.387 1 188 1" stroke="url(#stroke-right)" />
</svg>

                <p>یا</p>

                <svg
  width="188"
  height="2"
  viewBox="0 0 188 2"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path d="M0 1H187.5" stroke="url(#stroke-left)" />
</svg>
              </div>
              <div class="flex w-full gap-8 items-center justify-center">
                <div class="group/tooltip relative">
                  <button
                    type="button"
                    data-activate-form="one-time-pass"
                    class="login-with-app-btn w-[64px] h-[64px] rounded-lg cursor-pointer"
                    style="
                      background: linear-gradient(
                        199.93deg,
                        rgba(199, 210, 230, 0.5) 8.72%,
                        rgba(255, 255, 255, 0.5) 87.25%
                      );
                      box-shadow: 0px 5.4px 5.4px 0px #00000040;
                      backdrop-filter: blur(344.82763671875px);
                    "
                  >
                    <div class="w-full flex gap-2 items-center justify-center">
                      <svg
  width="65"
  height="64"
  viewBox="0 0 65 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#otp-btn)">
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M39.5 36.999V47.3983H9.5V11.9999C9.5 11.2039 9.816 10.4419 10.378 9.87998C10.94 9.31599 11.7039 9 12.4999 9H36.4998C37.2959 9 38.0579 9.31599 38.6218 9.87998C39.1838 10.442 39.4998 11.2039 39.4998 11.9999V15.7997H18.0998C17.1978 15.7997 16.3338 16.1577 15.6958 16.7956C15.0578 17.4336 14.6999 18.2976 14.6999 19.1995V33.5988C14.6999 34.5007 15.0578 35.3667 15.6958 36.0026C16.3338 36.6406 17.1998 36.9986 18.1018 36.9966H20.2998C20.2998 36.9966 20.2858 39.4104 20.2998 39.9984C20.3038 40.1744 20.3278 40.3584 20.3798 40.5264C20.4278 40.6864 20.5078 40.8544 20.6038 40.9984C20.6918 41.1284 20.8198 41.2744 20.9538 41.3863C21.0718 41.4823 21.2478 41.5943 21.4118 41.6623C21.5438 41.7163 21.7558 41.7743 21.9298 41.7903C22.0978 41.8063 22.3038 41.7923 22.4518 41.7643C22.6218 41.7303 22.8198 41.6563 22.9438 41.5883C23.0978 41.5063 23.2618 41.3803 23.3738 41.2704C24.1578 40.5064 27.6458 36.9985 27.6458 36.9985L39.5 36.999ZM39.5 48.5983V51.9981C39.5 52.7941 39.184 53.5581 38.622 54.12C38.06 54.684 37.2961 55 36.5001 55L12.5002 54.998C11.7042 54.998 10.9421 54.682 10.3782 54.118C9.81616 53.556 9.5002 52.7921 9.5002 51.9961V48.5963H39.5002L39.5 48.5983ZM27.1481 35.799L22.5241 40.4227C22.3521 40.5947 22.0941 40.6467 21.8701 40.5527C21.6461 40.4607 21.5001 40.2407 21.5001 39.9988V35.795L18.1021 35.797C17.5182 35.797 16.9581 35.565 16.5461 35.153C16.1341 34.741 15.9021 34.1811 15.9021 33.5971L15.9001 19.1999C15.9001 18.6159 16.1321 18.0579 16.5441 17.6439C16.9561 17.2319 17.5161 17 18.1001 17H53.3C53.884 17 54.444 17.2319 54.856 17.6439C55.268 18.0559 55.5 18.6159 55.5 19.1999V33.5991C55.5 34.8131 54.516 35.799 53.3 35.799L27.1481 35.799ZM38.6178 26.3996L36.9658 27.4996C36.6898 27.6836 36.6158 28.0555 36.7998 28.3315C36.9838 28.6075 37.3558 28.6795 37.6318 28.4975L39.0998 27.5195V28.7975C39.0998 29.1295 39.3698 29.3975 39.6998 29.3975C40.0318 29.3975 40.2998 29.1295 40.2998 28.7975V27.5195L41.7678 28.4975C42.0438 28.6815 42.4158 28.6075 42.5998 28.3315C42.7838 28.0555 42.7078 27.6836 42.4338 27.4996L40.7838 26.3996L42.4338 25.2997C42.7098 25.1157 42.7838 24.7438 42.5998 24.4678C42.4158 24.1918 42.0438 24.1178 41.7678 24.3018L40.2998 25.2797V23.9998C40.2998 23.6698 40.0318 23.3998 39.6998 23.3998C39.3698 23.3998 39.0998 23.6698 39.0998 23.9998V25.2797L37.6318 24.3018C37.3558 24.1178 36.9838 24.1918 36.7998 24.4678C36.6158 24.7437 36.6918 25.1157 36.9658 25.2997L38.6178 26.3996ZM22.6179 26.3996L20.9659 27.4996C20.6899 27.6836 20.6159 28.0555 20.7999 28.3315C20.9839 28.6075 21.3559 28.6795 21.6319 28.4975L23.0999 27.5195V28.7975C23.0999 29.1295 23.3679 29.3975 23.6999 29.3975C24.0299 29.3975 24.2999 29.1295 24.2999 28.7975V27.5195L25.7679 28.4975C26.0439 28.6815 26.4159 28.6075 26.5999 28.3315C26.7839 28.0555 26.7099 27.6836 26.4339 27.4996L24.7819 26.3996L26.4339 25.2997C26.7099 25.1157 26.7839 24.7438 26.5999 24.4678C26.4159 24.1918 26.0439 24.1178 25.7679 24.3018L24.2999 25.2797V23.9998C24.2999 23.6698 24.0319 23.3998 23.6999 23.3998C23.3679 23.3998 23.0999 23.6698 23.0999 23.9998V25.2797L21.6319 24.3018C21.3559 24.1178 20.9839 24.1918 20.7999 24.4678C20.6159 24.7437 20.6919 25.1157 20.9659 25.2997L22.6179 26.3996ZM30.6179 26.3996L28.9659 27.4996C28.6899 27.6836 28.6159 28.0555 28.7998 28.3315C28.9838 28.6075 29.3558 28.6795 29.6318 28.4975L31.0998 27.5195V28.7975C31.0998 29.1295 31.3678 29.3975 31.6999 29.3975C32.0319 29.3975 32.2999 29.1295 32.2999 28.7975V27.5195L33.7679 28.4975C34.0439 28.6815 34.4159 28.6075 34.5999 28.3315C34.7839 28.0555 34.7099 27.6836 34.4339 27.4996L32.7819 26.3996L34.4339 25.2997C34.7099 25.1157 34.7839 24.7438 34.5999 24.4678C34.4159 24.1918 34.0439 24.1178 33.7679 24.3018L32.2999 25.2797V23.9998C32.2999 23.6698 32.0319 23.3998 31.6999 23.3998C31.3699 23.3998 31.0998 23.6698 31.0998 23.9998V25.2797L29.6318 24.3018C29.3558 24.1178 28.9839 24.1918 28.7998 24.4678C28.6158 24.7437 28.6899 25.1157 28.9659 25.2997L30.6179 26.3996ZM46.6178 26.3996L44.9678 27.4996C44.6918 27.6836 44.6178 28.0555 44.8018 28.3315C44.9858 28.6075 45.3578 28.6795 45.6338 28.4975L47.1018 27.5195V28.7975C47.1018 29.1295 47.3718 29.3975 47.7018 29.3975C48.0318 29.3975 48.3018 29.1295 48.3018 28.7975V27.5195L49.7698 28.4975C50.0458 28.6815 50.4178 28.6075 50.6018 28.3315C50.7858 28.0555 50.7098 27.6836 50.4358 27.4996L48.7838 26.3996L50.4358 25.2997C50.7118 25.1157 50.7858 24.7438 50.6018 24.4678C50.4178 24.1918 50.0458 24.1178 49.7698 24.3018L48.3018 25.2797V23.9998C48.3018 23.6698 48.0318 23.3998 47.7018 23.3998C47.3718 23.3998 47.1018 23.6698 47.1018 23.9998V25.2797L45.6338 24.3018C45.3578 24.1178 44.9858 24.1918 44.8018 24.4678C44.6178 24.7437 44.6938 25.1157 44.9678 25.2997L46.6178 26.3996Z"
      fill="url(#paint0_linear_3876_137583)"
    />
  </g>
</svg>
                    </div>
                  </button>
                  <div
                    class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                  >
                    <p class="">ورود با رمز یکبارمصرف</p>
                  </div>
                </div>
                <div class="group/tooltip relative">
                  <button
                    type="button"
                    data-activate-form="application"
                    class="login-with-app-btn w-[64px] h-[64px] rounded-lg cursor-pointer"
                    style="
                      background: linear-gradient(
                        199.93deg,
                        rgba(199, 210, 230, 0.5) 8.72%,
                        rgba(255, 255, 255, 0.5) 87.25%
                      );
                      box-shadow: 0px 5.4px 5.4px 0px #00000040;
                      backdrop-filter: blur(344.82763671875px);
                    "
                  >
                    <div class="w-full flex gap-2 items-center justify-center">
                      <svg
  width="53"
  height="64"
  viewBox="0 0 53 64"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <g filter="url(#app-btn)">
    <path
      d="M33.7454 9H19.2547C15.5328 9 12.5 11.9647 12.5 15.6128V48.3872C12.5 52.0309 15.5283 55 19.2547 55H33.743C37.4648 55 40.4976 52.0353 40.4976 48.3872L40.5 15.6128C40.5 11.9714 37.467 9 33.7454 9ZM26.5037 49.645C25.2092 49.645 24.157 48.6149 24.157 47.3477C24.157 46.0804 25.2092 45.0503 26.5037 45.0503C27.7981 45.0503 28.8503 46.0804 28.8503 47.3477C28.8456 48.6149 27.7981 49.645 26.5037 49.645ZM29.4393 15.864H23.5585C22.8887 15.864 22.3424 15.3292 22.3424 14.6734C22.3424 14.0177 22.8887 13.4829 23.5585 13.4829H29.4393C30.1091 13.4829 30.6554 14.0177 30.6554 14.6734C30.6554 15.3269 30.1163 15.864 29.4393 15.864Z"
      fill="url(#paint0_linear_3769_168762)"
    />
  </g>
</svg>
                    </div>
                  </button>
                  <div
                    class="absolute bottom-full left-1/2 transform -translate-x-1/2 flex items-center gap-2 w-max -mb-4 px-2 py-1 text-sm text-white bg-[#bababa] rounded shadow-lg opacity-0 group-hover/tooltip:opacity-100"
                  >
                    <div class="w-[2px] h-4.5 bg-[#DADBDC] rounded-full"></div>
                    <p class="">ورود با اپلیکیشن</p>
                  </div>
                </div>
              </div>
            </form>

            <div
              class="w-full -bottom-1/5 text-center text-xs sm:text-sm font-medium sm:font-normal text-[#AAACAE] my-4 mt-[66px] pb-4"
            >
              کلیه حقوق این درگاه‌ها متعلق به دولت جمهوری اسلامی ایران است.
            </div>
          </div>
        </section>
      </div>
    </div>
    <div
      data-type="error"
      id="toast-wrapper"
      class="fixed md:gap-6 items-center opacity-0 transition-opacity duration-500 z-10 left-1/2 -translate-x-1/2 top-4 border-2 border-white w-full max-w-[620px] shadow-[-2px_-2px_16px_0px_rgba(0,0,0,0.08)] rounded-toast data-[type=error]:bg-[linear-gradient(215deg,#FFE3E2_2.43%,#FFFFFF_34.3%)] data-[type=success]:bg-[linear-gradient(215deg,#98EDB3_2.43%,#FFFFFF_34.3%)] data-[type=warning]:bg-[linear-gradient(215deg,#FFEDC2_2.43%,#FFFFFF_34.3%)] rounded-2xl"
    ></div>

    <!----------------------------------------------------  Scripts-------------------------------------------------------->
    <script>
      const otp = "false";
      const mob = "";
      const lang = "fa";

      const csrf_parameterName = "_csrf";
      const csrf_token = "yAtpHmp7GsEKyyxyDAxnFXTIEAYmsSd2_4TL6W_0UWedrJKP8Gpafw4eePMn8k9KOyFTcRHwPT9C0kNbmrH4iFuRaFeuzarp";
      const csrf_headerName = "_csrf.headerName";

      const languageProperties = {
        "messages.internalerrors": "خطای داخلی",
        "messages.tryagain": "لطفا مجدد تلاش نمایید",
        loading: "در حال بارگذاری ...",
        "messages.mobileIsInvalid": "شماره موبایل وارد شده باید 11 رقمی باشد و با 09 شروع شود",
        "messages.checkinput": "لطفا اطلاعات وارد شده را بررسی نمایید",
        "messages.captchaRequired": "لطفا کد فعال سازی را بصورت کامل وارد نمایید",
        "messages.captchaFetchError": "خطا در ریافت تصویر کد امنیتی",
        "messages.voiceCaptchaFetchError":
          "خطا در ریافت صوت کد امنیتی",
        "buttons.resendOtp": "ارسال مجدد کد فعال سازی",
        "messages.postalCodeIsInvalid": "کد پستی نا صحیح است",
        "messages.birthdateRequired": "تاریخ تولد را وارد نمایید",
        "messages.emailSent": "ارسال ایمیل با موفقیت انجام شد",
        "messages.success": "موفق",
        "messages.error": "خطا",
        "messages.foreginerDescription": "کاربر گرامی جهت احراز هویت تاریخ تولد و کد اتباع خود راوارد نمایید",
        "messages.nationalCodeValidate": "کد ملی باید 10 رقم باشد",
        "messages.postalCodeExistDescription":
          "کاربر گرامی جهت احراز هویت تاریخ تولد و کد ملی خود راوارد نمایید",
      };
    </script>
    <script src="../static/web/js/widget.module.min.js"></script>
    <script src="../static/web/assets/js/aaa.js?v=1"></script>
    <!--    <script src="../static/web/assets/js/app.js?v=11"></script>-->
    <script src="../static/web/assets/js/common.js?v=0.1"></script>
    <!--    <script-->
    <!--      src="../static/web/assets/js/accessible/vaslAccessible.js"-->
    <!--      async-->
    <!--    ></script>-->
    <!--<script src="../static/web/assets/js/a11y.js"></script>-->

    <script>
      const stateLang = "state1";
      const checks = {
        stateHasEn: stateLang?.includes("lang_en"),
        stateHasFa: stateLang?.includes("lang_fa"),
        searchIsEn: window.location.search.toLowerCase() === "?lang=en",
        searchIsFa: window.location.search.toLowerCase() === "?lang=fa",
        hasSearch: Boolean(window.location.search),
      };

      if (checks.stateHasEn && !checks.hasSearch && !checks.searchIsEn) {
        window.location.replace("/login?lang=en");
      } else if (checks.stateHasFa && !checks.hasSearch && !checks.searchIsFa) {
        window.location.replace("/login?lang=fa");
      }
    </script>
    <script>
      window.onFrcDone = function () {
        document.getElementById("simple-captcha").style.display = "block";
        const captchaDiv = document.querySelector(".frc-captcha");
        const holder = document.getElementById("frc-captcha-holder");
        if (captchaDiv && holder) {
          holder.appendChild(captchaDiv);
        }
      };
    </script>
    <script type="module">
      import Core from "../static/web/assets/js/core.js";
      import Utils from "../static/web/assets/js/utils.js";
      import Handlers from "../static/web/assets/js/handlers.js";

      document.addEventListener("DOMContentLoaded", () => {
        const utils = new Utils();
        const handlers = new Handlers({
          utils: utils,
          toastWrapperSelector: "#toast-wrapper",
        });

        handlers.sendOtpBaleHandler().sendOtpEmailHandler();

        utils.otpConfigUpdate();

        const core = new Core({
          toastWrapperSelector: "#toast-wrapper",
          utils,
          handlers,
        });

        core.loginSelectorsUpdate({}).loginMobile();
      });
    </script>
  </body>
</html>

```

---

[← Previous Step](step_02.md) | [Summary](summary.md)
