import "react"
import React, { useState } from "react"
import { handleRegister } from "../api/users/handleUsers"
import RequestError from "../api/models/errors"
import "./styles/signIn.css"
import { useNavigate } from "react-router-dom"

function Register() {
    const [name, setName] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const navigate = useNavigate()

    async function handleSubmit(e: React.FormEvent) {
        e.preventDefault()
        const response = await handleRegister({ "name": name, "username": username, "password": password })
        const data = await response.json()

        if ((response.status) != 201) {
            console.log("error creating user")
            console.log(data as RequestError)
        }

        if ((response.status) == 201) {
            navigate("/users")
        }
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