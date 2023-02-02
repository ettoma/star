import { Button, Header, Main } from 'grommet'
import { Home as HomeIcon } from 'grommet-icons'
import { Outlet, useNavigate } from "react-router"

function Root() {
    const navigate = useNavigate()
    return (
        <>
            <Header>
                <Button icon={<HomeIcon />} hoverIndicator onClick={() => { navigate('/') }} />
            </Header>
            <Main>
                <Outlet />
            </Main>
        </>
    )
}

export default Root