package views

<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <form action="/user/doAdd" method="post">
        ID     <input type="text" name="id" /> <br><br>
        用户名  <input type="text" name="username" /> <br><br>
        密 码   <input type="password" name="password" /> <br><br>

        爱 好   <input type="checkbox" value=1 id="label1" name="hobby"/> <label for="label1">吃饭</label>
                <input type="checkbox" value=2 id="label2" name="hobby"/> <label for="label2">睡觉</label>
                <input type="checkbox" value=3 id="label3" name="hobby"/> <label for="label3">敲代码</label>

        <input type="submit" value="提交">
    </form>
</body>
</html>
