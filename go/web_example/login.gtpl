<html>
<head>
<title>Test form</title>
</head>
<body>
<form action="/login" method="post">
  <input type="checkbox" name="interest" value="football"> Football
  <input type="checkbox" name="interest" value="basketball"> Basketball
  <input type="checkbox" name="interest" value="tennis"> Tennis
  <label>UserName: </label>
  <input type="text" name="username">
  <label>Password: </password>
  <input type="password" name="password">
  <input type="hidden" name="token" value="{{.}}">
  <input type="submit" value="ログイン">
</form>
</body>
</html>
