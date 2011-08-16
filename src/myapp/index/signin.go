package index

import (
	"fmt"
	"http"
	
)

func init() {
	http.HandleFunc("/signin", signin)
}

func signin(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, signIn)
}

const signIn = `<!DOCTYPE HTML>
<html>
	<head>
		<meta content="text/html; charset=UTF-8" http-equiv="content-type">
		<title>สมัครสมาชิก</title>
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
				<td colspan="1" rowspan="19" style="vertical-align: top; width: 849px; text-align: center;"><br><br>
				<center><span style="color: rgb(0, 0, 153); font-weight: bold;">สมัครสมาชิกระบบจองตั๋วรถออนไลน์<br>
				<br>
				</span>
				<form action="/register" method="post">
				<table style="text-align: left; width: 500px; height: 59px;" border="0" cellpadding="0" cellspacing="0">
				<tbody>
				<tr align="left">
					<td colspan="2" rowspan="1" style="vertical-align: top;">กรุณากรอกข้อมูลที่เป็นจริง และครบถ้วน<br><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">ชื่อผู้ใช้:<br>
					</td>
					<td style="vertical-align: top;"><input name="user" size="15"><span style="color: rgb(102, 102, 102);"> ใช้รหัสนักศึกษา เช่นB5123456</span>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">ชื่อ - สกุล: </td>
					<td style="vertical-align: top;"><input maxlength="250" size="30" name="name"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">รหัสผ่าน: <br>
					</td>
					<td style="vertical-align: top;"><input size="15" name="passwd" type="password"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">ยืนยันรหัสผ่าน: <br>
					</td>
					<td style="vertical-align: top;"><input size="15" name="repasswd" type="password"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">เบอร์โทรศัพท์: <br>
					</td>
					<td style="vertical-align: top;"><input maxlength="50" size="15" name="phone"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">อีเมลล์: <br>
					</td>
					<td style="vertical-align: top;"><input maxlength="100" size="20" name="email"><br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">สาขาวิชา: <br>
					</td>
					<td style="vertical-align: top;">
						<select name="study">
						<option value="">-เลือกสาขาวิชาที่เรียน-</option>
						<option>สำนักวิชาวิทยาศาสตร์</option>
						<option>สำนักวิชาพยาบาลศาสตร์</option>
						<option>เทคโนโลยีสารสนเทศ</option>
						<option>เทคโนโลยีการจัดการ</option>
						<option>เทคโนโลยีการผลิตพืช</option>
						<option>เทคโนโลยีการผลิตสัตว์</option>
						<option>เทคโนโลยีอาหาร</option>
						<option>อาชีวอนามัยและความปลอดภัย</option>
						<option>อนามัยสิ่งแวดล้อม</option>
						<option>แพทยศาสตร์</option>
						<option>วิศวกรรมเกษตร</option>
						<option>วิศวกรรมขนส่ง</option>
						<option>วิศวกรรมคอมพิวเตอร์</option>
						<option>วิศวกรรมเคมี</option>
						<option>วิศวกรรมเครื่องกล</option>
						<option>วิศวกรรมเซรามิก</option>
						<option>วิศวกรรมโทรคมนาคม</option>
						<option>วิศวกรรมพอลิเมอร์</option>
						<option>วิศวกรรมไฟฟ้า</option>
						<option>วิศวกรรมโยธา</option>
						<option>วิศวกรรมโลหการ</option>
						<option>วิศวกรรมสิ่งแวดล้อม</option>
						<option>วิศวกรรมอุตสาหการ</option>
						<option>เทคโนโลยีธรณี</option>
						<option>วิศวกรรมอิเล็กทรอนิกส์ </option>
						<option>วิศวกรรมการผลิต</option>
						<option>วิศวกรรมยานยนต์</option>
						<option>วิศวกรรมอากาศยาน</option>
						<option>แมคคาทรอนิกส์</option>
						</select>
					<br>
					</td>
				</tr>
				<tr>
					<td style="vertical-align: top; text-align: right;">ที่อยู่: <br>
					</td>
					<td style="vertical-align: top;"><textarea cols="40" rows="3" name="address"></textarea><br>
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
					<td style="vertical-align: top;"><input name="submit" value="สมัครสมาชิก" type="submit">&nbsp;&nbsp;&nbsp; <input
		name="reset" value="ล้างฟอร์ม" type="reset"><br>
					</td>
				</tr>
			</tbody>
		</table>
		</form>
			
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