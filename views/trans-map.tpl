 
{{if .transMap}}

{{range $key, $value := .transMap}}
    <option value="{{$key}}">{{$value}}</option>
{{end}}

{{end}}