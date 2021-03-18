<!DOCTYPE html>
<html lang="zh-CN">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <title>Pantum 翻译</title>

    <link rel="icon" href="/static/img/favicon.ico" sizes="64x64">

    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/js/ladda-bootstrap/ladda-themeless.min.css">

    <!-- HTML5 shim 和 Respond.js 是为了让 IE8 支持 HTML5 元素和媒体查询（media queries）功能 -->
    <!-- 警告：通过 file:// 协议（就是直接将 html 页面拖拽到浏览器中）访问页面时 Respond.js 不起作用 -->
    <!--[if lt IE 9]>
      <script src="https://cdn.jsdelivr.net/npm/html5shiv@3.7.3/dist/html5shiv.min.js"></script>
      <script src="https://cdn.jsdelivr.net/npm/respond.js@1.4.2/dest/respond.min.js"></script>
    <![endif]-->
    {{template "header.tpl" .}}
</head>

<body>
    <div class="well well-sm">
        <h3><a class="button2" href="/" role="" text-decoration:none>Pantum 翻译</a></h3>
    </div>

    <div class="container">
        <form id="postForm" class="form-group" action="/?sl=en&tl=zh-CN&op=docs" method="get">
            <div class="btn-group" data-toggle="buttons">
                <label class="btn button3" id="label1">
                    <input type="radio" id="trans" name="op" value="translate" {{if .IsDocs}}{{else}}checked="checked"{{end}}/>
                    <span class="glyphicon glyphicon glyphicon-font"></span> 文字
                </label>
                <label class="btn button3" id="label2">
                    <input type="radio" id="docs" name="op" value="docs" {{if .IsDocs}}checked="checked"{{else}}{{end}}/>
                    <span class="glyphicon glyphicon-file"></span> 文档
                </label>
            </div>
            <br><br>

            {{.LayoutContent}}
        </form>
    </div>

    <br><br>
    {{template "footer.tpl" .}}

    <!-- jQuery (Bootstrap 的所有 JavaScript 插件都依赖 jQuery，所以必须放在前边) -->
    <script src="/static/js/jquery.min.js"></script>
    <!-- 加载 Bootstrap 的所有 JavaScript 插件。你也可以根据需要只加载单个插件。 -->
    <script src="/static/js/bootstrap.min.js"></script>

    {{template "scripts.tpl" .}}

</body>

</html>