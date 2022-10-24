$(function () {
  var API = window.$API$;

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

  var instance = null;
  var container = document.getElementById("app");

  function InitializeEditor(rawData, newData) {
    var hidden = document.createElement("textarea");
    hidden.style.display = "none";
    hidden.value = newData;

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

  function InitializeDiffPage() {
    $.ajax({
      url: API.Diff,
      success: function (response) {
        showMessage("Should we update?");
        InitializeEditor(response.data, response.prepare);
      },
      error: function (response) {
        showMessage("Failed to get Hosts diff data.");
        echo(response);
      },
    });

    InitializeEditor();
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

  $('button#submit[data-action="confirm"]').on("click", function (e) {
    e.preventDefault();
    Submit(instance.edit.getValue(), true);
  });

  InitializeDiffPage();
});
