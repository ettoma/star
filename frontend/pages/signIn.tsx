import React, { useState } from "react"
import { useNavigate } from "react-router"
import LoginRequestData from "../api/models/errors"
import { handleLogin } from "../api/users/handleUsers"
import { Box, Button, Form, FormField, PageHeader, TextInput } from "grommet"

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


    if (data.success) {
      document.cookie = "token=" + data.token + "; secure; sameSite=Lax;"
      navigate("/users")
    }
  }

  const messages = {
    invalid: "invalid",
    required: "required"
  }


  return (
    <Box justify="center" gap="large" pad="large">
      <PageHeader title="Login" />
      <Box align="center">
        <Form onSubmit={handleSubmit} method="POST" validate="submit" messages={messages}>
          <FormField label="Username" name="username" >
            <TextInput type="text" name="username" onChange={e => setUsername(e.target.value)} required />
          </FormField>
          <FormField label="Password" name="password" required>
            <TextInput type="password" name="password" onChange={e => setPassword(e.target.value)} required />
          </FormField>
          <Box margin="large">
            <Button type="submit" primary label="Sign In" />
          </Box>
        </Form>
      </Box>
    </Box>
  )
}

export default SignIn