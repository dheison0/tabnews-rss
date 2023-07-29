const usernameInput = document.querySelector("#user")
const userAddButton = document.querySelector("#submit")
const registeredUsers = document.querySelector(".registered-users")
const loadingAnimation = registeredUsers.innerHTML.toString()

function generateURL(path, params) {
  const url = new URL(window.location + path)
  for (let key in params) {
    url.searchParams.append(key, params[key])
  }
  return url.toString()
}

function elementWaitToggle(e) {
  e.disabled = !e.disabled
  e.style.cursor = e.disabled ? "wait" : "default"
}

async function addRemoveUserEventListener() {
  const buttons = document.querySelectorAll(".delete-user")
  buttons.forEach((e) => e.addEventListener('click', removeUser))
}

async function loadAddedUsers() {
  registeredUsers.innerHTML = loadingAnimation
  const response = await fetch("/api/user")
  let data
  try {
    data = await response.json()
  } catch (err) {
    registeredUsers.innerHTML = `
      <div class="center error-box">
        <h4>Não foi possível carregar a lista de usuários!</h4>
        <button class="button" onclick="loadAddedUsers()">Tentar novamente</button>
      </div>
    `
    return
  }
  if (!data.users) {
    registeredUsers.innerHTML = "<p>Não existe nenhum usuário registrado!</p>"
    return
  }
  const items = data.users.map((user) => `
    <div class="user">
      <p>${user.name}</p>
      <button class="button delete-user" data-user="${user.name}">Remover</button>
    </div>
  `)
  registeredUsers.innerHTML = items.join("\n")
  addRemoveUserEventListener()
}

async function removeUser() {
  const user = this.getAttribute('data-user')
  elementWaitToggle(this)
  const url = generateURL('api/user', { user })
  let response
  try {
    response = await fetch(url, { method: 'DELETE' })
  } catch (err) {
    elementWaitToggle(this)
    return alert("Fallha ao pedir servidor para remover usuário: " + err.toString())
  }
  if (response.status == 200) {
    return loadAddedUsers()
  }
  const data = await response.json()
  alert("Falha ao remover usuário: " + data.error)
  elementWaitToggle(this)
}

userAddButton.onclick = async () => {
  const username = usernameInput.value.trim()
  if (username == "") {
    return alert("Você não me disse que usuário eu devo adicionar!")
  }
  elementWaitToggle(userAddButton)
  const url = generateURL("api/user", { user: username })
  let response, data;
  try {
    response = await fetch(url, { method: "POST" })
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
  elementWaitToggle(userAddButton)
}

loadAddedUsers()
