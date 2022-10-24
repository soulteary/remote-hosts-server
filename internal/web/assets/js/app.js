$(function () {
  var API = window.$API$;
  var instance = null;

  function showMessage(message) {
    var container = $(".submit-result");
    container.removeClass("show hide").text(message).show().addClass("show");
    setTimeout(function () {
      container.removeClass("show").addClass("hide");
    }, 3000);
  }

  function echo(message) {
    if (console && console.log) {
      console.log(message);
    }
  }

  function InitializeSimpleEditor(data) {
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

  function InitializeDiffEditor(rawData, newData) {
    var hidden = document.createElement("textarea");
    hidden.style.display = "none";
    hidden.value = newData;
    var container = document.getElementById("app");
    container.innerHTML = "";
    instance = CodeMirror.MergeView(container, {
      value: hidden.value,
      orig: rawData,
      lineNumbers: true,
      collapseIdentical: true,
      mode: "hosts",
      theme: "seti",
    });
  }

  function Submit(data, confirmReview) {
    $.ajax({
      type: "POST",
      url: API.Submit + (confirmReview ? "?confirm=ok" : ""),
      data: data,
      contentType: "text/plain",
      success: function (response) {
        if (!response) {
          showMessage("The server did not respond correctly");
          return;
        }
        if (response.code == 0 && response.next) {
          location.href = response.next;
        } else {
          if (response.message) {
            showMessage(response.message);
          }
        }
      },
      error: function (response) {
        showMessage("Failed to update Hosts data.");
        echo(response);
      },
    });
  }

  function InitializeHomepage() {
    $.ajax({
      url: API.Data,
      success: function (response) {
        showMessage("Get the latest data successfully.");
        InitializeSimpleEditor(response);
      },
      error: function (response) {
        showMessage("Failed to get Hosts data.");
        echo(response);
      },
    });
  }
  function InitializeDiffPage() {
    $.ajax({
      url: API.Diff,
      success: function (response) {
        showMessage("Should we update the original configuration?");
        InitializeDiffEditor(response.data, response.prepare);
      },
      error: function (response) {
        showMessage("Failed to get Hosts diff data.");
        echo(response);
      },
    });
  }

  var btnSubmit = $('button#submit[data-action="submit"]');
  if (btnSubmit.length) {
    btnSubmit.on("click", function (e) {
      e.preventDefault();
      Submit(instance.getValue());
    });
    InitializeHomepage();
  }

  var btnConfirm = $('button#submit[data-action="confirm"]');
  if (btnConfirm.length) {
    btnConfirm.on("click", function (e) {
      e.preventDefault();
      Submit(instance.edit.getValue(), true);
    });
    InitializeDiffPage();
  }
});
