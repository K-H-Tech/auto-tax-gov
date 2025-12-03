package taxregister

import (
	"fmt"
	"time"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
)

// ========================================
// DROPDOWN OPTIONS FROM register.tax.gov.ir/Pages/Preaction/PublicData
// Extracted from curl.md (captured HTML responses)
// ========================================

// FormDropdownOption represents a single dropdown option with value and label.
// Uses a different name than DropdownOption to avoid conflict with models.go.
type FormDropdownOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// PublicDataFormDefaults contains all dropdown options for the PublicData form.
var PublicDataFormDefaults = struct {
	// DDLNewRegistrationCause - دلیل ثبت نام
	RegistrationReasons []DropdownOption
	// DDLIsTejari - نوع فعالیت (تجاری/غیر تجاری)
	ActivityTypes []DropdownOption
	// DDLGroupOneTypes - دسته بندی 8 گانه
	EightCategoryJobs []DropdownOption
	// DDLPDNewLegalGroup - مجامع حرفه ای/تشکل های صنفی
	ProfessionalAssemblies []DropdownOption
	// DDLHasJobLicence - پروانه کسب
	BusinessLicenses []DropdownOption
	// DDLPDOwnership - نوع مالکیت
	OwnershipTypes []DropdownOption
	// DDLFinantialDayStart - روز آغاز سال مالی
	FinancialDays []DropdownOption
	// DDLFinantialMonthStart - ماه آغاز سال مالی
	FinancialMonths []DropdownOption
	// DDLFinantialSoratMali - صورت های مالی حسابرسی
	FinancialAuditOptions []DropdownOption
	// DDLFinantialGozareshMali - گزارشات مالی حسابرسی
	ReportsAuditOptions []DropdownOption
}{
	// DDLNewRegistrationCause
	RegistrationReasons: []DropdownOption{
		{Value: "-1", Label: "انتخاب کنید..."},
		{Value: "0", Label: "ایجاد یک کسب و کار جدید"},
		{Value: "1", Label: "تغییر در شماره پستی"},
	},
	// DDLIsTejari
	ActivityTypes: []DropdownOption{
		{Value: "-1", Label: "انتخاب کنید"},
		{Value: "1", Label: "تجاری (همچنین فعالیتهای معاف از مالیات)"},
		{Value: "0", Label: "غیر تجاری"},
	},
	// DDLGroupOneTypes
	EightCategoryJobs: []DropdownOption{
		{Value: "0", Label: "نامشخص"},
		{Value: "1", Label: "(دارندگان کارت بازرگانی) واردکنندگان و صادرکنندگان"},
		{Value: "2", Label: "صاحبان کارخـانه ها و واحـدهای تـولیدی و بهره برداری معادن دارای مجوز تاسیس و پروانه بهره برداری از وزارت خانه ذیربط"},
		{Value: "3", Label: "صاحبان هتل های سه ستـاره و بالاتر"},
		{Value: "4", Label: "صاحبان بیمارستان‌ها، زایشگاه‌ها، درمانگاه‌ها، کلینیک‌های تخصصی"},
		{Value: "5", Label: "صـاحبان مشاغل صرافی"},
		{Value: "6", Label: "فروشگاه های زنجیره ای دارای مجوز فعالیت از وزارتخانه ذی ربط"},
		{Value: "7", Label: "صاحبان مؤسسات حسابرسی، حسابداری و دفترداری،‌خدمات مالی و ارائه‌دهندگان خدمات مدیریتی، مشاوره‌ای، انفورماتیک و طراحی سیستم"},
		{Value: "8", Label: "صاحبان مؤسسات حمل و نقل موتوری، زمینی، دریایی و هوایی اعم از مسافری و یا باربری"},
		{Value: "100", Label: "سایر (گروه دوم و سوم)"},
	},
	// DDLPDNewLegalGroup
	ProfessionalAssemblies: []DropdownOption{
		{Value: "0", Label: "انتخاب کنید..."},
		{Value: "1", Label: "اتاق بازرگانی،صنایع و معادن و کشاورزی ایران"},
		{Value: "2", Label: "اتاق تعاون ایران"},
		{Value: "3", Label: "جامعه حسابداران رسمی ایران"},
		{Value: "4", Label: "جامعه مشاوران رسمی مالیاتی ایران"},
		{Value: "5", Label: "مجامع حرفه ای"},
		{Value: "6", Label: "اتاق اصناف ایران"},
		{Value: "7", Label: "شورای اسلامی شهر"},
		{Value: "8", Label: "کانون وکلا"},
		{Value: "9", Label: "مرکز وکلا"},
		{Value: "10", Label: "کارشناسان رسمی"},
		{Value: "20", Label: "سایر نهادهای قانونی"},
	},
	// DDLHasJobLicence
	BusinessLicenses: []DropdownOption{
		{Value: "1", Label: "پروانه و جواز کسب دارم"},
		{Value: "0", Label: "فعلا فاقد جواز کسب می باشم و در درست اقدام قرار دارد"},
	},
	// DDLPDOwnership
	OwnershipTypes: []DropdownOption{
		{Value: "0", Label: "نامشخص"},
		{Value: "1", Label: "ملکی"},
		{Value: "2", Label: "سرقفلی"},
		{Value: "3", Label: "اجاری"},
		{Value: "4", Label: "وقفی"},
	},
	// DDLFinantialDayStart (1-31)
	FinancialDays: func() []DropdownOption {
		days := make([]DropdownOption, 31)
		for i := 1; i <= 31; i++ {
			days[i-1] = DropdownOption{Value: fmt.Sprintf("%d", i), Label: fmt.Sprintf("%d", i)}
		}
		return days
	}(),
	// DDLFinantialMonthStart
	FinancialMonths: []DropdownOption{
		{Value: "1", Label: "فروردین"},
		{Value: "2", Label: "اردیبهشت"},
		{Value: "3", Label: "خرداد"},
		{Value: "4", Label: "تیر"},
		{Value: "5", Label: "مرداد"},
		{Value: "6", Label: "شهریور"},
		{Value: "7", Label: "مهر"},
		{Value: "8", Label: "آبان"},
		{Value: "9", Label: "اذر"},
		{Value: "10", Label: "دی"},
		{Value: "11", Label: "بهمن"},
		{Value: "12", Label: "اسفند"},
	},
	// DDLFinantialSoratMali
	FinancialAuditOptions: []DropdownOption{
		{Value: "0", Label: "نامشخص"},
		{Value: "1", Label: "بله"},
		{Value: "2", Label: "خیر"},
	},
	// DDLFinantialGozareshMali
	ReportsAuditOptions: []DropdownOption{
		{Value: "0", Label: "نامشخص"},
		{Value: "1", Label: "بله"},
		{Value: "2", Label: "خیر"},
	},
}

