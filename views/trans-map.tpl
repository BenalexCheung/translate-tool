 
{{if .TransMap}}

{{range $key, $value := .TransMap}}
    {{if eq "auto" $key}}
    <option value="{{$key}}" selected="selected">{{$value}}</option>
    {{else}}
    <option value="{{$key}}">{{$value}}</option>
    {{end}}
{{end}}

{{end}}