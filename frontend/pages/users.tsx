import React, { useState } from "react"
import User from "../api/models/user"
import { handleGetUsers, handleDeleteUser } from "../api/users/handleUsers"
import UserCard from "../components/users/userCard"


function Users() {

    const [users, setUsers] = useState([])

    async function getUsers() {
        const response = await (await handleGetUsers()).json()

        setUsers(response)
    }

    async function deleteUser(id: number) {
        const response = await (handleDeleteUser(id)).then((response) => response.json())
            .finally(() => getUsers())

        console.log(response)

    }

    return (

        <>
            <h2>Users</h2>
            <button onClick={() => getUsers()}>Get users</button>

            <ul>
                {(users as User[]).map((user) =>
                    <div>
                        <UserCard
                            name={user.name}
                            username={user.username}
                            password={user.password}
                            createdAt={user.createdAt}
                            id={user.id}
                            key={user.id}
                        />
                        <button onClick={() => deleteUser(user.id)}>delete</button>
                    </div>
                )}
            </ul>
        </>
    )
}

export default Users