// HappyPathDefaults contains the exact values needed for a successful registration.
// These values are tested against the live portal via Playwright manual testing.
var HappyPathDefaults = struct {
	// PublicData form field values (exact values to submit)
	RegistrationReason    string // DDLNewRegistrationCause
	IsTejari              string // DDLIsTejari (Commercial)
	GroupOneType          string // DDLGroupOneTypes (8-category job)
	EnferadiTypes         string // DDLEnferadiTypes (مشاغل انفرادی)
	LegalType             string // DDLPDLegalType (Professional Guild)
	NewLegalGroup         string // DDLPDNewLegalGroup (Professional Assembly)
	NewLegalType          string // DDLPDNewLegalType (Cascade: New Guild Union)
	HasJobLicense         string // DDLHasJobLicence
	Ownership             string // DDLPDOwnership
	FinancialDayStart     string // DDLFinantialDayStart
	FinancialMonthStart   string // DDLFinantialMonthStart
	FinancialSoratMali    string // DDLFinantialSoratMali
	FinancialGozareshMali string // DDLFinantialGozareshMali
}{
	RegistrationReason:    "0",    // ایجاد یک کسب و کار جدید
	IsTejari:              "1",    // تجاری (همچنین فعالیتهای معاف از مالیات)
	GroupOneType:          "100",  // سایر (گروه دوم و سوم)
	EnferadiTypes:         "1000", // سایر (غیر انفرادی)
	LegalType:             "4032", // مشاوره و خدمات
	NewLegalGroup:         "6",    // اتاق اصناف ایران
	NewLegalType:          "9090", // اتحادیه کشوری - سایر
	HasJobLicense:         "0",    // فعلا فاقد جواز کسب
	Ownership:             "3",    // اجاری
	FinancialDayStart:     "1",    // First day
	FinancialMonthStart:   "1",    // فروردین
	FinancialSoratMali:    "2",    // خیر
	FinancialGozareshMali: "2",    // خیر
}

