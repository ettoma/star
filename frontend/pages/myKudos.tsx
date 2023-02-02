import React, { useEffect, useState } from "react"
import { handleGetKudosPerUser, handleSendKudos } from "../api/kudos/handleKudos"
import { useNavigate, useParams } from "react-router"
import { useDispatch, useSelector } from "react-redux"
import { User } from "../api/models/user"
import { Kudo } from "../api/models/kudos"
import { handleGetUsers } from "../api/users/handleUsers"
import { unixToDate } from "../utils/parser"
import { setUserList } from "../reducers/userSlice"
import { RootState } from "../src/store"
import { Box, PageHeader, FormField, TextInput, Button, Form, Menu, Spinner, Carousel, Card, Text, Header, NameValueList, NameValuePair, Grid, Paragraph } from "grommet"
import SendKudosModal from "../components/modal/sendKudosModal"



function MyKudos() {
    const [kudos, setKudos] = useState([])

    const [isSuccess, setIsSuccess] = useState(false)

    const navigate = useNavigate()

    const { username } = useParams()

    const dispatch = useDispatch()
    const users: User[] = useSelector((state: RootState) => state.users.users)

    const [recipient, setRecipient] = useState("")
    const [message, setMessage] = useState("")

    const [show, setShow] = useState(false)


    useEffect(() => {
        getKudosPerUser(username!)
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
        const response = handleSendKudos({ receiver: recipient, content: message, sender: username! })
            .then((response) => response.json())
            .catch((error) => console.log(error))
        setShow(true)
    }

    function clear() {
        setMessage("")
        setRecipient("")
    }

    const messages = {
        invalid: "invalid",
        required: "required"
    }

    async function getKudosPerUser(receiver: string) {
        const response: Response = await (handleGetKudosPerUser(receiver))
        const res = await response.json()

        if (response.status == 200) {
            setKudos(res)
            setIsSuccess(true)
        } else if (res["message"] == "Token is expired" || res["message"] == "Token not provided") {
            console.error("expired or blank token")
            setIsSuccess(false)
            navigate("/signin")
        } else if (res["message"] == "Receiver not found") {
            console.log("receiver not found")
            setIsSuccess(false)
        }


    }

    return (

        <>
            <Box justify="center" gap="large" pad="large">
                <PageHeader title={"Hi, " + username} />
                {isSuccess === true ?
                    <Box align="center">
                        <Form onSubmit={handleSubmit} validate="submit" messages={messages}>
                            <FormField label="To">
                                <Menu items={formatUsers()} label={recipient} />
                            </FormField>
                            <FormField label="Message">
                                <TextInput onChange={(e) => setMessage(e.target.value)} />
                            </FormField>
                            <Box margin="large" direction='row' gap="small">
                                <Button label="Send" primary type="submit" onClick={handleSubmit} />
                                {show && (
                                    <SendKudosModal setShow={setShow} clear={clear} message={message} recipient={recipient} />
                                )}
                                <Button label="Clear" type='reset' onClick={() => {
                                    setRecipient('')
                                    setMessage('')
                                }} />
                            </Box>
                        </Form>
                        <Box height="small" width="medium" overflow="hidden">
                            <Carousel alignSelf="center" fill controls="selectors">
                                {kudos.length === 0 ? <Card>No kudos yet</Card> :
                                    kudos.map((k: Kudo) =>
                                        <Box key={k.id}
                                            background={{
                                                color: "#4B4B4B"
                                            }}
                                            round
                                            pad="medium"
                                            wrap>
                                            <Grid
                                                fill
                                                rows={['xxsmall', 'xsmall']}
                                                columns={['xsmall', 'small']}
                                                gap="small"
                                                areas={[
                                                    { name: 'date', start: [0, 0], end: [1, 0] },
                                                    { name: 'sender', start: [0, 1], end: [0, 1] },
                                                    { name: 'message', start: [1, 1], end: [1, 1] },
                                                ]}
                                            >
                                                <Box gridArea="date" round pad={{
                                                    horizontal: "medium",
                                                    vertical: "medium"
                                                }}>{unixToDate(k.timestamp * 1000)}</Box>
                                                <Box gridArea="sender" background="light-5" round>
                                                    <NameValueList pad={{
                                                        horizontal: "medium",
                                                        vertical: "small"
                                                    }}>
                                                        <NameValuePair name="sender">
                                                            <Text>{k.sender}</Text>
                                                        </NameValuePair>
                                                    </NameValueList>
                                                </Box>
                                                <Box gridArea="message" background="light-2" round>
                                                    <NameValueList pad={{
                                                        horizontal: "medium",
                                                        vertical: "small"
                                                    }}>
                                                        <NameValuePair name="message">
                                                            <Text>{k.content}</Text>
                                                        </NameValuePair>
                                                    </NameValueList>
                                                </Box>
                                            </Grid>
                                        </Box>)
                                }
                            </Carousel>
                        </Box>
                    </Box>
                    :
                    <Box align="center">
                        <Spinner size="xlarge" />
                    </Box>}
            </Box>
        </>
    )
}


export default MyKudos