<style>
    .layout-left{
        width: 190px;
        float: left;
        height: 100%;
    }
</style>
<div class="layout layout-left page-sidebar">
    <p>{{.userInfo.Nickname}}</p>
    <ul>
    {{printf "The projects length is %d" (.projects|len)}}
        {{/*{{if gt (.projects|len) 0}}
            {{range .projects}}
            <li><a href="/project/{{.ID}}">.Name</a></li>
            {{end}}
        {{end}}*/}}
    </ul>
</div>