// ProfessionalGuilds contains all options for DDLPDLegalType (صنف/شغل حرفه‌ای).
// This is a large list from the portal (200+ options).
var ProfessionalGuilds = []DropdownOption{
	{Value: "0", Label: "انتخاب کنید..."},
	{Value: "4124", Label: "کسب و کاراینترنتی"},
	{Value: "2001", Label: "آبکاران"},
	{Value: "2003", Label: "آرایشگران زنانه"},
	{Value: "2002", Label: "آرایشگران مردانه"},
	{Value: "2004", Label: "آسیابداران"},
	{Value: "1066", Label: "آموزشگاهها"},
	{Value: "2005", Label: "آهنسازان"},
	{Value: "2007", Label: "آهنگران اتومبیل"},
	{Value: "2006", Label: "آهنکاران ساختمان"},
	{Value: "1002", Label: "ابزار فروشان"},
	{Value: "4123", Label: "اتحادیه کسب و کار فضای مجازی"},
	{Value: "4118", Label: "اتحادیه کشوری سوختهای جایگزین و خدمات وابسته"},
	{Value: "2008", Label: "اتوسرویس"},
	{Value: "4042", Label: "اسباب بازی"},
	{Value: "4122", Label: "استفاده کننده از پایانه پرداخت (POS)"},
	{Value: "1086", Label: "اشیا قدیمی و صنایع دستی"},
	{Value: "1003", Label: "اغذیه فروشان"},
	{Value: "4095", Label: "الکتروموتور و سیم پیچ"},
	{Value: "2065", Label: "الکتریک و سیم کشی"},
	{Value: "4039", Label: "امور سینمائی و تاتر"},
	{Value: "4069", Label: "انبارداران کالاهای تجاری"},
	{Value: "2009", Label: "اوراق کنندگان اتومبیل"},
	{Value: "1007", Label: "بار فروشان"},
	{Value: "1004", Label: "بارکش شهری"},
	{Value: "4055", Label: "بازرگانان"},
	{Value: "2010", Label: "باطریساز و باطری فروش"},
	{Value: "4041", Label: "بافندگی(بجز فرش)"},
	{Value: "4119", Label: "برنج فروشان"},
	{Value: "2011", Label: "بستنی و فالوده و آبمیوه فروش"},
	{Value: "1005", Label: "بنکدار مواد غذایی"},
	{Value: "1085", Label: "بنکداران چای تهران"},
	{Value: "1084", Label: "بوفه داران سینما"},
	{Value: "4062", Label: "بیمه"},
	{Value: "1008", Label: "پرنده و ماهی"},
	{Value: "4002", Label: "پزشک عمومی"},
	{Value: "4001", Label: "پزشک متخصص"},
	{Value: "1067", Label: "پو ششکاران داخل ساختمان"},
	{Value: "4056", Label: "پیچ و مهره فروشان"},
	{Value: "2012", Label: "پیراهن دوزان و پیراهن فروشان"},
	{Value: "4061", Label: "پیمانکاری"},
	{Value: "2018", Label: "تابلوسازان نئون و پلاستیک"},
	{Value: "4068", Label: "تابلوسازان و موسسات برق صنعتی"},
	{Value: "4098", Label: "تاسیسات مکانیکی ساختمان"},
	{Value: "1011", Label: "تالار های پذیرایی"},
	{Value: "1010", Label: "تاکسی بار"},
	{Value: "4093", Label: "تراشکاران"},
	{Value: "2013", Label: "تزئینات داخل ساختمان"},
	{Value: "4016", Label: "تزئینات سفره عقد و لباس عروس"},
	{Value: "2064", Label: "تشکدوزان اتومبیل"},
	{Value: "2014", Label: "تعمیر کاران چرخ خیاطی"},
	{Value: "2017", Label: "تعمیر کاران دوچرخه و موتور"},
	{Value: "2015", Label: "تعمیر کاران قفل و کلید"},
	{Value: "2075", Label: "تعمیرکاران لوازم صوتی و تصویری"},
	{Value: "2076", Label: "تعویض روغن و آپاراتی"},
	{Value: "4027", Label: "تلفن و موبایل"},
	{Value: "1069", Label: "توزیع کنندگان یخ"},
	{Value: "2016", Label: "تولید کنندگان لوازم برقی"},
	{Value: "4092", Label: "تولیدو تعمیر کنندگان لوازم الکترومکانیک"},
	{Value: "4107", Label: "تولیدکنندگان و صادرکنندگان محصولات و مواد اولیه بیسکویت، شیرینی و شکلات"},
	{Value: "4110", Label: "تولیدکنندگان و فروشندگان مصنوعات سیمانی و لوازم فلزی ساختمانی تهران"},
	{Value: "1012", Label: "جراید"},
	{Value: "2019", Label: "جواهر _ طلا و نقره"},
	{Value: "4094", Label: "جوشکاران و درب و پنجره سازان آهنی"},
	{Value: "2020", Label: "چاپخانه داران"},
	{Value: "2061", Label: "چادر دوزان و برزنت فروشان"},
	{Value: "1070", Label: "چرم فروشان"},
	{Value: "1013", Label: "چرم و لوازم کفش"},
	{Value: "1015", Label: "چلو کباب و چلو خورش"},
	{Value: "1014", Label: "چوب و تخته و فیبر"},
	{Value: "4024", Label: "حسابدار و حسابرسان"},
	{Value: "1073", Label: "حفر چاه های عمیق و نیمه عمیق"},
	{Value: "2021", Label: "حلوا ساز و عصار"},
	{Value: "1072", Label: "حمل و نقل استان مرکز"},
	{Value: "4053", Label: "خبرنگاران"},
	{Value: "4034", Label: "خدمات درمانی"},
	{Value: "4091", Label: "خدمات فنی"},
	{Value: "4033", Label: "خدمات مهندسی _ نظام مهندسی"},
	{Value: "4057", Label: "خدمات و فروش رایانه"},
	{Value: "1018", Label: "خرازی"},
	{Value: "4108", Label: "خرده فروشی سکه طلا"},
	{Value: "1017", Label: "خشکبار و آجیل فروش"},
	{Value: "2022", Label: "خشکشوئی و لباسشوئی"},
	{Value: "1016", Label: "خوار بار فروش"},
	{Value: "2023", Label: "خیاطان"},
	{Value: "4015", Label: "داروخانه و پخش دارو"},
	{Value: "4023", Label: "دامداری"},
	{Value: "1074", Label: "دباغ و پوستی"},
	{Value: "4134", Label: "درگاه ملی مجوزها"},
	{Value: "2024", Label: "درود گران و مبلسازان"},
	{Value: "4101", Label: "دستگاه های مخابراتی و ارتباطی و لوازم جانبی"},
	{Value: "4103", Label: "دستگاه های مخابراتی، ارتباطی و لوازم جانبی تهران"},
	{Value: "1021", Label: "دستگاههای صوتی و تصویری"},
	{Value: "4196", Label: "دفاتر پیشخوان"},
	{Value: "4064", Label: "دفاتراسناد رسمی"},
	{Value: "4054", Label: "دفتر داران"},
	{Value: "1020", Label: "دل و جگر و قلوه"},
	{Value: "4036", Label: "دندانسازی و دندانساز تجربی"},
	{Value: "1019", Label: "دوچرخه و موتور سیکلت"},
	{Value: "1022", Label: "ذغالفروشان"},
	{Value: "2025", Label: "ذوب فلزات"},
	{Value: "1075", Label: "رادیو و لوازم یدکی (فروشندگان)"},
	{Value: "4026", Label: "رایانه و کافی نت و بازیهای رایانه"},
	{Value: "1023", Label: "رنگ فروشان"},
	{Value: "2066", Label: "رویفروشان و مسگران"},
	{Value: "4130", Label: "ساخت و ساز ساختمان"},
	{Value: "2062", Label: "سازندگان و فروشندگان ساعت"},
	{Value: "2036", Label: "سازندگان و فروشندگان عینک"},
	{Value: "2063", Label: "سازنده درب و پنجره آلومینیوم"},
	{Value: "1076", Label: "سالنهای پذیرائی و ظروف کرایه"},
	{Value: "2026", Label: "سراجان و تشک دوزان"},
	{Value: "1027", Label: "سرایداران"},
	{Value: "4100", Label: "سفال و سرامیک"},
	{Value: "2029", Label: "سماور ساز و چراغساز"},
	{Value: "1026", Label: "سمساران"},
	{Value: "1025", Label: "سموم دفع آفات"},
	{Value: "4003", Label: "سنج و طبل"},
	{Value: "2027", Label: "سنگبر و سنگتراش"},
	{Value: "1028", Label: "سوپر مارکتها"},
	{Value: "1024", Label: "سیسمونی"},
	{Value: "2028", Label: "سیمانکاران و موزائیک سازان"},
	{Value: "4106", Label: "شالیکوبان"},
	{Value: "4049", Label: "شبرنگ و کفپوش و دیوارکوب"},
	{Value: "4018", Label: "شمع و پارافین"},
	{Value: "1029", Label: "شوفاژ و تهویه مطبوع"},
	{Value: "1077", Label: "شیشه و آئینه"},
	{Value: "2030", Label: "صابونساز و صابون فروش"},
	{Value: "4096", Label: "صافکاران اتومبیل"},
	{Value: "2031", Label: "صباغان و گلزن"},
	{Value: "2032", Label: "صحافان و آلبوم سازان"},
	{Value: "4011", Label: "صرافان"},
	{Value: "2033", Label: "صنایع آلومینیوم"},
	{Value: "4067", Label: "صنایع برودتی،تهویه مطبوع و شوینده"},
	{Value: "4008", Label: "صنایع فلزی"},
	{Value: "4044", Label: "صنایع فلزی"},
	{Value: "4102", Label: "صنعتگران"},
	{Value: "1006", Label: "صنف پارچه"},
	{Value: "4050", Label: "ضایعات فلزی و غیر فلزی"},
	{Value: "2034", Label: "ظروف آلومینیوم"},
	{Value: "1030", Label: "ظروف بلور چینی و لوستر"},
	{Value: "1087", Label: "ظروف کرایه"},
	{Value: "1078", Label: "عتیقه فروشان و مصنوعات دستی"},
	{Value: "1031", Label: "عطار و سقط فروش"},
	{Value: "2035", Label: "عکاسان و فیلمبرداران"},
	{Value: "1032", Label: "غلات و حبوبات"},
	{Value: "4065", Label: "فاقد اتحادیه"},
	{Value: "1033", Label: "فتوکپی و اوزالید"},
	{Value: "2037", Label: "فخاران"},
	{Value: "4017", Label: "فرش دستباف و گلیم"},
	{Value: "1034", Label: "فرش فروشان"},
	{Value: "1001", Label: "فروشندگان آهن و فولاد"},
	{Value: "1036", Label: "فروشندگان گل"},
	{Value: "4099", Label: "فروشندگان لاستیک"},
	{Value: "1088", Label: "فروشندگان لوازم پزشکی و طبی"},
	{Value: "1090", Label: "فروشندگان محصولات گوشتی"},
	{Value: "1037", Label: "فروشندگان نفت"},
	{Value: "4063", Label: "فروشندگان و تعمیرکاران لوازم خانگی"},
	{Value: "4129", Label: "فعالیت های غیر تجاری"},
	{Value: "4040", Label: "فلزات آهنی و غیر آهنی"},
	{Value: "2038", Label: "قالیشویان"},
	{Value: "2070", Label: "قصابان"},
	{Value: "2039", Label: "قنادی و شیرینی و کافه قنادی"},
	{Value: "2040", Label: "قندریزان"},
	{Value: "1038", Label: "قهوه خانه داران"},
	{Value: "4031", Label: "گالری"},
	{Value: "1044", Label: "گچ و آهک و  مصالح ساختمانی"},
	{Value: "1043", Label: "گرمابه داران"},
	{Value: "2047", Label: "گلگیر و رادیاتور"},
	{Value: "1045", Label: "گوشت گاوی"},
	{Value: "1046", Label: "گوشت گوسفندی"},
	{Value: "4035", Label: "گونی و لوازم"},
	{Value: "2048", Label: "لاستیک و روغن"},
	{Value: "1052", Label: "لباس و پوشاک دوخته و پوشاک فروشان"},
	{Value: "1050", Label: "لبنیات"},
	{Value: "2069", Label: "لوازم آرایشی و بهداشتی"},
	{Value: "2067", Label: "لوازم التحریر"},
	{Value: "2049", Label: "لوازم الکتریک و سیم کشی"},
	{Value: "1047", Label: "لوازم بهداشتی و ساختمانی"},
	{Value: "2072", Label: "لوازم خانگی"},
	{Value: "1089", Label: "لوازم دندانپزشکی"},
	{Value: "4007", Label: "لوازم شکار و ماهیگیری"},
	{Value: "2073", Label: "لوازم صوتی و تصویری"},
	{Value: "1051", Label: "لوازم فلزی خانگی"},
	{Value: "2050", Label: "لوازم مسی"},
	{Value: "4043", Label: "لوازم مهندسی"},
	{Value: "4010", Label: "لوازم و مواد اولیه دامداری"},
	{Value: "1049", Label: "لوازم ورزشی"},
	{Value: "1048", Label: "لوازم یدکی اتومبیل"},
	{Value: "2071", Label: "لوله فروشان"},
}

