function create() {
    const title = document.getElementById("title").value
    const author = document.getElementById("author").value
    const category = document.getElementById("category").value
    const readership = document.getElementById("readership").value
    const description = document.getElementById("description").value
    const content = document.getElementById("content").value


    const article = {
        title: title,
        author: author,
        category: category,
        readership: readership,
        description: description,
        content: content,
    }
    console.log(article)

    fetch(`/add/create`, {
        method: "POST",
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