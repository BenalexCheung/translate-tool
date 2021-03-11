<div>
<button class="btn btn-default" type="button" onclick="loadFile()">导入翻译数据</button>
</div>
<br>

<div>
    <input id="file" name="file" class="bootstrap-filestyle" type="file" class="filestyle" data-buttonBefore="true"
        accept=".csv, .xls, .xlsx">
</div>
<br>

<div>
<button class="btn btn-default" type="button" onclick="transFile()">翻译</button>
<button class="btn btn-default" type="button" onclick="convertFile()">格式转换</button>
</div>
<br>

<div class="col-12">
    <ul class="list-group">
    </ul>
  </div>