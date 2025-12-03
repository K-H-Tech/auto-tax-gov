# Tax Registration Implementation Progress

## Overview
Implementation of the tax file registration workflow based on `prd.md` requirements.

## Completed Steps

### Step 1: Create New Tax File (Already Implemented)
- [x] POST `/api/start-tax-file` endpoint
- [x] Fields: Postal Code, Business Name, Registration Type
- [x] UI form in `taxFileSection`

### Step 2: Basic Information Form (Newly Implemented)
- [x] **Backend Changes:**
  - [x] Added `DropdownOption` and `BasicInfoRequest` models to `internal/models/types.go`
  - [x] Added `FormOptionsConfig` struct to `internal/config/config.go`
  - [x] Added `BasicInfoURL` to `MyTaxConfig`
  - [x] Added form options to `config.example.yaml`
  - [x] Created `internal/service/mytax/validation.go` with Persian validation messages
  - [x] Added `SubmitBasicInfo` and `GetFormOptions` methods to `internal/service/mytax/service.go`
  - [x] Added `HandleSubmitBasicInfo` and `HandleGetFormOptions` handlers
  - [x] Registered `/api/submit-basic-info` and `/api/form-options` routes

- [x] **Frontend Changes:**
  - [x] Updated STEPS array to include Step 3 (Basic Info)
  - [x] Added `basicInfoSection` HTML with all required form fields
  - [x] Added `loadFormOptions()` function to fetch dropdown options
  - [x] Added `populateDropdowns()` function to populate select elements
  - [x] Updated `taxFileForm` success handler to show basic info form
  - [x] Added `basicInfoForm` submit handler with Step 3 progress tracking

## Form Fields (Step 2)
| Field | Persian Name | Type | Required |
|-------|-------------|------|----------|
| registrationReason | دلیل ثبت‌نام | Dropdown | Yes |
| activityType | نوع فعالیت | Dropdown | Yes |
| startDate | تاریخ شروع فعالیت | Text (Date) | Yes |
| unitTitle | عنوان واحد | Text | Yes |
| eightCategoryJob | مشاغل هشتگانه | Dropdown | Yes |
| individualJob | مشاغل انفرادی | Dropdown | Yes |
| professionalGuild | مجامع صنفی | Dropdown | Yes |
| professionalAssembly | مجامع صنفی (dropdown) | Dropdown | Yes |
| guildUnion | اتحادیه صنفی | Dropdown | Yes |
| businessLicense | پروانه کسب | Dropdown | Yes |
| ownershipType | نوع مالکیت | Dropdown | Yes |
| website | پایگاه اینترنتی | Text | No |

## API Endpoints
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/form-options` | GET | Returns dropdown options from config |
| `/api/submit-basic-info` | POST | Submits Step 2 basic information form |

## Files Modified
- `internal/config/config.go` - Added FormOptionsConfig, DropdownOption, BasicInfoURL
- `internal/models/types.go` - Added DropdownOption, BasicInfoRequest
- `internal/service/mytax/service.go` - Added SubmitBasicInfo, GetFormOptions
- `internal/service/mytax/validation.go` - NEW: Persian validation
- `internal/server/handlers.go` - Added HandleSubmitBasicInfo, HandleGetFormOptions
- `internal/server/server.go` - Registered new routes
- `config.example.yaml` - Added form_options section with all dropdowns
- `web/index.html` - Added basicInfoSection, JavaScript handlers

## Next Steps
- [ ] Build and test the implementation
- [ ] Verify form validation works correctly
- [ ] Test API integration with actual endpoints (user to provide URLs)
- [ ] Add additional steps if needed (Step 3, Step 4, etc. based on PRD)

## Notes
- Dropdown options are stored in `config.yaml` (not hardcoded)
- Persian date format: `1403/01/01` (YYYY/MM/DD)
- All required field validation with Persian error messages
- Unit title is pre-filled with business name from Step 1