// GetCurrentJalaliDate returns current date in Jalali format (YYYY/MM/DD).
func GetCurrentJalaliDate() string {
	// Simple Gregorian to Jalali conversion for current date
	// For a production app, use a proper Jalali calendar library
	now := time.Now()
	year, month, day := gregorianToJalali(now.Year(), int(now.Month()), now.Day())
	return formatJalaliDate(year, month, day)
}

// gregorianToJalali converts Gregorian date to Jalali date.
// This is a simplified conversion - for production, use a proper library.
func gregorianToJalali(gy, gm, gd int) (jy, jm, jd int) {
	var g_days_in_month = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var j_days_in_month = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

	gy2 := gy - 1600
	gm2 := gm - 1
	gd2 := gd - 1

	g_day_no := 365*gy2 + (gy2+3)/4 - (gy2+99)/100 + (gy2+399)/400
	for i := 0; i < gm2; i++ {
		g_day_no += g_days_in_month[i]
	}
	if gm2 > 1 && ((gy%4 == 0 && gy%100 != 0) || gy%400 == 0) {
		g_day_no++
	}
	g_day_no += gd2

	j_day_no := g_day_no - 79

	j_np := j_day_no / 12053
	j_day_no = j_day_no % 12053

	jy = 979 + 33*j_np + 4*(j_day_no/1461)
	j_day_no = j_day_no % 1461

	if j_day_no >= 366 {
		jy += (j_day_no - 1) / 365
		j_day_no = (j_day_no - 1) % 365
	}

	for i := 0; i < 11 && j_day_no >= j_days_in_month[i]; i++ {
		j_day_no -= j_days_in_month[i]
		jm = i + 2
	}
	if jm == 0 {
		jm = 1
	}
	jd = j_day_no + 1

	return jy, jm, jd
}

