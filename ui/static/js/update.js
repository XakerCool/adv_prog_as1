function send() {
    const urlParams = new URLSearchParams(window.location.search)

    const id = urlParams.get('id')
    const title = document.getElementById("title").value
    const author = document.getElementById("author").value
    const category = document.getElementById("category").value
    const readership = document.getElementById("readership").value
    const description = document.getElementById("description").value
    const content = document.getElementById("content").value


    const article = {
        id: id,
        title: title,
        author: author,
        category: category,
        readership: readership,
        description: description,
        content: content,
        created: new Date()
    }
    console.log(article)
    fetch(`/update?id=${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(article)
    }).then(response => {
        console.log(response)
    }).catch(error => {
        console.log(error)
    })
}