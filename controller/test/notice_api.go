package test

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/test/notice"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

type notice_api struct{}

// 插入
func (d *notice_api) CreateNoticeApi(c *gin.Context) {
	var noticeList []notice.SysNotice
	err := c.ShouldBindJSON(&noticeList)
	if err != nil {
		fmt.Println("参数错误", err.Error())
		response.FailWithMessage("添加的参数错误", c)
		return
	}
	// 添加数据
	result := global.Global_Db.Create(&noticeList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *notice_api) DeleteNoticeApi(c *gin.Context) {
	var noticeList []notice.SysNotice
	err := c.ShouldBindJSON(&noticeList)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	//遍历查寻数据是否存在
	var notice notice.SysNotice
	for _, value := range noticeList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&notice)
		if err2.Error != nil {
			response.FailWithMessage("删除的数据不存在:", c)
			return
		}
	}
	for _, del := range noticeList {
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("该数据不存在,无法删除:", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *notice_api) UpdateNoticeApi(c *gin.Context) {
	var notice notice.SysNotice
	err := c.ShouldBindJSON(&notice)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	result := global.Global_Db.Model(&notice).Omit("id").Where("id=?", notice.Id).Updates(notice)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *notice_api) QueryNoticeApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	var total int64
	var Notice notice.SysNotice
	var noticeList []notice.SysNotice
	// 获取query
	rawUrl := c.Request.URL.String()
	u, er := url.Parse(rawUrl)
	if er != nil {
		fmt.Println("解析url错误")
	}
	queryParams := u.Query()
	fmt.Println("查寻字符串参数", queryParams)
	// 获取请求体数据

	condition := make(map[string]interface{})
	for index, value := range queryParams {
		key := utils.ToCamelCase(index)
		condition[key] = value
	}
	fmt.Println("condition", condition)

	// 分页数据
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		fmt.Println("分页数错误", er1.Error(), er2.Error())
	}
	offset = PageSize * (Page - 1)
	limit = PageSize
	fmt.Println(offset, limit)
	title := fmt.Sprintf("%v", condition["title"])
	title = strings.Trim(title, "[]")
	fmt.Println("titles5555数据", title)
	// 查寻数量

	count := global.Global_Db.Model(&Notice).Where("title LIKE ?", "%"+title+"%").Count(&total).Error
	if count != nil {
		fmt.Println("查寻数量错误")
		response.FailWithMessage("系统查询错误", c)
		return
	}
	// 查寻数据

	result := global.Global_Db.Where("title LIKE ?", "%"+title+"%").Limit(limit).Offset(offset).Order("timestamp desc").Find(&noticeList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     noticeList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Notice_api = new(notice_api)
