import User from "../../api/models/user";
import "./userCard.css"

function UserCard({ name, username, password, createdAt, id }: User): JSX.Element {
    return (
        <div className="user-card-container">
            <h3>{name}</h3>
            <ul className="user-card-container__list">
                <li>ID: <span>{id}</span></li>
                <li>Username: <span>{username}</span></li>
                <li>Password: <span>{password}</span></li>
                <li>Created at: <span>{createdAt}</span></li>
            </ul>
        </div>
    )
}

export default UserCard

