import { Box, Button, Form, FormField, Menu, PageHeader, TextInput } from 'grommet'
import { handleKudos } from "../api/kudos/handleKudos"
import React, { useCallback, useEffect, useState } from 'react'
import { handleGetUsers } from '../api/users/handleUsers'
import { User } from '../api/models/user'
import { useSelector } from 'react-redux'
import { RootState } from '../src/store'



function SendKudos() {

    // const [users, setUsers] = useState([])
    const [userFormat, setUserFormat] = useState([])
    const users: User[] = useSelector((state: RootState) => state.users.users)

    //TODO: implement Redux to propagate users to the whole app when they are fetched. Avoid fetching on every page

    useEffect(() => {
        // getUsers()
        formatUsers()
    }, [])

    function formatUsers() {
        const newA: {}[] = []
        users.map((user: User) => {
            const u = { label: user.username, onClick: () => setRecipient(user.username) }
            newA.push(u)
        })
        return newA
    }

    // const getUsers = useCallback(async () => {
    //     await (handleGetUsers())
    //         .then((res) => res.json())
    //         .then((data) => setUser(data))
    // }, [])



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
                        <Menu items={formatUsers()} label={recipient} />
                    </FormField>
                    <FormField label="Message">
                        <TextInput onChange={(e) => setMessage(e.target.value)} required />
                    </FormField>
                    <Box margin="large" direction='row' gap="small">
                        <Button label="Send" primary type="submit" />
                        <Button label="Clear" type='reset' onClick={() => setRecipient('')} />
                        <Button label="Clear" onClick={() => console.log(users)} />
                    </Box>
                    <Button onClick={() => formatUsers()} label="test" />
                </Form>
            </Box>
        </Box>
    )
}

export default SendKudos