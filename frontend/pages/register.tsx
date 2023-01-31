import "react"
import React, { useState } from "react"
import { handleRegister } from "../api/users/handleUsers"
import RequestData from "../api/models/errors"
import { useNavigate } from "react-router-dom"
import { Box, Button, Form, FormField, PageHeader, TextInput } from "grommet"

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

        //TODO : handle errors with modal
        // if ((response.status) != 201) {
        //     console.log("error creating user")
        //     console.log(data as RequestData)
        // }

        if ((response.status) == 201) {
            navigate("/users")
        }
    }

    const messages = {
        invalid: "invalid",
        required: "required"
    }

    return (
        <Box justify="center" gap="large" pad="large">
            <PageHeader title="Register" />
            <Box align="center">
                <Form onSubmit={handleSubmit} method="POST" validate="submit" messages={messages}>
                    <FormField label="Full name" name="fullname" required>
                        <TextInput type="text" name="fullname" onChange={e => setName(e.target.value)} />
                    </FormField>
                    <FormField label="Username" name="username" required>
                        <TextInput type="text" name="username" onChange={e => setUsername(e.target.value)} />
                    </FormField>
                    <FormField label="Password" name="password" required>
                        <TextInput type="password" name="password" onChange={e => setPassword(e.target.value)} />
                    </FormField>
                    <Box margin="large">
                        <Button type="submit" label="Register" primary />
                    </Box>
                </Form>
            </Box>
        </Box>
    )
}

export default Register