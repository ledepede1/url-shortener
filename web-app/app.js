async function CreateSubmitHandler() {

    event.preventDefault();

    const url = event.target.elements["shorturl-url"].value;

    if (url.length >= 1) { 
        const respone = await fetch("http://localhost:8080/backend/create", {
            method: "POST",
            body: JSON.stringify({
                url: url,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
        });

        if (respone.ok) {
            const data = await respone.json();
            const url = data.result;

            showCopyTextbox("localhost:8080/"+url);
        }
    } else {
        alert("You cannot leave the input box empty!")
    }

    return false;
}

function showCopyTextbox(text) {
    document.body.innerHTML = "";

    const container = document.createElement('div');
    container.className = 'copy-container';

    const textbox = document.createElement('input');
    textbox.type = 'text';
    textbox.value = text;
    textbox.readOnly = true;

    const copyButton = document.createElement('div');
    copyButton.className = 'copybutton';
    copyButton.textContent = '📋';

    copyButton.addEventListener('click', function () {
      navigator.clipboard.writeText(text).then(function () {
        alert('Text has been copied to the clipboard');
      }).catch(function (err) {
        console.error('Unable to copy text', err);
      });
    });

    container.appendChild(textbox);
    container.appendChild(copyButton);

    document.body.appendChild(container);
  }