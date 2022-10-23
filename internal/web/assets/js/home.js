$(function () {
  var API = {
    GET_DATA: "/api/data.txt",
    PREPARE: "/api/prepare",
  };

  var instance = null;
  function InitializeEditor(data) {
    var hidden = document.createElement("textarea");
    hidden.style.display = "none";
    hidden.value = data;

    var container = document.getElementById("app");
    container.innerHTML = "";
    instance = CodeMirror(container, {
      value: hidden.value,
      lineNumbers: true,
      matchBrackets: true,
      readOnly: false,
      mode: "hosts",
      theme: "seti",
    });
  }

  function PrepareUpdate(data) {
    $.ajax({
      type: "POST",
      url: API.PREPARE,
      data: data,
      contentType: "text/plain",
      success: function (response) {
        // TODO: 跳转下一步
        console.log(response);
      },
      error: function (response) {
        // TODO: 处理错误提示
        console.log(response);
      },
    });
  }

  $('button#submit[data-action="prepare"]').on("click", function (e) {
    e.preventDefault();
    PrepareUpdate(instance.getValue());
    console.log(instance.getValue());
  });

  function InitializeHomepage() {
    $.ajax({
      url: API.GET_DATA,
      success: function (data) {
        // TODO: 提示初始化成功
        InitializeEditor(data);
        console.log(data);
      },
      error: function (data) {
        // TODO: 处理错误提示
        console.log(data);
      },
    });
  }

  InitializeHomepage();
});
