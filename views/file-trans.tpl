<br>

<div>
    <input id="file" name="file" class="bootstrap-filestyle" type="file" class="filestyle" data-buttonBefore="true"
        accept=".csv, .xls, .xlsx">
</div>
<br>

<div>
    <button class="btn btn-default" type="button" onclick="transFile()">翻译</button>
    <button class="btn btn-default" type="button" onclick="convertFile()">格式转换</button>
    <button class="btn btn-default pull-right" type="button" onclick="loadFile()">导入翻译数据</button>

</div>
<br>

<div class="col-12">
    <ul class="list-group">
    </ul>
</div>

<div class="col-12">
    <div class="alert alert-success">① 翻译说明：在Excel表格的A列填入你所需要翻译的字段，其他列保持为空即可</div>
    <div class="alert alert-success">② 格式转换说明：将Excel表格的A列和B列内容转换成键值对应的Localizable.strings</div>
    <div class="alert alert-success">③ 导入翻译说明：Excel表格的A列为主键（英文），B列为对应翻译内容（其他语言），不同的本地化语言由文件名进行标识，如：en.xlsx</div>
    <div class="alert alert-info">注意！文档模板请点击<a class="text-white" href="static/upload/en.xlsx">这里</a>下载。</div>

</div>