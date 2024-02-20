package QueryData

type BookInfo struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		List []struct {
			OrderClassItemId    string  `json:"orderClassItemId"`
			CheckCode           string  `json:"checkCode"`
			Title               string  `json:"title"`
			Isbn                string  `json:"isbn"`
			Price               float64 `json:"price"`
			PressId             string  `json:"pressId"`
			PressName           string  `json:"pressName"`
			ReleaseDate         *string `json:"releaseDate"`
			ImgUrl              *string `json:"imgUrl"`
			CourseId            string  `json:"courseId"`
			CourseCode          string  `json:"courseCode"`
			CourseName          string  `json:"courseName"`
			CourseNature        string  `json:"courseNature"`
			CourseType          string  `json:"courseType"`
			CourseCollegeName   string  `json:"courseCollegeName"`
			MajorNames          string  `json:"majorNames"`
			TeacherNames        string  `json:"teacherNames"`
			TextbookDiscount    string  `json:"textbookDiscount"`
			TextbookFloatRatio  string  `json:"textbookFloatRatio"`
			DiscountedPrice     float64 `json:"discountedPrice"`
			IsCanOrder          int     `json:"isCanOrder"`
			IsMustOrder         int     `json:"isMustOrder"`
			IsOrderedAndPayed   int     `json:"isOrderedAndPayed"`
			AvailableStockCount int     `json:"availableStockCount"`
		} `json:"list"`
		IsHaveUnpaidSysGenerateOrder int `json:"isHaveUnpaidSysGenerateOrder"`
	} `json:"data"`
}

type MajorInfo struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		Grade       int    `json:"grade"`
		MajorId     string `json:"majorId"`
		MajorCode   string `json:"majorCode"`
		MajorName   string `json:"majorName"`
		CollegeName string `json:"collegeName"`
	} `json:"data"`
}

type SelectInfo struct {
	TermCode               int         `json:"termCode"`
	Grade                  int         `json:"grade"`
	MajorId                string      `json:"majorId"`
	PageClassId            interface{} `json:"pageClassId"`
	CourseId               interface{} `json:"courseId"`
	Title                  interface{} `json:"title"`
	Isbn                   interface{} `json:"isbn"`
	IsSearchTextbookType   bool        `json:"isSearchTextbookType"`
	IsSearchMajorGradeType bool        `json:"isSearchMajorGradeType"`
}

type IsbnMapInfo struct {
	Errcode int `json:"errcode"`
	Data    struct {
		Total int `json:"total"`
		Books []struct {
			Img       string `json:"img"`
			Author    string `json:"author"`
			Isbn      string `json:"isbn"`
			Language  string `json:"language"`
			Title     string `json:"title"`
			Price     string `json:"price"`
			Publisher string `json:"publisher"`
			Id        string `json:"id"`
			Category  string `json:"category"`
			Pubdate   string `json:"pubdate"`
		} `json:"Books"`
	} `json:"data"`
	Errmsg string `json:"errmsg"`
}

func (sel *SelectInfo) SetQueryGrade(grade int) {
	sel.Grade = grade
}

func (sel *SelectInfo) SetQueryMajorId(majorId string) {
	sel.MajorId = majorId
}
