后端使用golang + gin
使用golang渲染html页面
前端使用Vue3 + ElementPlus
不需要go build校验  

web\views存放html模板 
每个page使用独立的html模板

前台页面除了登录和注册页，其他页面都使用layout/layout.html模板，参考path /page/circle/detail，handler注意调用RenderIndexPage，注意不要出现index_layout_extra_scripts部分渲染两次的问题，使用layout中的.container样式，不要自己定义.container样式  
后台页面都使用admin/_layout.html模板  

---

# bishe - 知识分享平台

Go + Gin web application, server-side rendered HTML templates.

## Build & Run

```bash
go run main.go    # Start on :8080, watches web/views/ at init
```

No `go build` needed for validation (Go template runtime errors don't surface at compile time).

## Architecture

```
controller/     # HTTP handlers (page + API)
service/        # Business logic, init (DB, Kafka, logger)
dao/mysql/      # Database access
model/          # Struct definitions
middleware/     # Auth middleware (user + admin, page + API variants)
config/         # Configuration
cons/           # Constants
```

Routes in `main.go`: two main route groups — `/page` (page handlers, uses `PageUserLogin` middleware) and `/api` (API handlers, uses `ApiUserLogin` middleware). Admin routes similarly use `/page` + `/api` with admin middleware.

## Template System

Templates live in `web/views/`. At init, every `.html` file is parsed with its relative path as the template name (e.g., `feedback/index.html`).

### Frontend pages (user-facing)

Page templates using layout follow this structure:

```
{{template "index_layout_head" .}}
{{template "index_layout_body_start" .}}

{{block "index_layout_extra_styles" .}}<style>/* page CSS */</style>{{end}}

<!-- page HTML content -->

{{template "index_layout_body_end" .}}

{{define "index_layout_extra_scripts"}}<script>/* page JS */</script>{{end}}
```

**Do NOT:**
- Define `.container` in extra styles — use the layout's `.container`
- Define `escapeHtml()` in extra scripts — layout already provides it
- Include axios CDN — layout already includes it
- Write standalone `<html>`, `<head>`, `<body>` — layout provides them

**Exceptions** (standalone, no layout): login page, register page.

### Admin pages (backend)

调用 `RenderAdminPage(c, "adminxxx/page.html", data)`，same clone+parse pattern as index layout.

### Layout templates

| Template name | Purpose |
|---|---|
| `index_layout_head` | `<head>` with shared CSS, header nav, axios |
| `index_layout_body_start` | `<body>`, shared header bar, opens `.container` |
| `index_layout_body_end` | closes `.container`, shared search JS, `escapeHtml()`, `</body></html>` |

### Named blocks for page injection

| Block | Purpose |
|---|---|
| `index_layout_extra_head` | Extra `<head>` content (rarely used) |
| `index_layout_extra_styles` | Page-specific `<style>` |
| `index_layout_extra_scripts` | Page-specific `<script>` (use `{{define}}`, not `{{block}}`) |

## Key Conventions

- Backend: Go + Gin, server-side HTML rendering
- Frontend: Vue 3 + Element Plus (loaded via CDN in pages that need it)
- Each page uses its own independent HTML template
- Use `RenderIndexPage` / `RenderAdminPage` for layout-based pages, never `c.HTML()` for those
- Avoid rendering `index_layout_extra_scripts` twice (use `{{define}}`, not `{{block}}`)
- DAO → Service → Controller: thin pass-through from service to DAO for most operations
