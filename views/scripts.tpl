<script src="https://cdn.bootcss.com/bootstrap-filestyle/1.2.3/bootstrap-filestyle.min.js"></script>
<script>
    $(".bootstrap-filestyle").filestyle({ input: false });
    $(".bootstrap-filestyle").on('change', function () {
        alert()
    });
</script>

<script>
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

    function sub() {
        $.ajax({
            url: window.location.href.split("?")[0],
            type: "POST",
            data: $('#postForm').serialize(),
            success: function (data) {
                console.info(data);
                if (data.sentences) {
                    $('#textAreaTrans').empty();
                    data.sentences.forEach(element => {
                     console.info("data");
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
        sub()
        autoTextAreaHeight()
    }, 500));

    $('#inlineFormTargetSelect').on('click', debounce(function (e) {
        var state = { title: '', url: window.location.href.split("?")[0] };
        history.pushState(state, 0, '?' + $('#postForm').serialize());
        sub()
        autoTextAreaHeight()
    }, 500));
</script>