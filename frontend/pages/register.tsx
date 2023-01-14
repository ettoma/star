import "react"
import React, { useState } from "react"
import { handleRegister } from "../api/users/handleUsers"
import RequestData from "../api/models/errors"
// import "./styles/signIn.css"
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
        console.log((data as RequestData).status)

        // if ((response.status) != 201) {
        //     console.log("error creating user")
        //     console.log(data as RequestData)
        // }

        if ((response.status) == 201) {
            navigate("/users")
        }
    }

    return (
        <main>
            <div className="card">
                <div className="card_title">
                    <h1>New User</h1>
                    <p className="card_subtitle">Create a new user</p>
                </div>
                <form className="form" onSubmit={handleSubmit} method="POST">
                    <label>Full name</label>
                    <input type="text" name="name" placeholder="Enter your full name" onChange={e => setName(e.target.value)} />
                    <label>Username</label>
                    <input type="text" name="username" placeholder="Enter username" onChange={e => setUsername(e.target.value)} />
                    <label>Password</label>
                    <input type="password" name="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
                    <div className="button_group">

                        <button id="signin" type="submit">Register</button>
                    </div>

                </form>
            </div>
        </main>
    )
}

export default Register