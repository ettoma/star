import "react"
import React, { useState } from "react"
import "./styles/signIn.css"

function SignIn() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    console.log(username, password)
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

          <button>--- register ---</button>
        </form>
      </section>
    </main>
  )
}

export default SignIn