import "react"
import React, { useState } from "react"
import "./styles/signIn.css"

function Register() {
    const [name, setName] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    function handleSubmit(e: React.FormEvent) {
        e.preventDefault()
        console.log(name, username, password)
    }

    return (
        <main>
            <section className="card">
                <h3>New User</h3>
                <p>Create a new user</p>
                <form onSubmit={handleSubmit}>
                    <input type="text" name="name" placeholder="Enter your full name" onChange={e => setName(e.target.value)} />
                    <input type="text" name="username" placeholder="Enter username" onChange={e => setUsername(e.target.value)} />
                    <input type="password" name="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />

                    <button type="submit">Register</button>

                </form>
            </section>
        </main>
    )
}

export default Register