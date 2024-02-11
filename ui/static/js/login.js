window.addEventListener("load", () => {

    let userPasswordInput = document.getElementById("user_password")
    let userEmailInput = document.getElementById("user_email")


    let loginButton = document.getElementById("login_button")

    userPasswordInput.addEventListener('input', validateForm);
    userEmailInput.addEventListener('input', validateForm);

    function validateForm() {
        var passwordValid = userPasswordInput.value.trim() !== '' && userPasswordInput.value.length < 16;
        var userEmailValid = userEmailInput.value.trim() !== '' && userEmailInput.value.length < 30;
        loginButton.disabled = !(passwordValid && userEmailValid);
    }

    loginButton.addEventListener("click", (event) => {
        event.preventDefault()

        const loginUser = {
            email: userEmailInput.value.trim(),
            password: userPasswordInput.value.trim()
        }
        console.log(loginUser)

        fetch(`/login_page/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(loginUser)
        })
            .then(response => {
                if (!response.ok) {
                    if (response.status === 400) {
                        alert("Incorrect data or credentials");
                    } else {
                        console.log("Server returned an error:", response.status);
                    }
                }
            })
            .catch(error => {
                console.error('There was a problem with the fetch operation:', error);
            });


    })
})