// formatJalaliDate formats Jalali date as YYYY/MM/DD.
func formatJalaliDate(year, month, day int) string {
	return fmt.Sprintf("%04d/%02d/%02d", year, month, day)
}

// ApplyBasicInfoDefaults applies default values to empty basic info fields.
func ApplyBasicInfoDefaults(req *BasicInfoFormData, defaults config.BasicInfoDefaults) {
	if req.RegistrationReason == "" {
		req.RegistrationReason = defaults.RegistrationReason
	}
	if req.ActivityType == "" {
		req.ActivityType = defaults.ActivityType
	}
	if req.StartDate == "" || req.StartDate == "1xxx/xx/xx" {
		req.StartDate = GetCurrentJalaliDate()
	}
	if req.EightCategoryJob == "" || req.EightCategoryJob == "نامشخص" {
		req.EightCategoryJob = defaults.EightCategoryJob
	}
	if req.ProfessionalAssembly == "" {
		req.ProfessionalAssembly = defaults.ProfessionalAssembly
	}
	if req.BusinessLicense == "" {
		req.BusinessLicense = defaults.BusinessLicense
	}
	if req.OwnershipType == "" || req.OwnershipType == "نامشخص" {
		req.OwnershipType = defaults.OwnershipType
	}
}

// ApplyMemberDefaults applies default values to empty member fields.
func ApplyMemberDefaults(req *MemberSubmitRequest, defaults config.MemberDefaults) {
	if req.PersonType == "" {
		req.PersonType = defaults.PersonType
	}
	if req.Nationality == "" {
		req.Nationality = defaults.Nationality
	}
	if req.BirthCountry == "" {
		req.BirthCountry = defaults.BirthCountry
	}
	if req.NationalCardType == "" {
		req.NationalCardType = defaults.NationalCardType
	}
	// MembershipType maps to PartnershipType in config
	if req.MembershipType == "" {
		req.MembershipType = defaults.PartnershipType
	}
	// IsResponsible maps to IsEmployed in config
	if req.IsResponsible == "" {
		req.IsResponsible = defaults.IsEmployed
	}
	if req.SignatureAuthority == "" {
		req.SignatureAuthority = defaults.SignatureAuthority
	}
	if req.ResponsibilityType == "" {
		req.ResponsibilityType = defaults.ResponsibilityType
	}
	if req.Position == "" {
		req.Position = defaults.Position
	}
	if req.StartDate == "" || req.StartDate == "1xxx/xx/xx" {
		req.StartDate = GetCurrentJalaliDate()
	}
	if req.EndDate == "" || req.EndDate == "1xxx/xx/xx" {
		req.EndDate = defaults.EndDate
	}
}

