window.addEventListener("load", () => {
    const btns = Array.from(document.getElementsByClassName("del_btn"))
    btns.forEach(btn => {
        btn.addEventListener("click", (e) => {
            fetch(`/delete?id=${e.target.dataset.id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(response => {
                location.reload()
            }).catch(error => {
                console.error('Error:', error);
            })
        })
    })
})