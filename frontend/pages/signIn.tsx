import "react"
import React, { useState } from "react"
import { Route, Router, useNavigate } from "react-router"
import { Link } from "react-router-dom"
import RequestError from "../api/models/errors"
import { handleLogin } from "../api/users/handleUsers"
import "./styles/signIn.css"

function SignIn() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const navigate = useNavigate()


  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    const response = await handleLogin({ "username": username, "password": password })
    const data = await response.json()

    if ((response.status) != 200) {
      console.log("error creating user")
      console.log(data as RequestError)
    }

    if ((response.status) == 200) {
      navigate("/users")
    }
  }


  return (
    <main>
      <section className="card">
        <h3>Login</h3>
        <p>Enter your login details to get sign in to your account</p>
        <form onSubmit={handleSubmit}>
          <input type="text" name="username" placeholder="Enter username" onChange={e => setUsername(e.target.value)} />
          <input type="password" name="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />

          <button type="submit">Sign in</button>

          <Link to={"/register"}>
            <button>--- register ---</button>
          </Link>
        </form>
      </section>
    </main>
  )
}

export default SignIn