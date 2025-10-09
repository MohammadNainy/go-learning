package main

import "fmt"

// تعریف ساختار student برای ذخیره اطلاعات یک دانش‌آموز
type student struct {
	Name  string  // نام دانش‌آموز
	Age   int     // سن دانش‌آموز
	Grade float64 // نمره دانش‌آموز
}

// متد Display برای نمایش اطلاعات دانش‌آموز
func (s student) Display() {
	// چاپ اطلاعات دانش‌آموز با فرمت مشخص
	fmt.Printf("Student: %s | Age: %v | Grade: %.2f\n", s.Name, s.Age, s.Grade)
}

// متد IsAdult برای بررسی بزرگسال بودن دانش‌آموز
func (s student) IsAdult() bool {
	// اگر سن دانش‌آموز 18 یا بیشتر باشد، true برمی‌گرداند
	return s.Age >= 18
}

// متد UpdateGrade برای به‌روزرسانی نمره دانش‌آموز
// استفاده از اشاره‌گر (*student) برای تغییر مقدار واقعی ساختار
func (s *student) UpdateGrade(g float64) {
	s.Grade = g // به‌روزرسانی نمره
}

func main() {
	// ایجاد یک نمونه از ساختار student با نام Ali
	Ali := student{"Ali", 20, 18.00}
	
	// فراخوانی متد Display برای نمایش اطلاعات اولیه Ali
	Ali.Display()
	
	// بررسی بزرگسال بودن Ali و چاپ نتیجه
	fmt.Println(Ali.Name, " is an adult: ", Ali.IsAdult())
	
	// به‌روزرسانی نمره Ali به 12.5
	Ali.UpdateGrade(12.5)
	
	// نمایش اطلاعات به‌روزرسانی شده Ali
	Ali.Display()
}