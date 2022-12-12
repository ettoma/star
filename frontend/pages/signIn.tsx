import React, { useState } from "react"
import { useNavigate } from "react-router"
import { Link } from "react-router-dom"
import RequestData from "../api/models/errors"
import { handleLogin } from "../api/users/handleUsers"
import "./styles/signIn.css"

function SignIn() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const navigate = useNavigate()


  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    const response = await handleLogin({ "username": username, "password": password })
    var data = await response.json() as RequestData

    if (data.success != true) {
      console.log("error creating user")
      console.log(data.message)
    }

    if (data.success) {
      navigate("/users")
    }
  }


  return (
    <main>
      <div className="card">
        <div className="card_title">
          <h1>Login</h1>
          <p className="card_subtitle">Enter your login details to get sign in to your account</p>
        </div>
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