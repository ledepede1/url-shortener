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

            alert("Created link: localhost:8080/"+url)
        }
    } else {
        alert("You cannot leave the input box empty!")
    }

    return false;
}