import { useEffect, useState } from "react"
import { handleGetKudosPerUser, handleKudos } from "../api/kudos/handleKudos"
import { useNavigate, useParams } from "react-router"
import { useDispatch, useSelector } from "react-redux"
import { User } from "../api/models/user"
import { handleGetUsers } from "../api/users/handleUsers"
import { setUserList } from "../reducers/userSlice"
import { RootState } from "../src/store"
import { Box, PageHeader, FormField, TextInput, Button, Form, Menu } from "grommet"
import SendKudosModal from "../components/modal/sendKudosModal"

type Kudo = {
    sender: string,
    receiver: string,
    content: string,
    id: number
}

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
        getUsers()
        formatUsers()
        getKudosPerUser(username!)
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
        handleKudos(recipient, message, username!)
        setShow(true)
    }

    const messages = {
        invalid: "invalid",
        required: "required"
    }

    async function getKudosPerUser(recipient: string) {
        const response: Response = await (handleGetKudosPerUser(recipient))
        const res = await response.json()

        if (response.ok === true) {
            setKudos(res)
            setIsSuccess(true)
        }

        if (res["message"] == "Token is expired") {
            console.log("expired token")
            setIsSuccess(false)
            navigate("/signin")
        } else if (res["message"] == "Receiver not found") {
            console.log("receiver not found")
            setIsSuccess(false)
        }


    }

    return (

        <>
            {isSuccess === false ?
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
                    {kudos.map((k: Kudo) => <p key={k.id}>{k.content}</p>)}
                </Box>
                : <p>no</p>}
        </>
    )
}

export default MyKudos
