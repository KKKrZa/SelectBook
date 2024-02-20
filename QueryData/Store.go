package QueryData

// 保存查询相关全局变量

const (
	bookList  string = "https://api.quxuanshu.com/student/order/otherClassCourseOrder/listCourseBookInfo"
	majorList string = "https://api.quxuanshu.com/student/order/otherClassCourseOrder/listMajorInfo"
	isbnList  string = "https://api.ibook.tech/book/api/search/books"
)

var secret Secret

var (
	Books    BookInfo
	Majors   MajorInfo
	MySelect = SelectInfo{
		TermCode:               23242,
		Grade:                  0,
		MajorId:                "",
		PageClassId:            nil,
		CourseId:               nil,
		Title:                  nil,
		Isbn:                   nil,
		IsSearchTextbookType:   false,
		IsSearchMajorGradeType: true,
	}
)
