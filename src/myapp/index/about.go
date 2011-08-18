package index

import (
	"fmt"
	"http"
	
)

func init() {
	http.HandleFunc("/about", about)
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, aboutPage)
}

const aboutPage = `<!DOCTYPE HTML>
<html>
	<head>
		<meta content="text/html; charset=UTF-8" http-equiv="content-type">
		<title>เกี่ยวกับเรา</title>
		<meta content="chatchai" name="author">
		<meta content="ระบบจองตั๋วรถออนไลน์ มหาวิทยาลัยเทคโนโลยีสุรนารี"name="description">
	</head>
	<body>
		<center>
			<table style="text-align: left; height: 172px; width: 980px;" border="0" cellpadding="0" cellspacing="0">
			<tbody>
				<tr>
					<td colspan="5" rowspan="1" style="vertical-align: bottom; height: 172px; width: 980px; background-color: rgb(255, 164, 10);">
					<img style="width: 980px; height: 172px;" alt="" src="images/down_01.gif"></td>
				</tr>
				<tr>
					<td style="vertical-align: top; height: 41px; width: 143px;"><a href='/'>
					<img onmouseout='src="images/down_02.gif"' style="width: 165px; height: 38px;"onmouseover='src="images/up_02.gif"' alt="" src="images/down_02.gif"></a></td>
					<td style="vertical-align: top; height: 41px; width: 148px;">
					<img onmouseout='src="images/down_03.gif"' style="width: 234px; height: 38px;" onmouseover='src="images/up_03.gif"' alt="" src="images/down_03.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 69px;">
					<img onmouseout='src="images/down_04.gif"' onmouseover='src="images/up_04.gif"' style="width: 151px; height: 38px;" alt="" src="images/down_04.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 503px;">
					<img onmouseout='src="images/down_05.gif"' onmouseover='src="images/up_05.gif"' style="width: 229px; height: 38px;" alt="" src="images/down_05.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 42px;"><a href='/about'>
					<img onmouseout='src="images/down_06.jpg"' onmouseover='src="images/up_06.jpg"' style="width: 201px; height: 38px;" alt="" src="images/down_06.jpg"></a></td>
				</tr>
			</tbody>
			</table>
			<table style="text-align: left; width: 980px; height: 794px;" border="0" cellpadding="0" cellspacing="0">
			<tbody>
			<tr>
				<td style="vertical-align: top; width: 125px; background-color: rgb(255, 164, 10);">
				<img style="width: 200px; height: 30px;" alt="" src="images/login.jpg"><br>
				</td>
				<td colspan="1" rowspan="19" style="vertical-align: top; width: 849px; text-align: center;">
				<br>
				<br>
				<table style="text-align: left; width: 600px; height: 151px; margin-left: auto; margin-right: auto;" border="0" cellpadding="2" cellspacing="5">
					<tbody>
					<tr align="center">
					<td style="vertical-align: top;">
					<h2><span style="color: rgb(0, 0, 153);">เกี่ยวกับระบบจองตั๋วรถออนไลน์</span><br>
					</h2>
					</td>
					</tr>
					<tr>
					<td style="vertical-align: top;">ระบบจองตั๋วรถออนไลน์
		จัดทำขึ้นเพื่อใช้ในการศึกษาการเขียนโปรแกรมด้วยภาษา Go ของ Goole
		สำหรับทำงานบน Google App Engine และการใช้งาน datastore
		ซึ่งเป็นการศึกษาในรายวิชา 423412 ADVANCED TOPICS IN SYSTEM SOFTWARE
		ภาคการศึกษาที่ 1/2554 มหาวิทยาลัยเทคโนโลยีสุรนารี
		ระบบนี้จะเป็นการจำลองระบบการเดินรถภายในมหาวิทยาลัยเทคโนโลยีสุรนารี
		ซึ่งจะมีบางฟังก์ชั่นเท่านั้นที่สามารถทำงานได้สมบูรณ์<br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top;"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top;"><span style="color: rgb(0, 0, 153);">ผู้จัดทำ</span><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top;">นายชัชชัย บุญหลาย รหัส B512702 สาขาวิชาวิศวกรรมคอมพิวเตอร์<br>
					</td>
				</tr>
			</tbody>
		</table>
		
	
	</td>
	</td>
</tr>
<tr>
<td style="vertical-align: top; width: 125px; text-align: center; background-color: rgb(255, 164, 10);">
	<form action="/login" method="post">username<br>
		<input name="username" size="15"><br>
		password<br>
		<input name="password" size="15" type="password"><br>
		<input value="login" type="submit"></form>
		<br>
		<div style="text-align: left;"><a href='/signin' target="_top">สมัครสมาชิก</a><br>
		<a href="/forgetpass" target="_top">ลืมรหัสผ่าน</a><br>
		<a href="/view" target="_top">รายชื่อสมาชิก</a>
		<br>
		</div>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><img style="width: 200px; height: 30px;" alt="" src="images/mainmenu.jpg"><br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><a href='/' target="_top">หน้าแรก</a><br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);">ตรวจสอบการจอง<br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);">เส้นทาง</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);">คำถามที่พบบ่อย<br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><a href='/about' target="_top">เกี่ยวกับเรา</a><br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><img style="width: 200px; height: 30px;" alt="" src="images/links.jpg"><br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><a href="http://www.sut.ac.th/" target="_blank">มหาวิทยาลัยเทคโนโลยีสุรนารี</a></td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><a href="http://reg.sut.ac.th" target="_blank">ระบบลงทะเบียน</a></td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><a href="http://web.sut.ac.th/dsa/sutfund/sutfundNew/" target="_blank">งานทุนการศึกษา</a><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><a
	href="http://eng.sut.ac.th" target="_blank">สำนักวิชาวิศวกรรมศาสตร์</a><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><a
	href="http://library.sut.ac.th/" target="_blank">ศูนบรรณสาร</a><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><a
	href="http://web.sut.ac.th/coop" target="_blank">ศูนย์สหกิจศึกษา</a><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><a
	href="http://elearning.sut.ac.th/" target="_blank">SUT
	e-learning</a><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><br>
	</td>
	</tr>
	<tr>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10);"><br>
	<br>
	</td>
	<td
	style="vertical-align: top; background-color: rgb(255, 164, 10); text-align: center;">
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<span class="Apple-style-span"
	style="border-collapse: separate; color: rgb(0, 0, 0); font-family: arial,sans-serif; font-style: normal; font-variant: normal; font-weight: normal; letter-spacing: normal; line-height: normal; orphans: 2; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; font-size: small;"><span
	class="Apple-style-span" style="line-height: 16px;"><em
	style="color: rgb(204, 0, 51); font-style: normal; font-weight: normal;">Copyright</em><span
	class="Apple-converted-space">&nbsp;</span>©<span
	class="Apple-converted-space">&nbsp;</span><em
	style="color: rgb(204, 0, 51); font-style: normal; font-weight: normal;">2011</em></span></span>
	</td>
	</tr>
	<tr>
	</tr>
	</tbody>
	</table>
		</center>
	</body>
</html>
`