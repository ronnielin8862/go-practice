package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	GetTextLiveReq struct {
		RoomId int `json:"id"`
		LastId int `json:"last_id"`
		Size   int `json:"size"`
	}
)

func main() {
	var r GetTextLiveReq
	e := echo.New()

	e.POST("/api/tencentIm/getUserSig", func(c echo.Context) error {

		fmt.Println("hello 22222: ", c.Get("uid"))
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", func(c echo.Context) error {

		fmt.Println("hello : ", c.Get("uid"))
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/get_path", func(c echo.Context) error {
		m := c.QueryParams()
		fmt.Println("id = ", m["id"])
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/binding_path", func(c echo.Context) error {
		err := echo.QueryParamsBinder(c).Int("id", &r.RoomId).Int("last_id", &r.LastId).Int("size", &r.Size).BindErrors()
		if err != nil {
			fmt.Println("err = ", err)
		}

		fmt.Println("GetTextLiveReq = ", r)
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/post_path", func(c echo.Context) error {
		fmt.Println("hello : ", fmt.Sprint(c.Get("uid")))
		fmt.Println("post_path = ", c)
		return c.String(http.StatusOK, "post_path")
	})

	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":9487"))
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

//func GetMatchText(c echo.Context) (err error) {
//	var r globle.GetTextLiveReq
//	errs := echo.QueryParamsBinder(c).Int("id", &r.RoomId).Int("last_id", &r.LastId).Int("size", &r.Size).BindErrors()
//	if errs != nil || r.RoomId == 0 || r.LastId == 0 {
//		fmt.Println("GetMatchText param err : %v", errs)
//		return c.JSON(http.StatusBadRequest, "param err")
//	}
//
//	result, err := GetMatchTextService(r)
//	if err != nil {
//		fmt.Println("GetMatchText err : %v", err)
//		return c.JSON(http.StatusBadRequest, err)
//	}
//	return c.JSON(http.StatusOK, result)
//}

//func GetMatchTextService(r globle.GetTextLiveReq) (tl []messages.TextLiveStruct, err error) {
//	startId := r.LastId - r.Size
//	if startId < 1 {
//		startId = 1
//	}
//	sort := elastic.NewFieldSort("create_time").Order(false)
//	result, err := globle.GetEs().Search(messages.MatchLiveConst, "_", messages.TextConst, "_", strconv.Itoa(r.RoomId)).From(startId).Size(r.Size).SortBy(sort).Do(context.Background())
//	if err != nil {
//		return nil, err
//	}
//
//	fmt.Println("1 = ", result.Hits.TotalHits.Value)
//	return nil, nil
//}
