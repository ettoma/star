import React, { useState } from "react"
import { useNavigate } from "react-router"
import LoginRequestData from "../api/models/errors"
import { handleLogin } from "../api/users/handleUsers"
import ErrorModal from "../components/modal/loginErrorModal"
import { Box, Button, Form, FormField, PageHeader, TextInput } from "grommet"

function SignIn() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [showErrorModal, setShowErrorModal] = useState(false)
  const [errorMessage, setErrorMessage] = useState("")
  const navigate = useNavigate()

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    const response = await handleLogin({ "username": username, "password": password })
    var data = await response.json() as LoginRequestData

    document.cookie = "token=" + data.token
    document.cookie = "refresh_token=" + data.refreshToken

    if (data.success != true) {
      setShowErrorModal(true)
      setErrorMessage(data.message)
    }


    if (data.success) {
      navigate(`/kudos/${username}`)
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
            {showErrorModal && (
              <ErrorModal setShowErrorModal={setShowErrorModal} errorMessage={errorMessage} />
            )}
          </Box>
        </Form>
      </Box>
    </Box>
  )
}

export default SignIn