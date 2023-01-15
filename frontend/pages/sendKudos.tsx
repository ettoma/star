import { Box, Button, Form, FormField, Menu, PageHeader, TextInput } from 'grommet'
import { handleKudos } from "../api/kudos/handleKudos"
import React, { useCallback, useEffect, useState } from 'react'
import { handleGetUsers } from '../api/users/handleUsers'
import { User } from '../api/models/user'



function SendKudos() {

    const [users, setUsers] = useState([])

    useEffect(() => {
        getUsers()
    }, [])

    async function getUsers() {
        const response = await (handleGetUsers()).then((response) => response.json())
        setUsers(response)
    }

    const us = [
        {
            label: "ettore"
        },
        {
            label: "test"
        }
    ]



    const [recipient, setRecipient] = useState("")
    const [message, setMessage] = useState("")

    function handleSubmit(e: React.FormEvent) {
        e.preventDefault()
        handleKudos(recipient, message)
    }

    const messages = {
        invalid: "invalid",
        required: "required"
    }


    return (
        <Box justify="center" gap="large" pad="large">
            <PageHeader title="Send Kudos" />
            <Box align="center">
                <Form onSubmit={handleSubmit} validate="submit" messages={messages}>
                    <FormField label="To">
                        //TODO fix useEffect
                        {/* <Menu items={[users.map((user) =>
                            <Button label={(user as User).username} />
                        )]} /> */}
                    </FormField>

                    <FormField label="To">
                        <TextInput onChange={(e) => setRecipient(e.target.value)} required />
                    </FormField>
                    <FormField label="Message">
                        <TextInput onChange={(e) => setMessage(e.target.value)} required />
                    </FormField>
                    <Box margin="large" direction='row' gap="small">
                        <Button label="Send" primary type="submit" onClick={() => console.log(users)} />
                        <Button label="Clear" type='reset' />
                    </Box>
                </Form>
            </Box>
        </Box>
    )
}

export default SendKudos