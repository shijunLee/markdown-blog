<html>
<title> love zhao-xiaoying's Blog! </title>
<body>
<h1></h1>
<img src="/static/img/a.jpg">
{{range .posts}}
<a href="posts/{{.File}}"><h2>{{.Title}} ({{.Date}})</h2></a>
<p>{{str2html  .Summary}}</p>
{{end}}
</body>
</html>