// ApplyINTACodeDefaults applies default values to empty INTA code fields.
func ApplyINTACodeDefaults(req *INTACodeSubmitRequest, defaults config.INTACodeDefaults) {
	if req.Code == "" {
		req.Code = defaults.Code
	}
	if req.Description == "" {
		req.Description = defaults.Description
	}
	if req.Percent == 0 {
		req.Percent = defaults.Percent
	}
}

// BasicInfoFormData represents the basic info form data structure.
type BasicInfoFormData struct {
	RegistrationReason   string `json:"registrationReason"`
	ActivityType         string `json:"activityType"`
	StartDate            string `json:"startDate"`
	UnitTitle            string `json:"unitTitle"`
	EightCategoryJob     string `json:"eightCategoryJob"`
	IndividualJob        string `json:"individualJob"`
	ProfessionalGuild    string `json:"professionalGuild"`
	ProfessionalAssembly string `json:"professionalAssembly"`
	GuildUnion           string `json:"guildUnion"`
	NewGuildUnion        string `json:"newGuildUnion"`
	BusinessLicense      string `json:"businessLicense"`
	LicenseAuthority     string `json:"licenseAuthority"`
	OwnershipType        string `json:"ownershipType"`
	Phone                string `json:"phone"`
	AreaCode             string `json:"areaCode"`
	Fax                  string `json:"fax"`
	FaxAreaCode          string `json:"faxAreaCode"`
	Email                string `json:"email"`
	Website              string `json:"website"`
	HasBusinessCard      bool   `json:"hasBusinessCard"`
	HasCoinTax           bool   `json:"hasCoinTax"`
}

