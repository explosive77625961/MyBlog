{{define "nav"}}
    <ul class="nav nav-tabs">
        <li role="presentation" {{ if .IsHome }}class="active" {{end}}><a href="/">首页</a></li>
        <li role="presentation" {{ if .IsTopic }}class="active" {{end}}><a href="#">分类</a></li>
        <li role="presentation" {{ if .IsBlogs }}class="active" {{end}}><a href="#">文章</a></li>
        <li role="presentation" {{ if .IsOther }}class="active" {{end}}><a href="/other">其他</a></li>
    </ul>
{{end}}