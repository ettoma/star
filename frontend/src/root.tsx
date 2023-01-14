import "react"
import { Box, Button, Header, Main, Menu } from 'grommet'
import { Home as HomeIcon } from 'grommet-icons'
import { Outlet, useNavigate } from "react-router"

function Root() {
    const navigate = useNavigate()
    return (
        <>
            <Header>
                <Button icon={<HomeIcon />} hoverIndicator />
                <Menu
                    items={[{ label: 'home', onClick: () => { navigate('/') } }]} />
            </Header>
            <Main>
                <Outlet />
            </Main>
        </>
    )
}

export default Root