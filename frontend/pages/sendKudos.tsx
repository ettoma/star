import { Box, Button, Form, FormField, Layer, Menu, PageHeader, Tag, TextInput } from 'grommet'
import { handleKudos } from "../api/kudos/handleKudos"
import React, { useEffect, useState } from 'react'
import { handleGetUsers } from '../api/users/handleUsers'
import { User } from '../api/models/user'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../src/store'
import { setUserList } from '../reducers/userSlice'
import SendKudosModal from '../components/modal/sendKudosModal'



function SendKudos() {
    const dispatch = useDispatch()
    const users: User[] = useSelector((state: RootState) => state.users.users)

    const [recipient, setRecipient] = useState("")
    const [message, setMessage] = useState("")

    const [show, setShow] = useState(false)


    useEffect(() => {
        getUsers()
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

    async function getUsers() {
        const response = await (handleGetUsers())
            .then((response) => response.json())
            .catch((error) => console.log(error))

        dispatch(setUserList(response))
    }

    function handleSubmit(e: React.FormEvent) {
        e.preventDefault()
        handleKudos(recipient, message)
        setShow(true)
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
                        <TextInput onChange={(e) => setMessage(e.target.value)} />
                    </FormField>
                    <Box margin="large" direction='row' gap="small">
                        <Button label="Send" primary type="submit" />
                        {show && (
                            <SendKudosModal setShow={setShow} message={message} recipient={recipient} />
                        )}
                        <Button label="Clear" type='reset' onClick={() => {
                            setRecipient('')
                            setMessage('')
                        }} />
                    </Box>
                </Form>
            </Box>
        </Box>
    )
}

export default SendKudos