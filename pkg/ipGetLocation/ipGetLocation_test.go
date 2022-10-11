package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/kayon/iploc"
	"testing"
)

func Test1(t *testing.T) {
	ip := "101.226.168.228"
	iPGetLocation(ip)
}

func iPGetLocation(ip string) {
	loc, err := iploc.Open("/Users/ronnie/Documents/coding/code/go-practice/pkg/ipGetLocation/qqwry.dat/qqwry_lastest.dat")
	if err != nil {
		panic(err)
	}
	detail := loc.Find("8.8.8") // 补全为8.8.0.8, 参考 ping 工具
	fmt.Printf("IP:%s; 网段:%s - %s; %s\n", detail.IP, detail.Start, detail.End, detail)

	detail2 := loc.Find("8.8.3.1")
	fmt.Printf("%t %t\n", detail.In(detail2.IP.String()), detail.String() == detail2.String())

	// output
	// IP:8.8.0.8; 网段: 8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
	// true true

	detail = loc.Find(ip)
	fmt.Println(detail.String())
	fmt.Println(detail.Country, detail.Province, detail.City, detail.County)

	fmt.Println("22 = ", ConvertToString(detail.String(), "gbk", "utf-8"))
	fmt.Println("Country = ", ConvertToString(detail.Country, "gbk", "utf-8"))
	fmt.Println("Province = ", ConvertToString(detail.Province, "gbk", "utf-8"))
	fmt.Println("City = ", ConvertToString(detail.City, "gbk", "utf-8"))
	fmt.Println("County = ", ConvertToString(detail.County, "gbk", "utf-8"))
	fmt.Println("Region = ", ConvertToString(detail.Region, "gbk", "utf-8"))
	fmt.Println("ip = ", detail.IP)
	// output
	// 内蒙古锡林郭勒盟苏尼特右旗 联通
	// 中国 内蒙古 锡林郭勒盟 苏尼特右旗

}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
