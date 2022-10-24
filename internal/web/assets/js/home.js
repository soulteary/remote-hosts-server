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

  function Submit(data) {
    $.ajax({
      type: "POST",
      url: API.Submit,
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

  $('button#submit[data-action="submit"]').on("click", function (e) {
    e.preventDefault();
    Submit(instance.getValue());
  });

  function InitializeHomepage() {
    $.ajax({
      url: API.Data,
      success: function (response) {
        showMessage("Get the latest data successfully.");
        InitializeEditor(response);
      },
      error: function (response) {
        showMessage("Failed to get Hosts data.");
        echo(response);
      },
    });
  }

  InitializeHomepage();
});
