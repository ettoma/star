import React, { useState } from "react"
import { useNavigate } from "react-router"
import LoginRequestData from "../api/models/errors"
import { handleLogin } from "../api/users/handleUsers"
import { Anchor, PageHeader } from "grommet"
// import "./styles/signIn.css"

function SignIn() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const navigate = useNavigate()


  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    const token = document.cookie.split("=")[1]
    const response = await handleLogin({ "username": username, "password": password, "token": token })
    var data = await response.json() as LoginRequestData

    if (data.success != true) {
      console.log("error creating user")
      console.log(data.message)
    }

    //!!! fix the HttpOnly logic
    if (data.success) {
      document.cookie = "token=" + data.token + "; HttpOnly; secure; sameSite=Lax;"
      console.log(data.token)
      navigate("/users")
    }
  }


  return (
    <main>
      <div className="card">
        <PageHeader title="Login" parent={<Anchor label="home" onClick={() => navigate('/')} />} />


        <form className="form" onSubmit={handleSubmit} method="POST">
          <label>Username</label>
          <input type="text" name="username" placeholder="Enter username" onChange={e => setUsername(e.target.value)} />
          <label>Password</label>
          <input type="password" name="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
          <div className="button_group">
            <button type="submit" id="signin">Sign in</button>
            {/* <Link to={"/register"}>
              <button>--- register ---</button>
            </Link> */}
          </div>
        </form>
      </div>
    </main>
  )
}

export default SignIn