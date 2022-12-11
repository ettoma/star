import { useEffect, useState } from "react"
import { User } from "../api/models/user"
import { handleGetUsers, handleDeleteUser } from "../api/users/handleUsers"
import UserCard from "../components/users/userCard"
import "./styles/users.css"


function Users() {

    const [users, setUsers] = useState([])

    useEffect(() => {
        getUsers()
    }, [])


    async function getUsers() {
        const response = await (handleGetUsers()).then((response) => response.json())

        setUsers(response)
    }

    async function deleteUser(id: number) {
        const response = await (handleDeleteUser(id)).then((response) => response.json())
            .finally(() => getUsers())

        console.log(response)

    }

    return (
        <div>
            <h2>Users</h2>
            <button onClick={() => getUsers()}>Get users</button>

            <div>
                {!users ? <h2>no users</h2> :
                    (users as User[]).map((user) =>
                        <div className="user-card">

                            <UserCard
                                name={user.name}
                                username={user.username}
                                createdAt={user.createdAt}
                                id={user.id}

                            />
                            <button onClick={() => deleteUser(user.id)}>X delete</button>
                        </div>
                    )}
            </div>
        </div>
    )
}

export default Users