// INTACodeSubmitRequest represents an INTA code submission request.
type INTACodeSubmitRequest struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Percent     int    `json:"percent"`
}

// GetDefaultBasicInfo returns a BasicInfoFormData with all defaults applied.
func GetDefaultBasicInfo(defaults config.BasicInfoDefaults, businessName string) BasicInfoFormData {
	return BasicInfoFormData{
		RegistrationReason:   defaults.RegistrationReason,
		ActivityType:         defaults.ActivityType,
		StartDate:            GetCurrentJalaliDate(),
		UnitTitle:            businessName,
		EightCategoryJob:     defaults.EightCategoryJob,
		ProfessionalAssembly: defaults.ProfessionalAssembly,
		BusinessLicense:      defaults.BusinessLicense,
		OwnershipType:        defaults.OwnershipType,
	}
}

// GetDefaultMember returns a MemberSubmitRequest with all defaults applied.
func GetDefaultMember(defaults config.MemberDefaults) MemberSubmitRequest {
	return MemberSubmitRequest{
		PersonType:         defaults.PersonType,
		Nationality:        defaults.Nationality,
		BirthCountry:       defaults.BirthCountry,
		NationalCardType:   defaults.NationalCardType,
		MembershipType:     defaults.PartnershipType,
		IsResponsible:      defaults.IsEmployed,
		SignatureAuthority: defaults.SignatureAuthority,
		ResponsibilityType: defaults.ResponsibilityType,
		Position:           defaults.Position,
		StartDate:          GetCurrentJalaliDate(),
		EndDate:            defaults.EndDate,
	}
}

