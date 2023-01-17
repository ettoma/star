import { User } from "../../api/models/user";

function UserCard({ name, username, createdAt, id }: User): JSX.Element {
    return (
        <div className="user-card-container">
            <h3>{username}</h3>
            <ul className="user-card-container__list">
                <li>ID: <span>{id}</span></li>
                <li>Name: <span>{name}</span></li>
                <li>Created at: <span>{createdAt}</span></li>
            </ul>
        </div>
    )
}

export default UserCard

