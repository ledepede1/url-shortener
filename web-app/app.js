function CreateSubmitHandler() {

    event.preventDefault();

    const url = event.target.elements["shorturl-url"].value;

    fetch("http://localhost:8080/backend/create", {
        method: "POST",
        body: JSON.stringify({
            url: url,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    });

    return false;
}