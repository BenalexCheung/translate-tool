<br>

<div>
    <input id="file" name="file" class="bootstrap-filestyle" type="file" class="filestyle" data-buttonBefore="true"
        accept=".csv, .xls, .xlsx">
</div>
<br>

<div>
    <button class="btn btn-default form-submit ladda-button" type="button" data-style="expand-right" id="trans-file"><span class="ladda-label">翻译</span></button>
    <button class="btn btn-default form-submit ladda-button" type="button" data-style="expand-right" id="convert-file"><span class="ladda-label">Xls2Strings</span></button>
    <button class="btn btn-warning pull-right form-submit ladda-button" data-style="expand-right" type="button" id="load-file"><span class="ladda-label">导入翻译数据</span></button>
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
    <div class="alert alert-info">注意！导入按钮请谨慎操作，文档模板请点击<a class="text-white" href="static/upload/en.xlsx">这里</a>下载。</div>
</div>