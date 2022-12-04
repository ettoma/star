import React from "react";
import User from "../../api/models/user";

function UserCard({ name, username, password, createdAt, id }: User): JSX.Element {
    return (
        <div>
            <h3>{name}</h3>
            <ul>
                <li>ID: {id}</li>
                <li>Username: {username}</li>
                <li>Password: {password}</li>
                <li>Created at: {createdAt}</li>
            </ul>
        </div>
    )
}

export default UserCard

