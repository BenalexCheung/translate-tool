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

  <!-- HTML5 shim 和 Respond.js 是为了让 IE8 支持 HTML5 元素和媒体查询（media queries）功能 -->
  <!-- 警告：通过 file:// 协议（就是直接将 html 页面拖拽到浏览器中）访问页面时 Respond.js 不起作用 -->
  <!--[if lt IE 9]>
      <script src="https://cdn.jsdelivr.net/npm/html5shiv@3.7.3/dist/html5shiv.min.js"></script>
      <script src="https://cdn.jsdelivr.net/npm/respond.js@1.4.2/dest/respond.min.js"></script>
    <![endif]-->

  <style>
    .button3 {
      background-color: white;
      color: black;
    }

    .button3:hover {
      background-color: #368ff4;
      color: white;
    }
  </style>
</head>

<body>
  <div class="well well-sm">
    <h3>Pantum 翻译</h3>
  </div>

  <div class="container">
    <button type="button" class="btn button3"><span class="glyphicon glyphicon glyphicon-font"></span> 文字</button>
    <button type="button" class="btn button3"><span class="glyphicon glyphicon-file"></span> 文档</button>
    <br><br>
    <form>
      <div class="form-group row">
        <div class="col-xs-6">
          <ul class="nav nav-tabs">
            <li class="active"><a href="#home" data-toggle="tab">中文</a></li>
            <li><a href="#ios" data-toggle="tab">繁体</a></li>
            <li class="dropdown">
              <a href="#" id="myTabDrop1" class="dropdown-toggle" data-toggle="dropdown">英语<strong
                  class="caret"></strong></a>
              <ul class="dropdown-menu" role="menu" aria-labelledby="myTabDrop1">
                <li><a href="#jmeter" tabindex="-1" data-toggle="tab">俄语</a></li>
                <li><a href="#ejb" tabindex="-1" data-toggle="tab">法语</a></li>
              </ul>
            </li>
          </ul>
        </div>
        <div class="col-xs-6">
          <ul class="nav nav-tabs">
            <li class="active"><a href="#home" data-toggle="tab">中文</a></li>
            <li><a href="#ios" data-toggle="tab">繁体</a></li>
            <li class="dropdown">
              <a href="#" id="myTabDrop1" class="dropdown-toggle" data-toggle="dropdown">英语<strong
                  class="caret"></strong></a>
              <ul class="dropdown-menu" role="menu" aria-labelledby="myTabDrop1">
                <li><a href="#jmeter" tabindex="-1" data-toggle="tab">俄语</a></li>
                <li><a href="#ejb" tabindex="-1" data-toggle="tab">法语</a></li>
              </ul>
            </li>
          </ul>
        </div>

        <div class="col-xs-6">
          <label for="ex1">

          </label>
          <textarea class="form-control input-lg" id="ex1" type="text" rows="16"></textarea>
        </div>

        <div class="col-xs-6">
          <label for="ex2">

          </label>
          <textarea class="form-control input-lg" id="ex2" type="text" rows="16" style="min-width: 90%"></textarea>
        </div>
      </div>
    </form>

    <form enctype="multipart/form-data">
      <label for="file-0b">Test invalid input type</label>
      <div class="file-loading">
        <input id="bs-file-input" name="file" class="file" type="text" multiple>
      </div>
    </form>

    <div>
        <input class="bootstrap-filestyle" type="file" class="filestyle">
    </div>

    <!-- Footer -->
    <footer class="bg-dark text-center text-white">

      <!-- Section: Text -->
      <section class="mb-4">
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Sunt distinctio earum
          repellat quaerat voluptatibus placeat nam, commodi optio pariatur est quia magnam
          eum harum corrupti dicta, aliquam sequi voluptate quas.
        </p>
      </section>
      <!-- Section: Text -->
      <!-- Copyright -->
      <div class="text-center p-3">
        © 2021 Copyright:
        <a class="text-white" href="https://www.pantum.com/">Pantum</a>
      </div>
      <!-- Copyright -->
    </footer>
    <!-- Footer -->

    <!-- jQuery (Bootstrap 的所有 JavaScript 插件都依赖 jQuery，所以必须放在前边) -->
    <script src="/static/js/jquery.min.js"></script>
    <!-- 加载 Bootstrap 的所有 JavaScript 插件。你也可以根据需要只加载单个插件。 -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="https://cdn.bootcss.com/bootstrap-filestyle/1.2.3/bootstrap-filestyle.min.js"></script>

    <script type="text/javascript">
      $(function(){
          $('#bs-file-input').fileinput({
              language: 'zh',
              uploadUrl: '/upload-example-url',
              maxFileCount: 1,
              allowedFileExtensions: ['jpg', 'png', 'gif'],
      })
      })
      </script>
          <script>
              $(".bootstrap-filestyle").filestyle({input: false});
              $(".bootstrap-filestyle").on('change',function(){
                  alert()
              });
          </script>
</body>

</html>