// GetDefaultINTACode returns an INTACodeSubmitRequest with all defaults applied.
func GetDefaultINTACode(defaults config.INTACodeDefaults) INTACodeSubmitRequest {
	return INTACodeSubmitRequest{
		Code:        defaults.Code,
		Description: defaults.Description,
		Percent:     defaults.Percent,
	}
}

// GetHappyPathPublicData returns a PublicDataRequest with hardcoded happy path values.
// These values are extracted from curl.md and are known to work with the portal.
// Use this when you want guaranteed working form values.
func GetHappyPathPublicData(businessName string) *PublicDataRequest {
	return &PublicDataRequest{
		RegistrationCause:   HappyPathDefaults.RegistrationReason,  // "0" = ایجاد یک کسب و کار جدید
		IsTejari:            HappyPathDefaults.IsTejari,            // "1" = تجاری
		FinancialStartDate:  GetCurrentJalaliDate(),                // Current Jalali date
		BusinessName:        businessName,                          // User-provided business name
		GroupOneType:        HappyPathDefaults.GroupOneType,        // "100" = سایر (گروه دوم و سوم)
		LegalType:           HappyPathDefaults.LegalType,           // "4032" = مشاوره و خدمات
		NewLegalGroup:       HappyPathDefaults.NewLegalGroup,       // "6" = اتاق اصناف ایران
		NewLegalType:        HappyPathDefaults.NewLegalType,        // "9090" = اتحادیه کشوری - سایر
		HasJobLicense:       HappyPathDefaults.HasJobLicense,       // "0" = فعلا فاقد جواز کسب
		Ownership:           HappyPathDefaults.Ownership,           // "1" = ملکی
		FinancialDayStart:   HappyPathDefaults.FinancialDayStart,   // "1" = First day
		FinancialMonthStart: HappyPathDefaults.FinancialMonthStart, // "1" = فروردین
	}
}

// GetHappyPathINTACode returns an INTA code activity for happy path testing.
// Uses management consulting/services INTA code verified via Playwright testing.
// Cascade path: [3] خدمات → [11] خدمات امور اداری → [0] خدمات → [3110100]
func GetHappyPathINTACode() INTAActivity {
	return INTAActivity{
		Code:        "3110100", // مدیریت، مشاوره و نظارت در امور فنی، مدیریتی و اداری
		Description: "خدمات مشاوره مدیریتی و اداری",
		Percent:     100,
	}
}

// GetHappyPathVATStatus returns a VAT status request for happy path.
func GetHappyPathVATStatus() *VATStatusRequest {
	return &VATStatusRequest{
		EligibilityType: "1", // عدم مشمولیت
	}
}
