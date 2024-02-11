window.addEventListener("load", () => {

    let userFirstNameInput = document.getElementById("user_first_name")
    let userSecondNameInput = document.getElementById("user_second_name")
    let userPasswordInput = document.getElementById("user_password")
    let userConfirmPasswordInput = document.getElementById("user_confirm_password")
    let userEmailInput = document.getElementById("user_email")


    let registerButton = document.getElementById("register_button")

    userFirstNameInput.addEventListener('input', validateForm);
    userSecondNameInput.addEventListener('input', validateForm);
    userPasswordInput.addEventListener('input', validateForm);
    userConfirmPasswordInput.addEventListener('input', validateForm);
    userEmailInput.addEventListener('input', validateForm);

    function validateForm() {
        // Проверяем, заполнены ли обязательные поля
        var userFirstNameValid = userFirstNameInput.value.trim() !== '' && userFirstNameInput.value.length > 2 && userFirstNameInput.value.length < 16;
        var userSecondNameValid = userSecondNameInput.value.trim() !== '' && userSecondNameInput.value.length > 2 && userSecondNameInput.value.length < 16;
        var passwordValid = userPasswordInput.value.trim() !== '' && userPasswordInput.value.length > 8 && userPasswordInput.value.length < 16;
        var userConfirmPasswordValid = userConfirmPasswordInput.value.trim() !== '' && userConfirmPasswordInput.value.length > 8 && userConfirmPasswordInput.value.length < 16;
        var userEmailValid = userEmailInput.value.trim() !== '' && userEmailInput.value.length > 8 && userEmailInput.value.length < 30;

        // Если оба поля заполнены, делаем кнопку доступной, иначе - недоступной
        registerButton.disabled = !(userFirstNameValid && userSecondNameValid && passwordValid && userConfirmPasswordValid && userEmailValid);
    }

    registerButton.addEventListener("click", (event) => {
        event.preventDefault()

        let password = document.getElementById("user_password").value
        let confirmPassword = document.getElementById("user_confirm_password").value

        if (password !== confirmPassword) {
            alert("Passwords don't match!")
        }

        const newUser = {
            full_name: userFirstNameInput.value.trim() + " " + userSecondNameInput.value.trim(),
            email: userEmailInput.value.trim(),
            role: document.getElementById("user_role").value,
            password: userPasswordInput.value.trim(),
            approved: document.getElementById("user_role").value === 'student'
        }
        console.log(newUser)

        fetch(`/register_page/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(newUser)
        }).then(response => {
            if (response.status === 200) {
                alert(`User: ${newUser.full_name} was successfully registered!`)
                window.location.replace("/login_page")
            }
        }).catch(error => {
            console.log(error)
        })
    })
})