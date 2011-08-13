package index

import (
	"fmt"
	"http"
	
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
		<meta content="text/html; charset=UTF-8" http-equiv="content-type">
		<title>ระบบจองตั๋วรถออนไลน์</title>
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
					<td style="vertical-align: top; height: 41px; width: 143px;">
					<img onmouseout='src="images/down_02.gif"' style="width: 165px; height: 38px;"onmouseover='src="images/up_02.gif"' alt="" src="images/down_02.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 148px;">
					<img onmouseout='src="images/down_03.gif"' style="width: 234px; height: 38px;" onmouseover='src="images/up_03.gif"' alt="" src="images/down_03.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 69px;">
					<img onmouseout='src="images/down_04.gif"' onmouseover='src="images/up_04.gif"' style="width: 151px; height: 38px;" alt="" src="images/down_04.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 503px;">
					<img onmouseout='src="images/down_05.gif"' onmouseover='src="images/up_05.gif"' style="width: 229px; height: 38px;" alt="" src="images/down_05.gif"></td>
					<td style="vertical-align: top; height: 41px; width: 42px;">
					<img onmouseout='src="images/down_06.jpg"' onmouseover='src="images/up_06.jpg"' style="width: 201px; height: 38px;" alt="" src="images/down_06.jpg"></td>
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
				<h3><br>ยินดีต้อนรับเข้าสู่ระบบจองตั๋วรถออนไลน์<br></h3>
				<h4>ระบบนี้เป็นการจำลองการจองตั๋วรถเมล์ภายในมหาวิทยาลัยเทคโนโลยีสุรนารี<br>
				เพื่อเป็นต้นแบบในการนำไปประยุกต์ใช้งานต่อไป</h4>
				<br>
				<br>
				<hr style="width: 90%; height: 2px;"><br>
			<center>
				<form action="/search" method="post">
				<table style="text-align: left; width: 550px;" border="0" cellpadding="0" cellspacing="0">
					<tbody>
					<tr>
						<td colspan="2" rowspan="1" style="vertical-align: top; text-align: center;"><br>
						<big style="font-weight: bold;">ค้นหาเที่ยวรถ</big><br>
						<br>
						</td>
					</tr>
					<tr>
						<td style="vertical-align: top; text-align: right;"><span style="color: rgb(0, 0, 153);">จุดขึ้นรถ:</span> <br>
						</td>
						<td style="vertical-align: top;">
						<select name="start">
							<option value="">-เลือกจุดขึ้นรถ-</option>
							<option>สถานีขนส่ง</option>
							<option>อาคารบริหาร</option>
							<option>หอพักสุรนิเวศ 15</option>
							<option>หอพักสุรนิเวศ1-6</option>
							<option>หอพักสุรนิเวศ16</option>
							<option>โรงอาหารหอพัก</option>
							<option>หอพักสุรนิเวศ13</option>
							<option>หอพักสุรนิเวศ7-12</option>
							<option>อาคารเรียนรวม1</option>
							<option>อาคารเรียนรวม2</option>
							<option>บรรณสาร</option>
							<option>อาคารเครื่องมือ</option>
							<option></option>
						</select>
						<br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;"><span style="color: rgb(0, 0, 153);">จุดลงรถ: </span><br>
					</td>
					<td style="vertical-align: top;">
						<select name="stop">
							<option value="">-เลือกจุดลงรถ-</option>
							<option>สถานีขนส่ง</option>
							<option>อาคารบริหาร</option>
							<option>หอพักสุรนิเวศ 15</option>
							<option>หอพักสุรนิเวศ1-6</option>
							<option>หอพักสุรนิเวศ16</option>
							<option>โรงอาหารหอพัก</option>
							<option>หอพักสุรนิเวศ13</option>
							<option>หอพักสุรนิเวศ7-12</option>
							<option>อาคารเรียนรวม1</option>
							<option>อาคารเรียนรวม2</option>
							<option>บรรณสาร</option>
							<option>อาคารเครื่องมือ</option>
							<option></option>
						</select>
						<br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;"><span style="color: rgb(0, 0, 153);">วัน/เวลาที่เดินทาง: </span><br>
					</td>
					<td style="vertical-align: top;">
						<select name="datetime">
							<option value="">-เลือกวันที่เดินทาง-</option>
							<option>วันจันทร์</option>
							<option>วันอังคาร</option>
							<option>วันพุธ</option>
							<option>วันพฤหัสบดี</option>
							<option>วันศุกร์</option>
							<option>วันเสาร์</option>
							<option>วันอาทิตย์</option>
							<option></option>
						</select>
						<br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top;"><br>
					</td>
					<td style="vertical-align: top;"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top;"><br>
					</td>
					<td style="vertical-align: top;"><input name="submit" value="ค้นหาเที่ยวรถ" type="submit"><br>
					</td>
				</tr>
			</tbody>
		</table>
		</form>
	</center>
	<br>
	<hr style="width: 90%; height: 2px;"><br>
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
		<div style="text-align: left;"><a href="/signin" target="_top">สมัครสมาชิก</a><br>
		<a href="/forgetpass" target="_top">ลืมรหัสผ่าน</a><br>
		<br>
		</div>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);"><img style="width: 200px; height: 30px;" alt="" src="images/mainmenu.jpg"><br>
	</td>
	</tr>
	<tr>
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);">หน้าแรก<br>
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
	<td style="vertical-align: top; background-color: rgb(255, 164, 10);">เกี่ยวกับเรา<br>
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