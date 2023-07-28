const usernameInput = document.querySelector("#user")
const userAddButton = document.querySelector("#submit")
const registeredUsers = document.querySelector(".registered")

const loadingAnimation = toString(registeredUsers.innerHTML)

userAddButton.onclick = async () => {
    const username = usernameInput.value.trim()
    if (username == "") {
        alert("Você não me disse que usuário eu devo adicionar!")
        return
    }
    userAddButton.style.cursor = "wait"
    userAddButton.disabled = true
    const url = new URL(window.location + "api/user")
    url.searchParams.append("user", username)
    let response, data;
    try {
        response = await fetch(url.toString(), { method: "POST" })
        data = await response.json()
    } catch (err) {
        alert("Error on request: " + toString(err))
    }
    if (response.status == 200) {
        loadAddedUsers()
        usernameInput.value = ""
    } else if (data) {
        alert("Falha ao adicionar usuário: " + (data?.error || response.statusText))
    }
    userAddButton.style.cursor = "default"
    userAddButton.disabled = false
}

const loadAddedUsers = async () => {
    registeredUsers.innerHTML = loadingAnimation
    const response = await fetch("/api/user")
    const data = await response.json()
    if (!data.users) {
        registeredUsers.innerHTML = "<p>Não existe nenhum usuário registrado!</p>"
        return
    }
    let html = ""
    data.users.map((user) => {
        html += `
            <div class="user">
                <p>${user.name}</p>
                <button class="delete" data-user="${user.name}">Remover</button>
            </div>
        `
    })
    registeredUsers.innerHTML = html
    addRemoveUserEventListener()
}


async function addRemoveUserEventListener() {
    const buttons = document.querySelectorAll(".delete")
    buttons.forEach((e) => e.addEventListener('click', removeUser))
}


async function removeUser() {
    const user = this.getAttribute('data-user')
    this.style.cursor = 'wait'
    this.disabled = true
    const url = new URL(window.location + 'api/user')
    url.searchParams.append('user', user)
    let response
    try {
        response = await fetch(url.toString(), { method: 'DELETE' })
    } catch (err) {
        alert("Fallha ao pedir servidor para remover usuário: " + err.toString())
        return
    }
    if (response.status == 200) {
        loadAddedUsers()
        return
    }
    const data = await response.json()
    alert("Falha ao remover usuário: " + data.error)
}

loadAddedUsers()