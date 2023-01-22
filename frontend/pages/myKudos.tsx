import { useEffect, useState } from "react"
import { handleGetKudosPerUser } from "../api/kudos/handleKudos"

type Kudo = {
    sender: string,
    receiver: string,
    content: string,
    id: number
}

function MyKudos() {
    const [kudos, setKudos] = useState([])


    useEffect(() => {
        getKudosPerUser("ettore") //TODO: implement getKudosPerUser for the logged in user
    }, [])

    async function getKudosPerUser(recipient: string) {
        const response = await (handleGetKudosPerUser(recipient))
            .then((response) => response.json())
            .then((response) => setKudos(response))
            .catch((error) => console.log(error))

    }

    return (
        <div>
            {kudos.length === 0 ? <p>no kudos for you (yet)</p> : kudos.map((k: Kudo) => <p key={k.id}>{k.content}</p>)}
        </div>
    )
}

export default MyKudos
