<script src="https://cdn.bootcss.com/bootstrap-filestyle/1.2.3/bootstrap-filestyle.min.js"></script>
<script>
    function download(href, filename = '') {
        var ul = document.querySelector("ul");
        const a = document.createElement('a')
        a.download = filename
        a.href = href
        a.classList.add("btn");
        a.classList.add("btn-success");
        a.classList.add("btn-sm");
        a.classList.add("glyphicon");
        a.classList.add("glyphicon-download-alt");
        a.classList.add("pull-right");

        var li = document.createElement("li");
        // Add Bootstrap class to the list element
        li.classList.add("list-group-item");
        li.appendChild(document.createTextNode(filename));
        li.appendChild(a);
        ul.appendChild(li);
    }
</script>
<script>
    $("#label1").click(function (event) { document.getElementById("trans").checked = true; $("#postForm").submit() });
    $("#label2").click(function (event) { document.getElementById("docs").checked = true; $("#postForm").submit() });
    /**
    *
    * @param fn {Function}   实际要执行的函数
    * @param delay {Number}  延迟时间，也就是阈值，单位是毫秒（ms）
    *
    * @return {Function}     返回一个“去弹跳”了的函数
    */
    function debounce(fn, delay) {
        // 定时器，用来 setTimeout
        var timer
        // 返回一个函数，这个函数会在一个时间区间结束后的 delay 毫秒时执行 fn 函数
        return function () {
            // 保存函数调用时的上下文和参数，传递给 fn
            var context = this
            var args = arguments
            // 每次这个返回的函数被调用，就清除定时器，以保证不执行 fn
            clearTimeout(timer)
            // 当返回的函数被最后一次调用后（也就是用户停止了某个连续的操作），
            // 再过 delay 毫秒就执行 fn
            timer = setTimeout(function () {
                fn.apply(context, args)
            }, delay)
        }
    }

    //文本域自适应
    function autoTextAreaHeight() {
        let obj = $('textarea'),
            len = obj.length;
        for (let i = 0; i < len; i++) {
            obj[i].style.height = 'auto';
            obj[i].style.height = obj[i].scrollTop + obj[i].scrollHeight + 'px';
        }
    }

    var filePath = ""
    var fileName = ""
    function subFile() {
        // var fd = new FormData($('#postForm')[0]);  
        // var fd = $('#postForm').serialize()
        var fd = new FormData($('#postForm')[0]);
        console.info(fd);
        $.ajax({
            url: window.location.href.split("?")[0],
            type: "POST",
            data: fd,
            processData: false,
            contentType: false,
            success: function (data) {
                console.info(data);
                filePath = data.file_path
                fileName = data.file_name
            },
            error: function (data) {
                console.warn(data);
            }
        });
        return false;
    }

    function subText() {
        var fd = $('#postForm').serialize()
        console.info(fd);
        $.ajax({
            url: window.location.href.split("?")[0],
            type: "POST",
            data: $('#postForm').serialize(),
            success: function (data) {
                console.info(data);
                if (data.sentences) {
                    $('#textAreaTrans').empty();
                    data.sentences.forEach(element => {
                        $('#textAreaTrans').append(element.trans);
                    });
                } else {
                    $('#textAreaTrans').empty();
                }
            },
            error: function (data) {
                console.warn(data);
            }
        });
        return false;
    }

    $('.textareaClass').on('keyup', debounce(function (e) {
        var state = { title: '', url: window.location.href.split("?")[0] };
        history.pushState(state, 0, '?' + $('#postForm').serialize());
        subText()
        autoTextAreaHeight()
    }, 500));

    $('#inlineFormTargetSelect').on('click', debounce(function (e) {
        var state = { title: '', url: window.location.href.split("?")[0] };
        history.pushState(state, 0, '?' + $('#postForm').serialize());
        subText()
        autoTextAreaHeight()
    }, 500));

    $(".bootstrap-filestyle").filestyle({ buttonBefore: true });
    $(".bootstrap-filestyle").on('change', function () {
        subFile()
    });
</script>
<script>
    function transFile() {
        if (filePath == "") {
            alert("请先选择文件！")
            return
        }
        var fd = new FormData();
        fd.append("file_name", fileName);
        fd.append("file_path", filePath);
        fd.append("op", "translate");
        $.ajax({
            url: window.location.href.split("?")[0] + "convert",
            type: "POST",
            data: fd,
            processData: false,
            contentType: false,
            success: function (data) {
                console.info(data);
                download(data.file_path, data.file_name)
            },
            error: function (data) {
                console.warn(data);
            }
        });
        return false;
    }
    function convertFile() {
        if (filePath == "") {
            alert("请先选择文件！")
            return
        }
        var fd = new FormData();
        fd.append("file_name", fileName);
        fd.append("file_path", filePath);
        fd.append("op", "convert");
        $.ajax({
            url: window.location.href.split("?")[0] + "convert",
            type: "POST",
            data: fd,
            processData: false,
            contentType: false,
            success: function (data) {
                console.info(data);
                download(data.file_path, data.file_name)
            },
            error: function (data) {
                console.warn(data);
            }
        });
        return false;
    }
    function loadFile() {
        if (filePath == "") {
            alert("请先选择文件！")
            return
        }
        var fd = new FormData();
        fd.append("file_name", fileName);
        fd.append("file_path", filePath);
        fd.append("op", "loading");
        $.ajax({
            url: window.location.href.split("?")[0] + "convert",
            type: "POST",
            data: fd,
            processData: false,
            contentType: false,
            success: function (data) {
                console.info(data);
            },
            error: function (data) {
                console.warn(data);
            }
        });
        return false;
    }
</script>