 
{{if .TransMap}}

{{range $key, $value := .TransMap}}
    <option value="{{$key}}">{{$value}}</option>
{{end}}

{{end}}