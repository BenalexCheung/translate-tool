<div class="col-md-6">
    <select name="sl" class="form-control" id="inlineFormSourceSelect">
        {{template "trans-map.tpl" .}}
    </select>
</div>

<div class="col-md-6">
    <select name="tl" class="form-control" id="inlineFormTargetSelect">
        {{template "trans-map.tpl" .}}
    </select>
</div>

<div class="form-outline col-md-6">
    <textarea name="text" class="form-control textareaClass" id="textAreaSrc" rows="4"></textarea>
</div>
<div class="form-outline col-md-6">
    <textarea name="trans" class="form-control textareaClass" id="textAreaTrans" rows="4"></textarea>
</div>