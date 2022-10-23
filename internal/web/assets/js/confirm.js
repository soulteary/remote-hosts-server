$(function () {
  var API = window.$API$;

  var instance = null;
  var container = document.getElementById("app");

  function InitializeEditor(rawData, newData) {
    if (rawData == "") {
      // TODO: 可以直接保存，判断新数据是否存在，有效性
      return;
    }
    if (newData == "") {
      // TODO: 提示数据有问题
      return;
    }
    var hidden = document.createElement("textarea");
    hidden.style.display = "none";
    hidden.value = newData;

    container.innerHTML = "";
    instance = CodeMirror.MergeView(container, {
      value: hidden.value,
      orig: rawData,
      lineNumbers: true,
      collapseIdentical: true,
      // connect: "align",
      mode: "hosts",
      theme: "seti",
    });
  }

  function InitializeDiffPage() {
    $.ajax({
      url: API.Compare,
      success: function (response) {
        var data = response.data;
        var prepare = response.prepare;
        // TODO: 提示初始化成功
        InitializeEditor(data, prepare);
      },
      error: function (data) {
        // TODO: 处理错误提示
        console.log(data);
      },
    });

    InitializeEditor();
  }

  InitializeDiffPage();
});
