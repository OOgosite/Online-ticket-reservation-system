package index

import (
	"fmt"
	"http"
	//"template"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, mainPage)
}

const mainPage = `<!DOCTYPE HTML>
<html>
	<head>
		<title>ระบบจองตั๋วออนไลน์</title>
		<meta content="text/html; charset=UTF-8" http-equiv="content-type">
		<meta content="chatchai" name="author">
		<meta content="ระบบจองตั๋วออนไลน์ เป็นระบบจำลองการจองตั๋วรถ โดยใช้ภาษา Go ทำงานบน Google App Engine" name="description">
	</head>
  	<body>
  	<center>
  	<table style="text-align: left; width: 981px; height: 344px;" border="0"
cellpadding="0" cellspacing="0">
<tbody>
<tr align="center">
<td colspan="3" rowspan="1"
style="vertical-align: top; height: 193px; background-color: rgb(51, 51, 255); width: 440px;"><big
style="color: white;"><big><big><small><small><small><small><small><small><br>
</small></small></small></small></small></small>ระบบจองตั๋วออนไลน์<br>
<small><small><br>
มหาวิทยาลัยเทคโนโลยีสุรนารี</small></small><br>
</big></big></big></td>
</tr>
<tr>
<td
style="vertical-align: top; height: 11px; width: 198px; text-align: center;">เข้าสู่ระบบ<br>
</td>
<td colspan="2" rowspan="1"
style="vertical-align: top; height: 11px; width: 440px;">Home |
ตรวจสอบการจอง | เส้นทาง | คำถามที่พบบ่อย | เกี่ยวกับเรา<br>
</td>
</tr>
<tr>
<td style="vertical-align: top; width: 198px;"><br>
</td>
<td style="vertical-align: top; width: 440px;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
</tr>
<tr>
<td style="vertical-align: top;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
</tr>
<tr>
<td style="vertical-align: top;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
</tr>
<tr>
<td style="vertical-align: top;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
<td style="vertical-align: top;"><br>
</td>
</tr>
</tbody>
</table>
<br>  
	</center>
  	</body>
</html>
`